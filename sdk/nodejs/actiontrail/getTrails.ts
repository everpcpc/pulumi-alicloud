// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of action trail of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const trails = pulumi.output(alicloud.actiontrail.getTrails({
 *     nameRegex: "tf-testacc-actiontrail",
 * }, { async: true }));
 *
 * export const firstTrailName = trails.actiontrails[0].name;
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
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrails.
 */
export interface GetTrailsArgs {
    /**
     * A regex string to filter results action trail name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getTrails.
 */
export interface GetTrailsResult {
    /**
     * A list of actiontrails. Each element contains the following attributes:
     */
    readonly actiontrails: outputs.actiontrail.GetTrailsActiontrail[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    /**
     * A list of trail names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}
