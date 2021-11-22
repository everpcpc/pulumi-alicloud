// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Firewall Control Policy resource.
//
// For information about Cloud Firewall Control Policy Order and how to use it, see [What is Control Policy Order](https://www.alibabacloud.com/help/doc-detail/138867.htm).
//
// > **NOTE:** Available in v1.130.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example1, err := cloudfirewall.NewControlPolicy(ctx, "example1", &cloudfirewall.ControlPolicyArgs{
// 			ApplicationName: pulumi.String("ANY"),
// 			AclAction:       pulumi.String("accept"),
// 			Description:     pulumi.String("example"),
// 			DestinationType: pulumi.String("net"),
// 			Destination:     pulumi.String("100.1.1.0/24"),
// 			Direction:       pulumi.String("out"),
// 			Proto:           pulumi.String("ANY"),
// 			Source:          pulumi.String("1.2.3.0/24"),
// 			SourceType:      pulumi.String("net"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudfirewall.NewControlPolicyOrder(ctx, "example2", &cloudfirewall.ControlPolicyOrderArgs{
// 			AclUuid:   example1.AclUuid,
// 			Direction: example1.Direction,
// 			Order:     pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Cloud Firewall Control Policy Order can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cloudfirewall/controlPolicyOrder:ControlPolicyOrder example <acl_uuid>:<direction>
// ```
type ControlPolicyOrder struct {
	pulumi.CustomResourceState

	// The unique ID of the access control policy.
	AclUuid pulumi.StringOutput `pulumi:"aclUuid"`
	// Direction. Valid values: `in`, `out`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of -1 indicates the lowest priority.
	Order pulumi.IntPtrOutput `pulumi:"order"`
}

// NewControlPolicyOrder registers a new resource with the given unique name, arguments, and options.
func NewControlPolicyOrder(ctx *pulumi.Context,
	name string, args *ControlPolicyOrderArgs, opts ...pulumi.ResourceOption) (*ControlPolicyOrder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclUuid == nil {
		return nil, errors.New("invalid value for required argument 'AclUuid'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	var resource ControlPolicyOrder
	err := ctx.RegisterResource("alicloud:cloudfirewall/controlPolicyOrder:ControlPolicyOrder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlPolicyOrder gets an existing ControlPolicyOrder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlPolicyOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlPolicyOrderState, opts ...pulumi.ResourceOption) (*ControlPolicyOrder, error) {
	var resource ControlPolicyOrder
	err := ctx.ReadResource("alicloud:cloudfirewall/controlPolicyOrder:ControlPolicyOrder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlPolicyOrder resources.
type controlPolicyOrderState struct {
	// The unique ID of the access control policy.
	AclUuid *string `pulumi:"aclUuid"`
	// Direction. Valid values: `in`, `out`.
	Direction *string `pulumi:"direction"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of -1 indicates the lowest priority.
	Order *int `pulumi:"order"`
}

type ControlPolicyOrderState struct {
	// The unique ID of the access control policy.
	AclUuid pulumi.StringPtrInput
	// Direction. Valid values: `in`, `out`.
	Direction pulumi.StringPtrInput
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of -1 indicates the lowest priority.
	Order pulumi.IntPtrInput
}

func (ControlPolicyOrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyOrderState)(nil)).Elem()
}

type controlPolicyOrderArgs struct {
	// The unique ID of the access control policy.
	AclUuid string `pulumi:"aclUuid"`
	// Direction. Valid values: `in`, `out`.
	Direction string `pulumi:"direction"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of -1 indicates the lowest priority.
	Order *int `pulumi:"order"`
}

// The set of arguments for constructing a ControlPolicyOrder resource.
type ControlPolicyOrderArgs struct {
	// The unique ID of the access control policy.
	AclUuid pulumi.StringInput
	// Direction. Valid values: `in`, `out`.
	Direction pulumi.StringInput
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority. **NOTE:** The value of -1 indicates the lowest priority.
	Order pulumi.IntPtrInput
}

func (ControlPolicyOrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyOrderArgs)(nil)).Elem()
}

type ControlPolicyOrderInput interface {
	pulumi.Input

	ToControlPolicyOrderOutput() ControlPolicyOrderOutput
	ToControlPolicyOrderOutputWithContext(ctx context.Context) ControlPolicyOrderOutput
}

func (*ControlPolicyOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPolicyOrder)(nil))
}

func (i *ControlPolicyOrder) ToControlPolicyOrderOutput() ControlPolicyOrderOutput {
	return i.ToControlPolicyOrderOutputWithContext(context.Background())
}

func (i *ControlPolicyOrder) ToControlPolicyOrderOutputWithContext(ctx context.Context) ControlPolicyOrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOrderOutput)
}

