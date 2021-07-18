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

package status

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	apiv1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"

	"k8s.io/ingress-nginx/internal/ingress"
	"k8s.io/ingress-nginx/internal/ingress/controller/ingressclass"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/ingress-nginx/internal/task"
)

func buildLoadBalancerIngressByIP() []apiv1.LoadBalancerIngress {
	return []apiv1.LoadBalancerIngress{
		{
			IP:       "10.0.0.1",
			Hostname: "foo1",
		},
		{
			IP:       "10.0.0.2",
			Hostname: "foo2",
		},
		{
			IP:       "10.0.0.3",
			Hostname: "",
		},
		{
			IP:       "",
			Hostname: "foo4",
		},
	}
}

func buildSimpleClientSet() *testclient.Clientset {
	return testclient.NewSimpleClientset(
		&apiv1.PodList{Items: []apiv1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo1",
					Namespace: apiv1.NamespaceDefault,
					Labels: map[string]string{
						"label_sig": "foo_pod",
					},
				},
				Spec: apiv1.PodSpec{
					NodeName: "foo_node_2",
				},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodRunning,
					Conditions: []apiv1.PodCondition{
						{
							Type:   apiv1.PodReady,
							Status: apiv1.ConditionTrue,
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo1-unknown",
					Namespace: apiv1.NamespaceDefault,
				},
				Spec: apiv1.PodSpec{
					NodeName: "foo_node_1",
				},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodUnknown,
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo2",
					Namespace: apiv1.NamespaceDefault,
					Labels: map[string]string{
						"label_sig": "foo_no",
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo3",
					Namespace: metav1.NamespaceSystem,
					Labels: map[string]string{
						"label_sig": "foo_pod",
					},
				},
				Spec: apiv1.PodSpec{
					NodeName: "foo_node_2",
				},
				Status: apiv1.PodStatus{
					Phase: apiv1.PodRunning,
					Conditions: []apiv1.PodCondition{
						{
							Type:   apiv1.PodReady,
							Status: apiv1.ConditionTrue,
						},
					},
				},
			},
		}},
		&apiv1.ServiceList{Items: []apiv1.Service{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: apiv1.NamespaceDefault,
				},
				Status: apiv1.ServiceStatus{
					LoadBalancer: apiv1.LoadBalancerStatus{
						Ingress: buildLoadBalancerIngressByIP(),
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo_non_exist",
					Namespace: apiv1.NamespaceDefault,
				},
			},
		}},
		&apiv1.NodeList{Items: []apiv1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo_node_1",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						}, {
							Type:    apiv1.NodeExternalIP,
							Address: "10.0.0.2",
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo_node_2",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "11.0.0.1",
						},
						{
							Type:    apiv1.NodeExternalIP,
							Address: "11.0.0.2",
						},
					},
				},
			},
		}},
		&apiv1.EndpointsList{Items: []apiv1.Endpoints{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ingress-controller-leader",
					Namespace: apiv1.NamespaceDefault,
				},
			}}},
		&networking.IngressList{Items: buildExtensionsIngresses()},
	)
}

func fakeSynFn(interface{}) error {
	return nil
}

func buildExtensionsIngresses() []networking.Ingress {
	return []networking.Ingress{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo_ingress_1",
				Namespace: apiv1.NamespaceDefault,
			},
			Status: networking.IngressStatus{
				LoadBalancer: apiv1.LoadBalancerStatus{
					Ingress: []apiv1.LoadBalancerIngress{
						{
							IP:       "10.0.0.1",
							Hostname: "foo1",
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo_ingress_different_class",
				Namespace: metav1.NamespaceDefault,
				Annotations: map[string]string{
					ingressclass.IngressKey: "no-nginx",
				},
			},
			Status: networking.IngressStatus{
				LoadBalancer: apiv1.LoadBalancerStatus{
					Ingress: []apiv1.LoadBalancerIngress{
						{
							IP:       "0.0.0.0",
							Hostname: "foo.bar.com",
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo_ingress_2",
				Namespace: apiv1.NamespaceDefault,
			},
			Status: networking.IngressStatus{
				LoadBalancer: apiv1.LoadBalancerStatus{
					Ingress: []apiv1.LoadBalancerIngress{},
				},
			},
		},
	}
}

type testIngressLister struct {
}

func (til *testIngressLister) ListIngresses() []*ingress.Ingress {
	var ingresses []*ingress.Ingress
	ingresses = append(ingresses, &ingress.Ingress{
		Ingress: networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo_ingress_non_01",
				Namespace: apiv1.NamespaceDefault,
			}}})

	ingresses = append(ingresses, &ingress.Ingress{
		Ingress: networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo_ingress_1",
				Namespace: apiv1.NamespaceDefault,
			},
			Status: networking.IngressStatus{
				LoadBalancer: apiv1.LoadBalancerStatus{
					Ingress: buildLoadBalancerIngressByIP(),
				},
			},
		}})

	return ingresses
}

