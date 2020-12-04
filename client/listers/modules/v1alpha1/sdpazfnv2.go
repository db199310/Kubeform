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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// SDPAzFnv2Lister helps list SDPAzFnv2s.
type SDPAzFnv2Lister interface {
	// List lists all SDPAzFnv2s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzFnv2, err error)
	// SDPAzFnv2s returns an object that can list and get SDPAzFnv2s.
	SDPAzFnv2s(namespace string) SDPAzFnv2NamespaceLister
	SDPAzFnv2ListerExpansion
}

// sDPAzFnv2Lister implements the SDPAzFnv2Lister interface.
type sDPAzFnv2Lister struct {
	indexer cache.Indexer
}

// NewSDPAzFnv2Lister returns a new SDPAzFnv2Lister.
func NewSDPAzFnv2Lister(indexer cache.Indexer) SDPAzFnv2Lister {
	return &sDPAzFnv2Lister{indexer: indexer}
}

// List lists all SDPAzFnv2s in the indexer.
func (s *sDPAzFnv2Lister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzFnv2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzFnv2))
	})
	return ret, err
}

// SDPAzFnv2s returns an object that can list and get SDPAzFnv2s.
func (s *sDPAzFnv2Lister) SDPAzFnv2s(namespace string) SDPAzFnv2NamespaceLister {
	return sDPAzFnv2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzFnv2NamespaceLister helps list and get SDPAzFnv2s.
type SDPAzFnv2NamespaceLister interface {
	// List lists all SDPAzFnv2s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzFnv2, err error)
	// Get retrieves the SDPAzFnv2 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDPAzFnv2, error)
	SDPAzFnv2NamespaceListerExpansion
}

// sDPAzFnv2NamespaceLister implements the SDPAzFnv2NamespaceLister
// interface.
type sDPAzFnv2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzFnv2s in the indexer for a given namespace.
func (s sDPAzFnv2NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzFnv2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzFnv2))
	})
	return ret, err
}

// Get retrieves the SDPAzFnv2 from the indexer for a given namespace and name.
func (s sDPAzFnv2NamespaceLister) Get(name string) (*v1alpha1.SDPAzFnv2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdpazfnv2"), name)
	}
	return obj.(*v1alpha1.SDPAzFnv2), nil
}