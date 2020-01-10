// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides availability zones that can be accessed by an Alibaba Cloud account within the region configured in the provider.
// 
// 
// > **NOTE:** If one zone is sold out, it will not be exported.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/zones.html.markdown.
func LookupZones(ctx *pulumi.Context, args *GetZonesArgs) (*GetZonesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["availableDiskCategory"] = args.AvailableDiskCategory
		inputs["availableInstanceType"] = args.AvailableInstanceType
		inputs["availableResourceCreation"] = args.AvailableResourceCreation
		inputs["availableSlbAddressIpVersion"] = args.AvailableSlbAddressIpVersion
		inputs["availableSlbAddressType"] = args.AvailableSlbAddressType
		inputs["enableDetails"] = args.EnableDetails
		inputs["instanceChargeType"] = args.InstanceChargeType
		inputs["multi"] = args.Multi
		inputs["networkType"] = args.NetworkType
		inputs["outputFile"] = args.OutputFile
		inputs["spotStrategy"] = args.SpotStrategy
	}
	outputs, err := ctx.Invoke("alicloud:index/getZones:getZones", inputs)
	if err != nil {
		return nil, err
	}
	return &GetZonesResult{
		AvailableDiskCategory: outputs["availableDiskCategory"],
		AvailableInstanceType: outputs["availableInstanceType"],
		AvailableResourceCreation: outputs["availableResourceCreation"],
		AvailableSlbAddressIpVersion: outputs["availableSlbAddressIpVersion"],
		AvailableSlbAddressType: outputs["availableSlbAddressType"],
		EnableDetails: outputs["enableDetails"],
		Ids: outputs["ids"],
		InstanceChargeType: outputs["instanceChargeType"],
		Multi: outputs["multi"],
		NetworkType: outputs["networkType"],
		OutputFile: outputs["outputFile"],
		SpotStrategy: outputs["spotStrategy"],
		Zones: outputs["zones"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
	AvailableDiskCategory interface{}
	// Filter the results by a specific instance type.
	AvailableInstanceType interface{}
	// Filter the results by a specific resource type.
	// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
	AvailableResourceCreation interface{}
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	// > **NOTE:** The disk category `cloud` has been outdated and can only be used by non-I/O Optimized ECS instances. Many availability zones don't support it. It is recommended to use `cloudEfficiency` or `cloudSsd`.
	AvailableSlbAddressIpVersion interface{}
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType interface{}
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails interface{}
	// Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType interface{}
	// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
	Multi interface{}
	// Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
	NetworkType interface{}
	OutputFile interface{}
	// - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy interface{}
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	AvailableDiskCategory interface{}
	AvailableInstanceType interface{}
	// Type of resources that can be created.
	AvailableResourceCreation interface{}
	AvailableSlbAddressIpVersion interface{}
	AvailableSlbAddressType interface{}
	EnableDetails interface{}
	// A list of zone IDs.
	Ids interface{}
	InstanceChargeType interface{}
	Multi interface{}
	NetworkType interface{}
	OutputFile interface{}
	SpotStrategy interface{}
	// A list of availability zones. Each element contains the following attributes:
	Zones interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
