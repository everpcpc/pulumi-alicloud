// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chatbot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetAgentsAgent struct {
	// The agent id.
	AgentId string `pulumi:"agentId"`
	// Service space signature, which is used when PAAS interface specifies the service space.
	AgentKey string `pulumi:"agentKey"`
	// The name of the agent.
	AgentName string `pulumi:"agentName"`
	// ID of the agent.
	Id string `pulumi:"id"`
}

// GetAgentsAgentInput is an input type that accepts GetAgentsAgentArgs and GetAgentsAgentOutput values.
// You can construct a concrete instance of `GetAgentsAgentInput` via:
//
//	GetAgentsAgentArgs{...}
type GetAgentsAgentInput interface {
	pulumi.Input

	ToGetAgentsAgentOutput() GetAgentsAgentOutput
	ToGetAgentsAgentOutputWithContext(context.Context) GetAgentsAgentOutput
}

type GetAgentsAgentArgs struct {
	// The agent id.
	AgentId pulumi.StringInput `pulumi:"agentId"`
	// Service space signature, which is used when PAAS interface specifies the service space.
	AgentKey pulumi.StringInput `pulumi:"agentKey"`
	// The name of the agent.
	AgentName pulumi.StringInput `pulumi:"agentName"`
	// ID of the agent.
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetAgentsAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentsAgent)(nil)).Elem()
}

func (i GetAgentsAgentArgs) ToGetAgentsAgentOutput() GetAgentsAgentOutput {
	return i.ToGetAgentsAgentOutputWithContext(context.Background())
}

func (i GetAgentsAgentArgs) ToGetAgentsAgentOutputWithContext(ctx context.Context) GetAgentsAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAgentsAgentOutput)
}

// GetAgentsAgentArrayInput is an input type that accepts GetAgentsAgentArray and GetAgentsAgentArrayOutput values.
// You can construct a concrete instance of `GetAgentsAgentArrayInput` via:
//
//	GetAgentsAgentArray{ GetAgentsAgentArgs{...} }
type GetAgentsAgentArrayInput interface {
	pulumi.Input

	ToGetAgentsAgentArrayOutput() GetAgentsAgentArrayOutput
	ToGetAgentsAgentArrayOutputWithContext(context.Context) GetAgentsAgentArrayOutput
}

type GetAgentsAgentArray []GetAgentsAgentInput

func (GetAgentsAgentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAgentsAgent)(nil)).Elem()
}

func (i GetAgentsAgentArray) ToGetAgentsAgentArrayOutput() GetAgentsAgentArrayOutput {
	return i.ToGetAgentsAgentArrayOutputWithContext(context.Background())
}

func (i GetAgentsAgentArray) ToGetAgentsAgentArrayOutputWithContext(ctx context.Context) GetAgentsAgentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAgentsAgentArrayOutput)
}

type GetAgentsAgentOutput struct{ *pulumi.OutputState }

func (GetAgentsAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentsAgent)(nil)).Elem()
}

func (o GetAgentsAgentOutput) ToGetAgentsAgentOutput() GetAgentsAgentOutput {
	return o
}

func (o GetAgentsAgentOutput) ToGetAgentsAgentOutputWithContext(ctx context.Context) GetAgentsAgentOutput {
	return o
}

// The agent id.
func (o GetAgentsAgentOutput) AgentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsAgent) string { return v.AgentId }).(pulumi.StringOutput)
}

// Service space signature, which is used when PAAS interface specifies the service space.
func (o GetAgentsAgentOutput) AgentKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsAgent) string { return v.AgentKey }).(pulumi.StringOutput)
}

// The name of the agent.
func (o GetAgentsAgentOutput) AgentName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsAgent) string { return v.AgentName }).(pulumi.StringOutput)
}

// ID of the agent.
func (o GetAgentsAgentOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsAgent) string { return v.Id }).(pulumi.StringOutput)
}

type GetAgentsAgentArrayOutput struct{ *pulumi.OutputState }

func (GetAgentsAgentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAgentsAgent)(nil)).Elem()
}

func (o GetAgentsAgentArrayOutput) ToGetAgentsAgentArrayOutput() GetAgentsAgentArrayOutput {
	return o
}

func (o GetAgentsAgentArrayOutput) ToGetAgentsAgentArrayOutputWithContext(ctx context.Context) GetAgentsAgentArrayOutput {
	return o
}

func (o GetAgentsAgentArrayOutput) Index(i pulumi.IntInput) GetAgentsAgentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAgentsAgent {
		return vs[0].([]GetAgentsAgent)[vs[1].(int)]
	}).(GetAgentsAgentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetAgentsAgentInput)(nil)).Elem(), GetAgentsAgentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAgentsAgentArrayInput)(nil)).Elem(), GetAgentsAgentArray{})
	pulumi.RegisterOutputType(GetAgentsAgentOutput{})
	pulumi.RegisterOutputType(GetAgentsAgentArrayOutput{})
}
