// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of KMS keys in an Alibaba Cloud account according to the specified filters.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * // Declare the data source
 * const kmsKeysDs = alicloud.kms.getKeys({
 *     descriptionRegex: "Hello KMS",
 *     outputFile: "kms_keys.json",
 * });
 * 
 * export const firstKeyId = kmsKeysDs.keys[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kms_keys.html.markdown.
 */
export function getKeys(args?: GetKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetKeysResult> & GetKeysResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetKeysResult> = pulumi.runtime.invoke("alicloud:kms/getKeys:getKeys", {
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getKeys.
 */
export interface GetKeysArgs {
    /**
     * A regex string to filter the results by the KMS key description.
     */
    readonly descriptionRegex?: string;
    /**
     * A list of KMS key IDs.
     */
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getKeys.
 */
export interface GetKeysResult {
    readonly descriptionRegex?: string;
    /**
     * A list of KMS key IDs.
     */
    readonly ids: string[];
    /**
     * A list of KMS keys. Each element contains the following attributes:
     */
    readonly keys: outputs.kms.GetKeysKey[];
    readonly outputFile?: string;
    /**
     * Status of the key. Possible values: `Enabled`, `Disabled` and `PendingDeletion`.
     */
    readonly status?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
