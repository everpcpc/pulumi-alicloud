// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the disks of the current Alibaba Cloud user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const disksDs = pulumi.output(alicloud.ecs.getDisks({
 *     nameRegex: "sampleDisk",
 * }, { async: true }));
 * 
 * export const firstDiskId = disksDs.disks[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/disks.html.markdown.
 */
export function getDisks(args?: GetDisksArgs, opts?: pulumi.InvokeOptions): Promise<GetDisksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ecs/getDisks:getDisks", {
        "category": args.category,
        "encrypted": args.encrypted,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getDisks.
 */
export interface GetDisksArgs {
    /**
     * Disk category. Possible values: `cloud` (basic cloud disk), `cloudEfficiency` (ultra cloud disk), `ephemeralSsd` (local SSD cloud disk), `cloudSsd` (SSD cloud disk), and `cloudEssd` (ESSD cloud disk).
     */
    readonly category?: string;
    /**
     * Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
     */
    readonly encrypted?: string;
    /**
     * A list of disks IDs.
     */
    readonly ids?: string[];
    /**
     * Filter the results by the specified ECS instance ID.
     */
    readonly instanceId?: string;
    /**
     * A regex string to filter results by disk name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group which the disk belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * A map of tags assigned to the disks. It must be in the format:
     * ```
     * data "alicloud.ecs.getDisks" "disksDs" {
     * tags = {
     * tagKey1 = "tagValue1",
     * tagKey2 = "tagValue2"
     * }
     * }
     * ```
     */
    readonly tags?: {[key: string]: any};
    /**
     * Disk type. Possible values: `system` and `data`.
     */
    readonly type?: string;
}

/**
 * A collection of values returned by getDisks.
 */
export interface GetDisksResult {
    /**
     * Disk category. Possible values: `cloud` (basic cloud disk), `cloudEfficiency` (ultra cloud disk), `ephemeralSsd` (local SSD cloud disk), `cloudSsd` (SSD cloud disk), and `cloudEssd` (ESSD cloud disk).
     */
    readonly category?: string;
    /**
     * A list of disks. Each element contains the following attributes:
     */
    readonly disks: outputs.ecs.GetDisksDisk[];
    /**
     * Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
     */
    readonly encrypted?: string;
    readonly ids: string[];
    /**
     * ID of the related instance. It is `null` unless the `status` is `In_use`.
     */
    readonly instanceId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group.
     */
    readonly resourceGroupId?: string;
    /**
     * A map of tags assigned to the disk.
     */
    readonly tags?: {[key: string]: any};
    /**
     * Disk type. Possible values: `system` and `data`.
     */
    readonly type?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
