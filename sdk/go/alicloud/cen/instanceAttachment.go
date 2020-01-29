// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cen

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CEN child instance attachment resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_instance_attachment.html.markdown.
type InstanceAttachment struct {
	pulumi.CustomResourceState

	// The ID of the child instance to attach.
	ChildInstanceId pulumi.StringOutput `pulumi:"childInstanceId"`
	// The uid of the child instance. Only used when attach a child instance of other account.
	ChildInstanceOwnerId pulumi.StringOutput `pulumi:"childInstanceOwnerId"`
	// The region ID of the child instance to attach.
	ChildInstanceRegionId pulumi.StringOutput `pulumi:"childInstanceRegionId"`
	// The ID of the CEN.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewInstanceAttachment registers a new resource with the given unique name, arguments, and options.
func NewInstanceAttachment(ctx *pulumi.Context,
	name string, args *InstanceAttachmentArgs, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	if args == nil || args.ChildInstanceId == nil {
		return nil, errors.New("missing required argument 'ChildInstanceId'")
	}
	if args == nil || args.ChildInstanceRegionId == nil {
		return nil, errors.New("missing required argument 'ChildInstanceRegionId'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &InstanceAttachmentArgs{}
	}
	var resource InstanceAttachment
	err := ctx.RegisterResource("alicloud:cen/instanceAttachment:InstanceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceAttachment gets an existing InstanceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceAttachmentState, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	var resource InstanceAttachment
	err := ctx.ReadResource("alicloud:cen/instanceAttachment:InstanceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceAttachment resources.
type instanceAttachmentState struct {
	// The ID of the child instance to attach.
	ChildInstanceId *string `pulumi:"childInstanceId"`
	// The uid of the child instance. Only used when attach a child instance of other account.
	ChildInstanceOwnerId *string `pulumi:"childInstanceOwnerId"`
	// The region ID of the child instance to attach.
	ChildInstanceRegionId *string `pulumi:"childInstanceRegionId"`
	// The ID of the CEN.
	InstanceId *string `pulumi:"instanceId"`
}

type InstanceAttachmentState struct {
	// The ID of the child instance to attach.
	ChildInstanceId pulumi.StringPtrInput
	// The uid of the child instance. Only used when attach a child instance of other account.
	ChildInstanceOwnerId pulumi.StringPtrInput
	// The region ID of the child instance to attach.
	ChildInstanceRegionId pulumi.StringPtrInput
	// The ID of the CEN.
	InstanceId pulumi.StringPtrInput
}

func (InstanceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentState)(nil)).Elem()
}

type instanceAttachmentArgs struct {
	// The ID of the child instance to attach.
	ChildInstanceId string `pulumi:"childInstanceId"`
	// The uid of the child instance. Only used when attach a child instance of other account.
	ChildInstanceOwnerId *string `pulumi:"childInstanceOwnerId"`
	// The region ID of the child instance to attach.
	ChildInstanceRegionId string `pulumi:"childInstanceRegionId"`
	// The ID of the CEN.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a InstanceAttachment resource.
type InstanceAttachmentArgs struct {
	// The ID of the child instance to attach.
	ChildInstanceId pulumi.StringInput
	// The uid of the child instance. Only used when attach a child instance of other account.
	ChildInstanceOwnerId pulumi.StringPtrInput
	// The region ID of the child instance to attach.
	ChildInstanceRegionId pulumi.StringInput
	// The ID of the CEN.
	InstanceId pulumi.StringInput
}

func (InstanceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentArgs)(nil)).Elem()
}

