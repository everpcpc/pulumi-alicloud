// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Common Bandwidth Packages owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.36.0+.
//
// ## Public ip addresses Block
//   
//   The public ip addresses mapping supports the following:
//   
//   * `ipAddress`   - The address of the EIP.
//   * `allocationId` - The ID of the EIP instance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/common_bandwidth_packages.html.markdown.
func GetCommonBandwidthPackages(ctx *pulumi.Context, args *GetCommonBandwidthPackagesArgs, opts ...pulumi.InvokeOption) (*GetCommonBandwidthPackagesResult, error) {
	var rv GetCommonBandwidthPackagesResult
	err := ctx.Invoke("alicloud:vpc/getCommonBandwidthPackages:getCommonBandwidthPackages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCommonBandwidthPackages.
type GetCommonBandwidthPackagesArgs struct {
	// A list of Common Bandwidth Packages IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}


// A collection of values returned by getCommonBandwidthPackages.
type GetCommonBandwidthPackagesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of Common Bandwidth Packages IDs.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of Common Bandwidth Packages names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of Common Bandwidth Packages. Each element contains the following attributes:
	Packages []GetCommonBandwidthPackagesPackage `pulumi:"packages"`
	// The Id of resource group which the common bandwidth package belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

