// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sag

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetAclsAcl struct {
	// The ID of the ACL. For example "acl-xxx".
	Id string `pulumi:"id"`
	// The name of the Acl.
	Name string `pulumi:"name"`
}

// GetAclsAclInput is an input type that accepts GetAclsAclArgs and GetAclsAclOutput values.
// You can construct a concrete instance of `GetAclsAclInput` via:
//
//          GetAclsAclArgs{...}
type GetAclsAclInput interface {
	pulumi.Input

	ToGetAclsAclOutput() GetAclsAclOutput
	ToGetAclsAclOutputWithContext(context.Context) GetAclsAclOutput
}

type GetAclsAclArgs struct {
	// The ID of the ACL. For example "acl-xxx".
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the Acl.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetAclsAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAclsAcl)(nil)).Elem()
}

func (i GetAclsAclArgs) ToGetAclsAclOutput() GetAclsAclOutput {
	return i.ToGetAclsAclOutputWithContext(context.Background())
}

func (i GetAclsAclArgs) ToGetAclsAclOutputWithContext(ctx context.Context) GetAclsAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAclsAclOutput)
}

// GetAclsAclArrayInput is an input type that accepts GetAclsAclArray and GetAclsAclArrayOutput values.
// You can construct a concrete instance of `GetAclsAclArrayInput` via:
//
//          GetAclsAclArray{ GetAclsAclArgs{...} }
type GetAclsAclArrayInput interface {
	pulumi.Input

	ToGetAclsAclArrayOutput() GetAclsAclArrayOutput
	ToGetAclsAclArrayOutputWithContext(context.Context) GetAclsAclArrayOutput
}

type GetAclsAclArray []GetAclsAclInput

func (GetAclsAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAclsAcl)(nil)).Elem()
}

func (i GetAclsAclArray) ToGetAclsAclArrayOutput() GetAclsAclArrayOutput {
	return i.ToGetAclsAclArrayOutputWithContext(context.Background())
}

func (i GetAclsAclArray) ToGetAclsAclArrayOutputWithContext(ctx context.Context) GetAclsAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAclsAclArrayOutput)
}

type GetAclsAclOutput struct{ *pulumi.OutputState }

func (GetAclsAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAclsAcl)(nil)).Elem()
}

func (o GetAclsAclOutput) ToGetAclsAclOutput() GetAclsAclOutput {
	return o
}

func (o GetAclsAclOutput) ToGetAclsAclOutputWithContext(ctx context.Context) GetAclsAclOutput {
	return o
}

// The ID of the ACL. For example "acl-xxx".
func (o GetAclsAclOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAclsAcl) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Acl.
func (o GetAclsAclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAclsAcl) string { return v.Name }).(pulumi.StringOutput)
}

type GetAclsAclArrayOutput struct{ *pulumi.OutputState }

func (GetAclsAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAclsAcl)(nil)).Elem()
}

func (o GetAclsAclArrayOutput) ToGetAclsAclArrayOutput() GetAclsAclArrayOutput {
	return o
}

func (o GetAclsAclArrayOutput) ToGetAclsAclArrayOutputWithContext(ctx context.Context) GetAclsAclArrayOutput {
	return o
}

func (o GetAclsAclArrayOutput) Index(i pulumi.IntInput) GetAclsAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAclsAcl {
		return vs[0].([]GetAclsAcl)[vs[1].(int)]
	}).(GetAclsAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetAclsAclInput)(nil)).Elem(), GetAclsAclArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAclsAclArrayInput)(nil)).Elem(), GetAclsAclArray{})
	pulumi.RegisterOutputType(GetAclsAclOutput{})
	pulumi.RegisterOutputType(GetAclsAclArrayOutput{})
}
