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

package ingress

import (
	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/ingress-nginx/internal/ingress/annotations"
	"k8s.io/ingress-nginx/internal/ingress/annotations/auth"
	"k8s.io/ingress-nginx/internal/ingress/annotations/authreq"
	"k8s.io/ingress-nginx/internal/ingress/annotations/authtls"
	"k8s.io/ingress-nginx/internal/ingress/annotations/connection"
	"k8s.io/ingress-nginx/internal/ingress/annotations/cors"
	"k8s.io/ingress-nginx/internal/ingress/annotations/fastcgi"
	"k8s.io/ingress-nginx/internal/ingress/annotations/globalratelimit"
	"k8s.io/ingress-nginx/internal/ingress/annotations/influxdb"
	"k8s.io/ingress-nginx/internal/ingress/annotations/ipwhitelist"
	"k8s.io/ingress-nginx/internal/ingress/annotations/log"
	"k8s.io/ingress-nginx/internal/ingress/annotations/mirror"
	"k8s.io/ingress-nginx/internal/ingress/annotations/modsecurity"
	"k8s.io/ingress-nginx/internal/ingress/annotations/opentracing"
	"k8s.io/ingress-nginx/internal/ingress/annotations/proxy"
	"k8s.io/ingress-nginx/internal/ingress/annotations/proxyssl"
	"k8s.io/ingress-nginx/internal/ingress/annotations/ratelimit"
	"k8s.io/ingress-nginx/internal/ingress/annotations/redirect"
	"k8s.io/ingress-nginx/internal/ingress/annotations/rewrite"
)

var (
	// DefaultSSLDirectory defines the location where the SSL certificates will be generated
	// This directory contains all the SSL certificates that are specified in Ingress rules.
	// The name of each file is <namespace>-<secret name>.pem. The content is the concatenated
	// certificate and key.
	DefaultSSLDirectory = "/ingress-controller/ssl"
)

// Configuration holds the definition of all the parts required to describe all
// ingresses reachable by the ingress controller (using a filter by namespace)
type Configuration struct {
	// Backends are a list of backends used by all the Ingress rules in the
	// ingress controller. This list includes the default backend
	Backends []*Backend `json:"backends,omitempty"`
	// Servers save the website config
	Servers []*Server `json:"servers,omitempty"`
	// TCPEndpoints contain endpoints for tcp streams handled by this backend
	// +optional
	TCPEndpoints []L4Service `json:"tcpEndpoints,omitempty"`
	// UDPEndpoints contain endpoints for udp streams handled by this backend
	// +optional
	UDPEndpoints []L4Service `json:"udpEndpoints,omitempty"`
	// PassthroughBackends contains the backends used for SSL passthrough.
	// It contains information about the associated Server Name Indication (SNI).
	// +optional
	PassthroughBackends []*SSLPassthroughBackend `json:"passthroughBackends,omitempty"`

	// BackendConfigChecksum contains the particular checksum of a Configuration object
	BackendConfigChecksum string `json:"BackendConfigChecksum,omitempty"`

	// ConfigurationChecksum contains the particular checksum of a Configuration object
	ConfigurationChecksum string `json:"configurationChecksum,omitempty"`

	DefaultSSLCertificate *SSLCert `json:"-"`
}

// Backend describes one or more remote server/s (endpoints) associated with a service
// +k8s:deepcopy-gen=true
type Backend struct {
	// Name represents an unique apiv1.Service name formatted as <namespace>-<name>-<port>
	Name    string             `json:"name"`
	Service *apiv1.Service     `json:"service,omitempty"`
	Port    intstr.IntOrString `json:"port"`
	// SSLPassthrough indicates that Ingress controller will delegate TLS termination to the endpoints.
	SSLPassthrough bool `json:"sslPassthrough"`
	// Endpoints contains the list of endpoints currently running
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// StickySessionAffinitySession contains the StickyConfig object with stickiness configuration
	SessionAffinity SessionAffinityConfig `json:"sessionAffinityConfig"`
	// Consistent hashing by NGINX variable
	UpstreamHashBy UpstreamHashByConfig `json:"upstreamHashByConfig,omitempty"`
	// LB algorithm configuration per ingress
	LoadBalancing string `json:"load-balance,omitempty"`
	// Denotes if a backend has no server. The backend instead shares a server with another backend and acts as an
	// alternative backend.
	// This can be used to share multiple upstreams in the sam nginx server block.
	NoServer bool `json:"noServer"`
	// Policies to describe the characteristics of an alternative backend.
	// +optional
	TrafficShapingPolicy TrafficShapingPolicy `json:"trafficShapingPolicy,omitempty"`
	// Contains a list of backends without servers that are associated with this backend.
	// +optional
	AlternativeBackends []string `json:"alternativeBackends,omitempty"`
}

