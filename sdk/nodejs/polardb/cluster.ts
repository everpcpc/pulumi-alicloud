// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB cluster resource. A PolarDB cluster is an isolated database
 * environment in the cloud. A PolarDB cluster can contain multiple user-created
 * databases.
 * 
 * > **NOTE:** Available in v1.66.0+.
 * 
 * ## Example Usage
 * 
 * ### Create a PolarDB MySQL cluster
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "polardbClusterconfig";
 * const creation = config.get("creation") || "PolarDB";
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
 *     dbNodeClass: "rds.mysql.s2.large",
 *     dbType: "MySQL",
 *     dbVersion: "5.6",
 *     description: name,
 *     payType: "PostPaid",
 *     vswitchId: defaultSwitch.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_cluster.html.markdown.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:polardb/cluster:Cluster';

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
     * (Available in 1.81.0+) PolarDB cluster connection string. Only return the Primary endpoint address if securityIps is configured.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    public readonly dbNodeClass!: pulumi.Output<string>;
    /**
     * Database type. Value options: MySQL, Oracle, PostgreSQL.
     */
    public readonly dbType!: pulumi.Output<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
     */
    public readonly dbVersion!: pulumi.Output<string>;
    /**
     * The description of cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    public readonly maintainTime!: pulumi.Output<string>;
    /**
     * Use as `dbNodeClass` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
     */
    public readonly modifyType!: pulumi.Output<string | undefined>;
    /**
     * Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    public readonly parameters!: pulumi.Output<outputs.polardb.ClusterParameter[]>;
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
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The Zone to launch the DB cluster. it supports multiple zone.
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
            inputs["dbNodeClass"] = state ? state.dbNodeClass : undefined;
            inputs["dbType"] = state ? state.dbType : undefined;
            inputs["dbVersion"] = state ? state.dbVersion : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maintainTime"] = state ? state.maintainTime : undefined;
            inputs["modifyType"] = state ? state.modifyType : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            inputs["securityIps"] = state ? state.securityIps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.dbNodeClass === undefined) {
                throw new Error("Missing required property 'dbNodeClass'");
            }
            if (!args || args.dbType === undefined) {
                throw new Error("Missing required property 'dbType'");
            }
            if (!args || args.dbVersion === undefined) {
                throw new Error("Missing required property 'dbVersion'");
            }
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["dbNodeClass"] = args ? args.dbNodeClass : undefined;
            inputs["dbType"] = args ? args.dbType : undefined;
            inputs["dbVersion"] = args ? args.dbVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maintainTime"] = args ? args.maintainTime : undefined;
            inputs["modifyType"] = args ? args.modifyType : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
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
     * (Available in 1.81.0+) PolarDB cluster connection string. Only return the Primary endpoint address if securityIps is configured.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * The dbNodeClass of cluster node.
     */
    readonly dbNodeClass?: pulumi.Input<string>;
    /**
     * Database type. Value options: MySQL, Oracle, PostgreSQL.
     */
    readonly dbType?: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
     */
    readonly dbVersion?: pulumi.Input<string>;
    /**
     * The description of cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * Use as `dbNodeClass` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
     */
    readonly modifyType?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.polardb.ClusterParameter>[]>;
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
     * The Zone to launch the DB cluster. it supports multiple zone.
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
     * The dbNodeClass of cluster node.
     */
    readonly dbNodeClass: pulumi.Input<string>;
    /**
     * Database type. Value options: MySQL, Oracle, PostgreSQL.
     */
    readonly dbType: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
     */
    readonly dbVersion: pulumi.Input<string>;
    /**
     * The description of cluster.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * Use as `dbNodeClass` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
     */
    readonly modifyType?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.polardb.ClusterParameter>[]>;
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
     * The Zone to launch the DB cluster. it supports multiple zone.
     */
    readonly zoneId?: pulumi.Input<string>;
}
