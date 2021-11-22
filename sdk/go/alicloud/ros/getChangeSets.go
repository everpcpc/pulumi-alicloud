// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ros

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ros Change Sets of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.105.0+.
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
// 		opt0 := "the_resource_name"
// 		example, err := ros.GetChangeSets(ctx, &ros.GetChangeSetsArgs{
// 			StackId: "example_value",
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstRosChangeSetId", example.Sets[0].Id)
// 		return nil
// 	})
// }
// ```
func GetChangeSets(ctx *pulumi.Context, args *GetChangeSetsArgs, opts ...pulumi.InvokeOption) (*GetChangeSetsResult, error) {
	var rv GetChangeSetsResult
	err := ctx.Invoke("alicloud:ros/getChangeSets:getChangeSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getChangeSets.
type GetChangeSetsArgs struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName *string `pulumi:"changeSetName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Change Set IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Change Set name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId string `pulumi:"stackId"`
	// The status of the change set.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getChangeSets.
type GetChangeSetsResult struct {
	ChangeSetName *string `pulumi:"changeSetName"`
	EnableDetails *bool   `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string             `pulumi:"id"`
	Ids        []string           `pulumi:"ids"`
	NameRegex  *string            `pulumi:"nameRegex"`
	Names      []string           `pulumi:"names"`
	OutputFile *string            `pulumi:"outputFile"`
	Sets       []GetChangeSetsSet `pulumi:"sets"`
	StackId    string             `pulumi:"stackId"`
	Status     *string            `pulumi:"status"`
}

func GetChangeSetsOutput(ctx *pulumi.Context, args GetChangeSetsOutputArgs, opts ...pulumi.InvokeOption) GetChangeSetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetChangeSetsResult, error) {
			args := v.(GetChangeSetsArgs)
			r, err := GetChangeSets(ctx, &args, opts...)
			return *r, err
		}).(GetChangeSetsResultOutput)
}

// A collection of arguments for invoking getChangeSets.
type GetChangeSetsOutputArgs struct {
	// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	ChangeSetName pulumi.StringPtrInput `pulumi:"changeSetName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Change Set IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Change Set name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
	StackId pulumi.StringInput `pulumi:"stackId"`
	// The status of the change set.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetChangeSetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChangeSetsArgs)(nil)).Elem()
}

// A collection of values returned by getChangeSets.
type GetChangeSetsResultOutput struct{ *pulumi.OutputState }

func (GetChangeSetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChangeSetsResult)(nil)).Elem()
}

func (o GetChangeSetsResultOutput) ToGetChangeSetsResultOutput() GetChangeSetsResultOutput {
	return o
}

func (o GetChangeSetsResultOutput) ToGetChangeSetsResultOutputWithContext(ctx context.Context) GetChangeSetsResultOutput {
	return o
}

func (o GetChangeSetsResultOutput) ChangeSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChangeSetsResult) *string { return v.ChangeSetName }).(pulumi.StringPtrOutput)
}

func (o GetChangeSetsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetChangeSetsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetChangeSetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetChangeSetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetChangeSetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetChangeSetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetChangeSetsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChangeSetsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetChangeSetsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetChangeSetsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetChangeSetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChangeSetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetChangeSetsResultOutput) Sets() GetChangeSetsSetArrayOutput {
	return o.ApplyT(func(v GetChangeSetsResult) []GetChangeSetsSet { return v.Sets }).(GetChangeSetsSetArrayOutput)
}

func (o GetChangeSetsResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v GetChangeSetsResult) string { return v.StackId }).(pulumi.StringOutput)
}

func (o GetChangeSetsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChangeSetsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChangeSetsResultOutput{})
}
