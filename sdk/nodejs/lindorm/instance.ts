// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Lindorm Instance resource.
 *
 * For information about Lindorm Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/doc-detail/174640.html).
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Import
 *
 * Lindorm Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:lindorm/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:lindorm/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The cold storage capacity of the instance. Unit: GB.
     */
    public readonly coldStorage!: pulumi.Output<number>;
    /**
     * The core num.
     */
    public readonly coreNum!: pulumi.Output<number | undefined>;
    /**
     * The core spec.
     */
    public readonly coreSpec!: pulumi.Output<string | undefined>;
    /**
     * The deletion protection of instance.
     */
    public readonly deletionProection!: pulumi.Output<boolean>;
    /**
     * The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
     */
    public readonly diskCategory!: pulumi.Output<string>;
    /**
     * The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
     */
    public readonly duration!: pulumi.Output<string | undefined>;
    /**
     * The count of file engine.
     */
    public readonly fileEngineNodeCount!: pulumi.Output<number>;
    /**
     * The specification of file engine. Valid values: `lindorm.c.xlarge`.
     */
    public readonly fileEngineSpecification!: pulumi.Output<string>;
    /**
     * The group name.
     */
    public readonly groupName!: pulumi.Output<string | undefined>;
    /**
     * The name of the instance.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
     */
    public readonly instanceStorage!: pulumi.Output<string | undefined>;
    /**
     * The ip white list of instance.
     */
    public readonly ipWhiteLists!: pulumi.Output<string[] | undefined>;
    /**
     * The count of lindorm tunnel service.
     */
    public readonly ltsNodeCount!: pulumi.Output<number>;
    /**
     * The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
     */
    public readonly ltsNodeSpecification!: pulumi.Output<string>;
    /**
     * The billing method. Valid values: `PayAsYouGo` and `Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The count of phoenix.
     */
    public readonly phoenixNodeCount!: pulumi.Output<number>;
    /**
     * The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    public readonly phoenixNodeSpecification!: pulumi.Output<string>;
    /**
     * The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
     */
    public readonly pricingCycle!: pulumi.Output<string | undefined>;
    /**
     * The count of search engine.
     */
    public readonly searchEngineNodeCount!: pulumi.Output<number>;
    /**
     * The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    public readonly searchEngineSpecification!: pulumi.Output<string>;
    /**
     * The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The count of table engine.
     */
    public readonly tableEngineNodeCount!: pulumi.Output<number>;
    /**
     * The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    public readonly tableEngineSpecification!: pulumi.Output<string>;
    /**
     * The count of time series engine.
     */
    public readonly timeSeriesEngineNodeCount!: pulumi.Output<number>;
    /**
     * The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    public readonly timeSeriresEngineSpecification!: pulumi.Output<string>;
    /**
     * The upgrade type. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
     */
    public readonly upgradeType!: pulumi.Output<string | undefined>;
    /**
     * The vswitch id.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The zone ID of the instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["coldStorage"] = state ? state.coldStorage : undefined;
            inputs["coreNum"] = state ? state.coreNum : undefined;
            inputs["coreSpec"] = state ? state.coreSpec : undefined;
            inputs["deletionProection"] = state ? state.deletionProection : undefined;
            inputs["diskCategory"] = state ? state.diskCategory : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["fileEngineNodeCount"] = state ? state.fileEngineNodeCount : undefined;
            inputs["fileEngineSpecification"] = state ? state.fileEngineSpecification : undefined;
            inputs["groupName"] = state ? state.groupName : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceStorage"] = state ? state.instanceStorage : undefined;
            inputs["ipWhiteLists"] = state ? state.ipWhiteLists : undefined;
            inputs["ltsNodeCount"] = state ? state.ltsNodeCount : undefined;
            inputs["ltsNodeSpecification"] = state ? state.ltsNodeSpecification : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["phoenixNodeCount"] = state ? state.phoenixNodeCount : undefined;
            inputs["phoenixNodeSpecification"] = state ? state.phoenixNodeSpecification : undefined;
            inputs["pricingCycle"] = state ? state.pricingCycle : undefined;
            inputs["searchEngineNodeCount"] = state ? state.searchEngineNodeCount : undefined;
            inputs["searchEngineSpecification"] = state ? state.searchEngineSpecification : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tableEngineNodeCount"] = state ? state.tableEngineNodeCount : undefined;
            inputs["tableEngineSpecification"] = state ? state.tableEngineSpecification : undefined;
            inputs["timeSeriesEngineNodeCount"] = state ? state.timeSeriesEngineNodeCount : undefined;
            inputs["timeSeriresEngineSpecification"] = state ? state.timeSeriresEngineSpecification : undefined;
            inputs["upgradeType"] = state ? state.upgradeType : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.diskCategory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskCategory'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["coldStorage"] = args ? args.coldStorage : undefined;
            inputs["coreNum"] = args ? args.coreNum : undefined;
            inputs["coreSpec"] = args ? args.coreSpec : undefined;
            inputs["deletionProection"] = args ? args.deletionProection : undefined;
            inputs["diskCategory"] = args ? args.diskCategory : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["fileEngineNodeCount"] = args ? args.fileEngineNodeCount : undefined;
            inputs["fileEngineSpecification"] = args ? args.fileEngineSpecification : undefined;
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceStorage"] = args ? args.instanceStorage : undefined;
            inputs["ipWhiteLists"] = args ? args.ipWhiteLists : undefined;
            inputs["ltsNodeCount"] = args ? args.ltsNodeCount : undefined;
            inputs["ltsNodeSpecification"] = args ? args.ltsNodeSpecification : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["phoenixNodeCount"] = args ? args.phoenixNodeCount : undefined;
            inputs["phoenixNodeSpecification"] = args ? args.phoenixNodeSpecification : undefined;
            inputs["pricingCycle"] = args ? args.pricingCycle : undefined;
            inputs["searchEngineNodeCount"] = args ? args.searchEngineNodeCount : undefined;
            inputs["searchEngineSpecification"] = args ? args.searchEngineSpecification : undefined;
            inputs["tableEngineNodeCount"] = args ? args.tableEngineNodeCount : undefined;
            inputs["tableEngineSpecification"] = args ? args.tableEngineSpecification : undefined;
            inputs["timeSeriesEngineNodeCount"] = args ? args.timeSeriesEngineNodeCount : undefined;
            inputs["timeSeriresEngineSpecification"] = args ? args.timeSeriresEngineSpecification : undefined;
            inputs["upgradeType"] = args ? args.upgradeType : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The cold storage capacity of the instance. Unit: GB.
     */
    coldStorage?: pulumi.Input<number>;
    /**
     * The core num.
     */
    coreNum?: pulumi.Input<number>;
    /**
     * The core spec.
     */
    coreSpec?: pulumi.Input<string>;
    /**
     * The deletion protection of instance.
     */
    deletionProection?: pulumi.Input<boolean>;
    /**
     * The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
     */
    diskCategory?: pulumi.Input<string>;
    /**
     * The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
     */
    duration?: pulumi.Input<string>;
    /**
     * The count of file engine.
     */
    fileEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of file engine. Valid values: `lindorm.c.xlarge`.
     */
    fileEngineSpecification?: pulumi.Input<string>;
    /**
     * The group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
     */
    instanceStorage?: pulumi.Input<string>;
    /**
     * The ip white list of instance.
     */
    ipWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The count of lindorm tunnel service.
     */
    ltsNodeCount?: pulumi.Input<number>;
    /**
     * The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
     */
    ltsNodeSpecification?: pulumi.Input<string>;
    /**
     * The billing method. Valid values: `PayAsYouGo` and `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The count of phoenix.
     */
    phoenixNodeCount?: pulumi.Input<number>;
    /**
     * The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    phoenixNodeSpecification?: pulumi.Input<string>;
    /**
     * The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
     */
    pricingCycle?: pulumi.Input<string>;
    /**
     * The count of search engine.
     */
    searchEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    searchEngineSpecification?: pulumi.Input<string>;
    /**
     * The status of Instance, enumerative: Valid values: `ACTIVATION`, `DELETED`, `CREATING`, `CLASS_CHANGING`, `LOCKED`, `INSTANCE_LEVEL_MODIFY`, `NET_MODIFYING`, `RESIZING`, `RESTARTING`, `MINOR_VERSION_TRANSING`.
     */
    status?: pulumi.Input<string>;
    /**
     * The count of table engine.
     */
    tableEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    tableEngineSpecification?: pulumi.Input<string>;
    /**
     * The count of time series engine.
     */
    timeSeriesEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    timeSeriresEngineSpecification?: pulumi.Input<string>;
    /**
     * The upgrade type. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
     */
    upgradeType?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The cold storage capacity of the instance. Unit: GB.
     */
    coldStorage?: pulumi.Input<number>;
    /**
     * The core num.
     */
    coreNum?: pulumi.Input<number>;
    /**
     * The core spec.
     */
    coreSpec?: pulumi.Input<string>;
    /**
     * The deletion protection of instance.
     */
    deletionProection?: pulumi.Input<boolean>;
    /**
     * The disk type of instance. Valid values: `capacityCloudStorage`, `cloudEfficiency`, `cloudEssd`, `cloudSsd`.
     */
    diskCategory: pulumi.Input<string>;
    /**
     * The duration of paid. Valid when the `paymentType` is `Subscription`.  When `pricingCycle` set to `Month`, the valid value id `1` to `9`.  When `pricingCycle` set to `Year`, the valid value id `1` to `3`.
     */
    duration?: pulumi.Input<string>;
    /**
     * The count of file engine.
     */
    fileEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of file engine. Valid values: `lindorm.c.xlarge`.
     */
    fileEngineSpecification?: pulumi.Input<string>;
    /**
     * The group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The name of the instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
     */
    instanceStorage?: pulumi.Input<string>;
    /**
     * The ip white list of instance.
     */
    ipWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The count of lindorm tunnel service.
     */
    ltsNodeCount?: pulumi.Input<number>;
    /**
     * The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
     */
    ltsNodeSpecification?: pulumi.Input<string>;
    /**
     * The billing method. Valid values: `PayAsYouGo` and `Subscription`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * The count of phoenix.
     */
    phoenixNodeCount?: pulumi.Input<number>;
    /**
     * The specification of phoenix. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    phoenixNodeSpecification?: pulumi.Input<string>;
    /**
     * The pricing cycle. Valid when the `paymentType` is `Subscription`. Valid values: `Month` and `Year`.
     */
    pricingCycle?: pulumi.Input<string>;
    /**
     * The count of search engine.
     */
    searchEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    searchEngineSpecification?: pulumi.Input<string>;
    /**
     * The count of table engine.
     */
    tableEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.c.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    tableEngineSpecification?: pulumi.Input<string>;
    /**
     * The count of time series engine.
     */
    timeSeriesEngineNodeCount?: pulumi.Input<number>;
    /**
     * The specification of time series engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     */
    timeSeriresEngineSpecification?: pulumi.Input<string>;
    /**
     * The upgrade type. Valid values:  `open-lindorm-engine`, `open-phoenix-engine`, `open-search-engine`, `open-tsdb-engine`,  `upgrade-cold-storage`, `upgrade-disk-size`,  `upgrade-lindorm-core-num`, `upgrade-lindorm-engine`,  `upgrade-search-core-num`, `upgrade-search-engine`, `upgrade-tsdb-core-num`, `upgrade-tsdb-engine`.
     */
    upgradeType?: pulumi.Input<string>;
    /**
     * The vswitch id.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * The zone ID of the instance.
     */
    zoneId?: pulumi.Input<string>;
}
