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
    /// Provides a ECS Snapshot resource.
    /// 
    /// For information about ECS Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/25524.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.120.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new AliCloud.Ecs.EcsSnapshot("default", new()
    ///     {
    ///         Category = "standard",
    ///         Description = "Test For Terraform",
    ///         DiskId = "d-gw8csgxxxxxxxxx",
    ///         RetentionDays = 20,
    ///         SnapshotName = "tf-test",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "Acceptance-test" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Snapshot can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/ecsSnapshot:EcsSnapshot example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/ecsSnapshot:EcsSnapshot")]
    public partial class EcsSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The category of the snapshot. Valid Values: `standard` and `flash`.
        /// </summary>
        [Output("category")]
        public Output<string?> Category { get; private set; } = null!;

        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to forcibly delete the snapshot that has been used to create disks.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the instant access feature.
        /// </summary>
        [Output("instantAccess")]
        public Output<bool?> InstantAccess { get; private set; } = null!;

        /// <summary>
        /// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
        /// </summary>
        [Output("instantAccessRetentionDays")]
        public Output<int?> InstantAccessRetentionDays { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.120.0. New field `snapshot_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The retention period of the snapshot.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Output("snapshotName")]
        public Output<string> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// The status of snapshot.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the snapshot.
        /// 
        /// &gt; **NOTE:** If `force` is true, After an snapshot is deleted, the disks created from this snapshot cannot be re-initialized.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EcsSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EcsSnapshot(string name, EcsSnapshotArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, args ?? new EcsSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EcsSnapshot(string name, Input<string> id, EcsSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsSnapshot:EcsSnapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EcsSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EcsSnapshot Get(string name, Input<string> id, EcsSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new EcsSnapshot(name, id, state, options);
        }
    }

    public sealed class EcsSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the snapshot. Valid Values: `standard` and `flash`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        /// <summary>
        /// Specifies whether to forcibly delete the snapshot that has been used to create disks.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Specifies whether to enable the instant access feature.
        /// </summary>
        [Input("instantAccess")]
        public Input<bool>? InstantAccess { get; set; }

        /// <summary>
        /// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
        /// </summary>
        [Input("instantAccessRetentionDays")]
        public Input<int>? InstantAccessRetentionDays { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.120.0. New field `snapshot_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The retention period of the snapshot.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the snapshot.
        /// 
        /// &gt; **NOTE:** If `force` is true, After an snapshot is deleted, the disks created from this snapshot cannot be re-initialized.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsSnapshotArgs()
        {
        }
        public static new EcsSnapshotArgs Empty => new EcsSnapshotArgs();
    }

    public sealed class EcsSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the snapshot. Valid Values: `standard` and `flash`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the disk.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// Specifies whether to forcibly delete the snapshot that has been used to create disks.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Specifies whether to enable the instant access feature.
        /// </summary>
        [Input("instantAccess")]
        public Input<bool>? InstantAccess { get; set; }

        /// <summary>
        /// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
        /// </summary>
        [Input("instantAccessRetentionDays")]
        public Input<int>? InstantAccessRetentionDays { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.120.0. New field `snapshot_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The retention period of the snapshot.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The status of snapshot.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the snapshot.
        /// 
        /// &gt; **NOTE:** If `force` is true, After an snapshot is deleted, the disks created from this snapshot cannot be re-initialized.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsSnapshotState()
        {
        }
        public static new EcsSnapshotState Empty => new EcsSnapshotState();
    }
}
