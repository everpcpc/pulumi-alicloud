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
    /// Add a group of backend servers (ECS or ENI instance) to the Server Load Balancer or remove them from it.
    /// 
    /// &gt; **NOTE:** Available in 1.53.0+
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var name = config.Get("name") ?? "slbbackendservertest";
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableDiskCategory = "cloud_efficiency",
    ///             AvailableResourceCreation = "VSwitch",
    ///         }));
    ///         var defaultInstanceTypes = defaultZones.Apply(defaultZones =&gt; Output.Create(AliCloud.Ecs.GetInstanceTypes.InvokeAsync(new AliCloud.Ecs.GetInstanceTypesArgs
    ///         {
    ///             AvailabilityZone = defaultZones.Zones[0].Id,
    ///             CpuCoreCount = 1,
    ///             MemorySize = 2,
    ///         })));
    ///         var defaultImages = Output.Create(AliCloud.Ecs.GetImages.InvokeAsync(new AliCloud.Ecs.GetImagesArgs
    ///         {
    ///             MostRecent = true,
    ///             NameRegex = "^ubuntu_18.*64",
    ///             Owners = "system",
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             AvailabilityZone = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///             CidrBlock = "172.16.0.0/16",
    ///             VpcId = defaultNetwork.Id,
    ///         });
    ///         var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new AliCloud.Ecs.SecurityGroupArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///         });
    ///         var defaultInstance = new List&lt;AliCloud.Ecs.Instance&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             defaultInstance.Add(new AliCloud.Ecs.Instance($"defaultInstance-{range.Value}", new AliCloud.Ecs.InstanceArgs
    ///             {
    ///                 AvailabilityZone = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///                 ImageId = defaultImages.Apply(defaultImages =&gt; defaultImages.Images[0].Id),
    ///                 InstanceChargeType = "PostPaid",
    ///                 InstanceName = name,
    ///                 InstanceType = defaultInstanceTypes.Apply(defaultInstanceTypes =&gt; defaultInstanceTypes.InstanceTypes[0].Id),
    ///                 InternetChargeType = "PayByTraffic",
    ///                 InternetMaxBandwidthOut = 10,
    ///                 SecurityGroups = 
    ///                 {
    ///                     defaultSecurityGroup,
    ///                 }.Select(__item =&gt; __item.Id).ToList(),
    ///                 SystemDiskCategory = "cloud_efficiency",
    ///                 VswitchId = defaultSwitch.Id,
    ///             }));
    ///         }
    ///         var defaultLoadBalancer = new AliCloud.Slb.LoadBalancer("defaultLoadBalancer", new AliCloud.Slb.LoadBalancerArgs
    ///         {
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///         var defaultBackendServer = new AliCloud.Slb.BackendServer("defaultBackendServer", new AliCloud.Slb.BackendServerArgs
    ///         {
    ///             BackendServers = 
    ///             {
    ///                 new AliCloud.Slb.Inputs.BackendServerBackendServerArgs
    ///                 {
    ///                     ServerId = defaultInstance[0].Id,
    ///                     Weight = 100,
    ///                 },
    ///                 new AliCloud.Slb.Inputs.BackendServerBackendServerArgs
    ///                 {
    ///                     ServerId = defaultInstance[1].Id,
    ///                     Weight = 100,
    ///                 },
    ///             },
    ///             LoadBalancerId = defaultLoadBalancer.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Block servers
    /// 
    /// The servers mapping supports the following:
    /// 
    /// * `server_id` - (Required) A list backend server ID (ECS instance ID).
    /// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. 
    /// * `type` - (Optional) Type of the backend server. Valid value `ecs`, `eni`. Default to `ecs`.
    /// * `server_ip` - (Optional, Available in 1.93.0+) ServerIp of the backend server. This parameter can be specified when the type is `eni`. `ecs` type currently does not support adding `server_ip` parameter.
    /// </summary>
    public partial class BackendServer : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        [Output("backendServers")]
        public Output<ImmutableArray<Outputs.BackendServerBackendServer>> BackendServers { get; private set; } = null!;

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;


        /// <summary>
        /// Create a BackendServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendServer(string name, BackendServerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/backendServer:BackendServer", name, args ?? new BackendServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendServer(string name, Input<string> id, BackendServerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/backendServer:BackendServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendServer Get(string name, Input<string> id, BackendServerState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendServer(name, id, state, options);
        }
    }

    public sealed class BackendServerArgs : Pulumi.ResourceArgs
    {
        [Input("backendServers")]
        private InputList<Inputs.BackendServerBackendServerArgs>? _backendServers;

        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.BackendServerBackendServerArgs> BackendServers
        {
            get => _backendServers ?? (_backendServers = new InputList<Inputs.BackendServerBackendServerArgs>());
            set => _backendServers = value;
        }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        public BackendServerArgs()
        {
        }
    }

    public sealed class BackendServerState : Pulumi.ResourceArgs
    {
        [Input("backendServers")]
        private InputList<Inputs.BackendServerBackendServerGetArgs>? _backendServers;

        /// <summary>
        /// A list of instances to added backend server in the SLB. It contains three sub-fields as `Block server` follows.
        /// </summary>
        public InputList<Inputs.BackendServerBackendServerGetArgs> BackendServers
        {
            get => _backendServers ?? (_backendServers = new InputList<Inputs.BackendServerBackendServerGetArgs>());
            set => _backendServers = value;
        }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        public BackendServerState()
        {
        }
    }
}
