// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a ECS auto provisioning group resource which is a solution that uses preemptive instances and payAsYouGo instances to rapidly deploy clusters.
//
// > **NOTE:** Available in 1.79.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "auto_provisioning_group"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := "cloud_efficiency"
// 		opt1 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableDiskCategory:     &opt0,
// 			AvailableResourceCreation: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/24"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
// 			VpcId: defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt2 := "^ubuntu_18.*64"
// 		opt3 := true
// 		opt4 := "system"
// 		defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
// 			NameRegex:  &opt2,
// 			MostRecent: &opt3,
// 			Owners:     &opt4,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		template, err := ecs.NewLaunchTemplate(ctx, "template", &ecs.LaunchTemplateArgs{
// 			ImageId:         pulumi.String(defaultImages.Images[0].Id),
// 			InstanceType:    pulumi.String("ecs.n1.tiny"),
// 			SecurityGroupId: defaultSecurityGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewAutoProvisioningGroup(ctx, "defaultAutoProvisioningGroup", &ecs.AutoProvisioningGroupArgs{
// 			LaunchTemplateId:         template.ID(),
// 			TotalTargetCapacity:      pulumi.String("4"),
// 			PayAsYouGoTargetCapacity: pulumi.String("1"),
// 			SpotTargetCapacity:       pulumi.String("2"),
// 			LaunchTemplateConfigs: ecs.AutoProvisioningGroupLaunchTemplateConfigArray{
// 				&ecs.AutoProvisioningGroupLaunchTemplateConfigArgs{
// 					InstanceType:     pulumi.String("ecs.n1.small"),
// 					VswitchId:        defaultSwitch.ID(),
// 					WeightedCapacity: pulumi.String("2"),
// 					MaxPrice:         pulumi.String("2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Block config
//
// The config mapping supports the following:
// * `instanceType` - (Optional) The instance type of the Nth extended configurations of the launch template.
// * `maxPrice` - (Required) The maximum price of the instance type specified in the Nth extended configurations of the launch template.
// * `vswitchId` - (Required) The ID of the VSwitch in the Nth extended configurations of the launch template.
// * `weightedCapacity` - (Required) The weight of the instance type specified in the Nth extended configurations of the launch template.
// * `priority` - (Optional) The priority of the instance type specified in the Nth extended configurations of the launch template. A value of 0 indicates the highest priority.
//
// ## Import
//
// ECS auto provisioning group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup example asg-abc123456
// ```
type AutoProvisioningGroup struct {
	pulumi.CustomResourceState

	// The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
	AutoProvisioningGroupName pulumi.StringOutput `pulumi:"autoProvisioningGroupName"`
	// The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
	AutoProvisioningGroupType pulumi.StringPtrOutput `pulumi:"autoProvisioningGroupType"`
	// The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
	DefaultTargetCapacityType pulumi.StringPtrOutput `pulumi:"defaultTargetCapacityType"`
	// The description of the auto provisioning group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrOutput `pulumi:"excessCapacityTerminationPolicy"`
	// DataDisk mappings to attach to ecs instance. See Block config below for details.
	LaunchTemplateConfigs AutoProvisioningGroupLaunchTemplateConfigArrayOutput `pulumi:"launchTemplateConfigs"`
	// The ID of the instance launch template associated with the auto provisioning group.
	LaunchTemplateId pulumi.StringOutput `pulumi:"launchTemplateId"`
	// The version of the instance launch template associated with the auto provisioning group.
	LaunchTemplateVersion pulumi.StringOutput `pulumi:"launchTemplateVersion"`
	// The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
	MaxSpotPrice pulumi.Float64Output `pulumi:"maxSpotPrice"`
	// The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
	PayAsYouGoAllocationStrategy pulumi.StringPtrOutput `pulumi:"payAsYouGoAllocationStrategy"`
	// The target capacity of pay-as-you-go instances in the auto provisioning group.
	PayAsYouGoTargetCapacity pulumi.StringPtrOutput `pulumi:"payAsYouGoTargetCapacity"`
	// The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
	SpotAllocationStrategy pulumi.StringPtrOutput `pulumi:"spotAllocationStrategy"`
	// The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
	SpotInstanceInterruptionBehavior pulumi.StringPtrOutput `pulumi:"spotInstanceInterruptionBehavior"`
	// This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
	SpotInstancePoolsToUseCount pulumi.IntOutput `pulumi:"spotInstancePoolsToUseCount"`
	// The target capacity of preemptible instances in the auto provisioning group.
	SpotTargetCapacity pulumi.StringPtrOutput `pulumi:"spotTargetCapacity"`
	// Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
	TerminateInstances pulumi.BoolPtrOutput `pulumi:"terminateInstances"`
	// The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrOutput `pulumi:"terminateInstancesWithExpiration"`
	// The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
	TotalTargetCapacity pulumi.StringOutput `pulumi:"totalTargetCapacity"`
	// The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
	ValidFrom pulumi.StringOutput `pulumi:"validFrom"`
	// The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
	ValidUntil pulumi.StringOutput `pulumi:"validUntil"`
}

