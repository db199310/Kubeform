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

// FakeIotcentralApplications implements IotcentralApplicationInterface
type FakeIotcentralApplications struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var iotcentralapplicationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "iotcentralapplications"}

var iotcentralapplicationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "IotcentralApplication"}

// Get takes name of the iotcentralApplication, and returns the corresponding iotcentralApplication object, and an error if there is any.
func (c *FakeIotcentralApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.IotcentralApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iotcentralapplicationsResource, c.ns, name), &v1alpha1.IotcentralApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotcentralApplication), err
}

// List takes label and field selectors, and returns the list of IotcentralApplications that match those selectors.
func (c *FakeIotcentralApplications) List(opts v1.ListOptions) (result *v1alpha1.IotcentralApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iotcentralapplicationsResource, iotcentralapplicationsKind, c.ns, opts), &v1alpha1.IotcentralApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IotcentralApplicationList{ListMeta: obj.(*v1alpha1.IotcentralApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.IotcentralApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iotcentralApplications.
func (c *FakeIotcentralApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iotcentralapplicationsResource, c.ns, opts))

}

// Create takes the representation of a iotcentralApplication and creates it.  Returns the server's representation of the iotcentralApplication, and an error, if there is any.
func (c *FakeIotcentralApplications) Create(iotcentralApplication *v1alpha1.IotcentralApplication) (result *v1alpha1.IotcentralApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iotcentralapplicationsResource, c.ns, iotcentralApplication), &v1alpha1.IotcentralApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotcentralApplication), err
}

// Update takes the representation of a iotcentralApplication and updates it. Returns the server's representation of the iotcentralApplication, and an error, if there is any.
func (c *FakeIotcentralApplications) Update(iotcentralApplication *v1alpha1.IotcentralApplication) (result *v1alpha1.IotcentralApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iotcentralapplicationsResource, c.ns, iotcentralApplication), &v1alpha1.IotcentralApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotcentralApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIotcentralApplications) UpdateStatus(iotcentralApplication *v1alpha1.IotcentralApplication) (*v1alpha1.IotcentralApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iotcentralapplicationsResource, "status", c.ns, iotcentralApplication), &v1alpha1.IotcentralApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotcentralApplication), err
}

// Delete takes name of the iotcentralApplication and deletes it. Returns an error if one occurs.
func (c *FakeIotcentralApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iotcentralapplicationsResource, c.ns, name), &v1alpha1.IotcentralApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIotcentralApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iotcentralapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IotcentralApplicationList{})
	return err
}

// Patch applies the patch and returns the patched iotcentralApplication.
func (c *FakeIotcentralApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotcentralApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iotcentralapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IotcentralApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotcentralApplication), err
}