// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = new alicloud.vpc.Network("foo", {cidrBlock: "172.16.0.0/12"});
 * const _interface = new alicloud.vpc.RouterInterface("interface", {
 *     oppositeRegion: "cn-beijing",
 *     routerType: "VRouter",
 *     routerId: foo.routerId,
 *     role: "InitiatingSide",
 *     specification: "Large.2",
 *     description: "test1",
 * });
 * ```
 *
 * ## Import
 *
 * The router interface can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/routerInterface:RouterInterface interface ri-abc123456
 * ```
 */
export class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceState, opts?: pulumi.CustomResourceOptions): RouterInterface {
        return new RouterInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routerInterface:RouterInterface';

    /**
     * Returns true if the given object is an instance of RouterInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterInterface.__pulumiType;
    }

    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    public /*out*/ readonly accessPointId!: pulumi.Output<string>;
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    public readonly healthCheckSourceIp!: pulumi.Output<string | undefined>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    public readonly healthCheckTargetIp!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    public readonly oppositeAccessPointId!: pulumi.Output<string | undefined>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     */
    public /*out*/ readonly oppositeInterfaceId!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_owner_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_owner_id' instead.
     */
    public /*out*/ readonly oppositeInterfaceOwnerId!: pulumi.Output<string>;
    /**
     * The Region of peer side.
     */
    public readonly oppositeRegion!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_router_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     */
    public /*out*/ readonly oppositeRouterId!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     *
     * @deprecated Attribute 'opposite_router_type' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     */
    public /*out*/ readonly oppositeRouterType!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The Router ID.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    public readonly routerType!: pulumi.Output<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    public readonly specification!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterInterfaceArgs | RouterInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterInterfaceState | undefined;
            resourceInputs["accessPointId"] = state ? state.accessPointId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthCheckSourceIp"] = state ? state.healthCheckSourceIp : undefined;
            resourceInputs["healthCheckTargetIp"] = state ? state.healthCheckTargetIp : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oppositeAccessPointId"] = state ? state.oppositeAccessPointId : undefined;
            resourceInputs["oppositeInterfaceId"] = state ? state.oppositeInterfaceId : undefined;
            resourceInputs["oppositeInterfaceOwnerId"] = state ? state.oppositeInterfaceOwnerId : undefined;
            resourceInputs["oppositeRegion"] = state ? state.oppositeRegion : undefined;
            resourceInputs["oppositeRouterId"] = state ? state.oppositeRouterId : undefined;
            resourceInputs["oppositeRouterType"] = state ? state.oppositeRouterType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["routerId"] = state ? state.routerId : undefined;
            resourceInputs["routerType"] = state ? state.routerType : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
        } else {
            const args = argsOrState as RouterInterfaceArgs | undefined;
            if ((!args || args.oppositeRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oppositeRegion'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.routerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerId'");
            }
            if ((!args || args.routerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["healthCheckSourceIp"] = args ? args.healthCheckSourceIp : undefined;
            resourceInputs["healthCheckTargetIp"] = args ? args.healthCheckTargetIp : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oppositeAccessPointId"] = args ? args.oppositeAccessPointId : undefined;
            resourceInputs["oppositeRegion"] = args ? args.oppositeRegion : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["routerId"] = args ? args.routerId : undefined;
            resourceInputs["routerType"] = args ? args.routerType : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["accessPointId"] = undefined /*out*/;
            resourceInputs["oppositeInterfaceId"] = undefined /*out*/;
            resourceInputs["oppositeInterfaceOwnerId"] = undefined /*out*/;
            resourceInputs["oppositeRouterId"] = undefined /*out*/;
            resourceInputs["oppositeRouterType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterInterface resources.
 */
export interface RouterInterfaceState {
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    accessPointId?: pulumi.Input<string>;
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    description?: pulumi.Input<string>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    healthCheckSourceIp?: pulumi.Input<string>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    healthCheckTargetIp?: pulumi.Input<string>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    name?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    oppositeAccessPointId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     */
    oppositeInterfaceId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_owner_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_owner_id' instead.
     */
    oppositeInterfaceOwnerId?: pulumi.Input<string>;
    /**
     * The Region of peer side.
     */
    oppositeRegion?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_router_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     */
    oppositeRouterId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     *
     * @deprecated Attribute 'opposite_router_type' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     */
    oppositeRouterType?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    role?: pulumi.Input<string>;
    /**
     * The Router ID.
     */
    routerId?: pulumi.Input<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    routerType?: pulumi.Input<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    specification?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterInterface resource.
 */
export interface RouterInterfaceArgs {
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    description?: pulumi.Input<string>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    healthCheckSourceIp?: pulumi.Input<string>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    healthCheckTargetIp?: pulumi.Input<string>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    name?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    oppositeAccessPointId?: pulumi.Input<string>;
    /**
     * The Region of peer side.
     */
    oppositeRegion: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    role: pulumi.Input<string>;
    /**
     * The Router ID.
     */
    routerId: pulumi.Input<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    routerType: pulumi.Input<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    specification?: pulumi.Input<string>;
}
