// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Serverless App Engine (SAE) GreyTagRoute resource.
//
// For information about Serverless App Engine (SAE) GreyTagRoute and how to use it, see [What is GreyTagRoute](https://help.aliyun.com/document_detail/97792.html).
//
// > **NOTE:** Available in v1.160.0+.
//
// ## Import
//
// Serverless App Engine (SAE) GreyTagRoute can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sae/greyTagRoute:GreyTagRoute example <id>
//
// ```
type GreyTagRoute struct {
	pulumi.CustomResourceState

	// The ID  of the SAE Application.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The description of GreyTagRoute.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
	DubboRules GreyTagRouteDubboRuleArrayOutput `pulumi:"dubboRules"`
	// The name of GreyTagRoute.
	GreyTagRouteName pulumi.StringOutput `pulumi:"greyTagRouteName"`
	// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
	ScRules GreyTagRouteScRuleArrayOutput `pulumi:"scRules"`
}

// NewGreyTagRoute registers a new resource with the given unique name, arguments, and options.
func NewGreyTagRoute(ctx *pulumi.Context,
	name string, args *GreyTagRouteArgs, opts ...pulumi.ResourceOption) (*GreyTagRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.GreyTagRouteName == nil {
		return nil, errors.New("invalid value for required argument 'GreyTagRouteName'")
	}
	var resource GreyTagRoute
	err := ctx.RegisterResource("alicloud:sae/greyTagRoute:GreyTagRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGreyTagRoute gets an existing GreyTagRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGreyTagRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GreyTagRouteState, opts ...pulumi.ResourceOption) (*GreyTagRoute, error) {
	var resource GreyTagRoute
	err := ctx.ReadResource("alicloud:sae/greyTagRoute:GreyTagRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GreyTagRoute resources.
type greyTagRouteState struct {
	// The ID  of the SAE Application.
	AppId *string `pulumi:"appId"`
	// The description of GreyTagRoute.
	Description *string `pulumi:"description"`
	// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
	DubboRules []GreyTagRouteDubboRule `pulumi:"dubboRules"`
	// The name of GreyTagRoute.
	GreyTagRouteName *string `pulumi:"greyTagRouteName"`
	// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
	ScRules []GreyTagRouteScRule `pulumi:"scRules"`
}

type GreyTagRouteState struct {
	// The ID  of the SAE Application.
	AppId pulumi.StringPtrInput
	// The description of GreyTagRoute.
	Description pulumi.StringPtrInput
	// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
	DubboRules GreyTagRouteDubboRuleArrayInput
	// The name of GreyTagRoute.
	GreyTagRouteName pulumi.StringPtrInput
	// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
	ScRules GreyTagRouteScRuleArrayInput
}

func (GreyTagRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*greyTagRouteState)(nil)).Elem()
}

type greyTagRouteArgs struct {
	// The ID  of the SAE Application.
	AppId string `pulumi:"appId"`
	// The description of GreyTagRoute.
	Description *string `pulumi:"description"`
	// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
	DubboRules []GreyTagRouteDubboRule `pulumi:"dubboRules"`
	// The name of GreyTagRoute.
	GreyTagRouteName string `pulumi:"greyTagRouteName"`
	// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
	ScRules []GreyTagRouteScRule `pulumi:"scRules"`
}

// The set of arguments for constructing a GreyTagRoute resource.
type GreyTagRouteArgs struct {
	// The ID  of the SAE Application.
	AppId pulumi.StringInput
	// The description of GreyTagRoute.
	Description pulumi.StringPtrInput
	// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
	DubboRules GreyTagRouteDubboRuleArrayInput
	// The name of GreyTagRoute.
	GreyTagRouteName pulumi.StringInput
	// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
	ScRules GreyTagRouteScRuleArrayInput
}

func (GreyTagRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*greyTagRouteArgs)(nil)).Elem()
}

type GreyTagRouteInput interface {
	pulumi.Input

	ToGreyTagRouteOutput() GreyTagRouteOutput
	ToGreyTagRouteOutputWithContext(ctx context.Context) GreyTagRouteOutput
}

func (*GreyTagRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GreyTagRoute)(nil)).Elem()
}

func (i *GreyTagRoute) ToGreyTagRouteOutput() GreyTagRouteOutput {
	return i.ToGreyTagRouteOutputWithContext(context.Background())
}

func (i *GreyTagRoute) ToGreyTagRouteOutputWithContext(ctx context.Context) GreyTagRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GreyTagRouteOutput)
}

// GreyTagRouteArrayInput is an input type that accepts GreyTagRouteArray and GreyTagRouteArrayOutput values.
// You can construct a concrete instance of `GreyTagRouteArrayInput` via:
//
//	GreyTagRouteArray{ GreyTagRouteArgs{...} }
type GreyTagRouteArrayInput interface {
	pulumi.Input

	ToGreyTagRouteArrayOutput() GreyTagRouteArrayOutput
	ToGreyTagRouteArrayOutputWithContext(context.Context) GreyTagRouteArrayOutput
}

