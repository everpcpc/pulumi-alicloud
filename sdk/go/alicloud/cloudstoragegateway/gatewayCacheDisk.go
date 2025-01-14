// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Storage Gateway Gateway Cache Disk resource.
//
// For information about Cloud Storage Gateway Gateway Cache Disk and how to use it, see [What is Gateway Cache Disk](https://www.alibabacloud.com/help/zh/doc-detail/170294.htm).
//
// > **NOTE:** Available in v1.144.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudstoragegateway"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleStocks, err := cloudstoragegateway.GetStocks(ctx, &cloudstoragegateway.GetStocksArgs{
//				GatewayClass: pulumi.StringRef("Standard"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("example_value"),
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VpcId:       vpc.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/21"),
//				ZoneId:      *pulumi.String(exampleStocks.Stocks[0].ZoneId),
//				VswitchName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleStorageBundle, err := cloudstoragegateway.NewStorageBundle(ctx, "exampleStorageBundle", &cloudstoragegateway.StorageBundleArgs{
//				StorageBundleName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudstoragegateway.NewGateway(ctx, "exampleGateway", &cloudstoragegateway.GatewayArgs{
//				Description:            pulumi.String("tf-acctestDesalone"),
//				GatewayClass:           pulumi.String("Standard"),
//				Type:                   pulumi.String("File"),
//				PaymentType:            pulumi.String("PayAsYouGo"),
//				VswitchId:              exampleSwitch.ID(),
//				ReleaseAfterExpiration: pulumi.Bool(true),
//				PublicNetworkBandwidth: pulumi.Int(10),
//				StorageBundleId:        exampleStorageBundle.ID(),
//				Location:               pulumi.String("Cloud"),
//				GatewayName:            pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudstoragegateway.NewGatewayCacheDisk(ctx, "exampleGatewayCacheDisk", &cloudstoragegateway.GatewayCacheDiskArgs{
//				CacheDiskCategory: pulumi.String("cloud_efficiency"),
//				GatewayId:         pulumi.Any(alicloud_cloud_storage_gateway_gateways.Example.Id),
//				CacheDiskSizeInGb: pulumi.Int(50),
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
// Cloud Storage Gateway Gateway Cache Disk can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk example <gateway_id>:<cache_id>:<local_file_path>
//
// ```
type GatewayCacheDisk struct {
	pulumi.CustomResourceState

	// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
	CacheDiskCategory pulumi.StringOutput `pulumi:"cacheDiskCategory"`
	// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
	CacheDiskSizeInGb pulumi.IntOutput `pulumi:"cacheDiskSizeInGb"`
	// The ID of the cache.
	CacheId pulumi.StringOutput `pulumi:"cacheId"`
	// The ID of the gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The cache disk inside the device name.
	LocalFilePath pulumi.StringOutput `pulumi:"localFilePath"`
	// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewGatewayCacheDisk registers a new resource with the given unique name, arguments, and options.
func NewGatewayCacheDisk(ctx *pulumi.Context,
	name string, args *GatewayCacheDiskArgs, opts ...pulumi.ResourceOption) (*GatewayCacheDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheDiskSizeInGb == nil {
		return nil, errors.New("invalid value for required argument 'CacheDiskSizeInGb'")
	}
	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	var resource GatewayCacheDisk
	err := ctx.RegisterResource("alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayCacheDisk gets an existing GatewayCacheDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayCacheDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCacheDiskState, opts ...pulumi.ResourceOption) (*GatewayCacheDisk, error) {
	var resource GatewayCacheDisk
	err := ctx.ReadResource("alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayCacheDisk resources.
type gatewayCacheDiskState struct {
	// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
	CacheDiskCategory *string `pulumi:"cacheDiskCategory"`
	// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
	CacheDiskSizeInGb *int `pulumi:"cacheDiskSizeInGb"`
	// The ID of the cache.
	CacheId *string `pulumi:"cacheId"`
	// The ID of the gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// The cache disk inside the device name.
	LocalFilePath *string `pulumi:"localFilePath"`
	// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
	Status *int `pulumi:"status"`
}

type GatewayCacheDiskState struct {
	// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
	CacheDiskCategory pulumi.StringPtrInput
	// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
	CacheDiskSizeInGb pulumi.IntPtrInput
	// The ID of the cache.
	CacheId pulumi.StringPtrInput
	// The ID of the gateway.
	GatewayId pulumi.StringPtrInput
	// The cache disk inside the device name.
	LocalFilePath pulumi.StringPtrInput
	// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
	Status pulumi.IntPtrInput
}

func (GatewayCacheDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCacheDiskState)(nil)).Elem()
}

type gatewayCacheDiskArgs struct {
	// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
	CacheDiskCategory *string `pulumi:"cacheDiskCategory"`
	// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
	CacheDiskSizeInGb int `pulumi:"cacheDiskSizeInGb"`
	// The ID of the gateway.
	GatewayId string `pulumi:"gatewayId"`
}

// The set of arguments for constructing a GatewayCacheDisk resource.
type GatewayCacheDiskArgs struct {
	// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
	CacheDiskCategory pulumi.StringPtrInput
	// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
	CacheDiskSizeInGb pulumi.IntInput
	// The ID of the gateway.
	GatewayId pulumi.StringInput
}

func (GatewayCacheDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCacheDiskArgs)(nil)).Elem()
}

