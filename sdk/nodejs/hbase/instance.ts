// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a HBase instance resource supports replica set instances only. the HBase provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://help.aliyun.com/product/49055.html)
 *
 * > **NOTE:**  Available in 1.67.0+
 *
 * > **NOTE:**  The following regions don't support create Classic network HBase instance.
 * [`cn-hangzhou`,`cn-shanghai`,`cn-qingdao`,`cn-beijing`,`cn-shenzhen`,`ap-southeast-1a`,.....]
 * the official website mark  more regions. or you can call [DescribeRegions](https://help.aliyun.com/document_detail/144489.html)
 *
 * > **NOTE:**  Create HBase instance or change instance type and storage would cost 15 minutes. Please make full preparation
 *
 * ## Example Usage
 * ### Create a hbase instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.hbase.Instance("default", {
 *     coldStorageSize: 0,
 *     coreDiskSize: 400,
 *     coreDiskType: "cloud_efficiency",
 *     coreInstanceQuantity: 2,
 *     coreInstanceType: "hbase.sn1.large",
 *     engineVersion: "2.0",
 *     masterInstanceType: "hbase.sn1.large",
 *     payType: "PostPaid",
 *     zoneId: "cn-shenzhen-b",
 * });
 * ```
 *
 * this is a example for class netType instance. you can find more detail with the examples/hbase dir.
 *
 * ## Import
 *
 * HBase can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:hbase/instance:Instance example hb-wz96815u13k659fvd
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
    public static readonly __pulumiType = 'alicloud:hbase/instance:Instance';

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
     * The account of the cluster web ui.
     */
    public readonly account!: pulumi.Output<string | undefined>;
    /**
     * `true`, `false`, System default to `false`, valid when payType = PrePaid.
     */
    public readonly autoRenew!: pulumi.Output<boolean>;
    /**
     * 0 or 0+. 0 means isColdStorage = false. 0+ means isColdStorage = true
     */
    public readonly coldStorageSize!: pulumi.Output<number | undefined>;
    /**
     * User-defined HBase instance one core node's storage, Valid when engine=hbase/hbaseue, bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space; value range: [400, 8000]
     * - 40-GB increments.
     */
    public readonly coreDiskSize!: pulumi.Output<number | undefined>;
    /**
     * Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`. localDisk size is fixed.
     */
    public readonly coreDiskType!: pulumi.Output<string>;
    /**
     * Default=2. if coreInstanceQuantity > 1,this is cluster's instance.  if coreInstanceQuantity = 1,this is a single instance.
     */
    public readonly coreInstanceQuantity!: pulumi.Output<number | undefined>;
    public readonly coreInstanceType!: pulumi.Output<string>;
    /**
     * The switch of delete protection. true: delete protect, false: no delete protect. you must set false when you want to delete cluster.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60, valid when payType = PrePaid. unit: month.
     */
    public readonly duration!: pulumi.Output<number>;
    /**
     * "hbase/hbaseue/bds", The following types are supported after v1.73.0: `hbaseue` and `bds `
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The white ip list of the cluster.
     */
    public readonly ipWhite!: pulumi.Output<string>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    public readonly maintainEndTime!: pulumi.Output<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    public readonly maintainStartTime!: pulumi.Output<string>;
    /**
     * Count nodes of the master node.
     * * `masterInstanceType`、`coreInstanceType` - (Required, ForceNew) Instance specification. see [Instance specifications](https://help.aliyun.com/document_detail/53532.html). or you can call describeInstanceType api.
     */
    public /*out*/ readonly masterInstanceQuantity!: pulumi.Output<number>;
    public readonly masterInstanceType!: pulumi.Output<string>;
    /**
     * HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password of the cluster web ui account.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    public readonly payType!: pulumi.Output<string | undefined>;
    /**
     * The security group resource of the cluster.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * (Optional, Available in 1.105.0+) The slb service addresses of the cluster.
     */
    public /*out*/ readonly slbConnAddrs!: pulumi.Output<outputs.hbase.InstanceSlbConnAddr[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * (Optional, Available in 1.105.0+) The Web UI proxy addresses of the cluster.
     */
    public /*out*/ readonly uiProxyConnAddrs!: pulumi.Output<outputs.hbase.InstanceUiProxyConnAddr[]>;
    /**
     * If vswitchId is not empty, that mean netType = vpc and has a same region. if vswitchId is empty, net_type_classic
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * (Optional, Available in 1.105.0+) The zookeeper addresses of the cluster.
     */
    public /*out*/ readonly zkConnAddrs!: pulumi.Output<outputs.hbase.InstanceZkConnAddr[]>;
    /**
     * The Zone to launch the HBase instance. if vswitchId is not empty, this zoneId can be "" or consistent.
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
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["account"] = state ? state.account : undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["coldStorageSize"] = state ? state.coldStorageSize : undefined;
            inputs["coreDiskSize"] = state ? state.coreDiskSize : undefined;
            inputs["coreDiskType"] = state ? state.coreDiskType : undefined;
            inputs["coreInstanceQuantity"] = state ? state.coreInstanceQuantity : undefined;
            inputs["coreInstanceType"] = state ? state.coreInstanceType : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["duration"] = state ? state.duration : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["ipWhite"] = state ? state.ipWhite : undefined;
            inputs["maintainEndTime"] = state ? state.maintainEndTime : undefined;
            inputs["maintainStartTime"] = state ? state.maintainStartTime : undefined;
            inputs["masterInstanceQuantity"] = state ? state.masterInstanceQuantity : undefined;
            inputs["masterInstanceType"] = state ? state.masterInstanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["payType"] = state ? state.payType : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["slbConnAddrs"] = state ? state.slbConnAddrs : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uiProxyConnAddrs"] = state ? state.uiProxyConnAddrs : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zkConnAddrs"] = state ? state.zkConnAddrs : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.coreDiskType === undefined) {
                throw new Error("Missing required property 'coreDiskType'");
            }
            if (!args || args.coreInstanceType === undefined) {
                throw new Error("Missing required property 'coreInstanceType'");
            }
            if (!args || args.engineVersion === undefined) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if (!args || args.masterInstanceType === undefined) {
                throw new Error("Missing required property 'masterInstanceType'");
            }
            inputs["account"] = args ? args.account : undefined;
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["coldStorageSize"] = args ? args.coldStorageSize : undefined;
            inputs["coreDiskSize"] = args ? args.coreDiskSize : undefined;
            inputs["coreDiskType"] = args ? args.coreDiskType : undefined;
            inputs["coreInstanceQuantity"] = args ? args.coreInstanceQuantity : undefined;
            inputs["coreInstanceType"] = args ? args.coreInstanceType : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["duration"] = args ? args.duration : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["ipWhite"] = args ? args.ipWhite : undefined;
            inputs["maintainEndTime"] = args ? args.maintainEndTime : undefined;
            inputs["maintainStartTime"] = args ? args.maintainStartTime : undefined;
            inputs["masterInstanceType"] = args ? args.masterInstanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["payType"] = args ? args.payType : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["masterInstanceQuantity"] = undefined /*out*/;
            inputs["slbConnAddrs"] = undefined /*out*/;
            inputs["uiProxyConnAddrs"] = undefined /*out*/;
            inputs["zkConnAddrs"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The account of the cluster web ui.
     */
    readonly account?: pulumi.Input<string>;
    /**
     * `true`, `false`, System default to `false`, valid when payType = PrePaid.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * 0 or 0+. 0 means isColdStorage = false. 0+ means isColdStorage = true
     */
    readonly coldStorageSize?: pulumi.Input<number>;
    /**
     * User-defined HBase instance one core node's storage, Valid when engine=hbase/hbaseue, bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space; value range: [400, 8000]
     * - 40-GB increments.
     */
    readonly coreDiskSize?: pulumi.Input<number>;
    /**
     * Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`. localDisk size is fixed.
     */
    readonly coreDiskType?: pulumi.Input<string>;
    /**
     * Default=2. if coreInstanceQuantity > 1,this is cluster's instance.  if coreInstanceQuantity = 1,this is a single instance.
     */
    readonly coreInstanceQuantity?: pulumi.Input<number>;
    readonly coreInstanceType?: pulumi.Input<string>;
    /**
     * The switch of delete protection. true: delete protect, false: no delete protect. you must set false when you want to delete cluster.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60, valid when payType = PrePaid. unit: month.
     */
    readonly duration?: pulumi.Input<number>;
    /**
     * "hbase/hbaseue/bds", The following types are supported after v1.73.0: `hbaseue` and `bds `
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The white ip list of the cluster.
     */
    readonly ipWhite?: pulumi.Input<string>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainStartTime?: pulumi.Input<string>;
    /**
     * Count nodes of the master node.
     * * `masterInstanceType`、`coreInstanceType` - (Required, ForceNew) Instance specification. see [Instance specifications](https://help.aliyun.com/document_detail/53532.html). or you can call describeInstanceType api.
     */
    readonly masterInstanceQuantity?: pulumi.Input<number>;
    readonly masterInstanceType?: pulumi.Input<string>;
    /**
     * HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the cluster web ui account.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The security group resource of the cluster.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional, Available in 1.105.0+) The slb service addresses of the cluster.
     */
    readonly slbConnAddrs?: pulumi.Input<pulumi.Input<inputs.hbase.InstanceSlbConnAddr>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Optional, Available in 1.105.0+) The Web UI proxy addresses of the cluster.
     */
    readonly uiProxyConnAddrs?: pulumi.Input<pulumi.Input<inputs.hbase.InstanceUiProxyConnAddr>[]>;
    /**
     * If vswitchId is not empty, that mean netType = vpc and has a same region. if vswitchId is empty, net_type_classic
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * (Optional, Available in 1.105.0+) The zookeeper addresses of the cluster.
     */
    readonly zkConnAddrs?: pulumi.Input<pulumi.Input<inputs.hbase.InstanceZkConnAddr>[]>;
    /**
     * The Zone to launch the HBase instance. if vswitchId is not empty, this zoneId can be "" or consistent.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The account of the cluster web ui.
     */
    readonly account?: pulumi.Input<string>;
    /**
     * `true`, `false`, System default to `false`, valid when payType = PrePaid.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * 0 or 0+. 0 means isColdStorage = false. 0+ means isColdStorage = true
     */
    readonly coldStorageSize?: pulumi.Input<number>;
    /**
     * User-defined HBase instance one core node's storage, Valid when engine=hbase/hbaseue, bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space; value range: [400, 8000]
     * - 40-GB increments.
     */
    readonly coreDiskSize?: pulumi.Input<number>;
    /**
     * Valid values are `cloudSsd`, `cloudEfficiency`, `localHddPro`, `localSsdPro`. localDisk size is fixed.
     */
    readonly coreDiskType: pulumi.Input<string>;
    /**
     * Default=2. if coreInstanceQuantity > 1,this is cluster's instance.  if coreInstanceQuantity = 1,this is a single instance.
     */
    readonly coreInstanceQuantity?: pulumi.Input<number>;
    readonly coreInstanceType: pulumi.Input<string>;
    /**
     * The switch of delete protection. true: delete protect, false: no delete protect. you must set false when you want to delete cluster.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60, valid when payType = PrePaid. unit: month.
     */
    readonly duration?: pulumi.Input<number>;
    /**
     * "hbase/hbaseue/bds", The following types are supported after v1.73.0: `hbaseue` and `bds `
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
     */
    readonly engineVersion: pulumi.Input<string>;
    /**
     * The white ip list of the cluster.
     */
    readonly ipWhite?: pulumi.Input<string>;
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainEndTime?: pulumi.Input<string>;
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
     */
    readonly maintainStartTime?: pulumi.Input<string>;
    readonly masterInstanceType: pulumi.Input<string>;
    /**
     * HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the cluster web ui account.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     */
    readonly payType?: pulumi.Input<string>;
    /**
     * The security group resource of the cluster.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * If vswitchId is not empty, that mean netType = vpc and has a same region. if vswitchId is empty, net_type_classic
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the HBase instance. if vswitchId is not empty, this zoneId can be "" or consistent.
     */
    readonly zoneId?: pulumi.Input<string>;
}
