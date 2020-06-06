// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DomainLogHeader struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

// DomainLogHeaderInput is an input type that accepts DomainLogHeaderArgs and DomainLogHeaderOutput values.
// You can construct a concrete instance of `DomainLogHeaderInput` via:
//
// 		 DomainLogHeaderArgs{...}
//
type DomainLogHeaderInput interface {
	pulumi.Input

	ToDomainLogHeaderOutput() DomainLogHeaderOutput
	ToDomainLogHeaderOutputWithContext(context.Context) DomainLogHeaderOutput
}

type DomainLogHeaderArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DomainLogHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainLogHeader)(nil)).Elem()
}

func (i DomainLogHeaderArgs) ToDomainLogHeaderOutput() DomainLogHeaderOutput {
	return i.ToDomainLogHeaderOutputWithContext(context.Background())
}

func (i DomainLogHeaderArgs) ToDomainLogHeaderOutputWithContext(ctx context.Context) DomainLogHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainLogHeaderOutput)
}

// DomainLogHeaderArrayInput is an input type that accepts DomainLogHeaderArray and DomainLogHeaderArrayOutput values.
// You can construct a concrete instance of `DomainLogHeaderArrayInput` via:
//
// 		 DomainLogHeaderArray{ DomainLogHeaderArgs{...} }
//
type DomainLogHeaderArrayInput interface {
	pulumi.Input

	ToDomainLogHeaderArrayOutput() DomainLogHeaderArrayOutput
	ToDomainLogHeaderArrayOutputWithContext(context.Context) DomainLogHeaderArrayOutput
}

type DomainLogHeaderArray []DomainLogHeaderInput

func (DomainLogHeaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainLogHeader)(nil)).Elem()
}

func (i DomainLogHeaderArray) ToDomainLogHeaderArrayOutput() DomainLogHeaderArrayOutput {
	return i.ToDomainLogHeaderArrayOutputWithContext(context.Background())
}

func (i DomainLogHeaderArray) ToDomainLogHeaderArrayOutputWithContext(ctx context.Context) DomainLogHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainLogHeaderArrayOutput)
}

type DomainLogHeaderOutput struct{ *pulumi.OutputState }

func (DomainLogHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainLogHeader)(nil)).Elem()
}

func (o DomainLogHeaderOutput) ToDomainLogHeaderOutput() DomainLogHeaderOutput {
	return o
}

func (o DomainLogHeaderOutput) ToDomainLogHeaderOutputWithContext(ctx context.Context) DomainLogHeaderOutput {
	return o
}

func (o DomainLogHeaderOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainLogHeader) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o DomainLogHeaderOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainLogHeader) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DomainLogHeaderArrayOutput struct{ *pulumi.OutputState }

func (DomainLogHeaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainLogHeader)(nil)).Elem()
}

func (o DomainLogHeaderArrayOutput) ToDomainLogHeaderArrayOutput() DomainLogHeaderArrayOutput {
	return o
}

func (o DomainLogHeaderArrayOutput) ToDomainLogHeaderArrayOutputWithContext(ctx context.Context) DomainLogHeaderArrayOutput {
	return o
}

func (o DomainLogHeaderArrayOutput) Index(i pulumi.IntInput) DomainLogHeaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainLogHeader {
		return vs[0].([]DomainLogHeader)[vs[1].(int)]
	}).(DomainLogHeaderOutput)
}

type GetDomainsDomain struct {
	// Name of the domain.
	Domain string `pulumi:"domain"`
}

// GetDomainsDomainInput is an input type that accepts GetDomainsDomainArgs and GetDomainsDomainOutput values.
// You can construct a concrete instance of `GetDomainsDomainInput` via:
//
// 		 GetDomainsDomainArgs{...}
//
type GetDomainsDomainInput interface {
	pulumi.Input

	ToGetDomainsDomainOutput() GetDomainsDomainOutput
	ToGetDomainsDomainOutputWithContext(context.Context) GetDomainsDomainOutput
}

type GetDomainsDomainArgs struct {
	// Name of the domain.
	Domain pulumi.StringInput `pulumi:"domain"`
}

func (GetDomainsDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return i.ToGetDomainsDomainOutputWithContext(context.Background())
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainOutput)
}

// GetDomainsDomainArrayInput is an input type that accepts GetDomainsDomainArray and GetDomainsDomainArrayOutput values.
// You can construct a concrete instance of `GetDomainsDomainArrayInput` via:
//
// 		 GetDomainsDomainArray{ GetDomainsDomainArgs{...} }
//
type GetDomainsDomainArrayInput interface {
	pulumi.Input

	ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput
	ToGetDomainsDomainArrayOutputWithContext(context.Context) GetDomainsDomainArrayOutput
}

type GetDomainsDomainArray []GetDomainsDomainInput

func (GetDomainsDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return i.ToGetDomainsDomainArrayOutputWithContext(context.Background())
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainArrayOutput)
}

type GetDomainsDomainOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return o
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return o
}

// Name of the domain.
func (o GetDomainsDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Domain }).(pulumi.StringOutput)
}

type GetDomainsDomainArrayOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) Index(i pulumi.IntInput) GetDomainsDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDomainsDomain {
		return vs[0].([]GetDomainsDomain)[vs[1].(int)]
	}).(GetDomainsDomainOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainLogHeaderOutput{})
	pulumi.RegisterOutputType(DomainLogHeaderArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainArrayOutput{})
}