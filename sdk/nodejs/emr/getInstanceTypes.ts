// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * The `alicloud.emr.getInstanceTypes` data source provides a collection of ecs
 * instance types available in Alibaba Cloud account when create a emr cluster.
 *
 * > **NOTE:** Available in 1.59.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstanceTypes = pulumi.output(alicloud.emr.getInstanceTypes({
 *     clusterType: "HADOOP",
 *     destinationResource: "InstanceType",
 *     instanceChargeType: "PostPaid",
 *     instanceType: "ecs.g5.2xlarge",
 *     supportLocalStorage: false,
 *     supportNodeTypes: [
 *         "MASTER",
 *         "CORE",
 *     ],
 * }));
 *
 * export const firstInstanceType = defaultInstanceTypes.types[0].id;
 * ```
 */
export function getInstanceTypes(args: GetInstanceTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:emr/getInstanceTypes:getInstanceTypes", {
        "clusterType": args.clusterType,
        "destinationResource": args.destinationResource,
        "instanceChargeType": args.instanceChargeType,
        "instanceType": args.instanceType,
        "outputFile": args.outputFile,
        "supportLocalStorage": args.supportLocalStorage,
        "supportNodeTypes": args.supportNodeTypes,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypes.
 */
export interface GetInstanceTypesArgs {
    /**
     * The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
     */
    clusterType: string;
    /**
     * The destination resource of emr cluster instance
     */
    destinationResource: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     */
    instanceChargeType: string;
    /**
     * Filter the specific ecs instance type to create emr cluster.
     */
    instanceType?: string;
    outputFile?: string;
    /**
     * Whether the current storage disk is local or not.
     */
    supportLocalStorage?: boolean;
    /**
     * The specific supported node type list.
     * Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
     */
    supportNodeTypes?: string[];
    /**
     * The supported resources of specific zoneId.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getInstanceTypes.
 */
export interface GetInstanceTypesResult {
    readonly clusterType: string;
    readonly destinationResource: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of emr instance types IDs.
     */
    readonly ids: string[];
    readonly instanceChargeType: string;
    readonly instanceType?: string;
    readonly outputFile?: string;
    readonly supportLocalStorage?: boolean;
    readonly supportNodeTypes?: string[];
    /**
     * A list of emr instance types. Each element contains the following attributes:
     */
    readonly types: outputs.emr.GetInstanceTypesType[];
    /**
     * The available zone id in Alibaba Cloud account
     */
    readonly zoneId?: string;
}

export function getInstanceTypesOutput(args: GetInstanceTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTypesResult> {
    return pulumi.output(args).apply(a => getInstanceTypes(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTypes.
 */
export interface GetInstanceTypesOutputArgs {
    /**
     * The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
     */
    clusterType: pulumi.Input<string>;
    /**
     * The destination resource of emr cluster instance
     */
    destinationResource: pulumi.Input<string>;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     */
    instanceChargeType: pulumi.Input<string>;
    /**
     * Filter the specific ecs instance type to create emr cluster.
     */
    instanceType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * Whether the current storage disk is local or not.
     */
    supportLocalStorage?: pulumi.Input<boolean>;
    /**
     * The specific supported node type list.
     * Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
     */
    supportNodeTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The supported resources of specific zoneId.
     */
    zoneId?: pulumi.Input<string>;
}
