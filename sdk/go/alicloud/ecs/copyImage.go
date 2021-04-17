// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CopyImage struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Encrypted   pulumi.BoolPtrOutput   `pulumi:"encrypted"`
	Force       pulumi.BoolPtrOutput   `pulumi:"force"`
	ImageName   pulumi.StringOutput    `pulumi:"imageName"`
	KmsKeyId    pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name           pulumi.StringOutput `pulumi:"name"`
	SourceImageId  pulumi.StringOutput `pulumi:"sourceImageId"`
	SourceRegionId pulumi.StringOutput `pulumi:"sourceRegionId"`
	Tags           pulumi.MapOutput    `pulumi:"tags"`
}

// NewCopyImage registers a new resource with the given unique name, arguments, and options.
func NewCopyImage(ctx *pulumi.Context,
	name string, args *CopyImageArgs, opts ...pulumi.ResourceOption) (*CopyImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceImageId == nil {
		return nil, errors.New("invalid value for required argument 'SourceImageId'")
	}
	if args.SourceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'SourceRegionId'")
	}
	var resource CopyImage
	err := ctx.RegisterResource("alicloud:ecs/copyImage:CopyImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCopyImage gets an existing CopyImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCopyImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CopyImageState, opts ...pulumi.ResourceOption) (*CopyImage, error) {
	var resource CopyImage
	err := ctx.ReadResource("alicloud:ecs/copyImage:CopyImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CopyImage resources.
type copyImageState struct {
	Description *string `pulumi:"description"`
	Encrypted   *bool   `pulumi:"encrypted"`
	Force       *bool   `pulumi:"force"`
	ImageName   *string `pulumi:"imageName"`
	KmsKeyId    *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name           *string                `pulumi:"name"`
	SourceImageId  *string                `pulumi:"sourceImageId"`
	SourceRegionId *string                `pulumi:"sourceRegionId"`
	Tags           map[string]interface{} `pulumi:"tags"`
}

type CopyImageState struct {
	Description pulumi.StringPtrInput
	Encrypted   pulumi.BoolPtrInput
	Force       pulumi.BoolPtrInput
	ImageName   pulumi.StringPtrInput
	KmsKeyId    pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name           pulumi.StringPtrInput
	SourceImageId  pulumi.StringPtrInput
	SourceRegionId pulumi.StringPtrInput
	Tags           pulumi.MapInput
}

func (CopyImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*copyImageState)(nil)).Elem()
}

type copyImageArgs struct {
	Description *string `pulumi:"description"`
	Encrypted   *bool   `pulumi:"encrypted"`
	Force       *bool   `pulumi:"force"`
	ImageName   *string `pulumi:"imageName"`
	KmsKeyId    *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name           *string                `pulumi:"name"`
	SourceImageId  string                 `pulumi:"sourceImageId"`
	SourceRegionId string                 `pulumi:"sourceRegionId"`
	Tags           map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a CopyImage resource.
type CopyImageArgs struct {
	Description pulumi.StringPtrInput
	Encrypted   pulumi.BoolPtrInput
	Force       pulumi.BoolPtrInput
	ImageName   pulumi.StringPtrInput
	KmsKeyId    pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name           pulumi.StringPtrInput
	SourceImageId  pulumi.StringInput
	SourceRegionId pulumi.StringInput
	Tags           pulumi.MapInput
}

func (CopyImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*copyImageArgs)(nil)).Elem()
}

type CopyImageInput interface {
	pulumi.Input

	ToCopyImageOutput() CopyImageOutput
	ToCopyImageOutputWithContext(ctx context.Context) CopyImageOutput
}

func (*CopyImage) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyImage)(nil))
}

func (i *CopyImage) ToCopyImageOutput() CopyImageOutput {
	return i.ToCopyImageOutputWithContext(context.Background())
}

func (i *CopyImage) ToCopyImageOutputWithContext(ctx context.Context) CopyImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyImageOutput)
}

func (i *CopyImage) ToCopyImagePtrOutput() CopyImagePtrOutput {
	return i.ToCopyImagePtrOutputWithContext(context.Background())
}

func (i *CopyImage) ToCopyImagePtrOutputWithContext(ctx context.Context) CopyImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyImagePtrOutput)
}

type CopyImagePtrInput interface {
	pulumi.Input

	ToCopyImagePtrOutput() CopyImagePtrOutput
	ToCopyImagePtrOutputWithContext(ctx context.Context) CopyImagePtrOutput
}

type copyImagePtrType CopyImageArgs

func (*copyImagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CopyImage)(nil))
}

