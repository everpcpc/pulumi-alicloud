// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// ESS scaling configuration can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ess/scalingConfiguration:ScalingConfiguration example asg-abc123456
// ```
type ScalingConfiguration struct {
	pulumi.CustomResourceState

	// Whether active current scaling configuration in the specified scaling group. Default to `false`.
	Active pulumi.BoolOutput `pulumi:"active"`
	// Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.
	CreditSpecification pulumi.StringPtrOutput `pulumi:"creditSpecification"`
	// DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
	DataDisks ScalingConfigurationDataDiskArrayOutput `pulumi:"dataDisks"`
	// Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// ID of an image file, indicating the image resource selected when an instance is enabled.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Name of an image file, indicating the image resource selected when an instance is enabled.
	ImageName pulumi.StringPtrOutput `pulumi:"imageName"`
	// It has been deprecated from version 1.6.0. New resource `ess.Attachment` replaces it.
	//
	// Deprecated: Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// Resource type of an ECS instance.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// Resource types of an ECS instance.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
	InternetMaxBandwidthIn pulumi.IntOutput `pulumi:"internetMaxBandwidthIn"`
	// Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
	InternetMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"internetMaxBandwidthOut"`
	// It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
	//
	// Deprecated: Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
	IoOptimized pulumi.StringPtrOutput `pulumi:"ioOptimized"`
	// Whether to use outdated instance type. Default to false.
	IsOutdated pulumi.BoolPtrOutput `pulumi:"isOutdated"`
	// The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
	KeyName pulumi.StringPtrOutput `pulumi:"keyName"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// Indicates whether to overwrite the existing data. Default to false.
	Override pulumi.BoolPtrOutput `pulumi:"override"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
	PasswordInherit pulumi.BoolPtrOutput `pulumi:"passwordInherit"`
	// Instance RAM role name. The name is provided and maintained by RAM. You can use `ram.Role` to create a new one.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
	ScalingConfigurationName pulumi.StringOutput `pulumi:"scalingConfigurationName"`
	// ID of the scaling group of a scaling configuration.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
	Substitute pulumi.StringOutput `pulumi:"substitute"`
	// The id of auto snapshot policy for system disk.
	SystemDiskAutoSnapshotPolicyId pulumi.StringPtrOutput `pulumi:"systemDiskAutoSnapshotPolicyId"`
	// Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
	SystemDiskCategory pulumi.StringPtrOutput `pulumi:"systemDiskCategory"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	SystemDiskDescription pulumi.StringPtrOutput `pulumi:"systemDiskDescription"`
	// The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
	SystemDiskName pulumi.StringPtrOutput `pulumi:"systemDiskName"`
	// Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
	SystemDiskSize pulumi.IntPtrOutput `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
}

// NewScalingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewScalingConfiguration(ctx *pulumi.Context,
	name string, args *ScalingConfigurationArgs, opts ...pulumi.ResourceOption) (*ScalingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	var resource ScalingConfiguration
	err := ctx.RegisterResource("alicloud:ess/scalingConfiguration:ScalingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingConfiguration gets an existing ScalingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingConfigurationState, opts ...pulumi.ResourceOption) (*ScalingConfiguration, error) {
	var resource ScalingConfiguration
	err := ctx.ReadResource("alicloud:ess/scalingConfiguration:ScalingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingConfiguration resources.
type scalingConfigurationState struct {
	// Whether active current scaling configuration in the specified scaling group. Default to `false`.
	Active *bool `pulumi:"active"`
	// Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.
	CreditSpecification *string `pulumi:"creditSpecification"`
	// DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
	DataDisks []ScalingConfigurationDataDisk `pulumi:"dataDisks"`
	// Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
	Enable *bool `pulumi:"enable"`
	// The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
	ForceDelete *bool `pulumi:"forceDelete"`
	// ID of an image file, indicating the image resource selected when an instance is enabled.
	ImageId *string `pulumi:"imageId"`
	// Name of an image file, indicating the image resource selected when an instance is enabled.
	ImageName *string `pulumi:"imageName"`
	// It has been deprecated from version 1.6.0. New resource `ess.Attachment` replaces it.
	//
	// Deprecated: Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
	InstanceIds []string `pulumi:"instanceIds"`
	// Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
	InstanceName *string `pulumi:"instanceName"`
	// Resource type of an ECS instance.
	InstanceType *string `pulumi:"instanceType"`
	// Resource types of an ECS instance.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
	InternetMaxBandwidthIn *int `pulumi:"internetMaxBandwidthIn"`
	// Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
	//
	// Deprecated: Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
	IoOptimized *string `pulumi:"ioOptimized"`
	// Whether to use outdated instance type. Default to false.
	IsOutdated *bool `pulumi:"isOutdated"`
	// The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Indicates whether to overwrite the existing data. Default to false.
	Override *bool `pulumi:"override"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
	Password *string `pulumi:"password"`
	// Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
	PasswordInherit *bool `pulumi:"passwordInherit"`
	// Instance RAM role name. The name is provided and maintained by RAM. You can use `ram.Role` to create a new one.
	RoleName *string `pulumi:"roleName"`
	// Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
	ScalingConfigurationName *string `pulumi:"scalingConfigurationName"`
	// ID of the scaling group of a scaling configuration.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
	Substitute *string `pulumi:"substitute"`
	// The id of auto snapshot policy for system disk.
	SystemDiskAutoSnapshotPolicyId *string `pulumi:"systemDiskAutoSnapshotPolicyId"`
	// Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
	SystemDiskCategory *string `pulumi:"systemDiskCategory"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	SystemDiskDescription *string `pulumi:"systemDiskDescription"`
	// The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
	SystemDiskName *string `pulumi:"systemDiskName"`
	// Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
	UserData *string `pulumi:"userData"`
}

type ScalingConfigurationState struct {
	// Whether active current scaling configuration in the specified scaling group. Default to `false`.
	Active pulumi.BoolPtrInput
	// Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.
	CreditSpecification pulumi.StringPtrInput
	// DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
	DataDisks ScalingConfigurationDataDiskArrayInput
	// Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
	Enable pulumi.BoolPtrInput
	// The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
	ForceDelete pulumi.BoolPtrInput
	// ID of an image file, indicating the image resource selected when an instance is enabled.
	ImageId pulumi.StringPtrInput
	// Name of an image file, indicating the image resource selected when an instance is enabled.
	ImageName pulumi.StringPtrInput
	// It has been deprecated from version 1.6.0. New resource `ess.Attachment` replaces it.
	//
	// Deprecated: Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
	InstanceIds pulumi.StringArrayInput
	// Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
	InstanceName pulumi.StringPtrInput
	// Resource type of an ECS instance.
	InstanceType pulumi.StringPtrInput
	// Resource types of an ECS instance.
	InstanceTypes pulumi.StringArrayInput
	// Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
	InternetMaxBandwidthIn pulumi.IntPtrInput
	// Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
	//
	// Deprecated: Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
	IoOptimized pulumi.StringPtrInput
	// Whether to use outdated instance type. Default to false.
	IsOutdated pulumi.BoolPtrInput
	// The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Indicates whether to overwrite the existing data. Default to false.
	Override pulumi.BoolPtrInput
	// The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
	Password pulumi.StringPtrInput
	// Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
	PasswordInherit pulumi.BoolPtrInput
	// Instance RAM role name. The name is provided and maintained by RAM. You can use `ram.Role` to create a new one.
	RoleName pulumi.StringPtrInput
	// Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
	ScalingConfigurationName pulumi.StringPtrInput
	// ID of the scaling group of a scaling configuration.
	ScalingGroupId pulumi.StringPtrInput
	// ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
	SecurityGroupId pulumi.StringPtrInput
	// List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
	SecurityGroupIds pulumi.StringArrayInput
	// The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
	Substitute pulumi.StringPtrInput
	// The id of auto snapshot policy for system disk.
	SystemDiskAutoSnapshotPolicyId pulumi.StringPtrInput
	// Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
	SystemDiskCategory pulumi.StringPtrInput
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	SystemDiskDescription pulumi.StringPtrInput
	// The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
	SystemDiskName pulumi.StringPtrInput
	// Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
	SystemDiskSize pulumi.IntPtrInput
	// A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
	Tags pulumi.MapInput
	// User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
	UserData pulumi.StringPtrInput
}

func (ScalingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationState)(nil)).Elem()
}

type scalingConfigurationArgs struct {
	// Whether active current scaling configuration in the specified scaling group. Default to `false`.
	Active *bool `pulumi:"active"`
	// Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.
	CreditSpecification *string `pulumi:"creditSpecification"`
	// DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
	DataDisks []ScalingConfigurationDataDisk `pulumi:"dataDisks"`
	// Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
	Enable *bool `pulumi:"enable"`
	// The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
	ForceDelete *bool `pulumi:"forceDelete"`
	// ID of an image file, indicating the image resource selected when an instance is enabled.
	ImageId *string `pulumi:"imageId"`
	// Name of an image file, indicating the image resource selected when an instance is enabled.
	ImageName *string `pulumi:"imageName"`
	// It has been deprecated from version 1.6.0. New resource `ess.Attachment` replaces it.
	//
	// Deprecated: Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
	InstanceIds []string `pulumi:"instanceIds"`
	// Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
	InstanceName *string `pulumi:"instanceName"`
	// Resource type of an ECS instance.
	InstanceType *string `pulumi:"instanceType"`
	// Resource types of an ECS instance.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
	InternetMaxBandwidthIn *int `pulumi:"internetMaxBandwidthIn"`
	// Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
	//
	// Deprecated: Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
	IoOptimized *string `pulumi:"ioOptimized"`
	// Whether to use outdated instance type. Default to false.
	IsOutdated *bool `pulumi:"isOutdated"`
	// The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
	KeyName *string `pulumi:"keyName"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Indicates whether to overwrite the existing data. Default to false.
	Override *bool `pulumi:"override"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
	Password *string `pulumi:"password"`
	// Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
	PasswordInherit *bool `pulumi:"passwordInherit"`
	// Instance RAM role name. The name is provided and maintained by RAM. You can use `ram.Role` to create a new one.
	RoleName *string `pulumi:"roleName"`
	// Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
	ScalingConfigurationName *string `pulumi:"scalingConfigurationName"`
	// ID of the scaling group of a scaling configuration.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
	Substitute *string `pulumi:"substitute"`
	// The id of auto snapshot policy for system disk.
	SystemDiskAutoSnapshotPolicyId *string `pulumi:"systemDiskAutoSnapshotPolicyId"`
	// Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
	SystemDiskCategory *string `pulumi:"systemDiskCategory"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	SystemDiskDescription *string `pulumi:"systemDiskDescription"`
	// The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
	SystemDiskName *string `pulumi:"systemDiskName"`
	// Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
	UserData *string `pulumi:"userData"`
}

