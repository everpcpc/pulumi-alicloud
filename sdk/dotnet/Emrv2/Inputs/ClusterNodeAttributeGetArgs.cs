// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Emrv2.Inputs
{

    public sealed class ClusterNodeAttributeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the key pair.
        /// </summary>
        [Input("keyPairName", required: true)]
        public Input<string> KeyPairName { get; set; } = null!;

        /// <summary>
        /// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
        /// </summary>
        [Input("ramRole", required: true)]
        public Input<string> RamRole { get; set; } = null!;

        /// <summary>
        /// Security Group ID for Cluster.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// Used to retrieve instances belong to specified VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// Zone ID, e.g. cn-hangzhou-i
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public ClusterNodeAttributeGetArgs()
        {
        }
        public static new ClusterNodeAttributeGetArgs Empty => new ClusterNodeAttributeGetArgs();
    }
}
