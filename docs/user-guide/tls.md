# TLS/HTTPS

## TLS Secrets

Anytime we reference a TLS secret, we mean a PEM-encoded X.509, RSA (2048) secret.

You can generate a self-signed certificate and private key with:

```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${KEY_FILE} -out ${CERT_FILE} -subj "/CN=${HOST}/O=${HOST}"
```

Then create the secret in the cluster via:

```bash
kubectl create secret tls ${CERT_NAME} --key ${KEY_FILE} --cert ${CERT_FILE}
```

The resulting secret will be of type `kubernetes.io/tls`.

## Host names

Ensure that the relevant [ingress rules specify a matching host name](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls).

## Default SSL Certificate

NGINX provides the option to configure a server as a catch-all with
[server_name](http://nginx.org/en/docs/http/server_names.html)
for requests that do not match any of the configured server names.
This configuration works out-of-the-box for HTTP traffic.
For HTTPS, a certificate is naturally required.

For this reason the Ingress controller provides the flag `--default-ssl-certificate`.
The secret referred to by this flag contains the default certificate to be used when
accessing the catch-all server.
If this flag is not provided NGINX will use a self-signed certificate.

For instance, if you have a TLS secret `foo-tls` in the `default` namespace,
add `--default-ssl-certificate=default/foo-tls` in the `nginx-controller` deployment.

The default certificate will also be used for ingress `tls:` sections that do not
have a `secretName` option.

You may want your default Nginx server to not expose your
default-ssl-certificate to all requests.  You want any Ingress without a
`secretName` (or invalid `secretName`) to use the `default-ssl-certificate`,
but not serve that up to requests without a matching Server Name Indication.
For this situation, add `--default-server-uses-fake-certificate`. This will
keep the hostnames of your default certificate hidden to bots or attackers that
only know your IP address.


## SSL Passthrough

The [`--enable-ssl-passthrough`](cli-arguments.md) flag enables the SSL Passthrough feature, which is disabled by
default. This is required to enable passthrough backends in Ingress objects.

!!! warning
    This feature is implemented by intercepting **all traffic** on the configured HTTPS port (default: 443) and handing
    it over to a local TCP proxy. This bypasses NGINX completely and introduces a non-negligible performance penalty.

SSL Passthrough leverages [SNI][SNI] and reads the virtual domain from the TLS negotiation, which requires compatible
clients. After a connection has been accepted by the TLS listener, it is handled by the controller itself and piped back
and forth between the backend and the client.

If there is no hostname matching the requested host name, the request is handed over to NGINX on the configured
passthrough proxy port (default: 442), which proxies the request to the default backend.

!!! note
    Unlike HTTP backends, traffic to Passthrough backends is sent to the *clusterIP* of the backing Service instead of
    individual Endpoints.

## HTTP Strict Transport Security

HTTP Strict Transport Security (HSTS) is an opt-in security enhancement specified
through the use of a special response header. Once a supported browser receives
this header that browser will prevent any communications from being sent over
HTTP to the specified domain and will instead send all communications over HTTPS.

HSTS is enabled by default.

To disable this behavior use `hsts: "false"` in the configuration [ConfigMap][ConfigMap].

## Server-side HTTPS enforcement through redirect

By default the controller redirects HTTP clients to the HTTPS port
443 using a 308 Permanent Redirect response if TLS is enabled for that Ingress.

This can be disabled globally using `ssl-redirect: "false"` in the NGINX [config map][ConfigMap],
or per-Ingress with the `nginx.ingress.kubernetes.io/ssl-redirect: "false"`
annotation in the particular resource.

!!! tip
    When using SSL offloading outside of cluster (e.g. AWS ELB) it may be useful to enforce a
    redirect to HTTPS even when there is no TLS certificate available.
    This can be achieved by using the `nginx.ingress.kubernetes.io/force-ssl-redirect: "true"`
    annotation in the particular resource.

## Automated Certificate Management with Kube-Lego

!!! tip
    Kube-Lego has reached end-of-life and is being
    replaced by [cert-manager](https://github.com/jetstack/cert-manager/).

[Kube-Lego] automatically requests missing or expired certificates from [Let's Encrypt]
by monitoring ingress resources and their referenced secrets.

To enable this for an ingress resource you have to add an annotation:

```console
kubectl annotate ing ingress-demo kubernetes.io/tls-acme="true"
```

To setup Kube-Lego you can take a look at this [full example][full-kube-lego-example].
The first version to fully support Kube-Lego is Nginx Ingress controller 0.8.

## Default TLS Version and Ciphers

To provide the most secure baseline configuration possible,

nginx-ingress defaults to using TLS 1.2 and 1.3 only, with a [secure set of TLS ciphers][ssl-ciphers].

### Legacy TLS

The default configuration, though secure, does not support some older browsers and operating systems.

For instance, TLS 1.1+ is only enabled by default from Android 5.0 on. At the time of writing,
May 2018, [approximately 15% of Android devices](https://developer.android.com/about/dashboards/#Platform)
are not compatible with nginx-ingress's default configuration.

To change this default behavior, use a [ConfigMap][ConfigMap].

A sample ConfigMap fragment to allow these older clients to connect could look something like the following
(generated using the Mozilla SSL Configuration Generator)[mozilla-ssl-config-old]:

```
kind: ConfigMap
apiVersion: v1
metadata:
  name: nginx-config
data:
  ssl-ciphers: "ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:DHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES256-SHA256:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA"
  ssl-protocols: "TLSv1 TLSv1.1 TLSv1.2 TLSv1.3"
```



[full-kube-lego-example]:https://github.com/jetstack/kube-lego/tree/master/examples
[Kube-Lego]:https://github.com/jetstack/kube-lego
[Let's Encrypt]:https://letsencrypt.org
[ConfigMap]: ./nginx-configuration/configmap.md
[ssl-ciphers]: ./nginx-configuration/configmap.md#ssl-ciphers
[SNI]: https://en.wikipedia.org/wiki/Server_Name_Indication
[mozilla-ssl-config-old]: https://ssl-config.mozilla.org/#server=nginx&config=old
