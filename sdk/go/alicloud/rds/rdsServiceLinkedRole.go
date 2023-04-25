// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RDS Service Linked Role.
//
// For information about RDS Service Linked Role and how to use it, see [What is Service Linked Role.](https://www.alibabacloud.com/help/en/doc-detail/171226.htm).
//
// > **NOTE:** Available in v1.189.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.NewRdsServiceLinkedRole(ctx, "default", &rds.RdsServiceLinkedRoleArgs{
//				ServiceName: pulumi.String("AliyunServiceRoleForRdsPgsqlOnEcs"),
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
// RDS Service Linked Role can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rds/rdsServiceLinkedRole:RdsServiceLinkedRole default <service_name>
//
// ```
type RdsServiceLinkedRole struct {
	pulumi.CustomResourceState

	// The Alibaba Cloud Resource Name (ARN) of the role.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the role.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The name of the role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewRdsServiceLinkedRole registers a new resource with the given unique name, arguments, and options.
func NewRdsServiceLinkedRole(ctx *pulumi.Context,
	name string, args *RdsServiceLinkedRoleArgs, opts ...pulumi.ResourceOption) (*RdsServiceLinkedRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource RdsServiceLinkedRole
	err := ctx.RegisterResource("alicloud:rds/rdsServiceLinkedRole:RdsServiceLinkedRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsServiceLinkedRole gets an existing RdsServiceLinkedRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsServiceLinkedRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsServiceLinkedRoleState, opts ...pulumi.ResourceOption) (*RdsServiceLinkedRole, error) {
	var resource RdsServiceLinkedRole
	err := ctx.ReadResource("alicloud:rds/rdsServiceLinkedRole:RdsServiceLinkedRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsServiceLinkedRole resources.
type rdsServiceLinkedRoleState struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	Arn *string `pulumi:"arn"`
	// The ID of the role.
	RoleId *string `pulumi:"roleId"`
	// The name of the role.
	RoleName *string `pulumi:"roleName"`
	// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
	ServiceName *string `pulumi:"serviceName"`
}

type RdsServiceLinkedRoleState struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	Arn pulumi.StringPtrInput
	// The ID of the role.
	RoleId pulumi.StringPtrInput
	// The name of the role.
	RoleName pulumi.StringPtrInput
	// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
	ServiceName pulumi.StringPtrInput
}

func (RdsServiceLinkedRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsServiceLinkedRoleState)(nil)).Elem()
}

type rdsServiceLinkedRoleArgs struct {
	// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a RdsServiceLinkedRole resource.
type RdsServiceLinkedRoleArgs struct {
	// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
	ServiceName pulumi.StringInput
}

func (RdsServiceLinkedRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsServiceLinkedRoleArgs)(nil)).Elem()
}

type RdsServiceLinkedRoleInput interface {
	pulumi.Input

	ToRdsServiceLinkedRoleOutput() RdsServiceLinkedRoleOutput
	ToRdsServiceLinkedRoleOutputWithContext(ctx context.Context) RdsServiceLinkedRoleOutput
}

func (*RdsServiceLinkedRole) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsServiceLinkedRole)(nil)).Elem()
}

func (i *RdsServiceLinkedRole) ToRdsServiceLinkedRoleOutput() RdsServiceLinkedRoleOutput {
	return i.ToRdsServiceLinkedRoleOutputWithContext(context.Background())
}

func (i *RdsServiceLinkedRole) ToRdsServiceLinkedRoleOutputWithContext(ctx context.Context) RdsServiceLinkedRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsServiceLinkedRoleOutput)
}

// RdsServiceLinkedRoleArrayInput is an input type that accepts RdsServiceLinkedRoleArray and RdsServiceLinkedRoleArrayOutput values.
// You can construct a concrete instance of `RdsServiceLinkedRoleArrayInput` via:
//
//	RdsServiceLinkedRoleArray{ RdsServiceLinkedRoleArgs{...} }
type RdsServiceLinkedRoleArrayInput interface {
	pulumi.Input

	ToRdsServiceLinkedRoleArrayOutput() RdsServiceLinkedRoleArrayOutput
	ToRdsServiceLinkedRoleArrayOutputWithContext(context.Context) RdsServiceLinkedRoleArrayOutput
}

type RdsServiceLinkedRoleArray []RdsServiceLinkedRoleInput

func (RdsServiceLinkedRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsServiceLinkedRole)(nil)).Elem()
}

func (i RdsServiceLinkedRoleArray) ToRdsServiceLinkedRoleArrayOutput() RdsServiceLinkedRoleArrayOutput {
	return i.ToRdsServiceLinkedRoleArrayOutputWithContext(context.Background())
}

func (i RdsServiceLinkedRoleArray) ToRdsServiceLinkedRoleArrayOutputWithContext(ctx context.Context) RdsServiceLinkedRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsServiceLinkedRoleArrayOutput)
}

