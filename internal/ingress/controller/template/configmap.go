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
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"github.com/mitchellh/hashstructure"
	"github.com/mitchellh/mapstructure"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/ingress-nginx/internal/ingress/annotations/authreq"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/controller/config"
	ing_net "k8s.io/ingress-nginx/internal/net"
	"k8s.io/ingress-nginx/internal/runtime"
)

const (
	customHTTPErrors              = "custom-http-errors"
	skipAccessLogUrls             = "skip-access-log-urls"
	whitelistSourceRange          = "whitelist-source-range"
	proxyRealIPCIDR               = "proxy-real-ip-cidr"
	bindAddress                   = "bind-address"
	httpRedirectCode              = "http-redirect-code"
	blockCIDRs                    = "block-cidrs"
	blockUserAgents               = "block-user-agents"
	blockReferers                 = "block-referers"
	proxyStreamResponses          = "proxy-stream-responses"
	hideHeaders                   = "hide-headers"
	nginxStatusIpv4Whitelist      = "nginx-status-ipv4-whitelist"
	nginxStatusIpv6Whitelist      = "nginx-status-ipv6-whitelist"
	proxyHeaderTimeout            = "proxy-protocol-header-timeout"
	workerProcesses               = "worker-processes"
	globalAuthURL                 = "global-auth-url"
	globalAuthMethod              = "global-auth-method"
	globalAuthSignin              = "global-auth-signin"
	globalAuthSigninRedirectParam = "global-auth-signin-redirect-param"
	globalAuthResponseHeaders     = "global-auth-response-headers"
	globalAuthRequestRedirect     = "global-auth-request-redirect"
	globalAuthSnippet             = "global-auth-snippet"
	globalAuthCacheKey            = "global-auth-cache-key"
	globalAuthCacheDuration       = "global-auth-cache-duration"
	globalAuthTLSVerifyClient     = "global-auth-tls-verify-client"
	globalAuthTLSVerifyDepth      = "global-auth-tls-verify-depth"
	globalAuthTLSSecret           = "global-auth-tls-secret"
	luaSharedDictsKey             = "lua-shared-dicts"
	plugins                       = "plugins"
)

var (
	validRedirectCodes    = sets.NewInt([]int{301, 302, 307, 308}...)
	defaultLuaSharedDicts = map[string]int{
		"configuration_data":            20,
		"certificate_data":              20,
		"balancer_ewma":                 10,
		"balancer_ewma_last_touched_at": 10,
		"balancer_ewma_locks":           1,
		"certificate_servers":           5,
		"ocsp_response_cache":           5, // keep this same as certificate_servers
		"global_throttle_cache":         10,
	}
	defaultGlobalAuthRedirectParam = "rd"
)

const (
	maxAllowedLuaDictSize = 200
	maxNumberOfLuaDicts   = 100
)

