/*
Copyright 2018 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshotSchedules implements VolumeSnapshotScheduleInterface
type FakeVolumeSnapshotSchedules struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var volumesnapshotschedulesResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "volumesnapshotschedules"}

var volumesnapshotschedulesKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "VolumeSnapshotSchedule"}

// Get takes name of the volumeSnapshotSchedule, and returns the corresponding volumeSnapshotSchedule object, and an error if there is any.
func (c *FakeVolumeSnapshotSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.VolumeSnapshotSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotschedulesResource, c.ns, name), &v1alpha1.VolumeSnapshotSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotSchedule), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotSchedules that match those selectors.
func (c *FakeVolumeSnapshotSchedules) List(opts v1.ListOptions) (result *v1alpha1.VolumeSnapshotScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotschedulesResource, volumesnapshotschedulesKind, c.ns, opts), &v1alpha1.VolumeSnapshotScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeSnapshotScheduleList{ListMeta: obj.(*v1alpha1.VolumeSnapshotScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.VolumeSnapshotScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotSchedules.
func (c *FakeVolumeSnapshotSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotschedulesResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshotSchedule and creates it.  Returns the server's representation of the volumeSnapshotSchedule, and an error, if there is any.
func (c *FakeVolumeSnapshotSchedules) Create(volumeSnapshotSchedule *v1alpha1.VolumeSnapshotSchedule) (result *v1alpha1.VolumeSnapshotSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotschedulesResource, c.ns, volumeSnapshotSchedule), &v1alpha1.VolumeSnapshotSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotSchedule), err
}

// Update takes the representation of a volumeSnapshotSchedule and updates it. Returns the server's representation of the volumeSnapshotSchedule, and an error, if there is any.
func (c *FakeVolumeSnapshotSchedules) Update(volumeSnapshotSchedule *v1alpha1.VolumeSnapshotSchedule) (result *v1alpha1.VolumeSnapshotSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotschedulesResource, c.ns, volumeSnapshotSchedule), &v1alpha1.VolumeSnapshotSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshotSchedules) UpdateStatus(volumeSnapshotSchedule *v1alpha1.VolumeSnapshotSchedule) (*v1alpha1.VolumeSnapshotSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotschedulesResource, "status", c.ns, volumeSnapshotSchedule), &v1alpha1.VolumeSnapshotSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotSchedule), err
}

// Delete takes name of the volumeSnapshotSchedule and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotSchedules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesnapshotschedulesResource, c.ns, name), &v1alpha1.VolumeSnapshotSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotschedulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeSnapshotScheduleList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotSchedule.
func (c *FakeVolumeSnapshotSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeSnapshotSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotschedulesResource, c.ns, name, data, subresources...), &v1alpha1.VolumeSnapshotSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotSchedule), err
}