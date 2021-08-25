/*
Copyright The KubeEdge Authors.

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

	v1alpha1 "github.com/kubeedge/sedna/pkg/apis/sedna/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeObjectTrackingServices implements ObjectTrackingServiceInterface
type FakeObjectTrackingServices struct {
	Fake *FakeSednaV1alpha1
	ns   string
}

var objecttrackingservicesResource = schema.GroupVersionResource{Group: "sedna.io", Version: "v1alpha1", Resource: "objecttrackingservices"}

var objecttrackingservicesKind = schema.GroupVersionKind{Group: "sedna.io", Version: "v1alpha1", Kind: "ObjectTrackingService"}

// Get takes name of the objectTrackingService, and returns the corresponding objectTrackingService object, and an error if there is any.
func (c *FakeObjectTrackingServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectTrackingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objecttrackingservicesResource, c.ns, name), &v1alpha1.ObjectTrackingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectTrackingService), err
}

// List takes label and field selectors, and returns the list of ObjectTrackingServices that match those selectors.
func (c *FakeObjectTrackingServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectTrackingServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objecttrackingservicesResource, objecttrackingservicesKind, c.ns, opts), &v1alpha1.ObjectTrackingServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectTrackingServiceList{ListMeta: obj.(*v1alpha1.ObjectTrackingServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectTrackingServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectTrackingServices.
func (c *FakeObjectTrackingServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objecttrackingservicesResource, c.ns, opts))

}

// Create takes the representation of a objectTrackingService and creates it.  Returns the server's representation of the objectTrackingService, and an error, if there is any.
func (c *FakeObjectTrackingServices) Create(ctx context.Context, objectTrackingService *v1alpha1.ObjectTrackingService, opts v1.CreateOptions) (result *v1alpha1.ObjectTrackingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objecttrackingservicesResource, c.ns, objectTrackingService), &v1alpha1.ObjectTrackingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectTrackingService), err
}

// Update takes the representation of a objectTrackingService and updates it. Returns the server's representation of the objectTrackingService, and an error, if there is any.
func (c *FakeObjectTrackingServices) Update(ctx context.Context, objectTrackingService *v1alpha1.ObjectTrackingService, opts v1.UpdateOptions) (result *v1alpha1.ObjectTrackingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objecttrackingservicesResource, c.ns, objectTrackingService), &v1alpha1.ObjectTrackingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectTrackingService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectTrackingServices) UpdateStatus(ctx context.Context, objectTrackingService *v1alpha1.ObjectTrackingService, opts v1.UpdateOptions) (*v1alpha1.ObjectTrackingService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objecttrackingservicesResource, "status", c.ns, objectTrackingService), &v1alpha1.ObjectTrackingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectTrackingService), err
}

// Delete takes name of the objectTrackingService and deletes it. Returns an error if one occurs.
func (c *FakeObjectTrackingServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objecttrackingservicesResource, c.ns, name), &v1alpha1.ObjectTrackingService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectTrackingServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objecttrackingservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectTrackingServiceList{})
	return err
}

// Patch applies the patch and returns the patched objectTrackingService.
func (c *FakeObjectTrackingServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectTrackingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objecttrackingservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ObjectTrackingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectTrackingService), err
}