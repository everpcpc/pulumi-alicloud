// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide RDS cluster instance endpoint connection resources.
// > **NOTE:** Available in 1.203.0+.
//
// ## Block nodeItems
//
// The nodeItems mapping supports the following:
//
// * `nodeId` - (Required) The ID of the node.
// * `weight` - (Required) The weight of the node. Read requests are distributed based on the weight.Valid values: 0 to 100.
//
// ## Import
//
// RDS database endpoint feature can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint example <db_instance_id>:<db_instance_endpoint_id>
//
// ```
type DbInstanceEndpoint struct {
	pulumi.CustomResourceState

	// The internal endpoint.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The IP address of the internal endpoint.
	ConnectionStringPrefix pulumi.StringOutput `pulumi:"connectionStringPrefix"`
	// The user-defined description of the endpoint.
	DbInstanceEndpointDescription pulumi.StringOutput `pulumi:"dbInstanceEndpointDescription"`
	// The Endpoint ID of the instance.
	DbInstanceEndpointId pulumi.StringOutput `pulumi:"dbInstanceEndpointId"`
	// The type of the endpoint.
	DbInstanceEndpointType pulumi.StringOutput `pulumi:"dbInstanceEndpointType"`
	// The ID of the instance.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The type of the IP address.
	IpType pulumi.StringOutput `pulumi:"ipType"`
	// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
	NodeItems DbInstanceEndpointNodeItemArrayOutput `pulumi:"nodeItems"`
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
	Port pulumi.StringOutput `pulumi:"port"`
	// The IP address of the internal endpoint.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The virtual private cloud (VPC) ID of the internal endpoint.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The vSwitch ID of the internal endpoint.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewDbInstanceEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDbInstanceEndpoint(ctx *pulumi.Context,
	name string, args *DbInstanceEndpointArgs, opts ...pulumi.ResourceOption) (*DbInstanceEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionStringPrefix == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionStringPrefix'")
	}
	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.NodeItems == nil {
		return nil, errors.New("invalid value for required argument 'NodeItems'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource DbInstanceEndpoint
	err := ctx.RegisterResource("alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbInstanceEndpoint gets an existing DbInstanceEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbInstanceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbInstanceEndpointState, opts ...pulumi.ResourceOption) (*DbInstanceEndpoint, error) {
	var resource DbInstanceEndpoint
	err := ctx.ReadResource("alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbInstanceEndpoint resources.
type dbInstanceEndpointState struct {
	// The internal endpoint.
	ConnectionString *string `pulumi:"connectionString"`
	// The IP address of the internal endpoint.
	ConnectionStringPrefix *string `pulumi:"connectionStringPrefix"`
	// The user-defined description of the endpoint.
	DbInstanceEndpointDescription *string `pulumi:"dbInstanceEndpointDescription"`
	// The Endpoint ID of the instance.
	DbInstanceEndpointId *string `pulumi:"dbInstanceEndpointId"`
	// The type of the endpoint.
	DbInstanceEndpointType *string `pulumi:"dbInstanceEndpointType"`
	// The ID of the instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The type of the IP address.
	IpType *string `pulumi:"ipType"`
	// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
	NodeItems []DbInstanceEndpointNodeItem `pulumi:"nodeItems"`
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
	Port *string `pulumi:"port"`
	// The IP address of the internal endpoint.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The virtual private cloud (VPC) ID of the internal endpoint.
	VpcId *string `pulumi:"vpcId"`
	// The vSwitch ID of the internal endpoint.
	VswitchId *string `pulumi:"vswitchId"`
}

type DbInstanceEndpointState struct {
	// The internal endpoint.
	ConnectionString pulumi.StringPtrInput
	// The IP address of the internal endpoint.
	ConnectionStringPrefix pulumi.StringPtrInput
	// The user-defined description of the endpoint.
	DbInstanceEndpointDescription pulumi.StringPtrInput
	// The Endpoint ID of the instance.
	DbInstanceEndpointId pulumi.StringPtrInput
	// The type of the endpoint.
	DbInstanceEndpointType pulumi.StringPtrInput
	// The ID of the instance.
	DbInstanceId pulumi.StringPtrInput
	// The type of the IP address.
	IpType pulumi.StringPtrInput
	// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
	NodeItems DbInstanceEndpointNodeItemArrayInput
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
	Port pulumi.StringPtrInput
	// The IP address of the internal endpoint.
	PrivateIpAddress pulumi.StringPtrInput
	// The virtual private cloud (VPC) ID of the internal endpoint.
	VpcId pulumi.StringPtrInput
	// The vSwitch ID of the internal endpoint.
	VswitchId pulumi.StringPtrInput
}

func (DbInstanceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstanceEndpointState)(nil)).Elem()
}

type dbInstanceEndpointArgs struct {
	// The IP address of the internal endpoint.
	ConnectionStringPrefix string `pulumi:"connectionStringPrefix"`
	// The user-defined description of the endpoint.
	DbInstanceEndpointDescription *string `pulumi:"dbInstanceEndpointDescription"`
	// The ID of the instance.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
	NodeItems []DbInstanceEndpointNodeItem `pulumi:"nodeItems"`
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
	Port string `pulumi:"port"`
	// The virtual private cloud (VPC) ID of the internal endpoint.
	VpcId string `pulumi:"vpcId"`
	// The vSwitch ID of the internal endpoint.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a DbInstanceEndpoint resource.
type DbInstanceEndpointArgs struct {
	// The IP address of the internal endpoint.
	ConnectionStringPrefix pulumi.StringInput
	// The user-defined description of the endpoint.
	DbInstanceEndpointDescription pulumi.StringPtrInput
	// The ID of the instance.
	DbInstanceId pulumi.StringInput
	// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
	NodeItems DbInstanceEndpointNodeItemArrayInput
	// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
	Port pulumi.StringInput
	// The virtual private cloud (VPC) ID of the internal endpoint.
	VpcId pulumi.StringInput
	// The vSwitch ID of the internal endpoint.
	VswitchId pulumi.StringInput
}

func (DbInstanceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbInstanceEndpointArgs)(nil)).Elem()
}

type DbInstanceEndpointInput interface {
	pulumi.Input

	ToDbInstanceEndpointOutput() DbInstanceEndpointOutput
	ToDbInstanceEndpointOutputWithContext(ctx context.Context) DbInstanceEndpointOutput
}

func (*DbInstanceEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DbInstanceEndpoint)(nil)).Elem()
}

