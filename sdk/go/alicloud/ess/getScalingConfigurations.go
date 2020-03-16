// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ess

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides available scaling configuration resources. 
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scaling_configurations.html.markdown.
func GetScalingConfigurations(ctx *pulumi.Context, args *GetScalingConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetScalingConfigurationsResult, error) {
	var rv GetScalingConfigurationsResult
	err := ctx.Invoke("alicloud:ess/getScalingConfigurations:getScalingConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingConfigurations.
type GetScalingConfigurationsArgs struct {
	// A list of scaling configuration IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting scaling configurations by name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Scaling group id the scaling configurations belong to.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}


// A collection of values returned by getScalingConfigurations.
type GetScalingConfigurationsResult struct {
	// A list of scaling rules. Each element contains the following attributes:
	Configurations []GetScalingConfigurationsConfiguration `pulumi:"configurations"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of scaling configuration ids.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of scaling configuration names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// ID of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

