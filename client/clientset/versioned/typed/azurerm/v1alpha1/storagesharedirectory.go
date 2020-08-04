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

// StorageShareDirectoriesGetter has a method to return a StorageShareDirectoryInterface.
// A group's client should implement this interface.
type StorageShareDirectoriesGetter interface {
	StorageShareDirectories(namespace string) StorageShareDirectoryInterface
}

// StorageShareDirectoryInterface has methods to work with StorageShareDirectory resources.
type StorageShareDirectoryInterface interface {
	Create(*v1alpha1.StorageShareDirectory) (*v1alpha1.StorageShareDirectory, error)
	Update(*v1alpha1.StorageShareDirectory) (*v1alpha1.StorageShareDirectory, error)
	UpdateStatus(*v1alpha1.StorageShareDirectory) (*v1alpha1.StorageShareDirectory, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageShareDirectory, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageShareDirectoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageShareDirectory, err error)
	StorageShareDirectoryExpansion
}

// storageShareDirectories implements StorageShareDirectoryInterface
type storageShareDirectories struct {
	client rest.Interface
	ns     string
}

// newStorageShareDirectories returns a StorageShareDirectories
func newStorageShareDirectories(c *AzurermV1alpha1Client, namespace string) *storageShareDirectories {
	return &storageShareDirectories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageShareDirectory, and returns the corresponding storageShareDirectory object, and an error if there is any.
func (c *storageShareDirectories) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageShareDirectory, err error) {
	result = &v1alpha1.StorageShareDirectory{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageShareDirectories that match those selectors.
func (c *storageShareDirectories) List(opts v1.ListOptions) (result *v1alpha1.StorageShareDirectoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageShareDirectoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageShareDirectories.
func (c *storageShareDirectories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageShareDirectory and creates it.  Returns the server's representation of the storageShareDirectory, and an error, if there is any.
func (c *storageShareDirectories) Create(storageShareDirectory *v1alpha1.StorageShareDirectory) (result *v1alpha1.StorageShareDirectory, err error) {
	result = &v1alpha1.StorageShareDirectory{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		Body(storageShareDirectory).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageShareDirectory and updates it. Returns the server's representation of the storageShareDirectory, and an error, if there is any.
func (c *storageShareDirectories) Update(storageShareDirectory *v1alpha1.StorageShareDirectory) (result *v1alpha1.StorageShareDirectory, err error) {
	result = &v1alpha1.StorageShareDirectory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		Name(storageShareDirectory.Name).
		Body(storageShareDirectory).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageShareDirectories) UpdateStatus(storageShareDirectory *v1alpha1.StorageShareDirectory) (result *v1alpha1.StorageShareDirectory, err error) {
	result = &v1alpha1.StorageShareDirectory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		Name(storageShareDirectory.Name).
		SubResource("status").
		Body(storageShareDirectory).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageShareDirectory and deletes it. Returns an error if one occurs.
func (c *storageShareDirectories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageShareDirectories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagesharedirectories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageShareDirectory.
func (c *storageShareDirectories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageShareDirectory, err error) {
	result = &v1alpha1.StorageShareDirectory{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storagesharedirectories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}