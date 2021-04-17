// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides CEN VBR Health Checks available to the user.
//
// > **NOTE:** Available in 1.98.0+
func GetVbrHealthChecks(ctx *pulumi.Context, args *GetVbrHealthChecksArgs, opts ...pulumi.InvokeOption) (*GetVbrHealthChecksResult, error) {
	var rv GetVbrHealthChecksResult
	err := ctx.Invoke("alicloud:cen/getVbrHealthChecks:getVbrHealthChecks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVbrHealthChecks.
type GetVbrHealthChecksArgs struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId      *string `pulumi:"cenId"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the VBR instance.
	VbrInstanceId *string `pulumi:"vbrInstanceId"`
	// The User ID (UID) of the account to which the VBR instance belongs.
	VbrInstanceOwnerId *int `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region where the VBR instance is deployed.
	VbrInstanceRegionId string `pulumi:"vbrInstanceRegionId"`
}

// A collection of values returned by getVbrHealthChecks.
type GetVbrHealthChecksResult struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId *string `pulumi:"cenId"`
	// A list of CEN VBR Heath Checks. Each element contains the following attributes:
	Checks []GetVbrHealthChecksCheck `pulumi:"checks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the CEN VBR Heath Check IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the VBR instance.
	VbrInstanceId      *string `pulumi:"vbrInstanceId"`
	VbrInstanceOwnerId *int    `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region where the VBR instance is deployed.
	VbrInstanceRegionId string `pulumi:"vbrInstanceRegionId"`
}
