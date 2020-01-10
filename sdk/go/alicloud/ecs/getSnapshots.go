// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
func LookupSnapshots(ctx *pulumi.Context, args *GetSnapshotsArgs) (*GetSnapshotsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["diskId"] = args.DiskId
		inputs["encrypted"] = args.Encrypted
		inputs["ids"] = args.Ids
		inputs["instanceId"] = args.InstanceId
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["sourceDiskType"] = args.SourceDiskType
		inputs["status"] = args.Status
		inputs["tags"] = args.Tags
		inputs["type"] = args.Type
		inputs["usage"] = args.Usage
	}
	outputs, err := ctx.Invoke("alicloud:ecs/getSnapshots:getSnapshots", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSnapshotsResult{
		DiskId: outputs["diskId"],
		Encrypted: outputs["encrypted"],
		Ids: outputs["ids"],
		InstanceId: outputs["instanceId"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Snapshots: outputs["snapshots"],
		SourceDiskType: outputs["sourceDiskType"],
		Status: outputs["status"],
		Tags: outputs["tags"],
		Type: outputs["type"],
		Usage: outputs["usage"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsArgs struct {
	DiskId interface{}
	Encrypted interface{}
	Ids interface{}
	InstanceId interface{}
	NameRegex interface{}
	OutputFile interface{}
	SourceDiskType interface{}
	Status interface{}
	Tags interface{}
	Type interface{}
	Usage interface{}
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResult struct {
	DiskId interface{}
	// Whether the snapshot is encrypted or not.
	Encrypted interface{}
	// A list of snapshot IDs.
	Ids interface{}
	InstanceId interface{}
	NameRegex interface{}
	// A list of snapshots names.
	Names interface{}
	OutputFile interface{}
	// A list of snapshots. Each element contains the following attributes:
	Snapshots interface{}
	// Source disk attribute. Value range:
	// * System
	// * Data
	SourceDiskType interface{}
	// The snapshot status. Value range:
	// * progressing
	// * accomplished
	// * failed
	Status interface{}
	// A map of tags assigned to the snapshot.
	Tags interface{}
	Type interface{}
	// Whether the snapshots are used to create resources or not. Value range:
	// * image
	// * disk
	// * imageDisk
	// * none
	Usage interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
