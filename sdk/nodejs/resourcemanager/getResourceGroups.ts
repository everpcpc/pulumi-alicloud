// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides resource groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.84.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = pulumi.output(alicloud.resourcemanager.getResourceGroups({
 *     nameRegex: "tftest",
 * }, { async: true }));
 *
 * export const firstResourceGroupId = example.groups[0].id;
 * ```
 */
export function getResourceGroups(args?: GetResourceGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:resourcemanager/getResourceGroups:getResourceGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceGroups.
 */
export interface GetResourceGroupsArgs {
    /**
     * A list of resource group IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by resource group name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The status of the resource group. Possible values:`Creating`,`Deleted`,`OK` and `PendingDelete`.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getResourceGroups.
 */
export interface GetResourceGroupsResult {
    /**
     * A list of resource groups. Each element contains the following attributes:
     */
    readonly groups: outputs.resourcemanager.GetResourceGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of resource group IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of resource group names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of the resource group. Possible values:`Creating`,`Deleted`,`OK` and `PendingDelete`.
     */
    readonly status?: string;
}
