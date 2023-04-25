// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host Host Account Attachment resource to add list host accounts into one user group and one host group.
//
// > **NOTE:** Available in v1.135.0+.
//
// ## Import
//
// Bastion Host Host Account can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment example <instance_id>:<user_group_id>:<host_group_id>
//
// ```
type HostGroupAccountUserGroupAttachment struct {
	pulumi.CustomResourceState

	// A list names of the host account.
	HostAccountNames pulumi.StringArrayOutput `pulumi:"hostAccountNames"`
	// The ID of the host group.
	HostGroupId pulumi.StringOutput `pulumi:"hostGroupId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringOutput `pulumi:"userGroupId"`
}

// NewHostGroupAccountUserGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewHostGroupAccountUserGroupAttachment(ctx *pulumi.Context,
	name string, args *HostGroupAccountUserGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*HostGroupAccountUserGroupAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostAccountNames == nil {
		return nil, errors.New("invalid value for required argument 'HostAccountNames'")
	}
	if args.HostGroupId == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.UserGroupId == nil {
		return nil, errors.New("invalid value for required argument 'UserGroupId'")
	}
	var resource HostGroupAccountUserGroupAttachment
	err := ctx.RegisterResource("alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostGroupAccountUserGroupAttachment gets an existing HostGroupAccountUserGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostGroupAccountUserGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostGroupAccountUserGroupAttachmentState, opts ...pulumi.ResourceOption) (*HostGroupAccountUserGroupAttachment, error) {
	var resource HostGroupAccountUserGroupAttachment
	err := ctx.ReadResource("alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostGroupAccountUserGroupAttachment resources.
type hostGroupAccountUserGroupAttachmentState struct {
	// A list names of the host account.
	HostAccountNames []string `pulumi:"hostAccountNames"`
	// The ID of the host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId *string `pulumi:"userGroupId"`
}

type HostGroupAccountUserGroupAttachmentState struct {
	// A list names of the host account.
	HostAccountNames pulumi.StringArrayInput
	// The ID of the host group.
	HostGroupId pulumi.StringPtrInput
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringPtrInput
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringPtrInput
}

func (HostGroupAccountUserGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupAccountUserGroupAttachmentState)(nil)).Elem()
}

type hostGroupAccountUserGroupAttachmentArgs struct {
	// A list names of the host account.
	HostAccountNames []string `pulumi:"hostAccountNames"`
	// The ID of the host group.
	HostGroupId string `pulumi:"hostGroupId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId string `pulumi:"instanceId"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId string `pulumi:"userGroupId"`
}

// The set of arguments for constructing a HostGroupAccountUserGroupAttachment resource.
type HostGroupAccountUserGroupAttachmentArgs struct {
	// A list names of the host account.
	HostAccountNames pulumi.StringArrayInput
	// The ID of the host group.
	HostGroupId pulumi.StringInput
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringInput
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	UserGroupId pulumi.StringInput
}

func (HostGroupAccountUserGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupAccountUserGroupAttachmentArgs)(nil)).Elem()
}

type HostGroupAccountUserGroupAttachmentInput interface {
	pulumi.Input

	ToHostGroupAccountUserGroupAttachmentOutput() HostGroupAccountUserGroupAttachmentOutput
	ToHostGroupAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentOutput
}

func (*HostGroupAccountUserGroupAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (i *HostGroupAccountUserGroupAttachment) ToHostGroupAccountUserGroupAttachmentOutput() HostGroupAccountUserGroupAttachmentOutput {
	return i.ToHostGroupAccountUserGroupAttachmentOutputWithContext(context.Background())
}

func (i *HostGroupAccountUserGroupAttachment) ToHostGroupAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupAccountUserGroupAttachmentOutput)
}

// HostGroupAccountUserGroupAttachmentArrayInput is an input type that accepts HostGroupAccountUserGroupAttachmentArray and HostGroupAccountUserGroupAttachmentArrayOutput values.
// You can construct a concrete instance of `HostGroupAccountUserGroupAttachmentArrayInput` via:
//
//	HostGroupAccountUserGroupAttachmentArray{ HostGroupAccountUserGroupAttachmentArgs{...} }
type HostGroupAccountUserGroupAttachmentArrayInput interface {
	pulumi.Input

	ToHostGroupAccountUserGroupAttachmentArrayOutput() HostGroupAccountUserGroupAttachmentArrayOutput
	ToHostGroupAccountUserGroupAttachmentArrayOutputWithContext(context.Context) HostGroupAccountUserGroupAttachmentArrayOutput
}

type HostGroupAccountUserGroupAttachmentArray []HostGroupAccountUserGroupAttachmentInput

func (HostGroupAccountUserGroupAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (i HostGroupAccountUserGroupAttachmentArray) ToHostGroupAccountUserGroupAttachmentArrayOutput() HostGroupAccountUserGroupAttachmentArrayOutput {
	return i.ToHostGroupAccountUserGroupAttachmentArrayOutputWithContext(context.Background())
}

func (i HostGroupAccountUserGroupAttachmentArray) ToHostGroupAccountUserGroupAttachmentArrayOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupAccountUserGroupAttachmentArrayOutput)
}

