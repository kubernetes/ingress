1. [NGINX](https://github.com/kubernetes/ingress-nginx/tree/main/images/nginx)

* Open pull request
  
If you are updating any component in [build.sh](images/nginx/rootfs/build.sh) please also update the SHA256 checksum of that component as 
  well, the cloud build will fail with an exit 10 if not.  
Example [NGINX_VERSION](images/nginx/rootfs/build.sh#L21),
[SHA256](images/nginx/rootfs/build.sh#L124) 
* Merge
* Wait for [cloud build](https://console.cloud.google.com/cloud-build/builds?project=k8s-staging-ingress-nginx)

1a. Promote images:

Open pull request to promote staging image:
[add sha - version](https://github.com/kubernetes/k8s.io/blob/main/k8s.gcr.io/images/k8s-staging-ingress-nginx/images.yaml#L1)

The sha is available in output from [cloud build](https://console.cloud.google.com/cloud-build/builds?project=k8s-staging-ingress-nginx)

2. Change to images:
   
* [e2e](https://github.com/kubernetes/ingress-nginx/tree/main/images/test-runner)

    * [test-runner](https://github.com/kubernetes/ingress-nginx/tree/main/images/echo)
    * [echo](https://github.com/kubernetes/ingress-nginx/tree/main/images/echo)
    * [cfssl](https://github.com/kubernetes/ingress-nginx/tree/main/images/cfssl)
    * [fastcgi-helloserver](https://github.com/kubernetes/ingress-nginx/tree/main/images/fastcgi-helloserver)
    * [httpbin](https://github.com/kubernetes/ingress-nginx/tree/main/images/httpbin)

* Open pull request
* Merge
* Wait for [cloud build](https://console.cloud.google.com/cloud-build/builds?project=k8s-staging-ingress-nginx)

2a. Promote images:

* Open pull request to promote [staging image](https://github.com/kubernetes/k8s.io/blob/main/k8s.gcr.io/images/k8s-staging-ingress-nginx/images.yaml)
    * e2e-test-runner
    * e2e-test-cfssl
    * e2e-test-echo
    * e2e-test-fastcgi-helloserver
    * e2e-test-httpbin

3. Update references to e2e-test-runner image:

* [e2e-image](https://github.com/kubernetes/ingress-nginx/blob/main/test/e2e-image/Dockerfile#L1)
* [run-in-docker.sh](https://github.com/kubernetes/ingress-nginx/blob/ff60aa9e2b5377db1544091b98f475a90a630297/build/run-in-docker.sh#L37)

4. Prepare for a new release:

* Change [TAG](https://github.com/kubernetes/ingress-nginx/blob/main/TAG#L1)
* Open pull request
* Merge
* [Wait for cloud build](https://console.cloud.google.com/cloud-build/builds?project=k8s-staging-ingress-nginx)

4a. Promote images:

* Open pull request to promote [staging image](https://github.com/kubernetes/k8s.io/blob/main/k8s.gcr.io/images/k8s-staging-ingress-nginx/images.yaml)
  * controller

5. Release helm chart:

* Open pull request updating [Chart.yaml](https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/Chart.yaml#L3-L4)
* Merge
* [New helm chart is available](https://github.com/kubernetes/ingress-nginx/blob/main/.github/workflows/main.yaml#L47-L68)

6. New release:

* Update static scripts:
    * [generate-deploy-scripts.sh](https://github.com/kubernetes/ingress-nginx/blob/main/hack/generate-deploy-scripts.sh)
    * Open pull request with the updates
    * Merge

* Update Changelog and Documentation:
    * Open pull request updating [Changelog.md](https://github.com/kubernetes/ingress-nginx/blob/main/Changelog.md)
    * Update the version in [docs/deploy/index.md](docs/deploy/index.md)
    * Update Supported versions in the Support Versions table in the README.md 
    * Merge
      
7. Github release
