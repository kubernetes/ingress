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
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/internal/file"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeAnnotation("auth-tls-global", func() {
	f := framework.NewDefaultFramework("authtlsglobal")

	ginkgo.BeforeEach(func() {
		f.NewEchoDeploymentWithReplicas(2)
	})

	ginkgo.It("should use auth tls secret even when global ssl auth is set", func() {
		globalSecretName := "auth-tls-global-secret"
		host := "authtls.foo.com"
		globalAuthConfig := configureIngressWithGlobalSSLAuth(f, host, globalSecretName)

		f.UpdateNginxConfigMapData("global-ssl-client-certificate", "/etc/ingress-controller/ssl/global-auth-ssl/ca.crt")
		nameSpace := f.Namespace

		clientConfig, err := framework.CreateIngressMASecret(
			f.KubeClientSet,
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-secret":        nameSpace + "/" + host,
			"nginx.ingress.kubernetes.io/auth-tls-verify-client": "on",
		}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))
		assertSslClientCertificateConfig(f, host, "on", "1")

		framework.Sleep(30 * time.Minute)
		ioutil.WriteFile("/etc/ssl/nginx-host", []byte(f.GetURL(framework.HTTPS)), file.ReadWriteByUser)

		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
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

		f.HTTPTestClientWithTLSConfig(globalAuthConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusBadRequest)
	})

	ginkgo.It("should use global ssl auth when auth-tls-verify-client is set to on and auth-tls-secret is not defined", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"
		nameSpace := f.Namespace
		globalAuthConfig := configureIngressWithGlobalSSLAuth(f, host, globalSecretName)

		f.UpdateNginxConfigMapData("global-ssl-client-certificate", "/ssl/global-auth-ssl/ca.crt")

		clientConfig, err := framework.CreateIngressMASecret(
			f.KubeClientSet,
			host,
			host,
			nameSpace)
		assert.Nil(ginkgo.GinkgoT(), err)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-tls-verify-client": "on",
		}

		f.EnsureIngress(framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, nameSpace, framework.EchoService, 80, annotations))
		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "ssl_client_certificate /ssl/global-auth-ssl/ca.crt;") &&
					strings.Contains(server, "ssl_verify_client on;")
			})

		ioutil.WriteFile("/etc/ssl/nginx-host", []byte(f.GetURL(framework.HTTPS)), file.ReadWriteByUser)
		// framework.Sleep(20 * time.Minute)

		f.HTTPTestClientWithTLSConfig(&tls.Config{ServerName: host, InsecureSkipVerify: true}).
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

		f.HTTPTestClientWithTLSConfig(globalAuthConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)

	})

	ginkgo.It("should not use global ssl auth when auth-tls-verify-client is not defined", func() {
		host := "authtls.foo.com"
		globalSecretName := "auth-tls-global-secret"
		configureIngressWithGlobalSSLAuth(f, host, globalSecretName)

		f.UpdateNginxConfigMapData("global-ssl-client-certificate", "/ssl/global-auth-ssl/ca.crt")

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, nil)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK).
			Body().Contains(fmt.Sprintf("host=%v", host))
	})
})

func configureIngressWithGlobalSSLAuth(f *framework.Framework, host string, globalSecretName string) *tls.Config {
	globalAuthConfig, err := framework.CreateIngressMASecret(
		f.KubeClientSet,
		host,
		globalSecretName,
		f.Namespace)
	assert.Nil(ginkgo.GinkgoT(), err)

	globalAuthVolume := []corev1.Volume{
		{
			Name: "global-auth-ssl",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: globalSecretName,
				},
			},
		},
	}

	globalAuthVolumeMount := []corev1.VolumeMount{
		{
			Name:      "global-auth-ssl",
			MountPath: "/ssl/global-auth-ssl",
			ReadOnly:  true,
		},
	}
	err = f.UpdateIngressControllerDeployment(func(deployment *appsv1.Deployment) error {
		volumes := deployment.Spec.Template.Spec.Volumes
		volumes = append(volumes, globalAuthVolume...)
		deployment.Spec.Template.Spec.Volumes = volumes

		volumeMounts := deployment.Spec.Template.Spec.Containers[0].VolumeMounts
		volumeMounts = append(volumeMounts, globalAuthVolumeMount...)
		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
		assert.Nil(ginkgo.GinkgoT(), err, "updating deployment")
		_, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		return err
	})
	assert.Nil(ginkgo.GinkgoT(), err, "updating ingress controller deployment")
	return globalAuthConfig
}
