// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Private Link Vpc Endpoint resource.
//
// For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://help.aliyun.com/document_detail/120479.html).
//
// > **NOTE:** Available in v1.109.0+.
//
// ## Import
//
// Private Link Vpc Endpoint can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:privatelink/vpcEndpoint:VpcEndpoint example <endpoint_id>
// ```
type VpcEndpoint struct {
	pulumi.CustomResourceState

	// The Bandwidth.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The status of Connection.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// The dry run. Default to: `false`.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The status of Endpoint Business.
	EndpointBusinessStatus pulumi.StringOutput `pulumi:"endpointBusinessStatus"`
	// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
	EndpointDescription pulumi.StringPtrOutput `pulumi:"endpointDescription"`
	// The Endpoint Domain.
	EndpointDomain pulumi.StringOutput `pulumi:"endpointDomain"`
	// The security group associated with the terminal node network card.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The terminal node service associated with the terminal node.
	ServiceId pulumi.StringPtrOutput `pulumi:"serviceId"`
	// The name of the terminal node service associated with the terminal node.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The status of Vpc Endpoint.
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
	VpcEndpointName pulumi.StringPtrOutput `pulumi:"vpcEndpointName"`
	// The private network to which the terminal node belongs.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VpcEndpoint
	err := ctx.RegisterResource("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpoint gets an existing VpcEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointState, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	var resource VpcEndpoint
	err := ctx.ReadResource("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type vpcEndpointState struct {
	// The Bandwidth.
	Bandwidth *int `pulumi:"bandwidth"`
	// The status of Connection.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// The dry run. Default to: `false`.
	DryRun *bool `pulumi:"dryRun"`
	// The status of Endpoint Business.
	EndpointBusinessStatus *string `pulumi:"endpointBusinessStatus"`
	// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
	EndpointDescription *string `pulumi:"endpointDescription"`
	// The Endpoint Domain.
	EndpointDomain *string `pulumi:"endpointDomain"`
	// The security group associated with the terminal node network card.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The terminal node service associated with the terminal node.
	ServiceId *string `pulumi:"serviceId"`
	// The name of the terminal node service associated with the terminal node.
	ServiceName *string `pulumi:"serviceName"`
	// The status of Vpc Endpoint.
	Status *string `pulumi:"status"`
	// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
	VpcEndpointName *string `pulumi:"vpcEndpointName"`
	// The private network to which the terminal node belongs.
	VpcId *string `pulumi:"vpcId"`
}

type VpcEndpointState struct {
	// The Bandwidth.
	Bandwidth pulumi.IntPtrInput
	// The status of Connection.
	ConnectionStatus pulumi.StringPtrInput
	// The dry run. Default to: `false`.
	DryRun pulumi.BoolPtrInput
	// The status of Endpoint Business.
	EndpointBusinessStatus pulumi.StringPtrInput
	// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
	EndpointDescription pulumi.StringPtrInput
	// The Endpoint Domain.
	EndpointDomain pulumi.StringPtrInput
	// The security group associated with the terminal node network card.
	SecurityGroupIds pulumi.StringArrayInput
	// The terminal node service associated with the terminal node.
	ServiceId pulumi.StringPtrInput
	// The name of the terminal node service associated with the terminal node.
	ServiceName pulumi.StringPtrInput
	// The status of Vpc Endpoint.
	Status pulumi.StringPtrInput
	// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
	VpcEndpointName pulumi.StringPtrInput
	// The private network to which the terminal node belongs.
	VpcId pulumi.StringPtrInput
}

func (VpcEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointState)(nil)).Elem()
}

type vpcEndpointArgs struct {
	// The dry run. Default to: `false`.
	DryRun *bool `pulumi:"dryRun"`
	// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
	EndpointDescription *string `pulumi:"endpointDescription"`
	// The security group associated with the terminal node network card.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The terminal node service associated with the terminal node.
	ServiceId *string `pulumi:"serviceId"`
	// The name of the terminal node service associated with the terminal node.
	ServiceName *string `pulumi:"serviceName"`
	// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
	VpcEndpointName *string `pulumi:"vpcEndpointName"`
	// The private network to which the terminal node belongs.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// The dry run. Default to: `false`.
	DryRun pulumi.BoolPtrInput
	// The description of Vpc Endpoint. The length is 2~256 characters and cannot start with `http://` and `https://`.
	EndpointDescription pulumi.StringPtrInput
	// The security group associated with the terminal node network card.
	SecurityGroupIds pulumi.StringArrayInput
	// The terminal node service associated with the terminal node.
	ServiceId pulumi.StringPtrInput
	// The name of the terminal node service associated with the terminal node.
	ServiceName pulumi.StringPtrInput
	// The name of Vpc Endpoint. The length is between 2 and 128 characters, starting with English letters or Chinese, and can include numbers, hyphens (-) and underscores (_).
	VpcEndpointName pulumi.StringPtrInput
	// The private network to which the terminal node belongs.
	VpcId pulumi.StringInput
}

func (VpcEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointArgs)(nil)).Elem()
}

type VpcEndpointInput interface {
	pulumi.Input

	ToVpcEndpointOutput() VpcEndpointOutput
	ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput
}

func (*VpcEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpoint)(nil))
}

func (i *VpcEndpoint) ToVpcEndpointOutput() VpcEndpointOutput {
	return i.ToVpcEndpointOutputWithContext(context.Background())
}

func (i *VpcEndpoint) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointOutput)
}

func (i *VpcEndpoint) ToVpcEndpointPtrOutput() VpcEndpointPtrOutput {
	return i.ToVpcEndpointPtrOutputWithContext(context.Background())
}

