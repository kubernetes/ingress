/*
Copyright 2018 The Kubernetes Authors.

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
	"log"
	"net"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("[Status] status update", func() {
	f := framework.NewDefaultFramework("status-update")
	host := "status-update"
	address := getHostIP()

	ginkgo.It("should update status field after client-go reconnection", func() {
		port, cmd, err := f.KubectlProxy(0)
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error starting kubectl proxy")

		err = f.UpdateIngressControllerDeployment(func(deployment *appsv1.Deployment) error {
			args := []string{}
			// flags --publish-service and --publish-status-address are mutually exclusive

			for _, v := range deployment.Spec.Template.Spec.Containers[0].Args {
				if strings.Contains(v, "--publish-service") {
					continue
				}

				if strings.Contains(v, "--update-status") {
					continue
				}

				args = append(args, v)
			}

			args = append(args, fmt.Sprintf("--apiserver-host=http://%s:%d", address.String(), port))
			args = append(args, "--publish-status-address=1.1.0.0")

			deployment.Spec.Template.Spec.Containers[0].Args = args
			_, err := f.KubeClientSet.AppsV1().Deployments(f.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
			return err
		})
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error updating ingress controller deployment flags")

		f.NewEchoDeploymentWithReplicas(1)

		ing := f.EnsureIngress(framework.NewSingleIngress(host, "/", host, f.Namespace, framework.EchoService, 80, nil))

		f.WaitForNginxConfiguration(
			func(cfg string) bool {
				return strings.Contains(cfg, fmt.Sprintf("server_name %s", host))
			})

		framework.Logf("waiting for leader election and initial status update")
		framework.Sleep(30 * time.Second)

		err = cmd.Process.Kill()
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error terminating kubectl proxy")

		ing, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error getting %s/%v Ingress", f.Namespace, host)

		ing.Status.LoadBalancer.Ingress = []apiv1.LoadBalancerIngress{}
		_, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).UpdateStatus(context.TODO(), ing, metav1.UpdateOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error cleaning Ingress status")
		framework.Sleep(10 * time.Second)

		err = f.KubeClientSet.CoreV1().
			ConfigMaps(f.Namespace).
			Delete(context.TODO(), "ingress-controller-leader", metav1.DeleteOptions{})
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error deleting leader election configmap")

		_, cmd, err = f.KubectlProxy(port)
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error starting kubectl proxy")
		defer func() {
			defer ginkgo.GinkgoRecover()

			if cmd != nil {
				err := cmd.Process.Kill()
				assert.Nil(ginkgo.GinkgoT(), err, "unexpected error terminating kubectl proxy")
			}
		}()

		err = wait.Poll(5*time.Second, 4*time.Minute, func() (done bool, err error) {
			ing, err = f.KubeClientSet.NetworkingV1().Ingresses(f.Namespace).Get(context.TODO(), host, metav1.GetOptions{})
			if err != nil {
				return false, nil
			}

			if len(ing.Status.LoadBalancer.Ingress) != 1 {
				return false, nil
			}

			return true, nil
		})
		assert.Nil(ginkgo.GinkgoT(), err, "unexpected error waiting for ingress status")
		assert.Equal(ginkgo.GinkgoT(), ing.Status.LoadBalancer.Ingress, ([]apiv1.LoadBalancerIngress{
			{IP: "1.1.0.0"},
		}))
	})
})

func getHostIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
