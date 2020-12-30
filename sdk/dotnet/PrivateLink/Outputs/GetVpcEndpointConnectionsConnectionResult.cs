// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink.Outputs
{

    [OutputType]
    public sealed class GetVpcEndpointConnectionsConnectionResult
    {
        /// <summary>
        /// The Bandwidth.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The ID of the Vpc Endpoint.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The ID of the Vpc Endpoint Connection.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The status of Vpc Endpoint Connection.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetVpcEndpointConnectionsConnectionResult(
            int bandwidth,

            string endpointId,

            string id,

            string status)
        {
            Bandwidth = bandwidth;
            EndpointId = endpointId;
            Id = id;
            Status = status;
        }
    }
}
