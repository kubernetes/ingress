

# e2e test suite for [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/tree/main/)



### [[Default Backend] change default settings](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/with_hosts.go#L31)

- [should apply the annotation to the default backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/with_hosts.go#L39)

### [[Default Backend]](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/default_backend.go#L29)

- [should return 404 sending requests when only a default backend is running](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/default_backend.go#L32)
- [enables access logging for default backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/default_backend.go#L89)
- [disables access logging for default backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/default_backend.go#L106)

### [[Default Backend] custom service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/custom_default_backend.go#L33)

- [uses custom default backend that returns 200 as status code](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/custom_default_backend.go#L36)

### [[Default Backend] SSL](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/ssl.go#L26)

- [should return a self generated SSL certificate](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/defaultbackend/ssl.go#L29)

### [[TCP] tcp-services](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/tcpudp/tcp.go#L37)

- [should expose a TCP service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/tcpudp/tcp.go#L40)
- [should expose an ExternalName TCP service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/tcpudp/tcp.go#L98)

### [auth-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L39)

- [should return status code 200 when no authentication is configured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L46)
- [should return status code 503 when authentication is configured with an invalid secret](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L65)
- [should return status code 401 when authentication is configured but Authorization header is not configured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L89)
- [should return status code 401 when authentication is configured and Authorization header is sent with invalid credentials](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L116)
- [should return status code 200 when authentication is configured and Authorization header is sent](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L144)
- [should return status code 200 when authentication is configured with a map and Authorization header is sent](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L171)
- [should return status code 401 when authentication is configured with invalid content and Authorization header is sent](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L199)
- [proxy_set_header My-Custom-Header 42;](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L238)
- [proxy_set_header My-Custom-Header 42;](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L256)
- [proxy_set_header 'My-Custom-Header' '42';](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L273)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L294)
- [retains cookie set by external authentication server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L316)
- [should return status code 200 when signed in](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L413)
- [should redirect to signin url when not signed in](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L422)
- [keeps processing new ingresses even if one of the existing ingresses is misconfigured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L433)
- [should return status code 200 when signed in](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L490)
- [should redirect to signin url when not signed in](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L499)
- [keeps processing new ingresses even if one of the existing ingresses is misconfigured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L510)
- [should return status code 200 when signed in after auth backend is deleted ](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L584)
- [should deny login for different location on same server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L604)
- [should deny login for different servers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L632)
- [should redirect to signin url when not signed in](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/auth.go#L661)

### [affinitymode](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinitymode.go#L31)

- [Balanced affinity mode should balance](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinitymode.go#L34)
- [Check persistent affinity mode](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinitymode.go#L64)

### [proxy-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L28)

- [should set proxy_redirect to off](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L36)
- [should set proxy_redirect to default](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L52)
- [should set proxy_redirect to hello.com goodbye.com](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L68)
- [should set proxy client-max-body-size to 8m](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L85)
- [should not set proxy client-max-body-size to incorrect value](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L100)
- [should set valid proxy timeouts](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L115)
- [should not set invalid proxy timeouts](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L136)
- [should turn on proxy-buffering](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L157)
- [should turn off proxy-request-buffering](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L179)
- [should build proxy next upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L194)
- [should setup proxy cookies](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L215)
- [should change the default proxy HTTP version](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxy.go#L233)

### [affinity session-cookie-name](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L35)

- [should set sticky cookie SERVERID](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L42)
- [should change cookie name on ingress definition change](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L64)
- [should set the path to /something on the generated cookie](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L99)
- [does not set the path to / on the generated cookie if there's more than one rule referring to the same backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L121)
- [should set cookie with expires](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L182)
- [should work with use-regex annotation and session-cookie-path](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L213)
- [should warn user when use-regex is true and session-cookie-path is not set](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L237)
- [should not set affinity across all server locations when using separate ingresses](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L263)
- [should set sticky cookie without host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L295)
- [should work with server-alias annotation](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/affinity.go#L315)

### [mirror-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/mirror.go#L28)

- [should set mirror-target to http://localhost/mirror](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/mirror.go#L36)
- [should set mirror-target to https://test.env.com/$request_uri](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/mirror.go#L51)
- [should disable mirror-request-body](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/mirror.go#L67)

### [canary-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L35)

- [should response with a 200 status from the mainline upstream when requests are made to the mainline ingress](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L47)
- [should return 404 status for requests to the canary if no matching ingress is found](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L79)
- [should return the correct status codes when endpoints are unavailable](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L106)
- [should route requests to the correct upstream if mainline ingress is created before the canary ingress](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L160)
- [should route requests to the correct upstream if mainline ingress is created after the canary ingress](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L205)
- [should route requests to the correct upstream if the mainline ingress is modified](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L249)
- [should route requests to the correct upstream if the canary ingress is modified](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L306)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L361)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L415)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L479)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L521)
- [should routes to mainline upstream when the given Regex causes error](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L555)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L593)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L632)
- [should route requests to the correct upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L687)
- [should not use canary as a catch-all server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L743)
- [should not use canary with domain as a server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L771)
- [does not crash when canary ingress has multiple paths to the same non-matching backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/canary.go#L795)

### [limit-rate](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/limitrate.go#L29)

- [Check limit-rate annotation](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/limitrate.go#L37)

### [force-ssl-redirect](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/forcesslredirect.go#L27)

- [should redirect to https](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/forcesslredirect.go#L34)

### [http2-push-preload](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/http2pushpreload.go#L27)

- [enable the http2-push-preload directive](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/http2pushpreload.go#L34)

### [proxy-ssl-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L30)

- [should set valid proxy-ssl-secret](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L37)
- [should set valid proxy-ssl-secret, proxy-ssl-verify to on, proxy-ssl-verify-depth to 2, and proxy-ssl-server-name to on](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L64)
- [should set valid proxy-ssl-secret, proxy-ssl-ciphers to HIGH:!AES](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L94)
- [should set valid proxy-ssl-secret, proxy-ssl-protocols](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L122)
- [proxy-ssl-location-only flag should change the nginx config server part](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/proxyssl.go#L150)

### [modsecurity owasp](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L27)

- [should enable modsecurity](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L34)
- [should enable modsecurity with transaction ID and OWASP rules](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L52)
- [should disable modsecurity](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L73)
- [should enable modsecurity with snippet](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L90)
- [should enable modsecurity without using 'modsecurity on;'](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L109)
- [should disable modsecurity using 'modsecurity off;'](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L131)
- [should enable modsecurity with snippet and block requests](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L152)
- [should enable modsecurity globally and with modsecurity-snippet block requests](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/modsecurity.go#L186)

### [backend-protocol - GRPC](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/grpc.go#L38)

- [should use grpc_pass in the configuration file](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/grpc.go#L41)
- [should return OK for service with backend protocol GRPC](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/grpc.go#L66)
- [should return OK for service with backend protocol GRPCS](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/grpc.go#L124)

### [cors-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L28)

- [should enable cors](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L35)
- [should set cors methods to only allow POST, GET](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L60)
- [should set cors max-age](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L76)
- [should disable cors allow credentials](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L92)
- [should allow origin for cors](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L108)
- [should allow headers for cors](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L124)
- [should expose headers for cors](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/cors.go#L140)

### [influxdb-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/influxdb.go#L39)

- [should send the request metric to the influxdb server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/influxdb.go#L48)

### [Annotation - limit-connections](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/limitconnections.go#L31)

- [should limit-connections](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/limitconnections.go#L38)

### [client-body-buffer-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L28)

- [should set client_body_buffer_size to 1000](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L35)
- [should set client_body_buffer_size to 1K](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L51)
- [should set client_body_buffer_size to 1k](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L67)
- [should set client_body_buffer_size to 1m](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L83)
- [should set client_body_buffer_size to 1M](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L99)
- [should not set client_body_buffer_size to invalid 1b](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/clientbodybuffersize.go#L115)

### [default-backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/default_backend.go#L29)

- [should use a custom default backend as upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/default_backend.go#L37)

### [connection-proxy-header](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/connection.go#L29)

- [set connection header to keep-alive](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/connection.go#L36)

### [upstream-vhost](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/upstreamvhost.go#L27)

- [set host to upstreamvhost.bar.com](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/upstreamvhost.go#L34)

### [custom-http-errors](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/customhttperrors.go#L34)

- [configures Nginx correctly](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/customhttperrors.go#L41)

### [disable-access-log disable-http-access-log disable-stream-access-log](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/disableaccesslog.go#L27)

- [disable-access-log set access_log off](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/disableaccesslog.go#L34)
- [disable-http-access-log set access_log off](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/disableaccesslog.go#L46)

### [server-snippet](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/serversnippet.go#L27)

- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/serversnippet.go#L34)

### [rewrite-target use-regex enable-rewrite-log](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L30)

- [should write rewrite logs](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L37)
- [should use correct longest path match](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L66)
- [should use ~* location modifier if regex annotation is present](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L111)
- [should fail to use longest match for documented warning](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L158)
- [should allow for custom rewrite parameters](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/rewrite.go#L190)

### [app-root](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/approot.go#L28)

- [should redirect to /foo](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/approot.go#L35)

### [whitelist-source-range](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/ipwhitelist.go#L26)

- [should set valid ip whitelist range](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/ipwhitelist.go#L33)

### [enable-access-log enable-rewrite-log](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/log.go#L27)

- [set access_log off](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/log.go#L34)
- [set rewrite_log on](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/log.go#L49)

### [x-forwarded-prefix](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/xforwardedprefix.go#L28)

- [should set the X-Forwarded-Prefix to the annotation value](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/xforwardedprefix.go#L35)
- [should not add X-Forwarded-Prefix if the annotation value is empty](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/xforwardedprefix.go#L57)

### [configuration-snippet](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/snippet.go#L27)

- [ in all locations](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/snippet.go#L34)

### [backend-protocol - FastCGI](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fastcgi.go#L31)

- [should use fastcgi_pass in the configuration file](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fastcgi.go#L38)
- [should add fastcgi_index in the configuration file](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fastcgi.go#L55)
- [should add fastcgi_param in the configuration file](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fastcgi.go#L72)
- [should return OK for service with backend protocol FastCGI](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fastcgi.go#L105)

### [from-to-www-redirect](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fromtowwwredirect.go#L31)

- [should redirect from www HTTP to HTTP](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fromtowwwredirect.go#L38)
- [should redirect from www HTTPS to HTTPS](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/fromtowwwredirect.go#L64)

### [permanent-redirect permanent-redirect-code](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/redirect.go#L30)

- [should respond with a standard redirect code](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/redirect.go#L33)
- [should respond with a custom redirect code](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/redirect.go#L61)

### [upstream-hash-by-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/upstreamhashby.go#L76)

- [should connect to the same pod](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/upstreamhashby.go#L83)
- [should connect to the same subset of pods](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/upstreamhashby.go#L92)

### [annotation-global-rate-limit](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/globalratelimit.go#L30)

- [generates correct configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/globalratelimit.go#L38)

### [backend-protocol](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L27)

- [should set backend protocol to https:// and use proxy_pass](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L34)
- [should set backend protocol to grpc:// and use grpc_pass](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L49)
- [should set backend protocol to grpcs:// and use grpc_pass](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L64)
- [should set backend protocol to '' and use fastcgi_pass](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L79)
- [should set backend protocol to '' and use ajp_pass](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/backendprotocol.go#L94)

### [satisfy](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/satisfy.go#L35)

- [should configure satisfy directive correctly](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/satisfy.go#L42)
- [should allow multiple auth with satisfy any](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/satisfy.go#L84)

### [server-alias](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/alias.go#L29)

- [should return status code 200 for host 'foo' and 404 for 'bar'](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/alias.go#L36)
- [should return status code 200 for host 'foo' and 'bar'](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/alias.go#L62)
- [should return status code 200 for hosts defined in two ingresses, different path with one alias](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/alias.go#L87)

### [ssl-ciphers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/sslciphers.go#L27)

- [should change ssl ciphers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/sslciphers.go#L34)

### [auth-tls-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/authtls.go#L30)

- [should set valid auth-tls-secret](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/authtls.go#L37)
- [should set valid auth-tls-secret, sslVerify to off, and sslVerifyDepth to 2](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/authtls.go#L73)
- [should set valid auth-tls-secret, pass certificate to upstream, and error page](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/authtls.go#L103)
- [should validate auth-tls-verify-client](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/annotations/authtls.go#L153)

### [[Status] status update](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/status/update.go#L38)

- [should update status field after client-go reconnection](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/status/update.go#L43)

### [Debug CLI](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/dbg/main.go#L29)

- [should list the backend servers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/dbg/main.go#L37)
- [should get information for a specific backend server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/dbg/main.go#L56)
- [should produce valid JSON for /dbg general](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/dbg/main.go#L85)

### [[Memory Leak] Dynamic Certificates](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/leaks/lua_ssl.go#L35)

- [should not leak memory from ingress SSL certificates or configuration updates](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/leaks/lua_ssl.go#L42)

### [[Ingress] [PathType] mix Exact and Prefix paths](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_mixed.go#L30)

- [should choose the correct location](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_mixed.go#L39)

### [[Ingress] definition without host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/without_host.go#L32)

- [should set ingress details variables for ingresses without a host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/without_host.go#L35)
- [should set ingress details variables for ingresses with host without IngressRuleValue, only Backend](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/without_host.go#L56)

### [single ingress - multiple hosts](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/multiple_rules.go#L31)

- [should set the correct $service_name NGINX variable](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/multiple_rules.go#L39)

### [[Ingress] [PathType] exact](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_exact.go#L30)

- [should choose exact location for /exact](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_exact.go#L37)

### [[Ingress] [PathType] prefix checks](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_prefix.go#L28)

- [should return 404 when prefix /aaa does not match request /aaaccc](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/ingress/pathtype_prefix.go#L35)

### [[Security] request smuggling](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/security/request_smuggling.go#L32)

- [should not return body content from error_page](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/security/request_smuggling.go#L39)

### [[SSL] [Flag] default-ssl-certificate](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/default_ssl_certificate.go#L33)

- [uses default ssl certificate for catch-all ingress](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/default_ssl_certificate.go#L64)
- [uses default ssl certificate for host based ingress when configured certificate does not match host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/default_ssl_certificate.go#L80)

### [enable-real-ip](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/enable_real_ip.go#L30)

- [trusts X-Forwarded-For header only when setting is true](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/enable_real_ip.go#L40)
- [should not trust X-Forwarded-For header when setting is false](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/enable_real_ip.go#L78)

### [access-log](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L26)

- [use the default configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L31)
- [use the specified configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L39)
- [use the specified configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L51)
- [use the specified configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L63)
- [use the specified configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/access_log.go#L75)

### [[Lua] lua-shared-dicts](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/lua_shared_dicts.go#L26)

- [configures lua shared dicts](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/lua_shared_dicts.go#L29)

### [server-tokens](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/server_tokens.go#L30)

- [should not exists Server header in the response](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/server_tokens.go#L38)
- [should exists Server header in the response when is enabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/server_tokens.go#L50)

### [use-proxy-protocol](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_protocol.go#L36)

- [should respect port passed by the PROXY Protocol](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_protocol.go#L46)
- [should respect proto passed by the PROXY Protocol server port](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_protocol.go#L79)
- [should enable PROXY Protocol for HTTPS](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_protocol.go#L112)
- [should enable PROXY Protocol for TCP](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_protocol.go#L155)

### [[Flag] custom HTTP and HTTPS ports](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/listen_nondefault_ports.go#L32)

- [should set X-Forwarded-Port headers accordingly when listening on a non-default HTTP port](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/listen_nondefault_ports.go#L48)
- [should set X-Forwarded-Port header to 443](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/listen_nondefault_ports.go#L70)
- [should set the X-Forwarded-Port header to 443](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/listen_nondefault_ports.go#L100)

### [[Security] no-auth-locations](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_auth_locations.go#L34)

- [should return status code 401 when accessing '/' unauthentication](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_auth_locations.go#L55)
- [should return status code 200 when accessing '/' authentication](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_auth_locations.go#L69)
- [should return status code 200 when accessing '/noauth' unauthenticated](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_auth_locations.go#L83)

### [Dynamic $proxy_host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_host.go#L28)

- [should exist a proxy_host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_host.go#L36)
- [should exist a proxy_host using the upstream-vhost annotation value](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_host.go#L57)

### [proxy-connect-timeout](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_connect_timeout.go#L28)

- [should set valid proxy timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_connect_timeout.go#L36)
- [should not set invalid proxy timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_connect_timeout.go#L52)

### [[Security] Pod Security Policies](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/pod_security_policy.go#L40)

- [should be running with a Pod Security Policy](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/pod_security_policy.go#L43)

### [Geoip2](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/geoip2.go#L29)

- [should only allow requests from specific countries](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/geoip2.go#L38)

### [[Security] Pod Security Policies with volumes](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/pod_security_policy_volumes.go#L36)

- [should be running with a Pod Security Policy](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/pod_security_policy_volumes.go#L39)

### [enable-multi-accept](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/multi_accept.go#L27)

- [should be enabled by default](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/multi_accept.go#L31)
- [should be enabled when set to true](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/multi_accept.go#L39)
- [should be disabled when set to false](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/multi_accept.go#L49)

### [log-format-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L28)

- [should disable the log-format-escape-json by default](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L40)
- [should enable the log-format-escape-json](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L47)
- [should disable the log-format-escape-json](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L55)
- [log-format-escape-json enabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L66)
- [log-format-escape-json disabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/log-format.go#L89)

### [[Flag] ingress-class](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ingress_class.go#L39)

- [should ignore Ingress with class](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ingress_class.go#L91)
- [should ignore Ingress with no class](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ingress_class.go#L143)
- [should delete Ingress when class is removed](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ingress_class.go#L177)
- [check scenarios for IngressClass and ingress.class annotation](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ingress_class.go#L216)

### [ssl-ciphers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ssl_ciphers.go#L28)

- [Add ssl ciphers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ssl_ciphers.go#L31)

### [proxy-next-upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_next_upstream.go#L28)

- [should build proxy next upstream using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_next_upstream.go#L36)

### [[Security] global-auth-url](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L32)

- [should return status code 401 when request any protected service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L83)
- [should return status code 200 when request whitelisted (via no-auth-locations) service and 401 when request protected service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L100)
- [should return status code 200 when request whitelisted (via ingress annotation) service and 401 when request protected service](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L124)
- [should still return status code 200 after auth backend is deleted using cache ](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L153)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L195)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L208)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L221)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L235)
- [](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_external_auth.go#L248)

### [[Security] block-*](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_access_block.go#L28)

- [should block CIDRs defined in the ConfigMap](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_access_block.go#L38)
- [should block User-Agents defined in the ConfigMap](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_access_block.go#L55)
- [should block Referers defined in the ConfigMap](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/global_access_block.go#L88)

### [plugins](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/plugins.go#L28)

- [should exist a x-hello-world header](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/plugins.go#L35)

### [Configmap - limit-rate](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/limit_rate.go#L28)

- [Check limit-rate config](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/limit_rate.go#L36)

### [Configure OpenTracing](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L47)

- [should not exists opentracing directive](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L57)
- [should exists opentracing directive when is enabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L70)
- [should not exists opentracing_operation_name directive when is empty](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L84)
- [should exists opentracing_operation_name directive when is configured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L99)
- [should not exists opentracing_location_operation_name directive when is empty](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L114)
- [should exists opentracing_location_operation_name directive when is configured](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L129)
- [should enable opentracing using zipkin](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L144)
- [should enable opentracing using jaeger](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L156)
- [should enable opentracing using jaeger with sampler host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L168)
- [should propagate the w3c header when configured with jaeger](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L181)
- [should enable opentracing using datadog](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/opentracing.go#L225)

### [use-forwarded-headers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/forwarded_headers.go#L30)

- [should trust X-Forwarded headers when setting is true](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/forwarded_headers.go#L40)
- [should not trust X-Forwarded headers when setting is false](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/forwarded_headers.go#L90)

### [proxy-send-timeout](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_send_timeout.go#L28)

- [should set valid proxy send timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_send_timeout.go#L36)
- [should not set invalid proxy send timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_send_timeout.go#L52)

### [Add no tls redirect locations](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_tls_redirect_locations.go#L28)

- [Check no tls redirect locations config](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/no_tls_redirect_locations.go#L31)

### [settings-global-rate-limit](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/globalratelimit.go#L30)

- [generates correct NGINX configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/globalratelimit.go#L38)

### [add-headers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/custom_header.go#L30)

- [Add a custom header](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/custom_header.go#L40)
- [Add multiple custom headers](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/custom_header.go#L65)

### [hash size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L27)

- [should set server_names_hash_bucket_size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L40)
- [should set server_names_hash_max_size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L48)
- [should set proxy-headers-hash-bucket-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L60)
- [should set proxy-headers-hash-max-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L68)
- [should set variables-hash-bucket-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L80)
- [should set variables-hash-max-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L88)
- [should set vmap-hash-bucket-size](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/hash-size.go#L100)

### [keep-alive keep-alive-requests](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L28)

- [should set keepalive_timeout](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L40)
- [should set keepalive_requests](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L48)
- [should set keepalive connection to upstream server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L59)
- [should set keep alive connection timeout to upstream server](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L68)
- [should set the request count to upstream server through one keep alive connection](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/keep-alive.go#L77)

### [[Flag] disable-catch-all](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/disable_catch_all.go#L34)

- [should ignore catch all Ingress](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/disable_catch_all.go#L51)
- [should delete Ingress updated to catch-all](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/disable_catch_all.go#L70)
- [should allow Ingress with both a default backend and rules](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/disable_catch_all.go#L108)

### [main-snippet](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/main_snippet.go#L27)

- [should add value of main-snippet setting to nginx config](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/main_snippet.go#L31)

### [[SSL] TLS protocols, ciphers and headers)](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L31)

- [setting cipher suite](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L65)
- [enforcing TLS v1.0](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L87)
- [setting max-age parameter](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L133)
- [setting includeSubDomains parameter](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L149)
- [setting preload parameter](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L168)
- [overriding what's set from the upstream](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L188)
- [should not use ports during the HTTP to HTTPS redirection](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L209)
- [should not use ports or X-Forwarded-Host during the HTTP to HTTPS redirection](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/tls.go#L227)

### [Configmap change](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/configmap_change.go#L29)

- [should reload after an update in the configuration](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/configmap_change.go#L36)

### [proxy-read-timeout](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_read_timeout.go#L28)

- [should set valid proxy read timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_read_timeout.go#L36)
- [should not set invalid proxy read timeouts using configmap values](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/proxy_read_timeout.go#L52)

### [[Security] modsecurity-snippet](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/modsecurity_snippet.go#L27)

- [should add value of modsecurity-snippet setting to nginx config](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/modsecurity_snippet.go#L30)

### [OCSP](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ocsp/ocsp.go#L42)

- [should enable OCSP and contain stapling information in the connection](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/ocsp/ocsp.go#L49)

### [reuse-port](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/reuse-port.go#L27)

- [reuse port should be enabled by default](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/reuse-port.go#L38)
- [reuse port should be disabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/reuse-port.go#L44)
- [reuse port should be enabled](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/settings/reuse-port.go#L52)

### [[Shutdown] Graceful shutdown with pending request](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/slow_requests.go#L28)

- [should let slow requests finish before shutting down](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/slow_requests.go#L36)

### [[Shutdown] ingress controller](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/shutdown.go#L33)

- [should shutdown in less than 60 secons without pending connections](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/shutdown.go#L43)
- [should shutdown after waiting 60 seconds for pending connections to be closed](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/shutdown.go#L64)
- [should shutdown after waiting 150 seconds for pending connections to be closed](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/gracefulshutdown/shutdown.go#L109)

### [[Service] backend status code 503](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_backend.go#L32)

- [should return 503 when backend service does not exist](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_backend.go#L35)
- [should return 503 when all backend service endpoints are unavailable](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_backend.go#L53)

### [[Service] Type ExternalName](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L37)

- [works with external name set to incomplete fqdn](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L40)
- [should return 200 for service type=ExternalName without a port defined](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L73)
- [should return 200 for service type=ExternalName with a port defined](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L107)
- [should return status 502 for service type=ExternalName with an invalid host](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L148)
- [should return 200 for service type=ExternalName using a port name](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L179)
- [should return 200 for service type=ExternalName using FQDN with trailing dot](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L221)
- [should update the external name after a service update](https://github.com/kubernetes/ingress-nginx/tree/main/test/e2e/servicebackend/service_externalname.go#L252)
