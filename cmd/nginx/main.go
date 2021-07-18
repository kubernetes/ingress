/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"math/rand" // #nosec
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	discovery "k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/klog/v2"

	"k8s.io/ingress-nginx/internal/file"
	"k8s.io/ingress-nginx/internal/ingress/controller"
	"k8s.io/ingress-nginx/internal/ingress/metric"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/ingress-nginx/internal/net/ssl"
	"k8s.io/ingress-nginx/internal/nginx"
	"k8s.io/ingress-nginx/version"
)

func main() {
	klog.InitFlags(nil)

	rand.Seed(time.Now().UnixNano())

	fmt.Println(version.String())

	showVersion, conf, err := parseFlags()
	if showVersion {
		os.Exit(0)
	}

	if err != nil {
		klog.Fatal(err)
	}

	err = file.CreateRequiredDirectories()
	if err != nil {
		klog.Fatal(err)
	}

	kubeClient, err := createApiserverClient(conf.APIServerHost, conf.RootCAFile, conf.KubeConfigFile)
	if err != nil {
		handleFatalInitError(err)
	}

	if len(conf.DefaultService) > 0 {
		err := checkService(conf.DefaultService, kubeClient)
		if err != nil {
			klog.Fatal(err)
		}

		klog.InfoS("Valid default backend", "service", conf.DefaultService)
	}

	if len(conf.PublishService) > 0 {
		err := checkService(conf.PublishService, kubeClient)
		if err != nil {
			klog.Fatal(err)
		}
	}

	if conf.Namespace != "" {
		_, err = kubeClient.CoreV1().Namespaces().Get(context.TODO(), conf.Namespace, metav1.GetOptions{})
		if err != nil {
			klog.Fatalf("No namespace with name %v found: %v", conf.Namespace, err)
		}
	}

	conf.FakeCertificate = ssl.GetFakeSSLCert()
	klog.InfoS("SSL fake certificate created", "file", conf.FakeCertificate.PemFileName)

	if !k8s.NetworkingIngressAvailable(kubeClient) {
		klog.Fatalf("ingress-nginx requires Kubernetes v1.19.0 or higher")
	}

	_, err = kubeClient.NetworkingV1().IngressClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			if errors.IsUnauthorized(err) || !errors.IsForbidden(err) {
				klog.Fatalf("Error searching IngressClass: Please verify your RBAC and allow Ingress Controller to list and get Ingress Classes: %v", err)
			}
		}
	}

	conf.Client = kubeClient

	err = k8s.GetIngressPod(kubeClient)
	if err != nil {
		klog.Fatalf("Unexpected error obtaining ingress-nginx pod: %v", err)
	}

	reg := prometheus.NewRegistry()

	reg.MustRegister(prometheus.NewGoCollector())
	reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{
		PidFn:        func() (int, error) { return os.Getpid(), nil },
		ReportErrors: true,
	}))

	mc := metric.NewDummyCollector()
	if conf.EnableMetrics {
		mc, err = metric.NewCollector(conf.MetricsPerHost, reg, conf.IngressClassConfiguration.Controller)
		if err != nil {
			klog.Fatalf("Error creating prometheus collector:  %v", err)
		}
	}
	mc.Start()

	if conf.EnableProfiling {
		go registerProfiler()
	}

	ngx := controller.NewNGINXController(conf, mc)

	mux := http.NewServeMux()
	registerHealthz(nginx.HealthPath, ngx, mux)
	registerMetrics(reg, mux)

	go startHTTPServer(conf.ListenPorts.Health, mux)
	go ngx.Start()

	handleSigterm(ngx, func(code int) {
		os.Exit(code)
	})
}

type exiter func(code int)

func handleSigterm(ngx *controller.NGINXController, exit exiter) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	<-signalChan
	klog.InfoS("Received SIGTERM, shutting down")

	exitCode := 0
	if err := ngx.Stop(); err != nil {
		klog.Warningf("Error during shutdown: %v", err)
		exitCode = 1
	}

	klog.InfoS("Handled quit, awaiting Pod deletion")
	time.Sleep(10 * time.Second)

	klog.InfoS("Exiting", "code", exitCode)
	exit(exitCode)
}

