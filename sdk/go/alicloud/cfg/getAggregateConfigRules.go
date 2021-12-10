// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Config Aggregate Config Rules of the current Alibaba Cloud user.
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
// 		example, err := cfg.GetAggregateConfigRules(ctx, &cfg.GetAggregateConfigRulesArgs{
// 			AggregatorId: "ca-3a9b626622af001d****",
// 			Ids: []string{
// 				"cr-5154626622af0034****",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstConfigAggregateConfigRuleId", example.Rules[0].Id)
// 		return nil
// 	})
// }
// ```
func GetAggregateConfigRules(ctx *pulumi.Context, args *GetAggregateConfigRulesArgs, opts ...pulumi.InvokeOption) (*GetAggregateConfigRulesResult, error) {
	var rv GetAggregateConfigRulesResult
	err := ctx.Invoke("alicloud:cfg/getAggregateConfigRules:getAggregateConfigRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAggregateConfigRules.
type GetAggregateConfigRulesArgs struct {
	// The config rule name.
	AggregateConfigRuleName *string `pulumi:"aggregateConfigRuleName"`
	// The ID of aggregator.
	AggregatorId string `pulumi:"aggregatorId"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Aggregate Config Rule IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Aggregate Config Rule name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
	RiskLevel *int `pulumi:"riskLevel"`
	// The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAggregateConfigRules.
type GetAggregateConfigRulesResult struct {
	// The name of the rule.
	AggregateConfigRuleName *string `pulumi:"aggregateConfigRuleName"`
	// The ID of Aggregator.
	// * `compliance` -The Compliance information.
	AggregatorId  string `pulumi:"aggregatorId"`
	EnableDetails *bool  `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Aggregate Config Rule names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The risk level of the resources that are not compliant with the rule. Valid values: `1`: critical, `2`: warning, `3`: info.
	RiskLevel *int `pulumi:"riskLevel"`
	// A list of Config Aggregate Config Rules. Each element contains the following attributes:
	Rules []GetAggregateConfigRulesRule `pulumi:"rules"`
	// The status of the rule.
	Status *string `pulumi:"status"`
}

func GetAggregateConfigRulesOutput(ctx *pulumi.Context, args GetAggregateConfigRulesOutputArgs, opts ...pulumi.InvokeOption) GetAggregateConfigRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAggregateConfigRulesResult, error) {
			args := v.(GetAggregateConfigRulesArgs)
			r, err := GetAggregateConfigRules(ctx, &args, opts...)
			return *r, err
		}).(GetAggregateConfigRulesResultOutput)
}

// A collection of arguments for invoking getAggregateConfigRules.
type GetAggregateConfigRulesOutputArgs struct {
	// The config rule name.
	AggregateConfigRuleName pulumi.StringPtrInput `pulumi:"aggregateConfigRuleName"`
	// The ID of aggregator.
	AggregatorId pulumi.StringInput `pulumi:"aggregatorId"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Aggregate Config Rule IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Aggregate Config Rule name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
	RiskLevel pulumi.IntPtrInput `pulumi:"riskLevel"`
	// The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetAggregateConfigRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAggregateConfigRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAggregateConfigRules.
type GetAggregateConfigRulesResultOutput struct{ *pulumi.OutputState }

func (GetAggregateConfigRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAggregateConfigRulesResult)(nil)).Elem()
}

func (o GetAggregateConfigRulesResultOutput) ToGetAggregateConfigRulesResultOutput() GetAggregateConfigRulesResultOutput {
	return o
}

func (o GetAggregateConfigRulesResultOutput) ToGetAggregateConfigRulesResultOutputWithContext(ctx context.Context) GetAggregateConfigRulesResultOutput {
	return o
}

// The name of the rule.
func (o GetAggregateConfigRulesResultOutput) AggregateConfigRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *string { return v.AggregateConfigRuleName }).(pulumi.StringPtrOutput)
}

// The ID of Aggregator.
// * `compliance` -The Compliance information.
func (o GetAggregateConfigRulesResultOutput) AggregatorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) string { return v.AggregatorId }).(pulumi.StringOutput)
}

func (o GetAggregateConfigRulesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAggregateConfigRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAggregateConfigRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAggregateConfigRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Aggregate Config Rule names.
func (o GetAggregateConfigRulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAggregateConfigRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The risk level of the resources that are not compliant with the rule. Valid values: `1`: critical, `2`: warning, `3`: info.
func (o GetAggregateConfigRulesResultOutput) RiskLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *int { return v.RiskLevel }).(pulumi.IntPtrOutput)
}

// A list of Config Aggregate Config Rules. Each element contains the following attributes:
func (o GetAggregateConfigRulesResultOutput) Rules() GetAggregateConfigRulesRuleArrayOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) []GetAggregateConfigRulesRule { return v.Rules }).(GetAggregateConfigRulesRuleArrayOutput)
}

// The status of the rule.
func (o GetAggregateConfigRulesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateConfigRulesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAggregateConfigRulesResultOutput{})
}
