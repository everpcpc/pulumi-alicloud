// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Image Pipelines of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.163.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecs.getEcsImagePipeline({
 *     ids: ["example_value"],
 * });
 * export const ecsImagePipelineId1 = ids.then(ids => ids.pipelines?[0]?.id);
 * const nameRegex = alicloud.ecs.getEcsImagePipeline({
 *     nameRegex: "^my-ImagePipeline",
 * });
 * export const ecsImagePipelineId2 = nameRegex.then(nameRegex => nameRegex.pipelines?[0]?.id);
 * ```
 */
export function getEcsImagePipeline(args?: GetEcsImagePipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsImagePipelineResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:ecs/getEcsImagePipeline:getEcsImagePipeline", {
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsImagePipeline.
 */
export interface GetEcsImagePipelineArgs {
    /**
     * A list of Image Pipeline ids.
     */
    ids?: string[];
    /**
     * The name of the image template.
     */
    name?: string;
    /**
     * A regex string to filter results by Image Pipeline name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the resource group to which the image template belongs.
     */
    resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getEcsImagePipeline.
 */
export interface GetEcsImagePipelineResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly name?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly pipelines: outputs.ecs.GetEcsImagePipelinePipeline[];
    readonly resourceGroupId?: string;
    readonly tags?: {[key: string]: any};
}

export function getEcsImagePipelineOutput(args?: GetEcsImagePipelineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsImagePipelineResult> {
    return pulumi.output(args).apply(a => getEcsImagePipeline(a, opts))
}

/**
 * A collection of arguments for invoking getEcsImagePipeline.
 */
export interface GetEcsImagePipelineOutputArgs {
    /**
     * A list of Image Pipeline ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the image template.
     */
    name?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Image Pipeline name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the image template belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
