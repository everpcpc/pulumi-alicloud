// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SMS Short Url resource.
//
// For information about SMS Short Url and how to use it, see [What is Short Url](https://help.aliyun.com/document_detail/419291.html).
//
// > **NOTE:** Available in v1.178.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sms.NewShortUrl(ctx, "example", &sms.ShortUrlArgs{
//				EffectiveDays: pulumi.Int(30),
//				ShortUrlName:  pulumi.String("example_value"),
//				SourceUrl:     pulumi.String("example_value"),
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
// SMS Short Url can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sms/shortUrl:ShortUrl example <id>
//
// ```
type ShortUrl struct {
	pulumi.CustomResourceState

	// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
	EffectiveDays pulumi.IntOutput `pulumi:"effectiveDays"`
	// The name of the resource.
	ShortUrlName pulumi.StringOutput `pulumi:"shortUrlName"`
	// The original link address.
	SourceUrl pulumi.StringOutput `pulumi:"sourceUrl"`
	// Short chain status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewShortUrl registers a new resource with the given unique name, arguments, and options.
func NewShortUrl(ctx *pulumi.Context,
	name string, args *ShortUrlArgs, opts ...pulumi.ResourceOption) (*ShortUrl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EffectiveDays == nil {
		return nil, errors.New("invalid value for required argument 'EffectiveDays'")
	}
	if args.ShortUrlName == nil {
		return nil, errors.New("invalid value for required argument 'ShortUrlName'")
	}
	if args.SourceUrl == nil {
		return nil, errors.New("invalid value for required argument 'SourceUrl'")
	}
	var resource ShortUrl
	err := ctx.RegisterResource("alicloud:sms/shortUrl:ShortUrl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShortUrl gets an existing ShortUrl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShortUrl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShortUrlState, opts ...pulumi.ResourceOption) (*ShortUrl, error) {
	var resource ShortUrl
	err := ctx.ReadResource("alicloud:sms/shortUrl:ShortUrl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShortUrl resources.
type shortUrlState struct {
	// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
	EffectiveDays *int `pulumi:"effectiveDays"`
	// The name of the resource.
	ShortUrlName *string `pulumi:"shortUrlName"`
	// The original link address.
	SourceUrl *string `pulumi:"sourceUrl"`
	// Short chain status.
	Status *string `pulumi:"status"`
}

type ShortUrlState struct {
	// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
	EffectiveDays pulumi.IntPtrInput
	// The name of the resource.
	ShortUrlName pulumi.StringPtrInput
	// The original link address.
	SourceUrl pulumi.StringPtrInput
	// Short chain status.
	Status pulumi.StringPtrInput
}

func (ShortUrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*shortUrlState)(nil)).Elem()
}

type shortUrlArgs struct {
	// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
	EffectiveDays int `pulumi:"effectiveDays"`
	// The name of the resource.
	ShortUrlName string `pulumi:"shortUrlName"`
	// The original link address.
	SourceUrl string `pulumi:"sourceUrl"`
}

// The set of arguments for constructing a ShortUrl resource.
type ShortUrlArgs struct {
	// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
	EffectiveDays pulumi.IntInput
	// The name of the resource.
	ShortUrlName pulumi.StringInput
	// The original link address.
	SourceUrl pulumi.StringInput
}

func (ShortUrlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shortUrlArgs)(nil)).Elem()
}

type ShortUrlInput interface {
	pulumi.Input

	ToShortUrlOutput() ShortUrlOutput
	ToShortUrlOutputWithContext(ctx context.Context) ShortUrlOutput
}

func (*ShortUrl) ElementType() reflect.Type {
	return reflect.TypeOf((**ShortUrl)(nil)).Elem()
}

func (i *ShortUrl) ToShortUrlOutput() ShortUrlOutput {
	return i.ToShortUrlOutputWithContext(context.Background())
}

func (i *ShortUrl) ToShortUrlOutputWithContext(ctx context.Context) ShortUrlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShortUrlOutput)
}

// ShortUrlArrayInput is an input type that accepts ShortUrlArray and ShortUrlArrayOutput values.
// You can construct a concrete instance of `ShortUrlArrayInput` via:
//
//	ShortUrlArray{ ShortUrlArgs{...} }
type ShortUrlArrayInput interface {
	pulumi.Input

	ToShortUrlArrayOutput() ShortUrlArrayOutput
	ToShortUrlArrayOutputWithContext(context.Context) ShortUrlArrayOutput
}

