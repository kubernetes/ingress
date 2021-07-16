<!--
-----------------NOTICE------------------------
This file is referenced in code as
https://github.com/kubernetes/ingress-nginx/blob/main/docs/kubectl-plugin.md
Do not move it without providing redirects.
-----------------------------------------------
-->

# The ingress-nginx kubectl plugin

## Installation

Install [krew](https://github.com/GoogleContainerTools/krew), then run

```console
kubectl krew install ingress-nginx
```

to install the plugin. Then run

```console
kubectl ingress-nginx --help
```

to make sure the plugin is properly installed and to get a list of commands:

```console
kubectl ingress-nginx --help
A kubectl plugin for inspecting your ingress-nginx deployments

Usage:
  ingress-nginx [command]

Available Commands:
  backends    Inspect the dynamic backend information of an ingress-nginx instance
  certs       Output the certificate data stored in an ingress-nginx pod
  conf        Inspect the generated nginx.conf
  exec        Execute a command inside an ingress-nginx pod
  general     Inspect the other dynamic ingress-nginx information
  help        Help about any command
  info        Show information about the ingress-nginx service
  ingresses   Provide a short summary of all of the ingress definitions
  lint        Inspect kubernetes resources for possible issues
  logs        Get the kubernetes logs for an ingress-nginx pod
  ssh         ssh into a running ingress-nginx pod

Flags:
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default HTTP cache directory (default "/Users/alexkursell/.kube/http-cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -h, --help                           help for ingress-nginx
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use

Use "ingress-nginx [command] --help" for more information about a command.
```

## Common Flags

- Every subcommand supports the basic `kubectl` configuration flags like `--namespace`, `--context`, `--client-key` and so on.
- Subcommands that act on a particular `ingress-nginx` pod (`backends`, `certs`, `conf`, `exec`, `general`, `logs`, `ssh`), support the `--deployment <deployment>` and `--pod <pod>` flags to select either a pod from a deployment with the given name, or a pod with the given name. The `--deployment` flag defaults to `nginx-ingress-controller`.
- Subcommands that inspect resources (`ingresses`, `lint`) support the `--all-namespaces` flag, which causes them to inspect resources in every namespace.

## Subcommands

Note that `backends`, `general`, `certs`, and `conf` require `ingress-nginx` version `0.23.0` or higher.

### backends

Run `kubectl ingress-nginx backends` to get a JSON array of the backends that an ingress-nginx controller currently knows about:

```console
$ kubectl ingress-nginx backends -n ingress-nginx
[
  {
    "name": "default-apple-service-5678",
    "service": {
      "metadata": {
        "creationTimestamp": null
      },
      "spec": {
        "ports": [
          {
            "protocol": "TCP",
            "port": 5678,
            "targetPort": 5678
          }
        ],
        "selector": {
          "app": "apple"
        },
        "clusterIP": "10.97.230.121",
        "type": "ClusterIP",
        "sessionAffinity": "None"
      },
      "status": {
        "loadBalancer": {}
      }
    },
    "port": 0,
    "sslPassthrough": false,
    "endpoints": [
      {
        "address": "10.1.3.86",
        "port": "5678"
      }
    ],
    "sessionAffinityConfig": {
      "name": "",
      "cookieSessionAffinity": {
        "name": ""
      }
    },
    "upstreamHashByConfig": {
      "upstream-hash-by-subset-size": 3
    },
    "noServer": false,
    "trafficShapingPolicy": {
      "weight": 0,
      "header": "",
      "headerValue": "",
      "cookie": ""
    }
  },
  {
    "name": "default-echo-service-8080",
    ...
  },
  {
    "name": "upstream-default-backend",
    ...
  }
]
```

Add the `--list` option to show only the backend names. Add the `--backend <backend>` option to show only the backend with the given name.

### certs

Use `kubectl ingress-nginx certs --host <hostname>` to dump the SSL cert/key information for a given host.

**WARNING:** This command will dump sensitive private key information. Don't blindly share the output, and certainly don't log it anywhere.

```console
$ kubectl ingress-nginx certs -n ingress-nginx --host testaddr.local
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----

-----BEGIN RSA PRIVATE KEY-----
<REDACTED! DO NOT SHARE THIS!>
-----END RSA PRIVATE KEY-----
```

### conf

Use `kubectl ingress-nginx conf` to dump the generated `nginx.conf` file. Add the `--host <hostname>` option to view only the server block for that host:

```console
kubectl ingress-nginx conf -n ingress-nginx --host testaddr.local

	server {
		server_name testaddr.local ;

		listen 80;

		set $proxy_upstream_name "-";
		set $pass_access_scheme $scheme;
		set $pass_server_port $server_port;
		set $best_http_host $http_host;
		set $pass_port $pass_server_port;

		location / {

			set $namespace      "";
			set $ingress_name   "";
			set $service_name   "";
			set $service_port   "0";
			set $location_path  "/";

...
```

### exec

`kubectl ingress-nginx exec` is exactly the same as `kubectl exec`, with the same command flags. It will automatically choose an `ingress-nginx` pod to run the command in.

```console
$ kubectl ingress-nginx exec -i -n ingress-nginx -- ls /etc/nginx
fastcgi_params
geoip
lua
mime.types
modsecurity
modules
nginx.conf
opentracing.json
owasp-modsecurity-crs
template
```

### info

Shows the internal and external IP/CNAMES for an `ingress-nginx` service.

```console
$ kubectl ingress-nginx info -n ingress-nginx
Service cluster IP address: 10.187.253.31
LoadBalancer IP|CNAME: 35.123.123.123
```

Use the `--service <service>` flag if your `ingress-nginx` `LoadBalancer` service is not named `ingress-nginx`.

### ingresses

`kubectl ingress-nginx ingresses`, alternately `kubectl ingress-nginx ing`, shows a more detailed view of the ingress definitions in a namespace. Compare:

```console
$ kubectl get ingresses --all-namespaces
NAMESPACE   NAME               HOSTS                            ADDRESS     PORTS   AGE
default     example-ingress1   testaddr.local,testaddr2.local   localhost   80      5d
default     test-ingress-2     *                                localhost   80      5d
```

vs

```console
$ kubectl ingress-nginx ingresses --all-namespaces
NAMESPACE   INGRESS NAME       HOST+PATH                        ADDRESSES   TLS   SERVICE         SERVICE PORT   ENDPOINTS
default     example-ingress1   testaddr.local/etameta           localhost   NO    pear-service    5678           5
default     example-ingress1   testaddr2.local/otherpath        localhost   NO    apple-service   5678           1
default     example-ingress1   testaddr2.local/otherotherpath   localhost   NO    pear-service    5678           5
default     test-ingress-2     *                                localhost   NO    echo-service    8080           2
```

### lint

`kubectl ingress-nginx lint` can check a namespace or entire cluster for potential configuration issues. This command is especially useful when upgrading between `ingress-nginx` versions.

```console
$ kubectl ingress-nginx lint --all-namespaces --verbose
Checking ingresses...
✗ anamespace/this-nginx
  - Contains the removed session-cookie-hash annotation.
       Lint added for version 0.24.0
       https://github.com/kubernetes/ingress-nginx/issues/3743
✗ othernamespace/ingress-definition-blah
  - The rewrite-target annotation value does not reference a capture group
      Lint added for version 0.22.0
      https://github.com/kubernetes/ingress-nginx/issues/3174

Checking deployments...
✗ namespace2/nginx-ingress-controller
  - Uses removed config flag --sort-backends
      Lint added for version 0.22.0
      https://github.com/kubernetes/ingress-nginx/issues/3655
  - Uses removed config flag --enable-dynamic-certificates
      Lint added for version 0.24.0
      https://github.com/kubernetes/ingress-nginx/issues/3808
```

to show the lints added **only** for a particular `ingress-nginx` release, use the `--from-version` and `--to-version` flags:

```console
$ kubectl ingress-nginx lint --all-namespaces --verbose --from-version 0.24.0 --to-version 0.24.0
Checking ingresses...
✗ anamespace/this-nginx
  - Contains the removed session-cookie-hash annotation.
       Lint added for version 0.24.0
       https://github.com/kubernetes/ingress-nginx/issues/3743

Checking deployments...
✗ namespace2/nginx-ingress-controller
  - Uses removed config flag --enable-dynamic-certificates
      Lint added for version 0.24.0
      https://github.com/kubernetes/ingress-nginx/issues/3808
```

### logs

`kubectl ingress-nginx logs` is almost the same as `kubectl logs`, with fewer flags. It will automatically choose an `ingress-nginx` pod to read logs from.

```console
$ kubectl ingress-nginx logs -n ingress-nginx
-------------------------------------------------------------------------------
NGINX Ingress controller
  Release:    dev
  Build:      git-48dc3a867
  Repository: git@github.com:kubernetes/ingress-nginx.git
-------------------------------------------------------------------------------

W0405 16:53:46.061589       7 flags.go:214] SSL certificate chain completion is disabled (--enable-ssl-chain-completion=false)
nginx version: nginx/1.15.9
W0405 16:53:46.070093       7 client_config.go:549] Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.
I0405 16:53:46.070499       7 main.go:205] Creating API client for https://10.96.0.1:443
I0405 16:53:46.077784       7 main.go:249] Running in Kubernetes cluster version v1.10 (v1.10.11) - git (clean) commit 637c7e288581ee40ab4ca210618a89a555b6e7e9 - platform linux/amd64
I0405 16:53:46.183359       7 nginx.go:265] Starting NGINX Ingress controller
I0405 16:53:46.193913       7 event.go:209] Event(v1.ObjectReference{Kind:"ConfigMap", Namespace:"ingress-nginx", Name:"udp-services", UID:"82258915-563e-11e9-9c52-025000000001", APIVersion:"v1", ResourceVersion:"494", FieldPath:""}): type: 'Normal' reason: 'CREATE' ConfigMap ingress-nginx/udp-services
...
```

### ssh

`kubectl ingress-nginx ssh` is exactly the same as `kubectl ingress-nginx exec -it -- /bin/bash`. Use it when you want to quickly be dropped into a shell inside a running `ingress-nginx` container.

```console
$ kubectl ingress-nginx ssh -n ingress-nginx
www-data@nginx-ingress-controller-7cbf77c976-wx5pn:/etc/nginx$
```
