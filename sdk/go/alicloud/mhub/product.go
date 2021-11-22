// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MHUB Product resource.
//
// For information about MHUB Product and how to use it, see [What is Product](https://help.aliyun.com/product/65109.html).
//
// > **NOTE:** Available in v1.138.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mhub.NewProduct(ctx, "example", &mhub.ProductArgs{
// 			ProductName: pulumi.String("example_value"),
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
// MHUB Product can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:mhub/product:Product example <id>
// ```
type Product struct {
	pulumi.CustomResourceState

	// ProductName.
	ProductName pulumi.StringOutput `pulumi:"productName"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductName == nil {
		return nil, errors.New("invalid value for required argument 'ProductName'")
	}
	var resource Product
	err := ctx.RegisterResource("alicloud:mhub/product:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("alicloud:mhub/product:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
	// ProductName.
	ProductName *string `pulumi:"productName"`
}

type ProductState struct {
	// ProductName.
	ProductName pulumi.StringPtrInput
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// ProductName.
	ProductName string `pulumi:"productName"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// ProductName.
	ProductName pulumi.StringInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

func (i *Product) ToProductPtrOutput() ProductPtrOutput {
	return i.ToProductPtrOutputWithContext(context.Background())
}

func (i *Product) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductPtrOutput)
}

type ProductPtrInput interface {
	pulumi.Input

	ToProductPtrOutput() ProductPtrOutput
	ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput
}

type productPtrType ProductArgs

func (*productPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil))
}

func (i *productPtrType) ToProductPtrOutput() ProductPtrOutput {
	return i.ToProductPtrOutputWithContext(context.Background())
}

func (i *productPtrType) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductPtrOutput)
}

// ProductArrayInput is an input type that accepts ProductArray and ProductArrayOutput values.
// You can construct a concrete instance of `ProductArrayInput` via:
//
//          ProductArray{ ProductArgs{...} }
type ProductArrayInput interface {
	pulumi.Input

	ToProductArrayOutput() ProductArrayOutput
	ToProductArrayOutputWithContext(context.Context) ProductArrayOutput
}

type ProductArray []ProductInput

func (ProductArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Product)(nil)).Elem()
}

func (i ProductArray) ToProductArrayOutput() ProductArrayOutput {
	return i.ToProductArrayOutputWithContext(context.Background())
}

func (i ProductArray) ToProductArrayOutputWithContext(ctx context.Context) ProductArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductArrayOutput)
}

// ProductMapInput is an input type that accepts ProductMap and ProductMapOutput values.
// You can construct a concrete instance of `ProductMapInput` via:
//
//          ProductMap{ "key": ProductArgs{...} }
type ProductMapInput interface {
	pulumi.Input

	ToProductMapOutput() ProductMapOutput
	ToProductMapOutputWithContext(context.Context) ProductMapOutput
}

type ProductMap map[string]ProductInput

func (ProductMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Product)(nil)).Elem()
}

func (i ProductMap) ToProductMapOutput() ProductMapOutput {
	return i.ToProductMapOutputWithContext(context.Background())
}

func (i ProductMap) ToProductMapOutputWithContext(ctx context.Context) ProductMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductMapOutput)
}

type ProductOutput struct{ *pulumi.OutputState }

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Product)(nil))
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

func (o ProductOutput) ToProductPtrOutput() ProductPtrOutput {
	return o.ToProductPtrOutputWithContext(context.Background())
}

func (o ProductOutput) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Product) *Product {
		return &v
	}).(ProductPtrOutput)
}

type ProductPtrOutput struct{ *pulumi.OutputState }

func (ProductPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil))
}

func (o ProductPtrOutput) ToProductPtrOutput() ProductPtrOutput {
	return o
}

func (o ProductPtrOutput) ToProductPtrOutputWithContext(ctx context.Context) ProductPtrOutput {
	return o
}

func (o ProductPtrOutput) Elem() ProductOutput {
	return o.ApplyT(func(v *Product) Product {
		if v != nil {
			return *v
		}
		var ret Product
		return ret
	}).(ProductOutput)
}

type ProductArrayOutput struct{ *pulumi.OutputState }

func (ProductArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Product)(nil))
}

func (o ProductArrayOutput) ToProductArrayOutput() ProductArrayOutput {
	return o
}

func (o ProductArrayOutput) ToProductArrayOutputWithContext(ctx context.Context) ProductArrayOutput {
	return o
}

func (o ProductArrayOutput) Index(i pulumi.IntInput) ProductOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Product {
		return vs[0].([]Product)[vs[1].(int)]
	}).(ProductOutput)
}

type ProductMapOutput struct{ *pulumi.OutputState }

func (ProductMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Product)(nil))
}

func (o ProductMapOutput) ToProductMapOutput() ProductMapOutput {
	return o
}

func (o ProductMapOutput) ToProductMapOutputWithContext(ctx context.Context) ProductMapOutput {
	return o
}

func (o ProductMapOutput) MapIndex(k pulumi.StringInput) ProductOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Product {
		return vs[0].(map[string]Product)[vs[1].(string)]
	}).(ProductOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProductInput)(nil)).Elem(), &Product{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductPtrInput)(nil)).Elem(), &Product{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductArrayInput)(nil)).Elem(), ProductArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductMapInput)(nil)).Elem(), ProductMap{})
	pulumi.RegisterOutputType(ProductOutput{})
	pulumi.RegisterOutputType(ProductPtrOutput{})
	pulumi.RegisterOutputType(ProductArrayOutput{})
	pulumi.RegisterOutputType(ProductMapOutput{})
}
