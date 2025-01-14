// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Suspend/Resume processes to a specified scaling group.
//
// For information about scaling group suspend process, see [SuspendProcesses](https://www.alibabacloud.com/help/en/auto-scaling/latest/suspendprocesses).
//
// > NOTE: Available in v1.166.0+
//
// ## Import
//
// ### Timeouts The `timeouts` block allows you to specify timeouts for certain actions* `create` - (Defaults to 1 mins) Used when create the process. * `delete` - (Defaults to 1 mins) Used when delete the process.
type SuspendProcess struct {
	pulumi.CustomResourceState

	// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
	Process pulumi.StringOutput `pulumi:"process"`
	// ID of the scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
}

// NewSuspendProcess registers a new resource with the given unique name, arguments, and options.
func NewSuspendProcess(ctx *pulumi.Context,
	name string, args *SuspendProcessArgs, opts ...pulumi.ResourceOption) (*SuspendProcess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Process == nil {
		return nil, errors.New("invalid value for required argument 'Process'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	var resource SuspendProcess
	err := ctx.RegisterResource("alicloud:ess/suspendProcess:SuspendProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuspendProcess gets an existing SuspendProcess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuspendProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuspendProcessState, opts ...pulumi.ResourceOption) (*SuspendProcess, error) {
	var resource SuspendProcess
	err := ctx.ReadResource("alicloud:ess/suspendProcess:SuspendProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SuspendProcess resources.
type suspendProcessState struct {
	// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
	Process *string `pulumi:"process"`
	// ID of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

type SuspendProcessState struct {
	// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
	Process pulumi.StringPtrInput
	// ID of the scaling group.
	ScalingGroupId pulumi.StringPtrInput
}

func (SuspendProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*suspendProcessState)(nil)).Elem()
}

type suspendProcessArgs struct {
	// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
	Process string `pulumi:"process"`
	// ID of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// The set of arguments for constructing a SuspendProcess resource.
type SuspendProcessArgs struct {
	// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
	Process pulumi.StringInput
	// ID of the scaling group.
	ScalingGroupId pulumi.StringInput
}

func (SuspendProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suspendProcessArgs)(nil)).Elem()
}

type SuspendProcessInput interface {
	pulumi.Input

	ToSuspendProcessOutput() SuspendProcessOutput
	ToSuspendProcessOutputWithContext(ctx context.Context) SuspendProcessOutput
}

func (*SuspendProcess) ElementType() reflect.Type {
	return reflect.TypeOf((**SuspendProcess)(nil)).Elem()
}

func (i *SuspendProcess) ToSuspendProcessOutput() SuspendProcessOutput {
	return i.ToSuspendProcessOutputWithContext(context.Background())
}

func (i *SuspendProcess) ToSuspendProcessOutputWithContext(ctx context.Context) SuspendProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuspendProcessOutput)
}

// SuspendProcessArrayInput is an input type that accepts SuspendProcessArray and SuspendProcessArrayOutput values.
// You can construct a concrete instance of `SuspendProcessArrayInput` via:
//
//	SuspendProcessArray{ SuspendProcessArgs{...} }
type SuspendProcessArrayInput interface {
	pulumi.Input

	ToSuspendProcessArrayOutput() SuspendProcessArrayOutput
	ToSuspendProcessArrayOutputWithContext(context.Context) SuspendProcessArrayOutput
}

type SuspendProcessArray []SuspendProcessInput

func (SuspendProcessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SuspendProcess)(nil)).Elem()
}

func (i SuspendProcessArray) ToSuspendProcessArrayOutput() SuspendProcessArrayOutput {
	return i.ToSuspendProcessArrayOutputWithContext(context.Background())
}

func (i SuspendProcessArray) ToSuspendProcessArrayOutputWithContext(ctx context.Context) SuspendProcessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuspendProcessArrayOutput)
}

// SuspendProcessMapInput is an input type that accepts SuspendProcessMap and SuspendProcessMapOutput values.
// You can construct a concrete instance of `SuspendProcessMapInput` via:
//
//	SuspendProcessMap{ "key": SuspendProcessArgs{...} }
type SuspendProcessMapInput interface {
	pulumi.Input

	ToSuspendProcessMapOutput() SuspendProcessMapOutput
	ToSuspendProcessMapOutputWithContext(context.Context) SuspendProcessMapOutput
}

type SuspendProcessMap map[string]SuspendProcessInput

func (SuspendProcessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SuspendProcess)(nil)).Elem()
}

func (i SuspendProcessMap) ToSuspendProcessMapOutput() SuspendProcessMapOutput {
	return i.ToSuspendProcessMapOutputWithContext(context.Background())
}

func (i SuspendProcessMap) ToSuspendProcessMapOutputWithContext(ctx context.Context) SuspendProcessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuspendProcessMapOutput)
}

type SuspendProcessOutput struct{ *pulumi.OutputState }

func (SuspendProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuspendProcess)(nil)).Elem()
}

func (o SuspendProcessOutput) ToSuspendProcessOutput() SuspendProcessOutput {
	return o
}

func (o SuspendProcessOutput) ToSuspendProcessOutputWithContext(ctx context.Context) SuspendProcessOutput {
	return o
}

// Activity type N that you want to suspend. Valid values are: `SCALE_OUT`,`SCALE_IN`,`HealthCheck`,`AlarmNotification` and `ScheduledAction`.
func (o SuspendProcessOutput) Process() pulumi.StringOutput {
	return o.ApplyT(func(v *SuspendProcess) pulumi.StringOutput { return v.Process }).(pulumi.StringOutput)
}

// ID of the scaling group.
func (o SuspendProcessOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SuspendProcess) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

type SuspendProcessArrayOutput struct{ *pulumi.OutputState }

func (SuspendProcessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SuspendProcess)(nil)).Elem()
}

func (o SuspendProcessArrayOutput) ToSuspendProcessArrayOutput() SuspendProcessArrayOutput {
	return o
}

func (o SuspendProcessArrayOutput) ToSuspendProcessArrayOutputWithContext(ctx context.Context) SuspendProcessArrayOutput {
	return o
}

func (o SuspendProcessArrayOutput) Index(i pulumi.IntInput) SuspendProcessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SuspendProcess {
		return vs[0].([]*SuspendProcess)[vs[1].(int)]
	}).(SuspendProcessOutput)
}

type SuspendProcessMapOutput struct{ *pulumi.OutputState }

func (SuspendProcessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SuspendProcess)(nil)).Elem()
}

func (o SuspendProcessMapOutput) ToSuspendProcessMapOutput() SuspendProcessMapOutput {
	return o
}

func (o SuspendProcessMapOutput) ToSuspendProcessMapOutputWithContext(ctx context.Context) SuspendProcessMapOutput {
	return o
}

func (o SuspendProcessMapOutput) MapIndex(k pulumi.StringInput) SuspendProcessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SuspendProcess {
		return vs[0].(map[string]*SuspendProcess)[vs[1].(string)]
	}).(SuspendProcessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SuspendProcessInput)(nil)).Elem(), &SuspendProcess{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuspendProcessArrayInput)(nil)).Elem(), SuspendProcessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuspendProcessMapInput)(nil)).Elem(), SuspendProcessMap{})
	pulumi.RegisterOutputType(SuspendProcessOutput{})
	pulumi.RegisterOutputType(SuspendProcessArrayOutput{})
	pulumi.RegisterOutputType(SuspendProcessMapOutput{})
}
