// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Copies a custom image from one region to another. You can use copied images to perform operations in the target region, such as creating instances (RunInstances) and replacing system disks (ReplaceSystemDisk).
//
// > **NOTE:** You can only copy the custom image when it is in the Available state.
//
// > **NOTE:** You can only copy the image belonging to your Alibaba Cloud account. Images cannot be copied from one account to another.
//
// > **NOTE:** If the copying is not completed, you cannot call DeleteImage to delete the image but you can call CancelCopyImage to cancel the copying.
//
// > **NOTE:** Available in 1.66.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewImageCopy(ctx, "_default", &ecs.ImageCopyArgs{
// 			Description:    pulumi.String("test-image"),
// 			ImageName:      pulumi.String("test-image"),
// 			SourceImageId:  pulumi.String("m-bp1gxyhdswlsn18tu***"),
// 			SourceRegionId: pulumi.String("cn-hangzhou"),
// 			Tags: pulumi.AnyMap{
// 				"FinanceDept": pulumi.Any("FinanceDeptJoshua"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attributes Reference0
//
//  The following attributes are exported:
//
// * `id` - ID of the image.
//
// ## Import
//
// image can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ecs/imageCopy:ImageCopy default m-uf66871ape***yg1q***
// ```
type ImageCopy struct {
	pulumi.CustomResourceState

	DeleteAutoSnapshot pulumi.BoolPtrOutput `pulumi:"deleteAutoSnapshot"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The source image ID.
	SourceImageId pulumi.StringOutput `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringOutput `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewImageCopy registers a new resource with the given unique name, arguments, and options.
func NewImageCopy(ctx *pulumi.Context,
	name string, args *ImageCopyArgs, opts ...pulumi.ResourceOption) (*ImageCopy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceImageId == nil {
		return nil, errors.New("invalid value for required argument 'SourceImageId'")
	}
	if args.SourceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'SourceRegionId'")
	}
	var resource ImageCopy
	err := ctx.RegisterResource("alicloud:ecs/imageCopy:ImageCopy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageCopy gets an existing ImageCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageCopy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageCopyState, opts ...pulumi.ResourceOption) (*ImageCopy, error) {
	var resource ImageCopy
	err := ctx.ReadResource("alicloud:ecs/imageCopy:ImageCopy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageCopy resources.
type imageCopyState struct {
	DeleteAutoSnapshot *bool `pulumi:"deleteAutoSnapshot"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// The source image ID.
	SourceImageId *string `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId *string `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ImageCopyState struct {
	DeleteAutoSnapshot pulumi.BoolPtrInput
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// The source image ID.
	SourceImageId pulumi.StringPtrInput
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringPtrInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageCopyState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCopyState)(nil)).Elem()
}

