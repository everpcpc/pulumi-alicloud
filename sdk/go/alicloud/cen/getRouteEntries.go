// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cen

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides CEN Route Entries available to the user.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cen_route_entries.html.markdown.
func GetRouteEntries(ctx *pulumi.Context, args *GetRouteEntriesArgs, opts ...pulumi.InvokeOption) (*GetRouteEntriesResult, error) {
	var rv GetRouteEntriesResult
	err := ctx.Invoke("alicloud:cen/getRouteEntries:getRouteEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteEntries.
type GetRouteEntriesArgs struct {
	// The destination CIDR block of the route entry to query.
	CidrBlock *string `pulumi:"cidrBlock"`
	// ID of the CEN instance.
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// ID of the route table of the VPC or VBR.
	RouteTableId string `pulumi:"routeTableId"`
}


// A collection of values returned by getRouteEntries.
type GetRouteEntriesResult struct {
	// The destination CIDR block of the conflicted route entry.
	CidrBlock *string `pulumi:"cidrBlock"`
	// A list of CEN Route Entries. Each element contains the following attributes:
	Entries []GetRouteEntriesEntry `pulumi:"entries"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the CEN child instance.
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// ID of the route table.
	RouteTableId string `pulumi:"routeTableId"`
}

