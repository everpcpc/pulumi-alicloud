// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Alikms Key Version resource. For information about Alikms Key Version and how to use it, see [What is Resource Alikms Key Version](https://www.alibabacloud.com/help/doc-detail/133838.htm).
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
//			this, err := kms.NewKey(ctx, "this", nil)
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewKeyVersion(ctx, "keyversion", &kms.KeyVersionArgs{
//				KeyId: this.ID(),
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
// Alikms key version can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:kms/keyVersion:KeyVersion example 72da539a-2fa8-4f2d-b854-*****
//
// ```
type KeyVersion struct {
	pulumi.CustomResourceState

	// The id of the master key (CMK).
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The id of the Alikms key version.
	KeyVersionId pulumi.StringOutput `pulumi:"keyVersionId"`
}

// NewKeyVersion registers a new resource with the given unique name, arguments, and options.
func NewKeyVersion(ctx *pulumi.Context,
	name string, args *KeyVersionArgs, opts ...pulumi.ResourceOption) (*KeyVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	var resource KeyVersion
	err := ctx.RegisterResource("alicloud:kms/keyVersion:KeyVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyVersion gets an existing KeyVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyVersionState, opts ...pulumi.ResourceOption) (*KeyVersion, error) {
	var resource KeyVersion
	err := ctx.ReadResource("alicloud:kms/keyVersion:KeyVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyVersion resources.
type keyVersionState struct {
	// The id of the master key (CMK).
	KeyId *string `pulumi:"keyId"`
	// The id of the Alikms key version.
	KeyVersionId *string `pulumi:"keyVersionId"`
}

type KeyVersionState struct {
	// The id of the master key (CMK).
	KeyId pulumi.StringPtrInput
	// The id of the Alikms key version.
	KeyVersionId pulumi.StringPtrInput
}

func (KeyVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVersionState)(nil)).Elem()
}

type keyVersionArgs struct {
	// The id of the master key (CMK).
	KeyId string `pulumi:"keyId"`
}

// The set of arguments for constructing a KeyVersion resource.
type KeyVersionArgs struct {
	// The id of the master key (CMK).
	KeyId pulumi.StringInput
}

func (KeyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyVersionArgs)(nil)).Elem()
}

type KeyVersionInput interface {
	pulumi.Input

	ToKeyVersionOutput() KeyVersionOutput
	ToKeyVersionOutputWithContext(ctx context.Context) KeyVersionOutput
}

func (*KeyVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVersion)(nil)).Elem()
}

func (i *KeyVersion) ToKeyVersionOutput() KeyVersionOutput {
	return i.ToKeyVersionOutputWithContext(context.Background())
}

func (i *KeyVersion) ToKeyVersionOutputWithContext(ctx context.Context) KeyVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVersionOutput)
}

// KeyVersionArrayInput is an input type that accepts KeyVersionArray and KeyVersionArrayOutput values.
// You can construct a concrete instance of `KeyVersionArrayInput` via:
//
//	KeyVersionArray{ KeyVersionArgs{...} }
type KeyVersionArrayInput interface {
	pulumi.Input

	ToKeyVersionArrayOutput() KeyVersionArrayOutput
	ToKeyVersionArrayOutputWithContext(context.Context) KeyVersionArrayOutput
}

type KeyVersionArray []KeyVersionInput

func (KeyVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyVersion)(nil)).Elem()
}

func (i KeyVersionArray) ToKeyVersionArrayOutput() KeyVersionArrayOutput {
	return i.ToKeyVersionArrayOutputWithContext(context.Background())
}

func (i KeyVersionArray) ToKeyVersionArrayOutputWithContext(ctx context.Context) KeyVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVersionArrayOutput)
}

// KeyVersionMapInput is an input type that accepts KeyVersionMap and KeyVersionMapOutput values.
// You can construct a concrete instance of `KeyVersionMapInput` via:
//
//	KeyVersionMap{ "key": KeyVersionArgs{...} }
type KeyVersionMapInput interface {
	pulumi.Input

	ToKeyVersionMapOutput() KeyVersionMapOutput
	ToKeyVersionMapOutputWithContext(context.Context) KeyVersionMapOutput
}

type KeyVersionMap map[string]KeyVersionInput

func (KeyVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyVersion)(nil)).Elem()
}

func (i KeyVersionMap) ToKeyVersionMapOutput() KeyVersionMapOutput {
	return i.ToKeyVersionMapOutputWithContext(context.Background())
}

func (i KeyVersionMap) ToKeyVersionMapOutputWithContext(ctx context.Context) KeyVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVersionMapOutput)
}

type KeyVersionOutput struct{ *pulumi.OutputState }

func (KeyVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVersion)(nil)).Elem()
}

func (o KeyVersionOutput) ToKeyVersionOutput() KeyVersionOutput {
	return o
}

func (o KeyVersionOutput) ToKeyVersionOutputWithContext(ctx context.Context) KeyVersionOutput {
	return o
}

// The id of the master key (CMK).
func (o KeyVersionOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyVersion) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The id of the Alikms key version.
func (o KeyVersionOutput) KeyVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyVersion) pulumi.StringOutput { return v.KeyVersionId }).(pulumi.StringOutput)
}

type KeyVersionArrayOutput struct{ *pulumi.OutputState }

func (KeyVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyVersion)(nil)).Elem()
}

func (o KeyVersionArrayOutput) ToKeyVersionArrayOutput() KeyVersionArrayOutput {
	return o
}

func (o KeyVersionArrayOutput) ToKeyVersionArrayOutputWithContext(ctx context.Context) KeyVersionArrayOutput {
	return o
}

func (o KeyVersionArrayOutput) Index(i pulumi.IntInput) KeyVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyVersion {
		return vs[0].([]*KeyVersion)[vs[1].(int)]
	}).(KeyVersionOutput)
}

type KeyVersionMapOutput struct{ *pulumi.OutputState }

func (KeyVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyVersion)(nil)).Elem()
}

func (o KeyVersionMapOutput) ToKeyVersionMapOutput() KeyVersionMapOutput {
	return o
}

func (o KeyVersionMapOutput) ToKeyVersionMapOutputWithContext(ctx context.Context) KeyVersionMapOutput {
	return o
}

func (o KeyVersionMapOutput) MapIndex(k pulumi.StringInput) KeyVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyVersion {
		return vs[0].(map[string]*KeyVersion)[vs[1].(string)]
	}).(KeyVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVersionInput)(nil)).Elem(), &KeyVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVersionArrayInput)(nil)).Elem(), KeyVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVersionMapInput)(nil)).Elem(), KeyVersionMap{})
	pulumi.RegisterOutputType(KeyVersionOutput{})
	pulumi.RegisterOutputType(KeyVersionArrayOutput{})
	pulumi.RegisterOutputType(KeyVersionMapOutput{})
}
