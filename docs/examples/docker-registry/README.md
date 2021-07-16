# Docker registry

This example demonstrates how to deploy a [docker registry](https://github.com/docker/distribution) in the cluster and configure Ingress enable access from Internet

## Deployment

First we deploy the docker registry in the cluster:

```console
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/docker-registry/deployment.yaml
```

!!! Important
    **DO NOT RUN THIS IN PRODUCTION**

    This deployment uses `emptyDir` in the `volumeMount` which means the contents of the registry will be deleted when the pod dies.

The next required step is creation of the ingress rules. To do this we have two options: with and without TLS

### Without TLS

Download and edit the yaml deployment replacing `registry.<your domain>` with a valid DNS name pointing to the ingress controller:

```console
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/docker-registry/ingress-without-tls.yaml
```

!!! Important
    Running a docker registry without TLS requires we configure our local docker daemon with the insecure registry flag.

Please check [deploy a plain http registry](https://docs.docker.com/registry/insecure/#deploy-a-plain-http-registry)

### With TLS

Download and edit the yaml deployment replacing `registry.<your domain>` with a valid DNS name pointing to the ingress controller:

```console
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/docker-registry/ingress-with-tls.yaml
```

Deploy [kube lego](https://github.com/jetstack/kube-lego) use [Let's Encrypt](https://letsencrypt.org/) certificates or edit the ingress rule to use a secret with an existing SSL certificate.

### Testing

To test the registry is working correctly we download a known image from [docker hub](https://hub.docker.com), create a tag pointing to the new registry and upload the image:

```console
docker pull ubuntu:16.04
docker tag ubuntu:16.04 `registry.<your domain>/ubuntu:16.04`
docker push `registry.<your domain>/ubuntu:16.04`
```

Please replace `registry.<your domain>` with your domain.
