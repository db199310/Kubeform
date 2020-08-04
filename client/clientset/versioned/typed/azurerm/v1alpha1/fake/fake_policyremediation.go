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

// FakePolicyRemediations implements PolicyRemediationInterface
type FakePolicyRemediations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var policyremediationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "policyremediations"}

var policyremediationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PolicyRemediation"}

// Get takes name of the policyRemediation, and returns the corresponding policyRemediation object, and an error if there is any.
func (c *FakePolicyRemediations) Get(name string, options v1.GetOptions) (result *v1alpha1.PolicyRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policyremediationsResource, c.ns, name), &v1alpha1.PolicyRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyRemediation), err
}

// List takes label and field selectors, and returns the list of PolicyRemediations that match those selectors.
func (c *FakePolicyRemediations) List(opts v1.ListOptions) (result *v1alpha1.PolicyRemediationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policyremediationsResource, policyremediationsKind, c.ns, opts), &v1alpha1.PolicyRemediationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyRemediationList{ListMeta: obj.(*v1alpha1.PolicyRemediationList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyRemediationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyRemediations.
func (c *FakePolicyRemediations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policyremediationsResource, c.ns, opts))

}

// Create takes the representation of a policyRemediation and creates it.  Returns the server's representation of the policyRemediation, and an error, if there is any.
func (c *FakePolicyRemediations) Create(policyRemediation *v1alpha1.PolicyRemediation) (result *v1alpha1.PolicyRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policyremediationsResource, c.ns, policyRemediation), &v1alpha1.PolicyRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyRemediation), err
}

// Update takes the representation of a policyRemediation and updates it. Returns the server's representation of the policyRemediation, and an error, if there is any.
func (c *FakePolicyRemediations) Update(policyRemediation *v1alpha1.PolicyRemediation) (result *v1alpha1.PolicyRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policyremediationsResource, c.ns, policyRemediation), &v1alpha1.PolicyRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyRemediation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyRemediations) UpdateStatus(policyRemediation *v1alpha1.PolicyRemediation) (*v1alpha1.PolicyRemediation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policyremediationsResource, "status", c.ns, policyRemediation), &v1alpha1.PolicyRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyRemediation), err
}

// Delete takes name of the policyRemediation and deletes it. Returns an error if one occurs.
func (c *FakePolicyRemediations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policyremediationsResource, c.ns, name), &v1alpha1.PolicyRemediation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyRemediations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policyremediationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyRemediationList{})
	return err
}

// Patch applies the patch and returns the patched policyRemediation.
func (c *FakePolicyRemediations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PolicyRemediation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policyremediationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyRemediation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyRemediation), err
}