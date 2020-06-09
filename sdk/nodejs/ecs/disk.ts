// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskState, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/disk:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * The Zone to create the disk in.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    public readonly deleteAutoSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean | undefined>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    public readonly enableAutoSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The disk status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskArgs | DiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DiskState | undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["category"] = state ? state.category : undefined;
            inputs["deleteAutoSnapshot"] = state ? state.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableAutoSnapshot"] = state ? state.enableAutoSnapshot : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DiskArgs | undefined;
            if (!args || args.availabilityZone === undefined) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["deleteAutoSnapshot"] = args ? args.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableAutoSnapshot"] = args ? args.enableAutoSnapshot : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Disk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Disk resources.
 */
export interface DiskState {
    /**
     * The Zone to create the disk in.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    readonly category?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    readonly deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    readonly deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    readonly enableAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The disk status.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * The Zone to create the disk in.
     */
    readonly availabilityZone: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    readonly category?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
     */
    readonly deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance: Default value: false.
     */
    readonly deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
     */
    readonly enableAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     * > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    readonly size: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
