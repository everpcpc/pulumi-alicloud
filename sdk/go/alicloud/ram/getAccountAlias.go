// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func LookupAccountAlias(ctx *pulumi.Context, args *GetAccountAliasArgs) (*GetAccountAliasResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["outputFile"] = args.OutputFile
	}
	outputs, err := ctx.Invoke("alicloud:ram/getAccountAlias:getAccountAlias", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccountAliasResult{
		AccountAlias: outputs["accountAlias"],
		OutputFile: outputs["outputFile"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccountAlias.
type GetAccountAliasArgs struct {
	OutputFile interface{}
}

// A collection of values returned by getAccountAlias.
type GetAccountAliasResult struct {
	AccountAlias interface{}
	OutputFile interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
