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

package store

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/eapache/channels"
	corev1 "k8s.io/api/core/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	clientcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/utils/pointer"

	"k8s.io/ingress-nginx/internal/file"
	"k8s.io/ingress-nginx/internal/ingress"
	"k8s.io/ingress-nginx/internal/ingress/annotations"
	"k8s.io/ingress-nginx/internal/ingress/annotations/class"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	ngx_config "k8s.io/ingress-nginx/internal/ingress/controller/config"
	ngx_template "k8s.io/ingress-nginx/internal/ingress/controller/template"
	"k8s.io/ingress-nginx/internal/ingress/defaults"
	"k8s.io/ingress-nginx/internal/ingress/errors"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
	"k8s.io/ingress-nginx/internal/k8s"
	"k8s.io/ingress-nginx/internal/nginx"
)

// IngressFilterFunc decides if an Ingress should be omitted or not
type IngressFilterFunc func(*ingress.Ingress) bool

// Storer is the interface that wraps the required methods to gather information
// about ingresses, services, secrets and ingress annotations.
type Storer interface {
	// GetBackendConfiguration returns the nginx configuration stored in a configmap
	GetBackendConfiguration() ngx_config.Configuration

	// GetConfigMap returns the ConfigMap matching key.
	GetConfigMap(key string) (*corev1.ConfigMap, error)

	// GetSecret returns the Secret matching key.
	GetSecret(key string) (*corev1.Secret, error)

	// GetService returns the Service matching key.
	GetService(key string) (*corev1.Service, error)

	// GetServiceEndpoints returns the Endpoints of a Service matching key.
	GetServiceEndpoints(key string) (*corev1.Endpoints, error)

	// ListIngresses returns a list of all Ingresses in the store.
	ListIngresses() []*ingress.Ingress

	// GetLocalSSLCert returns the local copy of a SSLCert
	GetLocalSSLCert(name string) (*ingress.SSLCert, error)

	// ListLocalSSLCerts returns the list of local SSLCerts
	ListLocalSSLCerts() []*ingress.SSLCert

	// GetAuthCertificate resolves a given secret name into an SSL certificate.
	// The secret must contain 3 keys named:
	//   ca.crt: contains the certificate chain used for authentication
	GetAuthCertificate(string) (*resolver.AuthSSLCert, error)

	// GetDefaultBackend returns the default backend configuration
	GetDefaultBackend() defaults.Backend

	// Run initiates the synchronization of the controllers
	Run(stopCh chan struct{})
}

// EventType type of event associated with an informer
type EventType string

const (
	// CreateEvent event associated with new objects in an informer
	CreateEvent EventType = "CREATE"
	// UpdateEvent event associated with an object update in an informer
	UpdateEvent EventType = "UPDATE"
	// DeleteEvent event associated when an object is removed from an informer
	DeleteEvent EventType = "DELETE"
	// ConfigurationEvent event associated when a controller configuration object is created or updated
	ConfigurationEvent EventType = "CONFIGURATION"
)

// Event holds the context of an event.
type Event struct {
	Type EventType
	Obj  interface{}
}

// Informer defines the required SharedIndexInformers that interact with the API server.
type Informer struct {
	Ingress   cache.SharedIndexInformer
	Endpoint  cache.SharedIndexInformer
	Service   cache.SharedIndexInformer
	Secret    cache.SharedIndexInformer
	ConfigMap cache.SharedIndexInformer
}

// Lister contains object listers (stores).
type Lister struct {
	Ingress               IngressLister
	Service               ServiceLister
	Endpoint              EndpointLister
	Secret                SecretLister
	ConfigMap             ConfigMapLister
	IngressWithAnnotation IngressWithAnnotationsLister
}

// NotExistsError is returned when an object does not exist in a local store.
type NotExistsError string

// Error implements the error interface.
func (e NotExistsError) Error() string {
	return fmt.Sprintf("no object matching key %q in local store", string(e))
}

