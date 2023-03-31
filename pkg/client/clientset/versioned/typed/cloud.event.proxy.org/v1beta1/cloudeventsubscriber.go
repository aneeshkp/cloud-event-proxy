/*
Copyright 2023 The Kubernetes Authors

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/redhat-cne/cloud-event-proxy/pkg/api/cloud.event.proxy.org/v1beta1"
	scheme "github.com/redhat-cne/cloud-event-proxy/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudEventSubscribersGetter has a method to return a CloudEventSubscriberInterface.
// A group's client should implement this interface.
type CloudEventSubscribersGetter interface {
	CloudEventSubscribers(namespace string) CloudEventSubscriberInterface
}

// CloudEventSubscriberInterface has methods to work with CloudEventSubscriber resources.
type CloudEventSubscriberInterface interface {
	Create(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.CreateOptions) (*v1beta1.CloudEventSubscriber, error)
	Update(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (*v1beta1.CloudEventSubscriber, error)
	UpdateStatus(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (*v1beta1.CloudEventSubscriber, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CloudEventSubscriber, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CloudEventSubscriberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudEventSubscriber, err error)
	CloudEventSubscriberExpansion
}

// cloudEventSubscribers implements CloudEventSubscriberInterface
type cloudEventSubscribers struct {
	client rest.Interface
	ns     string
}

// newCloudEventSubscribers returns a CloudEventSubscribers
func newCloudEventSubscribers(c *CloudV1beta1Client, namespace string) *cloudEventSubscribers {
	return &cloudEventSubscribers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudEventSubscriber, and returns the corresponding cloudEventSubscriber object, and an error if there is any.
func (c *cloudEventSubscribers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	result = &v1beta1.CloudEventSubscriber{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudEventSubscribers that match those selectors.
func (c *cloudEventSubscribers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudEventSubscriberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CloudEventSubscriberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudEventSubscribers.
func (c *cloudEventSubscribers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudEventSubscriber and creates it.  Returns the server's representation of the cloudEventSubscriber, and an error, if there is any.
func (c *cloudEventSubscribers) Create(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.CreateOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	result = &v1beta1.CloudEventSubscriber{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudEventSubscriber).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudEventSubscriber and updates it. Returns the server's representation of the cloudEventSubscriber, and an error, if there is any.
func (c *cloudEventSubscribers) Update(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	result = &v1beta1.CloudEventSubscriber{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		Name(cloudEventSubscriber.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudEventSubscriber).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudEventSubscribers) UpdateStatus(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	result = &v1beta1.CloudEventSubscriber{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		Name(cloudEventSubscriber.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudEventSubscriber).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudEventSubscriber and deletes it. Returns an error if one occurs.
func (c *cloudEventSubscribers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudEventSubscribers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudEventSubscriber.
func (c *cloudEventSubscribers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudEventSubscriber, err error) {
	result = &v1beta1.CloudEventSubscriber{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudeventsubscribers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
