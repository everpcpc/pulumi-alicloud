// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ElasticSearch
{
    public static partial class Invokes
    {
        public static Task<GetInstancesResult> GetInstances(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:elasticsearch/getInstances:getInstances", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("version")]
        public string? Version { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly string? DescriptionRegex;
        public readonly ImmutableArray<string> Descriptions;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetInstancesInstancesResult> Instances;
        public readonly string? OutputFile;
        public readonly string? Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstancesResult(
            string? descriptionRegex,
            ImmutableArray<string> descriptions,
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetInstancesInstancesResult> instances,
            string? outputFile,
            string? version,
            string id)
        {
            DescriptionRegex = descriptionRegex;
            Descriptions = descriptions;
            Ids = ids;
            Instances = instances;
            OutputFile = outputFile;
            Version = version;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstancesInstancesResult
    {
        public readonly string CreatedAt;
        public readonly int DataNodeAmount;
        public readonly int DataNodeDiskSize;
        public readonly string DataNodeDiskType;
        public readonly string DataNodeSpec;
        public readonly string Description;
        public readonly string Id;
        public readonly string InstanceChargeType;
        public readonly string Status;
        public readonly string UpdatedAt;
        public readonly string Version;
        public readonly string VswitchId;

        [OutputConstructor]
        private GetInstancesInstancesResult(
            string createdAt,
            int dataNodeAmount,
            int dataNodeDiskSize,
            string dataNodeDiskType,
            string dataNodeSpec,
            string description,
            string id,
            string instanceChargeType,
            string status,
            string updatedAt,
            string version,
            string vswitchId)
        {
            CreatedAt = createdAt;
            DataNodeAmount = dataNodeAmount;
            DataNodeDiskSize = dataNodeDiskSize;
            DataNodeDiskType = dataNodeDiskType;
            DataNodeSpec = dataNodeSpec;
            Description = description;
            Id = id;
            InstanceChargeType = instanceChargeType;
            Status = status;
            UpdatedAt = updatedAt;
            Version = version;
            VswitchId = vswitchId;
        }
    }
    }
}