// NewAutoProvisioningGroup registers a new resource with the given unique name, arguments, and options.
func NewAutoProvisioningGroup(ctx *pulumi.Context,
	name string, args *AutoProvisioningGroupArgs, opts ...pulumi.ResourceOption) (*AutoProvisioningGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LaunchTemplateConfigs == nil {
		return nil, errors.New("invalid value for required argument 'LaunchTemplateConfigs'")
	}
	if args.LaunchTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'LaunchTemplateId'")
	}
	if args.TotalTargetCapacity == nil {
		return nil, errors.New("invalid value for required argument 'TotalTargetCapacity'")
	}
	var resource AutoProvisioningGroup
	err := ctx.RegisterResource("alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoProvisioningGroup gets an existing AutoProvisioningGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoProvisioningGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoProvisioningGroupState, opts ...pulumi.ResourceOption) (*AutoProvisioningGroup, error) {
	var resource AutoProvisioningGroup
	err := ctx.ReadResource("alicloud:ecs/autoProvisioningGroup:AutoProvisioningGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoProvisioningGroup resources.
type autoProvisioningGroupState struct {
	// The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
	AutoProvisioningGroupName *string `pulumi:"autoProvisioningGroupName"`
	// The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
	AutoProvisioningGroupType *string `pulumi:"autoProvisioningGroupType"`
	// The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
	DefaultTargetCapacityType *string `pulumi:"defaultTargetCapacityType"`
	// The description of the auto provisioning group.
	Description *string `pulumi:"description"`
	// The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
	ExcessCapacityTerminationPolicy *string `pulumi:"excessCapacityTerminationPolicy"`
	// DataDisk mappings to attach to ecs instance. See Block config below for details.
	LaunchTemplateConfigs []AutoProvisioningGroupLaunchTemplateConfig `pulumi:"launchTemplateConfigs"`
	// The ID of the instance launch template associated with the auto provisioning group.
	LaunchTemplateId *string `pulumi:"launchTemplateId"`
	// The version of the instance launch template associated with the auto provisioning group.
	LaunchTemplateVersion *string `pulumi:"launchTemplateVersion"`
	// The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
	MaxSpotPrice *float64 `pulumi:"maxSpotPrice"`
	// The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
	PayAsYouGoAllocationStrategy *string `pulumi:"payAsYouGoAllocationStrategy"`
	// The target capacity of pay-as-you-go instances in the auto provisioning group.
	PayAsYouGoTargetCapacity *string `pulumi:"payAsYouGoTargetCapacity"`
	// The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
	SpotAllocationStrategy *string `pulumi:"spotAllocationStrategy"`
	// The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
	SpotInstanceInterruptionBehavior *string `pulumi:"spotInstanceInterruptionBehavior"`
	// This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
	SpotInstancePoolsToUseCount *int `pulumi:"spotInstancePoolsToUseCount"`
	// The target capacity of preemptible instances in the auto provisioning group.
	SpotTargetCapacity *string `pulumi:"spotTargetCapacity"`
	// Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
	TerminateInstances *bool `pulumi:"terminateInstances"`
	// The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
	TerminateInstancesWithExpiration *bool `pulumi:"terminateInstancesWithExpiration"`
	// The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
	TotalTargetCapacity *string `pulumi:"totalTargetCapacity"`
	// The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
	ValidFrom *string `pulumi:"validFrom"`
	// The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
	ValidUntil *string `pulumi:"validUntil"`
}

type AutoProvisioningGroupState struct {
	// The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
	AutoProvisioningGroupName pulumi.StringPtrInput
	// The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
	AutoProvisioningGroupType pulumi.StringPtrInput
	// The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
	DefaultTargetCapacityType pulumi.StringPtrInput
	// The description of the auto provisioning group.
	Description pulumi.StringPtrInput
	// The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrInput
	// DataDisk mappings to attach to ecs instance. See Block config below for details.
	LaunchTemplateConfigs AutoProvisioningGroupLaunchTemplateConfigArrayInput
	// The ID of the instance launch template associated with the auto provisioning group.
	LaunchTemplateId pulumi.StringPtrInput
	// The version of the instance launch template associated with the auto provisioning group.
	LaunchTemplateVersion pulumi.StringPtrInput
	// The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
	MaxSpotPrice pulumi.Float64PtrInput
	// The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
	PayAsYouGoAllocationStrategy pulumi.StringPtrInput
	// The target capacity of pay-as-you-go instances in the auto provisioning group.
	PayAsYouGoTargetCapacity pulumi.StringPtrInput
	// The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
	SpotAllocationStrategy pulumi.StringPtrInput
	// The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
	SpotInstanceInterruptionBehavior pulumi.StringPtrInput
	// This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
	SpotInstancePoolsToUseCount pulumi.IntPtrInput
	// The target capacity of preemptible instances in the auto provisioning group.
	SpotTargetCapacity pulumi.StringPtrInput
	// Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
	TerminateInstances pulumi.BoolPtrInput
	// The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrInput
	// The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
	TotalTargetCapacity pulumi.StringPtrInput
	// The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
	ValidFrom pulumi.StringPtrInput
	// The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
	ValidUntil pulumi.StringPtrInput
}

func (AutoProvisioningGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoProvisioningGroupState)(nil)).Elem()
}

