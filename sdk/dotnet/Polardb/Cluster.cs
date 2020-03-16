// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB
{
    /// <summary>
    /// Provides a PolarDB cluster resource. A PolarDB cluster is an isolated database
    /// environment in the cloud. A PolarDB cluster can contain multiple user-created
    /// databases.
    /// 
    /// &gt; **NOTE:** Available in v1.66.0+.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_cluster.html.markdown.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Output("dbNodeClass")]
        public Output<string> DbNodeClass { get; private set; } = null!;

        /// <summary>
        /// Database type. Value options: MySQL, Oracle, PostgreSQL.
        /// </summary>
        [Output("dbType")]
        public Output<string> DbType { get; private set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
        /// </summary>
        [Output("dbVersion")]
        public Output<string> DbVersion { get; private set; } = null!;

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Output("maintainTime")]
        public Output<string> MaintainTime { get; private set; } = null!;

        /// <summary>
        /// Use as `db_node_class` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
        /// </summary>
        [Output("modifyType")]
        public Output<string?> ModifyType { get; private set; } = null!;

        /// <summary>
        /// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ClusterParameters>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Output("payType")]
        public Output<string?> PayType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Output("renewalStatus")]
        public Output<string?> RenewalStatus { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        [Output("securityIps")]
        public Output<ImmutableArray<string>> SecurityIps { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the DB cluster. it supports multiple zone.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:polardb/cluster:Cluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:polardb/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass", required: true)]
        public Input<string> DbNodeClass { get; set; } = null!;

        /// <summary>
        /// Database type. Value options: MySQL, Oracle, PostgreSQL.
        /// </summary>
        [Input("dbType", required: true)]
        public Input<string> DbType { get; set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
        /// </summary>
        [Input("dbVersion", required: true)]
        public Input<string> DbVersion { get; set; } = null!;

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// Use as `db_node_class` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
        /// </summary>
        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ClusterParametersArgs>? _parameters;

        /// <summary>
        /// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        public InputList<Inputs.ClusterParametersArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ClusterParametersArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster. it supports multiple zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass")]
        public Input<string>? DbNodeClass { get; set; }

        /// <summary>
        /// Database type. Value options: MySQL, Oracle, PostgreSQL.
        /// </summary>
        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
        /// </summary>
        [Input("dbVersion")]
        public Input<string>? DbVersion { get; set; }

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// Use as `db_node_class` change class , define upgrade or downgrade.  Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
        /// </summary>
        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ClusterParametersGetArgs>? _parameters;

        /// <summary>
        /// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        public InputList<Inputs.ClusterParametersGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ClusterParametersGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster. it supports multiple zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ClusterParametersArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterParametersArgs()
        {
        }
    }

    public sealed class ClusterParametersGetArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ClusterParametersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ClusterParameters
    {
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private ClusterParameters(
            string name,
            string value)
        {
            Name = name;
            Value = value;
        }
    }
    }
}
