// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of OOS Templates in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in v1.92.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oos"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := false
// 		opt1 := "test"
// 		opt2 := "Private"
// 		example, err := oos.GetTemplates(ctx, &oos.GetTemplatesArgs{
// 			HasTrigger: &opt0,
// 			NameRegex:  &opt1,
// 			ShareType:  &opt2,
// 			Tags: map[string]interface{}{
// 				"Created": "TF",
// 				"For":     "template Test",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstTemplateName", example.Templates[0].TemplateName)
// 		return nil
// 	})
// }
// ```
func GetTemplates(ctx *pulumi.Context, args *GetTemplatesArgs, opts ...pulumi.InvokeOption) (*GetTemplatesResult, error) {
	var rv GetTemplatesResult
	err := ctx.Invoke("alicloud:oos/getTemplates:getTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesArgs struct {
	// The category of template.
	Category *string `pulumi:"category"`
	// The creator of the template.
	CreatedBy *string `pulumi:"createdBy"`
	// The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.
	CreatedDate *string `pulumi:"createdDate"`
	// Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.
	CreatedDateAfter *string `pulumi:"createdDateAfter"`
	// Is it triggered successfully.
	HasTrigger *bool `pulumi:"hasTrigger"`
	// A list of OOS Template ids. Each element in the list is same as template_name.
	Ids []string `pulumi:"ids"`
	// A regex string to filter the results by the template_name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The sharing type of the template. Valid values: `Private`, `Public`.
	ShareType *string `pulumi:"shareType"`
	// Sort field. Valid values: `TotalExecutionCount`, `Popularity`, `TemplateName` and `CreatedDate`. Default to `TotalExecutionCount`.
	SortField *string `pulumi:"sortField"`
	// Sort order. Valid values: `Ascending`, `Descending`. Default to `Descending`
	SortOrder *string `pulumi:"sortOrder"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The format of the template. Valid values: `JSON`, `YAML`.
	TemplateFormat *string `pulumi:"templateFormat"`
	// The type of OOS Template.
	TemplateType *string `pulumi:"templateType"`
}

// A collection of values returned by getTemplates.
type GetTemplatesResult struct {
	Category         *string `pulumi:"category"`
	CreatedBy        *string `pulumi:"createdBy"`
	CreatedDate      *string `pulumi:"createdDate"`
	CreatedDateAfter *string `pulumi:"createdDateAfter"`
	HasTrigger       *bool   `pulumi:"hasTrigger"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of OOS Template ids. Each element in the list is same as template_name.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// (Available in v1.114.0+) A list of OOS Template names.
	Names          []string               `pulumi:"names"`
	OutputFile     *string                `pulumi:"outputFile"`
	ShareType      *string                `pulumi:"shareType"`
	SortField      *string                `pulumi:"sortField"`
	SortOrder      *string                `pulumi:"sortOrder"`
	Tags           map[string]interface{} `pulumi:"tags"`
	TemplateFormat *string                `pulumi:"templateFormat"`
	TemplateType   *string                `pulumi:"templateType"`
	// A list of OOS Templates. Each element contains the following attributes:
	Templates []GetTemplatesTemplate `pulumi:"templates"`
}
