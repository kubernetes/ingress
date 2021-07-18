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

package settings

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeSetting("[Security] no-auth-locations", func() {
	f := framework.NewDefaultFramework("no-auth-locations")

	setting := "no-auth-locations"
	username := "foo"
	password := "bar"
	secretName := "test-secret"
	host := "no-auth-locations"
	noAuthPath := "/noauth"

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()

		s := f.EnsureSecret(buildSecret(username, password, secretName, f.Namespace))

		f.UpdateNginxConfigMapData(setting, noAuthPath)

		bi := buildBasicAuthIngressWithSecondPath(host, f.Namespace, s.Name, noAuthPath)
		f.EnsureIngress(bi)
	})

	ginkgo.It("should return status code 401 when accessing '/' unauthentication", func() {
		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "test auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusUnauthorized).
			Body().Contains("401 Authorization Required")
	})

	ginkgo.It("should return status code 200 when accessing '/'  authentication", func() {
		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "test auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithBasicAuth(username, password).
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should return status code 200 when accessing '/noauth' unauthenticated", func() {
		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "test auth")
			})

		f.HTTPTestClient().
			GET("/noauth").
			WithHeader("Host", host).
			WithBasicAuth(username, password).
			Expect().
			Status(http.StatusOK)
	})
})

func buildBasicAuthIngressWithSecondPath(host, namespace, secretName, pathName string) *networking.Ingress {
	pathtype := networking.PathTypePrefix
	return &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      host,
			Namespace: namespace,
			Annotations: map[string]string{"nginx.ingress.kubernetes.io/auth-type": "basic",
				"nginx.ingress.kubernetes.io/auth-secret": secretName,
				"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
			},
		},
		Spec: networking.IngressSpec{
			IngressClassName: framework.GetIngressClassName(namespace),
			Rules: []networking.IngressRule{
				{
					Host: host,
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathtype,
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: framework.EchoService,
											Port: networking.ServiceBackendPort{
												Number: int32(80),
											},
										},
									},
								},
								{
									Path:     pathName,
									PathType: &pathtype,
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: framework.EchoService,
											Port: networking.ServiceBackendPort{
												Number: int32(80),
											},
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
}

func buildSecret(username, password, name, namespace string) *corev1.Secret {
	out, err := exec.Command("openssl", "passwd", "-crypt", password).CombinedOutput()
	assert.Nil(ginkgo.GinkgoT(), err, "creating password")

	encpass := fmt.Sprintf("%v:%s\n", username, out)

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:                       name,
			Namespace:                  namespace,
			DeletionGracePeriodSeconds: framework.NewInt64(1),
		},
		Data: map[string][]byte{
			"auth": []byte(encpass),
		},
		Type: corev1.SecretTypeOpaque,
	}
}
