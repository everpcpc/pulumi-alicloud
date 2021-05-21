// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Outputs
{

    [OutputType]
    public sealed class GetEcsSnapshotsSnapshotResult
    {
        /// <summary>
        /// The category of the snapshot.
        /// </summary>
        public readonly string Category;
        public readonly string CreationTime;
        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The source disk id.
        /// </summary>
        public readonly string DiskId;
        /// <summary>
        /// Whether the snapshot is encrypted.
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The ID of the Snapshot.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether snapshot speed availability is enabled.
        /// </summary>
        public readonly bool InstantAccess;
        /// <summary>
        /// Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
        /// </summary>
        public readonly int InstantAccessRetentionDays;
        public readonly string Name;
        /// <summary>
        /// The product number inherited from the mirror market.
        /// </summary>
        public readonly string ProductCode;
        /// <summary>
        /// Snapshot creation progress, in percentage.
        /// </summary>
        public readonly string Progress;
        /// <summary>
        /// Remaining completion time for the snapshot being created.
        /// </summary>
        public readonly int RemainTime;
        /// <summary>
        /// The resource group id.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// Automatic snapshot retention days.
        /// </summary>
        public readonly int RetentionDays;
        /// <summary>
        /// The snapshot id.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// Snapshot Display Name.
        /// </summary>
        public readonly string SnapshotName;
        /// <summary>
        /// The serial number of the snapshot.
        /// </summary>
        public readonly string SnapshotSn;
        /// <summary>
        /// Snapshot creation type.
        /// </summary>
        public readonly string SnapshotType;
        public readonly string SourceDiskId;
        /// <summary>
        /// Source disk capacity.
        /// </summary>
        public readonly string SourceDiskSize;
        /// <summary>
        /// Source disk attributes.
        /// </summary>
        public readonly string SourceDiskType;
        /// <summary>
        /// Original disk type.
        /// </summary>
        public readonly string SourceStorageType;
        /// <summary>
        /// The status of the snapshot.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tags.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string Type;
        /// <summary>
        /// A resource type that has a reference relationship.
        /// </summary>
        public readonly string Usage;

        [OutputConstructor]
        private GetEcsSnapshotsSnapshotResult(
            string category,

            string creationTime,

            string description,

            string diskId,

            bool encrypted,

            string id,

            bool instantAccess,

            int instantAccessRetentionDays,

            string name,

            string productCode,

            string progress,

            int remainTime,

            string resourceGroupId,

            int retentionDays,

            string snapshotId,

            string snapshotName,

            string snapshotSn,

            string snapshotType,

            string sourceDiskId,

            string sourceDiskSize,

            string sourceDiskType,

            string sourceStorageType,

            string status,

            ImmutableDictionary<string, object> tags,

            string type,

            string usage)
        {
            Category = category;
            CreationTime = creationTime;
            Description = description;
            DiskId = diskId;
            Encrypted = encrypted;
            Id = id;
            InstantAccess = instantAccess;
            InstantAccessRetentionDays = instantAccessRetentionDays;
            Name = name;
            ProductCode = productCode;
            Progress = progress;
            RemainTime = remainTime;
            ResourceGroupId = resourceGroupId;
            RetentionDays = retentionDays;
            SnapshotId = snapshotId;
            SnapshotName = snapshotName;
            SnapshotSn = snapshotSn;
            SnapshotType = snapshotType;
            SourceDiskId = sourceDiskId;
            SourceDiskSize = sourceDiskSize;
            SourceDiskType = sourceDiskType;
            SourceStorageType = sourceStorageType;
            Status = status;
            Tags = tags;
            Type = type;
            Usage = usage;
        }
    }
}