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

// CosmosdbCassandraKeyspaceLister helps list CosmosdbCassandraKeyspaces.
type CosmosdbCassandraKeyspaceLister interface {
	// List lists all CosmosdbCassandraKeyspaces in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbCassandraKeyspace, err error)
	// CosmosdbCassandraKeyspaces returns an object that can list and get CosmosdbCassandraKeyspaces.
	CosmosdbCassandraKeyspaces(namespace string) CosmosdbCassandraKeyspaceNamespaceLister
	CosmosdbCassandraKeyspaceListerExpansion
}

// cosmosdbCassandraKeyspaceLister implements the CosmosdbCassandraKeyspaceLister interface.
type cosmosdbCassandraKeyspaceLister struct {
	indexer cache.Indexer
}

// NewCosmosdbCassandraKeyspaceLister returns a new CosmosdbCassandraKeyspaceLister.
func NewCosmosdbCassandraKeyspaceLister(indexer cache.Indexer) CosmosdbCassandraKeyspaceLister {
	return &cosmosdbCassandraKeyspaceLister{indexer: indexer}
}

// List lists all CosmosdbCassandraKeyspaces in the indexer.
func (s *cosmosdbCassandraKeyspaceLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbCassandraKeyspace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbCassandraKeyspace))
	})
	return ret, err
}

// CosmosdbCassandraKeyspaces returns an object that can list and get CosmosdbCassandraKeyspaces.
func (s *cosmosdbCassandraKeyspaceLister) CosmosdbCassandraKeyspaces(namespace string) CosmosdbCassandraKeyspaceNamespaceLister {
	return cosmosdbCassandraKeyspaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CosmosdbCassandraKeyspaceNamespaceLister helps list and get CosmosdbCassandraKeyspaces.
type CosmosdbCassandraKeyspaceNamespaceLister interface {
	// List lists all CosmosdbCassandraKeyspaces in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbCassandraKeyspace, err error)
	// Get retrieves the CosmosdbCassandraKeyspace from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CosmosdbCassandraKeyspace, error)
	CosmosdbCassandraKeyspaceNamespaceListerExpansion
}

// cosmosdbCassandraKeyspaceNamespaceLister implements the CosmosdbCassandraKeyspaceNamespaceLister
// interface.
type cosmosdbCassandraKeyspaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CosmosdbCassandraKeyspaces in the indexer for a given namespace.
func (s cosmosdbCassandraKeyspaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbCassandraKeyspace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbCassandraKeyspace))
	})
	return ret, err
}

// Get retrieves the CosmosdbCassandraKeyspace from the indexer for a given namespace and name.
func (s cosmosdbCassandraKeyspaceNamespaceLister) Get(name string) (*v1alpha1.CosmosdbCassandraKeyspace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cosmosdbcassandrakeyspace"), name)
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), nil
}