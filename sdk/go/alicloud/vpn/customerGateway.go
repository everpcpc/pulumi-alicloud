// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

func (CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerGateway)(nil)).Elem()
}

func (i CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

type CustomerGatewayOutput struct {
	*pulumi.OutputState
}

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerGatewayOutput)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
}
