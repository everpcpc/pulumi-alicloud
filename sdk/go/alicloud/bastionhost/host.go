// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host Host resource.
//
// For information about Bastion Host Host and how to use it, see [What is Host](https://www.alibabacloud.com/help/en/doc-detail/201330.htm).
//
// > **NOTE:** Available in v1.135.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bastionhost.NewHost(ctx, "example", &bastionhost.HostArgs{
//				ActiveAddressType:  pulumi.String("Private"),
//				HostName:           pulumi.String("example_value"),
//				HostPrivateAddress: pulumi.String("172.16.0.10"),
//				InstanceId:         pulumi.String("bastionhost-cn-tl3xxxxxxx"),
//				OsType:             pulumi.String("Linux"),
//				Source:             pulumi.String("Local"),
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
// Bastion Host Host can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:bastionhost/host:Host example <instance_id>:<host_id>
//
// ```
type Host struct {
	pulumi.CustomResourceState

	// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
	ActiveAddressType pulumi.StringOutput `pulumi:"activeAddressType"`
	// Specify a host of notes, supports up to 500 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The host ID.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// Specify the new create a host name of the supports up to 128 characters.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
	HostPrivateAddress pulumi.StringPtrOutput `pulumi:"hostPrivateAddress"`
	// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
	HostPublicAddress pulumi.StringPtrOutput `pulumi:"hostPublicAddress"`
	// Specify the new create a host where the Bastion host ID of.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The instance region id.
	InstanceRegionId pulumi.StringPtrOutput `pulumi:"instanceRegionId"`
	// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// Specify the new create a host of source. Valid values:
	Source pulumi.StringOutput `pulumi:"source"`
	// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
	SourceInstanceId pulumi.StringPtrOutput `pulumi:"sourceInstanceId"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActiveAddressType == nil {
		return nil, errors.New("invalid value for required argument 'ActiveAddressType'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource Host
	err := ctx.RegisterResource("alicloud:bastionhost/host:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("alicloud:bastionhost/host:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
	// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
	ActiveAddressType *string `pulumi:"activeAddressType"`
	// Specify a host of notes, supports up to 500 characters.
	Comment *string `pulumi:"comment"`
	// The host ID.
	HostId *string `pulumi:"hostId"`
	// Specify the new create a host name of the supports up to 128 characters.
	HostName *string `pulumi:"hostName"`
	// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
	HostPrivateAddress *string `pulumi:"hostPrivateAddress"`
	// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
	HostPublicAddress *string `pulumi:"hostPublicAddress"`
	// Specify the new create a host where the Bastion host ID of.
	InstanceId *string `pulumi:"instanceId"`
	// The instance region id.
	InstanceRegionId *string `pulumi:"instanceRegionId"`
	// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
	OsType *string `pulumi:"osType"`
	// Specify the new create a host of source. Valid values:
	Source *string `pulumi:"source"`
	// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
	SourceInstanceId *string `pulumi:"sourceInstanceId"`
}

type HostState struct {
	// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
	ActiveAddressType pulumi.StringPtrInput
	// Specify a host of notes, supports up to 500 characters.
	Comment pulumi.StringPtrInput
	// The host ID.
	HostId pulumi.StringPtrInput
	// Specify the new create a host name of the supports up to 128 characters.
	HostName pulumi.StringPtrInput
	// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
	HostPrivateAddress pulumi.StringPtrInput
	// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
	HostPublicAddress pulumi.StringPtrInput
	// Specify the new create a host where the Bastion host ID of.
	InstanceId pulumi.StringPtrInput
	// The instance region id.
	InstanceRegionId pulumi.StringPtrInput
	// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
	OsType pulumi.StringPtrInput
	// Specify the new create a host of source. Valid values:
	Source pulumi.StringPtrInput
	// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
	SourceInstanceId pulumi.StringPtrInput
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
	ActiveAddressType string `pulumi:"activeAddressType"`
	// Specify a host of notes, supports up to 500 characters.
	Comment *string `pulumi:"comment"`
	// Specify the new create a host name of the supports up to 128 characters.
	HostName string `pulumi:"hostName"`
	// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
	HostPrivateAddress *string `pulumi:"hostPrivateAddress"`
	// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
	HostPublicAddress *string `pulumi:"hostPublicAddress"`
	// Specify the new create a host where the Bastion host ID of.
	InstanceId string `pulumi:"instanceId"`
	// The instance region id.
	InstanceRegionId *string `pulumi:"instanceRegionId"`
	// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
	OsType string `pulumi:"osType"`
	// Specify the new create a host of source. Valid values:
	Source string `pulumi:"source"`
	// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
	SourceInstanceId *string `pulumi:"sourceInstanceId"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
	ActiveAddressType pulumi.StringInput
	// Specify a host of notes, supports up to 500 characters.
	Comment pulumi.StringPtrInput
	// Specify the new create a host name of the supports up to 128 characters.
	HostName pulumi.StringInput
	// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
	HostPrivateAddress pulumi.StringPtrInput
	// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
	HostPublicAddress pulumi.StringPtrInput
	// Specify the new create a host where the Bastion host ID of.
	InstanceId pulumi.StringInput
	// The instance region id.
	InstanceRegionId pulumi.StringPtrInput
	// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
	OsType pulumi.StringInput
	// Specify the new create a host of source. Valid values:
	Source pulumi.StringInput
	// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
	SourceInstanceId pulumi.StringPtrInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

// HostArrayInput is an input type that accepts HostArray and HostArrayOutput values.
// You can construct a concrete instance of `HostArrayInput` via:
//
//	HostArray{ HostArgs{...} }
type HostArrayInput interface {
	pulumi.Input

	ToHostArrayOutput() HostArrayOutput
	ToHostArrayOutputWithContext(context.Context) HostArrayOutput
}

type HostArray []HostInput

func (HostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Host)(nil)).Elem()
}

