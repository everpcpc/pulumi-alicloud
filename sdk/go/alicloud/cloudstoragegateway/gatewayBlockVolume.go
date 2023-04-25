// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Storage Gateway Gateway Block Volume resource.
//
// For information about Cloud Storage Gateway Gateway Block Volume and how to use it, see [What is Gateway Block Volume](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
//
// > **NOTE:** Available in v1.144.0+.
//
// ## Import
//
// Cloud Storage Gateway Gateway Block Volume can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume example <gateway_id>:<index_id>
//
// ```
type GatewayBlockVolume struct {
	pulumi.CustomResourceState

	// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
	CacheMode pulumi.StringOutput `pulumi:"cacheMode"`
	// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
	ChapEnabled pulumi.BoolOutput `pulumi:"chapEnabled"`
	// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInPassword pulumi.StringPtrOutput `pulumi:"chapInPassword"`
	// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInUser pulumi.StringPtrOutput `pulumi:"chapInUser"`
	// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
	ChunkSize pulumi.IntOutput `pulumi:"chunkSize"`
	// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
	GatewayBlockVolumeName pulumi.StringOutput `pulumi:"gatewayBlockVolumeName"`
	// The Gateway ID.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The ID of the index.
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
	IsSourceDeletion pulumi.BoolPtrOutput `pulumi:"isSourceDeletion"`
	// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
	LocalPath pulumi.StringPtrOutput `pulumi:"localPath"`
	// The name of the OSS Bucket.
	OssBucketName pulumi.StringOutput `pulumi:"ossBucketName"`
	// Whether to enable SSL access your OSS Buckets. Default value: `true`.
	OssBucketSsl pulumi.BoolOutput `pulumi:"ossBucketSsl"`
	// The endpoint of the OSS Bucket.
	OssEndpoint pulumi.StringOutput `pulumi:"ossEndpoint"`
	// The Protocol. Valid values: `iSCSI`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The recovery.
	Recovery pulumi.BoolPtrOutput `pulumi:"recovery"`
	// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
	Size pulumi.IntOutput `pulumi:"size"`
	// The status of volume. Valid values:
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewGatewayBlockVolume registers a new resource with the given unique name, arguments, and options.
func NewGatewayBlockVolume(ctx *pulumi.Context,
	name string, args *GatewayBlockVolumeArgs, opts ...pulumi.ResourceOption) (*GatewayBlockVolume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayBlockVolumeName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayBlockVolumeName'")
	}
	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.OssBucketName == nil {
		return nil, errors.New("invalid value for required argument 'OssBucketName'")
	}
	if args.OssEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'OssEndpoint'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	var resource GatewayBlockVolume
	err := ctx.RegisterResource("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayBlockVolume gets an existing GatewayBlockVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayBlockVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayBlockVolumeState, opts ...pulumi.ResourceOption) (*GatewayBlockVolume, error) {
	var resource GatewayBlockVolume
	err := ctx.ReadResource("alicloud:cloudstoragegateway/gatewayBlockVolume:GatewayBlockVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayBlockVolume resources.
type gatewayBlockVolumeState struct {
	// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
	CacheMode *string `pulumi:"cacheMode"`
	// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
	ChapEnabled *bool `pulumi:"chapEnabled"`
	// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInPassword *string `pulumi:"chapInPassword"`
	// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInUser *string `pulumi:"chapInUser"`
	// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
	ChunkSize *int `pulumi:"chunkSize"`
	// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
	GatewayBlockVolumeName *string `pulumi:"gatewayBlockVolumeName"`
	// The Gateway ID.
	GatewayId *string `pulumi:"gatewayId"`
	// The ID of the index.
	IndexId *string `pulumi:"indexId"`
	// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
	IsSourceDeletion *bool `pulumi:"isSourceDeletion"`
	// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
	LocalPath *string `pulumi:"localPath"`
	// The name of the OSS Bucket.
	OssBucketName *string `pulumi:"ossBucketName"`
	// Whether to enable SSL access your OSS Buckets. Default value: `true`.
	OssBucketSsl *bool `pulumi:"ossBucketSsl"`
	// The endpoint of the OSS Bucket.
	OssEndpoint *string `pulumi:"ossEndpoint"`
	// The Protocol. Valid values: `iSCSI`.
	Protocol *string `pulumi:"protocol"`
	// The recovery.
	Recovery *bool `pulumi:"recovery"`
	// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
	Size *int `pulumi:"size"`
	// The status of volume. Valid values:
	Status *string `pulumi:"status"`
}

type GatewayBlockVolumeState struct {
	// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
	CacheMode pulumi.StringPtrInput
	// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
	ChapEnabled pulumi.BoolPtrInput
	// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInPassword pulumi.StringPtrInput
	// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInUser pulumi.StringPtrInput
	// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
	ChunkSize pulumi.IntPtrInput
	// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
	GatewayBlockVolumeName pulumi.StringPtrInput
	// The Gateway ID.
	GatewayId pulumi.StringPtrInput
	// The ID of the index.
	IndexId pulumi.StringPtrInput
	// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
	IsSourceDeletion pulumi.BoolPtrInput
	// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
	LocalPath pulumi.StringPtrInput
	// The name of the OSS Bucket.
	OssBucketName pulumi.StringPtrInput
	// Whether to enable SSL access your OSS Buckets. Default value: `true`.
	OssBucketSsl pulumi.BoolPtrInput
	// The endpoint of the OSS Bucket.
	OssEndpoint pulumi.StringPtrInput
	// The Protocol. Valid values: `iSCSI`.
	Protocol pulumi.StringPtrInput
	// The recovery.
	Recovery pulumi.BoolPtrInput
	// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
	Size pulumi.IntPtrInput
	// The status of volume. Valid values:
	Status pulumi.StringPtrInput
}

func (GatewayBlockVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayBlockVolumeState)(nil)).Elem()
}

type gatewayBlockVolumeArgs struct {
	// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
	CacheMode *string `pulumi:"cacheMode"`
	// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
	ChapEnabled *bool `pulumi:"chapEnabled"`
	// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInPassword *string `pulumi:"chapInPassword"`
	// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInUser *string `pulumi:"chapInUser"`
	// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
	ChunkSize *int `pulumi:"chunkSize"`
	// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
	GatewayBlockVolumeName string `pulumi:"gatewayBlockVolumeName"`
	// The Gateway ID.
	GatewayId string `pulumi:"gatewayId"`
	// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
	IsSourceDeletion *bool `pulumi:"isSourceDeletion"`
	// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
	LocalPath *string `pulumi:"localPath"`
	// The name of the OSS Bucket.
	OssBucketName string `pulumi:"ossBucketName"`
	// Whether to enable SSL access your OSS Buckets. Default value: `true`.
	OssBucketSsl *bool `pulumi:"ossBucketSsl"`
	// The endpoint of the OSS Bucket.
	OssEndpoint string `pulumi:"ossEndpoint"`
	// The Protocol. Valid values: `iSCSI`.
	Protocol string `pulumi:"protocol"`
	// The recovery.
	Recovery *bool `pulumi:"recovery"`
	// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
	Size *int `pulumi:"size"`
}

// The set of arguments for constructing a GatewayBlockVolume resource.
type GatewayBlockVolumeArgs struct {
	// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
	CacheMode pulumi.StringPtrInput
	// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
	ChapEnabled pulumi.BoolPtrInput
	// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInPassword pulumi.StringPtrInput
	// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
	ChapInUser pulumi.StringPtrInput
	// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
	ChunkSize pulumi.IntPtrInput
	// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
	GatewayBlockVolumeName pulumi.StringInput
	// The Gateway ID.
	GatewayId pulumi.StringInput
	// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
	IsSourceDeletion pulumi.BoolPtrInput
	// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
	LocalPath pulumi.StringPtrInput
	// The name of the OSS Bucket.
	OssBucketName pulumi.StringInput
	// Whether to enable SSL access your OSS Buckets. Default value: `true`.
	OssBucketSsl pulumi.BoolPtrInput
	// The endpoint of the OSS Bucket.
	OssEndpoint pulumi.StringInput
	// The Protocol. Valid values: `iSCSI`.
	Protocol pulumi.StringInput
	// The recovery.
	Recovery pulumi.BoolPtrInput
	// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
	Size pulumi.IntPtrInput
}

func (GatewayBlockVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayBlockVolumeArgs)(nil)).Elem()
}

