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
	"testing"

	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		ServiceName: "default-backend",
		ServicePort: intstr.FromInt(80),
	}

	return &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{
			Backend: &networking.IngressBackend{
				ServiceName: "default-backend",
				ServicePort: intstr.FromInt(80),
			},
			Rules: []networking.IngressRule{
				{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:    "/foo",
									Backend: defaultBackend,
								},
							},
						},
					},
				},
			},
		},
	}
}

// mocks the resolver for authTLS
type mockSecret struct {
	resolver.Mock
}

// GetAuthCertificate from mockSecret mocks the GetAuthCertificate for authTLS
func (m mockSecret) GetAuthCertificate(name string) (*resolver.AuthSSLCert, error) {
	if name != "default/demo-secret" {
		return nil, errors.Errorf("there is no secret with name %v", name)
	}

	return &resolver.AuthSSLCert{
		Secret:     "default/demo-secret",
		CAFileName: "/ssl/ca.crt",
		CASHA:      "abc",
	}, nil

}

func TestAnnotations(t *testing.T) {
	ing := buildIngress()
	data := map[string]string{}

	data[parser.GetAnnotationWithPrefix("auth-tls-secret")] = "default/demo-secret"
	data[parser.GetAnnotationWithPrefix("auth-tls-verify-client")] = "off"
	data[parser.GetAnnotationWithPrefix("auth-tls-verify-depth")] = "1"
	data[parser.GetAnnotationWithPrefix("auth-tls-error-page")] = "ok.com/error"
	data[parser.GetAnnotationWithPrefix("auth-tls-pass-certificate-to-upstream")] = "true"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp")] = "off"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp-responder")] = "http://alwaysok.com"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp-cache")] = "shared:foo:10m"

	ing.SetAnnotations(data)

	fakeSecret := &mockSecret{}
	i, err := NewParser(fakeSecret).Parse(ing)
	if err != nil {
		t.Errorf("Unexpected error with ingress: %v", err)
	}

	u, ok := i.(*Config)
	if !ok {
		t.Errorf("expected *Config but got %v", u)
	}

	secret, err := fakeSecret.GetAuthCertificate("default/demo-secret")
	if err != nil {
		t.Errorf("unexpected error getting secret %v", err)
	}

	if u.AuthSSLCert.Secret != secret.Secret {
		t.Errorf("expected %v but got %v", secret.Secret, u.AuthSSLCert.Secret)
	}
	if u.VerifyClient != "off" {
		t.Errorf("expected %v but got %v", "off", u.VerifyClient)
	}
	if u.ValidationDepth != 1 {
		t.Errorf("expected %v but got %v", 1, u.ValidationDepth)
	}
	if u.ErrorPage != "ok.com/error" {
		t.Errorf("expected %v but got %v", "ok.com/error", u.ErrorPage)
	}
	if u.PassCertToUpstream != true {
		t.Errorf("expected %v but got %v", true, u.PassCertToUpstream)
	}
	if u.OCSP != "off" {
		t.Errorf("expected %v but got %v", "off", u.OCSP)
	}
	if u.OCSPResponder != "http://alwaysok.com" {
		t.Errorf("expected %v but got %v", "http://alwaysok.com", u.OCSPResponder)
	}
	if u.OCSPCache != "shared:foo:10m" {
		t.Errorf("expected %v but got %v", "shared:foo:10m", u.OCSPCache)
	}
}

