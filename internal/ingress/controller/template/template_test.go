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

package template

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/ingress-nginx/internal/ingress"
	"k8s.io/ingress-nginx/internal/ingress/annotations/authreq"
	"k8s.io/ingress-nginx/internal/ingress/annotations/influxdb"
	"k8s.io/ingress-nginx/internal/ingress/annotations/modsecurity"
	"k8s.io/ingress-nginx/internal/ingress/annotations/opentracing"
	"k8s.io/ingress-nginx/internal/ingress/annotations/ratelimit"
	"k8s.io/ingress-nginx/internal/ingress/annotations/rewrite"
	"k8s.io/ingress-nginx/internal/ingress/controller/config"
	"k8s.io/ingress-nginx/internal/nginx"
)

func init() {
	// the default value of nginx.TemplatePath assumes the template exists in
	// the root filesystem and not in the rootfs directory
	path, err := filepath.Abs(filepath.Join("../../../../rootfs/", nginx.TemplatePath))
	if err == nil {
		nginx.TemplatePath = path
	}
}

var (
	// TODO: add tests for SSLPassthrough
	tmplFuncTestcases = map[string]struct {
		Path             string
		Target           string
		Location         string
		ProxyPass        string
		Sticky           bool
		XForwardedPrefix string
		SecureBackend    bool
		enforceRegex     bool
	}{
		"when secure backend enabled": {
			"/",
			"/",
			"/",
			"proxy_pass https://upstream_balancer;",
			false,
			"",
			true,
			false,
		},
		"when secure backend and dynamic config enabled": {
			"/",
			"/",
			"/",
			"proxy_pass https://upstream_balancer;",
			false,
			"",
			true,
			false,
		},
		"when secure backend, stickiness and dynamic config enabled": {
			"/",
			"/",
			"/",
			"proxy_pass https://upstream_balancer;",
			true,
			"",
			true,
			false,
		},
		"invalid redirect / to / with dynamic config enabled": {
			"/",
			"/",
			"/",
			"proxy_pass http://upstream_balancer;",
			false,
			"",
			false,
			false,
		},
		"invalid redirect / to /": {
			"/",
			"/",
			"/",
			"proxy_pass http://upstream_balancer;",
			false,
			"",
			false,
			false,
		},
		"redirect / to /jenkins": {
			"/",
			"/jenkins",
			`~* "^/"`,
			`
rewrite "(?i)/" /jenkins break;
proxy_pass http://upstream_balancer;`,
			false,
			"",
			false,
			true,
		},
		"redirect / to /something with sticky enabled": {
			"/",
			"/something",
			`~* "^/"`,
			`
rewrite "(?i)/" /something break;
proxy_pass http://upstream_balancer;`,
			true,
			"",
			false,
			true,
		},
		"redirect / to /something with sticky and dynamic config enabled": {
			"/",
			"/something",
			`~* "^/"`,
			`
rewrite "(?i)/" /something break;
proxy_pass http://upstream_balancer;`,
			true,
			"",
			false,
			true,
		},
		"add the X-Forwarded-Prefix header": {
			"/there",
			"/something",
			`~* "^/there"`,
			`
rewrite "(?i)/there" /something break;
proxy_set_header X-Forwarded-Prefix "/there";
proxy_pass http://upstream_balancer;`,
			true,
			"/there",
			false,
			true,
		},
		"use ~* location modifier when ingress does not use rewrite/regex target but at least one other ingress does": {
			"/something",
			"/something",
			`~* "^/something"`,
			"proxy_pass http://upstream_balancer;",
			false,
			"",
			false,
			true,
		},
	}
)

func TestBuildLuaSharedDictionaries(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""

	// config lua dict
	cfg := config.Configuration{
		LuaSharedDicts: map[string]int{
			"configuration_data": 10000, "certificate_data": 20000,
		},
	}
	actual := buildLuaSharedDictionaries(cfg, invalidType)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	servers := []*ingress.Server{
		{
			Hostname:  "foo.bar",
			Locations: []*ingress.Location{{Path: "/"}},
		},
		{
			Hostname:  "another.host",
			Locations: []*ingress.Location{{Path: "/"}},
		},
	}
	// returns value from config
	configuration := buildLuaSharedDictionaries(cfg, servers)
	if !strings.Contains(configuration, "lua_shared_dict configuration_data 10M;\n") {
		t.Errorf("expected to include 'configuration_data' but got %s", configuration)
	}
	if !strings.Contains(configuration, "lua_shared_dict certificate_data 20M;\n") {
		t.Errorf("expected to include 'certificate_data' but got %s", configuration)
	}
	// test invalid config
	configuration = buildLuaSharedDictionaries(invalidType, servers)
	if configuration != "" {
		t.Errorf("expected an empty string, but got %s", configuration)
	}

	if actual != expected {
		t.Errorf("Expected '%v' but returned '%v' ", expected, actual)
	}
}

func TestLuaConfigurationRequestBodySize(t *testing.T) {
	cfg := config.Configuration{
		LuaSharedDicts: map[string]int{
			"configuration_data": 10000, "certificate_data": 20000,
		},
	}

	size := luaConfigurationRequestBodySize(cfg)
	if size != "21M" {
		t.Errorf("expected the size to be 20 but got: %v", size)
	}
}

