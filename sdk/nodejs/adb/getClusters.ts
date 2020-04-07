// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.adb.getClusters` data source provides a collection of ADB clusters available in Alibaba Cloud account.
 * Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.
 * 
 * > **NOTE:** Available in v1.71.0+.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const adbClustersDs = alicloud.adb.getClusters({
 *     descriptionRegex: "am-\\w+",
 *     status: "Running",
 * });
 * 
 * export const firstAdbClusterId = adbClustersDs.clusters[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/adb_clusters.html.markdown.
 */
export function getClusters(args?: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> & GetClustersResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClustersResult> = pulumi.runtime.invoke("alicloud:adb/getClusters:getClusters", {
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    /**
     * A regex string to filter results by cluster description.
     */
    readonly descriptionRegex?: string;
    /**
     * A list of ADB cluster IDs. 
     */
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    /**
     * A list of ADB clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.adb.GetClustersCluster[];
    readonly descriptionRegex?: string;
    /**
     * A list of ADB cluster descriptions. 
     */
    readonly descriptions: string[];
    /**
     * A list of ADB cluster IDs. 
     */
    readonly ids: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}