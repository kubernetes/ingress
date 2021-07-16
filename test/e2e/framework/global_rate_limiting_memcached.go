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

package framework

import (
	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// NewGlobalRateLimitMemcachedDeployment creates a new single replica
// deployment of a memcached instance in a particular namespace
func (f *Framework) NewGlobalRateLimitMemcachedDeployment() {
	f.NewNewGlobalRateLimitMemcachedDeploymentWithReplicas(1)
}

// NewNewGlobalRateLimitMemcachedDeploymentWithReplicas creates a new deployment of the
// of a memcached instance in a particular namespace. Number of replicas is configurable
func (f *Framework) NewNewGlobalRateLimitMemcachedDeploymentWithReplicas(replicas int32) {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: f.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: NewInt32(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "memcached",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "memcached",
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: NewInt64(0),
					Containers: []corev1.Container{
						{
							Name:  "memcached",
							Image: "memcached:1.6.9",
							Env:   []corev1.EnvVar{},
							Ports: []corev1.ContainerPort{
								{
									Name:          "memcached",
									ContainerPort: 11211,
								},
							},
						},
					},
				},
			},
		},
	}

	d := f.EnsureDeployment(deployment)

	err := waitForPodsReady(f.KubeClientSet, DefaultTimeout, int(replicas), f.Namespace, metav1.ListOptions{
		LabelSelector: fields.SelectorFromSet(fields.Set(d.Spec.Template.ObjectMeta.Labels)).String(),
	})
	assert.Nil(ginkgo.GinkgoT(), err, "failed to wait for to become ready")

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: f.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "memcached",
					Port:       11211,
					TargetPort: intstr.FromInt(11211),
					Protocol:   "TCP",
				},
			},
			Selector: map[string]string{
				"app": "memcached",
			},
		},
	}

	f.EnsureService(service)
}
