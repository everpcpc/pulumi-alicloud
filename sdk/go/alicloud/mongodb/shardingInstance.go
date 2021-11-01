// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// MongoDB can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:mongodb/shardingInstance:ShardingInstance example dds-bp1291daeda44195
// ```
type ShardingInstance struct {
	pulumi.CustomResourceState

	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword pulumi.StringPtrOutput `pulumi:"accountPassword"`
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayOutput `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringOutput `pulumi:"backupTime"`
	// The node information list of config server. The details see Block `configServerList`. **NOTE:** Available in v1.140+.
	ConfigServerLists ShardingInstanceConfigServerListArrayOutput `pulumi:"configServerLists"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	MongoLists ShardingInstanceMongoListArrayOutput `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
	// * UPGRADE: The specifications are upgraded.
	// * DOWNGRADE: The specifications are downgraded.
	//   Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
	OrderType pulumi.StringPtrOutput `pulumi:"orderType"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntOutput `pulumi:"period"`
	// Instance log backup retention days. **NOTE:** Available in 1.42.0+.
	RetentionPeriod pulumi.IntOutput `pulumi:"retentionPeriod"`
	// The Security Group ID of ECS.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayOutput `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	ShardLists ShardingInstanceShardListArrayOutput `pulumi:"shardLists"`
	// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	StorageEngine pulumi.StringOutput `pulumi:"storageEngine"`
	Tags          pulumi.MapOutput    `pulumi:"tags"`
	// The TDE(Transparent Data Encryption) status.
	TdeStatus pulumi.StringPtrOutput `pulumi:"tdeStatus"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId pulumi.StringPtrOutput `pulumi:"zoneId"`
}

// NewShardingInstance registers a new resource with the given unique name, arguments, and options.
func NewShardingInstance(ctx *pulumi.Context,
	name string, args *ShardingInstanceArgs, opts ...pulumi.ResourceOption) (*ShardingInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.MongoLists == nil {
		return nil, errors.New("invalid value for required argument 'MongoLists'")
	}
	if args.ShardLists == nil {
		return nil, errors.New("invalid value for required argument 'ShardLists'")
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
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew *bool `pulumi:"autoRenew"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods []string `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime *string `pulumi:"backupTime"`
	// The node information list of config server. The details see Block `configServerList`. **NOTE:** Available in v1.140+.
	ConfigServerLists []ShardingInstanceConfigServerList `pulumi:"configServerLists"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	MongoLists []ShardingInstanceMongoList `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name *string `pulumi:"name"`
	// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
	// * UPGRADE: The specifications are upgraded.
	// * DOWNGRADE: The specifications are downgraded.
	//   Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
	OrderType *string `pulumi:"orderType"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period *int `pulumi:"period"`
	// Instance log backup retention days. **NOTE:** Available in 1.42.0+.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The Security Group ID of ECS.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	ShardLists []ShardingInstanceShardList `pulumi:"shardLists"`
	// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	StorageEngine *string                `pulumi:"storageEngine"`
	Tags          map[string]interface{} `pulumi:"tags"`
	// The TDE(Transparent Data Encryption) status.
	TdeStatus *string `pulumi:"tdeStatus"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId *string `pulumi:"zoneId"`
}

type ShardingInstanceState struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword pulumi.StringPtrInput
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew pulumi.BoolPtrInput
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayInput
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringPtrInput
	// The node information list of config server. The details see Block `configServerList`. **NOTE:** Available in v1.140+.
	ConfigServerLists ShardingInstanceConfigServerListArrayInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`.
	EngineVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
	InstanceChargeType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The mongo-node count can be purchased is in range of [2, 32].
	MongoLists ShardingInstanceMongoListArrayInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringPtrInput
	// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
	// * UPGRADE: The specifications are upgraded.
	// * DOWNGRADE: The specifications are downgraded.
	//   Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
	OrderType pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntPtrInput
	// Instance log backup retention days. **NOTE:** Available in 1.42.0+.
	RetentionPeriod pulumi.IntPtrInput
	// The Security Group ID of ECS.
	SecurityGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayInput
	// the shard-node count can be purchased is in range of [2, 32].
	ShardLists ShardingInstanceShardListArrayInput
	// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	StorageEngine pulumi.StringPtrInput
	Tags          pulumi.MapInput
	// The TDE(Transparent Data Encryption) status.
	TdeStatus pulumi.StringPtrInput
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
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew *bool `pulumi:"autoRenew"`
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods []string `pulumi:"backupPeriods"`
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime *string `pulumi:"backupTime"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`.
	EngineVersion string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The mongo-node count can be purchased is in range of [2, 32].
	MongoLists []ShardingInstanceMongoList `pulumi:"mongoLists"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Name *string `pulumi:"name"`
	// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
	// * UPGRADE: The specifications are upgraded.
	// * DOWNGRADE: The specifications are downgraded.
	//   Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
	OrderType *string `pulumi:"orderType"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period *int `pulumi:"period"`
	// The Security Group ID of ECS.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// the shard-node count can be purchased is in range of [2, 32].
	ShardLists []ShardingInstanceShardList `pulumi:"shardLists"`
	// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	StorageEngine *string                `pulumi:"storageEngine"`
	Tags          map[string]interface{} `pulumi:"tags"`
	// The TDE(Transparent Data Encryption) status.
	TdeStatus *string `pulumi:"tdeStatus"`
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
	// Auto renew for prepaid, true of false. Default is false.
	AutoRenew pulumi.BoolPtrInput
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods pulumi.StringArrayInput
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
	BackupTime pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) `EngineVersion`.
	EngineVersion pulumi.StringInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
	InstanceChargeType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The mongo-node count can be purchased is in range of [2, 32].
	MongoLists ShardingInstanceMongoListArrayInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Name pulumi.StringPtrInput
	// The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
	// * UPGRADE: The specifications are upgraded.
	// * DOWNGRADE: The specifications are downgraded.
	//   Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
	OrderType pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period pulumi.IntPtrInput
	// The Security Group ID of ECS.
	SecurityGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
	SecurityIpLists pulumi.StringArrayInput
	// the shard-node count can be purchased is in range of [2, 32].
	ShardLists ShardingInstanceShardListArrayInput
	// Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	StorageEngine pulumi.StringPtrInput
	Tags          pulumi.MapInput
	// The TDE(Transparent Data Encryption) status.
	TdeStatus pulumi.StringPtrInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	ZoneId pulumi.StringPtrInput
}

