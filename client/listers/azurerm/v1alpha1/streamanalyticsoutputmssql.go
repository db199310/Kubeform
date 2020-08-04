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

// StreamAnalyticsOutputMssqlLister helps list StreamAnalyticsOutputMssqls.
type StreamAnalyticsOutputMssqlLister interface {
	// List lists all StreamAnalyticsOutputMssqls in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputMssql, err error)
	// StreamAnalyticsOutputMssqls returns an object that can list and get StreamAnalyticsOutputMssqls.
	StreamAnalyticsOutputMssqls(namespace string) StreamAnalyticsOutputMssqlNamespaceLister
	StreamAnalyticsOutputMssqlListerExpansion
}

// streamAnalyticsOutputMssqlLister implements the StreamAnalyticsOutputMssqlLister interface.
type streamAnalyticsOutputMssqlLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsOutputMssqlLister returns a new StreamAnalyticsOutputMssqlLister.
func NewStreamAnalyticsOutputMssqlLister(indexer cache.Indexer) StreamAnalyticsOutputMssqlLister {
	return &streamAnalyticsOutputMssqlLister{indexer: indexer}
}

// List lists all StreamAnalyticsOutputMssqls in the indexer.
func (s *streamAnalyticsOutputMssqlLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputMssql, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputMssql))
	})
	return ret, err
}

// StreamAnalyticsOutputMssqls returns an object that can list and get StreamAnalyticsOutputMssqls.
func (s *streamAnalyticsOutputMssqlLister) StreamAnalyticsOutputMssqls(namespace string) StreamAnalyticsOutputMssqlNamespaceLister {
	return streamAnalyticsOutputMssqlNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsOutputMssqlNamespaceLister helps list and get StreamAnalyticsOutputMssqls.
type StreamAnalyticsOutputMssqlNamespaceLister interface {
	// List lists all StreamAnalyticsOutputMssqls in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputMssql, err error)
	// Get retrieves the StreamAnalyticsOutputMssql from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsOutputMssql, error)
	StreamAnalyticsOutputMssqlNamespaceListerExpansion
}

// streamAnalyticsOutputMssqlNamespaceLister implements the StreamAnalyticsOutputMssqlNamespaceLister
// interface.
type streamAnalyticsOutputMssqlNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsOutputMssqls in the indexer for a given namespace.
func (s streamAnalyticsOutputMssqlNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputMssql, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputMssql))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsOutputMssql from the indexer for a given namespace and name.
func (s streamAnalyticsOutputMssqlNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsOutputMssql, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsoutputmssql"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputMssql), nil
}