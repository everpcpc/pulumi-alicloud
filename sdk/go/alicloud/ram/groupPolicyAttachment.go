// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM Group Policy attachment resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			group, err := ram.NewGroup(ctx, "group", &ram.GroupArgs{
//				Comments: pulumi.String("this is a group comments."),
//				Force:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			policy, err := ram.NewPolicy(ctx, "policy", &ram.PolicyArgs{
//				Document:    pulumi.String("    {\n      \"Statement\": [\n        {\n          \"Action\": [\n            \"oss:ListObjects\",\n            \"oss:GetObject\"\n          ],\n          \"Effect\": \"Allow\",\n          \"Resource\": [\n            \"acs:oss:*:*:mybucket\",\n            \"acs:oss:*:*:mybucket/*\"\n          ]\n        }\n      ],\n        \"Version\": \"1\"\n    }\n"),
//				Description: pulumi.String("this is a policy test"),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewGroupPolicyAttachment(ctx, "attach", &ram.GroupPolicyAttachmentArgs{
//				PolicyName: policy.Name,
//				PolicyType: policy.Type,
//				GroupName:  group.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// RAM Group Policy attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/groupPolicyAttachment:GroupPolicyAttachment example group:my-policy:Custom:my-group
//
// ```
type GroupPolicyAttachment struct {
	pulumi.CustomResourceState

	// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
}

// NewGroupPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicyAttachment(ctx *pulumi.Context,
	name string, args *GroupPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	var resource GroupPolicyAttachment
	err := ctx.RegisterResource("alicloud:ram/groupPolicyAttachment:GroupPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicyAttachment gets an existing GroupPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyAttachmentState, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	var resource GroupPolicyAttachment
	err := ctx.ReadResource("alicloud:ram/groupPolicyAttachment:GroupPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicyAttachment resources.
type groupPolicyAttachmentState struct {
	// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	GroupName *string `pulumi:"groupName"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName *string `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType *string `pulumi:"policyType"`
}

type GroupPolicyAttachmentState struct {
	// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	GroupName pulumi.StringPtrInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringPtrInput
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringPtrInput
}

func (GroupPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentState)(nil)).Elem()
}

type groupPolicyAttachmentArgs struct {
	// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	GroupName string `pulumi:"groupName"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName string `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType string `pulumi:"policyType"`
}

// The set of arguments for constructing a GroupPolicyAttachment resource.
type GroupPolicyAttachmentArgs struct {
	// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	GroupName pulumi.StringInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringInput
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringInput
}

func (GroupPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentArgs)(nil)).Elem()
}

type GroupPolicyAttachmentInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput
	ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput
}

func (*GroupPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicyAttachment)(nil)).Elem()
}

func (i *GroupPolicyAttachment) ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput {
	return i.ToGroupPolicyAttachmentOutputWithContext(context.Background())
}

func (i *GroupPolicyAttachment) ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentOutput)
}

// GroupPolicyAttachmentArrayInput is an input type that accepts GroupPolicyAttachmentArray and GroupPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `GroupPolicyAttachmentArrayInput` via:
//
//	GroupPolicyAttachmentArray{ GroupPolicyAttachmentArgs{...} }
type GroupPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput
	ToGroupPolicyAttachmentArrayOutputWithContext(context.Context) GroupPolicyAttachmentArrayOutput
}

type GroupPolicyAttachmentArray []GroupPolicyAttachmentInput

func (GroupPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicyAttachment)(nil)).Elem()
}

func (i GroupPolicyAttachmentArray) ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput {
	return i.ToGroupPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i GroupPolicyAttachmentArray) ToGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) GroupPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentArrayOutput)
}

// GroupPolicyAttachmentMapInput is an input type that accepts GroupPolicyAttachmentMap and GroupPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `GroupPolicyAttachmentMapInput` via:
//
//	GroupPolicyAttachmentMap{ "key": GroupPolicyAttachmentArgs{...} }
type GroupPolicyAttachmentMapInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput
	ToGroupPolicyAttachmentMapOutputWithContext(context.Context) GroupPolicyAttachmentMapOutput
}

type GroupPolicyAttachmentMap map[string]GroupPolicyAttachmentInput

func (GroupPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicyAttachment)(nil)).Elem()
}

func (i GroupPolicyAttachmentMap) ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput {
	return i.ToGroupPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i GroupPolicyAttachmentMap) ToGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) GroupPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentMapOutput)
}

type GroupPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentOutput) ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput {
	return o
}

func (o GroupPolicyAttachmentOutput) ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput {
	return o
}

// Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
func (o GroupPolicyAttachmentOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicyAttachment) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
func (o GroupPolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// Type of the RAM policy. It must be `Custom` or `System`.
func (o GroupPolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

type GroupPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentArrayOutput) ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput {
	return o
}

func (o GroupPolicyAttachmentArrayOutput) ToGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) GroupPolicyAttachmentArrayOutput {
	return o
}

func (o GroupPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) GroupPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupPolicyAttachment {
		return vs[0].([]*GroupPolicyAttachment)[vs[1].(int)]
	}).(GroupPolicyAttachmentOutput)
}

type GroupPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentMapOutput) ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput {
	return o
}

func (o GroupPolicyAttachmentMapOutput) ToGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) GroupPolicyAttachmentMapOutput {
	return o
}

func (o GroupPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) GroupPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupPolicyAttachment {
		return vs[0].(map[string]*GroupPolicyAttachment)[vs[1].(string)]
	}).(GroupPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentInput)(nil)).Elem(), &GroupPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentArrayInput)(nil)).Elem(), GroupPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentMapInput)(nil)).Elem(), GroupPolicyAttachmentMap{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentMapOutput{})
}
