// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    /// <summary>
    /// Provides a Nlb Load Balancer Security Group Attachment resource.
    /// 
    /// For information about Nlb Load Balancer Security Group Attachment and how to use it, see [What is Load Balancer Security Group Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/loadbalancerjoinsecuritygroup).
    /// 
    /// &gt; **NOTE:** Available in v1.198.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultZones = AliCloud.Nlb.GetZones.Invoke();
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "^default-NODELETING$",
    ///     });
    /// 
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var default1 = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var default2 = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///     });
    /// 
    ///     var zoneId1 = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id);
    /// 
    ///     var vswitchId1 = default1.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]);
    /// 
    ///     var zoneId2 = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id);
    /// 
    ///     var vswitchId2 = default2.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]);
    /// 
    ///     var defaultLoadBalancer = new AliCloud.Nlb.LoadBalancer("defaultLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = @var.Name,
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         LoadBalancerType = "Network",
    ///         AddressType = "Internet",
    ///         AddressIpVersion = "Ipv4",
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneMappings = new[]
    ///         {
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = vswitchId1,
    ///                 ZoneId = zoneId1,
    ///             },
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = vswitchId2,
    ///                 ZoneId = zoneId2,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultLoadBalancerSecurityGroupAttachment = new AliCloud.Nlb.LoadBalancerSecurityGroupAttachment("defaultLoadBalancerSecurityGroupAttachment", new()
    ///     {
    ///         SecurityGroupId = defaultSecurityGroup.Id,
    ///         LoadBalancerId = defaultLoadBalancer.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Nlb Load Balancer Security Group Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment example &lt;LoadBalancerId&gt;:&lt;SecurityGroupId&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment")]
    public partial class LoadBalancerSecurityGroupAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the network-based server load balancer instance to be bound to the security group.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// The ID of security groups.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancerSecurityGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancerSecurityGroupAttachment(string name, LoadBalancerSecurityGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment", name, args ?? new LoadBalancerSecurityGroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancerSecurityGroupAttachment(string name, Input<string> id, LoadBalancerSecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nlb/loadBalancerSecurityGroupAttachment:LoadBalancerSecurityGroupAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancerSecurityGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancerSecurityGroupAttachment Get(string name, Input<string> id, LoadBalancerSecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancerSecurityGroupAttachment(name, id, state, options);
        }
    }

    public sealed class LoadBalancerSecurityGroupAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the network-based server load balancer instance to be bound to the security group.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// The ID of security groups.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public LoadBalancerSecurityGroupAttachmentArgs()
        {
        }
        public static new LoadBalancerSecurityGroupAttachmentArgs Empty => new LoadBalancerSecurityGroupAttachmentArgs();
    }

    public sealed class LoadBalancerSecurityGroupAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to PreCheck this request only. Value:-**true**: sends a check request and does not bind a security group to the instance. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly performs the operation.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the network-based server load balancer instance to be bound to the security group.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// The ID of security groups.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public LoadBalancerSecurityGroupAttachmentState()
        {
        }
        public static new LoadBalancerSecurityGroupAttachmentState Empty => new LoadBalancerSecurityGroupAttachmentState();
    }
}
