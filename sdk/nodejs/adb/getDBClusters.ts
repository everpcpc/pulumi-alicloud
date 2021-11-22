// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Adb DBClusters of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.121.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.adb.getDBClusters({
 *     descriptionRegex: "example",
 * });
 * export const firstAdbDbClusterId = example.then(example => example.clusters?[0]?.id);
 * ```
 */
export function getDBClusters(args?: GetDBClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetDBClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:adb/getDBClusters:getDBClusters", {
        "description": args.description,
        "descriptionRegex": args.descriptionRegex,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDBClusters.
 */
export interface GetDBClustersArgs {
    /**
     * The description of DBCluster.
     */
    description?: string;
    /**
     * A regex string to filter results by DBCluster description.
     */
    descriptionRegex?: string;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of DBCluster IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: string;
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * The tag of the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getDBClusters.
 */
export interface GetDBClustersResult {
    readonly clusters: outputs.adb.GetDBClustersCluster[];
    readonly description?: string;
    readonly descriptionRegex?: string;
    readonly descriptions: string[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly resourceGroupId?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
}

export function getDBClustersOutput(args?: GetDBClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBClustersResult> {
    return pulumi.output(args).apply(a => getDBClusters(a, opts))
}

/**
 * A collection of arguments for invoking getDBClusters.
 */
export interface GetDBClustersOutputArgs {
    /**
     * The description of DBCluster.
     */
    description?: pulumi.Input<string>;
    /**
     * A regex string to filter results by DBCluster description.
     */
    descriptionRegex?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of DBCluster IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag of the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
