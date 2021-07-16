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
	"testing"

	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		Service: &networking.IngressServiceBackend{
			Name: "default-backend",
			Port: networking.ServiceBackendPort{
				Number: 80,
			},
		},
	}

	return &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &networking.IngressBackend{
				Service: &networking.IngressServiceBackend{
					Name: "default-backend",
					Port: networking.ServiceBackendPort{
						Number: 80,
					},
				},
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

func TestAnnotation(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix("enable-global-tls-auth")] = "true"
	ing.SetAnnotations(data)

	i, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("Unexpected error with ingress: %v", err)
	}
	globalTLSConfig, ok := i.(*GlobalTLSConfig)
	if !ok {
		t.Errorf("Expected a GlobalTLSConfig type")
	}

	if globalTLSConfig.AuthTLSConfig.VerifyClient != "on" {
		t.Errorf("expected %v but got %v", "on", globalTLSConfig.AuthTLSConfig.VerifyClient)
	}

	if globalTLSConfig.AuthTLSConfig.ValidationDepth != 1 {
		t.Errorf("expected %v but got %v", 1, globalTLSConfig.AuthTLSConfig.ValidationDepth)
	}

	if globalTLSConfig.AuthTLSConfig.CAFileName != "/ssl/ca.crt" {
		t.Errorf("expected %v but got %v", "/ssl/ca.crt", globalTLSConfig.AuthTLSConfig.CAFileName)
	}
}

type mockResolver struct {
	resolver.Mock
}

func (m mockResolver) GetGlobalTLSAuth() resolver.GlobalTLSAuth {
	return resolver.GlobalTLSAuth{
		AuthTLSVerifyClient: "on",
		AuthTLSVerifyDepth:  1,
		AuthTLSSecret:       "default/demo-secret",
	}
}

// GetAuthCertificate from mockSecret mocks the GetAuthCertificate for authTLS
func (m mockResolver) GetAuthCertificate(name string) (*resolver.AuthSSLCert, error) {
	if name != "default/demo-secret" {
		return nil, errors.Errorf("there is no secret with name %v", name)
	}

	return &resolver.AuthSSLCert{
		Secret:     "default/demo-secret",
		CAFileName: "/ssl/ca.crt",
		CASHA:      "abc",
	}, nil
}
