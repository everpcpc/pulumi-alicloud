// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Outputs
{

    [OutputType]
    public sealed class GetEdgeKubernetesClustersClusterResult
    {
        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
        /// </summary>
        public readonly Outputs.GetEdgeKubernetesClustersClusterConnectionsResult Connections;
        /// <summary>
        /// ID of the node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Node name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of nat gateway used to launch kubernetes cluster.
        /// </summary>
        public readonly string NatGatewayId;
        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// List of cluster worker nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEdgeKubernetesClustersClusterWorkerNodeResult> WorkerNodes;

        [OutputConstructor]
        private GetEdgeKubernetesClustersClusterResult(
            string availabilityZone,

            Outputs.GetEdgeKubernetesClustersClusterConnectionsResult connections,

            string id,

            string name,

            string natGatewayId,

            string securityGroupId,

            string vpcId,

            ImmutableArray<Outputs.GetEdgeKubernetesClustersClusterWorkerNodeResult> workerNodes)
        {
            AvailabilityZone = availabilityZone;
            Connections = connections;
            Id = id;
            Name = name;
            NatGatewayId = natGatewayId;
            SecurityGroupId = securityGroupId;
            VpcId = vpcId;
            WorkerNodes = workerNodes;
        }
    }
}