// Run initiates the synchronization of the informers against the API server.
func (i *Informer) Run(stopCh chan struct{}) {
	go i.Secret.Run(stopCh)
	go i.Endpoint.Run(stopCh)
	go i.Service.Run(stopCh)
	go i.ConfigMap.Run(stopCh)

	// wait for all involved caches to be synced before processing items
	// from the queue
	if !cache.WaitForCacheSync(stopCh,
		i.Endpoint.HasSynced,
		i.Service.HasSynced,
		i.Secret.HasSynced,
		i.ConfigMap.HasSynced,
	) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
	}

	// in big clusters, deltas can keep arriving even after HasSynced
	// functions have returned 'true'
	time.Sleep(1 * time.Second)

	// we can start syncing ingress objects only after other caches are
	// ready, because ingress rules require content from other listers, and
	// 'add' events get triggered in the handlers during caches population.
	go i.Ingress.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh,
		i.Ingress.HasSynced,
	) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
	}
}

// k8sStore internal Storer implementation using informers and thread safe stores
type k8sStore struct {
	// backendConfig contains the running configuration from the configmap
	// this is required because this rarely changes but is a very expensive
	// operation to execute in each OnUpdate invocation
	backendConfig ngx_config.Configuration

	// informer contains the cache Informers
	informers *Informer

	// listers contains the cache.Store interfaces used in the ingress controller
	listers *Lister

	// sslStore local store of SSL certificates (certificates used in ingress)
	// this is required because the certificates must be present in the
	// container filesystem
	sslStore *SSLCertTracker

	annotations annotations.Extractor

	// secretIngressMap contains information about which ingress references a
	// secret in the annotations.
	secretIngressMap ObjectRefMap

	// updateCh
	updateCh *channels.RingChannel

	// syncSecretMu protects against simultaneous invocations of syncSecret
	syncSecretMu *sync.Mutex

	// backendConfigMu protects against simultaneous read/write of backendConfig
	backendConfigMu *sync.RWMutex

	defaultSSLCertificate string
}

