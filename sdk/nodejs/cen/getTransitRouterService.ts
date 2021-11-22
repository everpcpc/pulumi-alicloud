// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can open CEN Transit Router Service automatically. If the service has been opened, it will return opened.
 *
 * For information about CEN and how to use it, see [What is CEN](https://www.alibabacloud.com/help/en/doc-detail/59870.htm).
 *
 * > **NOTE:** Available in v1.139.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const open = pulumi.output(alicloud.cen.getTransitRouterService({
 *     enable: "On",
 * }));
 * ```
 */
export function getTransitRouterService(args?: GetTransitRouterServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRouterServiceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouterService:getTransitRouterService", {
        "enable": args.enable,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouterService.
 */
export interface GetTransitRouterServiceArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     */
    enable?: string;
}

/**
 * A collection of values returned by getTransitRouterService.
 */
export interface GetTransitRouterServiceResult {
    readonly enable?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The current service enable status.
     */
    readonly status: string;
}

export function getTransitRouterServiceOutput(args?: GetTransitRouterServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRouterServiceResult> {
    return pulumi.output(args).apply(a => getTransitRouterService(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouterService.
 */
export interface GetTransitRouterServiceOutputArgs {
    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
     */
    enable?: pulumi.Input<string>;
}
