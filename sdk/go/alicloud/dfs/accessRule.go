// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DFS Access Rule resource.
//
// For information about DFS Access Rule and how to use it, see [What is Access Rule](https://www.alibabacloud.com/help/doc-detail/207144.htm).
//
// > **NOTE:** Available in v1.140.0+.
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
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "exmaple_name"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		defaultAccessGroup, err := dfs.NewAccessGroup(ctx, "defaultAccessGroup", &dfs.AccessGroupArgs{
// 			NetworkType:     pulumi.String("VPC"),
// 			AccessGroupName: pulumi.String(name),
// 			Description:     pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dfs.NewAccessRule(ctx, "defaultAccessRule", &dfs.AccessRuleArgs{
// 			NetworkSegment: pulumi.String("192.0.2.0/24"),
// 			AccessGroupId:  defaultAccessGroup.ID(),
// 			Description:    pulumi.String(name),
// 			RwAccessType:   pulumi.String("RDWR"),
// 			Priority:       pulumi.Int(10),
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
// DFS Access Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dfs/accessRule:AccessRule example <access_group_id>:<access_rule_id>
// ```
type AccessRule struct {
	pulumi.CustomResourceState

	// The resource ID of Access Group.
	AccessGroupId pulumi.StringOutput `pulumi:"accessGroupId"`
	// The ID of the Access Rule.
	AccessRuleId pulumi.StringOutput `pulumi:"accessRuleId"`
	// The Description of the Access Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The NetworkSegment of the Access Rule.
	NetworkSegment pulumi.StringOutput `pulumi:"networkSegment"`
	// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
	RwAccessType pulumi.StringOutput `pulumi:"rwAccessType"`
}

// NewAccessRule registers a new resource with the given unique name, arguments, and options.
func NewAccessRule(ctx *pulumi.Context,
	name string, args *AccessRuleArgs, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupId'")
	}
	if args.NetworkSegment == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSegment'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.RwAccessType == nil {
		return nil, errors.New("invalid value for required argument 'RwAccessType'")
	}
	var resource AccessRule
	err := ctx.RegisterResource("alicloud:dfs/accessRule:AccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessRule gets an existing AccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessRuleState, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	var resource AccessRule
	err := ctx.ReadResource("alicloud:dfs/accessRule:AccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessRule resources.
type accessRuleState struct {
	// The resource ID of Access Group.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// The ID of the Access Rule.
	AccessRuleId *string `pulumi:"accessRuleId"`
	// The Description of the Access Rule.
	Description *string `pulumi:"description"`
	// The NetworkSegment of the Access Rule.
	NetworkSegment *string `pulumi:"networkSegment"`
	// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
	Priority *int `pulumi:"priority"`
	// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
	RwAccessType *string `pulumi:"rwAccessType"`
}

type AccessRuleState struct {
	// The resource ID of Access Group.
	AccessGroupId pulumi.StringPtrInput
	// The ID of the Access Rule.
	AccessRuleId pulumi.StringPtrInput
	// The Description of the Access Rule.
	Description pulumi.StringPtrInput
	// The NetworkSegment of the Access Rule.
	NetworkSegment pulumi.StringPtrInput
	// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
	Priority pulumi.IntPtrInput
	// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
	RwAccessType pulumi.StringPtrInput
}

func (AccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleState)(nil)).Elem()
}

type accessRuleArgs struct {
	// The resource ID of Access Group.
	AccessGroupId string `pulumi:"accessGroupId"`
	// The Description of the Access Rule.
	Description *string `pulumi:"description"`
	// The NetworkSegment of the Access Rule.
	NetworkSegment string `pulumi:"networkSegment"`
	// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
	Priority int `pulumi:"priority"`
	// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
	RwAccessType string `pulumi:"rwAccessType"`
}

// The set of arguments for constructing a AccessRule resource.
type AccessRuleArgs struct {
	// The resource ID of Access Group.
	AccessGroupId pulumi.StringInput
	// The Description of the Access Rule.
	Description pulumi.StringPtrInput
	// The NetworkSegment of the Access Rule.
	NetworkSegment pulumi.StringInput
	// The Priority of the Access Rule. Valid values: `1` to `100`. **NOTE:** When multiple rules are matched by the same authorized object, the high-priority rule takes effect. `1` is the highest priority.
	Priority pulumi.IntInput
	// The RWAccessType of the Access Rule. Valid values: `RDONLY`, `RDWR`.
	RwAccessType pulumi.StringInput
}

func (AccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleArgs)(nil)).Elem()
}

type AccessRuleInput interface {
	pulumi.Input

	ToAccessRuleOutput() AccessRuleOutput
	ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput
}

func (*AccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRule)(nil))
}

func (i *AccessRule) ToAccessRuleOutput() AccessRuleOutput {
	return i.ToAccessRuleOutputWithContext(context.Background())
}