// New creates a new object store to be used in the ingress controller
func New(
	namespace, configmap, tcp, udp, defaultSSLCertificate string,
	resyncPeriod time.Duration,
	client clientset.Interface,
	updateCh *channels.RingChannel,
	disableCatchAll bool) Storer {

	store := &k8sStore{
		informers:             &Informer{},
		listers:               &Lister{},
		sslStore:              NewSSLCertTracker(),
		updateCh:              updateCh,
		backendConfig:         ngx_config.NewDefault(),
		syncSecretMu:          &sync.Mutex{},
		backendConfigMu:       &sync.RWMutex{},
		secretIngressMap:      NewObjectRefMap(),
		defaultSSLCertificate: defaultSSLCertificate,
	}

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&clientcorev1.EventSinkImpl{
		Interface: client.CoreV1().Events(namespace),
	})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{
		Component: "nginx-ingress-controller",
	})

	// k8sStore fulfills resolver.Resolver interface
	store.annotations = annotations.NewAnnotationExtractor(store)

	store.listers.IngressWithAnnotation.Store = cache.NewStore(cache.DeletionHandlingMetaNamespaceKeyFunc)

	// As we currently do not filter out kubernetes objects we list, we can
	// retrieve a huge amount of data from the API server.
	// In a cluster using HELM < v3 configmaps are used to store binary data.
	// If you happen to have a lot of HELM releases in the cluster it will make
	// the memory consumption of nginx-ingress-controller explode.
	// In order to avoid that we filter out labels OWNER=TILLER.
	labelsTweakListOptionsFunc := func(options *metav1.ListOptions) {
		if len(options.LabelSelector) > 0 {
			options.LabelSelector += ",OWNER!=TILLER"
		} else {
			options.LabelSelector = "OWNER!=TILLER"
		}
	}

	// As of HELM >= v3 helm releases are stored using Secrets instead of ConfigMaps.
	// In order to avoid listing those secrets we discard type "helm.sh/release.v1"
	secretsTweakListOptionsFunc := func(options *metav1.ListOptions) {
		helmAntiSelector := fields.OneTermNotEqualSelector("type", "helm.sh/release.v1")
		baseSelector, err := fields.ParseSelector(options.FieldSelector)

		if err != nil {
			options.FieldSelector = helmAntiSelector.String()
		} else {
			options.FieldSelector = fields.AndSelectors(baseSelector, helmAntiSelector).String()
		}
	}

	// create informers factory, enable and assign required informers
	infFactory := informers.NewSharedInformerFactoryWithOptions(client, resyncPeriod,
		informers.WithNamespace(namespace),
	)

	// create informers factory for configmaps
	infFactoryConfigmaps := informers.NewSharedInformerFactoryWithOptions(client, resyncPeriod,
		informers.WithNamespace(namespace),
		informers.WithTweakListOptions(labelsTweakListOptionsFunc),
	)

	// create informers factory for secrets
	infFactorySecrets := informers.NewSharedInformerFactoryWithOptions(client, resyncPeriod,
		informers.WithNamespace(namespace),
		informers.WithTweakListOptions(secretsTweakListOptionsFunc),
	)

	store.informers.Ingress = infFactory.Networking().V1beta1().Ingresses().Informer()
	store.listers.Ingress.Store = store.informers.Ingress.GetStore()

	store.informers.Endpoint = infFactory.Core().V1().Endpoints().Informer()
	store.listers.Endpoint.Store = store.informers.Endpoint.GetStore()

	store.informers.Secret = infFactorySecrets.Core().V1().Secrets().Informer()
	store.listers.Secret.Store = store.informers.Secret.GetStore()

	store.informers.ConfigMap = infFactoryConfigmaps.Core().V1().ConfigMaps().Informer()
	store.listers.ConfigMap.Store = store.informers.ConfigMap.GetStore()

	store.informers.Service = infFactory.Core().V1().Services().Informer()
	store.listers.Service.Store = store.informers.Service.GetStore()

	ingDeleteHandler := func(obj interface{}) {
		ing, ok := toIngress(obj)
		if !ok {
			// If we reached here it means the ingress was deleted but its final state is unrecorded.
			tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				klog.ErrorS(nil, "Error obtaining object from tombstone", "key", obj)
				return
			}
			ing, ok = tombstone.Obj.(*networkingv1beta1.Ingress)
			if !ok {
				klog.Errorf("Tombstone contained object that is not an Ingress: %#v", obj)
				return
			}
		}

		if !class.IsValid(ing) {
			return
		}

		if isCatchAllIngress(ing.Spec) && disableCatchAll {
			klog.InfoS("Ignoring delete for catch-all because of --disable-catch-all", "ingress", klog.KObj(ing))
			return
		}

		store.listers.IngressWithAnnotation.Delete(ing)

		key := k8s.MetaNamespaceKey(ing)
		store.secretIngressMap.Delete(key)

		updateCh.In() <- Event{
			Type: DeleteEvent,
			Obj:  obj,
		}
	}

	ingEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ing, _ := toIngress(obj)
			if !class.IsValid(ing) {
				ingressClass, _ := parser.GetStringAnnotation(class.IngressKey, ing)
				klog.InfoS("Ignoring ingress", "ingress", klog.KObj(ing), "kubernetes.io/ingress.class", ingressClass, "ingressClassName", pointer.StringPtrDerefOr(ing.Spec.IngressClassName, ""))
				return
			}

			if isCatchAllIngress(ing.Spec) && disableCatchAll {
				klog.InfoS("Ignoring add for catch-all ingress because of --disable-catch-all", "ingress", klog.KObj(ing))
				return
			}

			recorder.Eventf(ing, corev1.EventTypeNormal, "Sync", "Scheduled for sync")

			store.syncIngress(ing)
			store.updateSecretIngressMap(ing)
			store.syncSecrets(ing)

			updateCh.In() <- Event{
				Type: CreateEvent,
				Obj:  obj,
			}
		},
		DeleteFunc: ingDeleteHandler,
		UpdateFunc: func(old, cur interface{}) {
			oldIng, _ := toIngress(old)
			curIng, _ := toIngress(cur)

			validOld := class.IsValid(oldIng)
			validCur := class.IsValid(curIng)
			if !validOld && validCur {
				if isCatchAllIngress(curIng.Spec) && disableCatchAll {
					klog.InfoS("ignoring update for catch-all ingress because of --disable-catch-all", "ingress", klog.KObj(curIng))
					return
				}

				klog.InfoS("creating ingress", "ingress", klog.KObj(curIng), "class", class.IngressKey)
				recorder.Eventf(curIng, corev1.EventTypeNormal, "Sync", "Scheduled for sync")
			} else if validOld && !validCur {
				klog.InfoS("removing ingress", "ingress", klog.KObj(curIng), "class", class.IngressKey)
				ingDeleteHandler(old)
				return
			} else if validCur && !reflect.DeepEqual(old, cur) {
				if isCatchAllIngress(curIng.Spec) && disableCatchAll {
					klog.InfoS("ignoring update for catch-all ingress and delete old one because of --disable-catch-all", "ingress", klog.KObj(curIng))
					ingDeleteHandler(old)
					return
				}

				recorder.Eventf(curIng, corev1.EventTypeNormal, "Sync", "Scheduled for sync")
			} else {
				klog.V(3).InfoS("No changes on ingress. Skipping update", "ingress", klog.KObj(curIng))
				return
			}

			store.syncIngress(curIng)
			store.updateSecretIngressMap(curIng)
			store.syncSecrets(curIng)

			updateCh.In() <- Event{
				Type: UpdateEvent,
				Obj:  cur,
			}
		},
	}

	secrEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sec := obj.(*corev1.Secret)
			key := k8s.MetaNamespaceKey(sec)

			if store.defaultSSLCertificate == key {
				store.syncSecret(store.defaultSSLCertificate)
			}

			// find references in ingresses and update local ssl certs
			if ings := store.secretIngressMap.Reference(key); len(ings) > 0 {
				klog.InfoS("Secret was added and it is used in ingress annotations. Parsing", "secret", key)
				for _, ingKey := range ings {
					ing, err := store.getIngress(ingKey)
					if err != nil {
						klog.Errorf("could not find Ingress %v in local store", ingKey)
						continue
					}
					store.syncIngress(ing)
					store.syncSecrets(ing)
				}
				updateCh.In() <- Event{
					Type: CreateEvent,
					Obj:  obj,
				}
			}
		},
		UpdateFunc: func(old, cur interface{}) {
			if !reflect.DeepEqual(old, cur) {
				sec := cur.(*corev1.Secret)
				key := k8s.MetaNamespaceKey(sec)

				if store.defaultSSLCertificate == key {
					store.syncSecret(store.defaultSSLCertificate)
				}

				// find references in ingresses and update local ssl certs
				if ings := store.secretIngressMap.Reference(key); len(ings) > 0 {
					klog.InfoS("secret was updated and it is used in ingress annotations. Parsing", "secret", key)
					for _, ingKey := range ings {
						ing, err := store.getIngress(ingKey)
						if err != nil {
							klog.ErrorS(err, "could not find Ingress in local store", "ingress", ingKey)
							continue
						}
						store.syncSecrets(ing)
						store.syncIngress(ing)
					}
					updateCh.In() <- Event{
						Type: UpdateEvent,
						Obj:  cur,
					}
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			sec, ok := obj.(*corev1.Secret)
			if !ok {
				// If we reached here it means the secret was deleted but its final state is unrecorded.
				tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
				if !ok {
					return
				}

				sec, ok = tombstone.Obj.(*corev1.Secret)
				if !ok {
					return
				}
			}

			store.sslStore.Delete(k8s.MetaNamespaceKey(sec))

			key := k8s.MetaNamespaceKey(sec)

			// find references in ingresses
			if ings := store.secretIngressMap.Reference(key); len(ings) > 0 {
				klog.InfoS("secret was deleted and it is used in ingress annotations. Parsing", "secret", key)
				for _, ingKey := range ings {
					ing, err := store.getIngress(ingKey)
					if err != nil {
						klog.Errorf("could not find Ingress %v in local store", ingKey)
						continue
					}
					store.syncIngress(ing)
				}

				updateCh.In() <- Event{
					Type: DeleteEvent,
					Obj:  obj,
				}
			}
		},
	}

	epEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			updateCh.In() <- Event{
				Type: CreateEvent,
				Obj:  obj,
			}
		},
		DeleteFunc: func(obj interface{}) {
			updateCh.In() <- Event{
				Type: DeleteEvent,
				Obj:  obj,
			}
		},
		UpdateFunc: func(old, cur interface{}) {
			oep := old.(*corev1.Endpoints)
			cep := cur.(*corev1.Endpoints)
			if !reflect.DeepEqual(cep.Subsets, oep.Subsets) {
				updateCh.In() <- Event{
					Type: UpdateEvent,
					Obj:  cur,
				}
			}
		},
	}

	// TODO: add e2e test to verify that changes to one or more configmap trigger an update
	changeTriggerUpdate := func(name string) bool {
		return name == configmap || name == tcp || name == udp
	}

	handleCfgMapEvent := func(key string, cfgMap *corev1.ConfigMap, eventName string) {
		// updates to configuration configmaps can trigger an update
		triggerUpdate := false
		if changeTriggerUpdate(key) {
			triggerUpdate = true
			recorder.Eventf(cfgMap, corev1.EventTypeNormal, eventName, fmt.Sprintf("ConfigMap %v", key))
			if key == configmap {
				store.setConfig(cfgMap)
			}
		}

		ings := store.listers.IngressWithAnnotation.List()
		for _, ingKey := range ings {
			key := k8s.MetaNamespaceKey(ingKey)
			ing, err := store.getIngress(key)
			if err != nil {
				klog.Errorf("could not find Ingress %v in local store: %v", key, err)
				continue
			}

			if parser.AnnotationsReferencesConfigmap(ing) {
				store.syncIngress(ing)
				continue
			}

			if triggerUpdate {
				store.syncIngress(ing)
			}
		}

		if triggerUpdate {
			updateCh.In() <- Event{
				Type: ConfigurationEvent,
				Obj:  cfgMap,
			}
		}
	}

	cmEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			cfgMap := obj.(*corev1.ConfigMap)
			key := k8s.MetaNamespaceKey(cfgMap)
			handleCfgMapEvent(key, cfgMap, "CREATE")
		},
		UpdateFunc: func(old, cur interface{}) {
			if reflect.DeepEqual(old, cur) {
				return
			}

			cfgMap := cur.(*corev1.ConfigMap)
			key := k8s.MetaNamespaceKey(cfgMap)
			handleCfgMapEvent(key, cfgMap, "UPDATE")
		},
	}

	serviceHandler := cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(old, cur interface{}) {
			oldSvc := old.(*corev1.Service)
			curSvc := cur.(*corev1.Service)

			if reflect.DeepEqual(oldSvc, curSvc) {
				return
			}

			updateCh.In() <- Event{
				Type: UpdateEvent,
				Obj:  cur,
			}
		},
	}

	store.informers.Ingress.AddEventHandler(ingEventHandler)
	store.informers.Endpoint.AddEventHandler(epEventHandler)
	store.informers.Secret.AddEventHandler(secrEventHandler)
	store.informers.ConfigMap.AddEventHandler(cmEventHandler)
	store.informers.Service.AddEventHandler(serviceHandler)

	// do not wait for informers to read the configmap configuration
	ns, name, _ := k8s.ParseNameNS(configmap)
	cm, err := client.CoreV1().ConfigMaps(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Warningf("Unexpected error reading configuration configmap: %v", err)
	}

	store.setConfig(cm)
	return store
}

