// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get a list of snapshot according to the specified filters in an Alibaba Cloud account.
// 
// For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
// 
// > **NOTE:**  Available in 1.40.0+.
// 
// ##  Argument Reference
// 
// The following arguments are supported:
// 
// * `instanceId` - (Optional) The specified instance ID.
// * `diskId` - (Optional) The specified disk ID.
// * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values:
//   * true: Encrypted snapshots.
//   * false: No encryption attribute limit.
//   
//   Default value: false.
// * `ids` - (Optional)  A list of snapshot IDs.
// * `nameRegex` - (Optional) A regex string to filter results by snapshot name.
// * `status` - (Optional) The specified snapshot status.
//   * The snapshot status. Optional values:
//   * progressing: The snapshots are being created.
//   * accomplished: The snapshots are ready to use.
//   * failed: The snapshot creation failed.
//   * all: All status.
//   
//   Default value: all.
// 
// * `type` - (Optional) The snapshot category. Optional values:
//   * auto: Auto snapshots.
//   * user: Manual snapshots.
//   * all: Auto and manual snapshots.
//   
//   Default value: all.
// * `sourceDiskType` - (Optional) The type of source disk:
//   * System: The snapshots are created for system disks.
//   * Data: The snapshots are created for data disks.
//   
// * `usage` - (Optional) The usage of the snapshot:
//   * image: The snapshots are used to create custom images.
//   * disk: The snapshots are used to CreateDisk.
//   * mage_disk: The snapshots are used to create custom images and data disks.
//   * none: The snapshots are not used yet.
// * `tags` - (Optional) A map of tags assigned to snapshots.
// * `outputFile` - (Optional) The name of output file that saves the filter results.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/snapshots.html.markdown.
func GetSnapshots(ctx *pulumi.Context, args *GetSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetSnapshotsResult, error) {
	var rv GetSnapshotsResult
	err := ctx.Invoke("alicloud:ecs/getSnapshots:getSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsArgs struct {
	DiskId *string `pulumi:"diskId"`
	Encrypted *bool `pulumi:"encrypted"`
	Ids []string `pulumi:"ids"`
	InstanceId *string `pulumi:"instanceId"`
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	SourceDiskType *string `pulumi:"sourceDiskType"`
	Status *string `pulumi:"status"`
	Tags map[string]interface{} `pulumi:"tags"`
	Type *string `pulumi:"type"`
	Usage *string `pulumi:"usage"`
}


// A collection of values returned by getSnapshots.
type GetSnapshotsResult struct {
	DiskId *string `pulumi:"diskId"`
	// Whether the snapshot is encrypted or not.
	Encrypted *bool `pulumi:"encrypted"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of snapshot IDs.
	Ids []string `pulumi:"ids"`
	InstanceId *string `pulumi:"instanceId"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of snapshots names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of snapshots. Each element contains the following attributes:
	Snapshots []GetSnapshotsSnapshot `pulumi:"snapshots"`
	// Source disk attribute. Value range:
	// * System
	// * Data
	SourceDiskType *string `pulumi:"sourceDiskType"`
	// The snapshot status. Value range:
	// * progressing
	// * accomplished
	// * failed
	Status *string `pulumi:"status"`
	// A map of tags assigned to the snapshot.
	Tags map[string]interface{} `pulumi:"tags"`
	Type *string `pulumi:"type"`
	// Whether the snapshots are used to create resources or not. Value range:
	// * image
	// * disk
	// * imageDisk
	// * none
	Usage *string `pulumi:"usage"`
}

