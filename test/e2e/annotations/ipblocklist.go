/*
Copyright 2021 The Kubernetes Authors.

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
	"bytes"
	"context"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"os/exec"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

const (
	_ngxConf = `#
events {
	worker_connections  1024;
	multi_accept on;
}

http {
	default_type 'text/plain';
	client_max_body_size 0;

	server {
		access_log /dev/stdout;

		listen 80;

		location / {
			proxy_set_header Connection "";
			proxy_set_header Host $http_host;
			proxy_http_version 1.1;
            proxy_pass http://nginx-ingress-controller:80;
		}
	}
}
`
)

func getTestClient(host string) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://" + host,
		Client:  http.DefaultClient,
	})
}

var _ = framework.DescribeAnnotation("blocklist-source-range", func() {
	f := framework.NewDefaultFramework("ipblocklist")

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})
	ginkgo.AfterEach(func() {
		// This is for debugging.
		var (
			execOut bytes.Buffer
			execErr bytes.Buffer
		)

		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("%v logs deployments/reverse-proxy --namespace %s", framework.KubectlPath, f.Namespace))
		cmd.Stdout = &execOut
		cmd.Stderr = &execErr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("could not execute '%s %s': %v\n", cmd.Path, cmd.Args, err)
			return
		}

		eout := strings.TrimSpace(execErr.String())
		if len(eout) > 0 {
			fmt.Printf("stderr: %v\n", eout)
			return
		}
		fmt.Println("stdout", execOut.String())
	})

	ginkgo.It("should set valid ip blocklist range", func() {
		f.NGINXWithConfigDeployment("reverse-proxy", _ngxConf, 1)

		e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), "reverse-proxy", metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets), 1, "expected at least one endpoint")
		assert.Equal(ginkgo.GinkgoT(), len(e.Subsets[0].Addresses), 1, "expected three address ready in the endpoint")

		host := "ipblocklist.foo.com"
		nameSpace := f.Namespace

		//allowed := e.Subsets[0].Addresses[0]
		//denied := e.Subsets[0].Addresses[1]
		//general := e.Subsets[0].Addresses[2]
		denied := e.Subsets[0].Addresses[0]

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/blocklist-source-range": denied.IP,
		}

		ing := framework.NewSingleIngress(host, "/", host, nameSpace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf("deny %s;", denied.IP)) &&
					strings.Contains(server, "allow all;")
			},
		)

		//getTestClient(allowed.String()).GET("/").WithHeader("Host", host).Expect().Status(200)
		getTestClient(denied.String()).GET("/").WithHeader("Host", host).Expect().Status(403)
		//getTestClient(general.String()).GET("/").WithHeader("Host", host).Expect().Status(200)
	})

	ginkgo.It("ignore ip blocklist range when whitelist range is set", func() {
		f.NGINXWithConfigDeployment("reverse-proxy", _ngxConf, 3)

		e, err := f.KubeClientSet.CoreV1().Endpoints(f.Namespace).Get(context.TODO(), "reverse-proxy", metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		assert.GreaterOrEqual(ginkgo.GinkgoT(), len(e.Subsets), 1, "expected at least one endpoint")
		assert.Equal(ginkgo.GinkgoT(), len(e.Subsets[0].Addresses), 3, "expected three address ready in the endpoint")

		allowed := e.Subsets[0].Addresses[0]
		denied := e.Subsets[0].Addresses[1]
		general := e.Subsets[0].Addresses[2]

		host := "ipblocklist.foo.com"
		nameSpace := f.Namespace

		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/blocklist-source-range": denied.IP,
			"nginx.ingress.kubernetes.io/whitelist-source-range": allowed.IP,
		}

		ing := framework.NewSingleIngress(host, "/", host, nameSpace, framework.EchoService, 80, annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf("allow %s;", allowed.IP)) &&
					strings.Contains(server, "deny all;")
			})
		getTestClient(allowed.String()).GET("/").WithHeader("Host", host).Expect().Status(200)
		getTestClient(denied.String()).GET("/").WithHeader("Host", host).Expect().Status(403)
		getTestClient(general.String()).GET("/").WithHeader("Host", host).Expect().Status(403)
	})
})
