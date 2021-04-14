// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type KeyPairAttachment struct {
	pulumi.CustomResourceState

	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// The name of key pair used to bind.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName     pulumi.StringOutput `pulumi:"keyName"`
	KeyPairName pulumi.StringOutput `pulumi:"keyPairName"`
}

// NewKeyPairAttachment registers a new resource with the given unique name, arguments, and options.
func NewKeyPairAttachment(ctx *pulumi.Context,
	name string, args *KeyPairAttachmentArgs, opts ...pulumi.ResourceOption) (*KeyPairAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceIds == nil {
		return nil, errors.New("invalid value for required argument 'InstanceIds'")
	}
	var resource KeyPairAttachment
	err := ctx.RegisterResource("alicloud:ecs/keyPairAttachment:KeyPairAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPairAttachment gets an existing KeyPairAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPairAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairAttachmentState, opts ...pulumi.ResourceOption) (*KeyPairAttachment, error) {
	var resource KeyPairAttachment
	err := ctx.ReadResource("alicloud:ecs/keyPairAttachment:KeyPairAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPairAttachment resources.
type keyPairAttachmentState struct {
	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force *bool `pulumi:"force"`
	// The list of ECS instance's IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The name of key pair used to bind.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName     *string `pulumi:"keyName"`
	KeyPairName *string `pulumi:"keyPairName"`
}

type KeyPairAttachmentState struct {
	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force pulumi.BoolPtrInput
	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayInput
	// The name of key pair used to bind.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName     pulumi.StringPtrInput
	KeyPairName pulumi.StringPtrInput
}

func (KeyPairAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairAttachmentState)(nil)).Elem()
}

type keyPairAttachmentArgs struct {
	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force *bool `pulumi:"force"`
	// The list of ECS instance's IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The name of key pair used to bind.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName     *string `pulumi:"keyName"`
	KeyPairName *string `pulumi:"keyPairName"`
}

// The set of arguments for constructing a KeyPairAttachment resource.
type KeyPairAttachmentArgs struct {
	// Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
	Force pulumi.BoolPtrInput
	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayInput
	// The name of key pair used to bind.
	//
	// Deprecated: Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
	KeyName     pulumi.StringPtrInput
	KeyPairName pulumi.StringPtrInput
}

func (KeyPairAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairAttachmentArgs)(nil)).Elem()
}

type KeyPairAttachmentInput interface {
	pulumi.Input

	ToKeyPairAttachmentOutput() KeyPairAttachmentOutput
	ToKeyPairAttachmentOutputWithContext(ctx context.Context) KeyPairAttachmentOutput
}

func (*KeyPairAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairAttachment)(nil))
}

func (i *KeyPairAttachment) ToKeyPairAttachmentOutput() KeyPairAttachmentOutput {
	return i.ToKeyPairAttachmentOutputWithContext(context.Background())
}

func (i *KeyPairAttachment) ToKeyPairAttachmentOutputWithContext(ctx context.Context) KeyPairAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAttachmentOutput)
}

func (i *KeyPairAttachment) ToKeyPairAttachmentPtrOutput() KeyPairAttachmentPtrOutput {
	return i.ToKeyPairAttachmentPtrOutputWithContext(context.Background())
}

func (i *KeyPairAttachment) ToKeyPairAttachmentPtrOutputWithContext(ctx context.Context) KeyPairAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAttachmentPtrOutput)
}

type KeyPairAttachmentPtrInput interface {
	pulumi.Input

	ToKeyPairAttachmentPtrOutput() KeyPairAttachmentPtrOutput
	ToKeyPairAttachmentPtrOutputWithContext(ctx context.Context) KeyPairAttachmentPtrOutput
}

type keyPairAttachmentPtrType KeyPairAttachmentArgs

func (*keyPairAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPairAttachment)(nil))
}

func (i *keyPairAttachmentPtrType) ToKeyPairAttachmentPtrOutput() KeyPairAttachmentPtrOutput {
	return i.ToKeyPairAttachmentPtrOutputWithContext(context.Background())
}

func (i *keyPairAttachmentPtrType) ToKeyPairAttachmentPtrOutputWithContext(ctx context.Context) KeyPairAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAttachmentPtrOutput)
}

// KeyPairAttachmentArrayInput is an input type that accepts KeyPairAttachmentArray and KeyPairAttachmentArrayOutput values.
// You can construct a concrete instance of `KeyPairAttachmentArrayInput` via:
//
//          KeyPairAttachmentArray{ KeyPairAttachmentArgs{...} }
type KeyPairAttachmentArrayInput interface {
	pulumi.Input

	ToKeyPairAttachmentArrayOutput() KeyPairAttachmentArrayOutput
	ToKeyPairAttachmentArrayOutputWithContext(context.Context) KeyPairAttachmentArrayOutput
}

