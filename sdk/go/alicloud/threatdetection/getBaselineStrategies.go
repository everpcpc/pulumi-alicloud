// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Threat Detection Baseline Strategy available to the user.[What is Baseline Strategy](https://www.alibabacloud.com/help/zh/security-center/latest/api-doc-sas-2018-12-03-api-doc-modifystrategy)
//
// > **NOTE:** Available in 1.195.0+
func GetBaselineStrategies(ctx *pulumi.Context, args *GetBaselineStrategiesArgs, opts ...pulumi.InvokeOption) (*GetBaselineStrategiesResult, error) {
	var rv GetBaselineStrategiesResult
	err := ctx.Invoke("alicloud:threatdetection/getBaselineStrategies:getBaselineStrategies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaselineStrategies.
type GetBaselineStrategiesArgs struct {
	// The type of policy. Value:-**common**: standard policy-**custom**: custom policy
	CustomType *string `pulumi:"customType"`
	// A list of Baseline Strategy IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  *string `pulumi:"outputFile"`
	StrategyIds *string `pulumi:"strategyIds"`
}

// A collection of values returned by getBaselineStrategies.
type GetBaselineStrategiesResult struct {
	// The type of policy. Value:
	// * **common**: standard policy
	// * **custom**: custom policy
	CustomType *string `pulumi:"customType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Baseline Strategy IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of name of Baseline Strategys.
	Names       []string                        `pulumi:"names"`
	OutputFile  *string                         `pulumi:"outputFile"`
	Strategies  []GetBaselineStrategiesStrategy `pulumi:"strategies"`
	StrategyIds *string                         `pulumi:"strategyIds"`
}

func GetBaselineStrategiesOutput(ctx *pulumi.Context, args GetBaselineStrategiesOutputArgs, opts ...pulumi.InvokeOption) GetBaselineStrategiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBaselineStrategiesResult, error) {
			args := v.(GetBaselineStrategiesArgs)
			r, err := GetBaselineStrategies(ctx, &args, opts...)
			var s GetBaselineStrategiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBaselineStrategiesResultOutput)
}

// A collection of arguments for invoking getBaselineStrategies.
type GetBaselineStrategiesOutputArgs struct {
	// The type of policy. Value:-**common**: standard policy-**custom**: custom policy
	CustomType pulumi.StringPtrInput `pulumi:"customType"`
	// A list of Baseline Strategy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Group Metric Rule name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
	StrategyIds pulumi.StringPtrInput `pulumi:"strategyIds"`
}

func (GetBaselineStrategiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaselineStrategiesArgs)(nil)).Elem()
}

// A collection of values returned by getBaselineStrategies.
type GetBaselineStrategiesResultOutput struct{ *pulumi.OutputState }

func (GetBaselineStrategiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaselineStrategiesResult)(nil)).Elem()
}

func (o GetBaselineStrategiesResultOutput) ToGetBaselineStrategiesResultOutput() GetBaselineStrategiesResultOutput {
	return o
}

func (o GetBaselineStrategiesResultOutput) ToGetBaselineStrategiesResultOutputWithContext(ctx context.Context) GetBaselineStrategiesResultOutput {
	return o
}

// The type of policy. Value:
// * **common**: standard policy
// * **custom**: custom policy
func (o GetBaselineStrategiesResultOutput) CustomType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) *string { return v.CustomType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBaselineStrategiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Baseline Strategy IDs.
func (o GetBaselineStrategiesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetBaselineStrategiesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Baseline Strategys.
func (o GetBaselineStrategiesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetBaselineStrategiesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBaselineStrategiesResultOutput) Strategies() GetBaselineStrategiesStrategyArrayOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) []GetBaselineStrategiesStrategy { return v.Strategies }).(GetBaselineStrategiesStrategyArrayOutput)
}

func (o GetBaselineStrategiesResultOutput) StrategyIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaselineStrategiesResult) *string { return v.StrategyIds }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBaselineStrategiesResultOutput{})
}
