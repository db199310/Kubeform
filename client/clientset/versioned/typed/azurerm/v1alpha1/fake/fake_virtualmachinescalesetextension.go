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

// FakeVirtualMachineScaleSetExtensions implements VirtualMachineScaleSetExtensionInterface
type FakeVirtualMachineScaleSetExtensions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var virtualmachinescalesetextensionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "virtualmachinescalesetextensions"}

var virtualmachinescalesetextensionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "VirtualMachineScaleSetExtension"}

// Get takes name of the virtualMachineScaleSetExtension, and returns the corresponding virtualMachineScaleSetExtension object, and an error if there is any.
func (c *FakeVirtualMachineScaleSetExtensions) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinescalesetextensionsResource, c.ns, name), &v1alpha1.VirtualMachineScaleSetExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), err
}

// List takes label and field selectors, and returns the list of VirtualMachineScaleSetExtensions that match those selectors.
func (c *FakeVirtualMachineScaleSetExtensions) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineScaleSetExtensionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinescalesetextensionsResource, virtualmachinescalesetextensionsKind, c.ns, opts), &v1alpha1.VirtualMachineScaleSetExtensionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineScaleSetExtensionList{ListMeta: obj.(*v1alpha1.VirtualMachineScaleSetExtensionList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineScaleSetExtensionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineScaleSetExtensions.
func (c *FakeVirtualMachineScaleSetExtensions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinescalesetextensionsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineScaleSetExtension and creates it.  Returns the server's representation of the virtualMachineScaleSetExtension, and an error, if there is any.
func (c *FakeVirtualMachineScaleSetExtensions) Create(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinescalesetextensionsResource, c.ns, virtualMachineScaleSetExtension), &v1alpha1.VirtualMachineScaleSetExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), err
}

// Update takes the representation of a virtualMachineScaleSetExtension and updates it. Returns the server's representation of the virtualMachineScaleSetExtension, and an error, if there is any.
func (c *FakeVirtualMachineScaleSetExtensions) Update(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinescalesetextensionsResource, c.ns, virtualMachineScaleSetExtension), &v1alpha1.VirtualMachineScaleSetExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineScaleSetExtensions) UpdateStatus(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (*v1alpha1.VirtualMachineScaleSetExtension, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinescalesetextensionsResource, "status", c.ns, virtualMachineScaleSetExtension), &v1alpha1.VirtualMachineScaleSetExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), err
}

// Delete takes name of the virtualMachineScaleSetExtension and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineScaleSetExtensions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinescalesetextensionsResource, c.ns, name), &v1alpha1.VirtualMachineScaleSetExtension{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineScaleSetExtensions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachinescalesetextensionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineScaleSetExtensionList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineScaleSetExtension.
func (c *FakeVirtualMachineScaleSetExtensions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinescalesetextensionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineScaleSetExtension{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), err
}