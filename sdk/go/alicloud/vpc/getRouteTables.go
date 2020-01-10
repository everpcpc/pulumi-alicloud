// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Route Tables owned by an Alibaba Cloud account.
// 
// > **NOTE:** Available in 1.36.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/route_tables.html.markdown.
func LookupRouteTables(ctx *pulumi.Context, args *GetRouteTablesArgs) (*GetRouteTablesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ids"] = args.Ids
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["resourceGroupId"] = args.ResourceGroupId
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	outputs, err := ctx.Invoke("alicloud:vpc/getRouteTables:getRouteTables", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRouteTablesResult{
		Ids: outputs["ids"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		ResourceGroupId: outputs["resourceGroupId"],
		Tables: outputs["tables"],
		Tags: outputs["tags"],
		VpcId: outputs["vpcId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRouteTables.
type GetRouteTablesArgs struct {
	// A list of Route Tables IDs.
	Ids interface{}
	// A regex string to filter route tables by name.
	NameRegex interface{}
	OutputFile interface{}
	// The Id of resource group which route tables belongs.
	ResourceGroupId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Vpc id of the route table.
	VpcId interface{}
}

// A collection of values returned by getRouteTables.
type GetRouteTablesResult struct {
	// (Optional) A list of Route Tables IDs.
	Ids interface{}
	NameRegex interface{}
	// A list of Route Tables names.
	Names interface{}
	OutputFile interface{}
	// The Id of resource group which route tables belongs.
	ResourceGroupId interface{}
	// A list of Route Tables. Each element contains the following attributes:
	Tables interface{}
	Tags interface{}
	VpcId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
