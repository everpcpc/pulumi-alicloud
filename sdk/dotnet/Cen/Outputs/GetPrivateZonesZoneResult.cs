// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen.Outputs
{

    [OutputType]
    public sealed class GetPrivateZonesZoneResult
    {
        /// <summary>
        /// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
        /// </summary>
        public readonly string AccessRegionId;
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        public readonly string CenId;
        /// <summary>
        /// The service region. The service region is the target region of the PrivateZone service accessed through CEN.
        /// </summary>
        public readonly string HostRegionId;
        /// <summary>
        /// The VPC that belongs to the service region.
        /// </summary>
        public readonly string HostVpcId;
        /// <summary>
        /// The DNS IP addresses of the PrivateZone service.
        /// </summary>
        public readonly string PrivateZoneDnsServers;
        /// <summary>
        /// The status of the PrivateZone service, including `Creating`, `Active` and `Deleting`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetPrivateZonesZoneResult(
            string accessRegionId,

            string cenId,

            string hostRegionId,

            string hostVpcId,

            string privateZoneDnsServers,

            string status)
        {
            AccessRegionId = accessRegionId;
            CenId = cenId;
            HostRegionId = hostRegionId;
            HostVpcId = hostVpcId;
            PrivateZoneDnsServers = privateZoneDnsServers;
            Status = status;
        }
    }
}
