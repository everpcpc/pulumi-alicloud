// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud, 
        /// other public images and the ones available on the image market. 
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/images.html.markdown.
        /// </summary>
        [Obsolete("Use GetImages.InvokeAsync() instead")]
        public static Task<GetImagesResult> GetImages(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("alicloud:ecs/getImages:getImages", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetImages
    {
        /// <summary>
        /// This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud, 
        /// other public images and the ones available on the image market. 
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/images.html.markdown.
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("alicloud:ecs/getImages:getImages", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If more than one result are returned, select the most recent one.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// A regex string to filter resulting images by name. 
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
        /// </summary>
        [Input("owners")]
        public string? Owners { get; set; }

        public GetImagesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetImagesResult
    {
        /// <summary>
        /// A list of image IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of images. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImagesResult> Images;
        public readonly bool? MostRecent;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? Owners;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetImagesResult(
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetImagesImagesResult> images,
            bool? mostRecent,
            string? nameRegex,
            string? outputFile,
            string? owners,
            string id)
        {
            Ids = ids;
            Images = images;
            MostRecent = mostRecent;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Owners = owners;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetImagesImagesDiskDeviceMappingsResult
    {
        /// <summary>
        /// Device information of the created disk: such as /dev/xvdb.
        /// </summary>
        public readonly string Device;
        /// <summary>
        /// Size of the created disk.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// Snapshot ID.
        /// </summary>
        public readonly string SnapshotId;

        [OutputConstructor]
        private GetImagesImagesDiskDeviceMappingsResult(
            string device,
            string size,
            string snapshotId)
        {
            Device = device;
            Size = size;
            SnapshotId = snapshotId;
        }
    }

    [OutputType]
    public sealed class GetImagesImagesResult
    {
        /// <summary>
        /// Platform type of the image system: i386 or x86_64.
        /// </summary>
        public readonly string Architecture;
        /// <summary>
        /// Time of creation.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Description of the image.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Description of the system with disks and snapshots under the image.
        /// </summary>
        public readonly ImmutableArray<GetImagesImagesDiskDeviceMappingsResult> DiskDeviceMappings;
        /// <summary>
        /// ID of the image.
        /// </summary>
        public readonly string Id;
        public readonly string ImageId;
        /// <summary>
        /// Alias of the image owner.
        /// </summary>
        public readonly string ImageOwnerAlias;
        /// <summary>
        /// Version of the image.
        /// </summary>
        public readonly string ImageVersion;
        public readonly bool IsCopied;
        public readonly string IsSelfShared;
        /// <summary>
        /// Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.
        /// </summary>
        public readonly bool IsSubscribed;
        public readonly bool IsSupportIoOptimized;
        public readonly string Name;
        /// <summary>
        /// Display Chinese name of the OS.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// Display English name of the OS.
        /// </summary>
        public readonly string OsNameEn;
        public readonly string OsType;
        public readonly string Platform;
        /// <summary>
        /// Product code of the image on the image market.
        /// </summary>
        public readonly string ProductCode;
        /// <summary>
        /// Progress of image creation, presented in percentages.
        /// </summary>
        public readonly string Progress;
        /// <summary>
        /// Size of the created disk.
        /// </summary>
        public readonly int Size;
        public readonly string State;
        /// <summary>
        /// Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
        /// </summary>
        public readonly string Status;
        public readonly ImmutableDictionary<string, object>? Tags;
        public readonly string Usage;

        [OutputConstructor]
        private GetImagesImagesResult(
            string architecture,
            string creationTime,
            string description,
            ImmutableArray<GetImagesImagesDiskDeviceMappingsResult> diskDeviceMappings,
            string id,
            string imageId,
            string imageOwnerAlias,
            string imageVersion,
            bool isCopied,
            string isSelfShared,
            bool isSubscribed,
            bool isSupportIoOptimized,
            string name,
            string osName,
            string osNameEn,
            string osType,
            string platform,
            string productCode,
            string progress,
            int size,
            string state,
            string status,
            ImmutableDictionary<string, object>? tags,
            string usage)
        {
            Architecture = architecture;
            CreationTime = creationTime;
            Description = description;
            DiskDeviceMappings = diskDeviceMappings;
            Id = id;
            ImageId = imageId;
            ImageOwnerAlias = imageOwnerAlias;
            ImageVersion = imageVersion;
            IsCopied = isCopied;
            IsSelfShared = isSelfShared;
            IsSubscribed = isSubscribed;
            IsSupportIoOptimized = isSupportIoOptimized;
            Name = name;
            OsName = osName;
            OsNameEn = osNameEn;
            OsType = osType;
            Platform = platform;
            ProductCode = productCode;
            Progress = progress;
            Size = size;
            State = state;
            Status = status;
            Tags = tags;
            Usage = usage;
        }
    }
    }
}
