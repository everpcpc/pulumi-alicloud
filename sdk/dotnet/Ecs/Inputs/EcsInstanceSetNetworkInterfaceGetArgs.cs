// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Inputs
{

    public sealed class EcsInstanceSetNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of ENI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of ENI.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// The primary private IP address of ENI.
        /// </summary>
        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        /// <summary>
        /// -(Required, ForceNew) The ID of the security group to which to assign secondary ENI.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the vSwitch to which to connect ENI.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public EcsInstanceSetNetworkInterfaceGetArgs()
        {
        }
    }
}
