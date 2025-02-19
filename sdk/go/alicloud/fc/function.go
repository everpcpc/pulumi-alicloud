// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Alicloud Function Compute Function resource. Function allows you to trigger execution of code in response to events in Alibaba Cloud. The Function itself includes source code and runtime configuration.
//
//	For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/doc-detail/52895.htm).
//
// > **NOTE:** The resource requires a provider field 'account_id'. See account_id.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fc"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "alicloudfcfunctionconfig"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultProject, err := log.NewProject(ctx, "defaultProject", &log.ProjectArgs{
//				Description: pulumi.String("tf unit test"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultStore, err := log.NewStore(ctx, "defaultStore", &log.StoreArgs{
//				Project:         defaultProject.Name,
//				RetentionPeriod: pulumi.Int(3000),
//				ShardCount:      pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRole, err := ram.NewRole(ctx, "defaultRole", &ram.RoleArgs{
//				Document:    pulumi.String("        {\n          \"Statement\": [\n            {\n              \"Action\": \"sts:AssumeRole\",\n              \"Effect\": \"Allow\",\n              \"Principal\": {\n                \"Service\": [\n                  \"fc.aliyuncs.com\"\n                ]\n              }\n            }\n          ],\n          \"Version\": \"1\"\n        }\n    \n"),
//				Description: pulumi.String("this is a test"),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRolePolicyAttachment, err := ram.NewRolePolicyAttachment(ctx, "defaultRolePolicyAttachment", &ram.RolePolicyAttachmentArgs{
//				RoleName:   defaultRole.Name,
//				PolicyName: pulumi.String("AliyunLogFullAccess"),
//				PolicyType: pulumi.String("System"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultService, err := fc.NewService(ctx, "defaultService", &fc.ServiceArgs{
//				Description: pulumi.String("tf unit test"),
//				LogConfig: &fc.ServiceLogConfigArgs{
//					Project:  defaultProject.Name,
//					Logstore: defaultStore.Name,
//				},
//				Role: defaultRole.Arn,
//			}, pulumi.DependsOn([]pulumi.Resource{
//				defaultRolePolicyAttachment,
//			}))
//			if err != nil {
//				return err
//			}
//			defaultBucket, err := oss.NewBucket(ctx, "defaultBucket", &oss.BucketArgs{
//				Bucket: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBucketObject, err := oss.NewBucketObject(ctx, "defaultBucketObject", &oss.BucketObjectArgs{
//				Bucket: defaultBucket.ID(),
//				Key:    pulumi.String("fc/hello.zip"),
//				Source: pulumi.String("./hello.zip"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fc.NewFunction(ctx, "foo", &fc.FunctionArgs{
//				Service:     defaultService.Name,
//				Description: pulumi.String("tf"),
//				OssBucket:   defaultBucket.ID(),
//				OssKey:      defaultBucketObject.Key,
//				MemorySize:  pulumi.Int(512),
//				Runtime:     pulumi.String("python2.7"),
//				Handler:     pulumi.String("hello.handler"),
//				EnvironmentVariables: pulumi.AnyMap{
//					"prefix": pulumi.Any("terraform"),
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
// ## Module Support
//
// You can use to the existing fc module
// to create a function quickly and set several triggers for it.
//
// ## Import
//
// Function Compute function can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:fc/function:Function foo my-fc-service:hello-world
//
// ```
type Function struct {
	pulumi.CustomResourceState

	// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
	CaPort pulumi.IntPtrOutput `pulumi:"caPort"`
	// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
	// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
	CodeChecksum pulumi.StringOutput `pulumi:"codeChecksum"`
	// The configuration for custom container runtime.
	CustomContainerConfig FunctionCustomContainerConfigPtrOutput `pulumi:"customContainerConfig"`
	// The Function Compute function description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A map that defines environment variables for the function.
	EnvironmentVariables pulumi.MapOutput `pulumi:"environmentVariables"`
	// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// The Function Compute service ID.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
	Handler pulumi.StringOutput `pulumi:"handler"`
	// The maximum length of time, in seconds, that the function's initialization should be run for.
	InitializationTimeout pulumi.IntPtrOutput `pulumi:"initializationTimeout"`
	// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
	Initializer pulumi.StringPtrOutput `pulumi:"initializer"`
	// The maximum number of requests can be executed concurrently within the single function instance.
	InstanceConcurrency pulumi.IntPtrOutput `pulumi:"instanceConcurrency"`
	// The instance type of the function.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The date this resource was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The configuration for layers.
	Layers pulumi.StringArrayOutput `pulumi:"layers"`
	// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
	MemorySize pulumi.IntPtrOutput `pulumi:"memorySize"`
	// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringOutput `pulumi:"name"`
	// Setting a prefix to get a only function name. It is conflict with "name".
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
	OssBucket pulumi.StringPtrOutput `pulumi:"ossBucket"`
	// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
	OssKey pulumi.StringPtrOutput `pulumi:"ossKey"`
	// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// The Function Compute service name.
	Service pulumi.StringOutput `pulumi:"service"`
	// The amount of time your function has to run in seconds.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Handler == nil {
		return nil, errors.New("invalid value for required argument 'Handler'")
	}
	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource Function
	err := ctx.RegisterResource("alicloud:fc/function:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("alicloud:fc/function:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
	CaPort *int `pulumi:"caPort"`
	// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
	// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
	CodeChecksum *string `pulumi:"codeChecksum"`
	// The configuration for custom container runtime.
	CustomContainerConfig *FunctionCustomContainerConfig `pulumi:"customContainerConfig"`
	// The Function Compute function description.
	Description *string `pulumi:"description"`
	// A map that defines environment variables for the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
	Filename *string `pulumi:"filename"`
	// The Function Compute service ID.
	FunctionId *string `pulumi:"functionId"`
	// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
	Handler *string `pulumi:"handler"`
	// The maximum length of time, in seconds, that the function's initialization should be run for.
	InitializationTimeout *int `pulumi:"initializationTimeout"`
	// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
	Initializer *string `pulumi:"initializer"`
	// The maximum number of requests can be executed concurrently within the single function instance.
	InstanceConcurrency *int `pulumi:"instanceConcurrency"`
	// The instance type of the function.
	InstanceType *string `pulumi:"instanceType"`
	// The date this resource was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The configuration for layers.
	Layers []string `pulumi:"layers"`
	// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
	MemorySize *int `pulumi:"memorySize"`
	// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only function name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
	OssBucket *string `pulumi:"ossBucket"`
	// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
	OssKey *string `pulumi:"ossKey"`
	// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
	Runtime *string `pulumi:"runtime"`
	// The Function Compute service name.
	Service *string `pulumi:"service"`
	// The amount of time your function has to run in seconds.
	Timeout *int `pulumi:"timeout"`
}

