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
    /// Provides an RDS instance resource. A DB instance is an isolated database environment in the cloud. A DB instance can contain multiple user-created databases.
    /// 
    /// For information about RDS and how to use it, see [What is ApsaraDB for RDS](https://www.alibabacloud.com/help/en/doc-detail/26092.htm).
    /// 
    /// ## Example Usage
    /// ### Create a RDS MySQL instance
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
    ///         var name = config.Get("name") ?? "tf-testaccdbinstance";
    ///         var creation = config.Get("creation") ?? "Rds";
    ///         var exampleZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableResourceCreation = creation,
    ///         }));
    ///         var exampleNetwork = new AliCloud.Vpc.Network("exampleNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var exampleSwitch = new AliCloud.Vpc.Switch("exampleSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = exampleNetwork.Id,
    ///             CidrBlock = "172.16.0.0/24",
    ///             AvailabilityZone = exampleZones.Apply(exampleZones =&gt; exampleZones.Zones[0].Id),
    ///         });
    ///         var exampleInstance = new AliCloud.Rds.Instance("exampleInstance", new AliCloud.Rds.InstanceArgs
    ///         {
    ///             Engine = "MySQL",
    ///             EngineVersion = "5.6",
    ///             InstanceType = "rds.mysql.s2.large",
    ///             InstanceStorage = 30,
    ///             InstanceChargeType = "Postpaid",
    ///             InstanceName = name,
    ///             VswitchId = exampleSwitch.Id,
    ///             MonitoringPeriod = 60,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// The upgrade method to use. Valid values:
        /// - Auto: Instances are automatically upgraded to a higher minor version.
        /// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
        /// </summary>
        [Output("autoUpgradeMinorVersion")]
        public Output<string> AutoUpgradeMinorVersion { get; private set; } = null!;

        /// <summary>
        /// RDS database connection string.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The storage type of the instance. Valid values:
        /// - local_ssd: specifies to use local SSDs. This value is recommended.
        /// - cloud_ssd: specifies to use standard SSDs.
        /// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
        /// </summary>
        [Output("dbInstanceStorageType")]
        public Output<string> DbInstanceStorageType { get; private set; } = null!;

        /// <summary>
        /// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Set it to true to make some parameter efficient when modifying them. Default to false.
        /// </summary>
        [Output("forceRestart")]
        public Output<bool?> ForceRestart { get; private set; } = null!;

        /// <summary>
        /// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Output("instanceName")]
        public Output<string?> InstanceName { get; private set; } = null!;

        /// <summary>
        /// User-defined DB instance storage space. Value range:
        /// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
        /// - [20,1000] for MySQL 5.7 basic single node edition;
        /// - [10, 2000] for SQL Server 2008R2;
        /// - [20,2000] for SQL Server 2012 basic single node edition
        /// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instance_storage`.
        /// </summary>
        [Output("instanceStorage")]
        public Output<int> InstanceStorage { get; private set; } = null!;

        /// <summary>
        /// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Output("maintainTime")]
        public Output<string> MaintainTime { get; private set; } = null!;

        /// <summary>
        /// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
        /// </summary>
        [Output("monitoringPeriod")]
        public Output<int> MonitoringPeriod { get; private set; } = null!;

        /// <summary>
        /// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.InstanceParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// RDS database connection port.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of resource group which the DB instance belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from 1.69.0 and use `security_group_ids` instead.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
        /// </summary>
        [Output("securityIpMode")]
        public Output<string?> SecurityIpMode { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        [Output("securityIps")]
        public Output<ImmutableArray<string>> SecurityIps { get; private set; } = null!;

        /// <summary>
        /// The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
        /// </summary>
        [Output("sqlCollectorConfigValue")]
        public Output<int?> SqlCollectorConfigValue { get; private set; } = null!;

        /// <summary>
        /// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
        /// </summary>
        [Output("sqlCollectorStatus")]
        public Output<string> SqlCollectorStatus { get; private set; } = null!;

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
        /// </summary>
        [Output("sslAction")]
        public Output<string> SslAction { get; private set; } = null!;

        /// <summary>
        /// Status of the SSL feature. `Yes`: SSL is turned on; `No`: SSL is turned off.
        /// </summary>
        [Output("sslStatus")]
        public Output<string> SslStatus { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
        /// </summary>
        [Output("tdeStatus")]
        public Output<string?> TdeStatus { get; private set; } = null!;

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in the one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;

        /// <summary>
        /// The region ID of the secondary instance if you create a secondary instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Output("zoneIdSlaveA")]
        public Output<string?> ZoneIdSlaveA { get; private set; } = null!;

        /// <summary>
        /// The region ID of the log instance if you create a log instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Output("zoneIdSlaveB")]
        public Output<string?> ZoneIdSlaveB { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The upgrade method to use. Valid values:
        /// - Auto: Instances are automatically upgraded to a higher minor version.
        /// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<string>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// The storage type of the instance. Valid values:
        /// - local_ssd: specifies to use local SSDs. This value is recommended.
        /// - cloud_ssd: specifies to use standard SSDs.
        /// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
        /// </summary>
        [Input("dbInstanceStorageType")]
        public Input<string>? DbInstanceStorageType { get; set; }

        /// <summary>
        /// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// Set it to true to make some parameter efficient when modifying them. Default to false.
        /// </summary>
        [Input("forceRestart")]
        public Input<bool>? ForceRestart { get; set; }

        /// <summary>
        /// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// User-defined DB instance storage space. Value range:
        /// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
        /// - [20,1000] for MySQL 5.7 basic single node edition;
        /// - [10, 2000] for SQL Server 2008R2;
        /// - [20,2000] for SQL Server 2012 basic single node edition
        /// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instance_storage`.
        /// </summary>
        [Input("instanceStorage", required: true)]
        public Input<int> InstanceStorage { get; set; } = null!;

        /// <summary>
        /// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
        /// </summary>
        [Input("monitoringPeriod")]
        public Input<int>? MonitoringPeriod { get; set; }

        [Input("parameters")]
        private InputList<Inputs.InstanceParameterArgs>? _parameters;

        /// <summary>
        /// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        public InputList<Inputs.InstanceParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.InstanceParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of resource group which the DB instance belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// It has been deprecated from 1.69.0 and use `security_group_ids` instead.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
        /// </summary>
        [Input("securityIpMode")]
        public Input<string>? SecurityIpMode { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        /// <summary>
        /// The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
        /// </summary>
        [Input("sqlCollectorConfigValue")]
        public Input<int>? SqlCollectorConfigValue { get; set; }

        /// <summary>
        /// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
        /// </summary>
        [Input("sqlCollectorStatus")]
        public Input<string>? SqlCollectorStatus { get; set; }

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
        /// </summary>
        [Input("sslAction")]
        public Input<string>? SslAction { get; set; }

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
        /// The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
        /// </summary>
        [Input("tdeStatus")]
        public Input<string>? TdeStatus { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in the one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The region ID of the secondary instance if you create a secondary instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Input("zoneIdSlaveA")]
        public Input<string>? ZoneIdSlaveA { get; set; }

        /// <summary>
        /// The region ID of the log instance if you create a log instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Input("zoneIdSlaveB")]
        public Input<string>? ZoneIdSlaveB { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is `PrePaid`. Default to `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is `PrePaid`. Valid value:[1~12], Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// The upgrade method to use. Valid values:
        /// - Auto: Instances are automatically upgraded to a higher minor version.
        /// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<string>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// RDS database connection string.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The storage type of the instance. Valid values:
        /// - local_ssd: specifies to use local SSDs. This value is recommended.
        /// - cloud_ssd: specifies to use standard SSDs.
        /// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
        /// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
        /// </summary>
        [Input("dbInstanceStorageType")]
        public Input<string>? DbInstanceStorageType { get; set; }

        /// <summary>
        /// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Set it to true to make some parameter efficient when modifying them. Default to false.
        /// </summary>
        [Input("forceRestart")]
        public Input<bool>? ForceRestart { get; set; }

        /// <summary>
        /// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// User-defined DB instance storage space. Value range:
        /// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
        /// - [20,1000] for MySQL 5.7 basic single node edition;
        /// - [10, 2000] for SQL Server 2008R2;
        /// - [20,2000] for SQL Server 2012 basic single node edition
        /// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instance_storage`.
        /// </summary>
        [Input("instanceStorage")]
        public Input<int>? InstanceStorage { get; set; }

        /// <summary>
        /// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
        /// </summary>
        [Input("monitoringPeriod")]
        public Input<int>? MonitoringPeriod { get; set; }

        [Input("parameters")]
        private InputList<Inputs.InstanceParameterGetArgs>? _parameters;

        /// <summary>
        /// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
        /// </summary>
        public InputList<Inputs.InstanceParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.InstanceParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// RDS database connection port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The ID of resource group which the DB instance belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// It has been deprecated from 1.69.0 and use `security_group_ids` instead.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
        /// </summary>
        [Input("securityIpMode")]
        public Input<string>? SecurityIpMode { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        /// <summary>
        /// The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
        /// </summary>
        [Input("sqlCollectorConfigValue")]
        public Input<int>? SqlCollectorConfigValue { get; set; }

        /// <summary>
        /// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
        /// </summary>
        [Input("sqlCollectorStatus")]
        public Input<string>? SqlCollectorStatus { get; set; }

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
        /// </summary>
        [Input("sslAction")]
        public Input<string>? SslAction { get; set; }

        /// <summary>
        /// Status of the SSL feature. `Yes`: SSL is turned on; `No`: SSL is turned off.
        /// </summary>
        [Input("sslStatus")]
        public Input<string>? SslStatus { get; set; }

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
        /// The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
        /// </summary>
        [Input("tdeStatus")]
        public Input<string>? TdeStatus { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in the one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The region ID of the secondary instance if you create a secondary instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Input("zoneIdSlaveA")]
        public Input<string>? ZoneIdSlaveA { get; set; }

        /// <summary>
        /// The region ID of the log instance if you create a log instance. If you set this parameter to the same value as the ZoneId parameter, the instance is deployed in a single zone. Otherwise, the instance is deployed in multiple zones.
        /// </summary>
        [Input("zoneIdSlaveB")]
        public Input<string>? ZoneIdSlaveB { get; set; }

        public InstanceState()
        {
        }
    }
}
