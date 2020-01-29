// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ots

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource will help you to bind a VPC to an OTS instance.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ots_instance_attachment.html.markdown.
type InstanceAttachment struct {
	pulumi.CustomResourceState

	// The name of the OTS instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The name of attaching VPC to instance.
	VpcName pulumi.StringOutput `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewInstanceAttachment registers a new resource with the given unique name, arguments, and options.
func NewInstanceAttachment(ctx *pulumi.Context,
	name string, args *InstanceAttachmentArgs, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.VpcName == nil {
		return nil, errors.New("missing required argument 'VpcName'")
	}
	if args == nil || args.VswitchId == nil {
		return nil, errors.New("missing required argument 'VswitchId'")
	}
	if args == nil {
		args = &InstanceAttachmentArgs{}
	}
	var resource InstanceAttachment
	err := ctx.RegisterResource("alicloud:ots/instanceAttachment:InstanceAttachment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:ots/instanceAttachment:InstanceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceAttachment resources.
type instanceAttachmentState struct {
	// The name of the OTS instance.
	InstanceName *string `pulumi:"instanceName"`
	// The ID of attaching VPC to instance.
	VpcId *string `pulumi:"vpcId"`
	// The name of attaching VPC to instance.
	VpcName *string `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId *string `pulumi:"vswitchId"`
}

type InstanceAttachmentState struct {
	// The name of the OTS instance.
	InstanceName pulumi.StringPtrInput
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringPtrInput
	// The name of attaching VPC to instance.
	VpcName pulumi.StringPtrInput
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringPtrInput
}

func (InstanceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentState)(nil)).Elem()
}

type instanceAttachmentArgs struct {
	// The name of the OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of attaching VPC to instance.
	VpcName string `pulumi:"vpcName"`
	// The ID of attaching VSwitch to instance.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a InstanceAttachment resource.
type InstanceAttachmentArgs struct {
	// The name of the OTS instance.
	InstanceName pulumi.StringInput
	// The name of attaching VPC to instance.
	VpcName pulumi.StringInput
	// The ID of attaching VSwitch to instance.
	VswitchId pulumi.StringInput
}

func (InstanceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentArgs)(nil)).Elem()
}