type autoProvisioningGroupArgs struct {
	// The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
	AutoProvisioningGroupName *string `pulumi:"autoProvisioningGroupName"`
	// The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
	AutoProvisioningGroupType *string `pulumi:"autoProvisioningGroupType"`
	// The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
	DefaultTargetCapacityType *string `pulumi:"defaultTargetCapacityType"`
	// The description of the auto provisioning group.
	Description *string `pulumi:"description"`
	// The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
	ExcessCapacityTerminationPolicy *string `pulumi:"excessCapacityTerminationPolicy"`
	// DataDisk mappings to attach to ecs instance. See Block config below for details.
	LaunchTemplateConfigs []AutoProvisioningGroupLaunchTemplateConfig `pulumi:"launchTemplateConfigs"`
	// The ID of the instance launch template associated with the auto provisioning group.
	LaunchTemplateId string `pulumi:"launchTemplateId"`
	// The version of the instance launch template associated with the auto provisioning group.
	LaunchTemplateVersion *string `pulumi:"launchTemplateVersion"`
	// The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
	MaxSpotPrice *float64 `pulumi:"maxSpotPrice"`
	// The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
	PayAsYouGoAllocationStrategy *string `pulumi:"payAsYouGoAllocationStrategy"`
	// The target capacity of pay-as-you-go instances in the auto provisioning group.
	PayAsYouGoTargetCapacity *string `pulumi:"payAsYouGoTargetCapacity"`
	// The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
	SpotAllocationStrategy *string `pulumi:"spotAllocationStrategy"`
	// The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
	SpotInstanceInterruptionBehavior *string `pulumi:"spotInstanceInterruptionBehavior"`
	// This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
	SpotInstancePoolsToUseCount *int `pulumi:"spotInstancePoolsToUseCount"`
	// The target capacity of preemptible instances in the auto provisioning group.
	SpotTargetCapacity *string `pulumi:"spotTargetCapacity"`
	// Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
	TerminateInstances *bool `pulumi:"terminateInstances"`
	// The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
	TerminateInstancesWithExpiration *bool `pulumi:"terminateInstancesWithExpiration"`
	// The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
	TotalTargetCapacity string `pulumi:"totalTargetCapacity"`
	// The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
	ValidFrom *string `pulumi:"validFrom"`
	// The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
	ValidUntil *string `pulumi:"validUntil"`
}

