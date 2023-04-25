// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Invocation resource.
//
// For information about ECS Invocation and how to use it, see [What is Invocation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/invokecommand#t9958.html).
//
// > **NOTE:** Available in v1.168.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultCommand, err := ecs.NewCommand(ctx, "defaultCommand", &ecs.CommandArgs{
//				CommandContent: pulumi.String("bHMK"),
//				Description:    pulumi.String("For Terraform Test"),
//				Type:           pulumi.String("RunShellScript"),
//				WorkingDir:     pulumi.String("/root"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstances, err := ecs.GetInstances(ctx, &ecs.GetInstancesArgs{
//				Status: pulumi.StringRef("Running"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsInvocation(ctx, "defaultEcsInvocation", &ecs.EcsInvocationArgs{
//				CommandId: defaultCommand.ID(),
//				InstanceIds: pulumi.StringArray{
//					*pulumi.String(defaultInstances.Ids[0]),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ECS Invocation can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/ecsInvocation:EcsInvocation example <id>
//
// ```
type EcsInvocation struct {
	pulumi.CustomResourceState

	// The ID of the command.
	CommandId pulumi.StringOutput `pulumi:"commandId"`
	// The schedule on which the recurring execution of the command takes place. Take note of the following items:
	// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
	// * When you set Timed to true, you must specify Frequency.
	// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// The list of instances to execute the command. You can specify up to 50 instance IDs.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
	RepeatMode pulumi.StringOutput `pulumi:"repeatMode"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies whether to periodically run the command. Default value: `false`.
	Timed pulumi.BoolOutput `pulumi:"timed"`
	// The username that is used to run the command on the ECS instance.
	// * For Linux instances, the root username is used.
	// * For Windows instances, the System username is used.
	// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
	Username pulumi.StringOutput `pulumi:"username"`
	// The name of the password used to run the command on a Windows instance.
	WindowsPasswordName pulumi.StringPtrOutput `pulumi:"windowsPasswordName"`
}

// NewEcsInvocation registers a new resource with the given unique name, arguments, and options.
func NewEcsInvocation(ctx *pulumi.Context,
	name string, args *EcsInvocationArgs, opts ...pulumi.ResourceOption) (*EcsInvocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommandId == nil {
		return nil, errors.New("invalid value for required argument 'CommandId'")
	}
	if args.InstanceIds == nil {
		return nil, errors.New("invalid value for required argument 'InstanceIds'")
	}
	var resource EcsInvocation
	err := ctx.RegisterResource("alicloud:ecs/ecsInvocation:EcsInvocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsInvocation gets an existing EcsInvocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsInvocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsInvocationState, opts ...pulumi.ResourceOption) (*EcsInvocation, error) {
	var resource EcsInvocation
	err := ctx.ReadResource("alicloud:ecs/ecsInvocation:EcsInvocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsInvocation resources.
type ecsInvocationState struct {
	// The ID of the command.
	CommandId *string `pulumi:"commandId"`
	// The schedule on which the recurring execution of the command takes place. Take note of the following items:
	// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
	// * When you set Timed to true, you must specify Frequency.
	// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
	Frequency *string `pulumi:"frequency"`
	// The list of instances to execute the command. You can specify up to 50 instance IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
	RepeatMode *string `pulumi:"repeatMode"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Specifies whether to periodically run the command. Default value: `false`.
	Timed *bool `pulumi:"timed"`
	// The username that is used to run the command on the ECS instance.
	// * For Linux instances, the root username is used.
	// * For Windows instances, the System username is used.
	// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
	Username *string `pulumi:"username"`
	// The name of the password used to run the command on a Windows instance.
	WindowsPasswordName *string `pulumi:"windowsPasswordName"`
}

type EcsInvocationState struct {
	// The ID of the command.
	CommandId pulumi.StringPtrInput
	// The schedule on which the recurring execution of the command takes place. Take note of the following items:
	// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
	// * When you set Timed to true, you must specify Frequency.
	// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
	Frequency pulumi.StringPtrInput
	// The list of instances to execute the command. You can specify up to 50 instance IDs.
	InstanceIds pulumi.StringArrayInput
	// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
	Parameters pulumi.MapInput
	// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
	RepeatMode pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Specifies whether to periodically run the command. Default value: `false`.
	Timed pulumi.BoolPtrInput
	// The username that is used to run the command on the ECS instance.
	// * For Linux instances, the root username is used.
	// * For Windows instances, the System username is used.
	// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
	Username pulumi.StringPtrInput
	// The name of the password used to run the command on a Windows instance.
	WindowsPasswordName pulumi.StringPtrInput
}

