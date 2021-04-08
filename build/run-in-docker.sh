#!/bin/bash

# Copyright 2018 The Kubernetes Authors.
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

# temporal directory for the /etc/ingress-controller directory
INGRESS_VOLUME=$(mktemp -d)

if [[ "$OSTYPE" == darwin* ]]; then
  INGRESS_VOLUME=/private$INGRESS_VOLUME
fi

function cleanup {
  rm -rf "${INGRESS_VOLUME}"
}
trap cleanup EXIT

E2E_IMAGE=${E2E_IMAGE:-k8s.gcr.io/ingress-nginx/e2e-test-runner:v20210326-gb52c538bb@sha256:b1b684ac3cc6a1ba68611707467fe2c9fe1c9c4a60f85e19ee10ea14b5343432}

DOCKER_OPTS=${DOCKER_OPTS:-}
DOCKER_IN_DOCKER_ENABLED=${DOCKER_IN_DOCKER_ENABLED:-}

KUBE_ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)

FLAGS=$@

PKG=k8s.io/ingress-nginx
ARCH=${ARCH:-}
if [[ -z "$ARCH" ]]; then
  ARCH=$(go env GOARCH)
fi

# create output directory as current user to avoid problem with docker.
mkdir -p "${KUBE_ROOT}/bin" "${KUBE_ROOT}/bin/${ARCH}"

if [[ "$DOCKER_IN_DOCKER_ENABLED" == "true" ]]; then
  /bin/bash -c "${FLAGS}"
else
  docker run                                            \
    --tty                                               \
    --rm                                                \
    ${DOCKER_OPTS}                                      \
    -e GOCACHE="/go/src/${PKG}/.cache"                  \
    -e DOCKER_IN_DOCKER_ENABLED="true"                  \
    -v "${HOME}/.kube:${HOME}/.kube"                    \
    -v "${KUBE_ROOT}:/go/src/${PKG}"                    \
    -v "${KUBE_ROOT}/bin/${ARCH}:/go/bin/linux_${ARCH}" \
    -v "/var/run/docker.sock:/var/run/docker.sock"      \
    -v "${INGRESS_VOLUME}:/etc/ingress-controller/"     \
    -w "/go/src/${PKG}"                                 \
    -u $(id -u ${USER}):$(id -g ${USER})                \
    ${E2E_IMAGE} /bin/bash -c "${FLAGS}"
fi
