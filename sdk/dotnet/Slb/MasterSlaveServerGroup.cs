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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var msServerGroupZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableDiskCategory = "cloud_efficiency",
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var msServerGroupInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = msServerGroupZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         EniAmount = 2,
    ///     });
    /// 
    ///     var image = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var config = new Config();
    ///     var slbMasterSlaveServerGroup = config.Get("slbMasterSlaveServerGroup") ?? "forSlbMasterSlaveServerGroup";
    ///     var mainNetwork = new AliCloud.Vpc.Network("mainNetwork", new()
    ///     {
    ///         VpcName = slbMasterSlaveServerGroup,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var mainSwitch = new AliCloud.Vpc.Switch("mainSwitch", new()
    ///     {
    ///         VpcId = mainNetwork.Id,
    ///         CidrBlock = "172.16.0.0/16",
    ///         ZoneId = msServerGroupZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = slbMasterSlaveServerGroup,
    ///     });
    /// 
    ///     var groupSecurityGroup = new AliCloud.Ecs.SecurityGroup("groupSecurityGroup", new()
    ///     {
    ///         VpcId = mainNetwork.Id,
    ///     });
    /// 
    ///     var msServerGroupInstance = new List&lt;AliCloud.Ecs.Instance&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         msServerGroupInstance.Add(new AliCloud.Ecs.Instance($"msServerGroupInstance-{range.Value}", new()
    ///         {
    ///             ImageId = image.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///             InstanceType = msServerGroupInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///             InstanceName = slbMasterSlaveServerGroup,
    ///             SecurityGroups = new[]
    ///             {
    ///                 groupSecurityGroup.Id,
    ///             },
    ///             InternetChargeType = "PayByTraffic",
    ///             InternetMaxBandwidthOut = 10,
    ///             AvailabilityZone = msServerGroupZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///             InstanceChargeType = "PostPaid",
    ///             SystemDiskCategory = "cloud_efficiency",
    ///             VswitchId = mainSwitch.Id,
    ///         }));
    ///     }
    ///     var msServerGroupApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("msServerGroupApplicationLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = slbMasterSlaveServerGroup,
    ///         VswitchId = mainSwitch.Id,
    ///         LoadBalancerSpec = "slb.s2.small",
    ///     });
    /// 
    ///     var msServerGroupEcsNetworkInterface = new AliCloud.Ecs.EcsNetworkInterface("msServerGroupEcsNetworkInterface", new()
    ///     {
    ///         NetworkInterfaceName = slbMasterSlaveServerGroup,
    ///         VswitchId = mainSwitch.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             groupSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var msServerGroupEcsNetworkInterfaceAttachment = new AliCloud.Ecs.EcsNetworkInterfaceAttachment("msServerGroupEcsNetworkInterfaceAttachment", new()
    ///     {
    ///         InstanceId = msServerGroupInstance[0].Id,
    ///         NetworkInterfaceId = msServerGroupEcsNetworkInterface.Id,
    ///     });
    /// 
    ///     var groupMasterSlaveServerGroup = new AliCloud.Slb.MasterSlaveServerGroup("groupMasterSlaveServerGroup", new()
    ///     {
    ///         LoadBalancerId = msServerGroupApplicationLoadBalancer.Id,
    ///         Servers = new[]
    ///         {
    ///             new AliCloud.Slb.Inputs.MasterSlaveServerGroupServerArgs
    ///             {
    ///                 ServerId = msServerGroupInstance[0].Id,
    ///                 Port = 100,
    ///                 Weight = 100,
    ///                 ServerType = "Master",
    ///             },
    ///             new AliCloud.Slb.Inputs.MasterSlaveServerGroupServerArgs
    ///             {
    ///                 ServerId = msServerGroupInstance[1].Id,
    ///                 Port = 100,
    ///                 Weight = 100,
    ///                 ServerType = "Slave",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var tcp = new AliCloud.Slb.Listener("tcp", new()
    ///     {
    ///         LoadBalancerId = msServerGroupApplicationLoadBalancer.Id,
    ///         MasterSlaveServerGroupId = groupMasterSlaveServerGroup.Id,
    ///         FrontendPort = 22,
    ///         Protocol = "tcp",
    ///         Bandwidth = 10,
    ///         HealthCheckType = "tcp",
    ///         PersistenceTimeout = 3600,
    ///         HealthyThreshold = 8,
    ///         UnhealthyThreshold = 8,
    ///         HealthCheckTimeout = 8,
    ///         HealthCheckInterval = 5,
    ///         HealthCheckHttpCode = "http_2xx",
    ///         HealthCheckConnectPort = 20,
    ///         HealthCheckUri = "/console",
    ///         EstablishedTimeout = 600,
    ///     });
    /// 
    /// });
    /// ```
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
    /// ## Import
    /// 
    /// Load balancer master slave server group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup")]
    public partial class MasterSlaveServerGroup : global::Pulumi.CustomResource
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
        public Output<ImmutableArray<Outputs.MasterSlaveServerGroupServer>> Servers { get; private set; } = null!;


        /// <summary>
        /// Create a MasterSlaveServerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MasterSlaveServerGroup(string name, MasterSlaveServerGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, args ?? new MasterSlaveServerGroupArgs(), MakeResourceOptions(options, ""))
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

    public sealed class MasterSlaveServerGroupArgs : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.MasterSlaveServerGroupServerArgs>? _servers;

        /// <summary>
        /// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.MasterSlaveServerGroupServerArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.MasterSlaveServerGroupServerArgs>());
            set => _servers = value;
        }

        public MasterSlaveServerGroupArgs()
        {
        }
        public static new MasterSlaveServerGroupArgs Empty => new MasterSlaveServerGroupArgs();
    }

    public sealed class MasterSlaveServerGroupState : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.MasterSlaveServerGroupServerGetArgs>? _servers;

        /// <summary>
        /// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.MasterSlaveServerGroupServerGetArgs> Servers
        {
            get => _servers ?? (_servers = new InputList<Inputs.MasterSlaveServerGroupServerGetArgs>());
            set => _servers = value;
        }

        public MasterSlaveServerGroupState()
        {
        }
        public static new MasterSlaveServerGroupState Empty => new MasterSlaveServerGroupState();
    }
}