func (ShardingInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shardingInstanceArgs)(nil)).Elem()
}

type ShardingInstanceInput interface {
	pulumi.Input

	ToShardingInstanceOutput() ShardingInstanceOutput
	ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput
}

func (*ShardingInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ShardingInstance)(nil))
}

func (i *ShardingInstance) ToShardingInstanceOutput() ShardingInstanceOutput {
	return i.ToShardingInstanceOutputWithContext(context.Background())
}

func (i *ShardingInstance) ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceOutput)
}

func (i *ShardingInstance) ToShardingInstancePtrOutput() ShardingInstancePtrOutput {
	return i.ToShardingInstancePtrOutputWithContext(context.Background())
}

func (i *ShardingInstance) ToShardingInstancePtrOutputWithContext(ctx context.Context) ShardingInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstancePtrOutput)
}

type ShardingInstancePtrInput interface {
	pulumi.Input

	ToShardingInstancePtrOutput() ShardingInstancePtrOutput
	ToShardingInstancePtrOutputWithContext(ctx context.Context) ShardingInstancePtrOutput
}

type shardingInstancePtrType ShardingInstanceArgs

func (*shardingInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingInstance)(nil))
}

func (i *shardingInstancePtrType) ToShardingInstancePtrOutput() ShardingInstancePtrOutput {
	return i.ToShardingInstancePtrOutputWithContext(context.Background())
}

func (i *shardingInstancePtrType) ToShardingInstancePtrOutputWithContext(ctx context.Context) ShardingInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstancePtrOutput)
}

// ShardingInstanceArrayInput is an input type that accepts ShardingInstanceArray and ShardingInstanceArrayOutput values.
// You can construct a concrete instance of `ShardingInstanceArrayInput` via:
//
//          ShardingInstanceArray{ ShardingInstanceArgs{...} }
type ShardingInstanceArrayInput interface {
	pulumi.Input

	ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput
	ToShardingInstanceArrayOutputWithContext(context.Context) ShardingInstanceArrayOutput
}

type ShardingInstanceArray []ShardingInstanceInput

