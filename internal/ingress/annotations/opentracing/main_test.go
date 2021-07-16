/*
Copyright 2019 The Kubernetes Authors.

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

package opentracing

import (
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

func TestIngressAnnotationOpentracingSetTrue(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix("enable-opentracing")] = "true"
	ing.SetAnnotations(data)

	val, _ := NewParser(&resolver.Mock{}).Parse(ing)
	openTracing, ok := val.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if !openTracing.Enabled {
		t.Errorf("expected annotation value to be true, got false")
	}
}

func TestIngressAnnotationOpentracingSetFalse(t *testing.T) {
	ing := buildIngress()

	// Test with explicitly set to false
	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix("enable-opentracing")] = "false"
	ing.SetAnnotations(data)

	val, _ := NewParser(&resolver.Mock{}).Parse(ing)
	openTracing, ok := val.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if openTracing.Enabled {
		t.Errorf("expected annotation value to be false, got true")
	}
}

func TestIngressAnnotationOpentracingTrustSetTrue(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix("enable-opentracing")] = "true"
	data[parser.GetAnnotationWithPrefix("opentracing-trust-incoming-span")] = "true"
	ing.SetAnnotations(data)

	val, _ := NewParser(&resolver.Mock{}).Parse(ing)
	openTracing, ok := val.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if !openTracing.Enabled {
		t.Errorf("expected annotation value to be true, got false")
	}

	if !openTracing.TrustEnabled {
		t.Errorf("expected annotation value to be true, got false")
	}
}

func TestIngressAnnotationOpentracingUnset(t *testing.T) {
	ing := buildIngress()

	// Test with no annotation specified
	data := map[string]string{}
	ing.SetAnnotations(data)

	val, _ := NewParser(&resolver.Mock{}).Parse(ing)
	_, ok := val.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}
}
