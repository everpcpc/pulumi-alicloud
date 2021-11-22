// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DFS File System resource.
//
// For information about DFS File System and how to use it, see [What is File System](https://www.alibabacloud.com/help/doc-detail/207144.htm).
//
// > **NOTE:** Available in v1.140.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dfs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "tf-testAccFileSystem"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		defaultZones, err := dfs.GetZones(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dfs.NewFileSystem(ctx, "defaultFileSystem", &dfs.FileSystemArgs{
// 			StorageType:    pulumi.String(defaultZones.Zones[0].Options[0].StorageType),
// 			ZoneId:         pulumi.String(defaultZones.Zones[0].ZoneId),
// 			ProtocolType:   pulumi.String("HDFS"),
// 			Description:    pulumi.String(name),
// 			FileSystemName: pulumi.String(name),
// 			ThroughputMode: pulumi.String("Standard"),
// 			SpaceCapacity:  pulumi.Int(1024),
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
// DFS File System can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dfs/fileSystem:FileSystem example <id>
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// The description of the File system.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the File system.
	FileSystemName pulumi.StringOutput `pulumi:"fileSystemName"`
	// The protocol type. Valid values: `HDFS`.
	ProtocolType pulumi.StringOutput `pulumi:"protocolType"`
	// The preset throughput of the File system. Valid values: `1` to `1024`, Unit: MB/s. **NOTE:** Only when `throughputMode` is `Provisioned`, this param is valid.
	ProvisionedThroughputInMiBps pulumi.IntPtrOutput `pulumi:"provisionedThroughputInMiBps"`
	// The capacity budget of the File system. **NOTE:** When the actual data storage reaches the file system capacity budget, the data cannot be written. The file system capacity budget does not support shrinking.
	SpaceCapacity pulumi.IntOutput `pulumi:"spaceCapacity"`
	// The storage specifications of the File system. Valid values: `PERFORMANCE`, `STANDARD`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// The throughput mode of the File system. Valid values: `Provisioned`, `Standard`.
	ThroughputMode pulumi.StringPtrOutput `pulumi:"throughputMode"`
	// The zone ID of the File system.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemName == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemName'")
	}
	if args.ProtocolType == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolType'")
	}
	if args.SpaceCapacity == nil {
		return nil, errors.New("invalid value for required argument 'SpaceCapacity'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource FileSystem
	err := ctx.RegisterResource("alicloud:dfs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("alicloud:dfs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// The description of the File system.
	Description *string `pulumi:"description"`
	// The name of the File system.
	FileSystemName *string `pulumi:"fileSystemName"`
	// The protocol type. Valid values: `HDFS`.
	ProtocolType *string `pulumi:"protocolType"`
	// The preset throughput of the File system. Valid values: `1` to `1024`, Unit: MB/s. **NOTE:** Only when `throughputMode` is `Provisioned`, this param is valid.
	ProvisionedThroughputInMiBps *int `pulumi:"provisionedThroughputInMiBps"`
	// The capacity budget of the File system. **NOTE:** When the actual data storage reaches the file system capacity budget, the data cannot be written. The file system capacity budget does not support shrinking.
	SpaceCapacity *int `pulumi:"spaceCapacity"`
	// The storage specifications of the File system. Valid values: `PERFORMANCE`, `STANDARD`.
	StorageType *string `pulumi:"storageType"`
	// The throughput mode of the File system. Valid values: `Provisioned`, `Standard`.
	ThroughputMode *string `pulumi:"throughputMode"`
	// The zone ID of the File system.
	ZoneId *string `pulumi:"zoneId"`
}

type FileSystemState struct {
	// The description of the File system.
	Description pulumi.StringPtrInput
	// The name of the File system.
	FileSystemName pulumi.StringPtrInput
	// The protocol type. Valid values: `HDFS`.
	ProtocolType pulumi.StringPtrInput
	// The preset throughput of the File system. Valid values: `1` to `1024`, Unit: MB/s. **NOTE:** Only when `throughputMode` is `Provisioned`, this param is valid.
	ProvisionedThroughputInMiBps pulumi.IntPtrInput
	// The capacity budget of the File system. **NOTE:** When the actual data storage reaches the file system capacity budget, the data cannot be written. The file system capacity budget does not support shrinking.
	SpaceCapacity pulumi.IntPtrInput
	// The storage specifications of the File system. Valid values: `PERFORMANCE`, `STANDARD`.
	StorageType pulumi.StringPtrInput
	// The throughput mode of the File system. Valid values: `Provisioned`, `Standard`.
	ThroughputMode pulumi.StringPtrInput
	// The zone ID of the File system.
	ZoneId pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// The description of the File system.
	Description *string `pulumi:"description"`
	// The name of the File system.
	FileSystemName string `pulumi:"fileSystemName"`
	// The protocol type. Valid values: `HDFS`.
	ProtocolType string `pulumi:"protocolType"`
	// The preset throughput of the File system. Valid values: `1` to `1024`, Unit: MB/s. **NOTE:** Only when `throughputMode` is `Provisioned`, this param is valid.
	ProvisionedThroughputInMiBps *int `pulumi:"provisionedThroughputInMiBps"`
	// The capacity budget of the File system. **NOTE:** When the actual data storage reaches the file system capacity budget, the data cannot be written. The file system capacity budget does not support shrinking.
	SpaceCapacity int `pulumi:"spaceCapacity"`
	// The storage specifications of the File system. Valid values: `PERFORMANCE`, `STANDARD`.
	StorageType string `pulumi:"storageType"`
	// The throughput mode of the File system. Valid values: `Provisioned`, `Standard`.
	ThroughputMode *string `pulumi:"throughputMode"`
	// The zone ID of the File system.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// The description of the File system.
	Description pulumi.StringPtrInput
	// The name of the File system.
	FileSystemName pulumi.StringInput
	// The protocol type. Valid values: `HDFS`.
	ProtocolType pulumi.StringInput
	// The preset throughput of the File system. Valid values: `1` to `1024`, Unit: MB/s. **NOTE:** Only when `throughputMode` is `Provisioned`, this param is valid.
	ProvisionedThroughputInMiBps pulumi.IntPtrInput
	// The capacity budget of the File system. **NOTE:** When the actual data storage reaches the file system capacity budget, the data cannot be written. The file system capacity budget does not support shrinking.
	SpaceCapacity pulumi.IntInput
	// The storage specifications of the File system. Valid values: `PERFORMANCE`, `STANDARD`.
	StorageType pulumi.StringInput
	// The throughput mode of the File system. Valid values: `Provisioned`, `Standard`.
	ThroughputMode pulumi.StringPtrInput
	// The zone ID of the File system.
	ZoneId pulumi.StringInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

func (i *FileSystem) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

type FileSystemPtrInput interface {
	pulumi.Input

	ToFileSystemPtrOutput() FileSystemPtrOutput
	ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput
}

type fileSystemPtrType FileSystemArgs

func (*fileSystemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (i *fileSystemPtrType) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *fileSystemPtrType) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//          FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//          FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct{ *pulumi.OutputState }

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o.ToFileSystemPtrOutputWithContext(context.Background())
}

func (o FileSystemOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystem) *FileSystem {
		return &v
	}).(FileSystemPtrOutput)
}

type FileSystemPtrOutput struct{ *pulumi.OutputState }

func (FileSystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o
}

func (o FileSystemPtrOutput) Elem() FileSystemOutput {
	return o.ApplyT(func(v *FileSystem) FileSystem {
		if v != nil {
			return *v
		}
		var ret FileSystem
		return ret
	}).(FileSystemOutput)
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileSystem)(nil))
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].([]FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FileSystem)(nil))
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].(map[string]FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemPtrInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemArrayInput)(nil)).Elem(), FileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemMapInput)(nil)).Elem(), FileSystemMap{})
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemPtrOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
