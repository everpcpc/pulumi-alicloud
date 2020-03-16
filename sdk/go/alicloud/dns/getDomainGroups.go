// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dns

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetDomainGroups(ctx *pulumi.Context, args *GetDomainGroupsArgs, opts ...pulumi.InvokeOption) (*GetDomainGroupsResult, error) {
	var rv GetDomainGroupsResult
	err := ctx.Invoke("alicloud:dns/getDomainGroups:getDomainGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainGroups.
type GetDomainGroupsArgs struct {
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getDomainGroups.
type GetDomainGroupsResult struct {
	Groups []GetDomainGroupsGroup `pulumi:"groups"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
}

