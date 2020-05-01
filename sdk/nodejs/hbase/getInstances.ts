// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
 * Filters support regular expression for the instance name, ids or availability_zone.
 * 
 * > **NOTE:**  Available in 1.67.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const hbase = pulumi.output(alicloud.hbase.getInstances({
 *     availabilityZone: "cn-shenzhen-b",
 *     nameRegex: "tf_testAccHBase",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/hbase_instances.html.markdown.
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:hbase/getInstances:getInstances", {
        "availabilityZone": args.availabilityZone,
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
     * Instance availability zone.
     */
    readonly availabilityZone?: string;
    /**
     * The ids list of HBase instances
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
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly availabilityZone?: string;
    /**
     * The ids list of HBase instances
     */
    readonly ids: string[];
    /**
     * A list of HBase instances. Its every element contains the following attributes:
     */
    readonly instances: outputs.hbase.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * The names list of HBase instances
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
