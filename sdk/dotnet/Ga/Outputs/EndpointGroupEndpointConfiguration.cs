// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Outputs
{

    [OutputType]
    public sealed class EndpointGroupEndpointConfiguration
    {
        /// <summary>
        /// Indicates whether client IP addresses are reserved. Valid values: `true`: Client IP addresses are reserved, `false`: Client IP addresses are not reserved. Default value is `false`.
        /// </summary>
        public readonly bool? EnableClientipPreservation;
        /// <summary>
        /// The IP address or domain name of Endpoint N in the endpoint group.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The type of Endpoint N in the endpoint group. Valid values: `Domain`: a custom domain name, `Ip`: a custom IP address, `PublicIp`: an Alibaba Cloud public IP address, `ECS`: an Alibaba Cloud Elastic Compute Service (ECS) instance, `SLB`: an Alibaba Cloud Server Load Balancer (SLB) instance.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The weight of Endpoint N in the endpoint group. Valid value is 0 to 255.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private EndpointGroupEndpointConfiguration(
            bool? enableClientipPreservation,

            string endpoint,

            string type,

            int weight)
        {
            EnableClientipPreservation = enableClientipPreservation;
            Endpoint = endpoint;
            Type = type;
            Weight = weight;
        }
    }
}
