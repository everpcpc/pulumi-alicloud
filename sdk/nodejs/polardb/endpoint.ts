// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB endpoint resource to manage endpoint of PolarDB cluster.
 *
 * > **NOTE:** After v1.80.0 and before v1.121.0, you can only use this resource to manage the custom endpoint. Since v1.121.0, you also can import the primary endpoint and the cluster endpoint, to modify their ssl status and so on.
 *
 * > **NOTE:** The primary endpoint and the default cluster endpoint can not be created or deleted manually.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "PolarDB";
 * const name = config.get("name") || "polardbconnectionbasic";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultCluster = new alicloud.polardb.Cluster("defaultCluster", {
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     dbNodeClass: "polar.mysql.x4.large",
 *     vswitchId: defaultSwitch.id,
 *     description: name,
 * });
 * const endpoint = new alicloud.polardb.Endpoint("endpoint", {
 *     dbClusterId: defaultCluster.id,
 *     endpointType: "Custom",
 * });
 * ```
 * ## Argument Reference
 *
 * The following arguments are supported:
 *
 * * `dbClusterId` - (Required, ForceNew) The Id of cluster that can run database.
 * * `endpointType` - (Required & ForceNew before v1.121.0, Optional in v1.121.0+) Type of the endpoint. Before v1.121.0, it only can be `Custom`. since v1.121.0, `Custom`, `Cluster`, `Primary` are valid, default to `Custom`. However when creating a new endpoint, it also only can be `Custom`.
 * * `readWriteMode` - (Optional) Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
 * * `nodes` - (Optional) Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
 * * `autoAddNewNodes` - (Optional) Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
 * * `endpointConfig` - (Optional) The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
 * * `sslEnabled` - (Optional, Available in v1.121.0+) Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
 * * `netType` - (Optional, Available in v1.121.0+) The network type of the endpoint address.\
 *     **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).\
 *     For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
 *
 * ## Import
 *
 * PolarDB endpoint can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:polardb/endpoint:Endpoint example pc-abc123456:pe-abc123456
 * ```
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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

    public readonly autoAddNewNodes!: pulumi.Output<string>;
    public readonly dbClusterId!: pulumi.Output<string>;
    public readonly endpointConfig!: pulumi.Output<{[key: string]: any}>;
    /**
     * Type of endpoint.
     */
    public readonly endpointType!: pulumi.Output<string | undefined>;
    public readonly netType!: pulumi.Output<string | undefined>;
    public readonly nodes!: pulumi.Output<string[]>;
    public readonly readWriteMode!: pulumi.Output<string>;
    /**
     * (Available in v1.121.0+) The SSL connection string.
     */
    public /*out*/ readonly sslConnectionString!: pulumi.Output<string>;
    public readonly sslEnabled!: pulumi.Output<string | undefined>;
    /**
     * (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    public /*out*/ readonly sslExpireTime!: pulumi.Output<string>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["autoAddNewNodes"] = state ? state.autoAddNewNodes : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["endpointConfig"] = state ? state.endpointConfig : undefined;
            inputs["endpointType"] = state ? state.endpointType : undefined;
            inputs["netType"] = state ? state.netType : undefined;
            inputs["nodes"] = state ? state.nodes : undefined;
            inputs["readWriteMode"] = state ? state.readWriteMode : undefined;
            inputs["sslConnectionString"] = state ? state.sslConnectionString : undefined;
            inputs["sslEnabled"] = state ? state.sslEnabled : undefined;
            inputs["sslExpireTime"] = state ? state.sslExpireTime : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if ((!args || args.dbClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            inputs["autoAddNewNodes"] = args ? args.autoAddNewNodes : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["endpointConfig"] = args ? args.endpointConfig : undefined;
            inputs["endpointType"] = args ? args.endpointType : undefined;
            inputs["netType"] = args ? args.netType : undefined;
            inputs["nodes"] = args ? args.nodes : undefined;
            inputs["readWriteMode"] = args ? args.readWriteMode : undefined;
            inputs["sslEnabled"] = args ? args.sslEnabled : undefined;
            inputs["sslConnectionString"] = undefined /*out*/;
            inputs["sslExpireTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    readonly autoAddNewNodes?: pulumi.Input<string>;
    readonly dbClusterId?: pulumi.Input<string>;
    readonly endpointConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of endpoint.
     */
    readonly endpointType?: pulumi.Input<string>;
    readonly netType?: pulumi.Input<string>;
    readonly nodes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly readWriteMode?: pulumi.Input<string>;
    /**
     * (Available in v1.121.0+) The SSL connection string.
     */
    readonly sslConnectionString?: pulumi.Input<string>;
    readonly sslEnabled?: pulumi.Input<string>;
    /**
     * (Available in v1.121.0+) The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    readonly sslExpireTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    readonly autoAddNewNodes?: pulumi.Input<string>;
    readonly dbClusterId: pulumi.Input<string>;
    readonly endpointConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of endpoint.
     */
    readonly endpointType?: pulumi.Input<string>;
    readonly netType?: pulumi.Input<string>;
    readonly nodes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly readWriteMode?: pulumi.Input<string>;
    readonly sslEnabled?: pulumi.Input<string>;
}
