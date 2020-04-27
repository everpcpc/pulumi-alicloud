// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB endpoint resource to allocate an Internet endpoint string for PolarDB instance.
 * 
 * > **NOTE:** Available in v1.80.0+. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
 *  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "PolarDB";
 * const name = config.get("name") || "polardbconnectionbasic";
 * 
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: creation,
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultCluster = new alicloud.polardb.Cluster("default", {
 *     dbNodeClass: "polar.mysql.x4.large",
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     description: name,
 *     payType: "PostPaid",
 *     vswitchId: defaultSwitch.id,
 * });
 * const endpoint = new alicloud.PolardbEndpoints("endpoint", {
 *     dbClusterId: defaultCluster.id,
 *     endpointType: "Custom",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_endpoint.html.markdown.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:polardb/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
     */
    public readonly autoAddNewNodes!: pulumi.Output<string | undefined>;
    /**
     * The Id of cluster that can run database.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * Advanced configuration of the cluster address.
     */
    public readonly endpointConfig!: pulumi.Output<{[key: string]: any}>;
    /**
     * Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
     */
    public readonly endpointType!: pulumi.Output<string>;
    /**
     * Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
     */
    public readonly nodes!: pulumi.Output<string[]>;
    /**
     * Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
     */
    public readonly readWriteMode!: pulumi.Output<string | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["autoAddNewNodes"] = state ? state.autoAddNewNodes : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["endpointConfig"] = state ? state.endpointConfig : undefined;
            inputs["endpointType"] = state ? state.endpointType : undefined;
            inputs["nodes"] = state ? state.nodes : undefined;
            inputs["readWriteMode"] = state ? state.readWriteMode : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if (!args || args.dbClusterId === undefined) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            if (!args || args.endpointType === undefined) {
                throw new Error("Missing required property 'endpointType'");
            }
            inputs["autoAddNewNodes"] = args ? args.autoAddNewNodes : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["endpointConfig"] = args ? args.endpointConfig : undefined;
            inputs["endpointType"] = args ? args.endpointType : undefined;
            inputs["nodes"] = args ? args.nodes : undefined;
            inputs["readWriteMode"] = args ? args.readWriteMode : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
     */
    readonly autoAddNewNodes?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId?: pulumi.Input<string>;
    /**
     * Advanced configuration of the cluster address.
     */
    readonly endpointConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
     */
    readonly endpointType?: pulumi.Input<string>;
    /**
     * Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
     */
    readonly nodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
     */
    readonly readWriteMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. Default to `Disable`.
     */
    readonly autoAddNewNodes?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId: pulumi.Input<string>;
    /**
     * Advanced configuration of the cluster address.
     */
    readonly endpointConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of endpoint. Valid value: `Custom`. Currently supported only `Custom`.
     */
    readonly endpointType: pulumi.Input<string>;
    /**
     * Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
     */
    readonly nodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. Default to `ReadOnly`.
     */
    readonly readWriteMode?: pulumi.Input<string>;
}
