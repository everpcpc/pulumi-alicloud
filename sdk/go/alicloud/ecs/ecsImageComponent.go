// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Image Component resource.
//
// For information about ECS Image Component and how to use it, see [What is Image Component](https://www.alibabacloud.com/help/en/doc-detail/200424.htm).
//
// > **NOTE:** Available in v1.159.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				NameRegex: pulumi.StringRef("default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsImageComponent(ctx, "example", &ecs.EcsImageComponentArgs{
//				ComponentType:      pulumi.String("Build"),
//				Content:            pulumi.String("RUN yum update -y"),
//				Description:        pulumi.String("example_value"),
//				ImageComponentName: pulumi.String("example_value"),
//				ResourceGroupId:    *pulumi.String(_default.Groups[0].Id),
//				SystemType:         pulumi.String("Linux"),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
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
//
// ## Import
//
// ECS Image Component can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/ecsImageComponent:EcsImageComponent example <id>
//
// ```
type EcsImageComponent struct {
	pulumi.CustomResourceState

	// The type of the image component. Only image building components are supported. Valid values: `Build`.
	ComponentType pulumi.StringOutput `pulumi:"componentType"`
	// The content of the image component. The content can consist of up to 127 commands.
	Content pulumi.StringOutput `pulumi:"content"`
	// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	ImageComponentName pulumi.StringOutput `pulumi:"imageComponentName"`
	// The ID of the resource group to which to assign the image component.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
	SystemType pulumi.StringOutput `pulumi:"systemType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEcsImageComponent registers a new resource with the given unique name, arguments, and options.
func NewEcsImageComponent(ctx *pulumi.Context,
	name string, args *EcsImageComponentArgs, opts ...pulumi.ResourceOption) (*EcsImageComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	var resource EcsImageComponent
	err := ctx.RegisterResource("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsImageComponent gets an existing EcsImageComponent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsImageComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsImageComponentState, opts ...pulumi.ResourceOption) (*EcsImageComponent, error) {
	var resource EcsImageComponent
	err := ctx.ReadResource("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsImageComponent resources.
type ecsImageComponentState struct {
	// The type of the image component. Only image building components are supported. Valid values: `Build`.
	ComponentType *string `pulumi:"componentType"`
	// The content of the image component. The content can consist of up to 127 commands.
	Content *string `pulumi:"content"`
	// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	ImageComponentName *string `pulumi:"imageComponentName"`
	// The ID of the resource group to which to assign the image component.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
	SystemType *string `pulumi:"systemType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EcsImageComponentState struct {
	// The type of the image component. Only image building components are supported. Valid values: `Build`.
	ComponentType pulumi.StringPtrInput
	// The content of the image component. The content can consist of up to 127 commands.
	Content pulumi.StringPtrInput
	// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	ImageComponentName pulumi.StringPtrInput
	// The ID of the resource group to which to assign the image component.
	ResourceGroupId pulumi.StringPtrInput
	// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
	SystemType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EcsImageComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsImageComponentState)(nil)).Elem()
}

