# Default backend

The default backend is a service which handles all URL paths and hosts the nginx controller doesn't understand
(i.e., all the requests that are not mapped with an Ingress).

Basically a default backend exposes two URLs:

- `/healthz` that returns 200
- `/` that returns 404

!!! example
    The sub-directory [`/images/custom-error-pages`](https://github.com/kubernetes/ingress-nginx/tree/main/images/custom-error-pages)
    provides an additional service for the purpose of customizing the error pages served via the default backend.
