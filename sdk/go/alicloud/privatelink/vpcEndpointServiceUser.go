// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Private Link Vpc Endpoint Service User resource.
//
// For information about Private Link Vpc Endpoint Service User and how to use it, see [What is Vpc Endpoint Service User](https://help.aliyun.com/document_detail/183545.html).
//
// > **NOTE:** Available in v1.110.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/privatelink"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := privatelink.NewVpcEndpointServiceUser(ctx, "example", &privatelink.VpcEndpointServiceUserArgs{
// 			ServiceId: pulumi.String("epsrv-gw81c6xxxxxx"),
// 			UserId:    pulumi.String("YourRamUserId"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Private Link Vpc Endpoint Service User can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser example <service_id>:<user_id>
// ```
type VpcEndpointServiceUser struct {
	pulumi.CustomResourceState

	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The Id of Vpc Endpoint Service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The Id of Ram User.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewVpcEndpointServiceUser registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceUser(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceUserArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource VpcEndpointServiceUser
	err := ctx.RegisterResource("alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceUser gets an existing VpcEndpointServiceUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceUserState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceUser, error) {
	var resource VpcEndpointServiceUser
	err := ctx.ReadResource("alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceUser resources.
type vpcEndpointServiceUserState struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The Id of Vpc Endpoint Service.
	ServiceId *string `pulumi:"serviceId"`
	// The Id of Ram User.
	UserId *string `pulumi:"userId"`
}

type VpcEndpointServiceUserState struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The Id of Vpc Endpoint Service.
	ServiceId pulumi.StringPtrInput
	// The Id of Ram User.
	UserId pulumi.StringPtrInput
}

func (VpcEndpointServiceUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceUserState)(nil)).Elem()
}

type vpcEndpointServiceUserArgs struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The Id of Vpc Endpoint Service.
	ServiceId string `pulumi:"serviceId"`
	// The Id of Ram User.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a VpcEndpointServiceUser resource.
type VpcEndpointServiceUserArgs struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The Id of Vpc Endpoint Service.
	ServiceId pulumi.StringInput
	// The Id of Ram User.
	UserId pulumi.StringInput
}

func (VpcEndpointServiceUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceUserArgs)(nil)).Elem()
}

type VpcEndpointServiceUserInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput
	ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput
}

func (*VpcEndpointServiceUser) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServiceUser)(nil))
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput {
	return i.ToVpcEndpointServiceUserOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserOutput)
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserPtrOutput() VpcEndpointServiceUserPtrOutput {
	return i.ToVpcEndpointServiceUserPtrOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceUser) ToVpcEndpointServiceUserPtrOutputWithContext(ctx context.Context) VpcEndpointServiceUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserPtrOutput)
}

type VpcEndpointServiceUserPtrInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserPtrOutput() VpcEndpointServiceUserPtrOutput
	ToVpcEndpointServiceUserPtrOutputWithContext(ctx context.Context) VpcEndpointServiceUserPtrOutput
}

type vpcEndpointServiceUserPtrType VpcEndpointServiceUserArgs

func (*vpcEndpointServiceUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceUser)(nil))
}

func (i *vpcEndpointServiceUserPtrType) ToVpcEndpointServiceUserPtrOutput() VpcEndpointServiceUserPtrOutput {
	return i.ToVpcEndpointServiceUserPtrOutputWithContext(context.Background())
}

func (i *vpcEndpointServiceUserPtrType) ToVpcEndpointServiceUserPtrOutputWithContext(ctx context.Context) VpcEndpointServiceUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserPtrOutput)
}

// VpcEndpointServiceUserArrayInput is an input type that accepts VpcEndpointServiceUserArray and VpcEndpointServiceUserArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceUserArrayInput` via:
//
//          VpcEndpointServiceUserArray{ VpcEndpointServiceUserArgs{...} }
type VpcEndpointServiceUserArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput
	ToVpcEndpointServiceUserArrayOutputWithContext(context.Context) VpcEndpointServiceUserArrayOutput
}