// isCatchAllIngress returns whether or not an ingress produces a
// catch-all server, and so should be ignored when --disable-catch-all is set
func isCatchAllIngress(spec networkingv1beta1.IngressSpec) bool {
	return spec.Backend != nil && len(spec.Rules) == 0
}

// syncIngress parses ingress annotations converting the value of the
// annotation to a go struct
func (s *k8sStore) syncIngress(ing *networkingv1beta1.Ingress) {
	key := k8s.MetaNamespaceKey(ing)
	klog.V(3).Infof("updating annotations information for ingress %v", key)

	copyIng := &networkingv1beta1.Ingress{}
	ing.ObjectMeta.DeepCopyInto(&copyIng.ObjectMeta)
	ing.Spec.DeepCopyInto(&copyIng.Spec)
	ing.Status.DeepCopyInto(&copyIng.Status)

	for ri, rule := range copyIng.Spec.Rules {
		if rule.HTTP == nil {
			continue
		}

		for pi, path := range rule.HTTP.Paths {
			if path.Path == "" {
				copyIng.Spec.Rules[ri].HTTP.Paths[pi].Path = "/"
			}
		}
	}

	k8s.SetDefaultNGINXPathType(copyIng)

	err := s.listers.IngressWithAnnotation.Update(&ingress.Ingress{
		Ingress:           *copyIng,
		ParsedAnnotations: s.annotations.Extract(ing),
	})
	if err != nil {
		klog.Error(err)
	}
}

