// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Config Compliance Packs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.124.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "the_resource_name"
// 		example, err := cfg.GetCompliancePacks(ctx, &cfg.GetCompliancePacksArgs{
// 			Ids: []string{
// 				"cp-152a626622af00bc****",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstConfigCompliancePackId", example.Packs[0].Id)
// 		return nil
// 	})
// }
// ```
func GetCompliancePacks(ctx *pulumi.Context, args *GetCompliancePacksArgs, opts ...pulumi.InvokeOption) (*GetCompliancePacksResult, error) {
	var rv GetCompliancePacksResult
	err := ctx.Invoke("alicloud:cfg/getCompliancePacks:getCompliancePacks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCompliancePacks.
type GetCompliancePacksArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Compliance Pack IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Compliance Pack name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getCompliancePacks.
type GetCompliancePacksResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                   `pulumi:"id"`
	Ids        []string                 `pulumi:"ids"`
	NameRegex  *string                  `pulumi:"nameRegex"`
	Names      []string                 `pulumi:"names"`
	OutputFile *string                  `pulumi:"outputFile"`
	Packs      []GetCompliancePacksPack `pulumi:"packs"`
	Status     *string                  `pulumi:"status"`
}

func GetCompliancePacksOutput(ctx *pulumi.Context, args GetCompliancePacksOutputArgs, opts ...pulumi.InvokeOption) GetCompliancePacksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCompliancePacksResult, error) {
			args := v.(GetCompliancePacksArgs)
			r, err := GetCompliancePacks(ctx, &args, opts...)
			return *r, err
		}).(GetCompliancePacksResultOutput)
}

// A collection of arguments for invoking getCompliancePacks.
type GetCompliancePacksOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Compliance Pack IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Compliance Pack name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetCompliancePacksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCompliancePacksArgs)(nil)).Elem()
}

// A collection of values returned by getCompliancePacks.
type GetCompliancePacksResultOutput struct{ *pulumi.OutputState }

func (GetCompliancePacksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCompliancePacksResult)(nil)).Elem()
}

func (o GetCompliancePacksResultOutput) ToGetCompliancePacksResultOutput() GetCompliancePacksResultOutput {
	return o
}

func (o GetCompliancePacksResultOutput) ToGetCompliancePacksResultOutputWithContext(ctx context.Context) GetCompliancePacksResultOutput {
	return o
}

func (o GetCompliancePacksResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCompliancePacksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCompliancePacksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCompliancePacksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetCompliancePacksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetCompliancePacksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCompliancePacksResultOutput) Packs() GetCompliancePacksPackArrayOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) []GetCompliancePacksPack { return v.Packs }).(GetCompliancePacksPackArrayOutput)
}

func (o GetCompliancePacksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCompliancePacksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCompliancePacksResultOutput{})
}
