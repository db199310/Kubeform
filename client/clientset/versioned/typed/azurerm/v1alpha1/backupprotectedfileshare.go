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

// BackupProtectedFileSharesGetter has a method to return a BackupProtectedFileShareInterface.
// A group's client should implement this interface.
type BackupProtectedFileSharesGetter interface {
	BackupProtectedFileShares(namespace string) BackupProtectedFileShareInterface
}

// BackupProtectedFileShareInterface has methods to work with BackupProtectedFileShare resources.
type BackupProtectedFileShareInterface interface {
	Create(*v1alpha1.BackupProtectedFileShare) (*v1alpha1.BackupProtectedFileShare, error)
	Update(*v1alpha1.BackupProtectedFileShare) (*v1alpha1.BackupProtectedFileShare, error)
	UpdateStatus(*v1alpha1.BackupProtectedFileShare) (*v1alpha1.BackupProtectedFileShare, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BackupProtectedFileShare, error)
	List(opts v1.ListOptions) (*v1alpha1.BackupProtectedFileShareList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupProtectedFileShare, err error)
	BackupProtectedFileShareExpansion
}

// backupProtectedFileShares implements BackupProtectedFileShareInterface
type backupProtectedFileShares struct {
	client rest.Interface
	ns     string
}

// newBackupProtectedFileShares returns a BackupProtectedFileShares
func newBackupProtectedFileShares(c *AzurermV1alpha1Client, namespace string) *backupProtectedFileShares {
	return &backupProtectedFileShares{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupProtectedFileShare, and returns the corresponding backupProtectedFileShare object, and an error if there is any.
func (c *backupProtectedFileShares) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupProtectedFileShare, err error) {
	result = &v1alpha1.BackupProtectedFileShare{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupProtectedFileShares that match those selectors.
func (c *backupProtectedFileShares) List(opts v1.ListOptions) (result *v1alpha1.BackupProtectedFileShareList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackupProtectedFileShareList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupProtectedFileShares.
func (c *backupProtectedFileShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a backupProtectedFileShare and creates it.  Returns the server's representation of the backupProtectedFileShare, and an error, if there is any.
func (c *backupProtectedFileShares) Create(backupProtectedFileShare *v1alpha1.BackupProtectedFileShare) (result *v1alpha1.BackupProtectedFileShare, err error) {
	result = &v1alpha1.BackupProtectedFileShare{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		Body(backupProtectedFileShare).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupProtectedFileShare and updates it. Returns the server's representation of the backupProtectedFileShare, and an error, if there is any.
func (c *backupProtectedFileShares) Update(backupProtectedFileShare *v1alpha1.BackupProtectedFileShare) (result *v1alpha1.BackupProtectedFileShare, err error) {
	result = &v1alpha1.BackupProtectedFileShare{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		Name(backupProtectedFileShare.Name).
		Body(backupProtectedFileShare).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backupProtectedFileShares) UpdateStatus(backupProtectedFileShare *v1alpha1.BackupProtectedFileShare) (result *v1alpha1.BackupProtectedFileShare, err error) {
	result = &v1alpha1.BackupProtectedFileShare{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		Name(backupProtectedFileShare.Name).
		SubResource("status").
		Body(backupProtectedFileShare).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupProtectedFileShare and deletes it. Returns an error if one occurs.
func (c *backupProtectedFileShares) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupProtectedFileShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupProtectedFileShare.
func (c *backupProtectedFileShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupProtectedFileShare, err error) {
	result = &v1alpha1.BackupProtectedFileShare{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupprotectedfileshares").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}