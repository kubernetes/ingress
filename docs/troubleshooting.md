<!--
-----------------NOTICE------------------------
This file is referenced in code as
https://github.com/kubernetes/ingress-nginx/blob/master/docs/troubleshooting.md
Do not move it without providing redirects.
-----------------------------------------------
-->

# Troubleshooting

## Ingress-Controller Logs and Events

There are many ways to troubleshoot the ingress-controller. The following are basic troubleshooting
methods to obtain more information.

Check the Ingress Resource Events

```console
$ kubectl get ing -n <namespace-of-ingress-resource>
NAME           HOSTS      ADDRESS     PORTS     AGE
cafe-ingress   cafe.com   10.0.2.15   80        25s

$ kubectl describe ing <ingress-resource-name> -n <namespace-of-ingress-resource>
Name:             cafe-ingress
Namespace:        default
Address:          10.0.2.15
Default backend:  default-http-backend:80 (172.17.0.5:8080)
Rules:
  Host      Path  Backends
  ----      ----  --------
  cafe.com
            /tea      tea-svc:80 (<none>)
            /coffee   coffee-svc:80 (<none>)
Annotations:
  kubectl.kubernetes.io/last-applied-configuration:  {"apiVersion":"networking.k8s.io/v1beta1","kind":"Ingress","metadata":{"annotations":{},"name":"cafe-ingress","namespace":"default","selfLink":"/apis/networking/v1beta1/namespaces/default/ingresses/cafe-ingress"},"spec":{"rules":[{"host":"cafe.com","http":{"paths":[{"backend":{"serviceName":"tea-svc","servicePort":80},"path":"/tea"},{"backend":{"serviceName":"coffee-svc","servicePort":80},"path":"/coffee"}]}}]},"status":{"loadBalancer":{"ingress":[{"ip":"169.48.142.110"}]}}}

Events:
  Type    Reason  Age   From                      Message
  ----    ------  ----  ----                      -------
  Normal  CREATE  1m    nginx-ingress-controller  Ingress default/cafe-ingress
  Normal  UPDATE  58s   nginx-ingress-controller  Ingress default/cafe-ingress
```

Check the Ingress Controller Logs

```console
$ kubectl get pods -n <namespace-of-ingress-controller>
NAME                                        READY     STATUS    RESTARTS   AGE
nginx-ingress-controller-67956bf89d-fv58j   1/1       Running   0          1m

$ kubectl logs -n <namespace> nginx-ingress-controller-67956bf89d-fv58j
-------------------------------------------------------------------------------
NGINX Ingress controller
  Release:    0.14.0
  Build:      git-734361d
  Repository: https://github.com/kubernetes/ingress-nginx
-------------------------------------------------------------------------------
....
```

Check the Nginx Configuration

```console
$ kubectl get pods -n <namespace-of-ingress-controller>
NAME                                        READY     STATUS    RESTARTS   AGE
nginx-ingress-controller-67956bf89d-fv58j   1/1       Running   0          1m

$ kubectl exec -it -n <namespace-of-ingress-controller> nginx-ingress-controller-67956bf89d-fv58j -- cat /etc/nginx/nginx.conf
daemon off;
worker_processes 2;
pid /run/nginx.pid;
worker_rlimit_nofile 523264;
worker_shutdown_timeout 240s;
events {
	multi_accept        on;
	worker_connections  16384;
	use                 epoll;
}
http {
....
```

Check if used Services Exist

```console
$ kubectl get svc --all-namespaces
NAMESPACE     NAME                   TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
default       coffee-svc             ClusterIP   10.106.154.35    <none>        80/TCP          18m
default       kubernetes             ClusterIP   10.96.0.1        <none>        443/TCP         30m
default       tea-svc                ClusterIP   10.104.172.12    <none>        80/TCP          18m
kube-system   default-http-backend   NodePort    10.108.189.236   <none>        80:30001/TCP    30m
kube-system   kube-dns               ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP   30m
kube-system   kubernetes-dashboard   NodePort    10.103.128.17    <none>        80:30000/TCP    30m
```

