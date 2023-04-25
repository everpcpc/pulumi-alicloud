// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dfs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DFS Mount Point resource.
//
// For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/doc-detail/207144.htm).
//
// > **NOTE:** Available in v1.140.0+.
//
// ## Import
//
// DFS Mount Point can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dfs/mountPoint:MountPoint example <file_system_id>:<mount_point_id>
//
// ```
type MountPoint struct {
	pulumi.CustomResourceState

	// The ID of the Access Group.
	AccessGroupId pulumi.StringOutput `pulumi:"accessGroupId"`
	// The description of the Mount Point.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the File System.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The ID of the Mount Point.
	MountPointId pulumi.StringOutput `pulumi:"mountPointId"`
	// The network type of the Mount Point. Valid values: `VPC`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// The status of the Mount Point. Valid values: `Active`, `Inactive`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The vpc id.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The vswitch id.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewMountPoint registers a new resource with the given unique name, arguments, and options.
func NewMountPoint(ctx *pulumi.Context,
	name string, args *MountPointArgs, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupId'")
	}
	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource MountPoint
	err := ctx.RegisterResource("alicloud:dfs/mountPoint:MountPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMountPoint gets an existing MountPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMountPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountPointState, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	var resource MountPoint
	err := ctx.ReadResource("alicloud:dfs/mountPoint:MountPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MountPoint resources.
type mountPointState struct {
	// The ID of the Access Group.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// The description of the Mount Point.
	Description *string `pulumi:"description"`
	// The ID of the File System.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The ID of the Mount Point.
	MountPointId *string `pulumi:"mountPointId"`
	// The network type of the Mount Point. Valid values: `VPC`.
	NetworkType *string `pulumi:"networkType"`
	// The status of the Mount Point. Valid values: `Active`, `Inactive`.
	Status *string `pulumi:"status"`
	// The vpc id.
	VpcId *string `pulumi:"vpcId"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
}

type MountPointState struct {
	// The ID of the Access Group.
	AccessGroupId pulumi.StringPtrInput
	// The description of the Mount Point.
	Description pulumi.StringPtrInput
	// The ID of the File System.
	FileSystemId pulumi.StringPtrInput
	// The ID of the Mount Point.
	MountPointId pulumi.StringPtrInput
	// The network type of the Mount Point. Valid values: `VPC`.
	NetworkType pulumi.StringPtrInput
	// The status of the Mount Point. Valid values: `Active`, `Inactive`.
	Status pulumi.StringPtrInput
	// The vpc id.
	VpcId pulumi.StringPtrInput
	// The vswitch id.
	VswitchId pulumi.StringPtrInput
}

func (MountPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointState)(nil)).Elem()
}

type mountPointArgs struct {
	// The ID of the Access Group.
	AccessGroupId string `pulumi:"accessGroupId"`
	// The description of the Mount Point.
	Description *string `pulumi:"description"`
	// The ID of the File System.
	FileSystemId string `pulumi:"fileSystemId"`
	// The network type of the Mount Point. Valid values: `VPC`.
	NetworkType string `pulumi:"networkType"`
	// The status of the Mount Point. Valid values: `Active`, `Inactive`.
	Status *string `pulumi:"status"`
	// The vpc id.
	VpcId string `pulumi:"vpcId"`
	// The vswitch id.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a MountPoint resource.
type MountPointArgs struct {
	// The ID of the Access Group.
	AccessGroupId pulumi.StringInput
	// The description of the Mount Point.
	Description pulumi.StringPtrInput
	// The ID of the File System.
	FileSystemId pulumi.StringInput
	// The network type of the Mount Point. Valid values: `VPC`.
	NetworkType pulumi.StringInput
	// The status of the Mount Point. Valid values: `Active`, `Inactive`.
	Status pulumi.StringPtrInput
	// The vpc id.
	VpcId pulumi.StringInput
	// The vswitch id.
	VswitchId pulumi.StringInput
}

func (MountPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointArgs)(nil)).Elem()
}

type MountPointInput interface {
	pulumi.Input

	ToMountPointOutput() MountPointOutput
	ToMountPointOutputWithContext(ctx context.Context) MountPointOutput
}

func (*MountPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (i *MountPoint) ToMountPointOutput() MountPointOutput {
	return i.ToMountPointOutputWithContext(context.Background())
}

func (i *MountPoint) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointOutput)
}

// MountPointArrayInput is an input type that accepts MountPointArray and MountPointArrayOutput values.
// You can construct a concrete instance of `MountPointArrayInput` via:
//
//	MountPointArray{ MountPointArgs{...} }
type MountPointArrayInput interface {
	pulumi.Input

	ToMountPointArrayOutput() MountPointArrayOutput
	ToMountPointArrayOutputWithContext(context.Context) MountPointArrayOutput
}

type MountPointArray []MountPointInput

func (MountPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (i MountPointArray) ToMountPointArrayOutput() MountPointArrayOutput {
	return i.ToMountPointArrayOutputWithContext(context.Background())
}

func (i MountPointArray) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointArrayOutput)
}

// MountPointMapInput is an input type that accepts MountPointMap and MountPointMapOutput values.
// You can construct a concrete instance of `MountPointMapInput` via:
//
//	MountPointMap{ "key": MountPointArgs{...} }
type MountPointMapInput interface {
	pulumi.Input

	ToMountPointMapOutput() MountPointMapOutput
	ToMountPointMapOutputWithContext(context.Context) MountPointMapOutput
}

type MountPointMap map[string]MountPointInput

func (MountPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (i MountPointMap) ToMountPointMapOutput() MountPointMapOutput {
	return i.ToMountPointMapOutputWithContext(context.Background())
}

func (i MountPointMap) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapOutput)
}

type MountPointOutput struct{ *pulumi.OutputState }

func (MountPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (o MountPointOutput) ToMountPointOutput() MountPointOutput {
	return o
}

func (o MountPointOutput) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return o
}

// The ID of the Access Group.
func (o MountPointOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.AccessGroupId }).(pulumi.StringOutput)
}

// The description of the Mount Point.
func (o MountPointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the File System.
func (o MountPointOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The ID of the Mount Point.
func (o MountPointOutput) MountPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.MountPointId }).(pulumi.StringOutput)
}

// The network type of the Mount Point. Valid values: `VPC`.
func (o MountPointOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// The status of the Mount Point. Valid values: `Active`, `Inactive`.
func (o MountPointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The vpc id.
func (o MountPointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The vswitch id.
func (o MountPointOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type MountPointArrayOutput struct{ *pulumi.OutputState }

func (MountPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (o MountPointArrayOutput) ToMountPointArrayOutput() MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) Index(i pulumi.IntInput) MountPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].([]*MountPoint)[vs[1].(int)]
	}).(MountPointOutput)
}

type MountPointMapOutput struct{ *pulumi.OutputState }

func (MountPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (o MountPointMapOutput) ToMountPointMapOutput() MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) MapIndex(k pulumi.StringInput) MountPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].(map[string]*MountPoint)[vs[1].(string)]
	}).(MountPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointInput)(nil)).Elem(), &MountPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointArrayInput)(nil)).Elem(), MountPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointMapInput)(nil)).Elem(), MountPointMap{})
	pulumi.RegisterOutputType(MountPointOutput{})
	pulumi.RegisterOutputType(MountPointArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapOutput{})
}
