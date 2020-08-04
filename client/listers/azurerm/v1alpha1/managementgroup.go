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

// ManagementGroupLister helps list ManagementGroups.
type ManagementGroupLister interface {
	// List lists all ManagementGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ManagementGroup, err error)
	// ManagementGroups returns an object that can list and get ManagementGroups.
	ManagementGroups(namespace string) ManagementGroupNamespaceLister
	ManagementGroupListerExpansion
}

// managementGroupLister implements the ManagementGroupLister interface.
type managementGroupLister struct {
	indexer cache.Indexer
}

// NewManagementGroupLister returns a new ManagementGroupLister.
func NewManagementGroupLister(indexer cache.Indexer) ManagementGroupLister {
	return &managementGroupLister{indexer: indexer}
}

// List lists all ManagementGroups in the indexer.
func (s *managementGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ManagementGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagementGroup))
	})
	return ret, err
}

// ManagementGroups returns an object that can list and get ManagementGroups.
func (s *managementGroupLister) ManagementGroups(namespace string) ManagementGroupNamespaceLister {
	return managementGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagementGroupNamespaceLister helps list and get ManagementGroups.
type ManagementGroupNamespaceLister interface {
	// List lists all ManagementGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ManagementGroup, err error)
	// Get retrieves the ManagementGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ManagementGroup, error)
	ManagementGroupNamespaceListerExpansion
}

// managementGroupNamespaceLister implements the ManagementGroupNamespaceLister
// interface.
type managementGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagementGroups in the indexer for a given namespace.
func (s managementGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagementGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagementGroup))
	})
	return ret, err
}

// Get retrieves the ManagementGroup from the indexer for a given namespace and name.
func (s managementGroupNamespaceLister) Get(name string) (*v1alpha1.ManagementGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managementgroup"), name)
	}
	return obj.(*v1alpha1.ManagementGroup), nil
}