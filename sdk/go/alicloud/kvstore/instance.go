// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ApsaraDB Redis / Memcache instance resource. A DB instance is an isolated database environment in the cloud. It can be associated with IP whitelists and backup configuration which are separate resource providers.
type Instance struct {
	pulumi.CustomResourceState

	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrOutput `pulumi:"autoRenewPeriod"`
	// The Zone to launch the DB instance.
	AvailabilityZone pulumi.StringOutput    `pulumi:"availabilityZone"`
	BackupId         pulumi.StringPtrOutput `pulumi:"backupId"`
	// Instance connection domain (only Intranet access supported).
	ConnectionDomain pulumi.StringOutput    `pulumi:"connectionDomain"`
	EngineVersion    pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	InstanceClass      pulumi.StringOutput    `pulumi:"instanceClass"`
	// The name of DB instance. It a string of 2 to 256 characters.
	// * `password`- (Optional, Sensitive) The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime pulumi.StringOutput `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime pulumi.StringOutput `pulumi:"maintainStartTime"`
	// Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
	Parameters InstanceParameterArrayOutput `pulumi:"parameters"`
	Password   pulumi.StringPtrOutput       `pulumi:"password"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period    pulumi.IntPtrOutput `pulumi:"period"`
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The Security Group ID of ECS.
	// * `privateIp`- (Optional) Set the instance's private IP.
	// * `backupId`- (Optional) If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
	// * `vpcAuthMode`- (Optional) Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
	SecurityGroupId pulumi.StringOutput      `pulumi:"securityGroupId"`
	SecurityIps     pulumi.StringArrayOutput `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	Tags        pulumi.MapOutput    `pulumi:"tags"`
	VpcAuthMode pulumi.StringOutput `pulumi:"vpcAuthMode"`
	// The ID of VSwitch.
	// * `engineVersion`- (Optional, ForceNew) Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
	// * `securityIps`- (Optional) Set the instance's IP whitelist of the default security group.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.InstanceClass == nil {
		return nil, errors.New("missing required argument 'InstanceClass'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:kvstore/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:kvstore/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The Zone to launch the DB instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	BackupId         *string `pulumi:"backupId"`
	// Instance connection domain (only Intranet access supported).
	ConnectionDomain *string `pulumi:"connectionDomain"`
	EngineVersion    *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	InstanceClass      *string `pulumi:"instanceClass"`
	// The name of DB instance. It a string of 2 to 256 characters.
	// * `password`- (Optional, Sensitive) The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	InstanceName *string `pulumi:"instanceName"`
	// The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
	InstanceType *string `pulumi:"instanceType"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime *string `pulumi:"maintainStartTime"`
	// Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
	Parameters []InstanceParameter `pulumi:"parameters"`
	Password   *string             `pulumi:"password"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period    *int    `pulumi:"period"`
	PrivateIp *string `pulumi:"privateIp"`
	// The Security Group ID of ECS.
	// * `privateIp`- (Optional) Set the instance's private IP.
	// * `backupId`- (Optional) If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
	// * `vpcAuthMode`- (Optional) Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
	SecurityGroupId *string  `pulumi:"securityGroupId"`
	SecurityIps     []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	Tags        map[string]interface{} `pulumi:"tags"`
	VpcAuthMode *string                `pulumi:"vpcAuthMode"`
	// The ID of VSwitch.
	// * `engineVersion`- (Optional, ForceNew) Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
	// * `securityIps`- (Optional) Set the instance's IP whitelist of the default security group.
	VswitchId *string `pulumi:"vswitchId"`
}

type InstanceState struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrInput
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// The Zone to launch the DB instance.
	AvailabilityZone pulumi.StringPtrInput
	BackupId         pulumi.StringPtrInput
	// Instance connection domain (only Intranet access supported).
	ConnectionDomain pulumi.StringPtrInput
	EngineVersion    pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	InstanceClass      pulumi.StringPtrInput
	// The name of DB instance. It a string of 2 to 256 characters.
	// * `password`- (Optional, Sensitive) The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	InstanceName pulumi.StringPtrInput
	// The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
	InstanceType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime pulumi.StringPtrInput
	// Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
	Parameters InstanceParameterArrayInput
	Password   pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period    pulumi.IntPtrInput
	PrivateIp pulumi.StringPtrInput
	// The Security Group ID of ECS.
	// * `privateIp`- (Optional) Set the instance's private IP.
	// * `backupId`- (Optional) If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
	// * `vpcAuthMode`- (Optional) Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
	SecurityGroupId pulumi.StringPtrInput
	SecurityIps     pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags        pulumi.MapInput
	VpcAuthMode pulumi.StringPtrInput
	// The ID of VSwitch.
	// * `engineVersion`- (Optional, ForceNew) Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
	// * `securityIps`- (Optional) Set the instance's IP whitelist of the default security group.
	VswitchId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The Zone to launch the DB instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	BackupId         *string `pulumi:"backupId"`
	EngineVersion    *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	InstanceClass      string  `pulumi:"instanceClass"`
	// The name of DB instance. It a string of 2 to 256 characters.
	// * `password`- (Optional, Sensitive) The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	InstanceName *string `pulumi:"instanceName"`
	// The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
	InstanceType *string `pulumi:"instanceType"`
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime *string `pulumi:"maintainStartTime"`
	// Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
	Parameters []InstanceParameter `pulumi:"parameters"`
	Password   *string             `pulumi:"password"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period    *int    `pulumi:"period"`
	PrivateIp *string `pulumi:"privateIp"`
	// The Security Group ID of ECS.
	// * `privateIp`- (Optional) Set the instance's private IP.
	// * `backupId`- (Optional) If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
	// * `vpcAuthMode`- (Optional) Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
	SecurityGroupId *string  `pulumi:"securityGroupId"`
	SecurityIps     []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	Tags        map[string]interface{} `pulumi:"tags"`
	VpcAuthMode *string                `pulumi:"vpcAuthMode"`
	// The ID of VSwitch.
	// * `engineVersion`- (Optional, ForceNew) Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
	// * `securityIps`- (Optional) Set the instance's IP whitelist of the default security group.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrInput
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// The Zone to launch the DB instance.
	AvailabilityZone pulumi.StringPtrInput
	BackupId         pulumi.StringPtrInput
	EngineVersion    pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	InstanceClass      pulumi.StringInput
	// The name of DB instance. It a string of 2 to 256 characters.
	// * `password`- (Optional, Sensitive) The password of the DB instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
	InstanceName pulumi.StringPtrInput
	// The engine to use: `Redis` or `Memcache`. Defaults to `Redis`.
	InstanceType pulumi.StringPtrInput
	// An KMS encrypts password used to a instance. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime pulumi.StringPtrInput
	// Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .
	Parameters InstanceParameterArrayInput
	Password   pulumi.StringPtrInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period    pulumi.IntPtrInput
	PrivateIp pulumi.StringPtrInput
	// The Security Group ID of ECS.
	// * `privateIp`- (Optional) Set the instance's private IP.
	// * `backupId`- (Optional) If an instance created based on a backup set generated by another instance is valid, this parameter indicates the ID of the generated backup set.
	// * `vpcAuthMode`- (Optional) Only meaningful if instanceType is `Redis` and network type is VPC. Valid values are `Close`, `Open`. Defaults to `Open`.  `Close` means the redis instance can be accessed without authentication. `Open` means authentication is required.
	SecurityGroupId pulumi.StringPtrInput
	SecurityIps     pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags        pulumi.MapInput
	VpcAuthMode pulumi.StringPtrInput
	// The ID of VSwitch.
	// * `engineVersion`- (Optional, ForceNew) Engine version. Supported values: 2.8, 4.0 and 5.0. Default value: 2.8. Only 2.8 can be supported for Memcache Instance.
	// * `securityIps`- (Optional) Set the instance's IP whitelist of the default security group.
	VswitchId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
