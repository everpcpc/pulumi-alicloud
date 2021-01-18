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
// 		_, err := ga.NewAccelerator(ctx, "exampleAccelerator", &ga.AcceleratorArgs{
// 			Duration:      pulumi.Int(1),
// 			AutoUseCoupon: pulumi.Bool(true),
// 			Spec:          pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ga.NewBandwidthPackage(ctx, "exampleBandwidthPackage", &ga.BandwidthPackageArgs{
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
// 			AcceleratorId:      pulumi.String("alicloud_ga_accelerator.example.id"),
// 			BandwidthPackageId: pulumi.String("alicloud_ga_bandwidth_package.example.id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ga.NewIpSet(ctx, "exampleIpSet", &ga.IpSetArgs{
// 			AccelerateRegionId: pulumi.String("cn-hangzhou"),
// 			Bandwidth:          pulumi.Int(5),
// 			AcceleratorId:      pulumi.String("alicloud_ga_accelerator.example.id"),
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

func (IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSet)(nil)).Elem()
}

func (i IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

type IpSetOutput struct {
	*pulumi.OutputState
}

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSetOutput)(nil)).Elem()
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IpSetOutput{})
}