func buildIngressLister() ingressLister {
	return &testIngressLister{}
}

func buildStatusSync() statusSync {
	return statusSync{
		syncQueue: task.NewTaskQueue(fakeSynFn),
		Config: Config{
			Client:         buildSimpleClientSet(),
			PublishService: apiv1.NamespaceDefault + "/" + "foo",
			IngressLister:  buildIngressLister(),
		},
	}
}

func TestStatusActions(t *testing.T) {
	// make sure election can be created
	os.Setenv("POD_NAME", "foo1")
	os.Setenv("POD_NAMESPACE", apiv1.NamespaceDefault)
	c := Config{
		Client:                 buildSimpleClientSet(),
		PublishService:         "",
		IngressLister:          buildIngressLister(),
		UpdateStatusOnShutdown: true,
	}

	k8s.IngressPodDetails = &k8s.PodInfo{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo_base_pod",
			Namespace: apiv1.NamespaceDefault,
			Labels: map[string]string{
				"label_sig": "foo_pod",
			},
		},
	}

	// create object
	fkSync := NewStatusSyncer(c)
	if fkSync == nil {
		t.Fatalf("expected a valid Sync")
	}

	fk := fkSync.(statusSync)

	// start it and wait for the election and syn actions
	stopCh := make(chan struct{})
	defer close(stopCh)

	go fk.Run(stopCh)
	//  wait for the election
	time.Sleep(100 * time.Millisecond)
	// execute sync
	fk.sync("just-test")
	// PublishService is empty, so the running address is: ["11.0.0.2"]
	// after updated, the ingress's ip should only be "11.0.0.2"
	newIPs := []apiv1.LoadBalancerIngress{{
		IP: "11.0.0.2",
	}}
	fooIngress1, err1 := fk.Client.NetworkingV1().Ingresses(apiv1.NamespaceDefault).Get(context.TODO(), "foo_ingress_1", metav1.GetOptions{})
	if err1 != nil {
		t.Fatalf("unexpected error")
	}
	fooIngress1CurIPs := fooIngress1.Status.LoadBalancer.Ingress
	if !ingressSliceEqual(fooIngress1CurIPs, newIPs) {
		t.Fatalf("returned %v but expected %v", fooIngress1CurIPs, newIPs)
	}

	time.Sleep(1 * time.Second)

	// execute shutdown
	fk.Shutdown()
	// ingress should be empty
	newIPs2 := []apiv1.LoadBalancerIngress{}
	fooIngress2, err2 := fk.Client.NetworkingV1().Ingresses(apiv1.NamespaceDefault).Get(context.TODO(), "foo_ingress_1", metav1.GetOptions{})
	if err2 != nil {
		t.Fatalf("unexpected error")
	}
	fooIngress2CurIPs := fooIngress2.Status.LoadBalancer.Ingress
	if !ingressSliceEqual(fooIngress2CurIPs, newIPs2) {
		t.Fatalf("returned %v but expected %v", fooIngress2CurIPs, newIPs2)
	}

	oic, err := fk.Client.NetworkingV1().Ingresses(metav1.NamespaceDefault).Get(context.TODO(), "foo_ingress_different_class", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("unexpected error")
	}
	if oic.Status.LoadBalancer.Ingress[0].IP != "0.0.0.0" && oic.Status.LoadBalancer.Ingress[0].Hostname != "foo.bar.com" {
		t.Fatalf("invalid ingress status for rule with different class")
	}
}

func TestCallback(t *testing.T) {
	buildStatusSync()
}