type GatewayBlockVolumeInput interface {
	pulumi.Input

	ToGatewayBlockVolumeOutput() GatewayBlockVolumeOutput
	ToGatewayBlockVolumeOutputWithContext(ctx context.Context) GatewayBlockVolumeOutput
}

func (*GatewayBlockVolume) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayBlockVolume)(nil)).Elem()
}

func (i *GatewayBlockVolume) ToGatewayBlockVolumeOutput() GatewayBlockVolumeOutput {
	return i.ToGatewayBlockVolumeOutputWithContext(context.Background())
}

func (i *GatewayBlockVolume) ToGatewayBlockVolumeOutputWithContext(ctx context.Context) GatewayBlockVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayBlockVolumeOutput)
}

// GatewayBlockVolumeArrayInput is an input type that accepts GatewayBlockVolumeArray and GatewayBlockVolumeArrayOutput values.
// You can construct a concrete instance of `GatewayBlockVolumeArrayInput` via:
//
//	GatewayBlockVolumeArray{ GatewayBlockVolumeArgs{...} }
type GatewayBlockVolumeArrayInput interface {
	pulumi.Input

	ToGatewayBlockVolumeArrayOutput() GatewayBlockVolumeArrayOutput
	ToGatewayBlockVolumeArrayOutputWithContext(context.Context) GatewayBlockVolumeArrayOutput
}