func TestInvalidAnnotations(t *testing.T) {
	ing := buildIngress()
	fakeSecret := &mockSecret{}
	data := map[string]string{}

	// No annotation
	_, err := NewParser(fakeSecret).Parse(ing)
	if err == nil {
		t.Errorf("Expected error with ingress but got nil")
	}

	// Invalid NameSpace
	data[parser.GetAnnotationWithPrefix("auth-tls-secret")] = "demo-secret"
	ing.SetAnnotations(data)
	_, err = NewParser(fakeSecret).Parse(ing)
	if err == nil {
		t.Errorf("Expected error with ingress but got nil")
	}

	// Invalid Auth Certificate
	data[parser.GetAnnotationWithPrefix("auth-tls-secret")] = "default/invalid-demo-secret"
	ing.SetAnnotations(data)
	_, err = NewParser(fakeSecret).Parse(ing)
	if err == nil {
		t.Errorf("Expected error with ingress but got nil")
	}

	// Enabling OCSP without client verification
	data[parser.GetAnnotationWithPrefix("auth-tls-secret")] = "default/demo-secret"
	data[parser.GetAnnotationWithPrefix("auth-tls-verify-client")] = "off"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp")] = "on"
	ing.SetAnnotations(data)

	_, err = NewParser(fakeSecret).Parse(ing)
	if err == nil || !strings.EqualFold(err.Error(), "the annotation auth-tls-ocsp does not contain a valid configuration: requires auth-tls-verify-client to be set on") {
		t.Errorf("Expected a auth-tls-ocsp error but got: %v", err)
	}

	// Invalid optional Annotations
	data[parser.GetAnnotationWithPrefix("auth-tls-secret")] = "default/demo-secret"
	data[parser.GetAnnotationWithPrefix("auth-tls-verify-client")] = "w00t"
	data[parser.GetAnnotationWithPrefix("auth-tls-verify-depth")] = "abcd"
	data[parser.GetAnnotationWithPrefix("auth-tls-pass-certificate-to-upstream")] = "nahh"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp")] = "1337"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp-responder")] = "https://totes.not.ok"
	data[parser.GetAnnotationWithPrefix("auth-tls-ocsp-cache")] = "should:totes:fail"
	ing.SetAnnotations(data)

	i, err := NewParser(fakeSecret).Parse(ing)
	if err != nil {
		t.Errorf("Unexpected error with ingress: %v", err)
	}
	u, ok := i.(*Config)
	if !ok {
		t.Errorf("expected *Config but got %v", u)
	}

	if u.VerifyClient != "on" {
		t.Errorf("expected %v but got %v", "on", u.VerifyClient)
	}
	if u.ValidationDepth != 1 {
		t.Errorf("expected %v but got %v", 1, u.ValidationDepth)
	}
	if u.PassCertToUpstream != false {
		t.Errorf("expected %v but got %v", false, u.PassCertToUpstream)
	}
	if u.OCSP != "off" {
		t.Errorf("expected %v but got %v", "off", u.OCSP)
	}
	if u.OCSPResponder != "" {
		t.Errorf("expected an empty string but got %v", u.OCSPResponder)
	}
	if u.OCSPCache != "off" {
		t.Errorf("execpted %v but got %v", "off", u.OCSPCache)
	}
}

func TestEquals(t *testing.T) {
	cfg1 := &Config{}
	cfg2 := &Config{}

	// Same config
	result := cfg1.Equal(cfg1)
	if result != true {
		t.Errorf("Expected true")
	}

	// compare nil
	result = cfg1.Equal(nil)
	if result != false {
		t.Errorf("Expected false")
	}

	// Different Certs
	sslCert1 := resolver.AuthSSLCert{
		Secret:     "default/demo-secret",
		CAFileName: "/ssl/ca.crt",
		CASHA:      "abc",
	}
	sslCert2 := resolver.AuthSSLCert{
		Secret:     "default/other-demo-secret",
		CAFileName: "/ssl/ca.crt",
		CASHA:      "abc",
	}
	cfg1.AuthSSLCert = sslCert1
	cfg2.AuthSSLCert = sslCert2
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.AuthSSLCert = sslCert1

	// Different Verify Client
	cfg1.VerifyClient = "on"
	cfg2.VerifyClient = "off"
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.VerifyClient = "on"

	// Different Validation Depth
	cfg1.ValidationDepth = 1
	cfg2.ValidationDepth = 2
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.ValidationDepth = 1

	// Different Error Page
	cfg1.ErrorPage = "error-1"
	cfg2.ErrorPage = "error-2"
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.ErrorPage = "error-1"

	// Different Pass to Upstream
	cfg1.PassCertToUpstream = true
	cfg2.PassCertToUpstream = false
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.PassCertToUpstream = true

	// Different OCSP
	cfg1.OCSP = "leaf"
	cfg2.OCSP = "on"
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.OCSP = "leaf"

	// Different OCSP Cache
	cfg1.OCSPCache = "shared:this:10m"
	cfg2.OCSPCache = "shared:that:5m"
	result = cfg1.Equal(cfg2)
	if result != false {
		t.Errorf("Expected false")
	}
	cfg2.OCSPCache = "shared:this:10m"

	// Equal Configs
	result = cfg1.Equal(cfg2)
	if result != true {
		t.Errorf("Expected true")
	}
}
