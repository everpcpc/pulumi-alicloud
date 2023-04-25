// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host Host Group resource.
//
// For information about Bastion Host Host Group and how to use it, see [What is Host Group](https://www.alibabacloud.com/help/en/doc-detail/204307.htm).
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bastionhost.NewHostGroup(ctx, "example", &bastionhost.HostGroupArgs{
//				HostGroupName: pulumi.String("example_value"),
//				InstanceId:    pulumi.String("bastionhost-cn-tl3xxxxxxx"),
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
// Bastion Host Host Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:bastionhost/hostGroup:HostGroup example <instance_id>:<host_group_id>
//
// ```
type HostGroup struct {
	pulumi.CustomResourceState

	// Specify the New Host Group of Notes, Supports up to 500 Characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Host Group ID.
	HostGroupId pulumi.StringOutput `pulumi:"hostGroupId"`
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName pulumi.StringOutput `pulumi:"hostGroupName"`
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewHostGroup registers a new resource with the given unique name, arguments, and options.
func NewHostGroup(ctx *pulumi.Context,
	name string, args *HostGroupArgs, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupName == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource HostGroup
	err := ctx.RegisterResource("alicloud:bastionhost/hostGroup:HostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostGroup gets an existing HostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostGroupState, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	var resource HostGroup
	err := ctx.ReadResource("alicloud:bastionhost/hostGroup:HostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostGroup resources.
type hostGroupState struct {
	// Specify the New Host Group of Notes, Supports up to 500 Characters.
	Comment *string `pulumi:"comment"`
	// Host Group ID.
	HostGroupId *string `pulumi:"hostGroupId"`
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName *string `pulumi:"hostGroupName"`
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId *string `pulumi:"instanceId"`
}

type HostGroupState struct {
	// Specify the New Host Group of Notes, Supports up to 500 Characters.
	Comment pulumi.StringPtrInput
	// Host Group ID.
	HostGroupId pulumi.StringPtrInput
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName pulumi.StringPtrInput
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId pulumi.StringPtrInput
}

func (HostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupState)(nil)).Elem()
}

type hostGroupArgs struct {
	// Specify the New Host Group of Notes, Supports up to 500 Characters.
	Comment *string `pulumi:"comment"`
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName string `pulumi:"hostGroupName"`
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a HostGroup resource.
type HostGroupArgs struct {
	// Specify the New Host Group of Notes, Supports up to 500 Characters.
	Comment pulumi.StringPtrInput
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName pulumi.StringInput
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId pulumi.StringInput
}

func (HostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupArgs)(nil)).Elem()
}

type HostGroupInput interface {
	pulumi.Input

	ToHostGroupOutput() HostGroupOutput
	ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput
}

func (*HostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroup)(nil)).Elem()
}

func (i *HostGroup) ToHostGroupOutput() HostGroupOutput {
	return i.ToHostGroupOutputWithContext(context.Background())
}

func (i *HostGroup) ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupOutput)
}

// HostGroupArrayInput is an input type that accepts HostGroupArray and HostGroupArrayOutput values.
// You can construct a concrete instance of `HostGroupArrayInput` via:
//
//	HostGroupArray{ HostGroupArgs{...} }
type HostGroupArrayInput interface {
	pulumi.Input

	ToHostGroupArrayOutput() HostGroupArrayOutput
	ToHostGroupArrayOutputWithContext(context.Context) HostGroupArrayOutput
}

type HostGroupArray []HostGroupInput

func (HostGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroup)(nil)).Elem()
}

func (i HostGroupArray) ToHostGroupArrayOutput() HostGroupArrayOutput {
	return i.ToHostGroupArrayOutputWithContext(context.Background())
}

func (i HostGroupArray) ToHostGroupArrayOutputWithContext(ctx context.Context) HostGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupArrayOutput)
}

// HostGroupMapInput is an input type that accepts HostGroupMap and HostGroupMapOutput values.
// You can construct a concrete instance of `HostGroupMapInput` via:
//
//	HostGroupMap{ "key": HostGroupArgs{...} }
type HostGroupMapInput interface {
	pulumi.Input

	ToHostGroupMapOutput() HostGroupMapOutput
	ToHostGroupMapOutputWithContext(context.Context) HostGroupMapOutput
}

type HostGroupMap map[string]HostGroupInput

func (HostGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroup)(nil)).Elem()
}

func (i HostGroupMap) ToHostGroupMapOutput() HostGroupMapOutput {
	return i.ToHostGroupMapOutputWithContext(context.Background())
}

func (i HostGroupMap) ToHostGroupMapOutputWithContext(ctx context.Context) HostGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupMapOutput)
}

type HostGroupOutput struct{ *pulumi.OutputState }

func (HostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroup)(nil)).Elem()
}

func (o HostGroupOutput) ToHostGroupOutput() HostGroupOutput {
	return o
}

func (o HostGroupOutput) ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput {
	return o
}

// Specify the New Host Group of Notes, Supports up to 500 Characters.
func (o HostGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Host Group ID.
func (o HostGroupOutput) HostGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.HostGroupId }).(pulumi.StringOutput)
}

// Specify the New Host Group Name, Supports up to 128 Characters.
func (o HostGroupOutput) HostGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.HostGroupName }).(pulumi.StringOutput)
}

// Specify the New Host Group Where the Bastion Host ID of.
func (o HostGroupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type HostGroupArrayOutput struct{ *pulumi.OutputState }

func (HostGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroup)(nil)).Elem()
}

func (o HostGroupArrayOutput) ToHostGroupArrayOutput() HostGroupArrayOutput {
	return o
}

func (o HostGroupArrayOutput) ToHostGroupArrayOutputWithContext(ctx context.Context) HostGroupArrayOutput {
	return o
}

func (o HostGroupArrayOutput) Index(i pulumi.IntInput) HostGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostGroup {
		return vs[0].([]*HostGroup)[vs[1].(int)]
	}).(HostGroupOutput)
}

type HostGroupMapOutput struct{ *pulumi.OutputState }

func (HostGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroup)(nil)).Elem()
}

func (o HostGroupMapOutput) ToHostGroupMapOutput() HostGroupMapOutput {
	return o
}

func (o HostGroupMapOutput) ToHostGroupMapOutputWithContext(ctx context.Context) HostGroupMapOutput {
	return o
}

func (o HostGroupMapOutput) MapIndex(k pulumi.StringInput) HostGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostGroup {
		return vs[0].(map[string]*HostGroup)[vs[1].(string)]
	}).(HostGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupInput)(nil)).Elem(), &HostGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupArrayInput)(nil)).Elem(), HostGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupMapInput)(nil)).Elem(), HostGroupMap{})
	pulumi.RegisterOutputType(HostGroupOutput{})
	pulumi.RegisterOutputType(HostGroupArrayOutput{})
	pulumi.RegisterOutputType(HostGroupMapOutput{})
}
