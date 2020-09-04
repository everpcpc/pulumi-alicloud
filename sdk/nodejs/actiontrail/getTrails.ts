// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ActionTrail Trails in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.95.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.actiontrail.getTrails({
 *     nameRegex: "tf-testacc-actiontrail",
 * });
 * export const trailName = _default.then(_default => _default.trails[0].id);
 * ```
 */
export function getTrails(args?: GetTrailsArgs, opts?: pulumi.InvokeOptions): Promise<GetTrailsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:actiontrail/getTrails:getTrails", {
        "ids": args.ids,
        "includeShadowTrails": args.includeShadowTrails,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrails.
 */
export interface GetTrailsArgs {
    /**
     * A list of ActionTrail Trail IDs. It is the same as trail name.
     */
    readonly ids?: string[];
    /**
     * Whether to show shadow tracking. Default to `false`.
     */
    readonly includeShadowTrails?: boolean;
    /**
     * A regex string to filter results by trail name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Filter the results by status of the ActionTrail Trail. Valid values: `Disable`, `Enable`, `Fresh`.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getTrails.
 */
export interface GetTrailsResult {
    /**
     * Field `actiontrails` has been deprecated from version 1.95.0. Use `trails` instead."
     *
     * @deprecated Field 'actiontrails' has been deprecated from version 1.95.0. Use 'trails' instead.
     */
    readonly actiontrails: outputs.actiontrail.GetTrailsActiontrail[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of ActionTrail Trail ids. It is the same as trail name.
     */
    readonly ids: string[];
    readonly includeShadowTrails?: boolean;
    readonly nameRegex?: string;
    /**
     * A list of trail names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of the ActionTrail Trail.
     */
    readonly status?: string;
    /**
     * A list of ActionTrail Trails. Each element contains the following attributes:
     */
    readonly trails: outputs.actiontrail.GetTrailsTrail[];
}
