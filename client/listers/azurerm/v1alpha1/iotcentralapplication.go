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

// IotcentralApplicationLister helps list IotcentralApplications.
type IotcentralApplicationLister interface {
	// List lists all IotcentralApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotcentralApplication, err error)
	// IotcentralApplications returns an object that can list and get IotcentralApplications.
	IotcentralApplications(namespace string) IotcentralApplicationNamespaceLister
	IotcentralApplicationListerExpansion
}

// iotcentralApplicationLister implements the IotcentralApplicationLister interface.
type iotcentralApplicationLister struct {
	indexer cache.Indexer
}

// NewIotcentralApplicationLister returns a new IotcentralApplicationLister.
func NewIotcentralApplicationLister(indexer cache.Indexer) IotcentralApplicationLister {
	return &iotcentralApplicationLister{indexer: indexer}
}

// List lists all IotcentralApplications in the indexer.
func (s *iotcentralApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.IotcentralApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotcentralApplication))
	})
	return ret, err
}

// IotcentralApplications returns an object that can list and get IotcentralApplications.
func (s *iotcentralApplicationLister) IotcentralApplications(namespace string) IotcentralApplicationNamespaceLister {
	return iotcentralApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotcentralApplicationNamespaceLister helps list and get IotcentralApplications.
type IotcentralApplicationNamespaceLister interface {
	// List lists all IotcentralApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotcentralApplication, err error)
	// Get retrieves the IotcentralApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotcentralApplication, error)
	IotcentralApplicationNamespaceListerExpansion
}

// iotcentralApplicationNamespaceLister implements the IotcentralApplicationNamespaceLister
// interface.
type iotcentralApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotcentralApplications in the indexer for a given namespace.
func (s iotcentralApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotcentralApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotcentralApplication))
	})
	return ret, err
}

// Get retrieves the IotcentralApplication from the indexer for a given namespace and name.
func (s iotcentralApplicationNamespaceLister) Get(name string) (*v1alpha1.IotcentralApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iotcentralapplication"), name)
	}
	return obj.(*v1alpha1.IotcentralApplication), nil
}