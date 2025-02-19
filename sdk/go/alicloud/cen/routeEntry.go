// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CEN route entry resource. Cloud Enterprise Network (CEN) supports publishing and withdrawing route entries of attached networks. You can publish a route entry of an attached VPC or VBR to a CEN instance, then other attached networks can learn the route if there is no route conflict. You can withdraw a published route entry when CEN does not need it any more.
//
// For information about CEN route entries publishment and how to use it, see [Manage network routes](https://www.alibabacloud.com/help/doc-detail/86980.htm).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alicloud.NewProvider(ctx, "hz", &alicloud.ProviderArgs{
//				Region: pulumi.String("cn-hangzhou"),
//			})
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			name := "tf-testAccCenRouteEntryConfig"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(defaultZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       vpc.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/21"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				Description: pulumi.String("foo"),
//				VpcId:       vpc.ID(),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := ecs.NewInstance(ctx, "defaultInstance", &ecs.InstanceArgs{
//				VswitchId:               defaultSwitch.ID(),
//				ImageId:                 *pulumi.String(defaultImages.Images[0].Id),
//				InstanceType:            *pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				SystemDiskCategory:      pulumi.String("cloud_efficiency"),
//				InternetChargeType:      pulumi.String("PayByTraffic"),
//				InternetMaxBandwidthOut: pulumi.Int(5),
//				SecurityGroups: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//				InstanceName: pulumi.String(name),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			cen, err := cen.NewInstance(ctx, "cen", nil)
//			if err != nil {
//				return err
//			}
//			attach, err := cen.NewInstanceAttachment(ctx, "attach", &cen.InstanceAttachmentArgs{
//				InstanceId:            cen.ID(),
//				ChildInstanceId:       vpc.ID(),
//				ChildInstanceType:     pulumi.String("VPC"),
//				ChildInstanceRegionId: pulumi.String("cn-hangzhou"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				defaultSwitch,
//			}))
//			if err != nil {
//				return err
//			}
//			route, err := vpc.NewRouteEntry(ctx, "route", &vpc.RouteEntryArgs{
//				RouteTableId:         vpc.RouteTableId,
//				DestinationCidrblock: pulumi.String("11.0.0.0/16"),
//				NexthopType:          pulumi.String("Instance"),
//				NexthopId:            defaultInstance.ID(),
//			}, pulumi.Provider(alicloud.Hz))
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewRouteEntry(ctx, "foo", &cen.RouteEntryArgs{
//				InstanceId:   cen.ID(),
//				RouteTableId: vpc.RouteTableId,
//				CidrBlock:    route.DestinationCidrblock,
//			}, pulumi.Provider(alicloud.Hz), pulumi.DependsOn([]pulumi.Resource{
//				attach,
//			}))
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
// CEN instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/routeEntry:RouteEntry example cen-abc123456:vtb-abc123:192.168.0.0/24
//
// ```
type RouteEntry struct {
	pulumi.CustomResourceState

	// The destination CIDR block of the route entry to publish.
	//
	// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
	//
	// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the CEN.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The route table of the attached VBR or VPC.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	var resource RouteEntry
	err := ctx.RegisterResource("alicloud:cen/routeEntry:RouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteEntryState, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	var resource RouteEntry
	err := ctx.ReadResource("alicloud:cen/routeEntry:RouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteEntry resources.
type routeEntryState struct {
	// The destination CIDR block of the route entry to publish.
	//
	// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
	//
	// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the CEN.
	InstanceId *string `pulumi:"instanceId"`
	// The route table of the attached VBR or VPC.
	RouteTableId *string `pulumi:"routeTableId"`
}

type RouteEntryState struct {
	// The destination CIDR block of the route entry to publish.
	//
	// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
	//
	// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
	CidrBlock pulumi.StringPtrInput
	// The ID of the CEN.
	InstanceId pulumi.StringPtrInput
	// The route table of the attached VBR or VPC.
	RouteTableId pulumi.StringPtrInput
}

func (RouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryState)(nil)).Elem()
}

type routeEntryArgs struct {
	// The destination CIDR block of the route entry to publish.
	//
	// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
	//
	// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
	CidrBlock string `pulumi:"cidrBlock"`
	// The ID of the CEN.
	InstanceId string `pulumi:"instanceId"`
	// The route table of the attached VBR or VPC.
	RouteTableId string `pulumi:"routeTableId"`
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// The destination CIDR block of the route entry to publish.
	//
	// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
	//
	// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
	CidrBlock pulumi.StringInput
	// The ID of the CEN.
	InstanceId pulumi.StringInput
	// The route table of the attached VBR or VPC.
	RouteTableId pulumi.StringInput
}

func (RouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryArgs)(nil)).Elem()
}

type RouteEntryInput interface {
	pulumi.Input

	ToRouteEntryOutput() RouteEntryOutput
	ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput
}

func (*RouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (i *RouteEntry) ToRouteEntryOutput() RouteEntryOutput {
	return i.ToRouteEntryOutputWithContext(context.Background())
}

func (i *RouteEntry) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryOutput)
}

// RouteEntryArrayInput is an input type that accepts RouteEntryArray and RouteEntryArrayOutput values.
// You can construct a concrete instance of `RouteEntryArrayInput` via:
//
//	RouteEntryArray{ RouteEntryArgs{...} }
type RouteEntryArrayInput interface {
	pulumi.Input

	ToRouteEntryArrayOutput() RouteEntryArrayOutput
	ToRouteEntryArrayOutputWithContext(context.Context) RouteEntryArrayOutput
}

type RouteEntryArray []RouteEntryInput

func (RouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryArray) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return i.ToRouteEntryArrayOutputWithContext(context.Background())
}

