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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// FakeAzureFnApps implements AzureFnAppInterface
type FakeAzureFnApps struct {
	Fake *FakeModulesV1alpha1
	ns   string
}

var azurefnappsResource = schema.GroupVersionResource{Group: "modules.kubeform.com", Version: "v1alpha1", Resource: "azurefnapps"}

var azurefnappsKind = schema.GroupVersionKind{Group: "modules.kubeform.com", Version: "v1alpha1", Kind: "AzureFnApp"}

// Get takes name of the azureFnApp, and returns the corresponding azureFnApp object, and an error if there is any.
func (c *FakeAzureFnApps) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureFnApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurefnappsResource, c.ns, name), &v1alpha1.AzureFnApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureFnApp), err
}

// List takes label and field selectors, and returns the list of AzureFnApps that match those selectors.
func (c *FakeAzureFnApps) List(opts v1.ListOptions) (result *v1alpha1.AzureFnAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurefnappsResource, azurefnappsKind, c.ns, opts), &v1alpha1.AzureFnAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureFnAppList{ListMeta: obj.(*v1alpha1.AzureFnAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureFnAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureFnApps.
func (c *FakeAzureFnApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurefnappsResource, c.ns, opts))

}

// Create takes the representation of a azureFnApp and creates it.  Returns the server's representation of the azureFnApp, and an error, if there is any.
func (c *FakeAzureFnApps) Create(azureFnApp *v1alpha1.AzureFnApp) (result *v1alpha1.AzureFnApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurefnappsResource, c.ns, azureFnApp), &v1alpha1.AzureFnApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureFnApp), err
}

// Update takes the representation of a azureFnApp and updates it. Returns the server's representation of the azureFnApp, and an error, if there is any.
func (c *FakeAzureFnApps) Update(azureFnApp *v1alpha1.AzureFnApp) (result *v1alpha1.AzureFnApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurefnappsResource, c.ns, azureFnApp), &v1alpha1.AzureFnApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureFnApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureFnApps) UpdateStatus(azureFnApp *v1alpha1.AzureFnApp) (*v1alpha1.AzureFnApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurefnappsResource, "status", c.ns, azureFnApp), &v1alpha1.AzureFnApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureFnApp), err
}

// Delete takes name of the azureFnApp and deletes it. Returns an error if one occurs.
func (c *FakeAzureFnApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurefnappsResource, c.ns, name), &v1alpha1.AzureFnApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureFnApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurefnappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureFnAppList{})
	return err
}

// Patch applies the patch and returns the patched azureFnApp.
func (c *FakeAzureFnApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureFnApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurefnappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureFnApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureFnApp), err
}