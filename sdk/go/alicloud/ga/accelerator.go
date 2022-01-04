// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Accelerator resource.
//
// For information about Global Accelerator (GA) Accelerator and how to use it, see [What is Accelerator](https://help.aliyun.com/document_detail/153235.html).
//
// > **NOTE:** At present, The `ga.Accelerator` cannot be deleted. you need to wait until the resource is outdated and released automatically.
//
// > **NOTE:** Available in v1.111.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ga.NewAccelerator(ctx, "example", &ga.AcceleratorArgs{
// 			AutoUseCoupon: pulumi.Bool(true),
// 			Duration:      pulumi.Int(1),
// 			Spec:          pulumi.String("1"),
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
// Ga Accelerator can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ga/accelerator:Accelerator example <accelerator_id>
// ```
type Accelerator struct {
	pulumi.CustomResourceState

	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrOutput `pulumi:"acceleratorName"`
	// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
	AutoRenewDuration pulumi.IntPtrOutput `pulumi:"autoRenewDuration"`
	// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
	AutoUseCoupon pulumi.BoolPtrOutput `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricingCycle` are both required.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
	// * `Month`: billed on a monthly basis.
	// * `Year`: billed on an annual basis.
	PricingCycle pulumi.StringOutput `pulumi:"pricingCycle"`
	// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringOutput `pulumi:"renewalStatus"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringOutput `pulumi:"spec"`
	// The status of the GA instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAccelerator registers a new resource with the given unique name, arguments, and options.
func NewAccelerator(ctx *pulumi.Context,
	name string, args *AcceleratorArgs, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource Accelerator
	err := ctx.RegisterResource("alicloud:ga/accelerator:Accelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccelerator gets an existing Accelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcceleratorState, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	var resource Accelerator
	err := ctx.ReadResource("alicloud:ga/accelerator:Accelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accelerator resources.
type acceleratorState struct {
	// The Name of the GA instance.
	AcceleratorName *string `pulumi:"acceleratorName"`
	// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description *string `pulumi:"description"`
	// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricingCycle` are both required.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
	Duration *int `pulumi:"duration"`
	// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
	// * `Month`: billed on a monthly basis.
	// * `Year`: billed on an annual basis.
	PricingCycle *string `pulumi:"pricingCycle"`
	// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec *string `pulumi:"spec"`
	// The status of the GA instance.
	Status *string `pulumi:"status"`
}

type AcceleratorState struct {
	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrInput
	// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
	AutoRenewDuration pulumi.IntPtrInput
	// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
	AutoUseCoupon pulumi.BoolPtrInput
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrInput
	// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricingCycle` are both required.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
	Duration pulumi.IntPtrInput
	// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
	// * `Month`: billed on a monthly basis.
	// * `Year`: billed on an annual basis.
	PricingCycle pulumi.StringPtrInput
	// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringPtrInput
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringPtrInput
	// The status of the GA instance.
	Status pulumi.StringPtrInput
}

func (AcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorState)(nil)).Elem()
}

type acceleratorArgs struct {
	// The Name of the GA instance.
	AcceleratorName *string `pulumi:"acceleratorName"`
	// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// Descriptive information of the global acceleration instance.
	Description *string `pulumi:"description"`
	// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricingCycle` are both required.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
	Duration int `pulumi:"duration"`
	// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
	// * `Month`: billed on a monthly basis.
	// * `Year`: billed on an annual basis.
	PricingCycle *string `pulumi:"pricingCycle"`
	// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec string `pulumi:"spec"`
}

// The set of arguments for constructing a Accelerator resource.
type AcceleratorArgs struct {
	// The Name of the GA instance.
	AcceleratorName pulumi.StringPtrInput
	// Auto renewal period of an instance, in the unit of month. The value range is 1-12.
	AutoRenewDuration pulumi.IntPtrInput
	// Use coupons to pay bills automatically. Default value is `false`. Valid value: `true`: Use, `false`: Not used.
	AutoUseCoupon pulumi.BoolPtrInput
	// Descriptive information of the global acceleration instance.
	Description pulumi.StringPtrInput
	// The subscription duration. **NOTE:** Starting from v1.150.0+, the `duration` and  `pricingCycle` are both required.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are 1 to 9.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are 1 to 3.
	Duration pulumi.IntInput
	// The billing cycle of the GA instance. Valid values: `Month`,`Year`. The default value: `Month`.
	// * `Month`: billed on a monthly basis.
	// * `Year`: billed on an annual basis.
	PricingCycle pulumi.StringPtrInput
	// Whether to renew an accelerator automatically or not. Default to "Normal". Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringPtrInput
	// The instance type of the GA instance. Specification of global acceleration instance, value:
	// `1`: Small 1.
	// `2`: Small 2.
	// `3`: Small 3.
	// `5`: Medium 1.
	// `8`: Medium 2.
	// `10`: Medium 3.
	Spec pulumi.StringInput
}

func (AcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorArgs)(nil)).Elem()
}

type AcceleratorInput interface {
	pulumi.Input

	ToAcceleratorOutput() AcceleratorOutput
	ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput
}

func (*Accelerator) ElementType() reflect.Type {
	return reflect.TypeOf((*Accelerator)(nil))
}

func (i *Accelerator) ToAcceleratorOutput() AcceleratorOutput {
	return i.ToAcceleratorOutputWithContext(context.Background())
}

func (i *Accelerator) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorOutput)
}

func (i *Accelerator) ToAcceleratorPtrOutput() AcceleratorPtrOutput {
	return i.ToAcceleratorPtrOutputWithContext(context.Background())
}