type GatewayBlockVolumeArray []GatewayBlockVolumeInput

func (GatewayBlockVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayBlockVolume)(nil)).Elem()
}

func (i GatewayBlockVolumeArray) ToGatewayBlockVolumeArrayOutput() GatewayBlockVolumeArrayOutput {
	return i.ToGatewayBlockVolumeArrayOutputWithContext(context.Background())
}

func (i GatewayBlockVolumeArray) ToGatewayBlockVolumeArrayOutputWithContext(ctx context.Context) GatewayBlockVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayBlockVolumeArrayOutput)
}

// GatewayBlockVolumeMapInput is an input type that accepts GatewayBlockVolumeMap and GatewayBlockVolumeMapOutput values.
// You can construct a concrete instance of `GatewayBlockVolumeMapInput` via:
//
//	GatewayBlockVolumeMap{ "key": GatewayBlockVolumeArgs{...} }
type GatewayBlockVolumeMapInput interface {
	pulumi.Input

	ToGatewayBlockVolumeMapOutput() GatewayBlockVolumeMapOutput
	ToGatewayBlockVolumeMapOutputWithContext(context.Context) GatewayBlockVolumeMapOutput
}

type GatewayBlockVolumeMap map[string]GatewayBlockVolumeInput

func (GatewayBlockVolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayBlockVolume)(nil)).Elem()
}

func (i GatewayBlockVolumeMap) ToGatewayBlockVolumeMapOutput() GatewayBlockVolumeMapOutput {
	return i.ToGatewayBlockVolumeMapOutputWithContext(context.Background())
}

func (i GatewayBlockVolumeMap) ToGatewayBlockVolumeMapOutputWithContext(ctx context.Context) GatewayBlockVolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayBlockVolumeMapOutput)
}

type GatewayBlockVolumeOutput struct{ *pulumi.OutputState }

func (GatewayBlockVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayBlockVolume)(nil)).Elem()
}

func (o GatewayBlockVolumeOutput) ToGatewayBlockVolumeOutput() GatewayBlockVolumeOutput {
	return o
}

func (o GatewayBlockVolumeOutput) ToGatewayBlockVolumeOutputWithContext(ctx context.Context) GatewayBlockVolumeOutput {
	return o
}

// The Block volume set mode to cache mode. Valid values: `Cache`, `WriteThrough`.
func (o GatewayBlockVolumeOutput) CacheMode() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.CacheMode }).(pulumi.StringOutput)
}

// Whether to enable iSCSI access of CHAP authentication, which currently supports both CHAP inbound authentication.  Default value: `false`.
func (o GatewayBlockVolumeOutput) ChapEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.BoolOutput { return v.ChapEnabled }).(pulumi.BoolOutput)
}

// The password for inbound authentication when the block volume enables iSCSI access to CHAP authentication. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
func (o GatewayBlockVolumeOutput) ChapInPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringPtrOutput { return v.ChapInPassword }).(pulumi.StringPtrOutput)
}

// The Inbound CHAP user. The `chapInUser` must be 1 to 32 characters in length, and can contain letters and digits. **NOTE:** When the `chapEnabled` is  `true` is,The `chapInPassword` is valid.
func (o GatewayBlockVolumeOutput) ChapInUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringPtrOutput { return v.ChapInUser }).(pulumi.StringPtrOutput)
}

