// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides CEN Bandwidth Limits available to the user.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cen_bandwidth_limits.html.markdown.
        /// </summary>
        public static Task<GetBandwidthLimitsResult> GetBandwidthLimits(GetBandwidthLimitsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthLimitsResult>("alicloud:cen/getBandwidthLimits:getBandwidthLimits", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetBandwidthLimitsArgs : Pulumi.InvokeArgs
    {
        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// A list of CEN instances IDs.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetBandwidthLimitsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetBandwidthLimitsResult
    {
        public readonly ImmutableArray<string> InstanceIds;
        /// <summary>
        /// A list of CEN Bandwidth Limits. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBandwidthLimitsLimitsResult> Limits;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBandwidthLimitsResult(
            ImmutableArray<string> instanceIds,
            ImmutableArray<Outputs.GetBandwidthLimitsLimitsResult> limits,
            string? outputFile,
            string id)
        {
            InstanceIds = instanceIds;
            Limits = limits;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetBandwidthLimitsLimitsResult
    {
        /// <summary>
        /// The bandwidth limit configured for the interconnected regions communication.
        /// </summary>
        public readonly int BandwidthLimit;
        /// <summary>
        /// ID of the CEN instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// ID of local region.
        /// </summary>
        public readonly string LocalRegionId;
        /// <summary>
        /// ID of opposite region.
        /// </summary>
        public readonly string OppositeRegionId;
        /// <summary>
        /// Status of the CEN Bandwidth Limit, including "Active" and "Modifying".
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetBandwidthLimitsLimitsResult(
            int bandwidthLimit,
            string instanceId,
            string localRegionId,
            string oppositeRegionId,
            string status)
        {
            BandwidthLimit = bandwidthLimit;
            InstanceId = instanceId;
            LocalRegionId = localRegionId;
            OppositeRegionId = oppositeRegionId;
            Status = status;
        }
    }
    }
}
