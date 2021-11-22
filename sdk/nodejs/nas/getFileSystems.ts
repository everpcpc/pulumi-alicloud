// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides FileSystems available to the user.
 *
 * > **NOTE**: Available in 1.35.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const fs = alicloud_nas_file_system_foo.description.apply(description => alicloud.nas.getFileSystems({
 *     descriptionRegex: description,
 *     protocolType: "NFS",
 * }));
 *
 * export const alicloudNasFileSystemsId = fs.systems[0].id;
 * ```
 */
export function getFileSystems(args?: GetFileSystemsArgs, opts?: pulumi.InvokeOptions): Promise<GetFileSystemsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:nas/getFileSystems:getFileSystems", {
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "protocolType": args.protocolType,
        "storageType": args.storageType,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsArgs {
    /**
     * A regex string to filter the results by the ：FileSystem description.
     */
    descriptionRegex?: string;
    /**
     * A list of FileSystemId.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `fileSystemType` is `standard`).
     */
    protocolType?: string;
    /**
     * The storage type of the file system.
     * * Valid values:
     */
    storageType?: string;
}

/**
 * A collection of values returned by getFileSystems.
 */
export interface GetFileSystemsResult {
    readonly descriptionRegex?: string;
    /**
     * A list of FileSystem descriptions.
     */
    readonly descriptions: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of FileSystem Id.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * ProtocolType block of the FileSystem
     */
    readonly protocolType?: string;
    /**
     * StorageType block of the FileSystem.
     */
    readonly storageType?: string;
    /**
     * A list of VPCs. Each element contains the following attributes:
     */
    readonly systems: outputs.nas.GetFileSystemsSystem[];
}

export function getFileSystemsOutput(args?: GetFileSystemsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileSystemsResult> {
    return pulumi.output(args).apply(a => getFileSystems(a, opts))
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsOutputArgs {
    /**
     * A regex string to filter the results by the ：FileSystem description.
     */
    descriptionRegex?: pulumi.Input<string>;
    /**
     * A list of FileSystemId.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The protocol type of the file system.
     * Valid values:
     * `NFS`,
     * `SMB` (Available when the `fileSystemType` is `standard`).
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The storage type of the file system.
     * * Valid values:
     */
    storageType?: pulumi.Input<string>;
}
