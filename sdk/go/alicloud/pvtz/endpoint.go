// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Private Zone Endpoint resource.
//
// For information about Private Zone Endpoint and how to use it, see [What is Endpoint](https://www.alibabacloud.com/help/en/doc-detail/64611.htm).
//
// > **NOTE:** Available in v1.143.0+.
//
// ## Import
//
// Private Zone Endpoint can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:pvtz/endpoint:Endpoint example <id>
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// The name of the resource.
	EndpointName pulumi.StringOutput `pulumi:"endpointName"`
	// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
	IpConfigs EndpointIpConfigArrayOutput `pulumi:"ipConfigs"`
	// The ID of the Security Group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The status of the resource. Valid values: `CHANGE_FAILED`, `CHANGE_INIT`, `EXCEPTION`, `FAILED`, `INIT`, `SUCCESS`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The Region of the VPC.
	VpcRegionId pulumi.StringOutput `pulumi:"vpcRegionId"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.IpConfigs == nil {
		return nil, errors.New("invalid value for required argument 'IpConfigs'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VpcRegionId == nil {
		return nil, errors.New("invalid value for required argument 'VpcRegionId'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("alicloud:pvtz/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("alicloud:pvtz/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// The name of the resource.
	EndpointName *string `pulumi:"endpointName"`
	// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
	IpConfigs []EndpointIpConfig `pulumi:"ipConfigs"`
	// The ID of the Security Group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The status of the resource. Valid values: `CHANGE_FAILED`, `CHANGE_INIT`, `EXCEPTION`, `FAILED`, `INIT`, `SUCCESS`.
	Status *string `pulumi:"status"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
	// The Region of the VPC.
	VpcRegionId *string `pulumi:"vpcRegionId"`
}

type EndpointState struct {
	// The name of the resource.
	EndpointName pulumi.StringPtrInput
	// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
	IpConfigs EndpointIpConfigArrayInput
	// The ID of the Security Group.
	SecurityGroupId pulumi.StringPtrInput
	// The status of the resource. Valid values: `CHANGE_FAILED`, `CHANGE_INIT`, `EXCEPTION`, `FAILED`, `INIT`, `SUCCESS`.
	Status pulumi.StringPtrInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
	// The Region of the VPC.
	VpcRegionId pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The name of the resource.
	EndpointName string `pulumi:"endpointName"`
	// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
	IpConfigs []EndpointIpConfig `pulumi:"ipConfigs"`
	// The ID of the Security Group.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
	// The Region of the VPC.
	VpcRegionId string `pulumi:"vpcRegionId"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The name of the resource.
	EndpointName pulumi.StringInput
	// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
	IpConfigs EndpointIpConfigArrayInput
	// The ID of the Security Group.
	SecurityGroupId pulumi.StringInput
	// The VPC ID.
	VpcId pulumi.StringInput
	// The Region of the VPC.
	VpcRegionId pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// The name of the resource.
func (o EndpointOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointName }).(pulumi.StringOutput)
}

// The Ip Configs. See the following `Block ipConfigs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
func (o EndpointOutput) IpConfigs() EndpointIpConfigArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointIpConfigArrayOutput { return v.IpConfigs }).(EndpointIpConfigArrayOutput)
}

// The ID of the Security Group.
func (o EndpointOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The status of the resource. Valid values: `CHANGE_FAILED`, `CHANGE_INIT`, `EXCEPTION`, `FAILED`, `INIT`, `SUCCESS`.
func (o EndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The VPC ID.
func (o EndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The Region of the VPC.
func (o EndpointOutput) VpcRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.VpcRegionId }).(pulumi.StringOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
