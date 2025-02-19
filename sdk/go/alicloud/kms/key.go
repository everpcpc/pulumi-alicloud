// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A kms key can help user to protect data security in the transmission process. For information about Alikms Key and how to use it, see [What is Resource Alikms Key](https://www.alibabacloud.com/help/doc-detail/28947.htm).
//
// > **NOTE:** Available in v1.85.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kms.NewKey(ctx, "key", &kms.KeyArgs{
//				Description:         pulumi.String("Hello KMS"),
//				PendingWindowInDays: pulumi.Int(7),
//				Status:              pulumi.String("Enabled"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Alikms key can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:kms/key:Key example abc123456
//
// ```
type Key struct {
	pulumi.CustomResourceState

	// The Alicloud Resource Name (ARN) of the key.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether to enable automatic key rotation. Valid values:
	// - Enabled
	// - Disabled (default value)
	//   **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
	AutomaticRotation pulumi.StringPtrOutput `pulumi:"automaticRotation"`
	// The date and time when the CMK was created. The time is displayed in UTC.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The creator of the CMK.
	Creator pulumi.StringOutput `pulumi:"creator"`
	// The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	DeleteDate pulumi.StringOutput `pulumi:"deleteDate"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntOutput `pulumi:"deletionWindowInDays"`
	// The description of the CMK. The description can be 0 to 8,192 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The instance ID of the exclusive KMS instance.
	DkmsInstanceId pulumi.StringPtrOutput `pulumi:"dkmsInstanceId"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The type of the CMK. Valid values:
	// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
	// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
	KeySpec pulumi.StringOutput `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// The usage of the CMK. Valid values:
	// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
	// - SIGN/VERIFY: generates or verifies a digital signature.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate pulumi.StringOutput `pulumi:"lastRotationDate"`
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime pulumi.StringOutput `pulumi:"materialExpireTime"`
	// The time the next rotation is scheduled for execution.
	NextRotationDate pulumi.StringOutput `pulumi:"nextRotationDate"`
	// The source of key material. Valid values:
	// - Aliyun_KMS (default value)
	// - EXTERNAL
	//   **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
	//   you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
	//   For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
	Origin pulumi.StringOutput `pulumi:"origin"`
	// The number of days before the CMK is deleted.
	// During this period, the CMK is in the PendingDeletion state.
	// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
	// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
	PendingWindowInDays pulumi.IntOutput `pulumi:"pendingWindowInDays"`
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion pulumi.StringOutput `pulumi:"primaryKeyVersion"`
	// The protection level of the CMK. Valid values:
	// - SOFTWARE (default value)
	// - HSM
	//   **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
	//   If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
	//   If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
	ProtectionLevel pulumi.StringPtrOutput `pulumi:"protectionLevel"`
	// The interval for automatic key rotation. Specify the value in the integer[unit] format.
	// The following units are supported: d (day), h (hour), m (minute), and s (second).
	// For example, you can use either 7d or 604800s to specify a seven-day interval.
	// The interval can range from 7 days to 730 days.
	// **NOTE**: It is Required when `automaticRotation = "Enabled"`
	//
	// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
	RotationInterval pulumi.StringPtrOutput `pulumi:"rotationInterval"`
	// The status of CMK. Valid Values:
	// - Disabled
	// - Enabled (default value)
	// - PendingDeletion
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		args = &KeyArgs{}
	}

	var resource Key
	err := ctx.RegisterResource("alicloud:kms/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("alicloud:kms/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// The Alicloud Resource Name (ARN) of the key.
	Arn *string `pulumi:"arn"`
	// Specifies whether to enable automatic key rotation. Valid values:
	// - Enabled
	// - Disabled (default value)
	//   **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
	AutomaticRotation *string `pulumi:"automaticRotation"`
	// The date and time when the CMK was created. The time is displayed in UTC.
	CreationDate *string `pulumi:"creationDate"`
	// The creator of the CMK.
	Creator *string `pulumi:"creator"`
	// The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	DeleteDate *string `pulumi:"deleteDate"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the CMK. The description can be 0 to 8,192 characters in length.
	Description *string `pulumi:"description"`
	// The instance ID of the exclusive KMS instance.
	DkmsInstanceId *string `pulumi:"dkmsInstanceId"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The type of the CMK. Valid values:
	// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
	// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
	KeySpec *string `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState *string `pulumi:"keyState"`
	// The usage of the CMK. Valid values:
	// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
	// - SIGN/VERIFY: generates or verifies a digital signature.
	KeyUsage *string `pulumi:"keyUsage"`
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate *string `pulumi:"lastRotationDate"`
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime *string `pulumi:"materialExpireTime"`
	// The time the next rotation is scheduled for execution.
	NextRotationDate *string `pulumi:"nextRotationDate"`
	// The source of key material. Valid values:
	// - Aliyun_KMS (default value)
	// - EXTERNAL
	//   **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
	//   you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
	//   For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
	Origin *string `pulumi:"origin"`
	// The number of days before the CMK is deleted.
	// During this period, the CMK is in the PendingDeletion state.
	// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
	// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
	PendingWindowInDays *int `pulumi:"pendingWindowInDays"`
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion *string `pulumi:"primaryKeyVersion"`
	// The protection level of the CMK. Valid values:
	// - SOFTWARE (default value)
	// - HSM
	//   **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
	//   If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
	//   If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The interval for automatic key rotation. Specify the value in the integer[unit] format.
	// The following units are supported: d (day), h (hour), m (minute), and s (second).
	// For example, you can use either 7d or 604800s to specify a seven-day interval.
	// The interval can range from 7 days to 730 days.
	// **NOTE**: It is Required when `automaticRotation = "Enabled"`
	//
	// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
	RotationInterval *string `pulumi:"rotationInterval"`
	// The status of CMK. Valid Values:
	// - Disabled
	// - Enabled (default value)
	// - PendingDeletion
	Status *string `pulumi:"status"`
}

