// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// This topic describes how to configure PrivateZone access.
    /// PrivateZone is a VPC-based resolution and management service for private domain names.
    /// After you set a PrivateZone access, the Cloud Connect Network (CCN) and Virtual Border Router (VBR) attached to a CEN instance can access the PrivateZone service through CEN.
    /// 
    /// For information about CEN Private Zone and how to use it, see [Manage CEN Private Zone](https://www.alibabacloud.com/help/en/doc-detail/106693.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.83.0+
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a cen Private Zone resource and use it.
    ///         var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new AliCloud.Cen.InstanceArgs
    ///         {
    ///         });
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///         });
    ///         var defaultInstanceAttachment = new AliCloud.Cen.InstanceAttachment("defaultInstanceAttachment", new AliCloud.Cen.InstanceAttachmentArgs
    ///         {
    ///             InstanceId = defaultInstance.Id,
    ///             ChildInstanceId = defaultNetwork.Id,
    ///             ChildInstanceType = "VPC",
    ///             ChildInstanceRegionId = "cn-hangzhou",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 defaultInstance,
    ///                 defaultNetwork,
    ///             },
    ///         });
    ///         var defaultPrivateZone = new AliCloud.Cen.PrivateZone("defaultPrivateZone", new AliCloud.Cen.PrivateZoneArgs
    ///         {
    ///             AccessRegionId = "cn-hangzhou",
    ///             CenId = defaultInstance.Id,
    ///             HostRegionId = "cn-hangzhou",
    ///             HostVpcId = defaultNetwork.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 defaultInstanceAttachment,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class PrivateZone : Pulumi.CustomResource
    {
        /// <summary>
        /// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
        /// </summary>
        [Output("accessRegionId")]
        public Output<string> AccessRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
        /// </summary>
        [Output("hostRegionId")]
        public Output<string> HostRegionId { get; private set; } = null!;

        /// <summary>
        /// The VPC that belongs to the service region.
        /// </summary>
        [Output("hostVpcId")]
        public Output<string> HostVpcId { get; private set; } = null!;

        /// <summary>
        /// The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateZone(string name, PrivateZoneArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/privateZone:PrivateZone", name, args ?? new PrivateZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateZone(string name, Input<string> id, PrivateZoneState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/privateZone:PrivateZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateZone Get(string name, Input<string> id, PrivateZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateZone(name, id, state, options);
        }
    }

    public sealed class PrivateZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
        /// </summary>
        [Input("accessRegionId", required: true)]
        public Input<string> AccessRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
        /// </summary>
        [Input("hostRegionId", required: true)]
        public Input<string> HostRegionId { get; set; } = null!;

        /// <summary>
        /// The VPC that belongs to the service region.
        /// </summary>
        [Input("hostVpcId", required: true)]
        public Input<string> HostVpcId { get; set; } = null!;

        public PrivateZoneArgs()
        {
        }
    }

    public sealed class PrivateZoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
        /// </summary>
        [Input("accessRegionId")]
        public Input<string>? AccessRegionId { get; set; }

        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.
        /// </summary>
        [Input("hostRegionId")]
        public Input<string>? HostRegionId { get; set; }

        /// <summary>
        /// The VPC that belongs to the service region.
        /// </summary>
        [Input("hostVpcId")]
        public Input<string>? HostVpcId { get; set; }

        /// <summary>
        /// The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateZoneState()
        {
        }
    }
}
