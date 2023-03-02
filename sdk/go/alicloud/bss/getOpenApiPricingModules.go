// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Bss Open Api Pricing Module available to the user.[What is Pricing Module](https://www.alibabacloud.com/help/en/bss-openapi/latest/describepricingmodule#doc-api-BssOpenApi-DescribePricingModule)
//
// > **NOTE:** Available in 1.195.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := bss.GetOpenApiPricingModules(ctx, &bss.GetOpenApiPricingModulesArgs{
//				NameRegex:        pulumi.StringRef("国内月均日峰值带宽"),
//				ProductCode:      "cdn",
//				ProductType:      pulumi.StringRef("CDN"),
//				SubscriptionType: "PayAsYouGo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudBssOpenapiPricingModuleExampleId", _default.Modules[0].Code)
//			return nil
//		})
//	}
//
// ```
func GetOpenApiPricingModules(ctx *pulumi.Context, args *GetOpenApiPricingModulesArgs, opts ...pulumi.InvokeOption) (*GetOpenApiPricingModulesResult, error) {
	var rv GetOpenApiPricingModulesResult
	err := ctx.Invoke("alicloud:bss/getOpenApiPricingModules:getOpenApiPricingModules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOpenApiPricingModules.
type GetOpenApiPricingModulesArgs struct {
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Property name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The product code.
	ProductCode string `pulumi:"productCode"`
	// The product type.
	ProductType *string `pulumi:"productType"`
	// Subscription type. Value:
	// * Subscription: Prepaid.
	// * PayAsYouGo: postpaid.
	SubscriptionType string `pulumi:"subscriptionType"`
}

// A collection of values returned by getOpenApiPricingModules.
type GetOpenApiPricingModulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// A list of Pricing Module Entries. Each element contains the following attributes:
	Modules   []GetOpenApiPricingModulesModule `pulumi:"modules"`
	NameRegex *string                          `pulumi:"nameRegex"`
	// A list of name of Pricing Modules.
	Names            []string `pulumi:"names"`
	OutputFile       *string  `pulumi:"outputFile"`
	ProductCode      string   `pulumi:"productCode"`
	ProductType      *string  `pulumi:"productType"`
	SubscriptionType string   `pulumi:"subscriptionType"`
}

func GetOpenApiPricingModulesOutput(ctx *pulumi.Context, args GetOpenApiPricingModulesOutputArgs, opts ...pulumi.InvokeOption) GetOpenApiPricingModulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOpenApiPricingModulesResult, error) {
			args := v.(GetOpenApiPricingModulesArgs)
			r, err := GetOpenApiPricingModules(ctx, &args, opts...)
			var s GetOpenApiPricingModulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOpenApiPricingModulesResultOutput)
}

// A collection of arguments for invoking getOpenApiPricingModules.
type GetOpenApiPricingModulesOutputArgs struct {
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Property name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The product code.
	ProductCode pulumi.StringInput `pulumi:"productCode"`
	// The product type.
	ProductType pulumi.StringPtrInput `pulumi:"productType"`
	// Subscription type. Value:
	// * Subscription: Prepaid.
	// * PayAsYouGo: postpaid.
	SubscriptionType pulumi.StringInput `pulumi:"subscriptionType"`
}

func (GetOpenApiPricingModulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpenApiPricingModulesArgs)(nil)).Elem()
}

// A collection of values returned by getOpenApiPricingModules.
type GetOpenApiPricingModulesResultOutput struct{ *pulumi.OutputState }

func (GetOpenApiPricingModulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpenApiPricingModulesResult)(nil)).Elem()
}

func (o GetOpenApiPricingModulesResultOutput) ToGetOpenApiPricingModulesResultOutput() GetOpenApiPricingModulesResultOutput {
	return o
}

func (o GetOpenApiPricingModulesResultOutput) ToGetOpenApiPricingModulesResultOutputWithContext(ctx context.Context) GetOpenApiPricingModulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOpenApiPricingModulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOpenApiPricingModulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of Pricing Module Entries. Each element contains the following attributes:
func (o GetOpenApiPricingModulesResultOutput) Modules() GetOpenApiPricingModulesModuleArrayOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) []GetOpenApiPricingModulesModule { return v.Modules }).(GetOpenApiPricingModulesModuleArrayOutput)
}

func (o GetOpenApiPricingModulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Pricing Modules.
func (o GetOpenApiPricingModulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetOpenApiPricingModulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetOpenApiPricingModulesResultOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) string { return v.ProductCode }).(pulumi.StringOutput)
}

func (o GetOpenApiPricingModulesResultOutput) ProductType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) *string { return v.ProductType }).(pulumi.StringPtrOutput)
}

func (o GetOpenApiPricingModulesResultOutput) SubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpenApiPricingModulesResult) string { return v.SubscriptionType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOpenApiPricingModulesResultOutput{})
}
