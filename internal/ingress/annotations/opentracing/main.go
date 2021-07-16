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
	networking "k8s.io/api/networking/v1beta1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

type opentracing struct {
	r resolver.Resolver
}

// Config contains the configuration to be used in the Ingress
type Config struct {
	Enabled      bool `json:"enabled"`
	Set          bool `json:"set"`
	TrustEnabled bool `json:"trust-enabled"`
	TrustSet     bool `json:"trust-set"`
}

// Equal tests for equality between two Config types
func (bd1 *Config) Equal(bd2 *Config) bool {
	if bd1.Set != bd2.Set {
		return false
	}

	if bd1.Enabled != bd2.Enabled {
		return false
	}

	if bd1.TrustSet != bd2.TrustSet {
		return false
	}

	if bd1.TrustEnabled != bd2.TrustEnabled {
		return false
	}

	return true
}

// NewParser creates a new serviceUpstream annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return opentracing{r}
}

func (s opentracing) Parse(ing *networking.Ingress) (interface{}, error) {
	enabled, err := parser.GetBoolAnnotation("enable-opentracing", ing)
	if err != nil {
		return &Config{}, nil
	}

	trustSpan, err := parser.GetBoolAnnotation("opentracing-trust-incoming-span", ing)
	if err != nil {
		return &Config{Set: true, Enabled: enabled}, nil
	}

	return &Config{Set: true, Enabled: enabled, TrustSet: true, TrustEnabled: trustSpan}, nil
}