func TestFormatIP(t *testing.T) {
	cases := map[string]struct {
		Input, Output string
	}{
		"ipv4-localhost": {"127.0.0.1", "127.0.0.1"},
		"ipv4-internet":  {"8.8.8.8", "8.8.8.8"},
		"ipv6-localhost": {"::1", "[::1]"},
		"ipv6-internet":  {"2001:4860:4860::8888", "[2001:4860:4860::8888]"},
		"invalid-ip":     {"nonsense", "nonsense"},
		"empty-ip":       {"", ""},
	}
	for k, tc := range cases {
		res := formatIP(tc.Input)
		if res != tc.Output {
			t.Errorf("%s: called formatIp('%s'); expected '%v' but returned '%v'", k, tc.Input, tc.Output, res)
		}
	}
}

func TestQuote(t *testing.T) {
	foo := "foo"
	cases := map[interface{}]string{
		"foo":      `"foo"`,
		"\"foo\"":  `"\"foo\""`,
		"foo\nbar": `"foo\nbar"`,
		&foo:       `"foo"`,
		10:         `"10"`,
	}
	for input, output := range cases {
		actual := quote(input)
		if actual != output {
			t.Errorf("quote('%s'): expected '%v' but returned '%v'", input, output, actual)
		}
	}
}

func TestBuildLocation(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := "/"
	actual := buildLocation(invalidType, true)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	for k, tc := range tmplFuncTestcases {
		loc := &ingress.Location{
			Path:    tc.Path,
			Rewrite: rewrite.Config{Target: tc.Target},
		}

		newLoc := buildLocation(loc, tc.enforceRegex)
		if tc.Location != newLoc {
			t.Errorf("%s: expected '%v' but returned %v", k, tc.Location, newLoc)
		}
	}
}

func TestBuildProxyPass(t *testing.T) {
	defaultBackend := "upstream-name"
	defaultHost := "example.com"

	for k, tc := range tmplFuncTestcases {
		loc := &ingress.Location{
			Path:             tc.Path,
			Rewrite:          rewrite.Config{Target: tc.Target},
			Backend:          defaultBackend,
			XForwardedPrefix: tc.XForwardedPrefix,
		}

		if tc.SecureBackend {
			loc.BackendProtocol = "HTTPS"
		}

		backend := &ingress.Backend{
			Name: defaultBackend,
		}

		if tc.Sticky {
			backend.SessionAffinity = ingress.SessionAffinityConfig{
				AffinityType: "cookie",
				CookieSessionAffinity: ingress.CookieSessionAffinity{
					Locations: map[string][]string{
						defaultHost: {tc.Path},
					},
				},
			}
		}

		backends := []*ingress.Backend{backend}

		pp := buildProxyPass(defaultHost, backends, loc)
		if !strings.EqualFold(tc.ProxyPass, pp) {
			t.Errorf("%s: expected \n'%v'\nbut returned \n'%v'", k, tc.ProxyPass, pp)
		}
	}
}

func TestBuildAuthLocation(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildAuthLocation(invalidType, "")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	authURL := "foo.com/auth"
	globalAuthURL := "foo.com/global-auth"

	loc := &ingress.Location{
		ExternalAuth: authreq.Config{
			URL: authURL,
		},
		Path:             "/cat",
		EnableGlobalAuth: true,
	}

	encodedAuthURL := strings.Replace(base64.URLEncoding.EncodeToString([]byte(loc.Path)), "=", "", -1)
	externalAuthPath := fmt.Sprintf("/_external-auth-%v-default", encodedAuthURL)

	testCases := []struct {
		title                    string
		authURL                  string
		globalAuthURL            string
		enableglobalExternalAuth bool
		expected                 string
	}{
		{"authURL, globalAuthURL and enabled", authURL, globalAuthURL, true, externalAuthPath},
		{"authURL, globalAuthURL and disabled", authURL, globalAuthURL, false, externalAuthPath},
		{"authURL, empty globalAuthURL and enabled", authURL, "", true, externalAuthPath},
		{"authURL, empty globalAuthURL and disabled", authURL, "", false, externalAuthPath},
		{"globalAuthURL and enabled", "", globalAuthURL, true, externalAuthPath},
		{"globalAuthURL and disabled", "", globalAuthURL, false, ""},
		{"all empty and enabled", "", "", true, ""},
		{"all empty and disabled", "", "", false, ""},
	}

	for _, testCase := range testCases {
		loc.ExternalAuth.URL = testCase.authURL
		loc.EnableGlobalAuth = testCase.enableglobalExternalAuth

		str := buildAuthLocation(loc, testCase.globalAuthURL)
		if str != testCase.expected {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.title, testCase.expected, str)
		}
	}
}

func TestShouldApplyGlobalAuth(t *testing.T) {

	authURL := "foo.com/auth"
	globalAuthURL := "foo.com/global-auth"

	loc := &ingress.Location{
		ExternalAuth: authreq.Config{
			URL: authURL,
		},
		Path:             "/cat",
		EnableGlobalAuth: true,
	}

	testCases := []struct {
		title                    string
		authURL                  string
		globalAuthURL            string
		enableglobalExternalAuth bool
		expected                 bool
	}{
		{"authURL, globalAuthURL and enabled", authURL, globalAuthURL, true, false},
		{"authURL, globalAuthURL and disabled", authURL, globalAuthURL, false, false},
		{"authURL, empty globalAuthURL and enabled", authURL, "", true, false},
		{"authURL, empty globalAuthURL and disabled", authURL, "", false, false},
		{"globalAuthURL and enabled", "", globalAuthURL, true, true},
		{"globalAuthURL and disabled", "", globalAuthURL, false, false},
		{"all empty and enabled", "", "", true, false},
		{"all empty and disabled", "", "", false, false},
	}

	for _, testCase := range testCases {
		loc.ExternalAuth.URL = testCase.authURL
		loc.EnableGlobalAuth = testCase.enableglobalExternalAuth

		result := shouldApplyGlobalAuth(loc, testCase.globalAuthURL)
		if result != testCase.expected {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.title, testCase.expected, result)
		}
	}
}