// createApiserverClient creates a new Kubernetes REST client. apiserverHost is
// the URL of the API server in the format protocol://address:port/pathPrefix,
// kubeConfig is the location of a kubeconfig file. If defined, the kubeconfig
// file is loaded first, the URL of the API server read from the file is then
// optionally overridden by the value of apiserverHost.
// If neither apiserverHost nor kubeConfig is passed in, we assume the
// controller runs inside Kubernetes and fallback to the in-cluster config. If
// the in-cluster config is missing or fails, we fallback to the default config.
func createApiserverClient(apiserverHost, rootCAFile, kubeConfig string) (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(apiserverHost, kubeConfig)
	if err != nil {
		return nil, err
	}

	// TODO: remove after k8s v1.22
	cfg.WarningHandler = rest.NoWarnings{}

	// Configure the User-Agent used for the HTTP requests made to the API server.
	cfg.UserAgent = fmt.Sprintf(
		"%s/%s (%s/%s) ingress-nginx/%s",
		filepath.Base(os.Args[0]),
		version.RELEASE,
		runtime.GOOS,
		runtime.GOARCH,
		version.COMMIT,
	)

	if apiserverHost != "" && rootCAFile != "" {
		tlsClientConfig := rest.TLSClientConfig{}

		if _, err := certutil.NewPool(rootCAFile); err != nil {
			klog.ErrorS(err, "Loading CA config", "file", rootCAFile)
		} else {
			tlsClientConfig.CAFile = rootCAFile
		}

		cfg.TLSClientConfig = tlsClientConfig
	}

	klog.InfoS("Creating API client", "host", cfg.Host)

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	var v *discovery.Info

	// The client may fail to connect to the API server in the first request.
	// https://github.com/kubernetes/ingress-nginx/issues/1968
	defaultRetry := wait.Backoff{
		Steps:    10,
		Duration: 1 * time.Second,
		Factor:   1.5,
		Jitter:   0.1,
	}

	var lastErr error
	retries := 0
	klog.V(2).InfoS("Trying to discover Kubernetes version")
	err = wait.ExponentialBackoff(defaultRetry, func() (bool, error) {
		v, err = client.Discovery().ServerVersion()

		if err == nil {
			return true, nil
		}

		lastErr = err
		klog.V(2).ErrorS(err, "Unexpected error discovering Kubernetes version", "attempt", retries)
		retries++
		return false, nil
	})

	// err is returned in case of timeout in the exponential backoff (ErrWaitTimeout)
	if err != nil {
		return nil, lastErr
	}

	// this should not happen, warn the user
	if retries > 0 {
		klog.Warningf("Initial connection to the Kubernetes API server was retried %d times.", retries)
	}

	klog.InfoS("Running in Kubernetes cluster",
		"major", v.Major,
		"minor", v.Minor,
		"git", v.GitVersion,
		"state", v.GitTreeState,
		"commit", v.GitCommit,
		"platform", v.Platform,
	)

	return client, nil
}

// Handler for fatal init errors. Prints a verbose error message and exits.
func handleFatalInitError(err error) {
	klog.Fatalf("Error while initiating a connection to the Kubernetes API server. "+
		"This could mean the cluster is misconfigured (e.g. it has invalid API server certificates "+
		"or Service Accounts configuration). Reason: %s\n"+
		"Refer to the troubleshooting guide for more information: "+
		"https://kubernetes.github.io/ingress-nginx/troubleshooting/",
		err)
}

func registerHealthz(healthPath string, ic *controller.NGINXController, mux *http.ServeMux) {
	// expose health check endpoint (/healthz)
	healthz.InstallPathHandler(mux,
		healthPath,
		healthz.PingHealthz,
		ic,
	)
}

func registerMetrics(reg *prometheus.Registry, mux *http.ServeMux) {
	mux.Handle(
		"/metrics",
		promhttp.InstrumentMetricHandler(
			reg,
			promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		),
	)
}

func registerProfiler() {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/heap", pprof.Index)
	mux.HandleFunc("/debug/pprof/mutex", pprof.Index)
	mux.HandleFunc("/debug/pprof/goroutine", pprof.Index)
	mux.HandleFunc("/debug/pprof/threadcreate", pprof.Index)
	mux.HandleFunc("/debug/pprof/block", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%v", nginx.ProfilerPort),
		Handler: mux,
	}
	klog.Fatal(server.ListenAndServe())
}

func startHTTPServer(port int, mux *http.ServeMux) {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", port),
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      300 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	klog.Fatal(server.ListenAndServe())
}

func checkService(key string, kubeClient *kubernetes.Clientset) error {
	ns, name, err := k8s.ParseNameNS(key)
	if err != nil {
		return err
	}

	_, err = kubeClient.CoreV1().Services(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsUnauthorized(err) || errors.IsForbidden(err) {
			return fmt.Errorf("✖ the cluster seems to be running with a restrictive Authorization mode and the Ingress controller does not have the required permissions to operate normally")
		}

		if errors.IsNotFound(err) {
			return fmt.Errorf("No service with name %v found in namespace %v: %v", name, ns, err)
		}

		return fmt.Errorf("Unexpected error searching service with name %v in namespace %v: %v", name, ns, err)
	}

	return nil
}
