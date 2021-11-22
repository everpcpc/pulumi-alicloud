// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.arms.getPrometheusAlertRules({
 *     clusterId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const armsPrometheusAlertRuleId1 = ids.then(ids => ids.rules?[0]?.id);
 * const nameRegex = alicloud.arms.getPrometheusAlertRules({
 *     clusterId: "example_value",
 *     nameRegex: "^my-PrometheusAlertRule",
 * });
 * export const armsPrometheusAlertRuleId2 = nameRegex.then(nameRegex => nameRegex.rules?[0]?.id);
 * ```
 */
export function getPrometheusAlertRules(args: GetPrometheusAlertRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetPrometheusAlertRulesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:arms/getPrometheusAlertRules:getPrometheusAlertRules", {
        "clusterId": args.clusterId,
        "ids": args.ids,
        "matchExpressions": args.matchExpressions,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrometheusAlertRules.
 */
export interface GetPrometheusAlertRulesArgs {
    /**
     * The ID of the cluster.
     */
    clusterId: string;
    /**
     * A list of Prometheus Alert Rule IDs.
     */
    ids?: string[];
    matchExpressions?: string;
    /**
     * A regex string to filter results by Prometheus Alert Rule name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of the resource. Valid values: `0`, `1`.
     * * `1`: open.
     * * `0`: off.
     */
    status?: number;
    /**
     * The type of the alert rule.
     */
    type?: string;
}

/**
 * A collection of values returned by getPrometheusAlertRules.
 */
export interface GetPrometheusAlertRulesResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly matchExpressions?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly rules: outputs.arms.GetPrometheusAlertRulesRule[];
    readonly status?: number;
    readonly type?: string;
}

export function getPrometheusAlertRulesOutput(args: GetPrometheusAlertRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrometheusAlertRulesResult> {
    return pulumi.output(args).apply(a => getPrometheusAlertRules(a, opts))
}

/**
 * A collection of arguments for invoking getPrometheusAlertRules.
 */
export interface GetPrometheusAlertRulesOutputArgs {
    /**
     * The ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * A list of Prometheus Alert Rule IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    matchExpressions?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Prometheus Alert Rule name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `0`, `1`.
     * * `1`: open.
     * * `0`: off.
     */
    status?: pulumi.Input<number>;
    /**
     * The type of the alert rule.
     */
    type?: pulumi.Input<string>;
}
