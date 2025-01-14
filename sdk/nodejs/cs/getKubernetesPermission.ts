// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Ram user permissions.
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const usersDs = alicloud.ram.getUsers({
 *     nameRegex: "your_user_name",
 * });
 * const default = usersDs.then(usersDs => alicloud.cs.getKubernetesPermission({
 *     uid: usersDs.users?.[0]?.id,
 * }));
 * export const permissions = _default.then(_default => _default.permissions);
 * ```
 */
export function getKubernetesPermission(args: GetKubernetesPermissionArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesPermissionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cs/getKubernetesPermission:getKubernetesPermission", {
        "permissions": args.permissions,
        "uid": args.uid,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesPermission.
 */
export interface GetKubernetesPermissionArgs {
    /**
     * A list of user permission.
     */
    permissions?: inputs.cs.GetKubernetesPermissionPermission[];
    /**
     * The ID of the RAM user. If you want to query the permissions of a RAM role, specify the ID of the RAM role.
     */
    uid: string;
}

/**
 * A collection of values returned by getKubernetesPermission.
 */
export interface GetKubernetesPermissionResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of user permission.
     */
    readonly permissions?: outputs.cs.GetKubernetesPermissionPermission[];
    /**
     * The ID of the RAM user. If you want to query the permissions of a RAM role, specify the ID of the RAM role.
     */
    readonly uid: string;
}
/**
 * This data source provides a list of Ram user permissions.
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const usersDs = alicloud.ram.getUsers({
 *     nameRegex: "your_user_name",
 * });
 * const default = usersDs.then(usersDs => alicloud.cs.getKubernetesPermission({
 *     uid: usersDs.users?.[0]?.id,
 * }));
 * export const permissions = _default.then(_default => _default.permissions);
 * ```
 */
export function getKubernetesPermissionOutput(args: GetKubernetesPermissionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubernetesPermissionResult> {
    return pulumi.output(args).apply((a: any) => getKubernetesPermission(a, opts))
}

/**
 * A collection of arguments for invoking getKubernetesPermission.
 */
export interface GetKubernetesPermissionOutputArgs {
    /**
     * A list of user permission.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.cs.GetKubernetesPermissionPermissionArgs>[]>;
    /**
     * The ID of the RAM user. If you want to query the permissions of a RAM role, specify the ID of the RAM role.
     */
    uid: pulumi.Input<string>;
}
