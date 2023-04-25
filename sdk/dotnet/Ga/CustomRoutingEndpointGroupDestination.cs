// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Custom Routing Endpoint Group Destination resource.
    /// 
    /// For information about Global Accelerator (GA) Custom Routing Endpoint Group Destination and how to use it, see [What is Custom Routing Endpoint Group Destination](https://www.alibabacloud.com/help/en/global-accelerator/latest/createcustomroutingendpointgroupdestinations).
    /// 
    /// &gt; **NOTE:** Available in v1.197.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new AliCloud.Ga.CustomRoutingEndpointGroupDestination("default", new()
    ///     {
    ///         EndpointGroupId = "your_custom_routing_endpoint_group_id",
    ///         FromPort = 1,
    ///         Protocols = new[]
    ///         {
    ///             "TCP",
    ///             "UDP",
    ///         },
    ///         ToPort = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Custom Routing Endpoint Group Destination can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination example &lt;endpoint_group_id&gt;:&lt;custom_routing_endpoint_group_destination_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination")]
    public partial class CustomRoutingEndpointGroupDestination : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the GA instance.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Custom Routing Endpoint Group Destination.
        /// </summary>
        [Output("customRoutingEndpointGroupDestinationId")]
        public Output<string> CustomRoutingEndpointGroupDestinationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the endpoint group.
        /// </summary>
        [Output("endpointGroupId")]
        public Output<string> EndpointGroupId { get; private set; } = null!;

        /// <summary>
        /// The start port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Output("fromPort")]
        public Output<int> FromPort { get; private set; } = null!;

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// The status of the Custom Routing Endpoint Group Destination.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The end port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Output("toPort")]
        public Output<int> ToPort { get; private set; } = null!;


        /// <summary>
        /// Create a CustomRoutingEndpointGroupDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomRoutingEndpointGroupDestination(string name, CustomRoutingEndpointGroupDestinationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination", name, args ?? new CustomRoutingEndpointGroupDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomRoutingEndpointGroupDestination(string name, Input<string> id, CustomRoutingEndpointGroupDestinationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomRoutingEndpointGroupDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomRoutingEndpointGroupDestination Get(string name, Input<string> id, CustomRoutingEndpointGroupDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomRoutingEndpointGroupDestination(name, id, state, options);
        }
    }

    public sealed class CustomRoutingEndpointGroupDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint group.
        /// </summary>
        [Input("endpointGroupId", required: true)]
        public Input<string> EndpointGroupId { get; set; } = null!;

        /// <summary>
        /// The start port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("protocols", required: true)]
        private InputList<string>? _protocols;

        /// <summary>
        /// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The end port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public CustomRoutingEndpointGroupDestinationArgs()
        {
        }
        public static new CustomRoutingEndpointGroupDestinationArgs Empty => new CustomRoutingEndpointGroupDestinationArgs();
    }

    public sealed class CustomRoutingEndpointGroupDestinationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the GA instance.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        /// <summary>
        /// The ID of the Custom Routing Endpoint Group Destination.
        /// </summary>
        [Input("customRoutingEndpointGroupDestinationId")]
        public Input<string>? CustomRoutingEndpointGroupDestinationId { get; set; }

        /// <summary>
        /// The ID of the endpoint group.
        /// </summary>
        [Input("endpointGroupId")]
        public Input<string>? EndpointGroupId { get; set; }

        /// <summary>
        /// The start port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// The status of the Custom Routing Endpoint Group Destination.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The end port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public CustomRoutingEndpointGroupDestinationState()
        {
        }
        public static new CustomRoutingEndpointGroupDestinationState Empty => new CustomRoutingEndpointGroupDestinationState();
    }
}
