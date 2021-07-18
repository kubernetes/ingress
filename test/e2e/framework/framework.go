/*
Copyright 2017 Jetstack Ltd.
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

package framework

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/onsi/ginkgo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/klog/v2"
)

// RequestScheme define a scheme used in a test request.
type RequestScheme string

// These are valid test request schemes.
const (
	HTTP  RequestScheme = "http"
	HTTPS RequestScheme = "https"
)

var (
	// KubectlPath defines the full path of the kubectl binary
	KubectlPath = "/usr/local/bin/kubectl"
)

// Framework supports common operations used by e2e tests; it will keep a client & a namespace for you.
type Framework struct {
	BaseName string

	IsIngressV1Ready bool

	// A Kubernetes and Service Catalog client
	KubeClientSet          kubernetes.Interface
	KubeConfig             *restclient.Config
	APIExtensionsClientSet apiextcs.Interface

	Namespace    string
	IngressClass string

	pod *corev1.Pod
}

// NewDefaultFramework makes a new framework and sets up a BeforeEach/AfterEach for
// you (you can write additional before/after each functions).
func NewDefaultFramework(baseName string) *Framework {
	defer ginkgo.GinkgoRecover()

	f := &Framework{
		BaseName: baseName,
	}

	ginkgo.BeforeEach(f.BeforeEach)
	ginkgo.AfterEach(f.AfterEach)

	return f
}

// BeforeEach gets a client and makes a namespace.
func (f *Framework) BeforeEach() {
	var err error

	if f.KubeClientSet == nil {
		f.KubeConfig, err = loadConfig()
		assert.Nil(ginkgo.GinkgoT(), err, "loading a kubernetes client configuration")

		// TODO: remove after k8s v1.22
		f.KubeConfig.WarningHandler = rest.NoWarnings{}

		f.KubeClientSet, err = kubernetes.NewForConfig(f.KubeConfig)
		assert.Nil(ginkgo.GinkgoT(), err, "creating a kubernetes client")

		f.IsIngressV1Ready = k8s.NetworkingIngressAvailable(f.KubeClientSet)
	}

	f.Namespace, err = CreateKubeNamespace(f.BaseName, f.KubeClientSet)
	assert.Nil(ginkgo.GinkgoT(), err, "creating namespace")

	f.IngressClass, err = CreateIngressClass(f.Namespace, f.KubeClientSet)
	assert.Nil(ginkgo.GinkgoT(), err, "creating IngressClass")

	err = f.newIngressController(f.Namespace, f.BaseName)
	assert.Nil(ginkgo.GinkgoT(), err, "deploying the ingress controller")

	err = f.updateIngressNGINXPod()
	assert.Nil(ginkgo.GinkgoT(), err, "updating ingress controller pod information")

	f.WaitForNginxListening(80)
}

// AfterEach deletes the namespace, after reading its events.
func (f *Framework) AfterEach() {
	defer func(kubeClient kubernetes.Interface, ns string) {
		go func() {
			defer ginkgo.GinkgoRecover()
			err := deleteKubeNamespace(kubeClient, ns)
			assert.Nil(ginkgo.GinkgoT(), err, "deleting namespace %v", f.Namespace)
		}()
	}(f.KubeClientSet, f.Namespace)

	defer func(kubeClient kubernetes.Interface, ingressclass string) {
		go func() {
			defer ginkgo.GinkgoRecover()
			err := deleteIngressClass(kubeClient, ingressclass)
			assert.Nil(ginkgo.GinkgoT(), err, "deleting IngressClass")
		}()
	}(f.KubeClientSet, f.IngressClass)

	if !ginkgo.CurrentGinkgoTestDescription().Failed {
		return
	}

	cmd := fmt.Sprintf("cat /etc/nginx/nginx.conf")
	o, err := f.ExecCommand(f.pod, cmd)
	if err != nil {
		Logf("Unexpected error obtaining nginx.conf file: %v", err)
		return
	}

	ginkgo.By("Dumping NGINX configuration after failure")
	Logf("%v", o)

	log, err := f.NginxLogs()
	if err != nil {
		Logf("Unexpected error obtaining NGINX logs: %v", err)
		return
	}

	ginkgo.By("Dumping NGINX logs")
	Logf("%v", log)

	o, err = f.NamespaceContent()
	if err != nil {
		Logf("Unexpected error obtaining namespace information: %v", err)
		return
	}

	ginkgo.By("Dumping namespace content")
	Logf("%v", o)
}

// IngressNginxDescribe wrapper function for ginkgo describe. Adds namespacing.
func IngressNginxDescribe(text string, body func()) bool {
	return ginkgo.Describe(text, body)
}

// DescribeAnnotation wrapper function for ginkgo describe. Adds namespacing.
func DescribeAnnotation(text string, body func()) bool {
	return ginkgo.Describe("[Annotations] "+text, body)
}

// DescribeSetting wrapper function for ginkgo describe. Adds namespacing.
func DescribeSetting(text string, body func()) bool {
	return ginkgo.Describe("[Setting] "+text, body)
}

// MemoryLeakIt is wrapper function for ginkgo It.  Adds "[MemoryLeak]" tag and makes static analysis easier.
func MemoryLeakIt(text string, body interface{}, timeout ...float64) bool {
	return ginkgo.It(text+" [MemoryLeak]", body, timeout...)
}

// GetNginxIP returns the number of TCP port where NGINX is running
func (f *Framework) GetNginxIP() string {
	s, err := f.KubeClientSet.
		CoreV1().
		Services(f.Namespace).
		Get(context.TODO(), "nginx-ingress-controller", metav1.GetOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "obtaining NGINX IP address")
	return s.Spec.ClusterIP
}

// GetNginxPodIP returns the IP addresses of the running pods
func (f *Framework) GetNginxPodIP() string {
	return f.pod.Status.PodIP
}

// GetURL returns the URL should be used to make a request to NGINX
func (f *Framework) GetURL(scheme RequestScheme) string {
	ip := f.GetNginxIP()
	return fmt.Sprintf("%v://%v", scheme, ip)
}

// GetIngressNGINXPod returns the ingress controller running pod
func (f *Framework) GetIngressNGINXPod() *corev1.Pod {
	return f.pod
}

// UpdateIngressNGINXPod search and updates the ingress controller running pod
func (f *Framework) updateIngressNGINXPod() error {
	var err error
	f.pod, err = getIngressNGINXPod(f.Namespace, f.KubeClientSet)
	return err
}

// WaitForNginxServer waits until the nginx configuration contains a particular server section.
// `cfg` passed to matcher is normalized by replacing all tabs and spaces with single space.
func (f *Framework) WaitForNginxServer(name string, matcher func(cfg string) bool) {
	err := wait.Poll(Poll, DefaultTimeout, f.matchNginxConditions(name, matcher))
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for nginx server condition/s")
	Sleep(1 * time.Second)
}

// WaitForNginxConfiguration waits until the nginx configuration contains a particular configuration
// `cfg` passed to matcher is normalized by replacing all tabs and spaces with single space.
func (f *Framework) WaitForNginxConfiguration(matcher func(cfg string) bool) {
	err := wait.Poll(Poll, DefaultTimeout, f.matchNginxConditions("", matcher))
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for nginx server condition/s")
	Sleep(1 * time.Second)
}

// WaitForNginxCustomConfiguration waits until the nginx configuration given part (from, to) contains a particular configuration
func (f *Framework) WaitForNginxCustomConfiguration(from string, to string, matcher func(cfg string) bool) {
	err := wait.Poll(Poll, DefaultTimeout, f.matchNginxCustomConditions(from, to, matcher))
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for nginx server condition/s")
}

// NginxLogs returns the logs of the nginx ingress controller pod running
func (f *Framework) NginxLogs() (string, error) {
	if isRunning, err := podRunningReady(f.pod); err == nil && isRunning {
		return Logs(f.KubeClientSet, f.Namespace, f.pod.Name)
	}

	return "", fmt.Errorf("no nginx ingress controller pod is running (logs)")
}

func (f *Framework) matchNginxConditions(name string, matcher func(cfg string) bool) wait.ConditionFunc {
	return func() (bool, error) {
		var cmd string
		if name == "" {
			cmd = fmt.Sprintf("cat /etc/nginx/nginx.conf")
		} else {
			cmd = fmt.Sprintf("cat /etc/nginx/nginx.conf | awk '/## start server %v/,/## end server %v/'", name, name)
		}

		o, err := f.ExecCommand(f.pod, cmd)
		if err != nil {
			return false, nil
		}

		if klog.V(10).Enabled() && len(o) > 0 {
			klog.InfoS("NGINX", "configuration", o)
		}

		// passes the nginx config to the passed function
		if matcher(strings.Join(strings.Fields(o), " ")) {
			return true, nil
		}

		return false, nil
	}
}

func (f *Framework) matchNginxCustomConditions(from string, to string, matcher func(cfg string) bool) wait.ConditionFunc {
	return func() (bool, error) {
		cmd := fmt.Sprintf("cat /etc/nginx/nginx.conf| awk '/%v/,/%v/'", from, to)

		o, err := f.ExecCommand(f.pod, cmd)
		if err != nil {
			return false, nil
		}

		if klog.V(10).Enabled() && len(o) > 0 {
			klog.InfoS("NGINX", "configuration", o)
		}

		// passes the nginx config to the passed function
		if matcher(strings.Join(strings.Fields(o), " ")) {
			return true, nil
		}

		return false, nil
	}
}

func (f *Framework) getConfigMap(name string) (*v1.ConfigMap, error) {
	if f.KubeClientSet == nil {
		return nil, fmt.Errorf("KubeClientSet not initialized")
	}

	config, err := f.KubeClientSet.
		CoreV1().
		ConfigMaps(f.Namespace).
		Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return config, err
}

// SetNginxConfigMapData sets ingress-nginx's nginx-ingress-controller configMap data
func (f *Framework) SetNginxConfigMapData(cmData map[string]string) {
	cfgMap, err := f.getConfigMap("nginx-ingress-controller")
	assert.Nil(ginkgo.GinkgoT(), err)
	assert.NotNil(ginkgo.GinkgoT(), cfgMap, "expected a configmap but none returned")

	cfgMap.Data = cmData

	fn := func() {
		_, err = f.KubeClientSet.
			CoreV1().
			ConfigMaps(f.Namespace).
			Update(context.TODO(), cfgMap, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "updating configuration configmap")
	}

	f.WaitForReload(fn)
}

// CreateConfigMap creates a new configmap in the current namespace
func (f *Framework) CreateConfigMap(name string, data map[string]string) {
	_, err := f.KubeClientSet.CoreV1().ConfigMaps(f.Namespace).Create(context.TODO(), &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Data: data,
	}, metav1.CreateOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "creating configMap")
}

// UpdateNginxConfigMapData updates single field in ingress-nginx's nginx-ingress-controller map data
func (f *Framework) UpdateNginxConfigMapData(key string, value string) {
	config, err := f.getConfigMap("nginx-ingress-controller")
	assert.Nil(ginkgo.GinkgoT(), err)
	assert.NotNil(ginkgo.GinkgoT(), config, "expected a configmap but none returned")

	config.Data[key] = value

	fn := func() {
		_, err = f.KubeClientSet.
			CoreV1().
			ConfigMaps(f.Namespace).
			Update(context.TODO(), config, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "updating configuration configmap")
	}

	f.WaitForReload(fn)
}

// WaitForReload calls the passed function and
// asser it has caused at least 1 reload.
func (f *Framework) WaitForReload(fn func()) {
	initialReloadCount := getReloadCount(f.pod, f.Namespace, f.KubeClientSet)

	fn()

	count := 0
	err := wait.Poll(1*time.Second, DefaultTimeout, func() (bool, error) {
		reloads := getReloadCount(f.pod, f.Namespace, f.KubeClientSet)
		// most of the cases reload the ingress controller
		// in cases where the value is not modified we could wait forever
		if count > 5 && reloads == initialReloadCount {
			return true, nil
		}

		count++

		return (reloads > initialReloadCount), nil
	})
	assert.Nil(ginkgo.GinkgoT(), err, "while waiting for ingress controller reload")
}

func getReloadCount(pod *corev1.Pod, namespace string, client kubernetes.Interface) int {
	events, err := client.CoreV1().Events(namespace).Search(scheme.Scheme, pod)
	assert.Nil(ginkgo.GinkgoT(), err, "obtaining NGINX Pod")

	reloadCount := 0
	for _, e := range events.Items {
		if e.Reason == "RELOAD" && e.Type == corev1.EventTypeNormal {
			reloadCount++
		}
	}

	return reloadCount
}

// DeleteNGINXPod deletes the currently running pod. It waits for the replacement pod to be up.
// Grace period to wait for pod shutdown is in seconds.
func (f *Framework) DeleteNGINXPod(grace int64) {
	ns := f.Namespace

	err := f.KubeClientSet.CoreV1().Pods(ns).Delete(context.TODO(), f.pod.GetName(), *metav1.NewDeleteOptions(grace))
	assert.Nil(ginkgo.GinkgoT(), err, "deleting ingress nginx pod")

	err = wait.Poll(Poll, DefaultTimeout, func() (bool, error) {
		err := f.updateIngressNGINXPod()
		if err != nil || f.pod == nil {
			return false, nil
		}
		return f.pod.GetName() != "", nil
	})
	assert.Nil(ginkgo.GinkgoT(), err, "while waiting for ingress nginx pod to come up again")
}

// HTTPTestClient returns a new httpexpect client for end-to-end HTTP testing.
func (f *Framework) HTTPTestClient() *httpexpect.Expect {
	return f.newTestClient(nil)
}

// HTTPTestClientWithTLSConfig returns a new httpexpect client for end-to-end
// HTTP testing with a custom TLS configuration.
func (f *Framework) HTTPTestClientWithTLSConfig(config *tls.Config) *httpexpect.Expect {
	return f.newTestClient(config)
}

func (f *Framework) newTestClient(config *tls.Config) *httpexpect.Expect {
	if config == nil {
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: f.GetURL(HTTP),
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: config,
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		Reporter: httpexpect.NewAssertReporter(
			httpexpect.NewAssertReporter(ginkgo.GinkgoT()),
		),
		Printers: []httpexpect.Printer{
			// TODO: enable conditionally?
			// httpexpect.NewDebugPrinter(ginkgo.GinkgoT(), false),
		},
	})
}

// WaitForNginxListening waits until NGINX starts accepting connections on a port
func (f *Framework) WaitForNginxListening(port int) {
	err := waitForPodsReady(f.KubeClientSet, DefaultTimeout, 1, f.Namespace, metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=ingress-nginx",
	})
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for ingress pods to be ready")

	podIP := f.GetNginxIP()
	err = wait.Poll(500*time.Millisecond, DefaultTimeout, func() (bool, error) {
		hostPort := net.JoinHostPort(podIP, fmt.Sprintf("%v", port))
		conn, err := net.Dial("tcp", hostPort)
		if err != nil {
			return false, nil
		}

		defer conn.Close()

		return true, nil
	})
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for ingress controller pod listening on port 80")
}

// UpdateDeployment runs the given updateFunc on the deployment and waits for it to be updated
func UpdateDeployment(kubeClientSet kubernetes.Interface, namespace string, name string, replicas int, updateFunc func(d *appsv1.Deployment) error) error {
	deployment, err := kubeClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if updateFunc != nil {
		if err := updateFunc(deployment); err != nil {
			return err
		}

		err = waitForDeploymentRollout(kubeClientSet, deployment)
		if err != nil {
			return err
		}
	}

	if *deployment.Spec.Replicas != int32(replicas) {
		deployment.Spec.Replicas = NewInt32(int32(replicas))
		_, err = kubeClientSet.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		if err != nil {
			return errors.Wrapf(err, "scaling the number of replicas to %v", replicas)
		}

		err = waitForDeploymentRollout(kubeClientSet, deployment)
		if err != nil {
			return err
		}
	}

	err = waitForPodsReady(kubeClientSet, DefaultTimeout, replicas, namespace, metav1.ListOptions{
		LabelSelector: fields.SelectorFromSet(fields.Set(deployment.Spec.Template.ObjectMeta.Labels)).String(),
	})
	if err != nil {
		return errors.Wrapf(err, "waiting for nginx-ingress-controller replica count to be %v", replicas)
	}

	return nil
}

func waitForDeploymentRollout(kubeClientSet kubernetes.Interface, resource *appsv1.Deployment) error {
	return wait.Poll(Poll, 5*time.Minute, func() (bool, error) {
		d, err := kubeClientSet.AppsV1().Deployments(resource.Namespace).Get(context.TODO(), resource.Name, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return false, nil
		}

		if err != nil {
			return false, nil
		}

		if d.DeletionTimestamp != nil {
			return false, fmt.Errorf("deployment %q is being deleted", resource.Name)
		}

		if d.Generation <= d.Status.ObservedGeneration && d.Status.UpdatedReplicas == d.Status.Replicas && d.Status.UnavailableReplicas == 0 {
			return true, nil
		}

		return false, nil
	})
}

// UpdateIngress runs the given updateFunc on the ingress
func UpdateIngress(kubeClientSet kubernetes.Interface, namespace string, name string, updateFunc func(d *networking.Ingress) error) error {
	ingress, err := kubeClientSet.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if ingress == nil {
		return fmt.Errorf("there is no ingress with name %v in namespace %v", name, namespace)
	}

	if ingress.ObjectMeta.Annotations == nil {
		ingress.ObjectMeta.Annotations = map[string]string{}
	}

	if err := updateFunc(ingress); err != nil {
		return err
	}

	_, err = kubeClientSet.NetworkingV1().Ingresses(namespace).Update(context.TODO(), ingress, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	Sleep(1 * time.Second)
	return nil
}

// NewSingleIngressWithTLS creates a simple ingress rule with TLS spec included
func NewSingleIngressWithTLS(name, path, host string, tlsHosts []string, ns, service string, port int, annotations map[string]string) *networking.Ingress {
	return newSingleIngressWithRules(name, path, host, ns, service, port, annotations, tlsHosts)
}

// NewSingleIngress creates a simple ingress rule
func NewSingleIngress(name, path, host, ns, service string, port int, annotations map[string]string) *networking.Ingress {
	return newSingleIngressWithRules(name, path, host, ns, service, port, annotations, nil)
}

// NewSingleIngressWithMultiplePaths creates a simple ingress rule with multiple paths
func NewSingleIngressWithMultiplePaths(name string, paths []string, host, ns, service string, port int, annotations map[string]string) *networking.Ingress {
	pathtype := networking.PathTypePrefix
	spec := networking.IngressSpec{
		IngressClassName: GetIngressClassName(ns),
		Rules: []networking.IngressRule{
			{
				Host: host,
				IngressRuleValue: networking.IngressRuleValue{
					HTTP: &networking.HTTPIngressRuleValue{},
				},
			},
		},
	}

	for _, path := range paths {
		spec.Rules[0].IngressRuleValue.HTTP.Paths = append(spec.Rules[0].IngressRuleValue.HTTP.Paths, networking.HTTPIngressPath{
			Path:     path,
			PathType: &pathtype,
			Backend: networking.IngressBackend{
				Service: &networking.IngressServiceBackend{
					Name: service,
					Port: networking.ServiceBackendPort{
						Number: int32(port),
					},
				},
			},
		})
	}

	return newSingleIngress(name, ns, annotations, spec)
}

func newSingleIngressWithRules(name, path, host, ns, service string, port int, annotations map[string]string, tlsHosts []string) *networking.Ingress {
	pathtype := networking.PathTypePrefix
	spec := networking.IngressSpec{
		IngressClassName: GetIngressClassName(ns),
		Rules: []networking.IngressRule{
			{
				IngressRuleValue: networking.IngressRuleValue{
					HTTP: &networking.HTTPIngressRuleValue{
						Paths: []networking.HTTPIngressPath{
							{
								Path:     path,
								PathType: &pathtype,
								Backend: networking.IngressBackend{
									Service: &networking.IngressServiceBackend{
										Name: service,
										Port: networking.ServiceBackendPort{
											Number: int32(port),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// allow ingresses without host field
	if host != "" {
		spec.Rules[0].Host = host
	}

	if len(tlsHosts) > 0 {
		spec.TLS = []networking.IngressTLS{
			{
				Hosts:      tlsHosts,
				SecretName: host,
			},
		}
	}

	return newSingleIngress(name, ns, annotations, spec)
}

// NewSingleIngressWithBackendAndRules creates an ingress with both a default backend and a rule
func NewSingleIngressWithBackendAndRules(name, path, host, ns, defaultService string, defaultPort int, service string, port int, annotations map[string]string) *networking.Ingress {
	pathtype := networking.PathTypePrefix
	spec := networking.IngressSpec{
		IngressClassName: GetIngressClassName(ns),
		DefaultBackend: &networking.IngressBackend{
			Service: &networking.IngressServiceBackend{
				Name: defaultService,
				Port: networking.ServiceBackendPort{
					Number: int32(defaultPort),
				},
			},
		},
		Rules: []networking.IngressRule{
			{
				Host: host,
				IngressRuleValue: networking.IngressRuleValue{
					HTTP: &networking.HTTPIngressRuleValue{
						Paths: []networking.HTTPIngressPath{
							{
								Path:     path,
								PathType: &pathtype,
								Backend: networking.IngressBackend{
									Service: &networking.IngressServiceBackend{
										Name: service,
										Port: networking.ServiceBackendPort{
											Number: int32(port),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return newSingleIngress(name, ns, annotations, spec)
}

// NewSingleCatchAllIngress creates a simple ingress with a catch-all backend
func NewSingleCatchAllIngress(name, ns, service string, port int, annotations map[string]string) *networking.Ingress {
	spec := networking.IngressSpec{
		IngressClassName: GetIngressClassName(ns),
		DefaultBackend: &networking.IngressBackend{
			Service: &networking.IngressServiceBackend{
				Name: service,
				Port: networking.ServiceBackendPort{
					Number: int32(port),
				},
			},
		},
	}
	return newSingleIngress(name, ns, annotations, spec)
}

func newSingleIngress(name, ns string, annotations map[string]string, spec networking.IngressSpec) *networking.Ingress {
	ing := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Spec: spec,
	}

	if annotations == nil {
		annotations = make(map[string]string)
	}

	ing.SetAnnotations(annotations)

	return ing
}

// defaultWaitDuration default sleep time for operations related
// to the API server and NGINX reloads.
var defaultWaitDuration = 5 * time.Second

// Sleep pauses the current goroutine for at least the duration d.
// If no duration is defined, it uses a default
func Sleep(duration ...time.Duration) {
	sleepFor := defaultWaitDuration
	if len(duration) != 0 {
		sleepFor = duration[0]
	}

	time.Sleep(sleepFor)
}

func loadConfig() (*restclient.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	config.UserAgent = "ingress-nginx-e2e"
	return config, nil
}
