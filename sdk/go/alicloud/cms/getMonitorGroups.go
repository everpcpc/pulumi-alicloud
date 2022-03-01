// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cms Monitor Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.113.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cms.GetMonitorGroups(ctx, &cms.GetMonitorGroupsArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: pulumi.StringRef("the_resource_name"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstCmsMonitorGroupId", example.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetMonitorGroups(ctx *pulumi.Context, args *GetMonitorGroupsArgs, opts ...pulumi.InvokeOption) (*GetMonitorGroupsResult, error) {
	var rv GetMonitorGroupsResult
	err := ctx.Invoke("alicloud:cms/getMonitorGroups:getMonitorGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMonitorGroups.
type GetMonitorGroupsArgs struct {
	// The ID of the tag rule.
	DynamicTagRuleId *string `pulumi:"dynamicTagRuleId"`
	// A list of Monitor Group IDs.
	Ids []string `pulumi:"ids"`
	// The include template history.
	IncludeTemplateHistory *bool `pulumi:"includeTemplateHistory"`
	// The keyword to be matched.
	Keyword *string `pulumi:"keyword"`
	// The name of the application group.
	MonitorGroupName *string `pulumi:"monitorGroupName"`
	// A regex string to filter results by Monitor Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The select contact groups.
	SelectContactGroups *bool `pulumi:"selectContactGroups"`
	// A map of tags assigned to the Cms Monitor Group.
	Tags map[string]interface{} `pulumi:"tags"`
	// The type of the application group.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getMonitorGroups.
type GetMonitorGroupsResult struct {
	DynamicTagRuleId *string                 `pulumi:"dynamicTagRuleId"`
	Groups           []GetMonitorGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string                 `pulumi:"id"`
	Ids                    []string               `pulumi:"ids"`
	IncludeTemplateHistory *bool                  `pulumi:"includeTemplateHistory"`
	Keyword                *string                `pulumi:"keyword"`
	MonitorGroupName       *string                `pulumi:"monitorGroupName"`
	NameRegex              *string                `pulumi:"nameRegex"`
	Names                  []string               `pulumi:"names"`
	OutputFile             *string                `pulumi:"outputFile"`
	SelectContactGroups    *bool                  `pulumi:"selectContactGroups"`
	Tags                   map[string]interface{} `pulumi:"tags"`
	Type                   *string                `pulumi:"type"`
}

func GetMonitorGroupsOutput(ctx *pulumi.Context, args GetMonitorGroupsOutputArgs, opts ...pulumi.InvokeOption) GetMonitorGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitorGroupsResult, error) {
			args := v.(GetMonitorGroupsArgs)
			r, err := GetMonitorGroups(ctx, &args, opts...)
			return *r, err
		}).(GetMonitorGroupsResultOutput)
}

// A collection of arguments for invoking getMonitorGroups.
type GetMonitorGroupsOutputArgs struct {
	// The ID of the tag rule.
	DynamicTagRuleId pulumi.StringPtrInput `pulumi:"dynamicTagRuleId"`
	// A list of Monitor Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The include template history.
	IncludeTemplateHistory pulumi.BoolPtrInput `pulumi:"includeTemplateHistory"`
	// The keyword to be matched.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// The name of the application group.
	MonitorGroupName pulumi.StringPtrInput `pulumi:"monitorGroupName"`
	// A regex string to filter results by Monitor Group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The select contact groups.
	SelectContactGroups pulumi.BoolPtrInput `pulumi:"selectContactGroups"`
	// A map of tags assigned to the Cms Monitor Group.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of the application group.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetMonitorGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getMonitorGroups.
type GetMonitorGroupsResultOutput struct{ *pulumi.OutputState }

func (GetMonitorGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorGroupsResult)(nil)).Elem()
}

func (o GetMonitorGroupsResultOutput) ToGetMonitorGroupsResultOutput() GetMonitorGroupsResultOutput {
	return o
}

func (o GetMonitorGroupsResultOutput) ToGetMonitorGroupsResultOutputWithContext(ctx context.Context) GetMonitorGroupsResultOutput {
	return o
}

func (o GetMonitorGroupsResultOutput) DynamicTagRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.DynamicTagRuleId }).(pulumi.StringPtrOutput)
}

func (o GetMonitorGroupsResultOutput) Groups() GetMonitorGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) []GetMonitorGroupsGroup { return v.Groups }).(GetMonitorGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMonitorGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMonitorGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetMonitorGroupsResultOutput) IncludeTemplateHistory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *bool { return v.IncludeTemplateHistory }).(pulumi.BoolPtrOutput)
}

func (o GetMonitorGroupsResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o GetMonitorGroupsResultOutput) MonitorGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.MonitorGroupName }).(pulumi.StringPtrOutput)
}

func (o GetMonitorGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetMonitorGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetMonitorGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetMonitorGroupsResultOutput) SelectContactGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *bool { return v.SelectContactGroups }).(pulumi.BoolPtrOutput)
}

func (o GetMonitorGroupsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetMonitorGroupsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorGroupsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitorGroupsResultOutput{})
}
