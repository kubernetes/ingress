#!/bin/bash

# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

if [ -n "$DEBUG" ]; then
	set -x
fi

set -o errexit
set -o nounset
set -o pipefail

DIR=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)

RELEASE_NAME=ingress-nginx
NAMESPACE=ingress-nginx

# ensures `ttlSecondsAfterFinished` gets set on Job specs
API_VERSIONS=batch/v1alpha1

NAMESPACE_VAR="
apiVersion: v1
kind: Namespace
metadata:
  name: $NAMESPACE
  labels:
    app.kubernetes.io/name: $RELEASE_NAME
    app.kubernetes.io/instance: ingress-nginx
"

# Baremetal
OUTPUT_FILE="${DIR}/deploy/static/provider/baremetal/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: NodePort

  publishService:
    enabled: false
EOF

echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}

# Cloud - generic
OUTPUT_FILE="${DIR}/deploy/static/provider/cloud/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local
EOF

echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}


# AWS - NLB
OUTPUT_FILE="${DIR}/deploy/static/provider/aws/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local
    annotations:
      service.beta.kubernetes.io/aws-load-balancer-backend-protocol: "tcp"
      service.beta.kubernetes.io/aws-load-balancer-type: nlb
      service.beta.kubernetes.io/aws-load-balancer-cross-zone-load-balancing-enabled: "true"
EOF

echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}


OUTPUT_FILE="${DIR}/deploy/static/provider/aws/deploy-tls-termination.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local

    annotations:
      service.beta.kubernetes.io/aws-load-balancer-backend-protocol: http
      service.beta.kubernetes.io/aws-load-balancer-cross-zone-load-balancing-enabled: 'true'
      service.beta.kubernetes.io/aws-load-balancer-ssl-ports: "https"
      service.beta.kubernetes.io/aws-load-balancer-ssl-cert: "arn:aws:acm:us-west-2:XXXXXXXX:certificate/XXXXXX-XXXXXXX-XXXXXXX-XXXXXXXX"
      service.beta.kubernetes.io/aws-load-balancer-type: elb
      # Ensure the ELB idle timeout is less than nginx keep-alive timeout. By default,
      # NGINX keep-alive is set to 75s. If using WebSockets, the value will need to be
      # increased to '3600' to avoid any potential issues.
      service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout: "60"

    targetPorts:
      http: tohttps
      https: http

  # Configures the ports the nginx-controller listens on
  containerPort:
    http: 80
    https: 80
    tohttps: 2443

  config:
    proxy-real-ip-cidr: XXX.XXX.XXX/XX
    use-forwarded-headers: "true"
    http-snippet: |
      server {
        listen 2443;
        return 308 https://\$host\$request_uri;
      }
EOF

echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}

# Kind - https://kind.sigs.k8s.io/docs/user/ingress/
OUTPUT_FILE="${DIR}/deploy/static/provider/kind/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  hostPort:
    enabled: true
  terminationGracePeriodSeconds: 0
  service:
    type: NodePort

  nodeSelector:
    ingress-ready: "true"
  tolerations:
    - key: "node-role.kubernetes.io/master"
      operator: "Equal"
      effect: "NoSchedule"

  publishService:
    enabled: false
  extraArgs:
    publish-status-address: localhost
EOF

# Digital Ocean
echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}

OUTPUT_FILE="${DIR}/deploy/static/provider/do/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local
    annotations:
      service.beta.kubernetes.io/do-loadbalancer-enable-proxy-protocol: "true"
  config:
    use-proxy-protocol: "true"

EOF

# Scaleway
echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}

OUTPUT_FILE="${DIR}/deploy/static/provider/scw/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local
    annotations:
      service.beta.kubernetes.io/scw-loadbalancer-proxy-protocol-v2: "true"
  config:
    use-proxy-protocol: "true"

EOF

# Exoscale
echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}

OUTPUT_FILE="${DIR}/deploy/static/provider/exoscale/deploy.yaml"
cat << EOF | helm template $RELEASE_NAME ${DIR}/charts/ingress-nginx --namespace $NAMESPACE --api-versions $API_VERSIONS --values - | $DIR/hack/add-namespace.py $NAMESPACE > ${OUTPUT_FILE}
controller:
  kind: DaemonSet
  service:
    type: LoadBalancer
    externalTrafficPolicy: Local
    annotations:
      service.beta.kubernetes.io/exoscale-loadbalancer-name: "nginx-ingress-controller"
      service.beta.kubernetes.io/exoscale-loadbalancer-description: "NGINX Ingress Controller load balancer"
      service.beta.kubernetes.io/exoscale-loadbalancer-service-strategy: "source-hash"
      service.beta.kubernetes.io/exoscale-loadbalancer-service-healthcheck-mode: "tcp"
      service.beta.kubernetes.io/exoscale-loadbalancer-service-healthcheck-interval: "10s"
      service.beta.kubernetes.io/exoscale-loadbalancer-service-healthcheck-timeout: "3s"
      service.beta.kubernetes.io/exoscale-loadbalancer-service-healthcheck-retries: "1"
  publishService:
      enabled: true
EOF

echo "${NAMESPACE_VAR}
$(cat ${OUTPUT_FILE})" > ${OUTPUT_FILE}