func (i *copyImagePtrType) ToCopyImagePtrOutput() CopyImagePtrOutput {
	return i.ToCopyImagePtrOutputWithContext(context.Background())
}

func (i *copyImagePtrType) ToCopyImagePtrOutputWithContext(ctx context.Context) CopyImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyImagePtrOutput)
}

// CopyImageArrayInput is an input type that accepts CopyImageArray and CopyImageArrayOutput values.
// You can construct a concrete instance of `CopyImageArrayInput` via:
//
//          CopyImageArray{ CopyImageArgs{...} }
type CopyImageArrayInput interface {
	pulumi.Input

	ToCopyImageArrayOutput() CopyImageArrayOutput
	ToCopyImageArrayOutputWithContext(context.Context) CopyImageArrayOutput
}

type CopyImageArray []CopyImageInput

func (CopyImageArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CopyImage)(nil))
}

func (i CopyImageArray) ToCopyImageArrayOutput() CopyImageArrayOutput {
	return i.ToCopyImageArrayOutputWithContext(context.Background())
}

func (i CopyImageArray) ToCopyImageArrayOutputWithContext(ctx context.Context) CopyImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyImageArrayOutput)
}

// CopyImageMapInput is an input type that accepts CopyImageMap and CopyImageMapOutput values.
// You can construct a concrete instance of `CopyImageMapInput` via:
//
//          CopyImageMap{ "key": CopyImageArgs{...} }
type CopyImageMapInput interface {
	pulumi.Input

	ToCopyImageMapOutput() CopyImageMapOutput
	ToCopyImageMapOutputWithContext(context.Context) CopyImageMapOutput
}

type CopyImageMap map[string]CopyImageInput

func (CopyImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CopyImage)(nil))
}

func (i CopyImageMap) ToCopyImageMapOutput() CopyImageMapOutput {
	return i.ToCopyImageMapOutputWithContext(context.Background())
}

func (i CopyImageMap) ToCopyImageMapOutputWithContext(ctx context.Context) CopyImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyImageMapOutput)
}

type CopyImageOutput struct {
	*pulumi.OutputState
}

func (CopyImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyImage)(nil))
}

func (o CopyImageOutput) ToCopyImageOutput() CopyImageOutput {
	return o
}

func (o CopyImageOutput) ToCopyImageOutputWithContext(ctx context.Context) CopyImageOutput {
	return o
}

func (o CopyImageOutput) ToCopyImagePtrOutput() CopyImagePtrOutput {
	return o.ToCopyImagePtrOutputWithContext(context.Background())
}

func (o CopyImageOutput) ToCopyImagePtrOutputWithContext(ctx context.Context) CopyImagePtrOutput {
	return o.ApplyT(func(v CopyImage) *CopyImage {
		return &v
	}).(CopyImagePtrOutput)
}

type CopyImagePtrOutput struct {
	*pulumi.OutputState
}

func (CopyImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CopyImage)(nil))
}

func (o CopyImagePtrOutput) ToCopyImagePtrOutput() CopyImagePtrOutput {
	return o
}

func (o CopyImagePtrOutput) ToCopyImagePtrOutputWithContext(ctx context.Context) CopyImagePtrOutput {
	return o
}

type CopyImageArrayOutput struct{ *pulumi.OutputState }

func (CopyImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyImage)(nil))
}

func (o CopyImageArrayOutput) ToCopyImageArrayOutput() CopyImageArrayOutput {
	return o
}

func (o CopyImageArrayOutput) ToCopyImageArrayOutputWithContext(ctx context.Context) CopyImageArrayOutput {
	return o
}

func (o CopyImageArrayOutput) Index(i pulumi.IntInput) CopyImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CopyImage {
		return vs[0].([]CopyImage)[vs[1].(int)]
	}).(CopyImageOutput)
}

type CopyImageMapOutput struct{ *pulumi.OutputState }

func (CopyImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CopyImage)(nil))
}

func (o CopyImageMapOutput) ToCopyImageMapOutput() CopyImageMapOutput {
	return o
}

func (o CopyImageMapOutput) ToCopyImageMapOutputWithContext(ctx context.Context) CopyImageMapOutput {
	return o
}

func (o CopyImageMapOutput) MapIndex(k pulumi.StringInput) CopyImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CopyImage {
		return vs[0].(map[string]CopyImage)[vs[1].(string)]
	}).(CopyImageOutput)
}

func init() {
	pulumi.RegisterOutputType(CopyImageOutput{})
	pulumi.RegisterOutputType(CopyImagePtrOutput{})
	pulumi.RegisterOutputType(CopyImageArrayOutput{})
	pulumi.RegisterOutputType(CopyImageMapOutput{})
}
