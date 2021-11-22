// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECS Disk resource.
 *
 * For information about ECS Disk and how to use it, see [What is Disk](https://www.alibabacloud.com/help/en/doc-detail/25513.htm).
 *
 * > **NOTE:** Available in v1.122.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ecs.EcsDisk("example", {
 *     category: "cloud_efficiency",
 *     description: "Hello ecs disk.",
 *     diskName: "tf-test",
 *     encrypted: true,
 *     kmsKeyId: "2a6767f0-a16c-4679-a60f-13bf*****",
 *     size: 30,
 *     tags: {
 *         Name: "TerraformTest",
 *     },
 *     zoneId: "cn-beijing-b",
 * });
 * ```
 *
 * ## Import
 *
 * ECS Disk can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsDisk:EcsDisk example d-abcd12345
 * ```
 */
export class EcsDisk extends pulumi.CustomResource {
    /**
     * Get an existing EcsDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsDiskState, opts?: pulumi.CustomResourceOptions): EcsDisk {
        return new EcsDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsDisk:EcsDisk';

    /**
     * Returns true if the given object is an instance of EcsDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsDisk.__pulumiType;
    }

    public readonly advancedFeatures!: pulumi.Output<string | undefined>;
    /**
     * Field `availabilityZone` has been deprecated from provider version 1.122.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    public readonly category!: pulumi.Output<string | undefined>;
    public readonly dedicatedBlockStorageClusterId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: `false`.
     */
    public readonly deleteAutoSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the disk is released together with the instance. Default value: `false`.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean | undefined>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with `http://` or `https://`. Default value is `null`.
     */
    public readonly diskName!: pulumi.Output<string>;
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * * `true`: The validity of the request is checked but the request is not made. Check items include the required parameters, request format, service limits, and available ECS resources. If the check fails, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
     * * `false`: The validity of the request is checked. If the check succeeds, a 2xx HTTP status code is returned and the request is made.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to enable creating snapshot automatically. Default value: `false`.
     */
    public readonly enableAutoSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly encryptAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the instance to which the created subscription disk is automatically attached.
     * * After you specify the instance ID, the specified `resourceGroupId`, `tags`, and `kmsKeyId` parameters are ignored.
     * * One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Field `name` has been deprecated from provider version 1.122.0. New field `diskName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Payment method for disk. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`. If you want to change the disk payment type, the `instanceId` is required.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    public readonly performanceLevel!: pulumi.Output<string>;
    /**
     * The Id of resource group which the disk belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    public readonly size!: pulumi.Output<number | undefined>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The disk status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the storage set.
     */
    public readonly storageSetId!: pulumi.Output<string | undefined>;
    /**
     * The number of partitions in the storage set.
     */
    public readonly storageSetPartitionNumber!: pulumi.Output<number | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The type to expand cloud disks. Valid Values: `online`, `offline`. Default to `offline`.
     * * `offline`: After you resize a disk offline, you must restart the instance by using the console or by calling the RebootInstance operation for the resizing operation to take effect. For more information, see Restart the instance and RebootInstance.
     * * `online`: After you resize a disk online, the resizing operation takes effect immediately and you do not need to restart the instance. You can resize ultra disks, standard SSDs, and ESSDs online.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * ID of the free zone to which the disk belongs. One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a EcsDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EcsDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsDiskArgs | EcsDiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsDiskState | undefined;
            inputs["advancedFeatures"] = state ? state.advancedFeatures : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["category"] = state ? state.category : undefined;
            inputs["dedicatedBlockStorageClusterId"] = state ? state.dedicatedBlockStorageClusterId : undefined;
            inputs["deleteAutoSnapshot"] = state ? state.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskName"] = state ? state.diskName : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["enableAutoSnapshot"] = state ? state.enableAutoSnapshot : undefined;
            inputs["encryptAlgorithm"] = state ? state.encryptAlgorithm : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["performanceLevel"] = state ? state.performanceLevel : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["storageSetId"] = state ? state.storageSetId : undefined;
            inputs["storageSetPartitionNumber"] = state ? state.storageSetPartitionNumber : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as EcsDiskArgs | undefined;
            inputs["advancedFeatures"] = args ? args.advancedFeatures : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["dedicatedBlockStorageClusterId"] = args ? args.dedicatedBlockStorageClusterId : undefined;
            inputs["deleteAutoSnapshot"] = args ? args.deleteAutoSnapshot : undefined;
            inputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskName"] = args ? args.diskName : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["enableAutoSnapshot"] = args ? args.enableAutoSnapshot : undefined;
            inputs["encryptAlgorithm"] = args ? args.encryptAlgorithm : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["performanceLevel"] = args ? args.performanceLevel : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["storageSetId"] = args ? args.storageSetId : undefined;
            inputs["storageSetPartitionNumber"] = args ? args.storageSetPartitionNumber : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EcsDisk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsDisk resources.
 */
export interface EcsDiskState {
    advancedFeatures?: pulumi.Input<string>;
    /**
     * Field `availabilityZone` has been deprecated from provider version 1.122.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    category?: pulumi.Input<string>;
    dedicatedBlockStorageClusterId?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: `false`.
     */
    deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default value: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with `http://` or `https://`. Default value is `null`.
     */
    diskName?: pulumi.Input<string>;
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * * `true`: The validity of the request is checked but the request is not made. Check items include the required parameters, request format, service limits, and available ECS resources. If the check fails, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
     * * `false`: The validity of the request is checked. If the check succeeds, a 2xx HTTP status code is returned and the request is made.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Indicates whether to enable creating snapshot automatically. Default value: `false`.
     */
    enableAutoSnapshot?: pulumi.Input<boolean>;
    encryptAlgorithm?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * The ID of the instance to which the created subscription disk is automatically attached.
     * * After you specify the instance ID, the specified `resourceGroupId`, `tags`, and `kmsKeyId` parameters are ignored.
     * * One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.122.0. New field `diskName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Payment method for disk. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`. If you want to change the disk payment type, the `instanceId` is required.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    performanceLevel?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    size?: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The disk status.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the storage set.
     */
    storageSetId?: pulumi.Input<string>;
    /**
     * The number of partitions in the storage set.
     */
    storageSetPartitionNumber?: pulumi.Input<number>;
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type to expand cloud disks. Valid Values: `online`, `offline`. Default to `offline`.
     * * `offline`: After you resize a disk offline, you must restart the instance by using the console or by calling the RebootInstance operation for the resizing operation to take effect. For more information, see Restart the instance and RebootInstance.
     * * `online`: After you resize a disk online, the resizing operation takes effect immediately and you do not need to restart the instance. You can resize ultra disks, standard SSDs, and ESSDs online.
     */
    type?: pulumi.Input<string>;
    /**
     * ID of the free zone to which the disk belongs. One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsDisk resource.
 */
export interface EcsDiskArgs {
    advancedFeatures?: pulumi.Input<string>;
    /**
     * Field `availabilityZone` has been deprecated from provider version 1.122.0. New field `zoneId` instead.
     *
     * @deprecated Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`. Default is `cloudEfficiency`.
     */
    category?: pulumi.Input<string>;
    dedicatedBlockStorageClusterId?: pulumi.Input<string>;
    /**
     * Indicates whether the automatic snapshot is deleted when the disk is released. Default value: `false`.
     */
    deleteAutoSnapshot?: pulumi.Input<boolean>;
    /**
     * Indicates whether the disk is released together with the instance. Default value: `false`.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with `http://` or `https://`. Default value is `null`.
     */
    diskName?: pulumi.Input<string>;
    /**
     * Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
     * * `true`: The validity of the request is checked but the request is not made. Check items include the required parameters, request format, service limits, and available ECS resources. If the check fails, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
     * * `false`: The validity of the request is checked. If the check succeeds, a 2xx HTTP status code is returned and the request is made.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Indicates whether to enable creating snapshot automatically. Default value: `false`.
     */
    enableAutoSnapshot?: pulumi.Input<boolean>;
    encryptAlgorithm?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted, conflict with `snapshotId`.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * The ID of the instance to which the created subscription disk is automatically attached.
     * * After you specify the instance ID, the specified `resourceGroupId`, `tags`, and `kmsKeyId` parameters are ignored.
     * * One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Field `name` has been deprecated from provider version 1.122.0. New field `diskName` instead.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Payment method for disk. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`. If you want to change the disk payment type, the `instanceId` is required.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Specifies the performance level of an ESSD when you create the ESSD. Valid values:                                                       
     * * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
     * * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
     * * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
     */
    performanceLevel?: pulumi.Input<string>;
    /**
     * The Id of resource group which the disk belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
     */
    size?: pulumi.Input<number>;
    /**
     * A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The ID of the storage set.
     */
    storageSetId?: pulumi.Input<string>;
    /**
     * The number of partitions in the storage set.
     */
    storageSetPartitionNumber?: pulumi.Input<number>;
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type to expand cloud disks. Valid Values: `online`, `offline`. Default to `offline`.
     * * `offline`: After you resize a disk offline, you must restart the instance by using the console or by calling the RebootInstance operation for the resizing operation to take effect. For more information, see Restart the instance and RebootInstance.
     * * `online`: After you resize a disk online, the resizing operation takes effect immediately and you do not need to restart the instance. You can resize ultra disks, standard SSDs, and ESSDs online.
     */
    type?: pulumi.Input<string>;
    /**
     * ID of the free zone to which the disk belongs. One of the `zoneId` and `instanceId` must be set but can not be set at the same time.
     */
    zoneId?: pulumi.Input<string>;
}
