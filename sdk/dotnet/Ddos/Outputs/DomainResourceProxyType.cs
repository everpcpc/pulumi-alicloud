// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos.Outputs
{

    [OutputType]
    public sealed class DomainResourceProxyType
    {
        /// <summary>
        /// the port number. This field is required and must be an integer.
        /// </summary>
        public readonly ImmutableArray<int> ProxyPorts;
        /// <summary>
        /// the protocol type. This field is required and must be a string. Valid values: `http`, `https`, `websocket`, and `websockets`.
        /// </summary>
        public readonly string? ProxyType;

        [OutputConstructor]
        private DomainResourceProxyType(
            ImmutableArray<int> proxyPorts,

            string? proxyType)
        {
            ProxyPorts = proxyPorts;
            ProxyType = proxyType;
        }
    }
}
