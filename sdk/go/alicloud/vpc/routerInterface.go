// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/router_interface.html.markdown.
type RouterInterface struct {
	s *pulumi.ResourceState
}

// NewRouterInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterInterface(ctx *pulumi.Context,
	name string, args *RouterInterfaceArgs, opts ...pulumi.ResourceOpt) (*RouterInterface, error) {
	if args == nil || args.OppositeRegion == nil {
		return nil, errors.New("missing required argument 'OppositeRegion'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.RouterId == nil {
		return nil, errors.New("missing required argument 'RouterId'")
	}
	if args == nil || args.RouterType == nil {
		return nil, errors.New("missing required argument 'RouterType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["healthCheckSourceIp"] = nil
		inputs["healthCheckTargetIp"] = nil
		inputs["instanceChargeType"] = nil
		inputs["name"] = nil
		inputs["oppositeAccessPointId"] = nil
		inputs["oppositeRegion"] = nil
		inputs["period"] = nil
		inputs["role"] = nil
		inputs["routerId"] = nil
		inputs["routerType"] = nil
		inputs["specification"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["healthCheckSourceIp"] = args.HealthCheckSourceIp
		inputs["healthCheckTargetIp"] = args.HealthCheckTargetIp
		inputs["instanceChargeType"] = args.InstanceChargeType
		inputs["name"] = args.Name
		inputs["oppositeAccessPointId"] = args.OppositeAccessPointId
		inputs["oppositeRegion"] = args.OppositeRegion
		inputs["period"] = args.Period
		inputs["role"] = args.Role
		inputs["routerId"] = args.RouterId
		inputs["routerType"] = args.RouterType
		inputs["specification"] = args.Specification
	}
	inputs["accessPointId"] = nil
	inputs["oppositeInterfaceId"] = nil
	inputs["oppositeInterfaceOwnerId"] = nil
	inputs["oppositeRouterId"] = nil
	inputs["oppositeRouterType"] = nil
	s, err := ctx.RegisterResource("alicloud:vpc/routerInterface:RouterInterface", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouterInterface{s: s}, nil
}

// GetRouterInterface gets an existing RouterInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterface(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouterInterfaceState, opts ...pulumi.ResourceOpt) (*RouterInterface, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessPointId"] = state.AccessPointId
		inputs["description"] = state.Description
		inputs["healthCheckSourceIp"] = state.HealthCheckSourceIp
		inputs["healthCheckTargetIp"] = state.HealthCheckTargetIp
		inputs["instanceChargeType"] = state.InstanceChargeType
		inputs["name"] = state.Name
		inputs["oppositeAccessPointId"] = state.OppositeAccessPointId
		inputs["oppositeInterfaceId"] = state.OppositeInterfaceId
		inputs["oppositeInterfaceOwnerId"] = state.OppositeInterfaceOwnerId
		inputs["oppositeRegion"] = state.OppositeRegion
		inputs["oppositeRouterId"] = state.OppositeRouterId
		inputs["oppositeRouterType"] = state.OppositeRouterType
		inputs["period"] = state.Period
		inputs["role"] = state.Role
		inputs["routerId"] = state.RouterId
		inputs["routerType"] = state.RouterType
		inputs["specification"] = state.Specification
	}
	s, err := ctx.ReadResource("alicloud:vpc/routerInterface:RouterInterface", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouterInterface{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RouterInterface) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RouterInterface) ID() pulumi.IDOutput {
	return r.s.ID()
}

// It has been deprecated from version 1.11.0.
func (r *RouterInterface) AccessPointId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessPointId"])
}

// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
func (r *RouterInterface) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
func (r *RouterInterface) HealthCheckSourceIp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["healthCheckSourceIp"])
}

// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
func (r *RouterInterface) HealthCheckTargetIp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["healthCheckTargetIp"])
}

// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
func (r *RouterInterface) InstanceChargeType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceChargeType"])
}

// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
func (r *RouterInterface) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// It has been deprecated from version 1.11.0.
func (r *RouterInterface) OppositeAccessPointId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeAccessPointId"])
}

// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
func (r *RouterInterface) OppositeInterfaceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeInterfaceId"])
}

// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
func (r *RouterInterface) OppositeInterfaceOwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeInterfaceOwnerId"])
}

// The Region of peer side.
func (r *RouterInterface) OppositeRegion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeRegion"])
}

// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
func (r *RouterInterface) OppositeRouterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeRouterId"])
}

// It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
func (r *RouterInterface) OppositeRouterType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["oppositeRouterType"])
}

// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
func (r *RouterInterface) Period() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["period"])
}

// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
func (r *RouterInterface) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// The Router ID.
func (r *RouterInterface) RouterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["routerId"])
}

// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
func (r *RouterInterface) RouterType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["routerType"])
}

// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
func (r *RouterInterface) Specification() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["specification"])
}

// Input properties used for looking up and filtering RouterInterface resources.
type RouterInterfaceState struct {
	// It has been deprecated from version 1.11.0.
	AccessPointId interface{}
	// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
	Description interface{}
	// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
	HealthCheckSourceIp interface{}
	// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
	HealthCheckTargetIp interface{}
	// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
	InstanceChargeType interface{}
	// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
	Name interface{}
	// It has been deprecated from version 1.11.0.
	OppositeAccessPointId interface{}
	// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
	OppositeInterfaceId interface{}
	// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
	OppositeInterfaceOwnerId interface{}
	// The Region of peer side.
	OppositeRegion interface{}
	// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
	OppositeRouterId interface{}
	// It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
	OppositeRouterType interface{}
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period interface{}
	// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
	Role interface{}
	// The Router ID.
	RouterId interface{}
	// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
	RouterType interface{}
	// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
	Specification interface{}
}

// The set of arguments for constructing a RouterInterface resource.
type RouterInterfaceArgs struct {
	// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
	Description interface{}
	// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
	HealthCheckSourceIp interface{}
	// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
	HealthCheckTargetIp interface{}
	// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
	InstanceChargeType interface{}
	// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
	Name interface{}
	// It has been deprecated from version 1.11.0.
	OppositeAccessPointId interface{}
	// The Region of peer side.
	OppositeRegion interface{}
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period interface{}
	// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
	Role interface{}
	// The Router ID.
	RouterId interface{}
	// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
	RouterType interface{}
	// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
	Specification interface{}
}
