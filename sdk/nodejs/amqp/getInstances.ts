// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Amqp Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.128.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.amqp.getInstances({
 *     ids: [
 *         "amqp-abc12345",
 *         "amqp-abc34567",
 *     ],
 * });
 * export const amqpInstanceId1 = ids.then(ids => ids.instances[0].id);
 * const nameRegex = alicloud.amqp.getInstances({
 *     nameRegex: "^my-Instance",
 * });
 * export const amqpInstanceId2 = nameRegex.then(nameRegex => nameRegex.instances[0].id);
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:amqp/getInstances:getInstances", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    readonly enableDetails?: boolean;
    /**
     * A list of Instance IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Instance name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The status of the resource.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instances: outputs.amqp.GetInstancesInstance[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
