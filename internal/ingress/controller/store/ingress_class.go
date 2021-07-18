/*
Copyright 2021 The Kubernetes Authors.

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

package store

import (
	networking "k8s.io/api/networking/v1"
	"k8s.io/client-go/tools/cache"
)

// IngressClassLister makes a Store that lists IngressClass.
type IngressClassLister struct {
	cache.Store
}

// ByKey returns the Ingress matching key in the local Ingress Store.
func (il IngressClassLister) ByKey(key string) (*networking.IngressClass, error) {
	i, exists, err := il.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotExistsError(key)
	}
	return i.(*networking.IngressClass), nil
}
