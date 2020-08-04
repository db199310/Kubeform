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

// LbRuleLister helps list LbRules.
type LbRuleLister interface {
	// List lists all LbRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbRule, err error)
	// LbRules returns an object that can list and get LbRules.
	LbRules(namespace string) LbRuleNamespaceLister
	LbRuleListerExpansion
}

// lbRuleLister implements the LbRuleLister interface.
type lbRuleLister struct {
	indexer cache.Indexer
}

// NewLbRuleLister returns a new LbRuleLister.
func NewLbRuleLister(indexer cache.Indexer) LbRuleLister {
	return &lbRuleLister{indexer: indexer}
}

// List lists all LbRules in the indexer.
func (s *lbRuleLister) List(selector labels.Selector) (ret []*v1alpha1.LbRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbRule))
	})
	return ret, err
}

// LbRules returns an object that can list and get LbRules.
func (s *lbRuleLister) LbRules(namespace string) LbRuleNamespaceLister {
	return lbRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbRuleNamespaceLister helps list and get LbRules.
type LbRuleNamespaceLister interface {
	// List lists all LbRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbRule, err error)
	// Get retrieves the LbRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbRule, error)
	LbRuleNamespaceListerExpansion
}

// lbRuleNamespaceLister implements the LbRuleNamespaceLister
// interface.
type lbRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbRules in the indexer for a given namespace.
func (s lbRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbRule))
	})
	return ret, err
}

// Get retrieves the LbRule from the indexer for a given namespace and name.
func (s lbRuleNamespaceLister) Get(name string) (*v1alpha1.LbRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbrule"), name)
	}
	return obj.(*v1alpha1.LbRule), nil
}