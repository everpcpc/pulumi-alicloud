// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the OSS buckets of the current Alibaba Cloud user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const ossBucketsDs = pulumi.output(alicloud.oss.getBuckets({
 *     nameRegex: "sampleOssBucket",
 * }, { async: true }));
 * 
 * export const firstOssBucketName = ossBucketsDs.buckets[0].name;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/oss_buckets.html.markdown.
 */
export function getBuckets(args?: GetBucketsArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:oss/getBuckets:getBuckets", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBuckets.
 */
export interface GetBucketsArgs {
    /**
     * A regex string to filter results by bucket name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getBuckets.
 */
export interface GetBucketsResult {
    /**
     * A list of buckets. Each element contains the following attributes:
     */
    readonly buckets: outputs.oss.GetBucketsBucket[];
    readonly nameRegex?: string;
    /**
     * A list of bucket names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