// updateSecretIngressMap takes an Ingress and updates all Secret objects it
// references in secretIngressMap.
func (s *k8sStore) updateSecretIngressMap(ing *networkingv1beta1.Ingress) {
	key := k8s.MetaNamespaceKey(ing)
	klog.V(3).Infof("updating references to secrets for ingress %v", key)

	// delete all existing references first
	s.secretIngressMap.Delete(key)

	var refSecrets []string

	for _, tls := range ing.Spec.TLS {
		secrName := tls.SecretName
		if secrName != "" {
			secrKey := fmt.Sprintf("%v/%v", ing.Namespace, secrName)
			refSecrets = append(refSecrets, secrKey)
		}
	}

	// We can not rely on cached ingress annotations because these are
	// discarded when the referenced secret does not exist in the local
	// store. As a result, adding a secret *after* the ingress(es) which
	// references it would not trigger a resync of that secret.
	secretAnnotations := []string{
		"auth-secret",
		"auth-tls-secret",
		"proxy-ssl-secret",
		"secure-verify-ca-secret",
	}
	for _, ann := range secretAnnotations {
		secrKey, err := objectRefAnnotationNsKey(ann, ing)
		if err != nil && !errors.IsMissingAnnotations(err) {
			klog.Errorf("error reading secret reference in annotation %q: %s", ann, err)
			continue
		}
		if secrKey != "" {
			refSecrets = append(refSecrets, secrKey)
		}
	}

	// populate map with all secret references
	s.secretIngressMap.Insert(key, refSecrets...)
}