// TrafficShapingPolicy describes the policies to put in place when a backend has no server and is used as an
// alternative backend
// +k8s:deepcopy-gen=true
type TrafficShapingPolicy struct {
	// Weight (0-100) of traffic to redirect to the backend.
	// e.g. Weight 20 means 20% of traffic will be redirected to the backend and 80% will remain
	// with the other backend. 0 weight will not send any traffic to this backend
	Weight int `json:"weight"`
	// Header on which to redirect requests to this backend
	Header string `json:"header"`
	// HeaderValue on which to redirect requests to this backend
	HeaderValue string `json:"headerValue"`
	// HeaderPattern the header value match pattern, support exact, regex.
	HeaderPattern string `json:"headerPattern"`
	// Cookie on which to redirect requests to this backend
	Cookie string `json:"cookie"`
	// Ip on which to redirect requests to this backend
	IpRange string `json:"ipRange"`
}

// HashInclude defines if a field should be used or not to calculate the hash
func (s Backend) HashInclude(field string, v interface{}) (bool, error) {
	switch field {
	case "Endpoints":
		return false, nil
	default:
		return true, nil
	}
}

// SessionAffinityConfig describes different affinity configurations for new sessions.
// Once a session is mapped to a backend based on some affinity setting, it
// retains that mapping till the backend goes down, or the ingress controller
// restarts. Exactly one of these values will be set on the upstream, since multiple
// affinity values are incompatible. Once set, the backend makes no guarantees
// about honoring updates.
// +k8s:deepcopy-gen=true
type SessionAffinityConfig struct {
	AffinityType          string                `json:"name"`
	AffinityMode          string                `json:"mode"`
	CookieSessionAffinity CookieSessionAffinity `json:"cookieSessionAffinity"`
}

// CookieSessionAffinity defines the structure used in Affinity configured by Cookies.
// +k8s:deepcopy-gen=true
type CookieSessionAffinity struct {
	Name                    string              `json:"name"`
	Expires                 string              `json:"expires,omitempty"`
	MaxAge                  string              `json:"maxage,omitempty"`
	Locations               map[string][]string `json:"locations,omitempty"`
	Path                    string              `json:"path,omitempty"`
	SameSite                string              `json:"samesite,omitempty"`
	ConditionalSameSiteNone bool                `json:"conditional_samesite_none,omitempty"`
	ChangeOnFailure         bool                `json:"change_on_failure,omitempty"`
}

// UpstreamHashByConfig described setting from the upstream-hash-by* annotations.
type UpstreamHashByConfig struct {
	UpstreamHashBy           string `json:"upstream-hash-by,omitempty"`
	UpstreamHashBySubset     bool   `json:"upstream-hash-by-subset,omitempty"`
	UpstreamHashBySubsetSize int    `json:"upstream-hash-by-subset-size,omitempty"`
}

// Endpoint describes a kubernetes endpoint in a backend
// +k8s:deepcopy-gen=true
type Endpoint struct {
	// Address IP address of the endpoint
	Address string `json:"address"`
	// Port number of the TCP port
	Port string `json:"port"`
	// Target returns a reference to the object providing the endpoint
	Target *apiv1.ObjectReference `json:"target,omitempty"`
}

// Server describes a website
type Server struct {
	// Hostname returns the FQDN of the server
	Hostname string `json:"hostname"`
	// SSLPassthrough indicates if the TLS termination is realized in
	// the server or in the remote endpoint
	SSLPassthrough bool `json:"sslPassthrough"`
	// SSLCert describes the certificate that will be used on the server
	SSLCert *SSLCert `json:"sslCert"`
	// Locations list of URIs configured in the server.
	Locations []*Location `json:"locations,omitempty"`
	// Aliases return the alias of the server name
	Aliases []string `json:"aliases,omitempty"`
	// RedirectFromToWWW returns if a redirect to/from prefix www is required
	RedirectFromToWWW bool `json:"redirectFromToWWW,omitempty"`
	// CertificateAuth indicates the this server requires mutual authentication
	// +optional
	CertificateAuth authtls.Config `json:"certificateAuth"`
	// ProxySSL indicates the this server uses client certificate to access backends
	// +optional
	ProxySSL proxyssl.Config `json:"proxySSL"`
	// ServerSnippet returns the snippet of server
	// +optional
	ServerSnippet string `json:"serverSnippet"`
	// SSLCiphers returns list of ciphers to be enabled
	SSLCiphers string `json:"sslCiphers,omitempty"`
	// SSLPreferServerCiphers indicates that server ciphers should be preferred
	// over client ciphers when using the SSLv3 and TLS protocols.
	SSLPreferServerCiphers string `json:"sslPreferServerCiphers,omitempty"`
	// AuthTLSError contains the reason why the access to a server should be denied
	AuthTLSError string `json:"authTLSError,omitempty"`
}

