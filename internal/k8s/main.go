/*
Copyright 2015 The Kubernetes Authors.

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

package k8s

import (
	"context"
	"fmt"
	"os"
	"strings"

	"k8s.io/klog/v2"

	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/version"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// ParseNameNS parses a string searching a namespace and name
func ParseNameNS(input string) (string, string, error) {
	nsName := strings.Split(input, "/")
	if len(nsName) != 2 {
		return "", "", fmt.Errorf("invalid format (namespace/name) found in '%v'", input)
	}

	return nsName[0], nsName[1], nil
}

// GetNodeIPOrName returns the IP address or the name of a node in the cluster
func GetNodeIPOrName(kubeClient clientset.Interface, name string, useInternalIP bool) string {
	node, err := kubeClient.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.ErrorS(err, "Error getting node", "name", name)
		return ""
	}

	defaultOrInternalIP := ""
	for _, address := range node.Status.Addresses {
		if address.Type == apiv1.NodeInternalIP {
			if address.Address != "" {
				defaultOrInternalIP = address.Address
				break
			}
		}
	}

	if useInternalIP {
		return defaultOrInternalIP
	}

	for _, address := range node.Status.Addresses {
		if address.Type == apiv1.NodeExternalIP {
			if address.Address != "" {
				return address.Address
			}
		}
	}

	return defaultOrInternalIP
}

var (
	// IngressPodDetails hold information about the ingress-nginx pod
	IngressPodDetails *PodInfo
)

// PodInfo contains runtime information about the pod running the Ingres controller
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PodInfo struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

// GetIngressPod load the ingress-nginx pod
func GetIngressPod(kubeClient clientset.Interface) error {
	podName := os.Getenv("POD_NAME")
	podNs := os.Getenv("POD_NAMESPACE")

	if podName == "" || podNs == "" {
		return fmt.Errorf("unable to get POD information (missing POD_NAME or POD_NAMESPACE environment variable")
	}

	pod, err := kubeClient.CoreV1().Pods(podNs).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("unable to get POD information: %v", err)
	}

	IngressPodDetails = &PodInfo{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Pod"},
	}

	pod.ObjectMeta.DeepCopyInto(&IngressPodDetails.ObjectMeta)
	IngressPodDetails.SetLabels(pod.GetLabels())

	return nil
}

// MetaNamespaceKey knows how to make keys for API objects which implement meta.Interface.
func MetaNamespaceKey(obj interface{}) string {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Warning(err)
	}

	return key
}

// IsIngressV1Ready indicates if the running Kubernetes version is at least v1.19.0
var IsIngressV1Ready bool

// IngressNGINXController defines the valid value of IngressClass
// Controller field for ingress-nginx
const IngressNGINXController = "k8s.io/ingress-nginx"

// NetworkingIngressAvailable checks if the package "k8s.io/api/networking/v1"
// is available or not and if Ingress V1 is supported (k8s >= v1.19.0)
func NetworkingIngressAvailable(client clientset.Interface) bool {
	version119, _ := version.ParseGeneric("v1.19.0")

	serverVersion, err := client.Discovery().ServerVersion()
	if err != nil {
		return false
	}

	runningVersion, err := version.ParseGeneric(serverVersion.String())
	if err != nil {
		klog.ErrorS(err, "unexpected error parsing running Kubernetes version")
		return false
	}

	return runningVersion.AtLeast(version119)
}

// default path type is Prefix to not break existing definitions
var defaultPathType = networkingv1.PathTypePrefix

// SetDefaultNGINXPathType sets a default PathType when is not defined.
func SetDefaultNGINXPathType(ing *networkingv1.Ingress) {
	for _, rule := range ing.Spec.Rules {
		if rule.IngressRuleValue.HTTP == nil {
			continue
		}

		for idx := range rule.IngressRuleValue.HTTP.Paths {
			p := &rule.IngressRuleValue.HTTP.Paths[idx]
			if p.PathType == nil {
				p.PathType = &defaultPathType
			}

			if *p.PathType == networkingv1.PathTypeImplementationSpecific {
				p.PathType = &defaultPathType
			}
		}
	}
}
