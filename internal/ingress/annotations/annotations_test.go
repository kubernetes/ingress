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
	"testing"

	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/defaults"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

var (
	annotationPassthrough            = parser.GetAnnotationWithPrefix("ssl-passthrough")
	annotationAffinityType           = parser.GetAnnotationWithPrefix("affinity")
	annotationAffinityMode           = parser.GetAnnotationWithPrefix("affinity-mode")
	annotationAffinityCanaryBehavior = parser.GetAnnotationWithPrefix("affinity-canary-behavior")
	annotationCorsEnabled            = parser.GetAnnotationWithPrefix("enable-cors")
	annotationCorsAllowMethods       = parser.GetAnnotationWithPrefix("cors-allow-methods")
	annotationCorsAllowHeaders       = parser.GetAnnotationWithPrefix("cors-allow-headers")
	annotationCorsExposeHeaders      = parser.GetAnnotationWithPrefix("cors-expose-headers")
	annotationCorsAllowCredentials   = parser.GetAnnotationWithPrefix("cors-allow-credentials")
	defaultCorsMethods               = "GET, PUT, POST, DELETE, PATCH, OPTIONS"
	defaultCorsHeaders               = "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization"
	annotationAffinityCookieName     = parser.GetAnnotationWithPrefix("session-cookie-name")
	annotationUpstreamHashBy         = parser.GetAnnotationWithPrefix("upstream-hash-by")
	annotationCustomHTTPErrors       = parser.GetAnnotationWithPrefix("custom-http-errors")
)

type mockCfg struct {
	resolver.Mock
	MockSecrets  map[string]*apiv1.Secret
	MockServices map[string]*apiv1.Service
}

func (m mockCfg) GetDefaultBackend() defaults.Backend {
	return defaults.Backend{}
}

func (m mockCfg) GetSecret(name string) (*apiv1.Secret, error) {
	return m.MockSecrets[name], nil
}

func (m mockCfg) GetService(name string) (*apiv1.Service, error) {
	return m.MockServices[name], nil
}

func (m mockCfg) GetAuthCertificate(name string) (*resolver.AuthSSLCert, error) {
	if secret, _ := m.GetSecret(name); secret != nil {
		return &resolver.AuthSSLCert{
			Secret:     name,
			CAFileName: "/opt/ca.pem",
			CASHA:      "123",
		}, nil
	}
	return nil, nil
}

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		ServiceName: "default-backend",
		ServicePort: intstr.FromInt(80),
	}

	return &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: apiv1.NamespaceDefault,
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

func TestSSLPassthrough(t *testing.T) {
	ec := NewAnnotationExtractor(mockCfg{})
	ing := buildIngress()

	fooAnns := []struct {
		annotations map[string]string
		er          bool
	}{
		{map[string]string{annotationPassthrough: "true"}, true},
		{map[string]string{annotationPassthrough: "false"}, false},
		{map[string]string{annotationPassthrough + "_no": "true"}, false},
		{map[string]string{}, false},
		{nil, false},
	}

	for _, foo := range fooAnns {
		ing.SetAnnotations(foo.annotations)
		r := ec.Extract(ing).SSLPassthrough
		if r != foo.er {
			t.Errorf("Returned %v but expected %v", r, foo.er)
		}
	}
}

func TestUpstreamHashBy(t *testing.T) {
	ec := NewAnnotationExtractor(mockCfg{})
	ing := buildIngress()

	fooAnns := []struct {
		annotations map[string]string
		er          string
	}{
		{map[string]string{annotationUpstreamHashBy: "$request_uri"}, "$request_uri"},
		{map[string]string{annotationUpstreamHashBy: "false"}, "false"},
		{map[string]string{annotationUpstreamHashBy + "_no": "true"}, ""},
		{map[string]string{}, ""},
		{nil, ""},
	}

	for _, foo := range fooAnns {
		ing.SetAnnotations(foo.annotations)
		r := ec.Extract(ing).UpstreamHashBy.UpstreamHashBy
		if r != foo.er {
			t.Errorf("Returned %v but expected %v", r, foo.er)
		}
	}
}

func TestAffinitySession(t *testing.T) {
	ec := NewAnnotationExtractor(mockCfg{})
	ing := buildIngress()

	fooAnns := []struct {
		annotations    map[string]string
		affinitytype   string
		affinitymode   string
		cookiename     string
		canarybehavior string
	}{
		{map[string]string{annotationAffinityType: "cookie", annotationAffinityMode: "balanced", annotationAffinityCookieName: "route", annotationAffinityCanaryBehavior: ""}, "cookie", "balanced", "route", ""},
		{map[string]string{annotationAffinityType: "cookie", annotationAffinityMode: "persistent", annotationAffinityCookieName: "route1", annotationAffinityCanaryBehavior: "sticky"}, "cookie", "persistent", "route1", "sticky"},
		{map[string]string{annotationAffinityType: "cookie", annotationAffinityMode: "balanced", annotationAffinityCookieName: "", annotationAffinityCanaryBehavior: "legacy"}, "cookie", "balanced", "INGRESSCOOKIE", "legacy"},
		{map[string]string{}, "", "", "", ""},
		{nil, "", "", "", ""},
	}

	for _, foo := range fooAnns {
		ing.SetAnnotations(foo.annotations)
		r := ec.Extract(ing).SessionAffinity
		t.Logf("Testing pass %v %v", foo.affinitytype, foo.cookiename)

		if r.Type != foo.affinitytype {
			t.Errorf("Returned %v but expected %v for Type", r.Type, foo.affinitytype)
		}

		if r.Mode != foo.affinitymode {
			t.Errorf("Returned %v but expected %v for Mode", r.Mode, foo.affinitymode)
		}

		if r.CanaryBehavior != foo.canarybehavior {
			t.Errorf("Returned %v but expected %v for CanaryBehavior", r.CanaryBehavior, foo.canarybehavior)
		}

		if r.Cookie.Name != foo.cookiename {
			t.Errorf("Returned %v but expected %v for Cookie.Name", r.Cookie.Name, foo.cookiename)
		}
	}
}

