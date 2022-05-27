// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Metric Rule Template resource.
//
// For information about Cloud Monitor Service Metric Rule Template and how to use it, see [What is Metric Rule Template](https://www.alibabacloud.com/help/doc-detail/114984.html).
//
// > **NOTE:** Available in v1.134.0+.
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
// 		_, err := cms.NewMetricRuleTemplate(ctx, "example", &cms.MetricRuleTemplateArgs{
// 			AlertTemplates: cms.MetricRuleTemplateAlertTemplateArray{
// 				&cms.MetricRuleTemplateAlertTemplateArgs{
// 					Category: pulumi.String("ecs"),
// 					Escalations: &cms.MetricRuleTemplateAlertTemplateEscalationsArgs{
// 						Critical: &cms.MetricRuleTemplateAlertTemplateEscalationsCriticalArgs{
// 							ComparisonOperator: pulumi.String("GreaterThanThreshold"),
// 							Statistics:         pulumi.String("Average"),
// 							Threshold:          pulumi.String("90"),
// 							Times:              pulumi.String("3"),
// 						},
// 					},
// 					MetricName: pulumi.String("cpu_total"),
// 					Namespace:  pulumi.String("acs_ecs_dashboard"),
// 					RuleName:   pulumi.String("tf_testAcc_new"),
// 				},
// 			},
// 			MetricRuleTemplateName: pulumi.String("example_value"),
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
// Cloud Monitor Service Metric Rule Template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cms/metricRuleTemplate:MetricRuleTemplate example <id>
// ```
type MetricRuleTemplate struct {
	pulumi.CustomResourceState

	// The details of alert rules that are generated based on the alert template. See the following `Block alertTemplates`.
	AlertTemplates MetricRuleTemplateAlertTemplateArrayOutput `pulumi:"alertTemplates"`
	// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	ApplyMode pulumi.StringPtrOutput `pulumi:"applyMode"`
	// The description of the alert template.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
	EnableEndTime pulumi.StringPtrOutput `pulumi:"enableEndTime"`
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
	EnableStartTime pulumi.StringPtrOutput `pulumi:"enableStartTime"`
	// The ID of the application group.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The name of the alert template.
	MetricRuleTemplateName pulumi.StringOutput `pulumi:"metricRuleTemplateName"`
	// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	NotifyLevel pulumi.StringPtrOutput `pulumi:"notifyLevel"`
	// The version of the alert template to be modified.
	RestVersion pulumi.StringOutput `pulumi:"restVersion"`
	// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
	SilenceTime pulumi.IntPtrOutput `pulumi:"silenceTime"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	Webhook pulumi.StringPtrOutput `pulumi:"webhook"`
}

// NewMetricRuleTemplate registers a new resource with the given unique name, arguments, and options.
func NewMetricRuleTemplate(ctx *pulumi.Context,
	name string, args *MetricRuleTemplateArgs, opts ...pulumi.ResourceOption) (*MetricRuleTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricRuleTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'MetricRuleTemplateName'")
	}
	var resource MetricRuleTemplate
	err := ctx.RegisterResource("alicloud:cms/metricRuleTemplate:MetricRuleTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricRuleTemplate gets an existing MetricRuleTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricRuleTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricRuleTemplateState, opts ...pulumi.ResourceOption) (*MetricRuleTemplate, error) {
	var resource MetricRuleTemplate
	err := ctx.ReadResource("alicloud:cms/metricRuleTemplate:MetricRuleTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricRuleTemplate resources.
type metricRuleTemplateState struct {
	// The details of alert rules that are generated based on the alert template. See the following `Block alertTemplates`.
	AlertTemplates []MetricRuleTemplateAlertTemplate `pulumi:"alertTemplates"`
	// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	ApplyMode *string `pulumi:"applyMode"`
	// The description of the alert template.
	Description *string `pulumi:"description"`
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
	EnableEndTime *string `pulumi:"enableEndTime"`
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
	EnableStartTime *string `pulumi:"enableStartTime"`
	// The ID of the application group.
	GroupId *string `pulumi:"groupId"`
	// The name of the alert template.
	MetricRuleTemplateName *string `pulumi:"metricRuleTemplateName"`
	// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	NotifyLevel *string `pulumi:"notifyLevel"`
	// The version of the alert template to be modified.
	RestVersion *string `pulumi:"restVersion"`
	// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
	SilenceTime *int `pulumi:"silenceTime"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	Webhook *string `pulumi:"webhook"`
}

