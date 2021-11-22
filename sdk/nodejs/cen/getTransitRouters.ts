// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Transit Routers available to the user.[What is Cen Transit Routers](https://help.aliyun.com/document_detail/261219.html)
 *
 * > **NOTE:** Available in 1.126.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cen.getTransitRouters({
 *     cenId: "cen-id1",
 * });
 * export const firstTransitRoutersType = _default.then(_default => _default.transitRouters?[0]?.type);
 * ```
 */
export function getTransitRouters(args: GetTransitRoutersArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRoutersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouters:getTransitRouters", {
        "cenId": args.cenId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "transitRouterId": args.transitRouterId,
        "transitRouterIds": args.transitRouterIds,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouters.
 */
export interface GetTransitRoutersArgs {
    /**
     * ID of the CEN instance.
     */
    cenId: string;
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of the resource. Valid values `Active`, `Creating`, `Deleting` and `Updating`.
     */
    status?: string;
    /**
     * ID of the transit router.
     */
    transitRouterId: string;
    /**
     * A list of ID of the transit router.
     */
    transitRouterIds?: string[];
}

/**
 * A collection of values returned by getTransitRouters.
 */
export interface GetTransitRoutersResult {
    /**
     * ID of the CEN instance.
     */
    readonly cenId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of the transit router attachment.
     */
    readonly status?: string;
    /**
     * ID of the transit router.
     */
    readonly transitRouterId: string;
    readonly transitRouterIds?: string[];
    /**
     * A list of CEN Transit Routers. Each element contains the following attributes:
     */
    readonly transitRouters: outputs.cen.GetTransitRoutersTransitRouter[];
}

export function getTransitRoutersOutput(args: GetTransitRoutersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRoutersResult> {
    return pulumi.output(args).apply(a => getTransitRouters(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouters.
 */
export interface GetTransitRoutersOutputArgs {
    /**
     * ID of the CEN instance.
     */
    cenId: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values `Active`, `Creating`, `Deleting` and `Updating`.
     */
    status?: pulumi.Input<string>;
    /**
     * ID of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * A list of ID of the transit router.
     */
    transitRouterIds?: pulumi.Input<pulumi.Input<string>[]>;
}
