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

// AppServiceEnvironmentLister helps list AppServiceEnvironments.
type AppServiceEnvironmentLister interface {
	// List lists all AppServiceEnvironments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppServiceEnvironment, err error)
	// AppServiceEnvironments returns an object that can list and get AppServiceEnvironments.
	AppServiceEnvironments(namespace string) AppServiceEnvironmentNamespaceLister
	AppServiceEnvironmentListerExpansion
}

// appServiceEnvironmentLister implements the AppServiceEnvironmentLister interface.
type appServiceEnvironmentLister struct {
	indexer cache.Indexer
}

// NewAppServiceEnvironmentLister returns a new AppServiceEnvironmentLister.
func NewAppServiceEnvironmentLister(indexer cache.Indexer) AppServiceEnvironmentLister {
	return &appServiceEnvironmentLister{indexer: indexer}
}

// List lists all AppServiceEnvironments in the indexer.
func (s *appServiceEnvironmentLister) List(selector labels.Selector) (ret []*v1alpha1.AppServiceEnvironment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppServiceEnvironment))
	})
	return ret, err
}

// AppServiceEnvironments returns an object that can list and get AppServiceEnvironments.
func (s *appServiceEnvironmentLister) AppServiceEnvironments(namespace string) AppServiceEnvironmentNamespaceLister {
	return appServiceEnvironmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppServiceEnvironmentNamespaceLister helps list and get AppServiceEnvironments.
type AppServiceEnvironmentNamespaceLister interface {
	// List lists all AppServiceEnvironments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppServiceEnvironment, err error)
	// Get retrieves the AppServiceEnvironment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppServiceEnvironment, error)
	AppServiceEnvironmentNamespaceListerExpansion
}

// appServiceEnvironmentNamespaceLister implements the AppServiceEnvironmentNamespaceLister
// interface.
type appServiceEnvironmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppServiceEnvironments in the indexer for a given namespace.
func (s appServiceEnvironmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppServiceEnvironment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppServiceEnvironment))
	})
	return ret, err
}

// Get retrieves the AppServiceEnvironment from the indexer for a given namespace and name.
func (s appServiceEnvironmentNamespaceLister) Get(name string) (*v1alpha1.AppServiceEnvironment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appserviceenvironment"), name)
	}
	return obj.(*v1alpha1.AppServiceEnvironment), nil
}