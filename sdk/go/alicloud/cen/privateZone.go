// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This topic describes how to configure PrivateZone access.
// PrivateZone is a VPC-based resolution and management service for private domain names.
// After you set a PrivateZone access, the Cloud Connect Network (CCN) and Virtual Border Router (VBR) attached to a CEN instance can access the PrivateZone service through CEN.
//
// For information about CEN Private Zone and how to use it, see [Manage CEN Private Zone](https://www.alibabacloud.com/help/en/doc-detail/106693.htm).
//
// > **NOTE:** Available in 1.83.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/cen"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultInstanceAttachment, err := cen.NewInstanceAttachment(ctx, "defaultInstanceAttachment", &cen.InstanceAttachmentArgs{
// 			InstanceId:            defaultInstance.ID(),
// 			ChildInstanceId:       defaultNetwork.ID(),
// 			ChildInstanceType:     pulumi.String("VPC"),
// 			ChildInstanceRegionId: pulumi.String("cn-hangzhou"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			defaultInstance,
// 			defaultNetwork,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cen.NewPrivateZone(ctx, "defaultPrivateZone", &cen.PrivateZoneArgs{
// 			AccessRegionId: pulumi.String("cn-hangzhou"),
// 			CenId:          defaultInstance.ID(),
// 			HostRegionId:   pulumi.String("cn-hangzhou"),
// 			HostVpcId:      defaultNetwork.ID(),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			defaultInstanceAttachment,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PrivateZone struct {
	pulumi.CustomResourceState

	// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
	AccessRegionId pulumi.StringOutput `pulumi:"accessRegionId"`
	// The ID of the CEN instance.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
	HostRegionId pulumi.StringOutput `pulumi:"hostRegionId"`
	// The VPC that belongs to the service region.
	HostVpcId pulumi.StringOutput `pulumi:"hostVpcId"`
	// The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewPrivateZone registers a new resource with the given unique name, arguments, and options.
func NewPrivateZone(ctx *pulumi.Context,
	name string, args *PrivateZoneArgs, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	if args == nil || args.AccessRegionId == nil {
		return nil, errors.New("missing required argument 'AccessRegionId'")
	}
	if args == nil || args.CenId == nil {
		return nil, errors.New("missing required argument 'CenId'")
	}
	if args == nil || args.HostRegionId == nil {
		return nil, errors.New("missing required argument 'HostRegionId'")
	}
	if args == nil || args.HostVpcId == nil {
		return nil, errors.New("missing required argument 'HostVpcId'")
	}
	if args == nil {
		args = &PrivateZoneArgs{}
	}
	var resource PrivateZone
	err := ctx.RegisterResource("alicloud:cen/privateZone:PrivateZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateZone gets an existing PrivateZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateZoneState, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	var resource PrivateZone
	err := ctx.ReadResource("alicloud:cen/privateZone:PrivateZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateZone resources.
type privateZoneState struct {
	// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
	AccessRegionId *string `pulumi:"accessRegionId"`
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
	HostRegionId *string `pulumi:"hostRegionId"`
	// The VPC that belongs to the service region.
	HostVpcId *string `pulumi:"hostVpcId"`
	// The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
	Status *string `pulumi:"status"`
}

type PrivateZoneState struct {
	// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
	AccessRegionId pulumi.StringPtrInput
	// The ID of the CEN instance.
	CenId pulumi.StringPtrInput
	// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
	HostRegionId pulumi.StringPtrInput
	// The VPC that belongs to the service region.
	HostVpcId pulumi.StringPtrInput
	// The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
	Status pulumi.StringPtrInput
}

func (PrivateZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneState)(nil)).Elem()
}

type privateZoneArgs struct {
	// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
	AccessRegionId string `pulumi:"accessRegionId"`
	// The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
	HostRegionId string `pulumi:"hostRegionId"`
	// The VPC that belongs to the service region.
	HostVpcId string `pulumi:"hostVpcId"`
}

// The set of arguments for constructing a PrivateZone resource.
type PrivateZoneArgs struct {
	// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
	AccessRegionId pulumi.StringInput
	// The ID of the CEN instance.
	CenId pulumi.StringInput
	// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
	HostRegionId pulumi.StringInput
	// The VPC that belongs to the service region.
	HostVpcId pulumi.StringInput
}

func (PrivateZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneArgs)(nil)).Elem()
}
