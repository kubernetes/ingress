# Changelog

### 0.45.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.45.0@sha256:c4390c53f348c3bd4e60a5dd6a11c35799ae78c49388090140b9d72ccede1755`

_New Features:_

- Update the Ingress Controller Image to correct OpenSSL CVEs
- Add support for Jaeger Endpoint [#6884](https://github.com/kubernetes/ingress-nginx/pull/6884)
- Allow Multiple Publish Status Addresses [#6856](https://github.com/kubernetes/ingress-nginx/pull/6856)

_Changes:_

- [X] [#6995](https://github.com/kubernetes/ingress-nginx/pull/6995) updating nginx base image across repo
- [X] [#6983](https://github.com/kubernetes/ingress-nginx/pull/6983) Expose Geo IP subdivision 1 as variables
- [X] [#6979](https://github.com/kubernetes/ingress-nginx/pull/6979) Changed servicePort value for metrics
- [X] [#6971](https://github.com/kubernetes/ingress-nginx/pull/6971) Fix crl not reload when crl got updated in the ca secret
- [X] [#6957](https://github.com/kubernetes/ingress-nginx/pull/6957) Add ability to specify automountServiceAccountToken
- [X] [#6956](https://github.com/kubernetes/ingress-nginx/pull/6956) update nginx base image, handle jaeger propagation format 
- [X] [#6936](https://github.com/kubernetes/ingress-nginx/pull/6936) update tracing libraries for opentracing 1.6.0 
- [X] [#6908](https://github.com/kubernetes/ingress-nginx/pull/6908) feat(chart) Add volumes to default-backend deployment #6908 
- [X] [#6884](https://github.com/kubernetes/ingress-nginx/pull/6884) jaeger-endpoint feature for non-agent trace collectors
- [X] [#6856](https://github.com/kubernetes/ingress-nginx/pull/6856) Allow multiple publish status addresses
- [X] [#6971](https://github.com/kubernetes/ingress-nginx/pull/6971) Fix bug related to CRL update

### 0.44.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a`

_New Features:_

- Update alpine to v3.13
- client-go v0.20.2
- New option to set a shutdown grace period [#5855](https://github.com/kubernetes/ingress-nginx/pull/5855)
- Support for Global/Distributed Rate Limiting [#6673](https://github.com/kubernetes/ingress-nginx/pull/6673
- Update kube-webhook-certgen image to v1.5.1

_Changes:_

- [X] [#5855](https://github.com/kubernetes/ingress-nginx/pull/5855) New option to set a shutdown grace period
- [X] [#6530](https://github.com/kubernetes/ingress-nginx/pull/6530) #6477 Allow specifying the loadBalancerIP for the internal load balancer
- [X] [#6667](https://github.com/kubernetes/ingress-nginx/pull/6667) Correct the value for setting a static IP
- [X] [#6669](https://github.com/kubernetes/ingress-nginx/pull/6669) chore: Add test to internal ingress resolver pkg
- [X] [#6673](https://github.com/kubernetes/ingress-nginx/pull/6673) Add Global/Distributed Rate Limiting support
- [X] [#6679](https://github.com/kubernetes/ingress-nginx/pull/6679) Bugfix: fix incomplete log
- [X] [#6695](https://github.com/kubernetes/ingress-nginx/pull/6695) include lua-resty-ipmatcher and lua-resty-global-throttle inn the base image
- [X] [#6705](https://github.com/kubernetes/ingress-nginx/pull/6705) Update nginx alpine image to 3.12
- [X] [#6708](https://github.com/kubernetes/ingress-nginx/pull/6708) fix generated code for the new year
- [X] [#6710](https://github.com/kubernetes/ingress-nginx/pull/6710) Update nginx base image
- [X] [#6711](https://github.com/kubernetes/ingress-nginx/pull/6711) Update test container images
- [X] [#6712](https://github.com/kubernetes/ingress-nginx/pull/6712) Update tag version
- [X] [#6717](https://github.com/kubernetes/ingress-nginx/pull/6717) fix ipmatcher installation
- [X] [#6718](https://github.com/kubernetes/ingress-nginx/pull/6718) generalize cidr parsing and improve lua tests
- [X] [#6719](https://github.com/kubernetes/ingress-nginx/pull/6719) Update nginx image
- [X] [#6720](https://github.com/kubernetes/ingress-nginx/pull/6720) Update test runner image
- [X] [#6724](https://github.com/kubernetes/ingress-nginx/pull/6724) e2e test for requiring memcached setting to configure global rate limit
- [X] [#6726](https://github.com/kubernetes/ingress-nginx/pull/6726) add body_filter_by_lua_block lua plugin to ingress-nginx
- [X] [#6730](https://github.com/kubernetes/ingress-nginx/pull/6730) Do not create HPA if default backend not enabled
- [X] [#6754](https://github.com/kubernetes/ingress-nginx/pull/6754) Fix KEDA autoscaler resource
- [X] [#6756](https://github.com/kubernetes/ingress-nginx/pull/6756) Dashboard Panel supposed to show average is showing sum.
- [X] [#6760](https://github.com/kubernetes/ingress-nginx/pull/6760) Update alpine to 3.13
- [X] [#6761](https://github.com/kubernetes/ingress-nginx/pull/6761) Adding quotes in the serviceAccount name in Helm values.
- [X] [#6764](https://github.com/kubernetes/ingress-nginx/pull/6764) Update nginx image
- [X] [#6767](https://github.com/kubernetes/ingress-nginx/pull/6767) Remove ClusterRole when scope option is enabled
- [X] [#6783](https://github.com/kubernetes/ingress-nginx/pull/6783) Add custom annotations to ScaledObject #6757
- [X] [#6785](https://github.com/kubernetes/ingress-nginx/pull/6785) Update kube-webhook-certgen image to v1.5.1
- [X] [#6786](https://github.com/kubernetes/ingress-nginx/pull/6786) Release helm chart v3.21.0
- [X] [#6787](https://github.com/kubernetes/ingress-nginx/pull/6787) Update static manifests
- [X] [#6788](https://github.com/kubernetes/ingress-nginx/pull/6788) Update go dependencies
- [X] [#6790](https://github.com/kubernetes/ingress-nginx/pull/6790) :bug: return error if tempconfig missing
- [X] [#6796](https://github.com/kubernetes/ingress-nginx/pull/6796) Updates to the custom default SSL certificate must trigger a reload
- [X] [#6802](https://github.com/kubernetes/ingress-nginx/pull/6802) Add value for configuring a custom Diffie-Hellman parameters file
- [X] [#6811](https://github.com/kubernetes/ingress-nginx/pull/6811) add --status-port flag to dbg
- [X] [#6815](https://github.com/kubernetes/ingress-nginx/pull/6815) Allow use of numeric namespaces in helm chart
- [X] [#6817](https://github.com/kubernetes/ingress-nginx/pull/6817) corrected nginx configuration doc path
- [X] [#6818](https://github.com/kubernetes/ingress-nginx/pull/6818) Release helm chart v3.22.0
- [X] [#6819](https://github.com/kubernetes/ingress-nginx/pull/6819) Update kind and kindest images
- [X] [#6830](https://github.com/kubernetes/ingress-nginx/pull/6830) Remove extra comma for Tracing Jaeger config json
- [X] [#6832](https://github.com/kubernetes/ingress-nginx/pull/6832) Fix e2e build
- [X] [#6842](https://github.com/kubernetes/ingress-nginx/pull/6842) Change chart-testing image

_Documentation:_

- [X] [#6723](https://github.com/kubernetes/ingress-nginx/pull/6723) fix link in annotation docs
- [X] [#6727](https://github.com/kubernetes/ingress-nginx/pull/6727) Fix the documentation for the proxy-ssl-secret and the auth-tls-secret annotations
- [X] [#6738](https://github.com/kubernetes/ingress-nginx/pull/6738) docs/deploy: fix grammar and inconsistencies
- [X] [#6741](https://github.com/kubernetes/ingress-nginx/pull/6741) Update mkdocs, fix nodeport link and add microk8s link
- [X] [#6746](https://github.com/kubernetes/ingress-nginx/pull/6746) Move Azure deploy note to right item on doc page
- [X] [#6749](https://github.com/kubernetes/ingress-nginx/pull/6749) Update README.md
- [X] [#6789](https://github.com/kubernetes/ingress-nginx/pull/6789) Update e2e tests link markdown
- [X] [#6813](https://github.com/kubernetes/ingress-nginx/pull/6813) Added docs to clear up PROXY definition
- [X] [#6823](https://github.com/kubernetes/ingress-nginx/pull/6823) Change break link for documentation

### 0.43.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.43.0@sha256:9bba603b99bf25f6d117cf1235b6598c16033ad027b143c90fa5b3cc583c5713`

_New Features:_

- Support local GeoIP database mirror [#6685](https://github.com/kubernetes/ingress-nginx/pull/6685)

_Changes:_

- [X] [#6676](https://github.com/kubernetes/ingress-nginx/pull/6676) include new resty lua libs in base image
- [X] [#6680](https://github.com/kubernetes/ingress-nginx/pull/6680) Update cloudbuild gcb-docker-gcloud image
- [X] [#6682](https://github.com/kubernetes/ingress-nginx/pull/6682) Update nginx image
- [X] [#6683](https://github.com/kubernetes/ingress-nginx/pull/6683) Remove dead code
- [X] [#6684](https://github.com/kubernetes/ingress-nginx/pull/6684) Update ingress-nginx test image
- [X] [#6685](https://github.com/kubernetes/ingress-nginx/pull/6685) Add GeoIP Local mirror support
- [X] [#6688](https://github.com/kubernetes/ingress-nginx/pull/6688) feat: allow volume-type emptyDir in controller podsecuritypolicy
- [X] [#6691](https://github.com/kubernetes/ingress-nginx/pull/6691) Helm: Ingress config change
- [X] [#6692](https://github.com/kubernetes/ingress-nginx/pull/6692) add string split function to template funcMap
- [X] [#6694](https://github.com/kubernetes/ingress-nginx/pull/6694) Release helm chart

### 0.42.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.42.0@sha256:f7187418c647af4a0039938b0ab36c2322ac3662d16be69f9cc178bfd25f7eee`

_New Features:_

- NGINX 1.19.6
- Go 1.15.6
- client-go v0.20.0
- Support [Keda Autoscaling](https://github.com/kubernetes/ingress-nginx/pull/6493) in helm chart
- Support custom redirect URL parameter in external authentication [#6294](https://github.com/kubernetes/ingress-nginx/pull/6294)

_Changes:_

- [X] [#6294](https://github.com/kubernetes/ingress-nginx/pull/6294) Allow customisation of redirect URL parameter in external auth redirects
- [X] [#6461](https://github.com/kubernetes/ingress-nginx/pull/6461) Update kind github action
- [X] [#6469](https://github.com/kubernetes/ingress-nginx/pull/6469) Allow custom service names for controller and backend (#6457)
- [X] [#6470](https://github.com/kubernetes/ingress-nginx/pull/6470) Added role binding for 'ingress-nginx-admission' to PSP example (#6018)
- [X] [#6473](https://github.com/kubernetes/ingress-nginx/pull/6473) Enable external auth e2e tests
- [X] [#6480](https://github.com/kubernetes/ingress-nginx/pull/6480) Refactor extraction of ingress pod details
- [X] [#6485](https://github.com/kubernetes/ingress-nginx/pull/6485) Fix sum value of nginx process connections
- [X] [#6493](https://github.com/kubernetes/ingress-nginx/pull/6493) Support Keda Autoscaling
- [X] [#6499](https://github.com/kubernetes/ingress-nginx/pull/6499) Fix opentracing propagation on auth-url
- [X] [#6505](https://github.com/kubernetes/ingress-nginx/pull/6505) Reorder HPA resource list to work with GitOps tooling
- [X] [#6514](https://github.com/kubernetes/ingress-nginx/pull/6514) Remove helm2 support and update docs
- [X] [#6526](https://github.com/kubernetes/ingress-nginx/pull/6526) Fix helm repo update command
- [X] [#6529](https://github.com/kubernetes/ingress-nginx/pull/6529) Fix ErrorLogLevel in stream contexts
- [X] [#6533](https://github.com/kubernetes/ingress-nginx/pull/6533) Update nginx to 1.19.5
- [X] [#6544](https://github.com/kubernetes/ingress-nginx/pull/6544) Fix the name of default backend variable
- [X] [#6546](https://github.com/kubernetes/ingress-nginx/pull/6546) bugfix: update trafficShapingPolicy not working in ewma load-balance
- [X] [#6550](https://github.com/kubernetes/ingress-nginx/pull/6550) Ensure any change in the charts directory trigger ci tests
- [X] [#6553](https://github.com/kubernetes/ingress-nginx/pull/6553) fixes: allow user to specify the maxmium number of retries in stream block
- [X] [#6558](https://github.com/kubernetes/ingress-nginx/pull/6558) Update kindest image
- [X] [#6560](https://github.com/kubernetes/ingress-nginx/pull/6560) Fix nginx ingress variables for definitions without hosts
- [X] [#6561](https://github.com/kubernetes/ingress-nginx/pull/6561) Disable HTTP/2 in the webhook server
- [X] [#6571](https://github.com/kubernetes/ingress-nginx/pull/6571) Add gosec action
- [X] [#6574](https://github.com/kubernetes/ingress-nginx/pull/6574) Release helm chart v3.13.0
- [X] [#6576](https://github.com/kubernetes/ingress-nginx/pull/6576) Fix nginx ingress variables for definitions with Backend
- [X] [#6577](https://github.com/kubernetes/ingress-nginx/pull/6577) Update build TAG
- [X] [#6578](https://github.com/kubernetes/ingress-nginx/pull/6578) Update helm chart-testing image
- [X] [#6580](https://github.com/kubernetes/ingress-nginx/pull/6580) fix for 6564
- [X] [#6586](https://github.com/kubernetes/ingress-nginx/pull/6586) fix(chart): Move 'maxmindLicenseKey' to `controller.maxmindLicenseKey`
- [X] [#6587](https://github.com/kubernetes/ingress-nginx/pull/6587) feat(nginx.tmpl): Add support for GeoLite2-Country and GeoIP2-Country databases
- [X] [#6592](https://github.com/kubernetes/ingress-nginx/pull/6592) Update github actions
- [X] [#6593](https://github.com/kubernetes/ingress-nginx/pull/6593) Rollback chart-releaser action
- [X] [#6594](https://github.com/kubernetes/ingress-nginx/pull/6594) Fix chart-releaser action
- [X] [#6595](https://github.com/kubernetes/ingress-nginx/pull/6595) Fix helm chart releaser action
- [X] [#6596](https://github.com/kubernetes/ingress-nginx/pull/6596) Fix helm chart releaser action
- [X] [#6600](https://github.com/kubernetes/ingress-nginx/pull/6600) Bugfix: some requests fail with 503 when nginx reload
- [X] [#6607](https://github.com/kubernetes/ingress-nginx/pull/6607) fix flaky lua tests
- [X] [#6608](https://github.com/kubernetes/ingress-nginx/pull/6608) make sure canary attributes are reset on ewma backend sync
- [X] [#6614](https://github.com/kubernetes/ingress-nginx/pull/6614) Refactor ingress nginx variables
- [X] [#6617](https://github.com/kubernetes/ingress-nginx/pull/6617) Allow FQDN for ExternalName Service
- [X] [#6620](https://github.com/kubernetes/ingress-nginx/pull/6620) Fix sticky session not set for host in server-alias annotation (#6448)
- [X] [#6621](https://github.com/kubernetes/ingress-nginx/pull/6621) Update go dependencies
- [X] [#6627](https://github.com/kubernetes/ingress-nginx/pull/6627) Update nginx to 1.19.6
- [X] [#6628](https://github.com/kubernetes/ingress-nginx/pull/6628) Update nginx image
- [X] [#6629](https://github.com/kubernetes/ingress-nginx/pull/6629) Add kind e2e tests support for k8s v1.20.0
- [X] [#6635](https://github.com/kubernetes/ingress-nginx/pull/6635) Update test images and go to 1.15.6
- [X] [#6638](https://github.com/kubernetes/ingress-nginx/pull/6638) Update test images
- [X] [#6639](https://github.com/kubernetes/ingress-nginx/pull/6639) Don't pick tried endpoint & count the latest in ewma balancer
- [X] [#6646](https://github.com/kubernetes/ingress-nginx/pull/6646) Adding LoadBalancerIP value for internal service to Helm chart
- [X] [#6652](https://github.com/kubernetes/ingress-nginx/pull/6652) Change helm chart tag name to disambiguate purpose
- [X] [#6660](https://github.com/kubernetes/ingress-nginx/pull/6660) Fix chart-releaser action

_Documentation:_

- [X] [#6481](https://github.com/kubernetes/ingress-nginx/pull/6481) docs(annotations): explicit redirect status code
- [X] [#6486](https://github.com/kubernetes/ingress-nginx/pull/6486) Fix typo
- [X] [#6528](https://github.com/kubernetes/ingress-nginx/pull/6528) Spelling
- [X] [#6494](https://github.com/kubernetes/ingress-nginx/pull/6494) Update development.md with docker version and experimental feature requirment
- [X] [#6501](https://github.com/kubernetes/ingress-nginx/pull/6501) Add Chart changelog instructions
- [X] [#6541](https://github.com/kubernetes/ingress-nginx/pull/6541) fixed misspell
- [X] [#6551](https://github.com/kubernetes/ingress-nginx/pull/6551) Add documentation to activate DHE based ciphers
- [X] [#6566](https://github.com/kubernetes/ingress-nginx/pull/6566) fix docs  log-format-upstream sample
- [X] [#6579](https://github.com/kubernetes/ingress-nginx/pull/6579) Update README.md
- [X] [#6598](https://github.com/kubernetes/ingress-nginx/pull/6598) Add a link to the helm ingress-nginx `CHANGELOG.md` file to the `README.md` file
- [X] [#6623](https://github.com/kubernetes/ingress-nginx/pull/6623) fix typo
- [X] [#6636](https://github.com/kubernetes/ingress-nginx/pull/6636) Fix link to kustomize docs
- [X] [#6642](https://github.com/kubernetes/ingress-nginx/pull/6642) Update README.md
- [X] [#6665](https://github.com/kubernetes/ingress-nginx/pull/6665) added AKS specific documentation

### 0.41.2

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.41.2@sha256:1f4f402b9c14f3ae92b11ada1dfe9893a88f0faeb0b2f4b903e2c67a0c3bf0de`

Fix regression introduced in 0.41.0 with external authentication

_Changes:_

- [X] [#6467](https://github.com/kubernetes/ingress-nginx/pull/6467) Add PathType details in external auth location

### 0.41.1

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.41.1@sha256:595f5c08aaa2bdfd1afdfb2e0f1a668bc85d96f80c097ddb3d4241f0c9122549`

Fix routing regression introduced in 0.41.0 with PathType Exact

_Changes:_

- [X] [#6417](https://github.com/kubernetes/ingress-nginx/pull/6417) Reload nginx when L4 proxy protocol change
- [X] [#6422](https://github.com/kubernetes/ingress-nginx/pull/6422) Add comment indicating server-snippet section
- [X] [#6423](https://github.com/kubernetes/ingress-nginx/pull/6423) Add Default backend HPA autoscaling.
- [X] [#6426](https://github.com/kubernetes/ingress-nginx/pull/6426) Alternate to respecting setting admissionWebhooks.failurePolicy in values.yaml
- [X] [#6443](https://github.com/kubernetes/ingress-nginx/pull/6443) Refactor handling of path Prefix and Exact
- [X] [#6445](https://github.com/kubernetes/ingress-nginx/pull/6445) fix: empty IngressClassName, Error handling
- [X] [#6447](https://github.com/kubernetes/ingress-nginx/pull/6447) Improve class.IsValid logs

### 0.41.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.41.0@sha256:e6019e536cfb921afb99408d5292fa88b017c49dd29d05fc8dbc456aa770d590`

_New Features:_

- NGINX 1.19.4
- Go 1.15.3
- client-go v0.19.3
- alpine 3.12
- jettech/kube-webhook-certgen v1.5.0

_Changes:_

- [X] [#6037](https://github.com/kubernetes/ingress-nginx/pull/6037) Do not append a trailing slash on redirects
- [X] [#6255](https://github.com/kubernetes/ingress-nginx/pull/6255) Update datadog opentracing plugin to v1.2.0
- [X] [#6260](https://github.com/kubernetes/ingress-nginx/pull/6260) Allow Helm Chart to customize admission webhook's annotations, timeoutSeconds, namespaceSelector, objectSelector and cert files locations
- [X] [#6261](https://github.com/kubernetes/ingress-nginx/pull/6261) Add datadog environment as a configuration option
- [X] [#6268](https://github.com/kubernetes/ingress-nginx/pull/6268) Update helm chart
- [X] [#6282](https://github.com/kubernetes/ingress-nginx/pull/6282) Enable e2e tests for k8s v1.16
- [X] [#6283](https://github.com/kubernetes/ingress-nginx/pull/6283) Start v0.41.0 cycle
- [X] [#6299](https://github.com/kubernetes/ingress-nginx/pull/6299) Fix helm chart release
- [X] [#6304](https://github.com/kubernetes/ingress-nginx/pull/6304) Update nginx image
- [X] [#6305](https://github.com/kubernetes/ingress-nginx/pull/6305) Add default linux nodeSelector
- [X] [#6316](https://github.com/kubernetes/ingress-nginx/pull/6316) Numerals in podAnnotations in quotes #6315
- [X] [#6319](https://github.com/kubernetes/ingress-nginx/pull/6319) Update static manifests
- [X] [#6325](https://github.com/kubernetes/ingress-nginx/pull/6325) Filter out secrets that belong to Helm v3
- [X] [#6326](https://github.com/kubernetes/ingress-nginx/pull/6326) Fix liveness and readiness probe path in daemonset chart
- [X] [#6331](https://github.com/kubernetes/ingress-nginx/pull/6331) fix for 6219
- [X] [#6348](https://github.com/kubernetes/ingress-nginx/pull/6348) Add validation for wildcard server names
- [X] [#6349](https://github.com/kubernetes/ingress-nginx/pull/6349) Refactor Exact path matching
- [X] [#6356](https://github.com/kubernetes/ingress-nginx/pull/6356) Add securitycontext settings on defaultbackend
- [X] [#6366](https://github.com/kubernetes/ingress-nginx/pull/6366) Enable validation of ingress definitions from extensions package
- [X] [#6372](https://github.com/kubernetes/ingress-nginx/pull/6372) Check pod is ready
- [X] [#6373](https://github.com/kubernetes/ingress-nginx/pull/6373) Remove k8s.io/kubernetes dependency
- [X] [#6375](https://github.com/kubernetes/ingress-nginx/pull/6375) Improve log messages
- [X] [#6377](https://github.com/kubernetes/ingress-nginx/pull/6377) Added loadBalancerSourceRanges for internal lbs
- [X] [#6382](https://github.com/kubernetes/ingress-nginx/pull/6382) fix: OWASP CoreRuleSet rules for NodeJS and Java
- [X] [#6383](https://github.com/kubernetes/ingress-nginx/pull/6383) Update nginx to 1.19.4
- [X] [#6384](https://github.com/kubernetes/ingress-nginx/pull/6384) Update nginx image to 1.19.4
- [X] [#6385](https://github.com/kubernetes/ingress-nginx/pull/6385) Update helm stable repo
- [X] [#6388](https://github.com/kubernetes/ingress-nginx/pull/6388) Update nginx image in project images
- [X] [#6389](https://github.com/kubernetes/ingress-nginx/pull/6389) Support prefix pathtype
- [X] [#6390](https://github.com/kubernetes/ingress-nginx/pull/6390) Update go to 1.15.3
- [X] [#6391](https://github.com/kubernetes/ingress-nginx/pull/6391) Update alpine packages
- [X] [#6392](https://github.com/kubernetes/ingress-nginx/pull/6392) Update test images
- [X] [#6395](https://github.com/kubernetes/ingress-nginx/pull/6395) Update jettech/kube-webhook-certgen image
- [X] [#6401](https://github.com/kubernetes/ingress-nginx/pull/6401) Fix controller service annotations
- [X] [#6412](https://github.com/kubernetes/ingress-nginx/pull/6412) Improve ingress class error message

_Documentation:_

- [X] [#6272](https://github.com/kubernetes/ingress-nginx/pull/6272) Update hardening guide doc
- [X] [#6278](https://github.com/kubernetes/ingress-nginx/pull/6278) Sync user guide with config defaults changes
- [X] [#6321](https://github.com/kubernetes/ingress-nginx/pull/6321) Fix typo
- [X] [#6403](https://github.com/kubernetes/ingress-nginx/pull/6403) Initial helm chart changelog

### 0.40.2

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.40.2@sha256:46ba23c3fbaafd9e5bd01ea85b2f921d9f2217be082580edc22e6c704a83f02f`

Improve HandleAdmission resiliency

_Changes:_

- [X] [#6284](https://github.com/kubernetes/ingress-nginx/pull/6284) Improve HandleAdmission resiliency

### 0.40.1

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.40.1@sha256:abffcf2d25e3e7c7b67a315a7c664ec79a1588c9c945d3c7a75637c2f55caec6`

Fix regression with clusters running v1.16 where AdmissionReview V1 is available but not enabled.

_Changes:_

- [X] [#6265](https://github.com/kubernetes/ingress-nginx/pull/6265) Add support for admission review v1beta1


### 0.40.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.40.0@sha256:b954d8ff1466eb236162c644bd64e9027a212c82b484cbe47cc21da45fe8bc59`

_Breaking Changes:_

Kubernetes v1.16 or higher is required.
Only ValidatingWebhookConfiguration AdmissionReviewVersions v1 is supported.

Following the [Ingress extensions/v1beta1](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16) deprecation, please use `networking.k8s.io/v1beta1` or `networking.k8s.io/v1` (Kubernetes v1.19 or higher) for new Ingress definitions


_New Features:_

- NGINX 1.19.3
- Go 1.15.2
- client-go v0.19.2
- New annotation `nginx.ingress.kubernetes.io/limit-burst-multiplier` to set value for burst multiplier on rate limit
- Switch to [structured logging](https://github.com/kubernetes/enhancements/tree/master/keps/sig-instrumentation/1602-structured-logging)
- NGINX reloads create Kubernetes events


_New defaults:_

- [server-tokens](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#server-tokens) is disabled
- [ssl-session-tickets](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-session-tickets) is disabled
- [use-gzip](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#use-gzip) is disabled
- [upstream-keepalive-requests](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#upstream-keepalive-requests) is now 10000
- [upstream-keepalive-connections](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#upstream-keepalive-connections) is now 320


_Changes:_

- [X] [#5348](https://github.com/kubernetes/ingress-nginx/pull/5348) Ability to separately disable access log in http and stream contexts
- [X] [#6087](https://github.com/kubernetes/ingress-nginx/pull/6087) Adding parameter for externalTrafficPolicy in internal controller service spec.
- [X] [#6097](https://github.com/kubernetes/ingress-nginx/pull/6097) Return unique addresses from service
- [X] [#6098](https://github.com/kubernetes/ingress-nginx/pull/6098) Update kubernetes kind e2e versions
- [X] [#6099](https://github.com/kubernetes/ingress-nginx/pull/6099) Update go dependencies
- [X] [#6101](https://github.com/kubernetes/ingress-nginx/pull/6101) Add annotation to set value for burst multiplier on rate limit
- [X] [#6104](https://github.com/kubernetes/ingress-nginx/pull/6104) Misc fixes for nginx-ingress chart for better keel and prom-oper
- [X] [#6111](https://github.com/kubernetes/ingress-nginx/pull/6111) Update kind node version to v1.19.0
- [X] [#6112](https://github.com/kubernetes/ingress-nginx/pull/6112) Update go dependencies to fix build logr errors
- [X] [#6113](https://github.com/kubernetes/ingress-nginx/pull/6113) Require Kubernetes v1.14 or higher and deprecate extensions
- [X] [#6114](https://github.com/kubernetes/ingress-nginx/pull/6114) Update go to 1.5.1 and add arm64
- [X] [#6115](https://github.com/kubernetes/ingress-nginx/pull/6115) Increase cloudbuild timeout value
- [X] [#6116](https://github.com/kubernetes/ingress-nginx/pull/6116) Update e2e image
- [X] [#6118](https://github.com/kubernetes/ingress-nginx/pull/6118) Cleanup github actions
- [X] [#6120](https://github.com/kubernetes/ingress-nginx/pull/6120) Use net.JoinHostPort to avoid IPV6 issues
- [X] [#6122](https://github.com/kubernetes/ingress-nginx/pull/6122) Update trace modules
- [X] [#6123](https://github.com/kubernetes/ingress-nginx/pull/6123) Add e2e tests to verify opentracing libraries
- [X] [#6125](https://github.com/kubernetes/ingress-nginx/pull/6125) Library dd-opentracing cannot be static
- [X] [#6129](https://github.com/kubernetes/ingress-nginx/pull/6129) Update test runner image
- [X] [#6143](https://github.com/kubernetes/ingress-nginx/pull/6143) Update default gzip level
- [X] [#6145](https://github.com/kubernetes/ingress-nginx/pull/6145) Switch modules to dynamic and remove http_dav_module
- [X] [#6146](https://github.com/kubernetes/ingress-nginx/pull/6146) Use dynamic load of modules
- [X] [#6148](https://github.com/kubernetes/ingress-nginx/pull/6148) Update cloudbuild jobs
- [X] [#6149](https://github.com/kubernetes/ingress-nginx/pull/6149) Update qemu-user-static image
- [X] [#6150](https://github.com/kubernetes/ingress-nginx/pull/6150) Reject ingresses that use the default annotation if a custom one was provided
- [X] [#6166](https://github.com/kubernetes/ingress-nginx/pull/6166) Update kind and kindest/node images
- [X] [#6167](https://github.com/kubernetes/ingress-nginx/pull/6167) Update chart requirements
- [X] [#6170](https://github.com/kubernetes/ingress-nginx/pull/6170) delete redundant NGINX config about X-Forwarded-Proto
- [X] [#6172](https://github.com/kubernetes/ingress-nginx/pull/6172) Update OSFileWatcher to support symlinks
- [X] [#6176](https://github.com/kubernetes/ingress-nginx/pull/6176) Update go modules
- [X] [#6184](https://github.com/kubernetes/ingress-nginx/pull/6184) Update static manifest yaml files
- [X] [#6186](https://github.com/kubernetes/ingress-nginx/pull/6186) Fix logr dependency issue
- [X] [#6187](https://github.com/kubernetes/ingress-nginx/pull/6187) Add validation support for networking.k8s.io/v1
- [X] [#6190](https://github.com/kubernetes/ingress-nginx/pull/6190) Change server-tokens default value to false
- [X] [#6196](https://github.com/kubernetes/ingress-nginx/pull/6196) disable session tickets by default
- [X] [#6198](https://github.com/kubernetes/ingress-nginx/pull/6198) Delete OCSP Response cache when certificate renewed
- [X] [#6200](https://github.com/kubernetes/ingress-nginx/pull/6200) Move ocsp_response_cache:delete after certificate_data:set
- [X] [#6203](https://github.com/kubernetes/ingress-nginx/pull/6203) Refactor key/value parsing
- [X] [#6211](https://github.com/kubernetes/ingress-nginx/pull/6211) Migrate to structured logging (klog)
- [X] [#6217](https://github.com/kubernetes/ingress-nginx/pull/6217) Add annotation to configure CORS Access-Control-Expose-Headers
- [X] [#6226](https://github.com/kubernetes/ingress-nginx/pull/6226) Change defaults
- [X] [#6231](https://github.com/kubernetes/ingress-nginx/pull/6231) Clear redundant Cross-Origin-Allow- headers from response
- [X] [#6233](https://github.com/kubernetes/ingress-nginx/pull/6233) Add admission controller e2e test
- [X] [#6234](https://github.com/kubernetes/ingress-nginx/pull/6234) Add events for NGINX reloads
- [X] [#6235](https://github.com/kubernetes/ingress-nginx/pull/6235) Update go dependencies

_Documentation:_

- [X] [#5757](https://github.com/kubernetes/ingress-nginx/pull/5757) feat: support to define trusted addresses for proxy protocol in stream block
- [X] [#5881](https://github.com/kubernetes/ingress-nginx/pull/5881) Doc: Adding initial hardening guide
- [X] [#6109](https://github.com/kubernetes/ingress-nginx/pull/6109) Fix documentation table layout
- [X] [#6127](https://github.com/kubernetes/ingress-nginx/pull/6127) AKS example of adding an internal loadbalancer
- [X] [#6128](https://github.com/kubernetes/ingress-nginx/pull/6128) Fixed proxy protocol link
- [X] [#6130](https://github.com/kubernetes/ingress-nginx/pull/6130) Update deploy instructions with corrections
- [X] [#6153](https://github.com/kubernetes/ingress-nginx/pull/6153) Add install command for Scaleway
- [X] [#6154](https://github.com/kubernetes/ingress-nginx/pull/6154) chart: add `topologySpreadConstraint` to controller
- [X] [#6162](https://github.com/kubernetes/ingress-nginx/pull/6162) Add helm chart options to expose metrics service as NodePort
- [X] [#6169](https://github.com/kubernetes/ingress-nginx/pull/6169) Fix Typo in example prometheus rules
- [X] [#6171](https://github.com/kubernetes/ingress-nginx/pull/6171) Docs: remove redundant --election-id arg from Multiple Ingresses
- [X] [#6177](https://github.com/kubernetes/ingress-nginx/pull/6177) Fix make help task to display options
- [X] [#6180](https://github.com/kubernetes/ingress-nginx/pull/6180) Fix helm chart admissionReviewVersions regression
- [X] [#6181](https://github.com/kubernetes/ingress-nginx/pull/6181) Improve prerequisite instructions
- [X] [#6214](https://github.com/kubernetes/ingress-nginx/pull/6214) Update annotations.md - improvements to the documentation of Client Certificate Authentication
- [X] [#6215](https://github.com/kubernetes/ingress-nginx/pull/6215) Update the comment for Makefile taks live-docs
- [X] [#6221](https://github.com/kubernetes/ingress-nginx/pull/6221) corrected reference for release


### 0.35.0

**Image:**

- `k8s.gcr.io/ingress-nginx/controller:v0.35.0@sha256:fc4979d8b8443a831c9789b5155cded454cb7de737a8b727bc2ba0106d2eae8b`

_New Features:_

- NGINX 1.19.2
- New configmap option `enable-real-ip` to enable [realip_module](http://nginx.org/en/docs/http/ngx_http_realip_module.html)
- Use k8s.gcr.io vanity domain
- Go 1.15
- client-go v0.18.6
- Migrate to klog v2

_Changes:_

- [X] [#5887](https://github.com/kubernetes/ingress-nginx/pull/5887) Add force-enable-realip-module
- [X] [#5888](https://github.com/kubernetes/ingress-nginx/pull/5888) Update dev-env.sh script
- [X] [#5923](https://github.com/kubernetes/ingress-nginx/pull/5923) Fix error in grpcbin deployment and enable e2e test
- [X] [#5924](https://github.com/kubernetes/ingress-nginx/pull/5924) Validate endpoints are ready in e2e tests
- [X] [#5931](https://github.com/kubernetes/ingress-nginx/pull/5931) Add opentracing operation name settings
- [X] [#5933](https://github.com/kubernetes/ingress-nginx/pull/5933) Update opentracing nginx module
- [X] [#5946](https://github.com/kubernetes/ingress-nginx/pull/5946) Do not add namespace to cluster-scoped resources
- [X] [#5951](https://github.com/kubernetes/ingress-nginx/pull/5951) Use env expansion to provide namespace in container args
- [X] [#5952](https://github.com/kubernetes/ingress-nginx/pull/5952) Refactor shutdown e2e tests
- [X] [#5957](https://github.com/kubernetes/ingress-nginx/pull/5957) bump fsnotify to v1.4.9
- [X] [#5958](https://github.com/kubernetes/ingress-nginx/pull/5958) Disable enable-access-log-for-default-backend e2e test
- [X] [#5984](https://github.com/kubernetes/ingress-nginx/pull/5984) Fix panic in ingress class validation
- [X] [#5986](https://github.com/kubernetes/ingress-nginx/pull/5986) Migrate to klog v2
- [X] [#5987](https://github.com/kubernetes/ingress-nginx/pull/5987) Fix wait times in e2e tests
- [X] [#5990](https://github.com/kubernetes/ingress-nginx/pull/5990) Fix nginx command env variable reference
- [X] [#6004](https://github.com/kubernetes/ingress-nginx/pull/6004) Update nginx to 1.19.2
- [X] [#6006](https://github.com/kubernetes/ingress-nginx/pull/6006) Update nginx image
- [X] [#6007](https://github.com/kubernetes/ingress-nginx/pull/6007) Update e2e-test-runner image
- [X] [#6008](https://github.com/kubernetes/ingress-nginx/pull/6008) Rollback update of Jaeger library to 0.5.0 and update datadog to 1.2.0
- [X] [#6014](https://github.com/kubernetes/ingress-nginx/pull/6014) Update go dependencies
- [X] [#6039](https://github.com/kubernetes/ingress-nginx/pull/6039) Add configurable serviceMonitor metricRelabelling and targetLabels
- [X] [#6046](https://github.com/kubernetes/ingress-nginx/pull/6046) Add new Dockerfile label org.opencontainers.image.revision
- [X] [#6047](https://github.com/kubernetes/ingress-nginx/pull/6047) Increase wait times in e2e tests
- [X] [#6049](https://github.com/kubernetes/ingress-nginx/pull/6049) Improve docs and logging for --ingress-class usage
- [X] [#6052](https://github.com/kubernetes/ingress-nginx/pull/6052) Fix flaky e2e test
- [X] [#6056](https://github.com/kubernetes/ingress-nginx/pull/6056) Rollback to Poll instead of PollImmediate
- [X] [#6062](https://github.com/kubernetes/ingress-nginx/pull/6062) Adjust e2e timeouts
- [X] [#6063](https://github.com/kubernetes/ingress-nginx/pull/6063) Remove file system paths executables
- [X] [#6080](https://github.com/kubernetes/ingress-nginx/pull/6080) Use k8s.gcr.io vanity domain

_Documentation:_

- [X] [#5911](https://github.com/kubernetes/ingress-nginx/pull/5911) Change image URL after switching to gcr.io in upgrade guide
- [X] [#5926](https://github.com/kubernetes/ingress-nginx/pull/5926) fixed some typos
- [X] [#5927](https://github.com/kubernetes/ingress-nginx/pull/5927) Simplify development doc
- [X] [#5942](https://github.com/kubernetes/ingress-nginx/pull/5942) Fix default backend flaking e2e test for default
- [X] [#5943](https://github.com/kubernetes/ingress-nginx/pull/5943) Add repo SECURITY.md
- [X] [#5965](https://github.com/kubernetes/ingress-nginx/pull/5965) feat(baremetal): Add kustomization.yaml
- [X] [#5971](https://github.com/kubernetes/ingress-nginx/pull/5971) Fixed typo "permanen"
- [X] [#5980](https://github.com/kubernetes/ingress-nginx/pull/5980) fix for 5590
- [X] [#5994](https://github.com/kubernetes/ingress-nginx/pull/5994) Clean up minikube installation instructions
- [X] [#6000](https://github.com/kubernetes/ingress-nginx/pull/6000) Fix error message formatting
- [X] [#6001](https://github.com/kubernetes/ingress-nginx/pull/6001) Update psp example
- [X] [#6020](https://github.com/kubernetes/ingress-nginx/pull/6020) Added missing backend protocol.
- [X] [#6022](https://github.com/kubernetes/ingress-nginx/pull/6022) fix typo in development.md
- [X] [#6038](https://github.com/kubernetes/ingress-nginx/pull/6038) Document migration path to ingress-nginx chart from stable/nginx-ingress
- [X] [#6042](https://github.com/kubernetes/ingress-nginx/pull/6042) Fix typo
- [X] [#6044](https://github.com/kubernetes/ingress-nginx/pull/6044) Chart readme fixes
- [X] [#6075](https://github.com/kubernetes/ingress-nginx/pull/6075) Sync helm chart affinity examples
- [X] [#6076](https://github.com/kubernetes/ingress-nginx/pull/6076) Update NLB idle timeout information

### 0.34.1

**Image:**

-   `us.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.1@sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20`
-   `eu.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.1@sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20`
- `asia.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.1@sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20`

Fix regression introduced in [#5691](https://github.com/kubernetes/ingress-nginx/pull/5691) related to annotations `use-regex` and `rewrite`.
Update go to [1.14.5](https://groups.google.com/g/golang-announce/c/XZNfaiwgt2w/m/E6gHDs32AQAJ)

_Changes:_

- [X] [#5896](https://github.com/kubernetes/ingress-nginx/pull/5896) Revert "use-regex annotation should be applied to only one Location"
- [X] [#5897](https://github.com/kubernetes/ingress-nginx/pull/5897) Update go version

### 0.34.0

**Image:**

-   `us.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.0@sha256:56633bd00dab33d92ba14c6e709126a762d54a75a6e72437adefeaaca0abb069`
-   `eu.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.0@sha256:56633bd00dab33d92ba14c6e709126a762d54a75a6e72437adefeaaca0abb069`
- `asia.gcr.io/k8s-artifacts-prod/ingress-nginx/controller:v0.34.0@sha256:56633bd00dab33d92ba14c6e709126a762d54a75a6e72437adefeaaca0abb069`

_Breaking Changes:_

The repository https://quay.io/repository/kubernetes-ingress-controller/nginx-ingress-controller is deprecated and read-only for historical purposes.

_New Features:_

- NGINX 1.19.1
- OWASP ModSecurity Core Rule Set [v3.3.0](https://github.com/coreruleset/coreruleset/releases/tag/v3.3.0)
- Configure User-Agent for [client-go](https://github.com/kubernetes/ingress-nginx/pull/5700)
- Switch to [gcr.io](https://cloud.google.com/container-registry/) as container registry
- Use cloud-build to build [container images](https://console.cloud.google.com/gcr/images/k8s-artifacts-prod/US/ingress-nginx)
- Publish images using [Container Image Promoter](https://github.com/kubernetes-sigs/k8s-container-image-promoter)
- Go 1.14.4
- client-go v0.18.5

_Changes:_

- [X] [#5671](https://github.com/kubernetes/ingress-nginx/pull/5671) Make liveness probe more resilient than readiness probe
- [X] [#5691](https://github.com/kubernetes/ingress-nginx/pull/5691) use-regex annotation should be applied to only one Location
- [X] [#5700](https://github.com/kubernetes/ingress-nginx/pull/5700) Configure User-Agent
- [X] [#5702](https://github.com/kubernetes/ingress-nginx/pull/5702) Filter out objects that belong to Helm
- [X] [#5703](https://github.com/kubernetes/ingress-nginx/pull/5703) Use build tags to make it compile on non linux platforms
- [X] [#5704](https://github.com/kubernetes/ingress-nginx/pull/5704) add custom metric to hpa template
- [X] [#5707](https://github.com/kubernetes/ingress-nginx/pull/5707) fix for #5666
- [X] [#5708](https://github.com/kubernetes/ingress-nginx/pull/5708) Add sysctl exemptions to controller PSP
- [X] [#5709](https://github.com/kubernetes/ingress-nginx/pull/5709) fix: remove duplicated X-Forwarded-Proto header.
- [X] [#5712](https://github.com/kubernetes/ingress-nginx/pull/5712) Add proxy-ssl-server-name to pass server name on SNI
- [X] [#5713](https://github.com/kubernetes/ingress-nginx/pull/5713) Fix static manifests location
- [X] [#5717](https://github.com/kubernetes/ingress-nginx/pull/5717) Add support for an internal load balancer along with an external one
- [X] [#5722](https://github.com/kubernetes/ingress-nginx/pull/5722) Remove deprecrated --generator flag
- [X] [#5725](https://github.com/kubernetes/ingress-nginx/pull/5725) Fix e2e externalName test
- [X] [#5726](https://github.com/kubernetes/ingress-nginx/pull/5726) Dynamic LB sync non-external backends only when necessary
- [X] [#5727](https://github.com/kubernetes/ingress-nginx/pull/5727) Adjust E2E_NODES variable
- [X] [#5733](https://github.com/kubernetes/ingress-nginx/pull/5733) Fix make-clean target
- [X] [#5742](https://github.com/kubernetes/ingress-nginx/pull/5742) Allow to use a custom arch to run e2e tests
- [X] [#5743](https://github.com/kubernetes/ingress-nginx/pull/5743) build/dev-env.sh: remove docker version check
- [X] [#5745](https://github.com/kubernetes/ingress-nginx/pull/5745) Add default-type as a configurable for default_type
- [X] [#5747](https://github.com/kubernetes/ingress-nginx/pull/5747) feat: enable stream_realip_module
- [X] [#5749](https://github.com/kubernetes/ingress-nginx/pull/5749) Be configurable max batch size of metrics
- [X] [#5752](https://github.com/kubernetes/ingress-nginx/pull/5752) Update ValidatingWebhook for Ingress to support --dry-run=server
- [X] [#5761](https://github.com/kubernetes/ingress-nginx/pull/5761) Improve buildx configuration
- [X] [#5762](https://github.com/kubernetes/ingress-nginx/pull/5762) Add cloudbuild job to build nginx image
- [X] [#5764](https://github.com/kubernetes/ingress-nginx/pull/5764) Remove vendor directory and enable go modules
- [X] [#5770](https://github.com/kubernetes/ingress-nginx/pull/5770) Use admissionregistration.k8s.io/v1beta1 to be k8s < 1.16 compatible
- [X] [#5773](https://github.com/kubernetes/ingress-nginx/pull/5773) Update go dependencies
- [X] [#5774](https://github.com/kubernetes/ingress-nginx/pull/5774) Build new nginx image
- [X] [#5777](https://github.com/kubernetes/ingress-nginx/pull/5777) Improve execution of prow jobs
- [X] [#5778](https://github.com/kubernetes/ingress-nginx/pull/5778) Simplify build of e2e test image
- [X] [#5779](https://github.com/kubernetes/ingress-nginx/pull/5779) Add version check form helm
- [X] [#5791](https://github.com/kubernetes/ingress-nginx/pull/5791) Trigger of cloudbuild for nginx image and push
- [X] [#5794](https://github.com/kubernetes/ingress-nginx/pull/5794) Remove use of terraform to build nginx and ingress-controller
- [X] [#5795](https://github.com/kubernetes/ingress-nginx/pull/5795) Use fully qualified images to avoid cri-o issues
- [X] [#5796](https://github.com/kubernetes/ingress-nginx/pull/5796) Fixup docs for the ingress-class flag.
- [X] [#5797](https://github.com/kubernetes/ingress-nginx/pull/5797) Add cloudbuild configuration for cfssl test image
- [X] [#5798](https://github.com/kubernetes/ingress-nginx/pull/5798) Add cloudbuild configuration for echo test image
- [X] [#5799](https://github.com/kubernetes/ingress-nginx/pull/5799) Add cloudbuild configuration for httpbin test image
- [X] [#5800](https://github.com/kubernetes/ingress-nginx/pull/5800) Add cloudbuild configuration for fastcgi test image
- [X] [#5804](https://github.com/kubernetes/ingress-nginx/pull/5804) Increase cloudbuild timeout
- [X] [#5806](https://github.com/kubernetes/ingress-nginx/pull/5806) Use a different machine type for httpbin cloudbuild
- [X] [#5807](https://github.com/kubernetes/ingress-nginx/pull/5807) Start using e2e test images from gcr.io
- [X] [#5808](https://github.com/kubernetes/ingress-nginx/pull/5808) Remove unused variables and verbose e2e logs
- [X] [#5811](https://github.com/kubernetes/ingress-nginx/pull/5811) Get ingress-controller pod name
- [X] [#5817](https://github.com/kubernetes/ingress-nginx/pull/5817) Update kind node image version
- [X] [#5818](https://github.com/kubernetes/ingress-nginx/pull/5818) Update e2e configuration and fix flaky test
- [X] [#5823](https://github.com/kubernetes/ingress-nginx/pull/5823) Add quoting to sysctls because numeric values need to be strings
- [X] [#5826](https://github.com/kubernetes/ingress-nginx/pull/5826) Use nginx image promoted from staging
- [X] [#5827](https://github.com/kubernetes/ingress-nginx/pull/5827) Extract version to file VERSION
- [X] [#5828](https://github.com/kubernetes/ingress-nginx/pull/5828) Switch to promoted e2e images in gcr
- [X] [#5832](https://github.com/kubernetes/ingress-nginx/pull/5832) fix: remove redundant health check to avoid liveness or readiness timeout
- [X] [#5836](https://github.com/kubernetes/ingress-nginx/pull/5836) Test pull requests using github actions
- [X] [#5840](https://github.com/kubernetes/ingress-nginx/pull/5840) Update nginx image registry
- [X] [#5843](https://github.com/kubernetes/ingress-nginx/pull/5843) Update jettech/kube-webhook-certgen image
- [X] [#5845](https://github.com/kubernetes/ingress-nginx/pull/5845) Update deploy manifests
- [X] [#5846](https://github.com/kubernetes/ingress-nginx/pull/5846) Filter github actions to be executed
- [X] [#5849](https://github.com/kubernetes/ingress-nginx/pull/5849) Fix github docs github action
- [X] [#5851](https://github.com/kubernetes/ingress-nginx/pull/5851) Fix conflicts of VERSION file
- [X] [#5853](https://github.com/kubernetes/ingress-nginx/pull/5853) fix json tag for SSLPreferServerCiphers
- [X] [#5856](https://github.com/kubernetes/ingress-nginx/pull/5856) Fix proxy ssl e2e test
- [X] [#5857](https://github.com/kubernetes/ingress-nginx/pull/5857) Custom default backend service must have ports
- [X] [#5859](https://github.com/kubernetes/ingress-nginx/pull/5859) Update nginx to 1.19.1
- [X] [#5861](https://github.com/kubernetes/ingress-nginx/pull/5861) Update nginx modules
- [X] [#5862](https://github.com/kubernetes/ingress-nginx/pull/5862) Adjust nginx cloudbuild timeout
- [X] [#5866](https://github.com/kubernetes/ingress-nginx/pull/5866) Update OWASP ModSecurity Core Rule Set
- [X] [#5867](https://github.com/kubernetes/ingress-nginx/pull/5867) Update nginx image
- [X] [#5868](https://github.com/kubernetes/ingress-nginx/pull/5868) Update go dependencies
- [X] [#5869](https://github.com/kubernetes/ingress-nginx/pull/5869) Changes in TAG file should trigger e2e testing

_Documentation:_

- [X] [#5706](https://github.com/kubernetes/ingress-nginx/pull/5706) Fix controller.publishService.enabled on README
- [X] [#5711](https://github.com/kubernetes/ingress-nginx/pull/5711) Update mkdocs
- [X] [#5724](https://github.com/kubernetes/ingress-nginx/pull/5724) Update troubleshooting.md
- [X] [#5729](https://github.com/kubernetes/ingress-nginx/pull/5729) docs: update development.md
- [X] [#5751](https://github.com/kubernetes/ingress-nginx/pull/5751) docs: update development.md to use ingress-nginx-*
- [X] [#5759](https://github.com/kubernetes/ingress-nginx/pull/5759) Update helm chart name in upgrade doc
- [X] [#5760](https://github.com/kubernetes/ingress-nginx/pull/5760) Update comment about restart of pod
- [X] [#5763](https://github.com/kubernetes/ingress-nginx/pull/5763) Update e2e test suite index list doc
- [X] [#5819](https://github.com/kubernetes/ingress-nginx/pull/5819) Update krew installation doc
- [X] [#5821](https://github.com/kubernetes/ingress-nginx/pull/5821) doc: update docs and fixed typos

### 0.33.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.33.0`

_New Features:_

- NGINX 1.19.0
- TLSv1.3 is enabled by default
- Experimental support for s390x
- Allow combination of NGINX variables in annotation [upstream-hash-by](https://github.com/kubernetes/ingress-nginx/pull/5571)
- New setting to configure different access logs for http and stream sections: [http-access-log-path](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#http-access-log-path) and [stream-access-log-path](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#stream-access-log-path) options in configMap

_Deprecations:_

- Setting [access-log-path](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#access-log-path) is deprecated and will be removed in 0.35.0. Please use [http-access-log-path](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#http-access-log-path) and [stream-access-log-path](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#stream-access-log-path)

_Changes:_

- [X] [#5463](https://github.com/kubernetes/ingress-nginx/pull/5463) Wait before any request to the ingress controller pod
- [X] [#5488](https://github.com/kubernetes/ingress-nginx/pull/5488) Update kind
- [X] [#5491](https://github.com/kubernetes/ingress-nginx/pull/5491) Actually enable TLSv1.3 by default
- [X] [#5494](https://github.com/kubernetes/ingress-nginx/pull/5494) Add configuration option for the runAsUser parameter of the webhook patch job
- [X] [#5503](https://github.com/kubernetes/ingress-nginx/pull/5503) Update job-patchWebhook.yaml
- [X] [#5504](https://github.com/kubernetes/ingress-nginx/pull/5504) Add configuration option for the imagePullSecrets in the webhook jobs
- [X] [#5505](https://github.com/kubernetes/ingress-nginx/pull/5505) Update helm chart
- [X] [#5516](https://github.com/kubernetes/ingress-nginx/pull/5516) build: remove unnecessary tag line in e2e
- [X] [#5522](https://github.com/kubernetes/ingress-nginx/pull/5522) Remove duplicate annotation parsing for annotationAffinityCookieChangeOnFailure
- [X] [#5534](https://github.com/kubernetes/ingress-nginx/pull/5534) Add annotation ssl-prefer-server-ciphers.
- [X] [#5536](https://github.com/kubernetes/ingress-nginx/pull/5536) Fix error setting $service_name NGINX variable
- [X] [#5553](https://github.com/kubernetes/ingress-nginx/pull/5553) Check service If publish-service flag is defined
- [X] [#5571](https://github.com/kubernetes/ingress-nginx/pull/5571) feat: support the combination of Nginx variables for annotation upstream-hash-by.
- [X] [#5572](https://github.com/kubernetes/ingress-nginx/pull/5572) [chart] Add toleration support for admission webhooks
- [X] [#5578](https://github.com/kubernetes/ingress-nginx/pull/5578) Use image promoter to push images to gcr
- [X] [#5582](https://github.com/kubernetes/ingress-nginx/pull/5582) Allow pulling images by digest
- [X] [#5584](https://github.com/kubernetes/ingress-nginx/pull/5584) Add note about initial delay during first start
- [X] [#5586](https://github.com/kubernetes/ingress-nginx/pull/5586) Add MaxMind GeoIP2 Anonymous IP support
- [X] [#5589](https://github.com/kubernetes/ingress-nginx/pull/5589) Do not reload NGINX if master process dies
- [X] [#5596](https://github.com/kubernetes/ingress-nginx/pull/5596) Update go dependencies
- [X] [#5603](https://github.com/kubernetes/ingress-nginx/pull/5603) Update nginx to 1.19.0
- [X] [#5604](https://github.com/kubernetes/ingress-nginx/pull/5604) Update debian-base image
- [X] [#5606](https://github.com/kubernetes/ingress-nginx/pull/5606) Update nginx image and go to 1.14.3
- [X] [#5613](https://github.com/kubernetes/ingress-nginx/pull/5613) fix oauth2-proxy image repository
- [X] [#5614](https://github.com/kubernetes/ingress-nginx/pull/5614) Add support for s390x
- [X] [#5619](https://github.com/kubernetes/ingress-nginx/pull/5619) Use new multi-arch nginx image
- [X] [#5621](https://github.com/kubernetes/ingress-nginx/pull/5621) Update terraform build images
- [X] [#5624](https://github.com/kubernetes/ingress-nginx/pull/5624) feat: add lj-releng tool to check Lua code for finding the potential problems
- [X] [#5625](https://github.com/kubernetes/ingress-nginx/pull/5625) Update nginx image to use alpine 3.12
- [X] [#5626](https://github.com/kubernetes/ingress-nginx/pull/5626) Update nginx image
- [X] [#5629](https://github.com/kubernetes/ingress-nginx/pull/5629) Build multi-arch images by default
- [X] [#5630](https://github.com/kubernetes/ingress-nginx/pull/5630) Fix makefile task names
- [X] [#5631](https://github.com/kubernetes/ingress-nginx/pull/5631) Update e2e image
- [X] [#5632](https://github.com/kubernetes/ingress-nginx/pull/5632) Update buildx progress configuration
- [X] [#5636](https://github.com/kubernetes/ingress-nginx/pull/5636) Enable coredumps for e2e tests
- [X] [#5637](https://github.com/kubernetes/ingress-nginx/pull/5637) Refactor build of docker images
- [X] [#5641](https://github.com/kubernetes/ingress-nginx/pull/5641) Add missing ARCH variable
- [X] [#5642](https://github.com/kubernetes/ingress-nginx/pull/5642) Fix dev-env makefile task
- [X] [#5643](https://github.com/kubernetes/ingress-nginx/pull/5643) Fix build of image on osx
- [X] [#5644](https://github.com/kubernetes/ingress-nginx/pull/5644) Remove copy of binaries and deprecated e2e task
- [X] [#5656](https://github.com/kubernetes/ingress-nginx/pull/5656) feat: add http-access-log-path and stream-access-log-path options in configMap
- [X] [#5659](https://github.com/kubernetes/ingress-nginx/pull/5659) Update cloud-build configuration
- [X] [#5660](https://github.com/kubernetes/ingress-nginx/pull/5660) Set missing USER in cloud-build
- [X] [#5661](https://github.com/kubernetes/ingress-nginx/pull/5661) Add missing REPO_INFO en variable to cloud-build
- [X] [#5662](https://github.com/kubernetes/ingress-nginx/pull/5662) Increase cloud-build timeout
- [X] [#5663](https://github.com/kubernetes/ingress-nginx/pull/5663) Fix cloud-timeout setting
- [X] [#5664](https://github.com/kubernetes/ingress-nginx/pull/5664) fix undefined variable $auth_cookie error due to when location is denied
- [X] [#5665](https://github.com/kubernetes/ingress-nginx/pull/5665) Fix: improve performance
- [X] [#5669](https://github.com/kubernetes/ingress-nginx/pull/5669) Serve correct TLS certificate for requests with uppercase host
- [X] [#5672](https://github.com/kubernetes/ingress-nginx/pull/5672) feat: enable lj-releng tool to lint lua code.
- [X] [#5684](https://github.com/kubernetes/ingress-nginx/pull/5684) Fix proxy_protocol duplication in listen definition

_Documentation:_

- [X] [#5487](https://github.com/kubernetes/ingress-nginx/pull/5487) Add note about firewall ports for admission webhook
- [X] [#5512](https://github.com/kubernetes/ingress-nginx/pull/5512) Wrong filename in documantation example
- [X] [#5563](https://github.com/kubernetes/ingress-nginx/pull/5563) Use ingress-nginx-* naming in docs to match the default deployments
- [X] [#5566](https://github.com/kubernetes/ingress-nginx/pull/5566) Update configmap name in custom-headers/README.md
- [X] [#5639](https://github.com/kubernetes/ingress-nginx/pull/5639) Update timeout to align values
- [X] [#5646](https://github.com/kubernetes/ingress-nginx/pull/5646) Add minor doc fixes to user guide and chart readme
- [X] [#5652](https://github.com/kubernetes/ingress-nginx/pull/5652) Add documentation for loading e2e tests without using minikube
- [X] [#5677](https://github.com/kubernetes/ingress-nginx/pull/5677) Add URL to official grafana dashboards
- [X] [#5682](https://github.com/kubernetes/ingress-nginx/pull/5682) Fix typo

### 0.32.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0`

Fix regression in validating webhook when the ingress controller is installed in Kubernetes v1.18

_Changes:_

- [X] [#4271](https://github.com/kubernetes/ingress-nginx/pull/4271) Add support for multi-arch images
- [X] [#5429](https://github.com/kubernetes/ingress-nginx/pull/5429) Update krew plugin configuration
- [X] [#5430](https://github.com/kubernetes/ingress-nginx/pull/5430) Use github actions to create releases and krew plugin assets
- [X] [#5432](https://github.com/kubernetes/ingress-nginx/pull/5432) Allow releases from a github action
- [X] [#5433](https://github.com/kubernetes/ingress-nginx/pull/5433) Avoid removal of index.yaml file
- [X] [#5434](https://github.com/kubernetes/ingress-nginx/pull/5434) Disable PR against krew repository
- [X] [#5436](https://github.com/kubernetes/ingress-nginx/pull/5436) Disable github release action
- [X] [#5439](https://github.com/kubernetes/ingress-nginx/pull/5439) Change action order
- [X] [#5453](https://github.com/kubernetes/ingress-nginx/pull/5453) Ensure alpine packages are up to date
- [X] [#5456](https://github.com/kubernetes/ingress-nginx/pull/5456) Case-insensitive TLS host matching
- [X] [#5459](https://github.com/kubernetes/ingress-nginx/pull/5459) Refactor ingress validation in webhook
- [X] [#5461](https://github.com/kubernetes/ingress-nginx/pull/5461) Fix helper for defaultbackend name
- [X] [#5462](https://github.com/kubernetes/ingress-nginx/pull/5462) Remove noisy dns log
- [X] [#5469](https://github.com/kubernetes/ingress-nginx/pull/5469) Changes on services must trigger a sync event
- [X] [#5472](https://github.com/kubernetes/ingress-nginx/pull/5472) Update admission webhook image
- [X] [#5474](https://github.com/kubernetes/ingress-nginx/pull/5474) Add install command for Digital Ocean
- [X] [#5476](https://github.com/kubernetes/ingress-nginx/pull/5476) Fix chart missing default backend name
- [X] [#5481](https://github.com/kubernetes/ingress-nginx/pull/5481) fix first backend sync
- [X] [#5483](https://github.com/kubernetes/ingress-nginx/pull/5483) Fix chart maxmindLicenseKey location
- [X] [#5484](https://github.com/kubernetes/ingress-nginx/pull/5484) Only load docker images in kind worker nodes

_Documentation:_

- [X] [#5404](https://github.com/kubernetes/ingress-nginx/pull/5404) update the helm v3 install way
- [X] [#5435](https://github.com/kubernetes/ingress-nginx/pull/5435) Fix deployment links
- [X] [#5438](https://github.com/kubernetes/ingress-nginx/pull/5438) Update chart instructions
- [X] [#5460](https://github.com/kubernetes/ingress-nginx/pull/5460) fix(Chart): Mismatch between README.md and values.yml (defaultBackend.enabled)
- [X] [#5465](https://github.com/kubernetes/ingress-nginx/pull/5465) Update helm v2 installation instructions
- [X] [#5468](https://github.com/kubernetes/ingress-nginx/pull/5468) Update admission webhook annotations
- [X] [#5479](https://github.com/kubernetes/ingress-nginx/pull/5479) Remove obsolete default backend settings
- [X] [#5480](https://github.com/kubernetes/ingress-nginx/pull/5480) docs(changelog): fix typo

### 0.31.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.31.1`

Fix regression in validating webhook

- [X] [#5445](https://github.com/kubernetes/ingress-nginx/pull/5445) Ensure webhook validation ingress has a PathType

### 0.31.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.31.0`

_New Features:_

- NGINX 1.17.10
- OpenSSL 1.1.1g - [CVE-2020-1967](https://cve.mitre.org/cgi-bin/cvename.cgi?name=2020-1967)
- OCSP stapling
- Helm chart [stable/nginx-ingress](https://github.com/helm/charts/tree/master/stable/nginx-ingress) is now maintained in the [ingress-nginx](https://github.com/kubernetes/ingress-nginx/tree/master/charts/ingress-nginx) repository
- Support for custom Maxmind GeoLite2 Databases [flag --maxmind-edition-ids](https://kubernetes.github.io/ingress-nginx/user-guide/cli-arguments/)
- New [PathType](https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types) and [IngressClass](https://kubernetes.io/docs/concepts/services-networking/ingress/#ingress-class) fields. Requires Kubernetes v1.18 or higher
- Enable configuration of lua plugins using the configuration configmap
- Go 1.14

_Changes:_

- [X] [#4632](https://github.com/kubernetes/ingress-nginx/pull/4632) run lua plugin tests
- [X] [#4958](https://github.com/kubernetes/ingress-nginx/pull/4958) Add a forwarded protocol map for included x-forwarded-proto.
- [X] [#4981](https://github.com/kubernetes/ingress-nginx/pull/4981) Applying proxy-ssl-* directives on locations only
- [X] [#5131](https://github.com/kubernetes/ingress-nginx/pull/5131) Add request handling performance dashboard
- [X] [#5133](https://github.com/kubernetes/ingress-nginx/pull/5133) Lua OCSP stapling
- [X] [#5157](https://github.com/kubernetes/ingress-nginx/pull/5157) Added limit-rate annotation test
- [X] [#5158](https://github.com/kubernetes/ingress-nginx/pull/5158) Fix push task
- [X] [#5159](https://github.com/kubernetes/ingress-nginx/pull/5159) Start migration of helm chart
- [X] [#5160](https://github.com/kubernetes/ingress-nginx/pull/5160) Fix e2e test run.sh
- [X] [#5165](https://github.com/kubernetes/ingress-nginx/pull/5165) Use local chart directory for dev-env and e2e tests
- [X] [#5166](https://github.com/kubernetes/ingress-nginx/pull/5166) proxy_ssl_name support
- [X] [#5169](https://github.com/kubernetes/ingress-nginx/pull/5169) Cleanup e2e directory
- [X] [#5170](https://github.com/kubernetes/ingress-nginx/pull/5170) Update go dependencies
- [X] [#5171](https://github.com/kubernetes/ingress-nginx/pull/5171) Sync chart PR #20984
- [X] [#5172](https://github.com/kubernetes/ingress-nginx/pull/5172) Add script to check helm chart
- [X] [#5173](https://github.com/kubernetes/ingress-nginx/pull/5173) Update go to 1.14
- [X] [#5174](https://github.com/kubernetes/ingress-nginx/pull/5174) Update e2e image
- [X] [#5175](https://github.com/kubernetes/ingress-nginx/pull/5175) Migrate the backends handle logic to function
- [X] [#5178](https://github.com/kubernetes/ingress-nginx/pull/5178) Adding annotations support to helm chart configmaps
- [X] [#5181](https://github.com/kubernetes/ingress-nginx/pull/5181) Fix public function comment
- [X] [#5182](https://github.com/kubernetes/ingress-nginx/pull/5182) Update go mod for 1.14
- [X] [#5183](https://github.com/kubernetes/ingress-nginx/pull/5183) Remove unused docker file
- [X] [#5185](https://github.com/kubernetes/ingress-nginx/pull/5185) [helm chart] Use recommended labels and label helpers
- [X] [#5190](https://github.com/kubernetes/ingress-nginx/pull/5190) Refactored test/e2e/annotations/proxy.go
- [X] [#5192](https://github.com/kubernetes/ingress-nginx/pull/5192) Update helm templates to match new chart name
- [X] [#5194](https://github.com/kubernetes/ingress-nginx/pull/5194) I found a typo :)
- [X] [#5201](https://github.com/kubernetes/ingress-nginx/pull/5201) Added TC for proxy connect, read, and send timeout
- [X] [#5202](https://github.com/kubernetes/ingress-nginx/pull/5202) Refactored client body buffer size TC-s.
- [X] [#5204](https://github.com/kubernetes/ingress-nginx/pull/5204) Cleanup chart code
- [X] [#5205](https://github.com/kubernetes/ingress-nginx/pull/5205) Add OWNERS file for helm chart
- [X] [#5207](https://github.com/kubernetes/ingress-nginx/pull/5207) [helm chart] Hardcode component names.
- [X] [#5211](https://github.com/kubernetes/ingress-nginx/pull/5211) Update NGINX to 1.17.9
- [X] [#5213](https://github.com/kubernetes/ingress-nginx/pull/5213) Make quote function to render pointers in the template properly
- [X] [#5216](https://github.com/kubernetes/ingress-nginx/pull/5216) Check go exists in $PATH
- [X] [#5217](https://github.com/kubernetes/ingress-nginx/pull/5217) Added affinity-mode tc and refactored affinity.go
- [X] [#5221](https://github.com/kubernetes/ingress-nginx/pull/5221) Update NGINX image
- [X] [#5225](https://github.com/kubernetes/ingress-nginx/pull/5225) Avoid secret without tls.crt and tls.key but a valid ca.crt
- [X] [#5226](https://github.com/kubernetes/ingress-nginx/pull/5226) Fix $service_name and $service_port variables values without host
- [X] [#5232](https://github.com/kubernetes/ingress-nginx/pull/5232) Refacored proxy ssl TC-s
- [X] [#5241](https://github.com/kubernetes/ingress-nginx/pull/5241) Fix controller container name
- [X] [#5246](https://github.com/kubernetes/ingress-nginx/pull/5246) Remove checks for older versions
- [X] [#5249](https://github.com/kubernetes/ingress-nginx/pull/5249) Add support for hostPort in Deployment
- [X] [#5250](https://github.com/kubernetes/ingress-nginx/pull/5250) Use rbac scope feature in e2e tests
- [X] [#5251](https://github.com/kubernetes/ingress-nginx/pull/5251) Add support for custom healthz path in helm chart
- [X] [#5252](https://github.com/kubernetes/ingress-nginx/pull/5252) Check chart controller image tag
- [X] [#5254](https://github.com/kubernetes/ingress-nginx/pull/5254) Switch dev-env script to deployment
- [X] [#5258](https://github.com/kubernetes/ingress-nginx/pull/5258) Cleanup of chart labels
- [X] [#5262](https://github.com/kubernetes/ingress-nginx/pull/5262) Add Maxmind Editions support
- [X] [#5264](https://github.com/kubernetes/ingress-nginx/pull/5264) Fix reference to DH param secret, recommend larger parameter size
- [X] [#5266](https://github.com/kubernetes/ingress-nginx/pull/5266) Redirect for app-root should preserve current scheme
- [X] [#5268](https://github.com/kubernetes/ingress-nginx/pull/5268) do not require go for building
- [X] [#5269](https://github.com/kubernetes/ingress-nginx/pull/5269) Ensure DeleteDeployment waits until there are no pods running
- [X] [#5276](https://github.com/kubernetes/ingress-nginx/pull/5276) Fix the ability to disable ModSecurity at location level
- [X] [#5277](https://github.com/kubernetes/ingress-nginx/pull/5277) refactoring: use more specific var name
- [X] [#5281](https://github.com/kubernetes/ingress-nginx/pull/5281) Remove unnecessary logs
- [X] [#5283](https://github.com/kubernetes/ingress-nginx/pull/5283) Add retries for dns in tcp e2e test
- [X] [#5284](https://github.com/kubernetes/ingress-nginx/pull/5284) Wait for update in tcp e2e test
- [X] [#5288](https://github.com/kubernetes/ingress-nginx/pull/5288) Update client-go methods to support context and and new options
- [X] [#5289](https://github.com/kubernetes/ingress-nginx/pull/5289) Update go and e2e image
- [X] [#5290](https://github.com/kubernetes/ingress-nginx/pull/5290) Add DS_PROMETHEUS datasource for templating
- [X] [#5296](https://github.com/kubernetes/ingress-nginx/pull/5296) Added proxy-ssl-location-only test.
- [X] [#5298](https://github.com/kubernetes/ingress-nginx/pull/5298) Increase e2e concurrency
- [X] [#5301](https://github.com/kubernetes/ingress-nginx/pull/5301) Forward X-Request-ID to auth service
- [X] [#5307](https://github.com/kubernetes/ingress-nginx/pull/5307) Migrate ingress.class annotation to new IngressClassName field
- [X] [#5308](https://github.com/kubernetes/ingress-nginx/pull/5308) Set new default PathType to prefix
- [X] [#5309](https://github.com/kubernetes/ingress-nginx/pull/5309) Fix condition in server-alias annotation
- [X] [#5310](https://github.com/kubernetes/ingress-nginx/pull/5310) Added auth-tls-verify-client testcase
- [X] [#5313](https://github.com/kubernetes/ingress-nginx/pull/5313) Add script to generate yaml files from helm
- [X] [#5314](https://github.com/kubernetes/ingress-nginx/pull/5314) Set default resource requests limits
- [X] [#5315](https://github.com/kubernetes/ingress-nginx/pull/5315) Fix definition order of modsecurity directives
- [X] [#5320](https://github.com/kubernetes/ingress-nginx/pull/5320) Change condition order that produces endless loop
- [X] [#5324](https://github.com/kubernetes/ingress-nginx/pull/5324) Add support for PathTypeExact
- [X] [#5329](https://github.com/kubernetes/ingress-nginx/pull/5329) Update e2e dev image to v1.18.0
- [X] [#5330](https://github.com/kubernetes/ingress-nginx/pull/5330) Set k8s version kind should use for dev environment
- [X] [#5331](https://github.com/kubernetes/ingress-nginx/pull/5331) Enable configuration of plugins using configmap
- [X] [#5332](https://github.com/kubernetes/ingress-nginx/pull/5332) Add lifecycle hook and option to enable mimalloc
- [X] [#5333](https://github.com/kubernetes/ingress-nginx/pull/5333) Remove duplicated annotations definition and refactor hostPort conf
- [X] [#5336](https://github.com/kubernetes/ingress-nginx/pull/5336) Fix deployment strategy
- [X] [#5340](https://github.com/kubernetes/ingress-nginx/pull/5340) fix: remove unnecessary if statement when redirect annotation is defined
- [X] [#5341](https://github.com/kubernetes/ingress-nginx/pull/5341) ensure make lua-test runs locally
- [X] [#5346](https://github.com/kubernetes/ingress-nginx/pull/5346) Ensure krew plugin includes license
- [X] [#5357](https://github.com/kubernetes/ingress-nginx/pull/5357) Fix broken symlink to mimalloc
- [X] [#5361](https://github.com/kubernetes/ingress-nginx/pull/5361) Cleanup parsing of annotations with lists
- [X] [#5362](https://github.com/kubernetes/ingress-nginx/pull/5362) Cleanup httpbin image
- [X] [#5363](https://github.com/kubernetes/ingress-nginx/pull/5363) Remove version dependency in mimalloc symlink
- [X] [#5369](https://github.com/kubernetes/ingress-nginx/pull/5369) Update luajit and nginx to 1.17.10
- [X] [#5371](https://github.com/kubernetes/ingress-nginx/pull/5371) Update e2e image
- [X] [#5372](https://github.com/kubernetes/ingress-nginx/pull/5372) Update Go to 1.14.2
- [X] [#5374](https://github.com/kubernetes/ingress-nginx/pull/5374) Add port for plain HTTP to HTTPS redirection
- [X] [#5375](https://github.com/kubernetes/ingress-nginx/pull/5375) Remove chart old podSecurityPolicy check
- [X] [#5380](https://github.com/kubernetes/ingress-nginx/pull/5380) Use official mkdocs image and github action
- [X] [#5381](https://github.com/kubernetes/ingress-nginx/pull/5381) Add e2e tests for helm chart
- [X] [#5387](https://github.com/kubernetes/ingress-nginx/pull/5387) Add e2e test for OCSP and new configmap setting
- [X] [#5388](https://github.com/kubernetes/ingress-nginx/pull/5388) Remove TODO that were done
- [X] [#5392](https://github.com/kubernetes/ingress-nginx/pull/5392) Add new cfssl image and update e2e tests to use it
- [X] [#5393](https://github.com/kubernetes/ingress-nginx/pull/5393) Fix dev-env script to use new hostPort setting
- [X] [#5403](https://github.com/kubernetes/ingress-nginx/pull/5403) staple only when OCSP response status is "good"
- [X] [#5407](https://github.com/kubernetes/ingress-nginx/pull/5407) Update go dependencies
- [X] [#5409](https://github.com/kubernetes/ingress-nginx/pull/5409) Removed wrong code
- [X] [#5410](https://github.com/kubernetes/ingress-nginx/pull/5410) Add support for IngressClass and ingress.class annotation
- [X] [#5414](https://github.com/kubernetes/ingress-nginx/pull/5414) Pin mimalloc version and update openssl
- [X] [#5415](https://github.com/kubernetes/ingress-nginx/pull/5415) Update nginx image to fix openssl CVE-2020-1967
- [X] [#5419](https://github.com/kubernetes/ingress-nginx/pull/5419) Improve build time of httpbin e2e test image

_Documentation:_

- [X] [#5162](https://github.com/kubernetes/ingress-nginx/pull/5162) Migrate release of docs from travis-ci to github actions
- [X] [#5163](https://github.com/kubernetes/ingress-nginx/pull/5163) Cleanup build of documentation and update to mkdocs 1.1
- [X] [#5114](https://github.com/kubernetes/ingress-nginx/pull/5114) Feat: add header-pattern annotation.
- [X] [#5274](https://github.com/kubernetes/ingress-nginx/pull/5274) [docs]: fix deploy Prerequisite section
- [X] [#5347](https://github.com/kubernetes/ingress-nginx/pull/5347) docs: fix use-gzip wrong markdown style
- [X] [#5349](https://github.com/kubernetes/ingress-nginx/pull/5349) Update doc for validating Webhook with helm
- [X] [#5351](https://github.com/kubernetes/ingress-nginx/pull/5351) Remove deprecated flags and update docs
- [X] [#5355](https://github.com/kubernetes/ingress-nginx/pull/5355) ingress-nginx lua plugins docs
- [X] [#5360](https://github.com/kubernetes/ingress-nginx/pull/5360) Update deployment documentation
- [X] [#5365](https://github.com/kubernetes/ingress-nginx/pull/5365) Fix broken link for Layer 2 configuration mode
- [X] [#5370](https://github.com/kubernetes/ingress-nginx/pull/5370) Fix plugin README.md link
- [X] [#5395](https://github.com/kubernetes/ingress-nginx/pull/5395) Fix from-to-www link
- [X] [#5399](https://github.com/kubernetes/ingress-nginx/pull/5399) Cleanup deploy docs and remove old yaml manifests
- [X] [#5400](https://github.com/kubernetes/ingress-nginx/pull/5400) Update images README.md
- [X] [#5408](https://github.com/kubernetes/ingress-nginx/pull/5408) Add manifest for kind documentation
- [X] [#5420](https://github.com/kubernetes/ingress-nginx/pull/5420) Remove lua-resty-waf docs
- [X] [#5422](https://github.com/kubernetes/ingress-nginx/pull/5422) update notes.txt example with networking.k8s.io

### 0.30.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0`

- Allow service type ExternalName with different port and targetPort
- Update datadog tracer to v1.1.3
- Update default variables_hash_bucket_size value to 256
- Enable Opentracing for authentication subrequests (auth_request)

_Changes:_

- [X] [#5080](https://github.com/kubernetes/ingress-nginx/pull/5080) Add label selector for plugin
- [X] [#5083](https://github.com/kubernetes/ingress-nginx/pull/5083) Cleanup docker build
- [X] [#5084](https://github.com/kubernetes/ingress-nginx/pull/5084) Cleanup docker build
- [X] [#5085](https://github.com/kubernetes/ingress-nginx/pull/5085) Cleanup build of nginx image
- [X] [#5086](https://github.com/kubernetes/ingress-nginx/pull/5086) Migration e2e installation to helm
- [X] [#5087](https://github.com/kubernetes/ingress-nginx/pull/5087) Fox docker opencontainers version label
- [X] [#5088](https://github.com/kubernetes/ingress-nginx/pull/5088) Remove .cache directory with make clean.
- [X] [#5089](https://github.com/kubernetes/ingress-nginx/pull/5089) Abort any task in case of errors running shell commands
- [X] [#5090](https://github.com/kubernetes/ingress-nginx/pull/5090) Cleanup and standardization of e2e test definitions
- [X] [#5091](https://github.com/kubernetes/ingress-nginx/pull/5091) Add case for when user agent is nil
- [X] [#5092](https://github.com/kubernetes/ingress-nginx/pull/5092) Print information about e2e suite tests
- [X] [#5094](https://github.com/kubernetes/ingress-nginx/pull/5094) Remove comment from e2e_test.go
- [X] [#5095](https://github.com/kubernetes/ingress-nginx/pull/5095) Update datadog tracer to v1.1.3
- [X] [#5097](https://github.com/kubernetes/ingress-nginx/pull/5097) New e2e test: log-format-escape-json and log-format-upstream
- [X] [#5098](https://github.com/kubernetes/ingress-nginx/pull/5098) Fix make dev-env
- [X] [#5100](https://github.com/kubernetes/ingress-nginx/pull/5100) Ensure make dev-env support rolling updates
- [X] [#5101](https://github.com/kubernetes/ingress-nginx/pull/5101) Add keep-alive config check test
- [X] [#5102](https://github.com/kubernetes/ingress-nginx/pull/5102) Migrate e2e libaries
- [X] [#5103](https://github.com/kubernetes/ingress-nginx/pull/5103) Added configmap test for no-tls-redirect-locations
- [X] [#5105](https://github.com/kubernetes/ingress-nginx/pull/5105) Reuse-port check e2e tc (config check only)
- [X] [#5109](https://github.com/kubernetes/ingress-nginx/pull/5109) Added basic limit-rate configmap test.
- [X] [#5111](https://github.com/kubernetes/ingress-nginx/pull/5111) ingress-path-matching: doc typo
- [X] [#5117](https://github.com/kubernetes/ingress-nginx/pull/5117) Hash size e2e check test case
- [X] [#5122](https://github.com/kubernetes/ingress-nginx/pull/5122) refactor ssl handling in preparation of OCSP stapling
- [X] [#5123](https://github.com/kubernetes/ingress-nginx/pull/5123) Ensure helm repository and charts are available
- [X] [#5124](https://github.com/kubernetes/ingress-nginx/pull/5124) make dev-env improvements
- [X] [#5125](https://github.com/kubernetes/ingress-nginx/pull/5125) Added tc for limit-connection annotation
- [X] [#5131](https://github.com/kubernetes/ingress-nginx/pull/5131) Add request handling performance dashboard
- [X] [#5132](https://github.com/kubernetes/ingress-nginx/pull/5132) Lint go code
- [X] [#5134](https://github.com/kubernetes/ingress-nginx/pull/5134) Update list of e2e tests
- [X] [#5136](https://github.com/kubernetes/ingress-nginx/pull/5136) Add upstream keep alive tests
- [X] [#5139](https://github.com/kubernetes/ingress-nginx/pull/5139) Fixes https://github.com/kubernetes/ingress-nginx/issues/5120
- [X] [#5140](https://github.com/kubernetes/ingress-nginx/pull/5140) Added configmap test for ssl-ciphers.
- [X] [#5141](https://github.com/kubernetes/ingress-nginx/pull/5141) Allow service type ExternalName with different port and targetPort
- [X] [#5145](https://github.com/kubernetes/ingress-nginx/pull/5145) Refactor the HSTS related test file and add config check to the HSTS tests
- [X] [#5149](https://github.com/kubernetes/ingress-nginx/pull/5149) Use helm template instead of update to install dev cluster
- [X] [#5150](https://github.com/kubernetes/ingress-nginx/pull/5150) Update default VariablesHashBucketSize value to 256
- [X] [#5151](https://github.com/kubernetes/ingress-nginx/pull/5151) Check there is a difference in the template besides the checksum
- [X] [#5152](https://github.com/kubernetes/ingress-nginx/pull/5152) Clean template
- [X] [#5153](https://github.com/kubernetes/ingress-nginx/pull/5153) Update nginx and e2e images

_Documentation:_

- [X] [#5018](https://github.com/kubernetes/ingress-nginx/pull/5018) Update developer document on dependency updates
- [X] [#5081](https://github.com/kubernetes/ingress-nginx/pull/5081) Fixed incorrect documentation of cli flag --default-backend-service
- [X] [#5093](https://github.com/kubernetes/ingress-nginx/pull/5093) Generate doc with list of e2e tests
- [X] [#5135](https://github.com/kubernetes/ingress-nginx/pull/5135) Correct spelling of the word "Original" in annotations documentation

### 0.29.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.29.0`

_New Features:_

- NGINX 1.17.8
- Add SameSite support for [Cookie Affinity](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#cookie-affinity) https://www.chromium.org/updates/same-site
- Refactor of [mirror](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#mirror) feature to remove additional annotations

_Changes:_

- [X] [#4949](https://github.com/kubernetes/ingress-nginx/pull/4949) Add SameSite support - omit None for old browsers
- [X] [#4973](https://github.com/kubernetes/ingress-nginx/pull/4973) Fix release script
- [X] [#4975](https://github.com/kubernetes/ingress-nginx/pull/4975) Fix docker installation in travis script
- [X] [#4976](https://github.com/kubernetes/ingress-nginx/pull/4976) Fix travis
- [X] [#4977](https://github.com/kubernetes/ingress-nginx/pull/4977) Fix image version
- [X] [#4983](https://github.com/kubernetes/ingress-nginx/pull/4983) Fix enable opentracing per location
- [X] [#4987](https://github.com/kubernetes/ingress-nginx/pull/4987) Dump kind logs after e2e tests
- [X] [#4993](https://github.com/kubernetes/ingress-nginx/pull/4993) Calculation algorithm for server_names_hash_bucket_size should consid…
- [X] [#4995](https://github.com/kubernetes/ingress-nginx/pull/4995) Cleanup main makefile and remove the need of sed
- [X] [#4996](https://github.com/kubernetes/ingress-nginx/pull/4996) Fix status update for clusters where networking.k8s.io is not available
- [X] [#4999](https://github.com/kubernetes/ingress-nginx/pull/4999) Fix limitrange definition
- [X] [#5000](https://github.com/kubernetes/ingress-nginx/pull/5000) Update python syntax in OAuth2 example
- [X] [#5003](https://github.com/kubernetes/ingress-nginx/pull/5003) Fix server aliases
- [X] [#5008](https://github.com/kubernetes/ingress-nginx/pull/5008) Fix docker buildx check in Makefile
- [X] [#5009](https://github.com/kubernetes/ingress-nginx/pull/5009) Move mod-security logic from template to go code
- [X] [#5010](https://github.com/kubernetes/ingress-nginx/pull/5010) Update nginx image
- [X] [#5011](https://github.com/kubernetes/ingress-nginx/pull/5011) Update nginx image, go to 1.13.7 and e2e image
- [X] [#5015](https://github.com/kubernetes/ingress-nginx/pull/5015) Refactor mirror feature
- [X] [#5016](https://github.com/kubernetes/ingress-nginx/pull/5016) Fix dep-ensure task
- [X] [#5023](https://github.com/kubernetes/ingress-nginx/pull/5023) Update metric dependencies and restore default Objectives
- [X] [#5028](https://github.com/kubernetes/ingress-nginx/pull/5028) Add echo image to avoid building and installing dependencies in each …
- [X] [#5031](https://github.com/kubernetes/ingress-nginx/pull/5031) Update kindest/node version to v1.17.2
- [X] [#5032](https://github.com/kubernetes/ingress-nginx/pull/5032) Fix fortune-teller app manifest
- [X] [#5035](https://github.com/kubernetes/ingress-nginx/pull/5035) Update github.com/paultag/sniff dependency
- [X] [#5036](https://github.com/kubernetes/ingress-nginx/pull/5036) Disable DIND in script run-in-docker.sh
- [X] [#5038](https://github.com/kubernetes/ingress-nginx/pull/5038) Update code to use pault.ag/go/sniff package
- [X] [#5042](https://github.com/kubernetes/ingress-nginx/pull/5042) Fix X-Forwarded-Proto based on proxy-protocol server port
- [X] [#5050](https://github.com/kubernetes/ingress-nginx/pull/5050) Add flag to allow custom ingress status update intervals
- [X] [#5052](https://github.com/kubernetes/ingress-nginx/pull/5052) Change the handling of ConfigMap creation
- [X] [#5053](https://github.com/kubernetes/ingress-nginx/pull/5053) Validation of header in authreq should be done only in the key
- [X] [#5055](https://github.com/kubernetes/ingress-nginx/pull/5055) Only set mirror source when a target is configured
- [X] [#5059](https://github.com/kubernetes/ingress-nginx/pull/5059) Remove minikube and only use kind
- [X] [#5060](https://github.com/kubernetes/ingress-nginx/pull/5060) Cleanup e2e tests
- [X] [#5061](https://github.com/kubernetes/ingress-nginx/pull/5061) Fix scripts to run in osx
- [X] [#5062](https://github.com/kubernetes/ingress-nginx/pull/5062) Ensure scripts and dev-env works in osx
- [X] [#5067](https://github.com/kubernetes/ingress-nginx/pull/5067) Make sure set-cookie is retained from external auth endpoint
- [X] [#5069](https://github.com/kubernetes/ingress-nginx/pull/5069) Enable grpc e2e tests
- [X] [#5070](https://github.com/kubernetes/ingress-nginx/pull/5070) Update go to 1.13.8
- [X] [#5071](https://github.com/kubernetes/ingress-nginx/pull/5071) Add gzip-min-length as a Configuration Option

_Documentation:_

- [X] [#4974](https://github.com/kubernetes/ingress-nginx/pull/4974) Add travis script for docs
- [X] [#4991](https://github.com/kubernetes/ingress-nginx/pull/4991) doc: added hint why regular expressions might not be accepted
- [X] [#5018](https://github.com/kubernetes/ingress-nginx/pull/5018) Update developer document on dependency updates
- [X] [#5020](https://github.com/kubernetes/ingress-nginx/pull/5020) docs(deploy): fix helm install command for helm v3
- [X] [#5037](https://github.com/kubernetes/ingress-nginx/pull/5037) Cleanup README.md
- [X] [#5040](https://github.com/kubernetes/ingress-nginx/pull/5040) Update documentation and remove hack fixed by upstream cookie library
- [X] [#5041](https://github.com/kubernetes/ingress-nginx/pull/5041) 36.94% size reduction of image assets using lossless compression from ImgBot
- [X] [#5043](https://github.com/kubernetes/ingress-nginx/pull/5043) Cleanup docs
- [X] [#5068](https://github.com/kubernetes/ingress-nginx/pull/5068) docs: reference buildx as a requirement for docker builds
- [X] [#5073](https://github.com/kubernetes/ingress-nginx/pull/5073) oauth-external-auth: README.md: Link to oauth2-proxy, dashboard-ingress.yaml

### 0.28.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.28.0`

Fix occasional prometheus `http: superfluous response.WriteHeader call...` error [#4943](https://github.com/kubernetes/ingress-nginx/pull/4943)
Remove prometheus socket before the start of metrics collector [#4961](https://github.com/kubernetes/ingress-nginx/pull/4961)
Reduce CPU utilization when the ingress controller is shutting down [#4959](https://github.com/kubernetes/ingress-nginx/pull/4959)
Fixes a flaw (CVE-2020-8553) when auth-type basic annotation is used [#4960](https://github.com/kubernetes/ingress-nginx/pull/4960)

_Changes:_

- [X] [#4912](https://github.com/kubernetes/ingress-nginx/pull/4912) Update README.md
- [X] [#4914](https://github.com/kubernetes/ingress-nginx/pull/4914) Disable docker in docker tasks in terraform release script
- [X] [#4932](https://github.com/kubernetes/ingress-nginx/pull/4932) Cleanup dev-env script
- [X] [#4943](https://github.com/kubernetes/ingress-nginx/pull/4943) Update client_golang dependency to v1.3.0
- [X] [#4956](https://github.com/kubernetes/ingress-nginx/pull/4956) Fix proxy protocol support for X-Forwarded-Port
- [X] [#4959](https://github.com/kubernetes/ingress-nginx/pull/4959) Refactor how to handle sigterm and nginx process goroutine
- [X] [#4960](https://github.com/kubernetes/ingress-nginx/pull/4960) Avoid overlap of configuration definitions
- [X] [#4961](https://github.com/kubernetes/ingress-nginx/pull/4961) Remove prometheus socket before listen
- [X] [#4962](https://github.com/kubernetes/ingress-nginx/pull/4962) Cleanup of e2e docker images
- [X] [#4965](https://github.com/kubernetes/ingress-nginx/pull/4965) Move opentracing configuration for location to go
- [X] [#4966](https://github.com/kubernetes/ingress-nginx/pull/4966) Add verification of docker buildx support
- [X] [#4967](https://github.com/kubernetes/ingress-nginx/pull/4967) Update go dependencies

### 0.27.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.27.1`

Fix regression in Jaeger opentracing module, incorrect UID in webhook AdmissionResponse in Kubernetes > 1.16.0.

_Changes:_

- [X] [#4920](https://github.com/kubernetes/ingress-nginx/pull/4920) Rollback jaeger module version
- [X] [#4922](https://github.com/kubernetes/ingress-nginx/pull/4922) Use docker buildx and remove qemu-static
- [X] [#4927](https://github.com/kubernetes/ingress-nginx/pull/4927) Fix incorrect UID in webhook AdmissionResponse

### 0.27.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.27.0`

_New Features:_

- NGINX 1.17.7
- Migration to alpinelinux.
- Global [Modsecurity Snippet via ConfigMap](https://github.com/kubernetes/ingress-nginx/pull/4087)
- Support Datadog sample rate with global trace sampling from configmap [#4897](https://github.com/kubernetes/ingress-nginx/pull/4897)
- Modsecurity CRS v3.2.0 [#4829](https://github.com/kubernetes/ingress-nginx/pull/4829)
- Modsecurity-nginx v1.0.1 [#4842](https://github.com/kubernetes/ingress-nginx/pull/4842)
- Allow enabling/disabling opentracing for ingresses [#4732](https://github.com/kubernetes/ingress-nginx/pull/4732)

_Breaking Changes:_

- Enable download of GeoLite2 databases [#4896](https://github.com/kubernetes/ingress-nginx/pull/4896)

  _From maxmind website:_

  ```
  Due to upcoming data privacy regulations, we are making significant changes to how you access free GeoLite2 databases starting December 30, 2019.
  Learn more on our blog https://blog.maxmind.com/2019/12/18/significant-changes-to-accessing-and-using-geolite2-databases/
  ```

  Because of this change, it is not clear we can provide the databases directly from the docker image.
  To enable the feature, we provide two options:
  - Add the flag `--maxmind-license-key` to download the databases when the ingress controller starts.
  - or add a volume to mount the files `GeoLite2-City.mmdb` and `GeoLite2-ASN.mmdb` in the directory `/etc/nginx/geoip`.

  **If any of these conditions are not met, the geoip2 module will be disabled**

- The feature `lua-resty-waf` was removed.

- Due to the migration to alpinelinux the uid of the user is different. Please make sure to update it `runAsUser: 101` or the ingress controller will not start (CrashLoopBackOff).

_Changes:_

- [X] [#4087](https://github.com/kubernetes/ingress-nginx/pull/4087) Define Modsecurity Snippet via ConfigMap
- [X] [#4603](https://github.com/kubernetes/ingress-nginx/pull/4603) optimize: local cache global variable and reduce string object creation.
- [X] [#4613](https://github.com/kubernetes/ingress-nginx/pull/4613) Terraform release
- [X] [#4619](https://github.com/kubernetes/ingress-nginx/pull/4619) Issue 4244
- [X] [#4620](https://github.com/kubernetes/ingress-nginx/pull/4620) ISSUE-4244 e2e test
- [X] [#4645](https://github.com/kubernetes/ingress-nginx/pull/4645) Bind ingress controller to linux nodes to avoid Windows scheduling on kubernetes cluster includes linux nodes and windows nodes
- [X] [#4650](https://github.com/kubernetes/ingress-nginx/pull/4650) Expose GeoIP2 Organization as variable $geoip2_org
- [X] [#4658](https://github.com/kubernetes/ingress-nginx/pull/4658) Need to quote expansion of `$cfg.LogFormatStream` in `log_stream` access log
- [X] [#4664](https://github.com/kubernetes/ingress-nginx/pull/4664) warn when ConfigMap is missing or not parsable instead of erroring
- [X] [#4669](https://github.com/kubernetes/ingress-nginx/pull/4669) Simplify initialization function of bytes.Buffer
- [X] [#4671](https://github.com/kubernetes/ingress-nginx/pull/4671) Discontinue use of a single DNS query to validate an endpoint name
- [X] [#4673](https://github.com/kubernetes/ingress-nginx/pull/4673) More helpful dns error
- [X] [#4678](https://github.com/kubernetes/ingress-nginx/pull/4678) Increase the kubernetes 1.14 version to the installation prompt
- [X] [#4689](https://github.com/kubernetes/ingress-nginx/pull/4689) Server-only authentication of backends and per-location SSL config
- [X] [#4693](https://github.com/kubernetes/ingress-nginx/pull/4693) Adding some documentation about the use of metrics-per-host and enabl…
- [X] [#4694](https://github.com/kubernetes/ingress-nginx/pull/4694) Enhancement : add remote_addr in TCP access log
- [X] [#4695](https://github.com/kubernetes/ingress-nginx/pull/4695) Removing secure-verify-ca-secret support
- [X] [#4700](https://github.com/kubernetes/ingress-nginx/pull/4700) adds hability to use externalIP when controller service is of type NodePort
- [X] [#4730](https://github.com/kubernetes/ingress-nginx/pull/4730) add configuration for http2_max_concurrent_streams
- [X] [#4732](https://github.com/kubernetes/ingress-nginx/pull/4732) Allow enabling/disabling opentracing for ingresses
- [X] [#4745](https://github.com/kubernetes/ingress-nginx/pull/4745) add cmluciano to owners
- [X] [#4747](https://github.com/kubernetes/ingress-nginx/pull/4747) Docker image: Add source code reference label
- [X] [#4766](https://github.com/kubernetes/ingress-nginx/pull/4766) dev-env.sh: fix for parsing `minikube status` output of newer versions, fix shellcheck lints
- [X] [#4779](https://github.com/kubernetes/ingress-nginx/pull/4779) Remove lua-resty-waf feature
- [X] [#4780](https://github.com/kubernetes/ingress-nginx/pull/4780) Update nginx image to use openresty master
- [X] [#4785](https://github.com/kubernetes/ingress-nginx/pull/4785) Update nginx image and Go to 1.13.4
- [X] [#4791](https://github.com/kubernetes/ingress-nginx/pull/4791) deploy: add protocol to all Container/ServicePorts
- [X] [#4793](https://github.com/kubernetes/ingress-nginx/pull/4793)  Fix issue in logic of modsec template
- [X] [#4794](https://github.com/kubernetes/ingress-nginx/pull/4794) Remove extra annotation when Enabling ModSecurity
- [X] [#4797](https://github.com/kubernetes/ingress-nginx/pull/4797) Add a datasource variable $DS_PROMETHEUS
- [X] [#4803](https://github.com/kubernetes/ingress-nginx/pull/4803) Update nginx image to fix regression in jaeger tracing
- [X] [#4805](https://github.com/kubernetes/ingress-nginx/pull/4805) Update nginx and e2e images
- [X] [#4806](https://github.com/kubernetes/ingress-nginx/pull/4806) Add log to parallel command to dump logs in case of errors
- [X] [#4807](https://github.com/kubernetes/ingress-nginx/pull/4807) Allow custom CA certificate when flag --api-server is specified
- [X] [#4813](https://github.com/kubernetes/ingress-nginx/pull/4813) Update default SSL ciphers
- [X] [#4816](https://github.com/kubernetes/ingress-nginx/pull/4816) apply default certificate again in cases of invalid or incomplete cert config
- [X] [#4823](https://github.com/kubernetes/ingress-nginx/pull/4823) Update go dependencies to v1.17.0
- [X] [#4826](https://github.com/kubernetes/ingress-nginx/pull/4826) regression test and fix for duplicate hsts bug
- [X] [#4827](https://github.com/kubernetes/ingress-nginx/pull/4827) Migrate ingress definitions from extensions to networking.k8s.io
- [X] [#4829](https://github.com/kubernetes/ingress-nginx/pull/4829) Update modsecurity crs to v3.2.0
- [X] [#4840](https://github.com/kubernetes/ingress-nginx/pull/4840) Return specific type
- [X] [#4842](https://github.com/kubernetes/ingress-nginx/pull/4842) Update Modsecurity-nginx to latest (v1.0.1)
- [X] [#4843](https://github.com/kubernetes/ingress-nginx/pull/4843) Define minimum limits to run the ingress controller
- [X] [#4848](https://github.com/kubernetes/ingress-nginx/pull/4848) Update nginx image
- [X] [#4859](https://github.com/kubernetes/ingress-nginx/pull/4859) Use a named location for authSignURL
- [X] [#4862](https://github.com/kubernetes/ingress-nginx/pull/4862) Update nginx image
- [X] [#4863](https://github.com/kubernetes/ingress-nginx/pull/4863) Switch to nginx again
- [X] [#4866](https://github.com/kubernetes/ingress-nginx/pull/4866) Improve issue and pull request template
- [X] [#4867](https://github.com/kubernetes/ingress-nginx/pull/4867) Fix sticky session for ingress without host
- [X] [#4870](https://github.com/kubernetes/ingress-nginx/pull/4870) Default backend protocol only supports http
- [X] [#4871](https://github.com/kubernetes/ingress-nginx/pull/4871) Fix ingress status regression introduced in #4490
- [X] [#4875](https://github.com/kubernetes/ingress-nginx/pull/4875) Remove /build endpoint
- [X] [#4880](https://github.com/kubernetes/ingress-nginx/pull/4880) Remove download of geoip databases
- [X] [#4882](https://github.com/kubernetes/ingress-nginx/pull/4882) Use yaml files from a particular tag, not from master
- [X] [#4883](https://github.com/kubernetes/ingress-nginx/pull/4883) Update e2e image
- [X] [#4884](https://github.com/kubernetes/ingress-nginx/pull/4884) Update e2e image
- [X] [#4886](https://github.com/kubernetes/ingress-nginx/pull/4886) Fix flaking e2e tests
- [X] [#4887](https://github.com/kubernetes/ingress-nginx/pull/4887) Master branch uses a master tag image
- [X] [#4891](https://github.com/kubernetes/ingress-nginx/pull/4891) Add help task
- [X] [#4893](https://github.com/kubernetes/ingress-nginx/pull/4893) Use docker to run makefile tasks
- [X] [#4894](https://github.com/kubernetes/ingress-nginx/pull/4894) Remove todo from lua test
- [X] [#4896](https://github.com/kubernetes/ingress-nginx/pull/4896) Enable download of GeoLite2 databases
- [X] [#4897](https://github.com/kubernetes/ingress-nginx/pull/4897) Support Datadog sample rate with global trace sampling from configmap
- [X] [#4907](https://github.com/kubernetes/ingress-nginx/pull/4907) Add script to check go version and fix output directory permissions

_Documentation:_

- [X] [#4623](https://github.com/kubernetes/ingress-nginx/pull/4623) remove duplicated line in docs
- [X] [#4681](https://github.com/kubernetes/ingress-nginx/pull/4681) Fix docs/development.md describing inaccurate issues
- [X] [#4683](https://github.com/kubernetes/ingress-nginx/pull/4683) Fixed upgrading example command
- [X] [#4708](https://github.com/kubernetes/ingress-nginx/pull/4708) add proxy-max-temp-file-size doc
- [X] [#4727](https://github.com/kubernetes/ingress-nginx/pull/4727) update docs, remove output in prometheus deploy command
- [X] [#4744](https://github.com/kubernetes/ingress-nginx/pull/4744) Fix generation of sitemap.xml file
- [X] [#4746](https://github.com/kubernetes/ingress-nginx/pull/4746) Fix broken links in documentation
- [X] [#4748](https://github.com/kubernetes/ingress-nginx/pull/4748) Update documentation for static ip example
- [X] [#4749](https://github.com/kubernetes/ingress-nginx/pull/4749) Update documentation for rate limiting
- [X] [#4765](https://github.com/kubernetes/ingress-nginx/pull/4765) Fix extra word
- [X] [#4777](https://github.com/kubernetes/ingress-nginx/pull/4777) [docs] Add info about x-forwarded-prefix breaking change
- [X] [#4800](https://github.com/kubernetes/ingress-nginx/pull/4800) Update sysctl example
- [X] [#4801](https://github.com/kubernetes/ingress-nginx/pull/4801) Fix markdown list
- [X] [#4849](https://github.com/kubernetes/ingress-nginx/pull/4849) Fixed documentation for FCGI annotation.
- [X] [#4885](https://github.com/kubernetes/ingress-nginx/pull/4885) Correct MetalLB setup instructions.

### 0.26.2

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.2`

_Changes:_

- [X] [#4859](https://github.com/kubernetes/ingress-nginx/pull/4859) Use a named location for authSignURL

### 0.26.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.1`

_Changes:_

- [X] [#4617](https://github.com/kubernetes/ingress-nginx/pull/4617) Fix ports collision when hostNetwork=true
- [X] [#4619](https://github.com/kubernetes/ingress-nginx/pull/4619) Fix issue #4244

### 0.26.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.0`

_New Features:_

- Add support for NGINX [proxy_ssl_* directives](https://github.com/kubernetes/ingress-nginx/pull/4327)
- Add support for [FastCGI backends](https://github.com/kubernetes/ingress-nginx/pull/4344)
- [Only support SSL dynamic mode](https://github.com/kubernetes/ingress-nginx/pull/4356)
- [Add nginx ssl_early_data option support](https://github.com/kubernetes/ingress-nginx/pull/4412)
- [Add support for multiple alias and remove duplication of SSL certificates](https://github.com/kubernetes/ingress-nginx/pull/4472)
- [Support configuring basic auth credentials as a map of user/password hashes](https://github.com/kubernetes/ingress-nginx/pull/4560)
- Caching support for external authentication annotation with new annotations [auth-cache-key and auth-cache-duration](https://github.com/kubernetes/ingress-nginx/pull/4278)
- Allow Requests to be [Mirrored to different backends](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#mirror) [#4379](https://github.com/kubernetes/ingress-nginx/pull/4379)
- Improve connection draining when ingress controller pod is deleted using a lifecycle hook:

  With this new hook, we increased the default `terminationGracePeriodSeconds` from 30 seconds to 300, allowing the draining of connections up to five minutes.

  If the active connections end before that, the pod will terminate gracefully at that time.

  To efectively take advantage of this feature, the Configmap feature [worker-shutdown-timeout](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#worker-shutdown-timeout) new value is `240s` instead of `10s`.

  **IMPORTANT:** this value has a side effect during reloads, consuming more memory until the old NGINX workers are replaced.

  ```yaml
  lifecycle:
    preStop:
      exec:
        command:
          - /wait-shutdown
  ```

- [mimalloc](https://github.com/microsoft/mimalloc) as a drop-in replacement for malloc.

  This feature can be enabled using the [LD_PRELOAD](http://man7.org/linux/man-pages/man8/ld.so.8.html) environment variable in the ingress controller deployment

  Example:

  ```yaml
  env:
  - name: LD_PRELOAD
    value: /usr/local/lib/libmimalloc.so
  ```

  Please check the additional [options](https://github.com/microsoft/mimalloc#environment-options) it provides.

_Breaking Changes:_

- The variable [$the_real_ip variable](https://github.com/kubernetes/ingress-nginx/pull/4557) was removed from template and default `log_format`.
- The default value of configmap setting [proxy-add-original-uri-header](https://github.com/kubernetes/ingress-nginx/pull/4604) is now `"false"`.

  When the setting `proxy-add-original-uri-header` is `"true"`, the ingress controller adds a new header `X-Original-Uri` with the value of NGINX variable `$request_uri`.

  In most of the cases this is not an issue but with request with long URLs it could lead to unexpected errors in the application defined in the Ingress serviceName,
  like issue 4593 - [431 Request Header Fields Too Large](https://github.com/kubernetes/ingress-nginx/issues/4593)

_Non-functional improvements:_

- [Removal of internal NGINX unix sockets](https://github.com/kubernetes/ingress-nginx/pull/4531)
- Automation of NGINX image using [terraform scripts](https://github.com/kubernetes/ingress-nginx/pull/4484)
- Removal of Go profiling on port `:10254` to use `localhost:10245`

  To profile the ingress controller Go binary, use:

  ```console
  INGRESS_PODS=($(kubectl get pods -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx -o 'jsonpath={..metadata.name}'))
  kubectl port-forward -n ingress-nginx pod/${INGRESS_PODS[0]} 10245
  ```

Using the URL http://localhost:10245/debug/pprof/ to reach the profiler.

_Changes:_

- [X] [#3164](https://github.com/kubernetes/ingress-nginx/pull/3164) Initial support for CRL in Ingress Controller
- [X] [#4086](https://github.com/kubernetes/ingress-nginx/pull/4086) Resolve #4038, move X-Forwarded-Port variable to the location context
- [X] [#4278](https://github.com/kubernetes/ingress-nginx/pull/4278) feat: auth-req caching
- [X] [#4286](https://github.com/kubernetes/ingress-nginx/pull/4286) fix lua lints
- [X] [#4287](https://github.com/kubernetes/ingress-nginx/pull/4287) Add script for luacheck
- [X] [#4288](https://github.com/kubernetes/ingress-nginx/pull/4288) added proxy-http-version annotation to override the HTTP/1.1 default …
- [X] [#4289](https://github.com/kubernetes/ingress-nginx/pull/4289) Apply fixes suggested by staticcheck
- [X] [#4290](https://github.com/kubernetes/ingress-nginx/pull/4290) Make dev-env.sh script work on Linux
- [X] [#4291](https://github.com/kubernetes/ingress-nginx/pull/4291) hack scripts do not need PKG var
- [X] [#4298](https://github.com/kubernetes/ingress-nginx/pull/4298) Fix RBAC issues with networking.k8s.io
- [X] [#4299](https://github.com/kubernetes/ingress-nginx/pull/4299) Fix scripts to be able to run tests in docker
- [X] [#4302](https://github.com/kubernetes/ingress-nginx/pull/4302) Squash rules regarding ingresses
- [X] [#4306](https://github.com/kubernetes/ingress-nginx/pull/4306) Remove unnecessary output
- [X] [#4307](https://github.com/kubernetes/ingress-nginx/pull/4307) Disable access log in stream section for configuration socket
- [X] [#4313](https://github.com/kubernetes/ingress-nginx/pull/4313) avoid warning during lua unit test
- [X] [#4322](https://github.com/kubernetes/ingress-nginx/pull/4322) Update go dependencies
- [X] [#4327](https://github.com/kubernetes/ingress-nginx/pull/4327) Add proxy_ssl_* directives
- [X] [#4333](https://github.com/kubernetes/ingress-nginx/pull/4333) Add  [$proxy_alternative_upstream_name]
- [X] [#4334](https://github.com/kubernetes/ingress-nginx/pull/4334) Refactor http client for unix sockets
- [X] [#4341](https://github.com/kubernetes/ingress-nginx/pull/4341) duplicate argument "--disable-catch-all"
- [X] [#4344](https://github.com/kubernetes/ingress-nginx/pull/4344) Add FastCGI backend support (#2982)
- [X] [#4356](https://github.com/kubernetes/ingress-nginx/pull/4356) Only support SSL dynamic mode
- [X] [#4365](https://github.com/kubernetes/ingress-nginx/pull/4365) memoize balancer for a request
- [X] [#4369](https://github.com/kubernetes/ingress-nginx/pull/4369) Fix broken test's filenames
- [X] [#4371](https://github.com/kubernetes/ingress-nginx/pull/4371) Update datadog tracing plugin to v1.0.1
- [X] [#4379](https://github.com/kubernetes/ingress-nginx/pull/4379) Allow Requests to be Mirrored to different backends
- [X] [#4383](https://github.com/kubernetes/ingress-nginx/pull/4383) Add support for psp
- [X] [#4386](https://github.com/kubernetes/ingress-nginx/pull/4386) Update go dependencies
- [X] [#4405](https://github.com/kubernetes/ingress-nginx/pull/4405) Lua shared cfg
- [X] [#4409](https://github.com/kubernetes/ingress-nginx/pull/4409) sort ingress by namespace and name when ingress.CreationTimestamp identical
- [X] [#4410](https://github.com/kubernetes/ingress-nginx/pull/4410) fix dev-env script
- [X] [#4412](https://github.com/kubernetes/ingress-nginx/pull/4412) Add nginx ssl_early_data option support
- [X] [#4415](https://github.com/kubernetes/ingress-nginx/pull/4415) more dev-env script improvements
- [X] [#4416](https://github.com/kubernetes/ingress-nginx/pull/4416) Remove invalid log "Failed to executing diff command: exit status 1"
- [X] [#4418](https://github.com/kubernetes/ingress-nginx/pull/4418) Remove dynamic TLS records
- [X] [#4420](https://github.com/kubernetes/ingress-nginx/pull/4420) Cleanup
- [X] [#4422](https://github.com/kubernetes/ingress-nginx/pull/4422) teach lua about search and ndots settings in resolv.conf
- [X] [#4423](https://github.com/kubernetes/ingress-nginx/pull/4423) Add quote function in template
- [X] [#4426](https://github.com/kubernetes/ingress-nginx/pull/4426) Update klog
- [X] [#4428](https://github.com/kubernetes/ingress-nginx/pull/4428) Add timezone value into $geoip2_time_zone variable
- [X] [#4435](https://github.com/kubernetes/ingress-nginx/pull/4435) Add option to use existing images
- [X] [#4437](https://github.com/kubernetes/ingress-nginx/pull/4437) Refactor version helper
- [X] [#4438](https://github.com/kubernetes/ingress-nginx/pull/4438) Add helper to extract prometheus metrics in e2e tests
- [X] [#4439](https://github.com/kubernetes/ingress-nginx/pull/4439) Move listen logic to go
- [X] [#4440](https://github.com/kubernetes/ingress-nginx/pull/4440) Fixes for CVE-2018-16843, CVE-2018-16844, CVE-2019-9511, CVE-2019-9513, and CVE-2019-9516
- [X] [#4443](https://github.com/kubernetes/ingress-nginx/pull/4443) Lua resolv conf parser
- [X] [#4445](https://github.com/kubernetes/ingress-nginx/pull/4445) use latest openresty with CVE patches
- [X] [#4446](https://github.com/kubernetes/ingress-nginx/pull/4446) lua-shared-dicts improvements, fixes and documentation
- [X] [#4448](https://github.com/kubernetes/ingress-nginx/pull/4448) ewma improvements
- [X] [#4449](https://github.com/kubernetes/ingress-nginx/pull/4449) Fix service type external name using the name
- [X] [#4450](https://github.com/kubernetes/ingress-nginx/pull/4450) Add nginx proxy_max_temp_file_size configuration option
- [X] [#4451](https://github.com/kubernetes/ingress-nginx/pull/4451) post data to Lua only if it changes
- [X] [#4452](https://github.com/kubernetes/ingress-nginx/pull/4452) Fix test description on error
- [X] [#4456](https://github.com/kubernetes/ingress-nginx/pull/4456) Fix file permissions to support volumes
- [X] [#4458](https://github.com/kubernetes/ingress-nginx/pull/4458) implementation proposal for zone aware routing
- [X] [#4459](https://github.com/kubernetes/ingress-nginx/pull/4459) cleanup logging message typos in rewrite.go
- [X] [#4460](https://github.com/kubernetes/ingress-nginx/pull/4460) cleanup: fix typos in framework.go
- [X] [#4463](https://github.com/kubernetes/ingress-nginx/pull/4463) Always set headers with add-headers option
- [X] [#4466](https://github.com/kubernetes/ingress-nginx/pull/4466) Add rate limit units and error status
- [X] [#4471](https://github.com/kubernetes/ingress-nginx/pull/4471) Lint code using staticcheck
- [X] [#4472](https://github.com/kubernetes/ingress-nginx/pull/4472) Add support for multiple alias and remove duplication of SSL certificates
- [X] [#4476](https://github.com/kubernetes/ingress-nginx/pull/4476) Initialize nginx process error channel
- [X] [#4478](https://github.com/kubernetes/ingress-nginx/pull/4478) Re-add Support for Wildcard Hosts with Sticky Sessions
- [X] [#4484](https://github.com/kubernetes/ingress-nginx/pull/4484) Add terraform scripts to build nginx image
- [X] [#4487](https://github.com/kubernetes/ingress-nginx/pull/4487) Refactor health checks and wait until NGINX process ends
- [X] [#4489](https://github.com/kubernetes/ingress-nginx/pull/4489) Fix log format markdown
- [X] [#4490](https://github.com/kubernetes/ingress-nginx/pull/4490) Refactor ingress status IP address
- [X] [#4492](https://github.com/kubernetes/ingress-nginx/pull/4492) fix lua certificate handling tests
- [X] [#4495](https://github.com/kubernetes/ingress-nginx/pull/4495) point users to kubectl ingress-nginx plugin
- [X] [#4500](https://github.com/kubernetes/ingress-nginx/pull/4500) Fix nginx variable service_port (nginx)
- [X] [#4501](https://github.com/kubernetes/ingress-nginx/pull/4501) Move nginx helper
- [X] [#4502](https://github.com/kubernetes/ingress-nginx/pull/4502) Remove hard-coded names from e2e test and use local docker dependencies
- [X] [#4506](https://github.com/kubernetes/ingress-nginx/pull/4506) Fix panic on multiple ingress mess up upstream is primary or not
- [X] [#4509](https://github.com/kubernetes/ingress-nginx/pull/4509) Update openresty and third party modules
- [X] [#4520](https://github.com/kubernetes/ingress-nginx/pull/4520) fix typo
- [X] [#4521](https://github.com/kubernetes/ingress-nginx/pull/4521) backward compatibility for k8s version < 1.14
- [X] [#4522](https://github.com/kubernetes/ingress-nginx/pull/4522) Fix relative links
- [X] [#4524](https://github.com/kubernetes/ingress-nginx/pull/4524) Update go dependencies
- [X] [#4527](https://github.com/kubernetes/ingress-nginx/pull/4527) Switch to official kind images
- [X] [#4528](https://github.com/kubernetes/ingress-nginx/pull/4528) Cleanup of docker images
- [X] [#4530](https://github.com/kubernetes/ingress-nginx/pull/4530) Update nginx image to 0.92
- [X] [#4531](https://github.com/kubernetes/ingress-nginx/pull/4531) Remove nginx unix sockets
- [X] [#4534](https://github.com/kubernetes/ingress-nginx/pull/4534) Show current reloads count, not total
- [X] [#4535](https://github.com/kubernetes/ingress-nginx/pull/4535) Improve the time to run e2e tests
- [X] [#4543](https://github.com/kubernetes/ingress-nginx/pull/4543) Correctly format ipv6 resolver config for lua
- [X] [#4545](https://github.com/kubernetes/ingress-nginx/pull/4545) Rollback luarocks version to 3.1.3
- [X] [#4547](https://github.com/kubernetes/ingress-nginx/pull/4547) Fix terraform build of nginx images
- [X] [#4548](https://github.com/kubernetes/ingress-nginx/pull/4548) regression test for the issue fixed in #4543
- [X] [#4549](https://github.com/kubernetes/ingress-nginx/pull/4549) Cleanup of docker build
- [X] [#4556](https://github.com/kubernetes/ingress-nginx/pull/4556) Allow multiple CA Certificates
- [X] [#4557](https://github.com/kubernetes/ingress-nginx/pull/4557) Remove the_real_ip variable
- [X] [#4560](https://github.com/kubernetes/ingress-nginx/pull/4560) Support configuring basic auth credentials as a map of user/password hashes
- [X] [#4569](https://github.com/kubernetes/ingress-nginx/pull/4569) allow to configure jaeger header names
- [X] [#4570](https://github.com/kubernetes/ingress-nginx/pull/4570) Update nginx image
- [X] [#4571](https://github.com/kubernetes/ingress-nginx/pull/4571) Increase log level for identical CreationTimestamp warning
- [X] [#4572](https://github.com/kubernetes/ingress-nginx/pull/4572) Fix log format after #4557
- [X] [#4575](https://github.com/kubernetes/ingress-nginx/pull/4575) Update go dependencies for kubernetes 1.16.0
- [X] [#4583](https://github.com/kubernetes/ingress-nginx/pull/4583) Disable go modules
- [X] [#4584](https://github.com/kubernetes/ingress-nginx/pull/4584) Remove retries to ExternalName
- [X] [#4586](https://github.com/kubernetes/ingress-nginx/pull/4586) Fix reload when a configmap changes
- [X] [#4587](https://github.com/kubernetes/ingress-nginx/pull/4587) Avoid unnecessary reloads generating lua_shared_dict directives
- [X] [#4591](https://github.com/kubernetes/ingress-nginx/pull/4591) optimize: local cache global variable and avoid single lines over 80
- [X] [#4592](https://github.com/kubernetes/ingress-nginx/pull/4592) refactor force ssl redirect logic
- [X] [#4594](https://github.com/kubernetes/ingress-nginx/pull/4594) cleanup unused certificates
- [X] [#4595](https://github.com/kubernetes/ingress-nginx/pull/4595) Rollback change of ModSecurity setting SecAuditLog
- [X] [#4596](https://github.com/kubernetes/ingress-nginx/pull/4596) sort auth proxy headers from configmap
- [X] [#4597](https://github.com/kubernetes/ingress-nginx/pull/4597) more meaningful assertion for tls hsts test
- [X] [#4598](https://github.com/kubernetes/ingress-nginx/pull/4598) delete redundant config
- [X] [#4600](https://github.com/kubernetes/ingress-nginx/pull/4600) Update nginx image
- [X] [#4601](https://github.com/kubernetes/ingress-nginx/pull/4601) Hsts refactoring
- [X] [#4602](https://github.com/kubernetes/ingress-nginx/pull/4602) fix bug with new and running configuration comparison
- [X] [#4604](https://github.com/kubernetes/ingress-nginx/pull/4604) Change default for proxy-add-original-uri-header
- [X] [#4606](https://github.com/kubernetes/ingress-nginx/pull/4606) Mount temporal directory volume for ingress controller
- [X] [#4611](https://github.com/kubernetes/ingress-nginx/pull/4611) Fix custom default backend switch to default

_Documentation:_

- [X] [#4277](https://github.com/kubernetes/ingress-nginx/pull/4277) doc: fix image link.
- [X] [#4316](https://github.com/kubernetes/ingress-nginx/pull/4316) Update how-it-works.md
- [X] [#4329](https://github.com/kubernetes/ingress-nginx/pull/4329) Update references to oauth2_proxy
- [X] [#4348](https://github.com/kubernetes/ingress-nginx/pull/4348) KEP process
- [X] [#4351](https://github.com/kubernetes/ingress-nginx/pull/4351) KEP: Remove static SSL configuration mode
- [X] [#4389](https://github.com/kubernetes/ingress-nginx/pull/4389) Fix docs build due to an invalid link
- [X] [#4455](https://github.com/kubernetes/ingress-nginx/pull/4455) KEP: availability zone aware routing
- [X] [#4581](https://github.com/kubernetes/ingress-nginx/pull/4581) Fix spelling and remove local reference of 404 docker image
- [X] [#4582](https://github.com/kubernetes/ingress-nginx/pull/4582) Update kubectl-plugin docs
- [X] [#4588](https://github.com/kubernetes/ingress-nginx/pull/4588) tls user guide --default-ssl-certificate clarification

### 0.25.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.25.1`

_Changes:_

- [X] [#4440](https://github.com/kubernetes/ingress-nginx/pull/4440) Fixes for CVE-2018-16843, CVE-2018-16844, CVE-2019-9511, CVE-2019-9513, and CVE-2019-9516

### 0.25.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.25.0`

_New Features:_

- Validating webhook for ingress sanity check [documentation](https://kubernetes.github.io/ingress-nginx/deploy/validating-webhook/)
- Migration from NGINX to [OpenResty](https://openresty.org/en/) 1.15.8
- [ARM image](https://quay.io/repository/kubernetes-ingress-controller/nginx-ingress-controller-arm?tab=logs)
- Improve external authorization concept from opt-in to secure-by-default [3506](https://github.com/kubernetes/ingress-nginx/pull/3506)
- Reduce memory footprint and cpu usage when modsecurity is enabled [4091](https://github.com/kubernetes/ingress-nginx/pull/4091)
- Support new `networking.k8s.io/v1beta1` package (for Kubernetes cluster > v1.14.0) [4127](https://github.com/kubernetes/ingress-nginx/pull/4127)
- New variable `$proxy_alternative_upstream_name` in the log to show a hit in a canary endpoint [#4246](https://github.com/kubernetes/ingress-nginx/pull/4246)

_Non-functional improvements:_

- Migration from travis-ci to [Prow](https://prow.k8s.io/tide-history?repo=kubernetes%2Fingress-nginx&branch=master)
- [Testgrid dashboards](https://testgrid.k8s.io/sig-network-ingress-nginx#Summary) for ingress-nginx
- Update kind to [v0.4.0](https://github.com/kubernetes-sigs/kind/releases/tag/v0.4.0)
- Switch to go modules
- Go v1.12.6
- Docker size image reduced by 20%

_Changes:_

- [X] [#3506](https://github.com/kubernetes/ingress-nginx/pull/3506) Improve the external authorization concept from opt-in to secure-by-default
- [X] [#3802](https://github.com/kubernetes/ingress-nginx/pull/3802) Add a validating webhook for ingress sanity check
- [X] [#3803](https://github.com/kubernetes/ingress-nginx/pull/3803) use nkeys for counting lua table elements
- [X] [#3852](https://github.com/kubernetes/ingress-nginx/pull/3852) Enable arm again
- [X] [#4004](https://github.com/kubernetes/ingress-nginx/pull/4004) Remove valgrind
- [X] [#4005](https://github.com/kubernetes/ingress-nginx/pull/4005) Support proxy_next_upstream_timeout
- [X] [#4008](https://github.com/kubernetes/ingress-nginx/pull/4008) refactor GetFakeSSLCert
- [X] [#4009](https://github.com/kubernetes/ingress-nginx/pull/4009) Update nginx to 1.15.12
- [X] [#4010](https://github.com/kubernetes/ingress-nginx/pull/4010) Update nginx image and Go to 1.12.4
- [X] [#4012](https://github.com/kubernetes/ingress-nginx/pull/4012) Switch to go modules
- [X] [#4022](https://github.com/kubernetes/ingress-nginx/pull/4022) Add e2e test coverage for mult-auth
- [X] [#4042](https://github.com/kubernetes/ingress-nginx/pull/4042) Release custom error pages image v0.4 [skip-ci]
- [X] [#4048](https://github.com/kubernetes/ingress-nginx/pull/4048) Change upstream on error when sticky session balancer is used
- [X] [#4055](https://github.com/kubernetes/ingress-nginx/pull/4055) Rearrange deployment files into kustomizations
- [X] [#4064](https://github.com/kubernetes/ingress-nginx/pull/4064) Update go to 1.12.5, kubectl to 1.14.1 and kind to 0.2.1
- [X] [#4067](https://github.com/kubernetes/ingress-nginx/pull/4067) Trim spaces from annotations that can contain multiple lines
- [X] [#4069](https://github.com/kubernetes/ingress-nginx/pull/4069) fix e2e-test make target
- [X] [#4070](https://github.com/kubernetes/ingress-nginx/pull/4070) Don't try to create e2e runner rbac resources twice
- [X] [#4080](https://github.com/kubernetes/ingress-nginx/pull/4080) Load modsecurity config with OWASP core rules
- [X] [#4088](https://github.com/kubernetes/ingress-nginx/pull/4088) Migrate to Prow
- [X] [#4091](https://github.com/kubernetes/ingress-nginx/pull/4091) reduce memory footprint and cpu usage when modsecurity and owasp rule
- [X] [#4100](https://github.com/kubernetes/ingress-nginx/pull/4100) Remove stop controller endpoint
- [X] [#4101](https://github.com/kubernetes/ingress-nginx/pull/4101) Refactor whitelist from map to standard allow directives
- [X] [#4102](https://github.com/kubernetes/ingress-nginx/pull/4102) Refactor ListIngresses to add filters
- [X] [#4105](https://github.com/kubernetes/ingress-nginx/pull/4105) UPT: Add variable to define custom sampler host and port
- [X] [#4108](https://github.com/kubernetes/ingress-nginx/pull/4108) Add retry to LookupHost used to check the content of ExternalName
- [X] [#4109](https://github.com/kubernetes/ingress-nginx/pull/4109) Use real apiserver
- [X] [#4110](https://github.com/kubernetes/ingress-nginx/pull/4110) Update e2e images
- [X] [#4113](https://github.com/kubernetes/ingress-nginx/pull/4113) Force GOOS to linux
- [X] [#4119](https://github.com/kubernetes/ingress-nginx/pull/4119) Only load module ngx_http_modsecurity_module.so when option enable-mo…
- [X] [#4120](https://github.com/kubernetes/ingress-nginx/pull/4120) log info when endpoints change for a balancer
- [X] [#4122](https://github.com/kubernetes/ingress-nginx/pull/4122) Update Nginx to 1.17.0 and upgrade some other modules
- [X] [#4123](https://github.com/kubernetes/ingress-nginx/pull/4123) Update nginx image to 0.86
- [X] [#4127](https://github.com/kubernetes/ingress-nginx/pull/4127) Migrate to new networking.k8s.io/v1beta1 package
- [X] [#4128](https://github.com/kubernetes/ingress-nginx/pull/4128) feature(collectors): Added services to collectorLabels
- [X] [#4133](https://github.com/kubernetes/ingress-nginx/pull/4133) Run PodSecurityPolicy E2E test in parallel
- [X] [#4135](https://github.com/kubernetes/ingress-nginx/pull/4135) Use apps/v1 api group in e2e tests
- [X] [#4140](https://github.com/kubernetes/ingress-nginx/pull/4140) update modsecurity to latest, libmodsecurity to v3.0.3 and owasp-scrs…
- [X] [#4150](https://github.com/kubernetes/ingress-nginx/pull/4150) Update nginx
- [X] [#4160](https://github.com/kubernetes/ingress-nginx/pull/4160) SSL expiration metrics cannot be tied to dynamic updates
- [X] [#4162](https://github.com/kubernetes/ingress-nginx/pull/4162) Add "text/javascript" to compressible MIME types
- [X] [#4164](https://github.com/kubernetes/ingress-nginx/pull/4164) fix source file mods
- [X] [#4166](https://github.com/kubernetes/ingress-nginx/pull/4166) Session Affinity ChangeOnFailure should be boolean
- [X] [#4169](https://github.com/kubernetes/ingress-nginx/pull/4169) simplify sticky balancer and fix a bug
- [X] [#4180](https://github.com/kubernetes/ingress-nginx/pull/4180) Service type=ExternalName can be defined with ports
- [X] [#4185](https://github.com/kubernetes/ingress-nginx/pull/4185) Fix: fillout missing health check timeout on health check.
- [X] [#4187](https://github.com/kubernetes/ingress-nginx/pull/4187) Add unit test cases for balancer lua module
- [X] [#4191](https://github.com/kubernetes/ingress-nginx/pull/4191) increase lua_shared_dict config data
- [X] [#4204](https://github.com/kubernetes/ingress-nginx/pull/4204) Add e2e test for service type=ExternalName
- [X] [#4212](https://github.com/kubernetes/ingress-nginx/pull/4212) Add e2e tests for grpc
- [X] [#4214](https://github.com/kubernetes/ingress-nginx/pull/4214) Update go dependencies
- [X] [#4219](https://github.com/kubernetes/ingress-nginx/pull/4219) Get AuthTLS annotation unit tests to 100%
- [X] [#4220](https://github.com/kubernetes/ingress-nginx/pull/4220) Migrate to openresty
- [X] [#4221](https://github.com/kubernetes/ingress-nginx/pull/4221) Switch to openresty image
- [X] [#4223](https://github.com/kubernetes/ingress-nginx/pull/4223) Remove travis-ci badge
- [X] [#4224](https://github.com/kubernetes/ingress-nginx/pull/4224) fix monitor test after move to openresty
- [X] [#4225](https://github.com/kubernetes/ingress-nginx/pull/4225) Update image dependencies
- [X] [#4226](https://github.com/kubernetes/ingress-nginx/pull/4226) Update nginx image
- [X] [#4227](https://github.com/kubernetes/ingress-nginx/pull/4227) Fix misspelled and e2e check
- [X] [#4229](https://github.com/kubernetes/ingress-nginx/pull/4229) Do not send empty certificates to nginx
- [X] [#4232](https://github.com/kubernetes/ingress-nginx/pull/4232) override least recently used entries when certificate_data dict is full
- [X] [#4233](https://github.com/kubernetes/ingress-nginx/pull/4233) Update nginx image to 0.90
- [X] [#4235](https://github.com/kubernetes/ingress-nginx/pull/4235) Add new lints
- [X] [#4236](https://github.com/kubernetes/ingress-nginx/pull/4236) Add e2e test suite to detect memory leaks in lua
- [X] [#4237](https://github.com/kubernetes/ingress-nginx/pull/4237) Update go dependencies
- [X] [#4246](https://github.com/kubernetes/ingress-nginx/pull/4246) introduce proxy_alternative_upstream_name Nginx var
- [X] [#4249](https://github.com/kubernetes/ingress-nginx/pull/4249) test to make sure dynamic cert works trailing dot in domains
- [X] [#4250](https://github.com/kubernetes/ingress-nginx/pull/4250) Lint shell scripts
- [X] [#4251](https://github.com/kubernetes/ingress-nginx/pull/4251) Refactor prometheus leader helper
- [X] [#4253](https://github.com/kubernetes/ingress-nginx/pull/4253) Remove kubeclient configuration
- [X] [#4254](https://github.com/kubernetes/ingress-nginx/pull/4254) Update kind to 0.4.0
- [X] [#4257](https://github.com/kubernetes/ingress-nginx/pull/4257) Fix error deleting temporal directory in case of errors
- [X] [#4258](https://github.com/kubernetes/ingress-nginx/pull/4258) Fix go imports
- [X] [#4267](https://github.com/kubernetes/ingress-nginx/pull/4267) More e2e tests
- [X] [#4270](https://github.com/kubernetes/ingress-nginx/pull/4270) GetLbAlgorithm helper func for e2e
- [X] [#4272](https://github.com/kubernetes/ingress-nginx/pull/4272) introduce ngx.var.balancer_ewma_score
- [X] [#4273](https://github.com/kubernetes/ingress-nginx/pull/4273) Check and complete intermediate SSL certificates
- [X] [#4274](https://github.com/kubernetes/ingress-nginx/pull/4274) Support trailing dot

_Documentation:_

- [X] [#3966](https://github.com/kubernetes/ingress-nginx/pull/3966) Documentation example code fix
- [X] [#3978](https://github.com/kubernetes/ingress-nginx/pull/3978) Fix CA certificate example docs
- [X] [#3981](https://github.com/kubernetes/ingress-nginx/pull/3981) Add missing PR in changelog [skip ci]
- [X] [#3982](https://github.com/kubernetes/ingress-nginx/pull/3982) Add kubectl plugin docs
- [X] [#3987](https://github.com/kubernetes/ingress-nginx/pull/3987) Link to kubectl plugin docs in nav
- [X] [#4014](https://github.com/kubernetes/ingress-nginx/pull/4014) Update plugin krew manifest
- [X] [#4034](https://github.com/kubernetes/ingress-nginx/pull/4034) 🔧 fix navigation error in file baremetal.md
- [X] [#4036](https://github.com/kubernetes/ingress-nginx/pull/4036) Docs have incorrect command in baremetal.md
- [X] [#4037](https://github.com/kubernetes/ingress-nginx/pull/4037) [doc] fixing regex in example of rewrite
- [X] [#4040](https://github.com/kubernetes/ingress-nginx/pull/4040) Fix default Content-Type for custom-error-pages example
- [X] [#4068](https://github.com/kubernetes/ingress-nginx/pull/4068) fix typo: deployement->deployment
- [X] [#4082](https://github.com/kubernetes/ingress-nginx/pull/4082) Explain references in custom-headers documentation
- [X] [#4089](https://github.com/kubernetes/ingress-nginx/pull/4089) Docs: configmap: use-gzip
- [X] [#4099](https://github.com/kubernetes/ingress-nginx/pull/4099) Docs - Update capture group `placeholder`
- [X] [#4098](https://github.com/kubernetes/ingress-nginx/pull/4098) Update configmap about adding custom locations
- [X] [#4107](https://github.com/kubernetes/ingress-nginx/pull/4107) Clear up some inconsistent / unclear wording
- [X] [#4132](https://github.com/kubernetes/ingress-nginx/pull/4132) Update README.md for external-auth Test 4
- [X] [#4153](https://github.com/kubernetes/ingress-nginx/pull/4153) Add clarification on how to enable path matching
- [X] [#4159](https://github.com/kubernetes/ingress-nginx/pull/4159) Partially revert usage of kustomize for installation
- [X] [#4217](https://github.com/kubernetes/ingress-nginx/pull/4217) Fix typo in annotations
- [X] [#4228](https://github.com/kubernetes/ingress-nginx/pull/4228) Add notes on timeouts while using long GRPC streams

### 0.24.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.24.1`

_Changes:_

- [X] [#3990](https://github.com/kubernetes/ingress-nginx/pull/3990) Fix dynamic cert issue with default-ssl-certificate
- [X] [#3980](https://github.com/kubernetes/ingress-nginx/pull/3980) Refactor isIterable
- [X] [#4000](https://github.com/kubernetes/ingress-nginx/pull/4000) Dynamic ssl improvements
- [X] [#4007](https://github.com/kubernetes/ingress-nginx/pull/4007) do not create empty access_by_lua_block

### 0.24.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.24.0`

_New Features:_

- NGINX 1.15.10

_Breaking changes:_

- `x-forwarded-prefix` annotation changed from a boolean to a string, see [#3786](https://github.com/kubernetes/ingress-nginx/pull/3786)

_Changes:_

- [X] [#3743](https://github.com/kubernetes/ingress-nginx/pull/3743) Remove session-cookie-hash annotation
- [X] [#3786](https://github.com/kubernetes/ingress-nginx/pull/3786) Fix x-forwarded-prefix annotation
- [X] [#3798](https://github.com/kubernetes/ingress-nginx/pull/3798) Move some configuration logic from Nginx config to Lua code
- [X] [#3806](https://github.com/kubernetes/ingress-nginx/pull/3806) Migrate e2e cluster to kind
- [X] [#3807](https://github.com/kubernetes/ingress-nginx/pull/3807) Lua plugin system - MVP
- [X] [#3808](https://github.com/kubernetes/ingress-nginx/pull/3808) make dynamic SSL mode default
- [X] [#3827](https://github.com/kubernetes/ingress-nginx/pull/3827) Fix plugin install location
- [X] [#3829](https://github.com/kubernetes/ingress-nginx/pull/3829) Prevent e2e-tests from running on non-local clusters
- [X] [#3833](https://github.com/kubernetes/ingress-nginx/pull/3833) bump luajit version to v2.1-20190228
- [X] [#3835](https://github.com/kubernetes/ingress-nginx/pull/3835) Update nginx image
- [X] [#3839](https://github.com/kubernetes/ingress-nginx/pull/3839) Fix panic on multiple non-matching canary
- [X] [#3846](https://github.com/kubernetes/ingress-nginx/pull/3846) Fix race condition in metric process collector test
- [X] [#3849](https://github.com/kubernetes/ingress-nginx/pull/3849) Use Gauge instead of Counter for connections_active Prometheus metric
- [X] [#3853](https://github.com/kubernetes/ingress-nginx/pull/3853) Remove authbind
- [X] [#3856](https://github.com/kubernetes/ingress-nginx/pull/3856) Fix ssl-dh-param issue when secret does not exit
- [X] [#3864](https://github.com/kubernetes/ingress-nginx/pull/3864) ing.Service with multiple hosts fix
- [X] [#3870](https://github.com/kubernetes/ingress-nginx/pull/3870) Improve kubectl plugin
- [X] [#3871](https://github.com/kubernetes/ingress-nginx/pull/3871) Fix name of field used to sort ingresses [skip-ci]
- [X] [#3875](https://github.com/kubernetes/ingress-nginx/pull/3875) Allow the use of a secret located in a different namespace
- [X] [#3882](https://github.com/kubernetes/ingress-nginx/pull/3882) Add support for IPV6 resolvers
- [X] [#3884](https://github.com/kubernetes/ingress-nginx/pull/3884) update GKE header to match link in contents
- [X] [#3885](https://github.com/kubernetes/ingress-nginx/pull/3885) Refactor status update
- [X] [#3886](https://github.com/kubernetes/ingress-nginx/pull/3886) Clean up ssl package and fix dynamic cert mode
- [X] [#3887](https://github.com/kubernetes/ingress-nginx/pull/3887) Remove useless nodeip calls and deprecate --force-namespace-isolation
- [X] [#3889](https://github.com/kubernetes/ingress-nginx/pull/3889) Separate out annotation assignment logic
- [X] [#3895](https://github.com/kubernetes/ingress-nginx/pull/3895) Correctly format ipv6 resolver config for lua
- [X] [#3900](https://github.com/kubernetes/ingress-nginx/pull/3900) Add lint subcommand to plugin
- [X] [#3907](https://github.com/kubernetes/ingress-nginx/pull/3907) Remove unnecessary copy of GeoIP databases
- [X] [#3908](https://github.com/kubernetes/ingress-nginx/pull/3908) Update nginx image
- [X] [#3918](https://github.com/kubernetes/ingress-nginx/pull/3918) Set `X-Request-ID` for the `default-backend`, too.
- [X] [#3927](https://github.com/kubernetes/ingress-nginx/pull/3927) Update apiVersion to apps/v1, drop duplicate line
- [X] [#3932](https://github.com/kubernetes/ingress-nginx/pull/3932) Fix dynamic SSL certificate for aliases and redirect-from-to-www
- [X] [#3933](https://github.com/kubernetes/ingress-nginx/pull/3933) Update nginx to 1.15.10
- [X] [#3934](https://github.com/kubernetes/ingress-nginx/pull/3934) Update nginx image
- [X] [#3943](https://github.com/kubernetes/ingress-nginx/pull/3943) Update dependencies
- [X] [#3947](https://github.com/kubernetes/ingress-nginx/pull/3947) Adds a log warning when falling back to default fake cert
- [X] [#3950](https://github.com/kubernetes/ingress-nginx/pull/3950) Fix forwarded host parsing
- [X] [#3954](https://github.com/kubernetes/ingress-nginx/pull/3954) Fix load-balance configmap value
- [X] [#3955](https://github.com/kubernetes/ingress-nginx/pull/3955) Plugin select deployment using replicaset name
- [X] [#3958](https://github.com/kubernetes/ingress-nginx/pull/3958) Refactor equals
- [X] [#3960](https://github.com/kubernetes/ingress-nginx/pull/3960) Fix segfault on reference to nonexistent configmap
- [X] [#3968](https://github.com/kubernetes/ingress-nginx/pull/3968) Update nginx image
- [X] [#3969](https://github.com/kubernetes/ingress-nginx/pull/3969) Update nginx image to 0.84

_Documentation:_

- [X] [#3841](https://github.com/kubernetes/ingress-nginx/pull/3841) Improve "Sticky session" docs
- [X] [#3836](https://github.com/kubernetes/ingress-nginx/pull/3836) Update mkdocs [skip ci]
- [X] [#3847](https://github.com/kubernetes/ingress-nginx/pull/3847) Add missing basic usage documentation link
- [X] [#3874](https://github.com/kubernetes/ingress-nginx/pull/3874) Update embargo doc link in SECURITY_CONTACTS and change PST to PSC
- [X] [#3890](https://github.com/kubernetes/ingress-nginx/pull/3890) Make sure cli-arguments doc is in alphabetical order
- [X] [#3945](https://github.com/kubernetes/ingress-nginx/pull/3945) fix typo: delete '`'

### 0.23.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.23.0`

_New Features:_

- NGINX 1.15.9
- New `canary-by-header-value` [annotation](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md#canary).
- New debug binary to get runtime information from lua [3686](https://github.com/kubernetes/ingress-nginx/pull/3686)
- Support for Opentracing with Datadog
- New [kubectl plugin](https://github.com/kubernetes/ingress-nginx/pull/3779) **Alpha**

_Breaking changes:_

- The NGINX server listening in port 18080 was removed. It was replaced by a server using an unix socket as port [#3684](https://github.com/kubernetes/ingress-nginx/pull/3684)
  This server was internal to the ingress controller. In case this was being acceded from the outside, you can restore the old server using the `http-snipet` feature in the configuration configmap like:

  ```yaml
  http-snippet: |
    server {
      listen 18080;

      location /nginx_status {
        allow 127.0.0.1;
        allow ::1;
        deny all;
        stub_status on;
      }

      location / {
        return 404;
      }
    }
  ```

_Changes:_

- [X] [#3619](https://github.com/kubernetes/ingress-nginx/pull/3619) add header-value annotation
- [X] [#3628](https://github.com/kubernetes/ingress-nginx/pull/3628) Fix 503 error generation on empty endpoints
- [X] [#3666](https://github.com/kubernetes/ingress-nginx/pull/3666) rename sysctlFSFileMax to rlimitMaxNumFiles to reflect what it actually does
- [X] [#3667](https://github.com/kubernetes/ingress-nginx/pull/3667) worker_connections should be less (3/4th) than worker_rlimit_nofile
- [X] [#3671](https://github.com/kubernetes/ingress-nginx/pull/3671) bugfix: fixed duplicated seeds.
- [X] [#3673](https://github.com/kubernetes/ingress-nginx/pull/3673) used table functions of LuaJIT for better performance.
- [X] [#3674](https://github.com/kubernetes/ingress-nginx/pull/3674) used cjson.safe instead of pcall.
- [X] [#3682](https://github.com/kubernetes/ingress-nginx/pull/3682) enable use-forwarded-headers for L7 LB
- [X] [#3684](https://github.com/kubernetes/ingress-nginx/pull/3684) Replace Status port using a socket
- [X] [#3686](https://github.com/kubernetes/ingress-nginx/pull/3686) Add debug binary to the docker image
- [X] [#3695](https://github.com/kubernetes/ingress-nginx/pull/3695) > Don't reload nginx when L4 endpoints changed
- [X] [#3696](https://github.com/kubernetes/ingress-nginx/pull/3696) Apply annotations to default location
- [X] [#3698](https://github.com/kubernetes/ingress-nginx/pull/3698) Fix --disable-catch-all
- [X] [#3702](https://github.com/kubernetes/ingress-nginx/pull/3702) Add params for access log
- [X] [#3704](https://github.com/kubernetes/ingress-nginx/pull/3704) make sure dev-env forces context to be minikube
- [X] [#3728](https://github.com/kubernetes/ingress-nginx/pull/3728) Fix flaky test
- [X] [#3730](https://github.com/kubernetes/ingress-nginx/pull/3730) Changes CustomHTTPErrors annotation to use custom default backend
- [X] [#3734](https://github.com/kubernetes/ingress-nginx/pull/3734) remove old unused lua dicts
- [X] [#3736](https://github.com/kubernetes/ingress-nginx/pull/3736) do not unnecessarily log
- [X] [#3737](https://github.com/kubernetes/ingress-nginx/pull/3737) Adjust probe timeouts
- [X] [#3739](https://github.com/kubernetes/ingress-nginx/pull/3739) dont log unnecessarily
- [X] [#3740](https://github.com/kubernetes/ingress-nginx/pull/3740) Fix ingress updating for session-cookie-* annotation changes
- [X] [#3747](https://github.com/kubernetes/ingress-nginx/pull/3747) Update nginx and modules
- [X] [#3748](https://github.com/kubernetes/ingress-nginx/pull/3748) Update nginx image
- [X] [#3749](https://github.com/kubernetes/ingress-nginx/pull/3749) Enhance Unit Tests for Annotations
- [X] [#3750](https://github.com/kubernetes/ingress-nginx/pull/3750) Update go dependencies
- [X] [#3751](https://github.com/kubernetes/ingress-nginx/pull/3751) Parse environment variables in OpenTracing configuration
- [X] [#3756](https://github.com/kubernetes/ingress-nginx/pull/3756) Create custom annotation for satisfy "value"
- [X] [#3757](https://github.com/kubernetes/ingress-nginx/pull/3757) Add mention of secure-backends to backend-protocol docs
- [X] [#3764](https://github.com/kubernetes/ingress-nginx/pull/3764) delete confusing CustomErrors attribute to make things more explicit
- [X] [#3765](https://github.com/kubernetes/ingress-nginx/pull/3765) simplify customhttperrors e2e test and add regression test and fix a bug
- [X] [#3766](https://github.com/kubernetes/ingress-nginx/pull/3766) Support Opentracing with Datadog - part 2
- [X] [#3767](https://github.com/kubernetes/ingress-nginx/pull/3767) Support Opentracing with Datadog - part 1
- [X] [#3771](https://github.com/kubernetes/ingress-nginx/pull/3771) Do not log unnecessarily
- [X] [#3772](https://github.com/kubernetes/ingress-nginx/pull/3772) Fix dashboard link [skip ci]
- [X] [#3775](https://github.com/kubernetes/ingress-nginx/pull/3775) Fix DNS lookup failures in L4 services
- [X] [#3779](https://github.com/kubernetes/ingress-nginx/pull/3779) Add kubectl plugin
- [X] [#3780](https://github.com/kubernetes/ingress-nginx/pull/3780) Enable access log for default backend
- [X] [#3781](https://github.com/kubernetes/ingress-nginx/pull/3781) feat: configurable proxy buffers number
- [X] [#3782](https://github.com/kubernetes/ingress-nginx/pull/3782) Lua bridge tracer
- [X] [#3784](https://github.com/kubernetes/ingress-nginx/pull/3784) use correct host for jaeger-collector-host in docs
- [X] [#3785](https://github.com/kubernetes/ingress-nginx/pull/3785) use latest base nginx image
- [X] [#3787](https://github.com/kubernetes/ingress-nginx/pull/3787) Use UsePortInRedirects only if enabled
- [X] [#3791](https://github.com/kubernetes/ingress-nginx/pull/3791) - remove annotations in nginxcontroller struct
- [X] [#3792](https://github.com/kubernetes/ingress-nginx/pull/3792) dont restart minikube when it is already running
- [X] [#3793](https://github.com/kubernetes/ingress-nginx/pull/3793) Update mergo dependency
- [X] [#3794](https://github.com/kubernetes/ingress-nginx/pull/3794) use use-context that actually changes the context
- [X] [#3795](https://github.com/kubernetes/ingress-nginx/pull/3795) do not warn when optional annotations arent set
- [X] [#3799](https://github.com/kubernetes/ingress-nginx/pull/3799) Add /dbg certs command
- [X] [#3800](https://github.com/kubernetes/ingress-nginx/pull/3800) Refactor e2e
- [X] [#3809](https://github.com/kubernetes/ingress-nginx/pull/3809) Upgrade openresty/lua-resty-balancer
- [X] [#3810](https://github.com/kubernetes/ingress-nginx/pull/3810) Update nginx image
- [X] [#3811](https://github.com/kubernetes/ingress-nginx/pull/3811) Fix e2e tests
- [X] [#3812](https://github.com/kubernetes/ingress-nginx/pull/3812) Removes unused const from customhttperrors e2e test
- [X] [#3813](https://github.com/kubernetes/ingress-nginx/pull/3813) Prevent dep from vendoring grpc-fortune-teller dependencies
- [X] [#3819](https://github.com/kubernetes/ingress-nginx/pull/3819) Fix e2e test in osx
- [X] [#3820](https://github.com/kubernetes/ingress-nginx/pull/3820) Update nginx image
- [X] [#3821](https://github.com/kubernetes/ingress-nginx/pull/3821) Update nginx to 1.15.9
- [X] [#3822](https://github.com/kubernetes/ingress-nginx/pull/3822) Set default for satisfy annotation to nothing

_Documentation:_

- [X] [#3680](https://github.com/kubernetes/ingress-nginx/pull/3680) mention rewrite-target change for 0.22.0
- [X] [#3693](https://github.com/kubernetes/ingress-nginx/pull/3693) Correcting links for gRPC Fortune Teller app
- [X] [#3701](https://github.com/kubernetes/ingress-nginx/pull/3701) Update usage documentation for default-backend annotation
- [X] [#3705](https://github.com/kubernetes/ingress-nginx/pull/3705) Increase Unit Test Coverage for Templates
- [X] [#3708](https://github.com/kubernetes/ingress-nginx/pull/3708) Update OWNERS
- [X] [#3731](https://github.com/kubernetes/ingress-nginx/pull/3731) Update a doc example that uses rewrite-target

_Deprecations:_

- The annotation `session-cookie-hash` is deprecated and will be removed in 0.24.
- Flag `--force-namespace-isolation` is deprecated and will be removed in 0.24. Currently this annotation is being replaced by `--watch-namespace`

### 0.22.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.22.0`

_New Features:_

- NGINX 1.15.8
- New balancer implementation: consistent hash subset
- Adds support for HTTP2 Push Preload annotation
- Allow to disable NGINX prometheus metrics
- New --disable-catch-all flag to ignore catch-all ingresses
- Add flag --metrics-per-host to make per-host metrics optional

_Breaking changes:_

- Annotation `nginx.ingress.kubernetes.io/rewrite-target` has changed and will not behave as expected if you don't update them.

  Refer to [https://kubernetes.github.io/ingress-nginx/examples/rewrite/#rewrite-target](https://kubernetes.github.io/ingress-nginx/examples/rewrite/#rewrite-target) on how to change it.

  Refer to [https://github.com/kubernetes/ingress-nginx/pull/3174#issuecomment-455665710](https://github.com/kubernetes/ingress-nginx/pull/3174#issuecomment-455665710) on how to do seamless migration.

- Annotations `nginx.ingress.kubernetes.io/add-base-url` and `nginx.ingress.kubernetes.io/base-url-scheme` were removed.

  Please check issue [#3174](https://github.com/kubernetes/ingress-nginx/pull/3174) for details.

- By default do not trust any client to extract true client IP address from X-Forwarded-For header using realip module (`use-forwarded-headers: "false"`)

_Changes:_

- [X] [#3174](https://github.com/kubernetes/ingress-nginx/pull/3174) Generalize Rewrite Block Creation and Deprecate AddBaseUrl (not backwards compatible)
- [X] [#3240](https://github.com/kubernetes/ingress-nginx/pull/3240) Adds support for HTTP2 Push Preload annotation
- [X] [#3333](https://github.com/kubernetes/ingress-nginx/pull/3333) breaking change: by default do not trust any client
- [X] [#3342](https://github.com/kubernetes/ingress-nginx/pull/3342) Allow privilege escalation
- [X] [#3363](https://github.com/kubernetes/ingress-nginx/pull/3363) Document for cookie expires annotation
- [X] [#3396](https://github.com/kubernetes/ingress-nginx/pull/3396) New balancer implementation: consistent hash subset
- [X] [#3446](https://github.com/kubernetes/ingress-nginx/pull/3446) add more testing for mergeAlternativeBackends
- [X] [#3453](https://github.com/kubernetes/ingress-nginx/pull/3453) Monitor fixes
- [X] [#3455](https://github.com/kubernetes/ingress-nginx/pull/3455) Watch controller Pods and make then available in k8sStore
- [X] [#3465](https://github.com/kubernetes/ingress-nginx/pull/3465) Bump nginx-opentracing for gRPC support
- [X] [#3467](https://github.com/kubernetes/ingress-nginx/pull/3467) store ewma stats per backend
- [X] [#3470](https://github.com/kubernetes/ingress-nginx/pull/3470) Use opentracing_grpc_propagate_context when necessary
- [X] [#3474](https://github.com/kubernetes/ingress-nginx/pull/3474) Improve parsing of annotations and use of Ingress wrapper
- [X] [#3476](https://github.com/kubernetes/ingress-nginx/pull/3476) Fix nginx directory permissions
- [X] [#3477](https://github.com/kubernetes/ingress-nginx/pull/3477) clarify canary ingress
- [X] [#3478](https://github.com/kubernetes/ingress-nginx/pull/3478) delete unused buildLoadBalancingConfig
- [X] [#3487](https://github.com/kubernetes/ingress-nginx/pull/3487) dynamic certificate mode should support widlcard hosts
- [X] [#3488](https://github.com/kubernetes/ingress-nginx/pull/3488) Add probes to deployments used in e2e tests
- [X] [#3492](https://github.com/kubernetes/ingress-nginx/pull/3492) Fix data size validations
- [X] [#3494](https://github.com/kubernetes/ingress-nginx/pull/3494) Since dynamic mode only checking for 'return 503' is not valid anymore
- [X] [#3495](https://github.com/kubernetes/ingress-nginx/pull/3495) Adjust default timeout for e2e tests
- [X] [#3497](https://github.com/kubernetes/ingress-nginx/pull/3497) Wait for the right number of endpoints
- [X] [#3498](https://github.com/kubernetes/ingress-nginx/pull/3498) Update godeps
- [X] [#3501](https://github.com/kubernetes/ingress-nginx/pull/3501) be consistent with what Nginx supports
- [X] [#3503](https://github.com/kubernetes/ingress-nginx/pull/3503) compare error with error types from k8s.io/apimachinery/pkg/api/errors
- [X] [#3504](https://github.com/kubernetes/ingress-nginx/pull/3504) fix an ewma unit test
- [X] [#3505](https://github.com/kubernetes/ingress-nginx/pull/3505) Update lua configuration_data when number of controller pod change
- [X] [#3507](https://github.com/kubernetes/ingress-nginx/pull/3507) Remove temporal configuration file after a while
- [X] [#3508](https://github.com/kubernetes/ingress-nginx/pull/3508) Update nginx to 1.15.7
- [X] [#3509](https://github.com/kubernetes/ingress-nginx/pull/3509) [1759] Ingress affinity session cookie with Secure flag for HTTPS
- [X] [#3512](https://github.com/kubernetes/ingress-nginx/pull/3512) Allow to disable NGINX metrics
- [X] [#3518](https://github.com/kubernetes/ingress-nginx/pull/3518) Fix log output format
- [X] [#3521](https://github.com/kubernetes/ingress-nginx/pull/3521) Fix a bug with Canary becoming main server
- [X] [#3522](https://github.com/kubernetes/ingress-nginx/pull/3522) {tcp,udp}-services cm appear twice
- [X] [#3525](https://github.com/kubernetes/ingress-nginx/pull/3525) make canary ingresses independent of the order they were applied
- [X] [#3530](https://github.com/kubernetes/ingress-nginx/pull/3530) Update nginx image
- [X] [#3532](https://github.com/kubernetes/ingress-nginx/pull/3532) Ignore updates of ingresses with invalid class
- [X] [#3536](https://github.com/kubernetes/ingress-nginx/pull/3536) Replace dockerfile entrypoint
- [X] [#3548](https://github.com/kubernetes/ingress-nginx/pull/3548) e2e test to ensure graceful shutdown does not lose requests
- [X] [#3551](https://github.com/kubernetes/ingress-nginx/pull/3551) Fix --enable-dynamic-certificates for nested subdomain
- [X] [#3553](https://github.com/kubernetes/ingress-nginx/pull/3553) handle_error_when_executing_diff
- [X] [#3562](https://github.com/kubernetes/ingress-nginx/pull/3562) Rename nginx.yaml to nginx.json
- [X] [#3566](https://github.com/kubernetes/ingress-nginx/pull/3566) Add Unit Tests for getIngressInformation
- [X] [#3569](https://github.com/kubernetes/ingress-nginx/pull/3569) fix status updated: make sure ingress.status is copied
- [X] [#3573](https://github.com/kubernetes/ingress-nginx/pull/3573) Update Certificate Generation Docs to not use MD5
- [X] [#3581](https://github.com/kubernetes/ingress-nginx/pull/3581) lua randomseed per worker
- [X] [#3582](https://github.com/kubernetes/ingress-nginx/pull/3582) Sort ingresses by creation timestamp
- [X] [#3584](https://github.com/kubernetes/ingress-nginx/pull/3584) Update go to 1.11.4
- [X] [#3586](https://github.com/kubernetes/ingress-nginx/pull/3586) Add --disable-catch-all option to disable catch-all server
- [X] [#3587](https://github.com/kubernetes/ingress-nginx/pull/3587) adjust dind istallation
- [X] [#3594](https://github.com/kubernetes/ingress-nginx/pull/3594) Add a flag to make per-host metrics optional
- [X] [#3596](https://github.com/kubernetes/ingress-nginx/pull/3596) Fix proxy_host variable configuration
- [X] [#3601](https://github.com/kubernetes/ingress-nginx/pull/3601) Update nginx to 1.15.8
- [X] [#3602](https://github.com/kubernetes/ingress-nginx/pull/3602) Update nginx image
- [X] [#3604](https://github.com/kubernetes/ingress-nginx/pull/3604) Add an option to automatically set worker_connections based on worker_rlimit_nofile
- [X] [#3615](https://github.com/kubernetes/ingress-nginx/pull/3615) Pass k8s `Service` data through to the TCP balancer script.
- [X] [#3620](https://github.com/kubernetes/ingress-nginx/pull/3620) Added server alias to metrics
- [X] [#3624](https://github.com/kubernetes/ingress-nginx/pull/3624) Update nginx to fix geoip database deprecation
- [X] [#3625](https://github.com/kubernetes/ingress-nginx/pull/3625) Update nginx image
- [X] [#3633](https://github.com/kubernetes/ingress-nginx/pull/3633) Fix a bug in Ingress update handler
- [X] [#3634](https://github.com/kubernetes/ingress-nginx/pull/3634) canary by cookie should support hypen in cookie name
- [X] [#3635](https://github.com/kubernetes/ingress-nginx/pull/3635) Fix duplicate alternative backend merging
- [X] [#3637](https://github.com/kubernetes/ingress-nginx/pull/3637) Add support for redirect https to https (from-to-www-redirect)
- [X] [#3640](https://github.com/kubernetes/ingress-nginx/pull/3640) add limit connection status code
- [X] [#3641](https://github.com/kubernetes/ingress-nginx/pull/3641) Replace deprecated apiVersion in deploy folder
- [X] [#3643](https://github.com/kubernetes/ingress-nginx/pull/3643) Update nginx
- [X] [#3644](https://github.com/kubernetes/ingress-nginx/pull/3644) Update nginx image
- [X] [#3648](https://github.com/kubernetes/ingress-nginx/pull/3648) Remove stickyness cookie domain from Lua balancer to match old behavior
- [X] [#3649](https://github.com/kubernetes/ingress-nginx/pull/3649) Empty access_by_lua_block breaks satisfy any
- [X] [#3655](https://github.com/kubernetes/ingress-nginx/pull/3655) Remove flag sort-backends
- [X] [#3656](https://github.com/kubernetes/ingress-nginx/pull/3656) Change default value of  flag for ssl chain completion
- [X] [#3660](https://github.com/kubernetes/ingress-nginx/pull/3660) Revert max-worker-connections default value
- [X] [#3664](https://github.com/kubernetes/ingress-nginx/pull/3664) Fix invalid validation creating prometheus valid host values

_Documentation:_

- [X] [#3513](https://github.com/kubernetes/ingress-nginx/pull/3513) Revert removal of TCP and UDP support configmaps in mandatroy manifest
- [X] [#3456](https://github.com/kubernetes/ingress-nginx/pull/3456) Revert TCP/UDP documentation removal and links
- [X] [#3482](https://github.com/kubernetes/ingress-nginx/pull/3482) Annotations doc links: minor fixes and unification
- [X] [#3491](https://github.com/kubernetes/ingress-nginx/pull/3491) Update example to use latest Dashboard version.
- [X] [#3510](https://github.com/kubernetes/ingress-nginx/pull/3510) Update mkdocs [skip ci]
- [X] [#3516](https://github.com/kubernetes/ingress-nginx/pull/3516) Fix error in configmap yaml definition
- [X] [#3575](https://github.com/kubernetes/ingress-nginx/pull/3575) Add documentation for spec.rules.host format
- [X] [#3577](https://github.com/kubernetes/ingress-nginx/pull/3577) Add standard labels to namespace specs
- [X] [#3592](https://github.com/kubernetes/ingress-nginx/pull/3592) Add inside the User Guide documentation section a basic usage section and example
- [X] [#3605](https://github.com/kubernetes/ingress-nginx/pull/3605) Fix CLA URLs
- [X] [#3627](https://github.com/kubernetes/ingress-nginx/pull/3627) Typo: docs/examples/rewrite/README.md
- [X] [#3632](https://github.com/kubernetes/ingress-nginx/pull/3632) Fixed: error parsing with-rbac.yaml: error converting YAML to JSON

### 0.21.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.21.0`

_New Features:_

- NGINX 1.15.6 with fixes for vulnerabilities in HTTP/2 ([CVE-2018-16843](http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16843), [CVE-2018-16844](http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16844))
- Support for TLSv1.3. Disabled by default. Use [ssl-protocols](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#ssl-protocols) `ssl-protocols: TLSv1.3 TLSv1.2`
- New annotation for [canary deployments](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#canary)
- Support for configuration snippets when the authentication annotation is used
- Support for custom ModSecurity configuration
- LUA upstream configuration for TCP and UDP services

_Changes:_

- [X] [#3156](https://github.com/kubernetes/ingress-nginx/pull/3156) [404-server] Removes 404 server
- [X] [#3170](https://github.com/kubernetes/ingress-nginx/pull/3170) Move mainSnippet before events to fix load_module issue.
- [X] [#3187](https://github.com/kubernetes/ingress-nginx/pull/3187) UPT: annotation enhancement for resty-lua-waf
- [X] [#3190](https://github.com/kubernetes/ingress-nginx/pull/3190) Refactor e2e Tests to use common helper
- [X] [#3193](https://github.com/kubernetes/ingress-nginx/pull/3193) Add E2E tests for HealthCheck
- [X] [#3194](https://github.com/kubernetes/ingress-nginx/pull/3194) Make literal $ character work in set $location_path
- [X] [#3195](https://github.com/kubernetes/ingress-nginx/pull/3195) Add e2e Tests for AuthTLS
- [X] [#3196](https://github.com/kubernetes/ingress-nginx/pull/3196) Remove default backend requirement
- [X] [#3197](https://github.com/kubernetes/ingress-nginx/pull/3197) Remove support for TCP and UDP services
- [X] [#3198](https://github.com/kubernetes/ingress-nginx/pull/3198) Only support dynamic configuration
- [X] [#3199](https://github.com/kubernetes/ingress-nginx/pull/3199) Remove duplication in files
- [X] [#3201](https://github.com/kubernetes/ingress-nginx/pull/3201) no data shows for config reloads charts when select to namespace or controller
- [X] [#3203](https://github.com/kubernetes/ingress-nginx/pull/3203) Remove annotations grpc-backend and secure-backend already deprecated
- [X] [#3204](https://github.com/kubernetes/ingress-nginx/pull/3204) Flags publish-service and publish-status-address are mutually exclusive
- [X] [#3205](https://github.com/kubernetes/ingress-nginx/pull/3205) Update OWNERS [skip ci]
- [X] [#3207](https://github.com/kubernetes/ingress-nginx/pull/3207) delete upstream healthcheck annotation
- [X] [#3209](https://github.com/kubernetes/ingress-nginx/pull/3209) Fix: update config map name
- [X] [#3212](https://github.com/kubernetes/ingress-nginx/pull/3212) Add some extra detail to the client cert auth example regarding potential gotcha
- [X] [#3213](https://github.com/kubernetes/ingress-nginx/pull/3213) Update deps
- [X] [#3214](https://github.com/kubernetes/ingress-nginx/pull/3214) Cleanup of nginx image
- [X] [#3219](https://github.com/kubernetes/ingress-nginx/pull/3219) Update nginx image
- [X] [#3222](https://github.com/kubernetes/ingress-nginx/pull/3222) Allow Ability to Configure Upstream Keepalive
- [X] [#3230](https://github.com/kubernetes/ingress-nginx/pull/3230) Retry initial backend configuration
- [X] [#3231](https://github.com/kubernetes/ingress-nginx/pull/3231) Improve dynamic lua configuration
- [X] [#3234](https://github.com/kubernetes/ingress-nginx/pull/3234) Added e2e tests for backend protocols
- [X] [#3247](https://github.com/kubernetes/ingress-nginx/pull/3247) Refactor probe url requests
- [X] [#3252](https://github.com/kubernetes/ingress-nginx/pull/3252) remove the command args of enable-dynamic-configuration
- [X] [#3257](https://github.com/kubernetes/ingress-nginx/pull/3257) Add e2e tests for upstream vhost
- [X] [#3260](https://github.com/kubernetes/ingress-nginx/pull/3260) fix logging calls
- [X] [#3261](https://github.com/kubernetes/ingress-nginx/pull/3261) Mount minikube volume to docker container
- [X] [#3265](https://github.com/kubernetes/ingress-nginx/pull/3265) Update kubeadm-dind-cluster
- [X] [#3266](https://github.com/kubernetes/ingress-nginx/pull/3266) fix two bugs with backend-protocol annotation
- [X] [#3267](https://github.com/kubernetes/ingress-nginx/pull/3267) Fix status update in case of connection errors
- [X] [#3270](https://github.com/kubernetes/ingress-nginx/pull/3270) Don't sort IngressStatus from each Goroutine(update for each ingress)
- [X] [#3277](https://github.com/kubernetes/ingress-nginx/pull/3277) Add e2e test for configuration snippet
- [X] [#3279](https://github.com/kubernetes/ingress-nginx/pull/3279) Fix usages of %q formatting for numbers (%d)
- [X] [#3280](https://github.com/kubernetes/ingress-nginx/pull/3280) Add e2e test for from-to-www-redirect
- [X] [#3281](https://github.com/kubernetes/ingress-nginx/pull/3281) Add e2e test for log
- [X] [#3285](https://github.com/kubernetes/ingress-nginx/pull/3285) Add health-check-timeout as command line argument
- [X] [#3286](https://github.com/kubernetes/ingress-nginx/pull/3286) fix bug with balancer.lua configuration
- [X] [#3295](https://github.com/kubernetes/ingress-nginx/pull/3295) Refactor EWMA to not use shared dictionaries
- [X] [#3296](https://github.com/kubernetes/ingress-nginx/pull/3296) Update nginx and add support for TLSv1.3
- [X] [#3297](https://github.com/kubernetes/ingress-nginx/pull/3297) Add e2e test for force-ssl-redirect
- [X] [#3301](https://github.com/kubernetes/ingress-nginx/pull/3301) Add e2e tests for IP Whitelist
- [X] [#3302](https://github.com/kubernetes/ingress-nginx/pull/3302) Add e2e test for server snippet
- [X] [#3304](https://github.com/kubernetes/ingress-nginx/pull/3304) Update kubeadm-dind-cluster script
- [X] [#3305](https://github.com/kubernetes/ingress-nginx/pull/3305) Add e2e test for app-root
- [X] [#3306](https://github.com/kubernetes/ingress-nginx/pull/3306) Update e2e test to verify redirect code
- [X] [#3309](https://github.com/kubernetes/ingress-nginx/pull/3309) Customize ModSecurity to be used in Locations
- [X] [#3310](https://github.com/kubernetes/ingress-nginx/pull/3310) Fix geoip2 db files
- [X] [#3313](https://github.com/kubernetes/ingress-nginx/pull/3313) Support cookie expires
- [X] [#3320](https://github.com/kubernetes/ingress-nginx/pull/3320) Update nginx image and QEMU version
- [X] [#3321](https://github.com/kubernetes/ingress-nginx/pull/3321) Add configuration for geoip2 module
- [X] [#3322](https://github.com/kubernetes/ingress-nginx/pull/3322) Remove e2e boilerplate
- [X] [#3324](https://github.com/kubernetes/ingress-nginx/pull/3324) Fix sticky session
- [X] [#3325](https://github.com/kubernetes/ingress-nginx/pull/3325) Fix e2e tests
- [X] [#3328](https://github.com/kubernetes/ingress-nginx/pull/3328) Code linting
- [X] [#3332](https://github.com/kubernetes/ingress-nginx/pull/3332) Update build-single-manifest-sh,remove tcp-services-configmap.yaml and udp-services-configmap.yaml
- [X] [#3338](https://github.com/kubernetes/ingress-nginx/pull/3338) Avoid reloads when endpoints are not available
- [X] [#3341](https://github.com/kubernetes/ingress-nginx/pull/3341) Add canary annotation and alternative backends for traffic shaping
- [X] [#3343](https://github.com/kubernetes/ingress-nginx/pull/3343) Auth snippet
- [X] [#3344](https://github.com/kubernetes/ingress-nginx/pull/3344) Adds CustomHTTPErrors ingress annotation and test
- [X] [#3345](https://github.com/kubernetes/ingress-nginx/pull/3345) update annotation
- [X] [#3346](https://github.com/kubernetes/ingress-nginx/pull/3346) Add e2e test for session-cookie-hash
- [X] [#3347](https://github.com/kubernetes/ingress-nginx/pull/3347) Add e2e test for ssl-redirect
- [X] [#3348](https://github.com/kubernetes/ingress-nginx/pull/3348) Update cli-arguments.md. Remove tcp and udp, add health-check-timeout.
- [X] [#3353](https://github.com/kubernetes/ingress-nginx/pull/3353) Update nginx modules
- [X] [#3354](https://github.com/kubernetes/ingress-nginx/pull/3354) Update nginx image
- [X] [#3356](https://github.com/kubernetes/ingress-nginx/pull/3356) Download latest dep releases instead of fetching from HEAD
- [X] [#3357](https://github.com/kubernetes/ingress-nginx/pull/3357) Add missing modsecurity unicode.mapping file
- [X] [#3367](https://github.com/kubernetes/ingress-nginx/pull/3367) Remove reloads when there is no endpoints
- [X] [#3372](https://github.com/kubernetes/ingress-nginx/pull/3372) Add annotation for session affinity path
- [X] [#3373](https://github.com/kubernetes/ingress-nginx/pull/3373) Update nginx
- [X] [#3374](https://github.com/kubernetes/ingress-nginx/pull/3374) Revert removal of support for TCP and UDP services
- [X] [#3383](https://github.com/kubernetes/ingress-nginx/pull/3383) Only set cookies on paths that enable session affinity
- [X] [#3387](https://github.com/kubernetes/ingress-nginx/pull/3387) Modify the wrong function name
- [X] [#3390](https://github.com/kubernetes/ingress-nginx/pull/3390) Add e2e test for round robin load balancing
- [X] [#3400](https://github.com/kubernetes/ingress-nginx/pull/3400) Add Snippet for ModSecurity
- [X] [#3404](https://github.com/kubernetes/ingress-nginx/pull/3404) Update nginx image
- [X] [#3405](https://github.com/kubernetes/ingress-nginx/pull/3405) Prevent X-Forwarded-Proto forward during external auth subrequest
- [X] [#3406](https://github.com/kubernetes/ingress-nginx/pull/3406) Update nginx and e2e image
- [X] [#3407](https://github.com/kubernetes/ingress-nginx/pull/3407) Restructure load balance e2e tests and update round robin test
- [X] [#3408](https://github.com/kubernetes/ingress-nginx/pull/3408) Fix modsecurity configuration file location
- [X] [#3409](https://github.com/kubernetes/ingress-nginx/pull/3409) Convert isValidClientBodyBufferSize to something more generic
- [X] [#3410](https://github.com/kubernetes/ingress-nginx/pull/3410) fix logging calls
- [X] [#3415](https://github.com/kubernetes/ingress-nginx/pull/3415) bugfix: set canary attributes when initializing balancer
- [X] [#3417](https://github.com/kubernetes/ingress-nginx/pull/3417) bugfix: do not merge catch-all canary backends with itself
- [X] [#3421](https://github.com/kubernetes/ingress-nginx/pull/3421) Fix X-Forwarded-Proto typo
- [X] [#3424](https://github.com/kubernetes/ingress-nginx/pull/3424) Update nginx image
- [X] [#3425](https://github.com/kubernetes/ingress-nginx/pull/3425) Update nginx modules
- [X] [#3428](https://github.com/kubernetes/ingress-nginx/pull/3428) Set proxy_host variable to avoid using default value from proxy_pass
- [X] [#3437](https://github.com/kubernetes/ingress-nginx/pull/3437) Use struct to pack Ingress and its annotations
- [X] [#3441](https://github.com/kubernetes/ingress-nginx/pull/3441) Match buffer
- [X] [#3442](https://github.com/kubernetes/ingress-nginx/pull/3442) Increase log level when there is an invalid size value
- [X] [#3453](https://github.com/kubernetes/ingress-nginx/pull/3453) Monitor fixes

_Documentation:_

- [X] [#3166](https://github.com/kubernetes/ingress-nginx/pull/3166) Added ingress tls values.yaml example to documentation
- [X] [#3215](https://github.com/kubernetes/ingress-nginx/pull/3215) align opentracing user-guide with nginx configmap configuration
- [X] [#3229](https://github.com/kubernetes/ingress-nginx/pull/3229) Fix documentation links [skip ci]
- [X] [#3232](https://github.com/kubernetes/ingress-nginx/pull/3232) Fix typo
- [X] [#3242](https://github.com/kubernetes/ingress-nginx/pull/3242) Add a note to the deployment into GKE
- [X] [#3249](https://github.com/kubernetes/ingress-nginx/pull/3249) Clarify mandatory script doc
- [X] [#3262](https://github.com/kubernetes/ingress-nginx/pull/3262) Add e2e test for connection
- [X] [#3263](https://github.com/kubernetes/ingress-nginx/pull/3263) "diretly" typo
- [X] [#3264](https://github.com/kubernetes/ingress-nginx/pull/3264) Add missing annotations to Docs
- [X] [#3271](https://github.com/kubernetes/ingress-nginx/pull/3271) the sample ingress spec error
- [X] [#3275](https://github.com/kubernetes/ingress-nginx/pull/3275) Add Better Documentation for using AuthTLS
- [X] [#3282](https://github.com/kubernetes/ingress-nginx/pull/3282) Fix some typos
- [X] [#3312](https://github.com/kubernetes/ingress-nginx/pull/3312) Delete some extra words
- [X] [#3319](https://github.com/kubernetes/ingress-nginx/pull/3319) Fix links in deploy index docs
- [X] [#3326](https://github.com/kubernetes/ingress-nginx/pull/3326) fix broken link
- [X] [#3349](https://github.com/kubernetes/ingress-nginx/pull/3349) fix typo
- [X] [#3364](https://github.com/kubernetes/ingress-nginx/pull/3364) Fix links format [skip-ci]
- [X] [#3366](https://github.com/kubernetes/ingress-nginx/pull/3366) Fix some typos
- [X] [#3369](https://github.com/kubernetes/ingress-nginx/pull/3369) Fix some typos
- [X] [#3370](https://github.com/kubernetes/ingress-nginx/pull/3370) Fix typo: whitlelist -> whitelist
- [X] [#3377](https://github.com/kubernetes/ingress-nginx/pull/3377) Fix typos and default value
- [X] [#3379](https://github.com/kubernetes/ingress-nginx/pull/3379) Fix typos
- [X] [#3382](https://github.com/kubernetes/ingress-nginx/pull/3382) Fix typos: reqrite -> rewrite
- [X] [#3388](https://github.com/kubernetes/ingress-nginx/pull/3388) Update annotations.md. Remove Duplication.
- [X] [#3392](https://github.com/kubernetes/ingress-nginx/pull/3392) Fix link in documentation [skip ci]
- [X] [#3395](https://github.com/kubernetes/ingress-nginx/pull/3395) Fix some documents issues

### 0.20.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.20.0`

_New Features:_

- NGINX 1.15.5
- Support for *regular expressions* in paths https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/ingress-path-matching.md
- Provide possibility to block IPs, User-Agents and Referers globally
- Remove --default-backend-service requirement. Use the flag only for custom default backends
- Valgrind and Openresty gdb tools

_Changes:_

- [X] [#2997](https://github.com/kubernetes/ingress-nginx/pull/2997) Provide possibility to block IPs, User-Agents and Referers globally
- [X] [#3016](https://github.com/kubernetes/ingress-nginx/pull/3016) Log Errors Missing in Internal
- [X] [#3017](https://github.com/kubernetes/ingress-nginx/pull/3017) Add e2e tests for CORS
- [X] [#3022](https://github.com/kubernetes/ingress-nginx/pull/3022) Add support for valgrind
- [X] [#3029](https://github.com/kubernetes/ingress-nginx/pull/3029) add support for http2-max-requests in configmap
- [X] [#3035](https://github.com/kubernetes/ingress-nginx/pull/3035) Fixup #2970: Add Missing Label `app.kubernetes.io/part-of: ingress-nginx`
- [X] [#3049](https://github.com/kubernetes/ingress-nginx/pull/3049) fix: Don't try and find local certs when secretName is not specified
- [X] [#3050](https://github.com/kubernetes/ingress-nginx/pull/3050) Add Ingress variable in Grafana dashboard
- [X] [#3062](https://github.com/kubernetes/ingress-nginx/pull/3062) Pass Host header for custom errors
- [X] [#3065](https://github.com/kubernetes/ingress-nginx/pull/3065) Join host/port with go helper (supports ipv6)
- [X] [#3067](https://github.com/kubernetes/ingress-nginx/pull/3067) fix missing datasource value
- [X] [#3069](https://github.com/kubernetes/ingress-nginx/pull/3069) Replace client-go deprecated method
- [X] [#3072](https://github.com/kubernetes/ingress-nginx/pull/3072) Update ingress service IP
- [X] [#3073](https://github.com/kubernetes/ingress-nginx/pull/3073) do not hardcode the path
- [X] [#3078](https://github.com/kubernetes/ingress-nginx/pull/3078) Fix Rewrite-Target Annotation Edge Case
- [X] [#3079](https://github.com/kubernetes/ingress-nginx/pull/3079) Openresty gdb tools
- [X] [#3080](https://github.com/kubernetes/ingress-nginx/pull/3080) Update nginx image to 0.62
- [X] [#3098](https://github.com/kubernetes/ingress-nginx/pull/3098) make upstream keepalive work for http
- [X] [#3100](https://github.com/kubernetes/ingress-nginx/pull/3100) update annotation name from rewrite-log to enable-rewrite-log
- [X] [#3118](https://github.com/kubernetes/ingress-nginx/pull/3118) Replace standard json encoding with jsoniter
- [X] [#3121](https://github.com/kubernetes/ingress-nginx/pull/3121) Typo fix: adresses -> addresses
- [X] [#3126](https://github.com/kubernetes/ingress-nginx/pull/3126) do not require --default-backend-service
- [X] [#3130](https://github.com/kubernetes/ingress-nginx/pull/3130) fix newlines location denied
- [X] [#3133](https://github.com/kubernetes/ingress-nginx/pull/3133) multi-tls readme example to reference the file
- [X] [#3134](https://github.com/kubernetes/ingress-nginx/pull/3134) Update nginx to 1.15.4
- [X] [#3135](https://github.com/kubernetes/ingress-nginx/pull/3135) Remove payload from post log
- [X] [#3136](https://github.com/kubernetes/ingress-nginx/pull/3136) Update nginx image
- [X] [#3137](https://github.com/kubernetes/ingress-nginx/pull/3137) Docker run as user
- [X] [#3143](https://github.com/kubernetes/ingress-nginx/pull/3143) Ensure monitoring for custom error pages
- [X] [#3144](https://github.com/kubernetes/ingress-nginx/pull/3144) Fix incorrect .DisableLua access.
- [X] [#3145](https://github.com/kubernetes/ingress-nginx/pull/3145) Add "use-regex" Annotation to Toggle Regular Expression Location Modifier
- [X] [#3146](https://github.com/kubernetes/ingress-nginx/pull/3146) Update default backend image
- [X] [#3147](https://github.com/kubernetes/ingress-nginx/pull/3147) Fix error publishing docs [skip ci]
- [X] [#3149](https://github.com/kubernetes/ingress-nginx/pull/3149) Add e2e Tests for Proxy Annotations
- [X] [#3151](https://github.com/kubernetes/ingress-nginx/pull/3151) Add e2e test for SSL-Ciphers
- [X] [#3159](https://github.com/kubernetes/ingress-nginx/pull/3159) Pass --shell to minikube docker-env
- [X] [#3178](https://github.com/kubernetes/ingress-nginx/pull/3178) Update nginx to 1.15.5
- [X] [#3179](https://github.com/kubernetes/ingress-nginx/pull/3179) Update nginx image
- [X] [#3182](https://github.com/kubernetes/ingress-nginx/pull/3182) Allow curly braces to be used in regex paths

_Documentation:_

- [X] [#3021](https://github.com/kubernetes/ingress-nginx/pull/3021) Fix documentation search
- [X] [#3027](https://github.com/kubernetes/ingress-nginx/pull/3027) Add documentation about running Ingress NGINX on bare-metal
- [X] [#3039](https://github.com/kubernetes/ingress-nginx/pull/3039) Remove link to invalid example [ci-skip]
- [X] [#3046](https://github.com/kubernetes/ingress-nginx/pull/3046) Document when to modify ELB idle timeouts and set default value to 60s
- [X] [#3059](https://github.com/kubernetes/ingress-nginx/pull/3059) fix some typos
- [X] [#3068](https://github.com/kubernetes/ingress-nginx/pull/3068) Complete documentation about SSL Passthrough
- [X] [#3074](https://github.com/kubernetes/ingress-nginx/pull/3074) Add MetalLB to bare-metal deployment page
- [X] [#3090](https://github.com/kubernetes/ingress-nginx/pull/3090) Add note about default namespace and merge behavior
- [X] [#3092](https://github.com/kubernetes/ingress-nginx/pull/3092) Update mkdocs and travis-ci
- [X] [#3094](https://github.com/kubernetes/ingress-nginx/pull/3094) Fix baremetal images [skip ci]
- [X] [#3097](https://github.com/kubernetes/ingress-nginx/pull/3097) Added notes to  regarding external access when using TCP/UDP proxy in Ingress
- [X] [#3102](https://github.com/kubernetes/ingress-nginx/pull/3102) Replace kubernetes-users mailing list links with discuss forum link
- [X] [#3111](https://github.com/kubernetes/ingress-nginx/pull/3111) doc issue related to monitor part
- [X] [#3113](https://github.com/kubernetes/ingress-nginx/pull/3113) fix typos
- [X] [#3115](https://github.com/kubernetes/ingress-nginx/pull/3115) Fixed link to aws elastic loadbalancer
- [X] [#3162](https://github.com/kubernetes/ingress-nginx/pull/3162) update name of config map in README.md
- [X] [#3175](https://github.com/kubernetes/ingress-nginx/pull/3175) Fix yaml indentation in annotations server-snippet doc

### 0.19.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.19.0`

_New Features:_

- NGINX 1.15.3
- Serve SSL certificates synamically instead of reloading NGINX when they are created, updated, or deleted.
  Feature behind the flag `--enable-dynamic-certificates`
- GDB binary is included in the image to help [troubleshooting issues](https://github.com/kubernetes/ingress-nginx/pull/3002)
- Adjust the number of CPUs when CGROUP limits are defined (`worker-processes=auto` uses all the availables)

_Changes:_

- [x] [#2616](https://github.com/kubernetes/ingress-nginx/pull/2616) Add use-forwarded-headers configmap option.
- [x] [#2857](https://github.com/kubernetes/ingress-nginx/pull/2857) remove unnecessary encoding/decoding also fix ipv6 issue
- [x] [#2884](https://github.com/kubernetes/ingress-nginx/pull/2884) [grafana] Rate over 2 minutes since default Prometheus interval is 1m
- [x] [#2889](https://github.com/kubernetes/ingress-nginx/pull/2889) Add Lua endpoint to support dynamic certificate serving functionality
- [x] [#2899](https://github.com/kubernetes/ingress-nginx/pull/2899) fixed rewrites for paths not ending in /
- [x] [#2923](https://github.com/kubernetes/ingress-nginx/pull/2923) Add dynamic certificate serving feature to controller
- [x] [#2925](https://github.com/kubernetes/ingress-nginx/pull/2925) Update nginx dependencies
- [x] [#2932](https://github.com/kubernetes/ingress-nginx/pull/2932) Fixed typo in flags.go
- [x] [#2934](https://github.com/kubernetes/ingress-nginx/pull/2934) Datasource input variable
- [x] [#2941](https://github.com/kubernetes/ingress-nginx/pull/2941) now actually using the $controller and $namespace variables
- [x] [#2942](https://github.com/kubernetes/ingress-nginx/pull/2942) Update nginx image
- [x] [#2946](https://github.com/kubernetes/ingress-nginx/pull/2946) Add unit tests to configuration_test.lua that cover Backends configuration
- [x] [#2955](https://github.com/kubernetes/ingress-nginx/pull/2955) Update nginx opentracing zipkin module
- [x] [#2956](https://github.com/kubernetes/ingress-nginx/pull/2956) Update nginx and e2e images
- [x] [#2957](https://github.com/kubernetes/ingress-nginx/pull/2957) Batch metrics and flush periodically
- [x] [#2964](https://github.com/kubernetes/ingress-nginx/pull/2964) fix variable parsing when key is number
- [x] [#2965](https://github.com/kubernetes/ingress-nginx/pull/2965) Add Lua module to serve SSL Certificates dynamically
- [x] [#2966](https://github.com/kubernetes/ingress-nginx/pull/2966) Add unit tests for sticky lua module
- [x] [#2970](https://github.com/kubernetes/ingress-nginx/pull/2970) Update labels
- [x] [#2972](https://github.com/kubernetes/ingress-nginx/pull/2972) consistently fallback to default certificate when TLS is configured
- [x] [#2977](https://github.com/kubernetes/ingress-nginx/pull/2977) Pass real source IP address to auth request
- [x] [#2979](https://github.com/kubernetes/ingress-nginx/pull/2979) clear dynamic configuration e2e tests
- [x] [#2987](https://github.com/kubernetes/ingress-nginx/pull/2987) cleanup dynamic cert e2e tests
- [x] [#2988](https://github.com/kubernetes/ingress-nginx/pull/2988) Update go to 1.11
- [x] [#2990](https://github.com/kubernetes/ingress-nginx/pull/2990) Check if cgroup cpu limits are defined to get the number of CPUs
- [x] [#3003](https://github.com/kubernetes/ingress-nginx/pull/3003) Update nginx to 1.15.3
- [x] [#3004](https://github.com/kubernetes/ingress-nginx/pull/3004) Update nginx image
- [x] [#3005](https://github.com/kubernetes/ingress-nginx/pull/3005) Fix gdb issue and update e2e image
- [x] [#3006](https://github.com/kubernetes/ingress-nginx/pull/3006) apply nginx patch to make ssl_certificate_by_lua_block work properly
- [x] [#3011](https://github.com/kubernetes/ingress-nginx/pull/3011) Update nginx image

_Documentation:_

- [x] [#2806](https://github.com/kubernetes/ingress-nginx/pull/2806) add help for tls prerequisite for ingress.yaml
- [x] [#2912](https://github.com/kubernetes/ingress-nginx/pull/2912) Add documentation to install prometheus and grafana
- [x] [#2928](https://github.com/kubernetes/ingress-nginx/pull/2928) docs: Precisations on the usage of the InfluxDB module
- [x] [#2962](https://github.com/kubernetes/ingress-nginx/pull/2962) Fix broken anchor link to GCE/GKE
- [x] [#2983](https://github.com/kubernetes/ingress-nginx/pull/2983) Add documentation for enable-dynamic-certificates feature
- [x] [#2998](https://github.com/kubernetes/ingress-nginx/pull/2998) fixed jsonpath command in examples
- [x] [#3002](https://github.com/kubernetes/ingress-nginx/pull/3002) Enhance Troubleshooting Documentation

### 0.18.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.18.0`

_New Features:_

- NGINX 1.15.2
- Dynamic configuration is enabled by default
- Support for AJP protocol
- Use of authbind to bind privileged ports
- Replace minikube with [kubeadm-dind-cluster](https://github.com/kubernetes-sigs/kubeadm-dind-cluster) to run e2e tests

_Changes:_

- [x] [#2789](https://github.com/kubernetes/ingress-nginx/pull/2789) Remove KubeConfig Dependency for Store Tests
- [x] [#2794](https://github.com/kubernetes/ingress-nginx/pull/2794) enable dynamic backend configuration by default
- [x] [#2795](https://github.com/kubernetes/ingress-nginx/pull/2795) start minikube before trying to build the image
- [x] [#2804](https://github.com/kubernetes/ingress-nginx/pull/2804) add support for ExternalName service type in dynamic mode
- [x] [#2808](https://github.com/kubernetes/ingress-nginx/pull/2808) fix the bug #2799, add prefix (?i) in rewrite statement.
- [x] [#2811](https://github.com/kubernetes/ingress-nginx/pull/2811) Escape $request_uri for external auth
- [x] [#2812](https://github.com/kubernetes/ingress-nginx/pull/2812) modified annotation name "rewrite-to" to "rewrite-target" in comments
- [x] [#2819](https://github.com/kubernetes/ingress-nginx/pull/2819) Catch errors waiting for controller deployment
- [x] [#2823](https://github.com/kubernetes/ingress-nginx/pull/2823) Multiple optimizations to build targets
- [x] [#2825](https://github.com/kubernetes/ingress-nginx/pull/2825) Refactoring of how we run as user
- [x] [#2826](https://github.com/kubernetes/ingress-nginx/pull/2826) Remove setcap from image and update nginx to 0.15.1
- [x] [#2827](https://github.com/kubernetes/ingress-nginx/pull/2827) Use nginx image as base and install go on top
- [x] [#2829](https://github.com/kubernetes/ingress-nginx/pull/2829) use resty-cli for running lua unit tests
- [x] [#2830](https://github.com/kubernetes/ingress-nginx/pull/2830) Remove lua mocks
- [x] [#2834](https://github.com/kubernetes/ingress-nginx/pull/2834) Added permanent-redirect-code
- [x] [#2844](https://github.com/kubernetes/ingress-nginx/pull/2844) Do not allow invalid latency values in metrics
- [x] [#2852](https://github.com/kubernetes/ingress-nginx/pull/2852) fix custom-error-pages functionality in dynamic mode
- [x] [#2853](https://github.com/kubernetes/ingress-nginx/pull/2853) improve annotations/default_backend e2e test
- [x] [#2858](https://github.com/kubernetes/ingress-nginx/pull/2858) Update build image
- [x] [#2859](https://github.com/kubernetes/ingress-nginx/pull/2859) Fix inconsistent metric labels
- [x] [#2863](https://github.com/kubernetes/ingress-nginx/pull/2863) Replace minikube for e2e tests
- [x] [#2867](https://github.com/kubernetes/ingress-nginx/pull/2867) fix bug with lua e2e test suite
- [x] [#2868](https://github.com/kubernetes/ingress-nginx/pull/2868) Use an existing e2e image
- [x] [#2869](https://github.com/kubernetes/ingress-nginx/pull/2869) describe under what circumstances and how we avoid Nginx reload
- [x] [#2871](https://github.com/kubernetes/ingress-nginx/pull/2871) Add support for AJP protocol
- [x] [#2872](https://github.com/kubernetes/ingress-nginx/pull/2872) Update nginx to 1.15.2
- [x] [#2874](https://github.com/kubernetes/ingress-nginx/pull/2874) Delay initial prometheus status metric
- [x] [#2876](https://github.com/kubernetes/ingress-nginx/pull/2876) Remove dashboard an tune sync-frequency
- [x] [#2877](https://github.com/kubernetes/ingress-nginx/pull/2877) Refactor entrypoint to avoid issues with volumes
- [x] [#2885](https://github.com/kubernetes/ingress-nginx/pull/2885) fix: Sort TCP/UDP upstream order
- [x] [#2888](https://github.com/kubernetes/ingress-nginx/pull/2888) Fix grafana datasources
- [x] [#2890](https://github.com/kubernetes/ingress-nginx/pull/2890) Usability improvements to build steps
- [x] [#2893](https://github.com/kubernetes/ingress-nginx/pull/2893) Update nginx image
- [x] [#2894](https://github.com/kubernetes/ingress-nginx/pull/2894) Use authbind to bind privileged ports
- [x] [#2895](https://github.com/kubernetes/ingress-nginx/pull/2895) support custom configuration to main context of nginx config
- [x] [#2896](https://github.com/kubernetes/ingress-nginx/pull/2896) support configuring multi_accept directive via configmap
- [x] [#2897](https://github.com/kubernetes/ingress-nginx/pull/2897) Enable reuse-port by default
- [x] [#2905](https://github.com/kubernetes/ingress-nginx/pull/2905) Fix IPV6 detection

_Documentation:_

- [x] [#2816](https://github.com/kubernetes/ingress-nginx/pull/2816) doc log-format: add variables about ingress
- [x] [#2866](https://github.com/kubernetes/ingress-nginx/pull/2866) Update index.md
- [x] [#2898](https://github.com/kubernetes/ingress-nginx/pull/2898) Fix default sync-period doc
- [x] [#2903](https://github.com/kubernetes/ingress-nginx/pull/2903) Very minor grammar fix

### 0.17.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.17.1`

_Changes:_

- [x] [#2782](https://github.com/kubernetes/ingress-nginx/pull/2782) Add Better Error Handling for SSLSessionTicketKey
- [x] [#2790](https://github.com/kubernetes/ingress-nginx/pull/2790) Update prometheus labels

_Documentation:_

- [x] [#2770](https://github.com/kubernetes/ingress-nginx/pull/2770) Basic-Auth doc misleading: fix double quotes leading to nginx config error

### 0.17.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.17.0`

_New Features:_

- [Grafana dashboards](https://github.com/kubernetes/ingress-nginx/tree/master/deploy/grafana/dashboards)

_Changes:_

- [x] [#2705](https://github.com/kubernetes/ingress-nginx/pull/2705) Remove duplicated securityContext
- [x] [#2719](https://github.com/kubernetes/ingress-nginx/pull/2719) Sample rate configmap option for zipkin in nginx-opentracing
- [x] [#2726](https://github.com/kubernetes/ingress-nginx/pull/2726) Cleanup prometheus metrics after a reload
- [x] [#2727](https://github.com/kubernetes/ingress-nginx/pull/2727) Add e2e tests for Client-Body-Buffer-Size
- [x] [#2732](https://github.com/kubernetes/ingress-nginx/pull/2732) Improve logging
- [x] [#2741](https://github.com/kubernetes/ingress-nginx/pull/2741) Add redirect uri for oauth2 login
- [x] [#2744](https://github.com/kubernetes/ingress-nginx/pull/2744) fix: Use the correct opentracing plugin for Jaeger
- [x] [#2747](https://github.com/kubernetes/ingress-nginx/pull/2747) Update opentracing-cpp and modsecurity
- [x] [#2748](https://github.com/kubernetes/ingress-nginx/pull/2748) Update nginx image to 0.54
- [x] [#2749](https://github.com/kubernetes/ingress-nginx/pull/2749) Use docker to build go binaries
- [x] [#2754](https://github.com/kubernetes/ingress-nginx/pull/2754) Allow gzip compression level to be controlled via ConfigMap
- [x] [#2760](https://github.com/kubernetes/ingress-nginx/pull/2760) Fix ingress rule parsing error
- [x] [#2767](https://github.com/kubernetes/ingress-nginx/pull/2767) Fix regression introduced in #2732
- [x] [#2771](https://github.com/kubernetes/ingress-nginx/pull/2771) Grafana Dashboard
- [x] [#2775](https://github.com/kubernetes/ingress-nginx/pull/2775) Simplify handler registration and updates prometheus
- [x] [#2776](https://github.com/kubernetes/ingress-nginx/pull/2776) Fix configuration hash calculation

_Documentation:_

- [x] [#2717](https://github.com/kubernetes/ingress-nginx/pull/2717) GCE/GKE proxy mentioned for Azure
- [x] [#2743](https://github.com/kubernetes/ingress-nginx/pull/2743) Clarify Installation Document by Separating Helm Steps
- [x] [#2761](https://github.com/kubernetes/ingress-nginx/pull/2761) Fix spelling mistake
- [x] [#2764](https://github.com/kubernetes/ingress-nginx/pull/2764) Use language neutral links to MDN
- [x] [#2765](https://github.com/kubernetes/ingress-nginx/pull/2765) Add FOSSA status badge
- [x] [#2777](https://github.com/kubernetes/ingress-nginx/pull/2777) Build docs using local docker image [ci skip]

### 0.16.2

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.16.2`

_Breaking changes:_

Running as user requires an update in the deployment manifest.

```yaml
securityContext:
  capabilities:
    drop:
      - ALL
    add:
      - NET_BIND_SERVICE
  # www-data -> 33
  runAsUser: 33
```

Note: the deploy [guide](https://kubernetes.github.io/ingress-nginx/deploy/#mandatory-command) contains this change

_Changes:_

- [x] [#2678](https://github.com/kubernetes/ingress-nginx/pull/2678) Refactor server type to include SSLCert
- [x] [#2685](https://github.com/kubernetes/ingress-nginx/pull/2685) Fix qemu docker build
- [x] [#2696](https://github.com/kubernetes/ingress-nginx/pull/2696) If server_tokens is disabled completely remove the Server header
- [x] [#2698](https://github.com/kubernetes/ingress-nginx/pull/2698) Improve best-cert guessing with empty tls.hosts
- [x] [#2701](https://github.com/kubernetes/ingress-nginx/pull/2701) Remove prometheus labels with high cardinality

_Documentation:_

- [x] [#2368](https://github.com/kubernetes/ingress-nginx/pull/2368) [aggregate] Fix typos across codebase
- [x] [#2681](https://github.com/kubernetes/ingress-nginx/pull/2681) Typo fix in error message: encounted->encountered
- [x] [#2697](https://github.com/kubernetes/ingress-nginx/pull/2697) Enhance Distributed Tracing Documentation

### 0.16.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.16.1`

_Breaking changes:_

Running as user requires an update in the deployment manifest.

```yaml
securityContext:
  capabilities:
    drop:
      - ALL
    add:
      - NET_BIND_SERVICE
  # www-data -> 33
  runAsUser: 33
```

Note: the deploy [guide](https://kubernetes.github.io/ingress-nginx/deploy/#mandatory-command) contains this change

_New Features:_

- Run as user dropping root privileges
- New prometheus metric implementation (VTS module was removed)
- [InfluxDB integration](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#influxdb)
- [Module GeoIP2](https://github.com/leev/ngx_http_geoip2_module)

_Changes:_

- [x] [#2692](https://github.com/kubernetes/ingress-nginx/pull/2692) Fix initial read of configuration configmap
- [x] [#2693](https://github.com/kubernetes/ingress-nginx/pull/2693) Revert #2669
- [x] [#2694](https://github.com/kubernetes/ingress-nginx/pull/2694) Add note about status update

### 0.16.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.16.1`

_Breaking changes:_

Running as user requires an update in the deployment manifest.

```yaml
securityContext:
  capabilities:
    drop:
      - ALL
    add:
      - NET_BIND_SERVICE
  # www-data -> 33
  runAsUser: 33
```

Note: the deploy [guide](https://kubernetes.github.io/ingress-nginx/deploy/#mandatory-command) contains this change

_New Features:_

- Run as user dropping root privileges
- New prometheus metric implementation (VTS module was removed)
- [InfluxDB integration](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#influxdb)
- [Module GeoIP2](https://github.com/leev/ngx_http_geoip2_module)

_Changes:_

- [x] [#2423](https://github.com/kubernetes/ingress-nginx/pull/2423) Resolves issue with proxy-redirect nginx configuration
- [x] [#2451](https://github.com/kubernetes/ingress-nginx/pull/2451) fix for #1930, make sessions sticky, for ingress with multiple rules …
- [x] [#2484](https://github.com/kubernetes/ingress-nginx/pull/2484) Fix bugs in Lua implementation of sticky sessions
- [x] [#2486](https://github.com/kubernetes/ingress-nginx/pull/2486) Extend kubernetes interrelation variables in nginx.tmpl
- [x] [#2504](https://github.com/kubernetes/ingress-nginx/pull/2504) Add Timeout For TLS Passthrough
- [x] [#2505](https://github.com/kubernetes/ingress-nginx/pull/2505) Annotations for the InfluxDB module
- [x] [#2517](https://github.com/kubernetes/ingress-nginx/pull/2517) Fix typo about the kind of request
- [x] [#2523](https://github.com/kubernetes/ingress-nginx/pull/2523) Add tests for bind-address
- [x] [#2524](https://github.com/kubernetes/ingress-nginx/pull/2524) Add support for grpc_set_header
- [x] [#2526](https://github.com/kubernetes/ingress-nginx/pull/2526) Fix upstream hash lua test
- [x] [#2528](https://github.com/kubernetes/ingress-nginx/pull/2528) Remove go-bindata
- [x] [#2533](https://github.com/kubernetes/ingress-nginx/pull/2533) NGINX image update: add the influxdb module
- [x] [#2534](https://github.com/kubernetes/ingress-nginx/pull/2534) Set Focus for E2E Tests
- [x] [#2537](https://github.com/kubernetes/ingress-nginx/pull/2537) Update nginx modules
- [x] [#2542](https://github.com/kubernetes/ingress-nginx/pull/2542) Instrument controller to show configReload metrics
- [x] [#2543](https://github.com/kubernetes/ingress-nginx/pull/2543) introduce a balancer interface
- [x] [#2548](https://github.com/kubernetes/ingress-nginx/pull/2548) Implement generate-request-id
- [x] [#2554](https://github.com/kubernetes/ingress-nginx/pull/2554) use better defaults for proxy-next-upstream(-tries)
- [x] [#2558](https://github.com/kubernetes/ingress-nginx/pull/2558) Update qemu to 2.12.0 [ci skip]
- [x] [#2559](https://github.com/kubernetes/ingress-nginx/pull/2559) Add geoip2 module and DB to nginx build
- [x] [#2564](https://github.com/kubernetes/ingress-nginx/pull/2564) Add security contacts file [ci skip]
- [x] [#2569](https://github.com/kubernetes/ingress-nginx/pull/2569) Update nginx modules to fix core dump [ci skip]
- [x] [#2570](https://github.com/kubernetes/ingress-nginx/pull/2570) Enable core dumps during tests
- [x] [#2573](https://github.com/kubernetes/ingress-nginx/pull/2573) Refactor e2e tests and update go dependencies
- [x] [#2574](https://github.com/kubernetes/ingress-nginx/pull/2574) Fix default-backend annotation
- [x] [#2575](https://github.com/kubernetes/ingress-nginx/pull/2575) Print information about NGINX version
- [x] [#2577](https://github.com/kubernetes/ingress-nginx/pull/2577) make sure ingress-nginx instances are watching their namespace only during test runs
- [x] [#2588](https://github.com/kubernetes/ingress-nginx/pull/2588) Update nginx dependencies
- [x] [#2590](https://github.com/kubernetes/ingress-nginx/pull/2590) Typo fix: muthual autentication -> mutual authentication
- [x] [#2591](https://github.com/kubernetes/ingress-nginx/pull/2591) Access log improvements
- [x] [#2597](https://github.com/kubernetes/ingress-nginx/pull/2597) Fix arm paths for liblua.so and lua_package_cpath
- [x] [#2598](https://github.com/kubernetes/ingress-nginx/pull/2598) Always sort upstream list to provide stable iteration order
- [x] [#2600](https://github.com/kubernetes/ingress-nginx/pull/2600) typo fix futher to further && preformance to performance
- [x] [#2602](https://github.com/kubernetes/ingress-nginx/pull/2602) Crossplat fixes
- [x] [#2603](https://github.com/kubernetes/ingress-nginx/pull/2603) Bump nginx influxdb module to f8732268d44aea706ecf8d9c6036e9b6dacc99b2
- [x] [#2608](https://github.com/kubernetes/ingress-nginx/pull/2608) Expose UDP message on /metrics endpoint
- [x] [#2611](https://github.com/kubernetes/ingress-nginx/pull/2611) Add metric emitter lua module
- [x] [#2614](https://github.com/kubernetes/ingress-nginx/pull/2614) fix nginx conf test error when not found active service endpoints
- [x] [#2617](https://github.com/kubernetes/ingress-nginx/pull/2617) Update go to 1.10.3
- [x] [#2618](https://github.com/kubernetes/ingress-nginx/pull/2618) Update nginx to 1.15.0 and remove VTS module
- [x] [#2619](https://github.com/kubernetes/ingress-nginx/pull/2619) Run as user dropping privileges
- [x] [#2623](https://github.com/kubernetes/ingress-nginx/pull/2623) Proofread cmd package and update flags description
- [x] [#2634](https://github.com/kubernetes/ingress-nginx/pull/2634) Disable resync period
- [x] [#2636](https://github.com/kubernetes/ingress-nginx/pull/2636) Add missing equality comparisons for ingress.Server
- [x] [#2638](https://github.com/kubernetes/ingress-nginx/pull/2638) Wait the result of the controller deployment before running any test
- [x] [#2639](https://github.com/kubernetes/ingress-nginx/pull/2639) Clarify log messages in controller package
- [x] [#2643](https://github.com/kubernetes/ingress-nginx/pull/2643) Remove VTS from the ingress controller
- [x] [#2644](https://github.com/kubernetes/ingress-nginx/pull/2644) Update nginx image version
- [x] [#2646](https://github.com/kubernetes/ingress-nginx/pull/2646) Rollback nginx 1.15.0 to 1.13.12
- [x] [#2649](https://github.com/kubernetes/ingress-nginx/pull/2649) Add support for IPV6 in stream upstream servers
- [x] [#2652](https://github.com/kubernetes/ingress-nginx/pull/2652) Use a unix socket instead udp for reception of metrics
- [x] [#2653](https://github.com/kubernetes/ingress-nginx/pull/2653) Remove dummy file watcher
- [x] [#2654](https://github.com/kubernetes/ingress-nginx/pull/2654) Hotfix: influxdb module enable disable toggle
- [x] [#2656](https://github.com/kubernetes/ingress-nginx/pull/2656) Improve configuration change detection
- [x] [#2658](https://github.com/kubernetes/ingress-nginx/pull/2658) Do not wait informer initialization to read configuration
- [x] [#2659](https://github.com/kubernetes/ingress-nginx/pull/2659) Update nginx image
- [x] [#2660](https://github.com/kubernetes/ingress-nginx/pull/2660) Change modsecurity directories
- [x] [#2661](https://github.com/kubernetes/ingress-nginx/pull/2661) Add additional header when debug is enabled
- [x] [#2664](https://github.com/kubernetes/ingress-nginx/pull/2664) refactor some lua code
- [x] [#2669](https://github.com/kubernetes/ingress-nginx/pull/2669) Remove unnecessary sync when the leader change
- [x] [#2672](https://github.com/kubernetes/ingress-nginx/pull/2672) After a configmap change parse ingress annotations (again)
- [x] [#2673](https://github.com/kubernetes/ingress-nginx/pull/2673) Add new approvers to the project
- [x] [#2674](https://github.com/kubernetes/ingress-nginx/pull/2674) Add e2e test for configmap change and reload
- [x] [#2675](https://github.com/kubernetes/ingress-nginx/pull/2675) Update opentracing nginx module
- [x] [#2676](https://github.com/kubernetes/ingress-nginx/pull/2676) Update opentracing configuration

_Documentation:_

- [x] [#2479](https://github.com/kubernetes/ingress-nginx/pull/2479) Document how the NGINX Ingress controller build nginx.conf
- [x] [#2515](https://github.com/kubernetes/ingress-nginx/pull/2515) Simplify installation and e2e manifests
- [x] [#2531](https://github.com/kubernetes/ingress-nginx/pull/2531) Mention the #ingress-nginx Slack channel
- [x] [#2540](https://github.com/kubernetes/ingress-nginx/pull/2540) DOCS: Correct ssl-passthrough annotation description.
- [x] [#2544](https://github.com/kubernetes/ingress-nginx/pull/2544) [docs] Fix manifest URL for GKE + Azure
- [x] [#2566](https://github.com/kubernetes/ingress-nginx/pull/2566) Fix wrong default value for `enable-brotli`
- [x] [#2581](https://github.com/kubernetes/ingress-nginx/pull/2581) Improved link in modsecurity.md
- [x] [#2583](https://github.com/kubernetes/ingress-nginx/pull/2583) docs: add secret scheme details to the example
- [x] [#2592](https://github.com/kubernetes/ingress-nginx/pull/2592) Typo fix: are be->are/to on->to
- [x] [#2595](https://github.com/kubernetes/ingress-nginx/pull/2595) Typo fix: successfull->successful
- [x] [#2601](https://github.com/kubernetes/ingress-nginx/pull/2601) fix changelog link in README.md
- [x] [#2624](https://github.com/kubernetes/ingress-nginx/pull/2624) Fix minor documentation example
- [x] [#2625](https://github.com/kubernetes/ingress-nginx/pull/2625) Add annotation doc on proxy buffer size
- [x] [#2630](https://github.com/kubernetes/ingress-nginx/pull/2630) Update documentation for custom error pages
- [x] [#2666](https://github.com/kubernetes/ingress-nginx/pull/2666) Add documentation for proxy-cookie-domain annotation (#2034)

### 0.15.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.15.0`

_Changes:_

- [x] [#2440](https://github.com/kubernetes/ingress-nginx/pull/2440) TLS tests
- [x] [#2443](https://github.com/kubernetes/ingress-nginx/pull/2443) improve build-dev-env.sh script
- [x] [#2446](https://github.com/kubernetes/ingress-nginx/pull/2446) always use x-request-id
- [x] [#2447](https://github.com/kubernetes/ingress-nginx/pull/2447) Add basic security context to deployment YAMLs
- [x] [#2453](https://github.com/kubernetes/ingress-nginx/pull/2453) Add google analytics [ci skip]
- [x] [#2456](https://github.com/kubernetes/ingress-nginx/pull/2456) Assert or install go-bindata before incanting
- [x] [#2472](https://github.com/kubernetes/ingress-nginx/pull/2472) Refactor Lua balancer
- [x] [#2477](https://github.com/kubernetes/ingress-nginx/pull/2477) Change TrimLeft for TrimPrefix on the from-to-www redirect
- [x] [#2490](https://github.com/kubernetes/ingress-nginx/pull/2490) add resty cookie
- [x] [#2495](https://github.com/kubernetes/ingress-nginx/pull/2495) [ci skip] bump nginx baseimage version
- [x] [#2501](https://github.com/kubernetes/ingress-nginx/pull/2501) Refactor update of status removing initial check for loadbalancer
- [x] [#2502](https://github.com/kubernetes/ingress-nginx/pull/2502) Update go version in fortune teller image
- [x] [#2511](https://github.com/kubernetes/ingress-nginx/pull/2511) force backend sync when worker starts
- [x] [#2512](https://github.com/kubernetes/ingress-nginx/pull/2512) Remove warning when secret is used only for authentication
- [x] [#2514](https://github.com/kubernetes/ingress-nginx/pull/2514) Fix and simplify local dev workflow and execution of e2e tests

_Documentation:_

- [x] [#2448](https://github.com/kubernetes/ingress-nginx/pull/2448) Update GitHub pull request template
- [x] [#2449](https://github.com/kubernetes/ingress-nginx/pull/2449) Improve documentation format
- [x] [#2454](https://github.com/kubernetes/ingress-nginx/pull/2454) Add gRPC annotation doc
- [x] [#2455](https://github.com/kubernetes/ingress-nginx/pull/2455) Adjust size of tables and only adjust the first column on mobile
- [x] [#2457](https://github.com/kubernetes/ingress-nginx/pull/2457) Add Getting the Code section to Quick Start
- [x] [#2464](https://github.com/kubernetes/ingress-nginx/pull/2464) Documentation fixes & improvements
- [x] [#2467](https://github.com/kubernetes/ingress-nginx/pull/2467) Fixed broken link in deploy README
- [x] [#2498](https://github.com/kubernetes/ingress-nginx/pull/2498) Add some clarification around multiple ingress controller behavior
- [x] [#2503](https://github.com/kubernetes/ingress-nginx/pull/2503) Add KubeCon Europe 2018 Video to documentation

### 0.14.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.14.0`

_New Features:_

- [Documentation web page](https://kubernetes.github.io/ingress-nginx/)
- Support for `upstream-hash-by` annotation in dynamic configuration mode
- Improved e2e test suite

_Changes:_

- [x] [#2346](https://github.com/kubernetes/ingress-nginx/pull/2346) Move ConfigMap updating methods into e2e/framework
- [x] [#2347](https://github.com/kubernetes/ingress-nginx/pull/2347) Update owners
- [x] [#2348](https://github.com/kubernetes/ingress-nginx/pull/2348) Use same convention, curl + kubectl for GKE
- [x] [#2350](https://github.com/kubernetes/ingress-nginx/pull/2350) Correct some returned messages in server_tokens.go
- [x] [#2352](https://github.com/kubernetes/ingress-nginx/pull/2352) Correct some info in flags.go
- [x] [#2353](https://github.com/kubernetes/ingress-nginx/pull/2353) Add proxy-add-original-uri-header config flag
- [x] [#2356](https://github.com/kubernetes/ingress-nginx/pull/2356) Add vts-sum-key config flag
- [x] [#2361](https://github.com/kubernetes/ingress-nginx/pull/2361) Check ingress rule contains HTTP paths
- [x] [#2363](https://github.com/kubernetes/ingress-nginx/pull/2363) Review $request_id
- [x] [#2365](https://github.com/kubernetes/ingress-nginx/pull/2365) Clean JSON before post request to update configuration
- [x] [#2369](https://github.com/kubernetes/ingress-nginx/pull/2369) Update nginx image to fix modsecurity crs issues
- [x] [#2370](https://github.com/kubernetes/ingress-nginx/pull/2370) Update nginx image
- [x] [#2374](https://github.com/kubernetes/ingress-nginx/pull/2374) Remove most of the time.Sleep from the e2e tests
- [x] [#2379](https://github.com/kubernetes/ingress-nginx/pull/2379) Add busted unit testing framework for lua code
- [x] [#2382](https://github.com/kubernetes/ingress-nginx/pull/2382) Accept ns/name Secret reference in annotations
- [x] [#2383](https://github.com/kubernetes/ingress-nginx/pull/2383) Improve speed of e2e tests
- [x] [#2385](https://github.com/kubernetes/ingress-nginx/pull/2385) include lua-resty-balancer in nginx image
- [x] [#2386](https://github.com/kubernetes/ingress-nginx/pull/2386) upstream-hash-by annotation support for dynamic configuraton mode
- [x] [#2388](https://github.com/kubernetes/ingress-nginx/pull/2388) Silence unnecessary MissingAnnotations errors
- [x] [#2392](https://github.com/kubernetes/ingress-nginx/pull/2392) Ensure dep fix fsnotify
- [x] [#2395](https://github.com/kubernetes/ingress-nginx/pull/2395) Fix flaky test
- [x] [#2396](https://github.com/kubernetes/ingress-nginx/pull/2396) Update go dependencies
- [x] [#2398](https://github.com/kubernetes/ingress-nginx/pull/2398) Allow tls section without hosts in Ingress rule
- [x] [#2399](https://github.com/kubernetes/ingress-nginx/pull/2399) Add test for store helper ListIngresses
- [x] [#2401](https://github.com/kubernetes/ingress-nginx/pull/2401) Add tests for controller getEndpoints
- [x] [#2408](https://github.com/kubernetes/ingress-nginx/pull/2408) Read backends data even if buffered to temp file
- [x] [#2410](https://github.com/kubernetes/ingress-nginx/pull/2410) Add balancer unit tests
- [x] [#2411](https://github.com/kubernetes/ingress-nginx/pull/2411) Update nginx-opentracing to 0.3.0
- [x] [#2414](https://github.com/kubernetes/ingress-nginx/pull/2414) Fix golint installation
- [x] [#2416](https://github.com/kubernetes/ingress-nginx/pull/2416) Update nginx image
- [x] [#2417](https://github.com/kubernetes/ingress-nginx/pull/2417) Automate building developer environment
- [x] [#2421](https://github.com/kubernetes/ingress-nginx/pull/2421) Apply gometalinter suggestions
- [x] [#2428](https://github.com/kubernetes/ingress-nginx/pull/2428) Add buffer configuration to external auth location config
- [x] [#2433](https://github.com/kubernetes/ingress-nginx/pull/2433) Remove data races from tests
- [x] [#2434](https://github.com/kubernetes/ingress-nginx/pull/2434) Check ginkgo is installed before running e2e tests
- [x] [#2437](https://github.com/kubernetes/ingress-nginx/pull/2437) Add annotation to enable rewrite logs in a location

_Documentation:_

- [x] [#2351](https://github.com/kubernetes/ingress-nginx/pull/2351) Typo fix in cli-arguments.md
- [x] [#2372](https://github.com/kubernetes/ingress-nginx/pull/2372) fix the default cookie name in doc
- [x] [#2377](https://github.com/kubernetes/ingress-nginx/pull/2377) DOCS: Add clarification regarding ssl passthrough
- [x] [#2409](https://github.com/kubernetes/ingress-nginx/pull/2409) Add deployment instructions for Docker for Mac (Edge)
- [x] [#2413](https://github.com/kubernetes/ingress-nginx/pull/2413) Reorganize documentation
- [x] [#2438](https://github.com/kubernetes/ingress-nginx/pull/2438) Update custom-errors.md
- [x] [#2439](https://github.com/kubernetes/ingress-nginx/pull/2439) Update README.md
- [x] [#2430](https://github.com/kubernetes/ingress-nginx/pull/2430) Add scripts and tasks to publish docs to github pages
- [x] [#2431](https://github.com/kubernetes/ingress-nginx/pull/2431) Improve readme file
- [x] [#2366](https://github.com/kubernetes/ingress-nginx/pull/2366) fix: fill missing patch yaml config.
- [x] [#2432](https://github.com/kubernetes/ingress-nginx/pull/2432) Fix broken links in the docs
- [x] [#2436](https://github.com/kubernetes/ingress-nginx/pull/2436) Update exposing-tcp-udp-services.md

### 0.13.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.13.0`

_New Features:_

- NGINX 1.13.12
- Support for gRPC:
  - The annotation `nginx.ingress.kubernetes.io/grpc-backend: "true"` enable this feature
  - If the gRPC service requires TLS `nginx.ingress.kubernetes.io/secure-backends: "true"`
- Configurable load balancing with EWMA
- Support for [lua-resty-waf](https://github.com/p0pr0ck5/lua-resty-waf) as alternative to ModSecurity. [Check configuration guide](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/annotations.md#lua-resty-waf)
- Support for session affinity when dynamic configuration is enabled.
- Add NoAuthLocations and default it to "/.well-known/acme-challenge"

_Changes:_

- [x] [#2078](https://github.com/kubernetes/ingress-nginx/pull/2078) Expose SSL client cert data to external auth provider.
- [x] [#2187](https://github.com/kubernetes/ingress-nginx/pull/2187) Managing a whitelist for \_/nginx_status
- [x] [#2208](https://github.com/kubernetes/ingress-nginx/pull/2208) Add missing lua bindata change
- [x] [#2209](https://github.com/kubernetes/ingress-nginx/pull/2209) fix go test TestSkipEnqueue error, move queue.Run
- [x] [#2210](https://github.com/kubernetes/ingress-nginx/pull/2210) allow ipv6 localhost when enabled
- [x] [#2212](https://github.com/kubernetes/ingress-nginx/pull/2212) Fix dynamic configuration when custom errors are enabled
- [x] [#2215](https://github.com/kubernetes/ingress-nginx/pull/2215) fix wrong config generation when upstream-hash-by is set
- [x] [#2220](https://github.com/kubernetes/ingress-nginx/pull/2220) fix: cannot set $service_name if use rewrite
- [x] [#2221](https://github.com/kubernetes/ingress-nginx/pull/2221) Update nginx to 1.13.10 and enable gRPC
- [x] [#2223](https://github.com/kubernetes/ingress-nginx/pull/2223) Add support for gRPC
- [x] [#2227](https://github.com/kubernetes/ingress-nginx/pull/2227) do not hardcode keepalive for upstream_balancer
- [x] [#2228](https://github.com/kubernetes/ingress-nginx/pull/2228) Fix broken links in multi-tls
- [x] [#2229](https://github.com/kubernetes/ingress-nginx/pull/2229) Configurable load balancing with EWMA
- [x] [#2232](https://github.com/kubernetes/ingress-nginx/pull/2232) Make proxy_next_upstream_tries configurable
- [x] [#2233](https://github.com/kubernetes/ingress-nginx/pull/2233) clean backends data before sending to Lua endpoint
- [x] [#2234](https://github.com/kubernetes/ingress-nginx/pull/2234) Update go dependencies
- [x] [#2235](https://github.com/kubernetes/ingress-nginx/pull/2235) add proxy header ssl-client-issuer-dn, fix #2178
- [x] [#2241](https://github.com/kubernetes/ingress-nginx/pull/2241) Revert "Get file max from fs/file-max. (#2050)"
- [x] [#2243](https://github.com/kubernetes/ingress-nginx/pull/2243) Add NoAuthLocations and default it to "/.well-known/acme-challenge"
- [x] [#2244](https://github.com/kubernetes/ingress-nginx/pull/2244) fix: empty ingress path
- [x] [#2246](https://github.com/kubernetes/ingress-nginx/pull/2246) Fix grpc json tag name
- [x] [#2254](https://github.com/kubernetes/ingress-nginx/pull/2254) e2e tests for dynamic configuration and Lua features and a bug fix
- [x] [#2263](https://github.com/kubernetes/ingress-nginx/pull/2263) clean up tmpl
- [x] [#2270](https://github.com/kubernetes/ingress-nginx/pull/2270) Revert deleted code in #2146
- [x] [#2271](https://github.com/kubernetes/ingress-nginx/pull/2271) Use SharedIndexInformers in place of Informers
- [x] [#2272](https://github.com/kubernetes/ingress-nginx/pull/2272) Disable opentracing for nginx internal urls
- [x] [#2273](https://github.com/kubernetes/ingress-nginx/pull/2273) Update go to 1.10.1
- [x] [#2280](https://github.com/kubernetes/ingress-nginx/pull/2280) Fix bug when auth req is enabled(external authentication)
- [x] [#2283](https://github.com/kubernetes/ingress-nginx/pull/2283) Fix flaky e2e tests
- [x] [#2285](https://github.com/kubernetes/ingress-nginx/pull/2285) Update controller.go
- [x] [#2290](https://github.com/kubernetes/ingress-nginx/pull/2290) Update nginx to 1.13.11
- [x] [#2294](https://github.com/kubernetes/ingress-nginx/pull/2294) Fix HSTS without preload
- [x] [#2296](https://github.com/kubernetes/ingress-nginx/pull/2296) Improve indentation of generated nginx.conf
- [x] [#2298](https://github.com/kubernetes/ingress-nginx/pull/2298) Disable dynamic configuration in s390x and ppc64le
- [x] [#2300](https://github.com/kubernetes/ingress-nginx/pull/2300) Fix race condition when Ingress does not contains a secret
- [x] [#2301](https://github.com/kubernetes/ingress-nginx/pull/2301) include lua-resty-waf and its dependencies in the base Nginx image
- [x] [#2303](https://github.com/kubernetes/ingress-nginx/pull/2303) More lua dependencies
- [x] [#2304](https://github.com/kubernetes/ingress-nginx/pull/2304) Lua resty waf controller
- [x] [#2305](https://github.com/kubernetes/ingress-nginx/pull/2305) Fix issues building nginx image in different platforms
- [x] [#2306](https://github.com/kubernetes/ingress-nginx/pull/2306) Disable lua waf where luajit is not available
- [x] [#2308](https://github.com/kubernetes/ingress-nginx/pull/2308) Add verification of lua load balancer to health check
- [x] [#2309](https://github.com/kubernetes/ingress-nginx/pull/2309) Configure upload limits for setup of lua load balancer
- [x] [#2314](https://github.com/kubernetes/ingress-nginx/pull/2314) annotation to ignore given list of WAF rulesets
- [x] [#2315](https://github.com/kubernetes/ingress-nginx/pull/2315) extra waf rules per ingress
- [x] [#2317](https://github.com/kubernetes/ingress-nginx/pull/2317) run lua-resty-waf in different modes
- [x] [#2327](https://github.com/kubernetes/ingress-nginx/pull/2327) Update nginx to 1.13.12
- [x] [#2328](https://github.com/kubernetes/ingress-nginx/pull/2328) Update nginx image
- [x] [#2331](https://github.com/kubernetes/ingress-nginx/pull/2331) fix nil pointer when ssl with ca.crt
- [x] [#2333](https://github.com/kubernetes/ingress-nginx/pull/2333) disable lua for arch s390x and ppc64le
- [x] [#2340](https://github.com/kubernetes/ingress-nginx/pull/2340) Fix buildupstream name to work with dynamic session affinity
- [x] [#2341](https://github.com/kubernetes/ingress-nginx/pull/2341) Add session affinity to custom load balancing
- [x] [#2342](https://github.com/kubernetes/ingress-nginx/pull/2342) Sync SSL certificates on events

_Documentation:_

- [x] [#2236](https://github.com/kubernetes/ingress-nginx/pull/2236) Add missing configuration in #2235
- [x] [#1785](https://github.com/kubernetes/ingress-nginx/pull/1785) Add deployment docs for AWS NLB
- [x] [#2213](https://github.com/kubernetes/ingress-nginx/pull/2213) Update cli-arguments.md
- [x] [#2219](https://github.com/kubernetes/ingress-nginx/pull/2219) Fix log format documentation
- [x] [#2238](https://github.com/kubernetes/ingress-nginx/pull/2238) Correct typo
- [x] [#2239](https://github.com/kubernetes/ingress-nginx/pull/2239) fix-link
- [x] [#2240](https://github.com/kubernetes/ingress-nginx/pull/2240) fix:"any value other" should be "any other value"
- [x] [#2255](https://github.com/kubernetes/ingress-nginx/pull/2255) Update annotations.md
- [x] [#2267](https://github.com/kubernetes/ingress-nginx/pull/2267) Update README.md
- [x] [#2274](https://github.com/kubernetes/ingress-nginx/pull/2274) Typo fixes in modsecurity.md
- [x] [#2276](https://github.com/kubernetes/ingress-nginx/pull/2276) Update README.md
- [x] [#2282](https://github.com/kubernetes/ingress-nginx/pull/2282) Fix nlb instructions

### 0.12.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.12.0`

_New Features:_

- Live NGINX configuration update without reloading using the flag `--enable-dynamic-configuration` (disabled by default).
- New flag `--publish-status-address` to manually set the Ingress status IP address.
- Add worker-cpu-affinity NGINX option.
- Enable remote logging using syslog.
- Do not redirect `/.well-known/acme-challenge` to HTTPS.

_Changes:_

- [x] [#2125](https://github.com/kubernetes/ingress-nginx/pull/2125) Add GCB config to build defaultbackend
- [x] [#2127](https://github.com/kubernetes/ingress-nginx/pull/2127) Revert deletion of dependency version override
- [x] [#2137](https://github.com/kubernetes/ingress-nginx/pull/2137) Updated log level to v2 for sysctlFSFileMax.
- [x] [#2140](https://github.com/kubernetes/ingress-nginx/pull/2140) Cors header should always be returned
- [x] [#2141](https://github.com/kubernetes/ingress-nginx/pull/2141) Fix error loading modules
- [x] [#2143](https://github.com/kubernetes/ingress-nginx/pull/2143) Only add HSTS headers in HTTPS
- [x] [#2144](https://github.com/kubernetes/ingress-nginx/pull/2144) Add annotation to disable logs in a location
- [x] [#2145](https://github.com/kubernetes/ingress-nginx/pull/2145) Add option in the configuration configmap to enable remote logging
- [x] [#2146](https://github.com/kubernetes/ingress-nginx/pull/2146) In case of TLS errors do not allow traffic
- [x] [#2148](https://github.com/kubernetes/ingress-nginx/pull/2148) Add publish-status-address flag
- [x] [#2155](https://github.com/kubernetes/ingress-nginx/pull/2155) Update nginx with new modules
- [x] [#2162](https://github.com/kubernetes/ingress-nginx/pull/2162) Remove duplicated BuildConfigFromFlags func
- [x] [#2163](https://github.com/kubernetes/ingress-nginx/pull/2163) include lua-upstream-nginx-module in Nginx build
- [x] [#2164](https://github.com/kubernetes/ingress-nginx/pull/2164) use the correct error channel
- [x] [#2167](https://github.com/kubernetes/ingress-nginx/pull/2167) configuring load balancing per ingress
- [x] [#2172](https://github.com/kubernetes/ingress-nginx/pull/2172) include lua-resty-lock in nginx image
- [x] [#2174](https://github.com/kubernetes/ingress-nginx/pull/2174) Live Nginx configuration update without reloading
- [x] [#2180](https://github.com/kubernetes/ingress-nginx/pull/2180) Include tests in golint checks, fix warnings
- [x] [#2181](https://github.com/kubernetes/ingress-nginx/pull/2181) change nginx process pgid
- [x] [#2185](https://github.com/kubernetes/ingress-nginx/pull/2185) Remove ProxyPassParams setting
- [x] [#2191](https://github.com/kubernetes/ingress-nginx/pull/2191) Add checker test for bad pid
- [x] [#2193](https://github.com/kubernetes/ingress-nginx/pull/2193) fix wrong json tag
- [x] [#2201](https://github.com/kubernetes/ingress-nginx/pull/2201) Add worker-cpu-affinity nginx option
- [x] [#2202](https://github.com/kubernetes/ingress-nginx/pull/2202) Allow config to disable geoip
- [x] [#2205](https://github.com/kubernetes/ingress-nginx/pull/2205) add luacheck to lint lua files

_Documentation:_

- [x] [#2124](https://github.com/kubernetes/ingress-nginx/pull/2124) Document how to provide list types in configmap
- [x] [#2133](https://github.com/kubernetes/ingress-nginx/pull/2133) fix limit-req-status-code doc
- [x] [#2139](https://github.com/kubernetes/ingress-nginx/pull/2139) Update documentation for nginx-ingress-role RBAC.
- [x] [#2165](https://github.com/kubernetes/ingress-nginx/pull/2165) Typo fix "api server " -> "API server"
- [x] [#2169](https://github.com/kubernetes/ingress-nginx/pull/2169) Add documentation about secure-verify-ca-secret
- [x] [#2200](https://github.com/kubernetes/ingress-nginx/pull/2200) fix grammer mistake

### 0.11.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.11.0`

_New Features:_

- NGINX 1.13.9

_Changes:_

- [x] [#1992](https://github.com/kubernetes/ingress-nginx/pull/1992) Added configmap option to disable IPv6 in nginx DNS resolver
- [x] [#1993](https://github.com/kubernetes/ingress-nginx/pull/1993) Enable Customization of Auth Request Redirect
- [x] [#1996](https://github.com/kubernetes/ingress-nginx/pull/1996) Use v3/dev/performance of ModSecurity because of performance
- [x] [#1997](https://github.com/kubernetes/ingress-nginx/pull/1997) fix var checked
- [x] [#1998](https://github.com/kubernetes/ingress-nginx/pull/1998) Add support to enable/disable proxy buffering
- [x] [#1999](https://github.com/kubernetes/ingress-nginx/pull/1999) Add connection-proxy-header annotation
- [x] [#2001](https://github.com/kubernetes/ingress-nginx/pull/2001) Add limit-request-status-code option
- [x] [#2005](https://github.com/kubernetes/ingress-nginx/pull/2005) fix typo error for server name \_
- [x] [#2006](https://github.com/kubernetes/ingress-nginx/pull/2006) Add support for enabling ssl_ciphers per host
- [x] [#2019](https://github.com/kubernetes/ingress-nginx/pull/2019) Update nginx image
- [x] [#2021](https://github.com/kubernetes/ingress-nginx/pull/2021) Add nginx_cookie_flag_module
- [x] [#2026](https://github.com/kubernetes/ingress-nginx/pull/2026) update KUBERNETES from v1.8.0 to 1.9.0
- [x] [#2027](https://github.com/kubernetes/ingress-nginx/pull/2027) Show pod information in http-svc example
- [x] [#2030](https://github.com/kubernetes/ingress-nginx/pull/2030) do not ignore $http_host and $http_x_forwarded_host
- [x] [#2031](https://github.com/kubernetes/ingress-nginx/pull/2031) The maximum number of open file descriptors should be maxOpenFiles.
- [x] [#2036](https://github.com/kubernetes/ingress-nginx/pull/2036) add matchLabels in Deployment yaml, that both API extensions/v1beta1 …
- [x] [#2050](https://github.com/kubernetes/ingress-nginx/pull/2050) Get file max from fs/file-max.
- [x] [#2063](https://github.com/kubernetes/ingress-nginx/pull/2063) Run one test at a time
- [x] [#2065](https://github.com/kubernetes/ingress-nginx/pull/2065) Always return an IP address
- [x] [#2069](https://github.com/kubernetes/ingress-nginx/pull/2069) Do not cancel the synchronization of secrets
- [x] [#2071](https://github.com/kubernetes/ingress-nginx/pull/2071) Update Go to 1.9.4
- [x] [#2082](https://github.com/kubernetes/ingress-nginx/pull/2082) Use a ring channel to avoid blocking write of events
- [x] [#2089](https://github.com/kubernetes/ingress-nginx/pull/2089) Retry initial connection to the Kubernetes cluster
- [x] [#2093](https://github.com/kubernetes/ingress-nginx/pull/2093) Only pods in running phase are vallid for status
- [x] [#2099](https://github.com/kubernetes/ingress-nginx/pull/2099) Added GeoIP Organisational data
- [x] [#2107](https://github.com/kubernetes/ingress-nginx/pull/2107) Enabled the dynamic reload of GeoIP data
- [x] [#2119](https://github.com/kubernetes/ingress-nginx/pull/2119) Remove deprecated flag disable-node-list
- [x] [#2120](https://github.com/kubernetes/ingress-nginx/pull/2120) Migrate to codecov.io

_Documentation:_

- [x] [#1987](https://github.com/kubernetes/ingress-nginx/pull/1987) add kube-system namespace for oauth2-proxy example
- [x] [#1991](https://github.com/kubernetes/ingress-nginx/pull/1991) Add comment about bolean and number values
- [x] [#2009](https://github.com/kubernetes/ingress-nginx/pull/2009) docs/user-guide/tls: remove duplicated section
- [x] [#2011](https://github.com/kubernetes/ingress-nginx/pull/2011) broken link for sticky-ingress.yaml
- [x] [#2014](https://github.com/kubernetes/ingress-nginx/pull/2014) Add document for connection-proxy-header annotation
- [x] [#2016](https://github.com/kubernetes/ingress-nginx/pull/2016) Minor link fix in deployment docs
- [x] [#2018](https://github.com/kubernetes/ingress-nginx/pull/2018) Added documentation for Permanent Redirect
- [x] [#2035](https://github.com/kubernetes/ingress-nginx/pull/2035) fix broken links in static-ip readme
- [x] [#2038](https://github.com/kubernetes/ingress-nginx/pull/2038) fix typo: appropiate -> [appropriate]
- [x] [#2039](https://github.com/kubernetes/ingress-nginx/pull/2039) fix typo stickyness to stickiness
- [x] [#2040](https://github.com/kubernetes/ingress-nginx/pull/2040) fix wrong annotation
- [x] [#2041](https://github.com/kubernetes/ingress-nginx/pull/2041) fix spell error reslover -> resolver
- [x] [#2046](https://github.com/kubernetes/ingress-nginx/pull/2046) Fix typos
- [x] [#2054](https://github.com/kubernetes/ingress-nginx/pull/2054) Adding documentation for helm with RBAC enabled
- [x] [#2075](https://github.com/kubernetes/ingress-nginx/pull/2075) Fix opentracing configuration when multiple options are configured
- [x] [#2076](https://github.com/kubernetes/ingress-nginx/pull/2076) Fix spelling errors
- [x] [#2077](https://github.com/kubernetes/ingress-nginx/pull/2077) Remove initContainer from default deployment

### 0.10.2

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.10.2`

_Changes:_

- [x] [#1978](https://github.com/kubernetes/ingress-nginx/pull/1978) Fix chain completion and default certificate flag issues

### 0.10.1

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.10.1`

_Changes:_

- [x] [#1945](https://github.com/kubernetes/ingress-nginx/pull/1945) When a secret is updated read ingress annotations (again)
- [x] [#1948](https://github.com/kubernetes/ingress-nginx/pull/1948) Update go to 1.9.3
- [x] [#1953](https://github.com/kubernetes/ingress-nginx/pull/1953) Added annotation for upstream-vhost
- [x] [#1960](https://github.com/kubernetes/ingress-nginx/pull/1960) Adjust sysctl values to improve nginx performance
- [x] [#1963](https://github.com/kubernetes/ingress-nginx/pull/1963) Fix tests
- [x] [#1969](https://github.com/kubernetes/ingress-nginx/pull/1969) Rollback #1854
- [x] [#1970](https://github.com/kubernetes/ingress-nginx/pull/1970) By default brotli is disabled

### 0.10.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.10.0`

_Breaking changes:_

Changed the names of default Nginx ingress prometheus metrics.
If you are scraping default Nginx ingress metrics with prometheus the metrics changes are as follows:

```
nginx_active_connections_total          -> nginx_connections_total{state="active"}
nginx_accepted_connections_total        -> nginx_connections_total{state="accepted"}
nginx_handled_connections_total         -> nginx_connections_total{state="handled"}
nginx_current_reading_connections_total -> nginx_connections{state="reading"}
nginx_current_writing_connections_total -> nginx_connections{state="writing"}
current_waiting_connections_total       -> nginx_connections{state="waiting"}
```

_New Features:_

- NGINX 1.13.8
- Support to hide headers from upstream servers
- Support for Jaeger
- CORS max age annotation

_Changes:_

- [x] [#1782](https://github.com/kubernetes/ingress-nginx/pull/1782) auth-tls-pass-certificate-to-upstream should be bool
- [x] [#1787](https://github.com/kubernetes/ingress-nginx/pull/1787) force external_auth requests to http/1.1
- [x] [#1800](https://github.com/kubernetes/ingress-nginx/pull/1800) Add control of the configuration refresh interval
- [x] [#1805](https://github.com/kubernetes/ingress-nginx/pull/1805) Add X-Forwarded-Prefix on rewrites
- [x] [#1844](https://github.com/kubernetes/ingress-nginx/pull/1844) Validate x-forwarded-proto and connection scheme before redirect to https
- [x] [#1852](https://github.com/kubernetes/ingress-nginx/pull/1852) Update nginx to v1.13.8 and update modules
- [x] [#1854](https://github.com/kubernetes/ingress-nginx/pull/1854) Fix redirect to ssl
- [x] [#1858](https://github.com/kubernetes/ingress-nginx/pull/1858) When upstream-hash-by annotation is used do not configure a lb algorithm
- [x] [#1861](https://github.com/kubernetes/ingress-nginx/pull/1861) Improve speed of tests execution
- [x] [#1869](https://github.com/kubernetes/ingress-nginx/pull/1869) "proxy_redirect default" should be placed after the "proxy_pass"
- [x] [#1870](https://github.com/kubernetes/ingress-nginx/pull/1870) Fix SSL Passthrough template issue and custom ports in redirect to HTTPS
- [x] [#1871](https://github.com/kubernetes/ingress-nginx/pull/1871) Update nginx image to 0.31
- [x] [#1872](https://github.com/kubernetes/ingress-nginx/pull/1872) Fix data race updating ingress status
- [x] [#1880](https://github.com/kubernetes/ingress-nginx/pull/1880) Update go dependencies and cleanup deprecated packages
- [x] [#1888](https://github.com/kubernetes/ingress-nginx/pull/1888) Add CORS max age annotation
- [x] [#1891](https://github.com/kubernetes/ingress-nginx/pull/1891) Refactor initial synchronization of ingress objects
- [x] [#1903](https://github.com/kubernetes/ingress-nginx/pull/1903) If server_tokens is disabled remove the Server header
- [x] [#1906](https://github.com/kubernetes/ingress-nginx/pull/1906) Random string function should only contains letters
- [x] [#1907](https://github.com/kubernetes/ingress-nginx/pull/1907) Fix custom port in redirects
- [x] [#1909](https://github.com/kubernetes/ingress-nginx/pull/1909) Release nginx 0.32
- [x] [#1910](https://github.com/kubernetes/ingress-nginx/pull/1910) updating prometheus metrics names according to naming best practices
- [x] [#1912](https://github.com/kubernetes/ingress-nginx/pull/1912) removing \_total prefix from nginx guage metrics
- [x] [#1914](https://github.com/kubernetes/ingress-nginx/pull/1914) Add --with-http_secure_link_module for the Nginx build configuration
- [x] [#1916](https://github.com/kubernetes/ingress-nginx/pull/1916) Add support for jaeger backend
- [x] [#1918](https://github.com/kubernetes/ingress-nginx/pull/1918) Update nginx image to 0.32
- [x] [#1919](https://github.com/kubernetes/ingress-nginx/pull/1919) Add option for reuseport in nginx listen section
- [x] [#1926](https://github.com/kubernetes/ingress-nginx/pull/1926) Do not use port from host header
- [x] [#1927](https://github.com/kubernetes/ingress-nginx/pull/1927) Remove sendfile configuration
- [x] [#1928](https://github.com/kubernetes/ingress-nginx/pull/1928) Add support to hide headers from upstream servers
- [x] [#1929](https://github.com/kubernetes/ingress-nginx/pull/1929) Refactoring of kubernetes informers and local caches
- [x] [#1933](https://github.com/kubernetes/ingress-nginx/pull/1933) Remove deploy of ingress controller from the example

_Documentation:_

- [x] [#1786](https://github.com/kubernetes/ingress-nginx/pull/1786) fix: some typo.
- [x] [#1792](https://github.com/kubernetes/ingress-nginx/pull/1792) Add note about annotation values
- [x] [#1814](https://github.com/kubernetes/ingress-nginx/pull/1814) Fix link to custom configuration
- [x] [#1826](https://github.com/kubernetes/ingress-nginx/pull/1826) Add note about websocket and load balancers
- [x] [#1840](https://github.com/kubernetes/ingress-nginx/pull/1840) Add note about default log files
- [x] [#1853](https://github.com/kubernetes/ingress-nginx/pull/1853) Clarify docs for add-headers and proxy-set-headers
- [x] [#1864](https://github.com/kubernetes/ingress-nginx/pull/1864) configmap.md: Convert hyphens in name column to non-breaking-hyphens
- [x] [#1865](https://github.com/kubernetes/ingress-nginx/pull/1865) Add docs for legacy TLS version and ciphers
- [x] [#1867](https://github.com/kubernetes/ingress-nginx/pull/1867) Fix publish-service patch and update README
- [x] [#1913](https://github.com/kubernetes/ingress-nginx/pull/1913) Missing r
- [x] [#1925](https://github.com/kubernetes/ingress-nginx/pull/1925) Fix doc links

### 0.9.0

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.9.0`

_Changes:_

- [x] [#1731](https://github.com/kubernetes/ingress-nginx/pull/1731) Allow configuration of proxy_responses value for tcp/udp configmaps
- [x] [#1766](https://github.com/kubernetes/ingress-nginx/pull/1766) Fix ingress typo
- [x] [#1768](https://github.com/kubernetes/ingress-nginx/pull/1768) Custom default backend must use annotations if present
- [x] [#1769](https://github.com/kubernetes/ingress-nginx/pull/1769) Use custom https port in redirects
- [x] [#1771](https://github.com/kubernetes/ingress-nginx/pull/1771) Add additional check for old SSL certificates
- [x] [#1776](https://github.com/kubernetes/ingress-nginx/pull/1776) Add option to configure the redirect code

### 0.9-beta.19

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.9.0-beta.19`

_Changes:_

- Fix regression with ingress.class annotation introduced in 0.9-beta.18

### 0.9-beta.18

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.9.0-beta.18`

_Breaking changes:_

- The NGINX ingress annotations contains a new prefix: **nginx.ingress.kubernetes.io**. This change is behind a flag to avoid breaking running deployments.
  To avoid breaking a running NGINX ingress controller add the flag **--annotations-prefix=ingress.kubernetes.io** to the nginx ingress controller deployment.
  There is one exception, the annotation `kubernetes.io/ingress.class` remains unchanged (this annotation is used in multiple ingress controllers)

_New Features:_

- NGINX 1.13.7
- Support for s390x
- e2e tests

_Changes:_

- [x] [#1648](https://github.com/kubernetes/ingress-nginx/pull/1648) Remove GenericController and add tests
- [x] [#1650](https://github.com/kubernetes/ingress-nginx/pull/1650) Fix misspell errors
- [x] [#1651](https://github.com/kubernetes/ingress-nginx/pull/1651) Remove node lister
- [x] [#1652](https://github.com/kubernetes/ingress-nginx/pull/1652) Remove node lister
- [x] [#1653](https://github.com/kubernetes/ingress-nginx/pull/1653) Fix diff execution
- [x] [#1654](https://github.com/kubernetes/ingress-nginx/pull/1654) Fix travis script and update kubernetes to 1.8.0
- [x] [#1658](https://github.com/kubernetes/ingress-nginx/pull/1658) Tests
- [x] [#1659](https://github.com/kubernetes/ingress-nginx/pull/1659) Add nginx helper tests
- [x] [#1662](https://github.com/kubernetes/ingress-nginx/pull/1662) Refactor annotations
- [x] [#1665](https://github.com/kubernetes/ingress-nginx/pull/1665) Add the original http request method to the auth request
- [x] [#1687](https://github.com/kubernetes/ingress-nginx/pull/1687) Fix use merge of annotations
- [x] [#1689](https://github.com/kubernetes/ingress-nginx/pull/1689) Enable s390x
- [x] [#1693](https://github.com/kubernetes/ingress-nginx/pull/1693) Fix docker build
- [x] [#1695](https://github.com/kubernetes/ingress-nginx/pull/1695) Update nginx to v0.29
- [x] [#1696](https://github.com/kubernetes/ingress-nginx/pull/1696) Always add cors headers when enabled
- [x] [#1697](https://github.com/kubernetes/ingress-nginx/pull/1697) Disable features not availables in some platforms
- [x] [#1698](https://github.com/kubernetes/ingress-nginx/pull/1698) Auth e2e tests
- [x] [#1699](https://github.com/kubernetes/ingress-nginx/pull/1699) Refactor SSL intermediate CA certificate check
- [x] [#1700](https://github.com/kubernetes/ingress-nginx/pull/1700) Add patch command to append publish-service flag
- [x] [#1701](https://github.com/kubernetes/ingress-nginx/pull/1701) fix: Core() is deprecated use CoreV1() instead.
- [x] [#1702](https://github.com/kubernetes/ingress-nginx/pull/1702) Fix TLS example [ci skip]
- [x] [#1704](https://github.com/kubernetes/ingress-nginx/pull/1704) Add e2e tests to verify the correct source IP address
- [x] [#1705](https://github.com/kubernetes/ingress-nginx/pull/1705) Add annotation for setting proxy_redirect
- [x] [#1706](https://github.com/kubernetes/ingress-nginx/pull/1706) Increase ELB idle timeouts [ci skip]
- [x] [#1710](https://github.com/kubernetes/ingress-nginx/pull/1710) Do not update a secret not referenced by ingress rules
- [x] [#1713](https://github.com/kubernetes/ingress-nginx/pull/1713) add --report-node-internal-ip-address describe to cli-arguments.md
- [x] [#1717](https://github.com/kubernetes/ingress-nginx/pull/1717) Fix command used to detect version
- [x] [#1720](https://github.com/kubernetes/ingress-nginx/pull/1720) Add docker-registry example [ci skip]
- [x] [#1722](https://github.com/kubernetes/ingress-nginx/pull/1722) Add annotation to enable passing the certificate to the upstream server
- [x] [#1723](https://github.com/kubernetes/ingress-nginx/pull/1723) Add timeouts to http server and additional pprof routes
- [x] [#1724](https://github.com/kubernetes/ingress-nginx/pull/1724) Cleanup main
- [x] [#1725](https://github.com/kubernetes/ingress-nginx/pull/1725) Enable all e2e tests
- [x] [#1726](https://github.com/kubernetes/ingress-nginx/pull/1726) fix: replace deprecated methods.
- [x] [#1734](https://github.com/kubernetes/ingress-nginx/pull/1734) Changes ssl-client-cert header
- [x] [#1737](https://github.com/kubernetes/ingress-nginx/pull/1737) Update nginx v1.13.7
- [x] [#1738](https://github.com/kubernetes/ingress-nginx/pull/1738) Cleanup
- [x] [#1739](https://github.com/kubernetes/ingress-nginx/pull/1739) Improve e2e checks
- [x] [#1740](https://github.com/kubernetes/ingress-nginx/pull/1740) Update nginx
- [x] [#1745](https://github.com/kubernetes/ingress-nginx/pull/1745) Simplify annotations
- [x] [#1746](https://github.com/kubernetes/ingress-nginx/pull/1746) Cleanup of e2e helpers

_Documentation:_

- [x] [#1657](https://github.com/kubernetes/ingress-nginx/pull/1657) Add better documentation for deploying for dev
- [x] [#1680](https://github.com/kubernetes/ingress-nginx/pull/1680) Add doc for log-format-escape-json [ci skip]
- [x] [#1685](https://github.com/kubernetes/ingress-nginx/pull/1685) Fix default SSL certificate flag docs [ci skip]
- [x] [#1686](https://github.com/kubernetes/ingress-nginx/pull/1686) Fix development doc [ci skip]
- [x] [#1727](https://github.com/kubernetes/ingress-nginx/pull/1727) fix: fix typos in docs.
- [x] [#1747](https://github.com/kubernetes/ingress-nginx/pull/1747) Add config-map usage and options to Documentation

### 0.9-beta.17

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.9.0-beta.17`

_Changes:_

- Fix regression with annotations introduced in 0.9-beta.16 (thanks @tomlanyon)

### 0.9-beta.16

**Image:** `quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.9.0-beta.16`

_New Features:_

- Images are published to [quay.io](https://quay.io/repository/kubernetes-ingress-controller)
- NGINX 1.13.6
- OpenTracing Jaeger support inNGINX
- [ModSecurity support](https://github.com/SpiderLabs/ModSecurity-nginx)
- Support for [brotli compression in NGINX](https://certsimple.com/blog/nginx-brotli)
- Return 503 error instead of 404 when no endpoint is available

_Breaking changes:_

- The default SSL configuration was updated to use `TLSv1.2` and the default cipher list is `ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256`

_Known issues:_

- When ModSecurity is enabled a segfault could occur - [ModSecurity#1590](https://github.com/SpiderLabs/ModSecurity/issues/1590)

_Changes:_

- [x] [#1489](https://github.com/kubernetes/ingress-nginx/pull/1489) Compute a real `X-Forwarded-For` header
- [x] [#1490](https://github.com/kubernetes/ingress-nginx/pull/1490) Introduce an upstream-hash-by annotation to support consistent hashing by nginx variable or text
- [x] [#1498](https://github.com/kubernetes/ingress-nginx/pull/1498) Add modsecurity module
- [x] [#1500](https://github.com/kubernetes/ingress-nginx/pull/1500) Enable modsecurity feature
- [x] [#1501](https://github.com/kubernetes/ingress-nginx/pull/1501) Request ingress controller version in issue template
- [x] [#1502](https://github.com/kubernetes/ingress-nginx/pull/1502) Force reload on template change
- [x] [#1503](https://github.com/kubernetes/ingress-nginx/pull/1503) Add falg to report node internal IP address in ingress status
- [x] [#1505](https://github.com/kubernetes/ingress-nginx/pull/1505) Increase size of variable hash bucket
- [x] [#1506](https://github.com/kubernetes/ingress-nginx/pull/1506) Update nginx ssl configuration
- [x] [#1507](https://github.com/kubernetes/ingress-nginx/pull/1507) Add tls session ticket key setting
- [x] [#1511](https://github.com/kubernetes/ingress-nginx/pull/1511) fix deprecated ssl_client_cert. add ssl_client_verify header
- [x] [#1513](https://github.com/kubernetes/ingress-nginx/pull/1513) Return 503 by default when no endpoint is available
- [x] [#1520](https://github.com/kubernetes/ingress-nginx/pull/1520) Change alias behaviour not to create new server section needlessly
- [x] [#1523](https://github.com/kubernetes/ingress-nginx/pull/1523) Include the serversnippet from the config map in server blocks
- [x] [#1533](https://github.com/kubernetes/ingress-nginx/pull/1533) Remove authentication send body annotation
- [x] [#1535](https://github.com/kubernetes/ingress-nginx/pull/1535) Remove auth-send-body [ci skip]
- [x] [#1538](https://github.com/kubernetes/ingress-nginx/pull/1538) Rename service-nodeport.yml to service-nodeport.yaml
- [x] [#1543](https://github.com/kubernetes/ingress-nginx/pull/1543) Fix glog initialization error
- [x] [#1544](https://github.com/kubernetes/ingress-nginx/pull/1544) Fix `make container` for OSX.
- [x] [#1547](https://github.com/kubernetes/ingress-nginx/pull/1547) fix broken GCE-GKE service descriptor
- [x] [#1550](https://github.com/kubernetes/ingress-nginx/pull/1550) Add e2e tests - default backend
- [x] [#1553](https://github.com/kubernetes/ingress-nginx/pull/1553) Cors features improvements
- [x] [#1554](https://github.com/kubernetes/ingress-nginx/pull/1554) Add missing unit test for nextPowerOf2 function
- [x] [#1556](https://github.com/kubernetes/ingress-nginx/pull/1556) fixed https port forwarding in Azure LB service
- [x] [#1566](https://github.com/kubernetes/ingress-nginx/pull/1566) Release nginx-slim 0.27
- [x] [#1568](https://github.com/kubernetes/ingress-nginx/pull/1568) update defaultbackend tag
- [x] [#1569](https://github.com/kubernetes/ingress-nginx/pull/1569) Update 404 server image
- [x] [#1570](https://github.com/kubernetes/ingress-nginx/pull/1570) Update nginx version
- [x] [#1571](https://github.com/kubernetes/ingress-nginx/pull/1571) Fix cors tests
- [x] [#1572](https://github.com/kubernetes/ingress-nginx/pull/1572) Certificate Auth Bugfix
- [x] [#1577](https://github.com/kubernetes/ingress-nginx/pull/1577) Do not use relative urls for yaml files
- [x] [#1580](https://github.com/kubernetes/ingress-nginx/pull/1580) Upgrade to use the latest version of nginx-opentracing.
- [x] [#1581](https://github.com/kubernetes/ingress-nginx/pull/1581) Fix Makefile to work in OSX.
- [x] [#1582](https://github.com/kubernetes/ingress-nginx/pull/1582) Add scripts to release from travis-ci
- [x] [#1584](https://github.com/kubernetes/ingress-nginx/pull/1584) Add missing probes in deployments
- [x] [#1585](https://github.com/kubernetes/ingress-nginx/pull/1585) Add version flag
- [x] [#1587](https://github.com/kubernetes/ingress-nginx/pull/1587) Use pass access scheme in signin url
- [x] [#1589](https://github.com/kubernetes/ingress-nginx/pull/1589) Fix upstream vhost Equal comparison
- [x] [#1590](https://github.com/kubernetes/ingress-nginx/pull/1590) Fix Equals Comparison for CORS annotation
- [x] [#1592](https://github.com/kubernetes/ingress-nginx/pull/1592) Update opentracing module and release image to quay.io
- [x] [#1593](https://github.com/kubernetes/ingress-nginx/pull/1593) Fix makefile default task
- [x] [#1605](https://github.com/kubernetes/ingress-nginx/pull/1605) Fix ExternalName services
- [x] [#1607](https://github.com/kubernetes/ingress-nginx/pull/1607) Add support for named ports with service-upstream. #1459
- [x] [#1608](https://github.com/kubernetes/ingress-nginx/pull/1608) Fix issue with clusterIP detection on service upstream. #1534
- [x] [#1610](https://github.com/kubernetes/ingress-nginx/pull/1610) Only set alias if not already set
- [x] [#1618](https://github.com/kubernetes/ingress-nginx/pull/1618) Fix full XFF with PROXY
- [x] [#1620](https://github.com/kubernetes/ingress-nginx/pull/1620) Add gzip_vary
- [x] [#1621](https://github.com/kubernetes/ingress-nginx/pull/1621) Fix path to ELB listener image
- [x] [#1627](https://github.com/kubernetes/ingress-nginx/pull/1627) Add brotli support
- [x] [#1629](https://github.com/kubernetes/ingress-nginx/pull/1629) Add ssl-client-dn header
- [x] [#1632](https://github.com/kubernetes/ingress-nginx/pull/1632) Rename OWNERS assignees: to approvers:
- [x] [#1635](https://github.com/kubernetes/ingress-nginx/pull/1635) Install dumb-init using apt-get
- [x] [#1636](https://github.com/kubernetes/ingress-nginx/pull/1636) Update go to 1.9.2
- [x] [#1640](https://github.com/kubernetes/ingress-nginx/pull/1640) Update nginx to 0.28 and enable brotli

_Documentation:_

- [x] [#1491](https://github.com/kubernetes/ingress-nginx/pull/1491) Note that GCE has moved to a new repo
- [x] [#1492](https://github.com/kubernetes/ingress-nginx/pull/1492) Cleanup readme.md
- [x] [#1494](https://github.com/kubernetes/ingress-nginx/pull/1494) Cleanup
- [x] [#1497](https://github.com/kubernetes/ingress-nginx/pull/1497) Cleanup examples directory
- [x] [#1504](https://github.com/kubernetes/ingress-nginx/pull/1504) Clean readme
- [x] [#1508](https://github.com/kubernetes/ingress-nginx/pull/1508) Fixed link in prometheus example
- [x] [#1527](https://github.com/kubernetes/ingress-nginx/pull/1527) Split documentation
- [x] [#1536](https://github.com/kubernetes/ingress-nginx/pull/1536) Update documentation and examples [ci skip]
- [x] [#1541](https://github.com/kubernetes/ingress-nginx/pull/1541) fix(documentation): Fix some typos
- [x] [#1548](https://github.com/kubernetes/ingress-nginx/pull/1548) link to prometheus docs
- [x] [#1562](https://github.com/kubernetes/ingress-nginx/pull/1562) Fix development guide link
- [x] [#1563](https://github.com/kubernetes/ingress-nginx/pull/1563) Add task to verify markdown links
- [x] [#1583](https://github.com/kubernetes/ingress-nginx/pull/1583) Add note for certificate authentication in Cloudflare
- [x] [#1617](https://github.com/kubernetes/ingress-nginx/pull/1617) fix typo in user-guide/annotations.md

### 0.9-beta.15

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.15`

_New Features:_

- Add OCSP support
- Configurable ssl_verify_client

_Changes:_

- [x] [#1468](https://github.com/kubernetes/ingress/pull/1468) Add the original URL to the auth request
- [x] [#1469](https://github.com/kubernetes/ingress/pull/1469) Typo: Add missing {{ }}
- [x] [#1472](https://github.com/kubernetes/ingress/pull/1472) Fix X-Auth-Request-Redirect value to reflect the request uri
- [x] [#1473](https://github.com/kubernetes/ingress/pull/1473) Fix proxy protocol check
- [x] [#1475](https://github.com/kubernetes/ingress/pull/1475) Add OCSP support
- [x] [#1477](https://github.com/kubernetes/ingress/pull/1477) Fix semicolons in global configuration
- [x] [#1478](https://github.com/kubernetes/ingress/pull/1478) Pass redirect field in login page to get a proper redirect
- [x] [#1480](https://github.com/kubernetes/ingress/pull/1480) configurable ssl_verify_client
- [x] [#1485](https://github.com/kubernetes/ingress/pull/1485) Fix source IP address
- [x] [#1486](https://github.com/kubernetes/ingress/pull/1486) Fix overwrite of custom configuration

_Documentation:_

- [x] [#1460](https://github.com/kubernetes/ingress/pull/1460) Expose UDP port in UDP ingress example
- [x] [#1465](https://github.com/kubernetes/ingress/pull/1465) review prometheus docs

### 0.9-beta.14

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.14`

_New Features:_

- Opentracing support for NGINX
- Setting upstream vhost for nginx
- Allow custom global configuration at multiple levels
- Add support for proxy protocol decoding and encoding in TCP services

_Changes:_

- [x] [#719](https://github.com/kubernetes/ingress/pull/719) Setting upstream vhost for nginx.
- [x] [#1321](https://github.com/kubernetes/ingress/pull/1321) Enable keepalive in upstreams
- [x] [#1322](https://github.com/kubernetes/ingress/pull/1322) parse real ip
- [x] [#1323](https://github.com/kubernetes/ingress/pull/1323) use $the_real_ip for rate limit whitelist
- [x] [#1326](https://github.com/kubernetes/ingress/pull/1326) Pass headers from the custom error backend
- [x] [#1328](https://github.com/kubernetes/ingress/pull/1328) update deprecated interface
- [x] [#1329](https://github.com/kubernetes/ingress/pull/1329) add example for nginx-ingress
- [x] [#1330](https://github.com/kubernetes/ingress/pull/1330) Increase coverage in template.go for nginx controller
- [x] [#1335](https://github.com/kubernetes/ingress/pull/1335) Configurable proxy_request_buffering per location..
- [x] [#1338](https://github.com/kubernetes/ingress/pull/1338) Fix multiple leader election
- [x] [#1339](https://github.com/kubernetes/ingress/pull/1339) Enable status port listening in all interfaces
- [x] [#1340](https://github.com/kubernetes/ingress/pull/1340) Update sha256sum of nginx substitutions
- [x] [#1341](https://github.com/kubernetes/ingress/pull/1341) Fix typos
- [x] [#1345](https://github.com/kubernetes/ingress/pull/1345) refactor controllers.go
- [x] [#1349](https://github.com/kubernetes/ingress/pull/1349) Force reload if a secret is updated
- [x] [#1363](https://github.com/kubernetes/ingress/pull/1363) Fix proxy request buffering default configuration
- [x] [#1365](https://github.com/kubernetes/ingress/pull/1365) Fix equals comparsion returing False if both objects have nil Targets or Services.
- [x] [#1367](https://github.com/kubernetes/ingress/pull/1367) Fix typos
- [x] [#1379](https://github.com/kubernetes/ingress/pull/1379) Fix catch all upstream server
- [x] [#1380](https://github.com/kubernetes/ingress/pull/1380) Cleanup
- [x] [#1381](https://github.com/kubernetes/ingress/pull/1381) Refactor X-Forwarded-\* headers
- [x] [#1382](https://github.com/kubernetes/ingress/pull/1382) Cleanup
- [x] [#1387](https://github.com/kubernetes/ingress/pull/1387) Improve resource usage in nginx controller
- [x] [#1392](https://github.com/kubernetes/ingress/pull/1392) Avoid issues with goroutines updating fields
- [x] [#1393](https://github.com/kubernetes/ingress/pull/1393) Limit the number of goroutines used for the update of ingress status
- [x] [#1394](https://github.com/kubernetes/ingress/pull/1394) Improve equals
- [x] [#1402](https://github.com/kubernetes/ingress/pull/1402) fix error when cert or key is nil
- [x] [#1403](https://github.com/kubernetes/ingress/pull/1403) Added tls ports to rbac nginx ingress controller and service
- [x] [#1404](https://github.com/kubernetes/ingress/pull/1404) Use nginx default value for SSLECDHCurve
- [x] [#1411](https://github.com/kubernetes/ingress/pull/1411) Add more descriptive logging in certificate loading
- [x] [#1412](https://github.com/kubernetes/ingress/pull/1412) Correct Error Handling to avoid panics and add more logging to template
- [x] [#1413](https://github.com/kubernetes/ingress/pull/1413) Validate external names
- [x] [#1418](https://github.com/kubernetes/ingress/pull/1418) Fix links after design proposals move
- [x] [#1419](https://github.com/kubernetes/ingress/pull/1419) Remove duplicated ingress check code
- [x] [#1420](https://github.com/kubernetes/ingress/pull/1420) Process queue items by time window
- [x] [#1423](https://github.com/kubernetes/ingress/pull/1423) Fix cast error
- [x] [#1424](https://github.com/kubernetes/ingress/pull/1424) Allow overriding the tag and registry
- [x] [#1426](https://github.com/kubernetes/ingress/pull/1426) Enhance Certificate Logging and Clearup Mutual Auth Docs
- [x] [#1430](https://github.com/kubernetes/ingress/pull/1430) Add support for proxy protocol decoding and encoding in TCP services
- [x] [#1434](https://github.com/kubernetes/ingress/pull/1434) Fix exec of readSecrets
- [x] [#1435](https://github.com/kubernetes/ingress/pull/1435) Add header to upstream server for external authentication
- [x] [#1438](https://github.com/kubernetes/ingress/pull/1438) Do not intercept errors from the custom error service
- [x] [#1439](https://github.com/kubernetes/ingress/pull/1439) Nginx master process killed thus no further reloads
- [x] [#1440](https://github.com/kubernetes/ingress/pull/1440) Kill worker processes to allow the restart of nginx
- [x] [#1445](https://github.com/kubernetes/ingress/pull/1445) Updated godeps
- [x] [#1450](https://github.com/kubernetes/ingress/pull/1450) Fix links
- [x] [#1451](https://github.com/kubernetes/ingress/pull/1451) Add example of server-snippet
- [x] [#1452](https://github.com/kubernetes/ingress/pull/1452) Fix sync of secrets (kube lego)
- [x] [#1454](https://github.com/kubernetes/ingress/pull/1454) Allow custom global configuration at multiple levels

_Documentation:_

- [x] [#1400](https://github.com/kubernetes/ingress/pull/1400) Fix ConfigMap link in doc
- [x] [#1422](https://github.com/kubernetes/ingress/pull/1422) Add docs for opentracing
- [x] [#1441](https://github.com/kubernetes/ingress/pull/1441) Improve custom error pages doc
- [x] [#1442](https://github.com/kubernetes/ingress/pull/1442) Opentracing docs
- [x] [#1446](https://github.com/kubernetes/ingress/pull/1446) Add custom timeout annotations doc

### 0.9-beta.13

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.13`

_New Features:_

- NGINX 1.3.5
- New flag to disable node listing
- Custom X-Forwarder-Header (CloudFlare uses `CF-Connecting-IP` as header)
- Custom error page in Client Certificate Authentication

_Changes:_

- [x] [#1272](https://github.com/kubernetes/ingress/pull/1272) Delete useless statement
- [x] [#1277](https://github.com/kubernetes/ingress/pull/1277) Add indent for nginx.conf
- [x] [#1278](https://github.com/kubernetes/ingress/pull/1278) Add proxy-pass-params annotation and Backend field
- [x] [#1282](https://github.com/kubernetes/ingress/pull/1282) Fix nginx stats
- [x] [#1288](https://github.com/kubernetes/ingress/pull/1288) Allow PATCH in enable-cors
- [x] [#1290](https://github.com/kubernetes/ingress/pull/1290) Add flag to disabling node listing
- [x] [#1293](https://github.com/kubernetes/ingress/pull/1293) Adds support for error page in Client Certificate Authentication
- [x] [#1308](https://github.com/kubernetes/ingress/pull/1308) A trivial typo in config
- [x] [#1310](https://github.com/kubernetes/ingress/pull/1310) Refactoring nginx configuration configmap
- [x] [#1311](https://github.com/kubernetes/ingress/pull/1311) Enable nginx async writes
- [x] [#1312](https://github.com/kubernetes/ingress/pull/1312) Allow custom forwarded for header
- [x] [#1313](https://github.com/kubernetes/ingress/pull/1313) Fix eol in nginx template
- [x] [#1315](https://github.com/kubernetes/ingress/pull/1315) Fix nginx custom error pages

_Documentation:_

- [x] [#1270](https://github.com/kubernetes/ingress/pull/1270) add missing yamls in controllers/nginx
- [x] [#1276](https://github.com/kubernetes/ingress/pull/1276) Link rbac sample from deployment docs
- [x] [#1291](https://github.com/kubernetes/ingress/pull/1291) fix link to conformance suite
- [x] [#1295](https://github.com/kubernetes/ingress/pull/1295) fix README of nginx-ingress-controller
- [x] [#1299](https://github.com/kubernetes/ingress/pull/1299) fix two doc issues in nginx/README
- [x] [#1306](https://github.com/kubernetes/ingress/pull/1306) Fix kubeconfig example for nginx deployment

### 0.9-beta.12

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.12`

_Breaking changes:_

- SSL passthrough is disabled by default. To enable the feature use `--enable-ssl-passthrough`

_New Features:_

- Support for arm64
- New flags to customize listen ports
- Per minute rate limiting
- Rate limit whitelist
- Configuration of nginx worker timeout (to avoid zombie nginx workers processes)
- Redirects from non-www to www
- Custom default backend (per Ingress)
- Graceful shutdown for NGINX

_Changes:_

- [x] [#977](https://github.com/kubernetes/ingress/pull/977) Add sort-backends command line option
- [x] [#981](https://github.com/kubernetes/ingress/pull/981) Add annotation to allow use of service ClusterIP for NGINX upstream.
- [x] [#991](https://github.com/kubernetes/ingress/pull/991) Remove secret sync loop
- [x] [#992](https://github.com/kubernetes/ingress/pull/992) Check errors generating pem files
- [x] [#993](https://github.com/kubernetes/ingress/pull/993) Fix the sed command to work on macOS
- [x] [#1013](https://github.com/kubernetes/ingress/pull/1013) The fields of vtsDate are unified in the form of plural
- [x] [#1025](https://github.com/kubernetes/ingress/pull/1025) Fix file watch
- [x] [#1027](https://github.com/kubernetes/ingress/pull/1027) Lint code
- [x] [#1031](https://github.com/kubernetes/ingress/pull/1031) Change missing secret name log level to V(3)
- [x] [#1032](https://github.com/kubernetes/ingress/pull/1032) Alternative syncSecret approach #1030
- [x] [#1042](https://github.com/kubernetes/ingress/pull/1042) Add function to allow custom values in Ingress status
- [x] [#1043](https://github.com/kubernetes/ingress/pull/1043) Return reference to object providing Endpoint
- [x] [#1046](https://github.com/kubernetes/ingress/pull/1046) Add field FileSHA in BasicDigest struct
- [x] [#1058](https://github.com/kubernetes/ingress/pull/1058) add per minute rate limiting
- [x] [#1060](https://github.com/kubernetes/ingress/pull/1060) Update fsnotify dependency to fix arm64 issue
- [x] [#1065](https://github.com/kubernetes/ingress/pull/1065) Add more descriptive steps in Dev Documentation
- [x] [#1073](https://github.com/kubernetes/ingress/pull/1073) Release nginx-slim 0.22
- [x] [#1074](https://github.com/kubernetes/ingress/pull/1074) Remove lua and use fastcgi to render errors
- [x] [#1075](https://github.com/kubernetes/ingress/pull/1075) (feat/ #374) support proxy timeout
- [x] [#1076](https://github.com/kubernetes/ingress/pull/1076) Add more ssl test cases
- [x] [#1078](https://github.com/kubernetes/ingress/pull/1078) fix the same udp port and tcp port, update nginx.conf error
- [x] [#1080](https://github.com/kubernetes/ingress/pull/1080) Disable platform s390x
- [x] [#1081](https://github.com/kubernetes/ingress/pull/1081) Spit Static check and Coverage in diff Stages of Travis CI
- [x] [#1082](https://github.com/kubernetes/ingress/pull/1082) Fix build tasks
- [x] [#1087](https://github.com/kubernetes/ingress/pull/1087) Release nginx-slim 0.23
- [x] [#1088](https://github.com/kubernetes/ingress/pull/1088) Configure nginx worker timeout
- [x] [#1089](https://github.com/kubernetes/ingress/pull/1089) Update nginx to 1.13.4
- [x] [#1098](https://github.com/kubernetes/ingress/pull/1098) Exposing the event recorder to allow other controllers to create events
- [x] [#1102](https://github.com/kubernetes/ingress/pull/1102) Fix lose SSL Passthrough
- [x] [#1104](https://github.com/kubernetes/ingress/pull/1104) Simplify verification of hostname in ssl certificates
- [x] [#1109](https://github.com/kubernetes/ingress/pull/1109) Cleanup remote address in nginx template
- [x] [#1110](https://github.com/kubernetes/ingress/pull/1110) Fix Endpoint comparison
- [x] [#1118](https://github.com/kubernetes/ingress/pull/1118) feat(#733)Support nginx bandwidth control
- [x] [#1124](https://github.com/kubernetes/ingress/pull/1124) check fields len in dns.go
- [x] [#1130](https://github.com/kubernetes/ingress/pull/1130) Update nginx.go
- [x] [#1134](https://github.com/kubernetes/ingress/pull/1134) replace deprecated interface with versioned ones
- [x] [#1136](https://github.com/kubernetes/ingress/pull/1136) Fix status update - changed in #1074
- [x] [#1138](https://github.com/kubernetes/ingress/pull/1138) update nginx.go: performance improve
- [x] [#1139](https://github.com/kubernetes/ingress/pull/1139) Fix Todo:convert sequence to table
- [x] [#1162](https://github.com/kubernetes/ingress/pull/1162) Optimize CI build time
- [x] [#1164](https://github.com/kubernetes/ingress/pull/1164) Use variable request_uri as redirect after auth
- [x] [#1179](https://github.com/kubernetes/ingress/pull/1179) Fix sticky upstream not used when enable rewrite
- [x] [#1184](https://github.com/kubernetes/ingress/pull/1184) Add support for temporal and permanent redirects
- [x] [#1185](https://github.com/kubernetes/ingress/pull/1185) Add more info about Server-Alias usage
- [x] [#1186](https://github.com/kubernetes/ingress/pull/1186) Add annotation for client-body-buffer-size per location
- [x] [#1190](https://github.com/kubernetes/ingress/pull/1190) Add flag to disable SSL passthrough
- [x] [#1193](https://github.com/kubernetes/ingress/pull/1193) fix broken link
- [x] [#1198](https://github.com/kubernetes/ingress/pull/1198) Add option for specific scheme for base url
- [x] [#1202](https://github.com/kubernetes/ingress/pull/1202) formatIP issue
- [x] [#1203](https://github.com/kubernetes/ingress/pull/1203) NGINX not reloading correctly
- [x] [#1204](https://github.com/kubernetes/ingress/pull/1204) Fix template error
- [x] [#1205](https://github.com/kubernetes/ingress/pull/1205) Add initial sync of secrets
- [x] [#1206](https://github.com/kubernetes/ingress/pull/1206) Update ssl-passthrough docs
- [x] [#1207](https://github.com/kubernetes/ingress/pull/1207) delete broken link
- [x] [#1208](https://github.com/kubernetes/ingress/pull/1208) fix some typo
- [x] [#1210](https://github.com/kubernetes/ingress/pull/1210) add rate limit whitelist
- [x] [#1215](https://github.com/kubernetes/ingress/pull/1215) Replace base64 encoding with random uuid
- [x] [#1218](https://github.com/kubernetes/ingress/pull/1218) Trivial fixes in core/pkg/net
- [x] [#1219](https://github.com/kubernetes/ingress/pull/1219) keep zones unique per ingress resource
- [x] [#1221](https://github.com/kubernetes/ingress/pull/1221) Move certificate authentication from location to server
- [x] [#1223](https://github.com/kubernetes/ingress/pull/1223) Add doc for non-www to www annotation
- [x] [#1224](https://github.com/kubernetes/ingress/pull/1224) refactor rate limit whitelist
- [x] [#1226](https://github.com/kubernetes/ingress/pull/1226) Remove useless variable in nginx.tmpl
- [x] [#1227](https://github.com/kubernetes/ingress/pull/1227) Update annotations doc with base-url-scheme
- [x] [#1233](https://github.com/kubernetes/ingress/pull/1233) Fix ClientBodyBufferSize annotation
- [x] [#1234](https://github.com/kubernetes/ingress/pull/1234) Lint code
- [x] [#1235](https://github.com/kubernetes/ingress/pull/1235) Fix Equal comparison
- [x] [#1236](https://github.com/kubernetes/ingress/pull/1236) Add Validation for Client Body Buffer Size
- [x] [#1238](https://github.com/kubernetes/ingress/pull/1238) Add support for 'client_body_timeout' and 'client_header_timeout'
- [x] [#1239](https://github.com/kubernetes/ingress/pull/1239) Add flags to customize listen ports and detect port collisions
- [x] [#1243](https://github.com/kubernetes/ingress/pull/1243) Add support for access-log-path and error-log-path
- [x] [#1244](https://github.com/kubernetes/ingress/pull/1244) Add custom default backend annotation
- [x] [#1246](https://github.com/kubernetes/ingress/pull/1246) Add additional headers when custom default backend is used
- [x] [#1247](https://github.com/kubernetes/ingress/pull/1247) Make Ingress annotations available in template
- [x] [#1248](https://github.com/kubernetes/ingress/pull/1248) Improve nginx controller performance
- [x] [#1254](https://github.com/kubernetes/ingress/pull/1254) fix Type transform panic
- [x] [#1257](https://github.com/kubernetes/ingress/pull/1257) Graceful shutdown for Nginx
- [x] [#1261](https://github.com/kubernetes/ingress/pull/1261) Add support for 'worker-shutdown-timeout'

_Documentation:_

- [x] [#976](https://github.com/kubernetes/ingress/pull/976) Update annotations doc
- [x] [#979](https://github.com/kubernetes/ingress/pull/979) Missing auth example
- [x] [#980](https://github.com/kubernetes/ingress/pull/980) Add nginx basic auth example
- [x] [#1001](https://github.com/kubernetes/ingress/pull/1001) examples/nginx/rbac: Give access to own namespace
- [x] [#1005](https://github.com/kubernetes/ingress/pull/1005) Update configuration.md
- [x] [#1018](https://github.com/kubernetes/ingress/pull/1018) add docs for `proxy-set-headers` and `add-headers`
- [x] [#1038](https://github.com/kubernetes/ingress/pull/1038) typo / spelling in README.md
- [x] [#1039](https://github.com/kubernetes/ingress/pull/1039) typo in examples/tcp/nginx/README.md
- [x] [#1049](https://github.com/kubernetes/ingress/pull/1049) Fix config name in the example.
- [x] [#1054](https://github.com/kubernetes/ingress/pull/1054) Fix link to UDP example
- [x] [#1084](https://github.com/kubernetes/ingress/pull/1084) (issue #310)Fix some broken link
- [x] [#1103](https://github.com/kubernetes/ingress/pull/1103) Add GoDoc Widget
- [x] [#1105](https://github.com/kubernetes/ingress/pull/1105) Make Readme file more readable
- [x] [#1106](https://github.com/kubernetes/ingress/pull/1106) Update annotations.md
- [x] [#1107](https://github.com/kubernetes/ingress/pull/1107) Fix Broken Link
- [x] [#1119](https://github.com/kubernetes/ingress/pull/1119) fix typos in controllers/nginx/README.md
- [x] [#1122](https://github.com/kubernetes/ingress/pull/1122) Fix broken link
- [x] [#1131](https://github.com/kubernetes/ingress/pull/1131) Add short help doc in configuration for nginx limit rate
- [x] [#1143](https://github.com/kubernetes/ingress/pull/1143) Minor Typo Fix
- [x] [#1144](https://github.com/kubernetes/ingress/pull/1144) Minor Typo fix
- [x] [#1145](https://github.com/kubernetes/ingress/pull/1145) Minor Typo fix
- [x] [#1146](https://github.com/kubernetes/ingress/pull/1146) Fix Minor Typo in Readme
- [x] [#1147](https://github.com/kubernetes/ingress/pull/1147) Minor Typo Fix
- [x] [#1148](https://github.com/kubernetes/ingress/pull/1148) Minor Typo Fix in Getting-Started.md
- [x] [#1149](https://github.com/kubernetes/ingress/pull/1149) Fix Minor Typo in TLS authentication
- [x] [#1150](https://github.com/kubernetes/ingress/pull/1150) Fix Minor Typo in Customize the HAProxy configuration
- [x] [#1151](https://github.com/kubernetes/ingress/pull/1151) Fix Minor Typo in customization custom-template
- [x] [#1152](https://github.com/kubernetes/ingress/pull/1152) Fix minor typo in HAProxy Multi TLS certificate termination
- [x] [#1153](https://github.com/kubernetes/ingress/pull/1153) Fix minor typo in Multi TLS certificate termination
- [x] [#1154](https://github.com/kubernetes/ingress/pull/1154) Fix minor typo in Role Based Access Control
- [x] [#1155](https://github.com/kubernetes/ingress/pull/1155) Fix minor typo in TCP loadbalancing
- [x] [#1156](https://github.com/kubernetes/ingress/pull/1156) Fix minor typo in UDP loadbalancing
- [x] [#1157](https://github.com/kubernetes/ingress/pull/1157) Fix minor typos in Prerequisites
- [x] [#1158](https://github.com/kubernetes/ingress/pull/1158) Fix minor typo in Ingress examples
- [x] [#1159](https://github.com/kubernetes/ingress/pull/1159) Fix minor typos in Ingress admin guide
- [x] [#1160](https://github.com/kubernetes/ingress/pull/1160) Fix a broken href and typo in Ingress FAQ
- [x] [#1165](https://github.com/kubernetes/ingress/pull/1165) Update CONTRIBUTING.md
- [x] [#1168](https://github.com/kubernetes/ingress/pull/1168) finx link to running-locally.md
- [x] [#1170](https://github.com/kubernetes/ingress/pull/1170) Update dead link in nginx/HTTPS section
- [x] [#1172](https://github.com/kubernetes/ingress/pull/1172) Update README.md
- [x] [#1173](https://github.com/kubernetes/ingress/pull/1173) Update admin.md
- [x] [#1174](https://github.com/kubernetes/ingress/pull/1174) fix several titles
- [x] [#1177](https://github.com/kubernetes/ingress/pull/1177) fix typos
- [x] [#1188](https://github.com/kubernetes/ingress/pull/1188) Fix minor typo
- [x] [#1189](https://github.com/kubernetes/ingress/pull/1189) Fix sign in URL redirect parameter
- [x] [#1192](https://github.com/kubernetes/ingress/pull/1192) Update README.md
- [x] [#1195](https://github.com/kubernetes/ingress/pull/1195) Update troubleshooting.md
- [x] [#1196](https://github.com/kubernetes/ingress/pull/1196) Update README.md
- [x] [#1209](https://github.com/kubernetes/ingress/pull/1209) Update README.md
- [x] [#1085](https://github.com/kubernetes/ingress/pull/1085) Fix ConfigMap's namespace in custom configuration example for nginx
- [x] [#1142](https://github.com/kubernetes/ingress/pull/1142) Fix typo in multiple docs
- [x] [#1228](https://github.com/kubernetes/ingress/pull/1228) Update release doc in getting-started.md
- [x] [#1230](https://github.com/kubernetes/ingress/pull/1230) Update godep guide link

### 0.9-beta.11

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.11`

Fixes NGINX [CVE-2017-7529](http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-7529)

_Changes:_

- [x] [#659](https://github.com/kubernetes/ingress/pull/659) [nginx] TCP configmap should allow listen proxy_protocol per service
- [x] [#730](https://github.com/kubernetes/ingress/pull/730) Add support for add_headers
- [x] [#808](https://github.com/kubernetes/ingress/pull/808) HTTP->HTTPS redirect does not work with use-proxy-protocol: "true"
- [x] [#921](https://github.com/kubernetes/ingress/pull/921) Make proxy-real-ip-cidr a comma separated list
- [x] [#930](https://github.com/kubernetes/ingress/pull/930) Add support for proxy protocol in TCP services
- [x] [#933](https://github.com/kubernetes/ingress/pull/933) Lint code
- [x] [#937](https://github.com/kubernetes/ingress/pull/937) Fix lint code errors
- [x] [#940](https://github.com/kubernetes/ingress/pull/940) Sets parameters for a shared memory zone of limit_conn_zone
- [x] [#949](https://github.com/kubernetes/ingress/pull/949) fix nginx version to 1.13.3 to fix integer overflow
- [x] [#956](https://github.com/kubernetes/ingress/pull/956) Simplify handling of ssl certificates
- [x] [#958](https://github.com/kubernetes/ingress/pull/958) Release ubuntu-slim:0.13
- [x] [#959](https://github.com/kubernetes/ingress/pull/959) Release nginx-slim 0.21
- [x] [#960](https://github.com/kubernetes/ingress/pull/960) Update nginx in ingress controller
- [x] [#964](https://github.com/kubernetes/ingress/pull/964) Support for proxy_headers_hash_bucket_size and proxy_headers_hash_max_size
- [x] [#966](https://github.com/kubernetes/ingress/pull/966) Fix error checking for pod name & NS
- [x] [#967](https://github.com/kubernetes/ingress/pull/967) Fix runningAddresses typo
- [x] [#968](https://github.com/kubernetes/ingress/pull/968) Fix missing hyphen in yaml for nginx RBAC example
- [x] [#973](https://github.com/kubernetes/ingress/pull/973) check number of servers in configuration comparator

### 0.9-beta.10

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.10`

Fix release 0.9-beta.9

### 0.9-beta.9

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.9`

_New Features:_

- Add support for arm and ppc64le

_Changes:_

- [x] [#548](https://github.com/kubernetes/ingress/pull/548) nginx: support multidomain certificates
- [x] [#620](https://github.com/kubernetes/ingress/pull/620) [nginx] Listening ports are not configurable, so ingress can't be run multiple times per node when using CNI
- [x] [#648](https://github.com/kubernetes/ingress/pull/648) publish-service argument isn't honored when ELB is internal only facing.
- [x] [#833](https://github.com/kubernetes/ingress/pull/833) WIP: Avoid reloads implementing Equals in structs
- [x] [#838](https://github.com/kubernetes/ingress/pull/838) Feature request: Add ingress annotation to enable upstream "keepalive" option
- [x] [#844](https://github.com/kubernetes/ingress/pull/844) ingress annotations affinity is not working
- [x] [#862](https://github.com/kubernetes/ingress/pull/862) Avoid reloads implementing Equaler interface
- [x] [#864](https://github.com/kubernetes/ingress/pull/864) Remove dead code
- [x] [#868](https://github.com/kubernetes/ingress/pull/868) Lint nginx code
- [x] [#871](https://github.com/kubernetes/ingress/pull/871) Add feature to allow sticky sessions per location
- [x] [#873](https://github.com/kubernetes/ingress/pull/873) Update README.md
- [x] [#876](https://github.com/kubernetes/ingress/pull/876) Add information about nginx controller flags
- [x] [#878](https://github.com/kubernetes/ingress/pull/878) Update go to 1.8.3
- [x] [#881](https://github.com/kubernetes/ingress/pull/881) Option to not remove loadBalancer status record?
- [x] [#882](https://github.com/kubernetes/ingress/pull/882) Add flag to skip the update of Ingress status on shutdown
- [x] [#885](https://github.com/kubernetes/ingress/pull/885) Don't use $proxy_protocol var which may be undefined.
- [x] [#886](https://github.com/kubernetes/ingress/pull/886) Add support for SubjectAltName in SSL certificates
- [x] [#888](https://github.com/kubernetes/ingress/pull/888) Update nginx-slim to 0.19
- [x] [#889](https://github.com/kubernetes/ingress/pull/889) Add PHOST to backend
- [x] [#890](https://github.com/kubernetes/ingress/pull/890) Improve variable configuration for source IP address
- [x] [#892](https://github.com/kubernetes/ingress/pull/892) Add upstream keepalive connections cache
- [x] [#897](https://github.com/kubernetes/ingress/pull/897) Update outdated ingress resource link
- [x] [#898](https://github.com/kubernetes/ingress/pull/898) add error check right when reload nginx fail
- [x] [#899](https://github.com/kubernetes/ingress/pull/899) Fix nginx error check
- [x] [#900](https://github.com/kubernetes/ingress/pull/900) After #862 changes in the configmap do not trigger a reload
- [x] [#901](https://github.com/kubernetes/ingress/pull/901) [doc] Update NGinX status port to 18080
- [x] [#902](https://github.com/kubernetes/ingress/pull/902) Always reload after a change in the configuration
- [x] [#904](https://github.com/kubernetes/ingress/pull/904) Fix nginx sticky sessions
- [x] [#906](https://github.com/kubernetes/ingress/pull/906) Fix race condition with closed channels
- [x] [#907](https://github.com/kubernetes/ingress/pull/907) nginx/proxy: allow specifying next upstream behaviour
- [x] [#910](https://github.com/kubernetes/ingress/pull/910) Feature request: use `X-Forwarded-Host` from the reverse proxy before
- [x] [#911](https://github.com/kubernetes/ingress/pull/911) Improve X-Forwarded-Host support
- [x] [#915](https://github.com/kubernetes/ingress/pull/915) Release nginx-slim 0.20
- [x] [#916](https://github.com/kubernetes/ingress/pull/916) Add arm and ppc64le support
- [x] [#919](https://github.com/kubernetes/ingress/pull/919) Apply the 'ssl-redirect' annotation per-location
- [x] [#922](https://github.com/kubernetes/ingress/pull/922) Add example of TLS termination using a classic ELB

### 0.9-beta.8

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.8`

_Changes:_

- [x] [#761](https://github.com/kubernetes/ingress/pull/761) NGINX TCP Ingresses do not bind on IPv6
- [x] [#850](https://github.com/kubernetes/ingress/pull/850) Fix IPv6 UDP stream section
- [x] [#851](https://github.com/kubernetes/ingress/pull/851) ensure private key and certificate match
- [x] [#852](https://github.com/kubernetes/ingress/pull/852) Don't expose certificate metrics for default server
- [x] [#846](https://github.com/kubernetes/ingress/pull/846) Match ServicePort to Endpoints by Name
- [x] [#854](https://github.com/kubernetes/ingress/pull/854) Document log-format-stream and log-format-upstream
- [x] [#847](https://github.com/kubernetes/ingress/pull/847) fix semicolon
- [x] [#848](https://github.com/kubernetes/ingress/pull/848) Add metric "ssl certificate expiration"
- [x] [#839](https://github.com/kubernetes/ingress/pull/839) "No endpoints" issue
- [x] [#845](https://github.com/kubernetes/ingress/pull/845) Fix no endpoints issue when named ports are used
- [x] [#822](https://github.com/kubernetes/ingress/pull/822) Release ubuntu-slim 0.11
- [x] [#824](https://github.com/kubernetes/ingress/pull/824) Update nginx-slim to 0.18
- [x] [#823](https://github.com/kubernetes/ingress/pull/823) Release nginx-slim 0.18
- [x] [#827](https://github.com/kubernetes/ingress/pull/827) Introduce working example of nginx controller with rbac
- [x] [#835](https://github.com/kubernetes/ingress/pull/835) Make log format json escaping configurable
- [x] [#843](https://github.com/kubernetes/ingress/pull/843) Avoid setting maximum number of open file descriptors lower than 1024
- [x] [#837](https://github.com/kubernetes/ingress/pull/837) Cleanup interface
- [x] [#836](https://github.com/kubernetes/ingress/pull/836) Make log format json escaping configurable
- [x] [#828](https://github.com/kubernetes/ingress/pull/828) Wrap IPv6 endpoints in []
- [x] [#821](https://github.com/kubernetes/ingress/pull/821) nginx-ingress: occasional 503 Service Temporarily Unavailable
- [x] [#829](https://github.com/kubernetes/ingress/pull/829) feat(template): wrap IPv6 addresses in []
- [x] [#786](https://github.com/kubernetes/ingress/pull/786) Update echoserver image version in examples
- [x] [#825](https://github.com/kubernetes/ingress/pull/825) Create or delete ingress based on class annotation
- [x] [#790](https://github.com/kubernetes/ingress/pull/790) #789 removing duplicate X-Real-IP header
- [x] [#792](https://github.com/kubernetes/ingress/pull/792) Avoid checking if the controllers are synced
- [x] [#798](https://github.com/kubernetes/ingress/pull/798) nginx: RBAC for leader election
- [x] [#799](https://github.com/kubernetes/ingress/pull/799) could not build variables_hash
- [x] [#809](https://github.com/kubernetes/ingress/pull/809) Fix dynamic variable name
- [x] [#804](https://github.com/kubernetes/ingress/pull/804) Fix #798 - RBAC for leader election
- [x] [#806](https://github.com/kubernetes/ingress/pull/806) fix ingress rbac roles
- [x] [#811](https://github.com/kubernetes/ingress/pull/811) external auth - proxy_pass_request_body off + big bodies give 500/413
- [x] [#785](https://github.com/kubernetes/ingress/pull/785) Publish echoheader image
- [x] [#813](https://github.com/kubernetes/ingress/pull/813) Added client_max_body_size to authPath location
- [x] [#814](https://github.com/kubernetes/ingress/pull/814) rbac-nginx: resourceNames cannot filter create verb
- [x] [#774](https://github.com/kubernetes/ingress/pull/774) Add IPv6 support in TCP and UDP stream section
- [x] [#784](https://github.com/kubernetes/ingress/pull/784) Allow customization of variables hash tables
- [x] [#782](https://github.com/kubernetes/ingress/pull/782) Set "proxy_pass_header Server;"
- [x] [#783](https://github.com/kubernetes/ingress/pull/783) nginx/README.md: clarify app-root and fix example hyperlink
- [x] [#787](https://github.com/kubernetes/ingress/pull/787) Add setting to allow returning the Server header from the backend

### 0.9-beta.7

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.7`

_Changes:_

- [x] [#777](https://github.com/kubernetes/ingress/pull/777) Update sniff parser to fix index out of bound error

### 0.9-beta.6

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.6`

_Changes:_

- [x] [#647](https://github.com/kubernetes/ingress/pull/647) ingress.class enhancement for debugging.
- [x] [#708](https://github.com/kubernetes/ingress/pull/708) ingress losing real source IP when tls enabled
- [x] [#760](https://github.com/kubernetes/ingress/pull/760) Change recorder event scheme
- [x] [#704](https://github.com/kubernetes/ingress/pull/704) fix nginx reload flags '-c'
- [x] [#757](https://github.com/kubernetes/ingress/pull/757) Replace use of endpoints as locks with configmap
- [x] [#752](https://github.com/kubernetes/ingress/pull/752) nginx ingress header config backwards
- [x] [#756](https://github.com/kubernetes/ingress/pull/756) Fix bad variable assignment in template nginx
- [x] [#729](https://github.com/kubernetes/ingress/pull/729) Release nginx-slim 0.17
- [x] [#755](https://github.com/kubernetes/ingress/pull/755) Fix server name hash maxSize default value
- [x] [#741](https://github.com/kubernetes/ingress/pull/741) Update golang dependencies
- [x] [#749](https://github.com/kubernetes/ingress/pull/749) Remove service annotation for namedPorts
- [x] [#740](https://github.com/kubernetes/ingress/pull/740) Refactoring whitelist source IP verification
- [x] [#734](https://github.com/kubernetes/ingress/pull/734) Specify nginx image arch
- [x] [#728](https://github.com/kubernetes/ingress/pull/728) Update nginx image
- [x] [#723](https://github.com/kubernetes/ingress/pull/723) update readme about vts metrics
- [x] [#726](https://github.com/kubernetes/ingress/pull/726) Release ubuntu-slim 0.10
- [x] [#727](https://github.com/kubernetes/ingress/pull/727) [nginx] whitelist-source-range doesn’t work on ssl port
- [x] [#709](https://github.com/kubernetes/ingress/pull/709) Add config for X-Forwarded-For trust
- [x] [#679](https://github.com/kubernetes/ingress/pull/679) add getenv
- [x] [#680](https://github.com/kubernetes/ingress/pull/680) nginx/pkg/config: delete unuseful variable
- [x] [#716](https://github.com/kubernetes/ingress/pull/716) Add secure-verify-ca-secret annotation
- [x] [#722](https://github.com/kubernetes/ingress/pull/722) Remove go-reap and use tini as process reaper
- [x] [#725](https://github.com/kubernetes/ingress/pull/725) Add keepalive_requests and client_body_buffer_size options
- [x] [#724](https://github.com/kubernetes/ingress/pull/724) change the directory of default-backend.yaml
- [x] [#656](https://github.com/kubernetes/ingress/pull/656) Nginx Ingress Controller - Specify load balancing method
- [x] [#717](https://github.com/kubernetes/ingress/pull/717) delete unuseful variable
- [x] [#712](https://github.com/kubernetes/ingress/pull/712) Set $proxy_upstream_name before location directive
- [x] [#715](https://github.com/kubernetes/ingress/pull/715) Corrected annotation ex `signin-url` to `auth-url`
- [x] [#718](https://github.com/kubernetes/ingress/pull/718) nodeController sync
- [x] [#694](https://github.com/kubernetes/ingress/pull/694) SSL-Passthrough broken in beta.5
- [x] [#678](https://github.com/kubernetes/ingress/pull/678) Convert CN SSL Certificate to lowercase before comparison
- [x] [#690](https://github.com/kubernetes/ingress/pull/690) Fix IP in logs for https traffic
- [x] [#673](https://github.com/kubernetes/ingress/pull/673) Override load balancer alg view config map
- [x] [#675](https://github.com/kubernetes/ingress/pull/675) Use proxy-protocol to pass through source IP to nginx
- [x] [#707](https://github.com/kubernetes/ingress/pull/707) use nginx vts module version 0.1.14
- [x] [#702](https://github.com/kubernetes/ingress/pull/702) Document passing of ssl_client_cert to backend
- [x] [#688](https://github.com/kubernetes/ingress/pull/688) Add example of UDP loadbalancing
- [x] [#696](https://github.com/kubernetes/ingress/pull/696) [nginx] pass non-SNI TLS hello to default backend, Fixes #693
- [x] [#685](https://github.com/kubernetes/ingress/pull/685) Fix error in generated nginx.conf for optional hsts-preload

### 0.9-beta.5

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.5`

_Changes:_

- [x] [#663](https://github.com/kubernetes/ingress/pull/663) Remove helper required in go < 1.8
- [x] [#662](https://github.com/kubernetes/ingress/pull/662) Add debug information about ingress class
- [x] [#661](https://github.com/kubernetes/ingress/pull/661) Avoid running nginx if the configuration file is empty
- [x] [#660](https://github.com/kubernetes/ingress/pull/660) Rollback queue refactoring
- [x] [#654](https://github.com/kubernetes/ingress/pull/654) Update go version to 1.8

### 0.9-beta.4

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.4`

_New Features:_

- Add support for services of type ExternalName

_Changes:_

- [x] [#635](https://github.com/kubernetes/ingress/pull/635) Allow configuration of features underscores_in_headers and ignore_invalid_headers
- [x] [#633](https://github.com/kubernetes/ingress/pull/633) Fix lint errors
- [x] [#630](https://github.com/kubernetes/ingress/pull/630) Add example of TCP loadbalancing
- [x] [#629](https://github.com/kubernetes/ingress/pull/629) Add support for services of type ExternalName
- [x] [#624](https://github.com/kubernetes/ingress/pull/624) Compute server_names_hash_bucket_size correctly
- [x] [#615](https://github.com/kubernetes/ingress/pull/615) Process exited cleanly before we hit wait4
- [x] [#614](https://github.com/kubernetes/ingress/pull/614) Refactor nginx ssl passthrough
- [x] [#613](https://github.com/kubernetes/ingress/pull/613) Status leader election must consired the ingress class
- [x] [#607](https://github.com/kubernetes/ingress/pull/607) Allow custom server_names_hash_max_size & server_names_hash_bucket_size
- [x] [#601](https://github.com/kubernetes/ingress/pull/601) add a judgment
- [x] [#601](https://github.com/kubernetes/ingress/pull/600) Replace custom child reap code with go-reap
- [x] [#597](https://github.com/kubernetes/ingress/pull/599) Add flag to force namespace isolation
- [x] [#595](https://github.com/kubernetes/ingress/pull/595) Remove Host header from auth_request proxy configuration
- [x] [#588](https://github.com/kubernetes/ingress/pull/588) Read resolv.conf file just once
- [x] [#586](https://github.com/kubernetes/ingress/pull/586) Updated instructions to create an ingress controller build
- [x] [#583](https://github.com/kubernetes/ingress/pull/583) fixed lua_package_path in nginx.tmpl
- [x] [#580](https://github.com/kubernetes/ingress/pull/580) Updated faq for running multiple ingress controller
- [x] [#579](https://github.com/kubernetes/ingress/pull/579) Detect if the ingress controller is running with multiple replicas
- [x] [#578](https://github.com/kubernetes/ingress/pull/578) Set different listeners per protocol version
- [x] [#577](https://github.com/kubernetes/ingress/pull/577) Avoid zombie child processes
- [x] [#576](https://github.com/kubernetes/ingress/pull/576) Replace secret workqueue
- [x] [#568](https://github.com/kubernetes/ingress/pull/568) Revert merge annotations to the implicit root context
- [x] [#563](https://github.com/kubernetes/ingress/pull/563) Add option to disable hsts preload
- [x] [#560](https://github.com/kubernetes/ingress/pull/560) Fix intermittent misconfiguration of backend.secure and SessionAffinity
- [x] [#556](https://github.com/kubernetes/ingress/pull/556) Update nginx version and remove dumb-init
- [x] [#551](https://github.com/kubernetes/ingress/pull/551) Build namespace and ingress class as label
- [x] [#546](https://github.com/kubernetes/ingress/pull/546) Fix a couple of 'does not contains' typos
- [x] [#542](https://github.com/kubernetes/ingress/pull/542) Fix lint errors
- [x] [#540](https://github.com/kubernetes/ingress/pull/540) Add Backends.SSLPassthrough attribute
- [x] [#539](https://github.com/kubernetes/ingress/pull/539) Migrate to client-go
- [x] [#536](https://github.com/kubernetes/ingress/pull/536) add unit test cases for core/pkg/ingress/controller/backend_ssl
- [x] [#535](https://github.com/kubernetes/ingress/pull/535) Add test for ingress status update
- [x] [#532](https://github.com/kubernetes/ingress/pull/532) Add setting to configure ecdh curve
- [x] [#531](https://github.com/kubernetes/ingress/pull/531) Fix link to examples
- [x] [#530](https://github.com/kubernetes/ingress/pull/530) Fix link to custom nginx configuration
- [x] [#528](https://github.com/kubernetes/ingress/pull/528) Add reference to apiserver-host flag
- [x] [#527](https://github.com/kubernetes/ingress/pull/527) Add annotations to location of default backend (root context)
- [x] [#525](https://github.com/kubernetes/ingress/pull/525) Avoid negative values configuring the max number of open files
- [x] [#523](https://github.com/kubernetes/ingress/pull/523) Fix a typo in an error message
- [x] [#521](https://github.com/kubernetes/ingress/pull/521) nginx-ingress-controller is built twice by docker-build target
- [x] [#517](https://github.com/kubernetes/ingress/pull/517) Use whitelist-source-range from configmap when no annotation on ingress
- [x] [#516](https://github.com/kubernetes/ingress/pull/516) Convert WorkerProcesses setting to string to allow the value auto
- [x] [#512](https://github.com/kubernetes/ingress/pull/512) Fix typos regarding the ssl-passthrough annotation documentation
- [x] [#505](https://github.com/kubernetes/ingress/pull/505) add unit test cases for core/pkg/ingress/controller/annotations
- [x] [#503](https://github.com/kubernetes/ingress/pull/503) Add example for nginx in aws
- [x] [#502](https://github.com/kubernetes/ingress/pull/502) Add information about SSL Passthrough annotation
- [x] [#500](https://github.com/kubernetes/ingress/pull/500) Improve TLS secret configuration
- [x] [#498](https://github.com/kubernetes/ingress/pull/498) Proper enqueue a secret on the secret queue
- [x] [#493](https://github.com/kubernetes/ingress/pull/493) Update nginx and vts module
- [x] [#490](https://github.com/kubernetes/ingress/pull/490) Add unit test case for named_port
- [x] [#488](https://github.com/kubernetes/ingress/pull/488) Adds support for CORS on error responses and Authorization header
- [x] [#485](https://github.com/kubernetes/ingress/pull/485) Fix typo nginx configMap vts metrics customization
- [x] [#481](https://github.com/kubernetes/ingress/pull/481) Remove unnecessary quote in nginx log format
- [x] [#471](https://github.com/kubernetes/ingress/pull/471) prometheus scrape annotations
- [x] [#460](https://github.com/kubernetes/ingress/pull/460) add example of 'run multiple haproxy ingress controllers as a deployment'
- [x] [#459](https://github.com/kubernetes/ingress/pull/459) Add information about SSL certificates in the default log level
- [x] [#456](https://github.com/kubernetes/ingress/pull/456) Avoid upstreams with multiple servers with the same port
- [x] [#454](https://github.com/kubernetes/ingress/pull/454) Pass request port to real server
- [x] [#450](https://github.com/kubernetes/ingress/pull/450) fix nginx-tcp-and-udp on same port
- [x] [#446](https://github.com/kubernetes/ingress/pull/446) remove configmap validations
- [x] [#445](https://github.com/kubernetes/ingress/pull/445) Remove snakeoil certificate generation
- [x] [#442](https://github.com/kubernetes/ingress/pull/442) Fix a few bugs in the nginx-ingress-controller Makefile
- [x] [#441](https://github.com/kubernetes/ingress/pull/441) skip validation when configmap is empty
- [x] [#439](https://github.com/kubernetes/ingress/pull/439) Avoid a nil-reference when the temporary file cannot be created
- [x] [#438](https://github.com/kubernetes/ingress/pull/438) Improve English in error messages
- [x] [#437](https://github.com/kubernetes/ingress/pull/437) Reference constant

### 0.9-beta.3

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.3`

_New Features:_

- Custom log formats using `log-format-upstream` directive in the configuration configmap.
- Force redirect to SSL using the annotation `ingress.kubernetes.io/force-ssl-redirect`
- Prometheus metric for VTS status module (transparent, just enable vts stats)
- Improved external authentication adding `ingress.kubernetes.io/auth-signin` annotation. Please check this [example](https://github.com/kubernetes/ingress/tree/master/examples/external-auth/nginx)

_Breaking changes:_

- `ssl-dh-param` configuration in configmap is now the name of a secret that contains the Diffie-Hellman key

_Changes:_

- [x] [#433](https://github.com/kubernetes/ingress/pull/433) close over the ingress variable or the last assignment will be used
- [x] [#424](https://github.com/kubernetes/ingress/pull/424) Manually sync secrets from certificate authentication annotations
- [x] [#423](https://github.com/kubernetes/ingress/pull/423) Scrap json metrics from nginx vts module when enabled
- [x] [#418](https://github.com/kubernetes/ingress/pull/418) Only update Ingress status for the configured class
- [x] [#415](https://github.com/kubernetes/ingress/pull/415) Improve external authentication docs
- [x] [#410](https://github.com/kubernetes/ingress/pull/410) Add support for "signin url"
- [x] [#409](https://github.com/kubernetes/ingress/pull/409) Allow custom http2 header sizes
- [x] [#408](https://github.com/kubernetes/ingress/pull/408) Review docs
- [x] [#406](https://github.com/kubernetes/ingress/pull/406) Add debug info and fix spelling
- [x] [#402](https://github.com/kubernetes/ingress/pull/402) allow specifying custom dh param
- [x] [#397](https://github.com/kubernetes/ingress/pull/397) Fix external auth
- [x] [#394](https://github.com/kubernetes/ingress/pull/394) Update README.md
- [x] [#392](https://github.com/kubernetes/ingress/pull/392) Fix http2 header size
- [x] [#391](https://github.com/kubernetes/ingress/pull/391) remove tmp nginx-diff files
- [x] [#390](https://github.com/kubernetes/ingress/pull/390) Fix RateLimit comment
- [x] [#385](https://github.com/kubernetes/ingress/pull/385) add Copyright
- [x] [#382](https://github.com/kubernetes/ingress/pull/382) Ingress Fake Certificate generation
- [x] [#380](https://github.com/kubernetes/ingress/pull/380) Fix custom log format
- [x] [#373](https://github.com/kubernetes/ingress/pull/373) Cleanup
- [x] [#371](https://github.com/kubernetes/ingress/pull/371) add configuration to disable listening on ipv6
- [x] [#370](https://github.com/kubernetes/ingress/pull/270) Add documentation for ingress.kubernetes.io/force-ssl-redirect
- [x] [#369](https://github.com/kubernetes/ingress/pull/369) Minor text fix for "ApiServer"
- [x] [#367](https://github.com/kubernetes/ingress/pull/367) BuildLogFormatUpstream was always using the default log-format
- [x] [#366](https://github.com/kubernetes/ingress/pull/366) add_judgment
- [x] [#365](https://github.com/kubernetes/ingress/pull/365) add ForceSSLRedirect ingress annotation
- [x] [#364](https://github.com/kubernetes/ingress/pull/364) Fix error caused by increasing proxy_buffer_size (#363)
- [x] [#362](https://github.com/kubernetes/ingress/pull/362) Fix ingress class
- [x] [#360](https://github.com/kubernetes/ingress/pull/360) add example of 'run multiple nginx ingress controllers as a deployment'
- [x] [#358](https://github.com/kubernetes/ingress/pull/358) Checks if the TLS secret contains a valid keypair structure
- [x] [#356](https://github.com/kubernetes/ingress/pull/356) Disable listen only on ipv6 and fix proxy_protocol
- [x] [#354](https://github.com/kubernetes/ingress/pull/354) add judgment
- [x] [#352](https://github.com/kubernetes/ingress/pull/352) Add ability to customize upstream and stream log format
- [x] [#351](https://github.com/kubernetes/ingress/pull/351) Enable custom election id for status sync.
- [x] [#347](https://github.com/kubernetes/ingress/pull/347) Fix client source IP address
- [x] [#345](https://github.com/kubernetes/ingress/pull/345) Fix lint error
- [x] [#344](https://github.com/kubernetes/ingress/pull/344) Refactoring of TCP and UDP services
- [x] [#343](https://github.com/kubernetes/ingress/pull/343) Fix node lister when --watch-namespace is used
- [x] [#341](https://github.com/kubernetes/ingress/pull/341) Do not run coverage check in the default target.
- [x] [#340](https://github.com/kubernetes/ingress/pull/340) Add support for specify proxy cookie path/domain
- [x] [#337](https://github.com/kubernetes/ingress/pull/337) Fix for formatting error introduced in #304
- [x] [#335](https://github.com/kubernetes/ingress/pull/335) Fix for vet complaints:
- [x] [#332](https://github.com/kubernetes/ingress/pull/332) Add annotation to customize nginx configuration
- [x] [#331](https://github.com/kubernetes/ingress/pull/331) Correct spelling mistake
- [x] [#328](https://github.com/kubernetes/ingress/pull/328) fix misspell "affinity" in main.go
- [x] [#326](https://github.com/kubernetes/ingress/pull/326) add nginx daemonset example
- [x] [#311](https://github.com/kubernetes/ingress/pull/311) Sort stream service ports to avoid extra reloads
- [x] [#307](https://github.com/kubernetes/ingress/pull/307) Add docs for body-size annotation
- [x] [#306](https://github.com/kubernetes/ingress/pull/306) modify nginx readme
- [x] [#304](https://github.com/kubernetes/ingress/pull/304) change 'buildSSPassthrouthUpstreams' to 'buildSSLPassthroughUpstreams'

### 0.9-beta.2

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.2`

_New Features:_

- New configuration flag `proxy-set-headers` to allow set custom headers before send traffic to backends. [Example here](https://github.com/kubernetes/ingress/tree/master/examples/customization/custom-headers/nginx)
- Disable directive access_log globally using `disable-access-log: "true"` in the configuration ConfigMap.
- Sticky session per Ingress rule using the annotation `ingress.kubernetes.io/affinity`. [Example here](https://github.com/kubernetes/ingress/tree/master/examples/affinity/cookie/nginx)

_Changes:_

- [x] [#300](https://github.com/kubernetes/ingress/pull/300) Change nginx variable to use in filter of access_log
- [x] [#296](https://github.com/kubernetes/ingress/pull/296) Fix rewrite regex to match the start of the URL and not a substring
- [x] [#293](https://github.com/kubernetes/ingress/pull/293) Update makefile gcloud docker command
- [x] [#290](https://github.com/kubernetes/ingress/pull/290) Update nginx version in ingress controller to 1.11.10
- [x] [#286](https://github.com/kubernetes/ingress/pull/286) Add logs to help debugging and simplify default upstream configuration
- [x] [#285](https://github.com/kubernetes/ingress/pull/285) Added a Node StoreLister type
- [x] [#281](https://github.com/kubernetes/ingress/pull/281) Add chmod up directory tree for world read/execute on directories
- [x] [#279](https://github.com/kubernetes/ingress/pull/279) fix wrong link in the file of examples/README.md
- [x] [#275](https://github.com/kubernetes/ingress/pull/275) Pass headers to custom error backend
- [x] [#272](https://github.com/kubernetes/ingress/pull/272) Fix error getting class information from Ingress annotations
- [x] [#268](https://github.com/kubernetes/ingress/pull/268) minor: Fix typo in nginx README
- [x] [#265](https://github.com/kubernetes/ingress/pull/265) Fix rewrite annotation parser
- [x] [#262](https://github.com/kubernetes/ingress/pull/262) Add nginx README and configuration docs back
- [x] [#261](https://github.com/kubernetes/ingress/pull/261) types.go: fix typo in godoc
- [x] [#258](https://github.com/kubernetes/ingress/pull/258) Nginx sticky annotations
- [x] [#255](https://github.com/kubernetes/ingress/pull/255) Adds support for disabling access_log globally
- [x] [#247](https://github.com/kubernetes/ingress/pull/247) Fix wrong URL in nginx ingress configuration
- [x] [#246](https://github.com/kubernetes/ingress/pull/246) Add support for custom proxy headers using a ConfigMap
- [x] [#244](https://github.com/kubernetes/ingress/pull/244) Add information about cors annotation
- [x] [#241](https://github.com/kubernetes/ingress/pull/241) correct a spell mistake
- [x] [#232](https://github.com/kubernetes/ingress/pull/232) Change searchs with searches
- [x] [#231](https://github.com/kubernetes/ingress/pull/231) Add information about proxy_protocol in port 442
- [x] [#228](https://github.com/kubernetes/ingress/pull/228) Fix worker check issue
- [x] [#227](https://github.com/kubernetes/ingress/pull/227) proxy_protocol on ssl_passthrough listener
- [x] [#223](https://github.com/kubernetes/ingress/pull/223) Fix panic if a tempfile cannot be created
- [x] [#220](https://github.com/kubernetes/ingress/pull/220) Fixes for minikube usage instructions.
- [x] [#219](https://github.com/kubernetes/ingress/pull/219) Fix typo, add a couple of links.
- [x] [#218](https://github.com/kubernetes/ingress/pull/218) Improve links from CONTRIBUTING.
- [x] [#217](https://github.com/kubernetes/ingress/pull/217) Fix an e2e link.
- [x] [#212](https://github.com/kubernetes/ingress/pull/212) Simplify code to obtain TCP or UDP services
- [x] [#208](https://github.com/kubernetes/ingress/pull/208) Fix nil HTTP field
- [x] [#198](https://github.com/kubernetes/ingress/pull/198) Add an example for static-ip and deployment

### 0.9-beta.1

**Image:** `gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.1`

_New Features:_

- SSL Passthrough
- New Flag `--publish-service` that set the Service fronting the ingress controllers
- Ingress status shows the correct IP/hostname address without duplicates
- Custom body sizes per Ingress
- Prometheus metrics

_Breaking changes:_

- Flag `--nginx-configmap` was replaced with `--configmap`
- Configmap field `body-size` was replaced with `proxy-body-size`

_Changes:_

- [x] [#184](https://github.com/kubernetes/ingress/pull/184) Fix template error
- [x] [#179](https://github.com/kubernetes/ingress/pull/179) Allows the usage of Default SSL Cert
- [x] [#178](https://github.com/kubernetes/ingress/pull/178) Add initialization of proxy variable
- [x] [#177](https://github.com/kubernetes/ingress/pull/177) Refactoring sysctlFSFileMax helper
- [x] [#176](https://github.com/kubernetes/ingress/pull/176) Fix TLS does not get updated when changed
- [x] [#174](https://github.com/kubernetes/ingress/pull/174) Update nginx to 1.11.9
- [x] [#172](https://github.com/kubernetes/ingress/pull/172) add some unit test cases for some packages under folder "core.pkg.ingress"
- [x] [#168](https://github.com/kubernetes/ingress/pull/168) Changes the SSL Temp file to something inside the same SSL Directory
- [x] [#165](https://github.com/kubernetes/ingress/pull/165) Fix rate limit issue when more than 2 servers enabled in ingress
- [x] [#161](https://github.com/kubernetes/ingress/pull/161) Document some missing parameters and their defaults for NGINX controller
- [x] [#158](https://github.com/kubernetes/ingress/pull/158) prefect unit test cases for annotation.proxy
- [x] [#156](https://github.com/kubernetes/ingress/pull/156) Fix issue for ratelimit
- [x] [#154](https://github.com/kubernetes/ingress/pull/154) add unit test cases for core.pkg.ingress.annotations.cors
- [x] [#151](https://github.com/kubernetes/ingress/pull/151) Port in redirect
- [x] [#150](https://github.com/kubernetes/ingress/pull/150) Add support for custom header sizes
- [x] [#149](https://github.com/kubernetes/ingress/pull/149) Add flag to allow switch off the update of Ingress status
- [x] [#148](https://github.com/kubernetes/ingress/pull/148) Add annotation to allow custom body sizes
- [x] [#145](https://github.com/kubernetes/ingress/pull/145) fix wrong links and punctuations
- [x] [#144](https://github.com/kubernetes/ingress/pull/144) add unit test cases for core.pkg.k8s
- [x] [#143](https://github.com/kubernetes/ingress/pull/143) Use protobuf instead of rest to connect to apiserver host and add troubleshooting doc
- [x] [#142](https://github.com/kubernetes/ingress/pull/142) Use system fs.max-files as limits instead of hard-coded value
- [x] [#141](https://github.com/kubernetes/ingress/pull/141) Add reuse port and backlog to port 80 and 443
- [x] [#138](https://github.com/kubernetes/ingress/pull/138) reference to const
- [x] [#136](https://github.com/kubernetes/ingress/pull/136) Add content and descriptions about nginx's configuration
- [x] [#135](https://github.com/kubernetes/ingress/pull/135) correct improper punctuation
- [x] [#134](https://github.com/kubernetes/ingress/pull/134) fix typo
- [x] [#133](https://github.com/kubernetes/ingress/pull/133) Add TCP and UDP services removed in migration
- [x] [#132](https://github.com/kubernetes/ingress/pull/132) Document nginx controller configuration tweaks
- [x] [#128](https://github.com/kubernetes/ingress/pull/128) Add tests and godebug to compare structs
- [x] [#126](https://github.com/kubernetes/ingress/pull/126) change the type of imagePullPolicy
- [x] [#123](https://github.com/kubernetes/ingress/pull/123) Add resolver configuration to nginx
- [x] [#119](https://github.com/kubernetes/ingress/pull/119) add unit test case for annotations.service
- [x] [#115](https://github.com/kubernetes/ingress/pull/115) add default_server to listen statement for default backend
- [x] [#114](https://github.com/kubernetes/ingress/pull/114) fix typo
- [x] [#113](https://github.com/kubernetes/ingress/pull/113) Add condition of enqueue and unit test cases for task.Queue
- [x] [#108](https://github.com/kubernetes/ingress/pull/108) annotations: print error and skip if malformed
- [x] [#107](https://github.com/kubernetes/ingress/pull/107) fix some wrong links of examples which to be used for nginx
- [x] [#103](https://github.com/kubernetes/ingress/pull/103) Update the nginx controller manifests
- [x] [#101](https://github.com/kubernetes/ingress/pull/101) Add unit test for strings.StringInSlice
- [x] [#99](https://github.com/kubernetes/ingress/pull/99) Update nginx to 1.11.8
- [x] [#97](https://github.com/kubernetes/ingress/pull/97) Fix gofmt
- [x] [#96](https://github.com/kubernetes/ingress/pull/96) Fix typo PassthrougBackends -> PassthroughBackends
- [x] [#95](https://github.com/kubernetes/ingress/pull/95) Deny location mapping in case of specific errors
- [x] [#94](https://github.com/kubernetes/ingress/pull/94) Add support to disable server_tokens directive
- [x] [#93](https://github.com/kubernetes/ingress/pull/93) Fix sort for catch all server
- [x] [#92](https://github.com/kubernetes/ingress/pull/92) Refactoring of nginx configuration deserialization
- [x] [#91](https://github.com/kubernetes/ingress/pull/91) Fix x-forwarded-port mapping
- [x] [#90](https://github.com/kubernetes/ingress/pull/90) fix the wrong link to build/test/release
- [x] [#89](https://github.com/kubernetes/ingress/pull/89) fix the wrong links to the examples and developer documentation
- [x] [#88](https://github.com/kubernetes/ingress/pull/88) Fix multiple tls hosts sharing the same secretName
- [x] [#86](https://github.com/kubernetes/ingress/pull/86) Update X-Forwarded-Port
- [x] [#82](https://github.com/kubernetes/ingress/pull/82) Fix incorrect X-Forwarded-Port for TLS
- [x] [#81](https://github.com/kubernetes/ingress/pull/81) Do not push containers to remote repo as part of test-e2e
- [x] [#78](https://github.com/kubernetes/ingress/pull/78) Fix #76: hardcode X-Forwarded-Port due to SSL Passthrough
- [x] [#77](https://github.com/kubernetes/ingress/pull/77) Add support for IPV6 in dns resolvers
- [x] [#66](https://github.com/kubernetes/ingress/pull/66) Start FAQ docs
- [x] [#65](https://github.com/kubernetes/ingress/pull/65) Support hostnames in Ingress status
- [x] [#64](https://github.com/kubernetes/ingress/pull/64) Sort whitelist list to avoid random orders
- [x] [#62](https://github.com/kubernetes/ingress/pull/62) Fix e2e make targets
- [x] [#61](https://github.com/kubernetes/ingress/pull/61) Ignore coverage profile files
- [x] [#58](https://github.com/kubernetes/ingress/pull/58) Fix "invalid port in upstream" on nginx controller
- [x] [#57](https://github.com/kubernetes/ingress/pull/57) Fix invalid port in upstream
- [x] [#54](https://github.com/kubernetes/ingress/pull/54) Expand developer docs
- [x] [#52](https://github.com/kubernetes/ingress/pull/52) fix typo in variable ProxyRealIPCIDR
- [x] [#44](https://github.com/kubernetes/ingress/pull/44) Bump nginx version to one higher than that in contrib
- [x] [#36](https://github.com/kubernetes/ingress/pull/36) Add nginx metrics to prometheus
- [x] [#34](https://github.com/kubernetes/ingress/pull/34) nginx: also listen on ipv6
- [x] [#32](https://github.com/kubernetes/ingress/pull/32) Restart nginx if master process dies
- [x] [#31](https://github.com/kubernetes/ingress/pull/31) Add healthz checker
- [x] [#25](https://github.com/kubernetes/ingress/pull/25) Fix a data race in TestFileWatcher
- [x] [#12](https://github.com/kubernetes/ingress/pull/12) Split implementations from generic code
- [x] [#10](https://github.com/kubernetes/ingress/pull/10) Copy Ingress history from kubernetes/contrib
- [x] [#1498](https://github.com/kubernetes/contrib/pull/1498) Refactoring of template handling
- [x] [#1571](https://github.com/kubernetes/contrib/pull/1571) use POD_NAMESPACE as a namespace in cli parameters
- [x] [#1591](https://github.com/kubernetes/contrib/pull/1591) Always listen on port 443, even without ingress rules
- [x] [#1596](https://github.com/kubernetes/contrib/pull/1596) Adapt nginx hash sizes to the number of ingress
- [x] [#1653](https://github.com/kubernetes/contrib/pull/1653) Update image version
- [x] [#1672](https://github.com/kubernetes/contrib/pull/1672) Add firewall rules and ing class clarifications
- [x] [#1711](https://github.com/kubernetes/contrib/pull/1711) Add function helpers to nginx template
- [x] [#1743](https://github.com/kubernetes/contrib/pull/1743) Allow customisation of the nginx proxy_buffer_size directive via ConfigMap
- [x] [#1749](https://github.com/kubernetes/contrib/pull/1749) Readiness probe that works behind a CP lb
- [x] [#1751](https://github.com/kubernetes/contrib/pull/1751) Add the name of the upstream in the log
- [x] [#1758](https://github.com/kubernetes/contrib/pull/1758) Update nginx to 1.11.4
- [x] [#1759](https://github.com/kubernetes/contrib/pull/1759) Add support for default backend in Ingress rule
- [x] [#1762](https://github.com/kubernetes/contrib/pull/1762) Add cloud detection
- [x] [#1766](https://github.com/kubernetes/contrib/pull/1766) Clarify the controller uses endpoints and not services
- [x] [#1767](https://github.com/kubernetes/contrib/pull/1767) Update godeps
- [x] [#1772](https://github.com/kubernetes/contrib/pull/1772) Avoid replacing nginx.conf file if the new configuration is invalid
- [x] [#1773](https://github.com/kubernetes/contrib/pull/1773) Add annotation to add CORS support
- [x] [#1786](https://github.com/kubernetes/contrib/pull/1786) Add docs about go template
- [x] [#1796](https://github.com/kubernetes/contrib/pull/1796) Add external authentication support using auth_request
- [x] [#1802](https://github.com/kubernetes/contrib/pull/1802) Initialize proxy_upstream_name variable
- [x] [#1806](https://github.com/kubernetes/contrib/pull/1806) Add docs about the log format
- [x] [#1808](https://github.com/kubernetes/contrib/pull/1808) WebSocket documentation
- [x] [#1847](https://github.com/kubernetes/contrib/pull/1847) Change structure of packages
- [x] Add annotation for custom upstream timeouts
- [x] Mutual TLS auth (https://github.com/kubernetes/contrib/issues/1870)

### 0.8.3

- [x] [#1450](https://github.com/kubernetes/contrib/pull/1450) Check for errors in nginx template
- [ ] [#1498](https://github.com/kubernetes/contrib/pull/1498) Refactoring of template handling
- [x] [#1467](https://github.com/kubernetes/contrib/pull/1467) Use ClientConfig to configure connection
- [x] [#1575](https://github.com/kubernetes/contrib/pull/1575) Update nginx to 1.11.3

### 0.8.2

- [x] [#1336](https://github.com/kubernetes/contrib/pull/1336) Add annotation to skip ingress rule
- [x] [#1338](https://github.com/kubernetes/contrib/pull/1338) Add HTTPS default backend
- [x] [#1351](https://github.com/kubernetes/contrib/pull/1351) Avoid generation of invalid ssl certificates
- [x] [#1379](https://github.com/kubernetes/contrib/pull/1379) improve nginx performance
- [x] [#1350](https://github.com/kubernetes/contrib/pull/1350) Improve performance (listen backlog=net.core.somaxconn)
- [x] [#1384](https://github.com/kubernetes/contrib/pull/1384) Unset Authorization header when proxying
- [x] [#1398](https://github.com/kubernetes/contrib/pull/1398) Mitigate HTTPoxy Vulnerability

### 0.8.1

- [x] [#1317](https://github.com/kubernetes/contrib/pull/1317) Fix duplicated real_ip_header
- [x] [#1315](https://github.com/kubernetes/contrib/pull/1315) Addresses #1314

### 0.8

- [x] [#1063](https://github.com/kubernetes/contrib/pull/1063) watches referenced tls secrets
- [x] [#850](https://github.com/kubernetes/contrib/pull/850) adds configurable SSL redirect nginx controller
- [x] [#1136](https://github.com/kubernetes/contrib/pull/1136) Fix nginx rewrite rule order
- [x] [#1144](https://github.com/kubernetes/contrib/pull/1144) Add cidr whitelist support
- [x] [#1230](https://github.com/kubernetes/contrib/pull/1130) Improve docs and examples
- [x] [#1258](https://github.com/kubernetes/contrib/pull/1258) Avoid sync without a reachable
- [x] [#1235](https://github.com/kubernetes/contrib/pull/1235) Fix stats by country in nginx status page
- [x] [#1236](https://github.com/kubernetes/contrib/pull/1236) Update nginx to add dynamic TLS records and spdy
- [x] [#1238](https://github.com/kubernetes/contrib/pull/1238) Add support for dynamic TLS records and spdy
- [x] [#1239](https://github.com/kubernetes/contrib/pull/1239) Add support for conditional log of urls
- [x] [#1253](https://github.com/kubernetes/contrib/pull/1253) Use delayed queue
- [x] [#1296](https://github.com/kubernetes/contrib/pull/1296) Fix formatting
- [x] [#1299](https://github.com/kubernetes/contrib/pull/1299) Fix formatting

### 0.7

- [x] [#898](https://github.com/kubernetes/contrib/pull/898) reorder locations. Location / must be the last one to avoid errors routing to subroutes
- [x] [#946](https://github.com/kubernetes/contrib/pull/946) Add custom authentication (Basic or Digest) to ingress rules
- [x] [#926](https://github.com/kubernetes/contrib/pull/926) Custom errors should be optional
- [x] [#1002](https://github.com/kubernetes/contrib/pull/1002) Use k8s probes (disable NGINX checks)
- [x] [#962](https://github.com/kubernetes/contrib/pull/962) Make optional http2
- [x] [#1054](https://github.com/kubernetes/contrib/pull/1054) force reload if some certificate change
- [x] [#958](https://github.com/kubernetes/contrib/pull/958) update NGINX to 1.11.0 and add digest module
- [x] [#960](https://github.com/kubernetes/contrib/issues/960) https://trac.nginx.org/nginx/changeset/ce94f07d50826fcc8d48f046fe19d59329420fdb/nginx
- [x] [#1057](https://github.com/kubernetes/contrib/pull/1057) Remove loadBalancer ip on shutdown
- [x] [#1079](https://github.com/kubernetes/contrib/pull/1079) path rewrite
- [x] [#1093](https://github.com/kubernetes/contrib/pull/1093) rate limiting
- [x] [#1102](https://github.com/kubernetes/contrib/pull/1102) geolocation of traffic in stats
- [x] [#884](https://github.com/kubernetes/contrib/issues/884) support services running ssl
- [x] [#930](https://github.com/kubernetes/contrib/issues/930) detect changes in configuration configmaps