func TestBuildAuthResponseHeaders(t *testing.T) {
	externalAuthResponseHeaders := []string{"h1", "H-With-Caps-And-Dashes"}
	expected := []string{
		"auth_request_set $authHeader0 $upstream_http_h1;",
		"proxy_set_header 'h1' $authHeader0;",
		"auth_request_set $authHeader1 $upstream_http_h_with_caps_and_dashes;",
		"proxy_set_header 'H-With-Caps-And-Dashes' $authHeader1;",
	}

	headers := buildAuthResponseHeaders(externalAuthResponseHeaders)

	if !reflect.DeepEqual(expected, headers) {
		t.Errorf("Expected \n'%v'\nbut returned \n'%v'", expected, headers)
	}
}

func TestBuildAuthProxySetHeaders(t *testing.T) {
	proxySetHeaders := map[string]string{
		"header1": "value1",
		"header2": "value2",
	}
	expected := []string{
		"proxy_set_header 'header1' 'value1';",
		"proxy_set_header 'header2' 'value2';",
	}

	headers := buildAuthProxySetHeaders(proxySetHeaders)

	if !reflect.DeepEqual(expected, headers) {
		t.Errorf("Expected \n'%v'\nbut returned \n'%v'", expected, headers)
	}
}

func TestTemplateWithData(t *testing.T) {
	pwd, _ := os.Getwd()
	f, err := os.Open(path.Join(pwd, "../../../../test/data/config.json"))
	if err != nil {
		t.Errorf("unexpected error reading json file: %v", err)
	}
	defer f.Close()
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Error("unexpected error reading json file: ", err)
	}
	var dat config.TemplateConfig
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &dat); err != nil {
		t.Errorf("unexpected error unmarshalling json: %v", err)
	}
	if dat.ListenPorts == nil {
		dat.ListenPorts = &config.ListenPorts{}
	}

	ngxTpl, err := NewTemplate(nginx.TemplatePath)
	if err != nil {
		t.Errorf("invalid NGINX template: %v", err)
	}

	dat.Cfg.DefaultSSLCertificate = &ingress.SSLCert{}

	rt, err := ngxTpl.Write(dat)
	if err != nil {
		t.Errorf("invalid NGINX template: %v", err)
	}

	if !strings.Contains(string(rt), "listen [2001:db8:a0b:12f0::1]") {
		t.Errorf("invalid NGINX template, expected IPV6 listen address not present")
	}

	if !strings.Contains(string(rt), "listen [3731:54:65fe:2::a7]") {
		t.Errorf("invalid NGINX template, expected IPV6 listen address not present")
	}

	if !strings.Contains(string(rt), "listen 2.2.2.2") {
		t.Errorf("invalid NGINX template, expected IPV4 listen address not present")
	}
}

func BenchmarkTemplateWithData(b *testing.B) {
	pwd, _ := os.Getwd()
	f, err := os.Open(path.Join(pwd, "../../../../test/data/config.json"))
	if err != nil {
		b.Errorf("unexpected error reading json file: %v", err)
	}
	defer f.Close()
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		b.Error("unexpected error reading json file: ", err)
	}
	var dat config.TemplateConfig
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &dat); err != nil {
		b.Errorf("unexpected error unmarshalling json: %v", err)
	}

	ngxTpl, err := NewTemplate(nginx.TemplatePath)
	if err != nil {
		b.Errorf("invalid NGINX template: %v", err)
	}

	for i := 0; i < b.N; i++ {
		ngxTpl.Write(dat)
	}
}

func TestBuildDenyVariable(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildDenyVariable(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	a := buildDenyVariable("host1.example.com_/.well-known/acme-challenge")
	b := buildDenyVariable("host1.example.com_/.well-known/acme-challenge")
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected '%v' but returned '%v'", a, b)
	}
}

func TestBuildByteSize(t *testing.T) {
	cases := []struct {
		value    interface{}
		isOffset bool
		expected bool
	}{
		{"1000", false, true},
		{"1000k", false, true},
		{"1m", false, true},
		{"10g", false, false},
		{" 1m ", false, true},
		{"1000kk", false, false},
		{"1000km", false, false},
		{"1mm", false, false},
		{nil, false, false},
		{"", false, false},
		{"    ", false, false},
		{"1G", true, true},
		{"1000kk", true, false},
		{"", true, false},
	}

	for _, tc := range cases {
		val := isValidByteSize(tc.value, tc.isOffset)
		if tc.expected != val {
			t.Errorf("Expected '%v' but returned '%v'", tc.expected, val)
		}
	}
}

