// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ECS snapshot resource.
//
// For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewSnapshot(ctx, "snapshot", &ecs.SnapshotArgs{
// 			DiskId:      pulumi.Any(alicloud_disk_attachment.Instance - attachment.Disk_id),
// 			Description: pulumi.String("this snapshot is created for testing"),
// 			Tags: pulumi.StringMap{
// 				"version": pulumi.String("1.2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The source disk ID.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.DiskId == nil {
		return nil, errors.New("missing required argument 'DiskId'")
	}
	if args == nil {
		args = &SnapshotArgs{}
	}
	var resource Snapshot
	err := ctx.RegisterResource("alicloud:ecs/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("alicloud:ecs/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// The source disk ID.
	DiskId *string `pulumi:"diskId"`
	// Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name *string `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type SnapshotState struct {
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// The source disk ID.
	DiskId pulumi.StringPtrInput
	// Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// The source disk ID.
	DiskId string `pulumi:"diskId"`
	// Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name *string `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// The source disk ID.
	DiskId pulumi.StringInput
	// Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	Name pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}