// The set of arguments for constructing a AutoProvisioningGroup resource.
type AutoProvisioningGroupArgs struct {
	// The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)
	AutoProvisioningGroupName pulumi.StringPtrInput
	// The type of the auto provisioning group. Valid values:`request` and `maintain`,Default value: `maintain`.
	AutoProvisioningGroupType pulumi.StringPtrInput
	// The type of supplemental instances. When the total value of `PayAsYouGoTargetCapacity` and `SpotTargetCapacity` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:`PayAsYouGo`: Pay-as-you-go instances; `Spot`: Preemptible instances, Default value: `Spot`.
	DefaultTargetCapacityType pulumi.StringPtrInput
	// The description of the auto provisioning group.
	Description pulumi.StringPtrInput
	// The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: `no-termination` and `termination`,Default value: `no-termination`.
	ExcessCapacityTerminationPolicy pulumi.StringPtrInput
	// DataDisk mappings to attach to ecs instance. See Block config below for details.
	LaunchTemplateConfigs AutoProvisioningGroupLaunchTemplateConfigArrayInput
	// The ID of the instance launch template associated with the auto provisioning group.
	LaunchTemplateId pulumi.StringInput
	// The version of the instance launch template associated with the auto provisioning group.
	LaunchTemplateVersion pulumi.StringPtrInput
	// The global maximum price for preemptible instances in the auto provisioning group. If both the `MaxSpotPrice` and `LaunchTemplateConfig.N.MaxPrice` parameters are specified, the maximum price is the lower value of the two.
	MaxSpotPrice pulumi.Float64PtrInput
	// The scale-out policy for pay-as-you-go instances. Valid values: `lowest-price` and `prioritized`,Default value: `lowest-price`.
	PayAsYouGoAllocationStrategy pulumi.StringPtrInput
	// The target capacity of pay-as-you-go instances in the auto provisioning group.
	PayAsYouGoTargetCapacity pulumi.StringPtrInput
	// The scale-out policy for preemptible instances. Valid values:`lowest-price` and `diversified`,Default value: `lowest-price`.
	SpotAllocationStrategy pulumi.StringPtrInput
	// The default behavior after preemptible instances are shut down. Value values: `stop` and `terminate`,Default value: `stop`.
	SpotInstanceInterruptionBehavior pulumi.StringPtrInput
	// This parameter takes effect when the `SpotAllocationStrategy` parameter is set to `lowest-price`. The auto provisioning group selects instance types of the lowest cost to create instances.
	SpotInstancePoolsToUseCount pulumi.IntPtrInput
	// The target capacity of preemptible instances in the auto provisioning group.
	SpotTargetCapacity pulumi.StringPtrInput
	// Specifies whether to release instances of the auto provisioning group. Valid values:`false` and `true`, default value: `false`.
	TerminateInstances pulumi.BoolPtrInput
	// The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: `false` and `true`, default value: `false`.
	TerminateInstancesWithExpiration pulumi.BoolPtrInput
	// The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.
	TotalTargetCapacity pulumi.StringInput
	// The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the `validUntil` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.
	ValidFrom pulumi.StringPtrInput
	// The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the `validFrom` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.
	ValidUntil pulumi.StringPtrInput
}