func (i *DbInstanceEndpoint) ToDbInstanceEndpointOutput() DbInstanceEndpointOutput {
	return i.ToDbInstanceEndpointOutputWithContext(context.Background())
}

func (i *DbInstanceEndpoint) ToDbInstanceEndpointOutputWithContext(ctx context.Context) DbInstanceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstanceEndpointOutput)
}

// DbInstanceEndpointArrayInput is an input type that accepts DbInstanceEndpointArray and DbInstanceEndpointArrayOutput values.
// You can construct a concrete instance of `DbInstanceEndpointArrayInput` via:
//
//	DbInstanceEndpointArray{ DbInstanceEndpointArgs{...} }
type DbInstanceEndpointArrayInput interface {
	pulumi.Input

	ToDbInstanceEndpointArrayOutput() DbInstanceEndpointArrayOutput
	ToDbInstanceEndpointArrayOutputWithContext(context.Context) DbInstanceEndpointArrayOutput
}

type DbInstanceEndpointArray []DbInstanceEndpointInput

func (DbInstanceEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbInstanceEndpoint)(nil)).Elem()
}

func (i DbInstanceEndpointArray) ToDbInstanceEndpointArrayOutput() DbInstanceEndpointArrayOutput {
	return i.ToDbInstanceEndpointArrayOutputWithContext(context.Background())
}

func (i DbInstanceEndpointArray) ToDbInstanceEndpointArrayOutputWithContext(ctx context.Context) DbInstanceEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstanceEndpointArrayOutput)
}

// DbInstanceEndpointMapInput is an input type that accepts DbInstanceEndpointMap and DbInstanceEndpointMapOutput values.
// You can construct a concrete instance of `DbInstanceEndpointMapInput` via:
//
//	DbInstanceEndpointMap{ "key": DbInstanceEndpointArgs{...} }
type DbInstanceEndpointMapInput interface {
	pulumi.Input

	ToDbInstanceEndpointMapOutput() DbInstanceEndpointMapOutput
	ToDbInstanceEndpointMapOutputWithContext(context.Context) DbInstanceEndpointMapOutput
}

type DbInstanceEndpointMap map[string]DbInstanceEndpointInput

func (DbInstanceEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbInstanceEndpoint)(nil)).Elem()
}

