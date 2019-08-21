/*
Copyright 2020 The OpenEBS Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	lvmv1alpha1 "github.com/pawanpraka1/dynamic-lvm/pkg/apis/openebs.io/lvm/v1alpha1"
	internalclientset "github.com/pawanpraka1/dynamic-lvm/pkg/generated/clientset/internalclientset"
	internalinterfaces "github.com/pawanpraka1/dynamic-lvm/pkg/generated/informer/externalversions/internalinterfaces"
	v1alpha1 "github.com/pawanpraka1/dynamic-lvm/pkg/generated/lister/lvm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LVMVolumeInformer provides access to a shared informer and lister for
// LVMVolumes.
type LVMVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LVMVolumeLister
}

type lVMVolumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLVMVolumeInformer constructs a new informer for LVMVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLVMVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLVMVolumeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLVMVolumeInformer constructs a new informer for LVMVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLVMVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalV1alpha1().LVMVolumes(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalV1alpha1().LVMVolumes(namespace).Watch(options)
			},
		},
		&lvmv1alpha1.LVMVolume{},
		resyncPeriod,
		indexers,
	)
}

func (f *lVMVolumeInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLVMVolumeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lVMVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lvmv1alpha1.LVMVolume{}, f.defaultInformer)
}

func (f *lVMVolumeInformer) Lister() v1alpha1.LVMVolumeLister {
	return v1alpha1.NewLVMVolumeLister(f.Informer().GetIndexer())
}
