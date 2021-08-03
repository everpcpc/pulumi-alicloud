// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetNamespacesNamespace struct {
	// The ID of the Namespace.
	Id string `pulumi:"id"`
	// The Description of Namespace.
	NamespaceDescription string `pulumi:"namespaceDescription"`
	// The Id of Namespace.It can contain 2 to 32 characters.The value is in format {RegionId}:{namespace}.
	NamespaceId string `pulumi:"namespaceId"`
	// The Name of Namespace.
	NamespaceName string `pulumi:"namespaceName"`
}

// GetNamespacesNamespaceInput is an input type that accepts GetNamespacesNamespaceArgs and GetNamespacesNamespaceOutput values.
// You can construct a concrete instance of `GetNamespacesNamespaceInput` via:
//
//          GetNamespacesNamespaceArgs{...}
type GetNamespacesNamespaceInput interface {
	pulumi.Input

	ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput
	ToGetNamespacesNamespaceOutputWithContext(context.Context) GetNamespacesNamespaceOutput
}

type GetNamespacesNamespaceArgs struct {
	// The ID of the Namespace.
	Id pulumi.StringInput `pulumi:"id"`
	// The Description of Namespace.
	NamespaceDescription pulumi.StringInput `pulumi:"namespaceDescription"`
	// The Id of Namespace.It can contain 2 to 32 characters.The value is in format {RegionId}:{namespace}.
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
	// The Name of Namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
}

func (GetNamespacesNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesNamespace)(nil)).Elem()
}

func (i GetNamespacesNamespaceArgs) ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput {
	return i.ToGetNamespacesNamespaceOutputWithContext(context.Background())
}

func (i GetNamespacesNamespaceArgs) ToGetNamespacesNamespaceOutputWithContext(ctx context.Context) GetNamespacesNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNamespacesNamespaceOutput)
}

// GetNamespacesNamespaceArrayInput is an input type that accepts GetNamespacesNamespaceArray and GetNamespacesNamespaceArrayOutput values.
// You can construct a concrete instance of `GetNamespacesNamespaceArrayInput` via:
//
//          GetNamespacesNamespaceArray{ GetNamespacesNamespaceArgs{...} }
type GetNamespacesNamespaceArrayInput interface {
	pulumi.Input

	ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput
	ToGetNamespacesNamespaceArrayOutputWithContext(context.Context) GetNamespacesNamespaceArrayOutput
}

type GetNamespacesNamespaceArray []GetNamespacesNamespaceInput

func (GetNamespacesNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNamespacesNamespace)(nil)).Elem()
}

func (i GetNamespacesNamespaceArray) ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput {
	return i.ToGetNamespacesNamespaceArrayOutputWithContext(context.Background())
}

func (i GetNamespacesNamespaceArray) ToGetNamespacesNamespaceArrayOutputWithContext(ctx context.Context) GetNamespacesNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNamespacesNamespaceArrayOutput)
}

type GetNamespacesNamespaceOutput struct{ *pulumi.OutputState }

func (GetNamespacesNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesNamespace)(nil)).Elem()
}

func (o GetNamespacesNamespaceOutput) ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput {
	return o
}

func (o GetNamespacesNamespaceOutput) ToGetNamespacesNamespaceOutputWithContext(ctx context.Context) GetNamespacesNamespaceOutput {
	return o
}

// The ID of the Namespace.
func (o GetNamespacesNamespaceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.Id }).(pulumi.StringOutput)
}

// The Description of Namespace.
func (o GetNamespacesNamespaceOutput) NamespaceDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceDescription }).(pulumi.StringOutput)
}

// The Id of Namespace.It can contain 2 to 32 characters.The value is in format {RegionId}:{namespace}.
func (o GetNamespacesNamespaceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// The Name of Namespace.
func (o GetNamespacesNamespaceOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceName }).(pulumi.StringOutput)
}

type GetNamespacesNamespaceArrayOutput struct{ *pulumi.OutputState }

func (GetNamespacesNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNamespacesNamespace)(nil)).Elem()
}

func (o GetNamespacesNamespaceArrayOutput) ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput {
	return o
}

func (o GetNamespacesNamespaceArrayOutput) ToGetNamespacesNamespaceArrayOutputWithContext(ctx context.Context) GetNamespacesNamespaceArrayOutput {
	return o
}

func (o GetNamespacesNamespaceArrayOutput) Index(i pulumi.IntInput) GetNamespacesNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNamespacesNamespace {
		return vs[0].([]GetNamespacesNamespace)[vs[1].(int)]
	}).(GetNamespacesNamespaceOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNamespacesNamespaceOutput{})
	pulumi.RegisterOutputType(GetNamespacesNamespaceArrayOutput{})
}