type FunctionState struct {
	// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
	CaPort pulumi.IntPtrInput
	// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
	// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
	CodeChecksum pulumi.StringPtrInput
	// The configuration for custom container runtime.
	CustomContainerConfig FunctionCustomContainerConfigPtrInput
	// The Function Compute function description.
	Description pulumi.StringPtrInput
	// A map that defines environment variables for the function.
	EnvironmentVariables pulumi.MapInput
	// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
	Filename pulumi.StringPtrInput
	// The Function Compute service ID.
	FunctionId pulumi.StringPtrInput
	// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
	Handler pulumi.StringPtrInput
	// The maximum length of time, in seconds, that the function's initialization should be run for.
	InitializationTimeout pulumi.IntPtrInput
	// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
	Initializer pulumi.StringPtrInput
	// The maximum number of requests can be executed concurrently within the single function instance.
	InstanceConcurrency pulumi.IntPtrInput
	// The instance type of the function.
	InstanceType pulumi.StringPtrInput
	// The date this resource was last modified.
	LastModified pulumi.StringPtrInput
	// The configuration for layers.
	Layers pulumi.StringArrayInput
	// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
	MemorySize pulumi.IntPtrInput
	// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only function name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
	OssBucket pulumi.StringPtrInput
	// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
	OssKey pulumi.StringPtrInput
	// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
	Runtime pulumi.StringPtrInput
	// The Function Compute service name.
	Service pulumi.StringPtrInput
	// The amount of time your function has to run in seconds.
	Timeout pulumi.IntPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
	CaPort *int `pulumi:"caPort"`
	// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
	// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
	CodeChecksum *string `pulumi:"codeChecksum"`
	// The configuration for custom container runtime.
	CustomContainerConfig *FunctionCustomContainerConfig `pulumi:"customContainerConfig"`
	// The Function Compute function description.
	Description *string `pulumi:"description"`
	// A map that defines environment variables for the function.
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
	Filename *string `pulumi:"filename"`
	// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
	Handler string `pulumi:"handler"`
	// The maximum length of time, in seconds, that the function's initialization should be run for.
	InitializationTimeout *int `pulumi:"initializationTimeout"`
	// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
	Initializer *string `pulumi:"initializer"`
	// The maximum number of requests can be executed concurrently within the single function instance.
	InstanceConcurrency *int `pulumi:"instanceConcurrency"`
	// The instance type of the function.
	InstanceType *string `pulumi:"instanceType"`
	// The configuration for layers.
	Layers []string `pulumi:"layers"`
	// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
	MemorySize *int `pulumi:"memorySize"`
	// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only function name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
	OssBucket *string `pulumi:"ossBucket"`
	// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
	OssKey *string `pulumi:"ossKey"`
	// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
	Runtime string `pulumi:"runtime"`
	// The Function Compute service name.
	Service string `pulumi:"service"`
	// The amount of time your function has to run in seconds.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
	CaPort pulumi.IntPtrInput
	// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
	// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
	CodeChecksum pulumi.StringPtrInput
	// The configuration for custom container runtime.
	CustomContainerConfig FunctionCustomContainerConfigPtrInput
	// The Function Compute function description.
	Description pulumi.StringPtrInput
	// A map that defines environment variables for the function.
	EnvironmentVariables pulumi.MapInput
	// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
	Filename pulumi.StringPtrInput
	// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
	Handler pulumi.StringInput
	// The maximum length of time, in seconds, that the function's initialization should be run for.
	InitializationTimeout pulumi.IntPtrInput
	// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
	Initializer pulumi.StringPtrInput
	// The maximum number of requests can be executed concurrently within the single function instance.
	InstanceConcurrency pulumi.IntPtrInput
	// The instance type of the function.
	InstanceType pulumi.StringPtrInput
	// The configuration for layers.
	Layers pulumi.StringArrayInput
	// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
	MemorySize pulumi.IntPtrInput
	// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only function name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
	OssBucket pulumi.StringPtrInput
	// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
	OssKey pulumi.StringPtrInput
	// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
	Runtime pulumi.StringInput
	// The Function Compute service name.
	Service pulumi.StringInput
	// The amount of time your function has to run in seconds.
	Timeout pulumi.IntPtrInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}