type GatewayCacheDiskInput interface {
	pulumi.Input

	ToGatewayCacheDiskOutput() GatewayCacheDiskOutput
	ToGatewayCacheDiskOutputWithContext(ctx context.Context) GatewayCacheDiskOutput
}

func (*GatewayCacheDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCacheDisk)(nil)).Elem()
}

func (i *GatewayCacheDisk) ToGatewayCacheDiskOutput() GatewayCacheDiskOutput {
	return i.ToGatewayCacheDiskOutputWithContext(context.Background())
}

func (i *GatewayCacheDisk) ToGatewayCacheDiskOutputWithContext(ctx context.Context) GatewayCacheDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCacheDiskOutput)
}

// GatewayCacheDiskArrayInput is an input type that accepts GatewayCacheDiskArray and GatewayCacheDiskArrayOutput values.
// You can construct a concrete instance of `GatewayCacheDiskArrayInput` via:
//
//	GatewayCacheDiskArray{ GatewayCacheDiskArgs{...} }
type GatewayCacheDiskArrayInput interface {
	pulumi.Input

	ToGatewayCacheDiskArrayOutput() GatewayCacheDiskArrayOutput
	ToGatewayCacheDiskArrayOutputWithContext(context.Context) GatewayCacheDiskArrayOutput
}

type GatewayCacheDiskArray []GatewayCacheDiskInput

func (GatewayCacheDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCacheDisk)(nil)).Elem()
}

func (i GatewayCacheDiskArray) ToGatewayCacheDiskArrayOutput() GatewayCacheDiskArrayOutput {
	return i.ToGatewayCacheDiskArrayOutputWithContext(context.Background())
}

func (i GatewayCacheDiskArray) ToGatewayCacheDiskArrayOutputWithContext(ctx context.Context) GatewayCacheDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCacheDiskArrayOutput)
}

// GatewayCacheDiskMapInput is an input type that accepts GatewayCacheDiskMap and GatewayCacheDiskMapOutput values.
// You can construct a concrete instance of `GatewayCacheDiskMapInput` via:
//
//	GatewayCacheDiskMap{ "key": GatewayCacheDiskArgs{...} }
type GatewayCacheDiskMapInput interface {
	pulumi.Input

	ToGatewayCacheDiskMapOutput() GatewayCacheDiskMapOutput
	ToGatewayCacheDiskMapOutputWithContext(context.Context) GatewayCacheDiskMapOutput
}

type GatewayCacheDiskMap map[string]GatewayCacheDiskInput

func (GatewayCacheDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCacheDisk)(nil)).Elem()
}

func (i GatewayCacheDiskMap) ToGatewayCacheDiskMapOutput() GatewayCacheDiskMapOutput {
	return i.ToGatewayCacheDiskMapOutputWithContext(context.Background())
}

