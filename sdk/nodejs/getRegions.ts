// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This data source provides Alibaba Cloud regions.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const currentRegionDs = pulumi.output(alicloud.getRegions({
 *     current: true,
 * }, { async: true }));
 * 
 * export const currentRegionId = currentRegionDs.regions[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/regions.html.markdown.
 */
export function getRegions(args?: GetRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:index/getRegions:getRegions", {
        "current": args.current,
        "name": args.name,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsArgs {
    /**
     * Set to true to match only the region configured in the provider.
     */
    readonly current?: boolean;
    /**
     * The name of the region to select, such as `eu-central-1`.
     */
    readonly name?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getRegions.
 */
export interface GetRegionsResult {
    readonly current: boolean;
    /**
     * A list of region IDs.
     */
    readonly ids: string[];
    readonly name: string;
    readonly outputFile?: string;
    /**
     * A list of regions. Each element contains the following attributes:
     */
    readonly regions: outputs.GetRegionsRegion[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