type MetricRuleTemplateState struct {
	// The details of alert rules that are generated based on the alert template. See the following `Block alertTemplates`.
	AlertTemplates MetricRuleTemplateAlertTemplateArrayInput
	// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	ApplyMode pulumi.StringPtrInput
	// The description of the alert template.
	Description pulumi.StringPtrInput
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
	EnableEndTime pulumi.StringPtrInput
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
	EnableStartTime pulumi.StringPtrInput
	// The ID of the application group.
	GroupId pulumi.StringPtrInput
	// The name of the alert template.
	MetricRuleTemplateName pulumi.StringPtrInput
	// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	NotifyLevel pulumi.StringPtrInput
	// The version of the alert template to be modified.
	RestVersion pulumi.StringPtrInput
	// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
	SilenceTime pulumi.IntPtrInput
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	Webhook pulumi.StringPtrInput
}

func (MetricRuleTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricRuleTemplateState)(nil)).Elem()
}

type metricRuleTemplateArgs struct {
	// The details of alert rules that are generated based on the alert template. See the following `Block alertTemplates`.
	AlertTemplates []MetricRuleTemplateAlertTemplate `pulumi:"alertTemplates"`
	// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	ApplyMode *string `pulumi:"applyMode"`
	// The description of the alert template.
	Description *string `pulumi:"description"`
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
	EnableEndTime *string `pulumi:"enableEndTime"`
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
	EnableStartTime *string `pulumi:"enableStartTime"`
	// The ID of the application group.
	GroupId *string `pulumi:"groupId"`
	// The name of the alert template.
	MetricRuleTemplateName string `pulumi:"metricRuleTemplateName"`
	// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	NotifyLevel *string `pulumi:"notifyLevel"`
	// The version of the alert template to be modified.
	RestVersion *string `pulumi:"restVersion"`
	// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
	SilenceTime *int `pulumi:"silenceTime"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	Webhook *string `pulumi:"webhook"`
}

// The set of arguments for constructing a MetricRuleTemplate resource.
type MetricRuleTemplateArgs struct {
	// The details of alert rules that are generated based on the alert template. See the following `Block alertTemplates`.
	AlertTemplates MetricRuleTemplateAlertTemplateArrayInput
	// The mode in which the alert template is applied. Valid values:`GROUP_INSTANCE_FIRST`or `ALARM_TEMPLATE_FIRST`. GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template. ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	ApplyMode pulumi.StringPtrInput
	// The description of the alert template.
	Description pulumi.StringPtrInput
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:59 and the value 23 indicates 23:59.
	EnableEndTime pulumi.StringPtrInput
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. The value 00 indicates 00:00 and the value 23 indicates 23:00.
	EnableStartTime pulumi.StringPtrInput
	// The ID of the application group.
	GroupId pulumi.StringPtrInput
	// The name of the alert template.
	MetricRuleTemplateName pulumi.StringInput
	// The alert notification method. Valid values:Set the value to 4. The value 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	NotifyLevel pulumi.StringPtrInput
	// The version of the alert template to be modified.
	RestVersion pulumi.StringPtrInput
	// The mute period during which notifications are not repeatedly sent for an alert.Valid values: 0 to 86400. Unit: seconds. Default value: `86400`.
	SilenceTime pulumi.IntPtrInput
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	Webhook pulumi.StringPtrInput
}

func (MetricRuleTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricRuleTemplateArgs)(nil)).Elem()
}

type MetricRuleTemplateInput interface {
	pulumi.Input

	ToMetricRuleTemplateOutput() MetricRuleTemplateOutput
	ToMetricRuleTemplateOutputWithContext(ctx context.Context) MetricRuleTemplateOutput
}

func (*MetricRuleTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricRuleTemplate)(nil)).Elem()
}

func (i *MetricRuleTemplate) ToMetricRuleTemplateOutput() MetricRuleTemplateOutput {
	return i.ToMetricRuleTemplateOutputWithContext(context.Background())
}

func (i *MetricRuleTemplate) ToMetricRuleTemplateOutputWithContext(ctx context.Context) MetricRuleTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRuleTemplateOutput)
}

// MetricRuleTemplateArrayInput is an input type that accepts MetricRuleTemplateArray and MetricRuleTemplateArrayOutput values.
// You can construct a concrete instance of `MetricRuleTemplateArrayInput` via:
//
//          MetricRuleTemplateArray{ MetricRuleTemplateArgs{...} }
type MetricRuleTemplateArrayInput interface {
	pulumi.Input

	ToMetricRuleTemplateArrayOutput() MetricRuleTemplateArrayOutput
	ToMetricRuleTemplateArrayOutputWithContext(context.Context) MetricRuleTemplateArrayOutput
}

type MetricRuleTemplateArray []MetricRuleTemplateInput

func (MetricRuleTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricRuleTemplate)(nil)).Elem()
}

func (i MetricRuleTemplateArray) ToMetricRuleTemplateArrayOutput() MetricRuleTemplateArrayOutput {
	return i.ToMetricRuleTemplateArrayOutputWithContext(context.Background())
}

func (i MetricRuleTemplateArray) ToMetricRuleTemplateArrayOutputWithContext(ctx context.Context) MetricRuleTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRuleTemplateArrayOutput)
}

// MetricRuleTemplateMapInput is an input type that accepts MetricRuleTemplateMap and MetricRuleTemplateMapOutput values.
// You can construct a concrete instance of `MetricRuleTemplateMapInput` via:
//
//          MetricRuleTemplateMap{ "key": MetricRuleTemplateArgs{...} }
type MetricRuleTemplateMapInput interface {
	pulumi.Input

	ToMetricRuleTemplateMapOutput() MetricRuleTemplateMapOutput
	ToMetricRuleTemplateMapOutputWithContext(context.Context) MetricRuleTemplateMapOutput
}

type MetricRuleTemplateMap map[string]MetricRuleTemplateInput

func (MetricRuleTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricRuleTemplate)(nil)).Elem()
}

func (i MetricRuleTemplateMap) ToMetricRuleTemplateMapOutput() MetricRuleTemplateMapOutput {
	return i.ToMetricRuleTemplateMapOutputWithContext(context.Background())
}

func (i MetricRuleTemplateMap) ToMetricRuleTemplateMapOutputWithContext(ctx context.Context) MetricRuleTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRuleTemplateMapOutput)
}

type MetricRuleTemplateOutput struct{ *pulumi.OutputState }

func (MetricRuleTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricRuleTemplate)(nil)).Elem()
}

func (o MetricRuleTemplateOutput) ToMetricRuleTemplateOutput() MetricRuleTemplateOutput {
	return o
}

func (o MetricRuleTemplateOutput) ToMetricRuleTemplateOutputWithContext(ctx context.Context) MetricRuleTemplateOutput {
	return o
}

type MetricRuleTemplateArrayOutput struct{ *pulumi.OutputState }

func (MetricRuleTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricRuleTemplate)(nil)).Elem()
}

func (o MetricRuleTemplateArrayOutput) ToMetricRuleTemplateArrayOutput() MetricRuleTemplateArrayOutput {
	return o
}

func (o MetricRuleTemplateArrayOutput) ToMetricRuleTemplateArrayOutputWithContext(ctx context.Context) MetricRuleTemplateArrayOutput {
	return o
}

func (o MetricRuleTemplateArrayOutput) Index(i pulumi.IntInput) MetricRuleTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetricRuleTemplate {
		return vs[0].([]*MetricRuleTemplate)[vs[1].(int)]
	}).(MetricRuleTemplateOutput)
}

type MetricRuleTemplateMapOutput struct{ *pulumi.OutputState }

func (MetricRuleTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricRuleTemplate)(nil)).Elem()
}

func (o MetricRuleTemplateMapOutput) ToMetricRuleTemplateMapOutput() MetricRuleTemplateMapOutput {
	return o
}

func (o MetricRuleTemplateMapOutput) ToMetricRuleTemplateMapOutputWithContext(ctx context.Context) MetricRuleTemplateMapOutput {
	return o
}

func (o MetricRuleTemplateMapOutput) MapIndex(k pulumi.StringInput) MetricRuleTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetricRuleTemplate {
		return vs[0].(map[string]*MetricRuleTemplate)[vs[1].(string)]
	}).(MetricRuleTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRuleTemplateInput)(nil)).Elem(), &MetricRuleTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRuleTemplateArrayInput)(nil)).Elem(), MetricRuleTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRuleTemplateMapInput)(nil)).Elem(), MetricRuleTemplateMap{})
	pulumi.RegisterOutputType(MetricRuleTemplateOutput{})
	pulumi.RegisterOutputType(MetricRuleTemplateArrayOutput{})
	pulumi.RegisterOutputType(MetricRuleTemplateMapOutput{})
}
