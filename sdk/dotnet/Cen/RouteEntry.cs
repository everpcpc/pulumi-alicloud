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
    /// Provides a CEN route entry resource. Cloud Enterprise Network (CEN) supports publishing and withdrawing route entries of attached networks. You can publish a route entry of an attached VPC or VBR to a CEN instance, then other attached networks can learn the route if there is no route conflict. You can withdraw a published route entry when CEN does not need it any more.
    /// 
    /// For information about CEN route entries publishment and how to use it, see [Manage network routes](https://www.alibabacloud.com/help/doc-detail/86980.htm).
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
    ///         // Create a cen_route_entry resource and use it to publish a route entry pointing to an ECS.
    ///         var hz = new AliCloud.Provider("hz", new AliCloud.ProviderArgs
    ///         {
    ///             Region = "cn-hangzhou",
    ///         });
    ///         var config = new Config();
    ///         var name = config.Get("name") ?? "tf-testAccCenRouteEntryConfig";
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
    ///         var vpc = new AliCloud.Vpc.Network("vpc", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             VpcName = name,
    ///             CidrBlock = "172.16.0.0/12",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "172.16.0.0/21",
    ///             ZoneId = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones?[0]?.Id),
    ///             VswitchName = name,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///         });
    ///         var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new AliCloud.Ecs.SecurityGroupArgs
    ///         {
    ///             Description = "foo",
    ///             VpcId = vpc.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///         });
    ///         var defaultInstance = new AliCloud.Ecs.Instance("defaultInstance", new AliCloud.Ecs.InstanceArgs
    ///         {
    ///             VswitchId = defaultSwitch.Id,
    ///             ImageId = defaultImages.Apply(defaultImages =&gt; defaultImages.Images?[0]?.Id),
    ///             InstanceType = defaultInstanceTypes.Apply(defaultInstanceTypes =&gt; defaultInstanceTypes.InstanceTypes?[0]?.Id),
    ///             SystemDiskCategory = "cloud_efficiency",
    ///             InternetChargeType = "PayByTraffic",
    ///             InternetMaxBandwidthOut = 5,
    ///             SecurityGroups = 
    ///             {
    ///                 defaultSecurityGroup.Id,
    ///             },
    ///             InstanceName = name,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///         });
    ///         var cen = new AliCloud.Cen.Instance("cen", new AliCloud.Cen.InstanceArgs
    ///         {
    ///         });
    ///         var attach = new AliCloud.Cen.InstanceAttachment("attach", new AliCloud.Cen.InstanceAttachmentArgs
    ///         {
    ///             InstanceId = cen.Id,
    ///             ChildInstanceId = vpc.Id,
    ///             ChildInstanceType = "VPC",
    ///             ChildInstanceRegionId = "cn-hangzhou",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 defaultSwitch,
    ///             },
    ///         });
    ///         var route = new AliCloud.Vpc.RouteEntry("route", new AliCloud.Vpc.RouteEntryArgs
    ///         {
    ///             RouteTableId = vpc.RouteTableId,
    ///             DestinationCidrblock = "11.0.0.0/16",
    ///             NexthopType = "Instance",
    ///             NexthopId = defaultInstance.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///         });
    ///         var foo = new AliCloud.Cen.RouteEntry("foo", new AliCloud.Cen.RouteEntryArgs
    ///         {
    ///             InstanceId = cen.Id,
    ///             RouteTableId = vpc.RouteTableId,
    ///             CidrBlock = route.DestinationCidrblock,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = alicloud.Hz,
    ///             DependsOn = 
    ///             {
    ///                 attach,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/routeEntry:RouteEntry example cen-abc123456:vtb-abc123:192.168.0.0/24
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/routeEntry:RouteEntry")]
    public partial class RouteEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteEntry(string name, RouteEntryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeEntry:RouteEntry", name, args ?? new RouteEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteEntry(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/routeEntry:RouteEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteEntry Get(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteEntry(name, id, state, options);
        }
    }

    public sealed class RouteEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public RouteEntryArgs()
        {
        }
    }

    public sealed class RouteEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry to publish.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The route table of the attached VBR or VPC.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        public RouteEntryState()
        {
        }
    }
}
