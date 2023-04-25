// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Provide RDS cluster instance endpoint connection resources.
    /// &gt; **NOTE:** Available in 1.203.0+.
    /// 
    /// ## Block node_items
    /// 
    /// The node_items mapping supports the following:
    /// 
    /// * `node_id` - (Required) The ID of the node.
    /// * `weight` - (Required) The weight of the node. Read requests are distributed based on the weight.Valid values: 0 to 100.
    /// 
    /// ## Import
    /// 
    /// RDS database endpoint feature can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint example &lt;db_instance_id&gt;:&lt;db_instance_endpoint_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint")]
    public partial class DbInstanceEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The internal endpoint.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The IP address of the internal endpoint.
        /// </summary>
        [Output("connectionStringPrefix")]
        public Output<string> ConnectionStringPrefix { get; private set; } = null!;

        /// <summary>
        /// The user-defined description of the endpoint.
        /// </summary>
        [Output("dbInstanceEndpointDescription")]
        public Output<string> DbInstanceEndpointDescription { get; private set; } = null!;

        /// <summary>
        /// The Endpoint ID of the instance.
        /// </summary>
        [Output("dbInstanceEndpointId")]
        public Output<string> DbInstanceEndpointId { get; private set; } = null!;

        /// <summary>
        /// The type of the endpoint.
        /// </summary>
        [Output("dbInstanceEndpointType")]
        public Output<string> DbInstanceEndpointType { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the IP address.
        /// </summary>
        [Output("ipType")]
        public Output<string> IpType { get; private set; } = null!;

        /// <summary>
        /// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        /// </summary>
        [Output("nodeItems")]
        public Output<ImmutableArray<Outputs.DbInstanceEndpointNodeItem>> NodeItems { get; private set; } = null!;

        /// <summary>
        /// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The IP address of the internal endpoint.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The virtual private cloud (VPC) ID of the internal endpoint.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vSwitch ID of the internal endpoint.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a DbInstanceEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbInstanceEndpoint(string name, DbInstanceEndpointArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint", name, args ?? new DbInstanceEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbInstanceEndpoint(string name, Input<string> id, DbInstanceEndpointState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DbInstanceEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbInstanceEndpoint Get(string name, Input<string> id, DbInstanceEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new DbInstanceEndpoint(name, id, state, options);
        }
    }

    public sealed class DbInstanceEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the internal endpoint.
        /// </summary>
        [Input("connectionStringPrefix", required: true)]
        public Input<string> ConnectionStringPrefix { get; set; } = null!;

        /// <summary>
        /// The user-defined description of the endpoint.
        /// </summary>
        [Input("dbInstanceEndpointDescription")]
        public Input<string>? DbInstanceEndpointDescription { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        [Input("nodeItems", required: true)]
        private InputList<Inputs.DbInstanceEndpointNodeItemArgs>? _nodeItems;

        /// <summary>
        /// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        /// </summary>
        public InputList<Inputs.DbInstanceEndpointNodeItemArgs> NodeItems
        {
            get => _nodeItems ?? (_nodeItems = new InputList<Inputs.DbInstanceEndpointNodeItemArgs>());
            set => _nodeItems = value;
        }

        /// <summary>
        /// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        /// </summary>
        [Input("port", required: true)]
        public Input<string> Port { get; set; } = null!;

        /// <summary>
        /// The virtual private cloud (VPC) ID of the internal endpoint.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The vSwitch ID of the internal endpoint.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public DbInstanceEndpointArgs()
        {
        }
        public static new DbInstanceEndpointArgs Empty => new DbInstanceEndpointArgs();
    }

    public sealed class DbInstanceEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The internal endpoint.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The IP address of the internal endpoint.
        /// </summary>
        [Input("connectionStringPrefix")]
        public Input<string>? ConnectionStringPrefix { get; set; }

        /// <summary>
        /// The user-defined description of the endpoint.
        /// </summary>
        [Input("dbInstanceEndpointDescription")]
        public Input<string>? DbInstanceEndpointDescription { get; set; }

        /// <summary>
        /// The Endpoint ID of the instance.
        /// </summary>
        [Input("dbInstanceEndpointId")]
        public Input<string>? DbInstanceEndpointId { get; set; }

        /// <summary>
        /// The type of the endpoint.
        /// </summary>
        [Input("dbInstanceEndpointType")]
        public Input<string>? DbInstanceEndpointType { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The type of the IP address.
        /// </summary>
        [Input("ipType")]
        public Input<string>? IpType { get; set; }

        [Input("nodeItems")]
        private InputList<Inputs.DbInstanceEndpointNodeItemGetArgs>? _nodeItems;

        /// <summary>
        /// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
        /// </summary>
        public InputList<Inputs.DbInstanceEndpointNodeItemGetArgs> NodeItems
        {
            get => _nodeItems ?? (_nodeItems = new InputList<Inputs.DbInstanceEndpointNodeItemGetArgs>());
            set => _nodeItems = value;
        }

        /// <summary>
        /// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The IP address of the internal endpoint.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The virtual private cloud (VPC) ID of the internal endpoint.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vSwitch ID of the internal endpoint.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public DbInstanceEndpointState()
        {
        }
        public static new DbInstanceEndpointState Empty => new DbInstanceEndpointState();
    }
}
