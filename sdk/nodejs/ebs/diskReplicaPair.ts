// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Ebs Disk Replica Pair resource.
 *
 * For information about Ebs Disk Replica Pair and how to use it, see [What is Disk Replica Pair](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/CreateDiskReplicaPair).
 *
 * > **NOTE:** Available in v1.196.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultEcsDisk = new alicloud.ecs.EcsDisk("defaultEcsDisk", {
 *     zoneId: "cn-hangzhou-onebox-nebula",
 *     category: "cloud_essd",
 *     deleteAutoSnapshot: true,
 *     deleteWithInstance: true,
 *     description: "Test For Terraform",
 *     diskName: _var.name,
 *     enableAutoSnapshot: true,
 *     encrypted: true,
 *     size: 500,
 *     tags: {
 *         Created: "TF",
 *         Environment: "Acceptance-test",
 *     },
 * });
 * const defaultone = new alicloud.ecs.EcsDisk("defaultone", {
 *     zoneId: "cn-hangzhou-onebox-nebula-b",
 *     category: "cloud_essd",
 *     deleteAutoSnapshot: true,
 *     deleteWithInstance: true,
 *     description: "Test For Terraform",
 *     diskName: _var.name,
 *     enableAutoSnapshot: true,
 *     encrypted: true,
 *     size: 500,
 *     tags: {
 *         Created: "TF",
 *         Environment: "Acceptance-test",
 *     },
 * });
 * const defaultDiskReplicaPair = new alicloud.ebs.DiskReplicaPair("defaultDiskReplicaPair", {
 *     destinationDiskId: defaultEcsDisk.id,
 *     destinationRegionId: "cn-hangzhou-onebox-nebula",
 *     bandwidth: "10240",
 *     destinationZoneId: "cn-hangzhou-onebox-nebula-e",
 *     sourceZoneId: "cn-hangzhou-onebox-nebula-b",
 *     diskId: defaultone.id,
 *     description: "abc",
 * });
 * ```
 *
 * ## Import
 *
 * Ebs Disk Replica Pair can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ebs/diskReplicaPair:DiskReplicaPair example <id>
 * ```
 */
