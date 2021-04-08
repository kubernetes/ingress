# ConfigMaps

ConfigMaps allow you to decouple configuration artifacts from image content to keep containerized applications portable.

The ConfigMap API resource stores configuration data as key-value pairs. The data provides the configurations for system
components for the nginx-controller.

In order to overwrite nginx-controller configuration values as seen in [config.go](https://github.com/kubernetes/ingress-nginx/blob/master/internal/ingress/controller/config/config.go),
you can add key-value pairs to the data section of the config-map. For Example:

```yaml
data:
  map-hash-bucket-size: "128"
  ssl-protocols: SSLv2
```

!!! important
    The key and values in a ConfigMap can only be strings.
    This means that we want a value with boolean values we need to quote the values, like "true" or "false".
    Same for numbers, like "100".

    "Slice" types (defined below as `[]string` or `[]int`) can be provided as a comma-delimited string.

## Configuration options

The following table shows a configuration option's name, type, and the default value:

|name|type|default|
|:---|:---|:------|
|[add-headers](#add-headers)|string|""|
|[allow-backend-server-header](#allow-backend-server-header)|bool|"false"|
|[hide-headers](#hide-headers)|string array|empty|
|[access-log-params](#access-log-params)|string|""|
|[access-log-path](#access-log-path)|string|"/var/log/nginx/access.log"|
|[http-access-log-path](#http-access-log-path)|string|""|
|[stream-access-log-path](#stream-access-log-path)|string|""|
|[enable-access-log-for-default-backend](#enable-access-log-for-default-backend)|bool|"false"|
|[error-log-path](#error-log-path)|string|"/var/log/nginx/error.log"|
|[enable-modsecurity](#enable-modsecurity)|bool|"false"|
|[modsecurity-snippet](#modsecurity-snippet)|string|""|
|[enable-owasp-modsecurity-crs](#enable-owasp-modsecurity-crs)|bool|"false"|
|[client-header-buffer-size](#client-header-buffer-size)|string|"1k"|
|[client-header-timeout](#client-header-timeout)|int|60|
|[client-body-buffer-size](#client-body-buffer-size)|string|"8k"|
|[client-body-timeout](#client-body-timeout)|int|60|
|[disable-access-log](#disable-access-log)|bool|false|
|[disable-ipv6](#disable-ipv6)|bool|false|
|[disable-ipv6-dns](#disable-ipv6-dns)|bool|false|
|[enable-underscores-in-headers](#enable-underscores-in-headers)|bool|false|
|[enable-ocsp](#enable-ocsp)|bool|false|
|[ignore-invalid-headers](#ignore-invalid-headers)|bool|true|
|[retry-non-idempotent](#retry-non-idempotent)|bool|"false"|
|[error-log-level](#error-log-level)|string|"notice"|
|[http2-max-field-size](#http2-max-field-size)|string|"4k"|
|[http2-max-header-size](#http2-max-header-size)|string|"16k"|
|[http2-max-requests](#http2-max-requests)|int|1000|
|[http2-max-concurrent-streams](#http2-max-concurrent-streams)|int|1000|
|[hsts](#hsts)|bool|"true"|
|[hsts-include-subdomains](#hsts-include-subdomains)|bool|"true"|
|[hsts-max-age](#hsts-max-age)|string|"15724800"|
|[hsts-preload](#hsts-preload)|bool|"false"|
|[keep-alive](#keep-alive)|int|75|
|[keep-alive-requests](#keep-alive-requests)|int|100|
|[large-client-header-buffers](#large-client-header-buffers)|string|"4 8k"|
|[log-format-escape-json](#log-format-escape-json)|bool|"false"|
|[log-format-upstream](#log-format-upstream)|string|`$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" $request_length $request_time [$proxy_upstream_name] [$proxy_alternative_upstream_name] $upstream_addr $upstream_response_length $upstream_response_time $upstream_status $req_id`|
|[log-format-stream](#log-format-stream)|string|`[$remote_addr] [$time_local] $protocol $status $bytes_sent $bytes_received $session_time`|
|[enable-multi-accept](#enable-multi-accept)|bool|"true"|
|[max-worker-connections](#max-worker-connections)|int|16384|
|[max-worker-open-files](#max-worker-open-files)|int|0|
|[map-hash-bucket-size](#max-hash-bucket-size)|int|64|
|[nginx-status-ipv4-whitelist](#nginx-status-ipv4-whitelist)|[]string|"127.0.0.1"|
|[nginx-status-ipv6-whitelist](#nginx-status-ipv6-whitelist)|[]string|"::1"|
|[proxy-real-ip-cidr](#proxy-real-ip-cidr)|[]string|"0.0.0.0/0"|
|[proxy-set-headers](#proxy-set-headers)|string|""|
|[server-name-hash-max-size](#server-name-hash-max-size)|int|1024|
|[server-name-hash-bucket-size](#server-name-hash-bucket-size)|int|`<size of the processor’s cache line>`
|[proxy-headers-hash-max-size](#proxy-headers-hash-max-size)|int|512|
|[proxy-headers-hash-bucket-size](#proxy-headers-hash-bucket-size)|int|64|
|[plugins](#plugins)|[]string| |
|[reuse-port](#reuse-port)|bool|"true"|
|[server-tokens](#server-tokens)|bool|"false"|
|[ssl-ciphers](#ssl-ciphers)|string|"ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384"|
|[ssl-ecdh-curve](#ssl-ecdh-curve)|string|"auto"|
|[ssl-dh-param](#ssl-dh-param)|string|""|
|[ssl-protocols](#ssl-protocols)|string|"TLSv1.2 TLSv1.3"|
|[ssl-session-cache](#ssl-session-cache)|bool|"true"|
|[ssl-session-cache-size](#ssl-session-cache-size)|string|"10m"|
|[ssl-session-tickets](#ssl-session-tickets)|bool|"false"|
|[ssl-session-ticket-key](#ssl-session-ticket-key)|string|`<Randomly Generated>`
|[ssl-session-timeout](#ssl-session-timeout)|string|"10m"|
|[ssl-buffer-size](#ssl-buffer-size)|string|"4k"|
|[use-proxy-protocol](#use-proxy-protocol)|bool|"false"|
|[proxy-protocol-header-timeout](#proxy-protocol-header-timeout)|string|"5s"|
|[use-gzip](#use-gzip)|bool|"false"|
|[use-geoip](#use-geoip)|bool|"true"|
|[use-geoip2](#use-geoip2)|bool|"false"|
|[enable-brotli](#enable-brotli)|bool|"false"|
|[brotli-level](#brotli-level)|int|4|
|[brotli-types](#brotli-types)|string|"application/xml+rss application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/javascript text/plain text/x-component"|
|[use-http2](#use-http2)|bool|"true"|
|[gzip-level](#gzip-level)|int|1|
|[gzip-types](#gzip-types)|string|"application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/javascript text/plain text/x-component"|
|[worker-processes](#worker-processes)|string|`<Number of CPUs>`|
|[worker-cpu-affinity](#worker-cpu-affinity)|string|""|
|[worker-shutdown-timeout](#worker-shutdown-timeout)|string|"240s"|
|[load-balance](#load-balance)|string|"round_robin"|
|[variables-hash-bucket-size](#variables-hash-bucket-size)|int|128|
|[variables-hash-max-size](#variables-hash-max-size)|int|2048|
|[upstream-keepalive-connections](#upstream-keepalive-connections)|int|320|
|[upstream-keepalive-timeout](#upstream-keepalive-timeout)|int|60|
|[upstream-keepalive-requests](#upstream-keepalive-requests)|int|10000|
|[limit-conn-zone-variable](#limit-conn-zone-variable)|string|"$binary_remote_addr"|
|[proxy-stream-timeout](#proxy-stream-timeout)|string|"600s"|
|[proxy-stream-next-upstream](#proxy-stream-next-upstream)|bool|"true"|
|[proxy-stream-next-upstream-timeout](#proxy-stream-next-upstream-timeout)|string|"600s"|
|[proxy-stream-next-upstream-tries](#proxy-stream-next-upstream-tries)|int|3|
|[proxy-stream-responses](#proxy-stream-responses)|int|1|
|[bind-address](#bind-address)|[]string|""|
|[use-forwarded-headers](#use-forwarded-headers)|bool|"false"|
|[enable-real-ip](#enable-real-ip)|bool|"false"|
|[forwarded-for-header](#forwarded-for-header)|string|"X-Forwarded-For"|
|[compute-full-forwarded-for](#compute-full-forwarded-for)|bool|"false"|
|[proxy-add-original-uri-header](#proxy-add-original-uri-header)|bool|"false"|
|[generate-request-id](#generate-request-id)|bool|"true"|
|[enable-opentracing](#enable-opentracing)|bool|"false"|
|[opentracing-operation-name](#opentracing-operation-name)|string|""|
|[opentracing-location-operation-name](#opentracing-location-operation-name)|string|""|
|[zipkin-collector-host](#zipkin-collector-host)|string|""|
|[zipkin-collector-port](#zipkin-collector-port)|int|9411|
|[zipkin-service-name](#zipkin-service-name)|string|"nginx"|
|[zipkin-sample-rate](#zipkin-sample-rate)|float|1.0|
|[jaeger-collector-host](#jaeger-collector-host)|string|""|
|[jaeger-collector-port](#jaeger-collector-port)|int|6831|
|[jaeger-endpoint](#jaeger-endpoint)|string|""|
|[jaeger-service-name](#jaeger-service-name)|string|"nginx"|
|[jaeger-propagation-format](#jaeger-propagation-format)|string|"jaeger"|
|[jaeger-sampler-type](#jaeger-sampler-type)|string|"const"|
|[jaeger-sampler-param](#jaeger-sampler-param)|string|"1"|
|[jaeger-sampler-host](#jaeger-sampler-host)|string|"http://127.0.0.1"|
|[jaeger-sampler-port](#jaeger-sampler-port)|int|5778|
|[jaeger-trace-context-header-name](#jaeger-trace-context-header-name)|string|uber-trace-id|
|[jaeger-debug-header](#jaeger-debug-header)|string|uber-debug-id|
|[jaeger-baggage-header](#jaeger-baggage-header)|string|jaeger-baggage|
|[jaeger-trace-baggage-header-prefix](#jaeger-trace-baggage-header-prefix)|string|uberctx-|
|[datadog-collector-host](#datadog-collector-host)|string|""|
|[datadog-collector-port](#datadog-collector-port)|int|8126|
|[datadog-service-name](#datadog-service-name)|string|"nginx"|
|[datadog-environment](#datadog-environment)|string|"prod"|
|[datadog-operation-name-override](#datadog-operation-name-override)|string|"nginx.handle"|
|[datadog-priority-sampling](#datadog-priority-sampling)|bool|"true"|
|[datadog-sample-rate](#datadog-sample-rate)|float|1.0|
|[main-snippet](#main-snippet)|string|""|
|[http-snippet](#http-snippet)|string|""|
|[server-snippet](#server-snippet)|string|""|
|[location-snippet](#location-snippet)|string|""|
|[custom-http-errors](#custom-http-errors)|[]int|[]int{}|
|[proxy-body-size](#proxy-body-size)|string|"1m"|
|[proxy-connect-timeout](#proxy-connect-timeout)|int|5|
|[proxy-read-timeout](#proxy-read-timeout)|int|60|
|[proxy-send-timeout](#proxy-send-timeout)|int|60|
|[proxy-buffers-number](#proxy-buffers-number)|int|4|
|[proxy-buffer-size](#proxy-buffer-size)|string|"4k"|
|[proxy-cookie-path](#proxy-cookie-path)|string|"off"|
|[proxy-cookie-domain](#proxy-cookie-domain)|string|"off"|
|[proxy-next-upstream](#proxy-next-upstream)|string|"error timeout"|
|[proxy-next-upstream-timeout](#proxy-next-upstream-timeout)|int|0|
|[proxy-next-upstream-tries](#proxy-next-upstream-tries)|int|3|
|[proxy-redirect-from](#proxy-redirect-from)|string|"off"|
|[proxy-request-buffering](#proxy-request-buffering)|string|"on"|
|[ssl-redirect](#ssl-redirect)|bool|"true"|
|[whitelist-source-range](#whitelist-source-range)|[]string|[]string{}|
|[skip-access-log-urls](#skip-access-log-urls)|[]string|[]string{}|
|[limit-rate](#limit-rate)|int|0|
|[limit-rate-after](#limit-rate-after)|int|0|
|[lua-shared-dicts](#lua-shared-dicts)|string|""|
|[http-redirect-code](#http-redirect-code)|int|308|
|[proxy-buffering](#proxy-buffering)|string|"off"|
|[limit-req-status-code](#limit-req-status-code)|int|503|
|[limit-conn-status-code](#limit-conn-status-code)|int|503|
|[no-tls-redirect-locations](#no-tls-redirect-locations)|string|"/.well-known/acme-challenge"|
|[global-auth-url](#global-auth-url)|string|""|
|[global-auth-method](#global-auth-method)|string|""|
|[global-auth-signin](#global-auth-signin)|string|""|
|[global-auth-signin-redirect-param](#global-auth-signin-redirect-param)|string|"rd"|
|[global-auth-response-headers](#global-auth-response-headers)|string|""|
|[global-auth-request-redirect](#global-auth-request-redirect)|string|""|
|[global-auth-snippet](#global-auth-snippet)|string|""|
|[global-auth-cache-key](#global-auth-cache-key)|string|""|
|[global-auth-cache-duration](#global-auth-cache-duration)|string|"200 202 401 5m"|
|[no-auth-locations](#no-auth-locations)|string|"/.well-known/acme-challenge"|
|[block-cidrs](#block-cidrs)|[]string|""|
|[block-user-agents](#block-user-agents)|[]string|""|
|[block-referers](#block-referers)|[]string|""|
|[proxy-ssl-location-only](#proxy-ssl-location-only)|bool|"false"|
|[default-type](#default-type)|string|"text/html"|
|[global-rate-limit-memcached-host](#global-rate-limit)|string|""|
|[global-rate-limit-memcached-port](#global-rate-limit)|int|11211|
|[global-rate-limit-memcached-connect-timeout](#global-rate-limit)|int|50|
|[global-rate-limit-memcached-max-idle-timeout](#global-rate-limit)|int|10000|
|[global-rate-limit-memcached-pool-size](#global-rate-limit)|int|50|
|[global-rate-limit-status-code](#global-rate-limit)|int|429|

## add-headers

Sets custom headers from named configmap before sending traffic to the client. See [proxy-set-headers](#proxy-set-headers). [example](https://github.com/kubernetes/ingress-nginx/tree/master/docs/examples/customization/custom-headers)

## allow-backend-server-header

Enables the return of the header Server from the backend instead of the generic nginx string. _**default:**_ is disabled

## hide-headers

Sets additional header that will not be passed from the upstream server to the client response.
_**default:**_ empty

_References:_
[http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_hide_header](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_hide_header)

## access-log-params

Additional params for access_log. For example, buffer=16k, gzip, flush=1m

_References:_
[http://nginx.org/en/docs/http/ngx_http_log_module.html#access_log](http://nginx.org/en/docs/http/ngx_http_log_module.html#access_log)

## access-log-path

Access log path for both http and stream context. Goes to `/var/log/nginx/access.log` by default.

__Note:__ the file `/var/log/nginx/access.log` is a symlink to `/dev/stdout`

## http-access-log-path

Access log path for http context globally.
_**default:**_ ""

__Note:__ If not specified, the `access-log-path` will be used.

## stream-access-log-path

Access log path for stream context globally.
_**default:**_ ""

__Note:__ If not specified, the `access-log-path` will be used.

## enable-access-log-for-default-backend

Enables logging access to default backend. _**default:**_ is disabled.

## error-log-path

Error log path. Goes to `/var/log/nginx/error.log` by default.

__Note:__ the file `/var/log/nginx/error.log` is a symlink to `/dev/stderr`

_References:_
[http://nginx.org/en/docs/ngx_core_module.html#error_log](http://nginx.org/en/docs/ngx_core_module.html#error_log)

## enable-modsecurity

Enables the modsecurity module for NGINX. _**default:**_ is disabled

## enable-owasp-modsecurity-crs

Enables the OWASP ModSecurity Core Rule Set (CRS). _**default:**_ is disabled

## modsecurity-snippet

Adds custom rules to modsecurity section of nginx configuration

## client-header-buffer-size

Allows to configure a custom buffer size for reading client request header.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#client_header_buffer_size](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_header_buffer_size)

## client-header-timeout

Defines a timeout for reading client request header, in seconds.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#client_header_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_header_timeout)

## client-body-buffer-size

Sets buffer size for reading client request body.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_buffer_size](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_buffer_size)

## client-body-timeout

Defines a timeout for reading client request body, in seconds.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout)

## disable-access-log

Disables the Access Log from the entire Ingress Controller. _**default:**_ '"false"'

_References:_
[http://nginx.org/en/docs/http/ngx_http_log_module.html#access_log](http://nginx.org/en/docs/http/ngx_http_log_module.html#access_log)

## disable-ipv6

Disable listening on IPV6. _**default:**_ `false`; IPv6 listening is enabled

## disable-ipv6-dns

Disable IPV6 for nginx DNS resolver. _**default:**_ `false`; IPv6 resolving enabled.

## enable-underscores-in-headers

Enables underscores in header names. _**default:**_ is disabled

## enable-ocsp

Enables [Online Certificate Status Protocol stapling](https://en.wikipedia.org/wiki/OCSP_stapling) (OCSP) support.
_**default:**_ is disabled

## ignore-invalid-headers

Set if header fields with invalid names should be ignored.
_**default:**_ is enabled

## retry-non-idempotent

Since 1.9.13 NGINX will not retry non-idempotent requests (POST, LOCK, PATCH) in case of an error in the upstream server. The previous behavior can be restored using the value "true".

## error-log-level

Configures the logging level of errors. Log levels above are listed in the order of increasing severity.

_References:_
[http://nginx.org/en/docs/ngx_core_module.html#error_log](http://nginx.org/en/docs/ngx_core_module.html#error_log)

## http2-max-field-size

Limits the maximum size of an HPACK-compressed request header field.

_References:_
[https://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_field_size](https://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_field_size)

## http2-max-header-size

Limits the maximum size of the entire request header list after HPACK decompression.

_References:_
[https://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_header_size](https://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_header_size)

## http2-max-requests

Sets the maximum number of requests (including push requests) that can be served through one HTTP/2 connection, after which the next client request will lead to connection closing and the need of establishing a new connection.

_References:_
[http://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_requests](http://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_requests)

## http2-max-concurrent-streams

Sets the maximum number of concurrent HTTP/2 streams in a connection.

_References:_
[http://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_concurrent_streams](http://nginx.org/en/docs/http/ngx_http_v2_module.html#http2_max_concurrent_streams)

## hsts

Enables or disables the header HSTS in servers running SSL.
HTTP Strict Transport Security (often abbreviated as HSTS) is a security feature (HTTP header) that tell browsers that it should only be communicated with using HTTPS, instead of using HTTP. It provides protection against protocol downgrade attacks and cookie theft.

_References:_

- [https://developer.mozilla.org/en-US/docs/Web/Security/HTTP_strict_transport_security](https://developer.mozilla.org/en-US/docs/Web/Security/HTTP_strict_transport_security)
- [https://blog.qualys.com/securitylabs/2016/03/28/the-importance-of-a-proper-http-strict-transport-security-implementation-on-your-web-server](https://blog.qualys.com/securitylabs/2016/03/28/the-importance-of-a-proper-http-strict-transport-security-implementation-on-your-web-server)

## hsts-include-subdomains

Enables or disables the use of HSTS in all the subdomains of the server-name.

## hsts-max-age

Sets the time, in seconds, that the browser should remember that this site is only to be accessed using HTTPS.

## hsts-preload

Enables or disables the preload attribute in the HSTS feature (when it is enabled) dd

## keep-alive

Sets the time during which a keep-alive client connection will stay open on the server side. The zero value disables keep-alive client connections.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_timeout](http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_timeout)

## keep-alive-requests

Sets the maximum number of requests that can be served through one keep-alive connection.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests](http://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests)

## large-client-header-buffers

Sets the maximum number and size of buffers used for reading large client request header. _**default:**_ 4 8k

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#large_client_header_buffers](http://nginx.org/en/docs/http/ngx_http_core_module.html#large_client_header_buffers)

## log-format-escape-json

Sets if the escape parameter allows JSON ("true") or default characters escaping in variables ("false") Sets the nginx [log format](http://nginx.org/en/docs/http/ngx_http_log_module.html#log_format).

## log-format-upstream

Sets the nginx [log format](http://nginx.org/en/docs/http/ngx_http_log_module.html#log_format).
Example for json output:

```json

log-format-upstream: '{"time": "$time_iso8601", "remote_addr": "$proxy_protocol_addr", "x_forward_for": "$proxy_add_x_forwarded_for", "request_id": "$req_id",
  "remote_user": "$remote_user", "bytes_sent": $bytes_sent, "request_time": $request_time, "status": $status, "vhost": "$host", "request_proto": "$server_protocol",
  "path": "$uri", "request_query": "$args", "request_length": $request_length, "duration": $request_time,"method": "$request_method", "http_referrer": "$http_referer",
  "http_user_agent": "$http_user_agent" }'
```

Please check the [log-format](log-format.md) for definition of each field.

## log-format-stream

Sets the nginx [stream format](https://nginx.org/en/docs/stream/ngx_stream_log_module.html#log_format).

## enable-multi-accept

If disabled, a worker process will accept one new connection at a time. Otherwise, a worker process will accept all new connections at a time.
_**default:**_ true

_References:_
[http://nginx.org/en/docs/ngx_core_module.html#multi_accept](http://nginx.org/en/docs/ngx_core_module.html#multi_accept)

## max-worker-connections

Sets the [maximum number of simultaneous connections](http://nginx.org/en/docs/ngx_core_module.html#worker_connections) that can be opened by each worker process.
0 will use the value of [max-worker-open-files](#max-worker-open-files).
_**default:**_ 16384

!!! tip
    Using 0 in scenarios of high load improves performance at the cost of increasing RAM utilization (even on idle).

## max-worker-open-files

Sets the [maximum number of files](http://nginx.org/en/docs/ngx_core_module.html#worker_rlimit_nofile) that can be opened by each worker process.
The default of 0 means "max open files (system's limit) / [worker-processes](#worker-processes) - 1024".
_**default:**_ 0

## map-hash-bucket-size

Sets the bucket size for the [map variables hash tables](http://nginx.org/en/docs/http/ngx_http_map_module.html#map_hash_bucket_size). The details of setting up hash tables are provided in a separate [document](http://nginx.org/en/docs/hash.html).

## proxy-real-ip-cidr

If use-forwarded-headers or use-proxy-protocol is enabled, proxy-real-ip-cidr defines the default the IP/network address of your external load balancer.

## proxy-set-headers

Sets custom headers from named configmap before sending traffic to backends. The value format is namespace/name.  See [example](https://github.com/kubernetes/ingress-nginx/tree/master/docs/examples/customization/custom-headers)

## server-name-hash-max-size

Sets the maximum size of the [server names hash tables](http://nginx.org/en/docs/http/ngx_http_core_module.html#server_names_hash_max_size) used in server names,map directive’s values, MIME types, names of request header strings, etc.

_References:_
[http://nginx.org/en/docs/hash.html](http://nginx.org/en/docs/hash.html)

## server-name-hash-bucket-size

Sets the size of the bucket for the server names hash tables.

_References:_

- [http://nginx.org/en/docs/hash.html](http://nginx.org/en/docs/hash.html)
- [http://nginx.org/en/docs/http/ngx_http_core_module.html#server_names_hash_bucket_size](http://nginx.org/en/docs/http/ngx_http_core_module.html#server_names_hash_bucket_size)

## proxy-headers-hash-max-size

Sets the maximum size of the proxy headers hash tables.

_References:_

- [http://nginx.org/en/docs/hash.html](http://nginx.org/en/docs/hash.html)
- [https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_headers_hash_max_size](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_headers_hash_max_size)

## reuse-port

Instructs NGINX to create an individual listening socket for each worker process (using the SO_REUSEPORT socket option), allowing a kernel to distribute incoming connections between worker processes
_**default:**_ true

## proxy-headers-hash-bucket-size

Sets the size of the bucket for the proxy headers hash tables.

_References:_

- [http://nginx.org/en/docs/hash.html](http://nginx.org/en/docs/hash.html)
- [https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_headers_hash_bucket_size](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_headers_hash_bucket_size)

## plugins

Activates plugins installed in `/etc/nginx/lua/plugins`. Refer to [ingress-nginx plugins README](https://github.com/kubernetes/ingress-nginx/blob/master/rootfs/etc/nginx/lua/plugins/README.md) for more information on how to write and install a plugin.

## server-tokens

Send NGINX Server header in responses and display NGINX version in error pages. _**default:**_ is disabled

## ssl-ciphers

Sets the [ciphers](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_ciphers) list to enable. The ciphers are specified in the format understood by the OpenSSL library.

The default cipher list is:
 `ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384`.

The ordering of a ciphersuite is very important because it decides which algorithms are going to be selected in priority. The recommendation above prioritizes algorithms that provide perfect [forward secrecy](https://wiki.mozilla.org/Security/Server_Side_TLS#Forward_Secrecy).

DHE-based cyphers will not be available until DH parameter is configured [Custom DH parameters for perfect forward secrecy](https://github.com/kubernetes/ingress-nginx/tree/master/docs/examples/customization/ssl-dh-param)

Please check the [Mozilla SSL Configuration Generator](https://mozilla.github.io/server-side-tls/ssl-config-generator/).

__Note:__ ssl_prefer_server_ciphers directive will be enabled by default for http context.

## ssl-ecdh-curve

Specifies a curve for ECDHE ciphers.

_References:_
[http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_ecdh_curve](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_ecdh_curve)

## ssl-dh-param

Sets the name of the secret that contains Diffie-Hellman key to help with "Perfect Forward Secrecy".

_References:_

- [https://wiki.openssl.org/index.php/Diffie-Hellman_parameters](https://wiki.openssl.org/index.php/Diffie-Hellman_parameters)
- [https://wiki.mozilla.org/Security/Server_Side_TLS#DHE_handshake_and_dhparam](https://wiki.mozilla.org/Security/Server_Side_TLS#DHE_handshake_and_dhparam)
- [http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_dhparam](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_dhparam)

## ssl-protocols

Sets the [SSL protocols](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_protocols) to use. The default is: `TLSv1.2 TLSv1.3`.

Please check the result of the configuration using `https://ssllabs.com/ssltest/analyze.html` or `https://testssl.sh`.

## ssl-early-data

Enables or disables TLS 1.3 [early data](https://tools.ietf.org/html/rfc8446#section-2.3), also known as Zero Round Trip
Time Resumption (0-RTT).

This requires `ssl-protocols` to have `TLSv1.3` enabled. Enable this with caution, because requests sent within early
data are subject to [replay attacks](https://tools.ietf.org/html/rfc8470).

[ssl_early_data](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_early_data). The default is: `false`.

## ssl-session-cache

Enables or disables the use of shared [SSL cache](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_session_cache) among worker processes.

## ssl-session-cache-size

Sets the size of the [SSL shared session cache](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_session_cache) between all worker processes.

## ssl-session-tickets

Enables or disables session resumption through [TLS session tickets](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_session_tickets).

## ssl-session-ticket-key

Sets the secret key used to encrypt and decrypt TLS session tickets. The value must be a valid base64 string.
To create a ticket: `openssl rand 80 | openssl enc -A -base64`

[TLS session ticket-key](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_session_tickets), by default, a randomly generated key is used.

## ssl-session-timeout

Sets the time during which a client may [reuse the session](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_session_timeout) parameters stored in a cache.

## ssl-buffer-size

Sets the size of the [SSL buffer](http://nginx.org/en/docs/http/ngx_http_ssl_module.html#ssl_buffer_size) used for sending data. The default of 4k helps NGINX to improve TLS Time To First Byte (TTTFB).

_References:_
[https://www.igvita.com/2013/12/16/optimizing-nginx-tls-time-to-first-byte/](https://www.igvita.com/2013/12/16/optimizing-nginx-tls-time-to-first-byte/)

## use-proxy-protocol

Enables or disables the [PROXY protocol](https://www.nginx.com/resources/admin-guide/proxy-protocol/) to receive client connection (real IP address) information passed through proxy servers and load balancers such as HAProxy and Amazon Elastic Load Balancer (ELB).

## proxy-protocol-header-timeout

Sets the timeout value for receiving the proxy-protocol headers. The default of 5 seconds prevents the TLS passthrough handler from waiting indefinitely on a dropped connection.
_**default:**_ 5s

## use-gzip

Enables or disables compression of HTTP responses using the ["gzip" module](http://nginx.org/en/docs/http/ngx_http_gzip_module.html). MIME types to compress are controlled by [gzip-types](#gzip-types). _**default:**_ false

## use-geoip

Enables or disables ["geoip" module](http://nginx.org/en/docs/http/ngx_http_geoip_module.html) that creates variables with values depending on the client IP address, using the precompiled MaxMind databases.
_**default:**_ true

> __Note:__ MaxMind legacy databases are discontinued and will not receive updates after 2019-01-02, cf. [discontinuation notice](https://support.maxmind.com/geolite-legacy-discontinuation-notice/). Consider [use-geoip2](#use-geoip2) below.

## use-geoip2

Enables the [geoip2 module](https://github.com/leev/ngx_http_geoip2_module) for NGINX.
Since `0.27.0` and due to a [change in the MaxMind databases](https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases) a license is required to have access to the databases.
For this reason, it is required to define a new flag `--maxmind-license-key` in the ingress controller deployment to download the databases needed during the initialization of the ingress controller.
Alternatively, it is possible to use a volume to mount the files `/etc/nginx/geoip/GeoLite2-City.mmdb` and `/etc/nginx/geoip/GeoLite2-ASN.mmdb`, avoiding the overhead of the download.

!!! important
    If the feature is enabled but the files are missing, GeoIP2 will not be enabled.

_**default:**_ false

## enable-brotli

Enables or disables compression of HTTP responses using the ["brotli" module](https://github.com/google/ngx_brotli).
The default mime type list to compress is: `application/xml+rss application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/plain text/x-component`. _**default:**_ is disabled

> __Note:__ Brotli does not works in Safari < 11. For more information see [https://caniuse.com/#feat=brotli](https://caniuse.com/#feat=brotli)

## brotli-level

Sets the Brotli Compression Level that will be used. _**default:**_ 4

## brotli-types

Sets the MIME Types that will be compressed on-the-fly by brotli.
_**default:**_ `application/xml+rss application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/plain text/x-component`

## use-http2

Enables or disables [HTTP/2](http://nginx.org/en/docs/http/ngx_http_v2_module.html) support in secure connections.

## gzip-level

Sets the gzip Compression Level that will be used. _**default:**_ 1

## gzip-min-length

Minimum length of responses to be returned to the client before it is eligible for gzip compression, in bytes. _**default:**_ 256

## gzip-types

Sets the MIME types in addition to "text/html" to compress. The special value "\*" matches any MIME type. Responses with the "text/html" type are always compressed if [`use-gzip`](#use-gzip) is enabled.
_**default:**_ `application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/plain text/x-component`.

## worker-processes

Sets the number of [worker processes](http://nginx.org/en/docs/ngx_core_module.html#worker_processes).
The default of "auto" means number of available CPU cores.

## worker-cpu-affinity

Binds worker processes to the sets of CPUs. [worker_cpu_affinity](http://nginx.org/en/docs/ngx_core_module.html#worker_cpu_affinity).
By default worker processes are not bound to any specific CPUs. The value can be:

- "": empty string indicate no affinity is applied.
- cpumask: e.g. `0001 0010 0100 1000` to bind processes to specific cpus.
- auto: binding worker processes automatically to available CPUs.

## worker-shutdown-timeout

Sets a timeout for Nginx to [wait for worker to gracefully shutdown](http://nginx.org/en/docs/ngx_core_module.html#worker_shutdown_timeout). _**default:**_ "240s"

## load-balance

Sets the algorithm to use for load balancing.
The value can either be:

- round_robin: to use the default round robin loadbalancer
- ewma: to use the Peak EWMA method for routing ([implementation](https://github.com/kubernetes/ingress-nginx/blob/master/rootfs/etc/nginx/lua/balancer/ewma.lua))

The default is `round_robin`.

- To load balance using consistent hashing of IP or other variables, consider the `nginx.ingress.kubernetes.io/upstream-hash-by` annotation.
- To load balance using session cookies, consider the `nginx.ingress.kubernetes.io/affinity` annotation.

_References:_
[http://nginx.org/en/docs/http/load_balancing.html](http://nginx.org/en/docs/http/load_balancing.html)

## variables-hash-bucket-size

Sets the bucket size for the variables hash table.

_References:_
[http://nginx.org/en/docs/http/ngx_http_map_module.html#variables_hash_bucket_size](http://nginx.org/en/docs/http/ngx_http_map_module.html#variables_hash_bucket_size)

## variables-hash-max-size

Sets the maximum size of the variables hash table.

_References:_
[http://nginx.org/en/docs/http/ngx_http_map_module.html#variables_hash_max_size](http://nginx.org/en/docs/http/ngx_http_map_module.html#variables_hash_max_size)

## upstream-keepalive-connections

Activates the cache for connections to upstream servers. The connections parameter sets the maximum number of idle
keepalive connections to upstream servers that are preserved in the cache of each worker process. When this number is
exceeded, the least recently used connections are closed.
_**default:**_ 320

_References:_
[http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive)


## upstream-keepalive-timeout

Sets a timeout during which an idle keepalive connection to an upstream server will stay open.
 _**default:**_ 60

_References:_
[http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive_timeout](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive_timeout)


## upstream-keepalive-requests

Sets the maximum number of requests that can be served through one keepalive connection. After the maximum number of
requests is made, the connection is closed.
_**default:**_ 10000


_References:_
[http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive_requests](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#keepalive_requests)


## limit-conn-zone-variable

Sets parameters for a shared memory zone that will keep states for various keys of [limit_conn_zone](http://nginx.org/en/docs/http/ngx_http_limit_conn_module.html#limit_conn_zone). The default of "$binary_remote_addr" variable’s size is always 4 bytes for IPv4 addresses or 16 bytes for IPv6 addresses.

## proxy-stream-timeout

Sets the timeout between two successive read or write operations on client or proxied server connections. If no data is transmitted within this time, the connection is closed.

_References:_
[http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_timeout](http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_timeout)

## proxy-stream-next-upstream

When a connection to the proxied server cannot be established, determines whether a client connection will be passed to the next server.

_References:_
[http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream](http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream)

## proxy-stream-next-upstream-timeout

Limits the time allowed to pass a connection to the next server. The 0 value turns off this limitation.

_References:_
[http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream_timeout](http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream_timeout)

## proxy-stream-next-upstream-tries

Limits the number of possible tries a request should be passed to the next server. The 0 value turns off this limitation.

_References:_
[http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream_tries](http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_next_upstream_timeout)

## proxy-stream-responses

Sets the number of datagrams expected from the proxied server in response to the client request if the UDP protocol is used.

_References:_
[http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_responses](http://nginx.org/en/docs/stream/ngx_stream_proxy_module.html#proxy_responses)

## bind-address

Sets the addresses on which the server will accept requests instead of *. It should be noted that these addresses must exist in the runtime environment or the controller will crash loop.

## use-forwarded-headers

If true, NGINX passes the incoming `X-Forwarded-*` headers to upstreams. Use this option when NGINX is behind another L7 proxy / load balancer that is setting these headers.

If false, NGINX ignores incoming `X-Forwarded-*` headers, filling them with the request information it sees. Use this option if NGINX is exposed directly to the internet, or it's behind a L3/packet-based load balancer that doesn't alter the source IP in the packets.

## enable-real-ip

`enable-real-ip` enables the configuration of [http://nginx.org/en/docs/http/ngx_http_realip_module.html](http://nginx.org/en/docs/http/ngx_http_realip_module.html). Specific attributes of the module can be configured further by using `forwarded-for-header` and `proxy-real-ip-cidr` settings.

## forwarded-for-header

Sets the header field for identifying the originating IP address of a client. _**default:**_ X-Forwarded-For

## compute-full-forwarded-for

Append the remote address to the X-Forwarded-For header instead of replacing it. When this option is enabled, the upstream application is responsible for extracting the client IP based on its own list of trusted proxies.

## proxy-add-original-uri-header

Adds an X-Original-Uri header with the original request URI to the backend request

## generate-request-id

Ensures that X-Request-ID is defaulted to a random value, if no X-Request-ID is present in the request

## enable-opentracing

Enables the nginx Opentracing extension. _**default:**_ is disabled

_References:_
[https://github.com/opentracing-contrib/nginx-opentracing](https://github.com/opentracing-contrib/nginx-opentracing)

## opentracing-operation-name

Specifies a custom name for the server span. _**default:**_ is empty

For example, set to "HTTP $request_method $uri".

## opentracing-location-operation-name

Specifies a custom name for the location span. _**default:**_ is empty

For example, set to "HTTP $request_method $uri".

## zipkin-collector-host

Specifies the host to use when uploading traces. It must be a valid URL.

## zipkin-collector-port

Specifies the port to use when uploading traces. _**default:**_ 9411

## zipkin-service-name

Specifies the service name to use for any traces created. _**default:**_ nginx

## zipkin-sample-rate

Specifies sample rate for any traces created. _**default:**_ 1.0

## jaeger-collector-host

Specifies the host to use when uploading traces. It must be a valid URL.

## jaeger-collector-port

Specifies the port to use when uploading traces. _**default:**_ 6831

## jaeger-endpoint

Specifies the endpoint to use when uploading traces to a collector. This takes priority over `jaeger-collector-host` if both are specified.

## jaeger-service-name

Specifies the service name to use for any traces created. _**default:**_ nginx

## jaeger-propagation-format

Specifies the traceparent/tracestate propagation format. _**default:**_ jaeger

## jaeger-sampler-type

Specifies the sampler to be used when sampling traces. The available samplers are: const, probabilistic, ratelimiting, remote. _**default:**_ const

## jaeger-sampler-param

Specifies the argument to be passed to the sampler constructor. Must be a number.
For const this should be 0 to never sample and 1 to always sample. _**default:**_ 1

## jaeger-sampler-host

Specifies the custom remote sampler host to be passed to the sampler constructor. Must be a valid URL.
Leave blank to use default value (localhost). _**default:**_ http://127.0.0.1

## jaeger-sampler-port

Specifies the custom remote sampler port to be passed to the sampler constructor. Must be a number. _**default:**_ 5778

## jaeger-trace-context-header-name

Specifies the header name used for passing trace context. _**default:**_ uber-trace-id

## jaeger-debug-header

Specifies the header name used for force sampling. _**default:**_ jaeger-debug-id

## jaeger-baggage-header

Specifies the header name used to submit baggage if there is no root span. _**default:**_ jaeger-baggage

## jaeger-tracer-baggage-header-prefix

Specifies the header prefix used to propagate baggage. _**default:**_ uberctx-

## datadog-collector-host

Specifies the datadog agent host to use when uploading traces. It must be a valid URL.

## datadog-collector-port

Specifies the port to use when uploading traces. _**default:**_ 8126

## datadog-service-name

Specifies the service name to use for any traces created. _**default:**_ nginx

## datadog-environment

Specifies the environment this trace belongs to. _**default:**_ prod

## datadog-operation-name-override

Overrides the operation naem to use for any traces crated. _**default:**_ nginx.handle

## datadog-priority-sampling

Specifies to use client-side sampling.
If true disables client-side sampling (thus ignoring `sample_rate`) and enables distributed priority sampling, where traces are sampled based on a combination of user-assigned priorities and configuration from the agent. _**default:**_ true

## datadog-sample-rate

Specifies sample rate for any traces created.
This is effective only when `datadog-priority-sampling` is `false` _**default:**_ 1.0

## main-snippet

Adds custom configuration to the main section of the nginx configuration.

## http-snippet

Adds custom configuration to the http section of the nginx configuration.

## server-snippet

Adds custom configuration to all the servers in the nginx configuration.

## location-snippet

Adds custom configuration to all the locations in the nginx configuration.

You can not use this to add new locations that proxy to the Kubernetes pods, as the snippet does not have access to the Go template functions. If you want to add custom locations you will have to [provide your own nginx.tmpl](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/custom-template/).

## custom-http-errors

Enables which HTTP codes should be passed for processing with the [error_page directive](http://nginx.org/en/docs/http/ngx_http_core_module.html#error_page)

Setting at least one code also enables [proxy_intercept_errors](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_intercept_errors) which are required to process error_page.

Example usage: `custom-http-errors: 404,415`

## proxy-body-size

Sets the maximum allowed size of the client request body.
See NGINX [client_max_body_size](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size).

## proxy-connect-timeout

Sets the timeout for [establishing a connection with a proxied server](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_connect_timeout). It should be noted that this timeout cannot usually exceed 75 seconds.

## proxy-read-timeout

Sets the timeout in seconds for [reading a response from the proxied server](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_read_timeout). The timeout is set only between two successive read operations, not for the transmission of the whole response.

## proxy-send-timeout

Sets the timeout in seconds for [transmitting a request to the proxied server](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_send_timeout). The timeout is set only between two successive write operations, not for the transmission of the whole request.

## proxy-buffers-number

Sets the number of the buffer used for [reading the first part of the response](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffers) received from the proxied server. This part usually contains a small response header.

## proxy-buffer-size

Sets the size of the buffer used for [reading the first part of the response](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffer_size) received from the proxied server. This part usually contains a small response header.

## proxy-cookie-path

Sets a text that [should be changed in the path attribute](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cookie_path) of the “Set-Cookie” header fields of a proxied server response.

## proxy-cookie-domain

Sets a text that [should be changed in the domain attribute](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cookie_domain) of the “Set-Cookie” header fields of a proxied server response.

## proxy-next-upstream

Specifies in [which cases](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream) a request should be passed to the next server.

## proxy-next-upstream-timeout

[Limits the time](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream_timeout) in seconds during which a request can be passed to the next server.

## proxy-next-upstream-tries

Limit the number of [possible tries](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream_tries) a request should be passed to the next server.

## proxy-redirect-from

Sets the original text that should be changed in the "Location" and "Refresh" header fields of a proxied server response. _**default:**_ off

_References:_
[http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect)

## proxy-request-buffering

Enables or disables [buffering of a client request body](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_request_buffering).

## ssl-redirect

Sets the global value of redirects (301) to HTTPS if the server has a TLS certificate (defined in an Ingress rule).
_**default:**_ "true"

## whitelist-source-range

Sets the default whitelisted IPs for each `server` block. This can be overwritten by an annotation on an Ingress rule.
See [ngx_http_access_module](http://nginx.org/en/docs/http/ngx_http_access_module.html).

## skip-access-log-urls

Sets a list of URLs that should not appear in the NGINX access log. This is useful with urls like `/health` or `health-check` that make "complex" reading the logs. _**default:**_ is empty

## limit-rate

Limits the rate of response transmission to a client. The rate is specified in bytes per second. The zero value disables rate limiting. The limit is set per a request, and so if a client simultaneously opens two connections, the overall rate will be twice as much as the specified limit.

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate](http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate)

## limit-rate-after

Sets the initial amount after which the further transmission of a response to a client will be rate limited.

## lua-shared-dicts

Customize default Lua shared dictionaries or define more. You can use the following syntax to do so:

```
lua-shared-dicts: "<my dict name>: <my dict size>, [<my dict name>: <my dict size>], ..."
```

For example following will set default `certificate_data` dictionary to `100M` and will introduce a new dictionary called
`my_custom_plugin`:

```
lua-shared-dicts: "certificate_data: 100, my_custom_plugin: 5"
```

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate_after](http://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate_after)

## http-redirect-code

Sets the HTTP status code to be used in redirects.
Supported codes are [301](https://developer.mozilla.org/docs/Web/HTTP/Status/301),[302](https://developer.mozilla.org/docs/Web/HTTP/Status/302),[307](https://developer.mozilla.org/docs/Web/HTTP/Status/307) and [308](https://developer.mozilla.org/docs/Web/HTTP/Status/308)
_**default:**_ 308

> __Why the default code is 308?__

> [RFC 7238](https://tools.ietf.org/html/rfc7238) was created to define the 308 (Permanent Redirect) status code that is similar to 301 (Moved Permanently) but it keeps the payload in the redirect. This is important if we send a redirect in methods like POST.

## proxy-buffering

Enables or disables [buffering of responses from the proxied server](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_buffering).

## limit-req-status-code

Sets the [status code to return in response to rejected requests](http://nginx.org/en/docs/http/ngx_http_limit_req_module.html#limit_req_status). _**default:**_ 503

## limit-conn-status-code

Sets the [status code to return in response to rejected connections](http://nginx.org/en/docs/http/ngx_http_limit_conn_module.html#limit_conn_status). _**default:**_ 503

## no-tls-redirect-locations

A comma-separated list of locations on which http requests will never get redirected to their https counterpart.
_**default:**_ "/.well-known/acme-challenge"

## global-auth-url

A url to an existing service that provides authentication for all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-url`.
Locations that should not get authenticated can be listed using `no-auth-locations` See [no-auth-locations](#no-auth-locations). In addition, each service can be excluded from authentication via annotation `enable-global-auth` set to "false".
_**default:**_ ""

_References:_ [https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md#external-authentication](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md#external-authentication)

## global-auth-method

A HTTP method to use for an existing service that provides authentication for all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-method`.
_**default:**_ ""

## global-auth-signin

Sets the location of the error page for an existing service that provides authentication for all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-signin`.
_**default:**_ ""

## global-auth-signin-redirect-param

Sets the query parameter in the error page signin URL which contains the original URL of the request that failed authentication.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-signin-redirect-param`.
_**default:**_ "rd"

## global-auth-response-headers

Sets the headers to pass to backend once authentication request completes. Applied to all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-response-headers`.
_**default:**_ ""

## global-auth-request-redirect

Sets the X-Auth-Request-Redirect header value. Applied to all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-request-redirect`.
_**default:**_ ""

## global-auth-snippet

Sets a custom snippet to use with external authentication. Applied to all the locations.
Similar to the Ingress rule annotation `nginx.ingress.kubernetes.io/auth-request-redirect`.
_**default:**_ ""

## global-auth-cache-key

Enables caching for global auth requests. Specify a lookup key for auth responses, e.g. `$remote_user$http_authorization`.

## global-auth-cache-duration

Set a caching time for auth responses based on their response codes, e.g. `200 202 30m`. See [proxy_cache_valid](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_cache_valid) for details. You may specify multiple, comma-separated values: `200 202 10m, 401 5m`. defaults to `200 202 401 5m`.

## no-auth-locations

A comma-separated list of locations that should not get authenticated.
_**default:**_ "/.well-known/acme-challenge"

## block-cidrs

A comma-separated list of IP addresses (or subnets), request from which have to be blocked globally.

_References:_
[http://nginx.org/en/docs/http/ngx_http_access_module.html#deny](http://nginx.org/en/docs/http/ngx_http_access_module.html#deny)

## block-user-agents

A comma-separated list of User-Agent, request from which have to be blocked globally.
It's possible to use here full strings and regular expressions. More details about valid patterns can be found at `map` Nginx directive documentation.

_References:_
[http://nginx.org/en/docs/http/ngx_http_map_module.html#map](http://nginx.org/en/docs/http/ngx_http_map_module.html#map)

## block-referers

A comma-separated list of Referers, request from which have to be blocked globally.
It's possible to use here full strings and regular expressions. More details about valid patterns can be found at `map` Nginx directive documentation.

_References:_
[http://nginx.org/en/docs/http/ngx_http_map_module.html#map](http://nginx.org/en/docs/http/ngx_http_map_module.html#map)

## proxy-ssl-location-only

Set if proxy-ssl parameters should be applied only on locations and not on servers.
_**default:**_ is disabled

## default-type

Sets the default MIME type of a response.
_**default:**_ text/html

_References:_
[http://nginx.org/en/docs/http/ngx_http_core_module.html#default_type](http://nginx.org/en/docs/http/ngx_http_core_module.html#default_type)

## global-rate-limit

* `global-rate-limit-status-code`: configure HTTP status code to return when rejecting requests. Defaults to 429.

Configure `memcached` client for [Global Rate Limiting](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md#global-rate-limiting).

* `global-rate-limit-memcached-host`: IP/FQDN of memcached server to use. Required to enable Global Rate Limiting.
* `global-rate-limit-memcached-port`: port of memcached server to use. Defaults default memcached port of `11211`.
* `global-rate-limit-memcached-connect-timeout`: configure timeout for connect, send and receive operations. Unit is millisecond. Defaults to 50ms.
* `global-rate-limit-memcached-max-idle-timeout`: configure timeout for cleaning idle connections. Unit is millisecond. Defaults to 50ms. 
* `global-rate-limit-memcached-pool-size`: configure number of max connections to keep alive. Make sure your `memcached` server can handle
`global-rate-limit-memcached-pool-size * worker-processes * <number of ingress-nginx replicas>` simultaneous connections.

These settings get used by [lua-resty-global-throttle](https://github.com/ElvinEfendi/lua-resty-global-throttle)
that ingress-nginx includes. Refer to the link to learn more about `lua-resty-global-throttle`.
