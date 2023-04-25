// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Storage Gateway Express Sync Share Attachment resource.
//
// For information about Cloud Storage Gateway Express Sync Share Attachment and how to use it, see [What is Express Sync Share Attachment](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tftest"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			region := "cn-shanghai"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			defaultStocks, err := cloudstoragegateway.GetStocks(ctx, &cloudstoragegateway.GetStocksArgs{
//				GatewayClass: pulumi.StringRef("Standard"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("192.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       vpc.ID(),
//				CidrBlock:   pulumi.String("192.16.0.0/21"),
//				ZoneId:      *pulumi.String(defaultStocks.Stocks[0].ZoneId),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultStorageBundle, err := cloudstoragegateway.NewStorageBundle(ctx, "defaultStorageBundle", &cloudstoragegateway.StorageBundleArgs{
//				StorageBundleName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGateway, err := cloudstoragegateway.NewGateway(ctx, "defaultGateway", &cloudstoragegateway.GatewayArgs{
//				Description:            pulumi.String("tf-acctestDesalone"),
//				GatewayClass:           pulumi.String("Standard"),
//				Type:                   pulumi.String("File"),
//				PaymentType:            pulumi.String("PayAsYouGo"),
//				VswitchId:              defaultSwitch.ID(),
//				ReleaseAfterExpiration: pulumi.Bool(true),
//				PublicNetworkBandwidth: pulumi.Int(10),
//				StorageBundleId:        defaultStorageBundle.ID(),
//				Location:               pulumi.String("Cloud"),
//				GatewayName:            pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGatewayCacheDisk, err := cloudstoragegateway.NewGatewayCacheDisk(ctx, "defaultGatewayCacheDisk", &cloudstoragegateway.GatewayCacheDiskArgs{
//				CacheDiskCategory: pulumi.String("cloud_efficiency"),
//				GatewayId:         defaultGateway.ID(),
//				CacheDiskSizeInGb: pulumi.Int(50),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBucket, err := oss.NewBucket(ctx, "defaultBucket", &oss.BucketArgs{
//				Bucket: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGatewayFileShare, err := cloudstoragegateway.NewGatewayFileShare(ctx, "defaultGatewayFileShare", &cloudstoragegateway.GatewayFileShareArgs{
//				GatewayFileShareName: pulumi.String(name),
//				GatewayId:            defaultGateway.ID(),
//				LocalPath:            defaultGatewayCacheDisk.LocalFilePath,
//				OssBucketName:        defaultBucket.Bucket,
//				OssEndpoint:          defaultBucket.ExtranetEndpoint,
//				Protocol:             pulumi.String("NFS"),
//				RemoteSync:           pulumi.Bool(false),
//				FeLimit:              pulumi.Int(0),
//				BackendLimit:         pulumi.Int(0),
//				CacheMode:            pulumi.String("Cache"),
//				Squash:               pulumi.String("none"),
//				LagPeriod:            pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			defaultExpressSync, err := cloudstoragegateway.NewExpressSync(ctx, "defaultExpressSync", &cloudstoragegateway.ExpressSyncArgs{
//				BucketName:      defaultGatewayFileShare.OssBucketName,
//				BucketRegion:    pulumi.String(region),
//				Description:     pulumi.String(name),
//				ExpressSyncName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudstoragegateway.NewExpressSyncShareAttachment(ctx, "defaultExpressSyncShareAttachment", &cloudstoragegateway.ExpressSyncShareAttachmentArgs{
//				ExpressSyncId: defaultExpressSync.ID(),
//				GatewayId:     defaultGateway.ID(),
//				ShareName:     defaultGatewayFileShare.GatewayFileShareName,
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
// Cloud Storage Gateway Express Sync Share Attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudstoragegateway/expressSyncShareAttachment:ExpressSyncShareAttachment example <express_sync_id>:<gateway_id>:<share_name>
//
// ```
type ExpressSyncShareAttachment struct {
	pulumi.CustomResourceState

	// The ID of the ExpressSync.
	ExpressSyncId pulumi.StringOutput `pulumi:"expressSyncId"`
	// The ID of the Gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
	ShareName pulumi.StringOutput `pulumi:"shareName"`
}

