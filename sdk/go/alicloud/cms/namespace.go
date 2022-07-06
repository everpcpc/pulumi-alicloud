// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Namespace resource.
//
// For information about Cloud Monitor Service Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/doc-detail/28608.htm).
//
// > **NOTE:** Available in v1.171.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cms.NewNamespace(ctx, "example", &cms.NamespaceArgs{
// 			Namespace:     pulumi.String("example_value"),
// 			Specification: pulumi.String("cms.s1.large"),
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
// Cloud Monitor Service Namespace can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cms/namespace:Namespace example <namespace>
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// Description of indicator warehouse.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicator warehouse name. The namespace can contain lowercase letters, digits, and hyphens (-).
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Data storage duration. Valid values: `cms.s1.12xlarge`, `cms.s1.2xlarge`, `cms.s1.3xlarge`, `cms.s1.6xlarge`, `cms.s1.large`, `cms.s1.xlarge`.
	// - `cms.s1.large`: Data storage duration is 15 days.
	// - `cms.s1.xlarge`: Data storage duration is 32 days.
	// - `cms.s1.2xlarge`: Data storage duration 63 days.
	// - `cms.s1.3xlarge`: (Default) Data storage duration 93 days.
	// - `cms.s1.6xlarge`: Data storage duration 185 days.
	// - `cms.s1.12xlarge`: Data storage duration 376 days.
	Specification pulumi.StringOutput `pulumi:"specification"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	var resource Namespace
	err := ctx.RegisterResource("alicloud:cms/namespace:Namespace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:cms/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Description of indicator warehouse.
	Description *string `pulumi:"description"`
	// Indicator warehouse name. The namespace can contain lowercase letters, digits, and hyphens (-).
	Namespace *string `pulumi:"namespace"`
	// Data storage duration. Valid values: `cms.s1.12xlarge`, `cms.s1.2xlarge`, `cms.s1.3xlarge`, `cms.s1.6xlarge`, `cms.s1.large`, `cms.s1.xlarge`.
	// - `cms.s1.large`: Data storage duration is 15 days.
	// - `cms.s1.xlarge`: Data storage duration is 32 days.
	// - `cms.s1.2xlarge`: Data storage duration 63 days.
	// - `cms.s1.3xlarge`: (Default) Data storage duration 93 days.
	// - `cms.s1.6xlarge`: Data storage duration 185 days.
	// - `cms.s1.12xlarge`: Data storage duration 376 days.
	Specification *string `pulumi:"specification"`
}

type NamespaceState struct {
	// Description of indicator warehouse.
	Description pulumi.StringPtrInput
	// Indicator warehouse name. The namespace can contain lowercase letters, digits, and hyphens (-).
	Namespace pulumi.StringPtrInput
	// Data storage duration. Valid values: `cms.s1.12xlarge`, `cms.s1.2xlarge`, `cms.s1.3xlarge`, `cms.s1.6xlarge`, `cms.s1.large`, `cms.s1.xlarge`.
	// - `cms.s1.large`: Data storage duration is 15 days.
	// - `cms.s1.xlarge`: Data storage duration is 32 days.
	// - `cms.s1.2xlarge`: Data storage duration 63 days.
	// - `cms.s1.3xlarge`: (Default) Data storage duration 93 days.
	// - `cms.s1.6xlarge`: Data storage duration 185 days.
	// - `cms.s1.12xlarge`: Data storage duration 376 days.
	Specification pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Description of indicator warehouse.
	Description *string `pulumi:"description"`
	// Indicator warehouse name. The namespace can contain lowercase letters, digits, and hyphens (-).
	Namespace string `pulumi:"namespace"`
	// Data storage duration. Valid values: `cms.s1.12xlarge`, `cms.s1.2xlarge`, `cms.s1.3xlarge`, `cms.s1.6xlarge`, `cms.s1.large`, `cms.s1.xlarge`.
	// - `cms.s1.large`: Data storage duration is 15 days.
	// - `cms.s1.xlarge`: Data storage duration is 32 days.
	// - `cms.s1.2xlarge`: Data storage duration 63 days.
	// - `cms.s1.3xlarge`: (Default) Data storage duration 93 days.
	// - `cms.s1.6xlarge`: Data storage duration 185 days.
	// - `cms.s1.12xlarge`: Data storage duration 376 days.
	Specification *string `pulumi:"specification"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Description of indicator warehouse.
	Description pulumi.StringPtrInput
	// Indicator warehouse name. The namespace can contain lowercase letters, digits, and hyphens (-).
	Namespace pulumi.StringInput
	// Data storage duration. Valid values: `cms.s1.12xlarge`, `cms.s1.2xlarge`, `cms.s1.3xlarge`, `cms.s1.6xlarge`, `cms.s1.large`, `cms.s1.xlarge`.
	// - `cms.s1.large`: Data storage duration is 15 days.
	// - `cms.s1.xlarge`: Data storage duration is 32 days.
	// - `cms.s1.2xlarge`: Data storage duration 63 days.
	// - `cms.s1.3xlarge`: (Default) Data storage duration 93 days.
	// - `cms.s1.6xlarge`: Data storage duration 185 days.
	// - `cms.s1.12xlarge`: Data storage duration 376 days.
	Specification pulumi.StringPtrInput
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
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
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
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
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
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}