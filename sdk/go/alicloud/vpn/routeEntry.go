// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/vpn_route_entry.html.markdown.
type RouteEntry struct {
	s *pulumi.ResourceState
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOpt) (*RouteEntry, error) {
	if args == nil || args.NextHop == nil {
		return nil, errors.New("missing required argument 'NextHop'")
	}
	if args == nil || args.PublishVpc == nil {
		return nil, errors.New("missing required argument 'PublishVpc'")
	}
	if args == nil || args.RouteDest == nil {
		return nil, errors.New("missing required argument 'RouteDest'")
	}
	if args == nil || args.VpnGatewayId == nil {
		return nil, errors.New("missing required argument 'VpnGatewayId'")
	}
	if args == nil || args.Weight == nil {
		return nil, errors.New("missing required argument 'Weight'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["nextHop"] = nil
		inputs["publishVpc"] = nil
		inputs["routeDest"] = nil
		inputs["vpnGatewayId"] = nil
		inputs["weight"] = nil
	} else {
		inputs["nextHop"] = args.NextHop
		inputs["publishVpc"] = args.PublishVpc
		inputs["routeDest"] = args.RouteDest
		inputs["vpnGatewayId"] = args.VpnGatewayId
		inputs["weight"] = args.Weight
	}
	s, err := ctx.RegisterResource("alicloud:vpn/routeEntry:RouteEntry", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteEntry{s: s}, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteEntryState, opts ...pulumi.ResourceOpt) (*RouteEntry, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["nextHop"] = state.NextHop
		inputs["publishVpc"] = state.PublishVpc
		inputs["routeDest"] = state.RouteDest
		inputs["vpnGatewayId"] = state.VpnGatewayId
		inputs["weight"] = state.Weight
	}
	s, err := ctx.ReadResource("alicloud:vpn/routeEntry:RouteEntry", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteEntry{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RouteEntry) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RouteEntry) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The next hop of the destination route.
func (r *RouteEntry) NextHop() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["nextHop"])
}

// Whether to issue the destination route to the VPC.
func (r *RouteEntry) PublishVpc() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["publishVpc"])
}

// The destination network segment of the destination route.
func (r *RouteEntry) RouteDest() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["routeDest"])
}

// The id of the vpn gateway.
func (r *RouteEntry) VpnGatewayId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpnGatewayId"])
}

// The value should be 0 or 100.
func (r *RouteEntry) Weight() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["weight"])
}

// Input properties used for looking up and filtering RouteEntry resources.
type RouteEntryState struct {
	// The next hop of the destination route.
	NextHop interface{}
	// Whether to issue the destination route to the VPC.
	PublishVpc interface{}
	// The destination network segment of the destination route.
	RouteDest interface{}
	// The id of the vpn gateway.
	VpnGatewayId interface{}
	// The value should be 0 or 100.
	Weight interface{}
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// The next hop of the destination route.
	NextHop interface{}
	// Whether to issue the destination route to the VPC.
	PublishVpc interface{}
	// The destination network segment of the destination route.
	RouteDest interface{}
	// The id of the vpn gateway.
	VpnGatewayId interface{}
	// The value should be 0 or 100.
	Weight interface{}
}
