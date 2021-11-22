// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ALB Server Group resource.
//
// For information about ALB Server Group and how to use it,
// see [What is Server Group](https://www.alibabacloud.com/help/doc-detail/213627.htm).
//
// > **NOTE:** Available in v1.131.0+.
//
// ## Import
//
// ALB Server Group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:alb/serverGroup:ServerGroup example <id>
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The configuration of health checks.
	HealthCheckConfig ServerGroupHealthCheckConfigPtrOutput `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTPS `, `HTTP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`.
	Scheduler pulumi.StringOutput `pulumi:"scheduler"`
	// The name of the resource.
	ServerGroupName pulumi.StringPtrOutput `pulumi:"serverGroupName"`
	// The backend server.
	Servers ServerGroupServerArrayOutput `pulumi:"servers"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The configuration of the sticky session.
	StickySessionConfig ServerGroupStickySessionConfigPtrOutput `pulumi:"stickySessionConfig"`
	Tags                pulumi.MapOutput                        `pulumi:"tags"`
	// The ID of the VPC that you want to access.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		args = &ServerGroupArgs{}
	}

	var resource ServerGroup
	err := ctx.RegisterResource("alicloud:alb/serverGroup:ServerGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:alb/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The configuration of health checks.
	HealthCheckConfig *ServerGroupHealthCheckConfig `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTPS `, `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`.
	Scheduler *string `pulumi:"scheduler"`
	// The name of the resource.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The backend server.
	Servers []ServerGroupServer `pulumi:"servers"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The configuration of the sticky session.
	StickySessionConfig *ServerGroupStickySessionConfig `pulumi:"stickySessionConfig"`
	Tags                map[string]interface{}          `pulumi:"tags"`
	// The ID of the VPC that you want to access.
	VpcId *string `pulumi:"vpcId"`
}

type ServerGroupState struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The configuration of health checks.
	HealthCheckConfig ServerGroupHealthCheckConfigPtrInput
	// The server protocol. Valid values: `  HTTPS `, `HTTP`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`.
	Scheduler pulumi.StringPtrInput
	// The name of the resource.
	ServerGroupName pulumi.StringPtrInput
	// The backend server.
	Servers ServerGroupServerArrayInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The configuration of the sticky session.
	StickySessionConfig ServerGroupStickySessionConfigPtrInput
	Tags                pulumi.MapInput
	// The ID of the VPC that you want to access.
	VpcId pulumi.StringPtrInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The configuration of health checks.
	HealthCheckConfig *ServerGroupHealthCheckConfig `pulumi:"healthCheckConfig"`
	// The server protocol. Valid values: `  HTTPS `, `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`.
	Scheduler *string `pulumi:"scheduler"`
	// The name of the resource.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The backend server.
	Servers []ServerGroupServer `pulumi:"servers"`
	// The configuration of the sticky session.
	StickySessionConfig *ServerGroupStickySessionConfig `pulumi:"stickySessionConfig"`
	Tags                map[string]interface{}          `pulumi:"tags"`
	// The ID of the VPC that you want to access.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The configuration of health checks.
	HealthCheckConfig ServerGroupHealthCheckConfigPtrInput
	// The server protocol. Valid values: `  HTTPS `, `HTTP`.
	Protocol pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The scheduling algorithm. Valid values: `  Sch `, `  Wlc `, `Wrr`.
	Scheduler pulumi.StringPtrInput
	// The name of the resource.
	ServerGroupName pulumi.StringPtrInput
	// The backend server.
	Servers ServerGroupServerArrayInput
	// The configuration of the sticky session.
	StickySessionConfig ServerGroupStickySessionConfigPtrInput
	Tags                pulumi.MapInput
	// The ID of the VPC that you want to access.
	VpcId pulumi.StringPtrInput
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
	return reflect.TypeOf((*ServerGroup)(nil))
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

func (i *ServerGroup) ToServerGroupPtrOutput() ServerGroupPtrOutput {
	return i.ToServerGroupPtrOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupPtrOutputWithContext(ctx context.Context) ServerGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPtrOutput)
}

type ServerGroupPtrInput interface {
	pulumi.Input

	ToServerGroupPtrOutput() ServerGroupPtrOutput
	ToServerGroupPtrOutputWithContext(ctx context.Context) ServerGroupPtrOutput
}

type serverGroupPtrType ServerGroupArgs

func (*serverGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil))
}

func (i *serverGroupPtrType) ToServerGroupPtrOutput() ServerGroupPtrOutput {
	return i.ToServerGroupPtrOutputWithContext(context.Background())
}

func (i *serverGroupPtrType) ToServerGroupPtrOutputWithContext(ctx context.Context) ServerGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPtrOutput)
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
	return reflect.TypeOf((*ServerGroup)(nil))
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupPtrOutput() ServerGroupPtrOutput {
	return o.ToServerGroupPtrOutputWithContext(context.Background())
}

func (o ServerGroupOutput) ToServerGroupPtrOutputWithContext(ctx context.Context) ServerGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerGroup) *ServerGroup {
		return &v
	}).(ServerGroupPtrOutput)
}

type ServerGroupPtrOutput struct{ *pulumi.OutputState }

func (ServerGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil))
}

func (o ServerGroupPtrOutput) ToServerGroupPtrOutput() ServerGroupPtrOutput {
	return o
}

func (o ServerGroupPtrOutput) ToServerGroupPtrOutputWithContext(ctx context.Context) ServerGroupPtrOutput {
	return o
}

func (o ServerGroupPtrOutput) Elem() ServerGroupOutput {
	return o.ApplyT(func(v *ServerGroup) ServerGroup {
		if v != nil {
			return *v
		}
		var ret ServerGroup
		return ret
	}).(ServerGroupOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerGroup)(nil))
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerGroup {
		return vs[0].([]ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServerGroup)(nil))
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServerGroup {
		return vs[0].(map[string]ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupPtrInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupPtrOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}
