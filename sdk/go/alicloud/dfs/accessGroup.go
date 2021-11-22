// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DFS Access Group resource.
//
// For information about DFS Access Group and how to use it, see [What is Access Group](https://www.alibabacloud.com/help/doc-detail/207144.htm).
//
// > **NOTE:** Available in v1.133.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dfs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dfs.NewAccessGroup(ctx, "example", &dfs.AccessGroupArgs{
// 			AccessGroupName: pulumi.String("example_value"),
// 			NetworkType:     pulumi.String("VPC"),
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
// DFS Access Group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dfs/accessGroup:AccessGroup example <id>
// ```
type AccessGroup struct {
	pulumi.CustomResourceState

	// The Name of Access Group.The length of `accessGroupName` does not exceed 100 bytes.
	AccessGroupName pulumi.StringOutput `pulumi:"accessGroupName"`
	// The Description of Access Group. The length of `description` does not exceed 100 bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The NetworkType of Access Group. Valid values: `VPC`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
}

// NewAccessGroup registers a new resource with the given unique name, arguments, and options.
func NewAccessGroup(ctx *pulumi.Context,
	name string, args *AccessGroupArgs, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupName'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	var resource AccessGroup
	err := ctx.RegisterResource("alicloud:dfs/accessGroup:AccessGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessGroup gets an existing AccessGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessGroupState, opts ...pulumi.ResourceOption) (*AccessGroup, error) {
	var resource AccessGroup
	err := ctx.ReadResource("alicloud:dfs/accessGroup:AccessGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessGroup resources.
type accessGroupState struct {
	// The Name of Access Group.The length of `accessGroupName` does not exceed 100 bytes.
	AccessGroupName *string `pulumi:"accessGroupName"`
	// The Description of Access Group. The length of `description` does not exceed 100 bytes.
	Description *string `pulumi:"description"`
	// The NetworkType of Access Group. Valid values: `VPC`.
	NetworkType *string `pulumi:"networkType"`
}

type AccessGroupState struct {
	// The Name of Access Group.The length of `accessGroupName` does not exceed 100 bytes.
	AccessGroupName pulumi.StringPtrInput
	// The Description of Access Group. The length of `description` does not exceed 100 bytes.
	Description pulumi.StringPtrInput
	// The NetworkType of Access Group. Valid values: `VPC`.
	NetworkType pulumi.StringPtrInput
}

func (AccessGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupState)(nil)).Elem()
}

type accessGroupArgs struct {
	// The Name of Access Group.The length of `accessGroupName` does not exceed 100 bytes.
	AccessGroupName string `pulumi:"accessGroupName"`
	// The Description of Access Group. The length of `description` does not exceed 100 bytes.
	Description *string `pulumi:"description"`
	// The NetworkType of Access Group. Valid values: `VPC`.
	NetworkType string `pulumi:"networkType"`
}

// The set of arguments for constructing a AccessGroup resource.
type AccessGroupArgs struct {
	// The Name of Access Group.The length of `accessGroupName` does not exceed 100 bytes.
	AccessGroupName pulumi.StringInput
	// The Description of Access Group. The length of `description` does not exceed 100 bytes.
	Description pulumi.StringPtrInput
	// The NetworkType of Access Group. Valid values: `VPC`.
	NetworkType pulumi.StringInput
}

func (AccessGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessGroupArgs)(nil)).Elem()
}

type AccessGroupInput interface {
	pulumi.Input

	ToAccessGroupOutput() AccessGroupOutput
	ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput
}

func (*AccessGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessGroup)(nil))
}

func (i *AccessGroup) ToAccessGroupOutput() AccessGroupOutput {
	return i.ToAccessGroupOutputWithContext(context.Background())
}

func (i *AccessGroup) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupOutput)
}

func (i *AccessGroup) ToAccessGroupPtrOutput() AccessGroupPtrOutput {
	return i.ToAccessGroupPtrOutputWithContext(context.Background())
}

func (i *AccessGroup) ToAccessGroupPtrOutputWithContext(ctx context.Context) AccessGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupPtrOutput)
}

type AccessGroupPtrInput interface {
	pulumi.Input

	ToAccessGroupPtrOutput() AccessGroupPtrOutput
	ToAccessGroupPtrOutputWithContext(ctx context.Context) AccessGroupPtrOutput
}

type accessGroupPtrType AccessGroupArgs

func (*accessGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil))
}

func (i *accessGroupPtrType) ToAccessGroupPtrOutput() AccessGroupPtrOutput {
	return i.ToAccessGroupPtrOutputWithContext(context.Background())
}

