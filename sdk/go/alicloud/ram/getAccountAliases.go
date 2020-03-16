// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ram

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides an alias for the Alibaba Cloud account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_account_aliases.html.markdown.
func GetAccountAliases(ctx *pulumi.Context, args *GetAccountAliasesArgs, opts ...pulumi.InvokeOption) (*GetAccountAliasesResult, error) {
	var rv GetAccountAliasesResult
	err := ctx.Invoke("alicloud:ram/getAccountAliases:getAccountAliases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountAliases.
type GetAccountAliasesArgs struct {
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getAccountAliases.
type GetAccountAliasesResult struct {
	// Alias of the account.
	AccountAlias string `pulumi:"accountAlias"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
}

