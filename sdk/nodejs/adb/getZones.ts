// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides availability zones for ADB that can be accessed by an Alibaba Cloud account within the region configured in the provider.
 * 
 * > **NOTE:** Available in v1.73.0+.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * // Declare the data source
 * const zonesIds = alicloud.adb.getZones();
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/adb_zones.html.markdown.
 */
export function getZones(args?: GetZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetZonesResult> & GetZonesResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetZonesResult> = pulumi.runtime.invoke("alicloud:adb/getZones:getZones", {
        "multi": args.multi,
        "outputFile": args.outputFile,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesArgs {
    /**
     * Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch ADB instances.
     */
    readonly multi?: boolean;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getZones.
 */
export interface GetZonesResult {
    /**
     * A list of zone IDs.
     */
    readonly ids: string[];
    readonly multi?: boolean;
    readonly outputFile?: string;
    /**
     * A list of availability zones. Each element contains the following attributes:
     */
    readonly zones: outputs.adb.GetZonesZone[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