type KeyState struct {
	// The Alicloud Resource Name (ARN) of the key.
	Arn pulumi.StringPtrInput
	// Specifies whether to enable automatic key rotation. Valid values:
	// - Enabled
	// - Disabled (default value)
	//   **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
	AutomaticRotation pulumi.StringPtrInput
	// The date and time when the CMK was created. The time is displayed in UTC.
	CreationDate pulumi.StringPtrInput
	// The creator of the CMK.
	Creator pulumi.StringPtrInput
	// The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
	DeleteDate pulumi.StringPtrInput
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the CMK. The description can be 0 to 8,192 characters in length.
	Description pulumi.StringPtrInput
	// The instance ID of the exclusive KMS instance.
	DkmsInstanceId pulumi.StringPtrInput
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrInput
	// The type of the CMK. Valid values:
	// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
	// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
	KeySpec pulumi.StringPtrInput
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringPtrInput
	// The usage of the CMK. Valid values:
	// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
	// - SIGN/VERIFY: generates or verifies a digital signature.
	KeyUsage pulumi.StringPtrInput
	// The date and time the last rotation was performed. The time is displayed in UTC.
	LastRotationDate pulumi.StringPtrInput
	// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
	MaterialExpireTime pulumi.StringPtrInput
	// The time the next rotation is scheduled for execution.
	NextRotationDate pulumi.StringPtrInput
	// The source of key material. Valid values:
	// - Aliyun_KMS (default value)
	// - EXTERNAL
	//   **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
	//   you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
	//   For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
	Origin pulumi.StringPtrInput
	// The number of days before the CMK is deleted.
	// During this period, the CMK is in the PendingDeletion state.
	// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
	// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
	PendingWindowInDays pulumi.IntPtrInput
	// The ID of the current primary key version of the symmetric CMK.
	PrimaryKeyVersion pulumi.StringPtrInput
	// The protection level of the CMK. Valid values:
	// - SOFTWARE (default value)
	// - HSM
	//   **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
	//   If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
	//   If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
	ProtectionLevel pulumi.StringPtrInput
	// The interval for automatic key rotation. Specify the value in the integer[unit] format.
	// The following units are supported: d (day), h (hour), m (minute), and s (second).
	// For example, you can use either 7d or 604800s to specify a seven-day interval.
	// The interval can range from 7 days to 730 days.
	// **NOTE**: It is Required when `automaticRotation = "Enabled"`
	//
	// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
	RotationInterval pulumi.StringPtrInput
	// The status of CMK. Valid Values:
	// - Disabled
	// - Enabled (default value)
	// - PendingDeletion
	Status pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Specifies whether to enable automatic key rotation. Valid values:
	// - Enabled
	// - Disabled (default value)
	//   **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
	AutomaticRotation *string `pulumi:"automaticRotation"`
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the CMK. The description can be 0 to 8,192 characters in length.
	Description *string `pulumi:"description"`
	// The instance ID of the exclusive KMS instance.
	DkmsInstanceId *string `pulumi:"dkmsInstanceId"`
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The type of the CMK. Valid values:
	// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
	// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
	KeySpec *string `pulumi:"keySpec"`
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState *string `pulumi:"keyState"`
	// The usage of the CMK. Valid values:
	// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
	// - SIGN/VERIFY: generates or verifies a digital signature.
	KeyUsage *string `pulumi:"keyUsage"`
	// The source of key material. Valid values:
	// - Aliyun_KMS (default value)
	// - EXTERNAL
	//   **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
	//   you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
	//   For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
	Origin *string `pulumi:"origin"`
	// The number of days before the CMK is deleted.
	// During this period, the CMK is in the PendingDeletion state.
	// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
	// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
	PendingWindowInDays *int `pulumi:"pendingWindowInDays"`
	// The protection level of the CMK. Valid values:
	// - SOFTWARE (default value)
	// - HSM
	//   **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
	//   If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
	//   If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The interval for automatic key rotation. Specify the value in the integer[unit] format.
	// The following units are supported: d (day), h (hour), m (minute), and s (second).
	// For example, you can use either 7d or 604800s to specify a seven-day interval.
	// The interval can range from 7 days to 730 days.
	// **NOTE**: It is Required when `automaticRotation = "Enabled"`
	//
	// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
	RotationInterval *string `pulumi:"rotationInterval"`
	// The status of CMK. Valid Values:
	// - Disabled
	// - Enabled (default value)
	// - PendingDeletion
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Specifies whether to enable automatic key rotation. Valid values:
	// - Enabled
	// - Disabled (default value)
	//   **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
	AutomaticRotation pulumi.StringPtrInput
	// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
	//
	// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the CMK. The description can be 0 to 8,192 characters in length.
	Description pulumi.StringPtrInput
	// The instance ID of the exclusive KMS instance.
	DkmsInstanceId pulumi.StringPtrInput
	// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
	//
	// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
	IsEnabled pulumi.BoolPtrInput
	// The type of the CMK. Valid values:
	// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
	// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
	KeySpec pulumi.StringPtrInput
	// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
	//
	// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
	KeyState pulumi.StringPtrInput
	// The usage of the CMK. Valid values:
	// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
	// - SIGN/VERIFY: generates or verifies a digital signature.
	KeyUsage pulumi.StringPtrInput
	// The source of key material. Valid values:
	// - Aliyun_KMS (default value)
	// - EXTERNAL
	//   **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
	//   you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
	//   For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
	Origin pulumi.StringPtrInput
	// The number of days before the CMK is deleted.
	// During this period, the CMK is in the PendingDeletion state.
	// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
	// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
	PendingWindowInDays pulumi.IntPtrInput
	// The protection level of the CMK. Valid values:
	// - SOFTWARE (default value)
	// - HSM
	//   **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
	//   If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
	//   If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
	ProtectionLevel pulumi.StringPtrInput
	// The interval for automatic key rotation. Specify the value in the integer[unit] format.
	// The following units are supported: d (day), h (hour), m (minute), and s (second).
	// For example, you can use either 7d or 604800s to specify a seven-day interval.
	// The interval can range from 7 days to 730 days.
	// **NOTE**: It is Required when `automaticRotation = "Enabled"`
	//
	// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
	RotationInterval pulumi.StringPtrInput
	// The status of CMK. Valid Values:
	// - Disabled
	// - Enabled (default value)
	// - PendingDeletion
	Status pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//	KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//	KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

// The Alicloud Resource Name (ARN) of the key.
func (o KeyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether to enable automatic key rotation. Valid values:
//   - Enabled
//   - Disabled (default value)
//     **NOTE**: If you set the origin parameter to EXTERNAL or the keySpec parameter to an asymmetric CMK type, automatic key rotation is unavailable.
func (o KeyOutput) AutomaticRotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.AutomaticRotation }).(pulumi.StringPtrOutput)
}