type KeyPairAttachmentArray []KeyPairAttachmentInput

func (KeyPairAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*KeyPairAttachment)(nil))
}

func (i KeyPairAttachmentArray) ToKeyPairAttachmentArrayOutput() KeyPairAttachmentArrayOutput {
	return i.ToKeyPairAttachmentArrayOutputWithContext(context.Background())
}

func (i KeyPairAttachmentArray) ToKeyPairAttachmentArrayOutputWithContext(ctx context.Context) KeyPairAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAttachmentArrayOutput)
}

// KeyPairAttachmentMapInput is an input type that accepts KeyPairAttachmentMap and KeyPairAttachmentMapOutput values.
// You can construct a concrete instance of `KeyPairAttachmentMapInput` via:
//
//          KeyPairAttachmentMap{ "key": KeyPairAttachmentArgs{...} }
type KeyPairAttachmentMapInput interface {
	pulumi.Input

	ToKeyPairAttachmentMapOutput() KeyPairAttachmentMapOutput
	ToKeyPairAttachmentMapOutputWithContext(context.Context) KeyPairAttachmentMapOutput
}

type KeyPairAttachmentMap map[string]KeyPairAttachmentInput

func (KeyPairAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*KeyPairAttachment)(nil))
}

func (i KeyPairAttachmentMap) ToKeyPairAttachmentMapOutput() KeyPairAttachmentMapOutput {
	return i.ToKeyPairAttachmentMapOutputWithContext(context.Background())
}

func (i KeyPairAttachmentMap) ToKeyPairAttachmentMapOutputWithContext(ctx context.Context) KeyPairAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairAttachmentMapOutput)
}

type KeyPairAttachmentOutput struct {
	*pulumi.OutputState
}

func (KeyPairAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairAttachment)(nil))
}

func (o KeyPairAttachmentOutput) ToKeyPairAttachmentOutput() KeyPairAttachmentOutput {
	return o
}

func (o KeyPairAttachmentOutput) ToKeyPairAttachmentOutputWithContext(ctx context.Context) KeyPairAttachmentOutput {
	return o
}

func (o KeyPairAttachmentOutput) ToKeyPairAttachmentPtrOutput() KeyPairAttachmentPtrOutput {
	return o.ToKeyPairAttachmentPtrOutputWithContext(context.Background())
}

func (o KeyPairAttachmentOutput) ToKeyPairAttachmentPtrOutputWithContext(ctx context.Context) KeyPairAttachmentPtrOutput {
	return o.ApplyT(func(v KeyPairAttachment) *KeyPairAttachment {
		return &v
	}).(KeyPairAttachmentPtrOutput)
}

type KeyPairAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (KeyPairAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPairAttachment)(nil))
}

func (o KeyPairAttachmentPtrOutput) ToKeyPairAttachmentPtrOutput() KeyPairAttachmentPtrOutput {
	return o
}

func (o KeyPairAttachmentPtrOutput) ToKeyPairAttachmentPtrOutputWithContext(ctx context.Context) KeyPairAttachmentPtrOutput {
	return o
}

type KeyPairAttachmentArrayOutput struct{ *pulumi.OutputState }

func (KeyPairAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyPairAttachment)(nil))
}

func (o KeyPairAttachmentArrayOutput) ToKeyPairAttachmentArrayOutput() KeyPairAttachmentArrayOutput {
	return o
}

func (o KeyPairAttachmentArrayOutput) ToKeyPairAttachmentArrayOutputWithContext(ctx context.Context) KeyPairAttachmentArrayOutput {
	return o
}

func (o KeyPairAttachmentArrayOutput) Index(i pulumi.IntInput) KeyPairAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyPairAttachment {
		return vs[0].([]KeyPairAttachment)[vs[1].(int)]
	}).(KeyPairAttachmentOutput)
}

type KeyPairAttachmentMapOutput struct{ *pulumi.OutputState }

func (KeyPairAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyPairAttachment)(nil))
}

func (o KeyPairAttachmentMapOutput) ToKeyPairAttachmentMapOutput() KeyPairAttachmentMapOutput {
	return o
}

func (o KeyPairAttachmentMapOutput) ToKeyPairAttachmentMapOutputWithContext(ctx context.Context) KeyPairAttachmentMapOutput {
	return o
}

func (o KeyPairAttachmentMapOutput) MapIndex(k pulumi.StringInput) KeyPairAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyPairAttachment {
		return vs[0].(map[string]KeyPairAttachment)[vs[1].(string)]
	}).(KeyPairAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyPairAttachmentOutput{})
	pulumi.RegisterOutputType(KeyPairAttachmentPtrOutput{})
	pulumi.RegisterOutputType(KeyPairAttachmentArrayOutput{})
	pulumi.RegisterOutputType(KeyPairAttachmentMapOutput{})
}
