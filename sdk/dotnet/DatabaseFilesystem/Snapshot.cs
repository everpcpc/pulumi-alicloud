// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DatabaseFilesystem
{
    /// <summary>
    /// Provides a DBFS Snapshot resource.
    /// 
    /// For information about DBFS Snapshot and how to use it, see [What is Snapshot](https://help.aliyun.com/document_detail/149726.html).
    /// 
    /// &gt; **NOTE:** Available in v1.156.0+.
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
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var zoneId = "cn-hangzhou-i";
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = zoneId,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         Description = "tf test",
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///     });
    /// 
    ///     var defaultImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         Owners = "system",
    ///         NameRegex = "^centos_8",
    ///         MostRecent = true,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Ecs.Instance("defaultInstance", new()
    ///     {
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceName = @var.Name,
    ///         InstanceType = "ecs.g7se.large",
    ///         AvailabilityZone = zoneId,
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         SystemDiskCategory = "cloud_essd",
    ///         SecurityGroups = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var defaultDatabasefilesystem_instanceInstance = new AliCloud.DatabaseFilesystem.Instance("defaultDatabasefilesystem/instanceInstance", new()
    ///     {
    ///         Category = "standard",
    ///         ZoneId = defaultInstance.AvailabilityZone,
    ///         PerformanceLevel = "PL1",
    ///         InstanceName = @var.Name,
    ///         Size = 100,
    ///     });
    /// 
    ///     var defaultInstanceAttachment = new AliCloud.DatabaseFilesystem.InstanceAttachment("defaultInstanceAttachment", new()
    ///     {
    ///         EcsId = defaultInstance.Id,
    ///         InstanceId = defaultDatabasefilesystem / instanceInstance.Id,
    ///     });
    /// 
    ///     var example = new AliCloud.DatabaseFilesystem.Snapshot("example", new()
    ///     {
    ///         InstanceId = data.Alicloud_dbfs_instances.Default.Ids[0],
    ///         SnapshotName = "example_value",
    ///         Description = "example_value",
    ///         RetentionDays = 30,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             defaultInstanceAttachment,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// DBFS Snapshot can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:databasefilesystem/snapshot:Snapshot example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:databasefilesystem/snapshot:Snapshot")]
    public partial class Snapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to force deletion of snapshots.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
        /// </summary>
        [Output("snapshotName")]
        public Output<string?> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// The status of the Snapshot.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/snapshot:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:databasefilesystem/snapshot:Snapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, SnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, state, options);
        }
    }

    public sealed class SnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to force deletion of snapshots.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        public SnapshotArgs()
        {
        }
        public static new SnapshotArgs Empty => new SnapshotArgs();
    }

    public sealed class SnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the snapshot. The description must be `2` to `256` characters in length. It must start with a letter, and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to force deletion of snapshots.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The ID of the database file system.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The retention time of the snapshot. Unit: days. Snapshots are automatically released after the retention time expires. Valid values: `1` to `65536`.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The display name of the snapshot. The length is `2` to `128` characters. It must start with a large or small letter or Chinese, and cannot start with `http://` and `https://`. It can contain numbers, colons (:), underscores (_), or hyphens (-). To prevent name conflicts with automatic snapshots, you cannot start with `auto`.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The status of the Snapshot.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SnapshotState()
        {
        }
        public static new SnapshotState Empty => new SnapshotState();
    }
}
