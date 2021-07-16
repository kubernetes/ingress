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

package defaults

import "net"

// Backend defines the mandatory configuration that an Ingress controller must provide
// The reason of this requirements is the annotations are generic. If some implementation do not supports
// one or more annotations it just can provides defaults
type Backend struct {
	// AppRoot contains the AppRoot for apps that doesn't exposes its content in the 'root' context
	AppRoot string `json:"app-root"`

	// enables which HTTP codes should be passed for processing with the error_page directive
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_intercept_errors
	// http://nginx.org/en/docs/http/ngx_http_core_module.html#error_page
	// By default this is disabled
	CustomHTTPErrors []int `json:"custom-http-errors"`

	// toggles whether or not to remove trailing slashes during TLS redirects
	PreserveTrailingSlash bool `json:"preserve-trailing-slash"`

	// http://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size
	// Sets the maximum allowed size of the client request body
	ProxyBodySize string `json:"proxy-body-size"`

	// Defines a timeout for establishing a connection with a proxied server.
	// It should be noted that this timeout cannot usually exceed 75 seconds.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_connect_timeout
	ProxyConnectTimeout int `json:"proxy-connect-timeout"`

	// Timeout in seconds for reading a response from the proxied server. The timeout is set only between
	// two successive read operations, not for the transmission of the whole response
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_read_timeout
	ProxyReadTimeout int `json:"proxy-read-timeout"`

	// Timeout in seconds for transmitting a request to the proxied server. The timeout is set only between
	// two successive write operations, not for the transmission of the whole request.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_send_timeout
	ProxySendTimeout int `json:"proxy-send-timeout"`

	// Sets the number of the buffers used for reading a response from the proxied server
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffers
	ProxyBuffersNumber int `json:"proxy-buffers-number"`

	// Sets the size of the buffer used for reading the first part of the response received from the
	// proxied server. This part usually contains a small response header.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffer_size)
	ProxyBufferSize string `json:"proxy-buffer-size"`

	// Sets a text that should be changed in the path attribute of the “Set-Cookie” header fields of
	// a proxied server response.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cookie_path
	ProxyCookiePath string `json:"proxy-cookie-path"`

	// Sets a text that should be changed in the domain attribute of the “Set-Cookie” header fields
	// of a proxied server response.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cookie_domain
	ProxyCookieDomain string `json:"proxy-cookie-domain"`

	// Specifies in which cases a request should be passed to the next server.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream
	ProxyNextUpstream string `json:"proxy-next-upstream"`

	// Limits the time during which a request can be passed to the next server.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream_timeout
	ProxyNextUpstreamTimeout int `json:"proxy-next-upstream-timeout"`

	// Limits the number of possible tries for passing a request to the next server.
	// https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream_tries
	ProxyNextUpstreamTries int `json:"proxy-next-upstream-tries"`

	// Sets the original text that should be changed in the "Location" and "Refresh" header fields of a proxied server response.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect
	// Default: off
	ProxyRedirectFrom string `json:"proxy-redirect-from"`

	// Sets the replacement text that should be changed in the "Location" and "Refresh" header fields of a proxied server response.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect
	// Default: off
	ProxyRedirectTo string `json:"proxy-redirect-to"`

	// Enables or disables buffering of a client request body.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_request_buffering
	ProxyRequestBuffering string `json:"proxy-request-buffering"`

	// Name server/s used to resolve names of upstream servers into IP addresses.
	// The file /etc/resolv.conf is used as DNS resolution configuration.
	Resolver []net.IP

	// SkipAccessLogURLs sets a list of URLs that should not appear in the NGINX access log
	// This is useful with urls like `/health` or `health-check` that make "complex" reading the logs
	// By default this list is empty
	SkipAccessLogURLs []string `json:"skip-access-log-urls"`

	// Enables or disables the redirect (301) to the HTTPS port
	SSLRedirect bool `json:"ssl-redirect"`

	// Enables or disables the redirect (301) to the HTTPS port even without TLS cert
	// This is useful if doing SSL offloading outside of cluster eg AWS ELB
	ForceSSLRedirect bool `json:"force-ssl-redirect"`

	// Enables or disables the specification of port in redirects
	// Default: false
	UsePortInRedirects bool `json:"use-port-in-redirects"`

	// Enable stickiness by client-server mapping based on a NGINX variable, text or a combination of both.
	// A consistent hashing method will be used which ensures only a few keys would be remapped to different
	// servers on upstream group changes
	// http://nginx.org/en/docs/http/ngx_http_upstream_module.html#hash
	UpstreamHashBy string `json:"upstream-hash-by"`

	// Consistent hashing subset flag.
	// Default: false
	UpstreamHashBySubset bool `json:"upstream-hash-by-subset"`

	// Subset consistent hashing, subset size.
	// Default 3
	UpstreamHashBySubsetSize int `json:"upstream-hash-by-subset-size"`

	// Let's us choose a load balancing algorithm per ingress
	LoadBalancing string `json:"load-balance"`

	// WhitelistSourceRange allows limiting access to certain client addresses
	// http://nginx.org/en/docs/http/ngx_http_access_module.html
	WhitelistSourceRange []string `json:"whitelist-source-range"`

	// BlocklistSourceRange denies access from certain client addresses
	// http://nginx.org/en/docs/http/ngx_http_access_module.html
	BlocklistSourceRange []string `json:"blocklist-source-range"`

	// Limits the rate of response transmission to a client.
	// The rate is specified in bytes per second. The zero value disables rate limiting.
	// The limit is set per a request, and so if a client simultaneously opens two connections,
	// the overall rate will be twice as much as the specified limit.
	// http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate
	LimitRate int `json:"limit-rate"`

	// Sets the initial amount after which the further transmission of a response to a client will be rate limited.
	// http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate_after
	LimitRateAfter int `json:"limit-rate-after"`

	// Enables or disables buffering of responses from the proxied server.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffering
	ProxyBuffering string `json:"proxy-buffering"`

	// Modifies the HTTP version the proxy uses to interact with the backend.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_http_version
	ProxyHTTPVersion string `json:"proxy-http-version"`

	// Sets the maximum temp file size when proxy-buffers capacity is exceeded.
	// http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_max_temp_file_size
	ProxyMaxTempFileSize string `json:"proxy-max-temp-file-size"`
}
