// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides a alarm rule resource and it can be used to monitor several cloud services according different metrics.
// Details for [alarm rule](https://www.alibabacloud.com/help/doc-detail/28608.htm).
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cms.NewAlarm(ctx, "basic", &cms.AlarmArgs{
// 			ContactGroups: pulumi.StringArray{
// 				pulumi.String("test-group"),
// 			},
// 			EffectiveInterval: pulumi.String("0:00-2:00"),
// 			EscalationsCritical: &cms.AlarmEscalationsCriticalArgs{
// 				ComparisonOperator: pulumi.String("<="),
// 				Statistics:         pulumi.String("Average"),
// 				Threshold:          pulumi.String("35"),
// 				Times:              pulumi.Int(2),
// 			},
// 			MetricDimensions: cms.AlarmMetricDimensionArray{
// 				&cms.AlarmMetricDimensionArgs{
// 					Key:   pulumi.String("instanceId"),
// 					Value: pulumi.String("i-bp1247jeep0y53nu3bnk"),
// 				},
// 				&cms.AlarmMetricDimensionArgs{
// 					Key:   pulumi.String("device"),
// 					Value: pulumi.String("/dev/vda1"),
// 				},
// 			},
// 			Period:  pulumi.Int(900),
// 			Project: pulumi.String("acs_ecs_dashboard"),
// 			Webhook: pulumi.String(fmt.Sprintf("%v%v%v", "https://", data.Alicloud_account.Current.Id, ".eu-central-1.fc.aliyuncs.com/2016-08-15/proxy/Terraform/AlarmEndpointMock/")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Alarm rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cms/alarm:Alarm alarm abc12345
// ```
type Alarm struct {
	pulumi.CustomResourceState

	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayOutput `pulumi:"contactGroups"`
	// Field `dimensions` has been deprecated from version 1.95.0. Use `metricDimensions` instead.
	//
	// Deprecated: Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead.
	Dimensions pulumi.MapOutput `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrOutput `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// A configuration of critical alarm (documented below).
	EscalationsCritical AlarmEscalationsCriticalPtrOutput `pulumi:"escalationsCritical"`
	// A configuration of critical info (documented below).
	EscalationsInfo AlarmEscalationsInfoPtrOutput `pulumi:"escalationsInfo"`
	// A configuration of critical warn (documented below).
	EscalationsWarn AlarmEscalationsWarnPtrOutput `pulumi:"escalationsWarn"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringOutput `pulumi:"metric"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	MetricDimensions AlarmMetricDimensionArrayOutput `pulumi:"metricDimensions"`
	// The alarm rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
	//
	// Deprecated: Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
	Operator pulumi.StringOutput `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringOutput `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrOutput `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Availability","Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
	//
	// Deprecated: Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
	Statistics pulumi.StringOutput `pulumi:"statistics"`
	// The current alarm rule status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Critical level alarm threshold value, which must be a numeric value currently.
	//
	// Deprecated: Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
	Threshold pulumi.StringOutput `pulumi:"threshold"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
	//
	// Deprecated: Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
	TriggeredCount pulumi.IntOutput `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrOutput `pulumi:"webhook"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactGroups == nil {
		return nil, errors.New("invalid value for required argument 'ContactGroups'")
	}
	if args.Metric == nil {
		return nil, errors.New("invalid value for required argument 'Metric'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Alarm
	err := ctx.RegisterResource("alicloud:cms/alarm:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("alicloud:cms/alarm:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups []string `pulumi:"contactGroups"`
	// Field `dimensions` has been deprecated from version 1.95.0. Use `metricDimensions` instead.
	//
	// Deprecated: Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead.
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled *bool `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	EndTime *int `pulumi:"endTime"`
	// A configuration of critical alarm (documented below).
	EscalationsCritical *AlarmEscalationsCritical `pulumi:"escalationsCritical"`
	// A configuration of critical info (documented below).
	EscalationsInfo *AlarmEscalationsInfo `pulumi:"escalationsInfo"`
	// A configuration of critical warn (documented below).
	EscalationsWarn *AlarmEscalationsWarn `pulumi:"escalationsWarn"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric *string `pulumi:"metric"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	MetricDimensions []AlarmMetricDimension `pulumi:"metricDimensions"`
	// The alarm rule name.
	Name *string `pulumi:"name"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
	//
	// Deprecated: Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
	Operator *string `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period *int `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project *string `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime *int `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	StartTime *int `pulumi:"startTime"`
	// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Availability","Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
	//
	// Deprecated: Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
	Statistics *string `pulumi:"statistics"`
	// The current alarm rule status.
	Status *string `pulumi:"status"`
	// Critical level alarm threshold value, which must be a numeric value currently.
	//
	// Deprecated: Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
	Threshold *string `pulumi:"threshold"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
	//
	// Deprecated: Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
	TriggeredCount *int `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook *string `pulumi:"webhook"`
}

type AlarmState struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayInput
	// Field `dimensions` has been deprecated from version 1.95.0. Use `metricDimensions` instead.
	//
	// Deprecated: Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead.
	Dimensions pulumi.MapInput
	// The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrInput
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	EndTime pulumi.IntPtrInput
	// A configuration of critical alarm (documented below).
	EscalationsCritical AlarmEscalationsCriticalPtrInput
	// A configuration of critical info (documented below).
	EscalationsInfo AlarmEscalationsInfoPtrInput
	// A configuration of critical warn (documented below).
	EscalationsWarn AlarmEscalationsWarnPtrInput
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringPtrInput
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	MetricDimensions AlarmMetricDimensionArrayInput
	// The alarm rule name.
	Name pulumi.StringPtrInput
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
	//
	// Deprecated: Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
	Operator pulumi.StringPtrInput
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrInput
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringPtrInput
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	StartTime pulumi.IntPtrInput
	// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Availability","Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
	//
	// Deprecated: Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
	Statistics pulumi.StringPtrInput
	// The current alarm rule status.
	Status pulumi.StringPtrInput
	// Critical level alarm threshold value, which must be a numeric value currently.
	//
	// Deprecated: Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
	Threshold pulumi.StringPtrInput
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
	//
	// Deprecated: Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
	TriggeredCount pulumi.IntPtrInput
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrInput
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups []string `pulumi:"contactGroups"`
	// Field `dimensions` has been deprecated from version 1.95.0. Use `metricDimensions` instead.
	//
	// Deprecated: Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead.
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled *bool `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	EndTime *int `pulumi:"endTime"`
	// A configuration of critical alarm (documented below).
	EscalationsCritical *AlarmEscalationsCritical `pulumi:"escalationsCritical"`
	// A configuration of critical info (documented below).
	EscalationsInfo *AlarmEscalationsInfo `pulumi:"escalationsInfo"`
	// A configuration of critical warn (documented below).
	EscalationsWarn *AlarmEscalationsWarn `pulumi:"escalationsWarn"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric string `pulumi:"metric"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	MetricDimensions []AlarmMetricDimension `pulumi:"metricDimensions"`
	// The alarm rule name.
	Name *string `pulumi:"name"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
	//
	// Deprecated: Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
	Operator *string `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period *int `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project string `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime *int `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	StartTime *int `pulumi:"startTime"`
	// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Availability","Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
	//
	// Deprecated: Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
	Statistics *string `pulumi:"statistics"`
	// Critical level alarm threshold value, which must be a numeric value currently.
	//
	// Deprecated: Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
	Threshold *string `pulumi:"threshold"`
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
	//
	// Deprecated: Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
	TriggeredCount *int `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook *string `pulumi:"webhook"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayInput
	// Field `dimensions` has been deprecated from version 1.95.0. Use `metricDimensions` instead.
	//
	// Deprecated: Field 'dimensions' has been deprecated from version 1.173.0. Use 'metric_dimensions' instead.
	Dimensions pulumi.MapInput
	// The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrInput
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	EndTime pulumi.IntPtrInput
	// A configuration of critical alarm (documented below).
	EscalationsCritical AlarmEscalationsCriticalPtrInput
	// A configuration of critical info (documented below).
	EscalationsInfo AlarmEscalationsInfoPtrInput
	// A configuration of critical warn (documented below).
	EscalationsWarn AlarmEscalationsWarnPtrInput
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringInput
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	MetricDimensions AlarmMetricDimensionArrayInput
	// The alarm rule name.
	Name pulumi.StringPtrInput
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
	//
	// Deprecated: Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
	Operator pulumi.StringPtrInput
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrInput
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringInput
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	//
	// Deprecated: Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
	StartTime pulumi.IntPtrInput
	// Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Availability","Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
	//
	// Deprecated: Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
	Statistics pulumi.StringPtrInput
	// Critical level alarm threshold value, which must be a numeric value currently.
	//
	// Deprecated: Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
	Threshold pulumi.StringPtrInput
	// It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
	//
	// Deprecated: Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
	TriggeredCount pulumi.IntPtrInput
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}

type AlarmInput interface {
	pulumi.Input

	ToAlarmOutput() AlarmOutput
	ToAlarmOutputWithContext(ctx context.Context) AlarmOutput
}

func (*Alarm) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

// AlarmArrayInput is an input type that accepts AlarmArray and AlarmArrayOutput values.
// You can construct a concrete instance of `AlarmArrayInput` via:
//
//          AlarmArray{ AlarmArgs{...} }
type AlarmArrayInput interface {
	pulumi.Input

	ToAlarmArrayOutput() AlarmArrayOutput
	ToAlarmArrayOutputWithContext(context.Context) AlarmArrayOutput
}

type AlarmArray []AlarmInput

func (AlarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (i AlarmArray) ToAlarmArrayOutput() AlarmArrayOutput {
	return i.ToAlarmArrayOutputWithContext(context.Background())
}

func (i AlarmArray) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmArrayOutput)
}

// AlarmMapInput is an input type that accepts AlarmMap and AlarmMapOutput values.
// You can construct a concrete instance of `AlarmMapInput` via:
//
//          AlarmMap{ "key": AlarmArgs{...} }
type AlarmMapInput interface {
	pulumi.Input

	ToAlarmMapOutput() AlarmMapOutput
	ToAlarmMapOutputWithContext(context.Context) AlarmMapOutput
}

type AlarmMap map[string]AlarmInput

func (AlarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (i AlarmMap) ToAlarmMapOutput() AlarmMapOutput {
	return i.ToAlarmMapOutputWithContext(context.Background())
}

func (i AlarmMap) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmMapOutput)
}

type AlarmOutput struct{ *pulumi.OutputState }

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

type AlarmArrayOutput struct{ *pulumi.OutputState }

func (AlarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alarm)(nil)).Elem()
}

func (o AlarmArrayOutput) ToAlarmArrayOutput() AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) ToAlarmArrayOutputWithContext(ctx context.Context) AlarmArrayOutput {
	return o
}

func (o AlarmArrayOutput) Index(i pulumi.IntInput) AlarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].([]*Alarm)[vs[1].(int)]
	}).(AlarmOutput)
}

type AlarmMapOutput struct{ *pulumi.OutputState }

func (AlarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alarm)(nil)).Elem()
}

func (o AlarmMapOutput) ToAlarmMapOutput() AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) ToAlarmMapOutputWithContext(ctx context.Context) AlarmMapOutput {
	return o
}

func (o AlarmMapOutput) MapIndex(k pulumi.StringInput) AlarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alarm {
		return vs[0].(map[string]*Alarm)[vs[1].(string)]
	}).(AlarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmInput)(nil)).Elem(), &Alarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmArrayInput)(nil)).Elem(), AlarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmMapInput)(nil)).Elem(), AlarmMap{})
	pulumi.RegisterOutputType(AlarmOutput{})
	pulumi.RegisterOutputType(AlarmArrayOutput{})
	pulumi.RegisterOutputType(AlarmMapOutput{})
}
