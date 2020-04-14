// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the ECS instance types of Alibaba Cloud.
//
// > **NOTE:** By default, only the upgraded instance types are returned. If you want to get outdated instance types, you must set `isOutdated` to true.
//
// > **NOTE:** If one instance type is sold out, it will not be exported.
func GetInstanceTypes(ctx *pulumi.Context, args *GetInstanceTypesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypesResult, error) {
	var rv GetInstanceTypesResult
	err := ctx.Invoke("alicloud:ecs/getInstanceTypes:getInstanceTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesArgs struct {
	// The zone where instance types are supported.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Filter the results to a specific number of cpu cores.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// Filter the result whose network interface number is no more than `eniAmount`.
	EniAmount *int `pulumi:"eniAmount"`
	// The GPU amount of an instance type.
	GpuAmount *int `pulumi:"gpuAmount"`
	// The GPU spec of an instance type.
	GpuSpec *string `pulumi:"gpuSpec"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Filter the results based on their family name. For example: 'ecs.n4'.
	InstanceTypeFamily *string `pulumi:"instanceTypeFamily"`
	// If true, outdated instance types are included in the results. Default to false.
	IsOutdated         *bool   `pulumi:"isOutdated"`
	KubernetesNodeRole *string `pulumi:"kubernetesNodeRole"`
	// Filter the results to a specific memory size in GB.
	MemorySize *float64 `pulumi:"memorySize"`
	// Filter the results by network type. Valid values: `Classic` and `Vpc`.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	SortedBy    *string `pulumi:"sortedBy"`
	// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy *string `pulumi:"spotStrategy"`
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Number of CPU cores.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// The maximum number of network interfaces that an instance type can be attached to.
	EniAmount *int    `pulumi:"eniAmount"`
	GpuAmount *int    `pulumi:"gpuAmount"`
	GpuSpec   *string `pulumi:"gpuSpec"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance type IDs.
	Ids                []string `pulumi:"ids"`
	InstanceChargeType *string  `pulumi:"instanceChargeType"`
	InstanceTypeFamily *string  `pulumi:"instanceTypeFamily"`
	// A list of image types. Each element contains the following attributes:
	InstanceTypes      []GetInstanceTypesInstanceType `pulumi:"instanceTypes"`
	IsOutdated         *bool                          `pulumi:"isOutdated"`
	KubernetesNodeRole *string                        `pulumi:"kubernetesNodeRole"`
	// Size of memory, measured in GB.
	MemorySize   *float64 `pulumi:"memorySize"`
	NetworkType  *string  `pulumi:"networkType"`
	OutputFile   *string  `pulumi:"outputFile"`
	SortedBy     *string  `pulumi:"sortedBy"`
	SpotStrategy *string  `pulumi:"spotStrategy"`
}
