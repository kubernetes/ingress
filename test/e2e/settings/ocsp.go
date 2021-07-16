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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"strings"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ocsp"

	"k8s.io/ingress-nginx/test/e2e/framework"
	ocspframework "k8s.io/ingress-nginx/test/e2e/framework/ocsp"
)

var _ = framework.DescribeSetting("OCSP", func() {
	f := framework.NewDefaultFramework("ocsp")
	o := ocspframework.NewFramework(f)

	ginkgo.BeforeEach(func() {
		f.NewEchoDeployment()
	})

	ginkgo.It("should enable OCSP and contain stapling information in the connection", func() {
		host := "www.example.com"

		f.UpdateNginxConfigMapData("enable-ocsp", "true")

		ing := framework.NewSingleIngressWithTLS(host, "/", host, []string{host}, f.Namespace, framework.EchoService, 80, nil)
		f.EnsureIngress(ing)

		err := o.CreateIngressOcspSecret(
			host,
			host,
			f.Namespace)
		assert.Nil(ginkgo.GinkgoT(), err)

		err = o.EnsureOCSPResponderDeployment(f.Namespace, "ocspserve")
		assert.NoError(ginkgo.GinkgoT(), err)

		err = framework.WaitForEndpoints(f.KubeClientSet, framework.DefaultTimeout, "ocspserve", f.Namespace, 1)
		assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")

		f.WaitForNginxConfiguration(func(cfg string) bool {
			return strings.Contains(cfg, "certificate.is_ocsp_stapling_enabled = true")
		})

		f.WaitForNginxServer(host,
			func(server string) bool {
				return strings.Contains(server, fmt.Sprintf(`server_name %v`, host))
			})

		tlsConfig := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		f.HTTPTestClientWithTLSConfig(tlsConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK).
			Raw()

		// give time the lua request to the OCSP
		// URL to finish and update the cache
		framework.Sleep()

		// TODO: is possible to avoid second request?
		resp := f.HTTPTestClientWithTLSConfig(tlsConfig).
			GET("/").
			WithURL(f.GetURL(framework.HTTPS)).
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK).
			Raw()

		state := resp.TLS
		assert.NotNil(ginkgo.GinkgoT(), state.OCSPResponse, "unexpected connection without OCSP response")

		var issuerCertificate *x509.Certificate
		var leafAuthorityKeyID string
		for index, certificate := range state.PeerCertificates {
			if index == 0 {
				leafAuthorityKeyID = string(certificate.AuthorityKeyId)
				continue
			}

			if leafAuthorityKeyID == string(certificate.SubjectKeyId) {
				issuerCertificate = certificate
			}
		}

		response, err := ocsp.ParseResponse(state.OCSPResponse, issuerCertificate)
		assert.Nil(ginkgo.GinkgoT(), err)
		assert.Equal(ginkgo.GinkgoT(), ocsp.Good, response.Status)
	})
})