func TestKeyfunc(t *testing.T) {
	fk := buildStatusSync()

	i := "foo_base_pod"
	r, err := fk.keyfunc(i)

	if err != nil {
		t.Fatalf("unexpected error")
	}
	if r != i {
		t.Errorf("returned %v but expected %v", r, i)
	}
}

func TestRunningAddressesWithPublishService(t *testing.T) {
	testCases := map[string]struct {
		fakeClient  *testclient.Clientset
		expected    []string
		errExpected bool
	}{
		"service type ClusterIP": {
			testclient.NewSimpleClientset(
				&apiv1.PodList{Items: []apiv1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.PodSpec{
							NodeName: "foo_node",
						},
						Status: apiv1.PodStatus{
							Phase: apiv1.PodRunning,
						},
					},
				},
				},
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.ServiceSpec{
							Type:      apiv1.ServiceTypeClusterIP,
							ClusterIP: "1.1.1.1",
						},
					},
				},
				},
			),
			[]string{"1.1.1.1"},
			false,
		},
		"service type NodePort": {
			testclient.NewSimpleClientset(
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.ServiceSpec{
							Type:      apiv1.ServiceTypeNodePort,
							ClusterIP: "1.1.1.1",
						},
					},
				},
				},
			),
			[]string{"1.1.1.1"},
			false,
		},
		"service type ExternalName": {
			testclient.NewSimpleClientset(
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.ServiceSpec{
							Type:         apiv1.ServiceTypeExternalName,
							ExternalName: "foo.bar",
						},
					},
				},
				},
			),
			[]string{"foo.bar"},
			false,
		},
		"service type LoadBalancer": {
			testclient.NewSimpleClientset(
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.ServiceSpec{
							Type: apiv1.ServiceTypeLoadBalancer,
						},
						Status: apiv1.ServiceStatus{
							LoadBalancer: apiv1.LoadBalancerStatus{
								Ingress: []apiv1.LoadBalancerIngress{
									{
										IP: "10.0.0.1",
									},
									{
										IP:       "",
										Hostname: "foo",
									},
								},
							},
						},
					},
				},
				},
			),
			[]string{"10.0.0.1", "foo"},
			false,
		},
		"service type LoadBalancer with same externalIP and ingress IP": {
			testclient.NewSimpleClientset(
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
						Spec: apiv1.ServiceSpec{
							Type:        apiv1.ServiceTypeLoadBalancer,
							ExternalIPs: []string{"10.0.0.1"},
						},
						Status: apiv1.ServiceStatus{
							LoadBalancer: apiv1.LoadBalancerStatus{
								Ingress: []apiv1.LoadBalancerIngress{
									{
										IP: "10.0.0.1",
									},
								},
							},
						},
					},
				},
				},
			),
			[]string{"10.0.0.1"},
			false,
		},
		"invalid service type": {
			testclient.NewSimpleClientset(
				&apiv1.ServiceList{Items: []apiv1.Service{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "foo",
							Namespace: apiv1.NamespaceDefault,
						},
					},
				},
				},
			),
			nil,
			true,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {

			fk := buildStatusSync()
			fk.Config.Client = tc.fakeClient

			ra, err := fk.runningAddresses()
			if err != nil {
				if tc.errExpected {
					return
				}

				t.Fatalf("%v: unexpected error obtaining running address/es: %v", title, err)
			}

			if ra == nil {
				t.Fatalf("returned nil but expected valid []string")
			}

			if !reflect.DeepEqual(tc.expected, ra) {
				t.Errorf("returned %v but expected %v", ra, tc.expected)
			}
		})
	}
}

func TestRunningAddressesWithPods(t *testing.T) {
	fk := buildStatusSync()
	fk.PublishService = ""

	r, _ := fk.runningAddresses()
	if r == nil {
		t.Fatalf("returned nil but expected valid []string")
	}
	rl := len(r)
	if len(r) != 1 {
		t.Fatalf("returned %v but expected %v", rl, 1)
	}
	rv := r[0]
	if rv != "11.0.0.2" {
		t.Errorf("returned %v but expected %v", rv, "11.0.0.2")
	}
}

