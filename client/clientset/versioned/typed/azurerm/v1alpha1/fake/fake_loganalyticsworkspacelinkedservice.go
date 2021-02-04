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

// FakeLogAnalyticsWorkspaceLinkedServices implements LogAnalyticsWorkspaceLinkedServiceInterface
type FakeLogAnalyticsWorkspaceLinkedServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var loganalyticsworkspacelinkedservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "loganalyticsworkspacelinkedservices"}

var loganalyticsworkspacelinkedservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogAnalyticsWorkspaceLinkedService"}

// Get takes name of the logAnalyticsWorkspaceLinkedService, and returns the corresponding logAnalyticsWorkspaceLinkedService object, and an error if there is any.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsWorkspaceLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loganalyticsworkspacelinkedservicesResource, c.ns, name), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedService), err
}

// List takes label and field selectors, and returns the list of LogAnalyticsWorkspaceLinkedServices that match those selectors.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogAnalyticsWorkspaceLinkedServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loganalyticsworkspacelinkedservicesResource, loganalyticsworkspacelinkedservicesKind, c.ns, opts), &v1alpha1.LogAnalyticsWorkspaceLinkedServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogAnalyticsWorkspaceLinkedServiceList{ListMeta: obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logAnalyticsWorkspaceLinkedServices.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loganalyticsworkspacelinkedservicesResource, c.ns, opts))

}

// Create takes the representation of a logAnalyticsWorkspaceLinkedService and creates it.  Returns the server's representation of the logAnalyticsWorkspaceLinkedService, and an error, if there is any.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Create(ctx context.Context, logAnalyticsWorkspaceLinkedService *v1alpha1.LogAnalyticsWorkspaceLinkedService, opts v1.CreateOptions) (result *v1alpha1.LogAnalyticsWorkspaceLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loganalyticsworkspacelinkedservicesResource, c.ns, logAnalyticsWorkspaceLinkedService), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedService), err
}

// Update takes the representation of a logAnalyticsWorkspaceLinkedService and updates it. Returns the server's representation of the logAnalyticsWorkspaceLinkedService, and an error, if there is any.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Update(ctx context.Context, logAnalyticsWorkspaceLinkedService *v1alpha1.LogAnalyticsWorkspaceLinkedService, opts v1.UpdateOptions) (result *v1alpha1.LogAnalyticsWorkspaceLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loganalyticsworkspacelinkedservicesResource, c.ns, logAnalyticsWorkspaceLinkedService), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogAnalyticsWorkspaceLinkedServices) UpdateStatus(ctx context.Context, logAnalyticsWorkspaceLinkedService *v1alpha1.LogAnalyticsWorkspaceLinkedService, opts v1.UpdateOptions) (*v1alpha1.LogAnalyticsWorkspaceLinkedService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loganalyticsworkspacelinkedservicesResource, "status", c.ns, logAnalyticsWorkspaceLinkedService), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedService), err
}

// Delete takes name of the logAnalyticsWorkspaceLinkedService and deletes it. Returns an error if one occurs.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loganalyticsworkspacelinkedservicesResource, c.ns, name), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loganalyticsworkspacelinkedservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogAnalyticsWorkspaceLinkedServiceList{})
	return err
}

// Patch applies the patch and returns the patched logAnalyticsWorkspaceLinkedService.
func (c *FakeLogAnalyticsWorkspaceLinkedServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogAnalyticsWorkspaceLinkedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loganalyticsworkspacelinkedservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogAnalyticsWorkspaceLinkedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspaceLinkedService), err
}