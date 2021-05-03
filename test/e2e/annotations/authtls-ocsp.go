/*
Copyright 2018 The Kubernetes Authors.

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

package annotations

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/ingress-nginx/test/e2e/framework"
	"k8s.io/ingress-nginx/test/e2e/framework/ocsp"
)

var _ = framework.DescribeAnnotation("auth-tls-ocsp", func() {
	f := framework.NewDefaultFramework("authtls-ocsp")
	o := ocsp.NewFramework(f)

	ginkgo.BeforeEach(func() {
		f.NewEchoDeploymentWithReplicas(2)
	})

	ginkgo.AfterEach(func() {
		if !ginkgo.CurrentGinkgoTestDescription().Failed {
			return
		}

		logs, err := OCSPResponderLogs(f)
		assert.Nil(ginkgo.GinkgoT(), err)

		ginkgo.By("Dumping OCSPServe logs")
		framework.Logf("%v", logs)
	})

	ginkgo.It("should set auth-tls-ocsp", func() {
		host := "authtls.foo.com"
		nameSpace := f.Namespace

		err := o.CreateIngressOcspSecret(
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		clientConfig, err := o.TlsConfig(host)
		assert.NoError(ginkgo.GinkgoT(), err)

		err = o.EnsureOCSPResponderDeployment(nameSpace, "ocspserve")
		assert.NoError(ginkgo.GinkgoT(), err)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-secret":        nameSpace + "/" + host,
			"nginx.ingress.kubernetes.io/auth-tls-verify-client": "on",
			"nginx.ingress.kubernetes.io/auth-tls-verify-depth":  "2",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp":          "on",
		}
		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "ssl_ocsp on;")
			})

		err = framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, "ocspserve", f.Namespace, 1)
		assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")

		f.HTTPTestClient().
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("Should set auth-tls-ocsp, and invalidate revoked certificates", func() {
		host := "authtls.foo.com"
		nameSpace := f.Namespace

		err := o.CreateIngressOcspSecret(
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		err = o.OcspSignCertificates(false, "smime.pem")
		assert.Nil(ginkgo.GinkgoT(), err)

		clientConfig, err := o.TlsConfig(host)
		assert.NoError(ginkgo.GinkgoT(), err)

		ginkgo.By("deploying ocsp")
		err = o.EnsureOCSPResponderDeployment(nameSpace, "ocspserve")
		assert.NoError(ginkgo.GinkgoT(), err)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-secret":        nameSpace + "/" + host,
			"nginx.ingress.kubernetes.io/auth-tls-verify-client": "on",
			"nginx.ingress.kubernetes.io/auth-tls-verify-depth":  "2",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp":          "on",
		}
		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "ssl_ocsp on;")
			})

		err = framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, "ocspserve", f.Namespace, 1)
		assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")

		f.HTTPTestClient().
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)
	})

	// Our ocsp url is set to http://ocspserve.namesapce.svc.cluster.local when generating the certificates. While we would
	// normally deploy ocsp responder as ocspserve, we need to test that the responder annotations overrides the url in the certs.
	//
	// We set the ocsp-responder url to http://responder.namespace.svc.cluster.local and deploy the ocspresonder as responder.
	ginkgo.It("Should set auth-tls-ocsp, auth-tls-ocsp-responder", func() {
		host := "authtls.foo.com"
		nameSpace := f.Namespace

		err := o.CreateIngressOcspSecret(
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		clientConfig, err := o.TlsConfig(host)
		assert.NoError(ginkgo.GinkgoT(), err)

		err = o.EnsureOCSPResponderDeployment(nameSpace, "responder")
		assert.NoError(ginkgo.GinkgoT(), err)

		responder := fmt.Sprintf("http://responder.%v.svc.cluster.local", f.Namespace)
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-secret":         nameSpace + "/" + host,
			"nginx.ingress.kubernetes.io/auth-tls-verify-client":  "on",
			"nginx.ingress.kubernetes.io/auth-tls-verify-depth":   "2",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp":           "on",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp-responder": responder,
		}
		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "ssl_ocsp on;") &&
					strings.Contains(server, fmt.Sprintf("ssl_ocsp_responder %v;", responder))
			})

		err = framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, "responder", f.Namespace, 1)
		assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")

		f.HTTPTestClient().
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("Should set auth-tls-ocsp, auth-tls-ocsp-cache", func() {
		host := "authtls.foo.com"
		nameSpace := f.Namespace

		err := o.CreateIngressOcspSecret(
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		clientConfig, err := o.TlsConfig(host)
		assert.NoError(ginkgo.GinkgoT(), err)

		err = o.EnsureOCSPResponderDeployment(nameSpace, "ocspserve")
		assert.NoError(ginkgo.GinkgoT(), err)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-secret":        nameSpace + "/" + host,
			"nginx.ingress.kubernetes.io/auth-tls-verify-client": "on",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp":          "on",
			"nginx.ingress.kubernetes.io/auth-tls-ocsp-cache":    "shared:foo:15m",
		}
		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "ssl_ocsp on;") &&
					strings.Contains(server, "ssl_ocsp_cache shared:foo:15m;")
			})

		err = framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, "ocspserve", f.Namespace, 1)
		assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")

		f.HTTPTestClient().
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)

		// We should be able to delete the ocsp responder, because we have a cache set.
		err = f.DeleteDeployment("ocspserve")
		assert.NoError(ginkgo.GinkgoT(), err)

		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})
})

func OCSPResponderLogs(f *framework.Framework) (string, error) {
	var pod *core.Pod
	err := wait.Poll(1*time.Second, 2*time.Minute, func() (bool, error) {
		l, err := f.KubeClientSet.
			CoreV1().
			Pods(f.Namespace).
			List(context.TODO(), metav1.ListOptions{LabelSelector: "service=ocspserve"})
		assert.Nil(ginkgo.GinkgoT(), err, "couldn't get ocspserve pods")

		for _, p := range l.Items {
			// make sure the pod is running
			if p.Status.Phase != core.PodRunning {
				continue
			}

			// make sure the pod is ready
			ready := false
			for _, condition := range p.Status.Conditions {
				if condition.Type != core.ContainersReady {
					continue
				}

				ready = condition.Status == core.ConditionTrue
				break
			}
			if !ready {
				continue
			}
			pod = &p
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return "", fmt.Errorf("error getting ocspresponder pod: %v", err)
	}

	logs, err := f.KubeClientSet.CoreV1().RESTClient().Get().
		Resource("pods").
		Namespace(f.Namespace).
		Name(pod.Name).SubResource("log").
		Param("container", "ocspserve").
		Do(context.TODO()).
		Raw()
	if err != nil {
		return "", fmt.Errorf("error getting logs from found ocspresponder pod: %v", err)
	}

	return string(logs), nil
}
