#  Basic usage - host based routing

ingress-nginx can be used for many use cases, inside various cloud provider and supports a lot of configurations. In this section you can find a common usage scenario where a single load balancer powered by ingress-nginx will route traffic to 2 different HTTP backend services based on the host name.

First of all follow the instructions to install ingress-nginx. Then imagine that you need to expose 2 HTTP services already installed: `myServiceA`, `myServiceB`. Let's say that you want to expose the first at `myServiceA.foo.org` and the second at `myServiceB.foo.org`. One possible solution is to create two **ingress** resources:

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-myservicea
  annotations:
    # use the shared ingress-nginx
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
  - host: myservicea.foo.org
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myservicea
            port:
              number: 80
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-myserviceb
  annotations:
    # use the shared ingress-nginx
    kubernetes.io/ingress.class: "nginx"
spec:
  rules:
  - host: myserviceb.foo.org
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myserviceb
            port:
              number: 80
```

When you apply this yaml, 2 ingress resources will be created managed by the **ingress-nginx** instance. Nginx is configured to automatically discover all ingress with the `kubernetes.io/ingress.class: "nginx"` annotation.
Please note that the ingress resource should be placed inside the same namespace of the backend resource.

If cluster version >= 1.19 the Ingress resource above will not work.
You need to deploy the following so that it match the new api version.
This will require changes to the Ingress objects as shown below, and a new object from type "kind: IngressClass"
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-myservicea
spec:
  rules:
  - host: myservicea.foo.org
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myservicea
            port: 
              number: 80
  ingressClassName: nginx
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-myserviceb
spec:
  rules:
  - host: myserviceb.foo.org
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myserviceb
            port: 
              number: 80
  ingressClassName: nginx
---
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  name: nginx
spec:
  controller: k8s.io/ingress-nginx
```

On many cloud providers ingress-nginx will also create the corresponding Load Balancer resource. All you have to do is get the external IP and add a DNS `A record` inside your DNS provider that point myServiceA.foo.org and myServiceB.foo.org to the nginx external IP. Get the external IP by running:

```
kubectl get services -n ingress-nginx
```