// Location describes an URI inside a server.
// Also contains additional information about annotations in the Ingress.
//
// In some cases when more than one annotations is defined a particular order in the execution
// is required.
// The chain in the execution order of annotations should be:
// - Whitelist
// - RateLimit
// - BasicDigestAuth
// - ExternalAuth
// - Redirect
type Location struct {
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1,
	// (i.e this follows the egrep/unix syntax, not the perl syntax)
	// matched against the path of an incoming request. Currently it can
	// contain characters disallowed from the conventional "path"
	// part of a URL as defined by RFC 3986. Paths must begin with
	// a '/'. If unspecified, the path defaults to a catch all sending
	// traffic to the backend.
	Path string `json:"path"`
	// PathType represents the type of path referred to by a HTTPIngressPath.
	PathType *networking.PathType `json:"pathType"`
	// IsDefBackend indicates if service specified in the Ingress
	// contains active endpoints or not. Returning true means the location
	// uses the default backend.
	IsDefBackend bool `json:"isDefBackend"`
	// Ingress returns the ingress from which this location was generated
	Ingress *Ingress `json:"ingress"`
	// IngressPath original path defined in the ingress rule
	IngressPath string `json:"ingressPath"`
	// Backend describes the name of the backend to use.
	Backend string `json:"backend"`
	// Service describes the referenced services from the ingress
	Service *apiv1.Service `json:"-"`
	// Port describes to which port from the service
	Port intstr.IntOrString `json:"port"`
	// Overwrite the Host header passed into the backend. Defaults to
	// vhost of the incoming request.
	// +optional
	UpstreamVhost string `json:"upstream-vhost"`
	// BasicDigestAuth returns authentication configuration for
	// an Ingress rule.
	// +optional
	BasicDigestAuth auth.Config `json:"basicDigestAuth,omitempty"`
	// Denied returns an error when this location cannot not be allowed
	// Requesting a denied location should return HTTP code 403.
	Denied *string `json:"denied,omitempty"`
	// CorsConfig returns the Cors Configuration for the ingress rule
	// +optional
	CorsConfig cors.Config `json:"corsConfig,omitempty"`
	// ExternalAuth indicates the access to this location requires
	// authentication using an external provider
	// +optional
	ExternalAuth authreq.Config `json:"externalAuth,omitempty"`
	// EnableGlobalAuth indicates if the access to this location requires
	// authentication using an external provider defined in controller's config
	EnableGlobalAuth bool `json:"enableGlobalAuth"`
	// HTTP2PushPreload allows to configure the HTTP2 Push Preload from backend
	// original location.
	// +optional
	HTTP2PushPreload bool `json:"http2PushPreload,omitempty"`
	// RateLimit describes a limit in the number of connections per IP
	// address or connections per second.
	// The Redirect annotation precedes RateLimit
	// +optional
	RateLimit ratelimit.Config `json:"rateLimit,omitempty"`
	// GlobalRateLimit similar to RateLimit
	// but this is applied globally across multiple replicas.
	// +optional
	GlobalRateLimit globalratelimit.Config `json:"globalRateLimit,omitempty"`
	// Redirect describes a temporal o permanent redirection this location.
	// +optional
	Redirect redirect.Config `json:"redirect,omitempty"`
	// Rewrite describes the redirection this location.
	// +optional
	Rewrite rewrite.Config `json:"rewrite,omitempty"`
	// Whitelist indicates only connections from certain client
	// addresses or networks are allowed.
	// +optional
	Whitelist ipwhitelist.SourceRange `json:"whitelist,omitempty"`
	// Proxy contains information about timeouts and buffer sizes
	// to be used in connections against endpoints
	// +optional
	Proxy proxy.Config `json:"proxy,omitempty"`
	// ProxySSL contains information about SSL configuration parameters
	// to be used in connections against endpoints
	// +optional
	ProxySSL proxyssl.Config `json:"proxySSL,omitempty"`
	// UsePortInRedirects indicates if redirects must specify the port
	// +optional
	UsePortInRedirects bool `json:"usePortInRedirects"`
	// ConfigurationSnippet contains additional configuration for the backend
	// to be considered in the configuration of the location
	ConfigurationSnippet string `json:"configurationSnippet"`
	// Connection contains connection header to override the default Connection header
	// to the request.
	// +optional
	Connection connection.Config `json:"connection"`
	// ClientBodyBufferSize allows for the configuration of the client body
	// buffer size for a specific location.
	// +optional
	ClientBodyBufferSize string `json:"clientBodyBufferSize,omitempty"`
	// DefaultBackend allows the use of a custom default backend for this location.
	// +optional
	DefaultBackend *apiv1.Service `json:"-"`
	// DefaultBackendUpstreamName is the upstream-formatted string for the name of
	// this location's custom default backend
	DefaultBackendUpstreamName string `json:"defaultBackendUpstreamName,omitempty"`
	// XForwardedPrefix allows to add a header X-Forwarded-Prefix to the request with the
	// original location.
	// +optional
	XForwardedPrefix string `json:"xForwardedPrefix,omitempty"`
	// Logs allows to enable or disable the nginx logs
	// By default access logs are enabled and rewrite logs are disabled
	Logs log.Config `json:"logs,omitempty"`
	// InfluxDB allows to monitor the incoming request by sending them to an influxdb database
	// +optional
	InfluxDB influxdb.Config `json:"influxDB,omitempty"`
	// BackendProtocol indicates which protocol should be used to communicate with the service
	// By default this is HTTP
	BackendProtocol string `json:"backend-protocol"`
	// FastCGI allows the ingress to act as a FastCGI client for a given location.
	// +optional
	FastCGI fastcgi.Config `json:"fastcgi,omitempty"`
	// CustomHTTPErrors specifies the error codes that should be intercepted.
	// +optional
	CustomHTTPErrors []int `json:"custom-http-errors"`
	// ModSecurity allows to enable and configure modsecurity
	// +optional
	ModSecurity modsecurity.Config `json:"modsecurity"`
	// Satisfy dictates allow access if any or all is set
	Satisfy string `json:"satisfy"`
	// Mirror allows you to mirror traffic to a "test" backend
	// +optional
	Mirror mirror.Config `json:"mirror,omitempty"`
	// Opentracing allows the global opentracing setting to be overridden for a location
	// +optional
	Opentracing opentracing.Config `json:"opentracing"`
}

