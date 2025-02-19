// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ESS alarm task resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ess"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				MostRecent: pulumi.BoolRef(true),
//				NameRegex:  pulumi.StringRef("^centos_6\\w{1,5}[64].*"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(defaultZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			fooNetwork, err := vpc.NewNetwork(ctx, "fooNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSwitch, err := vpc.NewSwitch(ctx, "fooSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("tf-testAccEssAlarm_basic_foo"),
//				VpcId:       fooNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := vpc.NewSwitch(ctx, "bar", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("tf-testAccEssAlarm_basic_bar"),
//				VpcId:       fooNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.1.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			fooScalingGroup, err := ess.NewScalingGroup(ctx, "fooScalingGroup", &ess.ScalingGroupArgs{
//				MinSize:          pulumi.Int(1),
//				MaxSize:          pulumi.Int(1),
//				ScalingGroupName: pulumi.String("tf-testAccEssAlarm_basic"),
//				RemovalPolicies: pulumi.StringArray{
//					pulumi.String("OldestInstance"),
//					pulumi.String("NewestInstance"),
//				},
//				VswitchIds: pulumi.StringArray{
//					fooSwitch.ID(),
//					bar.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooScalingRule, err := ess.NewScalingRule(ctx, "fooScalingRule", &ess.ScalingRuleArgs{
//				ScalingRuleName: pulumi.String("tf-testAccEssAlarm_basic"),
//				ScalingGroupId:  fooScalingGroup.ID(),
//				AdjustmentType:  pulumi.String("TotalCapacity"),
//				AdjustmentValue: pulumi.Int(2),
//				Cooldown:        pulumi.Int(60),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ess.NewAlarm(ctx, "fooAlarm", &ess.AlarmArgs{
//				Description: pulumi.String("Acc alarm test"),
//				AlarmActions: pulumi.StringArray{
//					fooScalingRule.Ari,
//				},
//				ScalingGroupId:     fooScalingGroup.ID(),
//				MetricType:         pulumi.String("system"),
//				MetricName:         pulumi.String("CpuUtilization"),
//				Period:             pulumi.Int(300),
//				Statistics:         pulumi.String("Average"),
//				Threshold:          pulumi.String("200.3"),
//				ComparisonOperator: pulumi.String(">="),
//				EvaluationCount:    pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Module Support
//
// You can use to the existing autoscaling-rule module
// to create alarm task, different type rules and scheduled task one-click.
//
// ## Block metricNames_and_dimensions
//
// Supported metric names and dimensions :
//
// | MetricName         | Dimensions                   |
// | ------------------ | ---------------------------- |
// | CpuUtilization     | user_id,scaling_group        |
// | ClassicInternetRx  | user_id,scaling_group        |
// | ClassicInternetTx  | user_id,scaling_group        |
// | VpcInternetRx      | user_id,scaling_group        |
// | VpcInternetTx      | user_id,scaling_group        |
// | IntranetRx         | user_id,scaling_group        |
// | IntranetTx         | user_id,scaling_group        |
// | LoadAverage        | user_id,scaling_group        |
// | MemoryUtilization  | user_id,scaling_group        |
// | SystemDiskReadBps  | user_id,scaling_group        |
// | SystemDiskWriteBps | user_id,scaling_group        |
// | SystemDiskReadOps  | user_id,scaling_group        |
// | SystemDiskWriteOps | user_id,scaling_group        |
// | PackagesNetIn      | user_id,scaling_group,device |
// | PackagesNetOut     | user_id,scaling_group,device |
// | TcpConnection      | user_id,scaling_group,state  |
//
// > **NOTE:** Dimension `userId` and `scalingGroup` is automatically filled, which means you only need to care about dimension `device` and `state` when needed.
//
// ## Import
//
// Ess alarm can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ess/alarm:Alarm example asg-2ze500_045efffe-4d05
//
// ```
type Alarm struct {
	pulumi.CustomResourceState

	// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
	AlarmActions pulumi.StringArrayOutput `pulumi:"alarmActions"`
	// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
	CloudMonitorGroupId pulumi.IntPtrOutput `pulumi:"cloudMonitorGroupId"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
	ComparisonOperator pulumi.StringPtrOutput `pulumi:"comparisonOperator"`
	// The description for the alarm.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
	Dimensions pulumi.MapOutput `pulumi:"dimensions"`
	// Whether to enable specific ess alarm. Default to true.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
	EvaluationCount pulumi.IntPtrOutput `pulumi:"evaluationCount"`
	// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
	MetricType pulumi.StringPtrOutput `pulumi:"metricType"`
	// The name for ess alarm.
	Name pulumi.StringOutput `pulumi:"name"`
	// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// The state of specified alarm.
	State pulumi.StringOutput `pulumi:"state"`
	// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
	Statistics pulumi.StringPtrOutput `pulumi:"statistics"`
	// The value against which the specified statistics is compared.
	Threshold pulumi.StringOutput `pulumi:"threshold"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlarmActions == nil {
		return nil, errors.New("invalid value for required argument 'AlarmActions'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	if args.Threshold == nil {
		return nil, errors.New("invalid value for required argument 'Threshold'")
	}
	var resource Alarm
	err := ctx.RegisterResource("alicloud:ess/alarm:Alarm", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:ess/alarm:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
	// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
	AlarmActions []string `pulumi:"alarmActions"`
	// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
	CloudMonitorGroupId *int `pulumi:"cloudMonitorGroupId"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
	ComparisonOperator *string `pulumi:"comparisonOperator"`
	// The description for the alarm.
	Description *string `pulumi:"description"`
	// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// Whether to enable specific ess alarm. Default to true.
	Enable *bool `pulumi:"enable"`
	// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
	EvaluationCount *int `pulumi:"evaluationCount"`
	// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
	MetricName *string `pulumi:"metricName"`
	// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
	MetricType *string `pulumi:"metricType"`
	// The name for ess alarm.
	Name *string `pulumi:"name"`
	// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
	Period *int `pulumi:"period"`
	// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The state of specified alarm.
	State *string `pulumi:"state"`
	// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
	Statistics *string `pulumi:"statistics"`
	// The value against which the specified statistics is compared.
	Threshold *string `pulumi:"threshold"`
}

type AlarmState struct {
	// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
	AlarmActions pulumi.StringArrayInput
	// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
	CloudMonitorGroupId pulumi.IntPtrInput
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
	ComparisonOperator pulumi.StringPtrInput
	// The description for the alarm.
	Description pulumi.StringPtrInput
	// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
	Dimensions pulumi.MapInput
	// Whether to enable specific ess alarm. Default to true.
	Enable pulumi.BoolPtrInput
	// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
	EvaluationCount pulumi.IntPtrInput
	// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
	MetricName pulumi.StringPtrInput
	// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
	MetricType pulumi.StringPtrInput
	// The name for ess alarm.
	Name pulumi.StringPtrInput
	// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
	Period pulumi.IntPtrInput
	// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
	ScalingGroupId pulumi.StringPtrInput
	// The state of specified alarm.
	State pulumi.StringPtrInput
	// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
	Statistics pulumi.StringPtrInput
	// The value against which the specified statistics is compared.
	Threshold pulumi.StringPtrInput
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
	AlarmActions []string `pulumi:"alarmActions"`
	// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
	CloudMonitorGroupId *int `pulumi:"cloudMonitorGroupId"`
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
	ComparisonOperator *string `pulumi:"comparisonOperator"`
	// The description for the alarm.
	Description *string `pulumi:"description"`
	// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// Whether to enable specific ess alarm. Default to true.
	Enable *bool `pulumi:"enable"`
	// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
	EvaluationCount *int `pulumi:"evaluationCount"`
	// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
	MetricName string `pulumi:"metricName"`
	// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
	MetricType *string `pulumi:"metricType"`
	// The name for ess alarm.
	Name *string `pulumi:"name"`
	// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
	Period *int `pulumi:"period"`
	// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
	Statistics *string `pulumi:"statistics"`
	// The value against which the specified statistics is compared.
	Threshold string `pulumi:"threshold"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
	AlarmActions pulumi.StringArrayInput
	// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
	CloudMonitorGroupId pulumi.IntPtrInput
	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
	ComparisonOperator pulumi.StringPtrInput
	// The description for the alarm.
	Description pulumi.StringPtrInput
	// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
	Dimensions pulumi.MapInput
	// Whether to enable specific ess alarm. Default to true.
	Enable pulumi.BoolPtrInput
	// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
	EvaluationCount pulumi.IntPtrInput
	// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
	MetricName pulumi.StringInput
	// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
	MetricType pulumi.StringPtrInput
	// The name for ess alarm.
	Name pulumi.StringPtrInput
	// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
	Period pulumi.IntPtrInput
	// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
	ScalingGroupId pulumi.StringInput
	// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
	Statistics pulumi.StringPtrInput
	// The value against which the specified statistics is compared.
	Threshold pulumi.StringInput
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
//	AlarmArray{ AlarmArgs{...} }
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
//	AlarmMap{ "key": AlarmArgs{...} }
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

// The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
func (o AlarmOutput) AlarmActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringArrayOutput { return v.AlarmActions }).(pulumi.StringArrayOutput)
}

// Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
func (o AlarmOutput) CloudMonitorGroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.CloudMonitorGroupId }).(pulumi.IntPtrOutput)
}

// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
func (o AlarmOutput) ComparisonOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.ComparisonOperator }).(pulumi.StringPtrOutput)
}

// The description for the alarm.
func (o AlarmOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
func (o AlarmOutput) Dimensions() pulumi.MapOutput {
	return o.ApplyT(func(v *Alarm) pulumi.MapOutput { return v.Dimensions }).(pulumi.MapOutput)
}

// Whether to enable specific ess alarm. Default to true.
func (o AlarmOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
func (o AlarmOutput) EvaluationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.EvaluationCount }).(pulumi.IntPtrOutput)
}

// The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
func (o AlarmOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

// The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
func (o AlarmOutput) MetricType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.MetricType }).(pulumi.StringPtrOutput)
}

// The name for ess alarm.
func (o AlarmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
func (o AlarmOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
func (o AlarmOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The state of specified alarm.
func (o AlarmOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
func (o AlarmOutput) Statistics() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.Statistics }).(pulumi.StringPtrOutput)
}

// The value against which the specified statistics is compared.
func (o AlarmOutput) Threshold() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Threshold }).(pulumi.StringOutput)
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
