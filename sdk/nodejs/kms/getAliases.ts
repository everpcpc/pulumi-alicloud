// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of KMS aliases in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in v1.79.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const kmsAliases = pulumi.output(alicloud.kms.getAliases({
 *     ids: ["d89e8a53-b708-41aa-8c67-6873axxx"],
 *     nameRegex: "alias/tf-testKmsAlias_123",
 * }));
 *
 * export const firstKeyId = alicloud_kms_keys_kms_keys_ds.keys.0.id;
 * ```
 */
export function getAliases(args?: GetAliasesArgs, opts?: pulumi.InvokeOptions): Promise<GetAliasesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:kms/getAliases:getAliases", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAliases.
 */
export interface GetAliasesArgs {
    /**
     * A list of KMS aliases IDs. The value is same as KMS alias_name.
     */
    ids?: string[];
    /**
     * A regex string to filter the results by the KMS alias name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getAliases.
 */
export interface GetAliasesResult {
    /**
     * A list of KMS User alias. Each element contains the following attributes:
     */
    readonly aliases: outputs.kms.GetAliasesAlias[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of kms aliases IDs. The value is same as KMS alias_name.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of KMS alias name.
     */
    readonly names: string[];
    readonly outputFile?: string;
}

export function getAliasesOutput(args?: GetAliasesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAliasesResult> {
    return pulumi.output(args).apply(a => getAliases(a, opts))
}

/**
 * A collection of arguments for invoking getAliases.
 */
export interface GetAliasesOutputArgs {
    /**
     * A list of KMS aliases IDs. The value is same as KMS alias_name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter the results by the KMS alias name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
