# How it works

The objective of this document is to explain how the NGINX Ingress controller works, in particular how the NGINX model is built and why we need one.

## NGINX configuration

The goal of this Ingress controller is the assembly of a configuration file (nginx.conf). The main implication of this requirement is the need to reload NGINX after any change in the configuration file. _Though it is important to note that we don't reload Nginx on changes that impact only an `upstream` configuration (i.e Endpoints change when you deploy your app)_. We use [lua-nginx-module](https://github.com/openresty/lua-nginx-module) to achieve this. Check [below](#avoiding-reloads-on-endpoints-changes) to learn more about how it's done.

## NGINX model

Usually, a Kubernetes Controller utilizes the [synchronization loop pattern][1] to check if the desired state in the controller is updated or a change is required. To this purpose, we need to build a model using different objects from the cluster, in particular (in no special order) Ingresses, Services, Endpoints, Secrets, and Configmaps to generate a point in time configuration file that reflects the state of the cluster.

To get this object from the cluster, we use [Kubernetes Informers][2], in particular, `FilteredSharedInformer`. This informers allows reacting to changes in using [callbacks][3] to individual changes when a new object is added, modified or removed. Unfortunately, there is no way to know if a particular change is going to affect the final configuration file. Therefore on every change, we have to rebuild a new model from scratch based on the state of cluster and compare it to the current model. If the new model equals to the current one, then we avoid generating a new NGINX configuration and triggering a reload. Otherwise, we check if the difference is only about Endpoints. If so we then send the new list of Endpoints to a Lua handler running inside Nginx using HTTP POST request and again avoid generating a new NGINX configuration and triggering a reload. If the difference between running and new model is about more than just Endpoints we create a new NGINX configuration based on the new model, replace the current model and trigger a reload.

One of the uses of the model is to avoid unnecessary reloads when there's no change in the state and to detect conflicts in definitions.

The final representation of the NGINX configuration is generated from a [Go template][6] using the new model as input for the variables required by the template.

## Building the NGINX model

Building a model is an expensive operation, for this reason, the use of the synchronization loop is a must. By using a [work queue][4] it is possible to not lose changes and remove the use of [sync.Mutex][5] to force a single execution of the sync loop and additionally it is possible to create a time window between the start and end of the sync loop that allows us to discard unnecessary updates. It is important to understand that any change in the cluster could generate events that the informer will send to the controller and one of the reasons for the [work queue][4].

Operations to build the model:

- Order Ingress rules by `CreationTimestamp` field, i.e., old rules first.

  - If the same path for the same host is defined in more than one Ingress, the oldest rule wins.
  - If more than one Ingress contains a TLS section for the same host, the oldest rule wins.
  - If multiple Ingresses define an annotation that affects the configuration of the Server block, the oldest rule wins.

- Create a list of NGINX Servers (per hostname)
- Create a list of NGINX Upstreams
- If multiple Ingresses define different paths for the same host, the ingress controller will merge the definitions.
- Annotations are applied to all the paths in the Ingress.
- Multiple Ingresses can define different annotations. These definitions are not shared between Ingresses.

## When a reload is required

The next list describes the scenarios when a reload is required:

- New Ingress Resource Created.
- TLS section is added to existing Ingress.
- Change in Ingress annotations that impacts more than just upstream configuration. For instance `load-balance` annotation does not require a reload.
- A path is added/removed from an Ingress.
- An Ingress, Service, Secret is removed.
- Some missing referenced object from the Ingress is available, like a Service or Secret.
- A Secret is updated.

## Avoiding reloads

In some cases, it is possible to avoid reloads, in particular when there is a change in the endpoints, i.e., a pod is started or replaced. It is out of the scope of this Ingress controller to remove reloads completely. This would require an incredible amount of work and at some point makes no sense. This can change only if NGINX changes the way new configurations are read, basically, new changes do not replace worker processes.

### Avoiding reloads on Endpoints changes

On every endpoint change the controller fetches endpoints from all the services it sees and generates corresponding Backend objects. It then sends these objects to a Lua handler running inside Nginx. The Lua code in turn stores those backends in a shared memory zone. Then for every request Lua code running in [`balancer_by_lua`](https://github.com/openresty/lua-resty-core/blob/master/lib/ngx/balancer.md) context detects what endpoints it should choose upstream peer from and applies the configured load balancing algorithm to choose the peer. Then Nginx takes care of the rest. This way we avoid reloading Nginx on endpoint changes. _Note_ that this includes annotation changes that affects only `upstream` configuration in Nginx as well.

In a relatively big cluster with frequently deploying apps this feature saves significant number of Nginx reloads which can otherwise affect response latency, load balancing quality (after every reload Nginx resets the state of load balancing) and so on.

### Avoiding outage from wrong configuration

Because the ingress controller works using the [synchronization loop pattern](https://coreos.com/kubernetes/docs/latest/replication-controller.html#the-reconciliation-loop-in-detail), it is applying the configuration for all matching objects. In case some Ingress objects have a broken configuration, for example a syntax error in the `nginx.ingress.kubernetes.io/configuration-snippet` annotation, the generated configuration becomes invalid, does not reload and hence no more ingresses will be taken into account.

To prevent this situation to happen, the nginx ingress controller optionally exposes a [validating admission webhook server][8] to ensure the validity of incoming ingress objects.
This webhook appends the incoming ingress objects to the list of ingresses, generates the configuration and calls nginx to ensure the configuration has no syntax errors.

[0]: https://github.com/openresty/lua-nginx-module/pull/1259
[1]: https://coreos.com/kubernetes/docs/latest/replication-controller.html#the-reconciliation-loop-in-detail
[2]: https://godoc.org/k8s.io/client-go/informers#NewFilteredSharedInformerFactory
[3]: https://godoc.org/k8s.io/client-go/tools/cache#ResourceEventHandlerFuncs
[4]: https://github.com/kubernetes/ingress-nginx/blob/main/internal/task/queue.go#L38
[5]: https://golang.org/pkg/sync/#Mutex
[6]: https://github.com/kubernetes/ingress-nginx/blob/main/rootfs/etc/nginx/template/nginx.tmpl
[7]: http://nginx.org/en/docs/beginners_guide.html#control
[8]: https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#validatingadmissionwebhook
