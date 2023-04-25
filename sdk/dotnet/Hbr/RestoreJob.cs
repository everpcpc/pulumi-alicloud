// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    /// <summary>
    /// Provides a Hybrid Backup Recovery (HBR) Restore Job resource.
    /// 
    /// For information about Hybrid Backup Recovery (HBR) Restore Job and how to use it, see [What is Restore Job](https://www.alibabacloud.com/help/doc-detail/186575.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.133.0+.
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
    ///     var defaultEcsBackupPlans = AliCloud.Hbr.GetEcsBackupPlans.Invoke(new()
    ///     {
    ///         NameRegex = "plan-tf-used-dont-delete",
    ///     });
    /// 
    ///     var defaultOssBackupPlans = AliCloud.Hbr.GetOssBackupPlans.Invoke(new()
    ///     {
    ///         NameRegex = "plan-tf-used-dont-delete",
    ///     });
    /// 
    ///     var defaultNasBackupPlans = AliCloud.Hbr.GetNasBackupPlans.Invoke(new()
    ///     {
    ///         NameRegex = "plan-tf-used-dont-delete",
    ///     });
    /// 
    ///     var ecsSnapshots = AliCloud.Hbr.GetSnapshots.Invoke(new()
    ///     {
    ///         SourceType = "ECS_FILE",
    ///         VaultId = defaultEcsBackupPlans.Apply(getEcsBackupPlansResult =&gt; getEcsBackupPlansResult.Plans[0]?.VaultId),
    ///         InstanceId = defaultEcsBackupPlans.Apply(getEcsBackupPlansResult =&gt; getEcsBackupPlansResult.Plans[0]?.InstanceId),
    ///     });
    /// 
    ///     var ossSnapshots = AliCloud.Hbr.GetSnapshots.Invoke(new()
    ///     {
    ///         SourceType = "OSS",
    ///         VaultId = defaultOssBackupPlans.Apply(getOssBackupPlansResult =&gt; getOssBackupPlansResult.Plans[0]?.VaultId),
    ///         Bucket = defaultOssBackupPlans.Apply(getOssBackupPlansResult =&gt; getOssBackupPlansResult.Plans[0]?.Bucket),
    ///     });
    /// 
    ///     var nasSnapshots = AliCloud.Hbr.GetSnapshots.Invoke(new()
    ///     {
    ///         SourceType = "NAS",
    ///         VaultId = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.VaultId),
    ///         FileSystemId = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.FileSystemId),
    ///         CreateTime = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.CreateTime),
    ///     });
    /// 
    ///     var nasJob = new AliCloud.Hbr.RestoreJob("nasJob", new()
    ///     {
    ///         SnapshotHash = nasSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotHash),
    ///         VaultId = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.VaultId),
    ///         SourceType = "NAS",
    ///         RestoreType = "NAS",
    ///         SnapshotId = nasSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotId),
    ///         TargetFileSystemId = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.FileSystemId),
    ///         TargetCreateTime = defaultNasBackupPlans.Apply(getNasBackupPlansResult =&gt; getNasBackupPlansResult.Plans[0]?.CreateTime),
    ///         TargetPath = "/",
    ///         Options = @"    {""includes"":[], ""excludes"":[]}
    /// ",
    ///     });
    /// 
    ///     var ossJob = new AliCloud.Hbr.RestoreJob("ossJob", new()
    ///     {
    ///         SnapshotHash = ossSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotHash),
    ///         VaultId = defaultOssBackupPlans.Apply(getOssBackupPlansResult =&gt; getOssBackupPlansResult.Plans[0]?.VaultId),
    ///         SourceType = "OSS",
    ///         RestoreType = "OSS",
    ///         SnapshotId = ossSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotId),
    ///         TargetBucket = defaultOssBackupPlans.Apply(getOssBackupPlansResult =&gt; getOssBackupPlansResult.Plans[0]?.Bucket),
    ///         TargetPrefix = "",
    ///         Options = @"    {""includes"":[], ""excludes"":[]}
    /// ",
    ///     });
    /// 
    ///     var ecsJob = new AliCloud.Hbr.RestoreJob("ecsJob", new()
    ///     {
    ///         SnapshotHash = ecsSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotHash),
    ///         VaultId = defaultEcsBackupPlans.Apply(getEcsBackupPlansResult =&gt; getEcsBackupPlansResult.Plans[0]?.VaultId),
    ///         SourceType = "ECS_FILE",
    ///         RestoreType = "ECS_FILE",
    ///         SnapshotId = ecsSnapshots.Apply(getSnapshotsResult =&gt; getSnapshotsResult.Snapshots[0]?.SnapshotId),
    ///         TargetInstanceId = defaultEcsBackupPlans.Apply(getEcsBackupPlansResult =&gt; getEcsBackupPlansResult.Plans[0]?.InstanceId),
    ///         TargetPath = "/",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **NOTE:** This resource can only be created, cannot be modified or deleted. Therefore, any modification of the resource attribute will not affect exist resource.
    /// 
    /// ## Import
    /// 
    /// Hybrid Backup Recovery (HBR) Restore Job can be imported using the id. Format to `&lt;restore_job_id&gt;:&lt;restore_type&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:hbr/restoreJob:RestoreJob example your_restore_job_id:your_restore_type
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:hbr/restoreJob:RestoreJob")]
    public partial class RestoreJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The role name created in the original account RAM backup by the cross account managed by the current account.
        /// </summary>
        [Output("crossAccountRoleName")]
        public Output<string?> CrossAccountRoleName { get; private set; } = null!;

        /// <summary>
        /// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        /// </summary>
        [Output("crossAccountType")]
        public Output<string> CrossAccountType { get; private set; } = null!;

        /// <summary>
        /// The original account ID of the cross account backup managed by the current account.
        /// </summary>
        [Output("crossAccountUserId")]
        public Output<int?> CrossAccountUserId { get; private set; } = null!;

        /// <summary>
        /// The exclude path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Output("exclude")]
        public Output<string?> Exclude { get; private set; } = null!;

        /// <summary>
        /// The include path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while source_type equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Output("include")]
        public Output<string?> Include { get; private set; } = null!;

        /// <summary>
        /// Recovery options. **NOTE:** Required while source_type equals `OSS` or `NAS`, invalid while source_type equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018}`.
        /// </summary>
        [Output("options")]
        public Output<string?> Options { get; private set; } = null!;

        /// <summary>
        /// The details about the Tablestore instance. See the following `Block ots_detail`.
        /// </summary>
        [Output("otsDetail")]
        public Output<Outputs.RestoreJobOtsDetail> OtsDetail { get; private set; } = null!;

        /// <summary>
        /// Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
        /// </summary>
        [Output("restoreJobId")]
        public Output<string> RestoreJobId { get; private set; } = null!;

        /// <summary>
        /// The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
        /// </summary>
        [Output("restoreType")]
        public Output<string> RestoreType { get; private set; } = null!;

        /// <summary>
        /// The hashcode of Snapshot.
        /// </summary>
        [Output("snapshotHash")]
        public Output<string> SnapshotHash { get; private set; } = null!;

        /// <summary>
        /// The ID of Snapshot.
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
        /// </summary>
        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        /// <summary>
        /// The Restore Job Status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The target name of OSS bucket. **NOTE:** Required while source_type equals `OSS`,
        /// </summary>
        [Output("targetBucket")]
        public Output<string?> TargetBucket { get; private set; } = null!;

        /// <summary>
        /// The target client ID.
        /// </summary>
        [Output("targetClientId")]
        public Output<string?> TargetClientId { get; private set; } = null!;

        /// <summary>
        /// The creation time of destination File System. **NOTE:** While source_type equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        /// </summary>
        [Output("targetCreateTime")]
        public Output<string?> TargetCreateTime { get; private set; } = null!;

        /// <summary>
        /// The target data source ID.
        /// </summary>
        [Output("targetDataSourceId")]
        public Output<string?> TargetDataSourceId { get; private set; } = null!;

        /// <summary>
        /// The ID of destination File System. **NOTE:** Required while source_type equals `NAS`
        /// </summary>
        [Output("targetFileSystemId")]
        public Output<string?> TargetFileSystemId { get; private set; } = null!;

        /// <summary>
        /// The target ID of ECS instance. **NOTE:** Required while source_type equals `ECS_FILE`
        /// </summary>
        [Output("targetInstanceId")]
        public Output<string?> TargetInstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the Table store instance to which you want to restore data.**WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Output("targetInstanceName")]
        public Output<string?> TargetInstanceName { get; private set; } = null!;

        /// <summary>
        /// The target file path of (ECS) instance. **WARNING:** Required while source_type equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Output("targetPath")]
        public Output<string?> TargetPath { get; private set; } = null!;

        /// <summary>
        /// The target prefix of the OSS object. **WARNING:** Required while source_type equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Output("targetPrefix")]
        public Output<string?> TargetPrefix { get; private set; } = null!;

        /// <summary>
        /// The name of the table that stores the restored data. **WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Output("targetTableName")]
        public Output<string?> TargetTableName { get; private set; } = null!;

        /// <summary>
        /// The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while source_type equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the End time of the snapshot.
        /// </summary>
        [Output("targetTime")]
        public Output<string?> TargetTime { get; private set; } = null!;

        /// <summary>
        /// The full machine backup details.
        /// </summary>
        [Output("udmDetail")]
        public Output<string?> UdmDetail { get; private set; } = null!;

        /// <summary>
        /// The ID of backup vault.
        /// </summary>
        [Output("vaultId")]
        public Output<string> VaultId { get; private set; } = null!;


        /// <summary>
        /// Create a RestoreJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestoreJob(string name, RestoreJobArgs args, CustomResourceOptions? options = null)
            : base("alicloud:hbr/restoreJob:RestoreJob", name, args ?? new RestoreJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestoreJob(string name, Input<string> id, RestoreJobState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:hbr/restoreJob:RestoreJob", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestoreJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestoreJob Get(string name, Input<string> id, RestoreJobState? state = null, CustomResourceOptions? options = null)
        {
            return new RestoreJob(name, id, state, options);
        }
    }

    public sealed class RestoreJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The role name created in the original account RAM backup by the cross account managed by the current account.
        /// </summary>
        [Input("crossAccountRoleName")]
        public Input<string>? CrossAccountRoleName { get; set; }

        /// <summary>
        /// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        /// </summary>
        [Input("crossAccountType")]
        public Input<string>? CrossAccountType { get; set; }

        /// <summary>
        /// The original account ID of the cross account backup managed by the current account.
        /// </summary>
        [Input("crossAccountUserId")]
        public Input<int>? CrossAccountUserId { get; set; }

        /// <summary>
        /// The exclude path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        /// <summary>
        /// The include path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while source_type equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("include")]
        public Input<string>? Include { get; set; }

        /// <summary>
        /// Recovery options. **NOTE:** Required while source_type equals `OSS` or `NAS`, invalid while source_type equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018}`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// The details about the Tablestore instance. See the following `Block ots_detail`.
        /// </summary>
        [Input("otsDetail")]
        public Input<Inputs.RestoreJobOtsDetailArgs>? OtsDetail { get; set; }

        /// <summary>
        /// Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
        /// </summary>
        [Input("restoreJobId")]
        public Input<string>? RestoreJobId { get; set; }

        /// <summary>
        /// The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
        /// </summary>
        [Input("restoreType", required: true)]
        public Input<string> RestoreType { get; set; } = null!;

        /// <summary>
        /// The hashcode of Snapshot.
        /// </summary>
        [Input("snapshotHash", required: true)]
        public Input<string> SnapshotHash { get; set; } = null!;

        /// <summary>
        /// The ID of Snapshot.
        /// </summary>
        [Input("snapshotId", required: true)]
        public Input<string> SnapshotId { get; set; } = null!;

        /// <summary>
        /// The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        /// <summary>
        /// The target name of OSS bucket. **NOTE:** Required while source_type equals `OSS`,
        /// </summary>
        [Input("targetBucket")]
        public Input<string>? TargetBucket { get; set; }

        /// <summary>
        /// The target client ID.
        /// </summary>
        [Input("targetClientId")]
        public Input<string>? TargetClientId { get; set; }

        /// <summary>
        /// The creation time of destination File System. **NOTE:** While source_type equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        /// </summary>
        [Input("targetCreateTime")]
        public Input<string>? TargetCreateTime { get; set; }

        /// <summary>
        /// The target data source ID.
        /// </summary>
        [Input("targetDataSourceId")]
        public Input<string>? TargetDataSourceId { get; set; }

        /// <summary>
        /// The ID of destination File System. **NOTE:** Required while source_type equals `NAS`
        /// </summary>
        [Input("targetFileSystemId")]
        public Input<string>? TargetFileSystemId { get; set; }

        /// <summary>
        /// The target ID of ECS instance. **NOTE:** Required while source_type equals `ECS_FILE`
        /// </summary>
        [Input("targetInstanceId")]
        public Input<string>? TargetInstanceId { get; set; }

        /// <summary>
        /// The name of the Table store instance to which you want to restore data.**WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Input("targetInstanceName")]
        public Input<string>? TargetInstanceName { get; set; }

        /// <summary>
        /// The target file path of (ECS) instance. **WARNING:** Required while source_type equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("targetPath")]
        public Input<string>? TargetPath { get; set; }

        /// <summary>
        /// The target prefix of the OSS object. **WARNING:** Required while source_type equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("targetPrefix")]
        public Input<string>? TargetPrefix { get; set; }

        /// <summary>
        /// The name of the table that stores the restored data. **WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Input("targetTableName")]
        public Input<string>? TargetTableName { get; set; }

        /// <summary>
        /// The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while source_type equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the End time of the snapshot.
        /// </summary>
        [Input("targetTime")]
        public Input<string>? TargetTime { get; set; }

        /// <summary>
        /// The full machine backup details.
        /// </summary>
        [Input("udmDetail")]
        public Input<string>? UdmDetail { get; set; }

        /// <summary>
        /// The ID of backup vault.
        /// </summary>
        [Input("vaultId", required: true)]
        public Input<string> VaultId { get; set; } = null!;

        public RestoreJobArgs()
        {
        }
        public static new RestoreJobArgs Empty => new RestoreJobArgs();
    }

    public sealed class RestoreJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The role name created in the original account RAM backup by the cross account managed by the current account.
        /// </summary>
        [Input("crossAccountRoleName")]
        public Input<string>? CrossAccountRoleName { get; set; }

        /// <summary>
        /// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
        /// </summary>
        [Input("crossAccountType")]
        public Input<string>? CrossAccountType { get; set; }

        /// <summary>
        /// The original account ID of the cross account backup managed by the current account.
        /// </summary>
        [Input("crossAccountUserId")]
        public Input<int>? CrossAccountUserId { get; set; }

        /// <summary>
        /// The exclude path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/excludePath]`, up to 255 characters. **WARNING:** If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        /// <summary>
        /// The include path. **NOTE:** Invalid while source_type equals `OSS` or `NAS`. It's a json string with format:`["/includePath"]`, Up to 255 characters. **WARNING:** The field is required while source_type equals `OTS_TABLE` which means source table name. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("include")]
        public Input<string>? Include { get; set; }

        /// <summary>
        /// Recovery options. **NOTE:** Required while source_type equals `OSS` or `NAS`, invalid while source_type equals `ECS_FILE`. It's a json string with format:`"{"includes":[],"excludes":[]}",`. Recovery options. When restores OTS_TABLE and real target time is the rangEnd time of the snapshot, it should be a string with format: `{"UI_TargetTime":1650032529018}`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// The details about the Tablestore instance. See the following `Block ots_detail`.
        /// </summary>
        [Input("otsDetail")]
        public Input<Inputs.RestoreJobOtsDetailGetArgs>? OtsDetail { get; set; }

        /// <summary>
        /// Restore Job ID. It's the unique key of this resource, if you want to set this argument by yourself, you must specify a unique keyword that never appears.
        /// </summary>
        [Input("restoreJobId")]
        public Input<string>? RestoreJobId { get; set; }

        /// <summary>
        /// The type of recovery destination. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS_ROLLBACK`. **Note**: Currently, there is a one-to-one correspondence between the data source type with the recovery destination type.
        /// </summary>
        [Input("restoreType")]
        public Input<string>? RestoreType { get; set; }

        /// <summary>
        /// The hashcode of Snapshot.
        /// </summary>
        [Input("snapshotHash")]
        public Input<string>? SnapshotHash { get; set; }

        /// <summary>
        /// The ID of Snapshot.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The type of data source. Valid values: `ECS_FILE`, `NAS`, `OSS`,`OTS_TABLE`,`UDM_ECS`.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// The Restore Job Status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The target name of OSS bucket. **NOTE:** Required while source_type equals `OSS`,
        /// </summary>
        [Input("targetBucket")]
        public Input<string>? TargetBucket { get; set; }

        /// <summary>
        /// The target client ID.
        /// </summary>
        [Input("targetClientId")]
        public Input<string>? TargetClientId { get; set; }

        /// <summary>
        /// The creation time of destination File System. **NOTE:** While source_type equals `NAS`, this parameter must be set. **Note:** The time format of the API adopts the ISO 8601 format, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
        /// </summary>
        [Input("targetCreateTime")]
        public Input<string>? TargetCreateTime { get; set; }

        /// <summary>
        /// The target data source ID.
        /// </summary>
        [Input("targetDataSourceId")]
        public Input<string>? TargetDataSourceId { get; set; }

        /// <summary>
        /// The ID of destination File System. **NOTE:** Required while source_type equals `NAS`
        /// </summary>
        [Input("targetFileSystemId")]
        public Input<string>? TargetFileSystemId { get; set; }

        /// <summary>
        /// The target ID of ECS instance. **NOTE:** Required while source_type equals `ECS_FILE`
        /// </summary>
        [Input("targetInstanceId")]
        public Input<string>? TargetInstanceId { get; set; }

        /// <summary>
        /// The name of the Table store instance to which you want to restore data.**WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Input("targetInstanceName")]
        public Input<string>? TargetInstanceName { get; set; }

        /// <summary>
        /// The target file path of (ECS) instance. **WARNING:** Required while source_type equals `NAS` or `ECS_FILE`, If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("targetPath")]
        public Input<string>? TargetPath { get; set; }

        /// <summary>
        /// The target prefix of the OSS object. **WARNING:** Required while source_type equals `OSS`. If this value filled in incorrectly, the task may not start correctly, so please check the parameters before executing the plan.
        /// </summary>
        [Input("targetPrefix")]
        public Input<string>? TargetPrefix { get; set; }

        /// <summary>
        /// The name of the table that stores the restored data. **WARNING:** Required while source_type equals `OTS_TABLE`.
        /// </summary>
        [Input("targetTableName")]
        public Input<string>? TargetTableName { get; set; }

        /// <summary>
        /// The time when data is restored to the Table store instance. This value is a UNIX timestamp. Unit: seconds. **WARNING:** Required while source_type equals `OTS_TABLE`. **Note:** The time when data is restored to the Tablestore instance. It should be 0 if restores data at the End time of the snapshot.
        /// </summary>
        [Input("targetTime")]
        public Input<string>? TargetTime { get; set; }

        /// <summary>
        /// The full machine backup details.
        /// </summary>
        [Input("udmDetail")]
        public Input<string>? UdmDetail { get; set; }

        /// <summary>
        /// The ID of backup vault.
        /// </summary>
        [Input("vaultId")]
        public Input<string>? VaultId { get; set; }

        public RestoreJobState()
        {
        }
        public static new RestoreJobState Empty => new RestoreJobState();
    }
}
