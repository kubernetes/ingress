# gRPC

This example demonstrates how to route traffic to a gRPC service through the nginx controller.

## Prerequisites

1. You have a kubernetes cluster running.
2. You have a domain name such as `example.com` that is configured to route traffic to the ingress controller.
3. You have the nginx-ingress controller installed as per docs, with gRPC support.
4. You have a backend application running a gRPC server and listening for TCP traffic.  If you want, you can use <https://github.com/grpc/grpc-go/blob/91e0aeb192456225adf27966d04ada4cf8599915/examples/features/reflection/server/main.go> as an example.
5. You're also responsible for provisioning an SSL certificate for the ingress. So you need to have a valid SSL certificate, deployed as a Kubernetes secret of type tls, in the same namespace as the gRPC application.

### Step 1: Create a Kubernetes `Deployment` for gRPC app

- Make sure your gRPC application pod is running and listening for connections. For example you can try a kubectl command like this below:
  ```
  $ kubectl get po -A -o wide | grep go-grpc-greeter-server
  ```
- If you have a gRPC app deployed in your cluster, then skip further notes in this Step 1, and continue from Step 2 below. 

- As an example gRPC application, we can use this app <https://github.com/grpc/grpc-go/blob/91e0aeb192456225adf27966d04ada4cf8599915/examples/features/reflection/server/main.go> .

- To create a container image for this app, you can use [this Dockerfile](../../../images/go-grpc-greeter-server/rootfs/Dockerfile). 

- If you use the Dockerfile mentioned above, to create a image, then given below is an example of a Kubernetes manifest, to create a deployment resource, that uses that image. If needed, then edit this manifest to suit your needs. Assuming the name of this yaml file is  `deployment.go-grpc-greeter-server.yaml` ;

  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    labels:
      app: go-grpc-greeter-server
    name: go-grpc-greeter-server
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: go-grpc-greeter-server
    template:
      metadata:
        labels:
          app: go-grpc-greeter-server
      spec:
        containers:
        - image: <reponame>/go-grpc-greeter-server   # Edit this for your reponame
          resources:
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 50m
              memory: 50Mi
          name: go-grpc-greeter-server
          ports:
          - containerPort: 50051
  EOF
  ```

### Step 2: Create the Kubernetes `Service` for the gRPC app

- You can use the following example manifest to create a service of type ClusterIP. Edit the name/namespace/label/port to match your deployment/pod ;
  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: v1
  kind: Service
  metadata:
    labels:
      app: go-grpc-greeter-server
    name: go-grpc-greeter-server
  spec:
    ports:
    - port: 80
      protocol: TCP
      targetPort: 50051
    selector:
      app: go-grpc-greeter-server
    type: ClusterIP
  EOF
  ```
- You can save the above example manifest to a file with name `service.go-grpc-greeter-server.yaml` and edit it to match your deployment/pod, if required. You can create the service resource with a kubectl command like this ;

  ```
  $ kubectl create -f service.go-grpc-greeter-server.yaml
  ```

### Step 3: Create the Kubernetes `Ingress` resource for the gRPC app

- Use the following example manifest of a ingress resource to create a ingress for your grpc app. If required, edit it to match your app's details like name, namespace, service, secret etc. Make sure you have the required SSL-Certificate, existing in your Kubernetes cluster, in the same namespace where the gRPC app is. The certificate must be available as a kubernetes secret resource, of type "kubernete.io/tls" https://kubernetes.io/docs/concepts/configuration/secret/#tls-secrets. This is because we are terminating TLS on the ingress;

  ```
  cat <<EOF | kubectl apply -f -
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    annotations:
      kubernetes.io/ingress.class: "nginx"
      nginx.ingress.kubernetes.io/ssl-redirect: "true"
      nginx.ingress.kubernetes.io/backend-protocol: "GRPC"
    name: fortune-ingress
    namespace: default
  spec:
    rules:
    - host: grpctest.dev.mydomain.com
      http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: go-grpc-greeter-server
              port:
                number: 80
    tls:
    # This secret must exist beforehand
    # The cert must also contain the subj-name grpctest.dev.mydomain.com
    # https://github.com/kubernetes/ingress-nginx/blob/master/docs/examples/PREREQUISITES.md#tls-certificates
    - secretName: wildcard.dev.mydomain.com
      hosts:
        - grpctest.dev.mydomain.com
  EOF
  ```

- If you save the above example manifest as a file named `ingress.go-grpc-greeter-server.yaml` and edit it to match your deployment and service, you can create the ingress like this ;

  ```
  $ kubectl create -f ingress.go-grpc-greeter-server.yaml
  ```

- The takeaway is that we are not doing any TLS configuration on the server (as we are terminating TLS at the ingress level, gRPC traffic will travel unencrypted inside the cluster and arrive "insecure").

- For your own application you may or may not want to do this.  If you prefer to forward encrypted traffic to your POD and terminate TLS at the gRPC server itself, add the ingress annotation `nginx.ingress.kubernetes.io/backend-protocol: "GRPCS"`.

- A few more things to note:

  - We've tagged the ingress with the annotation `nginx.ingress.kubernetes.io/backend-protocol: "GRPC"`.  This is the magic ingredient that sets up the appropriate nginx configuration to route http/2 traffic to our service.

  - We're terminating TLS at the ingress and have configured an SSL certificate `wildcard.dev.mydomain.com`.  The ingress matches traffic arriving as `https://grpctest.dev.mydomain.com:443` and routes unencrypted messages to the backend Kubernetes service.

### Step 4: test the connection

- Once we've applied our configuration to Kubernetes, it's time to test that we can actually talk to the backend.  To do this, we'll use the [grpcurl](https://github.com/fullstorydev/grpcurl) utility:

  ```
  $ grpcurl grpctest.dev.mydomain.com:443 helloworld.Greeter/SayHello
  {
    "message": "
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
