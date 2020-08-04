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

// WindowsVirtualMachineScaleSetLister helps list WindowsVirtualMachineScaleSets.
type WindowsVirtualMachineScaleSetLister interface {
	// List lists all WindowsVirtualMachineScaleSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WindowsVirtualMachineScaleSet, err error)
	// WindowsVirtualMachineScaleSets returns an object that can list and get WindowsVirtualMachineScaleSets.
	WindowsVirtualMachineScaleSets(namespace string) WindowsVirtualMachineScaleSetNamespaceLister
	WindowsVirtualMachineScaleSetListerExpansion
}

// windowsVirtualMachineScaleSetLister implements the WindowsVirtualMachineScaleSetLister interface.
type windowsVirtualMachineScaleSetLister struct {
	indexer cache.Indexer
}

// NewWindowsVirtualMachineScaleSetLister returns a new WindowsVirtualMachineScaleSetLister.
func NewWindowsVirtualMachineScaleSetLister(indexer cache.Indexer) WindowsVirtualMachineScaleSetLister {
	return &windowsVirtualMachineScaleSetLister{indexer: indexer}
}

// List lists all WindowsVirtualMachineScaleSets in the indexer.
func (s *windowsVirtualMachineScaleSetLister) List(selector labels.Selector) (ret []*v1alpha1.WindowsVirtualMachineScaleSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WindowsVirtualMachineScaleSet))
	})
	return ret, err
}

// WindowsVirtualMachineScaleSets returns an object that can list and get WindowsVirtualMachineScaleSets.
func (s *windowsVirtualMachineScaleSetLister) WindowsVirtualMachineScaleSets(namespace string) WindowsVirtualMachineScaleSetNamespaceLister {
	return windowsVirtualMachineScaleSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WindowsVirtualMachineScaleSetNamespaceLister helps list and get WindowsVirtualMachineScaleSets.
type WindowsVirtualMachineScaleSetNamespaceLister interface {
	// List lists all WindowsVirtualMachineScaleSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WindowsVirtualMachineScaleSet, err error)
	// Get retrieves the WindowsVirtualMachineScaleSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WindowsVirtualMachineScaleSet, error)
	WindowsVirtualMachineScaleSetNamespaceListerExpansion
}

// windowsVirtualMachineScaleSetNamespaceLister implements the WindowsVirtualMachineScaleSetNamespaceLister
// interface.
type windowsVirtualMachineScaleSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WindowsVirtualMachineScaleSets in the indexer for a given namespace.
func (s windowsVirtualMachineScaleSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WindowsVirtualMachineScaleSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WindowsVirtualMachineScaleSet))
	})
	return ret, err
}

// Get retrieves the WindowsVirtualMachineScaleSet from the indexer for a given namespace and name.
func (s windowsVirtualMachineScaleSetNamespaceLister) Get(name string) (*v1alpha1.WindowsVirtualMachineScaleSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("windowsvirtualmachinescaleset"), name)
	}
	return obj.(*v1alpha1.WindowsVirtualMachineScaleSet), nil
}