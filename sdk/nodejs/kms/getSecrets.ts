// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.86.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const kmsSecretsDs = alicloud.kms.getSecrets({
 *     fetchTags: true,
 *     nameRegex: "name_regex",
 *     tags: {
 *         "k-aa": "v-aa",
 *         "k-bb": "v-bb",
 *     },
 * });
 * export const firstSecretId = kmsSecretsDs.then(kmsSecretsDs => kmsSecretsDs.secrets?.[0]?.id);
 * ```
 */
export function getSecrets(args?: GetSecretsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:kms/getSecrets:getSecrets", {
        "enableDetails": args.enableDetails,
        "fetchTags": args.fetchTags,
        "filters": args.filters,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecrets.
 */
export interface GetSecretsArgs {
    /**
     * Default to `false`. Set it to true can output more details.
     */
    enableDetails?: boolean;
    /**
     * Whether to include the predetermined resource tag in the return value. Default to `false`.
     */
    fetchTags?: boolean;
    /**
     * The secret filter. The filter consists of one or more key-value pairs. 
     * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
     */
    filters?: string;
    /**
     * A list of KMS Secret ids. The value is same as KMS secret_name.
     */
    ids?: string[];
    /**
     * A regex string to filter the results by the KMS secret_name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getSecrets.
 */
export interface GetSecretsResult {
    readonly enableDetails?: boolean;
    readonly fetchTags?: boolean;
    readonly filters?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Kms Secret ids. The value is same as KMS secret_name.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of KMS Secret names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of KMS Secrets. Each element contains the following attributes:
     */
    readonly secrets: outputs.kms.GetSecretsSecret[];
    /**
     * (Optional) A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.86.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const kmsSecretsDs = alicloud.kms.getSecrets({
 *     fetchTags: true,
 *     nameRegex: "name_regex",
 *     tags: {
 *         "k-aa": "v-aa",
 *         "k-bb": "v-bb",
 *     },
 * });
 * export const firstSecretId = kmsSecretsDs.then(kmsSecretsDs => kmsSecretsDs.secrets?.[0]?.id);
 * ```
 */
export function getSecretsOutput(args?: GetSecretsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretsResult> {
    return pulumi.output(args).apply((a: any) => getSecrets(a, opts))
}

/**
 * A collection of arguments for invoking getSecrets.
 */
export interface GetSecretsOutputArgs {
    /**
     * Default to `false`. Set it to true can output more details.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * Whether to include the predetermined resource tag in the return value. Default to `false`.
     */
    fetchTags?: pulumi.Input<boolean>;
    /**
     * The secret filter. The filter consists of one or more key-value pairs. 
     * More details see API [ListSecrets](https://www.alibabacloud.com/help/en/key-management-service/latest/listsecrets).
     */
    filters?: pulumi.Input<string>;
    /**
     * A list of KMS Secret ids. The value is same as KMS secret_name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter the results by the KMS secret_name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
