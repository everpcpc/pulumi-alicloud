// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
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
    ///     var foo = new AliCloud.Vpc.Network("foo", new()
    ///     {
    ///         VpcName = "tf_test_foo12345",
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var @interface = new AliCloud.Vpc.RouterInterface("interface", new()
    ///     {
    ///         OppositeRegion = "cn-beijing",
    ///         RouterType = "VRouter",
    ///         RouterId = foo.RouterId,
    ///         Role = "InitiatingSide",
    ///         Specification = "Large.2",
    ///         Description = "test1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The router interface can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/routerInterface:RouterInterface interface ri-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/routerInterface:RouterInterface")]
    public partial class RouterInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// It has been deprecated from version 1.11.0.
        /// </summary>
        [Output("accessPointId")]
        public Output<string> AccessPointId { get; private set; } = null!;

        /// <summary>
        /// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_target_ip` must be specified at the same time.
        /// </summary>
        [Output("healthCheckSourceIp")]
        public Output<string?> HealthCheckSourceIp { get; private set; } = null!;

        /// <summary>
        /// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_source_ip` must be specified at the same time.
        /// </summary>
        [Output("healthCheckTargetIp")]
        public Output<string?> HealthCheckTargetIp { get; private set; } = null!;

        /// <summary>
        /// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and opposite_region are the same.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.11.0.
        /// </summary>
        [Output("oppositeAccessPointId")]
        public Output<string?> OppositeAccessPointId { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
        /// </summary>
        [Output("oppositeInterfaceId")]
        public Output<string> OppositeInterfaceId { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
        /// </summary>
        [Output("oppositeInterfaceOwnerId")]
        public Output<string> OppositeInterfaceOwnerId { get; private set; } = null!;

        /// <summary>
        /// The Region of peer side.
        /// </summary>
        [Output("oppositeRegion")]
        public Output<string> OppositeRegion { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
        /// </summary>
        [Output("oppositeRouterId")]
        public Output<string> OppositeRouterId { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
        /// </summary>
        [Output("oppositeRouterType")]
        public Output<string> OppositeRouterType { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The Router ID.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
        /// </summary>
        [Output("routerType")]
        public Output<string> RouterType { get; private set; } = null!;

        /// <summary>
        /// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
        /// </summary>
        [Output("specification")]
        public Output<string?> Specification { get; private set; } = null!;


        /// <summary>
        /// Create a RouterInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterInterface(string name, RouterInterfaceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routerInterface:RouterInterface", name, args ?? new RouterInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterInterface(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/routerInterface:RouterInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterInterface Get(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterInterface(name, id, state, options);
        }
    }

    public sealed class RouterInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_target_ip` must be specified at the same time.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_source_ip` must be specified at the same time.
        /// </summary>
        [Input("healthCheckTargetIp")]
        public Input<string>? HealthCheckTargetIp { get; set; }

        /// <summary>
        /// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and opposite_region are the same.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0.
        /// </summary>
        [Input("oppositeAccessPointId")]
        public Input<string>? OppositeAccessPointId { get; set; }

        /// <summary>
        /// The Region of peer side.
        /// </summary>
        [Input("oppositeRegion", required: true)]
        public Input<string> OppositeRegion { get; set; } = null!;

        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The Router ID.
        /// </summary>
        [Input("routerId", required: true)]
        public Input<string> RouterId { get; set; } = null!;

        /// <summary>
        /// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
        /// </summary>
        [Input("routerType", required: true)]
        public Input<string> RouterType { get; set; } = null!;

        /// <summary>
        /// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
        /// </summary>
        [Input("specification")]
        public Input<string>? Specification { get; set; }

        public RouterInterfaceArgs()
        {
        }
        public static new RouterInterfaceArgs Empty => new RouterInterfaceArgs();
    }

    public sealed class RouterInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// It has been deprecated from version 1.11.0.
        /// </summary>
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        /// <summary>
        /// Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_target_ip` must be specified at the same time.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `router_type` is `VBR`. The IP must be an unused IP in the local VPC. It and `health_check_source_ip` must be specified at the same time.
        /// </summary>
        [Input("healthCheckTargetIp")]
        public Input<string>? HealthCheckTargetIp { get; set; }

        /// <summary>
        /// The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and opposite_region are the same.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0.
        /// </summary>
        [Input("oppositeAccessPointId")]
        public Input<string>? OppositeAccessPointId { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
        /// </summary>
        [Input("oppositeInterfaceId")]
        public Input<string>? OppositeInterfaceId { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
        /// </summary>
        [Input("oppositeInterfaceOwnerId")]
        public Input<string>? OppositeInterfaceOwnerId { get; set; }

        /// <summary>
        /// The Region of peer side.
        /// </summary>
        [Input("oppositeRegion")]
        public Input<string>? OppositeRegion { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
        /// </summary>
        [Input("oppositeRouterId")]
        public Input<string>? OppositeRouterId { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
        /// </summary>
        [Input("oppositeRouterType")]
        public Input<string>? OppositeRouterType { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The Router ID.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
        /// </summary>
        [Input("routerType")]
        public Input<string>? RouterType { get; set; }

        /// <summary>
        /// Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
        /// </summary>
        [Input("specification")]
        public Input<string>? Specification { get; set; }

        public RouterInterfaceState()
        {
        }
        public static new RouterInterfaceState Empty => new RouterInterfaceState();
    }
}
