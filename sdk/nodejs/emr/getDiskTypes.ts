// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * The `alicloud.emr.getDiskTypes` data source provides a collection of data disk and
 * system disk types available in Alibaba Cloud account when create a emr cluster.
 *
 * > **NOTE:** Available in 1.60.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultDiskTypes = pulumi.output(alicloud.emr.getDiskTypes({
 *     clusterType: "HADOOP",
 *     destinationResource: "DataDisk",
 *     instanceChargeType: "PostPaid",
 *     instanceType: "ecs.g5.xlarge",
 *     zoneId: "cn-huhehaote-a",
 * }));
 *
 * export const dataDiskType = defaultDiskTypes.types[0].value;
 * ```
 */
export function getDiskTypes(args: GetDiskTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetDiskTypesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:emr/getDiskTypes:getDiskTypes", {
        "clusterType": args.clusterType,
        "destinationResource": args.destinationResource,
        "instanceChargeType": args.instanceChargeType,
        "instanceType": args.instanceType,
        "outputFile": args.outputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDiskTypes.
 */
export interface GetDiskTypesArgs {
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
     * The ecs instance type of create emr cluster instance.
     */
    instanceType: string;
    outputFile?: string;
    /**
     * The Zone to create emr cluster instance.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getDiskTypes.
 */
export interface GetDiskTypesResult {
    readonly clusterType: string;
    readonly destinationResource: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of data disk and system disk type IDs.
     */
    readonly ids: string[];
    readonly instanceChargeType: string;
    readonly instanceType: string;
    readonly outputFile?: string;
    /**
     * A list of emr instance types. Each element contains the following attributes:
     */
    readonly types: outputs.emr.GetDiskTypesType[];
    readonly zoneId?: string;
}

export function getDiskTypesOutput(args: GetDiskTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDiskTypesResult> {
    return pulumi.output(args).apply(a => getDiskTypes(a, opts))
}

/**
 * A collection of arguments for invoking getDiskTypes.
 */
export interface GetDiskTypesOutputArgs {
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
     * The ecs instance type of create emr cluster instance.
     */
    instanceType: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The Zone to create emr cluster instance.
     */
    zoneId?: pulumi.Input<string>;
}