// ReadConfig obtains the configuration defined by the user merged with the defaults.
func ReadConfig(src map[string]string) config.Configuration {
	conf := map[string]string{}
	// we need to copy the configmap data because the content is altered
	for k, v := range src {
		conf[k] = v
	}

	to := config.NewDefault()
	errors := make([]int, 0)
	skipUrls := make([]string, 0)
	whiteList := make([]string, 0)
	proxyList := make([]string, 0)
	hideHeadersList := make([]string, 0)

	bindAddressIpv4List := make([]string, 0)
	bindAddressIpv6List := make([]string, 0)

	blockCIDRList := make([]string, 0)
	blockUserAgentList := make([]string, 0)
	blockRefererList := make([]string, 0)
	responseHeaders := make([]string, 0)
	luaSharedDicts := make(map[string]int)

	//parse lua shared dict values
	if val, ok := conf[luaSharedDictsKey]; ok {
		delete(conf, luaSharedDictsKey)
		lsd := splitAndTrimSpace(val, ",")
		for _, v := range lsd {
			v = strings.Replace(v, " ", "", -1)
			results := strings.SplitN(v, ":", 2)
			dictName := results[0]
			size, err := strconv.Atoi(results[1])
			if err != nil {
				klog.Errorf("Ignoring non integer value %v for Lua dictionary %v: %v.", results[1], dictName, err)
				continue
			}
			if size > maxAllowedLuaDictSize {
				klog.Errorf("Ignoring %v for Lua dictionary %v: maximum size is %v.", size, dictName, maxAllowedLuaDictSize)
				continue
			}
			if len(luaSharedDicts)+1 > maxNumberOfLuaDicts {
				klog.Errorf("Ignoring %v for Lua dictionary %v: can not configure more than %v dictionaries.",
					size, dictName, maxNumberOfLuaDicts)
				continue
			}

			luaSharedDicts[dictName] = size
		}
	}
	// set default Lua shared dicts
	for k, v := range defaultLuaSharedDicts {
		if _, ok := luaSharedDicts[k]; !ok {
			luaSharedDicts[k] = v
		}
	}

	if val, ok := conf[customHTTPErrors]; ok {
		delete(conf, customHTTPErrors)
		for _, i := range splitAndTrimSpace(val, ",") {
			j, err := strconv.Atoi(i)
			if err != nil {
				klog.Warningf("%v is not a valid http code: %v", i, err)
			} else {
				errors = append(errors, j)
			}
		}
	}

	if val, ok := conf[hideHeaders]; ok {
		delete(conf, hideHeaders)
		hideHeadersList = splitAndTrimSpace(val, ",")
	}

	if val, ok := conf[skipAccessLogUrls]; ok {
		delete(conf, skipAccessLogUrls)
		skipUrls = splitAndTrimSpace(val, ",")
	}

	if val, ok := conf[whitelistSourceRange]; ok {
		delete(conf, whitelistSourceRange)
		whiteList = append(whiteList, splitAndTrimSpace(val, ",")...)
	}

	if val, ok := conf[proxyRealIPCIDR]; ok {
		delete(conf, proxyRealIPCIDR)
		proxyList = append(proxyList, splitAndTrimSpace(val, ",")...)
	} else {
		proxyList = append(proxyList, "0.0.0.0/0")
	}

	if val, ok := conf[bindAddress]; ok {
		delete(conf, bindAddress)
		for _, i := range splitAndTrimSpace(val, ",") {
			ns := net.ParseIP(i)
			if ns != nil {
				if ing_net.IsIPV6(ns) {
					bindAddressIpv6List = append(bindAddressIpv6List, fmt.Sprintf("[%v]", ns))
				} else {
					bindAddressIpv4List = append(bindAddressIpv4List, fmt.Sprintf("%v", ns))
				}
			} else {
				klog.Warningf("%v is not a valid textual representation of an IP address", i)
			}
		}
	}

	if val, ok := conf[blockCIDRs]; ok {
		delete(conf, blockCIDRs)
		blockCIDRList = splitAndTrimSpace(val, ",")
	}

	if val, ok := conf[blockUserAgents]; ok {
		delete(conf, blockUserAgents)
		blockUserAgentList = splitAndTrimSpace(val, ",")
	}

	if val, ok := conf[blockReferers]; ok {
		delete(conf, blockReferers)
		blockRefererList = splitAndTrimSpace(val, ",")
	}

	if val, ok := conf[httpRedirectCode]; ok {
		delete(conf, httpRedirectCode)
		j, err := strconv.Atoi(val)
		if err != nil {
			klog.Warningf("%v is not a valid HTTP code: %v", val, err)
		} else {
			if validRedirectCodes.Has(j) {
				to.HTTPRedirectCode = j
			} else {
				klog.Warningf("The code %v is not a valid as HTTP redirect code. Using the default.", val)
			}
		}
	}

	// Verify that the configured global external authorization URL is parsable as URL. if not, set the default value
	if val, ok := conf[globalAuthURL]; ok {
		delete(conf, globalAuthURL)

		authURL, message := parser.StringToURL(val)
		if authURL == nil {
			klog.Warningf("Global auth location denied - %v.", message)
		} else {
			to.GlobalExternalAuth.URL = val
			to.GlobalExternalAuth.Host = authURL.Hostname()
		}
	}

	// Verify that the configured global external authorization method is a valid HTTP method. if not, set the default value
	if val, ok := conf[globalAuthMethod]; ok {
		delete(conf, globalAuthMethod)

		if len(val) != 0 && !authreq.ValidMethod(val) {
			klog.Warningf("Global auth location denied - %v.", "invalid HTTP method")
		} else {
			to.GlobalExternalAuth.Method = val
		}
	}

	// Verify that the configured global external authorization error page is set and valid. if not, set the default value
	if val, ok := conf[globalAuthSignin]; ok {
		delete(conf, globalAuthSignin)

		signinURL, _ := parser.StringToURL(val)
		if signinURL == nil {
			klog.Warningf("Global auth location denied - %v.", "global-auth-signin setting is undefined and will not be set")
		} else {
			to.GlobalExternalAuth.SigninURL = val
		}
	}

	// Verify that the configured global external authorization error page redirection URL parameter is set and valid. if not, set the default value
	if val, ok := conf[globalAuthSigninRedirectParam]; ok {
		delete(conf, globalAuthSigninRedirectParam)

		redirectParam := strings.TrimSpace(val)
		dummySigninURL, _ := parser.StringToURL(fmt.Sprintf("%s?%s=dummy", to.GlobalExternalAuth.SigninURL, redirectParam))
		if dummySigninURL == nil {
			klog.Warningf("Global auth redirect parameter denied - %v.", "global-auth-signin-redirect-param setting is invalid and will not be set")
		} else {
			to.GlobalExternalAuth.SigninURLRedirectParam = redirectParam
		}
	}

	// Verify that the configured global external authorization response headers are valid. if not, set the default value
	if val, ok := conf[globalAuthResponseHeaders]; ok {
		delete(conf, globalAuthResponseHeaders)

		if len(val) != 0 {
			harr := splitAndTrimSpace(val, ",")
			for _, header := range harr {
				if !authreq.ValidHeader(header) {
					klog.Warningf("Global auth location denied - %v.", "invalid headers list")
				} else {
					responseHeaders = append(responseHeaders, header)
				}
			}
		}
		to.GlobalExternalAuth.ResponseHeaders = responseHeaders
	}

	if val, ok := conf[globalAuthRequestRedirect]; ok {
		delete(conf, globalAuthRequestRedirect)
		to.GlobalExternalAuth.RequestRedirect = val
	}

	if val, ok := conf[globalAuthSnippet]; ok {
		delete(conf, globalAuthSnippet)
		to.GlobalExternalAuth.AuthSnippet = val
	}

	if val, ok := conf[globalAuthCacheKey]; ok {
		delete(conf, globalAuthCacheKey)
		to.GlobalExternalAuth.AuthCacheKey = val
	}

	// Verify that the configured global external authorization cache duration is valid
	if val, ok := conf[globalAuthCacheDuration]; ok {
		delete(conf, globalAuthCacheDuration)

		cacheDurations, err := authreq.ParseStringToCacheDurations(val)
		if err != nil {
			klog.Warningf("Global auth location denied - %s", err)
		}
		to.GlobalExternalAuth.AuthCacheDuration = cacheDurations
	}

	if val, ok := conf[globalAuthTLSVerifyClient]; ok {
		delete(conf, globalAuthTLSVerifyClient)
		to.GlobalTLSAuth.AuthTLSVerifyClient = val
	}

	if val, ok := conf[globalAuthTLSVerifyDepth]; ok {
		delete(conf, globalAuthTLSVerifyDepth)
		i, err := strconv.Atoi(val)
		if err != nil {
			klog.Warningf("Invalid global-auth-tls-verify-depth %v. Setting to default value 1 instead.", val, err)
		} else {
			to.GlobalTLSAuth.AuthTLSVerifyDepth = i
		}
	}

	if val, ok := conf[globalAuthTLSSecret]; ok {
		delete(conf, globalAuthTLSSecret)
		to.GlobalTLSAuth.AuthTLSSecret = val
	}

	// Verify that the configured timeout is parsable as a duration. if not, set the default value
	if val, ok := conf[proxyHeaderTimeout]; ok {
		delete(conf, proxyHeaderTimeout)
		duration, err := time.ParseDuration(val)
		if err != nil {
			klog.Warningf("proxy-protocol-header-timeout of %v encountered an error while being parsed %v. Switching to use default value instead.", val, err)
		} else {
			to.ProxyProtocolHeaderTimeout = duration
		}
	}

	streamResponses := 1
	if val, ok := conf[proxyStreamResponses]; ok {
		delete(conf, proxyStreamResponses)
		j, err := strconv.Atoi(val)
		if err != nil {
			klog.Warningf("%v is not a valid number: %v", val, err)
		} else {
			streamResponses = j
		}
	}

	// Nginx Status whitelist
	if val, ok := conf[nginxStatusIpv4Whitelist]; ok {
		to.NginxStatusIpv4Whitelist = splitAndTrimSpace(val, ",")
		delete(conf, nginxStatusIpv4Whitelist)
	}

	if val, ok := conf[nginxStatusIpv6Whitelist]; ok {
		to.NginxStatusIpv6Whitelist = splitAndTrimSpace(val, ",")
		delete(conf, nginxStatusIpv6Whitelist)
	}

	if val, ok := conf[workerProcesses]; ok {
		to.WorkerProcesses = val
		if val == "auto" {
			to.WorkerProcesses = strconv.Itoa(runtime.NumCPU())
		}

		delete(conf, workerProcesses)
	}

	if val, ok := conf[plugins]; ok {
		to.Plugins = splitAndTrimSpace(val, ",")
		delete(conf, plugins)
	}

	to.CustomHTTPErrors = filterErrors(errors)
	to.SkipAccessLogURLs = skipUrls
	to.WhitelistSourceRange = whiteList
	to.ProxyRealIPCIDR = proxyList
	to.BindAddressIpv4 = bindAddressIpv4List
	to.BindAddressIpv6 = bindAddressIpv6List
	to.BlockCIDRs = blockCIDRList
	to.BlockUserAgents = blockUserAgentList
	to.BlockReferers = blockRefererList
	to.HideHeaders = hideHeadersList
	to.ProxyStreamResponses = streamResponses
	to.DisableIpv6DNS = !ing_net.IsIPv6Enabled()
	to.LuaSharedDicts = luaSharedDicts

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		WeaklyTypedInput: true,
		Result:           &to,
		TagName:          "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		klog.Warningf("unexpected error merging defaults: %v", err)
	}
	err = decoder.Decode(conf)
	if err != nil {
		klog.Warningf("unexpected error merging defaults: %v", err)
	}

	hash, err := hashstructure.Hash(to, &hashstructure.HashOptions{
		TagName: "json",
	})
	if err != nil {
		klog.Warningf("unexpected error obtaining hash: %v", err)
	}

	to.Checksum = fmt.Sprintf("%v", hash)

	return to
}

func filterErrors(codes []int) []int {
	var fa []int
	for _, code := range codes {
		if code > 299 && code < 600 {
			fa = append(fa, code)
		} else {
			klog.Warningf("error code %v is not valid for custom error pages", code)
		}
	}

	return fa
}

func splitAndTrimSpace(s, sep string) []string {
	f := func(c rune) bool {
		return strings.EqualFold(string(c), sep)
	}

	values := strings.FieldsFunc(s, f)
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}

	return values
}
