/*
Copyright 2019 The Kubernetes Authors.

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

package settings

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	networking "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/ingress-nginx/internal/ingress/controller/class"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("[Flag] ingress-class", func() {
	f := framework.NewDefaultFramework("ingress-class")

	var doOnce sync.Once

	testIngressClassName := "test-new-ingress-class"

	ginkgo.BeforeEach(func() {
		f.NewEchoDeploymentWithReplicas(1)

		doOnce.Do(func() {
			f.KubeClientSet.RbacV1().ClusterRoles().Create(context.TODO(), &rbacv1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{Name: "ingress-nginx-class"},
				Rules: []rbacv1.PolicyRule{{
					APIGroups: []string{"networking.k8s.io"},
					Resources: []string{"ingressclasses"},
					Verbs:     []string{"get", "list", "watch"},
				}},
			}, metav1.CreateOptions{})

			f.KubeClientSet.RbacV1().ClusterRoleBindings().Create(context.TODO(), &rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ingress-nginx-class",
				},
				RoleRef: rbacv1.RoleRef{
					APIGroup: "rbac.authorization.k8s.io",
					Kind:     "ClusterRole",
					Name:     "ingress-nginx-class",
				},
			}, metav1.CreateOptions{})

			_, err := f.KubeClientSet.NetworkingV1().IngressClasses().
				Create(context.TODO(), &networking.IngressClass{
					ObjectMeta: metav1.ObjectMeta{
						Name: testIngressClassName,
					},
					Spec: networking.IngressClassSpec{
						Controller: k8s.IngressNGINXController,
					},
				}, metav1.CreateOptions{})

			if !apierrors.IsAlreadyExists(err) {
				assert.Nil(ginkgo.GinkgoT(), err, "creating IngressClass")
			}
		})
	})

	ginkgo.Context("Without a specific ingress-class", func() {
		ginkgo.It("should ignore Ingress with class", func() {
			invalidHost := "foo"
			annotations := map[string]string{
				class.IngressKey: "testclass",
			}
			ing := framework.NewSingleIngress(invalidHost, "/", invalidHost, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			validHost := "bar"
			ing = framework.NewSingleIngress(validHost, "/", validHost, f.Namespace, framework.EchoService, 80, nil)
			f.EnsureIngress(ing)

			f.WaitForNginxConfiguration(func(cfg string) bool {
				return !strings.Contains(cfg, "server_name foo") &&
					strings.Contains(cfg, "server_name bar")
			})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", invalidHost).
				Expect().
				Status(http.StatusNotFound)

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", validHost).
				Expect().
				Status(http.StatusOK)
		})
	})

	ginkgo.Context("With a specific ingress-class", func() {
		ginkgo.BeforeEach(func() {
			err := f.UpdateIngressControllerDeployment(func(deployment *appsv1.Deployment) error {
				args := []string{}
				for _, v := range deployment.Spec.Template.Spec.Containers[0].Args {
					if strings.Contains(v, "--ingress-class") {
						continue
					}

					args = append(args, v)
				}

				args = append(args, "--ingress-class=testclass")
				deployment.Spec.Template.Spec.Containers[0].Args = args
				_, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})

				return err
			})
			assert.Nil(ginkgo.GinkgoT(), err, "updating ingress controller deployment flags")
		})

		ginkgo.It("should ignore Ingress with no class", func() {
			invalidHost := "bar"

			ing := framework.NewSingleIngress(invalidHost, "/", invalidHost, f.Namespace, framework.EchoService, 80, nil)
			f.EnsureIngress(ing)

			validHost := "foo"
			annotations := map[string]string{
				class.IngressKey: "testclass",
			}
			ing = framework.NewSingleIngress(validHost, "/", validHost, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(validHost, func(cfg string) bool {
				return strings.Contains(cfg, "server_name foo")
			})

			f.WaitForNginxConfiguration(func(cfg string) bool {
				return !strings.Contains(cfg, "server_name bar")
			})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", validHost).
				Expect().
				Status(http.StatusOK)

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", invalidHost).
				Expect().
				Status(http.StatusNotFound)
		})

		ginkgo.It("should delete Ingress when class is removed", func() {
			host := "foo"
			annotations := map[string]string{
				class.IngressKey: "testclass",
			}
			ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, annotations)
			f.EnsureIngress(ing)

			f.WaitForNginxServer(host, func(cfg string) bool {
				return strings.Contains(cfg, "server_name foo")
			})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				Expect().
				Status(http.StatusOK)

			ing, err := f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
			assert.Nil(ginkgo.GinkgoT(), err)

			delete(ing.Annotations, class.IngressKey)
			_, err = f.KubeClientSet.NetworkingV1().Ingresses(ing.Namespace).Update(context.TODO(), ing, metav1.UpdateOptions{})
			assert.Nil(ginkgo.GinkgoT(), err)

			framework.Sleep()

			f.WaitForNginxConfiguration(func(cfg string) bool {
				return !strings.Contains(cfg, "server_name foo")
			})

			f.HTTPTestClient().
				GET("/").
				WithHeader("Host", host).
				Expect().
				Status(http.StatusNotFound)
		})
	})

	ginkgo.It("check scenarios for IngressClass and ingress.class annotation", func() {

		pod := f.GetIngressNGINXPod()

		crb, err := f.KubeClientSet.RbacV1().ClusterRoleBindings().Get(context.Background(), "ingress-nginx-class", metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "searching cluster role binding")

		// add service of current namespace
		crb.Subjects = append(crb.Subjects, rbacv1.Subject{
			APIGroup:  "",
			Kind:      "ServiceAccount",
			Name:      pod.Spec.ServiceAccountName,
			Namespace: f.Namespace,
		})

		_, err = f.KubeClientSet.RbacV1().ClusterRoleBindings().Update(context.Background(), crb, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "searching cluster role binding")

		err = f.UpdateIngressControllerDeployment(func(deployment *appsv1.Deployment) error {
			args := []string{}
			for _, v := range deployment.Spec.Template.Spec.Containers[0].Args {
				if strings.Contains(v, "--ingress-class") {
					continue
				}

				args = append(args, v)
			}

			args = append(args, fmt.Sprintf("--ingress-class=%v", testIngressClassName))
			deployment.Spec.Template.Spec.Containers[0].Args = args
			_, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
			return err
		})
		assert.Nil(ginkgo.GinkgoT(), err, "updating ingress controller deployment flags")

		host := "ingress.class"

		ginkgo.By("only having IngressClassName")
		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, nil)
		ing.Spec.IngressClassName = &testIngressClassName
		f.EnsureIngress(ing)

		f.WaitForNginxConfiguration(func(cfg string) bool {
			return strings.Contains(cfg, fmt.Sprintf("server_name %v", host))
		})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)

		ginkgo.By("only having ingress.class annotation")
		ing, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		ing.Annotations = map[string]string{
			class.IngressKey: testIngressClassName,
		}
		ing.Spec.IngressClassName = nil

		_, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Update(context.TODO(), ing, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		f.WaitForNginxConfiguration(func(cfg string) bool {
			return strings.Contains(cfg, fmt.Sprintf("server_name %v", host))
		})

		framework.Sleep()

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusOK)

		ginkgo.By("having an invalid ingress.class annotation and no IngressClassName")
		ing, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		ing.Annotations = map[string]string{
			class.IngressKey: "invalid",
		}
		ing.Spec.IngressClassName = nil

		_, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Update(context.TODO(), ing, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		framework.Sleep()

		f.WaitForNginxConfiguration(func(cfg string) bool {
			return !strings.Contains(cfg, fmt.Sprintf("server_name %v", host))
		})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusNotFound)

		ginkgo.By("not having ingress.class annotation and invalid IngressClassName")
		ing, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)
		ing.Annotations = map[string]string{}
		invalidClassName := "invalidclass"
		ing.Spec.IngressClassName = &invalidClassName

		_, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Update(context.TODO(), ing, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err)

		framework.Sleep()

		f.WaitForNginxConfiguration(func(cfg string) bool {
			return !strings.Contains(cfg, fmt.Sprintf("server_name %v", host))
		})

		f.HTTPTestClient().
			GET("/").
			WithHeader("Host", host).
			Expect().
			Status(http.StatusNotFound)
	})
})
