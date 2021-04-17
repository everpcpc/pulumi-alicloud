// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SchedulerRuleRule struct {
	Priority  *int    `pulumi:"priority"`
	RegionId  *string `pulumi:"regionId"`
	Status    *int    `pulumi:"status"`
	Type      *string `pulumi:"type"`
	Value     *string `pulumi:"value"`
	ValueType *int    `pulumi:"valueType"`
}

// SchedulerRuleRuleInput is an input type that accepts SchedulerRuleRuleArgs and SchedulerRuleRuleOutput values.
// You can construct a concrete instance of `SchedulerRuleRuleInput` via:
//
//          SchedulerRuleRuleArgs{...}
type SchedulerRuleRuleInput interface {
	pulumi.Input

	ToSchedulerRuleRuleOutput() SchedulerRuleRuleOutput
	ToSchedulerRuleRuleOutputWithContext(context.Context) SchedulerRuleRuleOutput
}

type SchedulerRuleRuleArgs struct {
	Priority  pulumi.IntPtrInput    `pulumi:"priority"`
	RegionId  pulumi.StringPtrInput `pulumi:"regionId"`
	Status    pulumi.IntPtrInput    `pulumi:"status"`
	Type      pulumi.StringPtrInput `pulumi:"type"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
	ValueType pulumi.IntPtrInput    `pulumi:"valueType"`
}

func (SchedulerRuleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulerRuleRule)(nil)).Elem()
}

func (i SchedulerRuleRuleArgs) ToSchedulerRuleRuleOutput() SchedulerRuleRuleOutput {
	return i.ToSchedulerRuleRuleOutputWithContext(context.Background())
}

func (i SchedulerRuleRuleArgs) ToSchedulerRuleRuleOutputWithContext(ctx context.Context) SchedulerRuleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulerRuleRuleOutput)
}

// SchedulerRuleRuleArrayInput is an input type that accepts SchedulerRuleRuleArray and SchedulerRuleRuleArrayOutput values.
// You can construct a concrete instance of `SchedulerRuleRuleArrayInput` via:
//
//          SchedulerRuleRuleArray{ SchedulerRuleRuleArgs{...} }
type SchedulerRuleRuleArrayInput interface {
	pulumi.Input

	ToSchedulerRuleRuleArrayOutput() SchedulerRuleRuleArrayOutput
	ToSchedulerRuleRuleArrayOutputWithContext(context.Context) SchedulerRuleRuleArrayOutput
}

type SchedulerRuleRuleArray []SchedulerRuleRuleInput

func (SchedulerRuleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchedulerRuleRule)(nil)).Elem()
}

func (i SchedulerRuleRuleArray) ToSchedulerRuleRuleArrayOutput() SchedulerRuleRuleArrayOutput {
	return i.ToSchedulerRuleRuleArrayOutputWithContext(context.Background())
}

func (i SchedulerRuleRuleArray) ToSchedulerRuleRuleArrayOutputWithContext(ctx context.Context) SchedulerRuleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulerRuleRuleArrayOutput)
}

type SchedulerRuleRuleOutput struct{ *pulumi.OutputState }

func (SchedulerRuleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulerRuleRule)(nil)).Elem()
}

func (o SchedulerRuleRuleOutput) ToSchedulerRuleRuleOutput() SchedulerRuleRuleOutput {
	return o
}

func (o SchedulerRuleRuleOutput) ToSchedulerRuleRuleOutputWithContext(ctx context.Context) SchedulerRuleRuleOutput {
	return o
}

func (o SchedulerRuleRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SchedulerRuleRuleOutput) RegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *string { return v.RegionId }).(pulumi.StringPtrOutput)
}

func (o SchedulerRuleRuleOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o SchedulerRuleRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SchedulerRuleRuleOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o SchedulerRuleRuleOutput) ValueType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulerRuleRule) *int { return v.ValueType }).(pulumi.IntPtrOutput)
}

type SchedulerRuleRuleArrayOutput struct{ *pulumi.OutputState }

func (SchedulerRuleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchedulerRuleRule)(nil)).Elem()
}

func (o SchedulerRuleRuleArrayOutput) ToSchedulerRuleRuleArrayOutput() SchedulerRuleRuleArrayOutput {
	return o
}

func (o SchedulerRuleRuleArrayOutput) ToSchedulerRuleRuleArrayOutputWithContext(ctx context.Context) SchedulerRuleRuleArrayOutput {
	return o
}

func (o SchedulerRuleRuleArrayOutput) Index(i pulumi.IntInput) SchedulerRuleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SchedulerRuleRule {
		return vs[0].([]SchedulerRuleRule)[vs[1].(int)]
	}).(SchedulerRuleRuleOutput)
}

type GetDdosBgpInstancesInstance struct {
	// The instance's elastic defend bandwidth.
	Bandwidth int `pulumi:"bandwidth"`
	// The instance's base defend bandwidth.
	BaseBandwidth int `pulumi:"baseBandwidth"`
	// The instance's id.
	Id string `pulumi:"id"`
	// The instance's count of ip config.
	IpCount int `pulumi:"ipCount"`
	// The instance's IP version.
	IpType string `pulumi:"ipType"`
	// The instance's remark.
	Name string `pulumi:"name"`
	// A region of instance.
	Region string `pulumi:"region"`
	// The instance's type.
	Type string `pulumi:"type"`
}

// GetDdosBgpInstancesInstanceInput is an input type that accepts GetDdosBgpInstancesInstanceArgs and GetDdosBgpInstancesInstanceOutput values.
// You can construct a concrete instance of `GetDdosBgpInstancesInstanceInput` via:
//
//          GetDdosBgpInstancesInstanceArgs{...}
type GetDdosBgpInstancesInstanceInput interface {
	pulumi.Input

	ToGetDdosBgpInstancesInstanceOutput() GetDdosBgpInstancesInstanceOutput
	ToGetDdosBgpInstancesInstanceOutputWithContext(context.Context) GetDdosBgpInstancesInstanceOutput
}

type GetDdosBgpInstancesInstanceArgs struct {
	// The instance's elastic defend bandwidth.
	Bandwidth pulumi.IntInput `pulumi:"bandwidth"`
	// The instance's base defend bandwidth.
	BaseBandwidth pulumi.IntInput `pulumi:"baseBandwidth"`
	// The instance's id.
	Id pulumi.StringInput `pulumi:"id"`
	// The instance's count of ip config.
	IpCount pulumi.IntInput `pulumi:"ipCount"`
	// The instance's IP version.
	IpType pulumi.StringInput `pulumi:"ipType"`
	// The instance's remark.
	Name pulumi.StringInput `pulumi:"name"`
	// A region of instance.
	Region pulumi.StringInput `pulumi:"region"`
	// The instance's type.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetDdosBgpInstancesInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosBgpInstancesInstance)(nil)).Elem()
}

func (i GetDdosBgpInstancesInstanceArgs) ToGetDdosBgpInstancesInstanceOutput() GetDdosBgpInstancesInstanceOutput {
	return i.ToGetDdosBgpInstancesInstanceOutputWithContext(context.Background())
}

func (i GetDdosBgpInstancesInstanceArgs) ToGetDdosBgpInstancesInstanceOutputWithContext(ctx context.Context) GetDdosBgpInstancesInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDdosBgpInstancesInstanceOutput)
}

// GetDdosBgpInstancesInstanceArrayInput is an input type that accepts GetDdosBgpInstancesInstanceArray and GetDdosBgpInstancesInstanceArrayOutput values.
// You can construct a concrete instance of `GetDdosBgpInstancesInstanceArrayInput` via:
//
//          GetDdosBgpInstancesInstanceArray{ GetDdosBgpInstancesInstanceArgs{...} }
type GetDdosBgpInstancesInstanceArrayInput interface {
	pulumi.Input

	ToGetDdosBgpInstancesInstanceArrayOutput() GetDdosBgpInstancesInstanceArrayOutput
	ToGetDdosBgpInstancesInstanceArrayOutputWithContext(context.Context) GetDdosBgpInstancesInstanceArrayOutput
}

type GetDdosBgpInstancesInstanceArray []GetDdosBgpInstancesInstanceInput

func (GetDdosBgpInstancesInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDdosBgpInstancesInstance)(nil)).Elem()
}

func (i GetDdosBgpInstancesInstanceArray) ToGetDdosBgpInstancesInstanceArrayOutput() GetDdosBgpInstancesInstanceArrayOutput {
	return i.ToGetDdosBgpInstancesInstanceArrayOutputWithContext(context.Background())
}

func (i GetDdosBgpInstancesInstanceArray) ToGetDdosBgpInstancesInstanceArrayOutputWithContext(ctx context.Context) GetDdosBgpInstancesInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDdosBgpInstancesInstanceArrayOutput)
}

type GetDdosBgpInstancesInstanceOutput struct{ *pulumi.OutputState }

func (GetDdosBgpInstancesInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosBgpInstancesInstance)(nil)).Elem()
}

func (o GetDdosBgpInstancesInstanceOutput) ToGetDdosBgpInstancesInstanceOutput() GetDdosBgpInstancesInstanceOutput {
	return o
}

func (o GetDdosBgpInstancesInstanceOutput) ToGetDdosBgpInstancesInstanceOutputWithContext(ctx context.Context) GetDdosBgpInstancesInstanceOutput {
	return o
}

// The instance's elastic defend bandwidth.
func (o GetDdosBgpInstancesInstanceOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// The instance's base defend bandwidth.
func (o GetDdosBgpInstancesInstanceOutput) BaseBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) int { return v.BaseBandwidth }).(pulumi.IntOutput)
}

// The instance's id.
func (o GetDdosBgpInstancesInstanceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) string { return v.Id }).(pulumi.StringOutput)
}

// The instance's count of ip config.
func (o GetDdosBgpInstancesInstanceOutput) IpCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) int { return v.IpCount }).(pulumi.IntOutput)
}

// The instance's IP version.
func (o GetDdosBgpInstancesInstanceOutput) IpType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) string { return v.IpType }).(pulumi.StringOutput)
}

// The instance's remark.
func (o GetDdosBgpInstancesInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) string { return v.Name }).(pulumi.StringOutput)
}

// A region of instance.
func (o GetDdosBgpInstancesInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) string { return v.Region }).(pulumi.StringOutput)
}

// The instance's type.
func (o GetDdosBgpInstancesInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpInstancesInstance) string { return v.Type }).(pulumi.StringOutput)
}

type GetDdosBgpInstancesInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetDdosBgpInstancesInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDdosBgpInstancesInstance)(nil)).Elem()
}

func (o GetDdosBgpInstancesInstanceArrayOutput) ToGetDdosBgpInstancesInstanceArrayOutput() GetDdosBgpInstancesInstanceArrayOutput {
	return o
}

func (o GetDdosBgpInstancesInstanceArrayOutput) ToGetDdosBgpInstancesInstanceArrayOutputWithContext(ctx context.Context) GetDdosBgpInstancesInstanceArrayOutput {
	return o
}

func (o GetDdosBgpInstancesInstanceArrayOutput) Index(i pulumi.IntInput) GetDdosBgpInstancesInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDdosBgpInstancesInstance {
		return vs[0].([]GetDdosBgpInstancesInstance)[vs[1].(int)]
	}).(GetDdosBgpInstancesInstanceOutput)
}

type GetDdosCooInstancesInstance struct {
	// The instance's elastic defend bandwidth.
	Bandwidth int `pulumi:"bandwidth"`
	// The instance's base defend bandwidth.
	BaseBandwidth int `pulumi:"baseBandwidth"`
	// The instance's count of domain retransmission config.
	DomainCount int `pulumi:"domainCount"`
	// The instance's id.
	Id string `pulumi:"id"`
	// The instance's remark.
	Name string `pulumi:"name"`
	// The instance's count of port retransmission config.
	PortCount int `pulumi:"portCount"`
	// The instance's business bandwidth.
	ServiceBandwidth int `pulumi:"serviceBandwidth"`
}

// GetDdosCooInstancesInstanceInput is an input type that accepts GetDdosCooInstancesInstanceArgs and GetDdosCooInstancesInstanceOutput values.
// You can construct a concrete instance of `GetDdosCooInstancesInstanceInput` via:
//
//          GetDdosCooInstancesInstanceArgs{...}
type GetDdosCooInstancesInstanceInput interface {
	pulumi.Input

	ToGetDdosCooInstancesInstanceOutput() GetDdosCooInstancesInstanceOutput
	ToGetDdosCooInstancesInstanceOutputWithContext(context.Context) GetDdosCooInstancesInstanceOutput
}

type GetDdosCooInstancesInstanceArgs struct {
	// The instance's elastic defend bandwidth.
	Bandwidth pulumi.IntInput `pulumi:"bandwidth"`
	// The instance's base defend bandwidth.
	BaseBandwidth pulumi.IntInput `pulumi:"baseBandwidth"`
	// The instance's count of domain retransmission config.
	DomainCount pulumi.IntInput `pulumi:"domainCount"`
	// The instance's id.
	Id pulumi.StringInput `pulumi:"id"`
	// The instance's remark.
	Name pulumi.StringInput `pulumi:"name"`
	// The instance's count of port retransmission config.
	PortCount pulumi.IntInput `pulumi:"portCount"`
	// The instance's business bandwidth.
	ServiceBandwidth pulumi.IntInput `pulumi:"serviceBandwidth"`
}

func (GetDdosCooInstancesInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosCooInstancesInstance)(nil)).Elem()
}

func (i GetDdosCooInstancesInstanceArgs) ToGetDdosCooInstancesInstanceOutput() GetDdosCooInstancesInstanceOutput {
	return i.ToGetDdosCooInstancesInstanceOutputWithContext(context.Background())
}

func (i GetDdosCooInstancesInstanceArgs) ToGetDdosCooInstancesInstanceOutputWithContext(ctx context.Context) GetDdosCooInstancesInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDdosCooInstancesInstanceOutput)
}

// GetDdosCooInstancesInstanceArrayInput is an input type that accepts GetDdosCooInstancesInstanceArray and GetDdosCooInstancesInstanceArrayOutput values.
// You can construct a concrete instance of `GetDdosCooInstancesInstanceArrayInput` via:
//
//          GetDdosCooInstancesInstanceArray{ GetDdosCooInstancesInstanceArgs{...} }
type GetDdosCooInstancesInstanceArrayInput interface {
	pulumi.Input

	ToGetDdosCooInstancesInstanceArrayOutput() GetDdosCooInstancesInstanceArrayOutput
	ToGetDdosCooInstancesInstanceArrayOutputWithContext(context.Context) GetDdosCooInstancesInstanceArrayOutput
}

type GetDdosCooInstancesInstanceArray []GetDdosCooInstancesInstanceInput

func (GetDdosCooInstancesInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDdosCooInstancesInstance)(nil)).Elem()
}

func (i GetDdosCooInstancesInstanceArray) ToGetDdosCooInstancesInstanceArrayOutput() GetDdosCooInstancesInstanceArrayOutput {
	return i.ToGetDdosCooInstancesInstanceArrayOutputWithContext(context.Background())
}

func (i GetDdosCooInstancesInstanceArray) ToGetDdosCooInstancesInstanceArrayOutputWithContext(ctx context.Context) GetDdosCooInstancesInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDdosCooInstancesInstanceArrayOutput)
}

type GetDdosCooInstancesInstanceOutput struct{ *pulumi.OutputState }

func (GetDdosCooInstancesInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosCooInstancesInstance)(nil)).Elem()
}

func (o GetDdosCooInstancesInstanceOutput) ToGetDdosCooInstancesInstanceOutput() GetDdosCooInstancesInstanceOutput {
	return o
}

func (o GetDdosCooInstancesInstanceOutput) ToGetDdosCooInstancesInstanceOutputWithContext(ctx context.Context) GetDdosCooInstancesInstanceOutput {
	return o
}

// The instance's elastic defend bandwidth.
func (o GetDdosCooInstancesInstanceOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// The instance's base defend bandwidth.
func (o GetDdosCooInstancesInstanceOutput) BaseBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) int { return v.BaseBandwidth }).(pulumi.IntOutput)
}

// The instance's count of domain retransmission config.
func (o GetDdosCooInstancesInstanceOutput) DomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) int { return v.DomainCount }).(pulumi.IntOutput)
}

// The instance's id.
func (o GetDdosCooInstancesInstanceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) string { return v.Id }).(pulumi.StringOutput)
}

// The instance's remark.
func (o GetDdosCooInstancesInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) string { return v.Name }).(pulumi.StringOutput)
}

// The instance's count of port retransmission config.
func (o GetDdosCooInstancesInstanceOutput) PortCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) int { return v.PortCount }).(pulumi.IntOutput)
}

// The instance's business bandwidth.
func (o GetDdosCooInstancesInstanceOutput) ServiceBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetDdosCooInstancesInstance) int { return v.ServiceBandwidth }).(pulumi.IntOutput)
}

type GetDdosCooInstancesInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetDdosCooInstancesInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDdosCooInstancesInstance)(nil)).Elem()
}

func (o GetDdosCooInstancesInstanceArrayOutput) ToGetDdosCooInstancesInstanceArrayOutput() GetDdosCooInstancesInstanceArrayOutput {
	return o
}

func (o GetDdosCooInstancesInstanceArrayOutput) ToGetDdosCooInstancesInstanceArrayOutputWithContext(ctx context.Context) GetDdosCooInstancesInstanceArrayOutput {
	return o
}

func (o GetDdosCooInstancesInstanceArrayOutput) Index(i pulumi.IntInput) GetDdosCooInstancesInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDdosCooInstancesInstance {
		return vs[0].([]GetDdosCooInstancesInstance)[vs[1].(int)]
	}).(GetDdosCooInstancesInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(SchedulerRuleRuleOutput{})
	pulumi.RegisterOutputType(SchedulerRuleRuleArrayOutput{})
	pulumi.RegisterOutputType(GetDdosBgpInstancesInstanceOutput{})
	pulumi.RegisterOutputType(GetDdosBgpInstancesInstanceArrayOutput{})
	pulumi.RegisterOutputType(GetDdosCooInstancesInstanceOutput{})
	pulumi.RegisterOutputType(GetDdosCooInstancesInstanceArrayOutput{})
}