type FunctionInput interface {
	pulumi.Input

	ToFunctionOutput() FunctionOutput
	ToFunctionOutputWithContext(ctx context.Context) FunctionOutput
}

func (*Function) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (i *Function) ToFunctionOutput() FunctionOutput {
	return i.ToFunctionOutputWithContext(context.Background())
}

func (i *Function) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutput)
}

// FunctionArrayInput is an input type that accepts FunctionArray and FunctionArrayOutput values.
// You can construct a concrete instance of `FunctionArrayInput` via:
//
//	FunctionArray{ FunctionArgs{...} }
type FunctionArrayInput interface {
	pulumi.Input

	ToFunctionArrayOutput() FunctionArrayOutput
	ToFunctionArrayOutputWithContext(context.Context) FunctionArrayOutput
}

type FunctionArray []FunctionInput

func (FunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (i FunctionArray) ToFunctionArrayOutput() FunctionArrayOutput {
	return i.ToFunctionArrayOutputWithContext(context.Background())
}

func (i FunctionArray) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionArrayOutput)
}

// FunctionMapInput is an input type that accepts FunctionMap and FunctionMapOutput values.
// You can construct a concrete instance of `FunctionMapInput` via:
//
//	FunctionMap{ "key": FunctionArgs{...} }
type FunctionMapInput interface {
	pulumi.Input

	ToFunctionMapOutput() FunctionMapOutput
	ToFunctionMapOutputWithContext(context.Context) FunctionMapOutput
}

type FunctionMap map[string]FunctionInput

func (FunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (i FunctionMap) ToFunctionMapOutput() FunctionMapOutput {
	return i.ToFunctionMapOutputWithContext(context.Background())
}

func (i FunctionMap) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionMapOutput)
}

type FunctionOutput struct{ *pulumi.OutputState }

