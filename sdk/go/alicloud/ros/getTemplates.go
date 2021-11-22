// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ros

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ros Templates of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.108.0+.
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
// 		example, err := ros.GetTemplates(ctx, &ros.GetTemplatesArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstRosTemplateId", example.Templates[0].Id)
// 		return nil
// 	})
// }
// ```
func GetTemplates(ctx *pulumi.Context, args *GetTemplatesArgs, opts ...pulumi.InvokeOption) (*GetTemplatesResult, error) {
	var rv GetTemplatesResult
	err := ctx.Invoke("alicloud:ros/getTemplates:getTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Template IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Template name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Share Type.
	ShareType *string `pulumi:"shareType"`
	// Tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	TemplateName *string `pulumi:"templateName"`
}

// A collection of values returned by getTemplates.
type GetTemplatesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id           string                 `pulumi:"id"`
	Ids          []string               `pulumi:"ids"`
	NameRegex    *string                `pulumi:"nameRegex"`
	Names        []string               `pulumi:"names"`
	OutputFile   *string                `pulumi:"outputFile"`
	ShareType    *string                `pulumi:"shareType"`
	Tags         map[string]interface{} `pulumi:"tags"`
	TemplateName *string                `pulumi:"templateName"`
	Templates    []GetTemplatesTemplate `pulumi:"templates"`
}

func GetTemplatesOutput(ctx *pulumi.Context, args GetTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplatesResult, error) {
			args := v.(GetTemplatesArgs)
			r, err := GetTemplates(ctx, &args, opts...)
			return *r, err
		}).(GetTemplatesResultOutput)
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Template IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Template name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Share Type.
	ShareType pulumi.StringPtrInput `pulumi:"shareType"`
	// Tags.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The name of the template.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	TemplateName pulumi.StringPtrInput `pulumi:"templateName"`
}

func (GetTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getTemplates.
type GetTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesResult)(nil)).Elem()
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutput() GetTemplatesResultOutput {
	return o
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutputWithContext(ctx context.Context) GetTemplatesResultOutput {
	return o
}

func (o GetTemplatesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTemplatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTemplatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTemplatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetTemplatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTemplatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetTemplatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTemplatesResultOutput) ShareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.ShareType }).(pulumi.StringPtrOutput)
}

func (o GetTemplatesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetTemplatesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetTemplatesResultOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.TemplateName }).(pulumi.StringPtrOutput)
}

func (o GetTemplatesResultOutput) Templates() GetTemplatesTemplateArrayOutput {
	return o.ApplyT(func(v GetTemplatesResult) []GetTemplatesTemplate { return v.Templates }).(GetTemplatesTemplateArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplatesResultOutput{})
}
