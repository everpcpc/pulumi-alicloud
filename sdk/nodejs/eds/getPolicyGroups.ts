// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecd Policy Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.eds.EcdPolicyGroup("default", {
 *     policyGroupName: "my-policy-group",
 *     clipboard: "read",
 *     localDrive: "read",
 *     usbRedirect: "off",
 *     watermark: "off",
 *     authorizeAccessPolicyRules: [{
 *         description: "my-description1",
 *         cidrIp: "1.2.3.45/24",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         type: "inflow",
 *         policy: "accept",
 *         description: "my-description",
 *         portRange: "80/80",
 *         ipProtocol: "TCP",
 *         priority: "1",
 *         cidrIp: "1.2.3.4/24",
 *     }],
 * });
 * const nameRegex = alicloud.eds.getPolicyGroups({
 *     nameRegex: "^my-policy",
 * });
 * export const ecdPolicyGroupId = nameRegex.then(nameRegex => nameRegex.groups[0].id);
 * ```
 */
export function getPolicyGroups(args?: GetPolicyGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:eds/getPolicyGroups:getPolicyGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicyGroups.
 */
export interface GetPolicyGroupsArgs {
    /**
     * A list of Policy Group IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by Policy Group name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The status of policy.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getPolicyGroups.
 */
export interface GetPolicyGroupsResult {
    readonly groups: outputs.eds.GetPolicyGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
