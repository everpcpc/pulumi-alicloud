// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Global Accelerator (GA) Ip Set resource.
//
// For information about Global Accelerator (GA) Ip Set and how to use it, see [What is Ip Set](https://www.alibabacloud.com/help/en/doc-detail/153246.htm).
//
// > **NOTE:** Available in v1.113.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ga"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccelerator, err := ga.NewAccelerator(ctx, "exampleAccelerator", &ga.AcceleratorArgs{
// 			Duration:      pulumi.Int(1),
// 			AutoUseCoupon: pulumi.Bool(true),
// 			Spec:          pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "exampleBandwidthPackage", &ga.BandwidthPackageArgs{
// 			Bandwidth:     pulumi.Int(20),
// 			Type:          pulumi.String("Basic"),
// 			BandwidthType: pulumi.String("Basic"),
// 			Duration:      pulumi.String("1"),
// 			AutoPay:       pulumi.Bool(true),
// 			Ratio:         pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "exampleBandwidthPackageAttachment", &ga.BandwidthPackageAttachmentArgs{
// 			AcceleratorId:      exampleAccelerator.ID(),
// 			BandwidthPackageId: exampleBandwidthPackage.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ga.NewIpSet(ctx, "exampleIpSet", &ga.IpSetArgs{
// 			AccelerateRegionId: pulumi.String("cn-hangzhou"),
// 			Bandwidth:          pulumi.Int(5),
// 			AcceleratorId:      exampleAccelerator.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleBandwidthPackageAttachment,
// 		}))
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
// Ga Ip Set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ga/ipSet:IpSet example <id>
// ```
type IpSet struct {
	pulumi.CustomResourceState

	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringOutput `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists pulumi.StringArrayOutput `pulumi:"ipAddressLists"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The status of the acceleration region.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccelerateRegionId == nil {
		return nil, errors.New("invalid value for required argument 'AccelerateRegionId'")
	}
	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	var resource IpSet
	err := ctx.RegisterResource("alicloud:ga/ipSet:IpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSetState, opts ...pulumi.ResourceOption) (*IpSet, error) {
	var resource IpSet
	err := ctx.ReadResource("alicloud:ga/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// The ID of an acceleration region.
	AccelerateRegionId *string `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	Bandwidth *int `pulumi:"bandwidth"`
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists []string `pulumi:"ipAddressLists"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
	// The status of the acceleration region.
	Status *string `pulumi:"status"`
}

type IpSetState struct {
	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringPtrInput
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringPtrInput
	// The bandwidth allocated to the acceleration region.
	Bandwidth pulumi.IntPtrInput
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists pulumi.StringArrayInput
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
	IpVersion pulumi.StringPtrInput
	// The status of the acceleration region.
	Status pulumi.StringPtrInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// The ID of an acceleration region.
	AccelerateRegionId string `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	Bandwidth *int `pulumi:"bandwidth"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringInput
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringInput
	// The bandwidth allocated to the acceleration region.
	Bandwidth pulumi.IntPtrInput
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
	IpVersion pulumi.StringPtrInput
}

func (IpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetArgs)(nil)).Elem()
}

type IpSetInput interface {
	pulumi.Input

	ToIpSetOutput() IpSetOutput
	ToIpSetOutputWithContext(ctx context.Context) IpSetOutput
}

func (*IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSet)(nil))
}

func (i *IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i *IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

func (i *IpSet) ToIpSetPtrOutput() IpSetPtrOutput {
	return i.ToIpSetPtrOutputWithContext(context.Background())
}

func (i *IpSet) ToIpSetPtrOutputWithContext(ctx context.Context) IpSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetPtrOutput)
}

type IpSetPtrInput interface {
	pulumi.Input

	ToIpSetPtrOutput() IpSetPtrOutput
	ToIpSetPtrOutputWithContext(ctx context.Context) IpSetPtrOutput
}

type ipSetPtrType IpSetArgs

func (*ipSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil))
}

func (i *ipSetPtrType) ToIpSetPtrOutput() IpSetPtrOutput {
	return i.ToIpSetPtrOutputWithContext(context.Background())
}

func (i *ipSetPtrType) ToIpSetPtrOutputWithContext(ctx context.Context) IpSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetPtrOutput)
}

// IpSetArrayInput is an input type that accepts IpSetArray and IpSetArrayOutput values.
// You can construct a concrete instance of `IpSetArrayInput` via:
//
//          IpSetArray{ IpSetArgs{...} }
type IpSetArrayInput interface {
	pulumi.Input

	ToIpSetArrayOutput() IpSetArrayOutput
	ToIpSetArrayOutputWithContext(context.Context) IpSetArrayOutput
}

type IpSetArray []IpSetInput

func (IpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IpSet)(nil))
}

func (i IpSetArray) ToIpSetArrayOutput() IpSetArrayOutput {
	return i.ToIpSetArrayOutputWithContext(context.Background())
}

func (i IpSetArray) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetArrayOutput)
}

// IpSetMapInput is an input type that accepts IpSetMap and IpSetMapOutput values.
// You can construct a concrete instance of `IpSetMapInput` via:
//
//          IpSetMap{ "key": IpSetArgs{...} }
type IpSetMapInput interface {
	pulumi.Input

	ToIpSetMapOutput() IpSetMapOutput
	ToIpSetMapOutputWithContext(context.Context) IpSetMapOutput
}

type IpSetMap map[string]IpSetInput

func (IpSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IpSet)(nil))
}

func (i IpSetMap) ToIpSetMapOutput() IpSetMapOutput {
	return i.ToIpSetMapOutputWithContext(context.Background())
}

func (i IpSetMap) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetMapOutput)
}

type IpSetOutput struct {
	*pulumi.OutputState
}

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSet)(nil))
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetPtrOutput() IpSetPtrOutput {
	return o.ToIpSetPtrOutputWithContext(context.Background())
}

func (o IpSetOutput) ToIpSetPtrOutputWithContext(ctx context.Context) IpSetPtrOutput {
	return o.ApplyT(func(v IpSet) *IpSet {
		return &v
	}).(IpSetPtrOutput)
}

type IpSetPtrOutput struct {
	*pulumi.OutputState
}

func (IpSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil))
}

func (o IpSetPtrOutput) ToIpSetPtrOutput() IpSetPtrOutput {
	return o
}

func (o IpSetPtrOutput) ToIpSetPtrOutputWithContext(ctx context.Context) IpSetPtrOutput {
	return o
}

type IpSetArrayOutput struct{ *pulumi.OutputState }

func (IpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSet)(nil))
}

func (o IpSetArrayOutput) ToIpSetArrayOutput() IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) Index(i pulumi.IntInput) IpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSet {
		return vs[0].([]IpSet)[vs[1].(int)]
	}).(IpSetOutput)
}

type IpSetMapOutput struct{ *pulumi.OutputState }

func (IpSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IpSet)(nil))
}

func (o IpSetMapOutput) ToIpSetMapOutput() IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) MapIndex(k pulumi.StringInput) IpSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IpSet {
		return vs[0].(map[string]IpSet)[vs[1].(string)]
	}).(IpSetOutput)
}

func init() {
	pulumi.RegisterOutputType(IpSetOutput{})
	pulumi.RegisterOutputType(IpSetPtrOutput{})
	pulumi.RegisterOutputType(IpSetArrayOutput{})
	pulumi.RegisterOutputType(IpSetMapOutput{})
}
