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

// AutomationCredentialLister helps list AutomationCredentials.
type AutomationCredentialLister interface {
	// List lists all AutomationCredentials in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationCredential, err error)
	// AutomationCredentials returns an object that can list and get AutomationCredentials.
	AutomationCredentials(namespace string) AutomationCredentialNamespaceLister
	AutomationCredentialListerExpansion
}

// automationCredentialLister implements the AutomationCredentialLister interface.
type automationCredentialLister struct {
	indexer cache.Indexer
}

// NewAutomationCredentialLister returns a new AutomationCredentialLister.
func NewAutomationCredentialLister(indexer cache.Indexer) AutomationCredentialLister {
	return &automationCredentialLister{indexer: indexer}
}

// List lists all AutomationCredentials in the indexer.
func (s *automationCredentialLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationCredential, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationCredential))
	})
	return ret, err
}

// AutomationCredentials returns an object that can list and get AutomationCredentials.
func (s *automationCredentialLister) AutomationCredentials(namespace string) AutomationCredentialNamespaceLister {
	return automationCredentialNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutomationCredentialNamespaceLister helps list and get AutomationCredentials.
type AutomationCredentialNamespaceLister interface {
	// List lists all AutomationCredentials in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutomationCredential, err error)
	// Get retrieves the AutomationCredential from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutomationCredential, error)
	AutomationCredentialNamespaceListerExpansion
}

// automationCredentialNamespaceLister implements the AutomationCredentialNamespaceLister
// interface.
type automationCredentialNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutomationCredentials in the indexer for a given namespace.
func (s automationCredentialNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutomationCredential, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutomationCredential))
	})
	return ret, err
}

// Get retrieves the AutomationCredential from the indexer for a given namespace and name.
func (s automationCredentialNamespaceLister) Get(name string) (*v1alpha1.AutomationCredential, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("automationcredential"), name)
	}
	return obj.(*v1alpha1.AutomationCredential), nil
}