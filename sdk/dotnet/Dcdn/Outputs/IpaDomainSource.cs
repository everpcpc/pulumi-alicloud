// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn.Outputs
{

    [OutputType]
    public sealed class IpaDomainSource
    {
        /// <summary>
        /// The address of the origin server. You can specify an IP address or a domain name.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// The custom port number. Valid values: `0` to `65535`.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The priority of the origin server. Valid values: `20` and `30`. Default value: `20`. A value of 20 specifies that the origin is a primary origin. A value of 30 specifies that the origin is a secondary origin.
        /// </summary>
        public readonly string Priority;
        /// <summary>
        /// The type of the origin server. Valid values: `ipaddr`, `domain`, `oss`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The weight of the origin server. You must specify a value that is less than `100`. Default value: `10`.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private IpaDomainSource(
            string content,

            int port,

            string priority,

            string type,

            int weight)
        {
            Content = content;
            Port = port;
            Priority = priority;
            Type = type;
            Weight = weight;
        }
    }
}