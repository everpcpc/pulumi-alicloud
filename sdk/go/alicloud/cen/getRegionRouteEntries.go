// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cen

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides CEN Regional Route Entries available to the user.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cen_region_route_entries.html.markdown.
func GetRegionRouteEntries(ctx *pulumi.Context, args *GetRegionRouteEntriesArgs, opts ...pulumi.InvokeOption) (*GetRegionRouteEntriesResult, error) {
	var rv GetRegionRouteEntriesResult
	err := ctx.Invoke("alicloud:cen/getRegionRouteEntries:getRegionRouteEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionRouteEntries.
type GetRegionRouteEntriesArgs struct {
	// ID of the CEN instance.
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// ID of the region.
	RegionId string `pulumi:"regionId"`
}


// A collection of values returned by getRegionRouteEntries.
type GetRegionRouteEntriesResult struct {
	// A list of CEN Route Entries. Each element contains the following attributes:
	Entries []GetRegionRouteEntriesEntry `pulumi:"entries"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	RegionId string `pulumi:"regionId"`
}

