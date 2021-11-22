// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecs Deployment Sets of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.140.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.ecs.getEcsDeploymentSets({
 *     ids: ["example_id"],
 * });
 * export const ecsDeploymentSetId1 = ids.then(ids => ids.sets?[0]?.id);
 * const nameRegex = alicloud.ecs.getEcsDeploymentSets({
 *     nameRegex: "^my-DeploymentSet",
 * });
 * export const ecsDeploymentSetId2 = nameRegex.then(nameRegex => nameRegex.sets?[0]?.id);
 * ```
 */
export function getEcsDeploymentSets(args?: GetEcsDeploymentSetsArgs, opts?: pulumi.InvokeOptions): Promise<GetEcsDeploymentSetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ecs/getEcsDeploymentSets:getEcsDeploymentSets", {
        "deploymentSetName": args.deploymentSetName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "strategy": args.strategy,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcsDeploymentSets.
 */
export interface GetEcsDeploymentSetsArgs {
    /**
     * The name of the deployment set.
     */
    deploymentSetName?: string;
    /**
     * A list of Deployment Set IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Deployment Set name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The deployment strategy.
     */
    strategy?: string;
}

/**
 * A collection of values returned by getEcsDeploymentSets.
 */
export interface GetEcsDeploymentSetsResult {
    readonly deploymentSetName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly sets: outputs.ecs.GetEcsDeploymentSetsSet[];
    readonly strategy?: string;
}

export function getEcsDeploymentSetsOutput(args?: GetEcsDeploymentSetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEcsDeploymentSetsResult> {
    return pulumi.output(args).apply(a => getEcsDeploymentSets(a, opts))
}

/**
 * A collection of arguments for invoking getEcsDeploymentSets.
 */
export interface GetEcsDeploymentSetsOutputArgs {
    /**
     * The name of the deployment set.
     */
    deploymentSetName?: pulumi.Input<string>;
    /**
     * A list of Deployment Set IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Deployment Set name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The deployment strategy.
     */
    strategy?: pulumi.Input<string>;
}
