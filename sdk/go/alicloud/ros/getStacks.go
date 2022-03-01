// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ros

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ros"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ros.GetStacks(ctx, &ros.GetStacksArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: pulumi.StringRef("the_resource_name"),
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

func GetStacksOutput(ctx *pulumi.Context, args GetStacksOutputArgs, opts ...pulumi.InvokeOption) GetStacksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStacksResult, error) {
			args := v.(GetStacksArgs)
			r, err := GetStacks(ctx, &args, opts...)
			return *r, err
		}).(GetStacksResultOutput)
}

// A collection of arguments for invoking getStacks.
type GetStacksOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Stack IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Stack name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Parent Stack Id.
	ParentStackId pulumi.StringPtrInput `pulumi:"parentStackId"`
	// The show nested stack.
	ShowNestedStack pulumi.BoolPtrInput `pulumi:"showNestedStack"`
	// Stack Name.
	StackName pulumi.StringPtrInput `pulumi:"stackName"`
	// The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetStacksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStacksArgs)(nil)).Elem()
}

// A collection of values returned by getStacks.
type GetStacksResultOutput struct{ *pulumi.OutputState }

func (GetStacksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStacksResult)(nil)).Elem()
}

func (o GetStacksResultOutput) ToGetStacksResultOutput() GetStacksResultOutput {
	return o
}

func (o GetStacksResultOutput) ToGetStacksResultOutputWithContext(ctx context.Context) GetStacksResultOutput {
	return o
}

func (o GetStacksResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStacksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStacksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStacksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStacksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetStacksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetStacksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStacksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetStacksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetStacksResultOutput) ParentStackId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *string { return v.ParentStackId }).(pulumi.StringPtrOutput)
}

func (o GetStacksResultOutput) ShowNestedStack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *bool { return v.ShowNestedStack }).(pulumi.BoolPtrOutput)
}

func (o GetStacksResultOutput) StackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *string { return v.StackName }).(pulumi.StringPtrOutput)
}

func (o GetStacksResultOutput) Stacks() GetStacksStackArrayOutput {
	return o.ApplyT(func(v GetStacksResult) []GetStacksStack { return v.Stacks }).(GetStacksStackArrayOutput)
}

func (o GetStacksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStacksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetStacksResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetStacksResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStacksResultOutput{})
}
