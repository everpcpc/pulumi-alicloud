// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Hpc Clusters of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.116.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ecs.getHpcClusters({
 *     ids: ["hpc-bp1i09xxxxxxxx"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const firstEcsHpcClusterId = example.then(example => example.clusters?[0]?.id);
 * ```
 */
export function getHpcClusters(args?: GetHpcClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetHpcClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ecs/getHpcClusters:getHpcClusters", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getHpcClusters.
 */
export interface GetHpcClustersArgs {
    /**
     * A list of Hpc Cluster IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Hpc Cluster name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getHpcClusters.
 */
export interface GetHpcClustersResult {
    readonly clusters: outputs.ecs.GetHpcClustersCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}

export function getHpcClustersOutput(args?: GetHpcClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHpcClustersResult> {
    return pulumi.output(args).apply(a => getHpcClusters(a, opts))
}

/**
 * A collection of arguments for invoking getHpcClusters.
 */
export interface GetHpcClustersOutputArgs {
    /**
     * A list of Hpc Cluster IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Hpc Cluster name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
