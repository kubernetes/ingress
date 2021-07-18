# Ingress NGINX - Code Overview

This document provides an overview of Ingress NGINX code.


## Core Golang code

This part of the code is responsible for the main logic of Ingress NGINX. It contains all the logics that parses [Ingress Objects](https://kubernetes.io/docs/concepts/services-networking/ingress/), 
[annotations](https://kubernetes.io/docs/reference/glossary/?fundamental=true#term-annotation), watches Endpoints and turn them into usable nginx.conf configuration.


### Core Sync Logics:

Ingress-nginx has an internal model of the ingresses, secrets and endpoints in a given cluster. It maintains two copy of that (1) currently running configuration model and (2) the one generated in response to some changes in the cluster. 

The sync logic diffs the two models and if there's a change it tries to converge the running configuration to the new one. 

There are static and dynamic configuration changes. 

All endpoints and certificate changes are handled dynamically by posting the payload to an internal NGINX endpoint that is handled by Lua.

---

The following parts of the code can be found:

### Entrypoint

Is the `main` package, responsible for starting ingress-nginx program.

It can be found in [cmd/nginx](https://github.com/kubernetes/ingress-nginx/tree/main/cmd/nginx) directory.

### Version

Is the package of the code responsible for adding `version` subcommand, and can be found in [version](https://github.com/kubernetes/ingress-nginx/tree/main/version) directory.

### Internal code

This part of the code contains the internal logics that compose Ingress NGINX Controller, and it's split in:

#### Admission Controller

Contains the code of [Kubernetes Admission Controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) which validates the syntax of ingress objects before accepting it.

This code can be found in [internal/admission/controller](https://github.com/kubernetes/ingress-nginx/tree/main/internal/admission/controller) directory.


#### File functions

Contains auxiliary codes that deal with files, such as generating the SHA1 checksum of a file, or creating required directories.

This code can be found in [internal/file](https://github.com/kubernetes/ingress-nginx/blob/main/internal/file) directory.

#### Ingress functions

Contains all the logics from NGINX Ingress Controller, with some examples being:

* Expected Golang structures that will be used in templates and other parts of the codes - [internal/ingress/types.go](https://github.com/kubernetes/ingress-nginx/blob/main/internal/ingress/types.go).
* supported annotations and its parsing logics - [internal/ingress/annotations](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/annotations).
* reconciliation loops and logics - [internal/ingress/controller](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/controller)
* Defaults - define the default struct.
* Error interface and types implementation - [internal/ingress/errors](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/errors)
* Metrics collectors for Prometheus exporting - [internal/ingress/metric](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/metric).
* Resolver - Extracts information from a controller.
* Ingress Object status publisher - [internal/ingress/status](https://github.com/kubernetes/ingress-nginx/tree/main/internal/ingress/status).

And other parts of the code that will be written in this document in a future.

#### K8s functions

Contains helper functions for parsing Kubernetes objects.

This part of the code can be found in [internal/k8s](https://github.com/kubernetes/ingress-nginx/tree/main/internal/k8s) directory.

#### Networking functions

Contains helper functions for networking, such as IPv4 and IPv6 parsing, SSL certificate parsing, etc.

This part of the code can be found in [internal/net](https://github.com/kubernetes/ingress-nginx/tree/main/internal/net) directory.

#### NGINX functions

Contains helper function to deal with NGINX, such as verify if it's running and reading it's configuration file parts.

This part of the code can be found in [internal/nginx](https://github.com/kubernetes/ingress-nginx/tree/main/internal/nginx) directory.

#### Tasks / Queue

Contains the functions responsible for the sync queue part of the controller.

This part of the code can be found in [internal/task](https://github.com/kubernetes/ingress-nginx/tree/main/internal/task) directory.

#### Other parts of internal

Other parts of internal code might not be covered here, like runtime and watch but they can be added in a future.

## E2E Test

The e2e tests code is in [test](https://github.com/kubernetes/ingress-nginx/tree/main/test) directory.

## Other programs

Describe here `kubectl plugin`, `dbg`, `waitshutdown` and cover the hack scripts.

## Deploy files

This directory contains the `yaml` deploy files used as examples or references in the docs to deploy Ingress NGINX and other componentes.

Those files are in [deploy](https://github.com/kubernetes/ingress-nginx/tree/main/deploy) directory.

## Helm Chart

Used to generate the Helm chart published.

Code is in [charts/ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx).

## Documentation/Website

The documentation used to generate the website https://kubernetes.github.io/ingress-nginx/

This code is available in [docs](https://github.com/kubernetes/ingress-nginx/tree/main/docs) and it's main "language" is `Markdown`, used by [mkdocs](https://github.com/kubernetes/ingress-nginx/blob/main/mkdocs.yml) file to generate static pages.

## Container Images

Container images used to run ingress-nginx, or to build the final image.

### Base Images

Contains the `Dockerfiles` and scripts used to build base images that are used in other parts of the repo. They are present in [images](https://github.com/kubernetes/ingress-nginx/tree/main/images) repo. Some examples:
* [nginx](https://github.com/kubernetes/ingress-nginx/tree/main/images/nginx) - The base NGINX image ingress-nginx uses is not a vanilla NGINX. It bundles many libraries together and it is a job in itself to maintain that and keep things up-to-date.
* [custom-error-pages](https://github.com/kubernetes/ingress-nginx/tree/main/images/custom-error-pages) - Used on the custom error page examples.

There are other images inside this directory.

### Ingress Controller Image

The image used to build the final ingress controller, used in deploy scripts and Helm charts. 

This is NGINX with some Lua enhancement. We do dynamic certificate, endpoints handling, canary traffic split, custom load balancing etc at this component. One can also add new functionalities using Lua plugin system.

The files are in [rootfs](https://github.com/kubernetes/ingress-nginx/tree/main/rootfs) directory and contains:

* The Dockerfile
* [Auxiliary scripts](https://github.com/kubernetes/ingress-nginx/tree/main/rootfs/ingress-controller)

#### Ingress NGINX Lua Scripts

Ingress NGINX uses Lua Scripts to enable features like hot reloading, rate limiting and monitoring. Some are written using the [OpenResty](https://openresty.org/en/) helper.

The directory containing Lua scripts is [rootfs/etc/nginx/lua](https://github.com/kubernetes/ingress-nginx/tree/main/rootfs/etc/nginx/lua).

#### Nginx Go template file

One of the functions of Ingress NGINX is to turn [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) objects into nginx.conf file. 

To do so, the final step is to apply those configurations in [nginx.tmpl](https://github.com/kubernetes/ingress-nginx/tree/main/rootfs/etc/nginx/template) turning it into a final nginx.conf file.

