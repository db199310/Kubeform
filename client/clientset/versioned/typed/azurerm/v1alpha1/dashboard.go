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

// DashboardsGetter has a method to return a DashboardInterface.
// A group's client should implement this interface.
type DashboardsGetter interface {
	Dashboards(namespace string) DashboardInterface
}

// DashboardInterface has methods to work with Dashboard resources.
type DashboardInterface interface {
	Create(*v1alpha1.Dashboard) (*v1alpha1.Dashboard, error)
	Update(*v1alpha1.Dashboard) (*v1alpha1.Dashboard, error)
	UpdateStatus(*v1alpha1.Dashboard) (*v1alpha1.Dashboard, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Dashboard, error)
	List(opts v1.ListOptions) (*v1alpha1.DashboardList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Dashboard, err error)
	DashboardExpansion
}

// dashboards implements DashboardInterface
type dashboards struct {
	client rest.Interface
	ns     string
}

// newDashboards returns a Dashboards
func newDashboards(c *AzurermV1alpha1Client, namespace string) *dashboards {
	return &dashboards{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dashboard, and returns the corresponding dashboard object, and an error if there is any.
func (c *dashboards) Get(name string, options v1.GetOptions) (result *v1alpha1.Dashboard, err error) {
	result = &v1alpha1.Dashboard{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dashboards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Dashboards that match those selectors.
func (c *dashboards) List(opts v1.ListOptions) (result *v1alpha1.DashboardList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DashboardList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dashboards.
func (c *dashboards) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dashboard and creates it.  Returns the server's representation of the dashboard, and an error, if there is any.
func (c *dashboards) Create(dashboard *v1alpha1.Dashboard) (result *v1alpha1.Dashboard, err error) {
	result = &v1alpha1.Dashboard{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dashboards").
		Body(dashboard).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dashboard and updates it. Returns the server's representation of the dashboard, and an error, if there is any.
func (c *dashboards) Update(dashboard *v1alpha1.Dashboard) (result *v1alpha1.Dashboard, err error) {
	result = &v1alpha1.Dashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dashboards").
		Name(dashboard.Name).
		Body(dashboard).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dashboards) UpdateStatus(dashboard *v1alpha1.Dashboard) (result *v1alpha1.Dashboard, err error) {
	result = &v1alpha1.Dashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dashboards").
		Name(dashboard.Name).
		SubResource("status").
		Body(dashboard).
		Do().
		Into(result)
	return
}

// Delete takes name of the dashboard and deletes it. Returns an error if one occurs.
func (c *dashboards) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dashboards").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dashboards) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dashboards").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dashboard.
func (c *dashboards) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Dashboard, err error) {
	result = &v1alpha1.Dashboard{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dashboards").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}