// The Block volume storage allocation unit.  Valid values: `8192`, `16384`, `32768`, `65536`, `131072`. Default value: `32768`. Unit: `Byte`.
func (o GatewayBlockVolumeOutput) ChunkSize() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.IntOutput { return v.ChunkSize }).(pulumi.IntOutput)
}

// The Block volume name. The name must be 1 to 32 characters in length, and can contain lower case letters and digits.
func (o GatewayBlockVolumeOutput) GatewayBlockVolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.GatewayBlockVolumeName }).(pulumi.StringOutput)
}

// The Gateway ID.
func (o GatewayBlockVolumeOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// The ID of the index.
func (o GatewayBlockVolumeOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.IndexId }).(pulumi.StringOutput)
}

// Whether to delete the source data. Default value `true`. **NOTE:** When `isSourceDeletion` is `true`, the data in the OSS Bucket on the cloud is also deleted when deleting the block gateway volume. Please operate with caution.
func (o GatewayBlockVolumeOutput) IsSourceDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.BoolPtrOutput { return v.IsSourceDeletion }).(pulumi.BoolPtrOutput)
}

// The Cache disk to local path. **NOTE:**  When the `cacheMode` is  `Cache` is,The `chapInPassword` is valid.
func (o GatewayBlockVolumeOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringPtrOutput { return v.LocalPath }).(pulumi.StringPtrOutput)
}

// The name of the OSS Bucket.
func (o GatewayBlockVolumeOutput) OssBucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.OssBucketName }).(pulumi.StringOutput)
}

// Whether to enable SSL access your OSS Buckets. Default value: `true`.
func (o GatewayBlockVolumeOutput) OssBucketSsl() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.BoolOutput { return v.OssBucketSsl }).(pulumi.BoolOutput)
}

// The endpoint of the OSS Bucket.
func (o GatewayBlockVolumeOutput) OssEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.OssEndpoint }).(pulumi.StringOutput)
}

// The Protocol. Valid values: `iSCSI`.
func (o GatewayBlockVolumeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The recovery.
func (o GatewayBlockVolumeOutput) Recovery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.BoolPtrOutput { return v.Recovery }).(pulumi.BoolPtrOutput)
}

// The Volume size. Valid values: `1` to `262144`. Unit: `Byte`.
func (o GatewayBlockVolumeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The status of volume. Valid values:
func (o GatewayBlockVolumeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayBlockVolume) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type GatewayBlockVolumeArrayOutput struct{ *pulumi.OutputState }

func (GatewayBlockVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayBlockVolume)(nil)).Elem()
}

func (o GatewayBlockVolumeArrayOutput) ToGatewayBlockVolumeArrayOutput() GatewayBlockVolumeArrayOutput {
	return o
}

func (o GatewayBlockVolumeArrayOutput) ToGatewayBlockVolumeArrayOutputWithContext(ctx context.Context) GatewayBlockVolumeArrayOutput {
	return o
}

func (o GatewayBlockVolumeArrayOutput) Index(i pulumi.IntInput) GatewayBlockVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayBlockVolume {
		return vs[0].([]*GatewayBlockVolume)[vs[1].(int)]
	}).(GatewayBlockVolumeOutput)
}

type GatewayBlockVolumeMapOutput struct{ *pulumi.OutputState }

func (GatewayBlockVolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayBlockVolume)(nil)).Elem()
}

func (o GatewayBlockVolumeMapOutput) ToGatewayBlockVolumeMapOutput() GatewayBlockVolumeMapOutput {
	return o
}

func (o GatewayBlockVolumeMapOutput) ToGatewayBlockVolumeMapOutputWithContext(ctx context.Context) GatewayBlockVolumeMapOutput {
	return o
}

func (o GatewayBlockVolumeMapOutput) MapIndex(k pulumi.StringInput) GatewayBlockVolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayBlockVolume {
		return vs[0].(map[string]*GatewayBlockVolume)[vs[1].(string)]
	}).(GatewayBlockVolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayBlockVolumeInput)(nil)).Elem(), &GatewayBlockVolume{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayBlockVolumeArrayInput)(nil)).Elem(), GatewayBlockVolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayBlockVolumeMapInput)(nil)).Elem(), GatewayBlockVolumeMap{})
	pulumi.RegisterOutputType(GatewayBlockVolumeOutput{})
	pulumi.RegisterOutputType(GatewayBlockVolumeArrayOutput{})
	pulumi.RegisterOutputType(GatewayBlockVolumeMapOutput{})
}