// The date and time when the CMK was created. The time is displayed in UTC.
func (o KeyOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The creator of the CMK.
func (o KeyOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

// The scheduled date to delete CMK. The time is displayed in UTC. This value is returned only when the KeyState value is PendingDeletion.
func (o KeyOutput) DeleteDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.DeleteDate }).(pulumi.StringOutput)
}

// Field `deletionWindowInDays` has been deprecated from provider version 1.85.0. New field `pendingWindowInDays` instead.
//
// Deprecated: Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.
func (o KeyOutput) DeletionWindowInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Key) pulumi.IntOutput { return v.DeletionWindowInDays }).(pulumi.IntOutput)
}

// The description of the CMK. The description can be 0 to 8,192 characters in length.
func (o KeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The instance ID of the exclusive KMS instance.
func (o KeyOutput) DkmsInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.DkmsInstanceId }).(pulumi.StringPtrOutput)
}

// Field `isEnabled` has been deprecated from provider version 1.85.0. New field `keyState` instead.
//
// Deprecated: Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.
func (o KeyOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The type of the CMK. Valid values:
// "Aliyun_AES_256", "Aliyun_AES_128", "Aliyun_AES_192", "Aliyun_SM4", "RSA_2048", "RSA_3072", "EC_P256", "EC_P256K", "EC_SM2".
// Note: The default type of the CMK is Aliyun_AES_256. Only Dedicated KMS supports Aliyun_AES_128 and Aliyun_AES_192.
func (o KeyOutput) KeySpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeySpec }).(pulumi.StringOutput)
}

