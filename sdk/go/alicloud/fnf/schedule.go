// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fnf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Serverless Workflow Schedule resource.
//
// For information about Serverless Workflow Schedule and how to use it, see [What is Schedule](https://www.alibabacloud.com/help/en/doc-detail/168934.htm).
//
// > **NOTE:** Available in v1.105.0+.
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
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fnf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleFlow, err := fnf.NewFlow(ctx, "exampleFlow", &fnf.FlowArgs{
// 			Definition:  pulumi.String(fmt.Sprintf("%v%v%v%v%v", "  version: v1beta1\n", "  type: flow\n", "  steps:\n", "    - type: pass\n", "      name: helloworld\n")),
// 			Description: pulumi.String("tf-testaccFnFFlow983041"),
// 			Type:        pulumi.String("FDL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fnf.NewSchedule(ctx, "exampleSchedule", &fnf.ScheduleArgs{
// 			CronExpression: pulumi.String("30 9 * * * *"),
// 			Description:    pulumi.String("tf-testaccFnFSchedule983041"),
// 			Enable:         pulumi.Bool(true),
// 			FlowName:       exampleFlow.Name,
// 			Payload:        pulumi.String("{\"tf-test\": \"test success\"}"),
// 			ScheduleName:   pulumi.String("tf-testaccFnFSchedule983041"),
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
// Serverless Workflow Schedule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:fnf/schedule:Schedule example <schedule_name>:<flow_name>
// ```
type Schedule struct {
	pulumi.CustomResourceState

	// The CRON expression of the time-based schedule to be created.
	CronExpression pulumi.StringOutput `pulumi:"cronExpression"`
	// The description of the time-based schedule to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// The name of the flow bound to the time-based schedule you want to create.
	FlowName pulumi.StringOutput `pulumi:"flowName"`
	// The time when the time-based schedule was last updated.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The trigger message of the time-based schedule to be created. It must be in JSON object format.
	Payload pulumi.StringPtrOutput `pulumi:"payload"`
	// The ID of the time-based schedule.
	ScheduleId pulumi.StringOutput `pulumi:"scheduleId"`
	// The name of the time-based schedule to be created.
	ScheduleName pulumi.StringOutput `pulumi:"scheduleName"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CronExpression == nil {
		return nil, errors.New("invalid value for required argument 'CronExpression'")
	}
	if args.FlowName == nil {
		return nil, errors.New("invalid value for required argument 'FlowName'")
	}
	if args.ScheduleName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleName'")
	}
	var resource Schedule
	err := ctx.RegisterResource("alicloud:fnf/schedule:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("alicloud:fnf/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	// The CRON expression of the time-based schedule to be created.
	CronExpression *string `pulumi:"cronExpression"`
	// The description of the time-based schedule to be created.
	Description *string `pulumi:"description"`
	// Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
	Enable *bool `pulumi:"enable"`
	// The name of the flow bound to the time-based schedule you want to create.
	FlowName *string `pulumi:"flowName"`
	// The time when the time-based schedule was last updated.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The trigger message of the time-based schedule to be created. It must be in JSON object format.
	Payload *string `pulumi:"payload"`
	// The ID of the time-based schedule.
	ScheduleId *string `pulumi:"scheduleId"`
	// The name of the time-based schedule to be created.
	ScheduleName *string `pulumi:"scheduleName"`
}

type ScheduleState struct {
	// The CRON expression of the time-based schedule to be created.
	CronExpression pulumi.StringPtrInput
	// The description of the time-based schedule to be created.
	Description pulumi.StringPtrInput
	// Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
	Enable pulumi.BoolPtrInput
	// The name of the flow bound to the time-based schedule you want to create.
	FlowName pulumi.StringPtrInput
	// The time when the time-based schedule was last updated.
	LastModifiedTime pulumi.StringPtrInput
	// The trigger message of the time-based schedule to be created. It must be in JSON object format.
	Payload pulumi.StringPtrInput
	// The ID of the time-based schedule.
	ScheduleId pulumi.StringPtrInput
	// The name of the time-based schedule to be created.
	ScheduleName pulumi.StringPtrInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// The CRON expression of the time-based schedule to be created.
	CronExpression string `pulumi:"cronExpression"`
	// The description of the time-based schedule to be created.
	Description *string `pulumi:"description"`
	// Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
	Enable *bool `pulumi:"enable"`
	// The name of the flow bound to the time-based schedule you want to create.
	FlowName string `pulumi:"flowName"`
	// The trigger message of the time-based schedule to be created. It must be in JSON object format.
	Payload *string `pulumi:"payload"`
	// The name of the time-based schedule to be created.
	ScheduleName string `pulumi:"scheduleName"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The CRON expression of the time-based schedule to be created.
	CronExpression pulumi.StringInput
	// The description of the time-based schedule to be created.
	Description pulumi.StringPtrInput
	// Specifies whether to enable the time-based schedule you want to create. Valid values: `false`, `true`.
	Enable pulumi.BoolPtrInput
	// The name of the flow bound to the time-based schedule you want to create.
	FlowName pulumi.StringInput
	// The trigger message of the time-based schedule to be created. It must be in JSON object format.
	Payload pulumi.StringPtrInput
	// The name of the time-based schedule to be created.
	ScheduleName pulumi.StringInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (*Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil))
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i *Schedule) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *Schedule) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil))
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

