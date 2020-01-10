// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 * const k8sClusters = alicloud.cs.getManagedKubernetesClusters({
 *     nameRegex: "my-first-k8s",
 *     outputFile: "my-first-k8s-json",
 * });
 * 
 * export const output = k8sClusters.clusters;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cs_managed_kubernetes_clusters.html.markdown.
 */
export function getManagedKubernetesClusters(args?: GetManagedKubernetesClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedKubernetesClustersResult> & GetManagedKubernetesClustersResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetManagedKubernetesClustersResult> = pulumi.runtime.invoke("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getManagedKubernetesClusters.
 */
export interface GetManagedKubernetesClustersArgs {
    readonly enableDetails?: boolean;
    /**
     * Cluster IDs to filter.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by cluster name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
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
     * A list of matched Kubernetes clusters' ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of matched Kubernetes clusters' names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
