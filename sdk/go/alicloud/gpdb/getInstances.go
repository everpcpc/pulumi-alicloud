// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `gpdb.getInstances` data source provides a collection of AnalyticDB for PostgreSQL instances available in Alicloud account.
// Filters support regular expression for the instance name or availability_zone.
// 
// > **NOTE:**  Available in 1.47.0+
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/gpdb_instances.html.markdown.
func LookupInstances(ctx *pulumi.Context, args *GetInstancesArgs) (*GetInstancesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["ids"] = args.Ids
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["tags"] = args.Tags
		inputs["vswitchId"] = args.VswitchId
	}
	outputs, err := ctx.Invoke("alicloud:gpdb/getInstances:getInstances", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstancesResult{
		AvailabilityZone: outputs["availabilityZone"],
		Ids: outputs["ids"],
		Instances: outputs["instances"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Tags: outputs["tags"],
		VswitchId: outputs["vswitchId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Instance availability zone.
	AvailabilityZone interface{}
	// A list of instance IDs.
	Ids interface{}
	// A regex string to apply to the instance name.
	NameRegex interface{}
	OutputFile interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId interface{}
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// Instance availability zone.
	AvailabilityZone interface{}
	// The ids list of AnalyticDB for PostgreSQL instances.
	Ids interface{}
	// A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:
	Instances interface{}
	NameRegex interface{}
	// The names list of AnalyticDB for PostgreSQL instance.
	Names interface{}
	OutputFile interface{}
	Tags interface{}
	VswitchId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
