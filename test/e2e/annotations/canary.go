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
	"fmt"
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

const (
	canaryService = "echo-canary"
)

var _ = framework.DescribeAnnotation("canary", func() {
	f := framework.NewDefaultFramework("canary")

	ginkgo.BeforeEach(func() {
		// Deployment for main backend
		f.NewEchoDeployment()

		// Deployment for canary backend
		f.NewEchoDeploymentWithNameAndReplicas(canaryService, 1)
	})

	ginkgo.Context("when canary is created", func() {
		ginkgo.It("should response with a 200 status from the mainline upstream when requests are made to the mainline ingress", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})

		ginkgo.It("should return 404 status for requests to the canary if no matching ingress is found", func() {
			host := "foo"

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)

			f.EnsureIngress(canaryIng)

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusNotFound)
		})

		/*

			TODO: This test needs improvements made to the e2e framework so that deployment updates work in order to successfully run

			It("should return the correct status codes when endpoints are unavailable", func() {
				host := "foo"
				annotations := map[string]string{}

				ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
				f.EnsureIngress(ing)

				f.WaitForNginxServer(host,
					func(server string) bool {
						return strings.Contains(server,"server_name foo")
					})

				canaryAnnotations := map[string]string{
					"nginx.ingress.kubernetes.io/canary":           "true",
					"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
				}

				canaryIngName := fmt.Sprintf("%v-canary", host)

				canaryIng := framework.NewSingleIngress(canaryIngName, "/", host, f.Namespace, canaryService,
					80, canaryAnnotations)
				f.EnsureIngress(canaryIng)



				ginkgo.By("returning a 503 status when the mainline deployment has 0 replicas and a request is sent to the canary")

				f.NewEchoDeploymentWithReplicas(0)

				resp, _, errs := gorequest.New().
					Get(f.GetURL(framework.HTTP)).
					Set("Host", host).
					Set("CanaryByHeader", "always").
					End()

				Expect(errs).Should(BeEmpty())
				Expect(resp.StatusCode).Should(Equal(http.StatusServiceUnavailable))

				ginkgo.By("returning a 200 status when the canary deployment has 0 replicas and a request is sent to the mainline ingress")

				f.NewEchoDeploymentWithReplicas(1)
				f.NewDeployment(canaryService, "gcr.io/kubernetes-e2e-test-images/echoserver:2.2", 8080, 0)

				resp, _, errs = gorequest.New().
					Get(f.GetURL(framework.HTTP)).
					Set("Host", host).
					Set("CanaryByHeader", "never").
					End()

				Expect(errs).Should(BeEmpty())
				Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			})
		*/

		ginkgo.It("should route requests to the correct upstream if mainline ingress is created before the canary ingress", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests destined for the mainline ingress to the maineline upstream")

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests destined for the canary ingress to the canary upstream")

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)
		})

		ginkgo.It("should route requests to the correct upstream if mainline ingress is created after the canary ingress", func() {
			host := "foo"

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			ginkgo.By("routing requests destined for the mainline ingress to the mainelin upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests destined for the canary ingress to the canary upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)
		})

		ginkgo.It("should route requests to the correct upstream if the mainline ingress is modified", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			modAnnotations := map[string]string{
				"foo": "bar",
			}

			modIng := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, modAnnotations)

			f.UpdateIngress(modIng)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			ginkgo.By("routing requests destined fro the mainline ingress to the mainline upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests destined for the canary ingress to the canary upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)
		})

		ginkgo.It("should route requests to the correct upstream if the canary ingress is modified", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			newAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader2",
			}

			modIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, newAnnotations)

			f.UpdateIngress(modIng)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			ginkgo.By("routing requests destined for the mainline ingress to the mainline upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader2", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests destined for the canary ingress to the canary upstream")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader2", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)
		})
	})

	ginkgo.Context("when canaried by header with no value", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace,
				framework.EchoService, 80, nil)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the canary upstream when header is set to 'always'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header is set to 'never'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "never").
				Expect().
				Status(http.StatusOK).
				Body().
				Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header is set to anything else")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "badheadervalue").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
	})

	ginkgo.Context("when canaried by header with value", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host,
				f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":                 "true",
				"nginx.ingress.kubernetes.io/canary-by-header":       "CanaryByHeader",
				"nginx.ingress.kubernetes.io/canary-by-header-value": "DoCanary",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the canary upstream when header is set to 'DoCanary'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "DoCanary").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header is set to 'always'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header is set to 'never'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header is set to anything else")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "otherheadervalue").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
	})

	ginkgo.Context("when canaried by header with value and pattern", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":                   "true",
				"nginx.ingress.kubernetes.io/canary-by-header":         "CanaryByHeader",
				"nginx.ingress.kubernetes.io/canary-by-header-pattern": "^Do[A-Z][a-z]+y$",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host, f.Namespace, canaryService,
				80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the canary upstream when header pattern is matched")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "DoCanary").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when header failed to match header value")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "Docanary").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":                   "true",
				"nginx.ingress.kubernetes.io/canary-by-header":         "CanaryByHeader",
				"nginx.ingress.kubernetes.io/canary-by-header-value":   "DoCanary",
				"nginx.ingress.kubernetes.io/canary-by-header-pattern": "^Do[A-Z][a-z]+y$",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host, f.Namespace, canaryService,
				80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the mainline upstream when header is set to 'DoCananry' and header-value is 'DoCanary'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "DoCananry").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
		ginkgo.It("should routes to mainline upstream when the given Regex causes error", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":                   "true",
				"nginx.ingress.kubernetes.io/canary-by-header":         "CanaryByHeader",
				"nginx.ingress.kubernetes.io/canary-by-header-pattern": "[][**da?$&*",
				"nginx.ingress.kubernetes.io/canary-by-cookie":         "CanaryByCookie",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host, f.Namespace, canaryService,
				80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the mainline upstream when the given Regex causes error")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "DoCanary").
				WithCookie("CanaryByCookie", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
	})

	ginkgo.Context("when canaried by header with value and cookie", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host,
				f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":                 "true",
				"nginx.ingress.kubernetes.io/canary-by-header":       "CanaryByHeader",
				"nginx.ingress.kubernetes.io/canary-by-header-value": "DoCanary",
				"nginx.ingress.kubernetes.io/canary-by-cookie":       "CanaryByCookie",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the canary upstream when header value does not match and cookie is set to 'always'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithHeader("CanaryByHeader", "otherheadervalue").
				WithCookie("CanaryByCookie", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)
		})
	})

	ginkgo.Context("when canaried by cookie", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host,
				f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-cookie": "Canary-By-Cookie",
			}

			canaryIngName := fmt.Sprintf("%v-canary", host)

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			ginkgo.By("routing requests to the canary upstream when cookie is set to 'always'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithCookie("Canary-By-Cookie", "always").
				Expect().
				Status(http.StatusOK).
				Body().Contains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when cookie is set to 'never'")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithCookie("Canary-By-Cookie", "never").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)

			ginkgo.By("routing requests to the mainline upstream when cookie is set to anything else")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithCookie("Canary-By-Cookie", "badcookievalue").
				Expect().
				Status(http.StatusOK).
				Body().Contains(framework.EchoService).NotContains(canaryService)
		})
	})

	// TODO: add testing for canary-weight 0 < weight < 100
	ginkgo.Context("when canaried by weight", func() {
		ginkgo.It("should route requests to the correct upstream", func() {
			host := "foo"
			annotations := map[string]string{}

			ing := framework.NewSingleIngress(host, "/", host,
				f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			canaryIngName := fmt.Sprintf("%v-canary", host)
			canaryAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":        "true",
				"nginx.ingress.kubernetes.io/canary-weight": "0",
			}

			canaryIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, canaryAnnotations)
			f.EnsureIngress(canaryIng)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			ginkgo.By("returning requests from the mainline only when weight is equal to 0")
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK).
				Body().
				Contains(framework.EchoService).
				NotContains(canaryService)

			ginkgo.By("returning requests from the canary only when weight is equal to 100")

			newAnnotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":        "true",
				"nginx.ingress.kubernetes.io/canary-weight": "100",
			}

			modIng := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, newAnnotations)

			f.UpdateIngress(modIng)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "server_name foo")
				})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK).
				Body().
				Contains(canaryService)
		})
	})

	ginkgo.Context("Single canary Ingress", func() {
		ginkgo.It("should not use canary as a catch-all server", func() {
			host := "foo"
			canaryIngName := fmt.Sprintf("%v-canary", host)
			annotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			ing := framework.NewSingleCatchAllIngress(canaryIngName,
				f.Namespace, canaryService, 80, annotations)
			f.EnsureIngress(ing)

			ing = framework.NewSingleCatchAllIngress(host, f.Namespace,
				framework.EchoService, 80, nil)
			f.EnsureIngress(ing)

			f.WaitForNginxServer("_",
				func(server string) bool {
					upstreamName := fmt.Sprintf(`set $proxy_upstream_name "%s-%s-%s";`, f.Namespace, framework.EchoService, "80")
					canaryUpstreamName := fmt.Sprintf(`set $proxy_upstream_name "%s-%s-%s";`, f.Namespace, canaryService, "80")

					return strings.Contains(server, fmt.Sprintf(`set $ingress_name "%v";`, host)) &&
						!strings.Contains(server, `set $proxy_upstream_name "upstream-default-backend";`) &&
						!strings.Contains(server, canaryUpstreamName) &&
						strings.Contains(server, upstreamName)
				})
		})

		ginkgo.It("should not use canary with domain as a server", func() {
			host := "foo"
			canaryIngName := fmt.Sprintf("%v-canary", host)
			annotations := map[string]string{
				"nginx.ingress.kubernetes.io/canary":           "true",
				"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
			}

			ing := framework.NewSingleIngress(canaryIngName, "/", host,
				f.Namespace, canaryService, 80, annotations)
			f.EnsureIngress(ing)

			otherHost := "bar"
			ing = framework.NewSingleIngress(otherHost, "/", otherHost,
				f.Namespace, framework.EchoService, 80, nil)
			f.EnsureIngress(ing)

			f.WaitForNginxConfiguration(func(cfg string) bool {
				return strings.Contains(cfg, "server_name "+otherHost) &&
					!strings.Contains(cfg, "server_name "+host)
			})
		})
	})

	ginkgo.It("does not crash when canary ingress has multiple paths to the same non-matching backend", func() {
		host := "foo"
		canaryIngName := fmt.Sprintf("%v-canary", host)
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/canary":           "true",
			"nginx.ingress.kubernetes.io/canary-by-header": "CanaryByHeader",
		}

		paths := []string{"/foo", "/bar"}
		ing := framework.NewSingleIngressWithMultiplePaths(canaryIngName, paths, host,
			f.Namespace, "httpy-svc-canary", 80, annotations)
		f.EnsureIngress(ing)

		ing = framework.NewSingleIngress(host, "/", host, f.Namespace,
			framework.EchoService, 80, nil)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name foo")
			})
	})
})
