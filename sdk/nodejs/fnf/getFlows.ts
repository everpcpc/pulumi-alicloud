// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Fnf Flows of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.105.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.fnf.getFlows({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstFnfFlowId = example.then(example => example.flows[0].id);
 * ```
 */
export function getFlows(args?: GetFlowsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:fnf/getFlows:getFlows", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlows.
 */
export interface GetFlowsArgs {
    /**
     * A list of Flow IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Flow name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getFlows.
 */
export interface GetFlowsResult {
    readonly flows: outputs.fnf.GetFlowsFlow[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
