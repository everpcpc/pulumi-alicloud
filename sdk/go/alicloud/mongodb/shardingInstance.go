// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package mongodb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ShardingInstance struct {
	pulumi.CustomResourceState

	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword pulumi.StringPtrOutput `pulumi:"accountPassword"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayOutput `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringOutput `pulumi:"backupTime"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`. 
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	MongoLists ShardingInstanceMongoListArrayOutput `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntOutput `pulumi:"period"`
	// Instance log backup retention days. Available in 1.42.0+.
	RetentionPeriod pulumi.IntOutput `pulumi:"retentionPeriod"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayOutput `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	ShardLists ShardingInstanceShardListArrayOutput `pulumi:"shardLists"`
	StorageEngine pulumi.StringOutput `pulumi:"storageEngine"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId pulumi.StringPtrOutput `pulumi:"zoneId"`
}

// NewShardingInstance registers a new resource with the given unique name, arguments, and options.
func NewShardingInstance(ctx *pulumi.Context,
	name string, args *ShardingInstanceArgs, opts ...pulumi.ResourceOption) (*ShardingInstance, error) {
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.MongoLists == nil {
		return nil, errors.New("missing required argument 'MongoLists'")
	}
	if args == nil || args.ShardLists == nil {
		return nil, errors.New("missing required argument 'ShardLists'")
	}
	if args == nil {
		args = &ShardingInstanceArgs{}
	}
	var resource ShardingInstance
	err := ctx.RegisterResource("alicloud:mongodb/shardingInstance:ShardingInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShardingInstance gets an existing ShardingInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShardingInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShardingInstanceState, opts ...pulumi.ResourceOption) (*ShardingInstance, error) {
	var resource ShardingInstance
	err := ctx.ReadResource("alicloud:mongodb/shardingInstance:ShardingInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShardingInstance resources.
type shardingInstanceState struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword *string `pulumi:"accountPassword"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods []string `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime *string `pulumi:"backupTime"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`. 
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	EngineVersion *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	MongoLists []ShardingInstanceMongoList `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period *int `pulumi:"period"`
	// Instance log backup retention days. Available in 1.42.0+.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	ShardLists []ShardingInstanceShardList `pulumi:"shardLists"`
	StorageEngine *string `pulumi:"storageEngine"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId *string `pulumi:"zoneId"`
}

type ShardingInstanceState struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword pulumi.StringPtrInput
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayInput
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`. 
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	EngineVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The mongo-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	MongoLists ShardingInstanceMongoListArrayInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntPtrInput
	// Instance log backup retention days. Available in 1.42.0+.
	RetentionPeriod pulumi.IntPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayInput
	// the shard-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	ShardLists ShardingInstanceShardListArrayInput
	StorageEngine pulumi.StringPtrInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId pulumi.StringPtrInput
}

func (ShardingInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingInstanceState)(nil)).Elem()
}

type shardingInstanceArgs struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword *string `pulumi:"accountPassword"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods []string `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime *string `pulumi:"backupTime"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`. 
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	EngineVersion string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	MongoLists []ShardingInstanceMongoList `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period *int `pulumi:"period"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	ShardLists []ShardingInstanceShardList `pulumi:"shardLists"`
	StorageEngine *string `pulumi:"storageEngine"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ShardingInstance resource.
type ShardingInstanceArgs struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword pulumi.StringPtrInput
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayInput
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`. 
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	EngineVersion pulumi.StringInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The mongo-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	MongoLists ShardingInstanceMongoListArrayInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayInput
	// the shard-node count can be purchased is in range of [2, 32].
	// * `nodeClass` -(Required) Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	ShardLists ShardingInstanceShardListArrayInput
	StorageEngine pulumi.StringPtrInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId pulumi.StringPtrInput
}

func (ShardingInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingInstanceArgs)(nil)).Elem()
}

