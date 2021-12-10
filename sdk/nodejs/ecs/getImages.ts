// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
 * other public images and the ones available on the image market.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const imagesDs = pulumi.output(alicloud.ecs.getImages({
 *     nameRegex: "^centos_6",
 *     owners: "system",
 * }));
 *
 * export const firstImageId = imagesDs.images[0].id;
 * ```
 */
export function getImages(args?: GetImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ecs/getImages:getImages", {
        "actionType": args.actionType,
        "architecture": args.architecture,
        "dryRun": args.dryRun,
        "imageFamily": args.imageFamily,
        "imageId": args.imageId,
        "imageName": args.imageName,
        "instanceType": args.instanceType,
        "isSupportCloudInit": args.isSupportCloudInit,
        "isSupportIoOptimized": args.isSupportIoOptimized,
        "mostRecent": args.mostRecent,
        "nameRegex": args.nameRegex,
        "osType": args.osType,
        "outputFile": args.outputFile,
        "owners": args.owners,
        "resourceGroupId": args.resourceGroupId,
        "snapshotId": args.snapshotId,
        "status": args.status,
        "tags": args.tags,
        "usage": args.usage,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesArgs {
    /**
     * The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:                                                
     * * `CreateEcs`: instance creation.
     * * `ChangeOS`: replacement of the system disk or operating system.
     */
    actionType?: string;
    /**
     * The image architecture. Valid values: `i386` and `x8664`.
     */
    architecture?: string;
    /**
     * Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:                                           
     * * `true`: The validity of the request is checked but resources are not queried. Check items include whether your AccessKey pair is valid, whether RAM users are authorized, and whether the required parameters are specified. If the check fails, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
     * * `false`: The validity of the request is checked, and a 2XX HTTP status code is returned and resources are queried if the check succeeds.
     */
    dryRun?: boolean;
    /**
     * The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
     */
    imageFamily?: string;
    /**
     * The ID of the image.
     */
    imageId?: string;
    /**
     * The name of the image.
     */
    imageName?: string;
    /**
     * The instance type for which the image can be used.
     */
    instanceType?: string;
    /**
     * Specifies whether the image supports cloud-init.
     */
    isSupportCloudInit?: boolean;
    /**
     * Specifies whether the image can be used on I/O optimized instances.
     */
    isSupportIoOptimized?: boolean;
    /**
     * If more than one result are returned, select the most recent one.
     */
    mostRecent?: boolean;
    /**
     * A regex string to filter resulting images by name.
     */
    nameRegex?: string;
    /**
     * The operating system type of the image. Valid values: `windows` and `linux`.
     */
    osType?: string;
    outputFile?: string;
    /**
     * Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
     */
    owners?: string;
    /**
     * The ID of the resource group to which the custom image belongs.
     */
    resourceGroupId?: string;
    /**
     * The ID of the snapshot used to create the custom image.
     */
    snapshotId?: string;
    /**
     * The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values: 
     * * `Creating`: The image is being created.
     * * `Waiting`: The image is waiting to be processed.
     * * `Available`: The image is available.
     * * `UnAvailable`: The image is unavailable.
     * * `CreateFailed`: The image failed to be created.
     * * `Deprecated`: The image is discontinued.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
    /**
     * Specifies whether to check the validity of the request without actually making the request. Valid values:                                           
     * * `instance`: The image is already in use and running on an ECS instance.
     * * `none`: The image is not in use.
     */
    usage?: string;
}

/**
 * A collection of values returned by getImages.
 */
export interface GetImagesResult {
    readonly actionType?: string;
    /**
     * Platform type of the image system: i386 or x86_64.
     */
    readonly architecture?: string;
    readonly dryRun?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of image IDs.
     */
    readonly ids: string[];
    readonly imageFamily?: string;
    readonly imageId?: string;
    readonly imageName?: string;
    /**
     * A list of images. Each element contains the following attributes:
     */
    readonly images: outputs.ecs.GetImagesImage[];
    readonly instanceType?: string;
    readonly isSupportCloudInit?: boolean;
    readonly isSupportIoOptimized?: boolean;
    readonly mostRecent?: boolean;
    readonly nameRegex?: string;
    readonly osType?: string;
    readonly outputFile?: string;
    readonly owners?: string;
    readonly resourceGroupId?: string;
    /**
     * Snapshot ID.
     */
    readonly snapshotId?: string;
    /**
     * Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
     */
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    readonly usage?: string;
}

export function getImagesOutput(args?: GetImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImagesResult> {
    return pulumi.output(args).apply(a => getImages(a, opts))
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesOutputArgs {
    /**
     * The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:                                                
     * * `CreateEcs`: instance creation.
     * * `ChangeOS`: replacement of the system disk or operating system.
     */
    actionType?: pulumi.Input<string>;
    /**
     * The image architecture. Valid values: `i386` and `x8664`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:                                           
     * * `true`: The validity of the request is checked but resources are not queried. Check items include whether your AccessKey pair is valid, whether RAM users are authorized, and whether the required parameters are specified. If the check fails, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
     * * `false`: The validity of the request is checked, and a 2XX HTTP status code is returned and resources are queried if the check succeeds.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
     */
    imageFamily?: pulumi.Input<string>;
    /**
     * The ID of the image.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The name of the image.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The instance type for which the image can be used.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Specifies whether the image supports cloud-init.
     */
    isSupportCloudInit?: pulumi.Input<boolean>;
    /**
     * Specifies whether the image can be used on I/O optimized instances.
     */
    isSupportIoOptimized?: pulumi.Input<boolean>;
    /**
     * If more than one result are returned, select the most recent one.
     */
    mostRecent?: pulumi.Input<boolean>;
    /**
     * A regex string to filter resulting images by name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The operating system type of the image. Valid values: `windows` and `linux`.
     */
    osType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
     */
    owners?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the custom image belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The ID of the snapshot used to create the custom image.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values: 
     * * `Creating`: The image is being created.
     * * `Waiting`: The image is waiting to be processed.
     * * `Available`: The image is available.
     * * `UnAvailable`: The image is unavailable.
     * * `CreateFailed`: The image failed to be created.
     * * `Deprecated`: The image is discontinued.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies whether to check the validity of the request without actually making the request. Valid values:                                           
     * * `instance`: The image is already in use and running on an ECS instance.
     * * `none`: The image is not in use.
     */
    usage?: pulumi.Input<string>;
}
