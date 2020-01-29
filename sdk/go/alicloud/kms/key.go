// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A kms key can help user to protect data security in the transmission process.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/kms_key.html.markdown.
type Key struct {
	pulumi.CustomResourceState

	// The Alicloud Resource Name (ARN) of the key.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
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
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	Description *string `pulumi:"description"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
	KeyUsage *string `pulumi:"keyUsage"`
}

type KeyState struct {
	// The Alicloud Resource Name (ARN) of the key.
	Arn pulumi.StringPtrInput
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrInput
	Description pulumi.StringPtrInput
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	Description *string `pulumi:"description"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
	KeyUsage *string `pulumi:"keyUsage"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrInput
	Description pulumi.StringPtrInput
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.
	KeyUsage pulumi.StringPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

