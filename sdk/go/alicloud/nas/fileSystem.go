// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Nas File System resource.
// 
// After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
// 
// For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
// 
// > **NOTE:** Available in v1.33.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/nas_file_system.html.markdown.
type FileSystem struct {
	s *pulumi.ResourceState
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOpt) (*FileSystem, error) {
	if args == nil || args.ProtocolType == nil {
		return nil, errors.New("missing required argument 'ProtocolType'")
	}
	if args == nil || args.StorageType == nil {
		return nil, errors.New("missing required argument 'StorageType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["protocolType"] = nil
		inputs["storageType"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["protocolType"] = args.ProtocolType
		inputs["storageType"] = args.StorageType
	}
	s, err := ctx.RegisterResource("alicloud:nas/fileSystem:FileSystem", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FileSystem{s: s}, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FileSystemState, opts ...pulumi.ResourceOpt) (*FileSystem, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["protocolType"] = state.ProtocolType
		inputs["storageType"] = state.StorageType
	}
	s, err := ctx.ReadResource("alicloud:nas/fileSystem:FileSystem", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FileSystem{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FileSystem) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FileSystem) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The File System description.
func (r *FileSystem) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
func (r *FileSystem) ProtocolType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocolType"])
}

// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
func (r *FileSystem) StorageType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageType"])
}

// Input properties used for looking up and filtering FileSystem resources.
type FileSystemState struct {
	// The File System description.
	Description interface{}
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType interface{}
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType interface{}
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// The File System description.
	Description interface{}
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType interface{}
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType interface{}
}
