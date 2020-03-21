// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides availability zones for FunctionCompute that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.74.0+.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/fc_zones.html.markdown.
        /// </summary>
        public static Task<GetZonesResult> GetZones(GetZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:fc/getZones:getZones", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetZonesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetZonesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetZonesResult
    {
        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of availability zones. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZonesZonesResult> Zones;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetZonesResult(
            ImmutableArray<string> ids,
            string? outputFile,
            ImmutableArray<Outputs.GetZonesZonesResult> zones,
            string id)
        {
            Ids = ids;
            OutputFile = outputFile;
            Zones = zones;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetZonesZonesResult
    {
        /// <summary>
        /// ID of the zone.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetZonesZonesResult(string id)
        {
            Id = id;
        }
    }
    }
}
