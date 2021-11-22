// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
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
 * const k8sClusters = pulumi.output(alicloud.cs.getManagedKubernetesClusters({
 *     nameRegex: "my-first-k8s",
 *     outputFile: "my-first-k8s-json",
 * }));
 *
 * export const output = k8sClusters.clusters;
 * ```
 */
export function getManagedKubernetesClusters(args?: GetManagedKubernetesClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedKubernetesClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersArgs {
    enableDetails?: boolean;
    /**
     * Cluster IDs to filter.
     */
    ids?: string[];
    /**
     * A regex string to filter results by cluster name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersResult {
    /**
     * A list of matched Kubernetes clusters. Each element contains the following attributes:
     */
    readonly clusters: outputs.cs.GetManagedKubernetesClustersCluster[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of matched Kubernetes clusters' ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of matched Kubernetes clusters' names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}

export function getManagedKubernetesClustersOutput(args?: GetManagedKubernetesClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetManagedKubernetesClustersResult> {
    return pulumi.output(args).apply(a => getManagedKubernetesClusters(a, opts))
}

/**
 * A collection of arguments for invoking getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersOutputArgs {
    enableDetails?: pulumi.Input<boolean>;
    /**
     * Cluster IDs to filter.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by cluster name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
