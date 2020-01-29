// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
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
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:index/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
	AvailableDiskCategory *string `pulumi:"availableDiskCategory"`
	// Filter the results by a specific instance type.
	AvailableInstanceType *string `pulumi:"availableInstanceType"`
	// Filter the results by a specific resource type.
	// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
	AvailableResourceCreation *string `pulumi:"availableResourceCreation"`
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	// > **NOTE:** The disk category `cloud` has been outdated and can only be used by non-I/O Optimized ECS instances. Many availability zones don't support it. It is recommended to use `cloudEfficiency` or `cloudSsd`.
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType *string `pulumi:"availableSlbAddressType"`
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
	Multi *bool `pulumi:"multi"`
	// Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
	NetworkType *string `pulumi:"networkType"`
	OutputFile *string `pulumi:"outputFile"`
	// - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy *string `pulumi:"spotStrategy"`
}


// A collection of values returned by getZones.
type GetZonesResult struct {
	AvailableDiskCategory *string `pulumi:"availableDiskCategory"`
	AvailableInstanceType *string `pulumi:"availableInstanceType"`
	// Type of resources that can be created.
	AvailableResourceCreation *string `pulumi:"availableResourceCreation"`
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	AvailableSlbAddressType *string `pulumi:"availableSlbAddressType"`
	EnableDetails *bool `pulumi:"enableDetails"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone IDs.
	Ids []string `pulumi:"ids"`
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	Multi *bool `pulumi:"multi"`
	NetworkType *string `pulumi:"networkType"`
	OutputFile *string `pulumi:"outputFile"`
	SpotStrategy *string `pulumi:"spotStrategy"`
	// A list of availability zones. Each element contains the following attributes:
	Zones []GetZonesZone `pulumi:"zones"`
}

