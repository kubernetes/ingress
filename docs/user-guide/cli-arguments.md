# Command line arguments

The following command line arguments are accepted by the Ingress controller executable.

They are set in the container spec of the `nginx-ingress-controller` Deployment manifest

| Argument | Description |
|----------|-------------|
| `--add_dir_header`                 | If true, adds the file directory to the header |
| `--alsologtostderr`                | log to standard error as well as files |
| `--annotations-prefix`             | Prefix of the Ingress annotations specific to the NGINX controller. (default "nginx.ingress.kubernetes.io") |
| `--apiserver-host`                 | Address of the Kubernetes API server. Takes the form "protocol://address:port". If not specified, it is assumed the program runs inside a Kubernetes cluster and local discovery is attempted. |
| `--certificate-authority`          | Path to a cert file for the certificate authority. This certificate is used only when the flag --apiserver-host is specified. |
| `--configmap`                      | Name of the ConfigMap containing custom global configurations for the controller. |
| `--default-backend-service`        | Service used to serve HTTP requests not matching any known server name (catch-all). Takes the form "namespace/name". The controller configures NGINX to forward requests to the first port of this Service. |
| `--default-server-port`            | Port to use for exposing the default server (catch-all). (default 8181) |
| `--default-ssl-certificate`        | Secret containing a SSL certificate to be used by the default HTTPS server (catch-all). Takes the form "namespace/name". |
| `--default-server-uses-fake-certificate`        | The default Nginx server will serve the fake certificate instead of what is specified in `--default-ssl-certificate`|
| `--disable-catch-all`              | Disable support for catch-all Ingresses |
| `--election-id`                    | Election id to use for Ingress status updates. (default "ingress-controller-leader") |
| `--enable-metrics`                 | Enables the collection of NGINX metrics (default true) |
| `--enable-ssl-chain-completion`    | Autocomplete SSL certificate chains with missing intermediate CA certificates. Certificates uploaded to Kubernetes must have the "Authority Information Access" X.509 v3 extension for this to succeed. |
| `--enable-ssl-passthrough`         | Enable SSL Passthrough. |
| `--health-check-path`              | URL path of the health check endpoint. Configured inside the NGINX status server. All requests received on the port defined by the healthz-port parameter are forwarded internally to this path. (default "/healthz") |
| `--health-check-timeout`           | Time limit, in seconds, for a probe to health-check-path to succeed. (default 10) |
| `--healthz-port`                   | Port to use for the healthz endpoint. (default 10254) |
| `--http-port`                      | Port to use for servicing HTTP traffic. (default 80) |
| `--https-port`                     | Port to use for servicing HTTPS traffic. (default 443) |
| `--ingress-class`                  | Name of the ingress class this controller satisfies. The class of an Ingress object is set using the field IngressClassName in Kubernetes clusters version v1.18.0 or higher or the annotation "kubernetes.io/ingress.class" (deprecated). If this parameter is not set, or set to the default value of "nginx", it will handle ingresses with either an empty or "nginx" class name. |
| `--kubeconfig`                     | Path to a kubeconfig file containing authorization and API server information. |
| `--log_backtrace_at`               | when logging hits line file:N, emit a stack trace (default :0) |
| `--log_dir`                        | If non-empty, write log files in this directory |
| `--log_file`                       | If non-empty, use this log file |
| `--log_file_max_size`              | Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited. (default 1800) |
| `--logtostderr`                    | log to standard error instead of files (default true) |
| `--maxmind-edition-ids`            | Maxmind edition ids to download GeoLite2 Databases. (default "GeoLite2-City,GeoLite2-ASN") |
| `--maxmind-license-key`            | Maxmind license key to download GeoLite2 Databases. https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases |
| `--metrics-per-host`               | Export metrics per-host (default true) |
| `--profiler-port`                  | Port to use for expose the ingress controller Go profiler when it is enabled. (default 10245) |
| `--profiling`                      | Enable profiling via web interface host:port/debug/pprof/ (default true) |
| `--publish-service`                | Service fronting the Ingress controller. Takes the form "namespace/name". When used together with update-status, the controller mirrors the address of this service's endpoints to the load-balancer status of all Ingress objects it satisfies. |
| `--publish-status-address`         | Customized address (or addresses, separated by comma) to set as the load-balancer status of Ingress objects this controller satisfies. Requires the update-status parameter. |
| `--report-node-internal-ip-address`| Set the load-balancer status of Ingress objects to internal Node addresses instead of external. Requires the update-status parameter. |
| `--skip_headers`                   | If true, avoid header prefixes in the log messages |
| `--skip_log_headers`               | If true, avoid headers when opening log files |
| `--ssl-passthrough-proxy-port`     | Port to use internally for SSL Passthrough. (default 442) |
| `--status-port`                    | Port to use for the lua HTTP endpoint configuration. (default 10246) |
| `--status-update-interval`         | Time interval in seconds in which the status should check if an update is required. Default is 60 seconds (default 60) |
| `--stderrthreshold`                | logs at or above this threshold go to stderr (default 2) |
| `--stream-port`                    | Port to use for the lua TCP/UDP endpoint configuration. (default 10247) |
| `--sync-period`                    | Period at which the controller forces the repopulation of its local object stores. Disabled by default. |
| `--sync-rate-limit`                | Define the sync frequency upper limit (default 0.3) |
| `--tcp-services-configmap`         | Name of the ConfigMap containing the definition of the TCP services to expose. The key in the map indicates the external port to be used. The value is a reference to a Service in the form "namespace/name:port", where "port" can either be a port number or name. TCP ports 80 and 443 are reserved by the controller for servicing HTTP traffic. |
| `--udp-services-configmap`         | Name of the ConfigMap containing the definition of the UDP services to expose. The key in the map indicates the external port to be used. The value is a reference to a Service in the form "namespace/name:port", where "port" can either be a port name or number. |
| `--update-status`                  | Update the load-balancer status of Ingress objects this controller satisfies. Requires setting the publish-service parameter to a valid Service reference. (default true) |
| `--update-status-on-shutdown`      | Update the load-balancer status of Ingress objects when the controller shuts down. Requires the update-status parameter. (default true) |
| `--shutdown-grace-period`          | Seconds to wait after receiving the shutdown signal, before stopping the nginx process. |
| `-v, --v Level`                    | number for the log level verbosity |
| `--validating-webhook`             | The address to start an admission controller on to validate incoming ingresses. Takes the form "<host>:port". If not provided, no admission controller is started. |
| `--validating-webhook-certificate` | The path of the validating webhook certificate PEM. |
| `--validating-webhook-key`         | The path of the validating webhook key PEM. |
| `--version`                        | Show release information about the NGINX Ingress controller and exit. |
| `--vmodule`                        | comma-separated list of pattern=N settings for file-filtered logging |
| `--watch-namespace`                | Namespace the controller watches for updates to Kubernetes objects. This includes Ingresses, Services and all configuration resources. All namespaces are watched if this parameter is left empty. |
