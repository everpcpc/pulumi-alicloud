// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CEN transit router resource that associate the transitRouter with the CEN instance.[What is Cen Transit Router](https://help.aliyun.com/document_detail/261169.html)
//
// > **NOTE:** Available in 1.126.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-testAccCenTransitRouter"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String(name),
//				Description:     pulumi.String("terraform01"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouter(ctx, "defaultTransitRouter", &cen.TransitRouterArgs{
//				TransitRouterName: pulumi.String(name),
//				CenId:             defaultInstance.ID(),
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
// CEN instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/transitRouter:TransitRouter default cen-*****:tr-*******
//
// ```
type TransitRouter struct {
	pulumi.CustomResourceState

	// The ID of the CEN.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The associating status of the Transit Router.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
	SupportMulticast pulumi.BoolPtrOutput `pulumi:"supportMulticast"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The description of the transit router.
	TransitRouterDescription pulumi.StringPtrOutput `pulumi:"transitRouterDescription"`
	// The transit router id of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
	// The name of the transit router.
	TransitRouterName pulumi.StringPtrOutput `pulumi:"transitRouterName"`
	// The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTransitRouter registers a new resource with the given unique name, arguments, and options.
func NewTransitRouter(ctx *pulumi.Context,
	name string, args *TransitRouterArgs, opts ...pulumi.ResourceOption) (*TransitRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	var resource TransitRouter
	err := ctx.RegisterResource("alicloud:cen/transitRouter:TransitRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouter gets an existing TransitRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterState, opts ...pulumi.ResourceOption) (*TransitRouter, error) {
	var resource TransitRouter
	err := ctx.ReadResource("alicloud:cen/transitRouter:TransitRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouter resources.
type transitRouterState struct {
	// The ID of the CEN.
	CenId *string `pulumi:"cenId"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The associating status of the Transit Router.
	Status *string `pulumi:"status"`
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
	SupportMulticast *bool `pulumi:"supportMulticast"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The description of the transit router.
	TransitRouterDescription *string `pulumi:"transitRouterDescription"`
	// The transit router id of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
	// The name of the transit router.
	TransitRouterName *string `pulumi:"transitRouterName"`
	// The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
	Type *string `pulumi:"type"`
}

type TransitRouterState struct {
	// The ID of the CEN.
	CenId pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The associating status of the Transit Router.
	Status pulumi.StringPtrInput
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
	SupportMulticast pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The description of the transit router.
	TransitRouterDescription pulumi.StringPtrInput
	// The transit router id of the transit router.
	TransitRouterId pulumi.StringPtrInput
	// The name of the transit router.
	TransitRouterName pulumi.StringPtrInput
	// The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
	Type pulumi.StringPtrInput
}

func (TransitRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterState)(nil)).Elem()
}

type transitRouterArgs struct {
	// The ID of the CEN.
	CenId string `pulumi:"cenId"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
	SupportMulticast *bool `pulumi:"supportMulticast"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The description of the transit router.
	TransitRouterDescription *string `pulumi:"transitRouterDescription"`
	// The name of the transit router.
	TransitRouterName *string `pulumi:"transitRouterName"`
}

// The set of arguments for constructing a TransitRouter resource.
type TransitRouterArgs struct {
	// The ID of the CEN.
	CenId pulumi.StringInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
	SupportMulticast pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The description of the transit router.
	TransitRouterDescription pulumi.StringPtrInput
	// The name of the transit router.
	TransitRouterName pulumi.StringPtrInput
}

func (TransitRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterArgs)(nil)).Elem()
}

type TransitRouterInput interface {
	pulumi.Input

	ToTransitRouterOutput() TransitRouterOutput
	ToTransitRouterOutputWithContext(ctx context.Context) TransitRouterOutput
}

func (*TransitRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouter)(nil)).Elem()
}

func (i *TransitRouter) ToTransitRouterOutput() TransitRouterOutput {
	return i.ToTransitRouterOutputWithContext(context.Background())
}

func (i *TransitRouter) ToTransitRouterOutputWithContext(ctx context.Context) TransitRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterOutput)
}

// TransitRouterArrayInput is an input type that accepts TransitRouterArray and TransitRouterArrayOutput values.
// You can construct a concrete instance of `TransitRouterArrayInput` via:
//
//	TransitRouterArray{ TransitRouterArgs{...} }
type TransitRouterArrayInput interface {
	pulumi.Input

	ToTransitRouterArrayOutput() TransitRouterArrayOutput
	ToTransitRouterArrayOutputWithContext(context.Context) TransitRouterArrayOutput
}

type TransitRouterArray []TransitRouterInput

func (TransitRouterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouter)(nil)).Elem()
}

func (i TransitRouterArray) ToTransitRouterArrayOutput() TransitRouterArrayOutput {
	return i.ToTransitRouterArrayOutputWithContext(context.Background())
}

func (i TransitRouterArray) ToTransitRouterArrayOutputWithContext(ctx context.Context) TransitRouterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterArrayOutput)
}

// TransitRouterMapInput is an input type that accepts TransitRouterMap and TransitRouterMapOutput values.
// You can construct a concrete instance of `TransitRouterMapInput` via:
//
//	TransitRouterMap{ "key": TransitRouterArgs{...} }
type TransitRouterMapInput interface {
	pulumi.Input

	ToTransitRouterMapOutput() TransitRouterMapOutput
	ToTransitRouterMapOutputWithContext(context.Context) TransitRouterMapOutput
}

type TransitRouterMap map[string]TransitRouterInput

func (TransitRouterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouter)(nil)).Elem()
}

func (i TransitRouterMap) ToTransitRouterMapOutput() TransitRouterMapOutput {
	return i.ToTransitRouterMapOutputWithContext(context.Background())
}

func (i TransitRouterMap) ToTransitRouterMapOutputWithContext(ctx context.Context) TransitRouterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMapOutput)
}

type TransitRouterOutput struct{ *pulumi.OutputState }

func (TransitRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouter)(nil)).Elem()
}

func (o TransitRouterOutput) ToTransitRouterOutput() TransitRouterOutput {
	return o
}

func (o TransitRouterOutput) ToTransitRouterOutputWithContext(ctx context.Context) TransitRouterOutput {
	return o
}

// The ID of the CEN.
func (o TransitRouterOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The dry run.
func (o TransitRouterOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The associating status of the Transit Router.
func (o TransitRouterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values: `false`, `true`. Default Value: `false`. The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-listtransitrouteravailableresource) to query the regions that support multicast.
func (o TransitRouterOutput) SupportMulticast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.BoolPtrOutput { return v.SupportMulticast }).(pulumi.BoolPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o TransitRouterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The description of the transit router.
func (o TransitRouterOutput) TransitRouterDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringPtrOutput { return v.TransitRouterDescription }).(pulumi.StringPtrOutput)
}

// The transit router id of the transit router.
func (o TransitRouterOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

// The name of the transit router.
func (o TransitRouterOutput) TransitRouterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringPtrOutput { return v.TransitRouterName }).(pulumi.StringPtrOutput)
}

// The Type of the Transit Router. Valid values: `Enterprise`, `Basic`.
func (o TransitRouterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type TransitRouterArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouter)(nil)).Elem()
}

func (o TransitRouterArrayOutput) ToTransitRouterArrayOutput() TransitRouterArrayOutput {
	return o
}

func (o TransitRouterArrayOutput) ToTransitRouterArrayOutputWithContext(ctx context.Context) TransitRouterArrayOutput {
	return o
}

func (o TransitRouterArrayOutput) Index(i pulumi.IntInput) TransitRouterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouter {
		return vs[0].([]*TransitRouter)[vs[1].(int)]
	}).(TransitRouterOutput)
}

type TransitRouterMapOutput struct{ *pulumi.OutputState }

func (TransitRouterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouter)(nil)).Elem()
}

func (o TransitRouterMapOutput) ToTransitRouterMapOutput() TransitRouterMapOutput {
	return o
}

func (o TransitRouterMapOutput) ToTransitRouterMapOutputWithContext(ctx context.Context) TransitRouterMapOutput {
	return o
}

func (o TransitRouterMapOutput) MapIndex(k pulumi.StringInput) TransitRouterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouter {
		return vs[0].(map[string]*TransitRouter)[vs[1].(string)]
	}).(TransitRouterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterInput)(nil)).Elem(), &TransitRouter{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterArrayInput)(nil)).Elem(), TransitRouterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMapInput)(nil)).Elem(), TransitRouterMap{})
	pulumi.RegisterOutputType(TransitRouterOutput{})
	pulumi.RegisterOutputType(TransitRouterArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterMapOutput{})
}