// ScheduleArrayInput is an input type that accepts ScheduleArray and ScheduleArrayOutput values.
// You can construct a concrete instance of `ScheduleArrayInput` via:
//
//          ScheduleArray{ ScheduleArgs{...} }
type ScheduleArrayInput interface {
	pulumi.Input

	ToScheduleArrayOutput() ScheduleArrayOutput
	ToScheduleArrayOutputWithContext(context.Context) ScheduleArrayOutput
}

type ScheduleArray []ScheduleInput

func (ScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schedule)(nil)).Elem()
}

func (i ScheduleArray) ToScheduleArrayOutput() ScheduleArrayOutput {
	return i.ToScheduleArrayOutputWithContext(context.Background())
}

func (i ScheduleArray) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleArrayOutput)
}

// ScheduleMapInput is an input type that accepts ScheduleMap and ScheduleMapOutput values.
// You can construct a concrete instance of `ScheduleMapInput` via:
//
//          ScheduleMap{ "key": ScheduleArgs{...} }
type ScheduleMapInput interface {
	pulumi.Input

	ToScheduleMapOutput() ScheduleMapOutput
	ToScheduleMapOutputWithContext(context.Context) ScheduleMapOutput
}

type ScheduleMap map[string]ScheduleInput

func (ScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schedule)(nil)).Elem()
}

func (i ScheduleMap) ToScheduleMapOutput() ScheduleMapOutput {
	return i.ToScheduleMapOutputWithContext(context.Background())
}

func (i ScheduleMap) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMapOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil))
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil))
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

type ScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schedule)(nil))
}

func (o ScheduleArrayOutput) ToScheduleArrayOutput() ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) Index(i pulumi.IntInput) ScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Schedule {
		return vs[0].([]Schedule)[vs[1].(int)]
	}).(ScheduleOutput)
}

type ScheduleMapOutput struct{ *pulumi.OutputState }

func (ScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Schedule)(nil))
}

func (o ScheduleMapOutput) ToScheduleMapOutput() ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) MapIndex(k pulumi.StringInput) ScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Schedule {
		return vs[0].(map[string]Schedule)[vs[1].(string)]
	}).(ScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleInput)(nil)).Elem(), &Schedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchedulePtrInput)(nil)).Elem(), &Schedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleArrayInput)(nil)).Elem(), ScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleMapInput)(nil)).Elem(), ScheduleMap{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScheduleMapOutput{})
}