func TestIsLocationAllowed(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := false
	actual := isLocationAllowed(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	loc := ingress.Location{
		Denied: nil,
	}

	isAllowed := isLocationAllowed(&loc)
	if !isAllowed {
		t.Errorf("Expected '%v' but returned '%v'", true, isAllowed)
	}
}

func TestBuildForwardedFor(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildForwardedFor(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	inputStr := "X-Forwarded-For"
	expected = "$http_x_forwarded_for"
	actual = buildForwardedFor(inputStr)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestBuildResolvers(t *testing.T) {
	ipOne := net.ParseIP("192.0.0.1")
	ipTwo := net.ParseIP("2001:db8:1234:0000:0000:0000:0000:0000")
	ipList := []net.IP{ipOne, ipTwo}

	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildResolvers(invalidType, false)

	// Invalid Type for []net.IP
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	actual = buildResolvers(ipList, invalidType)

	// Invalid Type for bool
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	validResolver := "resolver 192.0.0.1 [2001:db8:1234::] valid=30s;"
	resolver := buildResolvers(ipList, false)

	if resolver != validResolver {
		t.Errorf("Expected '%v' but returned '%v'", validResolver, resolver)
	}

	validResolver = "resolver 192.0.0.1 valid=30s ipv6=off;"
	resolver = buildResolvers(ipList, true)

	if resolver != validResolver {
		t.Errorf("Expected '%v' but returned '%v'", validResolver, resolver)
	}
}

func TestBuildNextUpstream(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildNextUpstream(invalidType, "")

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cases := map[string]struct {
		NextUpstream  string
		NonIdempotent bool
		Output        string
	}{
		"default": {
			"timeout http_500 http_502",
			false,
			"timeout http_500 http_502",
		},
		"global": {
			"timeout http_500 http_502",
			true,
			"timeout http_500 http_502 non_idempotent",
		},
		"local": {
			"timeout http_500 http_502 non_idempotent",
			false,
			"timeout http_500 http_502 non_idempotent",
		},
	}

	for k, tc := range cases {
		nextUpstream := buildNextUpstream(tc.NextUpstream, tc.NonIdempotent)
		if nextUpstream != tc.Output {
			t.Errorf(
				"%s: called buildNextUpstream('%s', %v); expected '%v' but returned '%v'",
				k,
				tc.NextUpstream,
				tc.NonIdempotent,
				tc.Output,
				nextUpstream,
			)
		}
	}
}

func TestBuildRateLimit(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := []string{}
	actual := buildRateLimit(invalidType)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	loc := &ingress.Location{}

	loc.RateLimit.Connections.Name = "con"
	loc.RateLimit.Connections.Limit = 1

	loc.RateLimit.RPS.Name = "rps"
	loc.RateLimit.RPS.Limit = 1
	loc.RateLimit.RPS.Burst = 1

	loc.RateLimit.RPM.Name = "rpm"
	loc.RateLimit.RPM.Limit = 2
	loc.RateLimit.RPM.Burst = 2

	loc.RateLimit.LimitRateAfter = 1
	loc.RateLimit.LimitRate = 1

	validLimits := []string{
		"limit_conn con 1;",
		"limit_req zone=rps burst=1 nodelay;",
		"limit_req zone=rpm burst=2 nodelay;",
		"limit_rate_after 1k;",
		"limit_rate 1k;",
	}

	limits := buildRateLimit(loc)

	for i, limit := range limits {
		if limit != validLimits[i] {
			t.Errorf("Expected '%v' but returned '%v'", validLimits, limits)
		}
	}

	// Invalid limit
	limits = buildRateLimit(&ingress.Ingress{})
	if !reflect.DeepEqual(expected, limits) {
		t.Errorf("Expected '%v' but returned '%v'", expected, limits)
	}
}

// TODO: Needs more tests
func TestBuildRateLimitZones(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := []string{}
	actual := buildRateLimitZones(invalidType)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

// TODO: Needs more tests
func TestFilterRateLimits(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := []ratelimit.Config{}
	actual := filterRateLimits(invalidType)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestBuildAuthSignURL(t *testing.T) {
	cases := map[string]struct {
		Input, RedirectParam, Output string
	}{
		"default url and redirect":              {"http://google.com", "rd", "http://google.com?rd=$pass_access_scheme://$http_host$escaped_request_uri"},
		"default url and custom redirect":       {"http://google.com", "orig", "http://google.com?orig=$pass_access_scheme://$http_host$escaped_request_uri"},
		"with random field":                     {"http://google.com?cat=0", "rd", "http://google.com?cat=0&rd=$pass_access_scheme://$http_host$escaped_request_uri"},
		"with random field and custom redirect": {"http://google.com?cat=0", "orig", "http://google.com?cat=0&orig=$pass_access_scheme://$http_host$escaped_request_uri"},
		"with rd field":                         {"http://google.com?cat&rd=$request", "rd", "http://google.com?cat&rd=$request"},
		"with orig field":                       {"http://google.com?cat&orig=$request", "orig", "http://google.com?cat&orig=$request"},
	}
	for k, tc := range cases {
		res := buildAuthSignURL(tc.Input, tc.RedirectParam)
		if res != tc.Output {
			t.Errorf("%s: called buildAuthSignURL('%s','%s'); expected '%v' but returned '%v'", k, tc.Input, tc.RedirectParam, tc.Output, res)
		}
	}
}

func TestIsLocationInLocationList(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := false
	actual := isLocationInLocationList(invalidType, "")

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	testCases := []struct {
		location        *ingress.Location
		rawLocationList string
		expected        bool
	}{
		{&ingress.Location{Path: "/match"}, "/match", true},
		{&ingress.Location{Path: "/match"}, ",/match", true},
		{&ingress.Location{Path: "/match"}, "/dontmatch", false},
		{&ingress.Location{Path: "/match"}, ",/dontmatch", false},
		{&ingress.Location{Path: "/match"}, "/dontmatch,/match", true},
		{&ingress.Location{Path: "/match"}, "/dontmatch,/dontmatcheither", false},
	}

	for _, testCase := range testCases {
		result := isLocationInLocationList(testCase.location, testCase.rawLocationList)
		if result != testCase.expected {
			t.Errorf(" expected %v but return %v, path: '%s', rawLocation: '%s'", testCase.expected, result, testCase.location.Path, testCase.rawLocationList)
		}
	}
}

func TestBuildUpstreamName(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildUpstreamName(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	defaultBackend := "upstream-name"
	defaultHost := "example.com"

	for k, tc := range tmplFuncTestcases {
		loc := &ingress.Location{
			Path:             tc.Path,
			Rewrite:          rewrite.Config{Target: tc.Target},
			Backend:          defaultBackend,
			XForwardedPrefix: tc.XForwardedPrefix,
		}

		if tc.SecureBackend {
			loc.BackendProtocol = "HTTPS"
		}

		backend := &ingress.Backend{
			Name: defaultBackend,
		}

		expected := defaultBackend

		if tc.Sticky {
			backend.SessionAffinity = ingress.SessionAffinityConfig{
				AffinityType: "cookie",
				CookieSessionAffinity: ingress.CookieSessionAffinity{
					Locations: map[string][]string{
						defaultHost: {tc.Path},
					},
				},
			}
		}

		pp := buildUpstreamName(loc)
		if !strings.EqualFold(expected, pp) {
			t.Errorf("%s: expected \n'%v'\nbut returned \n'%v'", k, expected, pp)
		}
	}
}

func TestEscapeLiteralDollar(t *testing.T) {
	escapedPath := escapeLiteralDollar("/$")
	expected := "/${literal_dollar}"
	if escapedPath != expected {
		t.Errorf("Expected %v but returned %v", expected, escapedPath)
	}

	escapedPath = escapeLiteralDollar("/hello-$/world-$/")
	expected = "/hello-${literal_dollar}/world-${literal_dollar}/"
	if escapedPath != expected {
		t.Errorf("Expected %v but returned %v", expected, escapedPath)
	}

	leaveUnchanged := "/leave-me/unchanged"
	escapedPath = escapeLiteralDollar(leaveUnchanged)
	if escapedPath != leaveUnchanged {
		t.Errorf("Expected %v but returned %v", leaveUnchanged, escapedPath)
	}

	escapedPath = escapeLiteralDollar(false)
	expected = ""
	if escapedPath != expected {
		t.Errorf("Expected %v but returned %v", expected, escapedPath)
	}
}

func TestOpentracingPropagateContext(t *testing.T) {
	tests := map[*ingress.Location]string{
		{BackendProtocol: "HTTP"}:  "opentracing_propagate_context;",
		{BackendProtocol: "HTTPS"}: "opentracing_propagate_context;",
		{BackendProtocol: "GRPC"}:  "opentracing_grpc_propagate_context;",
		{BackendProtocol: "GRPCS"}: "opentracing_grpc_propagate_context;",
		{BackendProtocol: "AJP"}:   "opentracing_propagate_context;",
		{BackendProtocol: "FCGI"}:  "opentracing_propagate_context;",
		nil:                        "",
	}

	for loc, expectedDirective := range tests {
		actualDirective := opentracingPropagateContext(loc)
		if actualDirective != expectedDirective {
			t.Errorf("Expected %v but returned %v", expectedDirective, actualDirective)
		}
	}
}

func TestGetIngressInformation(t *testing.T) {

	testcases := map[string]struct {
		Ingress  interface{}
		Host     string
		Path     interface{}
		Expected *ingressInformation
	}{
		"wrong ingress type": {
			"wrongtype",
			"host1",
			"/ok",
			&ingressInformation{},
		},
		"wrong path type": {
			&ingress.Ingress{},
			"host1",
			10,
			&ingressInformation{},
		},
		"valid ingress definition with name validIng in namespace default": {
			&ingress.Ingress{
				Ingress: networking.Ingress{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "validIng",
						Namespace: apiv1.NamespaceDefault,
						Annotations: map[string]string{
							"ingress.annotation": "ok",
						},
					},
					Spec: networking.IngressSpec{
						Backend: &networking.IngressBackend{
							ServiceName: "a-svc",
						},
					},
				},
			},
			"host1",
			"",
			&ingressInformation{
				Namespace: "default",
				Rule:      "validIng",
				Annotations: map[string]string{
					"ingress.annotation": "ok",
				},
				Service: "a-svc",
			},
		},
		"valid ingress definition with name demo in namespace something and path /ok using a service with name b-svc port 80": {
			&ingress.Ingress{
				Ingress: networking.Ingress{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "demo",
						Namespace: "something",
						Annotations: map[string]string{
							"ingress.annotation": "ok",
						},
					},
					Spec: networking.IngressSpec{
						Rules: []networking.IngressRule{
							{
								Host: "foo.bar",
								IngressRuleValue: networking.IngressRuleValue{
									HTTP: &networking.HTTPIngressRuleValue{
										Paths: []networking.HTTPIngressPath{
											{
												Path: "/ok",
												Backend: networking.IngressBackend{
													ServiceName: "b-svc",
													ServicePort: intstr.FromInt(80),
												},
											},
										},
									},
								},
							},
							{},
						},
					},
				},
			},
			"foo.bar",
			"/ok",
			&ingressInformation{
				Namespace: "something",
				Rule:      "demo",
				Annotations: map[string]string{
					"ingress.annotation": "ok",
				},
				Service:     "b-svc",
				ServicePort: "80",
			},
		},
	}

	for title, testCase := range testcases {
		info := getIngressInformation(testCase.Ingress, testCase.Host, testCase.Path)

		if !info.Equal(testCase.Expected) {
			t.Fatalf("%s: expected '%v' but returned %v", title, testCase.Expected, info)
		}
	}
}

func TestBuildCustomErrorLocationsPerServer(t *testing.T) {
	testCases := []struct {
		server          interface{}
		expectedResults []errorLocation
	}{
		{ // Single ingress
			&ingress.Server{Locations: []*ingress.Location{
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-backend",
					CustomHTTPErrors:           []int{401, 402},
				},
			}},
			[]errorLocation{
				{
					UpstreamName: "custom-default-backend-test-backend",
					Codes:        []int{401, 402},
				},
			},
		},
		{ // Two ingresses, overlapping error codes, same backend
			&ingress.Server{Locations: []*ingress.Location{
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-backend",
					CustomHTTPErrors:           []int{401, 402},
				},
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-backend",
					CustomHTTPErrors:           []int{402, 403},
				},
			}},
			[]errorLocation{
				{
					UpstreamName: "custom-default-backend-test-backend",
					Codes:        []int{401, 402, 403},
				},
			},
		},
		{ // Two ingresses, overlapping error codes, different backends
			&ingress.Server{Locations: []*ingress.Location{
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-one",
					CustomHTTPErrors:           []int{401, 402},
				},
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-two",
					CustomHTTPErrors:           []int{402, 403},
				},
			}},
			[]errorLocation{
				{
					UpstreamName: "custom-default-backend-test-one",
					Codes:        []int{401, 402},
				},
				{
					UpstreamName: "custom-default-backend-test-two",
					Codes:        []int{402, 403},
				},
			},
		},
		{ // Many ingresses, overlapping error codes, different backends
			&ingress.Server{Locations: []*ingress.Location{
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-one",
					CustomHTTPErrors:           []int{401, 402},
				},
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-one",
					CustomHTTPErrors:           []int{501, 502},
				},
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-two",
					CustomHTTPErrors:           []int{409, 410},
				},
				{
					DefaultBackendUpstreamName: "custom-default-backend-test-two",
					CustomHTTPErrors:           []int{504, 505},
				},
			}},
			[]errorLocation{
				{
					UpstreamName: "custom-default-backend-test-one",
					Codes:        []int{401, 402, 501, 502},
				},
				{
					UpstreamName: "custom-default-backend-test-two",
					Codes:        []int{409, 410, 504, 505},
				},
			},
		},
	}

	for _, c := range testCases {
		response := buildCustomErrorLocationsPerServer(c.server)
		if !reflect.DeepEqual(c.expectedResults, response) {
			t.Errorf("Expected %+v but got %+v", c.expectedResults, response)
		}
	}
}

