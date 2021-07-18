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

package defaultbackend

import (
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"

	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("[Default Backend] change default settings", func() {
	f := framework.NewDefaultFramework("default-backend-hosts")
	host := "foo.com"

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})

	ginkgo.It("should apply the annotation to the default backend", func() {
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/proxy-buffer-size": "8k",
		}

		ing := &networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "default-backend-annotations",
				Namespace:   f.Namespace,
				Annotations: annotations,
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
						Host: host,
					},
				},
			},
		}

		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "proxy_buffer_size 8k;")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", "foo.com").
			Expect().
			Status(http.StatusOK)
	})
})