func (i *ControlPolicyOrder) ToControlPolicyOrderPtrOutput() ControlPolicyOrderPtrOutput {
	return i.ToControlPolicyOrderPtrOutputWithContext(context.Background())
}

func (i *ControlPolicyOrder) ToControlPolicyOrderPtrOutputWithContext(ctx context.Context) ControlPolicyOrderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOrderPtrOutput)
}

type ControlPolicyOrderPtrInput interface {
	pulumi.Input

	ToControlPolicyOrderPtrOutput() ControlPolicyOrderPtrOutput
	ToControlPolicyOrderPtrOutputWithContext(ctx context.Context) ControlPolicyOrderPtrOutput
}

type controlPolicyOrderPtrType ControlPolicyOrderArgs

func (*controlPolicyOrderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicyOrder)(nil))
}

func (i *controlPolicyOrderPtrType) ToControlPolicyOrderPtrOutput() ControlPolicyOrderPtrOutput {
	return i.ToControlPolicyOrderPtrOutputWithContext(context.Background())
}

func (i *controlPolicyOrderPtrType) ToControlPolicyOrderPtrOutputWithContext(ctx context.Context) ControlPolicyOrderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOrderPtrOutput)
}

// ControlPolicyOrderArrayInput is an input type that accepts ControlPolicyOrderArray and ControlPolicyOrderArrayOutput values.
// You can construct a concrete instance of `ControlPolicyOrderArrayInput` via:
//
//          ControlPolicyOrderArray{ ControlPolicyOrderArgs{...} }
type ControlPolicyOrderArrayInput interface {
	pulumi.Input

	ToControlPolicyOrderArrayOutput() ControlPolicyOrderArrayOutput
	ToControlPolicyOrderArrayOutputWithContext(context.Context) ControlPolicyOrderArrayOutput
}

type ControlPolicyOrderArray []ControlPolicyOrderInput

func (ControlPolicyOrderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPolicyOrder)(nil)).Elem()
}

func (i ControlPolicyOrderArray) ToControlPolicyOrderArrayOutput() ControlPolicyOrderArrayOutput {
	return i.ToControlPolicyOrderArrayOutputWithContext(context.Background())
}

func (i ControlPolicyOrderArray) ToControlPolicyOrderArrayOutputWithContext(ctx context.Context) ControlPolicyOrderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOrderArrayOutput)
}

// ControlPolicyOrderMapInput is an input type that accepts ControlPolicyOrderMap and ControlPolicyOrderMapOutput values.
// You can construct a concrete instance of `ControlPolicyOrderMapInput` via:
//
//          ControlPolicyOrderMap{ "key": ControlPolicyOrderArgs{...} }
type ControlPolicyOrderMapInput interface {
	pulumi.Input

	ToControlPolicyOrderMapOutput() ControlPolicyOrderMapOutput
	ToControlPolicyOrderMapOutputWithContext(context.Context) ControlPolicyOrderMapOutput
}

type ControlPolicyOrderMap map[string]ControlPolicyOrderInput

func (ControlPolicyOrderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPolicyOrder)(nil)).Elem()
}

func (i ControlPolicyOrderMap) ToControlPolicyOrderMapOutput() ControlPolicyOrderMapOutput {
	return i.ToControlPolicyOrderMapOutputWithContext(context.Background())
}

