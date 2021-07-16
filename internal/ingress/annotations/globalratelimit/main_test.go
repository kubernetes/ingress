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

package globalratelimit

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	ing_errors "k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

const UID = "31285d47-b150-4dcf-bd6f-12c46d769f6e"
const expectedUID = "31285d47b1504dcfbd6f12c46d769f6e"

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		ServiceName: "default-backend",
		ServicePort: intstr.FromInt(80),
	}

	return &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
			UID:       UID,
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

type mockBackend struct {
	resolver.Mock
}

func TestGlobalRateLimiting(t *testing.T) {
	ing := buildIngress()

	annRateLimit := parser.GetAnnotationWithPrefix("global-rate-limit")
	annRateLimitWindow := parser.GetAnnotationWithPrefix("global-rate-limit-window")
	annRateLimitKey := parser.GetAnnotationWithPrefix("global-rate-limit-key")
	annRateLimitIgnoredCIDRs := parser.GetAnnotationWithPrefix("global-rate-limit-ignored-cidrs")
	annRateLimitHeaderBased := parser.GetAnnotationWithPrefix("global-rate-limit-header-based")

	testCases := []struct {
		title          string
		annotations    map[string]string
		expectedConfig *Config
		expectedErr    error
	}{
		{
			"no annotation",
			nil,
			&Config{},
			nil,
		},
		{
			"minimum required annotations",
			map[string]string{
				annRateLimit:       "100",
				annRateLimitWindow: "2m",
			},
			&Config{
				Namespace:             expectedUID,
				Limit:                 100,
				WindowSize:            120,
				Key:                   "$remote_addr",
				IgnoredCIDRs:          make([]string, 0),
				HeaderBasedRateLimits: make([]HeaderBasedRateLimit, 0),
			},
			nil,
		},
		{
			"global-rate-limit-key annotation",
			map[string]string{
				annRateLimit:       "100",
				annRateLimitWindow: "2m",
				annRateLimitKey:    "$http_x_api_user",
			},
			&Config{
				Namespace:             expectedUID,
				Limit:                 100,
				WindowSize:            120,
				Key:                   "$http_x_api_user",
				IgnoredCIDRs:          make([]string, 0),
				HeaderBasedRateLimits: make([]HeaderBasedRateLimit, 0),
			},
			nil,
		},
		{
			"global-rate-limit-ignored-cidrs annotation",
			map[string]string{
				annRateLimit:             "100",
				annRateLimitWindow:       "2m",
				annRateLimitKey:          "$http_x_api_user",
				annRateLimitIgnoredCIDRs: "127.0.0.1, 200.200.24.0/24",
			},
			&Config{
				Namespace:             expectedUID,
				Limit:                 100,
				WindowSize:            120,
				Key:                   "$http_x_api_user",
				IgnoredCIDRs:          []string{"127.0.0.1", "200.200.24.0/24"},
				HeaderBasedRateLimits: make([]HeaderBasedRateLimit, 0),
			},
			nil,
		},
		{
			"global-rate-limit-header-based annotation",
			map[string]string{
				annRateLimit:            "100",
				annRateLimitWindow:      "2m",
				annRateLimitKey:         "$http_x_user_name",
				annRateLimitHeaderBased: `[{"header-name":"x-user-name","header-values":["foo","bar"],"limit":1,"window-size":60}]`,
			},
			&Config{
				Namespace:    expectedUID,
				Limit:        100,
				WindowSize:   120,
				Key:          "$http_x_user_name",
				IgnoredCIDRs: make([]string, 0),
				HeaderBasedRateLimits: []HeaderBasedRateLimit{{
					HeaderName:   "x-user-name",
					HeaderValues: []string{"foo", "bar"},
					Limit:        1,
					WindowSize:   60,
				}},
			},
			nil,
		},
		{
			"incorrect duration for window",
			map[string]string{
				annRateLimit:       "100",
				annRateLimitWindow: "2mb",
				annRateLimitKey:    "$http_x_api_user",
			},
			&Config{},
			ing_errors.LocationDenied{
				Reason: errors.Wrap(fmt.Errorf(`time: unknown unit "mb" in duration "2mb"`),
					"failed to parse 'global-rate-limit-window' value"),
			},
		},
	}

	for _, testCase := range testCases {
		ing.SetAnnotations(testCase.annotations)

		i, actualErr := NewParser(mockBackend{}).Parse(ing)
		if (testCase.expectedErr == nil || actualErr == nil) && testCase.expectedErr != actualErr {
			t.Errorf("expected error 'nil' but got '%v'", actualErr)
		} else if testCase.expectedErr != nil && actualErr != nil &&
			testCase.expectedErr.Error() != actualErr.Error() {
			t.Errorf("expected error '%v' but got '%v'", testCase.expectedErr, actualErr)
		}

		actualConfig := i.(*Config)
		if !testCase.expectedConfig.Equal(actualConfig) {
			expectedJSON, _ := json.Marshal(testCase.expectedConfig)
			actualJSON, _ := json.Marshal(actualConfig)
			t.Errorf("%v: expected config '%s' but got '%s'", testCase.title, expectedJSON, actualJSON)
		}
	}
}
