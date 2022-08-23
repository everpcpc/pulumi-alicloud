// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Vpc Prefix Lists of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.182.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := vpc.GetPrefixLists(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcPrefixListId1", ids.Lists[0].Id)
//			nameRegex, err := vpc.GetPrefixLists(ctx, &vpc.GetPrefixListsArgs{
//				NameRegex: pulumi.StringRef("^my-PrefixList"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcPrefixListId2", nameRegex.Lists[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetPrefixLists(ctx *pulumi.Context, args *GetPrefixListsArgs, opts ...pulumi.InvokeOption) (*GetPrefixListsResult, error) {
	var rv GetPrefixListsResult
	err := ctx.Invoke("alicloud:vpc/getPrefixLists:getPrefixLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixLists.
type GetPrefixListsArgs struct {
	// Default to `true`. Set it to `false` can hide the `entrys` to output.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Prefix List IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Prefix List name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of the prefix list.
	PrefixListName *string `pulumi:"prefixListName"`
}

// A collection of values returned by getPrefixLists.
type GetPrefixListsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id             string               `pulumi:"id"`
	Ids            []string             `pulumi:"ids"`
	Lists          []GetPrefixListsList `pulumi:"lists"`
	NameRegex      *string              `pulumi:"nameRegex"`
	Names          []string             `pulumi:"names"`
	OutputFile     *string              `pulumi:"outputFile"`
	PrefixListName *string              `pulumi:"prefixListName"`
}

func GetPrefixListsOutput(ctx *pulumi.Context, args GetPrefixListsOutputArgs, opts ...pulumi.InvokeOption) GetPrefixListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPrefixListsResult, error) {
			args := v.(GetPrefixListsArgs)
			r, err := GetPrefixLists(ctx, &args, opts...)
			var s GetPrefixListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPrefixListsResultOutput)
}

// A collection of arguments for invoking getPrefixLists.
type GetPrefixListsOutputArgs struct {
	// Default to `true`. Set it to `false` can hide the `entrys` to output.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Prefix List IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Prefix List name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the prefix list.
	PrefixListName pulumi.StringPtrInput `pulumi:"prefixListName"`
}

func (GetPrefixListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixListsArgs)(nil)).Elem()
}

// A collection of values returned by getPrefixLists.
type GetPrefixListsResultOutput struct{ *pulumi.OutputState }

func (GetPrefixListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixListsResult)(nil)).Elem()
}

func (o GetPrefixListsResultOutput) ToGetPrefixListsResultOutput() GetPrefixListsResultOutput {
	return o
}

func (o GetPrefixListsResultOutput) ToGetPrefixListsResultOutputWithContext(ctx context.Context) GetPrefixListsResultOutput {
	return o
}

func (o GetPrefixListsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPrefixListsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPrefixListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixListsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPrefixListsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPrefixListsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPrefixListsResultOutput) Lists() GetPrefixListsListArrayOutput {
	return o.ApplyT(func(v GetPrefixListsResult) []GetPrefixListsList { return v.Lists }).(GetPrefixListsListArrayOutput)
}

func (o GetPrefixListsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixListsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetPrefixListsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPrefixListsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetPrefixListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPrefixListsResultOutput) PrefixListName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixListsResult) *string { return v.PrefixListName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPrefixListsResultOutput{})
}
