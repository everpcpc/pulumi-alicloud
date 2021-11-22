// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Elastic Cloud Phone (ECP) Key Pair resource.
//
// For information about Elastic Cloud Phone (ECP) Key Pair and how to use it, see [What is Key Pair](https://help.aliyun.com/document_detail/257197.html).
//
// > **NOTE:** Available in v1.130.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecp.NewKeyPair(ctx, "example", &ecp.KeyPairArgs{
// 			KeyPairName:   pulumi.String("my-KeyPair"),
// 			PublicKeyBody: pulumi.String("ssh-rsa AAAAxxxxxxxxxxtyuudsfsg"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Elastic Cloud Phone (ECP) Key Pair can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ecp/keyPair:KeyPair example <key_pair_name>
// ```
type KeyPair struct {
	pulumi.CustomResourceState

	// The Key Name.
	KeyPairName pulumi.StringOutput `pulumi:"keyPairName"`
	// The public key body.
	PublicKeyBody pulumi.StringOutput `pulumi:"publicKeyBody"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyPairName == nil {
		return nil, errors.New("invalid value for required argument 'KeyPairName'")
	}
	if args.PublicKeyBody == nil {
		return nil, errors.New("invalid value for required argument 'PublicKeyBody'")
	}
	var resource KeyPair
	err := ctx.RegisterResource("alicloud:ecp/keyPair:KeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairState, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	var resource KeyPair
	err := ctx.ReadResource("alicloud:ecp/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	// The Key Name.
	KeyPairName *string `pulumi:"keyPairName"`
	// The public key body.
	PublicKeyBody *string `pulumi:"publicKeyBody"`
}

type KeyPairState struct {
	// The Key Name.
	KeyPairName pulumi.StringPtrInput
	// The public key body.
	PublicKeyBody pulumi.StringPtrInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The Key Name.
	KeyPairName string `pulumi:"keyPairName"`
	// The public key body.
	PublicKeyBody string `pulumi:"publicKeyBody"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The Key Name.
	KeyPairName pulumi.StringInput
	// The public key body.
	PublicKeyBody pulumi.StringInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}

type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput
}

func (*KeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (i *KeyPair) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

func (i *KeyPair) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

type KeyPairPtrInput interface {
	pulumi.Input

	ToKeyPairPtrOutput() KeyPairPtrOutput
	ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput
}

type keyPairPtrType KeyPairArgs

func (*keyPairPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (i *keyPairPtrType) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *keyPairPtrType) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

// KeyPairArrayInput is an input type that accepts KeyPairArray and KeyPairArrayOutput values.
// You can construct a concrete instance of `KeyPairArrayInput` via:
//
//          KeyPairArray{ KeyPairArgs{...} }
type KeyPairArrayInput interface {
	pulumi.Input

	ToKeyPairArrayOutput() KeyPairArrayOutput
	ToKeyPairArrayOutputWithContext(context.Context) KeyPairArrayOutput
}

type KeyPairArray []KeyPairInput

func (KeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPair)(nil)).Elem()
}

func (i KeyPairArray) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return i.ToKeyPairArrayOutputWithContext(context.Background())
}

func (i KeyPairArray) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairArrayOutput)
}

// KeyPairMapInput is an input type that accepts KeyPairMap and KeyPairMapOutput values.
// You can construct a concrete instance of `KeyPairMapInput` via:
//
//          KeyPairMap{ "key": KeyPairArgs{...} }
type KeyPairMapInput interface {
	pulumi.Input

	ToKeyPairMapOutput() KeyPairMapOutput
	ToKeyPairMapOutputWithContext(context.Context) KeyPairMapOutput
}

type KeyPairMap map[string]KeyPairInput

func (KeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPair)(nil)).Elem()
}

func (i KeyPairMap) ToKeyPairMapOutput() KeyPairMapOutput {
	return i.ToKeyPairMapOutputWithContext(context.Background())
}

func (i KeyPairMap) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairMapOutput)
}

type KeyPairOutput struct{ *pulumi.OutputState }

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o.ToKeyPairPtrOutputWithContext(context.Background())
}

func (o KeyPairOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyPair) *KeyPair {
		return &v
	}).(KeyPairPtrOutput)
}

type KeyPairPtrOutput struct{ *pulumi.OutputState }

func (KeyPairPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o
}

func (o KeyPairPtrOutput) Elem() KeyPairOutput {
	return o.ApplyT(func(v *KeyPair) KeyPair {
		if v != nil {
			return *v
		}
		var ret KeyPair
		return ret
	}).(KeyPairOutput)
}

type KeyPairArrayOutput struct{ *pulumi.OutputState }

func (KeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyPair)(nil))
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) Index(i pulumi.IntInput) KeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].([]KeyPair)[vs[1].(int)]
	}).(KeyPairOutput)
}

type KeyPairMapOutput struct{ *pulumi.OutputState }

func (KeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyPair)(nil))
}

func (o KeyPairMapOutput) ToKeyPairMapOutput() KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) MapIndex(k pulumi.StringInput) KeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].(map[string]KeyPair)[vs[1].(string)]
	}).(KeyPairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairInput)(nil)).Elem(), &KeyPair{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairPtrInput)(nil)).Elem(), &KeyPair{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairArrayInput)(nil)).Elem(), KeyPairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairMapInput)(nil)).Elem(), KeyPairMap{})
	pulumi.RegisterOutputType(KeyPairOutput{})
	pulumi.RegisterOutputType(KeyPairPtrOutput{})
	pulumi.RegisterOutputType(KeyPairArrayOutput{})
	pulumi.RegisterOutputType(KeyPairMapOutput{})
}