// Field `keyState` has been deprecated from provider version 1.123.1. New field `status` instead.
//
// Deprecated: Field 'key_state' has been deprecated from provider version 1.123.1. New field 'status' instead.
func (o KeyOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeyState }).(pulumi.StringOutput)
}

// The usage of the CMK. Valid values:
// - ENCRYPT/DECRYPT(default value): encrypts or decrypts data.
// - SIGN/VERIFY: generates or verifies a digital signature.
func (o KeyOutput) KeyUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyUsage }).(pulumi.StringPtrOutput)
}

// The date and time the last rotation was performed. The time is displayed in UTC.
func (o KeyOutput) LastRotationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.LastRotationDate }).(pulumi.StringOutput)
}

// The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.
func (o KeyOutput) MaterialExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.MaterialExpireTime }).(pulumi.StringOutput)
}

// The time the next rotation is scheduled for execution.
func (o KeyOutput) NextRotationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.NextRotationDate }).(pulumi.StringOutput)
}

// The source of key material. Valid values:
//   - Aliyun_KMS (default value)
//   - EXTERNAL
//     **NOTE**: The value of this parameter is case-sensitive. If you set the `keySpec` to an asymmetric CMK type,
//     you are not allowed to set the `origin` to EXTERNAL. If you set the `origin` to EXTERNAL, you must import key material.
//     For more information, see [import key material](https://www.alibabacloud.com/help/en/doc-detail/68523.htm).
func (o KeyOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Origin }).(pulumi.StringOutput)
}

// The number of days before the CMK is deleted.
// During this period, the CMK is in the PendingDeletion state.
// After this period ends, you cannot cancel the deletion. Valid values: 7 to 366. Unit: days.
// **NOTE:** From version 1.184.0, `pendingWindowInDays` can be set to `366`.
func (o KeyOutput) PendingWindowInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Key) pulumi.IntOutput { return v.PendingWindowInDays }).(pulumi.IntOutput)
}

// The ID of the current primary key version of the symmetric CMK.
func (o KeyOutput) PrimaryKeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.PrimaryKeyVersion }).(pulumi.StringOutput)
}

// The protection level of the CMK. Valid values:
//   - SOFTWARE (default value)
//   - HSM
//     **NOTE**: The value of this parameter is case-sensitive. Assume that you set this parameter to HSM.
//     If you set the origin parameter to Aliyun_KMS, the CMK is created in a managed hardware security module (HSM).
//     If you set the origin parameter to EXTERNA, you can import an external key to the managed HSM.
func (o KeyOutput) ProtectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.ProtectionLevel }).(pulumi.StringPtrOutput)
}

// The interval for automatic key rotation. Specify the value in the integer[unit] format.
// The following units are supported: d (day), h (hour), m (minute), and s (second).
// For example, you can use either 7d or 604800s to specify a seven-day interval.
// The interval can range from 7 days to 730 days.
// **NOTE**: It is Required when `automaticRotation = "Enabled"`
//
// > **NOTE:** When the pre-deletion days elapses, the key is permanently deleted and cannot be recovered.
func (o KeyOutput) RotationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.RotationInterval }).(pulumi.StringPtrOutput)
}

// The status of CMK. Valid Values:
// - Disabled
// - Enabled (default value)
// - PendingDeletion
func (o KeyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Key {
		return vs[0].([]*Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Key {
		return vs[0].(map[string]*Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyArrayInput)(nil)).Elem(), KeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyMapInput)(nil)).Elem(), KeyMap{})
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
