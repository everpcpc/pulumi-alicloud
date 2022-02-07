// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail.Outputs
{

    [OutputType]
    public sealed class GetInstancesInstanceAllowedListVpcListResult
    {
        /// <summary>
        /// The allowed ip list of the internet_list.
        /// </summary>
        public readonly ImmutableArray<string> AllowedIpLists;
        /// <summary>
        /// The port range of the internet_list.
        /// </summary>
        public readonly string PortRange;

        [OutputConstructor]
        private GetInstancesInstanceAllowedListVpcListResult(
            ImmutableArray<string> allowedIpLists,

            string portRange)
        {
            AllowedIpLists = allowedIpLists;
            PortRange = portRange;
        }
    }
}