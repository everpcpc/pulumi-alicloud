// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cen

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides CEN instances available to the user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cen_instances.html.markdown.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:cen/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of CEN instances IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter CEN instances by name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CEN instances IDs.
	Ids []string `pulumi:"ids"`
	// A list of CEN instances. Each element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of CEN instances names. 
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
}

