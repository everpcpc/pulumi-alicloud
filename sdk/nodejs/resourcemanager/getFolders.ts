// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the resource manager folders of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.84.0+.
 *
 * > **NOTE:**  You can view only the information of the first-level child folders of the specified folder.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = pulumi.output(alicloud.resourcemanager.getFolders({
 *     nameRegex: "tftest",
 * }, { async: true }));
 *
 * export const firstFolderId = example.folders[0].id;
 * ```
 */
export function getFolders(args?: GetFoldersArgs, opts?: pulumi.InvokeOptions): Promise<GetFoldersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:resourcemanager/getFolders:getFolders", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "parentFolderId": args.parentFolderId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolders.
 */
export interface GetFoldersArgs {
    /**
     * A list of resource manager folders IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by folder name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The ID of the parent folder.
     */
    readonly parentFolderId?: string;
}

/**
 * A collection of values returned by getFolders.
 */
export interface GetFoldersResult {
    /**
     * A list of folders. Each element contains the following attributes:
     */
    readonly folders: outputs.resourcemanager.GetFoldersFolder[];
    /**
     * A list of folder IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of folder names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly parentFolderId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
