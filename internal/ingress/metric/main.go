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

package metric

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/ingress-nginx/internal/ingress"
	"k8s.io/ingress-nginx/internal/ingress/metric/collectors"
)

// Collector defines the interface for a metric collector
type Collector interface {
	ConfigSuccess(uint64, bool)

	IncReloadCount()
	IncReloadErrorCount()

	OnStartedLeading(string)
	OnStoppedLeading(string)

	IncCheckCount(string, string)
	IncCheckErrorCount(string, string)

	RemoveMetrics(ingresses, endpoints []string)

	SetSSLExpireTime([]*ingress.Server)

	// SetHosts sets the hostnames that are being served by the ingress controller
	SetHosts(sets.String)

	Start()
	Stop()
}

type collector struct {
	nginxStatus  collectors.NGINXStatusCollector
	nginxProcess collectors.NGINXProcessCollector

	ingressController *collectors.Controller

	socket *collectors.SocketCollector

	registry *prometheus.Registry
}

// NewCollector creates a new metric collector the for ingress controller
func NewCollector(metricsPerHost bool, registry *prometheus.Registry, ingressclass string) (Collector, error) {
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		podNamespace = "default"
	}

	podName := os.Getenv("POD_NAME")

	nc, err := collectors.NewNGINXStatus(podName, podNamespace, ingressclass)
	if err != nil {
		return nil, err
	}

	pc, err := collectors.NewNGINXProcess(podName, podNamespace, ingressclass)
	if err != nil {
		return nil, err
	}

	s, err := collectors.NewSocketCollector(podName, podNamespace, ingressclass, metricsPerHost)
	if err != nil {
		return nil, err
	}

	ic := collectors.NewController(podName, podNamespace, ingressclass)

	return Collector(&collector{
		nginxStatus:  nc,
		nginxProcess: pc,

		ingressController: ic,

		socket: s,

		registry: registry,
	}), nil
}

func (c *collector) ConfigSuccess(hash uint64, success bool) {
	c.ingressController.ConfigSuccess(hash, success)
}

func (c *collector) IncCheckCount(namespace string, name string) {
	c.ingressController.IncCheckCount(namespace, name)
}

func (c *collector) IncCheckErrorCount(namespace string, name string) {
	c.ingressController.IncCheckErrorCount(namespace, name)
}

func (c *collector) IncReloadCount() {
	c.ingressController.IncReloadCount()
}

func (c *collector) IncReloadErrorCount() {
	c.ingressController.IncReloadErrorCount()
}

func (c *collector) RemoveMetrics(ingresses, hosts []string) {
	c.socket.RemoveMetrics(ingresses, c.registry)
	c.ingressController.RemoveMetrics(hosts, c.registry)
}

func (c *collector) Start() {
	c.registry.MustRegister(c.nginxStatus)
	c.registry.MustRegister(c.nginxProcess)
	c.registry.MustRegister(c.ingressController)
	c.registry.MustRegister(c.socket)

	// the default nginx.conf does not contains
	// a server section with the status port
	go func() {
		time.Sleep(5 * time.Second)
		c.nginxStatus.Start()
	}()
	go c.nginxProcess.Start()
	go c.socket.Start()
}

func (c *collector) Stop() {
	c.registry.Unregister(c.nginxStatus)
	c.registry.Unregister(c.nginxProcess)
	c.registry.Unregister(c.ingressController)
	c.registry.Unregister(c.socket)

	c.nginxStatus.Stop()
	c.nginxProcess.Stop()
	c.socket.Stop()
}

func (c *collector) SetSSLExpireTime(servers []*ingress.Server) {
	if !isLeader() {
		return
	}

	klog.V(2).InfoS("Updating ssl expiration metrics")
	c.ingressController.SetSSLExpireTime(servers)
}

func (c *collector) SetHosts(hosts sets.String) {
	c.socket.SetHosts(hosts)
}

// OnStartedLeading indicates the pod was elected as the leader
func (c *collector) OnStartedLeading(electionID string) {
	setLeader(true)
	c.ingressController.OnStartedLeading(electionID)
}

// OnStoppedLeading indicates the pod stopped being the leader
func (c *collector) OnStoppedLeading(electionID string) {
	setLeader(false)
	c.ingressController.OnStoppedLeading(electionID)
	c.ingressController.RemoveAllSSLExpireMetrics(c.registry)
}

var (
	currentLeader uint32
)

func setLeader(leader bool) {
	var i uint32
	if leader {
		i = 1
	}

	atomic.StoreUint32(&currentLeader, i)
}

func isLeader() bool {
	return atomic.LoadUint32(&currentLeader) != 0
}