export class DiskReplicaPair extends pulumi.CustomResource {
    /**
     * Get an existing DiskReplicaPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskReplicaPairState, opts?: pulumi.CustomResourceOptions): DiskReplicaPair {
        return new DiskReplicaPair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ebs/diskReplicaPair:DiskReplicaPair';

    /**
     * Returns true if the given object is an instance of DiskReplicaPair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskReplicaPair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskReplicaPair.__pulumiType;
    }

    /**
     * The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * The creation time of the resource
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the standby disk.
     */
    public readonly destinationDiskId!: pulumi.Output<string>;
    /**
     * The ID of the region to which the disaster recovery site belongs.
     */
    public readonly destinationRegionId!: pulumi.Output<string>;
    /**
     * The ID of the zone to which the disaster recovery site belongs.
     */
    public readonly destinationZoneId!: pulumi.Output<string>;
    /**
     * The ID of the primary disk.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
     */
    public readonly pairName!: pulumi.Output<string | undefined>;
    /**
     * The payment type of the resource
     */
    public readonly paymentType!: pulumi.Output<string | undefined>;
    /**
     * The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The first ID of the resource
     */
    public readonly replicaPairId!: pulumi.Output<string>;
    /**
     * The ID of the resource group
     */
    public /*out*/ readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
     */
    public readonly rpo!: pulumi.Output<string>;
    /**
     * The ID of the zone to which the production site belongs.
     */
    public readonly sourceZoneId!: pulumi.Output<string>;
    /**
     * The status of the resource
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a DiskReplicaPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskReplicaPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskReplicaPairArgs | DiskReplicaPairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskReplicaPairState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationDiskId"] = state ? state.destinationDiskId : undefined;
            resourceInputs["destinationRegionId"] = state ? state.destinationRegionId : undefined;
            resourceInputs["destinationZoneId"] = state ? state.destinationZoneId : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["pairName"] = state ? state.pairName : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["replicaPairId"] = state ? state.replicaPairId : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["rpo"] = state ? state.rpo : undefined;
            resourceInputs["sourceZoneId"] = state ? state.sourceZoneId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DiskReplicaPairArgs | undefined;
            if ((!args || args.destinationDiskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationDiskId'");
            }
            if ((!args || args.destinationRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationRegionId'");
            }
            if ((!args || args.destinationZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationZoneId'");
            }
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.sourceZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceZoneId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationDiskId"] = args ? args.destinationDiskId : undefined;
            resourceInputs["destinationRegionId"] = args ? args.destinationRegionId : undefined;
            resourceInputs["destinationZoneId"] = args ? args.destinationZoneId : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["pairName"] = args ? args.pairName : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["replicaPairId"] = args ? args.replicaPairId : undefined;
            resourceInputs["rpo"] = args ? args.rpo : undefined;
            resourceInputs["sourceZoneId"] = args ? args.sourceZoneId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["resourceGroupId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DiskReplicaPair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskReplicaPair resources.
 */
export interface DiskReplicaPairState {
    /**
     * The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * The creation time of the resource
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the standby disk.
     */
    destinationDiskId?: pulumi.Input<string>;
    /**
     * The ID of the region to which the disaster recovery site belongs.
     */
    destinationRegionId?: pulumi.Input<string>;
    /**
     * The ID of the zone to which the disaster recovery site belongs.
     */
    destinationZoneId?: pulumi.Input<string>;
    /**
     * The ID of the primary disk.
     */
    diskId?: pulumi.Input<string>;
    /**
     * The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
     */
    pairName?: pulumi.Input<string>;
    /**
     * The payment type of the resource
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     */
    period?: pulumi.Input<string>;
    /**
     * The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The first ID of the resource
     */
    replicaPairId?: pulumi.Input<string>;
    /**
     * The ID of the resource group
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
     */
    rpo?: pulumi.Input<string>;
    /**
     * The ID of the zone to which the production site belongs.
     */
    sourceZoneId?: pulumi.Input<string>;
    /**
     * The status of the resource
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DiskReplicaPair resource.
 */
export interface DiskReplicaPairArgs {
    /**
     * The bandwidth for asynchronous data replication between cloud disks. The unit is Kbps. Value range:-10240 Kbps: equal to 10 Mbps.-20480 Kbps: equal to 20 Mbps.-51200 Kbps: equal to 50 Mbps.-102400 Kbps: equal to 100 Mbps.Default value: 10240.This parameter cannot be specified when the ChargeType value is POSTPAY. The system value is 0, which indicates that the disk is dynamically allocated according to data write changes during asynchronous replication.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * The description of the asynchronous replication relationship. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the standby disk.
     */
    destinationDiskId: pulumi.Input<string>;
    /**
     * The ID of the region to which the disaster recovery site belongs.
     */
    destinationRegionId: pulumi.Input<string>;
    /**
     * The ID of the zone to which the disaster recovery site belongs.
     */
    destinationZoneId: pulumi.Input<string>;
    /**
     * The ID of the primary disk.
     */
    diskId: pulumi.Input<string>;
    /**
     * The name of the asynchronous replication relationship. The length must be 2 to 128 characters in length and must start with a letter or Chinese name. It cannot start with http:// or https. It can contain Chinese, English, numbers, half-width colons (:), underscores (_), half-width periods (.), or dashes (-).
     */
    pairName?: pulumi.Input<string>;
    /**
     * The payment type of the resource
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The length of the purchase for the asynchronous replication relationship. When ChargeType=PrePay, this parameter is mandatory. The unit of duration is specified by PeriodUnit and takes on a range of values. When PeriodUnit=Week, this parameter takes values in the range `1`, `2`, `3` and `4`. When PeriodUnit=Month, the parameter takes on the values `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     */
    period?: pulumi.Input<string>;
    /**
     * The units of asynchronous replication relationship purchase length. Valid values: `Week` and `Month`. Default value: `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The first ID of the resource
     */
    replicaPairId?: pulumi.Input<string>;
    /**
     * The RPO value set by the consistency group in seconds. Currently only 900 seconds are supported.
     */
    rpo?: pulumi.Input<string>;
    /**
     * The ID of the zone to which the production site belongs.
     */
    sourceZoneId: pulumi.Input<string>;
}