func (i *accessGroupPtrType) ToAccessGroupPtrOutputWithContext(ctx context.Context) AccessGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupPtrOutput)
}

// AccessGroupArrayInput is an input type that accepts AccessGroupArray and AccessGroupArrayOutput values.
// You can construct a concrete instance of `AccessGroupArrayInput` via:
//
//          AccessGroupArray{ AccessGroupArgs{...} }
type AccessGroupArrayInput interface {
	pulumi.Input

	ToAccessGroupArrayOutput() AccessGroupArrayOutput
	ToAccessGroupArrayOutputWithContext(context.Context) AccessGroupArrayOutput
}

type AccessGroupArray []AccessGroupInput

func (AccessGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupArray) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return i.ToAccessGroupArrayOutputWithContext(context.Background())
}

func (i AccessGroupArray) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupArrayOutput)
}

// AccessGroupMapInput is an input type that accepts AccessGroupMap and AccessGroupMapOutput values.
// You can construct a concrete instance of `AccessGroupMapInput` via:
//
//          AccessGroupMap{ "key": AccessGroupArgs{...} }
type AccessGroupMapInput interface {
	pulumi.Input

	ToAccessGroupMapOutput() AccessGroupMapOutput
	ToAccessGroupMapOutputWithContext(context.Context) AccessGroupMapOutput
}

type AccessGroupMap map[string]AccessGroupInput

func (AccessGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessGroup)(nil)).Elem()
}

func (i AccessGroupMap) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return i.ToAccessGroupMapOutputWithContext(context.Background())
}

func (i AccessGroupMap) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessGroupMapOutput)
}

type AccessGroupOutput struct{ *pulumi.OutputState }

func (AccessGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessGroup)(nil))
}

func (o AccessGroupOutput) ToAccessGroupOutput() AccessGroupOutput {
	return o
}

func (o AccessGroupOutput) ToAccessGroupOutputWithContext(ctx context.Context) AccessGroupOutput {
	return o
}

func (o AccessGroupOutput) ToAccessGroupPtrOutput() AccessGroupPtrOutput {
	return o.ToAccessGroupPtrOutputWithContext(context.Background())
}

func (o AccessGroupOutput) ToAccessGroupPtrOutputWithContext(ctx context.Context) AccessGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessGroup) *AccessGroup {
		return &v
	}).(AccessGroupPtrOutput)
}

type AccessGroupPtrOutput struct{ *pulumi.OutputState }

func (AccessGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessGroup)(nil))
}

func (o AccessGroupPtrOutput) ToAccessGroupPtrOutput() AccessGroupPtrOutput {
	return o
}

func (o AccessGroupPtrOutput) ToAccessGroupPtrOutputWithContext(ctx context.Context) AccessGroupPtrOutput {
	return o
}

func (o AccessGroupPtrOutput) Elem() AccessGroupOutput {
	return o.ApplyT(func(v *AccessGroup) AccessGroup {
		if v != nil {
			return *v
		}
		var ret AccessGroup
		return ret
	}).(AccessGroupOutput)
}

type AccessGroupArrayOutput struct{ *pulumi.OutputState }

func (AccessGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessGroup)(nil))
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutput() AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) ToAccessGroupArrayOutputWithContext(ctx context.Context) AccessGroupArrayOutput {
	return o
}

func (o AccessGroupArrayOutput) Index(i pulumi.IntInput) AccessGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessGroup {
		return vs[0].([]AccessGroup)[vs[1].(int)]
	}).(AccessGroupOutput)
}

type AccessGroupMapOutput struct{ *pulumi.OutputState }

func (AccessGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessGroup)(nil))
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutput() AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) ToAccessGroupMapOutputWithContext(ctx context.Context) AccessGroupMapOutput {
	return o
}

func (o AccessGroupMapOutput) MapIndex(k pulumi.StringInput) AccessGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccessGroup {
		return vs[0].(map[string]AccessGroup)[vs[1].(string)]
	}).(AccessGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupInput)(nil)).Elem(), &AccessGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupPtrInput)(nil)).Elem(), &AccessGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupArrayInput)(nil)).Elem(), AccessGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessGroupMapInput)(nil)).Elem(), AccessGroupMap{})
	pulumi.RegisterOutputType(AccessGroupOutput{})
	pulumi.RegisterOutputType(AccessGroupPtrOutput{})
	pulumi.RegisterOutputType(AccessGroupArrayOutput{})
	pulumi.RegisterOutputType(AccessGroupMapOutput{})
}
