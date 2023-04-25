// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).
//
// > **NOTE:** Available in v1.89.0+.
//
// > **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eci"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eci.NewImageCache(ctx, "example", &eci.ImageCacheArgs{
//				EipInstanceId:  pulumi.String("eip-uf60c7cqb2pcrkgxhxxxx"),
//				ImageCacheName: pulumi.String("tf-test"),
//				Images: pulumi.StringArray{
//					pulumi.String("registry.cn-beijing.aliyuncs.com/sceneplatform/sae-image-xxxx:latest"),
//				},
//				SecurityGroupId: pulumi.String("sg-2zeef68b66fxxxx"),
//				VswitchId:       pulumi.String("vsw-2zef9k7ng82xxxx"),
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
// ECI Image Cache can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:eci/imageCache:ImageCache example abc123456
//
// ```
type ImageCache struct {
	pulumi.CustomResourceState

	// The ID of the container group job that is used to create the image cache.
	ContainerGroupId pulumi.StringOutput `pulumi:"containerGroupId"`
	// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
	EipInstanceId pulumi.StringPtrOutput `pulumi:"eipInstanceId"`
	// The name of the image cache.
	ImageCacheName pulumi.StringOutput `pulumi:"imageCacheName"`
	// The size of the image cache. Default to `20`. Unit: GiB.
	ImageCacheSize pulumi.IntPtrOutput `pulumi:"imageCacheSize"`
	// The Image Registry parameters about the image to be cached.
	ImageRegistryCredentials ImageCacheImageRegistryCredentialArrayOutput `pulumi:"imageRegistryCredentials"`
	// The images to be cached. The image name must be versioned.
	Images pulumi.StringArrayOutput `pulumi:"images"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// The ID of the security group. You do not need to specify the same security group as the container group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The status of the image cache.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The zone id to cache image.
	ZoneId pulumi.StringPtrOutput `pulumi:"zoneId"`
}

// NewImageCache registers a new resource with the given unique name, arguments, and options.
func NewImageCache(ctx *pulumi.Context,
	name string, args *ImageCacheArgs, opts ...pulumi.ResourceOption) (*ImageCache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageCacheName == nil {
		return nil, errors.New("invalid value for required argument 'ImageCacheName'")
	}
	if args.Images == nil {
		return nil, errors.New("invalid value for required argument 'Images'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource ImageCache
	err := ctx.RegisterResource("alicloud:eci/imageCache:ImageCache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageCache gets an existing ImageCache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageCacheState, opts ...pulumi.ResourceOption) (*ImageCache, error) {
	var resource ImageCache
	err := ctx.ReadResource("alicloud:eci/imageCache:ImageCache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageCache resources.
type imageCacheState struct {
	// The ID of the container group job that is used to create the image cache.
	ContainerGroupId *string `pulumi:"containerGroupId"`
	// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
	EipInstanceId *string `pulumi:"eipInstanceId"`
	// The name of the image cache.
	ImageCacheName *string `pulumi:"imageCacheName"`
	// The size of the image cache. Default to `20`. Unit: GiB.
	ImageCacheSize *int `pulumi:"imageCacheSize"`
	// The Image Registry parameters about the image to be cached.
	ImageRegistryCredentials []ImageCacheImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The images to be cached. The image name must be versioned.
	Images []string `pulumi:"images"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
	RetentionDays *int `pulumi:"retentionDays"`
	// The ID of the security group. You do not need to specify the same security group as the container group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The status of the image cache.
	Status *string `pulumi:"status"`
	// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone id to cache image.
	ZoneId *string `pulumi:"zoneId"`
}

type ImageCacheState struct {
	// The ID of the container group job that is used to create the image cache.
	ContainerGroupId pulumi.StringPtrInput
	// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
	EipInstanceId pulumi.StringPtrInput
	// The name of the image cache.
	ImageCacheName pulumi.StringPtrInput
	// The size of the image cache. Default to `20`. Unit: GiB.
	ImageCacheSize pulumi.IntPtrInput
	// The Image Registry parameters about the image to be cached.
	ImageRegistryCredentials ImageCacheImageRegistryCredentialArrayInput
	// The images to be cached. The image name must be versioned.
	Images pulumi.StringArrayInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
	RetentionDays pulumi.IntPtrInput
	// The ID of the security group. You do not need to specify the same security group as the container group.
	SecurityGroupId pulumi.StringPtrInput
	// The status of the image cache.
	Status pulumi.StringPtrInput
	// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
	VswitchId pulumi.StringPtrInput
	// The zone id to cache image.
	ZoneId pulumi.StringPtrInput
}

