// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Threat Detection Anti Brute Force Rule available to the user.[What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-createantibruteforcerule)
//
// > **NOTE:** Available in 1.195.0+
func GetAntiBruteForceRules(ctx *pulumi.Context, args *GetAntiBruteForceRulesArgs, opts ...pulumi.InvokeOption) (*GetAntiBruteForceRulesResult, error) {
	var rv GetAntiBruteForceRulesResult
	err := ctx.Invoke("alicloud:threatdetection/getAntiBruteForceRules:getAntiBruteForceRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAntiBruteForceRules.
type GetAntiBruteForceRulesArgs struct {
	// A list of Anti-Brute Force Rule IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by the name of the defense rule.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAntiBruteForceRules.
type GetAntiBruteForceRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Anti Brute Force Rule IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of name of Anti Brute Force Rules.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of Anti Brute Force Rule Entries. Each element contains the following attributes:
	Rules []GetAntiBruteForceRulesRule `pulumi:"rules"`
}

func GetAntiBruteForceRulesOutput(ctx *pulumi.Context, args GetAntiBruteForceRulesOutputArgs, opts ...pulumi.InvokeOption) GetAntiBruteForceRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAntiBruteForceRulesResult, error) {
			args := v.(GetAntiBruteForceRulesArgs)
			r, err := GetAntiBruteForceRules(ctx, &args, opts...)
			var s GetAntiBruteForceRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAntiBruteForceRulesResultOutput)
}

// A collection of arguments for invoking getAntiBruteForceRules.
type GetAntiBruteForceRulesOutputArgs struct {
	// A list of Anti-Brute Force Rule IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by the name of the defense rule.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAntiBruteForceRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAntiBruteForceRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAntiBruteForceRules.
type GetAntiBruteForceRulesResultOutput struct{ *pulumi.OutputState }

func (GetAntiBruteForceRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAntiBruteForceRulesResult)(nil)).Elem()
}

func (o GetAntiBruteForceRulesResultOutput) ToGetAntiBruteForceRulesResultOutput() GetAntiBruteForceRulesResultOutput {
	return o
}

func (o GetAntiBruteForceRulesResultOutput) ToGetAntiBruteForceRulesResultOutputWithContext(ctx context.Context) GetAntiBruteForceRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAntiBruteForceRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Anti Brute Force Rule IDs.
func (o GetAntiBruteForceRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAntiBruteForceRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Anti Brute Force Rules.
func (o GetAntiBruteForceRulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAntiBruteForceRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of Anti Brute Force Rule Entries. Each element contains the following attributes:
func (o GetAntiBruteForceRulesResultOutput) Rules() GetAntiBruteForceRulesRuleArrayOutput {
	return o.ApplyT(func(v GetAntiBruteForceRulesResult) []GetAntiBruteForceRulesRule { return v.Rules }).(GetAntiBruteForceRulesRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAntiBruteForceRulesResultOutput{})
}
