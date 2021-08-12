// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Event Bridge Event Sources of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.eventbridge.getEventSources({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstEventBridgeEventSourceId = example.then(example => example.sources[0].id);
 * ```
 */
export function getEventSources(args?: GetEventSourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetEventSourcesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:eventbridge/getEventSources:getEventSources", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getEventSources.
 */
export interface GetEventSourcesArgs {
    /**
     * A list of Event Source IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Event Source name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getEventSources.
 */
export interface GetEventSourcesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly sources: outputs.eventbridge.GetEventSourcesSource[];
}
