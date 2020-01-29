// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/disk_attachment.html.markdown.
type DiskAttachment struct {
	pulumi.CustomResourceState

	// The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// ID of the Disk to be attached.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDiskAttachment registers a new resource with the given unique name, arguments, and options.
func NewDiskAttachment(ctx *pulumi.Context,
	name string, args *DiskAttachmentArgs, opts ...pulumi.ResourceOption) (*DiskAttachment, error) {
	if args == nil || args.DiskId == nil {
		return nil, errors.New("missing required argument 'DiskId'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &DiskAttachmentArgs{}
	}
	var resource DiskAttachment
	err := ctx.RegisterResource("alicloud:ecs/diskAttachment:DiskAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAttachment gets an existing DiskAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAttachmentState, opts ...pulumi.ResourceOption) (*DiskAttachment, error) {
	var resource DiskAttachment
	err := ctx.ReadResource("alicloud:ecs/diskAttachment:DiskAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAttachment resources.
type diskAttachmentState struct {
	// The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
	DeviceName *string `pulumi:"deviceName"`
	// ID of the Disk to be attached.
	DiskId *string `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId *string `pulumi:"instanceId"`
}

type DiskAttachmentState struct {
	// The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
	DeviceName pulumi.StringPtrInput
	// ID of the Disk to be attached.
	DiskId pulumi.StringPtrInput
	// ID of the Instance to attach to.
	InstanceId pulumi.StringPtrInput
}

func (DiskAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAttachmentState)(nil)).Elem()
}

type diskAttachmentArgs struct {
	// The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
	DeviceName *string `pulumi:"deviceName"`
	// ID of the Disk to be attached.
	DiskId string `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a DiskAttachment resource.
type DiskAttachmentArgs struct {
	// The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
	DeviceName pulumi.StringPtrInput
	// ID of the Disk to be attached.
	DiskId pulumi.StringInput
	// ID of the Instance to attach to.
	InstanceId pulumi.StringInput
}

func (DiskAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAttachmentArgs)(nil)).Elem()
}