type ShortUrlArray []ShortUrlInput

func (ShortUrlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShortUrl)(nil)).Elem()
}

func (i ShortUrlArray) ToShortUrlArrayOutput() ShortUrlArrayOutput {
	return i.ToShortUrlArrayOutputWithContext(context.Background())
}

func (i ShortUrlArray) ToShortUrlArrayOutputWithContext(ctx context.Context) ShortUrlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShortUrlArrayOutput)
}

// ShortUrlMapInput is an input type that accepts ShortUrlMap and ShortUrlMapOutput values.
// You can construct a concrete instance of `ShortUrlMapInput` via:
//
//	ShortUrlMap{ "key": ShortUrlArgs{...} }
type ShortUrlMapInput interface {
	pulumi.Input

	ToShortUrlMapOutput() ShortUrlMapOutput
	ToShortUrlMapOutputWithContext(context.Context) ShortUrlMapOutput
}

type ShortUrlMap map[string]ShortUrlInput

func (ShortUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShortUrl)(nil)).Elem()
}

func (i ShortUrlMap) ToShortUrlMapOutput() ShortUrlMapOutput {
	return i.ToShortUrlMapOutputWithContext(context.Background())
}

func (i ShortUrlMap) ToShortUrlMapOutputWithContext(ctx context.Context) ShortUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShortUrlMapOutput)
}

type ShortUrlOutput struct{ *pulumi.OutputState }

func (ShortUrlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShortUrl)(nil)).Elem()
}

func (o ShortUrlOutput) ToShortUrlOutput() ShortUrlOutput {
	return o
}

func (o ShortUrlOutput) ToShortUrlOutputWithContext(ctx context.Context) ShortUrlOutput {
	return o
}

// Short chain service use validity period. Valid values: `30`, `60`, `90`. The unit is days, and the maximum validity period is 90 days.
func (o ShortUrlOutput) EffectiveDays() pulumi.IntOutput {
	return o.ApplyT(func(v *ShortUrl) pulumi.IntOutput { return v.EffectiveDays }).(pulumi.IntOutput)
}

// The name of the resource.
func (o ShortUrlOutput) ShortUrlName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShortUrl) pulumi.StringOutput { return v.ShortUrlName }).(pulumi.StringOutput)
}

// The original link address.
func (o ShortUrlOutput) SourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ShortUrl) pulumi.StringOutput { return v.SourceUrl }).(pulumi.StringOutput)
}

// Short chain status.
func (o ShortUrlOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ShortUrl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ShortUrlArrayOutput struct{ *pulumi.OutputState }

func (ShortUrlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShortUrl)(nil)).Elem()
}

func (o ShortUrlArrayOutput) ToShortUrlArrayOutput() ShortUrlArrayOutput {
	return o
}

func (o ShortUrlArrayOutput) ToShortUrlArrayOutputWithContext(ctx context.Context) ShortUrlArrayOutput {
	return o
}

func (o ShortUrlArrayOutput) Index(i pulumi.IntInput) ShortUrlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShortUrl {
		return vs[0].([]*ShortUrl)[vs[1].(int)]
	}).(ShortUrlOutput)
}

type ShortUrlMapOutput struct{ *pulumi.OutputState }

func (ShortUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShortUrl)(nil)).Elem()
}

func (o ShortUrlMapOutput) ToShortUrlMapOutput() ShortUrlMapOutput {
	return o
}

func (o ShortUrlMapOutput) ToShortUrlMapOutputWithContext(ctx context.Context) ShortUrlMapOutput {
	return o
}

func (o ShortUrlMapOutput) MapIndex(k pulumi.StringInput) ShortUrlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShortUrl {
		return vs[0].(map[string]*ShortUrl)[vs[1].(string)]
	}).(ShortUrlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShortUrlInput)(nil)).Elem(), &ShortUrl{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShortUrlArrayInput)(nil)).Elem(), ShortUrlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShortUrlMapInput)(nil)).Elem(), ShortUrlMap{})
	pulumi.RegisterOutputType(ShortUrlOutput{})
	pulumi.RegisterOutputType(ShortUrlArrayOutput{})
	pulumi.RegisterOutputType(ShortUrlMapOutput{})
}