func (FunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Function)(nil)).Elem()
}

func (o FunctionOutput) ToFunctionOutput() FunctionOutput {
	return o
}

func (o FunctionOutput) ToFunctionOutputWithContext(ctx context.Context) FunctionOutput {
	return o
}

// The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
func (o FunctionOutput) CaPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.CaPort }).(pulumi.IntPtrOutput)
}

// The checksum (crc64) of the function code.Used to trigger updates.The value can be generated by data source alicloud_file_crc64_checksum.
// > **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).
func (o FunctionOutput) CodeChecksum() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.CodeChecksum }).(pulumi.StringOutput)
}

// The configuration for custom container runtime.
func (o FunctionOutput) CustomContainerConfig() FunctionCustomContainerConfigPtrOutput {
	return o.ApplyT(func(v *Function) FunctionCustomContainerConfigPtrOutput { return v.CustomContainerConfig }).(FunctionCustomContainerConfigPtrOutput)
}

// The Function Compute function description.
func (o FunctionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A map that defines environment variables for the function.
func (o FunctionOutput) EnvironmentVariables() pulumi.MapOutput {
	return o.ApplyT(func(v *Function) pulumi.MapOutput { return v.EnvironmentVariables }).(pulumi.MapOutput)
}

// The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
func (o FunctionOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Filename }).(pulumi.StringPtrOutput)
}

// The Function Compute service ID.
func (o FunctionOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
func (o FunctionOutput) Handler() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Handler }).(pulumi.StringOutput)
}

// The maximum length of time, in seconds, that the function's initialization should be run for.
func (o FunctionOutput) InitializationTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.InitializationTimeout }).(pulumi.IntPtrOutput)
}

// The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
func (o FunctionOutput) Initializer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.Initializer }).(pulumi.StringPtrOutput)
}

// The maximum number of requests can be executed concurrently within the single function instance.
func (o FunctionOutput) InstanceConcurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.InstanceConcurrency }).(pulumi.IntPtrOutput)
}

// The instance type of the function.
func (o FunctionOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// The date this resource was last modified.
func (o FunctionOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The configuration for layers.
func (o FunctionOutput) Layers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Function) pulumi.StringArrayOutput { return v.Layers }).(pulumi.StringArrayOutput)
}

// Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
func (o FunctionOutput) MemorySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.MemorySize }).(pulumi.IntPtrOutput)
}

// The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
func (o FunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Setting a prefix to get a only function name. It is conflict with "name".
func (o FunctionOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
func (o FunctionOutput) OssBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.OssBucket }).(pulumi.StringPtrOutput)
}

// The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
func (o FunctionOutput) OssKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.StringPtrOutput { return v.OssKey }).(pulumi.StringPtrOutput)
}

// See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
func (o FunctionOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Runtime }).(pulumi.StringOutput)
}

// The Function Compute service name.
func (o FunctionOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *Function) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// The amount of time your function has to run in seconds.
func (o FunctionOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Function) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type FunctionArrayOutput struct{ *pulumi.OutputState }

func (FunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Function)(nil)).Elem()
}

func (o FunctionArrayOutput) ToFunctionArrayOutput() FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) ToFunctionArrayOutputWithContext(ctx context.Context) FunctionArrayOutput {
	return o
}

func (o FunctionArrayOutput) Index(i pulumi.IntInput) FunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Function {
		return vs[0].([]*Function)[vs[1].(int)]
	}).(FunctionOutput)
}

type FunctionMapOutput struct{ *pulumi.OutputState }

func (FunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Function)(nil)).Elem()
}

func (o FunctionMapOutput) ToFunctionMapOutput() FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) ToFunctionMapOutputWithContext(ctx context.Context) FunctionMapOutput {
	return o
}

func (o FunctionMapOutput) MapIndex(k pulumi.StringInput) FunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Function {
		return vs[0].(map[string]*Function)[vs[1].(string)]
	}).(FunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionInput)(nil)).Elem(), &Function{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionArrayInput)(nil)).Elem(), FunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionMapInput)(nil)).Elem(), FunctionMap{})
	pulumi.RegisterOutputType(FunctionOutput{})
	pulumi.RegisterOutputType(FunctionArrayOutput{})
	pulumi.RegisterOutputType(FunctionMapOutput{})
}
