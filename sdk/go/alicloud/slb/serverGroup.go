// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A virtual server group contains several ECS instances. The virtual server group can help you to define multiple listening dimension,
// and to meet the personalized requirements of domain name and URL forwarding.
//
// > **NOTE:** One ECS instance can be added into multiple virtual server groups.
//
// > **NOTE:** One virtual server group can be attached with multiple listeners in one load balancer.
//
// > **NOTE:** One Classic and Internet load balancer, its virtual server group can add Classic and VPC ECS instances.
//
// > **NOTE:** One Classic and Intranet load balancer, its virtual server group can only add Classic ECS instances.
//
// > **NOTE:** One VPC load balancer, its virtual server group can only add the same VPC ECS instances.
//
// For information about server group and how to use it, see [Configure a server group](https://www.alibabacloud.com/help/en/doc-detail/35215.html).
//
// ## Block servers
//
// The servers mapping supports the following:
//
// * `serverIds` - (Required) A list backend server ID (ECS instance ID).
// * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
// * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
//
// ## Import
//
// Load balancer backend server group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/serverGroup:ServerGroup example abc123456
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new virtual server group.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
	//
	// Deprecated: Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
	Servers ServerGroupServerArrayOutput `pulumi:"servers"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	var resource ServerGroup
	err := ctx.RegisterResource("alicloud:slb/serverGroup:ServerGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:slb/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new virtual server group.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
	//
	// Deprecated: Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
	Servers []ServerGroupServer `pulumi:"servers"`
}

type ServerGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new virtual server group.
	LoadBalancerId pulumi.StringPtrInput
	// Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
	//
	// Deprecated: Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
	Servers ServerGroupServerArrayInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new virtual server group.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
	//
	// Deprecated: Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
	Servers []ServerGroupServer `pulumi:"servers"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new virtual server group.
	LoadBalancerId pulumi.StringInput
	// Name of the virtual server group. Our plugin provides a default name: "tf-server-group".
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. **NOTE:** Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as `Block server` follows.
	//
	// Deprecated: Field 'servers' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_slb_server_group_server_attachment'.
	Servers ServerGroupServerArrayInput
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
//          ServerGroupArray{ ServerGroupArgs{...} }
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
//          ServerGroupMap{ "key": ServerGroupArgs{...} }
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
