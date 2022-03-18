// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.138.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultEcsBackupPlans = alicloud.hbr.getEcsBackupPlans({
 *     nameRegex: "plan-name",
 * });
 * const defaultBackupJobs = Promise.all([defaultEcsBackupPlans, defaultEcsBackupPlans]).then(([defaultEcsBackupPlans, defaultEcsBackupPlans1]) => alicloud.hbr.getBackupJobs({
 *     sourceType: "ECS_FILE",
 *     filters: [
 *         {
 *             key: "VaultId",
 *             operator: "IN",
 *             values: [defaultEcsBackupPlans.plans?[0]?.vaultId],
 *         },
 *         {
 *             key: "InstanceId",
 *             operator: "IN",
 *             values: [defaultEcsBackupPlans1.plans?[0]?.instanceId],
 *         },
 *         {
 *             key: "CompleteTime",
 *             operator: "BETWEEN",
 *             values: [
 *                 "2021-08-23T14:17:15CST",
 *                 "2021-08-24T14:17:15CST",
 *             ],
 *         },
 *     ],
 * }));
 * const example = Promise.all([defaultEcsBackupPlans, defaultEcsBackupPlans]).then(([defaultEcsBackupPlans, defaultEcsBackupPlans1]) => alicloud.hbr.getBackupJobs({
 *     sourceType: "ECS_FILE",
 *     status: "COMPLETE",
 *     filters: [
 *         {
 *             key: "VaultId",
 *             operator: "IN",
 *             values: [defaultEcsBackupPlans.plans?[0]?.vaultId],
 *         },
 *         {
 *             key: "InstanceId",
 *             operator: "IN",
 *             values: [defaultEcsBackupPlans1.plans?[0]?.instanceId],
 *         },
 *         {
 *             key: "CompleteTime",
 *             operator: "LESS_THAN",
 *             values: ["2021-10-20T20:20:20CST"],
 *         },
 *     ],
 * }));
 * export const alicloudHbrBackupJobsDefault1 = defaultBackupJobs.then(defaultBackupJobs => defaultBackupJobs.jobs?[0]?.id);
 * export const alicloudHbrBackupJobsExample1 = example.then(example => example.jobs?[0]?.id);
 * ```
 */
export function getBackupJobs(args: GetBackupJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupJobsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:hbr/getBackupJobs:getBackupJobs", {
        "filters": args.filters,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "sortDirection": args.sortDirection,
        "sourceType": args.sourceType,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupJobs.
 */
export interface GetBackupJobsArgs {
    filters?: inputs.hbr.GetBackupJobsFilter[];
    /**
     * A list of Backup Job IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
     */
    sortDirection?: string;
    /**
     * The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
     */
    sourceType: string;
    /**
     * The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
     */
    status?: string;
}

/**
 * A collection of values returned by getBackupJobs.
 */
export interface GetBackupJobsResult {
    readonly filters?: outputs.hbr.GetBackupJobsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly jobs: outputs.hbr.GetBackupJobsJob[];
    readonly outputFile?: string;
    readonly sortDirection?: string;
    readonly sourceType: string;
    readonly status?: string;
}

export function getBackupJobsOutput(args: GetBackupJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupJobsResult> {
    return pulumi.output(args).apply(a => getBackupJobs(a, opts))
}

/**
 * A collection of arguments for invoking getBackupJobs.
 */
export interface GetBackupJobsOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.hbr.GetBackupJobsFilterArgs>[]>;
    /**
     * A list of Backup Job IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
     */
    sortDirection?: pulumi.Input<string>;
    /**
     * The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
     */
    sourceType: pulumi.Input<string>;
    /**
     * The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
     */
    status?: pulumi.Input<string>;
}
