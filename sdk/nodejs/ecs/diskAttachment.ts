// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DiskAttachment extends pulumi.CustomResource {
    /**
     * Get an existing DiskAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskAttachmentState, opts?: pulumi.CustomResourceOptions): DiskAttachment {
        return new DiskAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/diskAttachment:DiskAttachment';

    /**
     * Returns true if the given object is an instance of DiskAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskAttachment.__pulumiType;
    }

    /**
     * The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * ID of the Disk to be attached.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * ID of the Instance to attach to.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a DiskAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskAttachmentArgs | DiskAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DiskAttachmentState | undefined;
            inputs["deviceName"] = state ? state.deviceName : undefined;
            inputs["diskId"] = state ? state.diskId : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as DiskAttachmentArgs | undefined;
            if (!args || args.diskId === undefined) {
                throw new Error("Missing required property 'diskId'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["diskId"] = args ? args.diskId : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DiskAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskAttachment resources.
 */
export interface DiskAttachmentState {
    /**
     * The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
     * 
     * @deprecated Attribute device_name is deprecated on disk attachment resource. Suggest to remove it from your template.
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * ID of the Disk to be attached.
     */
    readonly diskId?: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    readonly instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DiskAttachment resource.
 */
export interface DiskAttachmentArgs {
    /**
     * The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz.
     * 
     * @deprecated Attribute device_name is deprecated on disk attachment resource. Suggest to remove it from your template.
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * ID of the Disk to be attached.
     */
    readonly diskId: pulumi.Input<string>;
    /**
     * ID of the Instance to attach to.
     */
    readonly instanceId: pulumi.Input<string>;
}
