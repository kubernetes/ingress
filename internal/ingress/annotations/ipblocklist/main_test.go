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

package ipblocklist

import (
	"k8s.io/ingress-nginx/internal/ingress/defaults"
	"reflect"
	"testing"

	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
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

func TestParseAnnotations(t *testing.T) {
	ing := buildIngress()
	tests := map[string]struct {
		net        string
		expectCidr []string
		expectErr  bool
		errOut     string
	}{
		"test parse a valid net": {
			net:        "10.0.0.0/24",
			expectCidr: []string{"10.0.0.0/24"},
			expectErr:  false,
		},
		"test parse a invalid net": {
			net:       "ww",
			expectErr: true,
			errOut:    "the annotation does not contain a valid IP address or network: invalid CIDR address: ww",
		},
		"test parse a empty net": {
			net:       "",
			expectErr: true,
			errOut:    "the annotation does not contain a valid IP address or network: invalid CIDR address: ",
		},
		"test parse multiple valid cidr": {
			net:        "2.2.2.2/32,1.1.1.1/32,3.3.3.0/24",
			expectCidr: []string{"1.1.1.1/32", "2.2.2.2/32", "3.3.3.0/24"},
			expectErr:  false,
		},
	}
	for testName, test := range tests {
		data := map[string]string{
			parser.GetAnnotationWithPrefix("blocklist-source-range"): test.net,
		}
		ing.SetAnnotations(data)
		p := NewParser(&resolver.Mock{})
		i, err := p.Parse(ing)
		if err != nil && !test.expectErr {
			t.Errorf("%v:unexpected error: %v", testName, err)
		}

		if !test.expectErr {
			sr, ok := i.(*SourceRange)
			if !ok {
				t.Errorf("%v:expected a SourceRange type", testName)
			}
			if !reflect.DeepEqual(sr.CIDR, test.expectCidr) {
				t.Errorf("%v:expected %v CIDR but return %v returned", testName, test.expectCidr, sr.CIDR)
			}
		}
	}
}

type mockBackend struct {
	resolver.Mock
}

// GetDefaultBackend returns the backend that must be used as default
func (m mockBackend) GetDefaultBackend() defaults.Backend {
	return defaults.Backend{
		BlocklistSourceRange: []string{"4.4.4.0/24", "1.2.3.4/32"},
	}
}

// Test that when we have a block set on the Backend that is used when we
// don't have the annotation
func TestParseAnnotationsWithDefaultConfig(t *testing.T) {
	ing := buildIngress()

	mockBackend := mockBackend{}

	tests := map[string]struct {
		net        string
		expectCidr []string
		expectErr  bool
		errOut     string
	}{
		"test parse a valid net": {
			net:        "10.0.0.0/24",
			expectCidr: []string{"10.0.0.0/24"},
			expectErr:  false,
		},
		"test parse a invalid net": {
			net:       "ww",
			expectErr: true,
			errOut:    "the annotation does not contain a valid IP address or network: invalid CIDR address: ww",
		},
		"test parse a empty net": {
			net:       "",
			expectErr: true,
			errOut:    "the annotation does not contain a valid IP address or network: invalid CIDR address: ",
		},
		"test parse multiple valid cidr": {
			net:        "2.2.2.2/32,1.1.1.1/32,3.3.3.0/24",
			expectCidr: []string{"1.1.1.1/32", "2.2.2.2/32", "3.3.3.0/24"},
			expectErr:  false,
		},
	}

	for testName, test := range tests {
		data := map[string]string{
			parser.GetAnnotationWithPrefix("blocklist-source-range"): test.net,
		}
		ing.SetAnnotations(data)
		p := NewParser(mockBackend)
		i, err := p.Parse(ing)
		if err != nil && !test.expectErr {
			t.Errorf("%v:unexpected error: %v", testName, err)
		}
		if test.expectErr {
			if err.Error() != test.errOut {
				t.Errorf("%v:expected error: %v but %v return", testName, test.errOut, err.Error())
			}
		}
		if !test.expectErr {
			sr, ok := i.(*SourceRange)
			if !ok {
				t.Errorf("%v:expected a SourceRange type", testName)
			}
			if !reflect.DeepEqual(sr.CIDR, test.expectCidr) {
				t.Errorf("%v:expected %v CIDR but %v returned", testName, test.expectCidr, sr.CIDR)
			}
		}
	}
}
