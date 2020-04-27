// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.emr.getInstanceTypes` data source provides a collection of ecs
 * instance types available in Alibaba Cloud account when create a emr cluster.
 * 
 * > **NOTE:** Available in 1.59.0+
 * 
 * ## Example Usage
 * 
 * 
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
 * }, { async: true }));
 * 
 * export const firstInstanceType = defaultInstanceTypes.types[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/emr_instance_types.html.markdown.
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
    readonly clusterType: string;
    /**
     * The destination resource of emr cluster instance
     */
    readonly destinationResource: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     */
    readonly instanceChargeType: string;
    /**
     * Filter the specific ecs instance type to create emr cluster.
     */
    readonly instanceType?: string;
    readonly outputFile?: string;
    /**
     * Whether the current storage disk is local or not.
     */
    readonly supportLocalStorage?: boolean;
    /**
     * The specific supported node type list.
     * Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
     */
    readonly supportNodeTypes?: string[];
    /**
     * The supported resources of specific zoneId.
     */
    readonly zoneId?: string;
}

/**
 * A collection of values returned by getInstanceTypes.
 */
export interface GetInstanceTypesResult {
    readonly clusterType: string;
    readonly destinationResource: string;
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
