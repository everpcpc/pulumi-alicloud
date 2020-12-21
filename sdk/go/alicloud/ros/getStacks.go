// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ros

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Ros Stacks of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.106.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ros"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "the_resource_name"
// 		example, err := ros.GetStacks(ctx, &ros.GetStacksArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstRosStackId", example.Stacks[0].Id)
// 		return nil
// 	})
// }
// ```
func GetStacks(ctx *pulumi.Context, args *GetStacksArgs, opts ...pulumi.InvokeOption) (*GetStacksResult, error) {
	var rv GetStacksResult
	err := ctx.Invoke("alicloud:ros/getStacks:getStacks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStacks.
type GetStacksArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Stack IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Stack name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Parent Stack Id.
	ParentStackId *string `pulumi:"parentStackId"`
	// The show nested stack.
	ShowNestedStack *bool `pulumi:"showNestedStack"`
	// Stack Name.
	StackName *string `pulumi:"stackName"`
	// The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
	Status *string `pulumi:"status"`
	// Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getStacks.
type GetStacksResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	NameRegex       *string                `pulumi:"nameRegex"`
	Names           []string               `pulumi:"names"`
	OutputFile      *string                `pulumi:"outputFile"`
	ParentStackId   *string                `pulumi:"parentStackId"`
	ShowNestedStack *bool                  `pulumi:"showNestedStack"`
	StackName       *string                `pulumi:"stackName"`
	Stacks          []GetStacksStack       `pulumi:"stacks"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
}