func (i RouteEntryArray) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryArrayOutput)
}

// RouteEntryMapInput is an input type that accepts RouteEntryMap and RouteEntryMapOutput values.
// You can construct a concrete instance of `RouteEntryMapInput` via:
//
//	RouteEntryMap{ "key": RouteEntryArgs{...} }
type RouteEntryMapInput interface {
	pulumi.Input

	ToRouteEntryMapOutput() RouteEntryMapOutput
	ToRouteEntryMapOutputWithContext(context.Context) RouteEntryMapOutput
}

type RouteEntryMap map[string]RouteEntryInput

func (RouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryMap) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return i.ToRouteEntryMapOutputWithContext(context.Background())
}

func (i RouteEntryMap) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryMapOutput)
}

type RouteEntryOutput struct{ *pulumi.OutputState }

func (RouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (o RouteEntryOutput) ToRouteEntryOutput() RouteEntryOutput {
	return o
}

func (o RouteEntryOutput) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return o
}

// The destination CIDR block of the route entry to publish.
//
// ->**NOTE:** The "alicloudCenInstanceRouteEntries" resource depends on the related "cen.InstanceAttachment" resource.
//
// ->**NOTE:** The "cen.InstanceAttachment" resource should depend on the related "vpc.Switch" resource.
func (o RouteEntryOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The ID of the CEN.
func (o RouteEntryOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The route table of the attached VBR or VPC.
func (o RouteEntryOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

type RouteEntryArrayOutput struct{ *pulumi.OutputState }

func (RouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) Index(i pulumi.IntInput) RouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].([]*RouteEntry)[vs[1].(int)]
	}).(RouteEntryOutput)
}

type RouteEntryMapOutput struct{ *pulumi.OutputState }

func (RouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) MapIndex(k pulumi.StringInput) RouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].(map[string]*RouteEntry)[vs[1].(string)]
	}).(RouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryInput)(nil)).Elem(), &RouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryArrayInput)(nil)).Elem(), RouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryMapInput)(nil)).Elem(), RouteEntryMap{})
	pulumi.RegisterOutputType(RouteEntryOutput{})
	pulumi.RegisterOutputType(RouteEntryArrayOutput{})
	pulumi.RegisterOutputType(RouteEntryMapOutput{})
}