type imageCopyArgs struct {
	DeleteAutoSnapshot *bool `pulumi:"deleteAutoSnapshot"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// The source image ID.
	SourceImageId string `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId string `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ImageCopy resource.
type ImageCopyArgs struct {
	DeleteAutoSnapshot pulumi.BoolPtrInput
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// The source image ID.
	SourceImageId pulumi.StringInput
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCopyArgs)(nil)).Elem()
}

type ImageCopyInput interface {
	pulumi.Input

	ToImageCopyOutput() ImageCopyOutput
	ToImageCopyOutputWithContext(ctx context.Context) ImageCopyOutput
}

func (*ImageCopy) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageCopy)(nil))
}

func (i *ImageCopy) ToImageCopyOutput() ImageCopyOutput {
	return i.ToImageCopyOutputWithContext(context.Background())
}

func (i *ImageCopy) ToImageCopyOutputWithContext(ctx context.Context) ImageCopyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCopyOutput)
}

func (i *ImageCopy) ToImageCopyPtrOutput() ImageCopyPtrOutput {
	return i.ToImageCopyPtrOutputWithContext(context.Background())
}

func (i *ImageCopy) ToImageCopyPtrOutputWithContext(ctx context.Context) ImageCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCopyPtrOutput)
}

type ImageCopyPtrInput interface {
	pulumi.Input

	ToImageCopyPtrOutput() ImageCopyPtrOutput
	ToImageCopyPtrOutputWithContext(ctx context.Context) ImageCopyPtrOutput
}

type imageCopyPtrType ImageCopyArgs

func (*imageCopyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageCopy)(nil))
}

func (i *imageCopyPtrType) ToImageCopyPtrOutput() ImageCopyPtrOutput {
	return i.ToImageCopyPtrOutputWithContext(context.Background())
}

func (i *imageCopyPtrType) ToImageCopyPtrOutputWithContext(ctx context.Context) ImageCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCopyPtrOutput)
}

// ImageCopyArrayInput is an input type that accepts ImageCopyArray and ImageCopyArrayOutput values.
// You can construct a concrete instance of `ImageCopyArrayInput` via:
//
//          ImageCopyArray{ ImageCopyArgs{...} }
type ImageCopyArrayInput interface {
	pulumi.Input

	ToImageCopyArrayOutput() ImageCopyArrayOutput
	ToImageCopyArrayOutputWithContext(context.Context) ImageCopyArrayOutput
}

type ImageCopyArray []ImageCopyInput

func (ImageCopyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageCopy)(nil)).Elem()
}

func (i ImageCopyArray) ToImageCopyArrayOutput() ImageCopyArrayOutput {
	return i.ToImageCopyArrayOutputWithContext(context.Background())
}

func (i ImageCopyArray) ToImageCopyArrayOutputWithContext(ctx context.Context) ImageCopyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCopyArrayOutput)
}

// ImageCopyMapInput is an input type that accepts ImageCopyMap and ImageCopyMapOutput values.
// You can construct a concrete instance of `ImageCopyMapInput` via:
//
//          ImageCopyMap{ "key": ImageCopyArgs{...} }
type ImageCopyMapInput interface {
	pulumi.Input

	ToImageCopyMapOutput() ImageCopyMapOutput
	ToImageCopyMapOutputWithContext(context.Context) ImageCopyMapOutput
}

type ImageCopyMap map[string]ImageCopyInput

func (ImageCopyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageCopy)(nil)).Elem()
}

func (i ImageCopyMap) ToImageCopyMapOutput() ImageCopyMapOutput {
	return i.ToImageCopyMapOutputWithContext(context.Background())
}

func (i ImageCopyMap) ToImageCopyMapOutputWithContext(ctx context.Context) ImageCopyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageCopyMapOutput)
}

type ImageCopyOutput struct{ *pulumi.OutputState }

func (ImageCopyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageCopy)(nil))
}

func (o ImageCopyOutput) ToImageCopyOutput() ImageCopyOutput {
	return o
}

func (o ImageCopyOutput) ToImageCopyOutputWithContext(ctx context.Context) ImageCopyOutput {
	return o
}

func (o ImageCopyOutput) ToImageCopyPtrOutput() ImageCopyPtrOutput {
	return o.ToImageCopyPtrOutputWithContext(context.Background())
}

func (o ImageCopyOutput) ToImageCopyPtrOutputWithContext(ctx context.Context) ImageCopyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageCopy) *ImageCopy {
		return &v
	}).(ImageCopyPtrOutput)
}

type ImageCopyPtrOutput struct{ *pulumi.OutputState }

func (ImageCopyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageCopy)(nil))
}

func (o ImageCopyPtrOutput) ToImageCopyPtrOutput() ImageCopyPtrOutput {
	return o
}

func (o ImageCopyPtrOutput) ToImageCopyPtrOutputWithContext(ctx context.Context) ImageCopyPtrOutput {
	return o
}

func (o ImageCopyPtrOutput) Elem() ImageCopyOutput {
	return o.ApplyT(func(v *ImageCopy) ImageCopy {
		if v != nil {
			return *v
		}
		var ret ImageCopy
		return ret
	}).(ImageCopyOutput)
}

type ImageCopyArrayOutput struct{ *pulumi.OutputState }

func (ImageCopyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageCopy)(nil))
}

func (o ImageCopyArrayOutput) ToImageCopyArrayOutput() ImageCopyArrayOutput {
	return o
}

func (o ImageCopyArrayOutput) ToImageCopyArrayOutputWithContext(ctx context.Context) ImageCopyArrayOutput {
	return o
}

func (o ImageCopyArrayOutput) Index(i pulumi.IntInput) ImageCopyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageCopy {
		return vs[0].([]ImageCopy)[vs[1].(int)]
	}).(ImageCopyOutput)
}

type ImageCopyMapOutput struct{ *pulumi.OutputState }

func (ImageCopyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageCopy)(nil))
}

func (o ImageCopyMapOutput) ToImageCopyMapOutput() ImageCopyMapOutput {
	return o
}

func (o ImageCopyMapOutput) ToImageCopyMapOutputWithContext(ctx context.Context) ImageCopyMapOutput {
	return o
}

func (o ImageCopyMapOutput) MapIndex(k pulumi.StringInput) ImageCopyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageCopy {
		return vs[0].(map[string]ImageCopy)[vs[1].(string)]
	}).(ImageCopyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCopyInput)(nil)).Elem(), &ImageCopy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCopyPtrInput)(nil)).Elem(), &ImageCopy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCopyArrayInput)(nil)).Elem(), ImageCopyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageCopyMapInput)(nil)).Elem(), ImageCopyMap{})
	pulumi.RegisterOutputType(ImageCopyOutput{})
	pulumi.RegisterOutputType(ImageCopyPtrOutput{})
	pulumi.RegisterOutputType(ImageCopyArrayOutput{})
	pulumi.RegisterOutputType(ImageCopyMapOutput{})
}
