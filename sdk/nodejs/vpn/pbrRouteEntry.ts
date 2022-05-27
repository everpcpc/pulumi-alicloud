// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPN Pbr Route Entry resource.
 *
 * > **NOTE:** Available in 1.162.0+.
 *
 * For information about VPN Pbr Route Entry and how to use it, see [What is VPN Pbr Route Entry](https://www.alibabacloud.com/help/en/doc-detail/127248.html).
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tfacc";
 * const defaultGateways = alicloud.vpn.getGateways({});
 * const defaultCustomerGateway = new alicloud.vpn.CustomerGateway("defaultCustomerGateway", {ipAddress: "192.168.1.1"});
 * const defaultConnection = new alicloud.vpn.Connection("defaultConnection", {
 *     customerGatewayId: defaultCustomerGateway.id,
 *     vpnGatewayId: defaultGateways.then(defaultGateways => defaultGateways.ids?[0]),
 *     localSubnets: ["192.168.2.0/24"],
 *     remoteSubnets: ["192.168.3.0/24"],
 * });
 * const defaultPbrRouteEntry = new alicloud.vpn.PbrRouteEntry("defaultPbrRouteEntry", {
 *     vpnGatewayId: defaultGateways.then(defaultGateways => defaultGateways.ids?[0]),
 *     routeSource: "192.168.1.0/24",
 *     routeDest: "10.0.0.0/24",
 *     nextHop: defaultConnection.id,
 *     weight: 0,
 *     publishVpc: false,
 * });
 * ```
 *
 * ## Import
 *
 * VPN Pbr route entry can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpn/pbrRouteEntry:PbrRouteEntry example <vpn_gateway_id>:<next_hop>:<route_source>:<route_dest>
 * ```
 */
export class PbrRouteEntry extends pulumi.CustomResource {
    /**
     * Get an existing PbrRouteEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PbrRouteEntryState, opts?: pulumi.CustomResourceOptions): PbrRouteEntry {
        return new PbrRouteEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/pbrRouteEntry:PbrRouteEntry';

    /**
     * Returns true if the given object is an instance of PbrRouteEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PbrRouteEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PbrRouteEntry.__pulumiType;
    }

    /**
     * The next hop of the policy-based route.
     */
    public readonly nextHop!: pulumi.Output<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    public readonly publishVpc!: pulumi.Output<boolean>;
    /**
     * The destination CIDR block of the policy-based route.
     */
    public readonly routeDest!: pulumi.Output<string>;
    /**
     * The source CIDR block of the policy-based route.
     */
    public readonly routeSource!: pulumi.Output<string>;
    /**
     * The status of the vpn pbr route entry.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the vpn gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string>;
    /**
     * The weight of the policy-based route. Valid values: 0 and 100.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a PbrRouteEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PbrRouteEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PbrRouteEntryArgs | PbrRouteEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PbrRouteEntryState | undefined;
            resourceInputs["nextHop"] = state ? state.nextHop : undefined;
            resourceInputs["publishVpc"] = state ? state.publishVpc : undefined;
            resourceInputs["routeDest"] = state ? state.routeDest : undefined;
            resourceInputs["routeSource"] = state ? state.routeSource : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as PbrRouteEntryArgs | undefined;
            if ((!args || args.nextHop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHop'");
            }
            if ((!args || args.publishVpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publishVpc'");
            }
            if ((!args || args.routeDest === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeDest'");
            }
            if ((!args || args.routeSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeSource'");
            }
            if ((!args || args.vpnGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnGatewayId'");
            }
            if ((!args || args.weight === undefined) && !opts.urn) {
                throw new Error("Missing required property 'weight'");
            }
            resourceInputs["nextHop"] = args ? args.nextHop : undefined;
            resourceInputs["publishVpc"] = args ? args.publishVpc : undefined;
            resourceInputs["routeDest"] = args ? args.routeDest : undefined;
            resourceInputs["routeSource"] = args ? args.routeSource : undefined;
            resourceInputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PbrRouteEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PbrRouteEntry resources.
 */
export interface PbrRouteEntryState {
    /**
     * The next hop of the policy-based route.
     */
    nextHop?: pulumi.Input<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    publishVpc?: pulumi.Input<boolean>;
    /**
     * The destination CIDR block of the policy-based route.
     */
    routeDest?: pulumi.Input<string>;
    /**
     * The source CIDR block of the policy-based route.
     */
    routeSource?: pulumi.Input<string>;
    /**
     * The status of the vpn pbr route entry.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the vpn gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
    /**
     * The weight of the policy-based route. Valid values: 0 and 100.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PbrRouteEntry resource.
 */
export interface PbrRouteEntryArgs {
    /**
     * The next hop of the policy-based route.
     */
    nextHop: pulumi.Input<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    publishVpc: pulumi.Input<boolean>;
    /**
     * The destination CIDR block of the policy-based route.
     */
    routeDest: pulumi.Input<string>;
    /**
     * The source CIDR block of the policy-based route.
     */
    routeSource: pulumi.Input<string>;
    /**
     * The ID of the vpn gateway.
     */
    vpnGatewayId: pulumi.Input<string>;
    /**
     * The weight of the policy-based route. Valid values: 0 and 100.
     */
    weight: pulumi.Input<number>;
}
