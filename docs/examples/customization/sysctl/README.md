# Sysctl tuning

This example aims to demonstrate the use of an Init Container to adjust sysctl default values using `kubectl patch`

```console
kubectl patch deployment -n ingress-nginx nginx-ingress-controller \
    --patch="$(curl https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/docs/examples/customization/sysctl/patch.json)"
```

**Changes:**

- Backlog Queue setting `net.core.somaxconn` from `128` to `32768`
- Ephemeral Ports setting `net.ipv4.ip_local_port_range` from `32768 60999` to `1024 65000`

In a [post from the NGINX blog](https://www.nginx.com/blog/tuning-nginx/), it is possible to see an explanation for the changes.
