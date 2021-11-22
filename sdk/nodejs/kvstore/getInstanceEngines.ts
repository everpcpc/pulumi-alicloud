// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the KVStore instance engines resource available info of Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.51.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const resourcesZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: "KVStore",
 * }));
 * const resourcesInstanceEngines = resourcesZones.apply(resourcesZones => alicloud.kvstore.getInstanceEngines({
 *     engine: "Redis",
 *     engineVersion: "5.0",
 *     instanceChargeType: "PrePaid",
 *     outputFile: "./engines.txt",
 *     zoneId: resourcesZones.zones[0].id,
 * }));
 *
 * export const firstKvstoreInstanceClass = resourcesInstanceEngines.instanceEngines[0].engine;
 * ```
 */
export function getInstanceEngines(args: GetInstanceEnginesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceEnginesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:kvstore/getInstanceEngines:getInstanceEngines", {
        "engine": args.engine,
        "engineVersion": args.engineVersion,
        "instanceChargeType": args.instanceChargeType,
        "outputFile": args.outputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceEngines.
 */
export interface GetInstanceEnginesArgs {
    /**
     * Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
     */
    engine?: string;
    /**
     * Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
     */
    engineVersion?: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
     */
    instanceChargeType?: string;
    outputFile?: string;
    /**
     * The Zone to launch the KVStore instance.
     */
    zoneId: string;
}

/**
 * A collection of values returned by getInstanceEngines.
 */
export interface GetInstanceEnginesResult {
    /**
     * Database type.
     */
    readonly engine?: string;
    /**
     * KVStore Instance version.
     */
    readonly engineVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceChargeType?: string;
    /**
     * A list of KVStore available instance engines. Each element contains the following attributes:
     */
    readonly instanceEngines: outputs.kvstore.GetInstanceEnginesInstanceEngine[];
    readonly outputFile?: string;
    /**
     * The Zone to launch the KVStore instance.
     */
    readonly zoneId: string;
}

export function getInstanceEnginesOutput(args: GetInstanceEnginesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceEnginesResult> {
    return pulumi.output(args).apply(a => getInstanceEngines(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceEngines.
 */
export interface GetInstanceEnginesOutputArgs {
    /**
     * Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
     */
    engine?: pulumi.Input<string>;
    /**
     * Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The Zone to launch the KVStore instance.
     */
    zoneId: pulumi.Input<string>;
}
