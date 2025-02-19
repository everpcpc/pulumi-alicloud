// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms.Outputs
{

    [OutputType]
    public sealed class GetPrometheisPrometheiResult
    {
        /// <summary>
        /// The ID of the cluster.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The type of the cluster.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// The ID of the Grafana workspace.
        /// </summary>
        public readonly string GrafanaInstanceId;
        /// <summary>
        /// The ID of the Prometheus.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// The ID of the security group.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The child instance json string of the globalView instance.
        /// </summary>
        public readonly string SubClustersJson;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The ID of the VSwitch.
        /// </summary>
        public readonly string VswitchId;

        [OutputConstructor]
        private GetPrometheisPrometheiResult(
            string clusterId,

            string clusterName,

            string clusterType,

            string grafanaInstanceId,

            string id,

            string resourceGroupId,

            string securityGroupId,

            string subClustersJson,

            ImmutableDictionary<string, object> tags,

            string vpcId,

            string vswitchId)
        {
            ClusterId = clusterId;
            ClusterName = clusterName;
            ClusterType = clusterType;
            GrafanaInstanceId = grafanaInstanceId;
            Id = id;
            ResourceGroupId = resourceGroupId;
            SecurityGroupId = securityGroupId;
            SubClustersJson = subClustersJson;
            Tags = tags;
            VpcId = vpcId;
            VswitchId = vswitchId;
        }
    }
}
