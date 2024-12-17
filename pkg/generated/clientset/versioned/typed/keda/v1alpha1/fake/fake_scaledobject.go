/*
Copyright 2024 The KEDA Authors

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

	v1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScaledObjects implements ScaledObjectInterface
type FakeScaledObjects struct {
	Fake *FakeKedaV1alpha1
	ns   string
}

var scaledobjectsResource = v1alpha1.SchemeGroupVersion.WithResource("scaledobjects")

var scaledobjectsKind = v1alpha1.SchemeGroupVersion.WithKind("ScaledObject")

// Get takes name of the scaledObject, and returns the corresponding scaledObject object, and an error if there is any.
func (c *FakeScaledObjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScaledObject, err error) {
	emptyResult := &v1alpha1.ScaledObject{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(scaledobjectsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScaledObject), err
}

// List takes label and field selectors, and returns the list of ScaledObjects that match those selectors.
func (c *FakeScaledObjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScaledObjectList, err error) {
	emptyResult := &v1alpha1.ScaledObjectList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(scaledobjectsResource, scaledobjectsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScaledObjectList{ListMeta: obj.(*v1alpha1.ScaledObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScaledObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scaledObjects.
func (c *FakeScaledObjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(scaledobjectsResource, c.ns, opts))

}

// Create takes the representation of a scaledObject and creates it.  Returns the server's representation of the scaledObject, and an error, if there is any.
func (c *FakeScaledObjects) Create(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.CreateOptions) (result *v1alpha1.ScaledObject, err error) {
	emptyResult := &v1alpha1.ScaledObject{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(scaledobjectsResource, c.ns, scaledObject, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScaledObject), err
}

// Update takes the representation of a scaledObject and updates it. Returns the server's representation of the scaledObject, and an error, if there is any.
func (c *FakeScaledObjects) Update(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.UpdateOptions) (result *v1alpha1.ScaledObject, err error) {
	emptyResult := &v1alpha1.ScaledObject{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(scaledobjectsResource, c.ns, scaledObject, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScaledObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScaledObjects) UpdateStatus(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.UpdateOptions) (result *v1alpha1.ScaledObject, err error) {
	emptyResult := &v1alpha1.ScaledObject{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(scaledobjectsResource, "status", c.ns, scaledObject, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScaledObject), err
}

// Delete takes name of the scaledObject and deletes it. Returns an error if one occurs.
func (c *FakeScaledObjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scaledobjectsResource, c.ns, name, opts), &v1alpha1.ScaledObject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScaledObjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(scaledobjectsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScaledObjectList{})
	return err
}

// Patch applies the patch and returns the patched scaledObject.
func (c *FakeScaledObjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScaledObject, err error) {
	emptyResult := &v1alpha1.ScaledObject{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(scaledobjectsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.ScaledObject), err
}
