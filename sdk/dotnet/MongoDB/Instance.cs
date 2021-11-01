// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB
{
    /// <summary>
    /// ## Import
    /// 
    /// MongoDB can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:mongodb/instance:Instance example dds-bp1291daeda44194
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:mongodb/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
        /// </summary>
        [Output("accountPassword")]
        public Output<string?> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// Auto renew for prepaid, true of false. Default is false.
        /// &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
        /// </summary>
        [Output("backupPeriods")]
        public Output<ImmutableArray<string>> BackupPeriods { get; private set; } = null!;

        /// <summary>
        /// MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
        /// </summary>
        [Output("backupTime")]
        public Output<string> BackupTime { get; private set; } = null!;

        /// <summary>
        /// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        /// </summary>
        [Output("dbInstanceClass")]
        public Output<string> DbInstanceClass { get; private set; } = null!;

        /// <summary>
        /// User-defined DB instance storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [10,2000]
        /// - 10-GB increments.
        /// </summary>
        [Output("dbInstanceStorage")]
        public Output<int> DbInstanceStorage { get; private set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Output("maintainEndTime")]
        public Output<string> MaintainEndTime { get; private set; } = null!;

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Output("maintainStartTime")]
        public Output<string> MaintainStartTime { get; private set; } = null!;

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
        /// * UPGRADE: The specifications are upgraded.
        /// * DOWNGRADE: The specifications are downgraded.
        /// Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
        /// </summary>
        [Output("orderType")]
        public Output<string?> OrderType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// The name of the mongo replica set
        /// </summary>
        [Output("replicaSetName")]
        public Output<string> ReplicaSetName { get; private set; } = null!;

        /// <summary>
        /// Replica set instance information. The details see Block replica_sets. **NOTE:** Available in v1.140+.
        /// </summary>
        [Output("replicaSets")]
        public Output<ImmutableArray<Outputs.InstanceReplicaSet>> ReplicaSets { get; private set; } = null!;

        /// <summary>
        /// Number of replica set nodes. Valid values: [1, 3, 5, 7]
        /// </summary>
        [Output("replicationFactor")]
        public Output<int> ReplicationFactor { get; private set; } = null!;

        /// <summary>
        /// Instance log backup retention days. Available in 1.42.0+.
        /// </summary>
        [Output("retentionPeriod")]
        public Output<int> RetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The Security Group ID of ECS.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        [Output("securityIpLists")]
        public Output<ImmutableArray<string>> SecurityIpLists { get; private set; } = null!;

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
        /// </summary>
        [Output("sslAction")]
        public Output<string> SslAction { get; private set; } = null!;

        /// <summary>
        /// Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
        /// </summary>
        [Output("sslStatus")]
        public Output<string> SslStatus { get; private set; } = null!;

        /// <summary>
        /// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
        /// </summary>
        [Output("storageEngine")]
        public Output<string> StorageEngine { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The TDE(Transparent Data Encryption) status.
        /// </summary>
        [Output("tdeStatus")]
        public Output<string?> TdeStatus { get; private set; } = null!;

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the DB instance. it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mongodb/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
        /// </summary>
        [Input("accountPassword")]
        public Input<string>? AccountPassword { get; set; }

        /// <summary>
        /// Auto renew for prepaid, true of false. Default is false.
        /// &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
        /// </summary>
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        /// </summary>
        [Input("dbInstanceClass", required: true)]
        public Input<string> DbInstanceClass { get; set; } = null!;

        /// <summary>
        /// User-defined DB instance storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [10,2000]
        /// - 10-GB increments.
        /// </summary>
        [Input("dbInstanceStorage", required: true)]
        public Input<int> DbInstanceStorage { get; set; } = null!;

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
        /// * UPGRADE: The specifications are upgraded.
        /// * DOWNGRADE: The specifications are downgraded.
        /// Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Number of replica set nodes. Valid values: [1, 3, 5, 7]
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        /// <summary>
        /// The Security Group ID of ECS.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("securityIpLists")]
        private InputList<string>? _securityIpLists;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIpLists
        {
            get => _securityIpLists ?? (_securityIpLists = new InputList<string>());
            set => _securityIpLists = value;
        }

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
        /// </summary>
        [Input("sslAction")]
        public Input<string>? SslAction { get; set; }

        /// <summary>
        /// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
        /// </summary>
        [Input("storageEngine")]
        public Input<string>? StorageEngine { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The TDE(Transparent Data Encryption) status.
        /// </summary>
        [Input("tdeStatus")]
        public Input<string>? TdeStatus { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB instance. it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
        /// </summary>
        [Input("accountPassword")]
        public Input<string>? AccountPassword { get; set; }

        /// <summary>
        /// Auto renew for prepaid, true of false. Default is false.
        /// &gt; **NOTE:** The start time to the end time must be 1 hour. For example, the MaintainStartTime is 01:00Z, then the MaintainEndTime must be 02:00Z.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// MongoDB Instance backup period. It is required when `backup_time` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
        /// </summary>
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// MongoDB instance backup time. It is required when `backup_period` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
        /// </summary>
        [Input("dbInstanceClass")]
        public Input<string>? DbInstanceClass { get; set; }

        /// <summary>
        /// User-defined DB instance storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [10,2000]
        /// - 10-GB increments.
        /// </summary>
        [Input("dbInstanceStorage")]
        public Input<int>? DbInstanceStorage { get; set; }

        /// <summary>
        /// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a instance. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating instance with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
        /// * UPGRADE: The specifications are upgraded.
        /// * DOWNGRADE: The specifications are downgraded.
        /// Note: This parameter is only applicable to instances when `instance_charge_type` is PrePaid.
        /// </summary>
        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        /// <summary>
        /// The duration that you will buy DB instance (in month). It is valid when instance_charge_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The name of the mongo replica set
        /// </summary>
        [Input("replicaSetName")]
        public Input<string>? ReplicaSetName { get; set; }

        [Input("replicaSets")]
        private InputList<Inputs.InstanceReplicaSetGetArgs>? _replicaSets;

        /// <summary>
        /// Replica set instance information. The details see Block replica_sets. **NOTE:** Available in v1.140+.
        /// </summary>
        public InputList<Inputs.InstanceReplicaSetGetArgs> ReplicaSets
        {
            get => _replicaSets ?? (_replicaSets = new InputList<Inputs.InstanceReplicaSetGetArgs>());
            set => _replicaSets = value;
        }

        /// <summary>
        /// Number of replica set nodes. Valid values: [1, 3, 5, 7]
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        /// <summary>
        /// Instance log backup retention days. Available in 1.42.0+.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        /// <summary>
        /// The Security Group ID of ECS.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("securityIpLists")]
        private InputList<string>? _securityIpLists;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIpLists
        {
            get => _securityIpLists ?? (_securityIpLists = new InputList<string>());
            set => _securityIpLists = value;
        }

        /// <summary>
        /// Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate.
        /// </summary>
        [Input("sslAction")]
        public Input<string>? SslAction { get; set; }

        /// <summary>
        /// Status of the SSL feature. `Open`: SSL is turned on; `Closed`: SSL is turned off.
        /// </summary>
        [Input("sslStatus")]
        public Input<string>? SslStatus { get; set; }

        /// <summary>
        /// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
        /// </summary>
        [Input("storageEngine")]
        public Input<string>? StorageEngine { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The TDE(Transparent Data Encryption) status.
        /// </summary>
        [Input("tdeStatus")]
        public Input<string>? TdeStatus { get; set; }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB instance. it supports multiple zone.
        /// If it is a multi-zone and `vswitch_id` is specified, the vswitch must in one of them.
        /// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud.getZones`.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
    }
}
