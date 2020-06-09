// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the apps of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const dataApigatway = pulumi.output(alicloud.apigateway.getApps({
 *     outputFile: "outapps",
 * }, { async: true }));
 *
 * export const firstAppId = dataApigatway.apps[0].id;
 * ```
 */
export function getApps(args?: GetAppsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:apigateway/getApps:getApps", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getApps.
 */
export interface GetAppsArgs {
    /**
     * A list of app IDs. 
     */
    readonly ids?: string[];
    /**
     * A regex string to filter apps by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getApps.
 */
export interface GetAppsResult {
    /**
     * A list of apps. Each element contains the following attributes:
     */
    readonly apps: outputs.apigateway.GetAppsApp[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of app IDs. 
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of app names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
}