// SSLPassthroughBackend describes a SSL upstream server configured
// as passthrough (no TLS termination in the ingress controller)
// The endpoints must provide the TLS termination exposing the required SSL certificate.
// The ingress controller only pipes the underlying TCP connection
type SSLPassthroughBackend struct {
	Service *apiv1.Service     `json:"-"`
	Port    intstr.IntOrString `json:"port"`
	// Backend describes the endpoints to use.
	Backend string `json:"namespace,omitempty"`
	// Hostname returns the FQDN of the server
	Hostname string `json:"hostname"`
}

// L4Service describes a L4 Ingress service.
type L4Service struct {
	// Port external port to expose
	Port int `json:"port"`
	// Backend of the service
	Backend L4Backend `json:"backend"`
	// Endpoints active endpoints of the service
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// k8s Service
	Service *apiv1.Service `json:"-"`
}

// L4Backend describes the kubernetes service behind L4 Ingress service
type L4Backend struct {
	Port      intstr.IntOrString `json:"port"`
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Protocol  apiv1.Protocol     `json:"protocol"`
	// +optional
	ProxyProtocol ProxyProtocol `json:"proxyProtocol"`
}

// ProxyProtocol describes the proxy protocol configuration
type ProxyProtocol struct {
	Decode bool `json:"decode"`
	Encode bool `json:"encode"`
}

// Ingress holds the definition of an Ingress plus its annotations
type Ingress struct {
	networking.Ingress `json:"-"`
	ParsedAnnotations  *annotations.Ingress `json:"parsedAnnotations"`
}

// GeneralConfig holds the definition of lua general configuration data
type GeneralConfig struct {
}
