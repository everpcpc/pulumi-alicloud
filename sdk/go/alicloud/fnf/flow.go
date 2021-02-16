// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fnf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Serverless Workflow Flow resource.
//
// For information about Serverless Workflow Flow and how to use it, see [What is Flow](https://www.alibabacloud.com/help/en/doc-detail/123079.htm).
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
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/fnf"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ram"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ram.NewRole(ctx, "_default", &ram.RoleArgs{
// 			Document: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "    \"Statement\": [\n", "      {\n", "        \"Action\": \"sts:AssumeRole\",\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\n", "          \"Service\": [\n", "            \"fnf.aliyuncs.com\"\n", "          ]\n", "        }\n", "      }\n", "    ],\n", "    \"Version\": \"1\"\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fnf.NewFlow(ctx, "example", &fnf.FlowArgs{
// 			Definition:  pulumi.String(fmt.Sprintf("%v%v%v%v%v", "  version: v1beta1\n", "  type: flow\n", "  steps:\n", "    - type: pass\n", "      name: helloworld\n")),
// 			RoleArn:     _default.Arn,
// 			Description: pulumi.String("Test for terraform fnf_flow."),
// 			Type:        pulumi.String("FDL"),
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
// Serverless Workflow Flow can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:fnf/flow:Flow example <name>
// ```
type Flow struct {
	pulumi.CustomResourceState

	// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
	Definition pulumi.StringOutput `pulumi:"definition"`
	// The description of the flow.
	Description pulumi.StringOutput `pulumi:"description"`
	// The unique ID of the flow.
	FlowId pulumi.StringOutput `pulumi:"flowId"`
	// The time when the flow was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the flow. The name must be unique in an Alibaba Cloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The type of the flow. Valid values are `FDL` or `DEFAULT`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Flow
	err := ctx.RegisterResource("alicloud:fnf/flow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("alicloud:fnf/flow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
	// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
	Definition *string `pulumi:"definition"`
	// The description of the flow.
	Description *string `pulumi:"description"`
	// The unique ID of the flow.
	FlowId *string `pulumi:"flowId"`
	// The time when the flow was last modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the flow. The name must be unique in an Alibaba Cloud account.
	Name *string `pulumi:"name"`
	// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
	RoleArn *string `pulumi:"roleArn"`
	// The type of the flow. Valid values are `FDL` or `DEFAULT`.
	Type *string `pulumi:"type"`
}

type FlowState struct {
	// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
	Definition pulumi.StringPtrInput
	// The description of the flow.
	Description pulumi.StringPtrInput
	// The unique ID of the flow.
	FlowId pulumi.StringPtrInput
	// The time when the flow was last modified.
	LastModifiedTime pulumi.StringPtrInput
	// The name of the flow. The name must be unique in an Alibaba Cloud account.
	Name pulumi.StringPtrInput
	// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
	RoleArn pulumi.StringPtrInput
	// The type of the flow. Valid values are `FDL` or `DEFAULT`.
	Type pulumi.StringPtrInput
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
	Definition string `pulumi:"definition"`
	// The description of the flow.
	Description string `pulumi:"description"`
	// The name of the flow. The name must be unique in an Alibaba Cloud account.
	Name *string `pulumi:"name"`
	// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
	RoleArn *string `pulumi:"roleArn"`
	// The type of the flow. Valid values are `FDL` or `DEFAULT`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
	Definition pulumi.StringInput
	// The description of the flow.
	Description pulumi.StringInput
	// The name of the flow. The name must be unique in an Alibaba Cloud account.
	Name pulumi.StringPtrInput
	// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
	RoleArn pulumi.StringPtrInput
	// The type of the flow. Valid values are `FDL` or `DEFAULT`.
	Type pulumi.StringInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}

type FlowInput interface {
	pulumi.Input

	ToFlowOutput() FlowOutput
	ToFlowOutputWithContext(ctx context.Context) FlowOutput
}

func (*Flow) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (i *Flow) ToFlowOutput() FlowOutput {
	return i.ToFlowOutputWithContext(context.Background())
}

func (i *Flow) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowOutput)
}

func (i *Flow) ToFlowPtrOutput() FlowPtrOutput {
	return i.ToFlowPtrOutputWithContext(context.Background())
}

func (i *Flow) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowPtrOutput)
}

