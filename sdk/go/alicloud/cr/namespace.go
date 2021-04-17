// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manager Container Registry namespaces.
//
// > **NOTE:** Available in v1.34.0+.
//
// > **NOTE:** You need to set your registry password in Container Registry console before use this resource.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cr.NewNamespace(ctx, "my_namespace", &cr.NamespaceArgs{
// 			AutoCreate:        pulumi.Bool(false),
// 			DefaultVisibility: pulumi.String("PUBLIC"),
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
// Container Registry namespace can be imported using the namespace, e.g.
//
// ```sh
//  $ pulumi import alicloud:cr/namespace:Namespace default my-namespace
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolOutput `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringOutput `pulumi:"defaultVisibility"`
	// Name of Container Registry namespace.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoCreate == nil {
		return nil, errors.New("invalid value for required argument 'AutoCreate'")
	}
	if args.DefaultVisibility == nil {
		return nil, errors.New("invalid value for required argument 'DefaultVisibility'")
	}
	var resource Namespace
	err := ctx.RegisterResource("alicloud:cr/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("alicloud:cr/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate *bool `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility *string `pulumi:"defaultVisibility"`
	// Name of Container Registry namespace.
	Name *string `pulumi:"name"`
}

type NamespaceState struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolPtrInput
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringPtrInput
	// Name of Container Registry namespace.
	Name pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate bool `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility string `pulumi:"defaultVisibility"`
	// Name of Container Registry namespace.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolInput
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringInput
	// Name of Container Registry namespace.
	Name pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

func (i *Namespace) ToNamespacePtrOutput() NamespacePtrOutput {
	return i.ToNamespacePtrOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePtrOutput)
}

type NamespacePtrInput interface {
	pulumi.Input

	ToNamespacePtrOutput() NamespacePtrOutput
	ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput
}

type namespacePtrType NamespaceArgs

func (*namespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil))
}

func (i *namespacePtrType) ToNamespacePtrOutput() NamespacePtrOutput {
	return i.ToNamespacePtrOutputWithContext(context.Background())
}

func (i *namespacePtrType) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePtrOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//          NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Namespace)(nil))
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//          NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Namespace)(nil))
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespacePtrOutput() NamespacePtrOutput {
	return o.ToNamespacePtrOutputWithContext(context.Background())
}

func (o NamespaceOutput) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return o.ApplyT(func(v Namespace) *Namespace {
		return &v
	}).(NamespacePtrOutput)
}

type NamespacePtrOutput struct {
	*pulumi.OutputState
}

func (NamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil))
}

func (o NamespacePtrOutput) ToNamespacePtrOutput() NamespacePtrOutput {
	return o
}

func (o NamespacePtrOutput) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return o
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Namespace)(nil))
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Namespace {
		return vs[0].([]Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Namespace)(nil))
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Namespace {
		return vs[0].(map[string]Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespacePtrOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
