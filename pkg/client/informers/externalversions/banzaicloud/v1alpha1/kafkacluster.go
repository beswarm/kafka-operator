// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	banzaicloudv1alpha1 "github.com/banzaicloud/kafka-operator/pkg/apis/banzaicloud/v1alpha1"
	versioned "github.com/banzaicloud/kafka-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/banzaicloud/kafka-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/banzaicloud/kafka-operator/pkg/client/listers/banzaicloud/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KafkaClusterInformer provides access to a shared informer and lister for
// KafkaClusters.
type KafkaClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KafkaClusterLister
}

type kafkaClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKafkaClusterInformer constructs a new informer for KafkaCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKafkaClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKafkaClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKafkaClusterInformer constructs a new informer for KafkaCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKafkaClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BanzaicloudV1alpha1().KafkaClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BanzaicloudV1alpha1().KafkaClusters(namespace).Watch(options)
			},
		},
		&banzaicloudv1alpha1.KafkaCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *kafkaClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKafkaClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kafkaClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&banzaicloudv1alpha1.KafkaCluster{}, f.defaultInformer)
}

func (f *kafkaClusterInformer) Lister() v1alpha1.KafkaClusterLister {
	return v1alpha1.NewKafkaClusterLister(f.Informer().GetIndexer())
}
