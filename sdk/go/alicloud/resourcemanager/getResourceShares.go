// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Resource Manager Resource Shares of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.111.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "the_resource_name"
// 		example, err := resourcemanager.GetResourceShares(ctx, &resourcemanager.GetResourceSharesArgs{
// 			ResourceShareOwner: "Self",
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstResourceManagerResourceShareId", example.Shares[0].Id)
// 		return nil
// 	})
// }
// ```
func GetResourceShares(ctx *pulumi.Context, args *GetResourceSharesArgs, opts ...pulumi.InvokeOption) (*GetResourceSharesResult, error) {
	var rv GetResourceSharesResult
	err := ctx.Invoke("alicloud:resourcemanager/getResourceShares:getResourceShares", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceShares.
type GetResourceSharesArgs struct {
	// A list of Resource Share IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Resource Share name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of resource share.
	ResourceShareName *string `pulumi:"resourceShareName"`
	// The owner of resource share.
	ResourceShareOwner string `pulumi:"resourceShareOwner"`
	// The status of resource share.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getResourceShares.
type GetResourceSharesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                   `pulumi:"id"`
	Ids                []string                 `pulumi:"ids"`
	NameRegex          *string                  `pulumi:"nameRegex"`
	Names              []string                 `pulumi:"names"`
	OutputFile         *string                  `pulumi:"outputFile"`
	ResourceShareName  *string                  `pulumi:"resourceShareName"`
	ResourceShareOwner string                   `pulumi:"resourceShareOwner"`
	Shares             []GetResourceSharesShare `pulumi:"shares"`
	Status             *string                  `pulumi:"status"`
}