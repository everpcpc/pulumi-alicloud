// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Logtial config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:log/logTailConfig:LogTailConfig example tf-log:tf-log-store:tf-log-config
// ```
type LogTailConfig struct {
	pulumi.CustomResourceState

	// The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
	InputDetail pulumi.StringOutput `pulumi:"inputDetail"`
	// The input type. Currently only two types of files and plugin are supported.
	InputType pulumi.StringOutput `pulumi:"inputType"`
	// （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
	LogSample pulumi.StringPtrOutput `pulumi:"logSample"`
	// The log store name to the query index belongs.
	Logstore pulumi.StringOutput `pulumi:"logstore"`
	// The Logtail configuration name, which is unique in the same project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The output type. Currently, only LogService is supported.
	OutputType pulumi.StringOutput `pulumi:"outputType"`
	// The project name to the log store belongs.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewLogTailConfig registers a new resource with the given unique name, arguments, and options.
func NewLogTailConfig(ctx *pulumi.Context,
	name string, args *LogTailConfigArgs, opts ...pulumi.ResourceOption) (*LogTailConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputDetail == nil {
		return nil, errors.New("invalid value for required argument 'InputDetail'")
	}
	if args.InputType == nil {
		return nil, errors.New("invalid value for required argument 'InputType'")
	}
	if args.Logstore == nil {
		return nil, errors.New("invalid value for required argument 'Logstore'")
	}
	if args.OutputType == nil {
		return nil, errors.New("invalid value for required argument 'OutputType'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource LogTailConfig
	err := ctx.RegisterResource("alicloud:log/logTailConfig:LogTailConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogTailConfig gets an existing LogTailConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogTailConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogTailConfigState, opts ...pulumi.ResourceOption) (*LogTailConfig, error) {
	var resource LogTailConfig
	err := ctx.ReadResource("alicloud:log/logTailConfig:LogTailConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogTailConfig resources.
type logTailConfigState struct {
	// The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
	InputDetail *string `pulumi:"inputDetail"`
	// The input type. Currently only two types of files and plugin are supported.
	InputType *string `pulumi:"inputType"`
	// （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
	LogSample *string `pulumi:"logSample"`
	// The log store name to the query index belongs.
	Logstore *string `pulumi:"logstore"`
	// The Logtail configuration name, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The output type. Currently, only LogService is supported.
	OutputType *string `pulumi:"outputType"`
	// The project name to the log store belongs.
	Project *string `pulumi:"project"`
}

type LogTailConfigState struct {
	// The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
	InputDetail pulumi.StringPtrInput
	// The input type. Currently only two types of files and plugin are supported.
	InputType pulumi.StringPtrInput
	// （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
	LogSample pulumi.StringPtrInput
	// The log store name to the query index belongs.
	Logstore pulumi.StringPtrInput
	// The Logtail configuration name, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The output type. Currently, only LogService is supported.
	OutputType pulumi.StringPtrInput
	// The project name to the log store belongs.
	Project pulumi.StringPtrInput
}

func (LogTailConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailConfigState)(nil)).Elem()
}

type logTailConfigArgs struct {
	// The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
	InputDetail string `pulumi:"inputDetail"`
	// The input type. Currently only two types of files and plugin are supported.
	InputType string `pulumi:"inputType"`
	// （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
	LogSample *string `pulumi:"logSample"`
	// The log store name to the query index belongs.
	Logstore string `pulumi:"logstore"`
	// The Logtail configuration name, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The output type. Currently, only LogService is supported.
	OutputType string `pulumi:"outputType"`
	// The project name to the log store belongs.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a LogTailConfig resource.
type LogTailConfigArgs struct {
	// The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm))
	InputDetail pulumi.StringInput
	// The input type. Currently only two types of files and plugin are supported.
	InputType pulumi.StringInput
	// （Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.
	LogSample pulumi.StringPtrInput
	// The log store name to the query index belongs.
	Logstore pulumi.StringInput
	// The Logtail configuration name, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The output type. Currently, only LogService is supported.
	OutputType pulumi.StringInput
	// The project name to the log store belongs.
	Project pulumi.StringInput
}

func (LogTailConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailConfigArgs)(nil)).Elem()
}

type LogTailConfigInput interface {
	pulumi.Input

	ToLogTailConfigOutput() LogTailConfigOutput
	ToLogTailConfigOutputWithContext(ctx context.Context) LogTailConfigOutput
}

func (*LogTailConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*LogTailConfig)(nil))
}

func (i *LogTailConfig) ToLogTailConfigOutput() LogTailConfigOutput {
	return i.ToLogTailConfigOutputWithContext(context.Background())
}

func (i *LogTailConfig) ToLogTailConfigOutputWithContext(ctx context.Context) LogTailConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailConfigOutput)
}

func (i *LogTailConfig) ToLogTailConfigPtrOutput() LogTailConfigPtrOutput {
	return i.ToLogTailConfigPtrOutputWithContext(context.Background())
}

func (i *LogTailConfig) ToLogTailConfigPtrOutputWithContext(ctx context.Context) LogTailConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailConfigPtrOutput)
}

type LogTailConfigPtrInput interface {
	pulumi.Input

	ToLogTailConfigPtrOutput() LogTailConfigPtrOutput
	ToLogTailConfigPtrOutputWithContext(ctx context.Context) LogTailConfigPtrOutput
}

type logTailConfigPtrType LogTailConfigArgs

func (*logTailConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTailConfig)(nil))
}

func (i *logTailConfigPtrType) ToLogTailConfigPtrOutput() LogTailConfigPtrOutput {
	return i.ToLogTailConfigPtrOutputWithContext(context.Background())
}

