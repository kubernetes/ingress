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

package ingress

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("[Ingress] definition without host", func() {
	f := framework.NewDefaultFramework("ingress-without-host")

	ginkgo.It("should set ingress details variables for ingresses without a host", func() {
		f.NewEchoDeployment()

		ing := framework.NewSingleIngress("default-host", "/", "", f.Namespace, framework.EchoService, 80, nil)
		f.EnsureIngress(ing)

		f.WaitForNginxServer("_",
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf(`set $namespace "%v";`, f.Namespace)) &&
					strings.Contains(server, fmt.Sprintf(`set $ingress_name "%v";`, ing.Name)) &&
					strings.Contains(server, fmt.Sprintf(`set $service_name "%v";`, framework.EchoService)) &&
					strings.Contains(server, `set $service_port "80";`) &&
					strings.Contains(server, `set $location_path "/";`)
			})

		f.HTTPTestClient().
			GET("/").
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should set ingress details variables for ingresses with host without IngressRuleValue, only Backend", func() {
		f.NewEchoDeployment()

		ing := &networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "backend",
				Namespace: f.Namespace,
			},
			Spec: networking.IngressSpec{
				IngressClassName: &f.IngressClass,
				DefaultBackend: &networking.IngressBackend{
					Service: &networking.IngressServiceBackend{
						Name: framework.EchoService,
						Port: networking.ServiceBackendPort{
							Number: int32(80),
						},
					},
				},
				Rules: []networking.IngressRule{
					{
						Host: "only-backend",
					},
				},
			},
		}
		f.EnsureIngress(ing)

		f.WaitForNginxServer("only-backend",
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf(`set $namespace "%v";`, f.Namespace)) &&
					strings.Contains(server, fmt.Sprintf(`set $ingress_name "%v";`, ing.Name)) &&
					strings.Contains(server, fmt.Sprintf(`set $service_name "%v";`, framework.EchoService)) &&
					strings.Contains(server, `set $service_port "80";`) &&
					strings.Contains(server, `set $location_path "/";`)
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", "only-backend").
			Expect().
			Status(http.StatusOK)
	})
})
