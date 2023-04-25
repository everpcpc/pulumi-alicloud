// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS Snapshot resource.
//
// For information about ECS Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/25524.htm).
//
// > **NOTE:** Available in v1.120.0+.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewEcsSnapshot(ctx, "default", &ecs.EcsSnapshotArgs{
//				Category:      pulumi.String("standard"),
//				Description:   pulumi.String("Test For Terraform"),
//				DiskId:        pulumi.String("d-gw8csgxxxxxxxxx"),
//				RetentionDays: pulumi.Int(20),
//				SnapshotName:  pulumi.String("tf-test"),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("Acceptance-test"),
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
// ECS Snapshot can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/ecsSnapshot:EcsSnapshot example <id>
//
// ```
type EcsSnapshot struct {
	pulumi.CustomResourceState

	// The category of the snapshot. Valid Values: `standard` and `flash`.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// The description of the snapshot.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the disk.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// Specifies whether to forcibly delete the snapshot that has been used to create disks.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Specifies whether to enable the instant access feature.
	InstantAccess pulumi.BoolPtrOutput `pulumi:"instantAccess"`
	// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
	InstantAccessRetentionDays pulumi.IntPtrOutput `pulumi:"instantAccessRetentionDays"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource group id.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The retention period of the snapshot.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// The name of the snapshot.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// The status of snapshot.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the snapshot.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEcsSnapshot registers a new resource with the given unique name, arguments, and options.
func NewEcsSnapshot(ctx *pulumi.Context,
	name string, args *EcsSnapshotArgs, opts ...pulumi.ResourceOption) (*EcsSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	var resource EcsSnapshot
	err := ctx.RegisterResource("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsSnapshot gets an existing EcsSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsSnapshotState, opts ...pulumi.ResourceOption) (*EcsSnapshot, error) {
	var resource EcsSnapshot
	err := ctx.ReadResource("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsSnapshot resources.
type ecsSnapshotState struct {
	// The category of the snapshot. Valid Values: `standard` and `flash`.
	Category *string `pulumi:"category"`
	// The description of the snapshot.
	Description *string `pulumi:"description"`
	// The ID of the disk.
	DiskId *string `pulumi:"diskId"`
	// Specifies whether to forcibly delete the snapshot that has been used to create disks.
	Force *bool `pulumi:"force"`
	// Specifies whether to enable the instant access feature.
	InstantAccess *bool `pulumi:"instantAccess"`
	// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
	InstantAccessRetentionDays *int `pulumi:"instantAccessRetentionDays"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name *string `pulumi:"name"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The retention period of the snapshot.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// The status of snapshot.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the snapshot.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EcsSnapshotState struct {
	// The category of the snapshot. Valid Values: `standard` and `flash`.
	Category pulumi.StringPtrInput
	// The description of the snapshot.
	Description pulumi.StringPtrInput
	// The ID of the disk.
	DiskId pulumi.StringPtrInput
	// Specifies whether to forcibly delete the snapshot that has been used to create disks.
	Force pulumi.BoolPtrInput
	// Specifies whether to enable the instant access feature.
	InstantAccess pulumi.BoolPtrInput
	// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
	InstantAccessRetentionDays pulumi.IntPtrInput
	// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringPtrInput
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput
	// The retention period of the snapshot.
	RetentionDays pulumi.IntPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput
	// The status of snapshot.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the snapshot.
	Tags pulumi.MapInput
}

func (EcsSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsSnapshotState)(nil)).Elem()
}

type ecsSnapshotArgs struct {
	// The category of the snapshot. Valid Values: `standard` and `flash`.
	Category *string `pulumi:"category"`
	// The description of the snapshot.
	Description *string `pulumi:"description"`
	// The ID of the disk.
	DiskId string `pulumi:"diskId"`
	// Specifies whether to forcibly delete the snapshot that has been used to create disks.
	Force *bool `pulumi:"force"`
	// Specifies whether to enable the instant access feature.
	InstantAccess *bool `pulumi:"instantAccess"`
	// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
	InstantAccessRetentionDays *int `pulumi:"instantAccessRetentionDays"`
	// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name *string `pulumi:"name"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The retention period of the snapshot.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// A mapping of tags to assign to the snapshot.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EcsSnapshot resource.
type EcsSnapshotArgs struct {
	// The category of the snapshot. Valid Values: `standard` and `flash`.
	Category pulumi.StringPtrInput
	// The description of the snapshot.
	Description pulumi.StringPtrInput
	// The ID of the disk.
	DiskId pulumi.StringInput
	// Specifies whether to forcibly delete the snapshot that has been used to create disks.
	Force pulumi.BoolPtrInput
	// Specifies whether to enable the instant access feature.
	InstantAccess pulumi.BoolPtrInput
	// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
	InstantAccessRetentionDays pulumi.IntPtrInput
	// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringPtrInput
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput
	// The retention period of the snapshot.
	RetentionDays pulumi.IntPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput
	// A mapping of tags to assign to the snapshot.
	Tags pulumi.MapInput
}

func (EcsSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsSnapshotArgs)(nil)).Elem()
}

type EcsSnapshotInput interface {
	pulumi.Input

	ToEcsSnapshotOutput() EcsSnapshotOutput
	ToEcsSnapshotOutputWithContext(ctx context.Context) EcsSnapshotOutput
}

func (*EcsSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsSnapshot)(nil)).Elem()
}

func (i *EcsSnapshot) ToEcsSnapshotOutput() EcsSnapshotOutput {
	return i.ToEcsSnapshotOutputWithContext(context.Background())
}

