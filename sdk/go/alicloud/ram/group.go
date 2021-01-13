// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// RAM group can be imported using the id or name, e.g.
//
// ```sh
//  $ pulumi import alicloud:ram/group:Group example my-group
// ```
type Group struct {
	pulumi.CustomResourceState

	// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("alicloud:ram/group:Group", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:ram/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
	Comments *string `pulumi:"comments"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
}

type GroupState struct {
	// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
	Comments *string `pulumi:"comments"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Comment of the RAM group. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil)).Elem()
}

func (i Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupOutput)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
