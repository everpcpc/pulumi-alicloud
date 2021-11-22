// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides availability zones that can be accessed by an Alibaba Cloud account within the region configured in the provider.
 *
 * > **NOTE:** If one zone is sold out, it will not be exported.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const zonesDs = pulumi.output(alicloud.getZones({
 *     availableDiskCategory: "cloud_ssd",
 *     availableInstanceType: "ecs.n4.large",
 * }));
 * // Create an ECS instance with the first matched zone
 * const instance = new alicloud.ecs.Instance("instance", {
 *     availabilityZone: zonesDs.zones[0].id,
 * });
 * ```
 */
export function getZones(args?: GetZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:index/getZones:getZones", {
        "availableDiskCategory": args.availableDiskCategory,
        "availableInstanceType": args.availableInstanceType,
        "availableResourceCreation": args.availableResourceCreation,
        "availableSlbAddressIpVersion": args.availableSlbAddressIpVersion,
        "availableSlbAddressType": args.availableSlbAddressType,
        "enableDetails": args.enableDetails,
        "instanceChargeType": args.instanceChargeType,
        "multi": args.multi,
        "networkType": args.networkType,
        "outputFile": args.outputFile,
        "spotStrategy": args.spotStrategy,
    }, opts);
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesArgs {
    /**
     * Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
     */
    availableDiskCategory?: string;
    /**
     * Filter the results by a specific instance type.
     */
    availableInstanceType?: string;
    /**
     * Filter the results by a specific resource type.
     * Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
     */
    availableResourceCreation?: string;
    /**
     * Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
     */
    availableSlbAddressIpVersion?: string;
    /**
     * Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
     */
    availableSlbAddressType?: string;
    /**
     * Default to false and only output `id` in the `zones` block. Set it to true can output more details.
     */
    enableDetails?: boolean;
    /**
     * Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     */
    instanceChargeType?: string;
    /**
     * Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
     */
    multi?: boolean;
    /**
     * Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
     */
    networkType?: string;
    outputFile?: string;
    /**
     * - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
     */
    spotStrategy?: string;
}

/**
 * A collection of values returned by getZones.
 */
export interface GetZonesResult {
    readonly availableDiskCategory?: string;
    readonly availableInstanceType?: string;
    /**
     * Type of resources that can be created.
     */
    readonly availableResourceCreation?: string;
    readonly availableSlbAddressIpVersion?: string;
    readonly availableSlbAddressType?: string;
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of zone IDs.
     */
    readonly ids: string[];
    readonly instanceChargeType?: string;
    readonly multi?: boolean;
    readonly networkType?: string;
    readonly outputFile?: string;
    readonly spotStrategy?: string;
    /**
     * A list of availability zones. Each element contains the following attributes:
     */
    readonly zones: outputs.GetZonesZone[];
}

export function getZonesOutput(args?: GetZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZonesResult> {
    return pulumi.output(args).apply(a => getZones(a, opts))
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesOutputArgs {
    /**
     * Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
     */
    availableDiskCategory?: pulumi.Input<string>;
    /**
     * Filter the results by a specific instance type.
     */
    availableInstanceType?: pulumi.Input<string>;
    /**
     * Filter the results by a specific resource type.
     * Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
     */
    availableResourceCreation?: pulumi.Input<string>;
    /**
     * Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
     */
    availableSlbAddressIpVersion?: pulumi.Input<string>;
    /**
     * Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
     */
    availableSlbAddressType?: pulumi.Input<string>;
    /**
     * Default to false and only output `id` in the `zones` block. Set it to true can output more details.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
     */
    multi?: pulumi.Input<boolean>;
    /**
     * Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
     */
    networkType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
     */
    spotStrategy?: pulumi.Input<string>;
}