func (i *logTailConfigPtrType) ToLogTailConfigPtrOutputWithContext(ctx context.Context) LogTailConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailConfigPtrOutput)
}

// LogTailConfigArrayInput is an input type that accepts LogTailConfigArray and LogTailConfigArrayOutput values.
// You can construct a concrete instance of `LogTailConfigArrayInput` via:
//
//          LogTailConfigArray{ LogTailConfigArgs{...} }
type LogTailConfigArrayInput interface {
	pulumi.Input

	ToLogTailConfigArrayOutput() LogTailConfigArrayOutput
	ToLogTailConfigArrayOutputWithContext(context.Context) LogTailConfigArrayOutput
}

type LogTailConfigArray []LogTailConfigInput

func (LogTailConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogTailConfig)(nil)).Elem()
}

func (i LogTailConfigArray) ToLogTailConfigArrayOutput() LogTailConfigArrayOutput {
	return i.ToLogTailConfigArrayOutputWithContext(context.Background())
}

func (i LogTailConfigArray) ToLogTailConfigArrayOutputWithContext(ctx context.Context) LogTailConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailConfigArrayOutput)
}

// LogTailConfigMapInput is an input type that accepts LogTailConfigMap and LogTailConfigMapOutput values.
// You can construct a concrete instance of `LogTailConfigMapInput` via:
//
//          LogTailConfigMap{ "key": LogTailConfigArgs{...} }
type LogTailConfigMapInput interface {
	pulumi.Input

	ToLogTailConfigMapOutput() LogTailConfigMapOutput
	ToLogTailConfigMapOutputWithContext(context.Context) LogTailConfigMapOutput
}

type LogTailConfigMap map[string]LogTailConfigInput

func (LogTailConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogTailConfig)(nil)).Elem()
}

func (i LogTailConfigMap) ToLogTailConfigMapOutput() LogTailConfigMapOutput {
	return i.ToLogTailConfigMapOutputWithContext(context.Background())
}

func (i LogTailConfigMap) ToLogTailConfigMapOutputWithContext(ctx context.Context) LogTailConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailConfigMapOutput)
}

type LogTailConfigOutput struct{ *pulumi.OutputState }

func (LogTailConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogTailConfig)(nil))
}

func (o LogTailConfigOutput) ToLogTailConfigOutput() LogTailConfigOutput {
	return o
}

func (o LogTailConfigOutput) ToLogTailConfigOutputWithContext(ctx context.Context) LogTailConfigOutput {
	return o
}

func (o LogTailConfigOutput) ToLogTailConfigPtrOutput() LogTailConfigPtrOutput {
	return o.ToLogTailConfigPtrOutputWithContext(context.Background())
}

func (o LogTailConfigOutput) ToLogTailConfigPtrOutputWithContext(ctx context.Context) LogTailConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogTailConfig) *LogTailConfig {
		return &v
	}).(LogTailConfigPtrOutput)
}

type LogTailConfigPtrOutput struct{ *pulumi.OutputState }

func (LogTailConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogTailConfig)(nil))
}

func (o LogTailConfigPtrOutput) ToLogTailConfigPtrOutput() LogTailConfigPtrOutput {
	return o
}

func (o LogTailConfigPtrOutput) ToLogTailConfigPtrOutputWithContext(ctx context.Context) LogTailConfigPtrOutput {
	return o
}

func (o LogTailConfigPtrOutput) Elem() LogTailConfigOutput {
	return o.ApplyT(func(v *LogTailConfig) LogTailConfig {
		if v != nil {
			return *v
		}
		var ret LogTailConfig
		return ret
	}).(LogTailConfigOutput)
}

type LogTailConfigArrayOutput struct{ *pulumi.OutputState }

func (LogTailConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogTailConfig)(nil))
}

func (o LogTailConfigArrayOutput) ToLogTailConfigArrayOutput() LogTailConfigArrayOutput {
	return o
}

func (o LogTailConfigArrayOutput) ToLogTailConfigArrayOutputWithContext(ctx context.Context) LogTailConfigArrayOutput {
	return o
}

func (o LogTailConfigArrayOutput) Index(i pulumi.IntInput) LogTailConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogTailConfig {
		return vs[0].([]LogTailConfig)[vs[1].(int)]
	}).(LogTailConfigOutput)
}

type LogTailConfigMapOutput struct{ *pulumi.OutputState }

func (LogTailConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogTailConfig)(nil))
}

func (o LogTailConfigMapOutput) ToLogTailConfigMapOutput() LogTailConfigMapOutput {
	return o
}

func (o LogTailConfigMapOutput) ToLogTailConfigMapOutputWithContext(ctx context.Context) LogTailConfigMapOutput {
	return o
}

func (o LogTailConfigMapOutput) MapIndex(k pulumi.StringInput) LogTailConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogTailConfig {
		return vs[0].(map[string]LogTailConfig)[vs[1].(string)]
	}).(LogTailConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogTailConfigInput)(nil)).Elem(), &LogTailConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTailConfigPtrInput)(nil)).Elem(), &LogTailConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTailConfigArrayInput)(nil)).Elem(), LogTailConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogTailConfigMapInput)(nil)).Elem(), LogTailConfigMap{})
	pulumi.RegisterOutputType(LogTailConfigOutput{})
	pulumi.RegisterOutputType(LogTailConfigPtrOutput{})
	pulumi.RegisterOutputType(LogTailConfigArrayOutput{})
	pulumi.RegisterOutputType(LogTailConfigMapOutput{})
}
