// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nlb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NLB Server Group resource.
//
// For information about NLB Server Group and how to use it, see [What is Server Group](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createservergroup-nlb).
//
// > **NOTE:** Available in v1.186.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nlb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = nlb.NewServerGroup(ctx, "defaultServerGroup", &nlb.ServerGroupArgs{
//				ResourceGroupId: *pulumi.String(defaultResourceGroups.Ids[0]),
//				ServerGroupName: pulumi.Any(_var.Name),
//				ServerGroupType: pulumi.String("Instance"),
//				VpcId:           *pulumi.String(defaultNetworks.Ids[0]),
//				Scheduler:       pulumi.String("Wrr"),
//				Protocol:        pulumi.String("TCP"),
//				HealthCheck: &nlb.ServerGroupHealthCheckArgs{
//					HealthCheckEnabled:        pulumi.Bool(true),
//					HealthCheckType:           pulumi.String("TCP"),
//					HealthCheckConnectPort:    pulumi.Int(0),
//					HealthyThreshold:          pulumi.Int(2),
//					UnhealthyThreshold:        pulumi.Int(2),
//					HealthCheckConnectTimeout: pulumi.Int(5),
//					HealthCheckInterval:       pulumi.Int(10),
//					HttpCheckMethod:           pulumi.String("GET"),
//					HealthCheckHttpCodes: pulumi.StringArray{
//						pulumi.String("http_2xx"),
//						pulumi.String("http_3xx"),
//						pulumi.String("http_4xx"),
//					},
//				},
//				ConnectionDrain:        pulumi.Bool(true),
//				ConnectionDrainTimeout: pulumi.Int(60),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//				},
//				AddressIpVersion: pulumi.String("Ipv4"),
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
// NLB Server Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:nlb/serverGroup:ServerGroup example <id>
//
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
	AddressIpVersion pulumi.StringOutput `pulumi:"addressIpVersion"`
	// Specifies whether to enable connection draining.
	ConnectionDrain pulumi.BoolOutput `pulumi:"connectionDrain"`
	// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
	ConnectionDrainTimeout pulumi.IntOutput `pulumi:"connectionDrainTimeout"`
	// HealthCheck. See the following `Block healthCheck`.
	HealthCheck ServerGroupHealthCheckOutput `pulumi:"healthCheck"`
	// Indicates whether client address retention is enabled.
	PreserveClientIpEnabled pulumi.BoolOutput `pulumi:"preserveClientIpEnabled"`
	// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The ID of the resource group to which the security group belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The routing algorithm. Valid values:
	Scheduler pulumi.StringOutput `pulumi:"scheduler"`
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName pulumi.StringOutput `pulumi:"serverGroupName"`
	// The type of the server group. Valid values:
	ServerGroupType pulumi.StringOutput `pulumi:"serverGroupType"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The id of the vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HealthCheck == nil {
		return nil, errors.New("invalid value for required argument 'HealthCheck'")
	}
	if args.ServerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource ServerGroup
	err := ctx.RegisterResource("alicloud:nlb/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("alicloud:nlb/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// Specifies whether to enable connection draining.
	ConnectionDrain *bool `pulumi:"connectionDrain"`
	// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
	ConnectionDrainTimeout *int `pulumi:"connectionDrainTimeout"`
	// HealthCheck. See the following `Block healthCheck`.
	HealthCheck *ServerGroupHealthCheck `pulumi:"healthCheck"`
	// Indicates whether client address retention is enabled.
	PreserveClientIpEnabled *bool `pulumi:"preserveClientIpEnabled"`
	// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group to which the security group belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The routing algorithm. Valid values:
	Scheduler *string `pulumi:"scheduler"`
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The type of the server group. Valid values:
	ServerGroupType *string `pulumi:"serverGroupType"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The id of the vpc.
	VpcId *string `pulumi:"vpcId"`
}

