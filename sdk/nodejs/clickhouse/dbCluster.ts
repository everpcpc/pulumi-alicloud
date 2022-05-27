// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Click House DBCluster resource.
 *
 * For information about Click House DBCluster and how to use it, see [What is DBCluster](https://www.alibabacloud.com/product/clickhouse).
 *
 * > **NOTE:** Available in v1.134.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultRegions = alicloud.clickhouse.getRegions({
 *     current: true,
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "default-NODELETING",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultRegions]).then(([defaultNetworks, defaultRegions]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?[0],
 *     zoneId: defaultRegions.regions?[0]?.zoneIds?[0]?.zoneId,
 * }));
 * const defaultDbCluster = new alicloud.clickhouse.DbCluster("defaultDbCluster", {
 *     dbClusterVersion: "20.3.10.75",
 *     category: "Basic",
 *     dbClusterClass: "S8",
 *     dbClusterNetworkType: "vpc",
 *     dbNodeGroupCount: 1,
 *     paymentType: "PayAsYouGo",
 *     dbNodeStorage: "500",
 *     storageType: "cloud_essd",
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?[0]),
 *     dbClusterAccessWhiteLists: [{
 *         dbClusterIpArrayAttribute: "test",
 *         dbClusterIpArrayName: "test",
 *         securityIpList: "192.168.0.1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Click House DBCluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:clickhouse/dbCluster:DbCluster example <id>
 * ```
 */
export class DbCluster extends pulumi.CustomResource {
    /**
     * Get an existing DbCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbClusterState, opts?: pulumi.CustomResourceOptions): DbCluster {
        return new DbCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:clickhouse/dbCluster:DbCluster';

    /**
     * Returns true if the given object is an instance of DbCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbCluster.__pulumiType;
    }

    /**
     * The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * The db cluster access white list.
     */
    public readonly dbClusterAccessWhiteLists!: pulumi.Output<outputs.clickhouse.DbClusterDbClusterAccessWhiteList[] | undefined>;
    /**
     * The DBCluster class. According to the category, dbClusterClass has two value ranges:
     * * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
     * * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
     */
    public readonly dbClusterClass!: pulumi.Output<string>;
    /**
     * The DBCluster description.
     */
    public readonly dbClusterDescription!: pulumi.Output<string>;
    /**
     * The DBCluster network type. Valid values: `vpc`.
     */
    public readonly dbClusterNetworkType!: pulumi.Output<string>;
    /**
     * The DBCluster version. Valid values: `20.3.10.75`, `20.8.7.15`, `21.8.10.19`. **NOTE:** `19.15.2.2` is no longer supported.
     */
    public readonly dbClusterVersion!: pulumi.Output<string>;
    /**
     * The db node group count. The number should between 1 and 48.
     */
    public readonly dbNodeGroupCount!: pulumi.Output<number>;
    /**
     * The db node storage.
     */
    public readonly dbNodeStorage!: pulumi.Output<string>;
    /**
     * Key management service KMS key ID.
     */
    public readonly encryptionKey!: pulumi.Output<string | undefined>;
    /**
     * Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
     */
    public readonly encryptionType!: pulumi.Output<string | undefined>;
    /**
     * The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
     */
    public readonly maintainTime!: pulumi.Output<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * The used time of DBCluster.
     */
    public readonly usedTime!: pulumi.Output<string | undefined>;
    /**
     * The vswitch id of DBCluster.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a DbCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbClusterArgs | DbClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbClusterState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["dbClusterAccessWhiteLists"] = state ? state.dbClusterAccessWhiteLists : undefined;
            resourceInputs["dbClusterClass"] = state ? state.dbClusterClass : undefined;
            resourceInputs["dbClusterDescription"] = state ? state.dbClusterDescription : undefined;
            resourceInputs["dbClusterNetworkType"] = state ? state.dbClusterNetworkType : undefined;
            resourceInputs["dbClusterVersion"] = state ? state.dbClusterVersion : undefined;
            resourceInputs["dbNodeGroupCount"] = state ? state.dbNodeGroupCount : undefined;
            resourceInputs["dbNodeStorage"] = state ? state.dbNodeStorage : undefined;
            resourceInputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["maintainTime"] = state ? state.maintainTime : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["usedTime"] = state ? state.usedTime : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as DbClusterArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            if ((!args || args.dbClusterClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterClass'");
            }
            if ((!args || args.dbClusterNetworkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterNetworkType'");
            }
            if ((!args || args.dbClusterVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterVersion'");
            }
            if ((!args || args.dbNodeGroupCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbNodeGroupCount'");
            }
            if ((!args || args.dbNodeStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbNodeStorage'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            if ((!args || args.storageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageType'");
            }
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["dbClusterAccessWhiteLists"] = args ? args.dbClusterAccessWhiteLists : undefined;
            resourceInputs["dbClusterClass"] = args ? args.dbClusterClass : undefined;
            resourceInputs["dbClusterDescription"] = args ? args.dbClusterDescription : undefined;
            resourceInputs["dbClusterNetworkType"] = args ? args.dbClusterNetworkType : undefined;
            resourceInputs["dbClusterVersion"] = args ? args.dbClusterVersion : undefined;
            resourceInputs["dbNodeGroupCount"] = args ? args.dbNodeGroupCount : undefined;
            resourceInputs["dbNodeStorage"] = args ? args.dbNodeStorage : undefined;
            resourceInputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["maintainTime"] = args ? args.maintainTime : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["usedTime"] = args ? args.usedTime : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbCluster resources.
 */
export interface DbClusterState {
    /**
     * The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
     */
    category?: pulumi.Input<string>;
    /**
     * The db cluster access white list.
     */
    dbClusterAccessWhiteLists?: pulumi.Input<pulumi.Input<inputs.clickhouse.DbClusterDbClusterAccessWhiteList>[]>;
    /**
     * The DBCluster class. According to the category, dbClusterClass has two value ranges:
     * * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
     * * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
     */
    dbClusterClass?: pulumi.Input<string>;
    /**
     * The DBCluster description.
     */
    dbClusterDescription?: pulumi.Input<string>;
    /**
     * The DBCluster network type. Valid values: `vpc`.
     */
    dbClusterNetworkType?: pulumi.Input<string>;
    /**
     * The DBCluster version. Valid values: `20.3.10.75`, `20.8.7.15`, `21.8.10.19`. **NOTE:** `19.15.2.2` is no longer supported.
     */
    dbClusterVersion?: pulumi.Input<string>;
    /**
     * The db node group count. The number should between 1 and 48.
     */
    dbNodeGroupCount?: pulumi.Input<number>;
    /**
     * The db node storage.
     */
    dbNodeStorage?: pulumi.Input<string>;
    /**
     * Key management service KMS key ID.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
     */
    maintainTime?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
     */
    period?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     */
    status?: pulumi.Input<string>;
    /**
     * Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The used time of DBCluster.
     */
    usedTime?: pulumi.Input<string>;
    /**
     * The vswitch id of DBCluster.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DbCluster resource.
 */
export interface DbClusterArgs {
    /**
     * The Category of DBCluster. Valid values: `Basic`,`HighAvailability`.
     */
    category: pulumi.Input<string>;
    /**
     * The db cluster access white list.
     */
    dbClusterAccessWhiteLists?: pulumi.Input<pulumi.Input<inputs.clickhouse.DbClusterDbClusterAccessWhiteList>[]>;
    /**
     * The DBCluster class. According to the category, dbClusterClass has two value ranges:
     * * Under the condition that the category is the `Basic`, Valid values: `S4-NEW`, `S8`, `S16`, `S32`, `S64`, `S104`.
     * * Under the condition that the category is the `HighAvailability`, Valid values: `C4-NEW`, `C8`, `C16`, `C32`, `C64`, `C104`.
     */
    dbClusterClass: pulumi.Input<string>;
    /**
     * The DBCluster description.
     */
    dbClusterDescription?: pulumi.Input<string>;
    /**
     * The DBCluster network type. Valid values: `vpc`.
     */
    dbClusterNetworkType: pulumi.Input<string>;
    /**
     * The DBCluster version. Valid values: `20.3.10.75`, `20.8.7.15`, `21.8.10.19`. **NOTE:** `19.15.2.2` is no longer supported.
     */
    dbClusterVersion: pulumi.Input<string>;
    /**
     * The db node group count. The number should between 1 and 48.
     */
    dbNodeGroupCount: pulumi.Input<number>;
    /**
     * The db node storage.
     */
    dbNodeStorage: pulumi.Input<string>;
    /**
     * Key management service KMS key ID.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * Currently only supports ECS disk encryption, with a value of CloudDisk, not encrypted when empty.
     */
    encryptionType?: pulumi.Input<string>;
    /**
     * The maintenance window of DBCluster. Valid format: `hh:mmZ-hh:mm Z`.
     */
    maintainTime?: pulumi.Input<string>;
    /**
     * The payment type of the resource. Valid values: `PayAsYouGo`,`Subscription`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * Pre-paid cluster of the pay-as-you-go cycle. Valid values: `Month`, `Year`.
     */
    period?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Running`,`Creating`,`Deleting`,`Restarting`,`Preparing`.
     */
    status?: pulumi.Input<string>;
    /**
     * Storage type of DBCluster. Valid values: `cloudEssd`, `cloudEfficiency`, `cloudEssdPl2`, `cloudEssdPl3`.
     */
    storageType: pulumi.Input<string>;
    /**
     * The used time of DBCluster.
     */
    usedTime?: pulumi.Input<string>;
    /**
     * The vswitch id of DBCluster.
     */
    vswitchId?: pulumi.Input<string>;
}
