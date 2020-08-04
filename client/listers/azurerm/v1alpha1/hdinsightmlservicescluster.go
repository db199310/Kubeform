/*
Copyright The Kubeform Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// HdinsightMlServicesClusterLister helps list HdinsightMlServicesClusters.
type HdinsightMlServicesClusterLister interface {
	// List lists all HdinsightMlServicesClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightMlServicesCluster, err error)
	// HdinsightMlServicesClusters returns an object that can list and get HdinsightMlServicesClusters.
	HdinsightMlServicesClusters(namespace string) HdinsightMlServicesClusterNamespaceLister
	HdinsightMlServicesClusterListerExpansion
}

// hdinsightMlServicesClusterLister implements the HdinsightMlServicesClusterLister interface.
type hdinsightMlServicesClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightMlServicesClusterLister returns a new HdinsightMlServicesClusterLister.
func NewHdinsightMlServicesClusterLister(indexer cache.Indexer) HdinsightMlServicesClusterLister {
	return &hdinsightMlServicesClusterLister{indexer: indexer}
}

// List lists all HdinsightMlServicesClusters in the indexer.
func (s *hdinsightMlServicesClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightMlServicesCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightMlServicesCluster))
	})
	return ret, err
}

// HdinsightMlServicesClusters returns an object that can list and get HdinsightMlServicesClusters.
func (s *hdinsightMlServicesClusterLister) HdinsightMlServicesClusters(namespace string) HdinsightMlServicesClusterNamespaceLister {
	return hdinsightMlServicesClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightMlServicesClusterNamespaceLister helps list and get HdinsightMlServicesClusters.
type HdinsightMlServicesClusterNamespaceLister interface {
	// List lists all HdinsightMlServicesClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightMlServicesCluster, err error)
	// Get retrieves the HdinsightMlServicesCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightMlServicesCluster, error)
	HdinsightMlServicesClusterNamespaceListerExpansion
}

// hdinsightMlServicesClusterNamespaceLister implements the HdinsightMlServicesClusterNamespaceLister
// interface.
type hdinsightMlServicesClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightMlServicesClusters in the indexer for a given namespace.
func (s hdinsightMlServicesClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightMlServicesCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightMlServicesCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightMlServicesCluster from the indexer for a given namespace and name.
func (s hdinsightMlServicesClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightMlServicesCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsightmlservicescluster"), name)
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), nil
}