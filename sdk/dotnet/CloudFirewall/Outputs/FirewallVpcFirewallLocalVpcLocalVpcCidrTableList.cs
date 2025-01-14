// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudFirewall.Outputs
{

    [OutputType]
    public sealed class FirewallVpcFirewallLocalVpcLocalVpcCidrTableList
    {
        /// <summary>
        /// The list of route entries of the local VPC.See the following `Block LocalRouteEntryList`.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> LocalRouteEntryLists;
        /// <summary>
        /// The ID of the route table of the local VPC.
        /// </summary>
        public readonly string LocalRouteTableId;

        [OutputConstructor]
        private FirewallVpcFirewallLocalVpcLocalVpcCidrTableList(
            ImmutableArray<Outputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListLocalRouteEntryList> localRouteEntryLists,

            string localRouteTableId)
        {
            LocalRouteEntryLists = localRouteEntryLists;
            LocalRouteTableId = localRouteTableId;
        }
    }
}
