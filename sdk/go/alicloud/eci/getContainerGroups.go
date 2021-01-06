// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Eci Container Groups of the current Alibaba Cloud user.
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
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/eci"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := eci.GetContainerGroups(ctx, &eci.GetContainerGroupsArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstEciContainerGroupId", example.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetContainerGroups(ctx *pulumi.Context, args *GetContainerGroupsArgs, opts ...pulumi.InvokeOption) (*GetContainerGroupsResult, error) {
	var rv GetContainerGroupsResult
	err := ctx.Invoke("alicloud:eci/getContainerGroups:getContainerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerGroups.
type GetContainerGroupsArgs struct {
	// The name of ContainerGroup.
	ContainerGroupName *string `pulumi:"containerGroupName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Container Group IDs.
	Ids []string `pulumi:"ids"`
	// The maximum number of resources returned in the response. Default value is `20`. Maximum value: `20`. The number of returned results is no greater than the specified number.
	Limit *int `pulumi:"limit"`
	// A regex string to filter results by Container Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group to which the container group belongs. If you have not specified a resource group for the container group, it is added to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of container.
	Status *string `pulumi:"status"`
	// The tags attached to the container group. Each tag is a key-value pair. You can attach up to 20 tags to a container group.
	Tags map[string]interface{} `pulumi:"tags"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
	WithEvent *bool   `pulumi:"withEvent"`
	// The IDs of the zones where the container groups are deployed. If this parameter is not set, the system automatically selects the zones. By default, no value is specified.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getContainerGroups.
type GetContainerGroupsResult struct {
	ContainerGroupName *string                   `pulumi:"containerGroupName"`
	EnableDetails      *bool                     `pulumi:"enableDetails"`
	Groups             []GetContainerGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	Limit           *int                   `pulumi:"limit"`
	NameRegex       *string                `pulumi:"nameRegex"`
	Names           []string               `pulumi:"names"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
	VswitchId       *string                `pulumi:"vswitchId"`
	WithEvent       *bool                  `pulumi:"withEvent"`
	ZoneId          *string                `pulumi:"zoneId"`
}