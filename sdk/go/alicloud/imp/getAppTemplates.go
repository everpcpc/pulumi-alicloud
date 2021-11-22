// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Imp App Templates of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.137.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/imp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := imp.GetAppTemplates(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("impAppTemplateId1", ids.Templates[0].Id)
// 		opt0 := "^my_AppTemplate"
// 		nameRegex, err := imp.GetAppTemplates(ctx, &imp.GetAppTemplatesArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("impAppTemplateId2", nameRegex.Templates[0].Id)
// 		return nil
// 	})
// }
// ```
func GetAppTemplates(ctx *pulumi.Context, args *GetAppTemplatesArgs, opts ...pulumi.InvokeOption) (*GetAppTemplatesResult, error) {
	var rv GetAppTemplatesResult
	err := ctx.Invoke("alicloud:imp/getAppTemplates:getAppTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppTemplates.
type GetAppTemplatesArgs struct {
	// A list of App Template IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by App Template name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Application template usage status.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAppTemplates.
type GetAppTemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                    `pulumi:"id"`
	Ids        []string                  `pulumi:"ids"`
	NameRegex  *string                   `pulumi:"nameRegex"`
	Names      []string                  `pulumi:"names"`
	OutputFile *string                   `pulumi:"outputFile"`
	Status     *string                   `pulumi:"status"`
	Templates  []GetAppTemplatesTemplate `pulumi:"templates"`
}

func GetAppTemplatesOutput(ctx *pulumi.Context, args GetAppTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetAppTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppTemplatesResult, error) {
			args := v.(GetAppTemplatesArgs)
			r, err := GetAppTemplates(ctx, &args, opts...)
			return *r, err
		}).(GetAppTemplatesResultOutput)
}

// A collection of arguments for invoking getAppTemplates.
type GetAppTemplatesOutputArgs struct {
	// A list of App Template IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by App Template name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Application template usage status.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetAppTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getAppTemplates.
type GetAppTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetAppTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppTemplatesResult)(nil)).Elem()
}

func (o GetAppTemplatesResultOutput) ToGetAppTemplatesResultOutput() GetAppTemplatesResultOutput {
	return o
}

func (o GetAppTemplatesResultOutput) ToGetAppTemplatesResultOutputWithContext(ctx context.Context) GetAppTemplatesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppTemplatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAppTemplatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAppTemplatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAppTemplatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAppTemplatesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetAppTemplatesResultOutput) Templates() GetAppTemplatesTemplateArrayOutput {
	return o.ApplyT(func(v GetAppTemplatesResult) []GetAppTemplatesTemplate { return v.Templates }).(GetAppTemplatesTemplateArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppTemplatesResultOutput{})
}