// NewExpressSyncShareAttachment registers a new resource with the given unique name, arguments, and options.
func NewExpressSyncShareAttachment(ctx *pulumi.Context,
	name string, args *ExpressSyncShareAttachmentArgs, opts ...pulumi.ResourceOption) (*ExpressSyncShareAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressSyncId == nil {
		return nil, errors.New("invalid value for required argument 'ExpressSyncId'")
	}
	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	var resource ExpressSyncShareAttachment
	err := ctx.RegisterResource("alicloud:cloudstoragegateway/expressSyncShareAttachment:ExpressSyncShareAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressSyncShareAttachment gets an existing ExpressSyncShareAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressSyncShareAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressSyncShareAttachmentState, opts ...pulumi.ResourceOption) (*ExpressSyncShareAttachment, error) {
	var resource ExpressSyncShareAttachment
	err := ctx.ReadResource("alicloud:cloudstoragegateway/expressSyncShareAttachment:ExpressSyncShareAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressSyncShareAttachment resources.
type expressSyncShareAttachmentState struct {
	// The ID of the ExpressSync.
	ExpressSyncId *string `pulumi:"expressSyncId"`
	// The ID of the Gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
	ShareName *string `pulumi:"shareName"`
}

type ExpressSyncShareAttachmentState struct {
	// The ID of the ExpressSync.
	ExpressSyncId pulumi.StringPtrInput
	// The ID of the Gateway.
	GatewayId pulumi.StringPtrInput
	// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
	ShareName pulumi.StringPtrInput
}

func (ExpressSyncShareAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressSyncShareAttachmentState)(nil)).Elem()
}

type expressSyncShareAttachmentArgs struct {
	// The ID of the ExpressSync.
	ExpressSyncId string `pulumi:"expressSyncId"`
	// The ID of the Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
	ShareName string `pulumi:"shareName"`
}

// The set of arguments for constructing a ExpressSyncShareAttachment resource.
type ExpressSyncShareAttachmentArgs struct {
	// The ID of the ExpressSync.
	ExpressSyncId pulumi.StringInput
	// The ID of the Gateway.
	GatewayId pulumi.StringInput
	// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
	ShareName pulumi.StringInput
}

func (ExpressSyncShareAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressSyncShareAttachmentArgs)(nil)).Elem()
}

type ExpressSyncShareAttachmentInput interface {
	pulumi.Input

	ToExpressSyncShareAttachmentOutput() ExpressSyncShareAttachmentOutput
	ToExpressSyncShareAttachmentOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentOutput
}

func (*ExpressSyncShareAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressSyncShareAttachment)(nil)).Elem()
}

func (i *ExpressSyncShareAttachment) ToExpressSyncShareAttachmentOutput() ExpressSyncShareAttachmentOutput {
	return i.ToExpressSyncShareAttachmentOutputWithContext(context.Background())
}

func (i *ExpressSyncShareAttachment) ToExpressSyncShareAttachmentOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressSyncShareAttachmentOutput)
}

// ExpressSyncShareAttachmentArrayInput is an input type that accepts ExpressSyncShareAttachmentArray and ExpressSyncShareAttachmentArrayOutput values.
// You can construct a concrete instance of `ExpressSyncShareAttachmentArrayInput` via:
//
//	ExpressSyncShareAttachmentArray{ ExpressSyncShareAttachmentArgs{...} }
type ExpressSyncShareAttachmentArrayInput interface {
	pulumi.Input

	ToExpressSyncShareAttachmentArrayOutput() ExpressSyncShareAttachmentArrayOutput
	ToExpressSyncShareAttachmentArrayOutputWithContext(context.Context) ExpressSyncShareAttachmentArrayOutput
}

type ExpressSyncShareAttachmentArray []ExpressSyncShareAttachmentInput

func (ExpressSyncShareAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExpressSyncShareAttachment)(nil)).Elem()
}

func (i ExpressSyncShareAttachmentArray) ToExpressSyncShareAttachmentArrayOutput() ExpressSyncShareAttachmentArrayOutput {
	return i.ToExpressSyncShareAttachmentArrayOutputWithContext(context.Background())
}

func (i ExpressSyncShareAttachmentArray) ToExpressSyncShareAttachmentArrayOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressSyncShareAttachmentArrayOutput)
}

