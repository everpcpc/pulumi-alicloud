// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Adb DBClusters of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.121.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/adb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example"
// 		example, err := adb.GetDBClusters(ctx, &adb.GetDBClustersArgs{
// 			DescriptionRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstAdbDbClusterId", example.Clusters[0].Id)
// 		return nil
// 	})
// }
// ```
func GetDBClusters(ctx *pulumi.Context, args *GetDBClustersArgs, opts ...pulumi.InvokeOption) (*GetDBClustersResult, error) {
	var rv GetDBClustersResult
	err := ctx.Invoke("alicloud:adb/getDBClusters:getDBClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDBClusters.
type GetDBClustersArgs struct {
	// The description of DBCluster.
	Description *string `pulumi:"description"`
	// A regex string to filter results by DBCluster description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of DBCluster IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The tag of the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getDBClusters.
type GetDBClustersResult struct {
	Clusters         []GetDBClustersCluster `pulumi:"clusters"`
	Description      *string                `pulumi:"description"`
	DescriptionRegex *string                `pulumi:"descriptionRegex"`
	Descriptions     []string               `pulumi:"descriptions"`
	EnableDetails    *bool                  `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
}