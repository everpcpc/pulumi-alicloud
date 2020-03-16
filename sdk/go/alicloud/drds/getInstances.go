// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package drds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

//  The `drds.Instance` data source provides a collection of DRDS instances available in Alibaba Cloud account.
// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
//
// > **NOTE:** Available in 1.35.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/drds_instances.html.markdown.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:drds/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of DRDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// A list of DRDS descriptions. 
	Descriptions []string `pulumi:"descriptions"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of DRDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of DRDS instances.
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

