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

// LogAnalyticsDatasourceWindowsEventsGetter has a method to return a LogAnalyticsDatasourceWindowsEventInterface.
// A group's client should implement this interface.
type LogAnalyticsDatasourceWindowsEventsGetter interface {
	LogAnalyticsDatasourceWindowsEvents(namespace string) LogAnalyticsDatasourceWindowsEventInterface
}

// LogAnalyticsDatasourceWindowsEventInterface has methods to work with LogAnalyticsDatasourceWindowsEvent resources.
type LogAnalyticsDatasourceWindowsEventInterface interface {
	Create(*v1alpha1.LogAnalyticsDatasourceWindowsEvent) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error)
	Update(*v1alpha1.LogAnalyticsDatasourceWindowsEvent) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error)
	UpdateStatus(*v1alpha1.LogAnalyticsDatasourceWindowsEvent) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LogAnalyticsDatasourceWindowsEvent, error)
	List(opts v1.ListOptions) (*v1alpha1.LogAnalyticsDatasourceWindowsEventList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error)
	LogAnalyticsDatasourceWindowsEventExpansion
}

// logAnalyticsDatasourceWindowsEvents implements LogAnalyticsDatasourceWindowsEventInterface
type logAnalyticsDatasourceWindowsEvents struct {
	client rest.Interface
	ns     string
}

// newLogAnalyticsDatasourceWindowsEvents returns a LogAnalyticsDatasourceWindowsEvents
func newLogAnalyticsDatasourceWindowsEvents(c *AzurermV1alpha1Client, namespace string) *logAnalyticsDatasourceWindowsEvents {
	return &logAnalyticsDatasourceWindowsEvents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the logAnalyticsDatasourceWindowsEvent, and returns the corresponding logAnalyticsDatasourceWindowsEvent object, and an error if there is any.
func (c *logAnalyticsDatasourceWindowsEvents) Get(name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEvent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LogAnalyticsDatasourceWindowsEvents that match those selectors.
func (c *logAnalyticsDatasourceWindowsEvents) List(opts v1.ListOptions) (result *v1alpha1.LogAnalyticsDatasourceWindowsEventList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEventList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logAnalyticsDatasourceWindowsEvents.
func (c *logAnalyticsDatasourceWindowsEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a logAnalyticsDatasourceWindowsEvent and creates it.  Returns the server's representation of the logAnalyticsDatasourceWindowsEvent, and an error, if there is any.
func (c *logAnalyticsDatasourceWindowsEvents) Create(logAnalyticsDatasourceWindowsEvent *v1alpha1.LogAnalyticsDatasourceWindowsEvent) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEvent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		Body(logAnalyticsDatasourceWindowsEvent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a logAnalyticsDatasourceWindowsEvent and updates it. Returns the server's representation of the logAnalyticsDatasourceWindowsEvent, and an error, if there is any.
func (c *logAnalyticsDatasourceWindowsEvents) Update(logAnalyticsDatasourceWindowsEvent *v1alpha1.LogAnalyticsDatasourceWindowsEvent) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEvent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		Name(logAnalyticsDatasourceWindowsEvent.Name).
		Body(logAnalyticsDatasourceWindowsEvent).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *logAnalyticsDatasourceWindowsEvents) UpdateStatus(logAnalyticsDatasourceWindowsEvent *v1alpha1.LogAnalyticsDatasourceWindowsEvent) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEvent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		Name(logAnalyticsDatasourceWindowsEvent.Name).
		SubResource("status").
		Body(logAnalyticsDatasourceWindowsEvent).
		Do().
		Into(result)
	return
}

// Delete takes name of the logAnalyticsDatasourceWindowsEvent and deletes it. Returns an error if one occurs.
func (c *logAnalyticsDatasourceWindowsEvents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *logAnalyticsDatasourceWindowsEvents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched logAnalyticsDatasourceWindowsEvent.
func (c *logAnalyticsDatasourceWindowsEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsDatasourceWindowsEvent, err error) {
	result = &v1alpha1.LogAnalyticsDatasourceWindowsEvent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loganalyticsdatasourcewindowsevents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}