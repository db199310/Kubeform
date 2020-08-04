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

// FakeApiManagementNamedValues implements ApiManagementNamedValueInterface
type FakeApiManagementNamedValues struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementnamedvaluesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementnamedvalues"}

var apimanagementnamedvaluesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementNamedValue"}

// Get takes name of the apiManagementNamedValue, and returns the corresponding apiManagementNamedValue object, and an error if there is any.
func (c *FakeApiManagementNamedValues) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementNamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementnamedvaluesResource, c.ns, name), &v1alpha1.ApiManagementNamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), err
}

// List takes label and field selectors, and returns the list of ApiManagementNamedValues that match those selectors.
func (c *FakeApiManagementNamedValues) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementNamedValueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementnamedvaluesResource, apimanagementnamedvaluesKind, c.ns, opts), &v1alpha1.ApiManagementNamedValueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementNamedValueList{ListMeta: obj.(*v1alpha1.ApiManagementNamedValueList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementNamedValueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementNamedValues.
func (c *FakeApiManagementNamedValues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementnamedvaluesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementNamedValue and creates it.  Returns the server's representation of the apiManagementNamedValue, and an error, if there is any.
func (c *FakeApiManagementNamedValues) Create(apiManagementNamedValue *v1alpha1.ApiManagementNamedValue) (result *v1alpha1.ApiManagementNamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementnamedvaluesResource, c.ns, apiManagementNamedValue), &v1alpha1.ApiManagementNamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), err
}

// Update takes the representation of a apiManagementNamedValue and updates it. Returns the server's representation of the apiManagementNamedValue, and an error, if there is any.
func (c *FakeApiManagementNamedValues) Update(apiManagementNamedValue *v1alpha1.ApiManagementNamedValue) (result *v1alpha1.ApiManagementNamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementnamedvaluesResource, c.ns, apiManagementNamedValue), &v1alpha1.ApiManagementNamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementNamedValues) UpdateStatus(apiManagementNamedValue *v1alpha1.ApiManagementNamedValue) (*v1alpha1.ApiManagementNamedValue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementnamedvaluesResource, "status", c.ns, apiManagementNamedValue), &v1alpha1.ApiManagementNamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), err
}

// Delete takes name of the apiManagementNamedValue and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementNamedValues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementnamedvaluesResource, c.ns, name), &v1alpha1.ApiManagementNamedValue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementNamedValues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementnamedvaluesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementNamedValueList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementNamedValue.
func (c *FakeApiManagementNamedValues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementNamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementnamedvaluesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementNamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), err
}