// The set of arguments for constructing a ScalingConfiguration resource.
type ScalingConfigurationArgs struct {
	// Whether active current scaling configuration in the specified scaling group. Default to `false`.
	Active pulumi.BoolPtrInput
	// Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.
	CreditSpecification pulumi.StringPtrInput
	// DataDisk mappings to attach to ecs instance. See Block datadisk below for details.
	DataDisks ScalingConfigurationDataDiskArrayInput
	// Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.
	Enable pulumi.BoolPtrInput
	// The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.
	ForceDelete pulumi.BoolPtrInput
	// ID of an image file, indicating the image resource selected when an instance is enabled.
	ImageId pulumi.StringPtrInput
	// Name of an image file, indicating the image resource selected when an instance is enabled.
	ImageName pulumi.StringPtrInput
	// It has been deprecated from version 1.6.0. New resource `ess.Attachment` replaces it.
	//
	// Deprecated: Field 'instance_ids' has been deprecated from provider version 1.6.0. New resource 'alicloud_ess_attachment' replaces it.
	InstanceIds pulumi.StringArrayInput
	// Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.
	InstanceName pulumi.StringPtrInput
	// Resource type of an ECS instance.
	InstanceType pulumi.StringPtrInput
	// Resource types of an ECS instance.
	InstanceTypes pulumi.StringArrayInput
	// Network billing type, Values: PayByBandwidth or PayByTraffic. Default to `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].
	InternetMaxBandwidthIn pulumi.IntPtrInput
	// Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.
	//
	// Deprecated: Attribute io_optimized has been deprecated on instance resource. All the launched alicloud instances will be IO optimized. Suggest to remove it from your template.
	IoOptimized pulumi.StringPtrInput
	// Whether to use outdated instance type. Default to false.
	IsOutdated pulumi.BoolPtrInput
	// The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.
	KeyName pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Indicates whether to overwrite the existing data. Default to false.
	Override pulumi.BoolPtrInput
	// The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `() ~!@#$%^&*-_+=\|{}[]:;'<>,.?/`, The password of Windows-based instances cannot start with a forward slash (/).
	Password pulumi.StringPtrInput
	// Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the `password` and `kmsEncryptedPassword` will be ignored. You must ensure that the selected image has a password configured.
	PasswordInherit pulumi.BoolPtrInput
	// Instance RAM role name. The name is provided and maintained by RAM. You can use `ram.Role` to create a new one.
	RoleName pulumi.StringPtrInput
	// Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is ScalingConfigurationId.
	ScalingConfigurationName pulumi.StringPtrInput
	// ID of the scaling group of a scaling configuration.
	ScalingGroupId pulumi.StringInput
	// ID of the security group used to create new instance. It is conflict with `securityGroupIds`.
	SecurityGroupId pulumi.StringPtrInput
	// List IDs of the security group used to create new instances. It is conflict with `securityGroupId`.
	SecurityGroupIds pulumi.StringArrayInput
	// The another scaling configuration which will be active automatically and replace current configuration when setting `active` to 'false'. It is invalid when `active` is 'true'.
	Substitute pulumi.StringPtrInput
	// The id of auto snapshot policy for system disk.
	SystemDiskAutoSnapshotPolicyId pulumi.StringPtrInput
	// Category of the system disk. The parameter value options are `ephemeralSsd`, `cloudEfficiency`, `cloudSsd`, `cloudEssd` and `cloud`. `cloud` only is used to some no I/O optimized instance. Default to `cloudEfficiency`.
	SystemDiskCategory pulumi.StringPtrInput
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	SystemDiskDescription pulumi.StringPtrInput
	// The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.
	SystemDiskName pulumi.StringPtrInput
	// Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.
	SystemDiskSize pulumi.IntPtrInput
	// A mapping of tags to assign to the resource. It will be applied for ECS instances finally.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.
	Tags pulumi.MapInput
	// User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.
	UserData pulumi.StringPtrInput
}

func (ScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationArgs)(nil)).Elem()
}

type ScalingConfigurationInput interface {
	pulumi.Input

	ToScalingConfigurationOutput() ScalingConfigurationOutput
	ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput
}

func (ScalingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingConfiguration)(nil)).Elem()
}

func (i ScalingConfiguration) ToScalingConfigurationOutput() ScalingConfigurationOutput {
	return i.ToScalingConfigurationOutputWithContext(context.Background())
}

func (i ScalingConfiguration) ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationOutput)
}

type ScalingConfigurationOutput struct {
	*pulumi.OutputState
}

func (ScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingConfigurationOutput)(nil)).Elem()
}

func (o ScalingConfigurationOutput) ToScalingConfigurationOutput() ScalingConfigurationOutput {
	return o
}

func (o ScalingConfigurationOutput) ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScalingConfigurationOutput{})
}
