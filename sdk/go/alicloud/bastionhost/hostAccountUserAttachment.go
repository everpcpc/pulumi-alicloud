// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host Host Account Attachment resource to add list host accounts into one user.
//
// > **NOTE:** Available in v1.135.0+.
//
// ## Import
//
// Bastion Host Host Account can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:bastionhost/hostAccountUserAttachment:HostAccountUserAttachment example <instance_id>:<user_id>:<host_id>
//
// ```
type HostAccountUserAttachment struct {
	pulumi.CustomResourceState

	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayOutput `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewHostAccountUserAttachment registers a new resource with the given unique name, arguments, and options.
func NewHostAccountUserAttachment(ctx *pulumi.Context,
	name string, args *HostAccountUserAttachmentArgs, opts ...pulumi.ResourceOption) (*HostAccountUserAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostAccountIds == nil {
		return nil, errors.New("invalid value for required argument 'HostAccountIds'")
	}
	if args.HostId == nil {
		return nil, errors.New("invalid value for required argument 'HostId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource HostAccountUserAttachment
	err := ctx.RegisterResource("alicloud:bastionhost/hostAccountUserAttachment:HostAccountUserAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostAccountUserAttachment gets an existing HostAccountUserAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostAccountUserAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostAccountUserAttachmentState, opts ...pulumi.ResourceOption) (*HostAccountUserAttachment, error) {
	var resource HostAccountUserAttachment
	err := ctx.ReadResource("alicloud:bastionhost/hostAccountUserAttachment:HostAccountUserAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostAccountUserAttachment resources.
type hostAccountUserAttachmentState struct {
	// A list IDs of the host account.
	HostAccountIds []string `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId *string `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	UserId *string `pulumi:"userId"`
}

type HostAccountUserAttachmentState struct {
	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayInput
	// The ID of the host.
	HostId pulumi.StringPtrInput
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringPtrInput
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	UserId pulumi.StringPtrInput
}

func (HostAccountUserAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAccountUserAttachmentState)(nil)).Elem()
}

type hostAccountUserAttachmentArgs struct {
	// A list IDs of the host account.
	HostAccountIds []string `pulumi:"hostAccountIds"`
	// The ID of the host.
	HostId string `pulumi:"hostId"`
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId string `pulumi:"instanceId"`
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a HostAccountUserAttachment resource.
type HostAccountUserAttachmentArgs struct {
	// A list IDs of the host account.
	HostAccountIds pulumi.StringArrayInput
	// The ID of the host.
	HostId pulumi.StringInput
	// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
	InstanceId pulumi.StringInput
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	UserId pulumi.StringInput
}

func (HostAccountUserAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAccountUserAttachmentArgs)(nil)).Elem()
}

type HostAccountUserAttachmentInput interface {
	pulumi.Input

	ToHostAccountUserAttachmentOutput() HostAccountUserAttachmentOutput
	ToHostAccountUserAttachmentOutputWithContext(ctx context.Context) HostAccountUserAttachmentOutput
}

func (*HostAccountUserAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAccountUserAttachment)(nil)).Elem()
}

func (i *HostAccountUserAttachment) ToHostAccountUserAttachmentOutput() HostAccountUserAttachmentOutput {
	return i.ToHostAccountUserAttachmentOutputWithContext(context.Background())
}

func (i *HostAccountUserAttachment) ToHostAccountUserAttachmentOutputWithContext(ctx context.Context) HostAccountUserAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserAttachmentOutput)
}

// HostAccountUserAttachmentArrayInput is an input type that accepts HostAccountUserAttachmentArray and HostAccountUserAttachmentArrayOutput values.
// You can construct a concrete instance of `HostAccountUserAttachmentArrayInput` via:
//
//	HostAccountUserAttachmentArray{ HostAccountUserAttachmentArgs{...} }
type HostAccountUserAttachmentArrayInput interface {
	pulumi.Input

	ToHostAccountUserAttachmentArrayOutput() HostAccountUserAttachmentArrayOutput
	ToHostAccountUserAttachmentArrayOutputWithContext(context.Context) HostAccountUserAttachmentArrayOutput
}

type HostAccountUserAttachmentArray []HostAccountUserAttachmentInput

func (HostAccountUserAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAccountUserAttachment)(nil)).Elem()
}

func (i HostAccountUserAttachmentArray) ToHostAccountUserAttachmentArrayOutput() HostAccountUserAttachmentArrayOutput {
	return i.ToHostAccountUserAttachmentArrayOutputWithContext(context.Background())
}

func (i HostAccountUserAttachmentArray) ToHostAccountUserAttachmentArrayOutputWithContext(ctx context.Context) HostAccountUserAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserAttachmentArrayOutput)
}

// HostAccountUserAttachmentMapInput is an input type that accepts HostAccountUserAttachmentMap and HostAccountUserAttachmentMapOutput values.
// You can construct a concrete instance of `HostAccountUserAttachmentMapInput` via:
//
//	HostAccountUserAttachmentMap{ "key": HostAccountUserAttachmentArgs{...} }
type HostAccountUserAttachmentMapInput interface {
	pulumi.Input

	ToHostAccountUserAttachmentMapOutput() HostAccountUserAttachmentMapOutput
	ToHostAccountUserAttachmentMapOutputWithContext(context.Context) HostAccountUserAttachmentMapOutput
}

type HostAccountUserAttachmentMap map[string]HostAccountUserAttachmentInput

func (HostAccountUserAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAccountUserAttachment)(nil)).Elem()
}

func (i HostAccountUserAttachmentMap) ToHostAccountUserAttachmentMapOutput() HostAccountUserAttachmentMapOutput {
	return i.ToHostAccountUserAttachmentMapOutputWithContext(context.Background())
}

func (i HostAccountUserAttachmentMap) ToHostAccountUserAttachmentMapOutputWithContext(ctx context.Context) HostAccountUserAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAccountUserAttachmentMapOutput)
}

type HostAccountUserAttachmentOutput struct{ *pulumi.OutputState }

func (HostAccountUserAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAccountUserAttachment)(nil)).Elem()
}

func (o HostAccountUserAttachmentOutput) ToHostAccountUserAttachmentOutput() HostAccountUserAttachmentOutput {
	return o
}

func (o HostAccountUserAttachmentOutput) ToHostAccountUserAttachmentOutputWithContext(ctx context.Context) HostAccountUserAttachmentOutput {
	return o
}

// A list IDs of the host account.
func (o HostAccountUserAttachmentOutput) HostAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostAccountUserAttachment) pulumi.StringArrayOutput { return v.HostAccountIds }).(pulumi.StringArrayOutput)
}

// The ID of the host.
func (o HostAccountUserAttachmentOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserAttachment) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

// The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
func (o HostAccountUserAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
func (o HostAccountUserAttachmentOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostAccountUserAttachment) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type HostAccountUserAttachmentArrayOutput struct{ *pulumi.OutputState }

func (HostAccountUserAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAccountUserAttachment)(nil)).Elem()
}

func (o HostAccountUserAttachmentArrayOutput) ToHostAccountUserAttachmentArrayOutput() HostAccountUserAttachmentArrayOutput {
	return o
}

func (o HostAccountUserAttachmentArrayOutput) ToHostAccountUserAttachmentArrayOutputWithContext(ctx context.Context) HostAccountUserAttachmentArrayOutput {
	return o
}

func (o HostAccountUserAttachmentArrayOutput) Index(i pulumi.IntInput) HostAccountUserAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostAccountUserAttachment {
		return vs[0].([]*HostAccountUserAttachment)[vs[1].(int)]
	}).(HostAccountUserAttachmentOutput)
}

type HostAccountUserAttachmentMapOutput struct{ *pulumi.OutputState }

func (HostAccountUserAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAccountUserAttachment)(nil)).Elem()
}

func (o HostAccountUserAttachmentMapOutput) ToHostAccountUserAttachmentMapOutput() HostAccountUserAttachmentMapOutput {
	return o
}

func (o HostAccountUserAttachmentMapOutput) ToHostAccountUserAttachmentMapOutputWithContext(ctx context.Context) HostAccountUserAttachmentMapOutput {
	return o
}

func (o HostAccountUserAttachmentMapOutput) MapIndex(k pulumi.StringInput) HostAccountUserAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostAccountUserAttachment {
		return vs[0].(map[string]*HostAccountUserAttachment)[vs[1].(string)]
	}).(HostAccountUserAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserAttachmentInput)(nil)).Elem(), &HostAccountUserAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserAttachmentArrayInput)(nil)).Elem(), HostAccountUserAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAccountUserAttachmentMapInput)(nil)).Elem(), HostAccountUserAttachmentMap{})
	pulumi.RegisterOutputType(HostAccountUserAttachmentOutput{})
	pulumi.RegisterOutputType(HostAccountUserAttachmentArrayOutput{})
	pulumi.RegisterOutputType(HostAccountUserAttachmentMapOutput{})
}