// RdsServiceLinkedRoleMapInput is an input type that accepts RdsServiceLinkedRoleMap and RdsServiceLinkedRoleMapOutput values.
// You can construct a concrete instance of `RdsServiceLinkedRoleMapInput` via:
//
//	RdsServiceLinkedRoleMap{ "key": RdsServiceLinkedRoleArgs{...} }
type RdsServiceLinkedRoleMapInput interface {
	pulumi.Input

	ToRdsServiceLinkedRoleMapOutput() RdsServiceLinkedRoleMapOutput
	ToRdsServiceLinkedRoleMapOutputWithContext(context.Context) RdsServiceLinkedRoleMapOutput
}

type RdsServiceLinkedRoleMap map[string]RdsServiceLinkedRoleInput

func (RdsServiceLinkedRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsServiceLinkedRole)(nil)).Elem()
}

func (i RdsServiceLinkedRoleMap) ToRdsServiceLinkedRoleMapOutput() RdsServiceLinkedRoleMapOutput {
	return i.ToRdsServiceLinkedRoleMapOutputWithContext(context.Background())
}

func (i RdsServiceLinkedRoleMap) ToRdsServiceLinkedRoleMapOutputWithContext(ctx context.Context) RdsServiceLinkedRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsServiceLinkedRoleMapOutput)
}

type RdsServiceLinkedRoleOutput struct{ *pulumi.OutputState }

func (RdsServiceLinkedRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsServiceLinkedRole)(nil)).Elem()
}

func (o RdsServiceLinkedRoleOutput) ToRdsServiceLinkedRoleOutput() RdsServiceLinkedRoleOutput {
	return o
}

func (o RdsServiceLinkedRoleOutput) ToRdsServiceLinkedRoleOutputWithContext(ctx context.Context) RdsServiceLinkedRoleOutput {
	return o
}

// The Alibaba Cloud Resource Name (ARN) of the role.
func (o RdsServiceLinkedRoleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsServiceLinkedRole) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the role.
func (o RdsServiceLinkedRoleOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsServiceLinkedRole) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// The name of the role.
func (o RdsServiceLinkedRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsServiceLinkedRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// The product name for SLR. RDS can automatically create the following service-linked roles: `AliyunServiceRoleForRdsPgsqlOnEcs`, `AliyunServiceRoleForRDSProxyOnEcs`.
func (o RdsServiceLinkedRoleOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsServiceLinkedRole) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type RdsServiceLinkedRoleArrayOutput struct{ *pulumi.OutputState }

func (RdsServiceLinkedRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsServiceLinkedRole)(nil)).Elem()
}

func (o RdsServiceLinkedRoleArrayOutput) ToRdsServiceLinkedRoleArrayOutput() RdsServiceLinkedRoleArrayOutput {
	return o
}

func (o RdsServiceLinkedRoleArrayOutput) ToRdsServiceLinkedRoleArrayOutputWithContext(ctx context.Context) RdsServiceLinkedRoleArrayOutput {
	return o
}

func (o RdsServiceLinkedRoleArrayOutput) Index(i pulumi.IntInput) RdsServiceLinkedRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsServiceLinkedRole {
		return vs[0].([]*RdsServiceLinkedRole)[vs[1].(int)]
	}).(RdsServiceLinkedRoleOutput)
}

type RdsServiceLinkedRoleMapOutput struct{ *pulumi.OutputState }

func (RdsServiceLinkedRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsServiceLinkedRole)(nil)).Elem()
}

func (o RdsServiceLinkedRoleMapOutput) ToRdsServiceLinkedRoleMapOutput() RdsServiceLinkedRoleMapOutput {
	return o
}

func (o RdsServiceLinkedRoleMapOutput) ToRdsServiceLinkedRoleMapOutputWithContext(ctx context.Context) RdsServiceLinkedRoleMapOutput {
	return o
}

func (o RdsServiceLinkedRoleMapOutput) MapIndex(k pulumi.StringInput) RdsServiceLinkedRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsServiceLinkedRole {
		return vs[0].(map[string]*RdsServiceLinkedRole)[vs[1].(string)]
	}).(RdsServiceLinkedRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsServiceLinkedRoleInput)(nil)).Elem(), &RdsServiceLinkedRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsServiceLinkedRoleArrayInput)(nil)).Elem(), RdsServiceLinkedRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsServiceLinkedRoleMapInput)(nil)).Elem(), RdsServiceLinkedRoleMap{})
	pulumi.RegisterOutputType(RdsServiceLinkedRoleOutput{})
	pulumi.RegisterOutputType(RdsServiceLinkedRoleArrayOutput{})
	pulumi.RegisterOutputType(RdsServiceLinkedRoleMapOutput{})
}