func (i ControlPolicyOrderMap) ToControlPolicyOrderMapOutputWithContext(ctx context.Context) ControlPolicyOrderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOrderMapOutput)
}

type ControlPolicyOrderOutput struct{ *pulumi.OutputState }

func (ControlPolicyOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControlPolicyOrder)(nil))
}

func (o ControlPolicyOrderOutput) ToControlPolicyOrderOutput() ControlPolicyOrderOutput {
	return o
}

func (o ControlPolicyOrderOutput) ToControlPolicyOrderOutputWithContext(ctx context.Context) ControlPolicyOrderOutput {
	return o
}

func (o ControlPolicyOrderOutput) ToControlPolicyOrderPtrOutput() ControlPolicyOrderPtrOutput {
	return o.ToControlPolicyOrderPtrOutputWithContext(context.Background())
}

func (o ControlPolicyOrderOutput) ToControlPolicyOrderPtrOutputWithContext(ctx context.Context) ControlPolicyOrderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ControlPolicyOrder) *ControlPolicyOrder {
		return &v
	}).(ControlPolicyOrderPtrOutput)
}

type ControlPolicyOrderPtrOutput struct{ *pulumi.OutputState }

func (ControlPolicyOrderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicyOrder)(nil))
}

func (o ControlPolicyOrderPtrOutput) ToControlPolicyOrderPtrOutput() ControlPolicyOrderPtrOutput {
	return o
}

func (o ControlPolicyOrderPtrOutput) ToControlPolicyOrderPtrOutputWithContext(ctx context.Context) ControlPolicyOrderPtrOutput {
	return o
}

func (o ControlPolicyOrderPtrOutput) Elem() ControlPolicyOrderOutput {
	return o.ApplyT(func(v *ControlPolicyOrder) ControlPolicyOrder {
		if v != nil {
			return *v
		}
		var ret ControlPolicyOrder
		return ret
	}).(ControlPolicyOrderOutput)
}

type ControlPolicyOrderArrayOutput struct{ *pulumi.OutputState }

func (ControlPolicyOrderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ControlPolicyOrder)(nil))
}

func (o ControlPolicyOrderArrayOutput) ToControlPolicyOrderArrayOutput() ControlPolicyOrderArrayOutput {
	return o
}

func (o ControlPolicyOrderArrayOutput) ToControlPolicyOrderArrayOutputWithContext(ctx context.Context) ControlPolicyOrderArrayOutput {
	return o
}

func (o ControlPolicyOrderArrayOutput) Index(i pulumi.IntInput) ControlPolicyOrderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ControlPolicyOrder {
		return vs[0].([]ControlPolicyOrder)[vs[1].(int)]
	}).(ControlPolicyOrderOutput)
}

type ControlPolicyOrderMapOutput struct{ *pulumi.OutputState }

func (ControlPolicyOrderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ControlPolicyOrder)(nil))
}

func (o ControlPolicyOrderMapOutput) ToControlPolicyOrderMapOutput() ControlPolicyOrderMapOutput {
	return o
}

func (o ControlPolicyOrderMapOutput) ToControlPolicyOrderMapOutputWithContext(ctx context.Context) ControlPolicyOrderMapOutput {
	return o
}

func (o ControlPolicyOrderMapOutput) MapIndex(k pulumi.StringInput) ControlPolicyOrderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ControlPolicyOrder {
		return vs[0].(map[string]ControlPolicyOrder)[vs[1].(string)]
	}).(ControlPolicyOrderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyOrderInput)(nil)).Elem(), &ControlPolicyOrder{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyOrderPtrInput)(nil)).Elem(), &ControlPolicyOrder{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyOrderArrayInput)(nil)).Elem(), ControlPolicyOrderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyOrderMapInput)(nil)).Elem(), ControlPolicyOrderMap{})
	pulumi.RegisterOutputType(ControlPolicyOrderOutput{})
	pulumi.RegisterOutputType(ControlPolicyOrderPtrOutput{})
	pulumi.RegisterOutputType(ControlPolicyOrderArrayOutput{})
	pulumi.RegisterOutputType(ControlPolicyOrderMapOutput{})
}