func (i DbInstanceEndpointMap) ToDbInstanceEndpointMapOutput() DbInstanceEndpointMapOutput {
	return i.ToDbInstanceEndpointMapOutputWithContext(context.Background())
}

func (i DbInstanceEndpointMap) ToDbInstanceEndpointMapOutputWithContext(ctx context.Context) DbInstanceEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstanceEndpointMapOutput)
}

type DbInstanceEndpointOutput struct{ *pulumi.OutputState }

func (DbInstanceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbInstanceEndpoint)(nil)).Elem()
}

func (o DbInstanceEndpointOutput) ToDbInstanceEndpointOutput() DbInstanceEndpointOutput {
	return o
}

func (o DbInstanceEndpointOutput) ToDbInstanceEndpointOutputWithContext(ctx context.Context) DbInstanceEndpointOutput {
	return o
}

// The internal endpoint.
func (o DbInstanceEndpointOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// The IP address of the internal endpoint.
func (o DbInstanceEndpointOutput) ConnectionStringPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.ConnectionStringPrefix }).(pulumi.StringOutput)
}

// The user-defined description of the endpoint.
func (o DbInstanceEndpointOutput) DbInstanceEndpointDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.DbInstanceEndpointDescription }).(pulumi.StringOutput)
}

// The Endpoint ID of the instance.
func (o DbInstanceEndpointOutput) DbInstanceEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.DbInstanceEndpointId }).(pulumi.StringOutput)
}

// The type of the endpoint.
func (o DbInstanceEndpointOutput) DbInstanceEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.DbInstanceEndpointType }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o DbInstanceEndpointOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The type of the IP address.
func (o DbInstanceEndpointOutput) IpType() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.IpType }).(pulumi.StringOutput)
}

// The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight).
func (o DbInstanceEndpointOutput) NodeItems() DbInstanceEndpointNodeItemArrayOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) DbInstanceEndpointNodeItemArrayOutput { return v.NodeItems }).(DbInstanceEndpointNodeItemArrayOutput)
}

// The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
func (o DbInstanceEndpointOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The IP address of the internal endpoint.
func (o DbInstanceEndpointOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

// The virtual private cloud (VPC) ID of the internal endpoint.
func (o DbInstanceEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The vSwitch ID of the internal endpoint.
func (o DbInstanceEndpointOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbInstanceEndpoint) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type DbInstanceEndpointArrayOutput struct{ *pulumi.OutputState }

func (DbInstanceEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbInstanceEndpoint)(nil)).Elem()
}

func (o DbInstanceEndpointArrayOutput) ToDbInstanceEndpointArrayOutput() DbInstanceEndpointArrayOutput {
	return o
}

func (o DbInstanceEndpointArrayOutput) ToDbInstanceEndpointArrayOutputWithContext(ctx context.Context) DbInstanceEndpointArrayOutput {
	return o
}

func (o DbInstanceEndpointArrayOutput) Index(i pulumi.IntInput) DbInstanceEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DbInstanceEndpoint {
		return vs[0].([]*DbInstanceEndpoint)[vs[1].(int)]
	}).(DbInstanceEndpointOutput)
}

type DbInstanceEndpointMapOutput struct{ *pulumi.OutputState }

func (DbInstanceEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbInstanceEndpoint)(nil)).Elem()
}

func (o DbInstanceEndpointMapOutput) ToDbInstanceEndpointMapOutput() DbInstanceEndpointMapOutput {
	return o
}

func (o DbInstanceEndpointMapOutput) ToDbInstanceEndpointMapOutputWithContext(ctx context.Context) DbInstanceEndpointMapOutput {
	return o
}

func (o DbInstanceEndpointMapOutput) MapIndex(k pulumi.StringInput) DbInstanceEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DbInstanceEndpoint {
		return vs[0].(map[string]*DbInstanceEndpoint)[vs[1].(string)]
	}).(DbInstanceEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstanceEndpointInput)(nil)).Elem(), &DbInstanceEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstanceEndpointArrayInput)(nil)).Elem(), DbInstanceEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstanceEndpointMapInput)(nil)).Elem(), DbInstanceEndpointMap{})
	pulumi.RegisterOutputType(DbInstanceEndpointOutput{})
	pulumi.RegisterOutputType(DbInstanceEndpointArrayOutput{})
	pulumi.RegisterOutputType(DbInstanceEndpointMapOutput{})
}