func (EcsInvocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsInvocationState)(nil)).Elem()
}

type ecsInvocationArgs struct {
	// The ID of the command.
	CommandId string `pulumi:"commandId"`
	// The schedule on which the recurring execution of the command takes place. Take note of the following items:
	// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
	// * When you set Timed to true, you must specify Frequency.
	// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
	Frequency *string `pulumi:"frequency"`
	// The list of instances to execute the command. You can specify up to 50 instance IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
	RepeatMode *string `pulumi:"repeatMode"`
	// Specifies whether to periodically run the command. Default value: `false`.
	Timed *bool `pulumi:"timed"`
	// The username that is used to run the command on the ECS instance.
	// * For Linux instances, the root username is used.
	// * For Windows instances, the System username is used.
	// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
	Username *string `pulumi:"username"`
	// The name of the password used to run the command on a Windows instance.
	WindowsPasswordName *string `pulumi:"windowsPasswordName"`
}

// The set of arguments for constructing a EcsInvocation resource.
type EcsInvocationArgs struct {
	// The ID of the command.
	CommandId pulumi.StringInput
	// The schedule on which the recurring execution of the command takes place. Take note of the following items:
	// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
	// * When you set Timed to true, you must specify Frequency.
	// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
	Frequency pulumi.StringPtrInput
	// The list of instances to execute the command. You can specify up to 50 instance IDs.
	InstanceIds pulumi.StringArrayInput
	// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
	Parameters pulumi.MapInput
	// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
	RepeatMode pulumi.StringPtrInput
	// Specifies whether to periodically run the command. Default value: `false`.
	Timed pulumi.BoolPtrInput
	// The username that is used to run the command on the ECS instance.
	// * For Linux instances, the root username is used.
	// * For Windows instances, the System username is used.
	// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
	Username pulumi.StringPtrInput
	// The name of the password used to run the command on a Windows instance.
	WindowsPasswordName pulumi.StringPtrInput
}

func (EcsInvocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsInvocationArgs)(nil)).Elem()
}

type EcsInvocationInput interface {
	pulumi.Input

	ToEcsInvocationOutput() EcsInvocationOutput
	ToEcsInvocationOutputWithContext(ctx context.Context) EcsInvocationOutput
}

func (*EcsInvocation) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsInvocation)(nil)).Elem()
}

func (i *EcsInvocation) ToEcsInvocationOutput() EcsInvocationOutput {
	return i.ToEcsInvocationOutputWithContext(context.Background())
}

func (i *EcsInvocation) ToEcsInvocationOutputWithContext(ctx context.Context) EcsInvocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsInvocationOutput)
}

// EcsInvocationArrayInput is an input type that accepts EcsInvocationArray and EcsInvocationArrayOutput values.
// You can construct a concrete instance of `EcsInvocationArrayInput` via:
//
//	EcsInvocationArray{ EcsInvocationArgs{...} }
type EcsInvocationArrayInput interface {
	pulumi.Input

	ToEcsInvocationArrayOutput() EcsInvocationArrayOutput
	ToEcsInvocationArrayOutputWithContext(context.Context) EcsInvocationArrayOutput
}

type EcsInvocationArray []EcsInvocationInput

func (EcsInvocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsInvocation)(nil)).Elem()
}

func (i EcsInvocationArray) ToEcsInvocationArrayOutput() EcsInvocationArrayOutput {
	return i.ToEcsInvocationArrayOutputWithContext(context.Background())
}

func (i EcsInvocationArray) ToEcsInvocationArrayOutputWithContext(ctx context.Context) EcsInvocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsInvocationArrayOutput)
}

// EcsInvocationMapInput is an input type that accepts EcsInvocationMap and EcsInvocationMapOutput values.
// You can construct a concrete instance of `EcsInvocationMapInput` via:
//
//	EcsInvocationMap{ "key": EcsInvocationArgs{...} }
type EcsInvocationMapInput interface {
	pulumi.Input

	ToEcsInvocationMapOutput() EcsInvocationMapOutput
	ToEcsInvocationMapOutputWithContext(context.Context) EcsInvocationMapOutput
}

type EcsInvocationMap map[string]EcsInvocationInput

func (EcsInvocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsInvocation)(nil)).Elem()
}

func (i EcsInvocationMap) ToEcsInvocationMapOutput() EcsInvocationMapOutput {
	return i.ToEcsInvocationMapOutputWithContext(context.Background())
}

func (i EcsInvocationMap) ToEcsInvocationMapOutputWithContext(ctx context.Context) EcsInvocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsInvocationMapOutput)
}

