// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ens

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ens Key Pairs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.133.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ens"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "^my-KeyPair"
// 		nameRegex, err := ens.GetKeyPairs(ctx, &ens.GetKeyPairsArgs{
// 			Version:   "example_value",
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ensKeyPairId1", nameRegex.Pairs[0].Id)
// 		return nil
// 	})
// }
// ```
func GetKeyPairs(ctx *pulumi.Context, args *GetKeyPairsArgs, opts ...pulumi.InvokeOption) (*GetKeyPairsResult, error) {
	var rv GetKeyPairsResult
	err := ctx.Invoke("alicloud:ens/getKeyPairs:getKeyPairs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeyPairs.
type GetKeyPairsArgs struct {
	// The name of the key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// A regex string to filter results by Key Pair name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The version number.
	Version string `pulumi:"version"`
}

// A collection of values returned by getKeyPairs.
type GetKeyPairsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	Ids         []string          `pulumi:"ids"`
	KeyPairName *string           `pulumi:"keyPairName"`
	NameRegex   *string           `pulumi:"nameRegex"`
	Names       []string          `pulumi:"names"`
	OutputFile  *string           `pulumi:"outputFile"`
	Pairs       []GetKeyPairsPair `pulumi:"pairs"`
	Version     string            `pulumi:"version"`
}