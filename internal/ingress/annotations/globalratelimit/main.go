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
	"strings"
	"time"

	"github.com/pkg/errors"
	networking "k8s.io/api/networking/v1beta1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	ing_errors "k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/internal/net"
	"k8s.io/ingress-nginx/internal/sets"

	"github.com/r3labs/diff/v2"
)

const defaultKey = "$remote_addr"
const defaultWindowSize = "0s"
const defaultHeaderBasedRateLimits = "[]"

// HeaderBasedRateLimit be able to set different rate limits based on header/value match
type HeaderBasedRateLimit struct {
	HeaderName   string   `json:"header-name"`
	HeaderValues []string `json:"header-values"`
	Limit        int      `json:"limit"`
	WindowSize   int      `json:"window-size"`
}

// Config encapsulates all global rate limit attributes
type Config struct {
	Namespace             string                 `json:"namespace"`
	Limit                 int                    `json:"limit"`
	WindowSize            int                    `json:"window-size"`
	Key                   string                 `json:"key"`
	IgnoredCIDRs          []string               `json:"ignored-cidrs"`
	HeaderBasedRateLimits []HeaderBasedRateLimit `json:"header-based-rate-limits"`
}

// Equal tests for equality between two Config types
func (l *Config) Equal(r *Config) bool {
	if l.Namespace != r.Namespace {
		return false
	}
	if l.Limit != r.Limit {
		return false
	}
	if l.WindowSize != r.WindowSize {
		return false
	}
	if l.Key != r.Key {
		return false
	}
	if len(l.IgnoredCIDRs) != len(r.IgnoredCIDRs) || !sets.StringElementsMatch(l.IgnoredCIDRs, r.IgnoredCIDRs) {
		return false
	}

	if len(l.HeaderBasedRateLimits) != len(r.HeaderBasedRateLimits) {
		return false
	}
	for i := range l.HeaderBasedRateLimits {
		change, err := diff.Diff(l.HeaderBasedRateLimits[i], r.HeaderBasedRateLimits[i])
		// Diff two structs - If err: mark as equal
		if err != nil {
			return true
		}
		if len(change) > 0 {
			return false
		}
	}

	return true
}

type globalratelimit struct {
	r resolver.Resolver
}

// NewParser creates a new globalratelimit annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return globalratelimit{r}
}

// Parse extracts globalratelimit annotations from the given ingress
// and returns them structured as Config type
func (a globalratelimit) Parse(ing *networking.Ingress) (interface{}, error) {
	config := &Config{}

	limit, _ := parser.GetIntAnnotation("global-rate-limit", ing)
	rawWindowSize, _ := parser.GetStringAnnotation("global-rate-limit-window", ing)
	if len(rawWindowSize) == 0 {
		rawWindowSize = defaultWindowSize
	}

	windowSize, err := time.ParseDuration(rawWindowSize)
	if err != nil {
		return config, ing_errors.LocationDenied{
			Reason: errors.Wrap(err, "failed to parse 'global-rate-limit-window' value"),
		}
	}

	key, _ := parser.GetStringAnnotation("global-rate-limit-key", ing)
	if len(key) == 0 {
		key = defaultKey
	}

	rawIgnoredCIDRs, _ := parser.GetStringAnnotation("global-rate-limit-ignored-cidrs", ing)
	ignoredCIDRs, err := net.ParseCIDRs(rawIgnoredCIDRs)
	if err != nil {
		return nil, err
	}

	rawHeaderBasedRateLimits, _ := parser.GetStringAnnotation("global-rate-limit-header-based", ing)
	if len(rawHeaderBasedRateLimits) == 0 {
		rawHeaderBasedRateLimits = defaultHeaderBasedRateLimits
	}
	var headerBaseRateLimits []HeaderBasedRateLimit
	err = json.Unmarshal([]byte(rawHeaderBasedRateLimits), &headerBaseRateLimits)
	if err != nil {
		return nil, ing_errors.LocationDenied{
			Reason: errors.Wrap(err, "failed to parse 'global-rate-limit-header-based' json value"),
		}
	}

	if limit == 0 && windowSize == 0 && len(headerBaseRateLimits) == 0 {
		return config, nil
	}

	config.Namespace = strings.Replace(string(ing.UID), "-", "", -1)
	config.Limit = limit
	config.WindowSize = int(windowSize.Seconds())
	config.Key = key
	config.IgnoredCIDRs = ignoredCIDRs
	config.HeaderBasedRateLimits = headerBaseRateLimits

	return config, nil
}
