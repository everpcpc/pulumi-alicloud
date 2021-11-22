// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// VPC can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/network:Network example vpc-abc123456
// ```
type Network struct {
	pulumi.CustomResourceState

	// The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
	CidrBlock pulumi.StringPtrOutput `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
	EnableIpv6 pulumi.BoolPtrOutput `pulumi:"enableIpv6"`
	// (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The route table ID of the router created by default on VPC creation.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The ID of the router created by default on VPC creation.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId pulumi.StringOutput `pulumi:"routerTableId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayOutput `pulumi:"secondaryCidrBlocks"`
	Status              pulumi.StringOutput      `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The user cidrs of the VPC.
	UserCidrs pulumi.StringArrayOutput `pulumi:"userCidrs"`
	// The name of the VPC. Defaults to null.
	VpcName pulumi.StringOutput `pulumi:"vpcName"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		args = &NetworkArgs{}
	}

	var resource Network
	err := ctx.RegisterResource("alicloud:vpc/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("alicloud:vpc/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description *string `pulumi:"description"`
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun *bool `pulumi:"dryRun"`
	// Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
	Name *string `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table ID of the router created by default on VPC creation.
	RouteTableId *string `pulumi:"routeTableId"`
	// The ID of the router created by default on VPC creation.
	RouterId *string `pulumi:"routerId"`
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId *string `pulumi:"routerTableId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks []string `pulumi:"secondaryCidrBlocks"`
	Status              *string  `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The user cidrs of the VPC.
	UserCidrs []string `pulumi:"userCidrs"`
	// The name of the VPC. Defaults to null.
	VpcName *string `pulumi:"vpcName"`
}

type NetworkState struct {
	// The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
	CidrBlock pulumi.StringPtrInput
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrInput
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun pulumi.BoolPtrInput
	// Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
	EnableIpv6 pulumi.BoolPtrInput
	// (Available in v1.119.0+) ) The ipv6 cidr block of VPC.
	Ipv6CidrBlock pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
	Name pulumi.StringPtrInput
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The route table ID of the router created by default on VPC creation.
	RouteTableId pulumi.StringPtrInput
	// The ID of the router created by default on VPC creation.
	RouterId pulumi.StringPtrInput
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId pulumi.StringPtrInput
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayInput
	Status              pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The user cidrs of the VPC.
	UserCidrs pulumi.StringArrayInput
	// The name of the VPC. Defaults to null.
	VpcName pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description *string `pulumi:"description"`
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun *bool `pulumi:"dryRun"`
	// Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
	Name *string `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks []string `pulumi:"secondaryCidrBlocks"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The user cidrs of the VPC.
	UserCidrs []string `pulumi:"userCidrs"`
	// The name of the VPC. Defaults to null.
	VpcName *string `pulumi:"vpcName"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The CIDR block for the VPC. The `cidrBlock` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
	CidrBlock pulumi.StringPtrInput
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrInput
	// Specifies whether to precheck this request only. Valid values: `true` and `false`.
	DryRun pulumi.BoolPtrInput
	// Specifies whether to enable the IPv6 CIDR block. Valid values: `false` (Default): disables IPv6 CIDR blocks. `true`: enables IPv6 CIDR blocks. If the `enableIpv6` is `true`, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56.
	EnableIpv6 pulumi.BoolPtrInput
	// Field `name` has been deprecated from provider version 1.119.0. New field `vpcName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
	Name pulumi.StringPtrInput
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The user cidrs of the VPC.
	UserCidrs pulumi.StringArrayInput
	// The name of the VPC. Defaults to null.
	VpcName pulumi.StringPtrInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

func (i *Network) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *Network) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

type NetworkPtrInput interface {
	pulumi.Input

	ToNetworkPtrOutput() NetworkPtrOutput
	ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput
}

type networkPtrType NetworkArgs

func (*networkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (i *networkPtrType) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *networkPtrType) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

// NetworkArrayInput is an input type that accepts NetworkArray and NetworkArrayOutput values.
// You can construct a concrete instance of `NetworkArrayInput` via:
//
//          NetworkArray{ NetworkArgs{...} }
type NetworkArrayInput interface {
	pulumi.Input

	ToNetworkArrayOutput() NetworkArrayOutput
	ToNetworkArrayOutputWithContext(context.Context) NetworkArrayOutput
}

type NetworkArray []NetworkInput

func (NetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network)(nil)).Elem()
}

func (i NetworkArray) ToNetworkArrayOutput() NetworkArrayOutput {
	return i.ToNetworkArrayOutputWithContext(context.Background())
}

func (i NetworkArray) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkArrayOutput)
}

// NetworkMapInput is an input type that accepts NetworkMap and NetworkMapOutput values.
// You can construct a concrete instance of `NetworkMapInput` via:
//
//          NetworkMap{ "key": NetworkArgs{...} }
type NetworkMapInput interface {
	pulumi.Input

	ToNetworkMapOutput() NetworkMapOutput
	ToNetworkMapOutputWithContext(context.Context) NetworkMapOutput
}

type NetworkMap map[string]NetworkInput

func (NetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network)(nil)).Elem()
}

func (i NetworkMap) ToNetworkMapOutput() NetworkMapOutput {
	return i.ToNetworkMapOutputWithContext(context.Background())
}

func (i NetworkMap) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMapOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o.ToNetworkPtrOutputWithContext(context.Background())
}

func (o NetworkOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Network) *Network {
		return &v
	}).(NetworkPtrOutput)
}

type NetworkPtrOutput struct{ *pulumi.OutputState }

func (NetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (o NetworkPtrOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) Elem() NetworkOutput {
	return o.ApplyT(func(v *Network) Network {
		if v != nil {
			return *v
		}
		var ret Network
		return ret
	}).(NetworkOutput)
}

type NetworkArrayOutput struct{ *pulumi.OutputState }

func (NetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Network)(nil))
}

func (o NetworkArrayOutput) ToNetworkArrayOutput() NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) Index(i pulumi.IntInput) NetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Network {
		return vs[0].([]Network)[vs[1].(int)]
	}).(NetworkOutput)
}

type NetworkMapOutput struct{ *pulumi.OutputState }

func (NetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Network)(nil))
}

func (o NetworkMapOutput) ToNetworkMapOutput() NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) MapIndex(k pulumi.StringInput) NetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Network {
		return vs[0].(map[string]Network)[vs[1].(string)]
	}).(NetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInput)(nil)).Elem(), &Network{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPtrInput)(nil)).Elem(), &Network{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkArrayInput)(nil)).Elem(), NetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkMapInput)(nil)).Elem(), NetworkMap{})
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkPtrOutput{})
	pulumi.RegisterOutputType(NetworkArrayOutput{})
	pulumi.RegisterOutputType(NetworkMapOutput{})
}
