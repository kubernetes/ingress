/*
Copyright 2015 The Kubernetes Authors.

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

package authtls

import (
	"strings"

	"github.com/pkg/errors"
	networking "k8s.io/api/networking/v1beta1"

	"regexp"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	ing_errors "k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/internal/k8s"
)

const (
	defaultAuthTLSDepth     = 1
	defaultAuthVerifyClient = "on"
	defaultOCSP             = "off"
)

var (
	authVerifyClientRegex = regexp.MustCompile(`on|off|optional|optional_no_ca`)
	authOCSPRegex         = regexp.MustCompile(`on|off|leaf`)
)

// Config contains the AuthSSLCert used for mutual authentication
// and the configured ValidationDepth
type Config struct {
	resolver.AuthSSLCert
	VerifyClient       string `json:"verify_client"`
	ValidationDepth    int    `json:"validationDepth"`
	ErrorPage          string `json:"errorPage"`
	PassCertToUpstream bool   `json:"passCertToUpstream"`
	OCSP               string `json:"ocsp"`
	AuthTLSError       string
}

// Equal tests for equality between two Config types
func (assl1 *Config) Equal(assl2 *Config) bool {
	if assl1 == assl2 {
		return true
	}
	if assl1 == nil || assl2 == nil {
		return false
	}
	if !(&assl1.AuthSSLCert).Equal(&assl2.AuthSSLCert) {
		return false
	}
	if assl1.VerifyClient != assl2.VerifyClient {
		return false
	}
	if assl1.ValidationDepth != assl2.ValidationDepth {
		return false
	}
	if assl1.ErrorPage != assl2.ErrorPage {
		return false
	}
	if assl1.PassCertToUpstream != assl2.PassCertToUpstream {
		return false
	}
	if assl1.OCSP != assl2.OCSP {
		return false
	}

	return true
}

// NewParser creates a new TLS authentication annotation parser
func NewParser(resolver resolver.Resolver) parser.IngressAnnotation {
	return authTLS{resolver}
}

type authTLS struct {
	r resolver.Resolver
}

// Parse parses the annotations contained in the ingress
// rule used to use a Certificate as authentication method
func (a authTLS) Parse(ing *networking.Ingress) (interface{}, error) {
	var err error
	config := &Config{}

	tlsauthsecret, err := parser.GetStringAnnotation("auth-tls-secret", ing)
	if err != nil {
		return &Config{}, err
	}

	_, _, err = k8s.ParseNameNS(tlsauthsecret)
	if err != nil {
		return &Config{}, ing_errors.NewLocationDenied(err.Error())
	}

	authCert, err := a.r.GetAuthCertificate(tlsauthsecret)
	if err != nil {
		e := errors.Wrap(err, "error obtaining certificate")
		return &Config{}, ing_errors.LocationDenied{Reason: e}
	}
	config.AuthSSLCert = *authCert

	config.VerifyClient, err = parser.GetStringAnnotation("auth-tls-verify-client", ing)
	if err != nil || !authVerifyClientRegex.MatchString(config.VerifyClient) {
		config.VerifyClient = defaultAuthVerifyClient
	}

	config.ValidationDepth, err = parser.GetIntAnnotation("auth-tls-verify-depth", ing)
	if err != nil || config.ValidationDepth == 0 {
		config.ValidationDepth = defaultAuthTLSDepth
	}

	config.ErrorPage, err = parser.GetStringAnnotation("auth-tls-error-page", ing)
	if err != nil {
		config.ErrorPage = ""
	}

	config.PassCertToUpstream, err = parser.GetBoolAnnotation("auth-tls-pass-certificate-to-upstream", ing)
	if err != nil {
		config.PassCertToUpstream = false
	}

	config.OCSP, err = parser.GetStringAnnotation("auth-tls-ocsp", ing)
	if err != nil || !authOCSPRegex.MatchString(config.OCSP) {
		config.OCSP = defaultOCSP
	}

	// OCSP requires auth-tls-verify-client to be set
	if strings.EqualFold(config.OCSP, "on") && !strings.EqualFold(config.VerifyClient, "on") {
		return nil, ing_errors.NewInvalidAnnotationConfiguration("auth-tls-ocsp", "requires auth-tls-verify-client to be set on")
	}

	return config, nil
}
