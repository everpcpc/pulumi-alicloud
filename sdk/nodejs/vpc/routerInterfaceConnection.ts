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
 * const config = new pulumi.Config();
 * const region = config.get("region") || "cn-hangzhou";
 * const name = config.get("name") || "alicloudRouterInterfaceConnectionBasic";
 * const fooNetwork = new alicloud.vpc.Network("fooNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const barNetwork = new alicloud.vpc.Network("barNetwork", {
 *     vpcName: name,
 *     cidrBlock: "192.168.0.0/16",
 * }, {
 *     provider: alicloud,
 * });
 * const initiate = new alicloud.vpc.RouterInterface("initiate", {
 *     oppositeRegion: region,
 *     routerType: "VRouter",
 *     routerId: fooNetwork.routerId,
 *     role: "InitiatingSide",
 *     specification: "Large.2",
 *     description: name,
 *     instanceChargeType: "PostPaid",
 * });
 * const opposite = new alicloud.vpc.RouterInterface("opposite", {
 *     oppositeRegion: region,
 *     routerType: "VRouter",
 *     routerId: barNetwork.routerId,
 *     role: "AcceptingSide",
 *     specification: "Large.1",
 *     description: `${name}-opposite`,
 * }, {
 *     provider: alicloud,
 * });
 * const barRouterInterfaceConnection = new alicloud.vpc.RouterInterfaceConnection("barRouterInterfaceConnection", {
 *     interfaceId: opposite.id,
 *     oppositeInterfaceId: initiate.id,
 * }, {
 *     provider: alicloud,
 * });
 * // A integrated router interface connection tunnel requires both InitiatingSide and AcceptingSide configuring opposite router interface.
 * const fooRouterInterfaceConnection = new alicloud.vpc.RouterInterfaceConnection("fooRouterInterfaceConnection", {
 *     interfaceId: initiate.id,
 *     oppositeInterfaceId: opposite.id,
 * }, {
 *     dependsOn: [barRouterInterfaceConnection],
 * });
 * // The connection must start from the accepting side.
 * ```
 *
 * ## Import
 *
 * The router interface connection can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/routerInterfaceConnection:RouterInterfaceConnection foo ri-abc123456
 * ```
 */
export class RouterInterfaceConnection extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterfaceConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceConnectionState, opts?: pulumi.CustomResourceOptions): RouterInterfaceConnection {
        return new RouterInterfaceConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routerInterfaceConnection:RouterInterfaceConnection';

    /**
     * Returns true if the given object is an instance of RouterInterfaceConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterInterfaceConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterInterfaceConnection.__pulumiType;
    }

    /**
     * One side router interface ID.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
     */
    public readonly oppositeInterfaceId!: pulumi.Output<string>;
    public readonly oppositeInterfaceOwnerId!: pulumi.Output<string>;
    /**
     * Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    public readonly oppositeRouterId!: pulumi.Output<string>;
    /**
     * Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    public readonly oppositeRouterType!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterInterfaceConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterInterfaceConnectionArgs | RouterInterfaceConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterInterfaceConnectionState | undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["oppositeInterfaceId"] = state ? state.oppositeInterfaceId : undefined;
            resourceInputs["oppositeInterfaceOwnerId"] = state ? state.oppositeInterfaceOwnerId : undefined;
            resourceInputs["oppositeRouterId"] = state ? state.oppositeRouterId : undefined;
            resourceInputs["oppositeRouterType"] = state ? state.oppositeRouterType : undefined;
        } else {
            const args = argsOrState as RouterInterfaceConnectionArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.oppositeInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oppositeInterfaceId'");
            }
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["oppositeInterfaceId"] = args ? args.oppositeInterfaceId : undefined;
            resourceInputs["oppositeInterfaceOwnerId"] = args ? args.oppositeInterfaceOwnerId : undefined;
            resourceInputs["oppositeRouterId"] = args ? args.oppositeRouterId : undefined;
            resourceInputs["oppositeRouterType"] = args ? args.oppositeRouterType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterInterfaceConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterInterfaceConnection resources.
 */
export interface RouterInterfaceConnectionState {
    /**
     * One side router interface ID.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
     */
    oppositeInterfaceId?: pulumi.Input<string>;
    oppositeInterfaceOwnerId?: pulumi.Input<string>;
    /**
     * Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    oppositeRouterId?: pulumi.Input<string>;
    /**
     * Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    oppositeRouterType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterInterfaceConnection resource.
 */
export interface RouterInterfaceConnectionArgs {
    /**
     * One side router interface ID.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * Another side router interface ID. It must belong the specified "oppositeInterfaceOwnerId" account.
     */
    oppositeInterfaceId: pulumi.Input<string>;
    oppositeInterfaceOwnerId?: pulumi.Input<string>;
    /**
     * Another side router ID. It must belong the specified "oppositeInterfaceOwnerId" account. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    oppositeRouterId?: pulumi.Input<string>;
    /**
     * Another side router Type. Optional value: VRouter, VBR. It is valid when field "oppositeInterfaceOwnerId" is specified.
     */
    oppositeRouterType?: pulumi.Input<string>;
}
