// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides available scaling configuration resources. 
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scaling_configurations.html.markdown.
func LookupScalingConfigurations(ctx *pulumi.Context, args *GetScalingConfigurationsArgs) (*GetScalingConfigurationsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ids"] = args.Ids
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["scalingGroupId"] = args.ScalingGroupId
	}
	outputs, err := ctx.Invoke("alicloud:ess/getScalingConfigurations:getScalingConfigurations", inputs)
	if err != nil {
		return nil, err
	}
	return &GetScalingConfigurationsResult{
		Configurations: outputs["configurations"],
		Ids: outputs["ids"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		ScalingGroupId: outputs["scalingGroupId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getScalingConfigurations.
type GetScalingConfigurationsArgs struct {
	// A list of scaling configuration IDs.
	Ids interface{}
	// A regex string to filter resulting scaling configurations by name.
	NameRegex interface{}
	OutputFile interface{}
	// Scaling group id the scaling configurations belong to.
	ScalingGroupId interface{}
}

// A collection of values returned by getScalingConfigurations.
type GetScalingConfigurationsResult struct {
	// A list of scaling rules. Each element contains the following attributes:
	Configurations interface{}
	// A list of scaling configuration ids.
	Ids interface{}
	NameRegex interface{}
	// A list of scaling configuration names.
	Names interface{}
	OutputFile interface{}
	// ID of the scaling group.
	ScalingGroupId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
