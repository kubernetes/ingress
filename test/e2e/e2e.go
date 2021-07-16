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

package e2e

import (
	"os"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"k8s.io/component-base/logs"

	// required
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/ingress-nginx/test/e2e/framework"

	// tests to run
	_ "k8s.io/ingress-nginx/test/e2e/admission"
	_ "k8s.io/ingress-nginx/test/e2e/annotations"
	_ "k8s.io/ingress-nginx/test/e2e/dbg"
	_ "k8s.io/ingress-nginx/test/e2e/defaultbackend"
	_ "k8s.io/ingress-nginx/test/e2e/gracefulshutdown"
	_ "k8s.io/ingress-nginx/test/e2e/ingress"
	_ "k8s.io/ingress-nginx/test/e2e/leaks"
	_ "k8s.io/ingress-nginx/test/e2e/loadbalance"
	_ "k8s.io/ingress-nginx/test/e2e/lua"
	_ "k8s.io/ingress-nginx/test/e2e/security"
	_ "k8s.io/ingress-nginx/test/e2e/servicebackend"
	_ "k8s.io/ingress-nginx/test/e2e/settings"
	_ "k8s.io/ingress-nginx/test/e2e/ssl"
	_ "k8s.io/ingress-nginx/test/e2e/status"
	_ "k8s.io/ingress-nginx/test/e2e/tcpudp"
)

// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
func RunE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	if os.Getenv("KUBECTL_PATH") != "" {
		framework.KubectlPath = os.Getenv("KUBECTL_PATH")
		framework.Logf("Using kubectl path '%s'", framework.KubectlPath)
	}

	framework.Logf("Starting e2e run %q on Ginkgo node %d", framework.RunID, config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecs(t, "nginx-ingress-controller e2e suite")
}