func TestRunningAddressesWithPublishStatusAddress(t *testing.T) {
	fk := buildStatusSync()
	fk.PublishStatusAddress = "127.0.0.1"

	ra, _ := fk.runningAddresses()
	if ra == nil {
		t.Fatalf("returned nil but expected valid []string")
	}
	rl := len(ra)
	if len(ra) != 1 {
		t.Errorf("returned %v but expected %v", rl, 1)
	}
	rv := ra[0]
	if rv != "127.0.0.1" {
		t.Errorf("returned %v but expected %v", rv, "127.0.0.1")
	}
}

func TestRunningAddressesWithPublishStatusAddresses(t *testing.T) {
	fk := buildStatusSync()
	fk.PublishStatusAddress = "127.0.0.1,1.1.1.1"

	ra, _ := fk.runningAddresses()
	if ra == nil {
		t.Fatalf("returned nil but expected valid []string")
	}
	rl := len(ra)
	if len(ra) != 2 {
		t.Errorf("returned %v but expected %v", rl, 2)
	}
	rv := ra[0]
	rv2 := ra[1]
	if rv != "127.0.0.1" {
		t.Errorf("returned %v but expected %v", rv, "127.0.0.1")
	}
	if rv2 != "1.1.1.1" {
		t.Errorf("returned %v but expected %v", rv2, "1.1.1.1")
	}
}

func TestRunningAddressesWithPublishStatusAddressesAndSpaces(t *testing.T) {
	fk := buildStatusSync()
	fk.PublishStatusAddress = "127.0.0.1,  1.1.1.1"

	ra, _ := fk.runningAddresses()
	if ra == nil {
		t.Fatalf("returned nil but expected valid []string")
	}
	rl := len(ra)
	if len(ra) != 2 {
		t.Errorf("returned %v but expected %v", rl, 2)
	}
	rv := ra[0]
	rv2 := ra[1]
	if rv != "127.0.0.1" {
		t.Errorf("returned %v but expected %v", rv, "127.0.0.1")
	}
	if rv2 != "1.1.1.1" {
		t.Errorf("returned %v but expected %v", rv2, "1.1.1.1")
	}
}

func TestSliceToStatus(t *testing.T) {
	fkEndpoints := []string{
		"10.0.0.1",
		"2001:db8::68",
		"opensource-k8s-ingress",
	}

	r := sliceToStatus(fkEndpoints)

	if r == nil {
		t.Fatalf("returned nil but expected a valid []apiv1.LoadBalancerIngress")
	}
	rl := len(r)
	if rl != 3 {
		t.Fatalf("returned %v but expected %v", rl, 3)
	}
	re1 := r[0]
	if re1.Hostname != "opensource-k8s-ingress" {
		t.Fatalf("returned %v but expected %v", re1, apiv1.LoadBalancerIngress{Hostname: "opensource-k8s-ingress"})
	}
	re2 := r[1]
	if re2.IP != "10.0.0.1" {
		t.Fatalf("returned %v but expected %v", re2, apiv1.LoadBalancerIngress{IP: "10.0.0.1"})
	}
	re3 := r[2]
	if re3.IP != "2001:db8::68" {
		t.Fatalf("returned %v but expected %v", re3, apiv1.LoadBalancerIngress{IP: "2001:db8::68"})
	}
}

func TestIngressSliceEqual(t *testing.T) {
	fk1 := buildLoadBalancerIngressByIP()
	fk2 := append(buildLoadBalancerIngressByIP(), apiv1.LoadBalancerIngress{
		IP:       "10.0.0.5",
		Hostname: "foo5",
	})
	fk3 := buildLoadBalancerIngressByIP()
	fk3[0].Hostname = "foo_no_01"
	fk4 := buildLoadBalancerIngressByIP()
	fk4[2].IP = "11.0.0.3"

	fooTests := []struct {
		lhs []apiv1.LoadBalancerIngress
		rhs []apiv1.LoadBalancerIngress
		er  bool
	}{
		{fk1, fk1, true},
		{fk2, fk1, false},
		{fk3, fk1, false},
		{fk4, fk1, false},
		{fk1, nil, false},
		{nil, nil, true},
		{[]apiv1.LoadBalancerIngress{}, []apiv1.LoadBalancerIngress{}, true},
	}

	for _, fooTest := range fooTests {
		r := ingressSliceEqual(fooTest.lhs, fooTest.rhs)
		if r != fooTest.er {
			t.Errorf("returned %v but expected %v", r, fooTest.er)
		}
	}
}
