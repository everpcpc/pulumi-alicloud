// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Cloud Monitor Service Group Metric Rule resource.
//
// For information about Cloud Monitor Service Group Metric Rule and how to use it, see [What is Group Metric Rule](https://www.alibabacloud.com/help/en/doc-detail/114943.htm).
//
// > **NOTE:** Available in v1.104.0+.
//
// ## Import
//
// Cloud Monitor Service Group Metric Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cms/groupMetricRule:GroupMetricRule example <rule_id>
// ```
type GroupMetricRule struct {
	pulumi.CustomResourceState

	// The abbreviation of the service name.
	Category pulumi.StringOutput `pulumi:"category"`
	// Alarm contact group.
	ContactGroups pulumi.StringOutput `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringOutput `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrOutput `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.                                         .
	EmailSubject pulumi.StringOutput `pulumi:"emailSubject"`
	// Alarm level. See the block for escalations.
	Escalations GroupMetricRuleEscalationsOutput `pulumi:"escalations"`
	// The ID of the application group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringOutput `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrOutput `pulumi:"interval"`
	// The name of the metric.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The namespace of the service.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrOutput `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The ID of the alert rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntPtrOutput `pulumi:"silenceTime"`
	// The status of Group Metric Rule.
	Status pulumi.StringOutput `pulumi:"status"`
	// The callback URL.
	Webhook pulumi.StringPtrOutput `pulumi:"webhook"`
}

// NewGroupMetricRule registers a new resource with the given unique name, arguments, and options.
func NewGroupMetricRule(ctx *pulumi.Context,
	name string, args *GroupMetricRuleArgs, opts ...pulumi.ResourceOption) (*GroupMetricRule, error) {
	if args == nil || args.Category == nil {
		return nil, errors.New("missing required argument 'Category'")
	}
	if args == nil || args.Escalations == nil {
		return nil, errors.New("missing required argument 'Escalations'")
	}
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil || args.GroupMetricRuleName == nil {
		return nil, errors.New("missing required argument 'GroupMetricRuleName'")
	}
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	if args == nil || args.Namespace == nil {
		return nil, errors.New("missing required argument 'Namespace'")
	}
	if args == nil || args.RuleId == nil {
		return nil, errors.New("missing required argument 'RuleId'")
	}
	if args == nil {
		args = &GroupMetricRuleArgs{}
	}
	var resource GroupMetricRule
	err := ctx.RegisterResource("alicloud:cms/groupMetricRule:GroupMetricRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMetricRule gets an existing GroupMetricRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMetricRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMetricRuleState, opts ...pulumi.ResourceOption) (*GroupMetricRule, error) {
	var resource GroupMetricRule
	err := ctx.ReadResource("alicloud:cms/groupMetricRule:GroupMetricRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMetricRule resources.
type groupMetricRuleState struct {
	// The abbreviation of the service name.
	Category *string `pulumi:"category"`
	// Alarm contact group.
	ContactGroups *string `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions *string `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.                                         .
	EmailSubject *string `pulumi:"emailSubject"`
	// Alarm level. See the block for escalations.
	Escalations *GroupMetricRuleEscalations `pulumi:"escalations"`
	// The ID of the application group.
	GroupId *string `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName *string `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval *string `pulumi:"interval"`
	// The name of the metric.
	MetricName *string `pulumi:"metricName"`
	// The namespace of the service.
	Namespace *string `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval *string `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period *int `pulumi:"period"`
	// The ID of the alert rule.
	RuleId *string `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime *int `pulumi:"silenceTime"`
	// The status of Group Metric Rule.
	Status *string `pulumi:"status"`
	// The callback URL.
	Webhook *string `pulumi:"webhook"`
}

type GroupMetricRuleState struct {
	// The abbreviation of the service name.
	Category pulumi.StringPtrInput
	// Alarm contact group.
	ContactGroups pulumi.StringPtrInput
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringPtrInput
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrInput
	// The subject of the alert notification email.                                         .
	EmailSubject pulumi.StringPtrInput
	// Alarm level. See the block for escalations.
	Escalations GroupMetricRuleEscalationsPtrInput
	// The ID of the application group.
	GroupId pulumi.StringPtrInput
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringPtrInput
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrInput
	// The name of the metric.
	MetricName pulumi.StringPtrInput
	// The namespace of the service.
	Namespace pulumi.StringPtrInput
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrInput
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntPtrInput
	// The ID of the alert rule.
	RuleId pulumi.StringPtrInput
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntPtrInput
	// The status of Group Metric Rule.
	Status pulumi.StringPtrInput
	// The callback URL.
	Webhook pulumi.StringPtrInput
}

func (GroupMetricRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMetricRuleState)(nil)).Elem()
}

type groupMetricRuleArgs struct {
	// The abbreviation of the service name.
	Category string `pulumi:"category"`
	// Alarm contact group.
	ContactGroups *string `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions *string `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.                                         .
	EmailSubject *string `pulumi:"emailSubject"`
	// Alarm level. See the block for escalations.
	Escalations GroupMetricRuleEscalations `pulumi:"escalations"`
	// The ID of the application group.
	GroupId string `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName string `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval *string `pulumi:"interval"`
	// The name of the metric.
	MetricName string `pulumi:"metricName"`
	// The namespace of the service.
	Namespace string `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval *string `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period *int `pulumi:"period"`
	// The ID of the alert rule.
	RuleId string `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime *int `pulumi:"silenceTime"`
	// The callback URL.
	Webhook *string `pulumi:"webhook"`
}

// The set of arguments for constructing a GroupMetricRule resource.
type GroupMetricRuleArgs struct {
	// The abbreviation of the service name.
	Category pulumi.StringInput
	// Alarm contact group.
	ContactGroups pulumi.StringPtrInput
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringPtrInput
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrInput
	// The subject of the alert notification email.                                         .
	EmailSubject pulumi.StringPtrInput
	// Alarm level. See the block for escalations.
	Escalations GroupMetricRuleEscalationsInput
	// The ID of the application group.
	GroupId pulumi.StringInput
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringInput
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrInput
	// The name of the metric.
	MetricName pulumi.StringInput
	// The namespace of the service.
	Namespace pulumi.StringInput
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrInput
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntPtrInput
	// The ID of the alert rule.
	RuleId pulumi.StringInput
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntPtrInput
	// The callback URL.
	Webhook pulumi.StringPtrInput
}

func (GroupMetricRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMetricRuleArgs)(nil)).Elem()
}

type GroupMetricRuleInput interface {
	pulumi.Input

	ToGroupMetricRuleOutput() GroupMetricRuleOutput
	ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput
}

func (GroupMetricRule) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMetricRule)(nil)).Elem()
}

func (i GroupMetricRule) ToGroupMetricRuleOutput() GroupMetricRuleOutput {
	return i.ToGroupMetricRuleOutputWithContext(context.Background())
}

func (i GroupMetricRule) ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMetricRuleOutput)
}

type GroupMetricRuleOutput struct {
	*pulumi.OutputState
}

func (GroupMetricRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMetricRuleOutput)(nil)).Elem()
}

func (o GroupMetricRuleOutput) ToGroupMetricRuleOutput() GroupMetricRuleOutput {
	return o
}

func (o GroupMetricRuleOutput) ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupMetricRuleOutput{})
}