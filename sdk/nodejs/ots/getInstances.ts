// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ots instances of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instancesDs = pulumi.output(alicloud.ots.getInstances({
 *     nameRegex: "sample-instance",
 *     outputFile: "instances.txt",
 * }, { async: true }));
 *
 * export const firstInstanceId = instancesDs.instances[0].id;
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
    return pulumi.runtime.invoke("alicloud:ots/getInstances:getInstances", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * A list of instance IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by instance name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * A map of tags assigned to the instance. It must be in the format:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as alicloud from "@pulumi/alicloud";
     *
     * const instancesDs = pulumi.output(alicloud.ots.getInstances({
     *     tags: {
     *         tagKey1: "tagValue1",
     *         tagKey2: "tagValue2",
     *     },
     * }, { async: true }));
     * ```
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instance IDs.
     */
    readonly ids: string[];
    /**
     * A list of instances. Each element contains the following attributes:
     */
    readonly instances: outputs.ots.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * A list of instance names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The tags of the instance.
     */
    readonly tags?: {[key: string]: any};
}