func (ShardingInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ShardingInstance)(nil))
}

func (i ShardingInstanceArray) ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput {
	return i.ToShardingInstanceArrayOutputWithContext(context.Background())
}

func (i ShardingInstanceArray) ToShardingInstanceArrayOutputWithContext(ctx context.Context) ShardingInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceArrayOutput)
}

// ShardingInstanceMapInput is an input type that accepts ShardingInstanceMap and ShardingInstanceMapOutput values.
// You can construct a concrete instance of `ShardingInstanceMapInput` via:
//
//          ShardingInstanceMap{ "key": ShardingInstanceArgs{...} }
type ShardingInstanceMapInput interface {
	pulumi.Input

	ToShardingInstanceMapOutput() ShardingInstanceMapOutput
	ToShardingInstanceMapOutputWithContext(context.Context) ShardingInstanceMapOutput
}

type ShardingInstanceMap map[string]ShardingInstanceInput

func (ShardingInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ShardingInstance)(nil))
}

func (i ShardingInstanceMap) ToShardingInstanceMapOutput() ShardingInstanceMapOutput {
	return i.ToShardingInstanceMapOutputWithContext(context.Background())
}

func (i ShardingInstanceMap) ToShardingInstanceMapOutputWithContext(ctx context.Context) ShardingInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShardingInstanceMapOutput)
}

type ShardingInstanceOutput struct {
	*pulumi.OutputState
}

func (ShardingInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShardingInstance)(nil))
}

func (o ShardingInstanceOutput) ToShardingInstanceOutput() ShardingInstanceOutput {
	return o
}

func (o ShardingInstanceOutput) ToShardingInstanceOutputWithContext(ctx context.Context) ShardingInstanceOutput {
	return o
}

func (o ShardingInstanceOutput) ToShardingInstancePtrOutput() ShardingInstancePtrOutput {
	return o.ToShardingInstancePtrOutputWithContext(context.Background())
}

func (o ShardingInstanceOutput) ToShardingInstancePtrOutputWithContext(ctx context.Context) ShardingInstancePtrOutput {
	return o.ApplyT(func(v ShardingInstance) *ShardingInstance {
		return &v
	}).(ShardingInstancePtrOutput)
}

type ShardingInstancePtrOutput struct {
	*pulumi.OutputState
}

func (ShardingInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShardingInstance)(nil))
}

func (o ShardingInstancePtrOutput) ToShardingInstancePtrOutput() ShardingInstancePtrOutput {
	return o
}

func (o ShardingInstancePtrOutput) ToShardingInstancePtrOutputWithContext(ctx context.Context) ShardingInstancePtrOutput {
	return o
}

type ShardingInstanceArrayOutput struct{ *pulumi.OutputState }

func (ShardingInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShardingInstance)(nil))
}

func (o ShardingInstanceArrayOutput) ToShardingInstanceArrayOutput() ShardingInstanceArrayOutput {
	return o
}

func (o ShardingInstanceArrayOutput) ToShardingInstanceArrayOutputWithContext(ctx context.Context) ShardingInstanceArrayOutput {
	return o
}

func (o ShardingInstanceArrayOutput) Index(i pulumi.IntInput) ShardingInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShardingInstance {
		return vs[0].([]ShardingInstance)[vs[1].(int)]
	}).(ShardingInstanceOutput)
}

type ShardingInstanceMapOutput struct{ *pulumi.OutputState }

func (ShardingInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ShardingInstance)(nil))
}

func (o ShardingInstanceMapOutput) ToShardingInstanceMapOutput() ShardingInstanceMapOutput {
	return o
}

func (o ShardingInstanceMapOutput) ToShardingInstanceMapOutputWithContext(ctx context.Context) ShardingInstanceMapOutput {
	return o
}

func (o ShardingInstanceMapOutput) MapIndex(k pulumi.StringInput) ShardingInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ShardingInstance {
		return vs[0].(map[string]ShardingInstance)[vs[1].(string)]
	}).(ShardingInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(ShardingInstanceOutput{})
	pulumi.RegisterOutputType(ShardingInstancePtrOutput{})
	pulumi.RegisterOutputType(ShardingInstanceArrayOutput{})
	pulumi.RegisterOutputType(ShardingInstanceMapOutput{})
}
