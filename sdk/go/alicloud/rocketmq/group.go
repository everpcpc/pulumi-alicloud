// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rocketmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ONS group resource.
// 
// For more information about how to use it, see [RocketMQ Group Management API](https://www.alibabacloud.com/help/doc-detail/29616.html). 
// 
// > **NOTE:** Available in 1.53.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ons_group.html.markdown.
type Group struct {
	pulumi.CustomResourceState

	// Name of the group. Two groups on a single instance cannot have the same name. A `groupId` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrOutput `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("alicloud:rocketmq/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("alicloud:rocketmq/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupId` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupId *string `pulumi:"groupId"`
	// ID of the ONS Instance that owns the groups.
	InstanceId *string `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable *bool `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark *string `pulumi:"remark"`
}

type GroupState struct {
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupId` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupId pulumi.StringPtrInput
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringPtrInput
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrInput
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupId` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupId string `pulumi:"groupId"`
	// ID of the ONS Instance that owns the groups.
	InstanceId string `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable *bool `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupId` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupId pulumi.StringInput
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringInput
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrInput
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

