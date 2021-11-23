// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpc Traffic Mirror Filter Egress Rules of the current Alibaba Cloud user.
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
 * const ids = alicloud.vpc.getTrafficMirrorFilterEgressRules({
 *     trafficMirrorFilterId: "example_traffic_mirror_filter_id",
 *     ids: ["example_id"],
 * });
 * export const vpcTrafficMirrorFilterEgressRuleId1 = ids.then(ids => ids.rules?[0]?.id);
 * const status = alicloud.vpc.getTrafficMirrorFilterEgressRules({
 *     trafficMirrorFilterId: "example_traffic_mirror_filter_id",
 *     ids: ["example_id"],
 *     status: "Created",
 * });
 * export const vpcTrafficMirrorFilterEgressRuleId2 = status.then(status => status.rules?[0]?.id);
 * ```
 */
export function getTrafficMirrorFilterEgressRules(args: GetTrafficMirrorFilterEgressRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficMirrorFilterEgressRulesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getTrafficMirrorFilterEgressRules:getTrafficMirrorFilterEgressRules", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "trafficMirrorFilterId": args.trafficMirrorFilterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrafficMirrorFilterEgressRules.
 */
export interface GetTrafficMirrorFilterEgressRulesArgs {
    /**
     * A list of Traffic Mirror Filter Egress Rule IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of the resource. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: string;
    /**
     * The ID of the filter associated with the outbound rule.
     */
    trafficMirrorFilterId: string;
}

/**
 * A collection of values returned by getTrafficMirrorFilterEgressRules.
 */
export interface GetTrafficMirrorFilterEgressRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly rules: outputs.vpc.GetTrafficMirrorFilterEgressRulesRule[];
    readonly status?: string;
    readonly trafficMirrorFilterId: string;
}

export function getTrafficMirrorFilterEgressRulesOutput(args: GetTrafficMirrorFilterEgressRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficMirrorFilterEgressRulesResult> {
    return pulumi.output(args).apply(a => getTrafficMirrorFilterEgressRules(a, opts))
}

/**
 * A collection of arguments for invoking getTrafficMirrorFilterEgressRules.
 */
export interface GetTrafficMirrorFilterEgressRulesOutputArgs {
    /**
     * A list of Traffic Mirror Filter Egress Rule IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the filter associated with the outbound rule.
     */
    trafficMirrorFilterId: pulumi.Input<string>;
}