/*
Copyright 2020 The Kubernetes Authors.

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
	"github.com/stretchr/testify/assert"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.DescribeAnnotation("annotation-global-rate-limit", func() {
	f := framework.NewDefaultFramework("global-rate-limit")
	host := "global-rate-limit-annotation"

	ginkgo.BeforeEach(func() {
		f.NewGlobalRateLimitMemcachedDeployment()
		f.NewEchoDeployment()
	})

	ginkgo.It("generates correct configuration", func() {
		annotations := make(map[string]string)
		annotations["nginx.ingress.kubernetes.io/global-rate-limit"] = "5"
		annotations["nginx.ingress.kubernetes.io/global-rate-limit-window"] = "2m"

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
		ing = f.EnsureIngress(ing)
		namespace := strings.Replace(string(ing.UID), "-", "", -1)

		serverConfig := ""
		f.WaitForNginxServer(host, func(server string) bool {
			serverConfig = server
			return true
		})
		assert.Contains(ginkgo.GinkgoT(), serverConfig,
			fmt.Sprintf(`global_throttle = { namespace = "%v", set_headers = false, `+
				`limit = 5, window_size = 120, key = { { nil, nil, "remote_addr", nil, }, }, `+
				`ignored_cidrs = { } }`,
				namespace))

		f.HTTPTestClient().GET("/").WithHeader("Host", host).Expect().Status(http.StatusOK)

		ginkgo.By("regenerating the correct configuration after update")
		annotations["nginx.ingress.kubernetes.io/global-rate-limit-key"] = "${remote_addr}${http_x_api_client}"
		annotations["nginx.ingress.kubernetes.io/global-rate-limit-ignored-cidrs"] = "192.168.1.1, 234.234.234.0/24"
		annotations["nginx.ingress.kubernetes.io/global-rate-limit-set-headers"] = "true"
		ing.SetAnnotations(annotations)

		f.WaitForReload(func() {
			ing = f.UpdateIngress(ing)
		})
		memcachedHost := `memcached.` + f.Namespace + `.svc.cluster.local`
		f.UpdateNginxConfigMapData("global-rate-limit-memcached-host", memcachedHost)

		serverConfig = ""
		f.WaitForNginxServer(host, func(server string) bool {
			serverConfig = server
			return true
		})

		assert.Contains(ginkgo.GinkgoT(), serverConfig,
			fmt.Sprintf(`global_throttle = { namespace = "%v", set_headers = true, `+
				`limit = 5, window_size = 120, `+
				`key = { { nil, "remote_addr", nil, nil, }, { nil, "http_x_api_client", nil, nil, }, }, `+
				`ignored_cidrs = { "192.168.1.1", "234.234.234.0/24", } }`,
				namespace))

		resp := f.HTTPTestClient().GET("/").
			WithHeader("Host", host).Expect().
			Status(http.StatusOK)

		ginkgo.By("verifying the associated headers are set")
		assert.Equal(ginkgo.GinkgoT(), resp.Header("Ratelimit-Limit").Raw(), fmt.Sprint(5))
		assert.Equal(ginkgo.GinkgoT(), resp.Header("Ratelimit-Remaining").Raw(), fmt.Sprint(5))
		assert.Equal(ginkgo.GinkgoT(), resp.Header("Ratelimit-Window").Raw(), fmt.Sprint(120))
		assert.Contains(ginkgo.GinkgoT(), resp.Headers().Keys().Contains("Ratelimit-Reset").Raw(), "Ratelimit-Reset")

	})

})
