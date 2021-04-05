// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga.Inputs
{

    public sealed class EndpointGroupEndpointConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether client IP addresses are reserved. Valid values: `true`: Client IP addresses are reserved, `false`: Client IP addresses are not reserved. Default value is `false`.
        /// </summary>
        [Input("enableClientipPreservation")]
        public Input<bool>? EnableClientipPreservation { get; set; }

        /// <summary>
        /// The IP address or domain name of Endpoint N in the endpoint group.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// The type of Endpoint N in the endpoint group. Valid values: `Domain`: a custom domain name, `Ip`: a custom IP address, `PublicIp`: an Alibaba Cloud public IP address, `ECS`: an Alibaba Cloud Elastic Compute Service (ECS) instance, `SLB`: an Alibaba Cloud Server Load Balancer (SLB) instance.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The weight of Endpoint N in the endpoint group. Valid value is 0 to 255.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public EndpointGroupEndpointConfigurationGetArgs()
        {
        }
    }
}
