// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eipanycast

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Eipanycast Anycast Eip Address resource.
//
// For information about Eipanycast Anycast Eip Address and how to use it, see [What is Anycast Eip Address](https://help.aliyun.com/document_detail/169284.html).
//
// > **NOTE:** Available in v1.113.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eipanycast"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eipanycast.NewAnycastEipAddress(ctx, "example", &eipanycast.AnycastEipAddressArgs{
//				ServiceLocation: pulumi.String("international"),
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
// Eipanycast Anycast Eip Address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:eipanycast/anycastEipAddress:AnycastEipAddress example <id>
//
// ```
type AnycastEipAddress struct {
	pulumi.CustomResourceState

	// Anycast EIP instance name.
	AnycastEipAddressName pulumi.StringPtrOutput `pulumi:"anycastEipAddressName"`
	// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// Anycast EIP instance description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrOutput `pulumi:"paymentType"`
	// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
	ServiceLocation pulumi.StringOutput `pulumi:"serviceLocation"`
	// The IP status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAnycastEipAddress registers a new resource with the given unique name, arguments, and options.
func NewAnycastEipAddress(ctx *pulumi.Context,
	name string, args *AnycastEipAddressArgs, opts ...pulumi.ResourceOption) (*AnycastEipAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceLocation == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLocation'")
	}
	var resource AnycastEipAddress
	err := ctx.RegisterResource("alicloud:eipanycast/anycastEipAddress:AnycastEipAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnycastEipAddress gets an existing AnycastEipAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnycastEipAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnycastEipAddressState, opts ...pulumi.ResourceOption) (*AnycastEipAddress, error) {
	var resource AnycastEipAddress
	err := ctx.ReadResource("alicloud:eipanycast/anycastEipAddress:AnycastEipAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnycastEipAddress resources.
type anycastEipAddressState struct {
	// Anycast EIP instance name.
	AnycastEipAddressName *string `pulumi:"anycastEipAddressName"`
	// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
	Bandwidth *int `pulumi:"bandwidth"`
	// Anycast EIP instance description.
	Description *string `pulumi:"description"`
	// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
	ServiceLocation *string `pulumi:"serviceLocation"`
	// The IP status.
	Status *string `pulumi:"status"`
}

type AnycastEipAddressState struct {
	// Anycast EIP instance name.
	AnycastEipAddressName pulumi.StringPtrInput
	// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
	Bandwidth pulumi.IntPtrInput
	// Anycast EIP instance description.
	Description pulumi.StringPtrInput
	// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
	ServiceLocation pulumi.StringPtrInput
	// The IP status.
	Status pulumi.StringPtrInput
}

func (AnycastEipAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressState)(nil)).Elem()
}

type anycastEipAddressArgs struct {
	// Anycast EIP instance name.
	AnycastEipAddressName *string `pulumi:"anycastEipAddressName"`
	// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
	Bandwidth *int `pulumi:"bandwidth"`
	// Anycast EIP instance description.
	Description *string `pulumi:"description"`
	// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
	ServiceLocation string `pulumi:"serviceLocation"`
}

// The set of arguments for constructing a AnycastEipAddress resource.
type AnycastEipAddressArgs struct {
	// Anycast EIP instance name.
	AnycastEipAddressName pulumi.StringPtrInput
	// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
	Bandwidth pulumi.IntPtrInput
	// Anycast EIP instance description.
	Description pulumi.StringPtrInput
	// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
	ServiceLocation pulumi.StringInput
}

func (AnycastEipAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressArgs)(nil)).Elem()
}

type AnycastEipAddressInput interface {
	pulumi.Input

	ToAnycastEipAddressOutput() AnycastEipAddressOutput
	ToAnycastEipAddressOutputWithContext(ctx context.Context) AnycastEipAddressOutput
}

func (*AnycastEipAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddress)(nil)).Elem()
}

func (i *AnycastEipAddress) ToAnycastEipAddressOutput() AnycastEipAddressOutput {
	return i.ToAnycastEipAddressOutputWithContext(context.Background())
}

func (i *AnycastEipAddress) ToAnycastEipAddressOutputWithContext(ctx context.Context) AnycastEipAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressOutput)
}

// AnycastEipAddressArrayInput is an input type that accepts AnycastEipAddressArray and AnycastEipAddressArrayOutput values.
// You can construct a concrete instance of `AnycastEipAddressArrayInput` via:
//
//	AnycastEipAddressArray{ AnycastEipAddressArgs{...} }
type AnycastEipAddressArrayInput interface {
	pulumi.Input

	ToAnycastEipAddressArrayOutput() AnycastEipAddressArrayOutput
	ToAnycastEipAddressArrayOutputWithContext(context.Context) AnycastEipAddressArrayOutput
}

