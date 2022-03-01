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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var name = config.Get("name") ?? "slbattachmenttest";
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableDiskCategory = "cloud_efficiency",
    ///             AvailableResourceCreation = "VSwitch",
    ///         }));
    ///         var defaultInstanceTypes = defaultZones.Apply(defaultZones =&gt; Output.Create(AliCloud.Ecs.GetInstanceTypes.InvokeAsync(new AliCloud.Ecs.GetInstanceTypesArgs
    ///         {
    ///             AvailabilityZone = defaultZones.Zones?[0]?.Id,
    ///             CpuCoreCount = 1,
    ///             MemorySize = 2,
    ///         })));
    ///         var defaultImages = Output.Create(AliCloud.Ecs.GetImages.InvokeAsync(new AliCloud.Ecs.GetImagesArgs
    ///         {
    ///             NameRegex = "^ubuntu_18.*64",
    ///             MostRecent = true,
    ///             Owners = "system",
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///             CidrBlock = "172.16.0.0/16",
    ///             ZoneId = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones?[0]?.Id),
    ///             VswitchName = name,
    ///         });
    ///         var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new AliCloud.Ecs.SecurityGroupArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///         });
    ///         var defaultInstance = new AliCloud.Ecs.Instance("defaultInstance", new AliCloud.Ecs.InstanceArgs
    ///         {
    ///             ImageId = defaultImages.Apply(defaultImages =&gt; defaultImages.Images?[0]?.Id),
    ///             InstanceType = defaultInstanceTypes.Apply(defaultInstanceTypes =&gt; defaultInstanceTypes.InstanceTypes?[0]?.Id),
    ///             InternetChargeType = "PayByTraffic",
    ///             InternetMaxBandwidthOut = 5,
    ///             SystemDiskCategory = "cloud_efficiency",
    ///             SecurityGroups = 
    ///             {
    ///                 defaultSecurityGroup.Id,
    ///             },
    ///             InstanceName = name,
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///         var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", new AliCloud.Slb.ApplicationLoadBalancerArgs
    ///         {
    ///             LoadBalancerName = name,
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///         var defaultAttachment = new AliCloud.Slb.Attachment("defaultAttachment", new AliCloud.Slb.AttachmentArgs
    ///         {
    ///             LoadBalancerId = defaultApplicationLoadBalancer.Id,
    ///             InstanceIds = 
    ///             {
    ///                 defaultInstance.Id,
    ///             },
    ///             Weight = 90,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load balancer attachment can be imported using the id or load balancer id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:slb/attachment:Attachment example lb-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/attachment:Attachment")]
    public partial class Attachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The backend servers of the load balancer.
        /// </summary>
        [Output("backendServers")]
        public Output<string> BackendServers { get; private set; } = null!;

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// A list of instance ids to added backend server in the SLB.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Type of the instances. Valid value ecs, eni. Default to ecs.
        /// </summary>
        [Output("serverType")]
        public Output<string?> ServerType { get; private set; } = null!;

        /// <summary>
        /// Weight of the instances. Valid value range: [0-100]. Default to 100.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attachment(string name, AttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/attachment:Attachment", name, args ?? new AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attachment(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/attachment:Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attachment Get(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Attachment(name, id, state, options);
        }
    }

    public sealed class AttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend servers of the load balancer.
        /// </summary>
        [Input("backendServers")]
        public Input<string>? BackendServers { get; set; }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// A list of instance ids to added backend server in the SLB.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// Type of the instances. Valid value ecs, eni. Default to ecs.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        /// <summary>
        /// Weight of the instances. Valid value range: [0-100]. Default to 100.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public AttachmentArgs()
        {
        }
    }

    public sealed class AttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend servers of the load balancer.
        /// </summary>
        [Input("backendServers")]
        public Input<string>? BackendServers { get; set; }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// A list of instance ids to added backend server in the SLB.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the load balancer.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Type of the instances. Valid value ecs, eni. Default to ecs.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        /// <summary>
        /// Weight of the instances. Valid value range: [0-100]. Default to 100.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public AttachmentState()
        {
        }
    }
}
