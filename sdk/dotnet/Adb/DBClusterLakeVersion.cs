// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    /// <summary>
    /// Provides a AnalyticDB for MySQL (ADB) DB Cluster Lake Version resource.
    /// 
    /// For information about AnalyticDB for MySQL (ADB) DB Cluster Lake Version and how to use it, see [What is DB Cluster Lake Version](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/what-is-analyticdb-for-mysql).
    /// 
    /// &gt; **NOTE:** Available in v1.190.0+.
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
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "^default-NODELETING",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = "example",
    ///     });
    /// 
    ///     var defaultDBClusterLakeVersion = new AliCloud.Adb.DBClusterLakeVersion("defaultDBClusterLakeVersion", new()
    ///     {
    ///         ComputeResource = "16ACU",
    ///         DbClusterVersion = "5.0",
    ///         PaymentType = "PayAsYouGo",
    ///         StorageResource = "0ACU",
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AnalyticDB for MySQL (ADB) DB Cluster Lake Version can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion")]
    public partial class DBClusterLakeVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the service.
        /// </summary>
        [Output("commodityCode")]
        public Output<string> CommodityCode { get; private set; } = null!;

        /// <summary>
        /// The computing resources of the cluster.
        /// </summary>
        [Output("computeResource")]
        public Output<string> ComputeResource { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the cluster.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The createTime of the cluster.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Output("dbClusterDescription")]
        public Output<string> DbClusterDescription { get; private set; } = null!;

        /// <summary>
        /// The version of the cluster. Value options: `5.0`.
        /// </summary>
        [Output("dbClusterVersion")]
        public Output<string> DbClusterVersion { get; private set; } = null!;

        /// <summary>
        /// Whether to enable default allocation of resources to user_default resource groups.
        /// </summary>
        [Output("enableDefaultResourceGroup")]
        public Output<bool?> EnableDefaultResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The engine of the database.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// The engine version of the database.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The time when the cluster expires.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the cluster has expired.
        /// </summary>
        [Output("expired")]
        public Output<string> Expired { get; private set; } = null!;

        /// <summary>
        /// The lock mode of the cluster.
        /// </summary>
        [Output("lockMode")]
        public Output<string> LockMode { get; private set; } = null!;

        /// <summary>
        /// The reason why the cluster is locked.
        /// </summary>
        [Output("lockReason")]
        public Output<string> LockReason { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The port that is used to access the cluster.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
        /// - IP addresses, such as 10.23.XX.XX.
        /// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
        /// </summary>
        [Output("securityIps")]
        public Output<string> SecurityIps { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The storage resources of the cluster.
        /// </summary>
        [Output("storageResource")]
        public Output<string> StorageResource { get; private set; } = null!;

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The zone ID of the resource.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DBClusterLakeVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBClusterLakeVersion(string name, DBClusterLakeVersionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion", name, args ?? new DBClusterLakeVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBClusterLakeVersion(string name, Input<string> id, DBClusterLakeVersionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:adb/dBClusterLakeVersion:DBClusterLakeVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DBClusterLakeVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBClusterLakeVersion Get(string name, Input<string> id, DBClusterLakeVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new DBClusterLakeVersion(name, id, state, options);
        }
    }

    public sealed class DBClusterLakeVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The computing resources of the cluster.
        /// </summary>
        [Input("computeResource", required: true)]
        public Input<string> ComputeResource { get; set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("dbClusterDescription")]
        public Input<string>? DbClusterDescription { get; set; }

        /// <summary>
        /// The version of the cluster. Value options: `5.0`.
        /// </summary>
        [Input("dbClusterVersion", required: true)]
        public Input<string> DbClusterVersion { get; set; } = null!;

        /// <summary>
        /// Whether to enable default allocation of resources to user_default resource groups.
        /// </summary>
        [Input("enableDefaultResourceGroup")]
        public Input<bool>? EnableDefaultResourceGroup { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo`.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
        /// - IP addresses, such as 10.23.XX.XX.
        /// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
        /// </summary>
        [Input("securityIps")]
        public Input<string>? SecurityIps { get; set; }

        /// <summary>
        /// The storage resources of the cluster.
        /// </summary>
        [Input("storageResource", required: true)]
        public Input<string> StorageResource { get; set; } = null!;

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The zone ID of the resource.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public DBClusterLakeVersionArgs()
        {
        }
        public static new DBClusterLakeVersionArgs Empty => new DBClusterLakeVersionArgs();
    }

    public sealed class DBClusterLakeVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the service.
        /// </summary>
        [Input("commodityCode")]
        public Input<string>? CommodityCode { get; set; }

        /// <summary>
        /// The computing resources of the cluster.
        /// </summary>
        [Input("computeResource")]
        public Input<string>? ComputeResource { get; set; }

        /// <summary>
        /// The endpoint of the cluster.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The createTime of the cluster.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("dbClusterDescription")]
        public Input<string>? DbClusterDescription { get; set; }

        /// <summary>
        /// The version of the cluster. Value options: `5.0`.
        /// </summary>
        [Input("dbClusterVersion")]
        public Input<string>? DbClusterVersion { get; set; }

        /// <summary>
        /// Whether to enable default allocation of resources to user_default resource groups.
        /// </summary>
        [Input("enableDefaultResourceGroup")]
        public Input<bool>? EnableDefaultResourceGroup { get; set; }

        /// <summary>
        /// The engine of the database.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The engine version of the database.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The time when the cluster expires.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Indicates whether the cluster has expired.
        /// </summary>
        [Input("expired")]
        public Input<string>? Expired { get; set; }

        /// <summary>
        /// The lock mode of the cluster.
        /// </summary>
        [Input("lockMode")]
        public Input<string>? LockMode { get; set; }

        /// <summary>
        /// The reason why the cluster is locked.
        /// </summary>
        [Input("lockReason")]
        public Input<string>? LockReason { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values are `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The port that is used to access the cluster.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
        /// - IP addresses, such as 10.23.XX.XX.
        /// - CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
        /// </summary>
        [Input("securityIps")]
        public Input<string>? SecurityIps { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The storage resources of the cluster.
        /// </summary>
        [Input("storageResource")]
        public Input<string>? StorageResource { get; set; }

        /// <summary>
        /// The vpc ID of the resource.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The zone ID of the resource.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DBClusterLakeVersionState()
        {
        }
        public static new DBClusterLakeVersionState Empty => new DBClusterLakeVersionState();
    }
}