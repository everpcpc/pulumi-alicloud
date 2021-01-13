// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// ESS schedule task can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ess/scheduledTask:ScheduledTask example abc123456
// ```
type ScheduledTask struct {
	pulumi.CustomResourceState

	// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
	Description pulumi.StringOutput `pulumi:"description"`
	// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
	DesiredCapacity pulumi.IntPtrOutput `pulumi:"desiredCapacity"`
	// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
	LaunchExpirationTime pulumi.IntPtrOutput `pulumi:"launchExpirationTime"`
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
	// If the `recurrenceType` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
	// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
	LaunchTime pulumi.StringPtrOutput `pulumi:"launchTime"`
	// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MaxValue pulumi.IntPtrOutput `pulumi:"maxValue"`
	// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MinValue pulumi.IntPtrOutput `pulumi:"minValue"`
	// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
	RecurrenceEndTime pulumi.StringOutput `pulumi:"recurrenceEndTime"`
	// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
	// - Daily: The scheduled task is executed once every specified number of days.
	// - Weekly: The scheduled task is executed on each specified day of a week.
	// - Monthly: The scheduled task is executed on each specified day of a month.
	// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
	RecurrenceType pulumi.StringOutput `pulumi:"recurrenceType"`
	// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrenceType`
	// - Daily: You can enter one value. Valid values: 1 to 31.
	// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
	// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
	// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
	RecurrenceValue pulumi.StringOutput `pulumi:"recurrenceValue"`
	// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScheduledAction pulumi.StringPtrOutput `pulumi:"scheduledAction"`
	// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
	ScheduledTaskName pulumi.StringPtrOutput `pulumi:"scheduledTaskName"`
	// Specifies whether to start the scheduled task. Default to true.
	TaskEnabled pulumi.BoolPtrOutput `pulumi:"taskEnabled"`
}