func (ImageCacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCacheState)(nil)).Elem()
}

type imageCacheArgs struct {
	// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
	EipInstanceId *string `pulumi:"eipInstanceId"`
	// The name of the image cache.
	ImageCacheName string `pulumi:"imageCacheName"`
	// The size of the image cache. Default to `20`. Unit: GiB.
	ImageCacheSize *int `pulumi:"imageCacheSize"`
	// The Image Registry parameters about the image to be cached.
	ImageRegistryCredentials []ImageCacheImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The images to be cached. The image name must be versioned.
	Images []string `pulumi:"images"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
	RetentionDays *int `pulumi:"retentionDays"`
	// The ID of the security group. You do not need to specify the same security group as the container group.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
	VswitchId string `pulumi:"vswitchId"`
	// The zone id to cache image.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ImageCache resource.
type ImageCacheArgs struct {
	// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
	EipInstanceId pulumi.StringPtrInput
	// The name of the image cache.
	ImageCacheName pulumi.StringInput
	// The size of the image cache. Default to `20`. Unit: GiB.
	ImageCacheSize pulumi.IntPtrInput
	// The Image Registry parameters about the image to be cached.
	ImageRegistryCredentials ImageCacheImageRegistryCredentialArrayInput
	// The images to be cached. The image name must be versioned.
	Images pulumi.StringArrayInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
	RetentionDays pulumi.IntPtrInput
	// The ID of the security group. You do not need to specify the same security group as the container group.
	SecurityGroupId pulumi.StringInput
	// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
	VswitchId pulumi.StringInput
	// The zone id to cache image.
	ZoneId pulumi.StringPtrInput
}

func (ImageCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCacheArgs)(nil)).Elem()
}

type ImageCacheInput interface {
	pulumi.Input

	ToImageCacheOutput() ImageCacheOutput
	ToImageCacheOutputWithContext(ctx context.Context) ImageCacheOutput
}

func (*ImageCache) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageCache)(nil)).Elem()
}

func (i *ImageCache) ToImageCacheOutput() ImageCacheOutput {
	return i.ToImageCacheOutputWithContext(context.Background())
}

func (i *ImageCache) ToImageCacheOutputWithContext(ctx context.Context) ImageCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCacheOutput)
}

// ImageCacheArrayInput is an input type that accepts ImageCacheArray and ImageCacheArrayOutput values.
// You can construct a concrete instance of `ImageCacheArrayInput` via:
//
//	ImageCacheArray{ ImageCacheArgs{...} }
type ImageCacheArrayInput interface {
	pulumi.Input

	ToImageCacheArrayOutput() ImageCacheArrayOutput
	ToImageCacheArrayOutputWithContext(context.Context) ImageCacheArrayOutput
}

type ImageCacheArray []ImageCacheInput

func (ImageCacheArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageCache)(nil)).Elem()
}

func (i ImageCacheArray) ToImageCacheArrayOutput() ImageCacheArrayOutput {
	return i.ToImageCacheArrayOutputWithContext(context.Background())
}

func (i ImageCacheArray) ToImageCacheArrayOutputWithContext(ctx context.Context) ImageCacheArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCacheArrayOutput)
}

// ImageCacheMapInput is an input type that accepts ImageCacheMap and ImageCacheMapOutput values.
// You can construct a concrete instance of `ImageCacheMapInput` via:
//
//	ImageCacheMap{ "key": ImageCacheArgs{...} }
type ImageCacheMapInput interface {
	pulumi.Input

	ToImageCacheMapOutput() ImageCacheMapOutput
	ToImageCacheMapOutputWithContext(context.Context) ImageCacheMapOutput
}

type ImageCacheMap map[string]ImageCacheInput

func (ImageCacheMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageCache)(nil)).Elem()
}

func (i ImageCacheMap) ToImageCacheMapOutput() ImageCacheMapOutput {
	return i.ToImageCacheMapOutputWithContext(context.Background())
}

func (i ImageCacheMap) ToImageCacheMapOutputWithContext(ctx context.Context) ImageCacheMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCacheMapOutput)
}

type ImageCacheOutput struct{ *pulumi.OutputState }

func (ImageCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageCache)(nil)).Elem()
}

func (o ImageCacheOutput) ToImageCacheOutput() ImageCacheOutput {
	return o
}

func (o ImageCacheOutput) ToImageCacheOutputWithContext(ctx context.Context) ImageCacheOutput {
	return o
}

// The ID of the container group job that is used to create the image cache.
func (o ImageCacheOutput) ContainerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringOutput { return v.ContainerGroupId }).(pulumi.StringOutput)
}

// The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
func (o ImageCacheOutput) EipInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringPtrOutput { return v.EipInstanceId }).(pulumi.StringPtrOutput)
}

// The name of the image cache.
func (o ImageCacheOutput) ImageCacheName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringOutput { return v.ImageCacheName }).(pulumi.StringOutput)
}

// The size of the image cache. Default to `20`. Unit: GiB.
func (o ImageCacheOutput) ImageCacheSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.IntPtrOutput { return v.ImageCacheSize }).(pulumi.IntPtrOutput)
}

// The Image Registry parameters about the image to be cached.
func (o ImageCacheOutput) ImageRegistryCredentials() ImageCacheImageRegistryCredentialArrayOutput {
	return o.ApplyT(func(v *ImageCache) ImageCacheImageRegistryCredentialArrayOutput { return v.ImageRegistryCredentials }).(ImageCacheImageRegistryCredentialArrayOutput)
}

// The images to be cached. The image name must be versioned.
func (o ImageCacheOutput) Images() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringArrayOutput { return v.Images }).(pulumi.StringArrayOutput)
}

// The ID of the resource group.
func (o ImageCacheOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
func (o ImageCacheOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// The ID of the security group. You do not need to specify the same security group as the container group.
func (o ImageCacheOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The status of the image cache.
func (o ImageCacheOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
func (o ImageCacheOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The zone id to cache image.
func (o ImageCacheOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageCache) pulumi.StringPtrOutput { return v.ZoneId }).(pulumi.StringPtrOutput)
}

type ImageCacheArrayOutput struct{ *pulumi.OutputState }

func (ImageCacheArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageCache)(nil)).Elem()
}

func (o ImageCacheArrayOutput) ToImageCacheArrayOutput() ImageCacheArrayOutput {
	return o
}

func (o ImageCacheArrayOutput) ToImageCacheArrayOutputWithContext(ctx context.Context) ImageCacheArrayOutput {
	return o
}

func (o ImageCacheArrayOutput) Index(i pulumi.IntInput) ImageCacheOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageCache {
		return vs[0].([]*ImageCache)[vs[1].(int)]
	}).(ImageCacheOutput)
}

type ImageCacheMapOutput struct{ *pulumi.OutputState }

func (ImageCacheMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageCache)(nil)).Elem()
}

func (o ImageCacheMapOutput) ToImageCacheMapOutput() ImageCacheMapOutput {
	return o
}

func (o ImageCacheMapOutput) ToImageCacheMapOutputWithContext(ctx context.Context) ImageCacheMapOutput {
	return o
}

func (o ImageCacheMapOutput) MapIndex(k pulumi.StringInput) ImageCacheOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageCache {
		return vs[0].(map[string]*ImageCache)[vs[1].(string)]
	}).(ImageCacheOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCacheInput)(nil)).Elem(), &ImageCache{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCacheArrayInput)(nil)).Elem(), ImageCacheArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCacheMapInput)(nil)).Elem(), ImageCacheMap{})
	pulumi.RegisterOutputType(ImageCacheOutput{})
	pulumi.RegisterOutputType(ImageCacheArrayOutput{})
	pulumi.RegisterOutputType(ImageCacheMapOutput{})
}
