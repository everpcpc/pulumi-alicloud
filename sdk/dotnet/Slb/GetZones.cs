// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static class GetZones
    {
        /// <summary>
        /// This data source provides availability zones for SLB that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.73.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var zonesIds = AliCloud.Slb.GetZones.Invoke(new()
        ///     {
        ///         AvailableSlbAddressIpVersion = "ipv4",
        ///         AvailableSlbAddressType = "vpc",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZonesResult> InvokeAsync(GetZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:slb/getZones:getZones", args ?? new GetZonesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides availability zones for SLB that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.73.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var zonesIds = AliCloud.Slb.GetZones.Invoke(new()
        ///     {
        ///         AvailableSlbAddressIpVersion = "ipv4",
        ///         AvailableSlbAddressType = "vpc",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZonesResult> Invoke(GetZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZonesResult>("alicloud:slb/getZones:getZones", args ?? new GetZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
        /// </summary>
        [Input("availableSlbAddressIpVersion")]
        public string? AvailableSlbAddressIpVersion { get; set; }

        /// <summary>
        /// Filter the results by a slb instance network type. Valid values:
        /// * vpc: an internal SLB instance that is deployed in a virtual private cloud (VPC).
        /// * classic_internet: a public-facing SLB instance.
        /// * classic_intranet: an internal SLB instance that is deployed in a classic network.
        /// </summary>
        [Input("availableSlbAddressType")]
        public string? AvailableSlbAddressType { get; set; }

        /// <summary>
        /// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// The primary zone.
        /// </summary>
        [Input("masterZoneId")]
        public string? MasterZoneId { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The secondary zone.
        /// </summary>
        [Input("slaveZoneId")]
        public string? SlaveZoneId { get; set; }

        public GetZonesArgs()
        {
        }
        public static new GetZonesArgs Empty => new GetZonesArgs();
    }

    public sealed class GetZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
        /// </summary>
        [Input("availableSlbAddressIpVersion")]
        public Input<string>? AvailableSlbAddressIpVersion { get; set; }

        /// <summary>
        /// Filter the results by a slb instance network type. Valid values:
        /// * vpc: an internal SLB instance that is deployed in a virtual private cloud (VPC).
        /// * classic_internet: a public-facing SLB instance.
        /// * classic_intranet: an internal SLB instance that is deployed in a classic network.
        /// </summary>
        [Input("availableSlbAddressType")]
        public Input<string>? AvailableSlbAddressType { get; set; }

        /// <summary>
        /// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        /// <summary>
        /// The primary zone.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The secondary zone.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        public GetZonesInvokeArgs()
        {
        }
        public static new GetZonesInvokeArgs Empty => new GetZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetZonesResult
    {
        public readonly string? AvailableSlbAddressIpVersion;
        public readonly string? AvailableSlbAddressType;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of primary zone IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// (Available in 1.157.0+) The primary zone.
        /// </summary>
        public readonly string? MasterZoneId;
        public readonly string? OutputFile;
        /// <summary>
        /// (Available in 1.157.0+) The secondary zone.
        /// </summary>
        public readonly string? SlaveZoneId;
        /// <summary>
        /// A list of availability zones. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZonesZoneResult> Zones;

        [OutputConstructor]
        private GetZonesResult(
            string? availableSlbAddressIpVersion,

            string? availableSlbAddressType,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? masterZoneId,

            string? outputFile,

            string? slaveZoneId,

            ImmutableArray<Outputs.GetZonesZoneResult> zones)
        {
            AvailableSlbAddressIpVersion = availableSlbAddressIpVersion;
            AvailableSlbAddressType = availableSlbAddressType;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            MasterZoneId = masterZoneId;
            OutputFile = outputFile;
            SlaveZoneId = slaveZoneId;
            Zones = zones;
        }
    }
}
