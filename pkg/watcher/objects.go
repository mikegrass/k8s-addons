package watcher

import (
	aci "github.com/appscode/k8s-addons/api"
	"github.com/appscode/k8s-addons/pkg/events"
	"github.com/appscode/log"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/apps"
	ext "k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/util/wait"
)

func (k *Watcher) Namespace() {
	log.Debugln("watching", events.Namespace.String())
	_, controller := k.Cache(events.Namespace, &kapi.Namespace{}, nil)
	go controller.Run(wait.NeverStop)
}

func (k *Watcher) Pod() {
	log.Debugln("watching", events.Pod.String())
	indexer, controller := k.CacheIndexer(events.Pod, &kapi.Pod{}, nil, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.PodStore = cache.StoreToPodLister{indexer}
}

func (k *Watcher) Service() {
	log.Debugln("watching", events.Service.String())
	indexer, controller := k.CacheIndexer(events.Service, &kapi.Service{}, nil, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.ServiceStore = cache.StoreToServiceLister{indexer}
}

func (k *Watcher) RC() {
	log.Debugln("watching", events.RC.String())
	indexer, controller := k.CacheIndexer(events.RC, &kapi.ReplicationController{}, nil, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.RcStore = cache.StoreToReplicationControllerLister{indexer}
}

func (k *Watcher) ReplicaSet() {
	log.Debugln("watching", events.ReplicaSet.String())
	lw := &cache.ListWatch{
		ListFunc:  replicaSetListFunc(k.Client),
		WatchFunc: replicaSetWatchFunc(k.Client),
	}
	indexer, controller := k.CacheIndexer(events.ReplicaSet, &ext.ReplicaSet{}, lw, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.ReplicaSetStore = cache.StoreToReplicaSetLister{indexer}
}

func (k *Watcher) PetSet() {
	log.Debugln("watching", events.StatefulSet.String())
	lw := &cache.ListWatch{
		ListFunc:  petSetListFunc(k.Client),
		WatchFunc: petSetWatchFunc(k.Client),
	}
	indexer, controller := k.CacheIndexer(events.StatefulSet, &apps.StatefulSet{}, lw, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.StatefulSetStore = cache.StoreToStatefulSetLister{indexer}
}

func (k *Watcher) DaemonSet() {
	log.Debugln("watching", events.DaemonSet.String())
	lw := &cache.ListWatch{
		ListFunc:  daemonSetListFunc(k.Client),
		WatchFunc: daemonSetWatchFunc(k.Client),
	}
	indexer, controller := k.CacheIndexer(events.DaemonSet, &ext.DaemonSet{}, lw, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.DaemonSetStore = cache.StoreToDaemonSetLister{indexer}
}

func (k *Watcher) Endpoint() {
	log.Debugln("watching", events.Endpoint.String())
	store, controller := k.CacheStore(events.Endpoint, &kapi.Endpoints{}, nil)
	go controller.Run(wait.NeverStop)
	k.Storage.EndpointStore = cache.StoreToEndpointsLister{store}
}

func (k *Watcher) Node() {
	log.Debugln("watching", events.Node.String())
	_, controller := k.CacheStore(events.Node, &kapi.Node{}, nil)
	go controller.Run(wait.NeverStop)
}

func (k *Watcher) Ingress() {
	log.Debugln("watching", events.Ingress.String())
	lw := &cache.ListWatch{
		ListFunc:  ingressListFunc(k.Client),
		WatchFunc: ingressWatchFunc(k.Client),
	}
	_, controller := k.Cache(events.Ingress, &ext.Ingress{}, lw)
	go controller.Run(wait.NeverStop)
}

func (k *Watcher) ExtendedIngress() {
	log.Debugln("watching", events.ExtendedIngress.String())
	lw := &cache.ListWatch{
		ListFunc:  extIngressListFunc(k.AppsCodeExtensionClient),
		WatchFunc: extIngressWatchFunc(k.AppsCodeExtensionClient),
	}
	_, controller := k.Cache(events.ExtendedIngress, &aci.Ingress{}, lw)
	go controller.Run(wait.NeverStop)
}

func (k *Watcher) Alert() {
	log.Debugln("watching", events.Alert.String())
	lw := &cache.ListWatch{
		ListFunc:  alertListFunc(k.AppsCodeExtensionClient),
		WatchFunc: alertWatchFunc(k.AppsCodeExtensionClient),
	}
	_, controller := k.Cache(events.Alert, &aci.Alert{}, lw)
	go controller.Run(wait.NeverStop)
}

func (k *Watcher) Certificate() {
	log.Debugln("watching", events.Certificate.String())
	lw := &cache.ListWatch{
		ListFunc:  certificateListFunc(k.AppsCodeExtensionClient),
		WatchFunc: certificateWatchFunc(k.AppsCodeExtensionClient),
	}
	_, controller := k.Cache(events.Certificate, &aci.Certificate{}, lw)
	go controller.Run(wait.NeverStop)
}
