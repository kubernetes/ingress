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
	"regexp"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	networking "k8s.io/api/networking/v1"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeSetting("[Security] global-auth-url", func() {
	f := framework.NewDefaultFramework("global-external-auth")

	host := "global-external-auth"

	echoServiceName := framework.EchoService

	globalExternalAuthURLSetting := "global-auth-url"

	fooPath := "/foo"
	barPath := "/bar"

	noAuthSetting := "no-auth-locations"
	noAuthLocations := barPath

	enableGlobalExternalAuthAnnotation := "nginx.ingress.kubernetes.io/enable-global-auth"

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
		f.NewHttpbinDeployment()
	})

	ginkgo.Context("when global external authentication is configured", func() {

		ginkgo.BeforeEach(func() {
			globalExternalAuthURL := fmt.Sprintf("http://%s.%s.svc.cluster.local:80/status/401", framework.HTTPBinService, f.Namespace)

			ginkgo.By("Adding an ingress rule for /foo")
			fooIng := framework.NewSingleIngress("foo-ingress", fooPath, host, f.Namespace, echoServiceName, 80, nil)
			f.EnsureIngress(fooIng)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "location /foo")
				})

			ginkgo.By("Adding an ingress rule for /bar")
			barIng := framework.NewSingleIngress("bar-ingress", barPath, host, f.Namespace, echoServiceName, 80, nil)
			f.EnsureIngress(barIng)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "location /bar")
				})

			ginkgo.By("Adding a global-auth-url to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthURLSetting, globalExternalAuthURL)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, globalExternalAuthURL)
				})
		})

		ginkgo.It("should return status code 401 when request any protected service", func() {

			ginkgo.By("Sending a request to protected service /foo")
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusUnauthorized)

			ginkgo.By("Sending a request to protected service /bar")
			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusUnauthorized)
		})

		ginkgo.It("should return status code 200 when request whitelisted (via no-auth-locations) service and 401 when request protected service", func() {

			ginkgo.By("Adding a no-auth-locations for /bar to configMap")
			f.UpdateNginxConfigMapData(noAuthSetting, noAuthLocations)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, noAuthLocations)
				})

			ginkgo.By("Sending a request to protected service /foo")
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusUnauthorized)

			ginkgo.By("Sending a request to whitelisted service /bar")
			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It("should return status code 200 when request whitelisted (via ingress annotation) service and 401 when request protected service", func() {

			ginkgo.By("Adding an ingress rule for /bar with annotation enable-global-auth = false")
			err := framework.UpdateIngress(f.KubeClientSet, f.Namespace, "bar-ingress", func(ingress *networking.Ingress) error {
				ingress.ObjectMeta.Annotations[enableGlobalExternalAuthAnnotation] = "false"
				return nil
			})
			assert.Nil(ginkgo.GinkgoT(), err)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "location /bar")
				})

			ginkgo.By("Sending a request to protected service /foo")
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusUnauthorized)

			ginkgo.By("Sending a request to whitelisted service /bar")
			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It("should still return status code 200 after auth backend is deleted using cache ", func() {

			globalExternalAuthCacheKeySetting := "global-auth-cache-key"
			globalExternalAuthCacheKey := "foo"
			globalExternalAuthCacheDurationSetting := "global-auth-cache-duration"
			globalExternalAuthCacheDuration := "200 201 401 30m"
			globalExternalAuthURL := fmt.Sprintf("http://%s.%s.svc.cluster.local:80/status/200", framework.HTTPBinService, f.Namespace)

			ginkgo.By("Adding a global-auth-cache-key to configMap")
			f.SetNginxConfigMapData(map[string]string{
				globalExternalAuthCacheKeySetting:      globalExternalAuthCacheKey,
				globalExternalAuthCacheDurationSetting: globalExternalAuthCacheDuration,
				globalExternalAuthURLSetting:           globalExternalAuthURL,
			})

			cacheRegex := regexp.MustCompile(`\$cache_key.*foo`)

			f.WaitForNginxServer(host,
				func(server string) bool {
					return cacheRegex.MatchString(server) &&
						strings.Contains(server, `proxy_cache_valid 200 201 401 30m;`)
				})

			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", host).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			err := f.DeleteDeployment(framework.HTTPBinService)
			assert.Nil(ginkgo.GinkgoT(), err)
			framework.Sleep()

			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", host).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It(`should proxy_method method when global-auth-method is configured`, func() {

			globalExternalAuthMethodSetting := "global-auth-method"
			globalExternalAuthMethod := "GET"

			ginkgo.By("Adding a global-auth-method to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthMethodSetting, globalExternalAuthMethod)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "proxy_method")
				})
		})

		ginkgo.It(`should add custom error page when global-auth-signin url is configured`, func() {

			globalExternalAuthSigninSetting := "global-auth-signin"
			globalExternalAuthSignin := "http://foo.com/global-error-page"

			ginkgo.By("Adding a global-auth-signin to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthSigninSetting, globalExternalAuthSignin)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "error_page 401 = ")
				})
		})

		ginkgo.It(`should add auth headers when global-auth-response-headers is configured`, func() {

			globalExternalAuthResponseHeadersSetting := "global-auth-response-headers"
			globalExternalAuthResponseHeaders := "Foo, Bar"

			ginkgo.By("Adding a global-auth-response-headers to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthResponseHeadersSetting, globalExternalAuthResponseHeaders)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, "auth_request_set $authHeader0 $upstream_http_foo;") &&
						strings.Contains(server, "auth_request_set $authHeader1 $upstream_http_bar;")
				})
		})

		ginkgo.It(`should set request-redirect when global-auth-request-redirect is configured`, func() {

			globalExternalAuthRequestRedirectSetting := "global-auth-request-redirect"
			globalExternalAuthRequestRedirect := "Foo-Redirect"

			ginkgo.By("Adding a global-auth-request-redirect to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthRequestRedirectSetting, globalExternalAuthRequestRedirect)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, globalExternalAuthRequestRedirect)
				})
		})

		ginkgo.It(`should set snippet when global external auth is configured`, func() {
			globalExternalAuthSnippetSetting := "global-auth-snippet"
			globalExternalAuthSnippet := "proxy_set_header My-Custom-Header 42;"

			ginkgo.By("Adding a global-auth-snippet to configMap")
			f.UpdateNginxConfigMapData(globalExternalAuthSnippetSetting, globalExternalAuthSnippet)
			f.WaitForNginxServer(host,
				func(server string) bool {
					return strings.Contains(server, globalExternalAuthSnippet)
				})
		})

	})

})
