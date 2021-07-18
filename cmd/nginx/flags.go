/*
Copyright 2017 The Kubernetes Authors.

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
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/controller"
	ngx_config "k8s.io/ingress-nginx/internal/ingress/controller/config"
	"k8s.io/ingress-nginx/internal/ingress/controller/ingressclass"
	"k8s.io/ingress-nginx/internal/ingress/status"
	ing_net "k8s.io/ingress-nginx/internal/net"
	"k8s.io/ingress-nginx/internal/nginx"
	klog "k8s.io/klog/v2"
)

func parseFlags() (bool, *controller.Configuration, error) {
	var (
		flags = pflag.NewFlagSet("", pflag.ExitOnError)

		apiserverHost = flags.String("apiserver-host", "",
			`Address of the Kubernetes API server.
Takes the form "protocol://address:port". If not specified, it is assumed the
program runs inside a Kubernetes cluster and local discovery is attempted.`)

		rootCAFile = flags.String("certificate-authority", "",
			`Path to a cert file for the certificate authority. This certificate is used
only when the flag --apiserver-host is specified.`)

		kubeConfigFile = flags.String("kubeconfig", "",
			`Path to a kubeconfig file containing authorization and API server information.`)

		defaultSvc = flags.String("default-backend-service", "",
			`Service used to serve HTTP requests not matching any known server name (catch-all).
Takes the form "namespace/name". The controller configures NGINX to forward
requests to the first port of this Service.`)

		ingressClassAnnotation = flags.String("ingress-class", ingressclass.DefaultAnnotationValue,
			`[IN DEPRECATION] Name of the ingress class this controller satisfies.
The class of an Ingress object is set using the annotation "kubernetes.io/ingress.class" (deprecated).
The parameter --controller-class has precedence over this.`)

		ingressClassController = flags.String("controller-class", ingressclass.DefaultControllerName,
			`Ingress Class Controller value this Ingress satisfies.
The class of an Ingress object is set using the field IngressClassName in Kubernetes clusters version v1.19.0 or higher. The .spec.controller value of the IngressClass 
referenced in an Ingress Object should be the same value specified here to make this object be watched.`)

		watchWithoutClass = flags.Bool("watch-ingress-without-class", false,
			`Define if Ingress Controller should also watch for Ingresses without an IngressClass or the annotation specified`)

		configMap = flags.String("configmap", "",
			`Name of the ConfigMap containing custom global configurations for the controller.`)

		publishSvc = flags.String("publish-service", "",
			`Service fronting the Ingress controller.
Takes the form "namespace/name". When used together with update-status, the
controller mirrors the address of this service's endpoints to the load-balancer
status of all Ingress objects it satisfies.`)

		tcpConfigMapName = flags.String("tcp-services-configmap", "",
			`Name of the ConfigMap containing the definition of the TCP services to expose.
The key in the map indicates the external port to be used. The value is a
reference to a Service in the form "namespace/name:port", where "port" can
either be a port number or name. TCP ports 80 and 443 are reserved by the
controller for servicing HTTP traffic.`)
		udpConfigMapName = flags.String("udp-services-configmap", "",
			`Name of the ConfigMap containing the definition of the UDP services to expose.
The key in the map indicates the external port to be used. The value is a
reference to a Service in the form "namespace/name:port", where "port" can
either be a port name or number.`)

		resyncPeriod = flags.Duration("sync-period", 0,
			`Period at which the controller forces the repopulation of its local object stores. Disabled by default.`)

		watchNamespace = flags.String("watch-namespace", apiv1.NamespaceAll,
			`Namespace the controller watches for updates to Kubernetes objects.
This includes Ingresses, Services and all configuration resources. All
namespaces are watched if this parameter is left empty.`)

		profiling = flags.Bool("profiling", true,
			`Enable profiling via web interface host:port/debug/pprof/`)

		defSSLCertificate = flags.String("default-ssl-certificate", "",
			`Secret containing a SSL certificate to be used by the default HTTPS server (catch-all).
Takes the form "namespace/name".`)

		defHealthzURL = flags.String("health-check-path", "/healthz",
			`URL path of the health check endpoint.
Configured inside the NGINX status server. All requests received on the port
defined by the healthz-port parameter are forwarded internally to this path.`)

		defHealthCheckTimeout = flags.Int("health-check-timeout", 10, `Time limit, in seconds, for a probe to health-check-path to succeed.`)

		updateStatus = flags.Bool("update-status", true,
			`Update the load-balancer status of Ingress objects this controller satisfies.
Requires setting the publish-service parameter to a valid Service reference.`)

		electionID = flags.String("election-id", "ingress-controller-leader",
			`Election id to use for Ingress status updates.`)

		updateStatusOnShutdown = flags.Bool("update-status-on-shutdown", true,
			`Update the load-balancer status of Ingress objects when the controller shuts down.
Requires the update-status parameter.`)

		useNodeInternalIP = flags.Bool("report-node-internal-ip-address", false,
			`Set the load-balancer status of Ingress objects to internal Node addresses instead of external.
Requires the update-status parameter.`)

		showVersion = flags.Bool("version", false,
			`Show release information about the NGINX Ingress controller and exit.`)

		enableSSLPassthrough = flags.Bool("enable-ssl-passthrough", false,
			`Enable SSL Passthrough.`)

		disableServiceExternalName = flags.Bool("disable-svc-external-name", false,
			`Disable support for Services of type ExternalName`)

		annotationsPrefix = flags.String("annotations-prefix", parser.DefaultAnnotationsPrefix,
			`Prefix of the Ingress annotations specific to the NGINX controller.`)

		enableSSLChainCompletion = flags.Bool("enable-ssl-chain-completion", false,
			`Autocomplete SSL certificate chains with missing intermediate CA certificates.
Certificates uploaded to Kubernetes must have the "Authority Information Access" X.509 v3
extension for this to succeed.`)

		syncRateLimit = flags.Float32("sync-rate-limit", 0.3,
			`Define the sync frequency upper limit`)

		publishStatusAddress = flags.String("publish-status-address", "",
			`Customized address (or addresses, separated by comma) to set as the load-balancer status of Ingress objects this controller satisfies.
Requires the update-status parameter.`)

		enableMetrics = flags.Bool("enable-metrics", true,
			`Enables the collection of NGINX metrics`)
		metricsPerHost = flags.Bool("metrics-per-host", true,
			`Export metrics per-host`)
		monitorMaxBatchSize = flags.Int("monitor-max-batch-size", 10000, "Max batch size of NGINX metrics")

		httpPort  = flags.Int("http-port", 80, `Port to use for servicing HTTP traffic.`)
		httpsPort = flags.Int("https-port", 443, `Port to use for servicing HTTPS traffic.`)

		sslProxyPort  = flags.Int("ssl-passthrough-proxy-port", 442, `Port to use internally for SSL Passthrough.`)
		defServerPort = flags.Int("default-server-port", 8181, `Port to use for exposing the default server (catch-all).`)
		healthzPort   = flags.Int("healthz-port", 10254, "Port to use for the healthz endpoint.")

		disableCatchAll = flags.Bool("disable-catch-all", false,
			`Disable support for catch-all Ingresses`)

		validationWebhook = flags.String("validating-webhook", "",
			`The address to start an admission controller on to validate incoming ingresses.
Takes the form "<host>:port". If not provided, no admission controller is started.`)
		validationWebhookCert = flags.String("validating-webhook-certificate", "",
			`The path of the validating webhook certificate PEM.`)
		validationWebhookKey = flags.String("validating-webhook-key", "",
			`The path of the validating webhook key PEM.`)

		statusPort = flags.Int("status-port", 10246, `Port to use for the lua HTTP endpoint configuration.`)
		streamPort = flags.Int("stream-port", 10247, "Port to use for the lua TCP/UDP endpoint configuration.")

		profilerPort = flags.Int("profiler-port", 10245, "Port to use for expose the ingress controller Go profiler when it is enabled.")

		statusUpdateInterval = flags.Int("status-update-interval", status.UpdateInterval, "Time interval in seconds in which the status should check if an update is required. Default is 60 seconds")

		shutdownGracePeriod = flags.Int("shutdown-grace-period", 0, "Seconds to wait after receiving the shutdown signal, before stopping the nginx process.")
	)

	flags.StringVar(&nginx.MaxmindMirror, "maxmind-mirror", "", `Maxmind mirror url (example: http://geoip.local/databases`)
	flags.StringVar(&nginx.MaxmindLicenseKey, "maxmind-license-key", "", `Maxmind license key to download GeoLite2 Databases.
https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases`)
	flags.StringVar(&nginx.MaxmindEditionIDs, "maxmind-edition-ids", "GeoLite2-City,GeoLite2-ASN", `Maxmind edition ids to download GeoLite2 Databases.`)
	flags.IntVar(&nginx.MaxmindRetriesCount, "maxmind-retries-count", 1, "Number of attempts to download the GeoIP DB.")
	flags.DurationVar(&nginx.MaxmindRetriesTimeout, "maxmind-retries-timeout", time.Second*0, "Maxmind downloading delay between 1st and 2nd attempt, 0s - do not retry to download if something went wrong.")

	flag.Set("logtostderr", "true")

	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)

	// Workaround for this issue:
	// https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})

	pflag.VisitAll(func(flag *pflag.Flag) {
		klog.V(2).InfoS("FLAG", flag.Name, flag.Value)
	})

	if *showVersion {
		return true, nil, nil
	}

	if *statusUpdateInterval < 5 {
		klog.Warningf("The defined time to update the Ingress status too low (%v seconds). Adjusting to 5 seconds", *statusUpdateInterval)
		status.UpdateInterval = 5
	} else {
		status.UpdateInterval = *statusUpdateInterval
	}

	parser.AnnotationsPrefix = *annotationsPrefix

	// check port collisions
	if !ing_net.IsPortAvailable(*httpPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --http-port", *httpPort)
	}

	if !ing_net.IsPortAvailable(*httpsPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --https-port", *httpsPort)
	}

	if !ing_net.IsPortAvailable(*defServerPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --default-server-port", *defServerPort)
	}

	if !ing_net.IsPortAvailable(*statusPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --status-port", *statusPort)
	}

	if !ing_net.IsPortAvailable(*streamPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --stream-port", *streamPort)
	}

	if !ing_net.IsPortAvailable(*profilerPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --profiler-port", *profilerPort)
	}

	nginx.StatusPort = *statusPort
	nginx.StreamPort = *streamPort
	nginx.ProfilerPort = *profilerPort

	if *enableSSLPassthrough && !ing_net.IsPortAvailable(*sslProxyPort) {
		return false, nil, fmt.Errorf("port %v is already in use. Please check the flag --ssl-passthrough-proxy-port", *sslProxyPort)
	}

	if *publishSvc != "" && *publishStatusAddress != "" {
		return false, nil, fmt.Errorf("flags --publish-service and --publish-status-address are mutually exclusive")
	}

	nginx.HealthPath = *defHealthzURL

	if *defHealthCheckTimeout > 0 {
		nginx.HealthCheckTimeout = time.Duration(*defHealthCheckTimeout) * time.Second
	}

	ngx_config.EnableSSLChainCompletion = *enableSSLChainCompletion

	config := &controller.Configuration{
		APIServerHost:              *apiserverHost,
		KubeConfigFile:             *kubeConfigFile,
		UpdateStatus:               *updateStatus,
		ElectionID:                 *electionID,
		EnableProfiling:            *profiling,
		EnableMetrics:              *enableMetrics,
		MetricsPerHost:             *metricsPerHost,
		MonitorMaxBatchSize:        *monitorMaxBatchSize,
		DisableServiceExternalName: *disableServiceExternalName,
		EnableSSLPassthrough:       *enableSSLPassthrough,
		ResyncPeriod:               *resyncPeriod,
		DefaultService:             *defaultSvc,
		Namespace:                  *watchNamespace,
		ConfigMapName:              *configMap,
		TCPConfigMapName:           *tcpConfigMapName,
		UDPConfigMapName:           *udpConfigMapName,
		DefaultSSLCertificate:      *defSSLCertificate,
		PublishService:             *publishSvc,
		PublishStatusAddress:       *publishStatusAddress,
		UpdateStatusOnShutdown:     *updateStatusOnShutdown,
		ShutdownGracePeriod:        *shutdownGracePeriod,
		UseNodeInternalIP:          *useNodeInternalIP,
		SyncRateLimit:              *syncRateLimit,
		ListenPorts: &ngx_config.ListenPorts{
			Default:  *defServerPort,
			Health:   *healthzPort,
			HTTP:     *httpPort,
			HTTPS:    *httpsPort,
			SSLProxy: *sslProxyPort,
		},
		IngressClassConfiguration: &ingressclass.IngressClassConfiguration{
			Controller:        *ingressClassController,
			AnnotationValue:   *ingressClassAnnotation,
			WatchWithoutClass: *watchWithoutClass,
		},
		DisableCatchAll:           *disableCatchAll,
		ValidationWebhook:         *validationWebhook,
		ValidationWebhookCertPath: *validationWebhookCert,
		ValidationWebhookKeyPath:  *validationWebhookKey,
	}

	if *apiserverHost != "" {
		config.RootCAFile = *rootCAFile
	}

	var err error
	if nginx.MaxmindEditionIDs != "" {
		if err = nginx.ValidateGeoLite2DBEditions(); err != nil {
			return false, nil, err
		}
		if nginx.MaxmindLicenseKey != "" || nginx.MaxmindMirror != "" {
			klog.InfoS("downloading maxmind GeoIP2 databases")
			if err = nginx.DownloadGeoLite2DB(nginx.MaxmindRetriesCount, nginx.MaxmindRetriesTimeout); err != nil {
				klog.ErrorS(err, "unexpected error downloading GeoIP2 database")
			}
		}
		config.MaxmindEditionFiles = &nginx.MaxmindEditionFiles
	}

	return false, config, err
}
