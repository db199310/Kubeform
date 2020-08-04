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

// FakeStorageAccountCustomerManagedKeys implements StorageAccountCustomerManagedKeyInterface
type FakeStorageAccountCustomerManagedKeys struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var storageaccountcustomermanagedkeysResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "storageaccountcustomermanagedkeys"}

var storageaccountcustomermanagedkeysKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StorageAccountCustomerManagedKey"}

// Get takes name of the storageAccountCustomerManagedKey, and returns the corresponding storageAccountCustomerManagedKey object, and an error if there is any.
func (c *FakeStorageAccountCustomerManagedKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageAccountCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storageaccountcustomermanagedkeysResource, c.ns, name), &v1alpha1.StorageAccountCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountCustomerManagedKey), err
}

// List takes label and field selectors, and returns the list of StorageAccountCustomerManagedKeys that match those selectors.
func (c *FakeStorageAccountCustomerManagedKeys) List(opts v1.ListOptions) (result *v1alpha1.StorageAccountCustomerManagedKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storageaccountcustomermanagedkeysResource, storageaccountcustomermanagedkeysKind, c.ns, opts), &v1alpha1.StorageAccountCustomerManagedKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageAccountCustomerManagedKeyList{ListMeta: obj.(*v1alpha1.StorageAccountCustomerManagedKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageAccountCustomerManagedKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageAccountCustomerManagedKeys.
func (c *FakeStorageAccountCustomerManagedKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storageaccountcustomermanagedkeysResource, c.ns, opts))

}

// Create takes the representation of a storageAccountCustomerManagedKey and creates it.  Returns the server's representation of the storageAccountCustomerManagedKey, and an error, if there is any.
func (c *FakeStorageAccountCustomerManagedKeys) Create(storageAccountCustomerManagedKey *v1alpha1.StorageAccountCustomerManagedKey) (result *v1alpha1.StorageAccountCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storageaccountcustomermanagedkeysResource, c.ns, storageAccountCustomerManagedKey), &v1alpha1.StorageAccountCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountCustomerManagedKey), err
}

// Update takes the representation of a storageAccountCustomerManagedKey and updates it. Returns the server's representation of the storageAccountCustomerManagedKey, and an error, if there is any.
func (c *FakeStorageAccountCustomerManagedKeys) Update(storageAccountCustomerManagedKey *v1alpha1.StorageAccountCustomerManagedKey) (result *v1alpha1.StorageAccountCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storageaccountcustomermanagedkeysResource, c.ns, storageAccountCustomerManagedKey), &v1alpha1.StorageAccountCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountCustomerManagedKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageAccountCustomerManagedKeys) UpdateStatus(storageAccountCustomerManagedKey *v1alpha1.StorageAccountCustomerManagedKey) (*v1alpha1.StorageAccountCustomerManagedKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storageaccountcustomermanagedkeysResource, "status", c.ns, storageAccountCustomerManagedKey), &v1alpha1.StorageAccountCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountCustomerManagedKey), err
}

// Delete takes name of the storageAccountCustomerManagedKey and deletes it. Returns an error if one occurs.
func (c *FakeStorageAccountCustomerManagedKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storageaccountcustomermanagedkeysResource, c.ns, name), &v1alpha1.StorageAccountCustomerManagedKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageAccountCustomerManagedKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storageaccountcustomermanagedkeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageAccountCustomerManagedKeyList{})
	return err
}

// Patch applies the patch and returns the patched storageAccountCustomerManagedKey.
func (c *FakeStorageAccountCustomerManagedKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageAccountCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storageaccountcustomermanagedkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageAccountCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountCustomerManagedKey), err
}