func (i *EcsSnapshot) ToEcsSnapshotOutputWithContext(ctx context.Context) EcsSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsSnapshotOutput)
}

// EcsSnapshotArrayInput is an input type that accepts EcsSnapshotArray and EcsSnapshotArrayOutput values.
// You can construct a concrete instance of `EcsSnapshotArrayInput` via:
//
//	EcsSnapshotArray{ EcsSnapshotArgs{...} }
type EcsSnapshotArrayInput interface {
	pulumi.Input

	ToEcsSnapshotArrayOutput() EcsSnapshotArrayOutput
	ToEcsSnapshotArrayOutputWithContext(context.Context) EcsSnapshotArrayOutput
}

type EcsSnapshotArray []EcsSnapshotInput

func (EcsSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsSnapshot)(nil)).Elem()
}

func (i EcsSnapshotArray) ToEcsSnapshotArrayOutput() EcsSnapshotArrayOutput {
	return i.ToEcsSnapshotArrayOutputWithContext(context.Background())
}

func (i EcsSnapshotArray) ToEcsSnapshotArrayOutputWithContext(ctx context.Context) EcsSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsSnapshotArrayOutput)
}

// EcsSnapshotMapInput is an input type that accepts EcsSnapshotMap and EcsSnapshotMapOutput values.
// You can construct a concrete instance of `EcsSnapshotMapInput` via:
//
//	EcsSnapshotMap{ "key": EcsSnapshotArgs{...} }
type EcsSnapshotMapInput interface {
	pulumi.Input

	ToEcsSnapshotMapOutput() EcsSnapshotMapOutput
	ToEcsSnapshotMapOutputWithContext(context.Context) EcsSnapshotMapOutput
}

type EcsSnapshotMap map[string]EcsSnapshotInput

func (EcsSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsSnapshot)(nil)).Elem()
}

func (i EcsSnapshotMap) ToEcsSnapshotMapOutput() EcsSnapshotMapOutput {
	return i.ToEcsSnapshotMapOutputWithContext(context.Background())
}

func (i EcsSnapshotMap) ToEcsSnapshotMapOutputWithContext(ctx context.Context) EcsSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsSnapshotMapOutput)
}

type EcsSnapshotOutput struct{ *pulumi.OutputState }

func (EcsSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsSnapshot)(nil)).Elem()
}

func (o EcsSnapshotOutput) ToEcsSnapshotOutput() EcsSnapshotOutput {
	return o
}

func (o EcsSnapshotOutput) ToEcsSnapshotOutputWithContext(ctx context.Context) EcsSnapshotOutput {
	return o
}

// The category of the snapshot. Valid Values: `standard` and `flash`.
func (o EcsSnapshotOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// The description of the snapshot.
func (o EcsSnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the disk.
func (o EcsSnapshotOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

// Specifies whether to forcibly delete the snapshot that has been used to create disks.
func (o EcsSnapshotOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// Specifies whether to enable the instant access feature.
func (o EcsSnapshotOutput) InstantAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.BoolPtrOutput { return v.InstantAccess }).(pulumi.BoolPtrOutput)
}

// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
func (o EcsSnapshotOutput) InstantAccessRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.IntPtrOutput { return v.InstantAccessRetentionDays }).(pulumi.IntPtrOutput)
}

// Field `name` has been deprecated from provider version 1.120.0. New field `snapshotName` instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
func (o EcsSnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource group id.
func (o EcsSnapshotOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The retention period of the snapshot.
func (o EcsSnapshotOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// The name of the snapshot.
func (o EcsSnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// The status of snapshot.
func (o EcsSnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the snapshot.
func (o EcsSnapshotOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsSnapshot) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type EcsSnapshotArrayOutput struct{ *pulumi.OutputState }

func (EcsSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsSnapshot)(nil)).Elem()
}

func (o EcsSnapshotArrayOutput) ToEcsSnapshotArrayOutput() EcsSnapshotArrayOutput {
	return o
}

func (o EcsSnapshotArrayOutput) ToEcsSnapshotArrayOutputWithContext(ctx context.Context) EcsSnapshotArrayOutput {
	return o
}

func (o EcsSnapshotArrayOutput) Index(i pulumi.IntInput) EcsSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsSnapshot {
		return vs[0].([]*EcsSnapshot)[vs[1].(int)]
	}).(EcsSnapshotOutput)
}

type EcsSnapshotMapOutput struct{ *pulumi.OutputState }

func (EcsSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsSnapshot)(nil)).Elem()
}

func (o EcsSnapshotMapOutput) ToEcsSnapshotMapOutput() EcsSnapshotMapOutput {
	return o
}

func (o EcsSnapshotMapOutput) ToEcsSnapshotMapOutputWithContext(ctx context.Context) EcsSnapshotMapOutput {
	return o
}

func (o EcsSnapshotMapOutput) MapIndex(k pulumi.StringInput) EcsSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsSnapshot {
		return vs[0].(map[string]*EcsSnapshot)[vs[1].(string)]
	}).(EcsSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsSnapshotInput)(nil)).Elem(), &EcsSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsSnapshotArrayInput)(nil)).Elem(), EcsSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsSnapshotMapInput)(nil)).Elem(), EcsSnapshotMap{})
	pulumi.RegisterOutputType(EcsSnapshotOutput{})
	pulumi.RegisterOutputType(EcsSnapshotArrayOutput{})
	pulumi.RegisterOutputType(EcsSnapshotMapOutput{})
}
