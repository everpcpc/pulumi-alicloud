// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getEnhancedNatAvailableZones(args?: GetEnhancedNatAvailableZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetEnhancedNatAvailableZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getEnhancedNatAvailableZones:getEnhancedNatAvailableZones", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnhancedNatAvailableZones.
 */
export interface GetEnhancedNatAvailableZonesArgs {
    outputFile?: string;
}

/**
 * A collection of values returned by getEnhancedNatAvailableZones.
 */
export interface GetEnhancedNatAvailableZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly zones: outputs.vpc.GetEnhancedNatAvailableZonesZone[];
}

export function getEnhancedNatAvailableZonesOutput(args?: GetEnhancedNatAvailableZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnhancedNatAvailableZonesResult> {
    return pulumi.output(args).apply(a => getEnhancedNatAvailableZones(a, opts))
}

/**
 * A collection of arguments for invoking getEnhancedNatAvailableZones.
 */
export interface GetEnhancedNatAvailableZonesOutputArgs {
    outputFile?: pulumi.Input<string>;
}
