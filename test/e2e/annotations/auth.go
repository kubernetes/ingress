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

package annotations

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeAnnotation("auth-*", func() {
	f := framework.NewDefaultFramework("auth")

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})

	ginkgo.It("should return status code 200 when no authentication is configured", func() {
		host := "auth"

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

	ginkgo.It("should return status code 503 when authentication is configured with an invalid secret", func() {
		host := "auth"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":   "basic",
			"nginx.ingress.kubernetes.io/auth-secret": "something",
			"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusServiceUnavailable).
			Body().Contains("503 Service Temporarily Unavailable")
	})

	ginkgo.It("should return status code 401 when authentication is configured but Authorization header is not configured", func() {
		host := "auth"

		s := f.EnsureSecret(buildSecret("foo", "bar", "test", f.Namespace))

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":   "basic",
			"nginx.ingress.kubernetes.io/auth-secret": s.Name,
			"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusUnauthorized).
			Body().Contains("401 Authorization Required")
	})

	ginkgo.It("should return status code 401 when authentication is configured and Authorization header is sent with invalid credentials", func() {
		host := "auth"

		s := f.EnsureSecret(buildSecret("foo", "bar", "test", f.Namespace))

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":   "basic",
			"nginx.ingress.kubernetes.io/auth-secret": s.Name,
			"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithBasicAuth("user", "pass").
			Expect().
			Status(http.StatusUnauthorized).
			Body().Contains("401 Authorization Required")
	})

	ginkgo.It("should return status code 200 when authentication is configured and Authorization header is sent", func() {
		host := "auth"

		s := f.EnsureSecret(buildSecret("foo", "bar", "test", f.Namespace))

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":   "basic",
			"nginx.ingress.kubernetes.io/auth-secret": s.Name,
			"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithBasicAuth("foo", "bar").
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should return status code 200 when authentication is configured with a map and Authorization header is sent", func() {
		host := "auth"

		s := f.EnsureSecret(buildMapSecret("foo", "bar", "test", f.Namespace))

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":        "basic",
			"nginx.ingress.kubernetes.io/auth-secret":      s.Name,
			"nginx.ingress.kubernetes.io/auth-secret-type": "auth-map",
			"nginx.ingress.kubernetes.io/auth-realm":       "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithBasicAuth("foo", "bar").
			Expect().
			Status(http.StatusOK)
	})

	ginkgo.It("should return status code 401 when authentication is configured with invalid content and Authorization header is sent", func() {
		host := "auth"

		s := f.EnsureSecret(
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: f.Namespace,
				},
				Data: map[string][]byte{
					// invalid content
					"auth": []byte("foo:"),
				},
				Type: corev1.SecretTypeOpaque,
			},
		)

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-type":   "basic",
			"nginx.ingress.kubernetes.io/auth-secret": s.Name,
			"nginx.ingress.kubernetes.io/auth-realm":  "test auth",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithBasicAuth("foo", "bar").
			Expect().
			Status(http.StatusUnauthorized)
	})

	ginkgo.It(`should set snippet "proxy_set_header My-Custom-Header 42;" when external auth is configured`, func() {
		host := "auth"

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-url": "http://foo.bar/basic-auth/user/password",
			"nginx.ingress.kubernetes.io/auth-snippet": `
				proxy_set_header My-Custom-Header 42;`,
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, `proxy_set_header My-Custom-Header 42;`)
			})
	})

	ginkgo.It(`should not set snippet "proxy_set_header My-Custom-Header 42;" when external auth is not configured`, func() {
		host := "auth"

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-snippet": `
				proxy_set_header My-Custom-Header 42;`,
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return !strings.Contains(server, `proxy_set_header My-Custom-Header 42;`)
			})
	})

	ginkgo.It(`should set "proxy_set_header 'My-Custom-Header' '42';" when auth-headers are set`, func() {
		host := "auth"

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-url":               "http://foo.bar/basic-auth/user/password",
			"nginx.ingress.kubernetes.io/auth-proxy-set-headers": f.Namespace + "/auth-headers",
		}

		f.CreateConfigMap("auth-headers", map[string]string{
			"My-Custom-Header": "42",
		})

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, `proxy_set_header 'My-Custom-Header' '42';`)
			})
	})

	ginkgo.It(`should set cache_key when external auth cache is configured`, func() {
		host := "auth"

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-url":            "http://foo.bar/basic-auth/user/password",
			"nginx.ingress.kubernetes.io/auth-cache-key":      "foo",
			"nginx.ingress.kubernetes.io/auth-cache-duration": "200 202 401 30m",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		cacheRegex := regexp.MustCompile(`\$cache_key.*foo`)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return cacheRegex.MatchString(server) &&
					strings.Contains(server, `proxy_cache_valid 200 202 401 30m;`)

			})
	})

	ginkgo.It("retains cookie set by external authentication server", func() {
		host := "auth-check-cookies"

		cfg := `#
events {
	worker_connections  1024;
	multi_accept on;
}

http {
	default_type 'text/plain';
	client_max_body_size 0;

	server {
		access_log on;
		access_log /dev/stdout;

		listen 80;

		location ~ ^/cookies/set/(?<key>.*)/(?<value>.*) {
			content_by_lua_block {
				ngx.header['Set-Cookie'] = {ngx.var.key.."="..ngx.var.value}
				ngx.say("OK")
			}
		}

		location / {
			return 200;
		}
	}
}
`

		f.NGINXWithConfigDeployment(framework.HTTPBinService, cfg, 1)

		e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), framework.HTTPBinService, metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets), 1, "expected at least one endpoint")
		assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets[0].Addresses), 1, "expected at least one address ready in the endpoint")

		httpbinIP := e.Subsets[0].Addresses[0].IP

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/auth-url":    fmt.Sprintf("http://%s/cookies/set/alma/armud", httpbinIP),
			"nginx.ingress.kubernetes.io/auth-signin": "http://$host/auth/start",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host, func(server string) bool {
			return strings.Contains(server, "server_name auth")
		})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			WithQuery("a", "b").
			WithQuery("c", "d").
			Expect().
			Status(http.StatusOK).
			Header("Set-Cookie").Contains("alma=armud")
	})

	ginkgo.Context("when external authentication is configured", func() {
		host := "auth"
		var annotations map[string]string
		var ing *networking.Ingress

		ginkgo.BeforeEach(func() {
			f.NewHttpbinDeployment()

			err := framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, framework.HTTPBinService, f.Namespace, 1)
			assert.Nil(ginkgo.GinkgoT(), err)

			e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), framework.HTTPBinService, metav1.GetOptions{})
			assert.Nil(ginkgo.GinkgoT(), err)

			assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets), 1, "expected at least one endpoint")
			assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets[0].Addresses), 1, "expected at least one address ready in the endpoint")

			httpbinIP := e.Subsets[0].Addresses[0].IP

			annotations = map[string]string{
				"nginx.ingress.kubernetes.io/auth-url":    fmt.Sprintf("http://%s/basic-auth/user/password", httpbinIP),
				"nginx.ingress.kubernetes.io/auth-signin": "http://$host/auth/start",
			}

			ing = framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host, func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})
		})

		ginkgo.It("should return status code 200 when signed in", func() {
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It("should redirect to signin url when not signed in", func() {
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithQuery("a", "b").
				WithQuery("c", "d").
				Expect().
				Status(http.StatusFound).
				Header("Location").Equal(fmt.Sprintf("http://%s/auth/start?rd=http://%s%s", host, host, url.QueryEscape("/?a=b&c=d")))
		})

		ginkgo.It("keeps processing new ingresses even if one of the existing ingresses is misconfigured", func() {
			annotations["nginx.ingress.kubernetes.io/auth-type"] = "basic"
			annotations["nginx.ingress.kubernetes.io/auth-secret"] = "something"
			annotations["nginx.ingress.kubernetes.io/auth-realm"] = "test auth"
			f.UpdateIngress(ing)

			anotherHost := "different"
			anotherAnnotations := map[string]string{}

			anotherIng := framework.NewSingleIngress(anotherHost, "/", anotherHost, f.Namespace, framework.EchoService, 80, anotherAnnotations)
			f.EnsureIngress(anotherIng)

			f.WaitForNginxServer(anotherHost,
				func(server string) bool {
					return strings.Contains(server, "server_name "+anotherHost)
				})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", anotherHost).
				Expect().
				Status(http.StatusOK)
		})
	})

	ginkgo.Context("when external authentication is configured with a custom redirect param", func() {
		host := "auth"
		var annotations map[string]string
		var ing *networking.Ingress

		ginkgo.BeforeEach(func() {
			f.NewHttpbinDeployment()

			var httpbinIP string

			err := framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, framework.HTTPBinService, f.Namespace, 1)
			assert.Nil(ginkgo.GinkgoT(), err)

			e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), framework.HTTPBinService, metav1.GetOptions{})
			assert.Nil(ginkgo.GinkgoT(), err)

			httpbinIP = e.Subsets[0].Addresses[0].IP

			annotations = map[string]string{
				"nginx.ingress.kubernetes.io/auth-url":                   fmt.Sprintf("http://%s/basic-auth/user/password", httpbinIP),
				"nginx.ingress.kubernetes.io/auth-signin":                "http://$host/auth/start",
				"nginx.ingress.kubernetes.io/auth-signin-redirect-param": "orig",
			}

			ing = framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host, func(server string) bool {
				return strings.Contains(server, "server_name auth")
			})
		})

		ginkgo.It("should return status code 200 when signed in", func() {
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It("should redirect to signin url when not signed in", func() {
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				WithQuery("a", "b").
				WithQuery("c", "d").
				Expect().
				Status(http.StatusFound).
				Header("Location").Equal(fmt.Sprintf("http://%s/auth/start?orig=http://%s%s", host, host, url.QueryEscape("/?a=b&c=d")))
		})

		ginkgo.It("keeps processing new ingresses even if one of the existing ingresses is misconfigured", func() {
			annotations["nginx.ingress.kubernetes.io/auth-type"] = "basic"
			annotations["nginx.ingress.kubernetes.io/auth-secret"] = "something"
			annotations["nginx.ingress.kubernetes.io/auth-realm"] = "test auth"
			f.UpdateIngress(ing)

			anotherHost := "different"
			anotherAnnotations := map[string]string{}

			anotherIng := framework.NewSingleIngress(anotherHost, "/", anotherHost, f.Namespace, framework.EchoService, 80, anotherAnnotations)
			f.EnsureIngress(anotherIng)

			f.WaitForNginxServer(anotherHost,
				func(server string) bool {
					return strings.Contains(server, "server_name "+anotherHost)
				})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", anotherHost).
				Expect().
				Status(http.StatusOK)
		})
	})

	ginkgo.Context("when external authentication with caching is configured", func() {
		thisHost := "auth"
		thatHost := "different"

		fooPath := "/foo"
		barPath := "/bar"

		ginkgo.BeforeEach(func() {
			f.NewHttpbinDeployment()

			err := framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, framework.HTTPBinService, f.Namespace, 1)
			assert.Nil(ginkgo.GinkgoT(), err)

			framework.Sleep(1 * time.Second)

			e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), framework.HTTPBinService, metav1.GetOptions{})
			assert.Nil(ginkgo.GinkgoT(), err)

			assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets), 1, "expected at least one endpoint")
			assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets[0].Addresses), 1, "expected at least one address ready in the endpoint")

			httpbinIP := e.Subsets[0].Addresses[0].IP

			annotations := map[string]string{
				"nginx.ingress.kubernetes.io/auth-url":            fmt.Sprintf("http://%s/basic-auth/user/password", httpbinIP),
				"nginx.ingress.kubernetes.io/auth-signin":         "http://$host/auth/start",
				"nginx.ingress.kubernetes.io/auth-cache-key":      "fixed",
				"nginx.ingress.kubernetes.io/auth-cache-duration": "200 201 401 30m",
			}

			for _, host := range []string{thisHost, thatHost} {
				ginkgo.By("Adding an ingress rule for /foo")
				fooIng := framework.NewSingleIngress(fmt.Sprintf("foo-%s-ing", host), fooPath, host, f.Namespace, framework.EchoService, 80, annotations)
				f.EnsureIngress(fooIng)
				f.WaitForNginxServer(host, func(server string) bool {
					return strings.Contains(server, "location /foo")
				})

				ginkgo.By("Adding an ingress rule for /bar")
				barIng := framework.NewSingleIngress(fmt.Sprintf("bar-%s-ing", host), barPath, host, f.Namespace, framework.EchoService, 80, annotations)
				f.EnsureIngress(barIng)
				f.WaitForNginxServer(host, func(server string) bool {
					return strings.Contains(server, "location /bar")
				})
			}

			framework.Sleep()
		})

		ginkgo.It("should return status code 200 when signed in after auth backend is deleted ", func() {
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			err := f.DeleteDeployment(framework.HTTPBinService)
			assert.Nil(ginkgo.GinkgoT(), err)
			framework.Sleep()

			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)
		})

		ginkgo.It("should deny login for different location on same server", func() {
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			err := f.DeleteDeployment(framework.HTTPBinService)
			assert.Nil(ginkgo.GinkgoT(), err)
			framework.Sleep()

			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			ginkgo.By("receiving an internal server error without cache on location /bar")
			f.HTTPTestClient().
				GET(barPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusInternalServerError)
		})

		ginkgo.It("should deny login for different servers", func() {
			ginkgo.By("logging into server thisHost /foo")
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			err := f.DeleteDeployment(framework.HTTPBinService)
			assert.Nil(ginkgo.GinkgoT(), err)
			framework.Sleep()

			ginkgo.By("receiving an internal server error without cache on thisHost location /bar")
			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thisHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusOK)

			f.HTTPTestClient().
				GET(fooPath).
				WithHeader("Host", thatHost).
				WithBasicAuth("user", "password").
				Expect().
				Status(http.StatusInternalServerError)
		})

		ginkgo.It("should redirect to signin url when not signed in", func() {
			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", thisHost).
				WithQuery("a", "b").
				WithQuery("c", "d").
				Expect().
				Status(http.StatusFound).
				Header("Location").Equal(fmt.Sprintf("http://%s/auth/start?rd=http://%s%s", thisHost, thisHost, url.QueryEscape("/?a=b&c=d")))
		})
	})
})

// TODO: test Digest Auth
//   401
//   Realm name
//   Auth ok
//   Auth error

func buildSecret(username, password, name, namespace string) *corev1.Secret {
	out, err := exec.Command("openssl", "passwd", "-crypt", password).CombinedOutput()
	encpass := fmt.Sprintf("%v:%s\n", username, out)
	assert.Nil(ginkgo.GinkgoT(), err)

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

func buildMapSecret(username, password, name, namespace string) *corev1.Secret {
	out, err := exec.Command("openssl", "passwd", "-crypt", password).CombinedOutput()
	assert.Nil(ginkgo.GinkgoT(), err)

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:                       name,
			Namespace:                  namespace,
			DeletionGracePeriodSeconds: framework.NewInt64(1),
		},
		Data: map[string][]byte{
			username: []byte(out),
		},
		Type: corev1.SecretTypeOpaque,
	}
}