type ecsImageComponentArgs struct {
	// The type of the image component. Only image building components are supported. Valid values: `Build`.
	ComponentType *string `pulumi:"componentType"`
	// The content of the image component. The content can consist of up to 127 commands.
	Content string `pulumi:"content"`
	// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	ImageComponentName *string `pulumi:"imageComponentName"`
	// The ID of the resource group to which to assign the image component.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
	SystemType *string `pulumi:"systemType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EcsImageComponent resource.
type EcsImageComponentArgs struct {
	// The type of the image component. Only image building components are supported. Valid values: `Build`.
	ComponentType pulumi.StringPtrInput
	// The content of the image component. The content can consist of up to 127 commands.
	Content pulumi.StringInput
	// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	ImageComponentName pulumi.StringPtrInput
	// The ID of the resource group to which to assign the image component.
	ResourceGroupId pulumi.StringPtrInput
	// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
	SystemType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EcsImageComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsImageComponentArgs)(nil)).Elem()
}

type EcsImageComponentInput interface {
	pulumi.Input

	ToEcsImageComponentOutput() EcsImageComponentOutput
	ToEcsImageComponentOutputWithContext(ctx context.Context) EcsImageComponentOutput
}

func (*EcsImageComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsImageComponent)(nil)).Elem()
}

func (i *EcsImageComponent) ToEcsImageComponentOutput() EcsImageComponentOutput {
	return i.ToEcsImageComponentOutputWithContext(context.Background())
}

func (i *EcsImageComponent) ToEcsImageComponentOutputWithContext(ctx context.Context) EcsImageComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsImageComponentOutput)
}

// EcsImageComponentArrayInput is an input type that accepts EcsImageComponentArray and EcsImageComponentArrayOutput values.
// You can construct a concrete instance of `EcsImageComponentArrayInput` via:
//
//	EcsImageComponentArray{ EcsImageComponentArgs{...} }
type EcsImageComponentArrayInput interface {
	pulumi.Input

	ToEcsImageComponentArrayOutput() EcsImageComponentArrayOutput
	ToEcsImageComponentArrayOutputWithContext(context.Context) EcsImageComponentArrayOutput
}

type EcsImageComponentArray []EcsImageComponentInput

func (EcsImageComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsImageComponent)(nil)).Elem()
}

func (i EcsImageComponentArray) ToEcsImageComponentArrayOutput() EcsImageComponentArrayOutput {
	return i.ToEcsImageComponentArrayOutputWithContext(context.Background())
}

func (i EcsImageComponentArray) ToEcsImageComponentArrayOutputWithContext(ctx context.Context) EcsImageComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsImageComponentArrayOutput)
}

// EcsImageComponentMapInput is an input type that accepts EcsImageComponentMap and EcsImageComponentMapOutput values.
// You can construct a concrete instance of `EcsImageComponentMapInput` via:
//
//	EcsImageComponentMap{ "key": EcsImageComponentArgs{...} }
type EcsImageComponentMapInput interface {
	pulumi.Input

	ToEcsImageComponentMapOutput() EcsImageComponentMapOutput
	ToEcsImageComponentMapOutputWithContext(context.Context) EcsImageComponentMapOutput
}

type EcsImageComponentMap map[string]EcsImageComponentInput

func (EcsImageComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsImageComponent)(nil)).Elem()
}

func (i EcsImageComponentMap) ToEcsImageComponentMapOutput() EcsImageComponentMapOutput {
	return i.ToEcsImageComponentMapOutputWithContext(context.Background())
}

func (i EcsImageComponentMap) ToEcsImageComponentMapOutputWithContext(ctx context.Context) EcsImageComponentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsImageComponentMapOutput)
}

type EcsImageComponentOutput struct{ *pulumi.OutputState }

func (EcsImageComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsImageComponent)(nil)).Elem()
}

func (o EcsImageComponentOutput) ToEcsImageComponentOutput() EcsImageComponentOutput {
	return o
}

func (o EcsImageComponentOutput) ToEcsImageComponentOutputWithContext(ctx context.Context) EcsImageComponentOutput {
	return o
}

// The type of the image component. Only image building components are supported. Valid values: `Build`.
func (o EcsImageComponentOutput) ComponentType() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringOutput { return v.ComponentType }).(pulumi.StringOutput)
}

// The content of the image component. The content can consist of up to 127 commands.
func (o EcsImageComponentOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The description of the image component. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
func (o EcsImageComponentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the image component. The name must be `2` to `128` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
func (o EcsImageComponentOutput) ImageComponentName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringOutput { return v.ImageComponentName }).(pulumi.StringOutput)
}

// The ID of the resource group to which to assign the image component.
func (o EcsImageComponentOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The operating system type supported by the image component. Only Linux is supported. Valid values: `Linux`.
func (o EcsImageComponentOutput) SystemType() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.StringOutput { return v.SystemType }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o EcsImageComponentOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsImageComponent) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type EcsImageComponentArrayOutput struct{ *pulumi.OutputState }

func (EcsImageComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsImageComponent)(nil)).Elem()
}

func (o EcsImageComponentArrayOutput) ToEcsImageComponentArrayOutput() EcsImageComponentArrayOutput {
	return o
}

func (o EcsImageComponentArrayOutput) ToEcsImageComponentArrayOutputWithContext(ctx context.Context) EcsImageComponentArrayOutput {
	return o
}

func (o EcsImageComponentArrayOutput) Index(i pulumi.IntInput) EcsImageComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsImageComponent {
		return vs[0].([]*EcsImageComponent)[vs[1].(int)]
	}).(EcsImageComponentOutput)
}

type EcsImageComponentMapOutput struct{ *pulumi.OutputState }

func (EcsImageComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsImageComponent)(nil)).Elem()
}

func (o EcsImageComponentMapOutput) ToEcsImageComponentMapOutput() EcsImageComponentMapOutput {
	return o
}

func (o EcsImageComponentMapOutput) ToEcsImageComponentMapOutputWithContext(ctx context.Context) EcsImageComponentMapOutput {
	return o
}

func (o EcsImageComponentMapOutput) MapIndex(k pulumi.StringInput) EcsImageComponentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsImageComponent {
		return vs[0].(map[string]*EcsImageComponent)[vs[1].(string)]
	}).(EcsImageComponentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsImageComponentInput)(nil)).Elem(), &EcsImageComponent{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsImageComponentArrayInput)(nil)).Elem(), EcsImageComponentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsImageComponentMapInput)(nil)).Elem(), EcsImageComponentMap{})
	pulumi.RegisterOutputType(EcsImageComponentOutput{})
	pulumi.RegisterOutputType(EcsImageComponentArrayOutput{})
	pulumi.RegisterOutputType(EcsImageComponentMapOutput{})
}