func (i *AccessRule) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleOutput)
}

func (i *AccessRule) ToAccessRulePtrOutput() AccessRulePtrOutput {
	return i.ToAccessRulePtrOutputWithContext(context.Background())
}

func (i *AccessRule) ToAccessRulePtrOutputWithContext(ctx context.Context) AccessRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRulePtrOutput)
}

type AccessRulePtrInput interface {
	pulumi.Input

	ToAccessRulePtrOutput() AccessRulePtrOutput
	ToAccessRulePtrOutputWithContext(ctx context.Context) AccessRulePtrOutput
}

type accessRulePtrType AccessRuleArgs

func (*accessRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil))
}

func (i *accessRulePtrType) ToAccessRulePtrOutput() AccessRulePtrOutput {
	return i.ToAccessRulePtrOutputWithContext(context.Background())
}

func (i *accessRulePtrType) ToAccessRulePtrOutputWithContext(ctx context.Context) AccessRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRulePtrOutput)
}

// AccessRuleArrayInput is an input type that accepts AccessRuleArray and AccessRuleArrayOutput values.
// You can construct a concrete instance of `AccessRuleArrayInput` via:
//
//          AccessRuleArray{ AccessRuleArgs{...} }
type AccessRuleArrayInput interface {
	pulumi.Input

	ToAccessRuleArrayOutput() AccessRuleArrayOutput
	ToAccessRuleArrayOutputWithContext(context.Context) AccessRuleArrayOutput
}

type AccessRuleArray []AccessRuleInput

func (AccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessRule)(nil)).Elem()
}

func (i AccessRuleArray) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return i.ToAccessRuleArrayOutputWithContext(context.Background())
}

func (i AccessRuleArray) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleArrayOutput)
}

// AccessRuleMapInput is an input type that accepts AccessRuleMap and AccessRuleMapOutput values.
// You can construct a concrete instance of `AccessRuleMapInput` via:
//
//          AccessRuleMap{ "key": AccessRuleArgs{...} }
type AccessRuleMapInput interface {
	pulumi.Input

	ToAccessRuleMapOutput() AccessRuleMapOutput
	ToAccessRuleMapOutputWithContext(context.Context) AccessRuleMapOutput
}

type AccessRuleMap map[string]AccessRuleInput

func (AccessRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessRule)(nil)).Elem()
}

func (i AccessRuleMap) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return i.ToAccessRuleMapOutputWithContext(context.Background())
}

func (i AccessRuleMap) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleMapOutput)
}

type AccessRuleOutput struct{ *pulumi.OutputState }

func (AccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRule)(nil))
}

func (o AccessRuleOutput) ToAccessRuleOutput() AccessRuleOutput {
	return o
}

func (o AccessRuleOutput) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return o
}

func (o AccessRuleOutput) ToAccessRulePtrOutput() AccessRulePtrOutput {
	return o.ToAccessRulePtrOutputWithContext(context.Background())
}

func (o AccessRuleOutput) ToAccessRulePtrOutputWithContext(ctx context.Context) AccessRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRule) *AccessRule {
		return &v
	}).(AccessRulePtrOutput)
}

type AccessRulePtrOutput struct{ *pulumi.OutputState }

func (AccessRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil))
}

func (o AccessRulePtrOutput) ToAccessRulePtrOutput() AccessRulePtrOutput {
	return o
}

func (o AccessRulePtrOutput) ToAccessRulePtrOutputWithContext(ctx context.Context) AccessRulePtrOutput {
	return o
}

func (o AccessRulePtrOutput) Elem() AccessRuleOutput {
	return o.ApplyT(func(v *AccessRule) AccessRule {
		if v != nil {
			return *v
		}
		var ret AccessRule
		return ret
	}).(AccessRuleOutput)
}

type AccessRuleArrayOutput struct{ *pulumi.OutputState }

func (AccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessRule)(nil))
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) Index(i pulumi.IntInput) AccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessRule {
		return vs[0].([]AccessRule)[vs[1].(int)]
	}).(AccessRuleOutput)
}

type AccessRuleMapOutput struct{ *pulumi.OutputState }

func (AccessRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessRule)(nil))
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) MapIndex(k pulumi.StringInput) AccessRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccessRule {
		return vs[0].(map[string]AccessRule)[vs[1].(string)]
	}).(AccessRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleInput)(nil)).Elem(), &AccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRulePtrInput)(nil)).Elem(), &AccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleArrayInput)(nil)).Elem(), AccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleMapInput)(nil)).Elem(), AccessRuleMap{})
	pulumi.RegisterOutputType(AccessRuleOutput{})
	pulumi.RegisterOutputType(AccessRulePtrOutput{})
	pulumi.RegisterOutputType(AccessRuleArrayOutput{})
	pulumi.RegisterOutputType(AccessRuleMapOutput{})
}
