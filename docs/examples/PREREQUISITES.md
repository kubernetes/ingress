# Prerequisites

Many of the examples in this directory have common prerequisites.

## TLS certificates

Unless otherwise mentioned, the TLS secret used in examples is a 2048 bit RSA
key/cert pair with an arbitrarily chosen hostname, created as follows

```console
$ openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=nginxsvc/O=nginxsvc"
Generating a 2048 bit RSA private key
................+++
................+++
writing new private key to 'tls.key'
-----

$ kubectl create secret tls tls-secret --key tls.key --cert tls.crt
secret "tls-secret" created
```

Note: If using CA Authentication, described below, you will need to sign the server certificate with the CA.

## Client Certificate Authentication

CA Authentication also known as Mutual Authentication allows both the server and client to verify each others
identity via a common CA.

We have a CA Certificate which we obtain usually from a Certificate Authority and use that to sign
both our server certificate and client certificate. Then every time we want to access our backend, we must
pass the client certificate.

These instructions are based on the following [blog](https://medium.com/@awkwardferny/configuring-certificate-based-mutual-authentication-with-kubernetes-ingress-nginx-20e7e38fdfca)

**Generate the CA Key and Certificate:**

```console
openssl req -x509 -sha256 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 356 -nodes -subj '/CN=My Cert Authority'
```

**Generate the Server Key, and Certificate and Sign with the CA Certificate:**

```console
openssl req -new -newkey rsa:4096 -keyout server.key -out server.csr -nodes -subj '/CN=mydomain.com'
openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
```

**Generate the Client Key, and Certificate and Sign with the CA Certificate:**

```console
openssl req -new -newkey rsa:4096 -keyout client.key -out client.csr -nodes -subj '/CN=My Client'
openssl x509 -req -sha256 -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 02 -out client.crt
```

Once this is complete you can continue to follow the instructions [here](./auth/client-certs/README.md#creating-certificate-secrets)

## Test HTTP Service

All examples that require a test HTTP Service use the standard http-svc pod,
which you can deploy as follows

```console
$ kubectl create -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/http-svc.yaml
service "http-svc" created
replicationcontroller "http-svc" created

$ kubectl get po
NAME             READY     STATUS    RESTARTS   AGE
http-svc-p1t3t   1/1       Running   0          1d

$ kubectl get svc
NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)            AGE
http-svc         10.0.122.116   <pending>     80:30301/TCP       1d
```

You can test that the HTTP Service works by exposing it temporarily

```console
$ kubectl patch svc http-svc -p '{"spec":{"type": "LoadBalancer"}}'
"http-svc" patched

$ kubectl get svc http-svc
NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)            AGE
http-svc         10.0.122.116   <pending>     80:30301/TCP       1d

$ kubectl describe svc http-svc
Name:				    http-svc
Namespace:			    default
Labels:			        app=http-svc
Selector:		        app=http-svc
Type:			        LoadBalancer
IP:			            10.0.122.116
LoadBalancer Ingress:	108.59.87.136
Port:			        http	80/TCP
NodePort:		        http	30301/TCP
Endpoints:		        10.180.1.6:8080
Session Affinity:	    None
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason			Message
  ---------	--------	-----	----			-------------	--------	------			-------
  1m		1m		1	{service-controller }			Normal		Type			ClusterIP -> LoadBalancer
  1m		1m		1	{service-controller }			Normal		CreatingLoadBalancer	Creating load balancer
  16s		16s		1	{service-controller }			Normal		CreatedLoadBalancer	Created load balancer

$ curl 108.59.87.136
CLIENT VALUES:
client_address=10.240.0.3
command=GET
real path=/
query=nil
request_version=1.1
request_uri=http://108.59.87.136:8080/

SERVER VALUES:
server_version=nginx: 1.9.11 - lua: 10001

HEADERS RECEIVED:
accept=*/*
host=108.59.87.136
user-agent=curl/7.46.0
BODY:
-no body in request-

$ kubectl patch svc http-svc -p '{"spec":{"type": "NodePort"}}'
"http-svc" patched
```
