// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * An ECI Image Cache can help user to solve the time-consuming problem of image pull. For information about Alicloud ECI Image Cache and how to use it, see [What is Resource Alicloud ECI Image Cache](https://www.alibabacloud.com/help/doc-detail/146891.htm).
 *
 * > **NOTE:** Available in v1.89.0+.
 *
 * > **NOTE:** Each image cache corresponds to a snapshot, and the user does not delete the snapshot directly, otherwise the cache will fail.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.eci.ImageCache("example", {
 *     eipInstanceId: "eip-uf60c7cqb2pcrkgxhxxxx",
 *     imageCacheName: "tf-test",
 *     images: ["registry.cn-beijing.aliyuncs.com/sceneplatform/sae-image-xxxx:latest"],
 *     securityGroupId: "sg-2zeef68b66fxxxx",
 *     vswitchId: "vsw-2zef9k7ng82xxxx",
 * });
 * ```
 *
 * ## Import
 *
 * ECI Image Cache can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eci/imageCache:ImageCache example abc123456
 * ```
 */
export class ImageCache extends pulumi.CustomResource {
    /**
     * Get an existing ImageCache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageCacheState, opts?: pulumi.CustomResourceOptions): ImageCache {
        return new ImageCache(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eci/imageCache:ImageCache';

    /**
     * Returns true if the given object is an instance of ImageCache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageCache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageCache.__pulumiType;
    }

    /**
     * The ID of the container group job that is used to create the image cache.
     * * `status` -The status of the image cache.
     */
    public /*out*/ readonly containerGroupId!: pulumi.Output<string>;
    /**
     * The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
     */
    public readonly eipInstanceId!: pulumi.Output<string | undefined>;
    /**
     * The name of the image cache.
     */
    public readonly imageCacheName!: pulumi.Output<string>;
    /**
     * The size of the image cache. Default to `20`. Unit: GiB.
     */
    public readonly imageCacheSize!: pulumi.Output<number | undefined>;
    /**
     * The Image Registry parameters about the image to be cached.
     */
    public readonly imageRegistryCredentials!: pulumi.Output<outputs.eci.ImageCacheImageRegistryCredential[] | undefined>;
    /**
     * The images to be cached. The image name must be versioned.
     */
    public readonly images!: pulumi.Output<string[]>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    /**
     * The ID of the security group. You do not need to specify the same security group as the container group.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The zone id to cache image.
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a ImageCache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageCacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageCacheArgs | ImageCacheState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ImageCacheState | undefined;
            inputs["containerGroupId"] = state ? state.containerGroupId : undefined;
            inputs["eipInstanceId"] = state ? state.eipInstanceId : undefined;
            inputs["imageCacheName"] = state ? state.imageCacheName : undefined;
            inputs["imageCacheSize"] = state ? state.imageCacheSize : undefined;
            inputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            inputs["images"] = state ? state.images : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["retentionDays"] = state ? state.retentionDays : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ImageCacheArgs | undefined;
            if ((!args || args.imageCacheName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'imageCacheName'");
            }
            if ((!args || args.images === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'images'");
            }
            if ((!args || args.securityGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.vswitchId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'vswitchId'");
            }
            inputs["eipInstanceId"] = args ? args.eipInstanceId : undefined;
            inputs["imageCacheName"] = args ? args.imageCacheName : undefined;
            inputs["imageCacheSize"] = args ? args.imageCacheSize : undefined;
            inputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            inputs["images"] = args ? args.images : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["retentionDays"] = args ? args.retentionDays : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["containerGroupId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ImageCache.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageCache resources.
 */
export interface ImageCacheState {
    /**
     * The ID of the container group job that is used to create the image cache.
     * * `status` -The status of the image cache.
     */
    readonly containerGroupId?: pulumi.Input<string>;
    /**
     * The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
     */
    readonly eipInstanceId?: pulumi.Input<string>;
    /**
     * The name of the image cache.
     */
    readonly imageCacheName?: pulumi.Input<string>;
    /**
     * The size of the image cache. Default to `20`. Unit: GiB.
     */
    readonly imageCacheSize?: pulumi.Input<number>;
    /**
     * The Image Registry parameters about the image to be cached.
     */
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.ImageCacheImageRegistryCredential>[]>;
    /**
     * The images to be cached. The image name must be versioned.
     */
    readonly images?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the resource group.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
     */
    readonly retentionDays?: pulumi.Input<number>;
    /**
     * The ID of the security group. You do not need to specify the same security group as the container group.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    /**
     * The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The zone id to cache image.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImageCache resource.
 */
export interface ImageCacheArgs {
    /**
     * The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)
     */
    readonly eipInstanceId?: pulumi.Input<string>;
    /**
     * The name of the image cache.
     */
    readonly imageCacheName: pulumi.Input<string>;
    /**
     * The size of the image cache. Default to `20`. Unit: GiB.
     */
    readonly imageCacheSize?: pulumi.Input<number>;
    /**
     * The Image Registry parameters about the image to be cached.
     */
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.ImageCacheImageRegistryCredential>[]>;
    /**
     * The images to be cached. The image name must be versioned.
     */
    readonly images: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the resource group.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.
     */
    readonly retentionDays?: pulumi.Input<number>;
    /**
     * The ID of the security group. You do not need to specify the same security group as the container group.
     */
    readonly securityGroupId: pulumi.Input<string>;
    /**
     * The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.
     */
    readonly vswitchId: pulumi.Input<string>;
    /**
     * The zone id to cache image.
     */
    readonly zoneId?: pulumi.Input<string>;
}
