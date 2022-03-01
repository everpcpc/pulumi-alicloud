// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vpc.NewCommonBandwithPackage(ctx, "foo", &vpc.CommonBandwithPackageArgs{
// 			Bandwidth:            pulumi.String("1000"),
// 			BandwidthPackageName: pulumi.String("test-common-bandwidth-package"),
// 			Description:          pulumi.String("test-common-bandwidth-package"),
// 			InternetChargeType:   pulumi.String("PayByBandwidth"),
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
// The common bandwidth package can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/commonBandwithPackage:CommonBandwithPackage foo cbwp-abc123456
// ```
type CommonBandwithPackage struct {
	pulumi.CustomResourceState

	// The bandwidth of the common bandwidth package. Unit: Mbps.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// The name of the common bandwidth package.
	BandwidthPackageName pulumi.StringOutput `pulumi:"bandwidthPackageName"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// The description of the common bandwidth package instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Force       pulumi.StringPtrOutput `pulumi:"force"`
	// The billing method of the common bandwidth package. Valid values are `PayByBandwidth` and `PayBy95` and `PayByTraffic`. `PayBy95` is pay by classic 95th percentile pricing. International Account doesn't supports `PayByBandwidth` and `PayBy95`. Default to `PayByTraffic`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// The type of the Internet Service Provider. Valid values: `BGP` and `BGP_PRO`. Default to `BGP`.
	Isp pulumi.StringPtrOutput `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `bandwidthPackageName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ratio of the common bandwidth package. It is valid when `internetChargeType` is `PayBy95`. Default to `100`. Valid values: [10-100].
	Ratio pulumi.IntPtrOutput `pulumi:"ratio"`
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// (Available in 1.120.0+) The status of bandwidth package.
	Status pulumi.StringOutput `pulumi:"status"`
	// The zone of bandwidth package.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewCommonBandwithPackage registers a new resource with the given unique name, arguments, and options.
func NewCommonBandwithPackage(ctx *pulumi.Context,
	name string, args *CommonBandwithPackageArgs, opts ...pulumi.ResourceOption) (*CommonBandwithPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	var resource CommonBandwithPackage
	err := ctx.RegisterResource("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommonBandwithPackage gets an existing CommonBandwithPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommonBandwithPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommonBandwithPackageState, opts ...pulumi.ResourceOption) (*CommonBandwithPackage, error) {
	var resource CommonBandwithPackage
	err := ctx.ReadResource("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommonBandwithPackage resources.
type commonBandwithPackageState struct {
	// The bandwidth of the common bandwidth package. Unit: Mbps.
	Bandwidth *string `pulumi:"bandwidth"`
	// The name of the common bandwidth package.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the common bandwidth package instance.
	Description *string `pulumi:"description"`
	Force       *string `pulumi:"force"`
	// The billing method of the common bandwidth package. Valid values are `PayByBandwidth` and `PayBy95` and `PayByTraffic`. `PayBy95` is pay by classic 95th percentile pricing. International Account doesn't supports `PayByBandwidth` and `PayBy95`. Default to `PayByTraffic`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The type of the Internet Service Provider. Valid values: `BGP` and `BGP_PRO`. Default to `BGP`.
	Isp *string `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `bandwidthPackageName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name *string `pulumi:"name"`
	// Ratio of the common bandwidth package. It is valid when `internetChargeType` is `PayBy95`. Default to `100`. Valid values: [10-100].
	Ratio *int `pulumi:"ratio"`
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// (Available in 1.120.0+) The status of bandwidth package.
	Status *string `pulumi:"status"`
	// The zone of bandwidth package.
	Zone *string `pulumi:"zone"`
}

type CommonBandwithPackageState struct {
	// The bandwidth of the common bandwidth package. Unit: Mbps.
	Bandwidth pulumi.StringPtrInput
	// The name of the common bandwidth package.
	BandwidthPackageName pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// The description of the common bandwidth package instance.
	Description pulumi.StringPtrInput
	Force       pulumi.StringPtrInput
	// The billing method of the common bandwidth package. Valid values are `PayByBandwidth` and `PayBy95` and `PayByTraffic`. `PayBy95` is pay by classic 95th percentile pricing. International Account doesn't supports `PayByBandwidth` and `PayBy95`. Default to `PayByTraffic`.
	InternetChargeType pulumi.StringPtrInput
	// The type of the Internet Service Provider. Valid values: `BGP` and `BGP_PRO`. Default to `BGP`.
	Isp pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.120.0. New field `bandwidthPackageName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringPtrInput
	// Ratio of the common bandwidth package. It is valid when `internetChargeType` is `PayBy95`. Default to `100`. Valid values: [10-100].
	Ratio pulumi.IntPtrInput
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId pulumi.StringPtrInput
	// (Available in 1.120.0+) The status of bandwidth package.
	Status pulumi.StringPtrInput
	// The zone of bandwidth package.
	Zone pulumi.StringPtrInput
}

func (CommonBandwithPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageState)(nil)).Elem()
}

