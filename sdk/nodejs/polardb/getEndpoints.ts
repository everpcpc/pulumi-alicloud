// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * The `alicloud.polardb.getEndpoints` data source provides a collection of PolarDB endpoints available in Alibaba Cloud account.
 * Filters support regular expression for the cluster name, searches by clusterId, and other filters which are listed below.
 *
 * > **NOTE:** Available in v1.68.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const polardbClustersDs = alicloud.polardb.getClusters({
 *     descriptionRegex: "pc-\\w+",
 *     status: "Running",
 * });
 * const default = polardbClustersDs.then(polardbClustersDs => alicloud.polardb.getEndpoints({
 *     dbClusterId: polardbClustersDs.clusters?[0]?.id,
 * }));
 * export const endpoint = _default.then(_default => _default.endpoints?[0]?.dbEndpointId);
 * ```
 */
export function getEndpoints(args: GetEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:polardb/getEndpoints:getEndpoints", {
        "dbClusterId": args.dbClusterId,
        "dbEndpointId": args.dbEndpointId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoints.
 */
export interface GetEndpointsArgs {
    /**
     * PolarDB cluster ID.
     */
    dbClusterId: string;
    /**
     * endpoint of the cluster.
     */
    dbEndpointId?: string;
}

/**
 * A collection of values returned by getEndpoints.
 */
export interface GetEndpointsResult {
    readonly dbClusterId: string;
    /**
     * The endpoint ID.
     */
    readonly dbEndpointId?: string;
    /**
     * A list of PolarDB cluster endpoints. Each element contains the following attributes:
     */
    readonly endpoints: outputs.polardb.GetEndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

export function getEndpointsOutput(args: GetEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointsResult> {
    return pulumi.output(args).apply(a => getEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking getEndpoints.
 */
export interface GetEndpointsOutputArgs {
    /**
     * PolarDB cluster ID.
     */
    dbClusterId: pulumi.Input<string>;
    /**
     * endpoint of the cluster.
     */
    dbEndpointId?: pulumi.Input<string>;
}
