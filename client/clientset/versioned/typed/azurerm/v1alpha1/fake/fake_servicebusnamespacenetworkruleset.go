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

// FakeServicebusNamespaceNetworkRuleSets implements ServicebusNamespaceNetworkRuleSetInterface
type FakeServicebusNamespaceNetworkRuleSets struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicebusnamespacenetworkrulesetsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicebusnamespacenetworkrulesets"}

var servicebusnamespacenetworkrulesetsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicebusNamespaceNetworkRuleSet"}

// Get takes name of the servicebusNamespaceNetworkRuleSet, and returns the corresponding servicebusNamespaceNetworkRuleSet object, and an error if there is any.
func (c *FakeServicebusNamespaceNetworkRuleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.ServicebusNamespaceNetworkRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebusnamespacenetworkrulesetsResource, c.ns, name), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSet), err
}

// List takes label and field selectors, and returns the list of ServicebusNamespaceNetworkRuleSets that match those selectors.
func (c *FakeServicebusNamespaceNetworkRuleSets) List(opts v1.ListOptions) (result *v1alpha1.ServicebusNamespaceNetworkRuleSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebusnamespacenetworkrulesetsResource, servicebusnamespacenetworkrulesetsKind, c.ns, opts), &v1alpha1.ServicebusNamespaceNetworkRuleSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicebusNamespaceNetworkRuleSetList{ListMeta: obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicebusNamespaceNetworkRuleSets.
func (c *FakeServicebusNamespaceNetworkRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebusnamespacenetworkrulesetsResource, c.ns, opts))

}

// Create takes the representation of a servicebusNamespaceNetworkRuleSet and creates it.  Returns the server's representation of the servicebusNamespaceNetworkRuleSet, and an error, if there is any.
func (c *FakeServicebusNamespaceNetworkRuleSets) Create(servicebusNamespaceNetworkRuleSet *v1alpha1.ServicebusNamespaceNetworkRuleSet) (result *v1alpha1.ServicebusNamespaceNetworkRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebusnamespacenetworkrulesetsResource, c.ns, servicebusNamespaceNetworkRuleSet), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSet), err
}

// Update takes the representation of a servicebusNamespaceNetworkRuleSet and updates it. Returns the server's representation of the servicebusNamespaceNetworkRuleSet, and an error, if there is any.
func (c *FakeServicebusNamespaceNetworkRuleSets) Update(servicebusNamespaceNetworkRuleSet *v1alpha1.ServicebusNamespaceNetworkRuleSet) (result *v1alpha1.ServicebusNamespaceNetworkRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebusnamespacenetworkrulesetsResource, c.ns, servicebusNamespaceNetworkRuleSet), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicebusNamespaceNetworkRuleSets) UpdateStatus(servicebusNamespaceNetworkRuleSet *v1alpha1.ServicebusNamespaceNetworkRuleSet) (*v1alpha1.ServicebusNamespaceNetworkRuleSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebusnamespacenetworkrulesetsResource, "status", c.ns, servicebusNamespaceNetworkRuleSet), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSet), err
}

// Delete takes name of the servicebusNamespaceNetworkRuleSet and deletes it. Returns an error if one occurs.
func (c *FakeServicebusNamespaceNetworkRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicebusnamespacenetworkrulesetsResource, c.ns, name), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicebusNamespaceNetworkRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebusnamespacenetworkrulesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicebusNamespaceNetworkRuleSetList{})
	return err
}

// Patch applies the patch and returns the patched servicebusNamespaceNetworkRuleSet.
func (c *FakeServicebusNamespaceNetworkRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusNamespaceNetworkRuleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebusnamespacenetworkrulesetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicebusNamespaceNetworkRuleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceNetworkRuleSet), err
}