func (i HostArray) ToHostArrayOutput() HostArrayOutput {
	return i.ToHostArrayOutputWithContext(context.Background())
}

func (i HostArray) ToHostArrayOutputWithContext(ctx context.Context) HostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostArrayOutput)
}

// HostMapInput is an input type that accepts HostMap and HostMapOutput values.
// You can construct a concrete instance of `HostMapInput` via:
//
//	HostMap{ "key": HostArgs{...} }
type HostMapInput interface {
	pulumi.Input

	ToHostMapOutput() HostMapOutput
	ToHostMapOutputWithContext(context.Context) HostMapOutput
}

type HostMap map[string]HostInput

func (HostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Host)(nil)).Elem()
}

func (i HostMap) ToHostMapOutput() HostMapOutput {
	return i.ToHostMapOutputWithContext(context.Background())
}

func (i HostMap) ToHostMapOutputWithContext(ctx context.Context) HostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostMapOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

// Specify the new create a host of address types. Valid values: `Public`: the IP address of a Public network. `Private`: Private network address.
func (o HostOutput) ActiveAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.ActiveAddressType }).(pulumi.StringOutput)
}

// Specify a host of notes, supports up to 500 characters.
func (o HostOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The host ID.
func (o HostOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

// Specify the new create a host name of the supports up to 128 characters.
func (o HostOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Specify the new create a host of the private network address, it is possible to use the domain name or IP ADDRESS. **NOTE:**  This parameter is required if the `activeAddressType` parameter is set to `Private`.
func (o HostOutput) HostPrivateAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.HostPrivateAddress }).(pulumi.StringPtrOutput)
}

// Specify the new create a host of the IP address of a public network, it is possible to use the domain name or IP ADDRESS.
func (o HostOutput) HostPublicAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.HostPublicAddress }).(pulumi.StringPtrOutput)
}

// Specify the new create a host where the Bastion host ID of.
func (o HostOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The instance region id.
func (o HostOutput) InstanceRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.InstanceRegionId }).(pulumi.StringPtrOutput)
}

// Specify the new create the host's operating system. Valid values: `Linux`,`Windows`.
func (o HostOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// Specify the new create a host of source. Valid values:
func (o HostOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Specify the newly created ECS instance ID or dedicated cluster host ID. **NOTE:** This parameter is required if the `source` parameter is set to `Ecs` or `Rds`.
func (o HostOutput) SourceInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.SourceInstanceId }).(pulumi.StringPtrOutput)
}

type HostArrayOutput struct{ *pulumi.OutputState }

func (HostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Host)(nil)).Elem()
}

func (o HostArrayOutput) ToHostArrayOutput() HostArrayOutput {
	return o
}

func (o HostArrayOutput) ToHostArrayOutputWithContext(ctx context.Context) HostArrayOutput {
	return o
}

func (o HostArrayOutput) Index(i pulumi.IntInput) HostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Host {
		return vs[0].([]*Host)[vs[1].(int)]
	}).(HostOutput)
}

type HostMapOutput struct{ *pulumi.OutputState }

func (HostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Host)(nil)).Elem()
}

func (o HostMapOutput) ToHostMapOutput() HostMapOutput {
	return o
}

func (o HostMapOutput) ToHostMapOutputWithContext(ctx context.Context) HostMapOutput {
	return o
}

func (o HostMapOutput) MapIndex(k pulumi.StringInput) HostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Host {
		return vs[0].(map[string]*Host)[vs[1].(string)]
	}).(HostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostInput)(nil)).Elem(), &Host{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostArrayInput)(nil)).Elem(), HostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostMapInput)(nil)).Elem(), HostMap{})
	pulumi.RegisterOutputType(HostOutput{})
	pulumi.RegisterOutputType(HostArrayOutput{})
	pulumi.RegisterOutputType(HostMapOutput{})
}