func TestCors(t *testing.T) {
	ec := NewAnnotationExtractor(mockCfg{})
	ing := buildIngress()

	fooAnns := []struct {
		annotations map[string]string
		corsenabled bool
		methods     string
		headers     string
		origin      string
		credentials bool
		expose      string
	}{
		{map[string]string{annotationCorsEnabled: "true"}, true, defaultCorsMethods, defaultCorsHeaders, "*", true, ""},
		{map[string]string{annotationCorsEnabled: "true", annotationCorsAllowMethods: "POST, GET, OPTIONS", annotationCorsAllowHeaders: "$nginx_version", annotationCorsAllowCredentials: "false", annotationCorsExposeHeaders: "X-CustomResponseHeader"}, true, "POST, GET, OPTIONS", defaultCorsHeaders, "*", false, "X-CustomResponseHeader"},
		{map[string]string{annotationCorsEnabled: "true", annotationCorsAllowCredentials: "false"}, true, defaultCorsMethods, defaultCorsHeaders, "*", false, ""},
		{map[string]string{}, false, defaultCorsMethods, defaultCorsHeaders, "*", true, ""},
		{nil, false, defaultCorsMethods, defaultCorsHeaders, "*", true, ""},
	}

	for _, foo := range fooAnns {
		ing.SetAnnotations(foo.annotations)
		r := ec.Extract(ing).CorsConfig
		t.Logf("Testing pass %v %v %v %v %v", foo.corsenabled, foo.methods, foo.headers, foo.origin, foo.credentials)

		if r.CorsEnabled != foo.corsenabled {
			t.Errorf("Returned %v but expected %v for Cors Enabled", r.CorsEnabled, foo.corsenabled)
		}

		if r.CorsAllowHeaders != foo.headers {
			t.Errorf("Returned %v but expected %v for Cors Headers", r.CorsAllowHeaders, foo.headers)
		}

		if r.CorsAllowMethods != foo.methods {
			t.Errorf("Returned %v but expected %v for Cors Methods", r.CorsAllowMethods, foo.methods)
		}

		if r.CorsAllowOrigin != foo.origin {
			t.Errorf("Returned %v but expected %v for Cors Methods", r.CorsAllowOrigin, foo.origin)
		}

		if r.CorsAllowCredentials != foo.credentials {
			t.Errorf("Returned %v but expected %v for Cors Methods", r.CorsAllowCredentials, foo.credentials)
		}

	}
}
func TestCustomHTTPErrors(t *testing.T) {
	ec := NewAnnotationExtractor(mockCfg{})
	ing := buildIngress()

	fooAnns := []struct {
		annotations map[string]string
		er          []int
	}{
		{map[string]string{annotationCustomHTTPErrors: "404,415"}, []int{404, 415}},
		{map[string]string{annotationCustomHTTPErrors: "404"}, []int{404}},
		{map[string]string{annotationCustomHTTPErrors: ""}, []int{}},
		{map[string]string{annotationCustomHTTPErrors + "_no": "404"}, []int{}},
		{map[string]string{}, []int{}},
		{nil, []int{}},
	}

	for _, foo := range fooAnns {
		ing.SetAnnotations(foo.annotations)
		r := ec.Extract(ing).CustomHTTPErrors

		// Check that expected codes were created
		for i := range foo.er {
			if r[i] != foo.er[i] {
				t.Errorf("Returned %v but expected %v", r, foo.er)
			}
		}

		// Check that no unexpected codes were also created
		for i := range r {
			if r[i] != foo.er[i] {
				t.Errorf("Returned %v but expected %v", r, foo.er)
			}
		}
	}
}

/*
func TestMergeLocationAnnotations(t *testing.T) {
	// initial parameters
	keys := []string{"BasicDigestAuth", "CorsConfig", "ExternalAuth", "RateLimit", "Redirect", "Rewrite", "Whitelist", "Proxy", "UsePortInRedirects"}

	loc := ingress.Location{}
	annotations := &Ingress{
		BasicDigestAuth:    &auth.Config{},
		CorsConfig:         &cors.Config{},
		ExternalAuth:       &authreq.Config{},
		RateLimit:          &ratelimit.Config{},
		Redirect:           &redirect.Config{},
		Rewrite:            &rewrite.Config{},
		Whitelist:          &ipwhitelist.SourceRange{},
		Proxy:              &proxy.Config{},
		UsePortInRedirects: true,
	}

	// create test table
	type fooMergeLocationAnnotationsStruct struct {
		fName string
		er    interface{}
	}
	fooTests := []fooMergeLocationAnnotationsStruct{}
	for name, value := range keys {
		fva := fooMergeLocationAnnotationsStruct{name, value}
		fooTests = append(fooTests, fva)
	}

	// execute test
	MergeWithLocation(&loc, annotations)

	// check result
	for _, foo := range fooTests {
		fv := reflect.ValueOf(loc).FieldByName(foo.fName).Interface()
		if !reflect.DeepEqual(fv, foo.er) {
			t.Errorf("Returned %v but expected %v for the field %s", fv, foo.er, foo.fName)
		}
	}
	if _, ok := annotations[DeniedKeyName]; ok {
		t.Errorf("%s should be removed after mergeLocationAnnotations", DeniedKeyName)
	}
}
*/
