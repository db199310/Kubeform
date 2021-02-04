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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzureadApplications implements AzureadApplicationInterface
type FakeAzureadApplications struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var azureadapplicationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azureadapplications"}

var azureadapplicationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzureadApplication"}

// Get takes name of the azureadApplication, and returns the corresponding azureadApplication object, and an error if there is any.
func (c *FakeAzureadApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureadApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureadapplicationsResource, c.ns, name), &v1alpha1.AzureadApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadApplication), err
}

// List takes label and field selectors, and returns the list of AzureadApplications that match those selectors.
func (c *FakeAzureadApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureadApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureadapplicationsResource, azureadapplicationsKind, c.ns, opts), &v1alpha1.AzureadApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureadApplicationList{ListMeta: obj.(*v1alpha1.AzureadApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureadApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureadApplications.
func (c *FakeAzureadApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureadapplicationsResource, c.ns, opts))

}

// Create takes the representation of a azureadApplication and creates it.  Returns the server's representation of the azureadApplication, and an error, if there is any.
func (c *FakeAzureadApplications) Create(ctx context.Context, azureadApplication *v1alpha1.AzureadApplication, opts v1.CreateOptions) (result *v1alpha1.AzureadApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureadapplicationsResource, c.ns, azureadApplication), &v1alpha1.AzureadApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadApplication), err
}

// Update takes the representation of a azureadApplication and updates it. Returns the server's representation of the azureadApplication, and an error, if there is any.
func (c *FakeAzureadApplications) Update(ctx context.Context, azureadApplication *v1alpha1.AzureadApplication, opts v1.UpdateOptions) (result *v1alpha1.AzureadApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureadapplicationsResource, c.ns, azureadApplication), &v1alpha1.AzureadApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureadApplications) UpdateStatus(ctx context.Context, azureadApplication *v1alpha1.AzureadApplication, opts v1.UpdateOptions) (*v1alpha1.AzureadApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureadapplicationsResource, "status", c.ns, azureadApplication), &v1alpha1.AzureadApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadApplication), err
}

// Delete takes name of the azureadApplication and deletes it. Returns an error if one occurs.
func (c *FakeAzureadApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureadapplicationsResource, c.ns, name), &v1alpha1.AzureadApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureadApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureadapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureadApplicationList{})
	return err
}

// Patch applies the patch and returns the patched azureadApplication.
func (c *FakeAzureadApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureadApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureadapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureadApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadApplication), err
}