// ExpressSyncShareAttachmentMapInput is an input type that accepts ExpressSyncShareAttachmentMap and ExpressSyncShareAttachmentMapOutput values.
// You can construct a concrete instance of `ExpressSyncShareAttachmentMapInput` via:
//
//	ExpressSyncShareAttachmentMap{ "key": ExpressSyncShareAttachmentArgs{...} }
type ExpressSyncShareAttachmentMapInput interface {
	pulumi.Input

	ToExpressSyncShareAttachmentMapOutput() ExpressSyncShareAttachmentMapOutput
	ToExpressSyncShareAttachmentMapOutputWithContext(context.Context) ExpressSyncShareAttachmentMapOutput
}

type ExpressSyncShareAttachmentMap map[string]ExpressSyncShareAttachmentInput

func (ExpressSyncShareAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExpressSyncShareAttachment)(nil)).Elem()
}

func (i ExpressSyncShareAttachmentMap) ToExpressSyncShareAttachmentMapOutput() ExpressSyncShareAttachmentMapOutput {
	return i.ToExpressSyncShareAttachmentMapOutputWithContext(context.Background())
}

func (i ExpressSyncShareAttachmentMap) ToExpressSyncShareAttachmentMapOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressSyncShareAttachmentMapOutput)
}

type ExpressSyncShareAttachmentOutput struct{ *pulumi.OutputState }

func (ExpressSyncShareAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressSyncShareAttachment)(nil)).Elem()
}

func (o ExpressSyncShareAttachmentOutput) ToExpressSyncShareAttachmentOutput() ExpressSyncShareAttachmentOutput {
	return o
}

func (o ExpressSyncShareAttachmentOutput) ToExpressSyncShareAttachmentOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentOutput {
	return o
}

// The ID of the ExpressSync.
func (o ExpressSyncShareAttachmentOutput) ExpressSyncId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressSyncShareAttachment) pulumi.StringOutput { return v.ExpressSyncId }).(pulumi.StringOutput)
}

// The ID of the Gateway.
func (o ExpressSyncShareAttachmentOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressSyncShareAttachment) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// The name of the GatewayFileShare. **NOTE:** When GatewayFileShare is associated with a speed sync group, its reverse synchronization function will be turned off by default.
func (o ExpressSyncShareAttachmentOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressSyncShareAttachment) pulumi.StringOutput { return v.ShareName }).(pulumi.StringOutput)
}

type ExpressSyncShareAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ExpressSyncShareAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExpressSyncShareAttachment)(nil)).Elem()
}

func (o ExpressSyncShareAttachmentArrayOutput) ToExpressSyncShareAttachmentArrayOutput() ExpressSyncShareAttachmentArrayOutput {
	return o
}

func (o ExpressSyncShareAttachmentArrayOutput) ToExpressSyncShareAttachmentArrayOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentArrayOutput {
	return o
}

func (o ExpressSyncShareAttachmentArrayOutput) Index(i pulumi.IntInput) ExpressSyncShareAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExpressSyncShareAttachment {
		return vs[0].([]*ExpressSyncShareAttachment)[vs[1].(int)]
	}).(ExpressSyncShareAttachmentOutput)
}

type ExpressSyncShareAttachmentMapOutput struct{ *pulumi.OutputState }

func (ExpressSyncShareAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExpressSyncShareAttachment)(nil)).Elem()
}

func (o ExpressSyncShareAttachmentMapOutput) ToExpressSyncShareAttachmentMapOutput() ExpressSyncShareAttachmentMapOutput {
	return o
}

func (o ExpressSyncShareAttachmentMapOutput) ToExpressSyncShareAttachmentMapOutputWithContext(ctx context.Context) ExpressSyncShareAttachmentMapOutput {
	return o
}

func (o ExpressSyncShareAttachmentMapOutput) MapIndex(k pulumi.StringInput) ExpressSyncShareAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExpressSyncShareAttachment {
		return vs[0].(map[string]*ExpressSyncShareAttachment)[vs[1].(string)]
	}).(ExpressSyncShareAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressSyncShareAttachmentInput)(nil)).Elem(), &ExpressSyncShareAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressSyncShareAttachmentArrayInput)(nil)).Elem(), ExpressSyncShareAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressSyncShareAttachmentMapInput)(nil)).Elem(), ExpressSyncShareAttachmentMap{})
	pulumi.RegisterOutputType(ExpressSyncShareAttachmentOutput{})
	pulumi.RegisterOutputType(ExpressSyncShareAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ExpressSyncShareAttachmentMapOutput{})
}
