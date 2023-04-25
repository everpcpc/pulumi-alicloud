// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE:** Available in v1.163.0+.
//
// For information about server group server attachment and how to use it, see [Configure a server group server attachment](https://www.alibabacloud.com/help/en/doc-detail/35218.html).
//
// > **NOTE:** Applying this resource may conflict with applying `slb.Listener`,
// and the `slb.Listener` block should use `dependsOn = [alicloud_slb_server_group_server_attachment.xxx]` to avoid it.
//
// ## Import
//
// Load balancer backend server group server attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment example <server_group_id>:<server_id>:<port>
//
// ```
type ServerGroupServerAttachment struct {
	pulumi.CustomResourceState

	// The description of the backend server.
	Description pulumi.StringOutput `pulumi:"description"`
	// The port that is used by the backend server. Valid values: `1` to `65535`.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the server group.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
	// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The type of backend server. Valid values: `ecs`, `eni`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewServerGroupServerAttachment registers a new resource with the given unique name, arguments, and options.
func NewServerGroupServerAttachment(ctx *pulumi.Context,
	name string, args *ServerGroupServerAttachmentArgs, opts ...pulumi.ResourceOption) (*ServerGroupServerAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServerGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupId'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	var resource ServerGroupServerAttachment
	err := ctx.RegisterResource("alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroupServerAttachment gets an existing ServerGroupServerAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroupServerAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupServerAttachmentState, opts ...pulumi.ResourceOption) (*ServerGroupServerAttachment, error) {
	var resource ServerGroupServerAttachment
	err := ctx.ReadResource("alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroupServerAttachment resources.
type serverGroupServerAttachmentState struct {
	// The description of the backend server.
	Description *string `pulumi:"description"`
	// The port that is used by the backend server. Valid values: `1` to `65535`.
	Port *int `pulumi:"port"`
	// The ID of the server group.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	ServerId *string `pulumi:"serverId"`
	// The type of backend server. Valid values: `ecs`, `eni`.
	Type *string `pulumi:"type"`
	// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
	Weight *int `pulumi:"weight"`
}

type ServerGroupServerAttachmentState struct {
	// The description of the backend server.
	Description pulumi.StringPtrInput
	// The port that is used by the backend server. Valid values: `1` to `65535`.
	Port pulumi.IntPtrInput
	// The ID of the server group.
	ServerGroupId pulumi.StringPtrInput
	// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	ServerId pulumi.StringPtrInput
	// The type of backend server. Valid values: `ecs`, `eni`.
	Type pulumi.StringPtrInput
	// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
	Weight pulumi.IntPtrInput
}

func (ServerGroupServerAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupServerAttachmentState)(nil)).Elem()
}

type serverGroupServerAttachmentArgs struct {
	// The description of the backend server.
	Description *string `pulumi:"description"`
	// The port that is used by the backend server. Valid values: `1` to `65535`.
	Port int `pulumi:"port"`
	// The ID of the server group.
	ServerGroupId string `pulumi:"serverGroupId"`
	// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	ServerId string `pulumi:"serverId"`
	// The type of backend server. Valid values: `ecs`, `eni`.
	Type *string `pulumi:"type"`
	// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a ServerGroupServerAttachment resource.
type ServerGroupServerAttachmentArgs struct {
	// The description of the backend server.
	Description pulumi.StringPtrInput
	// The port that is used by the backend server. Valid values: `1` to `65535`.
	Port pulumi.IntInput
	// The ID of the server group.
	ServerGroupId pulumi.StringInput
	// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	ServerId pulumi.StringInput
	// The type of backend server. Valid values: `ecs`, `eni`.
	Type pulumi.StringPtrInput
	// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
	Weight pulumi.IntPtrInput
}

func (ServerGroupServerAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupServerAttachmentArgs)(nil)).Elem()
}

type ServerGroupServerAttachmentInput interface {
	pulumi.Input

	ToServerGroupServerAttachmentOutput() ServerGroupServerAttachmentOutput
	ToServerGroupServerAttachmentOutputWithContext(ctx context.Context) ServerGroupServerAttachmentOutput
}

func (*ServerGroupServerAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupServerAttachment)(nil)).Elem()
}

func (i *ServerGroupServerAttachment) ToServerGroupServerAttachmentOutput() ServerGroupServerAttachmentOutput {
	return i.ToServerGroupServerAttachmentOutputWithContext(context.Background())
}

func (i *ServerGroupServerAttachment) ToServerGroupServerAttachmentOutputWithContext(ctx context.Context) ServerGroupServerAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerAttachmentOutput)
}

// ServerGroupServerAttachmentArrayInput is an input type that accepts ServerGroupServerAttachmentArray and ServerGroupServerAttachmentArrayOutput values.
// You can construct a concrete instance of `ServerGroupServerAttachmentArrayInput` via:
//
//	ServerGroupServerAttachmentArray{ ServerGroupServerAttachmentArgs{...} }
type ServerGroupServerAttachmentArrayInput interface {
	pulumi.Input

	ToServerGroupServerAttachmentArrayOutput() ServerGroupServerAttachmentArrayOutput
	ToServerGroupServerAttachmentArrayOutputWithContext(context.Context) ServerGroupServerAttachmentArrayOutput
}