type commonBandwithPackageArgs struct {
	// The bandwidth of the common bandwidth package. Unit: Mbps.
	Bandwidth string `pulumi:"bandwidth"`
	// The name of the common bandwidth package.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the common bandwidth package instance.
	Description *string `pulumi:"description"`
	Force       *string `pulumi:"force"`
	// The billing method of the common bandwidth package. Valid values are `PayByBandwidth` and `PayBy95` and `PayByTraffic`. `PayBy95` is pay by classic 95th percentile pricing. International Account doesn't supports `PayByBandwidth` and `PayBy95`. Default to `PayByTraffic`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The type of the Internet Service Provider. Valid values: `BGP` and `BGP_PRO`. Default to `BGP`.
	Isp *string `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `bandwidthPackageName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name *string `pulumi:"name"`
	// Ratio of the common bandwidth package. It is valid when `internetChargeType` is `PayBy95`. Default to `100`. Valid values: [10-100].
	Ratio *int `pulumi:"ratio"`
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The zone of bandwidth package.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a CommonBandwithPackage resource.
type CommonBandwithPackageArgs struct {
	// The bandwidth of the common bandwidth package. Unit: Mbps.
	Bandwidth pulumi.StringInput
	// The name of the common bandwidth package.
	BandwidthPackageName pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// The description of the common bandwidth package instance.
	Description pulumi.StringPtrInput
	Force       pulumi.StringPtrInput
	// The billing method of the common bandwidth package. Valid values are `PayByBandwidth` and `PayBy95` and `PayByTraffic`. `PayBy95` is pay by classic 95th percentile pricing. International Account doesn't supports `PayByBandwidth` and `PayBy95`. Default to `PayByTraffic`.
	InternetChargeType pulumi.StringPtrInput
	// The type of the Internet Service Provider. Valid values: `BGP` and `BGP_PRO`. Default to `BGP`.
	Isp pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.120.0. New field `bandwidthPackageName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringPtrInput
	// Ratio of the common bandwidth package. It is valid when `internetChargeType` is `PayBy95`. Default to `100`. Valid values: [10-100].
	Ratio pulumi.IntPtrInput
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The zone of bandwidth package.
	Zone pulumi.StringPtrInput
}

func (CommonBandwithPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageArgs)(nil)).Elem()
}

type CommonBandwithPackageInput interface {
	pulumi.Input

	ToCommonBandwithPackageOutput() CommonBandwithPackageOutput
	ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput
}

func (*CommonBandwithPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackage)(nil)).Elem()
}

func (i *CommonBandwithPackage) ToCommonBandwithPackageOutput() CommonBandwithPackageOutput {
	return i.ToCommonBandwithPackageOutputWithContext(context.Background())
}

func (i *CommonBandwithPackage) ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageOutput)
}

// CommonBandwithPackageArrayInput is an input type that accepts CommonBandwithPackageArray and CommonBandwithPackageArrayOutput values.
// You can construct a concrete instance of `CommonBandwithPackageArrayInput` via:
//
//          CommonBandwithPackageArray{ CommonBandwithPackageArgs{...} }
type CommonBandwithPackageArrayInput interface {
	pulumi.Input

	ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput
	ToCommonBandwithPackageArrayOutputWithContext(context.Context) CommonBandwithPackageArrayOutput
}

type CommonBandwithPackageArray []CommonBandwithPackageInput

func (CommonBandwithPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackage)(nil)).Elem()
}

func (i CommonBandwithPackageArray) ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput {
	return i.ToCommonBandwithPackageArrayOutputWithContext(context.Background())
}

func (i CommonBandwithPackageArray) ToCommonBandwithPackageArrayOutputWithContext(ctx context.Context) CommonBandwithPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageArrayOutput)
}

// CommonBandwithPackageMapInput is an input type that accepts CommonBandwithPackageMap and CommonBandwithPackageMapOutput values.
// You can construct a concrete instance of `CommonBandwithPackageMapInput` via:
//
//          CommonBandwithPackageMap{ "key": CommonBandwithPackageArgs{...} }
type CommonBandwithPackageMapInput interface {
	pulumi.Input

	ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput
	ToCommonBandwithPackageMapOutputWithContext(context.Context) CommonBandwithPackageMapOutput
}

type CommonBandwithPackageMap map[string]CommonBandwithPackageInput

func (CommonBandwithPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackage)(nil)).Elem()
}

func (i CommonBandwithPackageMap) ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput {
	return i.ToCommonBandwithPackageMapOutputWithContext(context.Background())
}

func (i CommonBandwithPackageMap) ToCommonBandwithPackageMapOutputWithContext(ctx context.Context) CommonBandwithPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageMapOutput)
}

type CommonBandwithPackageOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageOutput) ToCommonBandwithPackageOutput() CommonBandwithPackageOutput {
	return o
}

func (o CommonBandwithPackageOutput) ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput {
	return o
}

type CommonBandwithPackageArrayOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageArrayOutput) ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput {
	return o
}

func (o CommonBandwithPackageArrayOutput) ToCommonBandwithPackageArrayOutputWithContext(ctx context.Context) CommonBandwithPackageArrayOutput {
	return o
}

func (o CommonBandwithPackageArrayOutput) Index(i pulumi.IntInput) CommonBandwithPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CommonBandwithPackage {
		return vs[0].([]*CommonBandwithPackage)[vs[1].(int)]
	}).(CommonBandwithPackageOutput)
}

type CommonBandwithPackageMapOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageMapOutput) ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput {
	return o
}

func (o CommonBandwithPackageMapOutput) ToCommonBandwithPackageMapOutputWithContext(ctx context.Context) CommonBandwithPackageMapOutput {
	return o
}

func (o CommonBandwithPackageMapOutput) MapIndex(k pulumi.StringInput) CommonBandwithPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CommonBandwithPackage {
		return vs[0].(map[string]*CommonBandwithPackage)[vs[1].(string)]
	}).(CommonBandwithPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageInput)(nil)).Elem(), &CommonBandwithPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageArrayInput)(nil)).Elem(), CommonBandwithPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageMapInput)(nil)).Elem(), CommonBandwithPackageMap{})
	pulumi.RegisterOutputType(CommonBandwithPackageOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageArrayOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageMapOutput{})
}
