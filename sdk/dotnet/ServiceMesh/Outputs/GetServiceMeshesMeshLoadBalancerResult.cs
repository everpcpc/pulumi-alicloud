// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh.Outputs
{

    [OutputType]
    public sealed class GetServiceMeshesMeshLoadBalancerResult
    {
        /// <summary>
        /// The IP address of a public network exposed API Server corresponding to the Load Balance.
        /// </summary>
        public readonly string ApiServerLoadbalancerId;
        /// <summary>
        /// Whether to use the IP address of a public network exposed the API Server.
        /// </summary>
        public readonly bool ApiServerPublicEip;
        /// <summary>
        /// Whether to use the IP address of a public network exposure the Istio Pilot.
        /// </summary>
        public readonly bool PilotPublicEip;
        /// <summary>
        /// The IP address of a public network exposure Istio Pilot corresponds to the Load Balance.
        /// </summary>
        public readonly string PilotPublicLoadbalancerId;

        [OutputConstructor]
        private GetServiceMeshesMeshLoadBalancerResult(
            string apiServerLoadbalancerId,

            bool apiServerPublicEip,

            bool pilotPublicEip,

            string pilotPublicLoadbalancerId)
        {
            ApiServerLoadbalancerId = apiServerLoadbalancerId;
            ApiServerPublicEip = apiServerPublicEip;
            PilotPublicEip = pilotPublicEip;
            PilotPublicLoadbalancerId = pilotPublicLoadbalancerId;
        }
    }
}