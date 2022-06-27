// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new ECS disk.
    ///         var ecsDisk = new AliCloud.Ecs.Disk("ecsDisk", new AliCloud.Ecs.DiskArgs
    ///         {
    ///             AvailabilityZone = "cn-beijing-b",
    ///             Category = "cloud_efficiency",
    ///             Description = "Hello ecs disk.",
    ///             Encrypted = true,
    ///             KmsKeyId = "2a6767f0-a16c-4679-a60f-13bf*****",
    ///             Size = 30,
    ///             Tags = 
    ///             {
    ///                 { "Name", "TerraformTest" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud disk can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/disk:Disk example d-abc12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/disk:Disk")]
    public partial class Disk : Pulumi.CustomResource
    {
        [Output("advancedFeatures")]
        public Output<string?> AdvancedFeatures { get; private set; } = null!;

        /// <summary>
        /// The Zone to create the disk in.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`. Default is `cloud_efficiency`.
        /// </summary>
        [Output("category")]
        public Output<string?> Category { get; private set; } = null!;

        [Output("dedicatedBlockStorageClusterId")]
        public Output<string?> DedicatedBlockStorageClusterId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
        /// </summary>
        [Output("deleteAutoSnapshot")]
        public Output<bool?> DeleteAutoSnapshot { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the disk is released together with the instance: Default value: false.
        /// </summary>
        [Output("deleteWithInstance")]
        public Output<bool> DeleteWithInstance { get; private set; } = null!;

        /// <summary>
        /// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("diskName")]
        public Output<string> DiskName { get; private set; } = null!;

        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
        /// </summary>
        [Output("enableAutoSnapshot")]
        public Output<bool> EnableAutoSnapshot { get; private set; } = null!;

        [Output("encryptAlgorithm")]
        public Output<string?> EncryptAlgorithm { get; private set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted, conflict with `snapshot_id`.
        /// </summary>
        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
        /// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
        /// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
        /// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
        /// </summary>
        [Output("performanceLevel")]
        public Output<string> PerformanceLevel { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
        /// </summary>
        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;

        /// <summary>
        /// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The disk status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("storageSetId")]
        public Output<string?> StorageSetId { get; private set; } = null!;

        [Output("storageSetPartitionNumber")]
        public Output<int?> StorageSetPartitionNumber { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/disk:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/disk:Disk", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, state, options);
        }
    }

    public sealed class DiskArgs : Pulumi.ResourceArgs
    {
        [Input("advancedFeatures")]
        public Input<string>? AdvancedFeatures { get; set; }

        /// <summary>
        /// The Zone to create the disk in.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`. Default is `cloud_efficiency`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("dedicatedBlockStorageClusterId")]
        public Input<string>? DedicatedBlockStorageClusterId { get; set; }

        /// <summary>
        /// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
        /// </summary>
        [Input("deleteAutoSnapshot")]
        public Input<bool>? DeleteAutoSnapshot { get; set; }

        /// <summary>
        /// Indicates whether the disk is released together with the instance: Default value: false.
        /// </summary>
        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        /// <summary>
        /// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
        /// </summary>
        [Input("enableAutoSnapshot")]
        public Input<bool>? EnableAutoSnapshot { get; set; }

        [Input("encryptAlgorithm")]
        public Input<string>? EncryptAlgorithm { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted, conflict with `snapshot_id`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
        /// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
        /// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
        /// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
        /// </summary>
        [Input("performanceLevel")]
        public Input<string>? PerformanceLevel { get; set; }

        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("storageSetId")]
        public Input<string>? StorageSetId { get; set; }

        [Input("storageSetPartitionNumber")]
        public Input<int>? StorageSetPartitionNumber { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DiskArgs()
        {
        }
    }

    public sealed class DiskState : Pulumi.ResourceArgs
    {
        [Input("advancedFeatures")]
        public Input<string>? AdvancedFeatures { get; set; }

        /// <summary>
        /// The Zone to create the disk in.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Category of the disk. Valid values are `cloud`, `cloud_efficiency`, `cloud_ssd`, `cloud_essd`. Default is `cloud_efficiency`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("dedicatedBlockStorageClusterId")]
        public Input<string>? DedicatedBlockStorageClusterId { get; set; }

        /// <summary>
        /// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
        /// </summary>
        [Input("deleteAutoSnapshot")]
        public Input<bool>? DeleteAutoSnapshot { get; set; }

        /// <summary>
        /// Indicates whether the disk is released together with the instance: Default value: false.
        /// </summary>
        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        /// <summary>
        /// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
        /// </summary>
        [Input("enableAutoSnapshot")]
        public Input<bool>? EnableAutoSnapshot { get; set; }

        [Input("encryptAlgorithm")]
        public Input<string>? EncryptAlgorithm { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted, conflict with `snapshot_id`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:                                                       
        /// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
        /// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
        /// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
        /// </summary>
        [Input("performanceLevel")]
        public Input<string>? PerformanceLevel { get; set; }

        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// &gt; **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloud_efficiency` and `cloud_ssd` disk.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The disk status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("storageSetId")]
        public Input<string>? StorageSetId { get; set; }

        [Input("storageSetPartitionNumber")]
        public Input<int>? StorageSetPartitionNumber { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DiskState()
        {
        }
    }
}