## Debug Logging

Using the flag `--v=XX` it is possible to increase the level of logging. This is performed by editing
the deployment.

```console
$ kubectl get deploy -n <namespace-of-ingress-controller>
NAME                       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
default-http-backend       1         1         1            1           35m
nginx-ingress-controller   1         1         1            1           35m

$ kubectl edit deploy -n <namespace-of-ingress-controller> nginx-ingress-controller
# Add --v=X to "- args", where X is an integer
```

- `--v=2` shows details using `diff` about the changes in the configuration in nginx
- `--v=3` shows details about the service, Ingress rule, endpoint changes and it dumps the nginx configuration in JSON format
- `--v=5` configures NGINX in [debug mode](http://nginx.org/en/docs/debugging_log.html)

## Authentication to the Kubernetes API Server

A number of components are involved in the authentication process and the first step is to narrow
down the source of the problem, namely whether it is a problem with service authentication or
with the kubeconfig file.

Both authentications must work:

```
+-------------+   service          +------------+
|             |   authentication   |            |
+  apiserver  +<-------------------+  ingress   |
|             |                    | controller |
+-------------+                    +------------+
```

**Service authentication**

The Ingress controller needs information from apiserver. Therefore, authentication is required, which can be achieved in two different ways:

1. _Service Account:_ This is recommended, because nothing has to be configured. The Ingress controller will use information provided by the system to communicate with the API server. See 'Service Account' section for details.

2. _Kubeconfig file:_ In some Kubernetes environments service accounts are not available. In this case a manual configuration is required. The Ingress controller binary can be started with the `--kubeconfig` flag. The value of the flag is a path to a file specifying how to connect to the API server. Using the `--kubeconfig` does not requires the flag `--apiserver-host`.
   The format of the file is identical to `~/.kube/config` which is used by kubectl to connect to the API server. See 'kubeconfig' section for details.

3. _Using the flag `--apiserver-host`:_ Using this flag `--apiserver-host=http://localhost:8080` it is possible to specify an unsecured API server or reach a remote kubernetes cluster using [kubectl proxy](https://kubernetes.io/docs/user-guide/kubectl/kubectl_proxy/).
   Please do not use this approach in production.

In the diagram below you can see the full authentication flow with all options, starting with the browser
on the lower left hand side.

```
Kubernetes                                                  Workstation
+---------------------------------------------------+     +------------------+
|                                                   |     |                  |
|  +-----------+   apiserver        +------------+  |     |  +------------+  |
|  |           |   proxy            |            |  |     |  |            |  |
|  | apiserver |                    |  ingress   |  |     |  |  ingress   |  |
|  |           |                    | controller |  |     |  | controller |  |
|  |           |                    |            |  |     |  |            |  |
|  |           |                    |            |  |     |  |            |  |
|  |           |  service account/  |            |  |     |  |            |  |
|  |           |  kubeconfig        |            |  |     |  |            |  |
|  |           +<-------------------+            |  |     |  |            |  |
|  |           |                    |            |  |     |  |            |  |
|  +------+----+      kubeconfig    +------+-----+  |     |  +------+-----+  |
|         |<--------------------------------------------------------|        |
|                                                   |     |                  |
+---------------------------------------------------+     +------------------+
```

### Service Account

If using a service account to connect to the API server, the ingress-controller expects the file
`/var/run/secrets/kubernetes.io/serviceaccount/token` to be present. It provides a secret
token that is required to authenticate with the API server.

Verify with the following commands:

```console
# start a container that contains curl
$ kubectl run -it --rm test --image=curlimages/curl --restart=Never -- /bin/sh

# check if secret exists
/ $ ls /var/run/secrets/kubernetes.io/serviceaccount/
ca.crt     namespace  token
/ $

# check base connectivity from cluster inside
/ $ curl -k https://kubernetes.default.svc.cluster.local
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {

  },
  "status": "Failure",
  "message": "forbidden: User \"system:anonymous\" cannot get path \"/\"",
  "reason": "Forbidden",
  "details": {

  },
  "code": 403
}/ $

# connect using tokens
}/ $ curl --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt -H  "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" https://kubernetes.default.svc.cluster.local
&& echo
{
  "paths": [
    "/api",
    "/api/v1",
    "/apis",
    "/apis/",
    ... TRUNCATED
    "/readyz/shutdown",
    "/version"
  ]
}
/ $

# when you type `exit` or `^D` the test pod will be deleted.
```

If it is not working, there are two possible reasons:

1. The contents of the tokens are invalid. Find the secret name with `kubectl get secrets | grep service-account` and
   delete it with `kubectl delete secret <name>`. It will automatically be recreated.

2. You have a non-standard Kubernetes installation and the file containing the token may not be present.
   The API server will mount a volume containing this file, but only if the API server is configured to use
   the ServiceAccount admission controller.
   If you experience this error, verify that your API server is using the ServiceAccount admission controller.
   If you are configuring the API server by hand, you can set this with the `--admission-control` parameter.
   > Note that you should use other admission controllers as well. Before configuring this option, you should read about admission controllers.

More information:

- [User Guide: Service Accounts](http://kubernetes.io/docs/user-guide/service-accounts/)
- [Cluster Administrator Guide: Managing Service Accounts](http://kubernetes.io/docs/admin/service-accounts-admin/)

## Kube-Config

If you want to use a kubeconfig file for authentication, follow the [deploy procedure](deploy/index.md) and
add the flag `--kubeconfig=/etc/kubernetes/kubeconfig.yaml` to the args section of the deployment.

## Using GDB with Nginx

[Gdb](https://www.gnu.org/software/gdb/) can be used to with nginx to perform a configuration
dump. This allows us to see which configuration is being used, as well as older configurations.

Note: The below is based on the nginx [documentation](https://docs.nginx.com/nginx/admin-guide/monitoring/debugging/#dumping-nginx-configuration-from-a-running-process).

1. SSH into the worker

```console
$ ssh user@workerIP
```

2. Obtain the Docker Container Running nginx

```console
$ docker ps | grep nginx-ingress-controller
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
d9e1d243156a        quay.io/kubernetes-ingress-controller/nginx-ingress-controller   "/usr/bin/dumb-init …"   19 minutes ago      Up 19 minutes                                                                            k8s_nginx-ingress-controller_nginx-ingress-controller-67956bf89d-mqxzt_kube-system_079f31ec-aa37-11e8-ad39-080027a227db_0
```

3. Exec into the container

```console
$ docker exec -it --user=0 --privileged d9e1d243156a bash
```

4. Make sure nginx is running in `--with-debug`

```console
$ nginx -V 2>&1 | grep -- '--with-debug'
```

5. Get list of processes running on container

```console
$ ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 20:23 ?        00:00:00 /usr/bin/dumb-init /nginx-ingres
root         5     1  0 20:23 ?        00:00:05 /nginx-ingress-controller --defa
root        21     5  0 20:23 ?        00:00:00 nginx: master process /usr/sbin/
nobody     106    21  0 20:23 ?        00:00:00 nginx: worker process
nobody     107    21  0 20:23 ?        00:00:00 nginx: worker process
root       172     0  0 20:43 pts/0    00:00:00 bash
```

7. Attach gdb to the nginx master process

```console
$ gdb -p 21
....
Attaching to process 21
Reading symbols from /usr/sbin/nginx...done.
....
(gdb)
```

8. Copy and paste the following:

```console
set $cd = ngx_cycle->config_dump
set $nelts = $cd.nelts
set $elts = (ngx_conf_dump_t*)($cd.elts)
while ($nelts-- > 0)
set $name = $elts[$nelts]->name.data
printf "Dumping %s to nginx_conf.txt\n", $name
append memory nginx_conf.txt \
        $elts[$nelts]->buffer.start $elts[$nelts]->buffer.end
end
```

9. Quit GDB by pressing CTRL+D

10. Open nginx_conf.txt

```console
cat nginx_conf.txt
```
