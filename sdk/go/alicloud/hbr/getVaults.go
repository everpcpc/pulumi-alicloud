// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Hbr Vaults of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.129.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "^my-Vault"
// 		ids, err := hbr.GetVaults(ctx, &hbr.GetVaultsArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("hbrVaultId1", ids.Vaults[0].Id)
// 		return nil
// 	})
// }
// ```
func GetVaults(ctx *pulumi.Context, args *GetVaultsArgs, opts ...pulumi.InvokeOption) (*GetVaultsResult, error) {
	var rv GetVaultsResult
	err := ctx.Invoke("alicloud:hbr/getVaults:getVaults", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVaults.
type GetVaultsArgs struct {
	// A list of Vault IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Vault name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
	Status *string `pulumi:"status"`
	// The type of Vault. Valid values: `STANDARD`.
	VaultType *string `pulumi:"vaultType"`
}

// A collection of values returned by getVaults.
type GetVaultsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string           `pulumi:"id"`
	Ids        []string         `pulumi:"ids"`
	NameRegex  *string          `pulumi:"nameRegex"`
	Names      []string         `pulumi:"names"`
	OutputFile *string          `pulumi:"outputFile"`
	Status     *string          `pulumi:"status"`
	VaultType  *string          `pulumi:"vaultType"`
	Vaults     []GetVaultsVault `pulumi:"vaults"`
}
