// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ADB cluster resource. A ADB cluster is an isolated database
 * environment in the cloud. A ADB cluster can contain multiple user-created
 * databases.
 *
 * > **NOTE:** Available in v1.71.0+.
 *
 * ## Example Usage
 * ### Create a ADB MySQL cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "adbClusterconfig";
 * const creation = config.get("creation") || "ADB";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultCluster = new alicloud.adb.Cluster("defaultCluster", {
 *     dbClusterVersion: "3.0",
 *     dbClusterCategory: "Cluster",
 *     dbNodeClass: "C8",
 *     dbNodeCount: 2,
 *     dbNodeStorage: 200,
 *     payType: "PostPaid",
 *     description: name,
 *     vswitchId: defaultSwitch.id,
 * });
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:adb/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * (Available in 1.93.0+) The connection string of the ADB cluster.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    public readonly dbClusterCategory!: pulumi.Output<string>;
    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     */
    public readonly dbClusterVersion!: pulumi.Output<string | undefined>;
    /**
     * The dbNodeClass of cluster node.
     */
    public readonly dbNodeClass!: pulumi.Output<string>;
    /**
     * The dbNodeCount of cluster node.
     */
    public readonly dbNodeCount!: pulumi.Output<number>;
    /**
     * The dbNodeStorage of cluster node.
     */
    public readonly dbNodeStorage!: pulumi.Output<number>;
    /**
     * The description of cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    public readonly maintainTime!: pulumi.Output<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    public readonly payType!: pulumi.Output<string | undefined>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    public readonly renewalStatus!: pulumi.Output<string | undefined>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    public readonly securityIps!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone to launch the DB cluster.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dbClusterCategory"] = state ? state.dbClusterCategory : undefined;
            inputs["dbClusterVersion"] = state ? state.dbClusterVersion : undefined;
            inputs["dbNodeClass"] = state ? state.dbNodeClass : undefined;
            inputs["dbNodeCount"] = state ? state.dbNodeCount : undefined;
            inputs["dbNodeStorage"] = state ? state.dbNodeStorage : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maintainTime"] = state ? state.maintainTime : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            inputs["securityIps"] = state ? state.securityIps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.dbClusterCategory === undefined) {
                throw new Error("Missing required property 'dbClusterCategory'");
            }
            if (!args || args.dbNodeClass === undefined) {
                throw new Error("Missing required property 'dbNodeClass'");
            }
            if (!args || args.dbNodeCount === undefined) {
                throw new Error("Missing required property 'dbNodeCount'");
            }
            if (!args || args.dbNodeStorage === undefined) {
                throw new Error("Missing required property 'dbNodeStorage'");
            }
            if (!args || args.vswitchId === undefined) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["dbClusterCategory"] = args ? args.dbClusterCategory : undefined;
            inputs["dbClusterVersion"] = args ? args.dbClusterVersion : undefined;
            inputs["dbNodeClass"] = args ? args.dbNodeClass : undefined;
            inputs["dbNodeCount"] = args ? args.dbNodeCount : undefined;
            inputs["dbNodeStorage"] = args ? args.dbNodeStorage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maintainTime"] = args ? args.maintainTime : undefined;
            inputs["payType"] = args ? args.payType : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            inputs["securityIps"] = args ? args.securityIps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["connectionString"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * (Available in 1.93.0+) The connection string of the ADB cluster.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    readonly dbClusterCategory?: pulumi.Input<string>;
    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     */
    readonly dbClusterVersion?: pulumi.Input<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    readonly dbNodeClass?: pulumi.Input<string>;
    /**
     * The dbNodeCount of cluster node.
     */
    readonly dbNodeCount?: pulumi.Input<number>;
    /**
     * The dbNodeStorage of cluster node.
     */
    readonly dbNodeStorage?: pulumi.Input<number>;
    /**
     * The description of cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    readonly renewalStatus?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB cluster.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    readonly dbClusterCategory: pulumi.Input<string>;
    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     */
    readonly dbClusterVersion?: pulumi.Input<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    readonly dbNodeClass: pulumi.Input<string>;
    /**
     * The dbNodeCount of cluster node.
     */
    readonly dbNodeCount: pulumi.Input<number>;
    /**
     * The dbNodeStorage of cluster node.
     */
    readonly dbNodeStorage: pulumi.Input<number>;
    /**
     * The description of cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    readonly renewalStatus?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    readonly vswitchId: pulumi.Input<string>;
    /**
     * The Zone to launch the DB cluster.
     */
    readonly zoneId?: pulumi.Input<string>;
}
