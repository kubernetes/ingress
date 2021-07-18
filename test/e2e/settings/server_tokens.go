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
	"strings"

	"github.com/onsi/ginkgo"

	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeSetting("server-tokens", func() {
	pathtype := networking.PathTypePrefix
	f := framework.NewDefaultFramework("server-tokens")
	serverTokens := "server-tokens"

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})

	ginkgo.It("should not exists Server header in the response", func() {
		f.UpdateNginxConfigMapData(serverTokens, "false")

		f.EnsureIngress(framework.NewSingleIngress(serverTokens, "/", serverTokens, f.Namespace, framework.EchoService, 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, "server_tokens off") &&
					strings.Contains(cfg, "more_clear_headers Server;")
			})
	})

	ginkgo.It("should exists Server header in the response when is enabled", func() {
		f.UpdateNginxConfigMapData(serverTokens, "true")

		f.EnsureIngress(&networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:        serverTokens,
				Namespace:   f.Namespace,
				Annotations: map[string]string{},
			},
			Spec: networking.IngressSpec{
				IngressClassName: &f.IngressClass,
				Rules: []networking.IngressRule{
					{
						Host: serverTokens,
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
								},
							},
						},
					},
				},
			},
		})

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, "server_tokens on")
			})
	})
})