type ServerGroupState struct {
	// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
	AddressIpVersion pulumi.StringPtrInput
	// Specifies whether to enable connection draining.
	ConnectionDrain pulumi.BoolPtrInput
	// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
	ConnectionDrainTimeout pulumi.IntPtrInput
	// HealthCheck. See the following `Block healthCheck`.
	HealthCheck ServerGroupHealthCheckPtrInput
	// Indicates whether client address retention is enabled.
	PreserveClientIpEnabled pulumi.BoolPtrInput
	// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group to which the security group belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The routing algorithm. Valid values:
	Scheduler pulumi.StringPtrInput
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName pulumi.StringPtrInput
	// The type of the server group. Valid values:
	ServerGroupType pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The id of the vpc.
	VpcId pulumi.StringPtrInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// Specifies whether to enable connection draining.
	ConnectionDrain *bool `pulumi:"connectionDrain"`
	// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
	ConnectionDrainTimeout *int `pulumi:"connectionDrainTimeout"`
	// HealthCheck. See the following `Block healthCheck`.
	HealthCheck ServerGroupHealthCheck `pulumi:"healthCheck"`
	// Indicates whether client address retention is enabled.
	PreserveClientIpEnabled *bool `pulumi:"preserveClientIpEnabled"`
	// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group to which the security group belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The routing algorithm. Valid values:
	Scheduler *string `pulumi:"scheduler"`
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName string `pulumi:"serverGroupName"`
	// The type of the server group. Valid values:
	ServerGroupType *string `pulumi:"serverGroupType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The id of the vpc.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
	AddressIpVersion pulumi.StringPtrInput
	// Specifies whether to enable connection draining.
	ConnectionDrain pulumi.BoolPtrInput
	// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
	ConnectionDrainTimeout pulumi.IntPtrInput
	// HealthCheck. See the following `Block healthCheck`.
	HealthCheck ServerGroupHealthCheckInput
	// Indicates whether client address retention is enabled.
	PreserveClientIpEnabled pulumi.BoolPtrInput
	// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group to which the security group belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The routing algorithm. Valid values:
	Scheduler pulumi.StringPtrInput
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	ServerGroupName pulumi.StringInput
	// The type of the server group. Valid values:
	ServerGroupType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The id of the vpc.
	VpcId pulumi.StringInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

// ServerGroupArrayInput is an input type that accepts ServerGroupArray and ServerGroupArrayOutput values.
// You can construct a concrete instance of `ServerGroupArrayInput` via:
//
//	ServerGroupArray{ ServerGroupArgs{...} }
type ServerGroupArrayInput interface {
	pulumi.Input

	ToServerGroupArrayOutput() ServerGroupArrayOutput
	ToServerGroupArrayOutputWithContext(context.Context) ServerGroupArrayOutput
}

type ServerGroupArray []ServerGroupInput

func (ServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupArray) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return i.ToServerGroupArrayOutputWithContext(context.Background())
}

func (i ServerGroupArray) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupArrayOutput)
}

// ServerGroupMapInput is an input type that accepts ServerGroupMap and ServerGroupMapOutput values.
// You can construct a concrete instance of `ServerGroupMapInput` via:
//
//	ServerGroupMap{ "key": ServerGroupArgs{...} }
type ServerGroupMapInput interface {
	pulumi.Input

	ToServerGroupMapOutput() ServerGroupMapOutput
	ToServerGroupMapOutputWithContext(context.Context) ServerGroupMapOutput
}

type ServerGroupMap map[string]ServerGroupInput

func (ServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupMap) ToServerGroupMapOutput() ServerGroupMapOutput {
	return i.ToServerGroupMapOutputWithContext(context.Background())
}

func (i ServerGroupMap) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupMapOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

// The protocol version. Valid values: `Ipv4` (default), `DualStack`.
func (o ServerGroupOutput) AddressIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.AddressIpVersion }).(pulumi.StringOutput)
}

// Specifies whether to enable connection draining.
func (o ServerGroupOutput) ConnectionDrain() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolOutput { return v.ConnectionDrain }).(pulumi.BoolOutput)
}

// The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
func (o ServerGroupOutput) ConnectionDrainTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.IntOutput { return v.ConnectionDrainTimeout }).(pulumi.IntOutput)
}

// HealthCheck. See the following `Block healthCheck`.
func (o ServerGroupOutput) HealthCheck() ServerGroupHealthCheckOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroupHealthCheckOutput { return v.HealthCheck }).(ServerGroupHealthCheckOutput)
}

// Indicates whether client address retention is enabled.
func (o ServerGroupOutput) PreserveClientIpEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.BoolOutput { return v.PreserveClientIpEnabled }).(pulumi.BoolOutput)
}

// The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
func (o ServerGroupOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The ID of the resource group to which the security group belongs.
func (o ServerGroupOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The routing algorithm. Valid values:
func (o ServerGroupOutput) Scheduler() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Scheduler }).(pulumi.StringOutput)
}

// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
func (o ServerGroupOutput) ServerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupName }).(pulumi.StringOutput)
}

// The type of the server group. Valid values:
func (o ServerGroupOutput) ServerGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupType }).(pulumi.StringOutput)
}

// The status of the resource.
func (o ServerGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o ServerGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The id of the vpc.
func (o ServerGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].([]*ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].(map[string]*ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}