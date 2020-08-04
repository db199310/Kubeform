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

// LogAnalyticsDatasourceWindowsEventLister helps list LogAnalyticsDatasourceWindowsEvents.
type LogAnalyticsDatasourceWindowsEventLister interface {
	// List lists all LogAnalyticsDatasourceWindowsEvents in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error)
	// LogAnalyticsDatasourceWindowsEvents returns an object that can list and get LogAnalyticsDatasourceWindowsEvents.
	LogAnalyticsDatasourceWindowsEvents(namespace string) LogAnalyticsDatasourceWindowsEventNamespaceLister
	LogAnalyticsDatasourceWindowsEventListerExpansion
}

// logAnalyticsDatasourceWindowsEventLister implements the LogAnalyticsDatasourceWindowsEventLister interface.
type logAnalyticsDatasourceWindowsEventLister struct {
	indexer cache.Indexer
}

// NewLogAnalyticsDatasourceWindowsEventLister returns a new LogAnalyticsDatasourceWindowsEventLister.
func NewLogAnalyticsDatasourceWindowsEventLister(indexer cache.Indexer) LogAnalyticsDatasourceWindowsEventLister {
	return &logAnalyticsDatasourceWindowsEventLister{indexer: indexer}
}

// List lists all LogAnalyticsDatasourceWindowsEvents in the indexer.
func (s *logAnalyticsDatasourceWindowsEventLister) List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogAnalyticsDatasourceWindowsEvent))
	})
	return ret, err
}

// LogAnalyticsDatasourceWindowsEvents returns an object that can list and get LogAnalyticsDatasourceWindowsEvents.
func (s *logAnalyticsDatasourceWindowsEventLister) LogAnalyticsDatasourceWindowsEvents(namespace string) LogAnalyticsDatasourceWindowsEventNamespaceLister {
	return logAnalyticsDatasourceWindowsEventNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogAnalyticsDatasourceWindowsEventNamespaceLister helps list and get LogAnalyticsDatasourceWindowsEvents.
type LogAnalyticsDatasourceWindowsEventNamespaceLister interface {
	// List lists all LogAnalyticsDatasourceWindowsEvents in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error)
	// Get retrieves the LogAnalyticsDatasourceWindowsEvent from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error)
	LogAnalyticsDatasourceWindowsEventNamespaceListerExpansion
}

// logAnalyticsDatasourceWindowsEventNamespaceLister implements the LogAnalyticsDatasourceWindowsEventNamespaceLister
// interface.
type logAnalyticsDatasourceWindowsEventNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogAnalyticsDatasourceWindowsEvents in the indexer for a given namespace.
func (s logAnalyticsDatasourceWindowsEventNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogAnalyticsDatasourceWindowsEvent))
	})
	return ret, err
}

// Get retrieves the LogAnalyticsDatasourceWindowsEvent from the indexer for a given namespace and name.
func (s logAnalyticsDatasourceWindowsEventNamespaceLister) Get(name string) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loganalyticsdatasourcewindowsevent"), name)
	}
	return obj.(*v1alpha1.LogAnalyticsDatasourceWindowsEvent), nil
}