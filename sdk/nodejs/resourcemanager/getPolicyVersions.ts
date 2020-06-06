// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Resource Manager Policy Versions of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.85.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultPolicyVersions = pulumi.output(alicloud.resourcemanager.getPolicyVersions({
 *     policyName: "tftest",
 *     policyType: "Custom",
 * }, { async: true }));
 *
 * export const firstPolicyVersionId = defaultPolicyVersions.versions[0].id;
 * ```
 */
export function getPolicyVersions(args: GetPolicyVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyVersionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:resourcemanager/getPolicyVersions:getPolicyVersions", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "policyName": args.policyName,
        "policyType": args.policyType,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyVersions.
 */
export interface GetPolicyVersionsArgs {
    /**
     * A list of policy version IDs.
     */
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The name of the policy.
     */
    readonly policyName: string;
    /**
     * The type of the policy. Valid values:`Custom` and `System`.
     */
    readonly policyType: string;
}

/**
 * A collection of values returned by getPolicyVersions.
 */
export interface GetPolicyVersionsResult {
    /**
     * A list of policy version IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    readonly policyName: string;
    readonly policyType: string;
    /**
     * A list of policy versions. Each element contains the following attributes:
     */
    readonly versions: outputs.resourcemanager.GetPolicyVersionsVersion[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}