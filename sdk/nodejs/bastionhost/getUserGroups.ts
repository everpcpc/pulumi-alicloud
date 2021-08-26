// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Bastionhost User Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.bastionhost.getUserGroups({
 *     instanceId: "bastionhost-cn-xxxx",
 *     ids: [
 *         "1",
 *         "2",
 *     ],
 * });
 * export const bastionhostUserGroupId1 = ids.then(ids => ids.groups[0].id);
 * const nameRegex = alicloud.bastionhost.getUserGroups({
 *     instanceId: "bastionhost-cn-xxxx",
 *     nameRegex: "^my-UserGroup",
 * });
 * export const bastionhostUserGroupId2 = nameRegex.then(nameRegex => nameRegex.groups[0].id);
 * ```
 */
export function getUserGroups(args: GetUserGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserGroupsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:bastionhost/getUserGroups:getUserGroups", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "userGroupName": args.userGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserGroups.
 */
export interface GetUserGroupsArgs {
    /**
     * A list of User Group self IDs.
     */
    readonly ids?: string[];
    /**
     * Specify the New Group of the Bastion Host of Instance Id.
     */
    readonly instanceId: string;
    /**
     * A regex string to filter results by User Group name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Specify the New Group Name. Supports up to 128 Characters.
     */
    readonly userGroupName?: string;
}

/**
 * A collection of values returned by getUserGroups.
 */
export interface GetUserGroupsResult {
    readonly groups: outputs.bastionhost.GetUserGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly userGroupName?: string;
}
