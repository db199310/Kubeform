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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeDataFactoryLinkedServiceSQLServers implements DataFactoryLinkedServiceSQLServerInterface
type FakeDataFactoryLinkedServiceSQLServers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datafactorylinkedservicesqlserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactorylinkedservicesqlservers"}

var datafactorylinkedservicesqlserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactoryLinkedServiceSQLServer"}

// Get takes name of the dataFactoryLinkedServiceSQLServer, and returns the corresponding dataFactoryLinkedServiceSQLServer object, and an error if there is any.
func (c *FakeDataFactoryLinkedServiceSQLServers) Get(name string, options v1.GetOptions) (result *v1alpha1.DataFactoryLinkedServiceSQLServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafactorylinkedservicesqlserversResource, c.ns, name), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceSQLServer), err
}

// List takes label and field selectors, and returns the list of DataFactoryLinkedServiceSQLServers that match those selectors.
func (c *FakeDataFactoryLinkedServiceSQLServers) List(opts v1.ListOptions) (result *v1alpha1.DataFactoryLinkedServiceSQLServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafactorylinkedservicesqlserversResource, datafactorylinkedservicesqlserversKind, c.ns, opts), &v1alpha1.DataFactoryLinkedServiceSQLServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryLinkedServiceSQLServerList{ListMeta: obj.(*v1alpha1.DataFactoryLinkedServiceSQLServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryLinkedServiceSQLServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactoryLinkedServiceSQLServers.
func (c *FakeDataFactoryLinkedServiceSQLServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafactorylinkedservicesqlserversResource, c.ns, opts))

}

// Create takes the representation of a dataFactoryLinkedServiceSQLServer and creates it.  Returns the server's representation of the dataFactoryLinkedServiceSQLServer, and an error, if there is any.
func (c *FakeDataFactoryLinkedServiceSQLServers) Create(dataFactoryLinkedServiceSQLServer *v1alpha1.DataFactoryLinkedServiceSQLServer) (result *v1alpha1.DataFactoryLinkedServiceSQLServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafactorylinkedservicesqlserversResource, c.ns, dataFactoryLinkedServiceSQLServer), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceSQLServer), err
}

// Update takes the representation of a dataFactoryLinkedServiceSQLServer and updates it. Returns the server's representation of the dataFactoryLinkedServiceSQLServer, and an error, if there is any.
func (c *FakeDataFactoryLinkedServiceSQLServers) Update(dataFactoryLinkedServiceSQLServer *v1alpha1.DataFactoryLinkedServiceSQLServer) (result *v1alpha1.DataFactoryLinkedServiceSQLServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafactorylinkedservicesqlserversResource, c.ns, dataFactoryLinkedServiceSQLServer), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceSQLServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactoryLinkedServiceSQLServers) UpdateStatus(dataFactoryLinkedServiceSQLServer *v1alpha1.DataFactoryLinkedServiceSQLServer) (*v1alpha1.DataFactoryLinkedServiceSQLServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafactorylinkedservicesqlserversResource, "status", c.ns, dataFactoryLinkedServiceSQLServer), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceSQLServer), err
}

// Delete takes name of the dataFactoryLinkedServiceSQLServer and deletes it. Returns an error if one occurs.
func (c *FakeDataFactoryLinkedServiceSQLServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datafactorylinkedservicesqlserversResource, c.ns, name), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactoryLinkedServiceSQLServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafactorylinkedservicesqlserversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryLinkedServiceSQLServerList{})
	return err
}

// Patch applies the patch and returns the patched dataFactoryLinkedServiceSQLServer.
func (c *FakeDataFactoryLinkedServiceSQLServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataFactoryLinkedServiceSQLServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafactorylinkedservicesqlserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataFactoryLinkedServiceSQLServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceSQLServer), err
}