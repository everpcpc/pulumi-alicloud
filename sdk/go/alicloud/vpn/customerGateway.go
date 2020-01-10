// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_customer_gateway.html.markdown.
type CustomerGateway struct {
	s *pulumi.ResourceState
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOpt) (*CustomerGateway, error) {
	if args == nil || args.IpAddress == nil {
		return nil, errors.New("missing required argument 'IpAddress'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["ipAddress"] = nil
		inputs["name"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["ipAddress"] = args.IpAddress
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("alicloud:vpn/customerGateway:CustomerGateway", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomerGateway{s: s}, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CustomerGatewayState, opts ...pulumi.ResourceOpt) (*CustomerGateway, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["ipAddress"] = state.IpAddress
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("alicloud:vpn/customerGateway:CustomerGateway", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomerGateway{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CustomerGateway) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CustomerGateway) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The description of the VPN customer gateway instance.
func (r *CustomerGateway) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The IP address of the customer gateway.
func (r *CustomerGateway) IpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipAddress"])
}

// The name of the VPN customer gateway. Defaults to null.
func (r *CustomerGateway) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering CustomerGateway resources.
type CustomerGatewayState struct {
	// The description of the VPN customer gateway instance.
	Description interface{}
	// The IP address of the customer gateway.
	IpAddress interface{}
	// The name of the VPN customer gateway. Defaults to null.
	Name interface{}
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The description of the VPN customer gateway instance.
	Description interface{}
	// The IP address of the customer gateway.
	IpAddress interface{}
	// The name of the VPN customer gateway. Defaults to null.
	Name interface{}
}
