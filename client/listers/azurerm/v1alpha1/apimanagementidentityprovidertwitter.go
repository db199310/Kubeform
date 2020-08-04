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

// ApiManagementIdentityProviderTwitterLister helps list ApiManagementIdentityProviderTwitters.
type ApiManagementIdentityProviderTwitterLister interface {
	// List lists all ApiManagementIdentityProviderTwitters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderTwitter, err error)
	// ApiManagementIdentityProviderTwitters returns an object that can list and get ApiManagementIdentityProviderTwitters.
	ApiManagementIdentityProviderTwitters(namespace string) ApiManagementIdentityProviderTwitterNamespaceLister
	ApiManagementIdentityProviderTwitterListerExpansion
}

// apiManagementIdentityProviderTwitterLister implements the ApiManagementIdentityProviderTwitterLister interface.
type apiManagementIdentityProviderTwitterLister struct {
	indexer cache.Indexer
}

// NewApiManagementIdentityProviderTwitterLister returns a new ApiManagementIdentityProviderTwitterLister.
func NewApiManagementIdentityProviderTwitterLister(indexer cache.Indexer) ApiManagementIdentityProviderTwitterLister {
	return &apiManagementIdentityProviderTwitterLister{indexer: indexer}
}

// List lists all ApiManagementIdentityProviderTwitters in the indexer.
func (s *apiManagementIdentityProviderTwitterLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderTwitter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementIdentityProviderTwitter))
	})
	return ret, err
}

// ApiManagementIdentityProviderTwitters returns an object that can list and get ApiManagementIdentityProviderTwitters.
func (s *apiManagementIdentityProviderTwitterLister) ApiManagementIdentityProviderTwitters(namespace string) ApiManagementIdentityProviderTwitterNamespaceLister {
	return apiManagementIdentityProviderTwitterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementIdentityProviderTwitterNamespaceLister helps list and get ApiManagementIdentityProviderTwitters.
type ApiManagementIdentityProviderTwitterNamespaceLister interface {
	// List lists all ApiManagementIdentityProviderTwitters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderTwitter, err error)
	// Get retrieves the ApiManagementIdentityProviderTwitter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementIdentityProviderTwitter, error)
	ApiManagementIdentityProviderTwitterNamespaceListerExpansion
}

// apiManagementIdentityProviderTwitterNamespaceLister implements the ApiManagementIdentityProviderTwitterNamespaceLister
// interface.
type apiManagementIdentityProviderTwitterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementIdentityProviderTwitters in the indexer for a given namespace.
func (s apiManagementIdentityProviderTwitterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderTwitter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementIdentityProviderTwitter))
	})
	return ret, err
}

// Get retrieves the ApiManagementIdentityProviderTwitter from the indexer for a given namespace and name.
func (s apiManagementIdentityProviderTwitterNamespaceLister) Get(name string) (*v1alpha1.ApiManagementIdentityProviderTwitter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementidentityprovidertwitter"), name)
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderTwitter), nil
}