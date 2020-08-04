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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// FunctionAppSlotsGetter has a method to return a FunctionAppSlotInterface.
// A group's client should implement this interface.
type FunctionAppSlotsGetter interface {
	FunctionAppSlots(namespace string) FunctionAppSlotInterface
}

// FunctionAppSlotInterface has methods to work with FunctionAppSlot resources.
type FunctionAppSlotInterface interface {
	Create(*v1alpha1.FunctionAppSlot) (*v1alpha1.FunctionAppSlot, error)
	Update(*v1alpha1.FunctionAppSlot) (*v1alpha1.FunctionAppSlot, error)
	UpdateStatus(*v1alpha1.FunctionAppSlot) (*v1alpha1.FunctionAppSlot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FunctionAppSlot, error)
	List(opts v1.ListOptions) (*v1alpha1.FunctionAppSlotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FunctionAppSlot, err error)
	FunctionAppSlotExpansion
}

// functionAppSlots implements FunctionAppSlotInterface
type functionAppSlots struct {
	client rest.Interface
	ns     string
}

// newFunctionAppSlots returns a FunctionAppSlots
func newFunctionAppSlots(c *AzurermV1alpha1Client, namespace string) *functionAppSlots {
	return &functionAppSlots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the functionAppSlot, and returns the corresponding functionAppSlot object, and an error if there is any.
func (c *functionAppSlots) Get(name string, options v1.GetOptions) (result *v1alpha1.FunctionAppSlot, err error) {
	result = &v1alpha1.FunctionAppSlot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functionappslots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FunctionAppSlots that match those selectors.
func (c *functionAppSlots) List(opts v1.ListOptions) (result *v1alpha1.FunctionAppSlotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FunctionAppSlotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functionappslots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested functionAppSlots.
func (c *functionAppSlots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("functionappslots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a functionAppSlot and creates it.  Returns the server's representation of the functionAppSlot, and an error, if there is any.
func (c *functionAppSlots) Create(functionAppSlot *v1alpha1.FunctionAppSlot) (result *v1alpha1.FunctionAppSlot, err error) {
	result = &v1alpha1.FunctionAppSlot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("functionappslots").
		Body(functionAppSlot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a functionAppSlot and updates it. Returns the server's representation of the functionAppSlot, and an error, if there is any.
func (c *functionAppSlots) Update(functionAppSlot *v1alpha1.FunctionAppSlot) (result *v1alpha1.FunctionAppSlot, err error) {
	result = &v1alpha1.FunctionAppSlot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functionappslots").
		Name(functionAppSlot.Name).
		Body(functionAppSlot).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *functionAppSlots) UpdateStatus(functionAppSlot *v1alpha1.FunctionAppSlot) (result *v1alpha1.FunctionAppSlot, err error) {
	result = &v1alpha1.FunctionAppSlot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functionappslots").
		Name(functionAppSlot.Name).
		SubResource("status").
		Body(functionAppSlot).
		Do().
		Into(result)
	return
}

// Delete takes name of the functionAppSlot and deletes it. Returns an error if one occurs.
func (c *functionAppSlots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functionappslots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *functionAppSlots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functionappslots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched functionAppSlot.
func (c *functionAppSlots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FunctionAppSlot, err error) {
	result = &v1alpha1.FunctionAppSlot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("functionappslots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}