type VpcEndpointServiceUserArray []VpcEndpointServiceUserInput

func (VpcEndpointServiceUserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpcEndpointServiceUser)(nil))
}

func (i VpcEndpointServiceUserArray) ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput {
	return i.ToVpcEndpointServiceUserArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceUserArray) ToVpcEndpointServiceUserArrayOutputWithContext(ctx context.Context) VpcEndpointServiceUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserArrayOutput)
}

// VpcEndpointServiceUserMapInput is an input type that accepts VpcEndpointServiceUserMap and VpcEndpointServiceUserMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceUserMapInput` via:
//
//          VpcEndpointServiceUserMap{ "key": VpcEndpointServiceUserArgs{...} }
type VpcEndpointServiceUserMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput
	ToVpcEndpointServiceUserMapOutputWithContext(context.Context) VpcEndpointServiceUserMapOutput
}

type VpcEndpointServiceUserMap map[string]VpcEndpointServiceUserInput

func (VpcEndpointServiceUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpcEndpointServiceUser)(nil))
}

func (i VpcEndpointServiceUserMap) ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput {
	return i.ToVpcEndpointServiceUserMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceUserMap) ToVpcEndpointServiceUserMapOutputWithContext(ctx context.Context) VpcEndpointServiceUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceUserMapOutput)
}

type VpcEndpointServiceUserOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointServiceUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServiceUser)(nil))
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserOutput() VpcEndpointServiceUserOutput {
	return o
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserOutputWithContext(ctx context.Context) VpcEndpointServiceUserOutput {
	return o
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserPtrOutput() VpcEndpointServiceUserPtrOutput {
	return o.ToVpcEndpointServiceUserPtrOutputWithContext(context.Background())
}

func (o VpcEndpointServiceUserOutput) ToVpcEndpointServiceUserPtrOutputWithContext(ctx context.Context) VpcEndpointServiceUserPtrOutput {
	return o.ApplyT(func(v VpcEndpointServiceUser) *VpcEndpointServiceUser {
		return &v
	}).(VpcEndpointServiceUserPtrOutput)
}

type VpcEndpointServiceUserPtrOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointServiceUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceUser)(nil))
}

func (o VpcEndpointServiceUserPtrOutput) ToVpcEndpointServiceUserPtrOutput() VpcEndpointServiceUserPtrOutput {
	return o
}

func (o VpcEndpointServiceUserPtrOutput) ToVpcEndpointServiceUserPtrOutputWithContext(ctx context.Context) VpcEndpointServiceUserPtrOutput {
	return o
}

type VpcEndpointServiceUserArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointServiceUser)(nil))
}

func (o VpcEndpointServiceUserArrayOutput) ToVpcEndpointServiceUserArrayOutput() VpcEndpointServiceUserArrayOutput {
	return o
}

func (o VpcEndpointServiceUserArrayOutput) ToVpcEndpointServiceUserArrayOutputWithContext(ctx context.Context) VpcEndpointServiceUserArrayOutput {
	return o
}

func (o VpcEndpointServiceUserArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpointServiceUser {
		return vs[0].([]VpcEndpointServiceUser)[vs[1].(int)]
	}).(VpcEndpointServiceUserOutput)
}

type VpcEndpointServiceUserMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcEndpointServiceUser)(nil))
}

func (o VpcEndpointServiceUserMapOutput) ToVpcEndpointServiceUserMapOutput() VpcEndpointServiceUserMapOutput {
	return o
}

func (o VpcEndpointServiceUserMapOutput) ToVpcEndpointServiceUserMapOutputWithContext(ctx context.Context) VpcEndpointServiceUserMapOutput {
	return o
}

func (o VpcEndpointServiceUserMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcEndpointServiceUser {
		return vs[0].(map[string]VpcEndpointServiceUser)[vs[1].(string)]
	}).(VpcEndpointServiceUserOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointServiceUserOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserPtrOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceUserMapOutput{})
}
