// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides available scaling rule resources. 
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const scalingrulesDs = pulumi.output(alicloud.ess.getScalingRules({
 *     ids: [
 *         "scalingRuleId1",
 *         "scalingRuleId2",
 *     ],
 *     nameRegex: "scalingRuleName",
 *     scalingGroupId: "scalingGroupId",
 * }, { async: true }));
 * 
 * export const firstScalingRule = scalingrulesDs.rules[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scaling_rules.html.markdown.
 */
export function getScalingRules(args?: GetScalingRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetScalingRulesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ess/getScalingRules:getScalingRules", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scalingGroupId": args.scalingGroupId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getScalingRules.
 */
export interface GetScalingRulesArgs {
    /**
     * A list of scaling rule IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter resulting scaling rules by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Scaling group id the scaling rules belong to.
     */
    readonly scalingGroupId?: string;
    /**
     * Type of scaling rule.
     */
    readonly type?: string;
}

/**
 * A collection of values returned by getScalingRules.
 */
export interface GetScalingRulesResult {
    /**
     * A list of scaling rule ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of scaling rule names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of scaling rules. Each element contains the following attributes:
     */
    readonly rules: outputs.ess.GetScalingRulesRule[];
    /**
     * ID of the scaling group.
     */
    readonly scalingGroupId?: string;
    /**
     * Type of the scaling rule.
     */
    readonly type?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
