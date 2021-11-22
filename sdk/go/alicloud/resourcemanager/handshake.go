// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Manager handshake resource. You can invite accounts to join a resource directory for unified management.
// For information about Resource Manager handshake and how to use it, see [What is Resource Manager handshake](https://www.alibabacloud.com/help/en/doc-detail/135287.htm).
//
// > **NOTE:** Available in v1.82.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcemanager.NewHandshake(ctx, "example", &resourcemanager.HandshakeArgs{
// 			Note:         pulumi.String("test resource manager handshake"),
// 			TargetEntity: pulumi.String("1182775234******"),
// 			TargetType:   pulumi.String("Account"),
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
// Resource Manager handshake can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/handshake:Handshake example h-QmdexeFm1kE*****
// ```
type Handshake struct {
	pulumi.CustomResourceState

	// The expiration time of the invitation.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Resource account master account ID.
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`
	// The name of the main account of the resource directory.
	MasterAccountName pulumi.StringOutput `pulumi:"masterAccountName"`
	// The modification time of the invitation.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// Remarks. The maximum length is 1024 characters.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringOutput `pulumi:"resourceDirectoryId"`
	// Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Invited account ID or login email.
	TargetEntity pulumi.StringOutput `pulumi:"targetEntity"`
	// Type of account being invited. Valid values: `Account`, `Email`.
	TargetType pulumi.StringOutput `pulumi:"targetType"`
}

// NewHandshake registers a new resource with the given unique name, arguments, and options.
func NewHandshake(ctx *pulumi.Context,
	name string, args *HandshakeArgs, opts ...pulumi.ResourceOption) (*Handshake, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetEntity == nil {
		return nil, errors.New("invalid value for required argument 'TargetEntity'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	var resource Handshake
	err := ctx.RegisterResource("alicloud:resourcemanager/handshake:Handshake", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHandshake gets an existing Handshake resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHandshake(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HandshakeState, opts ...pulumi.ResourceOption) (*Handshake, error) {
	var resource Handshake
	err := ctx.ReadResource("alicloud:resourcemanager/handshake:Handshake", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Handshake resources.
type handshakeState struct {
	// The expiration time of the invitation.
	ExpireTime *string `pulumi:"expireTime"`
	// Resource account master account ID.
	MasterAccountId *string `pulumi:"masterAccountId"`
	// The name of the main account of the resource directory.
	MasterAccountName *string `pulumi:"masterAccountName"`
	// The modification time of the invitation.
	ModifyTime *string `pulumi:"modifyTime"`
	// Remarks. The maximum length is 1024 characters.
	Note *string `pulumi:"note"`
	// Resource directory ID.
	ResourceDirectoryId *string `pulumi:"resourceDirectoryId"`
	// Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
	Status *string `pulumi:"status"`
	// Invited account ID or login email.
	TargetEntity *string `pulumi:"targetEntity"`
	// Type of account being invited. Valid values: `Account`, `Email`.
	TargetType *string `pulumi:"targetType"`
}

type HandshakeState struct {
	// The expiration time of the invitation.
	ExpireTime pulumi.StringPtrInput
	// Resource account master account ID.
	MasterAccountId pulumi.StringPtrInput
	// The name of the main account of the resource directory.
	MasterAccountName pulumi.StringPtrInput
	// The modification time of the invitation.
	ModifyTime pulumi.StringPtrInput
	// Remarks. The maximum length is 1024 characters.
	Note pulumi.StringPtrInput
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringPtrInput
	// Invitation status. Valid values: `Pending` waiting for confirmation, `Accepted`, `Cancelled`, `Declined`, `Expired`.
	Status pulumi.StringPtrInput
	// Invited account ID or login email.
	TargetEntity pulumi.StringPtrInput
	// Type of account being invited. Valid values: `Account`, `Email`.
	TargetType pulumi.StringPtrInput
}

func (HandshakeState) ElementType() reflect.Type {
	return reflect.TypeOf((*handshakeState)(nil)).Elem()
}

type handshakeArgs struct {
	// Remarks. The maximum length is 1024 characters.
	Note *string `pulumi:"note"`
	// Invited account ID or login email.
	TargetEntity string `pulumi:"targetEntity"`
	// Type of account being invited. Valid values: `Account`, `Email`.
	TargetType string `pulumi:"targetType"`
}

// The set of arguments for constructing a Handshake resource.
type HandshakeArgs struct {
	// Remarks. The maximum length is 1024 characters.
	Note pulumi.StringPtrInput
	// Invited account ID or login email.
	TargetEntity pulumi.StringInput
	// Type of account being invited. Valid values: `Account`, `Email`.
	TargetType pulumi.StringInput
}

func (HandshakeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*handshakeArgs)(nil)).Elem()
}

type HandshakeInput interface {
	pulumi.Input

	ToHandshakeOutput() HandshakeOutput
	ToHandshakeOutputWithContext(ctx context.Context) HandshakeOutput
}

func (*Handshake) ElementType() reflect.Type {
	return reflect.TypeOf((*Handshake)(nil))
}

func (i *Handshake) ToHandshakeOutput() HandshakeOutput {
	return i.ToHandshakeOutputWithContext(context.Background())
}

func (i *Handshake) ToHandshakeOutputWithContext(ctx context.Context) HandshakeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandshakeOutput)
}

func (i *Handshake) ToHandshakePtrOutput() HandshakePtrOutput {
	return i.ToHandshakePtrOutputWithContext(context.Background())
}

func (i *Handshake) ToHandshakePtrOutputWithContext(ctx context.Context) HandshakePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandshakePtrOutput)
}

type HandshakePtrInput interface {
	pulumi.Input

	ToHandshakePtrOutput() HandshakePtrOutput
	ToHandshakePtrOutputWithContext(ctx context.Context) HandshakePtrOutput
}

type handshakePtrType HandshakeArgs

func (*handshakePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Handshake)(nil))
}

func (i *handshakePtrType) ToHandshakePtrOutput() HandshakePtrOutput {
	return i.ToHandshakePtrOutputWithContext(context.Background())
}

func (i *handshakePtrType) ToHandshakePtrOutputWithContext(ctx context.Context) HandshakePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandshakePtrOutput)
}

// HandshakeArrayInput is an input type that accepts HandshakeArray and HandshakeArrayOutput values.
// You can construct a concrete instance of `HandshakeArrayInput` via:
//
//          HandshakeArray{ HandshakeArgs{...} }
type HandshakeArrayInput interface {
	pulumi.Input

	ToHandshakeArrayOutput() HandshakeArrayOutput
	ToHandshakeArrayOutputWithContext(context.Context) HandshakeArrayOutput
}

type HandshakeArray []HandshakeInput

func (HandshakeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Handshake)(nil)).Elem()
}

func (i HandshakeArray) ToHandshakeArrayOutput() HandshakeArrayOutput {
	return i.ToHandshakeArrayOutputWithContext(context.Background())
}

func (i HandshakeArray) ToHandshakeArrayOutputWithContext(ctx context.Context) HandshakeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandshakeArrayOutput)
}

// HandshakeMapInput is an input type that accepts HandshakeMap and HandshakeMapOutput values.
// You can construct a concrete instance of `HandshakeMapInput` via:
//
//          HandshakeMap{ "key": HandshakeArgs{...} }
type HandshakeMapInput interface {
	pulumi.Input

	ToHandshakeMapOutput() HandshakeMapOutput
	ToHandshakeMapOutputWithContext(context.Context) HandshakeMapOutput
}

type HandshakeMap map[string]HandshakeInput

func (HandshakeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Handshake)(nil)).Elem()
}

func (i HandshakeMap) ToHandshakeMapOutput() HandshakeMapOutput {
	return i.ToHandshakeMapOutputWithContext(context.Background())
}

func (i HandshakeMap) ToHandshakeMapOutputWithContext(ctx context.Context) HandshakeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandshakeMapOutput)
}

type HandshakeOutput struct{ *pulumi.OutputState }

func (HandshakeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Handshake)(nil))
}

func (o HandshakeOutput) ToHandshakeOutput() HandshakeOutput {
	return o
}

func (o HandshakeOutput) ToHandshakeOutputWithContext(ctx context.Context) HandshakeOutput {
	return o
}

func (o HandshakeOutput) ToHandshakePtrOutput() HandshakePtrOutput {
	return o.ToHandshakePtrOutputWithContext(context.Background())
}

func (o HandshakeOutput) ToHandshakePtrOutputWithContext(ctx context.Context) HandshakePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Handshake) *Handshake {
		return &v
	}).(HandshakePtrOutput)
}

type HandshakePtrOutput struct{ *pulumi.OutputState }

func (HandshakePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Handshake)(nil))
}

func (o HandshakePtrOutput) ToHandshakePtrOutput() HandshakePtrOutput {
	return o
}

func (o HandshakePtrOutput) ToHandshakePtrOutputWithContext(ctx context.Context) HandshakePtrOutput {
	return o
}

func (o HandshakePtrOutput) Elem() HandshakeOutput {
	return o.ApplyT(func(v *Handshake) Handshake {
		if v != nil {
			return *v
		}
		var ret Handshake
		return ret
	}).(HandshakeOutput)
}

type HandshakeArrayOutput struct{ *pulumi.OutputState }

func (HandshakeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Handshake)(nil))
}

func (o HandshakeArrayOutput) ToHandshakeArrayOutput() HandshakeArrayOutput {
	return o
}

func (o HandshakeArrayOutput) ToHandshakeArrayOutputWithContext(ctx context.Context) HandshakeArrayOutput {
	return o
}

func (o HandshakeArrayOutput) Index(i pulumi.IntInput) HandshakeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Handshake {
		return vs[0].([]Handshake)[vs[1].(int)]
	}).(HandshakeOutput)
}

type HandshakeMapOutput struct{ *pulumi.OutputState }

func (HandshakeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Handshake)(nil))
}

func (o HandshakeMapOutput) ToHandshakeMapOutput() HandshakeMapOutput {
	return o
}

func (o HandshakeMapOutput) ToHandshakeMapOutputWithContext(ctx context.Context) HandshakeMapOutput {
	return o
}

func (o HandshakeMapOutput) MapIndex(k pulumi.StringInput) HandshakeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Handshake {
		return vs[0].(map[string]Handshake)[vs[1].(string)]
	}).(HandshakeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HandshakeInput)(nil)).Elem(), &Handshake{})
	pulumi.RegisterInputType(reflect.TypeOf((*HandshakePtrInput)(nil)).Elem(), &Handshake{})
	pulumi.RegisterInputType(reflect.TypeOf((*HandshakeArrayInput)(nil)).Elem(), HandshakeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HandshakeMapInput)(nil)).Elem(), HandshakeMap{})
	pulumi.RegisterOutputType(HandshakeOutput{})
	pulumi.RegisterOutputType(HandshakePtrOutput{})
	pulumi.RegisterOutputType(HandshakeArrayOutput{})
	pulumi.RegisterOutputType(HandshakeMapOutput{})
}
