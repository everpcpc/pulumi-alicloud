// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the disks of the current Alibaba Cloud user.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/disks.html.markdown.
        /// </summary>
        public static Task<GetDisksResult> GetDisks(GetDisksArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDisksResult>("alicloud:ecs/getDisks:getDisks", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDisksArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk).
        /// </summary>
        [Input("category")]
        public string? Category { get; set; }

        /// <summary>
        /// Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
        /// </summary>
        [Input("encrypted")]
        public string? Encrypted { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of disks IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Filter the results by the specified ECS instance ID.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// A regex string to filter results by disk name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of resource group which the disk belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

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
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// Disk type. Possible values: `system` and `data`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetDisksArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDisksResult
    {
        /// <summary>
        /// Disk category. Possible values: `cloud` (basic cloud disk), `cloud_efficiency` (ultra cloud disk), `ephemeral_ssd` (local SSD cloud disk), `cloud_ssd` (SSD cloud disk), and `cloud_essd` (ESSD cloud disk).
        /// </summary>
        public readonly string? Category;
        /// <summary>
        /// A list of disks. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDisksDisksResult> Disks;
        /// <summary>
        /// Indicate whether the disk is encrypted or not. Possible values: `on` and `off`.
        /// </summary>
        public readonly string? Encrypted;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// ID of the related instance. It is `null` unless the `status` is `In_use`.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The Id of resource group.
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// A map of tags assigned to the disk.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// Disk type. Possible values: `system` and `data`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDisksResult(
            string? category,
            ImmutableArray<Outputs.GetDisksDisksResult> disks,
            string? encrypted,
            ImmutableArray<string> ids,
            string? instanceId,
            string? nameRegex,
            string? outputFile,
            string? resourceGroupId,
            ImmutableDictionary<string, object>? tags,
            string? type,
            string id)
        {
            Category = category;
            Disks = disks;
            Encrypted = encrypted;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            Tags = tags;
            Type = type;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetDisksDisksResult
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
        private GetDisksDisksResult(
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
}
