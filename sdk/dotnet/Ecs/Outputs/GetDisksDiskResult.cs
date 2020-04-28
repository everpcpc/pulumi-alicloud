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
    public sealed class GetDisksDiskResult
    {
        /// <summary>
        /// Disk attachment time.
        /// </summary>
        public readonly string AttachedTime;
        /// <summary>
        /// Availability zone of the disk.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk).
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Disk creation time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Disk description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Disk detachment time.
        /// </summary>
        public readonly string DetachedTime;
        /// <summary>
        /// Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
        /// </summary>
        public readonly string Encrypted;
        /// <summary>
        /// Disk expiration time.
        /// </summary>
        public readonly string ExpirationTime;
        /// <summary>
        /// ID of the disk.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the image from which the disk is created. It is null unless the disk is created using an image.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Filter the results by the specified ECS instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Disk name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Region ID the disk belongs to.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// Disk size in GiB.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// Current status. Possible values: `In_use`, `Available`, `Attaching`, `Detaching`, `Creating` and `ReIniting`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A map of tags assigned to the disks. It must be in the format:
        /// ```
        /// data "alicloud.ecs.getDisks" "disks_ds" {
        /// tags = {
        /// tagKey1 = "tagValue1",
        /// tagKey2 = "tagValue2"
        /// }
        /// }
        /// ```
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// Disk type. Possible values: `system` and `data`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDisksDiskResult(
            string attachedTime,

            string availabilityZone,

            string category,

            string creationTime,

            string description,

            string detachedTime,

            string encrypted,

            string expirationTime,

            string id,

            string imageId,

            string instanceId,

            string name,

            string regionId,

            string resourceGroupId,

            int size,

            string snapshotId,

            string status,

            ImmutableDictionary<string, object>? tags,

            string type)
        {
            AttachedTime = attachedTime;
            AvailabilityZone = availabilityZone;
            Category = category;
            CreationTime = creationTime;
            Description = description;
            DetachedTime = detachedTime;
            Encrypted = encrypted;
            ExpirationTime = expirationTime;
            Id = id;
            ImageId = imageId;
            InstanceId = instanceId;
            Name = name;
            RegionId = regionId;
            ResourceGroupId = resourceGroupId;
            Size = size;
            SnapshotId = snapshotId;
            Status = status;
            Tags = tags;
            Type = type;
        }
    }
}