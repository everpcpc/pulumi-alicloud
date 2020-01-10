// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.gpdb.getInstances` data source provides a collection of AnalyticDB for PostgreSQL instances available in Alicloud account.
 * Filters support regular expression for the instance name or availability_zone.
 * 
 * > **NOTE:**  Available in 1.47.0+
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const gpdb = alicloud.gpdb.getInstances({
 *     availabilityZone: "cn-beijing-c",
 *     nameRegex: "gp-.+\\d+",
 *     outputFile: "instances.txt",
 * });
 * 
 * export const instanceId = gpdb.instances[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/gpdb_instances.html.markdown.
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> & GetInstancesResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetInstancesResult> = pulumi.runtime.invoke("alicloud:gpdb/getInstances:getInstances", {
        "availabilityZone": args.availabilityZone,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
        "vswitchId": args.vswitchId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Instance availability zone.
     */
    readonly availabilityZone?: string;
    /**
     * A list of instance IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to apply to the instance name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
    /**
     * Used to retrieve instances belong to specified `vswitch` resources.
     */
    readonly vswitchId?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * Instance availability zone.
     */
    readonly availabilityZone?: string;
    /**
     * The ids list of AnalyticDB for PostgreSQL instances.
     */
    readonly ids: string[];
    /**
     * A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:
     */
    readonly instances: outputs.gpdb.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * The names list of AnalyticDB for PostgreSQL instance.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
    readonly vswitchId?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
