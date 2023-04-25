// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// ESS scaling rule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ess/scalingRule:ScalingRule example abc123456
//
// ```
type ScalingRule struct {
	pulumi.CustomResourceState

	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrOutput `pulumi:"adjustmentType"`
	// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrOutput `pulumi:"adjustmentValue"`
	Ari             pulumi.StringOutput `pulumi:"ari"`
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrOutput `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrOutput `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntOutput `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrOutput `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringOutput `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrOutput `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayOutput `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue pulumi.Float64PtrOutput `pulumi:"targetValue"`
}

// NewScalingRule registers a new resource with the given unique name, arguments, and options.
func NewScalingRule(ctx *pulumi.Context,
	name string, args *ScalingRuleArgs, opts ...pulumi.ResourceOption) (*ScalingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	var resource ScalingRule
	err := ctx.RegisterResource("alicloud:ess/scalingRule:ScalingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingRule gets an existing ScalingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingRuleState, opts ...pulumi.ResourceOption) (*ScalingRule, error) {
	var resource ScalingRule
	err := ctx.ReadResource("alicloud:ess/scalingRule:ScalingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingRule resources.
type scalingRuleState struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue *int    `pulumi:"adjustmentValue"`
	Ari             *string `pulumi:"ari"`
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown *int `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn *bool `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName *string `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName *string `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType *string `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments []ScalingRuleStepAdjustment `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue *float64 `pulumi:"targetValue"`
}

type ScalingRuleState struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrInput
	// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrInput
	Ari             pulumi.StringPtrInput
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrInput
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrInput
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringPtrInput
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringPtrInput
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrInput
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayInput
	// The target value for the metric.
	TargetValue pulumi.Float64PtrInput
}

func (ScalingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingRuleState)(nil)).Elem()
}

type scalingRuleArgs struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue *int `pulumi:"adjustmentValue"`
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown *int `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn *bool `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName *string `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName *string `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType *string `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments []ScalingRuleStepAdjustment `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue *float64 `pulumi:"targetValue"`
}

// The set of arguments for constructing a ScalingRule resource.
type ScalingRuleArgs struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrInput
	// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrInput
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrInput
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrInput
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringInput
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringPtrInput
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrInput
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayInput
	// The target value for the metric.
	TargetValue pulumi.Float64PtrInput
}

func (ScalingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingRuleArgs)(nil)).Elem()
}

type ScalingRuleInput interface {
	pulumi.Input

	ToScalingRuleOutput() ScalingRuleOutput
	ToScalingRuleOutputWithContext(ctx context.Context) ScalingRuleOutput
}

func (*ScalingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingRule)(nil)).Elem()
}

func (i *ScalingRule) ToScalingRuleOutput() ScalingRuleOutput {
	return i.ToScalingRuleOutputWithContext(context.Background())
}

func (i *ScalingRule) ToScalingRuleOutputWithContext(ctx context.Context) ScalingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingRuleOutput)
}

// ScalingRuleArrayInput is an input type that accepts ScalingRuleArray and ScalingRuleArrayOutput values.
// You can construct a concrete instance of `ScalingRuleArrayInput` via:
//
//	ScalingRuleArray{ ScalingRuleArgs{...} }
type ScalingRuleArrayInput interface {
	pulumi.Input

	ToScalingRuleArrayOutput() ScalingRuleArrayOutput
	ToScalingRuleArrayOutputWithContext(context.Context) ScalingRuleArrayOutput
}

type ScalingRuleArray []ScalingRuleInput

func (ScalingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingRule)(nil)).Elem()
}

func (i ScalingRuleArray) ToScalingRuleArrayOutput() ScalingRuleArrayOutput {
	return i.ToScalingRuleArrayOutputWithContext(context.Background())
}

func (i ScalingRuleArray) ToScalingRuleArrayOutputWithContext(ctx context.Context) ScalingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingRuleArrayOutput)
}

// ScalingRuleMapInput is an input type that accepts ScalingRuleMap and ScalingRuleMapOutput values.
// You can construct a concrete instance of `ScalingRuleMapInput` via:
//
//	ScalingRuleMap{ "key": ScalingRuleArgs{...} }
type ScalingRuleMapInput interface {
	pulumi.Input

	ToScalingRuleMapOutput() ScalingRuleMapOutput
	ToScalingRuleMapOutputWithContext(context.Context) ScalingRuleMapOutput
}

type ScalingRuleMap map[string]ScalingRuleInput

func (ScalingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingRule)(nil)).Elem()
}

func (i ScalingRuleMap) ToScalingRuleMapOutput() ScalingRuleMapOutput {
	return i.ToScalingRuleMapOutputWithContext(context.Background())
}

func (i ScalingRuleMap) ToScalingRuleMapOutputWithContext(ctx context.Context) ScalingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingRuleMapOutput)
}

type ScalingRuleOutput struct{ *pulumi.OutputState }

func (ScalingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingRule)(nil)).Elem()
}

func (o ScalingRuleOutput) ToScalingRuleOutput() ScalingRuleOutput {
	return o
}

func (o ScalingRuleOutput) ToScalingRuleOutputWithContext(ctx context.Context) ScalingRuleOutput {
	return o
}

// Adjustment mode of a scaling rule. Optional values:
// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
func (o ScalingRuleOutput) AdjustmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringPtrOutput { return v.AdjustmentType }).(pulumi.StringPtrOutput)
}

// The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
// - TotalCapacity：[0, 1000]
func (o ScalingRuleOutput) AdjustmentValue() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.IntPtrOutput { return v.AdjustmentValue }).(pulumi.IntPtrOutput)
}

func (o ScalingRuleOutput) Ari() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringOutput { return v.Ari }).(pulumi.StringOutput)
}

// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
func (o ScalingRuleOutput) Cooldown() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.IntPtrOutput { return v.Cooldown }).(pulumi.IntPtrOutput)
}

// Indicates whether scale in by the target tracking policy is disabled. Default to false.
func (o ScalingRuleOutput) DisableScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.BoolPtrOutput { return v.DisableScaleIn }).(pulumi.BoolPtrOutput)
}

// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
func (o ScalingRuleOutput) EstimatedInstanceWarmup() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.IntOutput { return v.EstimatedInstanceWarmup }).(pulumi.IntOutput)
}

// A CloudMonitor metric name.
func (o ScalingRuleOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringPtrOutput { return v.MetricName }).(pulumi.StringPtrOutput)
}

// ID of the scaling group of a scaling rule.
func (o ScalingRuleOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
func (o ScalingRuleOutput) ScalingRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringOutput { return v.ScalingRuleName }).(pulumi.StringOutput)
}

// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
func (o ScalingRuleOutput) ScalingRuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.StringPtrOutput { return v.ScalingRuleType }).(pulumi.StringPtrOutput)
}

// Steps for StepScalingRule. See Block stepAdjustment below for details.
func (o ScalingRuleOutput) StepAdjustments() ScalingRuleStepAdjustmentArrayOutput {
	return o.ApplyT(func(v *ScalingRule) ScalingRuleStepAdjustmentArrayOutput { return v.StepAdjustments }).(ScalingRuleStepAdjustmentArrayOutput)
}

// The target value for the metric.
func (o ScalingRuleOutput) TargetValue() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ScalingRule) pulumi.Float64PtrOutput { return v.TargetValue }).(pulumi.Float64PtrOutput)
}

type ScalingRuleArrayOutput struct{ *pulumi.OutputState }

func (ScalingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingRule)(nil)).Elem()
}

func (o ScalingRuleArrayOutput) ToScalingRuleArrayOutput() ScalingRuleArrayOutput {
	return o
}

func (o ScalingRuleArrayOutput) ToScalingRuleArrayOutputWithContext(ctx context.Context) ScalingRuleArrayOutput {
	return o
}

func (o ScalingRuleArrayOutput) Index(i pulumi.IntInput) ScalingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingRule {
		return vs[0].([]*ScalingRule)[vs[1].(int)]
	}).(ScalingRuleOutput)
}

type ScalingRuleMapOutput struct{ *pulumi.OutputState }

func (ScalingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingRule)(nil)).Elem()
}

func (o ScalingRuleMapOutput) ToScalingRuleMapOutput() ScalingRuleMapOutput {
	return o
}

func (o ScalingRuleMapOutput) ToScalingRuleMapOutputWithContext(ctx context.Context) ScalingRuleMapOutput {
	return o
}

func (o ScalingRuleMapOutput) MapIndex(k pulumi.StringInput) ScalingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingRule {
		return vs[0].(map[string]*ScalingRule)[vs[1].(string)]
	}).(ScalingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingRuleInput)(nil)).Elem(), &ScalingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingRuleArrayInput)(nil)).Elem(), ScalingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingRuleMapInput)(nil)).Elem(), ScalingRuleMap{})
	pulumi.RegisterOutputType(ScalingRuleOutput{})
	pulumi.RegisterOutputType(ScalingRuleArrayOutput{})
	pulumi.RegisterOutputType(ScalingRuleMapOutput{})
}
