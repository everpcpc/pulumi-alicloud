// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RabbitMQ (AMQP) Instance resource.
//
// For information about RabbitMQ (AMQP) Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/doc-detail/101631.htm).
//
// > **NOTE:** Available in v1.128.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := amqp.NewInstance(ctx, "professional", &amqp.InstanceArgs{
//				InstanceType:  pulumi.String("professional"),
//				MaxEipTps:     pulumi.String("128"),
//				MaxTps:        pulumi.String("1000"),
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				QueueCapacity: pulumi.String("50"),
//				SupportEip:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = amqp.NewInstance(ctx, "vip", &amqp.InstanceArgs{
//				InstanceType:  pulumi.String("vip"),
//				MaxEipTps:     pulumi.String("128"),
//				MaxTps:        pulumi.String("5000"),
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				QueueCapacity: pulumi.String("50"),
//				StorageSize:   pulumi.String("700"),
//				SupportEip:    pulumi.Bool(true),
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
// RabbitMQ (AMQP) Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:amqp/instance:Instance example <id>
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
	InstanceType pulumi.StringOutput    `pulumi:"instanceType"`
	Logistics    pulumi.StringPtrOutput `pulumi:"logistics"`
	// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
	MaxEipTps pulumi.StringPtrOutput `pulumi:"maxEipTps"`
	// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
	MaxTps pulumi.StringOutput `pulumi:"maxTps"`
	// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// The payment type. Valid values: `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The queue capacity. The smallest value is 50 and the step size 5.
	QueueCapacity pulumi.StringOutput `pulumi:"queueCapacity"`
	// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
	RenewalDuration pulumi.IntPtrOutput `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrOutput `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringOutput `pulumi:"renewalStatus"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The storage size. It is valid when `instanceType` is vip.
	StorageSize pulumi.StringPtrOutput `pulumi:"storageSize"`
	// Whether to support EIP.
	SupportEip pulumi.BoolOutput `pulumi:"supportEip"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.MaxTps == nil {
		return nil, errors.New("invalid value for required argument 'MaxTps'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.QueueCapacity == nil {
		return nil, errors.New("invalid value for required argument 'QueueCapacity'")
	}
	if args.SupportEip == nil {
		return nil, errors.New("invalid value for required argument 'SupportEip'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:amqp/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:amqp/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The instance name.
	InstanceName *string `pulumi:"instanceName"`
	// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
	InstanceType *string `pulumi:"instanceType"`
	Logistics    *string `pulumi:"logistics"`
	// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
	MaxEipTps *string `pulumi:"maxEipTps"`
	// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
	MaxTps *string `pulumi:"maxTps"`
	// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type. Valid values: `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
	Period *int `pulumi:"period"`
	// The queue capacity. The smallest value is 50 and the step size 5.
	QueueCapacity *string `pulumi:"queueCapacity"`
	// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
	RenewalDuration *int `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit *string `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The storage size. It is valid when `instanceType` is vip.
	StorageSize *string `pulumi:"storageSize"`
	// Whether to support EIP.
	SupportEip *bool `pulumi:"supportEip"`
}

type InstanceState struct {
	// The instance name.
	InstanceName pulumi.StringPtrInput
	// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
	InstanceType pulumi.StringPtrInput
	Logistics    pulumi.StringPtrInput
	// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
	MaxEipTps pulumi.StringPtrInput
	// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
	MaxTps pulumi.StringPtrInput
	// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
	ModifyType pulumi.StringPtrInput
	// The payment type. Valid values: `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
	Period pulumi.IntPtrInput
	// The queue capacity. The smallest value is 50 and the step size 5.
	QueueCapacity pulumi.StringPtrInput
	// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
	RenewalDuration pulumi.IntPtrInput
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrInput
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The storage size. It is valid when `instanceType` is vip.
	StorageSize pulumi.StringPtrInput
	// Whether to support EIP.
	SupportEip pulumi.BoolPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The instance name.
	InstanceName *string `pulumi:"instanceName"`
	// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
	InstanceType string  `pulumi:"instanceType"`
	Logistics    *string `pulumi:"logistics"`
	// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
	MaxEipTps *string `pulumi:"maxEipTps"`
	// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
	MaxTps string `pulumi:"maxTps"`
	// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
	ModifyType *string `pulumi:"modifyType"`
	// The payment type. Valid values: `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
	Period *int `pulumi:"period"`
	// The queue capacity. The smallest value is 50 and the step size 5.
	QueueCapacity string `pulumi:"queueCapacity"`
	// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
	RenewalDuration *int `pulumi:"renewalDuration"`
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit *string `pulumi:"renewalDurationUnit"`
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The storage size. It is valid when `instanceType` is vip.
	StorageSize *string `pulumi:"storageSize"`
	// Whether to support EIP.
	SupportEip bool `pulumi:"supportEip"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The instance name.
	InstanceName pulumi.StringPtrInput
	// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
	InstanceType pulumi.StringInput
	Logistics    pulumi.StringPtrInput
	// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
	MaxEipTps pulumi.StringPtrInput
	// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
	MaxTps pulumi.StringInput
	// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
	ModifyType pulumi.StringPtrInput
	// The payment type. Valid values: `Subscription`.
	PaymentType pulumi.StringInput
	// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
	Period pulumi.IntPtrInput
	// The queue capacity. The smallest value is 50 and the step size 5.
	QueueCapacity pulumi.StringInput
	// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
	RenewalDuration pulumi.IntPtrInput
	// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
	RenewalDurationUnit pulumi.StringPtrInput
	// Whether to renew an instance automatically or not. Default to "ManualRenewal".
	RenewalStatus pulumi.StringPtrInput
	// The storage size. It is valid when `instanceType` is vip.
	StorageSize pulumi.StringPtrInput
	// Whether to support EIP.
	SupportEip pulumi.BoolInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The instance name.
func (o InstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
func (o InstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InstanceOutput) Logistics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Logistics }).(pulumi.StringPtrOutput)
}

// The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
func (o InstanceOutput) MaxEipTps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.MaxEipTps }).(pulumi.StringPtrOutput)
}

// The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
func (o InstanceOutput) MaxTps() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.MaxTps }).(pulumi.StringOutput)
}

// The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
func (o InstanceOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// The payment type. Valid values: `Subscription`.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The queue capacity. The smallest value is 50 and the step size 5.
func (o InstanceOutput) QueueCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.QueueCapacity }).(pulumi.StringOutput)
}

// RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
func (o InstanceOutput) RenewalDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RenewalDuration }).(pulumi.IntPtrOutput)
}

// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
func (o InstanceOutput) RenewalDurationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RenewalDurationUnit }).(pulumi.StringPtrOutput)
}

// Whether to renew an instance automatically or not. Default to "ManualRenewal".
func (o InstanceOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

// The status of the resource.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The storage size. It is valid when `instanceType` is vip.
func (o InstanceOutput) StorageSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.StorageSize }).(pulumi.StringPtrOutput)
}

// Whether to support EIP.
func (o InstanceOutput) SupportEip() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.SupportEip }).(pulumi.BoolOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
