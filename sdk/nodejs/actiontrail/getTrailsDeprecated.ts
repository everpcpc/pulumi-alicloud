// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/** @deprecated DataSource has been renamed to `getTrails` */
export function getTrailsDeprecated(args?: GetTrailsDeprecatedArgs, opts?: pulumi.InvokeOptions): Promise<GetTrailsDeprecatedResult> {
    pulumi.log.warn("getTrailsDeprecated is deprecated: DataSource has been renamed to `getTrails`")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:actiontrail/getTrailsDeprecated:getTrailsDeprecated", {
        "ids": args.ids,
        "includeShadowTrails": args.includeShadowTrails,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrailsDeprecated.
 */
export interface GetTrailsDeprecatedArgs {
    readonly ids?: string[];
    readonly includeShadowTrails?: boolean;
    /**
     * A regex string to filter results action trail name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly status?: string;
}

/**
 * A collection of values returned by getTrailsDeprecated.
 */
export interface GetTrailsDeprecatedResult {
    /**
     * A list of actiontrails. Each element contains the following attributes:
     *
     * @deprecated Field 'actiontrails' has been deprecated from version 1.95.0. Use 'trails' instead.
     */
    readonly actiontrails: outputs.actiontrail.GetTrailsDeprecatedActiontrail[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly includeShadowTrails?: boolean;
    readonly nameRegex?: string;
    /**
     * A list of trail names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
    readonly trails: outputs.actiontrail.GetTrailsDeprecatedTrail[];
}