type EcsInvocationOutput struct{ *pulumi.OutputState }

func (EcsInvocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsInvocation)(nil)).Elem()
}

func (o EcsInvocationOutput) ToEcsInvocationOutput() EcsInvocationOutput {
	return o
}

func (o EcsInvocationOutput) ToEcsInvocationOutputWithContext(ctx context.Context) EcsInvocationOutput {
	return o
}

// The ID of the command.
func (o EcsInvocationOutput) CommandId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringOutput { return v.CommandId }).(pulumi.StringOutput)
}

// The schedule on which the recurring execution of the command takes place. Take note of the following items:
// * The interval between two consecutive executions must be 10 seconds or longer. The minimum interval cannot be less than the timeout period of the execution.
// * When you set Timed to true, you must specify Frequency.
// * The value of the Frequency parameter is a cron expression. For more information, see [Cron expression](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/cron-expression).
func (o EcsInvocationOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringPtrOutput { return v.Frequency }).(pulumi.StringPtrOutput)
}

// The list of instances to execute the command. You can specify up to 50 instance IDs.
func (o EcsInvocationOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringArrayOutput { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// The key-value pairs of custom parameters to be passed in when the custom parameter feature is enabled.  Number of custom parameters: 0 to 10.
func (o EcsInvocationOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.MapOutput { return v.Parameters }).(pulumi.MapOutput)
}

// Specifies how to run the command. Valid values: `Once`, `Period`, `NextRebootOnly`, `EveryReboot`. Default value: When `timed` is set to false and Frequency is not specified, the default value of `repeatMode` is `Once`. When `Timed` is set to true and Frequency is specified, `period` is used as the value of RepeatMode regardless of whether `repeatMode` is specified.
func (o EcsInvocationOutput) RepeatMode() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringOutput { return v.RepeatMode }).(pulumi.StringOutput)
}

// The status of the resource.
func (o EcsInvocationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies whether to periodically run the command. Default value: `false`.
func (o EcsInvocationOutput) Timed() pulumi.BoolOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.BoolOutput { return v.Timed }).(pulumi.BoolOutput)
}

// The username that is used to run the command on the ECS instance.
// * For Linux instances, the root username is used.
// * For Windows instances, the System username is used.
// * You can also specify other usernames that already exist in the ECS instance to run the command. It is more secure to run Cloud Assistant commands as a regular user. For more information, see [Configure a regular user to run Cloud Assistant commands](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/run-cloud-assistant-commands-as-a-regular-user).
func (o EcsInvocationOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// The name of the password used to run the command on a Windows instance.
func (o EcsInvocationOutput) WindowsPasswordName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsInvocation) pulumi.StringPtrOutput { return v.WindowsPasswordName }).(pulumi.StringPtrOutput)
}

type EcsInvocationArrayOutput struct{ *pulumi.OutputState }

func (EcsInvocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsInvocation)(nil)).Elem()
}

func (o EcsInvocationArrayOutput) ToEcsInvocationArrayOutput() EcsInvocationArrayOutput {
	return o
}

func (o EcsInvocationArrayOutput) ToEcsInvocationArrayOutputWithContext(ctx context.Context) EcsInvocationArrayOutput {
	return o
}

func (o EcsInvocationArrayOutput) Index(i pulumi.IntInput) EcsInvocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsInvocation {
		return vs[0].([]*EcsInvocation)[vs[1].(int)]
	}).(EcsInvocationOutput)
}

type EcsInvocationMapOutput struct{ *pulumi.OutputState }

func (EcsInvocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsInvocation)(nil)).Elem()
}

func (o EcsInvocationMapOutput) ToEcsInvocationMapOutput() EcsInvocationMapOutput {
	return o
}

func (o EcsInvocationMapOutput) ToEcsInvocationMapOutputWithContext(ctx context.Context) EcsInvocationMapOutput {
	return o
}

func (o EcsInvocationMapOutput) MapIndex(k pulumi.StringInput) EcsInvocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsInvocation {
		return vs[0].(map[string]*EcsInvocation)[vs[1].(string)]
	}).(EcsInvocationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsInvocationInput)(nil)).Elem(), &EcsInvocation{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsInvocationArrayInput)(nil)).Elem(), EcsInvocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsInvocationMapInput)(nil)).Elem(), EcsInvocationMap{})
	pulumi.RegisterOutputType(EcsInvocationOutput{})
	pulumi.RegisterOutputType(EcsInvocationArrayOutput{})
	pulumi.RegisterOutputType(EcsInvocationMapOutput{})
}
