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

package servicebackend

import (
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

var pathtype = networking.PathTypePrefix
var _ = framework.IngressNginxDescribe("[Service] backend status code 503", func() {
	f := framework.NewDefaultFramework("service-backend")

	ginkgo.It("should return 503 when backend service does not exist", func() {
		host := "nonexistent.svc.com"

		bi := buildIngressWithNonexistentService(host, f.Namespace, "/")
		f.EnsureIngress(bi)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "proxy_pass http://upstream_balancer;")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusServiceUnavailable)
	})

	ginkgo.It("should return 503 when all backend service endpoints are unavailable", func() {
		host := "unavailable.svc.com"

		bi, bs := buildIngressWithUnavailableServiceEndpoints(host, f.Namespace, "/")

		f.EnsureService(bs)
		f.EnsureIngress(bi)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "proxy_pass http://upstream_balancer;")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusServiceUnavailable)
	})
})

func buildIngressWithNonexistentService(host, namespace, path string) *networking.Ingress {
	backendService := "nonexistent-svc"
	return &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      host,
			Namespace: namespace,
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
									Path:     path,
									PathType: &pathtype,
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: backendService,
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

func buildIngressWithUnavailableServiceEndpoints(host, namespace, path string) (*networking.Ingress, *corev1.Service) {
	backendService := "unavailable-svc"
	return &networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      host,
				Namespace: namespace,
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
										Path:     path,
										PathType: &pathtype,
										Backend: networking.IngressBackend{
											Service: &networking.IngressServiceBackend{
												Name: backendService,
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
		}, &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      backendService,
				Namespace: namespace,
			},
			Spec: corev1.ServiceSpec{Ports: []corev1.ServicePort{
				{
					Name:       "tcp",
					Port:       80,
					TargetPort: intstr.FromInt(80),
					Protocol:   "TCP",
				},
			},
				Selector: map[string]string{
					"app": backendService,
				},
			},
		}
}
