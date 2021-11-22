// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.
 *
 * > **NOTE:** You can only share your own custom images to other Alibaba Cloud users.
 *
 * > **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.
 *
 * > **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.
 *
 * > **NOTE:** Available in 1.68.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultImageSharePermission = new alicloud.ecs.ImageSharePermission("default", {
 *     accountId: "1234567890",
 *     imageId: "m-bp1gxyh***",
 * });
 * ```
 * ## Attributes Reference0
 *
 *  The following attributes are exported:
 *
 * * `id` - ID of the image. It formats as `<image_id>:<account_id>`
 *
 * ## Import
 *
 * image can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/imageSharePermission:ImageSharePermission default m-uf66yg1q:123456789
 * ```
 */
export class ImageSharePermission extends pulumi.CustomResource {
    /**
     * Get an existing ImageSharePermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageSharePermissionState, opts?: pulumi.CustomResourceOptions): ImageSharePermission {
        return new ImageSharePermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/imageSharePermission:ImageSharePermission';

    /**
     * Returns true if the given object is an instance of ImageSharePermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageSharePermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageSharePermission.__pulumiType;
    }

    /**
     * Alibaba Cloud Account ID. It is used to share images.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The source image ID.
     */
    public readonly imageId!: pulumi.Output<string>;

    /**
     * Create a ImageSharePermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageSharePermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageSharePermissionArgs | ImageSharePermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageSharePermissionState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
        } else {
            const args = argsOrState as ImageSharePermissionArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ImageSharePermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageSharePermission resources.
 */
export interface ImageSharePermissionState {
    /**
     * Alibaba Cloud Account ID. It is used to share images.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The source image ID.
     */
    imageId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImageSharePermission resource.
 */
export interface ImageSharePermissionArgs {
    /**
     * Alibaba Cloud Account ID. It is used to share images.
     */
    accountId: pulumi.Input<string>;
    /**
     * The source image ID.
     */
    imageId: pulumi.Input<string>;
}
