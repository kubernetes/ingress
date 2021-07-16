# Pod Security Policy (PSP)

In most clusters today, by default, all resources (e.g. Deployments and ReplicatSets)
have permissions to create pods.
Kubernetes however provides a more fine-grained authorization policy called
[Pod Security Policy (PSP)](https://kubernetes.io/docs/concepts/policy/pod-security-policy/).

PSP allows the cluster owner to define the permission of each object, for example creating a pod.
If you have PSP enabled on the cluster, and you deploy ingress-nginx,
you will need to provide the Deployment with the permissions to create pods.

Before applying any objects, first apply the PSP permissions by running:
```console
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/psp/psp.yaml
```

Note: PSP permissions must be granted before to the creation of the Deployment and the ReplicaSet.
