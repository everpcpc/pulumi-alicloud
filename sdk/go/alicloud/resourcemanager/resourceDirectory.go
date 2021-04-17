// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Manager Resource Directory resource. Resource Directory enables you to establish an organizational structure for the resources used by applications of your enterprise. You can plan, build, and manage the resources in a centralized manner by using only one resource directory.
//
// For information about Resource Manager Resource Directory and how to use it, see [What is Resource Manager Resource Directory](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
//
// > **NOTE:** Available in v1.84.0+.
//
// > **NOTE:** An account can only be used to enable a resource directory after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.
//
// > **NOTE:** Before you destroy the resource, make sure that the following requirements are met:
//   - All member accounts must be removed from the resource directory.
//   - All folders except the root folder must be deleted from the resource directory.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcemanager.NewResourceDirectory(ctx, "example", &resourcemanager.ResourceDirectoryArgs{
// 			Status: pulumi.String("Enabled"),
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
// Resource Manager Resource Directory can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/resourceDirectory:ResourceDirectory example rd-s3****
// ```
type ResourceDirectory struct {
	pulumi.CustomResourceState

	// The ID of the master account.
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`
	// The name of the master account.
	MasterAccountName pulumi.StringOutput `pulumi:"masterAccountName"`
	// The ID of the root folder.
	RootFolderId pulumi.StringOutput `pulumi:"rootFolderId"`
	// The status of control policy. Valid values:`Enabled` and `Disabled`. Default value is `Disabled`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewResourceDirectory registers a new resource with the given unique name, arguments, and options.
func NewResourceDirectory(ctx *pulumi.Context,
	name string, args *ResourceDirectoryArgs, opts ...pulumi.ResourceOption) (*ResourceDirectory, error) {
	if args == nil {
		args = &ResourceDirectoryArgs{}
	}

	var resource ResourceDirectory
	err := ctx.RegisterResource("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDirectory gets an existing ResourceDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDirectoryState, opts ...pulumi.ResourceOption) (*ResourceDirectory, error) {
	var resource ResourceDirectory
	err := ctx.ReadResource("alicloud:resourcemanager/resourceDirectory:ResourceDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDirectory resources.
type resourceDirectoryState struct {
	// The ID of the master account.
	MasterAccountId *string `pulumi:"masterAccountId"`
	// The name of the master account.
	MasterAccountName *string `pulumi:"masterAccountName"`
	// The ID of the root folder.
	RootFolderId *string `pulumi:"rootFolderId"`
	// The status of control policy. Valid values:`Enabled` and `Disabled`. Default value is `Disabled`.
	Status *string `pulumi:"status"`
}

type ResourceDirectoryState struct {
	// The ID of the master account.
	MasterAccountId pulumi.StringPtrInput
	// The name of the master account.
	MasterAccountName pulumi.StringPtrInput
	// The ID of the root folder.
	RootFolderId pulumi.StringPtrInput
	// The status of control policy. Valid values:`Enabled` and `Disabled`. Default value is `Disabled`.
	Status pulumi.StringPtrInput
}

func (ResourceDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDirectoryState)(nil)).Elem()
}

type resourceDirectoryArgs struct {
	// The status of control policy. Valid values:`Enabled` and `Disabled`. Default value is `Disabled`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a ResourceDirectory resource.
type ResourceDirectoryArgs struct {
	// The status of control policy. Valid values:`Enabled` and `Disabled`. Default value is `Disabled`.
	Status pulumi.StringPtrInput
}

func (ResourceDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDirectoryArgs)(nil)).Elem()
}

type ResourceDirectoryInput interface {
	pulumi.Input

	ToResourceDirectoryOutput() ResourceDirectoryOutput
	ToResourceDirectoryOutputWithContext(ctx context.Context) ResourceDirectoryOutput
}

func (*ResourceDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDirectory)(nil))
}

func (i *ResourceDirectory) ToResourceDirectoryOutput() ResourceDirectoryOutput {
	return i.ToResourceDirectoryOutputWithContext(context.Background())
}

func (i *ResourceDirectory) ToResourceDirectoryOutputWithContext(ctx context.Context) ResourceDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDirectoryOutput)
}

func (i *ResourceDirectory) ToResourceDirectoryPtrOutput() ResourceDirectoryPtrOutput {
	return i.ToResourceDirectoryPtrOutputWithContext(context.Background())
}

func (i *ResourceDirectory) ToResourceDirectoryPtrOutputWithContext(ctx context.Context) ResourceDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDirectoryPtrOutput)
}

type ResourceDirectoryPtrInput interface {
	pulumi.Input

	ToResourceDirectoryPtrOutput() ResourceDirectoryPtrOutput
	ToResourceDirectoryPtrOutputWithContext(ctx context.Context) ResourceDirectoryPtrOutput
}

type resourceDirectoryPtrType ResourceDirectoryArgs

func (*resourceDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDirectory)(nil))
}

func (i *resourceDirectoryPtrType) ToResourceDirectoryPtrOutput() ResourceDirectoryPtrOutput {
	return i.ToResourceDirectoryPtrOutputWithContext(context.Background())
}

func (i *resourceDirectoryPtrType) ToResourceDirectoryPtrOutputWithContext(ctx context.Context) ResourceDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDirectoryPtrOutput)
}

// ResourceDirectoryArrayInput is an input type that accepts ResourceDirectoryArray and ResourceDirectoryArrayOutput values.
// You can construct a concrete instance of `ResourceDirectoryArrayInput` via:
//
//          ResourceDirectoryArray{ ResourceDirectoryArgs{...} }
type ResourceDirectoryArrayInput interface {
	pulumi.Input

	ToResourceDirectoryArrayOutput() ResourceDirectoryArrayOutput
	ToResourceDirectoryArrayOutputWithContext(context.Context) ResourceDirectoryArrayOutput
}

type ResourceDirectoryArray []ResourceDirectoryInput

func (ResourceDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResourceDirectory)(nil))
}

func (i ResourceDirectoryArray) ToResourceDirectoryArrayOutput() ResourceDirectoryArrayOutput {
	return i.ToResourceDirectoryArrayOutputWithContext(context.Background())
}

func (i ResourceDirectoryArray) ToResourceDirectoryArrayOutputWithContext(ctx context.Context) ResourceDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDirectoryArrayOutput)
}

// ResourceDirectoryMapInput is an input type that accepts ResourceDirectoryMap and ResourceDirectoryMapOutput values.
// You can construct a concrete instance of `ResourceDirectoryMapInput` via:
//
//          ResourceDirectoryMap{ "key": ResourceDirectoryArgs{...} }
type ResourceDirectoryMapInput interface {
	pulumi.Input

	ToResourceDirectoryMapOutput() ResourceDirectoryMapOutput
	ToResourceDirectoryMapOutputWithContext(context.Context) ResourceDirectoryMapOutput
}

type ResourceDirectoryMap map[string]ResourceDirectoryInput

func (ResourceDirectoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResourceDirectory)(nil))
}

func (i ResourceDirectoryMap) ToResourceDirectoryMapOutput() ResourceDirectoryMapOutput {
	return i.ToResourceDirectoryMapOutputWithContext(context.Background())
}

func (i ResourceDirectoryMap) ToResourceDirectoryMapOutputWithContext(ctx context.Context) ResourceDirectoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDirectoryMapOutput)
}

type ResourceDirectoryOutput struct {
	*pulumi.OutputState
}

func (ResourceDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDirectory)(nil))
}

func (o ResourceDirectoryOutput) ToResourceDirectoryOutput() ResourceDirectoryOutput {
	return o
}

func (o ResourceDirectoryOutput) ToResourceDirectoryOutputWithContext(ctx context.Context) ResourceDirectoryOutput {
	return o
}

func (o ResourceDirectoryOutput) ToResourceDirectoryPtrOutput() ResourceDirectoryPtrOutput {
	return o.ToResourceDirectoryPtrOutputWithContext(context.Background())
}

func (o ResourceDirectoryOutput) ToResourceDirectoryPtrOutputWithContext(ctx context.Context) ResourceDirectoryPtrOutput {
	return o.ApplyT(func(v ResourceDirectory) *ResourceDirectory {
		return &v
	}).(ResourceDirectoryPtrOutput)
}

type ResourceDirectoryPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDirectory)(nil))
}

func (o ResourceDirectoryPtrOutput) ToResourceDirectoryPtrOutput() ResourceDirectoryPtrOutput {
	return o
}

func (o ResourceDirectoryPtrOutput) ToResourceDirectoryPtrOutputWithContext(ctx context.Context) ResourceDirectoryPtrOutput {
	return o
}

type ResourceDirectoryArrayOutput struct{ *pulumi.OutputState }

func (ResourceDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceDirectory)(nil))
}

func (o ResourceDirectoryArrayOutput) ToResourceDirectoryArrayOutput() ResourceDirectoryArrayOutput {
	return o
}

func (o ResourceDirectoryArrayOutput) ToResourceDirectoryArrayOutputWithContext(ctx context.Context) ResourceDirectoryArrayOutput {
	return o
}

func (o ResourceDirectoryArrayOutput) Index(i pulumi.IntInput) ResourceDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceDirectory {
		return vs[0].([]ResourceDirectory)[vs[1].(int)]
	}).(ResourceDirectoryOutput)
}

type ResourceDirectoryMapOutput struct{ *pulumi.OutputState }

func (ResourceDirectoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceDirectory)(nil))
}

func (o ResourceDirectoryMapOutput) ToResourceDirectoryMapOutput() ResourceDirectoryMapOutput {
	return o
}

func (o ResourceDirectoryMapOutput) ToResourceDirectoryMapOutputWithContext(ctx context.Context) ResourceDirectoryMapOutput {
	return o
}

func (o ResourceDirectoryMapOutput) MapIndex(k pulumi.StringInput) ResourceDirectoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceDirectory {
		return vs[0].(map[string]ResourceDirectory)[vs[1].(string)]
	}).(ResourceDirectoryOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceDirectoryOutput{})
	pulumi.RegisterOutputType(ResourceDirectoryPtrOutput{})
	pulumi.RegisterOutputType(ResourceDirectoryArrayOutput{})
	pulumi.RegisterOutputType(ResourceDirectoryMapOutput{})
}
