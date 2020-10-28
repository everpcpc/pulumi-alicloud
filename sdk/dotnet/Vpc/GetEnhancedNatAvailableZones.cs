// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static class GetEnhancedNatAvailableZones
    {
        public static Task<GetEnhancedNatAvailableZonesResult> InvokeAsync(GetEnhancedNatAvailableZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnhancedNatAvailableZonesResult>("alicloud:vpc/getEnhancedNatAvailableZones:getEnhancedNatAvailableZones", args ?? new GetEnhancedNatAvailableZonesArgs(), options.WithVersion());
    }


    public sealed class GetEnhancedNatAvailableZonesArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetEnhancedNatAvailableZonesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnhancedNatAvailableZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetEnhancedNatAvailableZonesZoneResult> Zones;

        [OutputConstructor]
        private GetEnhancedNatAvailableZonesResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetEnhancedNatAvailableZonesZoneResult> zones)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Zones = zones;
        }
    }
}
