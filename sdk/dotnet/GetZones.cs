// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides availability zones that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// 
        /// &gt; **NOTE:** If one zone is sold out, it will not be exported.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/zones.html.markdown.
        /// </summary>
        public static Task<GetZonesResult> GetZones(GetZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZonesResult>("alicloud:index/getZones:getZones", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetZonesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter the results by a specific disk category. Can be either `cloud`, `cloud_efficiency`, `cloud_ssd`, `ephemeral_ssd`.
        /// </summary>
        [Input("availableDiskCategory")]
        public string? AvailableDiskCategory { get; set; }

        /// <summary>
        /// Filter the results by a specific instance type.
        /// </summary>
        [Input("availableInstanceType")]
        public string? AvailableInstanceType { get; set; }

        /// <summary>
        /// Filter the results by a specific resource type.
        /// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
        /// </summary>
        [Input("availableResourceCreation")]
        public string? AvailableResourceCreation { get; set; }

        /// <summary>
        /// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
        /// &gt; **NOTE:** The disk category `cloud` has been outdated and can only be used by non-I/O Optimized ECS instances. Many availability zones don't support it. It is recommended to use `cloud_efficiency` or `cloud_ssd`.
        /// </summary>
        [Input("availableSlbAddressIpVersion")]
        public string? AvailableSlbAddressIpVersion { get; set; }

        /// <summary>
        /// Filter the results by a slb instance address type. Can be either `Vpc`, `classic_internet` or `classic_intranet`
        /// </summary>
        [Input("availableSlbAddressType")]
        public string? AvailableSlbAddressType { get; set; }

        /// <summary>
        /// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        /// <summary>
        /// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
        /// </summary>
        [Input("multi")]
        public bool? Multi { get; set; }

        /// <summary>
        /// Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
        /// </summary>
        [Input("networkType")]
        public string? NetworkType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
        /// </summary>
        [Input("spotStrategy")]
        public string? SpotStrategy { get; set; }

        public GetZonesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetZonesResult
    {
        public readonly string? AvailableDiskCategory;
        public readonly string? AvailableInstanceType;
        /// <summary>
        /// Type of resources that can be created.
        /// </summary>
        public readonly string? AvailableResourceCreation;
        public readonly string? AvailableSlbAddressIpVersion;
        public readonly string? AvailableSlbAddressType;
        public readonly bool? EnableDetails;
        /// <summary>
        /// A list of zone IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceChargeType;
        public readonly bool? Multi;
        public readonly string? NetworkType;
        public readonly string? OutputFile;
        public readonly string? SpotStrategy;
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
            string? availableDiskCategory,
            string? availableInstanceType,
            string? availableResourceCreation,
            string? availableSlbAddressIpVersion,
            string? availableSlbAddressType,
            bool? enableDetails,
            ImmutableArray<string> ids,
            string? instanceChargeType,
            bool? multi,
            string? networkType,
            string? outputFile,
            string? spotStrategy,
            ImmutableArray<Outputs.GetZonesZonesResult> zones,
            string id)
        {
            AvailableDiskCategory = availableDiskCategory;
            AvailableInstanceType = availableInstanceType;
            AvailableResourceCreation = availableResourceCreation;
            AvailableSlbAddressIpVersion = availableSlbAddressIpVersion;
            AvailableSlbAddressType = availableSlbAddressType;
            EnableDetails = enableDetails;
            Ids = ids;
            InstanceChargeType = instanceChargeType;
            Multi = multi;
            NetworkType = networkType;
            OutputFile = outputFile;
            SpotStrategy = spotStrategy;
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
        /// Set of supported disk categories.
        /// </summary>
        public readonly ImmutableArray<string> AvailableDiskCategories;
        /// <summary>
        /// Allowed instance types.
        /// </summary>
        public readonly ImmutableArray<string> AvailableInstanceTypes;
        /// <summary>
        /// Filter the results by a specific resource type.
        /// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
        /// </summary>
        public readonly ImmutableArray<string> AvailableResourceCreations;
        /// <summary>
        /// ID of the zone.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the zone in the local language.
        /// </summary>
        public readonly string LocalName;
        /// <summary>
        /// A list of zone ids in which the multi zone.
        /// </summary>
        public readonly ImmutableArray<string> MultiZoneIds;
        /// <summary>
        /// A list of slb slave zone ids in which the slb master zone.
        /// </summary>
        public readonly ImmutableArray<string> SlbSlaveZoneIds;

        [OutputConstructor]
        private GetZonesZonesResult(
            ImmutableArray<string> availableDiskCategories,
            ImmutableArray<string> availableInstanceTypes,
            ImmutableArray<string> availableResourceCreations,
            string id,
            string localName,
            ImmutableArray<string> multiZoneIds,
            ImmutableArray<string> slbSlaveZoneIds)
        {
            AvailableDiskCategories = availableDiskCategories;
            AvailableInstanceTypes = availableInstanceTypes;
            AvailableResourceCreations = availableResourceCreations;
            Id = id;
            LocalName = localName;
            MultiZoneIds = multiZoneIds;
            SlbSlaveZoneIds = slbSlaveZoneIds;
        }
    }
    }
}
