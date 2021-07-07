// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Nat gateway can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/natGateway:NatGateway example ngw-abc123456
// ```
type NatGateway struct {
	pulumi.CustomResourceState

	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	DryRun      pulumi.BoolPtrOutput   `pulumi:"dryRun"`
	Force       pulumi.BoolPtrOutput   `pulumi:"force"`
	// The nat gateway will auto create a forward item.
	ForwardTableIds pulumi.StringOutput `pulumi:"forwardTableIds"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`, default value is `PayBySpec`. The `PayByLcu` is only support enhanced NAT.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringOutput `pulumi:"natGatewayName"`
	// The type of NAT gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
	NatType pulumi.StringPtrOutput `pulumi:"natType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	Period      pulumi.IntPtrOutput `pulumi:"period"`
	// The nat gateway will auto create a snat item.
	SnatTableIds pulumi.StringOutput `pulumi:"snatTableIds"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Effective when `internetChargeType` is `PayBySpec`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
	Specification pulumi.StringPtrOutput `pulumi:"specification"`
	// (Available in 1.121.0+) The status of NAT gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of NAT gateway.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource NatGateway
	err := ctx.RegisterResource("alicloud:vpc/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("alicloud:vpc/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description *string `pulumi:"description"`
	DryRun      *bool   `pulumi:"dryRun"`
	Force       *bool   `pulumi:"force"`
	// The nat gateway will auto create a forward item.
	ForwardTableIds *string `pulumi:"forwardTableIds"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`, default value is `PayBySpec`. The `PayByLcu` is only support enhanced NAT.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name *string `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The type of NAT gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
	NatType *string `pulumi:"natType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	Period      *int    `pulumi:"period"`
	// The nat gateway will auto create a snat item.
	SnatTableIds *string `pulumi:"snatTableIds"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Effective when `internetChargeType` is `PayBySpec`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
	Specification *string `pulumi:"specification"`
	// (Available in 1.121.0+) The status of NAT gateway.
	Status *string `pulumi:"status"`
	// The tags of NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type NatGatewayState struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrInput
	DryRun      pulumi.BoolPtrInput
	Force       pulumi.BoolPtrInput
	// The nat gateway will auto create a forward item.
	ForwardTableIds pulumi.StringPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringPtrInput
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`, default value is `PayBySpec`. The `PayByLcu` is only support enhanced NAT.
	InternetChargeType pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringPtrInput
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringPtrInput
	// The type of NAT gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
	NatType pulumi.StringPtrInput
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	Period      pulumi.IntPtrInput
	// The nat gateway will auto create a snat item.
	SnatTableIds pulumi.StringPtrInput
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Effective when `internetChargeType` is `PayBySpec`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
	Specification pulumi.StringPtrInput
	// (Available in 1.121.0+) The status of NAT gateway.
	Status pulumi.StringPtrInput
	// The tags of NAT gateway.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description *string `pulumi:"description"`
	DryRun      *bool   `pulumi:"dryRun"`
	Force       *bool   `pulumi:"force"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`, default value is `PayBySpec`. The `PayByLcu` is only support enhanced NAT.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name *string `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The type of NAT gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
	NatType *string `pulumi:"natType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	Period      *int    `pulumi:"period"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Effective when `internetChargeType` is `PayBySpec`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
	Specification *string `pulumi:"specification"`
	// The tags of NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrInput
	DryRun      pulumi.BoolPtrInput
	Force       pulumi.BoolPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringPtrInput
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`, default value is `PayBySpec`. The `PayByLcu` is only support enhanced NAT.
	InternetChargeType pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringPtrInput
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringPtrInput
	// The type of NAT gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
	NatType pulumi.StringPtrInput
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	Period      pulumi.IntPtrInput
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Effective when `internetChargeType` is `PayBySpec`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
	Specification pulumi.StringPtrInput
	// The tags of NAT gateway.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGateway)(nil))
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

func (i *NatGateway) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return i.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPtrOutput)
}

type NatGatewayPtrInput interface {
	pulumi.Input

	ToNatGatewayPtrOutput() NatGatewayPtrOutput
	ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput
}

type natGatewayPtrType NatGatewayArgs

func (*natGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil))
}

func (i *natGatewayPtrType) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return i.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (i *natGatewayPtrType) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPtrOutput)
}

// NatGatewayArrayInput is an input type that accepts NatGatewayArray and NatGatewayArrayOutput values.
// You can construct a concrete instance of `NatGatewayArrayInput` via:
//
//          NatGatewayArray{ NatGatewayArgs{...} }
type NatGatewayArrayInput interface {
	pulumi.Input

	ToNatGatewayArrayOutput() NatGatewayArrayOutput
	ToNatGatewayArrayOutputWithContext(context.Context) NatGatewayArrayOutput
}

type NatGatewayArray []NatGatewayInput

func (NatGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NatGateway)(nil))
}

func (i NatGatewayArray) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return i.ToNatGatewayArrayOutputWithContext(context.Background())
}

func (i NatGatewayArray) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayArrayOutput)
}

// NatGatewayMapInput is an input type that accepts NatGatewayMap and NatGatewayMapOutput values.
// You can construct a concrete instance of `NatGatewayMapInput` via:
//
//          NatGatewayMap{ "key": NatGatewayArgs{...} }
type NatGatewayMapInput interface {
	pulumi.Input

	ToNatGatewayMapOutput() NatGatewayMapOutput
	ToNatGatewayMapOutputWithContext(context.Context) NatGatewayMapOutput
}

type NatGatewayMap map[string]NatGatewayInput

func (NatGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NatGateway)(nil))
}

func (i NatGatewayMap) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return i.ToNatGatewayMapOutputWithContext(context.Background())
}

func (i NatGatewayMap) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayMapOutput)
}

type NatGatewayOutput struct {
	*pulumi.OutputState
}

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGateway)(nil))
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return o.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (o NatGatewayOutput) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return o.ApplyT(func(v NatGateway) *NatGateway {
		return &v
	}).(NatGatewayPtrOutput)
}

type NatGatewayPtrOutput struct {
	*pulumi.OutputState
}

func (NatGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil))
}

func (o NatGatewayPtrOutput) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return o
}

func (o NatGatewayPtrOutput) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return o
}

type NatGatewayArrayOutput struct{ *pulumi.OutputState }

func (NatGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NatGateway)(nil))
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) Index(i pulumi.IntInput) NatGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NatGateway {
		return vs[0].([]NatGateway)[vs[1].(int)]
	}).(NatGatewayOutput)
}

type NatGatewayMapOutput struct{ *pulumi.OutputState }

func (NatGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NatGateway)(nil))
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) MapIndex(k pulumi.StringInput) NatGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NatGateway {
		return vs[0].(map[string]NatGateway)[vs[1].(string)]
	}).(NatGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(NatGatewayOutput{})
	pulumi.RegisterOutputType(NatGatewayPtrOutput{})
	pulumi.RegisterOutputType(NatGatewayArrayOutput{})
	pulumi.RegisterOutputType(NatGatewayMapOutput{})
}