type GreyTagRouteArray []GreyTagRouteInput

func (GreyTagRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GreyTagRoute)(nil)).Elem()
}

func (i GreyTagRouteArray) ToGreyTagRouteArrayOutput() GreyTagRouteArrayOutput {
	return i.ToGreyTagRouteArrayOutputWithContext(context.Background())
}

func (i GreyTagRouteArray) ToGreyTagRouteArrayOutputWithContext(ctx context.Context) GreyTagRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GreyTagRouteArrayOutput)
}

// GreyTagRouteMapInput is an input type that accepts GreyTagRouteMap and GreyTagRouteMapOutput values.
// You can construct a concrete instance of `GreyTagRouteMapInput` via:
//
//	GreyTagRouteMap{ "key": GreyTagRouteArgs{...} }
type GreyTagRouteMapInput interface {
	pulumi.Input

	ToGreyTagRouteMapOutput() GreyTagRouteMapOutput
	ToGreyTagRouteMapOutputWithContext(context.Context) GreyTagRouteMapOutput
}

type GreyTagRouteMap map[string]GreyTagRouteInput

func (GreyTagRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GreyTagRoute)(nil)).Elem()
}

func (i GreyTagRouteMap) ToGreyTagRouteMapOutput() GreyTagRouteMapOutput {
	return i.ToGreyTagRouteMapOutputWithContext(context.Background())
}

func (i GreyTagRouteMap) ToGreyTagRouteMapOutputWithContext(ctx context.Context) GreyTagRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GreyTagRouteMapOutput)
}

type GreyTagRouteOutput struct{ *pulumi.OutputState }

func (GreyTagRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GreyTagRoute)(nil)).Elem()
}

func (o GreyTagRouteOutput) ToGreyTagRouteOutput() GreyTagRouteOutput {
	return o
}

func (o GreyTagRouteOutput) ToGreyTagRouteOutputWithContext(ctx context.Context) GreyTagRouteOutput {
	return o
}

// The ID  of the SAE Application.
func (o GreyTagRouteOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *GreyTagRoute) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The description of GreyTagRoute.
func (o GreyTagRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GreyTagRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The grayscale rule created for Dubbo Application. The details see Block `dubboRules`.
func (o GreyTagRouteOutput) DubboRules() GreyTagRouteDubboRuleArrayOutput {
	return o.ApplyT(func(v *GreyTagRoute) GreyTagRouteDubboRuleArrayOutput { return v.DubboRules }).(GreyTagRouteDubboRuleArrayOutput)
}

// The name of GreyTagRoute.
func (o GreyTagRouteOutput) GreyTagRouteName() pulumi.StringOutput {
	return o.ApplyT(func(v *GreyTagRoute) pulumi.StringOutput { return v.GreyTagRouteName }).(pulumi.StringOutput)
}

// The grayscale rule created for SpringCloud Application. The details see Block `scRules`.
func (o GreyTagRouteOutput) ScRules() GreyTagRouteScRuleArrayOutput {
	return o.ApplyT(func(v *GreyTagRoute) GreyTagRouteScRuleArrayOutput { return v.ScRules }).(GreyTagRouteScRuleArrayOutput)
}

type GreyTagRouteArrayOutput struct{ *pulumi.OutputState }

func (GreyTagRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GreyTagRoute)(nil)).Elem()
}

func (o GreyTagRouteArrayOutput) ToGreyTagRouteArrayOutput() GreyTagRouteArrayOutput {
	return o
}

func (o GreyTagRouteArrayOutput) ToGreyTagRouteArrayOutputWithContext(ctx context.Context) GreyTagRouteArrayOutput {
	return o
}

func (o GreyTagRouteArrayOutput) Index(i pulumi.IntInput) GreyTagRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GreyTagRoute {
		return vs[0].([]*GreyTagRoute)[vs[1].(int)]
	}).(GreyTagRouteOutput)
}

type GreyTagRouteMapOutput struct{ *pulumi.OutputState }

func (GreyTagRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GreyTagRoute)(nil)).Elem()
}

func (o GreyTagRouteMapOutput) ToGreyTagRouteMapOutput() GreyTagRouteMapOutput {
	return o
}

func (o GreyTagRouteMapOutput) ToGreyTagRouteMapOutputWithContext(ctx context.Context) GreyTagRouteMapOutput {
	return o
}

func (o GreyTagRouteMapOutput) MapIndex(k pulumi.StringInput) GreyTagRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GreyTagRoute {
		return vs[0].(map[string]*GreyTagRoute)[vs[1].(string)]
	}).(GreyTagRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GreyTagRouteInput)(nil)).Elem(), &GreyTagRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GreyTagRouteArrayInput)(nil)).Elem(), GreyTagRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GreyTagRouteMapInput)(nil)).Elem(), GreyTagRouteMap{})
	pulumi.RegisterOutputType(GreyTagRouteOutput{})
	pulumi.RegisterOutputType(GreyTagRouteArrayOutput{})
	pulumi.RegisterOutputType(GreyTagRouteMapOutput{})
}