func (AutoProvisioningGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoProvisioningGroupArgs)(nil)).Elem()
}

type AutoProvisioningGroupInput interface {
	pulumi.Input

	ToAutoProvisioningGroupOutput() AutoProvisioningGroupOutput
	ToAutoProvisioningGroupOutputWithContext(ctx context.Context) AutoProvisioningGroupOutput
}

func (*AutoProvisioningGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvisioningGroup)(nil))
}

func (i *AutoProvisioningGroup) ToAutoProvisioningGroupOutput() AutoProvisioningGroupOutput {
	return i.ToAutoProvisioningGroupOutputWithContext(context.Background())
}

func (i *AutoProvisioningGroup) ToAutoProvisioningGroupOutputWithContext(ctx context.Context) AutoProvisioningGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoProvisioningGroupOutput)
}

func (i *AutoProvisioningGroup) ToAutoProvisioningGroupPtrOutput() AutoProvisioningGroupPtrOutput {
	return i.ToAutoProvisioningGroupPtrOutputWithContext(context.Background())
}

func (i *AutoProvisioningGroup) ToAutoProvisioningGroupPtrOutputWithContext(ctx context.Context) AutoProvisioningGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoProvisioningGroupPtrOutput)
}

type AutoProvisioningGroupPtrInput interface {
	pulumi.Input

	ToAutoProvisioningGroupPtrOutput() AutoProvisioningGroupPtrOutput
	ToAutoProvisioningGroupPtrOutputWithContext(ctx context.Context) AutoProvisioningGroupPtrOutput
}

type autoProvisioningGroupPtrType AutoProvisioningGroupArgs

func (*autoProvisioningGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoProvisioningGroup)(nil))
}

func (i *autoProvisioningGroupPtrType) ToAutoProvisioningGroupPtrOutput() AutoProvisioningGroupPtrOutput {
	return i.ToAutoProvisioningGroupPtrOutputWithContext(context.Background())
}

func (i *autoProvisioningGroupPtrType) ToAutoProvisioningGroupPtrOutputWithContext(ctx context.Context) AutoProvisioningGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoProvisioningGroupPtrOutput)
}

// AutoProvisioningGroupArrayInput is an input type that accepts AutoProvisioningGroupArray and AutoProvisioningGroupArrayOutput values.
// You can construct a concrete instance of `AutoProvisioningGroupArrayInput` via:
//
//          AutoProvisioningGroupArray{ AutoProvisioningGroupArgs{...} }
type AutoProvisioningGroupArrayInput interface {
	pulumi.Input

	ToAutoProvisioningGroupArrayOutput() AutoProvisioningGroupArrayOutput
	ToAutoProvisioningGroupArrayOutputWithContext(context.Context) AutoProvisioningGroupArrayOutput
}

type AutoProvisioningGroupArray []AutoProvisioningGroupInput

func (AutoProvisioningGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AutoProvisioningGroup)(nil))
}

func (i AutoProvisioningGroupArray) ToAutoProvisioningGroupArrayOutput() AutoProvisioningGroupArrayOutput {
	return i.ToAutoProvisioningGroupArrayOutputWithContext(context.Background())
}

func (i AutoProvisioningGroupArray) ToAutoProvisioningGroupArrayOutputWithContext(ctx context.Context) AutoProvisioningGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoProvisioningGroupArrayOutput)
}

// AutoProvisioningGroupMapInput is an input type that accepts AutoProvisioningGroupMap and AutoProvisioningGroupMapOutput values.
// You can construct a concrete instance of `AutoProvisioningGroupMapInput` via:
//
//          AutoProvisioningGroupMap{ "key": AutoProvisioningGroupArgs{...} }
type AutoProvisioningGroupMapInput interface {
	pulumi.Input

	ToAutoProvisioningGroupMapOutput() AutoProvisioningGroupMapOutput
	ToAutoProvisioningGroupMapOutputWithContext(context.Context) AutoProvisioningGroupMapOutput
}

type AutoProvisioningGroupMap map[string]AutoProvisioningGroupInput