type ServerGroupServerAttachmentArray []ServerGroupServerAttachmentInput

func (ServerGroupServerAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroupServerAttachment)(nil)).Elem()
}

func (i ServerGroupServerAttachmentArray) ToServerGroupServerAttachmentArrayOutput() ServerGroupServerAttachmentArrayOutput {
	return i.ToServerGroupServerAttachmentArrayOutputWithContext(context.Background())
}

func (i ServerGroupServerAttachmentArray) ToServerGroupServerAttachmentArrayOutputWithContext(ctx context.Context) ServerGroupServerAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerAttachmentArrayOutput)
}

// ServerGroupServerAttachmentMapInput is an input type that accepts ServerGroupServerAttachmentMap and ServerGroupServerAttachmentMapOutput values.
// You can construct a concrete instance of `ServerGroupServerAttachmentMapInput` via:
//
//	ServerGroupServerAttachmentMap{ "key": ServerGroupServerAttachmentArgs{...} }
type ServerGroupServerAttachmentMapInput interface {
	pulumi.Input

	ToServerGroupServerAttachmentMapOutput() ServerGroupServerAttachmentMapOutput
	ToServerGroupServerAttachmentMapOutputWithContext(context.Context) ServerGroupServerAttachmentMapOutput
}

type ServerGroupServerAttachmentMap map[string]ServerGroupServerAttachmentInput

func (ServerGroupServerAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroupServerAttachment)(nil)).Elem()
}

func (i ServerGroupServerAttachmentMap) ToServerGroupServerAttachmentMapOutput() ServerGroupServerAttachmentMapOutput {
	return i.ToServerGroupServerAttachmentMapOutputWithContext(context.Background())
}

func (i ServerGroupServerAttachmentMap) ToServerGroupServerAttachmentMapOutputWithContext(ctx context.Context) ServerGroupServerAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerAttachmentMapOutput)
}

type ServerGroupServerAttachmentOutput struct{ *pulumi.OutputState }

func (ServerGroupServerAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupServerAttachment)(nil)).Elem()
}

func (o ServerGroupServerAttachmentOutput) ToServerGroupServerAttachmentOutput() ServerGroupServerAttachmentOutput {
	return o
}

func (o ServerGroupServerAttachmentOutput) ToServerGroupServerAttachmentOutputWithContext(ctx context.Context) ServerGroupServerAttachmentOutput {
	return o
}

// The description of the backend server.
func (o ServerGroupServerAttachmentOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The port that is used by the backend server. Valid values: `1` to `65535`.
func (o ServerGroupServerAttachmentOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The ID of the server group.
func (o ServerGroupServerAttachmentOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.StringOutput { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
func (o ServerGroupServerAttachmentOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The type of backend server. Valid values: `ecs`, `eni`.
func (o ServerGroupServerAttachmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
func (o ServerGroupServerAttachmentOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerGroupServerAttachment) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type ServerGroupServerAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupServerAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroupServerAttachment)(nil)).Elem()
}

func (o ServerGroupServerAttachmentArrayOutput) ToServerGroupServerAttachmentArrayOutput() ServerGroupServerAttachmentArrayOutput {
	return o
}

func (o ServerGroupServerAttachmentArrayOutput) ToServerGroupServerAttachmentArrayOutputWithContext(ctx context.Context) ServerGroupServerAttachmentArrayOutput {
	return o
}

func (o ServerGroupServerAttachmentArrayOutput) Index(i pulumi.IntInput) ServerGroupServerAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroupServerAttachment {
		return vs[0].([]*ServerGroupServerAttachment)[vs[1].(int)]
	}).(ServerGroupServerAttachmentOutput)
}

type ServerGroupServerAttachmentMapOutput struct{ *pulumi.OutputState }

func (ServerGroupServerAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroupServerAttachment)(nil)).Elem()
}

func (o ServerGroupServerAttachmentMapOutput) ToServerGroupServerAttachmentMapOutput() ServerGroupServerAttachmentMapOutput {
	return o
}

func (o ServerGroupServerAttachmentMapOutput) ToServerGroupServerAttachmentMapOutputWithContext(ctx context.Context) ServerGroupServerAttachmentMapOutput {
	return o
}

func (o ServerGroupServerAttachmentMapOutput) MapIndex(k pulumi.StringInput) ServerGroupServerAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroupServerAttachment {
		return vs[0].(map[string]*ServerGroupServerAttachment)[vs[1].(string)]
	}).(ServerGroupServerAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerAttachmentInput)(nil)).Elem(), &ServerGroupServerAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerAttachmentArrayInput)(nil)).Elem(), ServerGroupServerAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerAttachmentMapInput)(nil)).Elem(), ServerGroupServerAttachmentMap{})
	pulumi.RegisterOutputType(ServerGroupServerAttachmentOutput{})
	pulumi.RegisterOutputType(ServerGroupServerAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupServerAttachmentMapOutput{})
}
