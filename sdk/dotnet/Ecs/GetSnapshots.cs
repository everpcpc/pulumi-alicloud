// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetSnapshots
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var snapshots = AliCloud.Ecs.GetSnapshots.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "s-123456890abcdef",
        ///         },
        ///         NameRegex = "tf-testAcc-snapshot",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ##  Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// * `instance_id` - (Optional) The specified instance ID.
        /// * `disk_id` - (Optional) The specified disk ID.
        /// * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values: `true`: Encrypted snapshots. `false`: No encryption attribute limit. Default value: `false`.
        /// * `ids` - (Optional)  A list of snapshot IDs.
        /// * `name_regex` - (Optional) A regex string to filter results by snapshot name.
        /// * `status` - (Optional) The specified snapshot status. Default value: `all`. Optional values:
        ///   * progressing: The snapshots are being created.
        ///   * accomplished: The snapshots are ready to use.
        ///   * failed: The snapshot creation failed.
        ///   * all: All status.
        /// * `type` - (Optional) The snapshot category. Default value: `all`. Optional values:
        ///   * auto: Auto snapshots.
        ///   * user: Manual snapshots.
        ///   * all: Auto and manual snapshots.
        /// * `source_disk_type` - (Optional) The type of source disk:
        ///   * System: The snapshots are created for system disks.
        ///   * Data: The snapshots are created for data disks.
        /// * `usage` - (Optional) The usage of the snapshot:
        ///   * image: The snapshots are used to create custom images.
        ///   * disk: The snapshots are used to CreateDisk.
        ///   * mage_disk: The snapshots are used to create custom images and data disks.
        ///   * none: The snapshots are not used yet.
        /// * `tags` - (Optional) A map of tags assigned to snapshots.
        /// * `output_file` - (Optional) The name of output file that saves the filter results.
        /// </summary>
        public static Task<GetSnapshotsResult> InvokeAsync(GetSnapshotsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotsResult>("alicloud:ecs/getSnapshots:getSnapshots", args ?? new GetSnapshotsArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var snapshots = AliCloud.Ecs.GetSnapshots.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "s-123456890abcdef",
        ///         },
        ///         NameRegex = "tf-testAcc-snapshot",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ##  Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// * `instance_id` - (Optional) The specified instance ID.
        /// * `disk_id` - (Optional) The specified disk ID.
        /// * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values: `true`: Encrypted snapshots. `false`: No encryption attribute limit. Default value: `false`.
        /// * `ids` - (Optional)  A list of snapshot IDs.
        /// * `name_regex` - (Optional) A regex string to filter results by snapshot name.
        /// * `status` - (Optional) The specified snapshot status. Default value: `all`. Optional values:
        ///   * progressing: The snapshots are being created.
        ///   * accomplished: The snapshots are ready to use.
        ///   * failed: The snapshot creation failed.
        ///   * all: All status.
        /// * `type` - (Optional) The snapshot category. Default value: `all`. Optional values:
        ///   * auto: Auto snapshots.
        ///   * user: Manual snapshots.
        ///   * all: Auto and manual snapshots.
        /// * `source_disk_type` - (Optional) The type of source disk:
        ///   * System: The snapshots are created for system disks.
        ///   * Data: The snapshots are created for data disks.
        /// * `usage` - (Optional) The usage of the snapshot:
        ///   * image: The snapshots are used to create custom images.
        ///   * disk: The snapshots are used to CreateDisk.
        ///   * mage_disk: The snapshots are used to create custom images and data disks.
        ///   * none: The snapshots are not used yet.
        /// * `tags` - (Optional) A map of tags assigned to snapshots.
        /// * `output_file` - (Optional) The name of output file that saves the filter results.
        /// </summary>
        public static Output<GetSnapshotsResult> Invoke(GetSnapshotsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotsResult>("alicloud:ecs/getSnapshots:getSnapshots", args ?? new GetSnapshotsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotsArgs : global::Pulumi.InvokeArgs
    {
        [Input("category")]
        public string? Category { get; set; }

        [Input("dryRun")]
        public bool? DryRun { get; set; }

        /// <summary>
        /// Whether the snapshot is encrypted or not.
        /// </summary>
        [Input("encrypted")]
        public bool? Encrypted { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of snapshot IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("kmsKeyId")]
        public string? KmsKeyId { get; set; }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("snapshotLinkId")]
        public string? SnapshotLinkId { get; set; }

        [Input("snapshotName")]
        public string? SnapshotName { get; set; }

        [Input("snapshotType")]
        public string? SnapshotType { get; set; }

        /// <summary>
        /// Source disk attribute. Value range: `System`,`Data`.
        /// </summary>
        [Input("sourceDiskType")]
        public string? SourceDiskType { get; set; }

        /// <summary>
        /// The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the snapshot.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
        /// </summary>
        [Input("usage")]
        public string? Usage { get; set; }

        public GetSnapshotsArgs()
        {
        }
        public static new GetSnapshotsArgs Empty => new GetSnapshotsArgs();
    }

    public sealed class GetSnapshotsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether the snapshot is encrypted or not.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of snapshot IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("snapshotLinkId")]
        public Input<string>? SnapshotLinkId { get; set; }

        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        [Input("snapshotType")]
        public Input<string>? SnapshotType { get; set; }

        /// <summary>
        /// Source disk attribute. Value range: `System`,`Data`.
        /// </summary>
        [Input("sourceDiskType")]
        public Input<string>? SourceDiskType { get; set; }

        /// <summary>
        /// The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the snapshot.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        public GetSnapshotsInvokeArgs()
        {
        }
        public static new GetSnapshotsInvokeArgs Empty => new GetSnapshotsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotsResult
    {
        public readonly string? Category;
        public readonly bool? DryRun;
        /// <summary>
        /// Whether the snapshot is encrypted or not.
        /// </summary>
        public readonly bool? Encrypted;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of snapshot IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? KmsKeyId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of snapshots names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly string? SnapshotLinkId;
        public readonly string? SnapshotName;
        public readonly string? SnapshotType;
        /// <summary>
        /// A list of snapshots. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotsSnapshotResult> Snapshots;
        /// <summary>
        /// Source disk attribute. Value range: `System`,`Data`.
        /// </summary>
        public readonly string? SourceDiskType;
        /// <summary>
        /// The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A map of tags assigned to the snapshot.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string? Type;
        /// <summary>
        /// Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
        /// </summary>
        public readonly string? Usage;

        [OutputConstructor]
        private GetSnapshotsResult(
            string? category,

            bool? dryRun,

            bool? encrypted,

            string id,

            ImmutableArray<string> ids,

            string? kmsKeyId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? resourceGroupId,

            string? snapshotLinkId,

            string? snapshotName,

            string? snapshotType,

            ImmutableArray<Outputs.GetSnapshotsSnapshotResult> snapshots,

            string? sourceDiskType,

            string? status,

            ImmutableDictionary<string, object>? tags,

            string? type,

            string? usage)
        {
            Category = category;
            DryRun = dryRun;
            Encrypted = encrypted;
            Id = id;
            Ids = ids;
            KmsKeyId = kmsKeyId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            SnapshotLinkId = snapshotLinkId;
            SnapshotName = snapshotName;
            SnapshotType = snapshotType;
            Snapshots = snapshots;
            SourceDiskType = sourceDiskType;
            Status = status;
            Tags = tags;
            Type = type;
            Usage = usage;
        }
    }
}
