// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pvtz

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/pvtz_zones.html.markdown.
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:pvtz/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// A list of zone IDs. 
	Ids []string `pulumi:"ids"`
	// keyword for zone name.
	Keyword *string `pulumi:"keyword"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getZones.
type GetZonesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone IDs. 
	Ids []string `pulumi:"ids"`
	Keyword *string `pulumi:"keyword"`
	// A list of zone names. 
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of zones. Each element contains the following attributes:
	Zones []GetZonesZone `pulumi:"zones"`
}