type AnycastEipAddressArray []AnycastEipAddressInput

func (AnycastEipAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnycastEipAddress)(nil)).Elem()
}

func (i AnycastEipAddressArray) ToAnycastEipAddressArrayOutput() AnycastEipAddressArrayOutput {
	return i.ToAnycastEipAddressArrayOutputWithContext(context.Background())
}

func (i AnycastEipAddressArray) ToAnycastEipAddressArrayOutputWithContext(ctx context.Context) AnycastEipAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressArrayOutput)
}

// AnycastEipAddressMapInput is an input type that accepts AnycastEipAddressMap and AnycastEipAddressMapOutput values.
// You can construct a concrete instance of `AnycastEipAddressMapInput` via:
//
//	AnycastEipAddressMap{ "key": AnycastEipAddressArgs{...} }
type AnycastEipAddressMapInput interface {
	pulumi.Input

	ToAnycastEipAddressMapOutput() AnycastEipAddressMapOutput
	ToAnycastEipAddressMapOutputWithContext(context.Context) AnycastEipAddressMapOutput
}

type AnycastEipAddressMap map[string]AnycastEipAddressInput

func (AnycastEipAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnycastEipAddress)(nil)).Elem()
}

func (i AnycastEipAddressMap) ToAnycastEipAddressMapOutput() AnycastEipAddressMapOutput {
	return i.ToAnycastEipAddressMapOutputWithContext(context.Background())
}

func (i AnycastEipAddressMap) ToAnycastEipAddressMapOutputWithContext(ctx context.Context) AnycastEipAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressMapOutput)
}

type AnycastEipAddressOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddress)(nil)).Elem()
}

func (o AnycastEipAddressOutput) ToAnycastEipAddressOutput() AnycastEipAddressOutput {
	return o
}

func (o AnycastEipAddressOutput) ToAnycastEipAddressOutputWithContext(ctx context.Context) AnycastEipAddressOutput {
	return o
}

// Anycast EIP instance name.
func (o AnycastEipAddressOutput) AnycastEipAddressName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringPtrOutput { return v.AnycastEipAddressName }).(pulumi.StringPtrOutput)
}

// The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
func (o AnycastEipAddressOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// Anycast EIP instance description.
func (o AnycastEipAddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
func (o AnycastEipAddressOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringPtrOutput { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

// The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
func (o AnycastEipAddressOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringPtrOutput { return v.PaymentType }).(pulumi.StringPtrOutput)
}

// Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
func (o AnycastEipAddressOutput) ServiceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringOutput { return v.ServiceLocation }).(pulumi.StringOutput)
}

// The IP status.
func (o AnycastEipAddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddress) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AnycastEipAddressArrayOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnycastEipAddress)(nil)).Elem()
}

func (o AnycastEipAddressArrayOutput) ToAnycastEipAddressArrayOutput() AnycastEipAddressArrayOutput {
	return o
}

func (o AnycastEipAddressArrayOutput) ToAnycastEipAddressArrayOutputWithContext(ctx context.Context) AnycastEipAddressArrayOutput {
	return o
}

func (o AnycastEipAddressArrayOutput) Index(i pulumi.IntInput) AnycastEipAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnycastEipAddress {
		return vs[0].([]*AnycastEipAddress)[vs[1].(int)]
	}).(AnycastEipAddressOutput)
}

type AnycastEipAddressMapOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnycastEipAddress)(nil)).Elem()
}

func (o AnycastEipAddressMapOutput) ToAnycastEipAddressMapOutput() AnycastEipAddressMapOutput {
	return o
}

func (o AnycastEipAddressMapOutput) ToAnycastEipAddressMapOutputWithContext(ctx context.Context) AnycastEipAddressMapOutput {
	return o
}

func (o AnycastEipAddressMapOutput) MapIndex(k pulumi.StringInput) AnycastEipAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnycastEipAddress {
		return vs[0].(map[string]*AnycastEipAddress)[vs[1].(string)]
	}).(AnycastEipAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressInput)(nil)).Elem(), &AnycastEipAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressArrayInput)(nil)).Elem(), AnycastEipAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressMapInput)(nil)).Elem(), AnycastEipAddressMap{})
	pulumi.RegisterOutputType(AnycastEipAddressOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressArrayOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressMapOutput{})
}
