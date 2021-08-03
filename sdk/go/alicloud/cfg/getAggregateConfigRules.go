// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
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
	// The name of the rule.
	AggregateConfigRuleName *string `pulumi:"aggregateConfigRuleName"`
	// The ID of Aggregator.
	AggregatorId string `pulumi:"aggregatorId"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Aggregate Config Rule IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Aggregate Config Rule name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The risk level of the resources that are not compliant with the rule. Valid values: `1`: critical, `2`: warning, `3`: info.
	RiskLevel *int `pulumi:"riskLevel"`
	// The status of the rule.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAggregateConfigRules.
type GetAggregateConfigRulesResult struct {
	AggregateConfigRuleName *string `pulumi:"aggregateConfigRuleName"`
	AggregatorId            string  `pulumi:"aggregatorId"`
	EnableDetails           *bool   `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                        `pulumi:"id"`
	Ids        []string                      `pulumi:"ids"`
	NameRegex  *string                       `pulumi:"nameRegex"`
	Names      []string                      `pulumi:"names"`
	OutputFile *string                       `pulumi:"outputFile"`
	RiskLevel  *int                          `pulumi:"riskLevel"`
	Rules      []GetAggregateConfigRulesRule `pulumi:"rules"`
	Status     *string                       `pulumi:"status"`
}