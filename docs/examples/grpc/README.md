# gRPC

This example demonstrates how to route traffic to a gRPC service through the
nginx controller.

## Prerequisites

1. You have a kubernetes cluster running.
2. You have a domain name such as `example.com` that is configured to route
   traffic to the ingress controller.  Replace references to
   `fortune-teller.stack.build` (the domain name used in this example) to your
   own domain name (you're also responsible for provisioning an SSL certificate
   for the ingress).
3. You have the nginx-ingress controller installed in typical fashion (must be
   at least
   [quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0](https://quay.io/kubernetes-ingress-controller/nginx-ingress-controller)
   for grpc support.
4. You have a backend application running a gRPC server and listening for TCP
   traffic.  If you prefer, you can use the
   [fortune-teller](https://github.com/kubernetes/ingress-nginx/tree/main/images/grpc-fortune-teller)
   application provided here as an example.

### Step 1: kubernetes `Deployment`

```sh
$ kubectl create -f app.yaml
```

This is a standard kubernetes deployment object.  It is running a grpc service
listening on port `50051`.

The sample application
[fortune-teller-app](https://github.com/kubernetes/ingress-nginx/tree/main/images/grpc-fortune-teller)
is a grpc server implemented in go. Here's the stripped-down implementation:

```go
func main() {
	grpcServer := grpc.NewServer()
	fortune.RegisterFortuneTellerServer(grpcServer, &FortuneTeller{})
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer.Serve(lis)
}
```

The takeaway is that we are not doing any TLS configuration on the server (as we
are terminating TLS at the ingress level, grpc traffic will travel unencrypted
inside the cluster and arrive "insecure").

For your own application you may or may not want to do this.  If you prefer to
forward encrypted traffic to your POD and terminate TLS at the gRPC server
itself, add the ingress annotation `nginx.ingress.kubernetes.io/backend-protocol: "GRPCS"`.

### Step 2: the kubernetes `Service`

```sh
$ kubectl create -f svc.yaml
```

Here we have a typical service. Nothing special, just routing traffic to the
backend application on port `50051`.

### Step 3: the kubernetes `Ingress`

```sh
$ kubectl create -f ingress.yaml
```

A few things to note:

1. We've tagged the ingress with the annotation
   `nginx.ingress.kubernetes.io/backend-protocol: "GRPC"`.  This is the magic
   ingredient that sets up the appropriate nginx configuration to route http/2
   traffic to our service.
1. We're terminating TLS at the ingress and have configured an SSL certificate
   `fortune-teller.stack.build`.  The ingress matches traffic arriving as
   `https://fortune-teller.stack.build:443` and routes unencrypted messages to
   our kubernetes service.

### Step 4: test the connection

Once we've applied our configuration to kubernetes, it's time to test that we
can actually talk to the backend.  To do this, we'll use the
[grpcurl](https://github.com/fullstorydev/grpcurl) utility:

```sh
$ grpcurl fortune-teller.stack.build:443 build.stack.fortune.FortuneTeller/Predict
{
  "message": "Let us endeavor so to live that when we come to die even the undertaker will be sorry.\n\t\t-- Mark Twain, \"Pudd'nhead Wilson's Calendar\""
}
```

### Debugging Hints

1. Obviously, watch the logs on your app.
2. Watch the logs for the nginx-ingress-controller (increasing verbosity as
   needed).
3. Double-check your address and ports.
4. Set the `GODEBUG=http2debug=2` environment variable to get detailed http/2
   logging on the client and/or server.
5. Study RFC 7540 (http/2) <https://tools.ietf.org/html/rfc7540>.

> If you are developing public gRPC endpoints, check out
> https://proto.stack.build, a protocol buffer / gRPC build service that can use
> to help make it easier for your users to consume your API.

> See also the specific GRPC settings of NGINX: https://nginx.org/en/docs/http/ngx_http_grpc_module.html

### Notes on using response/request streams

1. If your server does only response streaming and you expect a stream to be open longer than 60 seconds, you will have to change the `grpc_read_timeout` to accommodate for this.
2. If your service does only request streaming and you expect a stream to be open longer than 60 seconds, you have to change the
`grpc_send_timeout` and the `client_body_timeout`.
3. If you do both response and request streaming with an open stream longer than 60 seconds, you have to change all three timeouts: `grpc_read_timeout`, `grpc_send_timeout` and `client_body_timeout`.

Values for the timeouts must be specified as e.g. `"1200s"`.

> On the most recent versions of nginx-ingress, changing these timeouts requires using the `nginx.ingress.kubernetes.io/server-snippet` annotation. There are plans for future releases to allow using the Kubernetes annotations to define each timeout separately.
