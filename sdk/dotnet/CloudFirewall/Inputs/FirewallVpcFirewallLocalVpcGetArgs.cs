// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudFirewall.Inputs
{

    public sealed class FirewallVpcFirewallLocalVpcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the instance of the Eni in the local VPC.
        /// </summary>
        [Input("eniId")]
        public Input<string>? EniId { get; set; }

        /// <summary>
        /// The private IP address of the elastic network card in the local VPC.
        /// </summary>
        [Input("eniPrivateIpAddress")]
        public Input<string>? EniPrivateIpAddress { get; set; }

        [Input("localVpcCidrTableLists", required: true)]
        private InputList<Inputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListGetArgs>? _localVpcCidrTableLists;

        /// <summary>
        /// The network segment list of the local VPC.See the following `Block LocalVpcCidrTableList`.
        /// </summary>
        public InputList<Inputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListGetArgs> LocalVpcCidrTableLists
        {
            get => _localVpcCidrTableLists ?? (_localVpcCidrTableLists = new InputList<Inputs.FirewallVpcFirewallLocalVpcLocalVpcCidrTableListGetArgs>());
            set => _localVpcCidrTableLists = value;
        }

        /// <summary>
        /// The region ID of the local VPC.
        /// </summary>
        [Input("regionNo", required: true)]
        public Input<string> RegionNo { get; set; } = null!;

        /// <summary>
        /// The ID of the router interface in the local VPC.
        /// </summary>
        [Input("routerInterfaceId")]
        public Input<string>? RouterInterfaceId { get; set; }

        /// <summary>
        /// The ID of the local VPC instance.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The instance name of the local VPC.
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public FirewallVpcFirewallLocalVpcGetArgs()
        {
        }
        public static new FirewallVpcFirewallLocalVpcGetArgs Empty => new FirewallVpcFirewallLocalVpcGetArgs();
    }
}