// objectRefAnnotationNsKey returns an object reference formatted as a
// 'namespace/name' key from the given annotation name.
func objectRefAnnotationNsKey(ann string, ing *networkingv1beta1.Ingress) (string, error) {
	annValue, err := parser.GetStringAnnotation(ann, ing)
	if err != nil {
		return "", err
	}

	secrNs, secrName, err := cache.SplitMetaNamespaceKey(annValue)
	if secrName == "" {
		return "", err
	}

	if secrNs == "" {
		return fmt.Sprintf("%v/%v", ing.Namespace, secrName), nil
	}
	return annValue, nil
}

// syncSecrets synchronizes data from all Secrets referenced by the given
// Ingress with the local store and file system.
func (s *k8sStore) syncSecrets(ing *networkingv1beta1.Ingress) {
	key := k8s.MetaNamespaceKey(ing)
	for _, secrKey := range s.secretIngressMap.ReferencedBy(key) {
		s.syncSecret(secrKey)
	}
}

// GetSecret returns the Secret matching key.
func (s *k8sStore) GetSecret(key string) (*corev1.Secret, error) {
	return s.listers.Secret.ByKey(key)
}

// ListLocalSSLCerts returns the list of local SSLCerts
func (s *k8sStore) ListLocalSSLCerts() []*ingress.SSLCert {
	var certs []*ingress.SSLCert
	for _, item := range s.sslStore.List() {
		if s, ok := item.(*ingress.SSLCert); ok {
			certs = append(certs, s)
		}
	}

	return certs
}

// GetService returns the Service matching key.
func (s *k8sStore) GetService(key string) (*corev1.Service, error) {
	return s.listers.Service.ByKey(key)
}

// getIngress returns the Ingress matching key.
func (s *k8sStore) getIngress(key string) (*networkingv1beta1.Ingress, error) {
	ing, err := s.listers.IngressWithAnnotation.ByKey(key)
	if err != nil {
		return nil, err
	}

	return &ing.Ingress, nil
}

func sortIngressSlice(ingresses []*ingress.Ingress) {
	// sort Ingresses using the CreationTimestamp field
	sort.SliceStable(ingresses, func(i, j int) bool {
		ir := ingresses[i].CreationTimestamp
		jr := ingresses[j].CreationTimestamp
		if ir.Equal(&jr) {
			in := fmt.Sprintf("%v/%v", ingresses[i].Namespace, ingresses[i].Name)
			jn := fmt.Sprintf("%v/%v", ingresses[j].Namespace, ingresses[j].Name)
			klog.V(3).Infof("Ingress %v and %v have identical CreationTimestamp", in, jn)
			return in > jn
		}
		return ir.Before(&jr)
	})
}

// ListIngresses returns the list of Ingresses
func (s *k8sStore) ListIngresses() []*ingress.Ingress {
	// filter ingress rules
	ingresses := make([]*ingress.Ingress, 0)
	for _, item := range s.listers.IngressWithAnnotation.List() {
		ing := item.(*ingress.Ingress)
		ingresses = append(ingresses, ing)
	}

	sortIngressSlice(ingresses)

	return ingresses
}

