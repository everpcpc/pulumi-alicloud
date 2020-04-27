// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of RAM Roles in an Alibaba Cloud account according to the specified filters.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const rolesDs = pulumi.output(alicloud.ram.getRoles({
 *     nameRegex: ".*test.*",
 *     outputFile: "roles.txt",
 *     policyName: "AliyunACSDefaultAccess",
 *     policyType: "Custom",
 * }, { async: true }));
 * 
 * export const firstRoleId = rolesDs.roles[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_roles.html.markdown.
 */
export function getRoles(args?: GetRolesArgs, opts?: pulumi.InvokeOptions): Promise<GetRolesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ram/getRoles:getRoles", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "policyName": args.policyName,
        "policyType": args.policyType,
    }, opts);
}

/**
 * A collection of arguments for invoking getRoles.
 */
export interface GetRolesArgs {
    /**
     * A list of ram role IDs. 
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by the role name.
     * * `ids` (Optional, Available 1.53.0+) - A list of ram role IDs.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Filter results by a specific policy name. If you set this parameter without setting `policyType`, the later will be automatically set to `System`. The resulting roles will be attached to the specified policy.
     */
    readonly policyName?: string;
    /**
     * Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policyName` as well.
     */
    readonly policyType?: string;
}

/**
 * A collection of values returned by getRoles.
 */
export interface GetRolesResult {
    /**
     * A list of ram role IDs. 
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of ram role names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly policyName?: string;
    readonly policyType?: string;
    /**
     * A list of roles. Each element contains the following attributes:
     */
    readonly roles: outputs.ram.GetRolesRole[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
