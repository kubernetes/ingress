/*
Copyright 2021 The Kubernetes Authors.
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
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeAnnotation("auth-tls-global", func() {
	f := framework.NewDefaultFramework("authtlsglobal")

	ginkgo.BeforeEach(func() {
		f.NewEchoDeploymentWithReplicas(2)
	})

	ginkgo.It("should should use global-auth-tls-secret when enable-global-tls-auth set to true", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"

		clientConfig, _ := configureGlobalTLSConfig(f, host, globalSecretName)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/enable-global-tls-auth": "true",
		}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, f.Namespace, framework.EchoService, 80, annotations))
		assertGlobalSslClientCertificateConfig(f, host, globalSecretName, "on", "1")

		// Send Request without Client Certs
		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		// Send Request Passing the Client Certs
		f.HTTPTestClientWithTLSConfig(clientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should should not use global-auth-tls-secret when enable-global-tls-auth set to false", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"

		configureGlobalTLSConfig(f, host, globalSecretName)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/enable-global-tls-auth": "false",
		}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, f.Namespace, framework.EchoService, 80, annotations))

		// Send Request without Client Certs
		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should should not use global-auth-tls-secret when enable-global-tls-auth is not set", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"

		configureGlobalTLSConfig(f, host, globalSecretName)

		annotations := map[string]string{}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, f.Namespace, framework.EchoService, 80, annotations))

		// Send Request without Client Certs
		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should should use auth-tls-secret when enable-global-tls-auth set to true and auth-tls-secret is set", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"

		globalClientConfig, tlsClientConfig := configureGlobalTLSConfig(f, host, globalSecretName)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/enable-global-tls-auth": "true",
			"nginx.ingress.kubernetes.io/auth-tls-secret":        f.Namespace + "/" + host,
		}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, f.Namespace, framework.EchoService, 80, annotations))
		assertGlobalSslClientCertificateConfig(f, host, host, "on", "1")

		// Send Request without Client Certs
		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		// Send Request Passing the global Client Certs
		f.HTTPTestClientWithTLSConfig(globalClientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)

		f.HTTPTestClientWithTLSConfig(tlsClientConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)
	})
})

func configureGlobalTLSConfig(f *framework.Framework, host string, globalSecretName string) (*tls.Config, *tls.Config) {
	globalAuthConfig, err := framework.CreateIngressMASecret(
		f.KubeClientSet,
		host,
		globalSecretName,
		f.Namespace)
	assert.Nil(ginkgo.GinkgoT(), err)

	f.UpdateNginxConfigMapData("global-auth-tls-secret", f.Namespace+"/"+globalSecretName)
	f.UpdateNginxConfigMapData("global-auth-tls-verify-client", "on")
	f.UpdateNginxConfigMapData("global-auth-tls-verify-depth", "1")

	serverTLSConfig, err := framework.CreateIngressMASecret(
		f.KubeClientSet,
		host,
		host,
		f.Namespace)
	assert.Nil(ginkgo.GinkgoT(), err)

	return globalAuthConfig, serverTLSConfig
}

func assertGlobalSslClientCertificateConfig(f *framework.Framework, host string, globalSecretName string, verifyClient string, verifyDepth string) {
	sslClientCertDirective := fmt.Sprintf("ssl_client_certificate /etc/ingress-controller/ssl/%s-%s.pem;", f.Namespace, globalSecretName)
	sslVerify := fmt.Sprintf("ssl_verify_client %s;", verifyClient)
	sslVerifyDepth := fmt.Sprintf("ssl_verify_depth %s;", verifyDepth)

	f.WaitForNginxServer(host,
		func(server string) bool {
			return strings.Contains(server, sslClientCertDirective) &&
				strings.Contains(server, sslVerify) &&
				strings.Contains(server, sslVerifyDepth)
		})
}
