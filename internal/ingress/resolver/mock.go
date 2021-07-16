/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resolver

import (
	"errors"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/ingress-nginx/internal/ingress/defaults"
)

// Mock implements the Resolver interface
type Mock struct {
	ConfigMaps map[string]*apiv1.ConfigMap
}

// GetDefaultBackend returns the backend that must be used as default
func (m Mock) GetDefaultBackend() defaults.Backend {
	return defaults.Backend{}
}

// GetSecret searches for secrets contenating the namespace and name using a the character /
func (m Mock) GetSecret(string) (*apiv1.Secret, error) {
	return nil, nil
}

// GetAuthCertificate resolves a given secret name into an SSL certificate.
// The secret must contain 3 keys named:
//   ca.crt: contains the certificate chain used for authentication
func (m Mock) GetAuthCertificate(string) (*AuthSSLCert, error) {
	return nil, nil
}

func (m Mock) GetGlobalTLSAuth() GlobalTLSAuth {
	return GlobalTLSAuth{}
}

// GetService searches for services contenating the namespace and name using a the character /
func (m Mock) GetService(string) (*apiv1.Service, error) {
	return nil, nil
}

// GetConfigMap searches for configMaps contenating the namespace and name using a the character /
func (m Mock) GetConfigMap(name string) (*apiv1.ConfigMap, error) {
	if v, ok := m.ConfigMaps[name]; ok {
		return v, nil
	}
	return nil, errors.New("no configmap")
}
