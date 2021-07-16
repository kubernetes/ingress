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

package authtlsglobal

import (
	"regexp"

	"github.com/pkg/errors"
	networking "k8s.io/api/networking/v1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/authtls"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	ing_errors "k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/klog/v2"
)

const (
	defaultAuthTLSDepth     = 1
	defaultAuthVerifyClient = "on"
)

var (
	authVerifyClientRegex = regexp.MustCompile(`on|off|optional|optional_no_ca`)
)

type GlobalTLSConfig struct {
	EnableGlobalTLSAuth bool
	AuthTLSConfig       authtls.Config
}

// NewParser creates a new authentication request annotation parser
func NewParser(resolver resolver.Resolver) parser.IngressAnnotation {
	return authTLSGlobal{resolver}
}

type authTLSGlobal struct {
	r resolver.Resolver
}

// ParseAnnotations parses the annotations contained in the ingress
// rule used to enable or disable global external authentication
func (a authTLSGlobal) Parse(ing *networking.Ingress) (interface{}, error) {
	var err error
	globalConfig := &GlobalTLSConfig{false, authtls.Config{}}

	enableGlobalTLSAuth, err := parser.GetBoolAnnotation("enable-global-tls-auth", ing)
	if err != nil {
		return globalConfig, err
	}

	globalConfig.EnableGlobalTLSAuth = enableGlobalTLSAuth
	if !globalConfig.EnableGlobalTLSAuth {
		return globalConfig, nil
	}

	globalTLSAuth := a.r.GetGlobalTLSAuth()
	_, _, err = k8s.ParseNameNS(globalTLSAuth.AuthTLSSecret)
	if err != nil {
		return globalConfig, ing_errors.NewLocationDenied(err.Error())
	}

	authCert, err := a.r.GetAuthCertificate(globalTLSAuth.AuthTLSSecret)
	if err != nil {
		e := errors.Wrap(err, "error obtaining certificate")
		return globalConfig, ing_errors.LocationDenied{Reason: e}
	}
	globalConfig.AuthTLSConfig.AuthSSLCert = *authCert

	globalConfig.AuthTLSConfig.VerifyClient = globalTLSAuth.AuthTLSVerifyClient
	if !authVerifyClientRegex.MatchString(globalConfig.AuthTLSConfig.VerifyClient) {
		globalConfig.AuthTLSConfig.VerifyClient = defaultAuthVerifyClient
	}

	globalConfig.AuthTLSConfig.ValidationDepth = globalTLSAuth.AuthTLSVerifyDepth
	if globalConfig.AuthTLSConfig.ValidationDepth == 0 {
		globalConfig.AuthTLSConfig.ValidationDepth = defaultAuthTLSDepth
	}

	klog.InfoS("globalconfiganns", globalTLSAuth.AuthTLSSecret, authCert.CAFileName)
	return globalConfig, nil
}
