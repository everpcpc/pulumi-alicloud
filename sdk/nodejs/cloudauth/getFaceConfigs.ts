// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Cloudauth Face Configs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.137.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultFaceConfig = new alicloud.cloudauth.FaceConfig("defaultFaceConfig", {
 *     bizName: "example-value",
 *     bizType: "example-value",
 * });
 * const defaultFaceConfigs = pulumi.all([defaultFaceConfig.id, defaultFaceConfig.bizName]).apply(([id, bizName]) => alicloud.cloudauth.getFaceConfigs({
 *     ids: [id],
 *     nameRegex: bizName,
 * }));
 * export const faceConfig = defaultFaceConfigs.apply(defaultFaceConfigs => defaultFaceConfigs.configs?[0]);
 * ```
 */
export function getFaceConfigs(args?: GetFaceConfigsArgs, opts?: pulumi.InvokeOptions): Promise<GetFaceConfigsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cloudauth/getFaceConfigs:getFaceConfigs", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getFaceConfigs.
 */
export interface GetFaceConfigsArgs {
    /**
     * A list of Face Config IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by biz_name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getFaceConfigs.
 */
export interface GetFaceConfigsResult {
    readonly configs: outputs.cloudauth.GetFaceConfigsConfig[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}

export function getFaceConfigsOutput(args?: GetFaceConfigsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFaceConfigsResult> {
    return pulumi.output(args).apply(a => getFaceConfigs(a, opts))
}

/**
 * A collection of arguments for invoking getFaceConfigs.
 */
export interface GetFaceConfigsOutputArgs {
    /**
     * A list of Face Config IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by biz_name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
