// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides MountTargets available to the user.
 * 
 * > NOTE: Available in 1.35.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const mt = alicloud.nas.getMountTargets({
 *     accessGroupName: "tf-testAccNasConfig",
 *     fileSystemId: "1a2sc4d",
 * });
 * 
 * export const alicloudNasMountTargetsId = mt.targets[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/nas_mount_targets.html.markdown.
 */
export function getMountTargets(args: GetMountTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:nas/getMountTargets:getMountTargets", {
        "accessGroupName": args.accessGroupName,
        "fileSystemId": args.fileSystemId,
        "ids": args.ids,
        "mountTargetDomain": args.mountTargetDomain,
        "outputFile": args.outputFile,
        "type": args.type,
        "vpcId": args.vpcId,
        "vswitchId": args.vswitchId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountTargets.
 */
export interface GetMountTargetsArgs {
    /**
     * Filter results by a specific AccessGroupName.
     */
    readonly accessGroupName?: string;
    /**
     * The ID of the FileSystem that owns the MountTarget.
     */
    readonly fileSystemId: string;
    /**
     * A list of MountTargetDomain.
     */
    readonly ids?: string[];
    /**
     * Filter results by a specific MountTargetDomain.
     * 
     * @deprecated Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
     */
    readonly mountTargetDomain?: string;
    readonly outputFile?: string;
    /**
     * Filter results by a specific NetworkType.
     */
    readonly type?: string;
    /**
     * Filter results by a specific VpcId.
     */
    readonly vpcId?: string;
    /**
     * Filter results by a specific VSwitchId.
     */
    readonly vswitchId?: string;
}

/**
 * A collection of values returned by getMountTargets.
 */
export interface GetMountTargetsResult {
    /**
     * AccessGroup of The MountTarget.
     */
    readonly accessGroupName?: string;
    readonly fileSystemId: string;
    /**
     * A list of MountTargetDomain.
     */
    readonly ids: string[];
    /**
     * MountTargetDomain of the MountTarget.
     * * `type`- NetworkType of The MountTarget.
     * 
     * @deprecated Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
     */
    readonly mountTargetDomain?: string;
    readonly outputFile?: string;
    /**
     * A list of MountTargetDomains. Each element contains the following attributes:
     */
    readonly targets: outputs.nas.GetMountTargetsTarget[];
    readonly type?: string;
    /**
     * VpcId of The MountTarget.
     */
    readonly vpcId?: string;
    /**
     * VSwitchId of The MountTarget.
     */
    readonly vswitchId?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
