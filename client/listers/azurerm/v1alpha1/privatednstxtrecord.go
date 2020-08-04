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

// PrivateDNSTxtRecordLister helps list PrivateDNSTxtRecords.
type PrivateDNSTxtRecordLister interface {
	// List lists all PrivateDNSTxtRecords in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSTxtRecord, err error)
	// PrivateDNSTxtRecords returns an object that can list and get PrivateDNSTxtRecords.
	PrivateDNSTxtRecords(namespace string) PrivateDNSTxtRecordNamespaceLister
	PrivateDNSTxtRecordListerExpansion
}

// privateDNSTxtRecordLister implements the PrivateDNSTxtRecordLister interface.
type privateDNSTxtRecordLister struct {
	indexer cache.Indexer
}

// NewPrivateDNSTxtRecordLister returns a new PrivateDNSTxtRecordLister.
func NewPrivateDNSTxtRecordLister(indexer cache.Indexer) PrivateDNSTxtRecordLister {
	return &privateDNSTxtRecordLister{indexer: indexer}
}

// List lists all PrivateDNSTxtRecords in the indexer.
func (s *privateDNSTxtRecordLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSTxtRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSTxtRecord))
	})
	return ret, err
}

// PrivateDNSTxtRecords returns an object that can list and get PrivateDNSTxtRecords.
func (s *privateDNSTxtRecordLister) PrivateDNSTxtRecords(namespace string) PrivateDNSTxtRecordNamespaceLister {
	return privateDNSTxtRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateDNSTxtRecordNamespaceLister helps list and get PrivateDNSTxtRecords.
type PrivateDNSTxtRecordNamespaceLister interface {
	// List lists all PrivateDNSTxtRecords in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSTxtRecord, err error)
	// Get retrieves the PrivateDNSTxtRecord from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PrivateDNSTxtRecord, error)
	PrivateDNSTxtRecordNamespaceListerExpansion
}

// privateDNSTxtRecordNamespaceLister implements the PrivateDNSTxtRecordNamespaceLister
// interface.
type privateDNSTxtRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateDNSTxtRecords in the indexer for a given namespace.
func (s privateDNSTxtRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSTxtRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSTxtRecord))
	})
	return ret, err
}

// Get retrieves the PrivateDNSTxtRecord from the indexer for a given namespace and name.
func (s privateDNSTxtRecordNamespaceLister) Get(name string) (*v1alpha1.PrivateDNSTxtRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatednstxtrecord"), name)
	}
	return obj.(*v1alpha1.PrivateDNSTxtRecord), nil
}