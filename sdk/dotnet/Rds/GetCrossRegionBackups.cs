// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    public static class GetCrossRegionBackups
    {
        /// <summary>
        /// This data source provides the Rds Parameter Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.196.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var backups = AliCloud.Rds.GetCrossRegionBackups.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         StartTime = "2022-12-01T00:00:00Z",
        ///         EndTime = "2022-12-16T00:00:00Z",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstRdsCrossRegionBackups"] = backups.Apply(getCrossRegionBackupsResult =&gt; getCrossRegionBackupsResult.Backups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCrossRegionBackupsResult> InvokeAsync(GetCrossRegionBackupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCrossRegionBackupsResult>("alicloud:rds/getCrossRegionBackups:getCrossRegionBackups", args ?? new GetCrossRegionBackupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Rds Parameter Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.196.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var backups = AliCloud.Rds.GetCrossRegionBackups.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///         StartTime = "2022-12-01T00:00:00Z",
        ///         EndTime = "2022-12-16T00:00:00Z",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstRdsCrossRegionBackups"] = backups.Apply(getCrossRegionBackupsResult =&gt; getCrossRegionBackupsResult.Backups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCrossRegionBackupsResult> Invoke(GetCrossRegionBackupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCrossRegionBackupsResult>("alicloud:rds/getCrossRegionBackups:getCrossRegionBackups", args ?? new GetCrossRegionBackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCrossRegionBackupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cross-region data backup file.
        /// </summary>
        [Input("backupId")]
        public string? BackupId { get; set; }

        /// <summary>
        /// The ID of the cross-region data backup file.
        /// </summary>
        [Input("crossBackupId")]
        public string? CrossBackupId { get; set; }

        /// <summary>
        /// The ID of the destination region where the cross-region data backup file of the instance is stored.
        /// </summary>
        [Input("crossBackupRegion")]
        public string? CrossBackupRegion { get; set; }

        /// <summary>
        /// The db instance id.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Cross Region Backup IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("startTime")]
        public string? StartTime { get; set; }

        public GetCrossRegionBackupsArgs()
        {
        }
        public static new GetCrossRegionBackupsArgs Empty => new GetCrossRegionBackupsArgs();
    }

    public sealed class GetCrossRegionBackupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the cross-region data backup file.
        /// </summary>
        [Input("backupId")]
        public Input<string>? BackupId { get; set; }

        /// <summary>
        /// The ID of the cross-region data backup file.
        /// </summary>
        [Input("crossBackupId")]
        public Input<string>? CrossBackupId { get; set; }

        /// <summary>
        /// The ID of the destination region where the cross-region data backup file of the instance is stored.
        /// </summary>
        [Input("crossBackupRegion")]
        public Input<string>? CrossBackupRegion { get; set; }

        /// <summary>
        /// The db instance id.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Cross Region Backup IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public GetCrossRegionBackupsInvokeArgs()
        {
        }
        public static new GetCrossRegionBackupsInvokeArgs Empty => new GetCrossRegionBackupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCrossRegionBackupsResult
    {
        public readonly string? BackupId;
        public readonly ImmutableArray<Outputs.GetCrossRegionBackupsBackupResult> Backups;
        public readonly string? CrossBackupId;
        public readonly string? CrossBackupRegion;
        public readonly string DbInstanceId;
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly string? StartTime;

        [OutputConstructor]
        private GetCrossRegionBackupsResult(
            string? backupId,

            ImmutableArray<Outputs.GetCrossRegionBackupsBackupResult> backups,

            string? crossBackupId,

            string? crossBackupRegion,

            string dbInstanceId,

            string? endTime,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? resourceGroupId,

            string? startTime)
        {
            BackupId = backupId;
            Backups = backups;
            CrossBackupId = crossBackupId;
            CrossBackupRegion = crossBackupRegion;
            DbInstanceId = dbInstanceId;
            EndTime = endTime;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            StartTime = startTime;
        }
    }
}