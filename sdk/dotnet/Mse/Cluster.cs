// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mse
{
    /// <summary>
    /// Provides a MSE Cluster resource. It is a one-stop microservice platform for the industry's mainstream open source microservice frameworks Spring Cloud and Dubbo, providing governance center, managed registry and managed configuration center.
    /// 
    /// &gt; **NOTE:** Available in 1.94.0+.
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
    ///     var example = new AliCloud.Mse.Cluster("example", new()
    ///     {
    ///         AclEntryLists = new[]
    ///         {
    ///             "127.0.0.1/32",
    ///         },
    ///         ClusterAliasName = "tf-testAccMseCluster",
    ///         ClusterSpecification = "MSE_SC_1_2_200_c",
    ///         ClusterType = "Nacos-Ans",
    ///         ClusterVersion = "NACOS_ANS_1_2_1",
    ///         InstanceCount = 1,
    ///         MseVersion = "mse_dev",
    ///         NetType = "privatenet",
    ///         PubNetworkFlow = "1",
    ///         VswitchId = "vsw-123456",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MSE Cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:mse/cluster:Cluster example mse-cn-0d9xxxx
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:mse/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        /// </summary>
        [Output("aclEntryLists")]
        public Output<ImmutableArray<string>> AclEntryLists { get; private set; } = null!;

        /// <summary>
        /// The alias of MSE Cluster.
        /// </summary>
        [Output("clusterAliasName")]
        public Output<string?> ClusterAliasName { get; private set; } = null!;

        /// <summary>
        /// (Available in v1.162.0+)  The id of Cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The engine specification of MSE Cluster. **NOTE:** From version 1.188.0, `cluster_specification` can be modified. Valid values:
        /// </summary>
        [Output("clusterSpecification")]
        public Output<string> ClusterSpecification { get; private set; } = null!;

        /// <summary>
        /// The type of MSE Cluster.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// The version of MSE Cluster. See [details](https://www.alibabacloud.com/help/en/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createcluster)
        /// </summary>
        [Output("clusterVersion")]
        public Output<string> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// The connection type. Valid values: `slb`.
        /// </summary>
        [Output("connectionType")]
        public Output<string> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// The type of Disk.
        /// </summary>
        [Output("diskType")]
        public Output<string?> DiskType { get; private set; } = null!;

        /// <summary>
        /// The count of instance. **NOTE:** From version 1.188.0, `instance_count` can be modified.
        /// </summary>
        [Output("instanceCount")]
        public Output<int> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// The version of MSE. Valid values: `mse_dev` or `mse_pro`.
        /// </summary>
        [Output("mseVersion")]
        public Output<string> MseVersion { get; private set; } = null!;

        /// <summary>
        /// The type of network. Valid values: "privatenet" and "pubnet".
        /// </summary>
        [Output("netType")]
        public Output<string> NetType { get; private set; } = null!;

        /// <summary>
        /// The specification of private network SLB.
        /// </summary>
        [Output("privateSlbSpecification")]
        public Output<string?> PrivateSlbSpecification { get; private set; } = null!;

        /// <summary>
        /// The public network bandwidth. `0` means no access to the public network.
        /// </summary>
        [Output("pubNetworkFlow")]
        public Output<string> PubNetworkFlow { get; private set; } = null!;

        /// <summary>
        /// The specification of public network SLB.
        /// </summary>
        [Output("pubSlbSpecification")]
        public Output<string?> PubSlbSpecification { get; private set; } = null!;

        /// <summary>
        /// The extended request parameters in the JSON format.
        /// </summary>
        [Output("requestPars")]
        public Output<string?> RequestPars { get; private set; } = null!;

        /// <summary>
        /// The status of MSE Cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The id of VSwitch.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mse/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mse/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("aclEntryLists")]
        private InputList<string>? _aclEntryLists;

        /// <summary>
        /// The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        /// </summary>
        public InputList<string> AclEntryLists
        {
            get => _aclEntryLists ?? (_aclEntryLists = new InputList<string>());
            set => _aclEntryLists = value;
        }

        /// <summary>
        /// The alias of MSE Cluster.
        /// </summary>
        [Input("clusterAliasName")]
        public Input<string>? ClusterAliasName { get; set; }

        /// <summary>
        /// The engine specification of MSE Cluster. **NOTE:** From version 1.188.0, `cluster_specification` can be modified. Valid values:
        /// </summary>
        [Input("clusterSpecification", required: true)]
        public Input<string> ClusterSpecification { get; set; } = null!;

        /// <summary>
        /// The type of MSE Cluster.
        /// </summary>
        [Input("clusterType", required: true)]
        public Input<string> ClusterType { get; set; } = null!;

        /// <summary>
        /// The version of MSE Cluster. See [details](https://www.alibabacloud.com/help/en/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createcluster)
        /// </summary>
        [Input("clusterVersion", required: true)]
        public Input<string> ClusterVersion { get; set; } = null!;

        /// <summary>
        /// The connection type. Valid values: `slb`.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The type of Disk.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The count of instance. **NOTE:** From version 1.188.0, `instance_count` can be modified.
        /// </summary>
        [Input("instanceCount", required: true)]
        public Input<int> InstanceCount { get; set; } = null!;

        /// <summary>
        /// The version of MSE. Valid values: `mse_dev` or `mse_pro`.
        /// </summary>
        [Input("mseVersion")]
        public Input<string>? MseVersion { get; set; }

        /// <summary>
        /// The type of network. Valid values: "privatenet" and "pubnet".
        /// </summary>
        [Input("netType", required: true)]
        public Input<string> NetType { get; set; } = null!;

        /// <summary>
        /// The specification of private network SLB.
        /// </summary>
        [Input("privateSlbSpecification")]
        public Input<string>? PrivateSlbSpecification { get; set; }

        /// <summary>
        /// The public network bandwidth. `0` means no access to the public network.
        /// </summary>
        [Input("pubNetworkFlow", required: true)]
        public Input<string> PubNetworkFlow { get; set; } = null!;

        /// <summary>
        /// The specification of public network SLB.
        /// </summary>
        [Input("pubSlbSpecification")]
        public Input<string>? PubSlbSpecification { get; set; }

        /// <summary>
        /// The extended request parameters in the JSON format.
        /// </summary>
        [Input("requestPars")]
        public Input<string>? RequestPars { get; set; }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The id of VSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        [Input("aclEntryLists")]
        private InputList<string>? _aclEntryLists;

        /// <summary>
        /// The whitelist. **NOTE:** This attribute is invalid when the value of `pub_network_flow` is `0` and the value of `net_type` is `privatenet`.
        /// </summary>
        public InputList<string> AclEntryLists
        {
            get => _aclEntryLists ?? (_aclEntryLists = new InputList<string>());
            set => _aclEntryLists = value;
        }

        /// <summary>
        /// The alias of MSE Cluster.
        /// </summary>
        [Input("clusterAliasName")]
        public Input<string>? ClusterAliasName { get; set; }

        /// <summary>
        /// (Available in v1.162.0+)  The id of Cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The engine specification of MSE Cluster. **NOTE:** From version 1.188.0, `cluster_specification` can be modified. Valid values:
        /// </summary>
        [Input("clusterSpecification")]
        public Input<string>? ClusterSpecification { get; set; }

        /// <summary>
        /// The type of MSE Cluster.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// The version of MSE Cluster. See [details](https://www.alibabacloud.com/help/en/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createcluster)
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// The connection type. Valid values: `slb`.
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The type of Disk.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The count of instance. **NOTE:** From version 1.188.0, `instance_count` can be modified.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The version of MSE. Valid values: `mse_dev` or `mse_pro`.
        /// </summary>
        [Input("mseVersion")]
        public Input<string>? MseVersion { get; set; }

        /// <summary>
        /// The type of network. Valid values: "privatenet" and "pubnet".
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        /// <summary>
        /// The specification of private network SLB.
        /// </summary>
        [Input("privateSlbSpecification")]
        public Input<string>? PrivateSlbSpecification { get; set; }

        /// <summary>
        /// The public network bandwidth. `0` means no access to the public network.
        /// </summary>
        [Input("pubNetworkFlow")]
        public Input<string>? PubNetworkFlow { get; set; }

        /// <summary>
        /// The specification of public network SLB.
        /// </summary>
        [Input("pubSlbSpecification")]
        public Input<string>? PubSlbSpecification { get; set; }

        /// <summary>
        /// The extended request parameters in the JSON format.
        /// </summary>
        [Input("requestPars")]
        public Input<string>? RequestPars { get; set; }

        /// <summary>
        /// The status of MSE Cluster.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The id of VSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