// GetLocalSSLCert returns the local copy of a SSLCert
func (s *k8sStore) GetLocalSSLCert(key string) (*ingress.SSLCert, error) {
	return s.sslStore.ByKey(key)
}

// GetConfigMap returns the ConfigMap matching key.
func (s *k8sStore) GetConfigMap(key string) (*corev1.ConfigMap, error) {
	return s.listers.ConfigMap.ByKey(key)
}

// GetServiceEndpoints returns the Endpoints of a Service matching key.
func (s *k8sStore) GetServiceEndpoints(key string) (*corev1.Endpoints, error) {
	return s.listers.Endpoint.ByKey(key)
}

// GetAuthCertificate is used by the auth-tls annotations to get a cert from a secret
func (s *k8sStore) GetAuthCertificate(name string) (*resolver.AuthSSLCert, error) {
	if _, err := s.GetLocalSSLCert(name); err != nil {
		s.syncSecret(name)
	}

	cert, err := s.GetLocalSSLCert(name)
	if err != nil {
		return nil, err
	}

	return &resolver.AuthSSLCert{
		Secret:      name,
		CAFileName:  cert.CAFileName,
		CASHA:       cert.CASHA,
		CRLFileName: cert.CRLFileName,
		CRLSHA:      cert.CRLSHA,
		PemFileName: cert.PemFileName,
	}, nil
}

func (s *k8sStore) writeSSLSessionTicketKey(cmap *corev1.ConfigMap, fileName string) {
	ticketString := ngx_template.ReadConfig(cmap.Data).SSLSessionTicketKey
	s.backendConfig.SSLSessionTicketKey = ""

	if ticketString != "" {
		ticketBytes := base64.StdEncoding.WithPadding(base64.StdPadding).DecodedLen(len(ticketString))

		// 81 used instead of 80 because of padding
		if !(ticketBytes == 48 || ticketBytes == 81) {
			klog.Warningf("ssl-session-ticket-key must contain either 48 or 80 bytes")
		}

		decodedTicket, err := base64.StdEncoding.DecodeString(ticketString)
		if err != nil {
			klog.Errorf("unexpected error decoding ssl-session-ticket-key: %v", err)
			return
		}

		err = ioutil.WriteFile(fileName, decodedTicket, file.ReadWriteByUser)
		if err != nil {
			klog.Errorf("unexpected error writing ssl-session-ticket-key to %s: %v", fileName, err)
			return
		}

		s.backendConfig.SSLSessionTicketKey = ticketString
	}
}

// GetDefaultBackend returns the default backend
func (s *k8sStore) GetDefaultBackend() defaults.Backend {
	return s.GetBackendConfiguration().Backend
}

func (s *k8sStore) GetBackendConfiguration() ngx_config.Configuration {
	s.backendConfigMu.RLock()
	defer s.backendConfigMu.RUnlock()

	return s.backendConfig
}

func (s *k8sStore) setConfig(cmap *corev1.ConfigMap) {
	s.backendConfigMu.Lock()
	defer s.backendConfigMu.Unlock()

	if cmap == nil {
		return
	}

	s.backendConfig = ngx_template.ReadConfig(cmap.Data)
	if s.backendConfig.UseGeoIP2 && !nginx.GeoLite2DBExists() {
		klog.Warning("The GeoIP2 feature is enabled but the databases are missing. Disabling")
		s.backendConfig.UseGeoIP2 = false
	}

	s.writeSSLSessionTicketKey(cmap, "/etc/nginx/tickets.key")
}

// Run initiates the synchronization of the informers and the initial
// synchronization of the secrets.
func (s *k8sStore) Run(stopCh chan struct{}) {
	// start informers
	s.informers.Run(stopCh)
}

var runtimeScheme = k8sruntime.NewScheme()

func init() {
	utilruntime.Must(networkingv1beta1.AddToScheme(runtimeScheme))
}

func toIngress(obj interface{}) (*networkingv1beta1.Ingress, bool) {
	if ing, ok := obj.(*networkingv1beta1.Ingress); ok {
		k8s.SetDefaultNGINXPathType(ing)
		return ing, true
	}

	return nil, false
}
