// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the master slave server groups related to a server load balancer.
        /// 
        /// &gt; **NOTE:** Available in 1.54.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_master_slave_server_groups.html.markdown.
        /// </summary>
        public static Task<GetMasterSlaveServerGroupsResult> GetMasterSlaveServerGroups(GetMasterSlaveServerGroupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMasterSlaveServerGroupsResult>("alicloud:slb/getMasterSlaveServerGroups:getMasterSlaveServerGroups", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetMasterSlaveServerGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of master slave server group IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the SLB.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public string LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by master slave server group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetMasterSlaveServerGroupsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetMasterSlaveServerGroupsResult
    {
        /// <summary>
        /// A list of SLB master slave server groups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMasterSlaveServerGroupsGroupsResult> Groups;
        /// <summary>
        /// A list of SLB master slave server groups IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string LoadBalancerId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of SLB master slave server groups names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetMasterSlaveServerGroupsResult(
            ImmutableArray<Outputs.GetMasterSlaveServerGroupsGroupsResult> groups,
            ImmutableArray<string> ids,
            string loadBalancerId,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Groups = groups;
            Ids = ids;
            LoadBalancerId = loadBalancerId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetMasterSlaveServerGroupsGroupsResult
    {
        /// <summary>
        /// master slave server group ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// master slave server group name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ECS instances associated to the group. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<GetMasterSlaveServerGroupsGroupsServersResult> Servers;

        [OutputConstructor]
        private GetMasterSlaveServerGroupsGroupsResult(
            string id,
            string name,
            ImmutableArray<GetMasterSlaveServerGroupsGroupsServersResult> servers)
        {
            Id = id;
            Name = name;
            Servers = servers;
        }
    }

    [OutputType]
    public sealed class GetMasterSlaveServerGroupsGroupsServersResult
    {
        /// <summary>
        /// ID of the attached ECS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The port used by the master slave server group.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The server type of the attached ECS instance.
        /// </summary>
        public readonly string ServerType;
        /// <summary>
        /// Weight associated to the ECS instance.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private GetMasterSlaveServerGroupsGroupsServersResult(
            string instanceId,
            int port,
            string serverType,
            int weight)
        {
            InstanceId = instanceId;
            Port = port;
            ServerType = serverType;
            Weight = weight;
        }
    }
    }
}
