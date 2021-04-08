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
	"context"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EchoService name of the deployment for the echo app
const EchoService = "echo"

// SlowEchoService name of the deployment for the echo app
const SlowEchoService = "slow-echo"

// HTTPBinService name of the deployment for the httpbin app
const HTTPBinService = "httpbin"

// NewEchoDeployment creates a new single replica deployment of the echoserver image in a particular namespace
func (f *Framework) NewEchoDeployment() {
	f.NewEchoDeploymentWithReplicas(1)
}

// NewEchoDeploymentWithReplicas creates a new deployment of the echoserver image in a particular namespace. Number of
// replicas is configurable
func (f *Framework) NewEchoDeploymentWithReplicas(replicas int) {
	f.NewEchoDeploymentWithNameAndReplicas(EchoService, replicas)
}

// NewEchoDeploymentWithNameAndReplicas creates a new deployment of the echoserver image in a particular namespace. Number of
// replicas is configurable and
// name is configurable
func (f *Framework) NewEchoDeploymentWithNameAndReplicas(name string, replicas int) {
	deployment := newDeployment(name, f.Namespace, "k8s.gcr.io/ingress-nginx/e2e-test-echo@sha256:b13e44f7bb6852b90633957e743e4a2b34f32f1694da556a9131b43950b8b2b1", 80, int32(replicas),
		nil,
		[]corev1.VolumeMount{},
		[]corev1.Volume{},
	)

	f.EnsureDeployment(deployment)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(80),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	f.EnsureService(service)

	err := WaitForEndpoints(f.KubeClientSet, DefaultTimeout, name, f.Namespace, replicas)
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")
}

// NewSlowEchoDeployment creates a new deployment of the slow echo server image in a particular namespace.
func (f *Framework) NewSlowEchoDeployment() {
	cfg := `#
events {
	worker_connections  1024;
	multi_accept on;
}

http {
	default_type 'text/plain';
	client_max_body_size 0;

	server {
		access_log on;
		access_log /dev/stdout;

		listen 80;

		location / {
			content_by_lua_block {
				ngx.print("ok")
			}
		}

		location ~ ^/sleep/(?<sleepTime>[0-9]+)$ {
			content_by_lua_block {
				ngx.sleep(ngx.var.sleepTime)
				ngx.print("ok after " .. ngx.var.sleepTime .. " seconds")
			}
		}
	}
}

`

	f.NGINXWithConfigDeployment(SlowEchoService, cfg)
}

// NGINXWithConfigDeployment creates an NGINX deployment using a configmap containing the nginx.conf configuration
func (f *Framework) NGINXWithConfigDeployment(name string, cfg string) {
	cfgMap := map[string]string{
		"nginx.conf": cfg,
	}

	_, err := f.KubeClientSet.CoreV1().ConfigMaps(f.Namespace).Create(context.TODO(), &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Data: cfgMap,
	}, metav1.CreateOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "creating configmap")

	deployment := newDeployment(name, f.Namespace, "k8s.gcr.io/ingress-nginx/nginx:v20210324-g8baef769d@sha256:fcfa3e9d1f8ec3141efedbf77cf659640f452a9c22165c78006ea462b84d06f6", 80, 1,
		nil,
		[]corev1.VolumeMount{
			{
				Name:      name,
				MountPath: "/etc/nginx/nginx.conf",
				SubPath:   "nginx.conf",
				ReadOnly:  true,
			},
		},
		[]corev1.Volume{
			{
				Name: name,
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: name,
						},
					},
				},
			},
		},
	)

	f.EnsureDeployment(deployment)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(80),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	f.EnsureService(service)

	err = WaitForEndpoints(f.KubeClientSet, DefaultTimeout, name, f.Namespace, 1)
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")
}

// NewGRPCBinDeployment creates a new deployment of the
// moul/grpcbin image for GRPC tests
func (f *Framework) NewGRPCBinDeployment() {
	name := "grpcbin"

	probe := &corev1.Probe{
		InitialDelaySeconds: 1,
		PeriodSeconds:       1,
		SuccessThreshold:    1,
		TimeoutSeconds:      1,
		Handler: corev1.Handler{
			TCPSocket: &corev1.TCPSocketAction{
				Port: intstr.FromInt(9000),
			},
		},
	}

	sel := map[string]string{
		"app": name,
	}

	f.EnsureDeployment(&appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: NewInt32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: sel,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: sel,
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: NewInt64(0),
					Containers: []corev1.Container{
						{
							Name:  name,
							Image: "moul/grpcbin",
							Env:   []corev1.EnvVar{},
							Ports: []corev1.ContainerPort{
								{
									Name:          "insecure",
									ContainerPort: 9000,
									Protocol:      corev1.ProtocolTCP,
								},
								{
									Name:          "secure",
									ContainerPort: 9001,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							ReadinessProbe: probe,
							LivenessProbe:  probe,
						},
					},
				},
			},
		},
	})

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "insecure",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   corev1.ProtocolTCP,
				},
				{
					Name:       "secure",
					Port:       9001,
					TargetPort: intstr.FromInt(9001),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: sel,
		},
	}

	f.EnsureService(service)

	err := WaitForEndpoints(f.KubeClientSet, DefaultTimeout, name, f.Namespace, 1)
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")
}

