// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudconnect

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides Cloud Connect Networks available to the user.
// 
// > **NOTE:** Available in 1.59.0+
// 
// > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cloud_connect_networks.html.markdown.
func GetNetworks(ctx *pulumi.Context, args *GetNetworksArgs, opts ...pulumi.InvokeOption) (*GetNetworksResult, error) {
	var rv GetNetworksResult
	err := ctx.Invoke("alicloud:cloudconnect/getNetworks:getNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworks.
type GetNetworksArgs struct {
	// A list of CCN instances IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter CCN instances by name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getNetworks.
type GetNetworksResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CCN instances IDs.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of CCN instances names. 
	Names []string `pulumi:"names"`
	// A list of CCN instances. Each element contains the following attributes:
	Networks []GetNetworksNetwork `pulumi:"networks"`
	OutputFile *string `pulumi:"outputFile"`
}