// HostGroupAccountUserGroupAttachmentMapInput is an input type that accepts HostGroupAccountUserGroupAttachmentMap and HostGroupAccountUserGroupAttachmentMapOutput values.
// You can construct a concrete instance of `HostGroupAccountUserGroupAttachmentMapInput` via:
//
//	HostGroupAccountUserGroupAttachmentMap{ "key": HostGroupAccountUserGroupAttachmentArgs{...} }
type HostGroupAccountUserGroupAttachmentMapInput interface {
	pulumi.Input

	ToHostGroupAccountUserGroupAttachmentMapOutput() HostGroupAccountUserGroupAttachmentMapOutput
	ToHostGroupAccountUserGroupAttachmentMapOutputWithContext(context.Context) HostGroupAccountUserGroupAttachmentMapOutput
}

type HostGroupAccountUserGroupAttachmentMap map[string]HostGroupAccountUserGroupAttachmentInput

func (HostGroupAccountUserGroupAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (i HostGroupAccountUserGroupAttachmentMap) ToHostGroupAccountUserGroupAttachmentMapOutput() HostGroupAccountUserGroupAttachmentMapOutput {
	return i.ToHostGroupAccountUserGroupAttachmentMapOutputWithContext(context.Background())
}

func (i HostGroupAccountUserGroupAttachmentMap) ToHostGroupAccountUserGroupAttachmentMapOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupAccountUserGroupAttachmentMapOutput)
}

type HostGroupAccountUserGroupAttachmentOutput struct{ *pulumi.OutputState }

func (HostGroupAccountUserGroupAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostGroupAccountUserGroupAttachmentOutput) ToHostGroupAccountUserGroupAttachmentOutput() HostGroupAccountUserGroupAttachmentOutput {
	return o
}

func (o HostGroupAccountUserGroupAttachmentOutput) ToHostGroupAccountUserGroupAttachmentOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentOutput {
	return o
}

// A list names of the host account.
func (o HostGroupAccountUserGroupAttachmentOutput) HostAccountNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostGroupAccountUserGroupAttachment) pulumi.StringArrayOutput { return v.HostAccountNames }).(pulumi.StringArrayOutput)
}

// The ID of the host group.
func (o HostGroupAccountUserGroupAttachmentOutput) HostGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroupAccountUserGroupAttachment) pulumi.StringOutput { return v.HostGroupId }).(pulumi.StringOutput)
}

// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
func (o HostGroupAccountUserGroupAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroupAccountUserGroupAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
func (o HostGroupAccountUserGroupAttachmentOutput) UserGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroupAccountUserGroupAttachment) pulumi.StringOutput { return v.UserGroupId }).(pulumi.StringOutput)
}

type HostGroupAccountUserGroupAttachmentArrayOutput struct{ *pulumi.OutputState }

func (HostGroupAccountUserGroupAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostGroupAccountUserGroupAttachmentArrayOutput) ToHostGroupAccountUserGroupAttachmentArrayOutput() HostGroupAccountUserGroupAttachmentArrayOutput {
	return o
}

func (o HostGroupAccountUserGroupAttachmentArrayOutput) ToHostGroupAccountUserGroupAttachmentArrayOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentArrayOutput {
	return o
}

func (o HostGroupAccountUserGroupAttachmentArrayOutput) Index(i pulumi.IntInput) HostGroupAccountUserGroupAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostGroupAccountUserGroupAttachment {
		return vs[0].([]*HostGroupAccountUserGroupAttachment)[vs[1].(int)]
	}).(HostGroupAccountUserGroupAttachmentOutput)
}

type HostGroupAccountUserGroupAttachmentMapOutput struct{ *pulumi.OutputState }

func (HostGroupAccountUserGroupAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroupAccountUserGroupAttachment)(nil)).Elem()
}

func (o HostGroupAccountUserGroupAttachmentMapOutput) ToHostGroupAccountUserGroupAttachmentMapOutput() HostGroupAccountUserGroupAttachmentMapOutput {
	return o
}

func (o HostGroupAccountUserGroupAttachmentMapOutput) ToHostGroupAccountUserGroupAttachmentMapOutputWithContext(ctx context.Context) HostGroupAccountUserGroupAttachmentMapOutput {
	return o
}

func (o HostGroupAccountUserGroupAttachmentMapOutput) MapIndex(k pulumi.StringInput) HostGroupAccountUserGroupAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostGroupAccountUserGroupAttachment {
		return vs[0].(map[string]*HostGroupAccountUserGroupAttachment)[vs[1].(string)]
	}).(HostGroupAccountUserGroupAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupAccountUserGroupAttachmentInput)(nil)).Elem(), &HostGroupAccountUserGroupAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupAccountUserGroupAttachmentArrayInput)(nil)).Elem(), HostGroupAccountUserGroupAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupAccountUserGroupAttachmentMapInput)(nil)).Elem(), HostGroupAccountUserGroupAttachmentMap{})
	pulumi.RegisterOutputType(HostGroupAccountUserGroupAttachmentOutput{})
	pulumi.RegisterOutputType(HostGroupAccountUserGroupAttachmentArrayOutput{})
	pulumi.RegisterOutputType(HostGroupAccountUserGroupAttachmentMapOutput{})
}
