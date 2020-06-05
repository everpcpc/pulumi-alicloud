// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Handshakes of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.86.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = pulumi.output(alicloud.resourcemanager.getHandshakes({ async: true }));
 *
 * export const firstHandshakeId = example.handshakes[0].id;
 * ```
 */
export function getHandshakes(args?: GetHandshakesArgs, opts?: pulumi.InvokeOptions): Promise<GetHandshakesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:resourcemanager/getHandshakes:getHandshakes", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getHandshakes.
 */
export interface GetHandshakesArgs {
    /**
     * A list of Resource Manager Handshake IDs.
     */
    readonly ids?: string[];
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getHandshakes.
 */
export interface GetHandshakesResult {
    /**
     * A list of Resource Manager Handshakes. Each element contains the following attributes:
     */
    readonly handshakes: outputs.resourcemanager.GetHandshakesHandshake[];
    /**
     * A list of Resource Manager Handshake IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
