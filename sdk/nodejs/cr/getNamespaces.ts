// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Registry namespaces on Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.35.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Declare the data source
 * const myNamespaces = pulumi.output(alicloud.cr.getNamespaces({
 *     nameRegex: "my-namespace",
 *     outputFile: "my-namespace-json",
 * }));
 *
 * export const output = myNamespaces.namespaces;
 * ```
 */
export function getNamespaces(args?: GetNamespacesArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespacesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cr/getNamespaces:getNamespaces", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamespaces.
 */
export interface GetNamespacesArgs {
    /**
     * A regex string to filter results by namespace name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getNamespaces.
 */
export interface GetNamespacesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of matched Container Registry namespaces. Its element is a namespace name.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of namespace names.
     */
    readonly names: string[];
    /**
     * A list of matched Container Registry namespaces. Each element contains the following attributes:
     */
    readonly namespaces: outputs.cr.GetNamespacesNamespace[];
    readonly outputFile?: string;
}

export function getNamespacesOutput(args?: GetNamespacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamespacesResult> {
    return pulumi.output(args).apply(a => getNamespaces(a, opts))
}

/**
 * A collection of arguments for invoking getNamespaces.
 */
export interface GetNamespacesOutputArgs {
    /**
     * A regex string to filter results by namespace name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