// NewScheduledTask registers a new resource with the given unique name, arguments, and options.
func NewScheduledTask(ctx *pulumi.Context,
	name string, args *ScheduledTaskArgs, opts ...pulumi.ResourceOption) (*ScheduledTask, error) {
	if args == nil {
		args = &ScheduledTaskArgs{}
	}

	var resource ScheduledTask
	err := ctx.RegisterResource("alicloud:ess/scheduledTask:ScheduledTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledTask gets an existing ScheduledTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledTaskState, opts ...pulumi.ResourceOption) (*ScheduledTask, error) {
	var resource ScheduledTask
	err := ctx.ReadResource("alicloud:ess/scheduledTask:ScheduledTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledTask resources.
type scheduledTaskState struct {
	// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
	Description *string `pulumi:"description"`
	// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
	LaunchExpirationTime *int `pulumi:"launchExpirationTime"`
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
	// If the `recurrenceType` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
	// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
	LaunchTime *string `pulumi:"launchTime"`
	// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MaxValue *int `pulumi:"maxValue"`
	// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MinValue *int `pulumi:"minValue"`
	// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
	RecurrenceEndTime *string `pulumi:"recurrenceEndTime"`
	// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
	// - Daily: The scheduled task is executed once every specified number of days.
	// - Weekly: The scheduled task is executed on each specified day of a week.
	// - Monthly: The scheduled task is executed on each specified day of a month.
	// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
	RecurrenceType *string `pulumi:"recurrenceType"`
	// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrenceType`
	// - Daily: You can enter one value. Valid values: 1 to 31.
	// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
	// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
	// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
	RecurrenceValue *string `pulumi:"recurrenceValue"`
	// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScheduledAction *string `pulumi:"scheduledAction"`
	// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
	ScheduledTaskName *string `pulumi:"scheduledTaskName"`
	// Specifies whether to start the scheduled task. Default to true.
	TaskEnabled *bool `pulumi:"taskEnabled"`
}

type ScheduledTaskState struct {
	// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
	Description pulumi.StringPtrInput
	// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
	DesiredCapacity pulumi.IntPtrInput
	// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
	LaunchExpirationTime pulumi.IntPtrInput
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
	// If the `recurrenceType` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
	// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
	LaunchTime pulumi.StringPtrInput
	// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MaxValue pulumi.IntPtrInput
	// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MinValue pulumi.IntPtrInput
	// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
	RecurrenceEndTime pulumi.StringPtrInput
	// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
	// - Daily: The scheduled task is executed once every specified number of days.
	// - Weekly: The scheduled task is executed on each specified day of a week.
	// - Monthly: The scheduled task is executed on each specified day of a month.
	// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
	RecurrenceType pulumi.StringPtrInput
	// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrenceType`
	// - Daily: You can enter one value. Valid values: 1 to 31.
	// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
	// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
	// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
	RecurrenceValue pulumi.StringPtrInput
	// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScalingGroupId pulumi.StringPtrInput
	// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScheduledAction pulumi.StringPtrInput
	// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
	ScheduledTaskName pulumi.StringPtrInput
	// Specifies whether to start the scheduled task. Default to true.
	TaskEnabled pulumi.BoolPtrInput
}

func (ScheduledTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledTaskState)(nil)).Elem()
}

type scheduledTaskArgs struct {
	// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
	Description *string `pulumi:"description"`
	// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
	LaunchExpirationTime *int `pulumi:"launchExpirationTime"`
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
	// If the `recurrenceType` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
	// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
	LaunchTime *string `pulumi:"launchTime"`
	// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MaxValue *int `pulumi:"maxValue"`
	// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MinValue *int `pulumi:"minValue"`
	// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
	RecurrenceEndTime *string `pulumi:"recurrenceEndTime"`
	// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
	// - Daily: The scheduled task is executed once every specified number of days.
	// - Weekly: The scheduled task is executed on each specified day of a week.
	// - Monthly: The scheduled task is executed on each specified day of a month.
	// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
	RecurrenceType *string `pulumi:"recurrenceType"`
	// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrenceType`
	// - Daily: You can enter one value. Valid values: 1 to 31.
	// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
	// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
	// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
	RecurrenceValue *string `pulumi:"recurrenceValue"`
	// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScheduledAction *string `pulumi:"scheduledAction"`
	// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
	ScheduledTaskName *string `pulumi:"scheduledTaskName"`
	// Specifies whether to start the scheduled task. Default to true.
	TaskEnabled *bool `pulumi:"taskEnabled"`
}

// The set of arguments for constructing a ScheduledTask resource.
type ScheduledTaskArgs struct {
	// Description of the scheduled task, which is 2-200 characters (English or Chinese) long.
	Description pulumi.StringPtrInput
	// The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group. **NOTE:** You must specify the `DesiredCapacity` parameter when you create the scaling group.
	DesiredCapacity pulumi.IntPtrInput
	// The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600
	LaunchExpirationTime pulumi.IntPtrInput
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	// The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation.
	// If the `recurrenceType` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime.
	// Otherwise, the task is only executed once at the date and time specified by LaunchTime.
	LaunchTime pulumi.StringPtrInput
	// The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MaxValue pulumi.IntPtrInput
	// The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.
	MinValue pulumi.IntPtrInput
	// Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	// The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time.
	RecurrenceEndTime pulumi.StringPtrInput
	// Specifies the recurrence type of the scheduled task. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. Valid values:
	// - Daily: The scheduled task is executed once every specified number of days.
	// - Weekly: The scheduled task is executed on each specified day of a week.
	// - Monthly: The scheduled task is executed on each specified day of a month.
	// - Cron: (Available in 1.60.0+) The scheduled task is executed based on the specified cron expression.
	RecurrenceType pulumi.StringPtrInput
	// Specifies how often a scheduled task recurs. **NOTE:** You must specify `RecurrenceType`, `RecurrenceValue`, and `RecurrenceEndTime` at the same time. The valid value depends on `recurrenceType`
	// - Daily: You can enter one value. Valid values: 1 to 31.
	// - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday.
	// - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A.
	// - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), and the L and W letters.
	RecurrenceValue pulumi.StringPtrInput
	// The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the `ScalingGroupId` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScalingGroupId pulumi.StringPtrInput
	// The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule. **NOTE:** You cannot specify `scheduledAction` and `scalingGroupId` at the same time.
	ScheduledAction pulumi.StringPtrInput
	// Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.
	ScheduledTaskName pulumi.StringPtrInput
	// Specifies whether to start the scheduled task. Default to true.
	TaskEnabled pulumi.BoolPtrInput
}

func (ScheduledTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledTaskArgs)(nil)).Elem()
}

type ScheduledTaskInput interface {
	pulumi.Input

	ToScheduledTaskOutput() ScheduledTaskOutput
	ToScheduledTaskOutputWithContext(ctx context.Context) ScheduledTaskOutput
}

func (ScheduledTask) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledTask)(nil)).Elem()
}

func (i ScheduledTask) ToScheduledTaskOutput() ScheduledTaskOutput {
	return i.ToScheduledTaskOutputWithContext(context.Background())
}

func (i ScheduledTask) ToScheduledTaskOutputWithContext(ctx context.Context) ScheduledTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledTaskOutput)
}

type ScheduledTaskOutput struct {
	*pulumi.OutputState
}

func (ScheduledTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledTaskOutput)(nil)).Elem()
}

func (o ScheduledTaskOutput) ToScheduledTaskOutput() ScheduledTaskOutput {
	return o
}

func (o ScheduledTaskOutput) ToScheduledTaskOutputWithContext(ctx context.Context) ScheduledTaskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScheduledTaskOutput{})
}