func newDeployment(name, namespace, image string, port int32, replicas int32, command []string,
	volumeMounts []corev1.VolumeMount, volumes []corev1.Volume) *appsv1.Deployment {
	probe := &corev1.Probe{
		InitialDelaySeconds: 2,
		PeriodSeconds:       1,
		SuccessThreshold:    1,
		TimeoutSeconds:      2,
		FailureThreshold:    6,
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Port: intstr.FromString("http"),
				Path: "/",
			},
		},
	}

	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: NewInt32(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: NewInt64(0),
					Containers: []corev1.Container{
						{
							Name:  name,
							Image: image,
							Env:   []corev1.EnvVar{},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: port,
								},
							},
							ReadinessProbe: probe,
							LivenessProbe:  probe,
							VolumeMounts:   volumeMounts,
						},
					},
					Volumes: volumes,
				},
			},
		},
	}

	if len(command) > 0 {
		d.Spec.Template.Spec.Containers[0].Command = command
	}

	return d
}

// NewHttpbinDeployment creates a new single replica deployment of the httpbin image in a particular namespace.
func (f *Framework) NewHttpbinDeployment() {
	f.NewDeployment(HTTPBinService, "k8s.gcr.io/ingress-nginx/e2e-test-httpbin@sha256:c6372ef57a775b95f18e19d4c735a9819f2e7bb4641e5e3f27287d831dfeb7e8", 80, 1)
}

// NewDeployment creates a new deployment in a particular namespace.
func (f *Framework) NewDeployment(name, image string, port int32, replicas int32) {
	deployment := newDeployment(name, f.Namespace, image, port, replicas, nil, nil, nil)

	f.EnsureDeployment(deployment)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt(int(port)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	f.EnsureService(service)

	err := WaitForEndpoints(f.KubeClientSet, DefaultTimeout, name, f.Namespace, int(replicas))
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for endpoints to become ready")
}

// DeleteDeployment deletes a deployment with a particular name and waits for the pods to be deleted
func (f *Framework) DeleteDeployment(name string) error {
	d, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "getting deployment")

	grace := int64(0)
	err = f.KubeClientSet.AppsV1().Deployments(f.Namespace).Delete(context.TODO(), name, metav1.DeleteOptions{
		GracePeriodSeconds: &grace,
	})
	assert.Nil(ginkgo.GinkgoT(), err, "deleting deployment")

	return waitForPodsDeleted(f.KubeClientSet, 2*time.Minute, f.Namespace, metav1.ListOptions{
		LabelSelector: labelSelectorToString(d.Spec.Selector.MatchLabels),
	})
}

// ScaleDeploymentToZero scales a deployment with a particular name and waits for the pods to be deleted
func (f *Framework) ScaleDeploymentToZero(name string) {
	d, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Get(context.TODO(), name, metav1.GetOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "getting deployment")
	assert.NotNil(ginkgo.GinkgoT(), d, "expected a deployment but none returned")

	d.Spec.Replicas = NewInt32(0)

	d, err = f.KubeClientSet.AppsV1().Deployments(f.Namespace).Update(context.TODO(), d, metav1.UpdateOptions{})
	assert.Nil(ginkgo.GinkgoT(), err, "getting deployment")
	assert.NotNil(ginkgo.GinkgoT(), d, "expected a deployment but none returned")

	err = WaitForEndpoints(f.KubeClientSet, DefaultTimeout, name, f.Namespace, 0)
	assert.Nil(ginkgo.GinkgoT(), err, "waiting for no endpoints")
}

// UpdateIngressControllerDeployment updates the ingress-nginx deployment
func (f *Framework) UpdateIngressControllerDeployment(fn func(deployment *appsv1.Deployment) error) error {
	err := UpdateDeployment(f.KubeClientSet, f.Namespace, "nginx-ingress-controller", 1, fn)
	if err != nil {
		return err
	}

	return f.updateIngressNGINXPod()
}