func (i *VpcEndpoint) ToVpcEndpointPtrOutputWithContext(ctx context.Context) VpcEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointPtrOutput)
}

type VpcEndpointPtrInput interface {
	pulumi.Input

	ToVpcEndpointPtrOutput() VpcEndpointPtrOutput
	ToVpcEndpointPtrOutputWithContext(ctx context.Context) VpcEndpointPtrOutput
}

type vpcEndpointPtrType VpcEndpointArgs

func (*vpcEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil))
}

func (i *vpcEndpointPtrType) ToVpcEndpointPtrOutput() VpcEndpointPtrOutput {
	return i.ToVpcEndpointPtrOutputWithContext(context.Background())
}

func (i *vpcEndpointPtrType) ToVpcEndpointPtrOutputWithContext(ctx context.Context) VpcEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointPtrOutput)
}

// VpcEndpointArrayInput is an input type that accepts VpcEndpointArray and VpcEndpointArrayOutput values.
// You can construct a concrete instance of `VpcEndpointArrayInput` via:
//
//          VpcEndpointArray{ VpcEndpointArgs{...} }
type VpcEndpointArrayInput interface {
	pulumi.Input

	ToVpcEndpointArrayOutput() VpcEndpointArrayOutput
	ToVpcEndpointArrayOutputWithContext(context.Context) VpcEndpointArrayOutput
}

type VpcEndpointArray []VpcEndpointInput

func (VpcEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpoint)(nil)).Elem()
}

func (i VpcEndpointArray) ToVpcEndpointArrayOutput() VpcEndpointArrayOutput {
	return i.ToVpcEndpointArrayOutputWithContext(context.Background())
}

func (i VpcEndpointArray) ToVpcEndpointArrayOutputWithContext(ctx context.Context) VpcEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointArrayOutput)
}

// VpcEndpointMapInput is an input type that accepts VpcEndpointMap and VpcEndpointMapOutput values.
// You can construct a concrete instance of `VpcEndpointMapInput` via:
//
//          VpcEndpointMap{ "key": VpcEndpointArgs{...} }
type VpcEndpointMapInput interface {
	pulumi.Input

	ToVpcEndpointMapOutput() VpcEndpointMapOutput
	ToVpcEndpointMapOutputWithContext(context.Context) VpcEndpointMapOutput
}

type VpcEndpointMap map[string]VpcEndpointInput

func (VpcEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpoint)(nil)).Elem()
}

func (i VpcEndpointMap) ToVpcEndpointMapOutput() VpcEndpointMapOutput {
	return i.ToVpcEndpointMapOutputWithContext(context.Background())
}

func (i VpcEndpointMap) ToVpcEndpointMapOutputWithContext(ctx context.Context) VpcEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointMapOutput)
}

type VpcEndpointOutput struct{ *pulumi.OutputState }

func (VpcEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpoint)(nil))
}

func (o VpcEndpointOutput) ToVpcEndpointOutput() VpcEndpointOutput {
	return o
}

func (o VpcEndpointOutput) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return o
}

func (o VpcEndpointOutput) ToVpcEndpointPtrOutput() VpcEndpointPtrOutput {
	return o.ToVpcEndpointPtrOutputWithContext(context.Background())
}

func (o VpcEndpointOutput) ToVpcEndpointPtrOutputWithContext(ctx context.Context) VpcEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpcEndpoint) *VpcEndpoint {
		return &v
	}).(VpcEndpointPtrOutput)
}

type VpcEndpointPtrOutput struct{ *pulumi.OutputState }

func (VpcEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil))
}

func (o VpcEndpointPtrOutput) ToVpcEndpointPtrOutput() VpcEndpointPtrOutput {
	return o
}

func (o VpcEndpointPtrOutput) ToVpcEndpointPtrOutputWithContext(ctx context.Context) VpcEndpointPtrOutput {
	return o
}

func (o VpcEndpointPtrOutput) Elem() VpcEndpointOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpoint {
		if v != nil {
			return *v
		}
		var ret VpcEndpoint
		return ret
	}).(VpcEndpointOutput)
}

type VpcEndpointArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpoint)(nil))
}

func (o VpcEndpointArrayOutput) ToVpcEndpointArrayOutput() VpcEndpointArrayOutput {
	return o
}

func (o VpcEndpointArrayOutput) ToVpcEndpointArrayOutputWithContext(ctx context.Context) VpcEndpointArrayOutput {
	return o
}

func (o VpcEndpointArrayOutput) Index(i pulumi.IntInput) VpcEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpoint {
		return vs[0].([]VpcEndpoint)[vs[1].(int)]
	}).(VpcEndpointOutput)
}

type VpcEndpointMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcEndpoint)(nil))
}

func (o VpcEndpointMapOutput) ToVpcEndpointMapOutput() VpcEndpointMapOutput {
	return o
}

func (o VpcEndpointMapOutput) ToVpcEndpointMapOutputWithContext(ctx context.Context) VpcEndpointMapOutput {
	return o
}

func (o VpcEndpointMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcEndpoint {
		return vs[0].(map[string]VpcEndpoint)[vs[1].(string)]
	}).(VpcEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointInput)(nil)).Elem(), &VpcEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointPtrInput)(nil)).Elem(), &VpcEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointArrayInput)(nil)).Elem(), VpcEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointMapInput)(nil)).Elem(), VpcEndpointMap{})
	pulumi.RegisterOutputType(VpcEndpointOutput{})
	pulumi.RegisterOutputType(VpcEndpointPtrOutput{})
	pulumi.RegisterOutputType(VpcEndpointArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointMapOutput{})
}
