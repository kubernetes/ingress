
# Hardening Guide

## Overview
There are several ways to do hardening and securing of nginx. In this documentation two guides are used, the guides are
overlapping in some points:

- [nginx CIS Benchmark](https://www.cisecurity.org/benchmark/nginx/)
- [cipherlist.eu](https://cipherlist.eu/) (one of many forks of the now dead project cipherli.st)

This guide describes, what of the different configurations described in those guides is already implemented as default
in the nginx implementation of kubernetes ingress, what needs to be configured, what is obsolete due to the fact that
the nginx is running as container (the CIS benchmark relates to a non-containerized installation) and what is difficult
or not possible.

Be aware that this is only a guide and you are responsible for your own implementation. Some of the configurations may
lead to have specific clients unable to reach your site or similar consequences.

This guide refers to chapters in the CIS Benchmark. For full explanation you should refer to the benchmark document itself

## Configuration Guide

| Chapter in CIS benchmark | Status | Default | Action to do if not default|
|:-------------------------|:-------|:--------|:---------------------------|
| __1 Initial Setup__ ||| |
| ||| |
| __1.1 Installation__||| |
| 1.1.1 Ensure NGINX is installed (Scored)| OK | done through helm charts / following documentation to deploy nginx ingress | |
| 1.1.2 Ensure NGINX is installed from source (Not Scored)| OK | done through helm charts / following documentation to deploy nginx ingress | |
| ||| |
| __1.2 Configure Software Updates__||| |
| 1.2.1 Ensure package manager repositories are properly configured (Not Scored) | OK | done via helm, nginx version could be overwritten, however compatibility is not ensured then| |
| 1.2.2 Ensure the latest software package is installed (Not Scored)| ACTION NEEDED | done via helm, nginx version could be overwritten, however compatibility is not ensured then| Plan for periodic updates |
| ||| |
| __2 Basic Configuration__ ||| |
| ||| |
| __2.1 Minimize NGINX Modules__||| |
| 2.1.1 Ensure only required modules are installed (Not Scored) | OK | Already only needed modules are installed, however proposals for further reduction are welcome | |
| 2.1.2 Ensure HTTP WebDAV module is not installed (Scored) | OK | | |
| 2.1.3 Ensure modules with gzip functionality are disabled (Scored)| OK | | |
| 2.1.4 Ensure the autoindex module is disabled (Scored)| OK | No autoindex configs so far in ingress defaults| |
| ||| |
| __2.2 Account Security__||| |
| 2.2.1 Ensure that NGINX is run using a non-privileged, dedicated service account (Not Scored) | OK | Pod configured as user www-data: [See this line in helm chart values](https://github.com/kubernetes/ingress-nginx/blob/0cbe783f43a9313c9c26136e888324b1ee91a72f/charts/ingress-nginx/values.yaml#L10). Compiled with user www-data: [See this line in build script](https://github.com/kubernetes/ingress-nginx/blob/5d67794f4fbf38ec6575476de46201b068eabf87/images/nginx/rootfs/build.sh#L529) | |
| 2.2.2 Ensure the NGINX service account is locked (Scored) | OK | Docker design ensures this | |
| 2.2.3 Ensure the NGINX service account has an invalid shell (Scored)| OK | Shell is nologin: [see this line in build script](https://github.com/kubernetes/ingress-nginx/blob/5d67794f4fbf38ec6575476de46201b068eabf87/images/nginx/rootfs/build.sh#L613)| |
| ||| |
| __2.3 Permissions and Ownership__ ||| |
| 2.3.1 Ensure NGINX directories and files are owned by root (Scored) | OK | Obsolete through docker-design and ingress controller needs to update the configs dynamically| |
| 2.3.2 Ensure access to NGINX directories and files is restricted (Scored) | OK | See previous answer| |
| 2.3.3 Ensure the NGINX process ID (PID) file is secured (Scored)| OK | No PID-File due to docker design | |
| 2.3.4 Ensure the core dump directory is secured (Not Scored)| OK | No working_directory configured by default | |
| ||| |
| __2.4 Network Configuration__ ||| |
| 2.4.1 Ensure NGINX only listens for network connections on authorized ports (Not Scored)| OK | Ensured by automatic nginx.conf configuration| |
| 2.4.2 Ensure requests for unknown host names are rejected (Not Scored)| OK | They are not rejected but send to the "default backend" delivering appropriate errors (mostly 404)| |
| 2.4.3 Ensure keepalive_timeout is 10 seconds or less, but not 0 (Scored)| ACTION NEEDED| Default is 75s | configure keep-alive to 10 seconds [according to this documentation](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#keep-alive) |
| 2.4.4 Ensure send_timeout is set to 10 seconds or less, but not 0 (Scored)| RISK TO BE ACCEPTED| Not configured, however the nginx default is 60s| Not configurable|
| ||| |
| __2.5 Information Disclosure__||| |
| 2.5.1 Ensure server_tokens directive is set to `off` (Scored) | OK | server_tokens is configured to off by default| |
| 2.5.2 Ensure default error and index.html pages do not reference NGINX (Scored) | ACTION NEEDED| 404 shows no version at all, 503 and 403 show "nginx", which is hardcoded [see this line in nginx source code](https://github.com/nginx/nginx/blob/master/src/http/ngx_http_special_response.c#L36) | configure custom error pages at least for 403, 404 and 503 and 500|
| 2.5.3 Ensure hidden file serving is disabled (Not Scored) | ACTION NEEDED | config not set | configure a config.server-snippet Snippet, but beware of .well-known challenges or similar. Refer to the benchmark here please |
| 2.5.4 Ensure the NGINX reverse proxy does not enable information disclosure (Scored)| ACTION NEEDED| hide not configured| configure hide-headers with array of "X-Powered-By" and "Server": [according to this documentation](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#hide-headers) |
| ||| |
| __3 Logging__ ||| |
| ||| |
| 3.1 Ensure detailed logging is enabled (Not Scored) | OK | nginx ingress has a very detailed log format by default | |
| 3.2 Ensure access logging is enabled (Scored) | OK | Access log is enabled by default | |
| 3.3 Ensure error logging is enabled and set to the info logging level (Scored)| OK | Error log is configured by default. The log level does not matter, because it is all sent to STDOUT anyway | |
| 3.4 Ensure log files are rotated (Scored) | OBSOLETE | Log file handling is not part of the nginx ingress and should be handled separately | |
| 3.5 Ensure error logs are sent to a remote syslog server (Not Scored) | OBSOLETE | See previous answer| |
| 3.6 Ensure access logs are sent to a remote syslog server (Not Scored)| OBSOLETE | See previous answer| |
| 3.7 Ensure proxies pass source IP information (Scored)| OK | Headers are set by default | |
| ||| |
| __4 Encryption__ ||| |
| ||| |
| __4.1 TLS / SSL Configuration__ ||| |
| 4.1.1 Ensure HTTP is redirected to HTTPS (Scored) | OK | Redirect to TLS is default | |
| 4.1.2 Ensure a trusted certificate and trust chain is installed (Not Scored)| ACTION NEEDED| For installing certs there are enough manuals in the web. A good way is to use lets encrypt through cert-manager | Install proper certificates or use lets encrypt with cert-manager |
| 4.1.3 Ensure private key permissions are restricted (Scored)| ACTION NEEDED| See previous answer| |
| 4.1.4 Ensure only modern TLS protocols are used (Scored)| OK/ACTION NEEDED | Default is TLS 1.2 + 1.3, while this is okay for CIS Benchmark, cipherlist.eu only recommends 1.3. This may cut off old OS's | Set controller.config.ssl-protocols to "TLSv1.3"|
| 4.1.5 Disable weak ciphers (Scored) | ACTION NEEDED| Default ciphers are already good, but cipherlist.eu recommends even stronger ciphers | Set controller.config.ssl-ciphers to "EECDH+AESGCM:EDH+AESGCM"|
| 4.1.6 Ensure custom Diffie-Hellman parameters are used (Scored) | ACTION NEEDED| No custom DH parameters are generated| Generate dh parameters for each ingress deployment you use - [see here for a how to](https://kubernetes.github.io/ingress-nginx/examples/customization/ssl-dh-param/) |
| 4.1.7 Ensure Online Certificate Status Protocol (OCSP) stapling is enabled (Scored) | ACTION NEEDED | Not enabled | set via [this configuration parameter](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#enable-ocsp) |
| 4.1.8 Ensure HTTP Strict Transport Security (HSTS) is enabled (Scored)| OK | HSTS is enabled by default | |
| 4.1.9 Ensure HTTP Public Key Pinning is enabled (Not Scored)| ACTION NEEDED / RISK TO BE ACCEPTED | HKPK not enabled by default | If lets encrypt is not used, set correct HPKP header. There are several ways to implement this - with the helm charts it works via controller.add-headers. If lets encrypt is used, this is complicated, a solution here is yet unknown |
| 4.1.10 Ensure upstream server traffic is authenticated with a client certificate (Scored) | DEPENDS ON BACKEND | Highly dependent on backends, not every backend allows configuring this, can also be mitigated via a service mesh| If backend allows it, [manual is here](https://kubernetes.github.io/ingress-nginx/examples/auth/client-certs/)|
| 4.1.11 Ensure the upstream traffic server certificate is trusted (Not Scored) | DEPENDS ON BACKEND | Highly dependent on backends, not every backend allows configuring this, can also be mitigated via a service mesh| If backend allows it, [see configuration here](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md#backend-certificate-authentication) |
| 4.1.12 Ensure your domain is preloaded (Not Scored) | ACTION NEEDED| Preload is not active by default | Set controller.config.hsts-preload to true|
| 4.1.13 Ensure session resumption is disabled to enable perfect forward security (Scored)| OK | Session tickets are disabled by default | |
| 4.1.14 Ensure HTTP/2.0 is used (Not Scored) | OK | http2 is set by default| |
| ||| |
| __5 Request Filtering and Restrictions__||| |
| ||| |
| __5.1 Access Control__||| |
| 5.1.1 Ensure allow and deny filters limit access to specific IP addresses (Not Scored)| OK/ACTION NEEDED | Depends on use case, geo ip module is compiled into nginx ingress controller, there are several ways to use it | If needed set IP restrictions via annotations or work with config snippets (be careful with lets-encrypt-http-challenge!) |
| 5.1.2 Ensure only whitelisted HTTP methods are allowed (Not Scored) | OK/ACTION NEEDED | Depends on use case| If required it can be set via config snippet|
| ||| |
| __5.2 Request Limits__||| |
| 5.2.1 Ensure timeout values for reading the client header and body are set correctly (Scored) | ACTION NEEDED| Default timeout is 60s | Set via [this configuration parameter](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#client-header-timeout) and respective body equivalent|
| 5.2.2 Ensure the maximum request body size is set correctly (Scored)| ACTION NEEDED| Default is 1m| set via [this configuration parameter](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#proxy-body-size)|
| 5.2.3 Ensure the maximum buffer size for URIs is defined (Scored) | ACTION NEEDED| Default is 4 8k| Set via [this configuration parameter](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/configmap.md#large-client-header-buffers)|
| 5.2.4 Ensure the number of connections per IP address is limited (Not Scored) | OK/ACTION NEEDED| No limit set| Depends on use case, limit can be set via [these annotations](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#rate-limiting)|
| 5.2.5 Ensure rate limits by IP address are set (Not Scored) | OK/ACTION NEEDED| No limit set| Depends on use case, limit can be set via [these annotations](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#rate-limiting)|
| ||| |
| __5.3 Browser Security__||| |
| 5.3.1 Ensure X-Frame-Options header is configured and enabled (Scored)| ACTION NEEDED| Header not set by default| Several ways to implement this - with the helm charts it works via controller.add-headers |
| 5.3.2 Ensure X-Content-Type-Options header is configured and enabled (Scored) | ACTION NEEDED| See previous answer| See previous answer |
| 5.3.3 Ensure the X-XSS-Protection Header is enabled and configured properly (Scored)| ACTION NEEDED| See previous answer| See previous answer |
| 5.3.4 Ensure that Content Security Policy (CSP) is enabled and configured properly (Not Scored) | ACTION NEEDED| See previous answer| See previous answer |
| 5.3.5 Ensure the Referrer Policy is enabled and configured properly (Not Scored)| ACTION NEEDED | Depends on application. It should be handled in the applications webserver itself, not in the load balancing ingress | check backend webserver |
| ||| |
| __6 Mandatory Access Control__| n/a| too high level, depends on backends | |

<style type="text/css" rel="stylesheet">
@media only screen and (min-width: 768px) {
	td:nth-child(1){
		white-space:normal !important;
    }

    .md-typeset table:not([class]) td {
        padding: .2rem .3rem;
    }
}
</style>
