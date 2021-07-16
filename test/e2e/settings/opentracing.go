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

package settings

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

const (
	enableOpentracing            = "enable-opentracing"
	opentracingTrustIncomingSpan = "opentracing-trust-incoming-span"

	zipkinCollectorHost = "zipkin-collector-host"

	jaegerCollectorHost     = "jaeger-collector-host"
	jaegerSamplerHost       = "jaeger-sampler-host"
	jaegerPropagationFormat = "jaeger-propagation-format"
	// jaegerEndpoint      = "jaeger-endpoint"

	datadogCollectorHost = "datadog-collector-host"

	opentracingOperationName         = "opentracing-operation-name"
	opentracingLocationOperationName = "opentracing-location-operation-name"
)

var _ = framework.IngressNginxDescribe("Configure OpenTracing", func() {
	f := framework.NewDefaultFramework("enable-opentracing")

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})

	ginkgo.AfterEach(func() {
	})

	ginkgo.It("should not exists opentracing directive", func() {
		config := map[string]string{}
		config[enableOpentracing] = "false"
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return !strings.Contains(cfg, "opentracing on")
			})
	})

	ginkgo.It("should exists opentracing directive when is enabled", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, "opentracing on")
			})
	})

	ginkgo.It("should include opentracing_trust_incoming_span off directive when disabled", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[opentracingTrustIncomingSpan] = "false"
		config[zipkinCollectorHost] = "127.0.0.1"
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, "opentracing_trust_incoming_span off")
			})
	})

	ginkgo.It("should not exists opentracing_operation_name directive when is empty", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		config[opentracingOperationName] = ""
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return !strings.Contains(cfg, "opentracing_operation_name")
			})
	})

	ginkgo.It("should exists opentracing_operation_name directive when is configured", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		config[opentracingOperationName] = "HTTP $request_method $uri"
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, `opentracing_operation_name "HTTP $request_method $uri"`)
			})
	})

	ginkgo.It("should not exists opentracing_location_operation_name directive when is empty", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		config[opentracingLocationOperationName] = ""
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return !strings.Contains(cfg, "opentracing_location_operation_name")
			})
	})

	ginkgo.It("should exists opentracing_location_operation_name directive when is configured", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		config[opentracingLocationOperationName] = "HTTP $request_method $uri"
		f.SetNginxConfigMapData(config)

		f.EnsureIngress(framework.NewSingleIngress(enableOpentracing, "/", enableOpentracing, f.Namespace, "http-svc", 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, "opentracing_location_operation_name \"HTTP $request_method $uri\"")
			})
	})

	ginkgo.It("should enable opentracing using zipkin", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[zipkinCollectorHost] = "127.0.0.1"
		f.SetNginxConfigMapData(config)

		framework.Sleep(10 * time.Second)
		log, err := f.NginxLogs()
		assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
		assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")
	})

	ginkgo.It("should enable opentracing using jaeger", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[jaegerCollectorHost] = "127.0.0.1"
		f.SetNginxConfigMapData(config)

		framework.Sleep(10 * time.Second)
		log, err := f.NginxLogs()
		assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
		assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")
	})

	ginkgo.It("should enable opentracing using jaeger with sampler host", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[jaegerCollectorHost] = "127.0.0.1"
		config[jaegerSamplerHost] = "127.0.0.1"
		f.SetNginxConfigMapData(config)

		framework.Sleep(10 * time.Second)
		log, err := f.NginxLogs()
		assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
		assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")
	})

	ginkgo.It("should propagate the w3c header when configured with jaeger", func() {
		host := "jaeger-w3c"
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[jaegerCollectorHost] = "127.0.0.1"
		config[jaegerPropagationFormat] = "w3c"
		f.SetNginxConfigMapData(config)

		framework.Sleep(10 * time.Second)
		log, err := f.NginxLogs()
		assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
		assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, nil)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf("server_name %s ;", host))
			})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK).
			Body().
			Match("traceparent=[0-9a-f]{2}-[0-9a-f]{32}-[0-9a-f]{16}-[0-9a-f]{2}")
	})

	/*
		ginkgo.It("should enable opentracing using jaeger with an HTTP endpoint", func() {
			config := map[string]string{}
			config[enableOpentracing] = "true"
			config[jaegerEndpoint] = "http://127.0.0.1/api/traces"
			f.SetNginxConfigMapData(config)

			framework.Sleep(10 * time.Second)
			log, err := f.NginxLogs()
			assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
			assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")
		})
	*/

	ginkgo.It("should enable opentracing using datadog", func() {
		config := map[string]string{}
		config[enableOpentracing] = "true"
		config[datadogCollectorHost] = "http://127.0.0.1"
		f.SetNginxConfigMapData(config)

		framework.Sleep(10 * time.Second)
		log, err := f.NginxLogs()
		assert.Nil(ginkgo.GinkgoT(), err, "obtaining nginx logs")
		assert.NotContains(ginkgo.GinkgoT(), log, "Unexpected failure reloading the backend", "reloading nginx after a configmap change")
	})
})
