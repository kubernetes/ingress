/*
Copyright 2016 The Kubernetes Authors.

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

package sessionaffinity

import (
	"regexp"

	networking "k8s.io/api/networking/v1beta1"
	"k8s.io/klog/v2"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

const (
	annotationAffinityType           = "affinity"
	annotationAffinityMode           = "affinity-mode"
	annotationAffinityCanaryBehavior = "affinity-canary-behavior"

	// If a cookie with this name exists,
	// its value is used as an index into the list of available backends.
	annotationAffinityCookieName = "session-cookie-name"

	defaultAffinityCookieName = "INGRESSCOOKIE"

	// This is used to control the cookie expires, its value is a number of seconds until the
	// cookie expires
	annotationAffinityCookieExpires = "session-cookie-expires"

	// This is used to control the cookie expires, its value is a number of seconds until the
	// cookie expires
	annotationAffinityCookieMaxAge = "session-cookie-max-age"

	// This is used to control the cookie path when use-regex is set to true
	annotationAffinityCookiePath = "session-cookie-path"

	// This is used to control the SameSite attribute of the cookie
	annotationAffinityCookieSameSite = "session-cookie-samesite"

	// This is used to control whether SameSite=None should be conditionally applied based on the User-Agent
	annotationAffinityCookieConditionalSameSiteNone = "session-cookie-conditional-samesite-none"

	// This is used to control the cookie change after request failure
	annotationAffinityCookieChangeOnFailure = "session-cookie-change-on-failure"
)

var (
	affinityCookieExpiresRegex = regexp.MustCompile(`(^0|-?[1-9]\d*$)`)
)

// Config describes the per ingress session affinity config
type Config struct {
	// The type of affinity that will be used
	Type string `json:"type"`
	// The affinity mode, i.e. how sticky a session is
	Mode string `json:"mode"`
	// Affinity behavior for canaries (sticky or legacy)
	CanaryBehavior string `json:"canaryBehavior"`
	Cookie
}

// Cookie describes the Config of cookie type affinity
type Cookie struct {
	// The name of the cookie that will be used in case of cookie affinity type.
	Name string `json:"name"`
	// The time duration to control cookie expires
	Expires string `json:"expires"`
	// The number of seconds until the cookie expires
	MaxAge string `json:"maxage"`
	// The path that a cookie will be set on
	Path string `json:"path"`
	// Flag that allows cookie regeneration on request failure
	ChangeOnFailure bool `json:"changeonfailure"`
	// SameSite attribute value
	SameSite string `json:"samesite"`
	// Flag that conditionally applies SameSite=None attribute on cookie if user agent accepts it.
	ConditionalSameSiteNone bool `json:"conditional-samesite-none"`
}

// cookieAffinityParse gets the annotation values related to Cookie Affinity
// It also sets default values when no value or incorrect value is found
func (a affinity) cookieAffinityParse(ing *networking.Ingress) *Cookie {
	var err error

	cookie := &Cookie{}

	cookie.Name, err = parser.GetStringAnnotation(annotationAffinityCookieName, ing)
	if err != nil {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieName, "default", defaultAffinityCookieName)
		cookie.Name = defaultAffinityCookieName
	}

	cookie.Expires, err = parser.GetStringAnnotation(annotationAffinityCookieExpires, ing)
	if err != nil || !affinityCookieExpiresRegex.MatchString(cookie.Expires) {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieExpires)
		cookie.Expires = ""
	}

	cookie.MaxAge, err = parser.GetStringAnnotation(annotationAffinityCookieMaxAge, ing)
	if err != nil || !affinityCookieExpiresRegex.MatchString(cookie.MaxAge) {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieMaxAge)
		cookie.MaxAge = ""
	}

	cookie.Path, err = parser.GetStringAnnotation(annotationAffinityCookiePath, ing)
	if err != nil {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookiePath)
	}

	cookie.SameSite, err = parser.GetStringAnnotation(annotationAffinityCookieSameSite, ing)
	if err != nil {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieSameSite)
	}

	cookie.ConditionalSameSiteNone, err = parser.GetBoolAnnotation(annotationAffinityCookieConditionalSameSiteNone, ing)
	if err != nil {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieConditionalSameSiteNone)
	}

	cookie.ChangeOnFailure, err = parser.GetBoolAnnotation(annotationAffinityCookieChangeOnFailure, ing)
	if err != nil {
		klog.V(3).InfoS("Invalid or no annotation value found. Ignoring", "ingress", klog.KObj(ing), "annotation", annotationAffinityCookieChangeOnFailure)
	}

	return cookie
}

// NewParser creates a new Affinity annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return affinity{r}
}

type affinity struct {
	r resolver.Resolver
}

// ParseAnnotations parses the annotations contained in the ingress
// rule used to configure the affinity directives
func (a affinity) Parse(ing *networking.Ingress) (interface{}, error) {
	cookie := &Cookie{}
	// Check the type of affinity that will be used
	at, err := parser.GetStringAnnotation(annotationAffinityType, ing)
	if err != nil {
		at = ""
	}

	// Check the affinity mode that will be used
	am, err := parser.GetStringAnnotation(annotationAffinityMode, ing)
	if err != nil {
		am = ""
	}

	cb, err := parser.GetStringAnnotation(annotationAffinityCanaryBehavior, ing)
	if err != nil {
		cb = ""
	}

	switch at {
	case "cookie":
		cookie = a.cookieAffinityParse(ing)
	default:
		klog.V(3).InfoS("No default affinity found", "ingress", ing.Name)

	}

	return &Config{
		Type:           at,
		Mode:           am,
		CanaryBehavior: cb,
		Cookie:         *cookie,
	}, nil
}
