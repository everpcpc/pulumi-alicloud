// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.35.0+
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cs_managed_kubernetes_clusters.html.markdown.
        /// </summary>
        [Obsolete("Use GetManagedKubernetesClusters.InvokeAsync() instead")]
        public static Task<GetManagedKubernetesClustersResult> GetManagedKubernetesClusters(GetManagedKubernetesClustersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedKubernetesClustersResult>("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetManagedKubernetesClusters
    {
        /// <summary>
        /// This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in v1.35.0+
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cs_managed_kubernetes_clusters.html.markdown.
        /// </summary>
        public static Task<GetManagedKubernetesClustersResult> InvokeAsync(GetManagedKubernetesClustersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagedKubernetesClustersResult>("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetManagedKubernetesClustersArgs : Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// Cluster IDs to filter.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by cluster name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetManagedKubernetesClustersArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetManagedKubernetesClustersResult
    {
        /// <summary>
        /// A list of matched Kubernetes clusters. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetManagedKubernetesClustersClustersResult> Clusters;
        public readonly bool? EnableDetails;
        /// <summary>
        /// A list of matched Kubernetes clusters' ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of matched Kubernetes clusters' names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetManagedKubernetesClustersResult(
            ImmutableArray<Outputs.GetManagedKubernetesClustersClustersResult> clusters,
            bool? enableDetails,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Clusters = clusters;
            EnableDetails = enableDetails;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetManagedKubernetesClustersClustersConnectionsResult
    {
        /// <summary>
        /// API Server Internet endpoint.
        /// </summary>
        public readonly string ApiServerInternet;
        /// <summary>
        /// API Server Intranet endpoint.
        /// </summary>
        public readonly string ApiServerIntranet;
        /// <summary>
        /// Master node SSH IP address.
        /// </summary>
        public readonly string MasterPublicIp;
        /// <summary>
        /// Service Access Domain.
        /// </summary>
        public readonly string ServiceDomain;

        [OutputConstructor]
        private GetManagedKubernetesClustersClustersConnectionsResult(
            string apiServerInternet,
            string apiServerIntranet,
            string masterPublicIp,
            string serviceDomain)
        {
            ApiServerInternet = apiServerInternet;
            ApiServerIntranet = apiServerIntranet;
            MasterPublicIp = masterPublicIp;
            ServiceDomain = serviceDomain;
        }
    }

    [OutputType]
    public sealed class GetManagedKubernetesClustersClustersResult
    {
        /// <summary>
        /// The ID of availability zone.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// Map of kubernetes cluster connection information. It contains several attributes to `Block Connections`.
        /// </summary>
        public readonly GetManagedKubernetesClustersClustersConnectionsResult Connections;
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
        public readonly ImmutableArray<GetManagedKubernetesClustersClustersWorkerNodesResult> WorkerNodes;

        [OutputConstructor]
        private GetManagedKubernetesClustersClustersResult(
            string availabilityZone,
            GetManagedKubernetesClustersClustersConnectionsResult connections,
            string id,
            string name,
            string natGatewayId,
            string securityGroupId,
            string vpcId,
            ImmutableArray<GetManagedKubernetesClustersClustersWorkerNodesResult> workerNodes)
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

    [OutputType]
    public sealed class GetManagedKubernetesClustersClustersWorkerNodesResult
    {
        /// <summary>
        /// ID of the node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Node name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The private IP address of node.
        /// </summary>
        public readonly string PrivateIp;

        [OutputConstructor]
        private GetManagedKubernetesClustersClustersWorkerNodesResult(
            string id,
            string name,
            string privateIp)
        {
            Id = id;
            Name = name;
            PrivateIp = privateIp;
        }
    }
    }
}
