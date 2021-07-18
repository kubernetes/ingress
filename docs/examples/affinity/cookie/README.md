# Sticky sessions

This example demonstrates how to achieve session affinity using cookies.

## Deployment

Session affinity can be configured using the following annotations:

|Name|Description|Value|
| --- | --- | --- |
|nginx.ingress.kubernetes.io/affinity|Type of the affinity, set this to `cookie` to enable session affinity|string (NGINX only supports `cookie`)|
|nginx.ingress.kubernetes.io/affinity-mode|The affinity mode defines how sticky a session is. Use `balanced` to redistribute some sessions when scaling pods or `persistent` for maximum stickiness.|`balanced` (default) or `persistent`|
|nginx.ingress.kubernetes.io/affinity-canary-behavior|Defines session affinity behavior of canaries. By default the behavior is `sticky`, and canaries respect session affinity configuration. Set this to `legacy` to restore original canary behavior, when session affinity parameters were not respected.|`sticky` (default) or `legacy`|
|nginx.ingress.kubernetes.io/session-cookie-name|Name of the cookie that will be created|string (defaults to `INGRESSCOOKIE`)|
|nginx.ingress.kubernetes.io/session-cookie-path|Path that will be set on the cookie (required if your [Ingress paths][ingress-paths] use regular expressions)|string (defaults to the currently [matched path][ingress-paths])|
|nginx.ingress.kubernetes.io/session-cookie-samesite|SameSite attribute to apply to the cookie|Browser accepted values are `None`, `Lax`, and `Strict`|
|nginx.ingress.kubernetes.io/session-cookie-conditional-samesite-none|Will omit `SameSite=None` attribute for older browsers which reject the more-recently defined `SameSite=None` value|`"true"` or `"false"`
|nginx.ingress.kubernetes.io/session-cookie-max-age|Time until the cookie expires, corresponds to the `Max-Age` cookie directive|number of seconds|
|nginx.ingress.kubernetes.io/session-cookie-expires|Legacy version of the previous annotation for compatibility with older browsers, generates an `Expires` cookie directive by adding the seconds to the current date|number of seconds|
|nginx.ingress.kubernetes.io/session-cookie-change-on-failure|When set to `false` nginx ingress will send request to upstream pointed by sticky cookie even if previous attempt failed. When set to `true` and previous attempt failed, sticky cookie will be changed to point to another upstream.|`true` or `false` (defaults to `false`)|

You can create the [example Ingress](ingress.yaml) to test this:

```console
kubectl create -f ingress.yaml
```

## Validation

You can confirm that the Ingress works:

```console
$ kubectl describe ing nginx-test
Name:			nginx-test
Namespace:		default
Address:
Default backend:	default-http-backend:80 (10.180.0.4:8080,10.240.0.2:8080)
Rules:
  Host	                        Path	Backends
  ----	                        ----	--------
  stickyingress.example.com
                                /   	 nginx-service:80 (<none>)
Annotations:
  affinity:	cookie
  session-cookie-name:		INGRESSCOOKIE
  session-cookie-expires: 172800
  session-cookie-max-age: 172800
Events:
  FirstSeen	LastSeen	Count	From				SubObjectPath	Type		Reason	Message
  ---------	--------	-----	----				-------------	--------	------	-------
  7s		7s		1	{nginx-ingress-controller }			Normal		CREATE	default/nginx-test


$ curl -I http://stickyingress.example.com
HTTP/1.1 200 OK
Server: nginx/1.11.9
Date: Fri, 10 Feb 2017 14:11:12 GMT
Content-Type: text/html
Content-Length: 612
Connection: keep-alive
Set-Cookie: INGRESSCOOKIE=a9907b79b248140b56bb13723f72b67697baac3d; Expires=Sun, 12-Feb-17 14:11:12 GMT; Max-Age=172800; Path=/; HttpOnly
Last-Modified: Tue, 24 Jan 2017 14:02:19 GMT
ETag: "58875e6b-264"
Accept-Ranges: bytes
```

In the example above, you can see that the response contains a `Set-Cookie` header with the settings we have defined.
This cookie is created by NGINX, it contains a randomly generated key corresponding to the upstream used for that request (selected using [consistent hashing][consistent-hashing]) and has an `Expires` directive.
If the user changes this cookie, NGINX creates a new one and redirects the user to another upstream.

If the backend pool grows NGINX will keep sending the requests through the same server of the first request, even if it's overloaded.

When the backend server is removed, the requests are re-routed to another upstream server. This does not require the cookie to be updated because the key's [consistent hash][consistent-hashing] will change.

When you have a Service pointing to more than one Ingress, with only one containing affinity configuration, the first created Ingress will be used.
This means that you can face the situation that you've configured session affinity on one Ingress and it doesn't work because the Service is pointing to another Ingress that doesn't configure this.

[ingress-paths]: ../../../user-guide/ingress-path-matching.md
[consistent-hashing]: https://en.wikipedia.org/wiki/Consistent_hashing