type FlowPtrInput interface {
	pulumi.Input

	ToFlowPtrOutput() FlowPtrOutput
	ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput
}

type flowPtrType FlowArgs

func (*flowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil))
}

func (i *flowPtrType) ToFlowPtrOutput() FlowPtrOutput {
	return i.ToFlowPtrOutputWithContext(context.Background())
}

func (i *flowPtrType) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowPtrOutput)
}

// FlowArrayInput is an input type that accepts FlowArray and FlowArrayOutput values.
// You can construct a concrete instance of `FlowArrayInput` via:
//
//          FlowArray{ FlowArgs{...} }
type FlowArrayInput interface {
	pulumi.Input

	ToFlowArrayOutput() FlowArrayOutput
	ToFlowArrayOutputWithContext(context.Context) FlowArrayOutput
}

type FlowArray []FlowInput

func (FlowArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Flow)(nil))
}

func (i FlowArray) ToFlowArrayOutput() FlowArrayOutput {
	return i.ToFlowArrayOutputWithContext(context.Background())
}

func (i FlowArray) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowArrayOutput)
}

// FlowMapInput is an input type that accepts FlowMap and FlowMapOutput values.
// You can construct a concrete instance of `FlowMapInput` via:
//
//          FlowMap{ "key": FlowArgs{...} }
type FlowMapInput interface {
	pulumi.Input

	ToFlowMapOutput() FlowMapOutput
	ToFlowMapOutputWithContext(context.Context) FlowMapOutput
}

type FlowMap map[string]FlowInput

func (FlowMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Flow)(nil))
}

func (i FlowMap) ToFlowMapOutput() FlowMapOutput {
	return i.ToFlowMapOutputWithContext(context.Background())
}

func (i FlowMap) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowMapOutput)
}

type FlowOutput struct {
	*pulumi.OutputState
}

func (FlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Flow)(nil))
}

func (o FlowOutput) ToFlowOutput() FlowOutput {
	return o
}

func (o FlowOutput) ToFlowOutputWithContext(ctx context.Context) FlowOutput {
	return o
}

func (o FlowOutput) ToFlowPtrOutput() FlowPtrOutput {
	return o.ToFlowPtrOutputWithContext(context.Background())
}

func (o FlowOutput) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return o.ApplyT(func(v Flow) *Flow {
		return &v
	}).(FlowPtrOutput)
}

type FlowPtrOutput struct {
	*pulumi.OutputState
}

func (FlowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flow)(nil))
}

func (o FlowPtrOutput) ToFlowPtrOutput() FlowPtrOutput {
	return o
}

func (o FlowPtrOutput) ToFlowPtrOutputWithContext(ctx context.Context) FlowPtrOutput {
	return o
}

type FlowArrayOutput struct{ *pulumi.OutputState }

func (FlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Flow)(nil))
}

func (o FlowArrayOutput) ToFlowArrayOutput() FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) ToFlowArrayOutputWithContext(ctx context.Context) FlowArrayOutput {
	return o
}

func (o FlowArrayOutput) Index(i pulumi.IntInput) FlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Flow {
		return vs[0].([]Flow)[vs[1].(int)]
	}).(FlowOutput)
}

type FlowMapOutput struct{ *pulumi.OutputState }

func (FlowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Flow)(nil))
}

func (o FlowMapOutput) ToFlowMapOutput() FlowMapOutput {
	return o
}

func (o FlowMapOutput) ToFlowMapOutputWithContext(ctx context.Context) FlowMapOutput {
	return o
}

func (o FlowMapOutput) MapIndex(k pulumi.StringInput) FlowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Flow {
		return vs[0].(map[string]Flow)[vs[1].(string)]
	}).(FlowOutput)
}

func init() {
	pulumi.RegisterOutputType(FlowOutput{})
	pulumi.RegisterOutputType(FlowPtrOutput{})
	pulumi.RegisterOutputType(FlowArrayOutput{})
	pulumi.RegisterOutputType(FlowMapOutput{})
}
