// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openshift/aws-ebs-csi-driver-operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEBSCSIDrivers implements EBSCSIDriverInterface
type FakeEBSCSIDrivers struct {
	Fake *FakeCsiV1alpha1
}

var ebscsidriversResource = schema.GroupVersionResource{Group: "csi.ebs.aws.com", Version: "v1alpha1", Resource: "ebscsidrivers"}

var ebscsidriversKind = schema.GroupVersionKind{Group: "csi.ebs.aws.com", Version: "v1alpha1", Kind: "EBSCSIDriver"}

// Get takes name of the eBSCSIDriver, and returns the corresponding eBSCSIDriver object, and an error if there is any.
func (c *FakeEBSCSIDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ebscsidriversResource, name), &v1alpha1.EBSCSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EBSCSIDriver), err
}

// List takes label and field selectors, and returns the list of EBSCSIDrivers that match those selectors.
func (c *FakeEBSCSIDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EBSCSIDriverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ebscsidriversResource, ebscsidriversKind, opts), &v1alpha1.EBSCSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EBSCSIDriverList{ListMeta: obj.(*v1alpha1.EBSCSIDriverList).ListMeta}
	for _, item := range obj.(*v1alpha1.EBSCSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eBSCSIDrivers.
func (c *FakeEBSCSIDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ebscsidriversResource, opts))
}

// Create takes the representation of a eBSCSIDriver and creates it.  Returns the server's representation of the eBSCSIDriver, and an error, if there is any.
func (c *FakeEBSCSIDrivers) Create(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.CreateOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ebscsidriversResource, eBSCSIDriver), &v1alpha1.EBSCSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EBSCSIDriver), err
}

// Update takes the representation of a eBSCSIDriver and updates it. Returns the server's representation of the eBSCSIDriver, and an error, if there is any.
func (c *FakeEBSCSIDrivers) Update(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ebscsidriversResource, eBSCSIDriver), &v1alpha1.EBSCSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EBSCSIDriver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEBSCSIDrivers) UpdateStatus(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (*v1alpha1.EBSCSIDriver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ebscsidriversResource, "status", eBSCSIDriver), &v1alpha1.EBSCSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EBSCSIDriver), err
}

// Delete takes name of the eBSCSIDriver and deletes it. Returns an error if one occurs.
func (c *FakeEBSCSIDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ebscsidriversResource, name), &v1alpha1.EBSCSIDriver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEBSCSIDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ebscsidriversResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EBSCSIDriverList{})
	return err
}

// Patch applies the patch and returns the patched eBSCSIDriver.
func (c *FakeEBSCSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EBSCSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ebscsidriversResource, name, pt, data, subresources...), &v1alpha1.EBSCSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EBSCSIDriver), err
}