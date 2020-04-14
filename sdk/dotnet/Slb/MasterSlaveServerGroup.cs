// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    /// <summary>
    /// A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.
    /// 
    /// &gt; **NOTE:** One ECS instance can be added into multiple master slave server groups.
    /// 
    /// &gt; **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.
    /// 
    /// &gt; **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.
    /// 
    /// &gt; **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.
    /// 
    /// &gt; **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.
    /// 
    /// &gt; **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.
    /// 
    /// &gt; **NOTE:** Available in 1.54.0+
    /// 
    /// 
    /// ## Block servers
    /// 
    /// The servers mapping supports the following:
    /// 
    /// * `server_ids` - (Required) A list backend server ID (ECS instance ID).
    /// * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
    /// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
    /// * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
    /// * `server_type` - (Optional) The server type of the backend server. Valid value Master, Slave.
    /// * `is_backup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1. 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_master_slave_server_group.html.markdown.
    /// </summary>
    public partial class MasterSlaveServerGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// The Load Balancer ID which is used to launch a new master slave server group.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Name of the master slave server group. 
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        /// </summary>
        [Output("servers")]
        public Output<ImmutableArray<Outputs.MasterSlaveServerGroupServers>> Servers { get; private set; } = null!;


        /// <summary>
        /// Create a MasterSlaveServerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MasterSlaveServerGroup(string name, MasterSlaveServerGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MasterSlaveServerGroup(string name, Input<string> id, MasterSlaveServerGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MasterSlaveServerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MasterSlaveServerGroup Get(string name, Input<string> id, MasterSlaveServerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new MasterSlaveServerGroup(name, id, state, options);
        }
    }

    public sealed class MasterSlaveServerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// The Load Balancer ID which is used to launch a new master slave server group.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// Name of the master slave server group. 
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("servers")]
        private InputList<Inputs.MasterSlaveServerGroupServersArgs>? _servers;

        /// <summary>
        /// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.MasterSlaveServerGroupServersArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.MasterSlaveServerGroupServersArgs>());
            set => _servers = value;
        }

        public MasterSlaveServerGroupArgs()
        {
        }
    }

    public sealed class MasterSlaveServerGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// The Load Balancer ID which is used to launch a new master slave server group.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Name of the master slave server group. 
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("servers")]
        private InputList<Inputs.MasterSlaveServerGroupServersGetArgs>? _servers;

        /// <summary>
        /// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.MasterSlaveServerGroupServersGetArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.MasterSlaveServerGroupServersGetArgs>());
            set => _servers = value;
        }

        public MasterSlaveServerGroupState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class MasterSlaveServerGroupServersArgs : Pulumi.ResourceArgs
    {
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public MasterSlaveServerGroupServersArgs()
        {
        }
    }

    public sealed class MasterSlaveServerGroupServersGetArgs : Pulumi.ResourceArgs
    {
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public MasterSlaveServerGroupServersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class MasterSlaveServerGroupServers
    {
        public readonly int Port;
        public readonly string ServerId;
        public readonly string? ServerType;
        public readonly string? Type;
        public readonly int? Weight;

        [OutputConstructor]
        private MasterSlaveServerGroupServers(
            int port,
            string serverId,
            string? serverType,
            string? type,
            int? weight)
        {
            Port = port;
            ServerId = serverId;
            ServerType = serverType;
            Type = type;
            Weight = weight;
        }
    }
    }
}