func (i *Accelerator) ToAcceleratorPtrOutputWithContext(ctx context.Context) AcceleratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorPtrOutput)
}

type AcceleratorPtrInput interface {
	pulumi.Input

	ToAcceleratorPtrOutput() AcceleratorPtrOutput
	ToAcceleratorPtrOutputWithContext(ctx context.Context) AcceleratorPtrOutput
}

type acceleratorPtrType AcceleratorArgs

func (*acceleratorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Accelerator)(nil))
}

func (i *acceleratorPtrType) ToAcceleratorPtrOutput() AcceleratorPtrOutput {
	return i.ToAcceleratorPtrOutputWithContext(context.Background())
}

func (i *acceleratorPtrType) ToAcceleratorPtrOutputWithContext(ctx context.Context) AcceleratorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorPtrOutput)
}

// AcceleratorArrayInput is an input type that accepts AcceleratorArray and AcceleratorArrayOutput values.
// You can construct a concrete instance of `AcceleratorArrayInput` via:
//
//          AcceleratorArray{ AcceleratorArgs{...} }
type AcceleratorArrayInput interface {
	pulumi.Input

	ToAcceleratorArrayOutput() AcceleratorArrayOutput
	ToAcceleratorArrayOutputWithContext(context.Context) AcceleratorArrayOutput
}

type AcceleratorArray []AcceleratorInput

func (AcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Accelerator)(nil)).Elem()
}

func (i AcceleratorArray) ToAcceleratorArrayOutput() AcceleratorArrayOutput {
	return i.ToAcceleratorArrayOutputWithContext(context.Background())
}

func (i AcceleratorArray) ToAcceleratorArrayOutputWithContext(ctx context.Context) AcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorArrayOutput)
}

// AcceleratorMapInput is an input type that accepts AcceleratorMap and AcceleratorMapOutput values.
// You can construct a concrete instance of `AcceleratorMapInput` via:
//
//          AcceleratorMap{ "key": AcceleratorArgs{...} }
type AcceleratorMapInput interface {
	pulumi.Input

	ToAcceleratorMapOutput() AcceleratorMapOutput
	ToAcceleratorMapOutputWithContext(context.Context) AcceleratorMapOutput
}

type AcceleratorMap map[string]AcceleratorInput

func (AcceleratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Accelerator)(nil)).Elem()
}

func (i AcceleratorMap) ToAcceleratorMapOutput() AcceleratorMapOutput {
	return i.ToAcceleratorMapOutputWithContext(context.Background())
}

func (i AcceleratorMap) ToAcceleratorMapOutputWithContext(ctx context.Context) AcceleratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorMapOutput)
}

type AcceleratorOutput struct{ *pulumi.OutputState }

func (AcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Accelerator)(nil))
}

func (o AcceleratorOutput) ToAcceleratorOutput() AcceleratorOutput {
	return o
}

func (o AcceleratorOutput) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return o
}

func (o AcceleratorOutput) ToAcceleratorPtrOutput() AcceleratorPtrOutput {
	return o.ToAcceleratorPtrOutputWithContext(context.Background())
}

func (o AcceleratorOutput) ToAcceleratorPtrOutputWithContext(ctx context.Context) AcceleratorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Accelerator) *Accelerator {
		return &v
	}).(AcceleratorPtrOutput)
}

type AcceleratorPtrOutput struct{ *pulumi.OutputState }

func (AcceleratorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accelerator)(nil))
}

func (o AcceleratorPtrOutput) ToAcceleratorPtrOutput() AcceleratorPtrOutput {
	return o
}

func (o AcceleratorPtrOutput) ToAcceleratorPtrOutputWithContext(ctx context.Context) AcceleratorPtrOutput {
	return o
}

func (o AcceleratorPtrOutput) Elem() AcceleratorOutput {
	return o.ApplyT(func(v *Accelerator) Accelerator {
		if v != nil {
			return *v
		}
		var ret Accelerator
		return ret
	}).(AcceleratorOutput)
}

type AcceleratorArrayOutput struct{ *pulumi.OutputState }

func (AcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Accelerator)(nil))
}

func (o AcceleratorArrayOutput) ToAcceleratorArrayOutput() AcceleratorArrayOutput {
	return o
}

func (o AcceleratorArrayOutput) ToAcceleratorArrayOutputWithContext(ctx context.Context) AcceleratorArrayOutput {
	return o
}

func (o AcceleratorArrayOutput) Index(i pulumi.IntInput) AcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Accelerator {
		return vs[0].([]Accelerator)[vs[1].(int)]
	}).(AcceleratorOutput)
}

type AcceleratorMapOutput struct{ *pulumi.OutputState }

func (AcceleratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Accelerator)(nil))
}

func (o AcceleratorMapOutput) ToAcceleratorMapOutput() AcceleratorMapOutput {
	return o
}

func (o AcceleratorMapOutput) ToAcceleratorMapOutputWithContext(ctx context.Context) AcceleratorMapOutput {
	return o
}

func (o AcceleratorMapOutput) MapIndex(k pulumi.StringInput) AcceleratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Accelerator {
		return vs[0].(map[string]Accelerator)[vs[1].(string)]
	}).(AcceleratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorInput)(nil)).Elem(), &Accelerator{})
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorPtrInput)(nil)).Elem(), &Accelerator{})
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorArrayInput)(nil)).Elem(), AcceleratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AcceleratorMapInput)(nil)).Elem(), AcceleratorMap{})
	pulumi.RegisterOutputType(AcceleratorOutput{})
	pulumi.RegisterOutputType(AcceleratorPtrOutput{})
	pulumi.RegisterOutputType(AcceleratorArrayOutput{})
	pulumi.RegisterOutputType(AcceleratorMapOutput{})
}