func (AutoProvisioningGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AutoProvisioningGroup)(nil))
}

func (i AutoProvisioningGroupMap) ToAutoProvisioningGroupMapOutput() AutoProvisioningGroupMapOutput {
	return i.ToAutoProvisioningGroupMapOutputWithContext(context.Background())
}

func (i AutoProvisioningGroupMap) ToAutoProvisioningGroupMapOutputWithContext(ctx context.Context) AutoProvisioningGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoProvisioningGroupMapOutput)
}

type AutoProvisioningGroupOutput struct {
	*pulumi.OutputState
}

func (AutoProvisioningGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvisioningGroup)(nil))
}

func (o AutoProvisioningGroupOutput) ToAutoProvisioningGroupOutput() AutoProvisioningGroupOutput {
	return o
}

func (o AutoProvisioningGroupOutput) ToAutoProvisioningGroupOutputWithContext(ctx context.Context) AutoProvisioningGroupOutput {
	return o
}

func (o AutoProvisioningGroupOutput) ToAutoProvisioningGroupPtrOutput() AutoProvisioningGroupPtrOutput {
	return o.ToAutoProvisioningGroupPtrOutputWithContext(context.Background())
}

func (o AutoProvisioningGroupOutput) ToAutoProvisioningGroupPtrOutputWithContext(ctx context.Context) AutoProvisioningGroupPtrOutput {
	return o.ApplyT(func(v AutoProvisioningGroup) *AutoProvisioningGroup {
		return &v
	}).(AutoProvisioningGroupPtrOutput)
}

type AutoProvisioningGroupPtrOutput struct {
	*pulumi.OutputState
}

func (AutoProvisioningGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoProvisioningGroup)(nil))
}

func (o AutoProvisioningGroupPtrOutput) ToAutoProvisioningGroupPtrOutput() AutoProvisioningGroupPtrOutput {
	return o
}

func (o AutoProvisioningGroupPtrOutput) ToAutoProvisioningGroupPtrOutputWithContext(ctx context.Context) AutoProvisioningGroupPtrOutput {
	return o
}

type AutoProvisioningGroupArrayOutput struct{ *pulumi.OutputState }

func (AutoProvisioningGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoProvisioningGroup)(nil))
}

func (o AutoProvisioningGroupArrayOutput) ToAutoProvisioningGroupArrayOutput() AutoProvisioningGroupArrayOutput {
	return o
}

func (o AutoProvisioningGroupArrayOutput) ToAutoProvisioningGroupArrayOutputWithContext(ctx context.Context) AutoProvisioningGroupArrayOutput {
	return o
}

func (o AutoProvisioningGroupArrayOutput) Index(i pulumi.IntInput) AutoProvisioningGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoProvisioningGroup {
		return vs[0].([]AutoProvisioningGroup)[vs[1].(int)]
	}).(AutoProvisioningGroupOutput)
}

type AutoProvisioningGroupMapOutput struct{ *pulumi.OutputState }

func (AutoProvisioningGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AutoProvisioningGroup)(nil))
}

func (o AutoProvisioningGroupMapOutput) ToAutoProvisioningGroupMapOutput() AutoProvisioningGroupMapOutput {
	return o
}

func (o AutoProvisioningGroupMapOutput) ToAutoProvisioningGroupMapOutputWithContext(ctx context.Context) AutoProvisioningGroupMapOutput {
	return o
}

func (o AutoProvisioningGroupMapOutput) MapIndex(k pulumi.StringInput) AutoProvisioningGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AutoProvisioningGroup {
		return vs[0].(map[string]AutoProvisioningGroup)[vs[1].(string)]
	}).(AutoProvisioningGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoProvisioningGroupOutput{})
	pulumi.RegisterOutputType(AutoProvisioningGroupPtrOutput{})
	pulumi.RegisterOutputType(AutoProvisioningGroupArrayOutput{})
	pulumi.RegisterOutputType(AutoProvisioningGroupMapOutput{})
}