func (i GatewayCacheDiskMap) ToGatewayCacheDiskMapOutputWithContext(ctx context.Context) GatewayCacheDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCacheDiskMapOutput)
}

type GatewayCacheDiskOutput struct{ *pulumi.OutputState }

func (GatewayCacheDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCacheDisk)(nil)).Elem()
}

func (o GatewayCacheDiskOutput) ToGatewayCacheDiskOutput() GatewayCacheDiskOutput {
	return o
}

func (o GatewayCacheDiskOutput) ToGatewayCacheDiskOutputWithContext(ctx context.Context) GatewayCacheDiskOutput {
	return o
}

// The cache disk type. Valid values: `cloudEfficiency`, `cloudSsd`.
func (o GatewayCacheDiskOutput) CacheDiskCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.StringOutput { return v.CacheDiskCategory }).(pulumi.StringOutput)
}

// size of the cache disk. Unit: `GB`. The upper limit of the basic gateway cache disk is `1` TB (`1024` GB), that of the standard gateway is `2` TB (`2048` GB), and that of other gateway cache disks is `32` TB (`32768` GB). The lower limit for the file gateway cache disk capacity is `40` GB, and the lower limit for the block gateway cache disk capacity is `20` GB.
func (o GatewayCacheDiskOutput) CacheDiskSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.IntOutput { return v.CacheDiskSizeInGb }).(pulumi.IntOutput)
}

// The ID of the cache.
func (o GatewayCacheDiskOutput) CacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.StringOutput { return v.CacheId }).(pulumi.StringOutput)
}

// The ID of the gateway.
func (o GatewayCacheDiskOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// The cache disk inside the device name.
func (o GatewayCacheDiskOutput) LocalFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.StringOutput { return v.LocalFilePath }).(pulumi.StringOutput)
}

// The status of the resource. Valid values: `0`, `1`, `2`. `0`: Normal. `1`: Is about to expire. `2`: Has expired.
func (o GatewayCacheDiskOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayCacheDisk) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type GatewayCacheDiskArrayOutput struct{ *pulumi.OutputState }

func (GatewayCacheDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCacheDisk)(nil)).Elem()
}

func (o GatewayCacheDiskArrayOutput) ToGatewayCacheDiskArrayOutput() GatewayCacheDiskArrayOutput {
	return o
}

func (o GatewayCacheDiskArrayOutput) ToGatewayCacheDiskArrayOutputWithContext(ctx context.Context) GatewayCacheDiskArrayOutput {
	return o
}

func (o GatewayCacheDiskArrayOutput) Index(i pulumi.IntInput) GatewayCacheDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayCacheDisk {
		return vs[0].([]*GatewayCacheDisk)[vs[1].(int)]
	}).(GatewayCacheDiskOutput)
}

type GatewayCacheDiskMapOutput struct{ *pulumi.OutputState }

func (GatewayCacheDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCacheDisk)(nil)).Elem()
}

func (o GatewayCacheDiskMapOutput) ToGatewayCacheDiskMapOutput() GatewayCacheDiskMapOutput {
	return o
}

func (o GatewayCacheDiskMapOutput) ToGatewayCacheDiskMapOutputWithContext(ctx context.Context) GatewayCacheDiskMapOutput {
	return o
}

func (o GatewayCacheDiskMapOutput) MapIndex(k pulumi.StringInput) GatewayCacheDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayCacheDisk {
		return vs[0].(map[string]*GatewayCacheDisk)[vs[1].(string)]
	}).(GatewayCacheDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCacheDiskInput)(nil)).Elem(), &GatewayCacheDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCacheDiskArrayInput)(nil)).Elem(), GatewayCacheDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCacheDiskMapInput)(nil)).Elem(), GatewayCacheDiskMap{})
	pulumi.RegisterOutputType(GatewayCacheDiskOutput{})
	pulumi.RegisterOutputType(GatewayCacheDiskArrayOutput{})
	pulumi.RegisterOutputType(GatewayCacheDiskMapOutput{})
}
