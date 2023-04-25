// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Alidns Address Pool resource.
//
// For information about Alidns Address Pool and how to use it, see [What is Address Pool](https://www.alibabacloud.com/help/doc-detail/189621.html).
//
// > **NOTE:** Available in v1.152.0+.
//
// ## Import
//
// Alidns Address Pool can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dns/addressPool:AddressPool example <id>
//
// ```
type AddressPool struct {
	pulumi.CustomResourceState

	// The name of the address pool.
	AddressPoolName pulumi.StringOutput `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See the following `Block address`.
	Addresses AddressPoolAddressArrayOutput `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringOutput `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAddressPool registers a new resource with the given unique name, arguments, and options.
func NewAddressPool(ctx *pulumi.Context,
	name string, args *AddressPoolArgs, opts ...pulumi.ResourceOption) (*AddressPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressPoolName == nil {
		return nil, errors.New("invalid value for required argument 'AddressPoolName'")
	}
	if args.Addresses == nil {
		return nil, errors.New("invalid value for required argument 'Addresses'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.LbaStrategy == nil {
		return nil, errors.New("invalid value for required argument 'LbaStrategy'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource AddressPool
	err := ctx.RegisterResource("alicloud:dns/addressPool:AddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressPool gets an existing AddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressPoolState, opts ...pulumi.ResourceOption) (*AddressPool, error) {
	var resource AddressPool
	err := ctx.ReadResource("alicloud:dns/addressPool:AddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressPool resources.
type addressPoolState struct {
	// The name of the address pool.
	AddressPoolName *string `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See the following `Block address`.
	Addresses []AddressPoolAddress `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy *string `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type *string `pulumi:"type"`
}

type AddressPoolState struct {
	// The name of the address pool.
	AddressPoolName pulumi.StringPtrInput
	// The address lists of the Address Pool. See the following `Block address`.
	Addresses AddressPoolAddressArrayInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringPtrInput
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringPtrInput
}

func (AddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressPoolState)(nil)).Elem()
}

type addressPoolArgs struct {
	// The name of the address pool.
	AddressPoolName string `pulumi:"addressPoolName"`
	// The address lists of the Address Pool. See the following `Block address`.
	Addresses []AddressPoolAddress `pulumi:"addresses"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy string `pulumi:"lbaStrategy"`
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AddressPool resource.
type AddressPoolArgs struct {
	// The name of the address pool.
	AddressPoolName pulumi.StringInput
	// The address lists of the Address Pool. See the following `Block address`.
	Addresses AddressPoolAddressArrayInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
	LbaStrategy pulumi.StringInput
	// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
	Type pulumi.StringInput
}

func (AddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressPoolArgs)(nil)).Elem()
}

type AddressPoolInput interface {
	pulumi.Input

	ToAddressPoolOutput() AddressPoolOutput
	ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput
}

func (*AddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPool)(nil)).Elem()
}

func (i *AddressPool) ToAddressPoolOutput() AddressPoolOutput {
	return i.ToAddressPoolOutputWithContext(context.Background())
}

func (i *AddressPool) ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolOutput)
}

// AddressPoolArrayInput is an input type that accepts AddressPoolArray and AddressPoolArrayOutput values.
// You can construct a concrete instance of `AddressPoolArrayInput` via:
//
//	AddressPoolArray{ AddressPoolArgs{...} }
type AddressPoolArrayInput interface {
	pulumi.Input

	ToAddressPoolArrayOutput() AddressPoolArrayOutput
	ToAddressPoolArrayOutputWithContext(context.Context) AddressPoolArrayOutput
}

type AddressPoolArray []AddressPoolInput

func (AddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressPool)(nil)).Elem()
}

func (i AddressPoolArray) ToAddressPoolArrayOutput() AddressPoolArrayOutput {
	return i.ToAddressPoolArrayOutputWithContext(context.Background())
}

func (i AddressPoolArray) ToAddressPoolArrayOutputWithContext(ctx context.Context) AddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolArrayOutput)
}

// AddressPoolMapInput is an input type that accepts AddressPoolMap and AddressPoolMapOutput values.
// You can construct a concrete instance of `AddressPoolMapInput` via:
//
//	AddressPoolMap{ "key": AddressPoolArgs{...} }
type AddressPoolMapInput interface {
	pulumi.Input

	ToAddressPoolMapOutput() AddressPoolMapOutput
	ToAddressPoolMapOutputWithContext(context.Context) AddressPoolMapOutput
}

type AddressPoolMap map[string]AddressPoolInput

func (AddressPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressPool)(nil)).Elem()
}

func (i AddressPoolMap) ToAddressPoolMapOutput() AddressPoolMapOutput {
	return i.ToAddressPoolMapOutputWithContext(context.Background())
}

func (i AddressPoolMap) ToAddressPoolMapOutputWithContext(ctx context.Context) AddressPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPoolMapOutput)
}

type AddressPoolOutput struct{ *pulumi.OutputState }

func (AddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPool)(nil)).Elem()
}

func (o AddressPoolOutput) ToAddressPoolOutput() AddressPoolOutput {
	return o
}

func (o AddressPoolOutput) ToAddressPoolOutputWithContext(ctx context.Context) AddressPoolOutput {
	return o
}

// The name of the address pool.
func (o AddressPoolOutput) AddressPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.AddressPoolName }).(pulumi.StringOutput)
}

// The address lists of the Address Pool. See the following `Block address`.
func (o AddressPoolOutput) Addresses() AddressPoolAddressArrayOutput {
	return o.ApplyT(func(v *AddressPool) AddressPoolAddressArrayOutput { return v.Addresses }).(AddressPoolAddressArrayOutput)
}

// The ID of the instance.
func (o AddressPoolOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
func (o AddressPoolOutput) LbaStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.LbaStrategy }).(pulumi.StringOutput)
}

// The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
func (o AddressPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AddressPoolArrayOutput struct{ *pulumi.OutputState }

func (AddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressPool)(nil)).Elem()
}

func (o AddressPoolArrayOutput) ToAddressPoolArrayOutput() AddressPoolArrayOutput {
	return o
}

func (o AddressPoolArrayOutput) ToAddressPoolArrayOutputWithContext(ctx context.Context) AddressPoolArrayOutput {
	return o
}

func (o AddressPoolArrayOutput) Index(i pulumi.IntInput) AddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AddressPool {
		return vs[0].([]*AddressPool)[vs[1].(int)]
	}).(AddressPoolOutput)
}

type AddressPoolMapOutput struct{ *pulumi.OutputState }

func (AddressPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressPool)(nil)).Elem()
}

func (o AddressPoolMapOutput) ToAddressPoolMapOutput() AddressPoolMapOutput {
	return o
}

func (o AddressPoolMapOutput) ToAddressPoolMapOutputWithContext(ctx context.Context) AddressPoolMapOutput {
	return o
}

func (o AddressPoolMapOutput) MapIndex(k pulumi.StringInput) AddressPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AddressPool {
		return vs[0].(map[string]*AddressPool)[vs[1].(string)]
	}).(AddressPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolInput)(nil)).Elem(), &AddressPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolArrayInput)(nil)).Elem(), AddressPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPoolMapInput)(nil)).Elem(), AddressPoolMap{})
	pulumi.RegisterOutputType(AddressPoolOutput{})
	pulumi.RegisterOutputType(AddressPoolArrayOutput{})
	pulumi.RegisterOutputType(AddressPoolMapOutput{})
}
