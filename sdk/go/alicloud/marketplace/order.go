// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Market order can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:marketplace/order:Order order your-order-id
// ```
type Order struct {
	pulumi.CustomResourceState

	// Service providers customize additional components.
	Components pulumi.MapOutput `pulumi:"components"`
	// The coupon id of the market product.
	CouponId pulumi.StringPtrOutput `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion pulumi.StringOutput `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrOutput `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringOutput `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrOutput `pulumi:"quantity"`
}

// NewOrder registers a new resource with the given unique name, arguments, and options.
func NewOrder(ctx *pulumi.Context,
	name string, args *OrderArgs, opts ...pulumi.ResourceOption) (*Order, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PackageVersion == nil {
		return nil, errors.New("invalid value for required argument 'PackageVersion'")
	}
	if args.PricingCycle == nil {
		return nil, errors.New("invalid value for required argument 'PricingCycle'")
	}
	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	var resource Order
	err := ctx.RegisterResource("alicloud:marketplace/order:Order", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrder gets an existing Order resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderState, opts ...pulumi.ResourceOption) (*Order, error) {
	var resource Order
	err := ctx.ReadResource("alicloud:marketplace/order:Order", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Order resources.
type orderState struct {
	// Service providers customize additional components.
	Components map[string]interface{} `pulumi:"components"`
	// The coupon id of the market product.
	CouponId *string `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration *int `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion *string `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType *string `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode *string `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity *int `pulumi:"quantity"`
}

type OrderState struct {
	// Service providers customize additional components.
	Components pulumi.MapInput
	// The coupon id of the market product.
	CouponId pulumi.StringPtrInput
	// The number of purchase cycles.
	Duration pulumi.IntPtrInput
	// The package version of the market product.
	PackageVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrInput
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringPtrInput
	// The productCode of market place product.
	ProductCode pulumi.StringPtrInput
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrInput
}

func (OrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderState)(nil)).Elem()
}

type orderArgs struct {
	// Service providers customize additional components.
	Components map[string]interface{} `pulumi:"components"`
	// The coupon id of the market product.
	CouponId *string `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration *int `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion string `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType *string `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle string `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode string `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity *int `pulumi:"quantity"`
}

// The set of arguments for constructing a Order resource.
type OrderArgs struct {
	// Service providers customize additional components.
	Components pulumi.MapInput
	// The coupon id of the market product.
	CouponId pulumi.StringPtrInput
	// The number of purchase cycles.
	Duration pulumi.IntPtrInput
	// The package version of the market product.
	PackageVersion pulumi.StringInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrInput
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringInput
	// The productCode of market place product.
	ProductCode pulumi.StringInput
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrInput
}

func (OrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderArgs)(nil)).Elem()
}

type OrderInput interface {
	pulumi.Input

	ToOrderOutput() OrderOutput
	ToOrderOutputWithContext(ctx context.Context) OrderOutput
}

func (Order) ElementType() reflect.Type {
	return reflect.TypeOf((*Order)(nil)).Elem()
}

func (i Order) ToOrderOutput() OrderOutput {
	return i.ToOrderOutputWithContext(context.Background())
}

func (i Order) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderOutput)
}

type OrderOutput struct {
	*pulumi.OutputState
}

func (OrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderOutput)(nil)).Elem()
}

func (o OrderOutput) ToOrderOutput() OrderOutput {
	return o
}

func (o OrderOutput) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrderOutput{})
}
