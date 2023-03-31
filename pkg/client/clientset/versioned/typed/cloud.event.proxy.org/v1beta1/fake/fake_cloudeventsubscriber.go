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

package fake

import (
	"context"

	v1beta1 "github.com/redhat-cne/cloud-event-proxy/pkg/api/cloud.event.proxy.org/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudEventSubscribers implements CloudEventSubscriberInterface
type FakeCloudEventSubscribers struct {
	Fake *FakeCloudV1beta1
	ns   string
}

var cloudeventsubscribersResource = schema.GroupVersionResource{Group: "cloud.event.proxy.org", Version: "v1beta1", Resource: "cloudeventsubscribers"}

var cloudeventsubscribersKind = schema.GroupVersionKind{Group: "cloud.event.proxy.org", Version: "v1beta1", Kind: "CloudEventSubscriber"}

// Get takes name of the cloudEventSubscriber, and returns the corresponding cloudEventSubscriber object, and an error if there is any.
func (c *FakeCloudEventSubscribers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudeventsubscribersResource, c.ns, name), &v1beta1.CloudEventSubscriber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudEventSubscriber), err
}

// List takes label and field selectors, and returns the list of CloudEventSubscribers that match those selectors.
func (c *FakeCloudEventSubscribers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudEventSubscriberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudeventsubscribersResource, cloudeventsubscribersKind, c.ns, opts), &v1beta1.CloudEventSubscriberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudEventSubscriberList{ListMeta: obj.(*v1beta1.CloudEventSubscriberList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudEventSubscriberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudEventSubscribers.
func (c *FakeCloudEventSubscribers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudeventsubscribersResource, c.ns, opts))

}

// Create takes the representation of a cloudEventSubscriber and creates it.  Returns the server's representation of the cloudEventSubscriber, and an error, if there is any.
func (c *FakeCloudEventSubscribers) Create(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.CreateOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudeventsubscribersResource, c.ns, cloudEventSubscriber), &v1beta1.CloudEventSubscriber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudEventSubscriber), err
}

// Update takes the representation of a cloudEventSubscriber and updates it. Returns the server's representation of the cloudEventSubscriber, and an error, if there is any.
func (c *FakeCloudEventSubscribers) Update(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (result *v1beta1.CloudEventSubscriber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudeventsubscribersResource, c.ns, cloudEventSubscriber), &v1beta1.CloudEventSubscriber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudEventSubscriber), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudEventSubscribers) UpdateStatus(ctx context.Context, cloudEventSubscriber *v1beta1.CloudEventSubscriber, opts v1.UpdateOptions) (*v1beta1.CloudEventSubscriber, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudeventsubscribersResource, "status", c.ns, cloudEventSubscriber), &v1beta1.CloudEventSubscriber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudEventSubscriber), err
}

// Delete takes name of the cloudEventSubscriber and deletes it. Returns an error if one occurs.
func (c *FakeCloudEventSubscribers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudeventsubscribersResource, c.ns, name, opts), &v1beta1.CloudEventSubscriber{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudEventSubscribers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudeventsubscribersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudEventSubscriberList{})
	return err
}

// Patch applies the patch and returns the patched cloudEventSubscriber.
func (c *FakeCloudEventSubscribers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudEventSubscriber, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudeventsubscribersResource, c.ns, name, pt, data, subresources...), &v1beta1.CloudEventSubscriber{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudEventSubscriber), err
}
