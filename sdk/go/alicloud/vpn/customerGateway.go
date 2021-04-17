// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// VPN customer gateway can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example cgw-abc123456
// ```
type CustomerGateway struct {
	pulumi.CustomResourceState

	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	var resource CustomerGateway
	err := ctx.RegisterResource("alicloud:vpn/customerGateway:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("alicloud:vpn/customerGateway:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
	// The description of the VPN customer gateway instance.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress *string `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name *string `pulumi:"name"`
}

type CustomerGatewayState struct {
	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringPtrInput
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringPtrInput
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// The description of the VPN customer gateway instance.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress string `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringInput
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringPtrInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerGateway)(nil))
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

func (i *CustomerGateway) ToCustomerGatewayPtrOutput() CustomerGatewayPtrOutput {
	return i.ToCustomerGatewayPtrOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayPtrOutputWithContext(ctx context.Context) CustomerGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayPtrOutput)
}

type CustomerGatewayPtrInput interface {
	pulumi.Input

	ToCustomerGatewayPtrOutput() CustomerGatewayPtrOutput
	ToCustomerGatewayPtrOutputWithContext(ctx context.Context) CustomerGatewayPtrOutput
}

type customerGatewayPtrType CustomerGatewayArgs

func (*customerGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil))
}

func (i *customerGatewayPtrType) ToCustomerGatewayPtrOutput() CustomerGatewayPtrOutput {
	return i.ToCustomerGatewayPtrOutputWithContext(context.Background())
}

func (i *customerGatewayPtrType) ToCustomerGatewayPtrOutputWithContext(ctx context.Context) CustomerGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayPtrOutput)
}

// CustomerGatewayArrayInput is an input type that accepts CustomerGatewayArray and CustomerGatewayArrayOutput values.
// You can construct a concrete instance of `CustomerGatewayArrayInput` via:
//
//          CustomerGatewayArray{ CustomerGatewayArgs{...} }
type CustomerGatewayArrayInput interface {
	pulumi.Input

	ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput
	ToCustomerGatewayArrayOutputWithContext(context.Context) CustomerGatewayArrayOutput
}

type CustomerGatewayArray []CustomerGatewayInput

func (CustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CustomerGateway)(nil))
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return i.ToCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayArrayOutput)
}

// CustomerGatewayMapInput is an input type that accepts CustomerGatewayMap and CustomerGatewayMapOutput values.
// You can construct a concrete instance of `CustomerGatewayMapInput` via:
//
//          CustomerGatewayMap{ "key": CustomerGatewayArgs{...} }
type CustomerGatewayMapInput interface {
	pulumi.Input

	ToCustomerGatewayMapOutput() CustomerGatewayMapOutput
	ToCustomerGatewayMapOutputWithContext(context.Context) CustomerGatewayMapOutput
}

type CustomerGatewayMap map[string]CustomerGatewayInput

func (CustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CustomerGateway)(nil))
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return i.ToCustomerGatewayMapOutputWithContext(context.Background())
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayMapOutput)
}

type CustomerGatewayOutput struct {
	*pulumi.OutputState
}

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerGateway)(nil))
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayPtrOutput() CustomerGatewayPtrOutput {
	return o.ToCustomerGatewayPtrOutputWithContext(context.Background())
}

func (o CustomerGatewayOutput) ToCustomerGatewayPtrOutputWithContext(ctx context.Context) CustomerGatewayPtrOutput {
	return o.ApplyT(func(v CustomerGateway) *CustomerGateway {
		return &v
	}).(CustomerGatewayPtrOutput)
}

type CustomerGatewayPtrOutput struct {
	*pulumi.OutputState
}

func (CustomerGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil))
}

func (o CustomerGatewayPtrOutput) ToCustomerGatewayPtrOutput() CustomerGatewayPtrOutput {
	return o
}

func (o CustomerGatewayPtrOutput) ToCustomerGatewayPtrOutputWithContext(ctx context.Context) CustomerGatewayPtrOutput {
	return o
}

type CustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerGateway)(nil))
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) Index(i pulumi.IntInput) CustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerGateway {
		return vs[0].([]CustomerGateway)[vs[1].(int)]
	}).(CustomerGatewayOutput)
}

type CustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (CustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomerGateway)(nil))
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) CustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomerGateway {
		return vs[0].(map[string]CustomerGateway)[vs[1].(string)]
	}).(CustomerGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayPtrOutput{})
	pulumi.RegisterOutputType(CustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayMapOutput{})
}
