// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * ADB cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:adb/cluster:Cluster example am-abc12345678
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
    public readonly computeResource!: pulumi.Output<string | undefined>;
    /**
     * (Available in 1.93.0+) The connection string of the ADB cluster.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    public readonly dbClusterCategory!: pulumi.Output<string>;
    /**
     * @deprecated It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     */
    public readonly dbClusterClass!: pulumi.Output<string | undefined>;
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
    public readonly description!: pulumi.Output<string>;
    public readonly elasticIoResource!: pulumi.Output<number | undefined>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    public readonly maintainTime!: pulumi.Output<string>;
    public readonly mode!: pulumi.Output<string>;
    public readonly modifyType!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    public readonly payType!: pulumi.Output<string>;
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    public readonly renewalStatus!: pulumi.Output<string | undefined>;
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    public readonly securityIps!: pulumi.Output<string[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["computeResource"] = state ? state.computeResource : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dbClusterCategory"] = state ? state.dbClusterCategory : undefined;
            inputs["dbClusterClass"] = state ? state.dbClusterClass : undefined;
            inputs["dbClusterVersion"] = state ? state.dbClusterVersion : undefined;
            inputs["dbNodeClass"] = state ? state.dbNodeClass : undefined;
            inputs["dbNodeCount"] = state ? state.dbNodeCount : undefined;
            inputs["dbNodeStorage"] = state ? state.dbNodeStorage : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["elasticIoResource"] = state ? state.elasticIoResource : undefined;
            inputs["maintainTime"] = state ? state.maintainTime : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["modifyType"] = state ? state.modifyType : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["securityIps"] = state ? state.securityIps : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.dbClusterCategory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterCategory'");
            }
            if ((!args || args.mode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mode'");
            }
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["computeResource"] = args ? args.computeResource : undefined;
            inputs["dbClusterCategory"] = args ? args.dbClusterCategory : undefined;
            inputs["dbClusterClass"] = args ? args.dbClusterClass : undefined;
            inputs["dbClusterVersion"] = args ? args.dbClusterVersion : undefined;
            inputs["dbNodeClass"] = args ? args.dbNodeClass : undefined;
            inputs["dbNodeCount"] = args ? args.dbNodeCount : undefined;
            inputs["dbNodeStorage"] = args ? args.dbNodeStorage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["elasticIoResource"] = args ? args.elasticIoResource : undefined;
            inputs["maintainTime"] = args ? args.maintainTime : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["modifyType"] = args ? args.modifyType : undefined;
            inputs["payType"] = args ? args.payType : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["securityIps"] = args ? args.securityIps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["connectionString"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
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
    autoRenewPeriod?: pulumi.Input<number>;
    computeResource?: pulumi.Input<string>;
    /**
     * (Available in 1.93.0+) The connection string of the ADB cluster.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    dbClusterCategory?: pulumi.Input<string>;
    /**
     * @deprecated It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     */
    dbClusterClass?: pulumi.Input<string>;
    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     */
    dbClusterVersion?: pulumi.Input<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    dbNodeClass?: pulumi.Input<string>;
    /**
     * The dbNodeCount of cluster node.
     */
    dbNodeCount?: pulumi.Input<number>;
    /**
     * The dbNodeStorage of cluster node.
     */
    dbNodeStorage?: pulumi.Input<number>;
    /**
     * The description of cluster.
     */
    description?: pulumi.Input<string>;
    elasticIoResource?: pulumi.Input<number>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    maintainTime?: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
    modifyType?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    payType?: pulumi.Input<string>;
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
    resourceGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB cluster.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    computeResource?: pulumi.Input<string>;
    /**
     * Cluster category. Value options: `Basic`, `Cluster`.
     */
    dbClusterCategory: pulumi.Input<string>;
    /**
     * @deprecated It duplicates with attribute db_node_class and is deprecated from 1.121.2.
     */
    dbClusterClass?: pulumi.Input<string>;
    /**
     * Cluster version. Value options: `3.0`, Default to `3.0`.
     */
    dbClusterVersion?: pulumi.Input<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    dbNodeClass?: pulumi.Input<string>;
    /**
     * The dbNodeCount of cluster node.
     */
    dbNodeCount?: pulumi.Input<number>;
    /**
     * The dbNodeStorage of cluster node.
     */
    dbNodeStorage?: pulumi.Input<number>;
    /**
     * The description of cluster.
     */
    description?: pulumi.Input<string>;
    elasticIoResource?: pulumi.Input<number>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    maintainTime?: pulumi.Input<string>;
    mode: pulumi.Input<string>;
    modifyType?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
     */
    payType?: pulumi.Input<string>;
    paymentType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
    resourceGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB cluster.
     */
    zoneId?: pulumi.Input<string>;
}
