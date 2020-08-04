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

// FakePacketCaptures implements PacketCaptureInterface
type FakePacketCaptures struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var packetcapturesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "packetcaptures"}

var packetcapturesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PacketCapture"}

// Get takes name of the packetCapture, and returns the corresponding packetCapture object, and an error if there is any.
func (c *FakePacketCaptures) Get(name string, options v1.GetOptions) (result *v1alpha1.PacketCapture, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packetcapturesResource, c.ns, name), &v1alpha1.PacketCapture{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketCapture), err
}

// List takes label and field selectors, and returns the list of PacketCaptures that match those selectors.
func (c *FakePacketCaptures) List(opts v1.ListOptions) (result *v1alpha1.PacketCaptureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packetcapturesResource, packetcapturesKind, c.ns, opts), &v1alpha1.PacketCaptureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PacketCaptureList{ListMeta: obj.(*v1alpha1.PacketCaptureList).ListMeta}
	for _, item := range obj.(*v1alpha1.PacketCaptureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packetCaptures.
func (c *FakePacketCaptures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packetcapturesResource, c.ns, opts))

}

// Create takes the representation of a packetCapture and creates it.  Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *FakePacketCaptures) Create(packetCapture *v1alpha1.PacketCapture) (result *v1alpha1.PacketCapture, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packetcapturesResource, c.ns, packetCapture), &v1alpha1.PacketCapture{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketCapture), err
}

// Update takes the representation of a packetCapture and updates it. Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *FakePacketCaptures) Update(packetCapture *v1alpha1.PacketCapture) (result *v1alpha1.PacketCapture, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packetcapturesResource, c.ns, packetCapture), &v1alpha1.PacketCapture{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketCapture), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePacketCaptures) UpdateStatus(packetCapture *v1alpha1.PacketCapture) (*v1alpha1.PacketCapture, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packetcapturesResource, "status", c.ns, packetCapture), &v1alpha1.PacketCapture{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketCapture), err
}

// Delete takes name of the packetCapture and deletes it. Returns an error if one occurs.
func (c *FakePacketCaptures) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(packetcapturesResource, c.ns, name), &v1alpha1.PacketCapture{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePacketCaptures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packetcapturesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PacketCaptureList{})
	return err
}

// Patch applies the patch and returns the patched packetCapture.
func (c *FakePacketCaptures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketCapture, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packetcapturesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PacketCapture{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PacketCapture), err
}