func TestProxySetHeader(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := "proxy_set_header"
	actual := proxySetHeader(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	grpcBackend := &ingress.Location{
		BackendProtocol: "GRPC",
	}

	expected = "grpc_set_header"
	actual = proxySetHeader(grpcBackend)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestBuildInfluxDB(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildInfluxDB(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cfg := influxdb.Config{
		InfluxDBEnabled:     true,
		InfluxDBServerName:  "ok.com",
		InfluxDBHost:        "host.com",
		InfluxDBPort:        "5252",
		InfluxDBMeasurement: "ok",
	}
	expected = "influxdb server_name=ok.com host=host.com port=5252 measurement=ok enabled=true;"
	actual = buildInfluxDB(cfg)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestBuildOpenTracing(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := ""
	actual := buildOpentracing(invalidType, []*ingress.Server{})

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cfgJaeger := config.Configuration{
		EnableOpentracing:   true,
		JaegerCollectorHost: "jaeger-host.com",
	}
	expected = "opentracing_load_tracer /usr/local/lib/libjaegertracing_plugin.so /etc/nginx/opentracing.json;\r\n"
	actual = buildOpentracing(cfgJaeger, []*ingress.Server{})

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cfgZipkin := config.Configuration{
		EnableOpentracing:   true,
		ZipkinCollectorHost: "zipkin-host.com",
	}
	expected = "opentracing_load_tracer /usr/local/lib/libzipkin_opentracing_plugin.so /etc/nginx/opentracing.json;\r\n"
	actual = buildOpentracing(cfgZipkin, []*ingress.Server{})

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cfgDatadog := config.Configuration{
		EnableOpentracing:    true,
		DatadogCollectorHost: "datadog-host.com",
	}
	expected = "opentracing_load_tracer /usr/local/lib64/libdd_opentracing.so /etc/nginx/opentracing.json;\r\n"
	actual = buildOpentracing(cfgDatadog, []*ingress.Server{})

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	cfgOpenTracing := config.Configuration{
		EnableOpentracing:                true,
		DatadogCollectorHost:             "datadog-host.com",
		OpentracingOperationName:         "my-operation-name",
		OpentracingLocationOperationName: "my-location-operation-name",
	}
	expected = "opentracing_load_tracer /usr/local/lib64/libdd_opentracing.so /etc/nginx/opentracing.json;\r\n"
	expected += "opentracing_operation_name \"my-operation-name\";\n"
	expected += "opentracing_location_operation_name \"my-location-operation-name\";\n"
	actual = buildOpentracing(cfgOpenTracing, []*ingress.Server{})

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

}

func TestEnforceRegexModifier(t *testing.T) {
	invalidType := &ingress.Ingress{}
	expected := false
	actual := enforceRegexModifier(invalidType)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	locs := []*ingress.Location{
		{
			Rewrite: rewrite.Config{
				Target:   "/alright",
				UseRegex: true,
			},
			Path: "/ok",
		},
	}
	expected = true
	actual = enforceRegexModifier(locs)

	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestShouldLoadModSecurityModule(t *testing.T) {
	// ### Invalid argument type tests ###
	// The first tests should return false.
	expected := false

	invalidType := &ingress.Ingress{}
	actual := shouldLoadModSecurityModule(config.Configuration{}, invalidType)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	actual = shouldLoadModSecurityModule(invalidType, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	// ### Functional tests ###
	actual = shouldLoadModSecurityModule(config.Configuration{}, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	// All further tests should return true.
	expected = true

	configuration := config.Configuration{EnableModsecurity: true}
	actual = shouldLoadModSecurityModule(configuration, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	servers := []*ingress.Server{
		{
			Locations: []*ingress.Location{
				{
					ModSecurity: modsecurity.Config{
						Enable: true,
					},
				},
			},
		},
	}
	actual = shouldLoadModSecurityModule(config.Configuration{}, servers)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	actual = shouldLoadModSecurityModule(configuration, servers)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestOpentracingForLocation(t *testing.T) {
	trueVal := true

	loadOT := `opentracing on;
opentracing_propagate_context;`
	testCases := []struct {
		description string
		globalOT    bool
		isSetInLoc  bool
		isOTInLoc   *bool
		expected    string
	}{
		{"globally enabled, without annotation", true, false, nil, loadOT},
		{"globally enabled and enabled in location", true, true, &trueVal, loadOT},
		{"globally disabled and not enabled in location", false, false, nil, ""},
		{"globally disabled but enabled in location", false, true, &trueVal, loadOT},
		{"globally disabled, enabled in location but false", false, true, &trueVal, loadOT},
	}

	for _, testCase := range testCases {
		il := &ingress.Location{
			Opentracing: opentracing.Config{Set: testCase.isSetInLoc},
		}
		if il.Opentracing.Set {
			il.Opentracing.Enabled = *testCase.isOTInLoc
		}

		actual := buildOpentracingForLocation(testCase.globalOT, il)

		if testCase.expected != actual {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.description, testCase.expected, actual)
		}
	}
}

func TestShouldLoadOpentracingModule(t *testing.T) {
	// ### Invalid argument type tests ###
	// The first tests should return false.
	expected := false

	invalidType := &ingress.Ingress{}
	actual := shouldLoadOpentracingModule(config.Configuration{}, invalidType)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	actual = shouldLoadOpentracingModule(invalidType, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	// ### Functional tests ###
	actual = shouldLoadOpentracingModule(config.Configuration{}, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	// All further tests should return true.
	expected = true

	configuration := config.Configuration{EnableOpentracing: true}
	actual = shouldLoadOpentracingModule(configuration, []*ingress.Server{})
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	servers := []*ingress.Server{
		{
			Locations: []*ingress.Location{
				{
					Opentracing: opentracing.Config{
						Enabled: true,
					},
				},
			},
		},
	}
	actual = shouldLoadOpentracingModule(config.Configuration{}, servers)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}

	actual = shouldLoadOpentracingModule(configuration, servers)
	if expected != actual {
		t.Errorf("Expected '%v' but returned '%v'", expected, actual)
	}
}

func TestModSecurityForLocation(t *testing.T) {
	loadModule := `modsecurity on;
`

	modSecCfg := `modsecurity_rules_file /etc/nginx/modsecurity/modsecurity.conf;
`

	modsecOff := "modsecurity off;"

	modsecRule := `modsecurity_rules '
#RULE#
';
`
	owaspRules := `modsecurity_rules_file /etc/nginx/owasp-modsecurity-crs/nginx-modsecurity.conf;
`

	transactionID := "0000000"
	transactionCfg := fmt.Sprintf(`modsecurity_transaction_id "%v";
`, transactionID)

	testRule := "#RULE#"

	testCases := []struct {
		description         string
		isEnabledInCM       bool
		isOwaspEnabledInCM  bool
		isEnabledInLoc      bool
		isEnableSetInLoc    bool
		isOwaspEnabledInLoc bool
		snippet             string
		transactionID       string
		expected            string
	}{
		{"configmap enabled, configmap OWASP disabled, without annotation, snippet or transaction ID", true, false, false, false, false, "", "", ""},
		{"configmap enabled, configmap OWASP disabled, without annotation, snippet and with transaction ID", true, false, false, false, false, "", transactionID, transactionCfg},
		{"configmap enabled, configmap OWASP enabled, without annotation, OWASP enabled", true, true, false, false, false, "", "", ""},
		{"configmap enabled, configmap OWASP enabled, without annotation, OWASP disabled, with snippet and no transaction ID", true, true, false, false, false, testRule, "", modsecRule},
		{"configmap enabled, configmap OWASP enabled, without annotation, OWASP disabled, with snippet and transaction ID", true, true, false, false, false, testRule, transactionID, fmt.Sprintf("%v%v", modsecRule, transactionCfg)},
		{"configmap enabled, configmap OWASP disabled, annotation enabled, OWASP disabled", true, false, true, true, false, "", "", ""},
		{"configmap enabled, configmap OWASP disabled, annotation enabled, OWASP enabled, no snippet and no transaction ID", true, false, true, true, true, "", "", owaspRules},
		{"configmap enabled, configmap OWASP disabled, annotation disabled, OWASP disabled, no snippet and no transaction ID", true, false, false, true, false, "", "", modsecOff},
		{"configmap enabled, configmap OWASP disabled, annotation enabled, OWASP enabled, with snippet and no transaction ID", true, false, true, true, true, "", "", owaspRules},
		{"configmap enabled, configmap OWASP disabled, annotation enabled, OWASP enabled, with snippet and transaction ID", true, false, true, true, true, "", transactionID, fmt.Sprintf("%v%v", transactionCfg, owaspRules)},
		{"configmap enabled, configmap OWASP enabled, annotation enabled, OWASP disabled", true, true, true, true, false, "", "", ""},
		{"configmap disabled, annotation enabled, OWASP disabled", false, false, true, true, false, "", "", fmt.Sprintf("%v%v", loadModule, modSecCfg)},
		{"configmap disabled, annotation disabled, OWASP disabled", false, false, false, true, false, "", "", ""},
		{"configmap disabled, annotation enabled, OWASP disabled", false, false, true, true, false, testRule, "", fmt.Sprintf("%v%v%v", loadModule, modsecRule, modSecCfg)},
		{"configmap disabled, annotation enabled, OWASP enabled", false, false, true, true, false, testRule, "", fmt.Sprintf("%v%v%v", loadModule, modsecRule, modSecCfg)},
	}

	for _, testCase := range testCases {
		il := &ingress.Location{
			ModSecurity: modsecurity.Config{
				Enable:        testCase.isEnabledInLoc,
				EnableSet:     testCase.isEnableSetInLoc,
				OWASPRules:    testCase.isOwaspEnabledInLoc,
				Snippet:       testCase.snippet,
				TransactionID: testCase.transactionID,
			},
		}

		actual := buildModSecurityForLocation(config.Configuration{
			EnableModsecurity:    testCase.isEnabledInCM,
			EnableOWASPCoreRules: testCase.isOwaspEnabledInCM,
		}, il)

		if testCase.expected != actual {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.description, testCase.expected, actual)
		}
	}
}

func TestBuildServerName(t *testing.T) {

	testCases := []struct {
		title    string
		hostname string
		expected string
	}{
		{"simple domain", "foo.bar", "foo.bar"},
		{"simple www domain", "www.foo.bar", "www.foo.bar"},
		{"wildcard domain", "*.foo.bar", "~^(?<subdomain>[\\w-]+)\\.foo\\.bar$"},
		{"wildcard two levels domain", "*.sub.foo.bar", "~^(?<subdomain>[\\w-]+)\\.sub\\.foo\\.bar$"},
	}

	for _, testCase := range testCases {
		result := buildServerName(testCase.hostname)
		if result != testCase.expected {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.title, testCase.expected, result)
		}
	}
}

func TestParseComplexNginxVarIntoLuaTable(t *testing.T) {
	testCases := []struct {
		ngxVar           string
		expectedLuaTable string
	}{
		{"foo", `{ { nil, nil, nil, "foo", }, }`},
		{"$foo", `{ { nil, nil, "foo", nil, }, }`},
		{"${foo}", `{ { nil, "foo", nil, nil, }, }`},
		{"\\$foo", `{ { "\$foo", nil, nil, nil, }, }`},
		{
			"foo\\$bar$baz${daz}xiyar$pomidor",
			`{ { nil, nil, nil, "foo", }, { "\$bar", nil, nil, nil, }, { nil, nil, "baz", nil, }, ` +
				`{ nil, "daz", nil, nil, }, { nil, nil, nil, "xiyar", }, { nil, nil, "pomidor", nil, }, }`,
		},
	}

	for _, testCase := range testCases {
		actualLuaTable := parseComplexNginxVarIntoLuaTable(testCase.ngxVar)
		if actualLuaTable != testCase.expectedLuaTable {
			t.Errorf("expected %v but returned %v", testCase.expectedLuaTable, actualLuaTable)
		}
	}
}

func TestConvertGoSliceIntoLuaTablet(t *testing.T) {
	testCases := []struct {
		title            string
		goSlice          interface{}
		emptyStringAsNil bool
		expectedLuaTable string
		expectedErr      error
	}{
		{
			"flat string slice",
			[]string{"one", "two", "three"},
			false,
			`{ "one", "two", "three", }`,
			nil,
		},
		{
			"nested string slice",
			[][]string{{"one", "", "three"}, {"foo", "bar"}},
			false,
			`{ { "one", "", "three", }, { "foo", "bar", }, }`,
			nil,
		},
		{
			"converts empty string to nil when enabled",
			[][]string{{"one", "", "three"}, {"foo", "bar"}},
			true,
			`{ { "one", nil, "three", }, { "foo", "bar", }, }`,
			nil,
		},
		{
			"boolean slice",
			[]bool{true, true, false},
			false,
			`{ true, true, false, }`,
			nil,
		},
		{
			"integer slice",
			[]int{4, 3, 6},
			false,
			`{ 4, 3, 6, }`,
			nil,
		},
	}

	for _, testCase := range testCases {
		actualLuaTable, err := convertGoSliceIntoLuaTable(testCase.goSlice, testCase.emptyStringAsNil)
		if testCase.expectedErr != nil && err != nil && testCase.expectedErr.Error() != err.Error() {
			t.Errorf("expected error '%v' but returned '%v'", testCase.expectedErr, err)
		}
		if testCase.expectedErr == nil && err != nil {
			t.Errorf("expected error to be nil but returned '%v'", err)
		}
		if testCase.expectedLuaTable != actualLuaTable {
			t.Errorf("%v: expected '%v' but returned '%v'", testCase.title, testCase.expectedLuaTable, actualLuaTable)
		}
	}
}
