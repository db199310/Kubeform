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

// PrivateDNSZoneLister helps list PrivateDNSZones.
type PrivateDNSZoneLister interface {
	// List lists all PrivateDNSZones in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSZone, err error)
	// PrivateDNSZones returns an object that can list and get PrivateDNSZones.
	PrivateDNSZones(namespace string) PrivateDNSZoneNamespaceLister
	PrivateDNSZoneListerExpansion
}

// privateDNSZoneLister implements the PrivateDNSZoneLister interface.
type privateDNSZoneLister struct {
	indexer cache.Indexer
}

// NewPrivateDNSZoneLister returns a new PrivateDNSZoneLister.
func NewPrivateDNSZoneLister(indexer cache.Indexer) PrivateDNSZoneLister {
	return &privateDNSZoneLister{indexer: indexer}
}

// List lists all PrivateDNSZones in the indexer.
func (s *privateDNSZoneLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSZone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSZone))
	})
	return ret, err
}

// PrivateDNSZones returns an object that can list and get PrivateDNSZones.
func (s *privateDNSZoneLister) PrivateDNSZones(namespace string) PrivateDNSZoneNamespaceLister {
	return privateDNSZoneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateDNSZoneNamespaceLister helps list and get PrivateDNSZones.
type PrivateDNSZoneNamespaceLister interface {
	// List lists all PrivateDNSZones in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSZone, err error)
	// Get retrieves the PrivateDNSZone from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PrivateDNSZone, error)
	PrivateDNSZoneNamespaceListerExpansion
}

// privateDNSZoneNamespaceLister implements the PrivateDNSZoneNamespaceLister
// interface.
type privateDNSZoneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateDNSZones in the indexer for a given namespace.
func (s privateDNSZoneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSZone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSZone))
	})
	return ret, err
}

// Get retrieves the PrivateDNSZone from the indexer for a given namespace and name.
func (s privateDNSZoneNamespaceLister) Get(name string) (*v1alpha1.PrivateDNSZone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatednszone"), name)
	}
	return obj.(*v1alpha1.PrivateDNSZone), nil
}