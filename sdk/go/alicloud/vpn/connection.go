// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_connection.html.markdown.
type Connection struct {
	s *pulumi.ResourceState
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOpt) (*Connection, error) {
	if args == nil || args.CustomerGatewayId == nil {
		return nil, errors.New("missing required argument 'CustomerGatewayId'")
	}
	if args == nil || args.LocalSubnets == nil {
		return nil, errors.New("missing required argument 'LocalSubnets'")
	}
	if args == nil || args.RemoteSubnets == nil {
		return nil, errors.New("missing required argument 'RemoteSubnets'")
	}
	if args == nil || args.VpnGatewayId == nil {
		return nil, errors.New("missing required argument 'VpnGatewayId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["customerGatewayId"] = nil
		inputs["effectImmediately"] = nil
		inputs["ikeConfigs"] = nil
		inputs["ipsecConfigs"] = nil
		inputs["localSubnets"] = nil
		inputs["name"] = nil
		inputs["remoteSubnets"] = nil
		inputs["vpnGatewayId"] = nil
	} else {
		inputs["customerGatewayId"] = args.CustomerGatewayId
		inputs["effectImmediately"] = args.EffectImmediately
		inputs["ikeConfigs"] = args.IkeConfigs
		inputs["ipsecConfigs"] = args.IpsecConfigs
		inputs["localSubnets"] = args.LocalSubnets
		inputs["name"] = args.Name
		inputs["remoteSubnets"] = args.RemoteSubnets
		inputs["vpnGatewayId"] = args.VpnGatewayId
	}
	inputs["status"] = nil
	s, err := ctx.RegisterResource("alicloud:vpn/connection:Connection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Connection{s: s}, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConnectionState, opts ...pulumi.ResourceOpt) (*Connection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["customerGatewayId"] = state.CustomerGatewayId
		inputs["effectImmediately"] = state.EffectImmediately
		inputs["ikeConfigs"] = state.IkeConfigs
		inputs["ipsecConfigs"] = state.IpsecConfigs
		inputs["localSubnets"] = state.LocalSubnets
		inputs["name"] = state.Name
		inputs["remoteSubnets"] = state.RemoteSubnets
		inputs["status"] = state.Status
		inputs["vpnGatewayId"] = state.VpnGatewayId
	}
	s, err := ctx.ReadResource("alicloud:vpn/connection:Connection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Connection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Connection) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Connection) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the customer gateway.
func (r *Connection) CustomerGatewayId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customerGatewayId"])
}

// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
func (r *Connection) EffectImmediately() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["effectImmediately"])
}

// The configurations of phase-one negotiation.
func (r *Connection) IkeConfigs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ikeConfigs"])
}

// The configurations of phase-two negotiation.
func (r *Connection) IpsecConfigs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipsecConfigs"])
}

// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
func (r *Connection) LocalSubnets() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["localSubnets"])
}

// The name of the IPsec connection.
func (r *Connection) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
func (r *Connection) RemoteSubnets() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["remoteSubnets"])
}

// The status of VPN connection.
func (r *Connection) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// The ID of the VPN gateway.
func (r *Connection) VpnGatewayId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpnGatewayId"])
}

// Input properties used for looking up and filtering Connection resources.
type ConnectionState struct {
	// The ID of the customer gateway.
	CustomerGatewayId interface{}
	// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
	EffectImmediately interface{}
	// The configurations of phase-one negotiation.
	IkeConfigs interface{}
	// The configurations of phase-two negotiation.
	IpsecConfigs interface{}
	// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
	LocalSubnets interface{}
	// The name of the IPsec connection.
	Name interface{}
	// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
	RemoteSubnets interface{}
	// The status of VPN connection.
	Status interface{}
	// The ID of the VPN gateway.
	VpnGatewayId interface{}
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The ID of the customer gateway.
	CustomerGatewayId interface{}
	// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
	EffectImmediately interface{}
	// The configurations of phase-one negotiation.
	IkeConfigs interface{}
	// The configurations of phase-two negotiation.
	IpsecConfigs interface{}
	// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
	LocalSubnets interface{}
	// The name of the IPsec connection.
	Name interface{}
	// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
	RemoteSubnets interface{}
	// The ID of the VPN gateway.
	VpnGatewayId interface{}
}
