// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation is provided to scale out an EDAS application.
//
// > **NOTE:** Available in 1.82.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/edas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := edas.NewApplicationScale(ctx, "default", &edas.ApplicationScaleArgs{
//				AppId:       pulumi.Any(_var.App_id),
//				DeployGroup: pulumi.Any(_var.Deploy_group),
//				EcuInfos:    pulumi.Any(_var.Ecu_info),
//				ForceStatus: pulumi.Any(_var.Force_status),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ApplicationScale struct {
	pulumi.CustomResourceState

	// The ID of the application that you want to deploy.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The ID of the instance group to which you want to add ECS instances to scale out the application.
	DeployGroup pulumi.StringOutput `pulumi:"deployGroup"`
	// The ecc information of the resource supplied above. The value is formulated as `<ecc1,ecc2>`.
	EccInfo pulumi.StringOutput `pulumi:"eccInfo"`
	// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayOutput `pulumi:"ecuInfos"`
	// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
	ForceStatus pulumi.BoolPtrOutput `pulumi:"forceStatus"`
}

// NewApplicationScale registers a new resource with the given unique name, arguments, and options.
func NewApplicationScale(ctx *pulumi.Context,
	name string, args *ApplicationScaleArgs, opts ...pulumi.ResourceOption) (*ApplicationScale, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.DeployGroup == nil {
		return nil, errors.New("invalid value for required argument 'DeployGroup'")
	}
	if args.EcuInfos == nil {
		return nil, errors.New("invalid value for required argument 'EcuInfos'")
	}
	var resource ApplicationScale
	err := ctx.RegisterResource("alicloud:edas/applicationScale:ApplicationScale", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationScale gets an existing ApplicationScale resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationScale(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationScaleState, opts ...pulumi.ResourceOption) (*ApplicationScale, error) {
	var resource ApplicationScale
	err := ctx.ReadResource("alicloud:edas/applicationScale:ApplicationScale", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationScale resources.
type applicationScaleState struct {
	// The ID of the application that you want to deploy.
	AppId *string `pulumi:"appId"`
	// The ID of the instance group to which you want to add ECS instances to scale out the application.
	DeployGroup *string `pulumi:"deployGroup"`
	// The ecc information of the resource supplied above. The value is formulated as `<ecc1,ecc2>`.
	EccInfo *string `pulumi:"eccInfo"`
	// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
	ForceStatus *bool `pulumi:"forceStatus"`
}

type ApplicationScaleState struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringPtrInput
	// The ID of the instance group to which you want to add ECS instances to scale out the application.
	DeployGroup pulumi.StringPtrInput
	// The ecc information of the resource supplied above. The value is formulated as `<ecc1,ecc2>`.
	EccInfo pulumi.StringPtrInput
	// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
	ForceStatus pulumi.BoolPtrInput
}

func (ApplicationScaleState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationScaleState)(nil)).Elem()
}

type applicationScaleArgs struct {
	// The ID of the application that you want to deploy.
	AppId string `pulumi:"appId"`
	// The ID of the instance group to which you want to add ECS instances to scale out the application.
	DeployGroup string `pulumi:"deployGroup"`
	// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
	ForceStatus *bool `pulumi:"forceStatus"`
}

// The set of arguments for constructing a ApplicationScale resource.
type ApplicationScaleArgs struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringInput
	// The ID of the instance group to which you want to add ECS instances to scale out the application.
	DeployGroup pulumi.StringInput
	// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
	ForceStatus pulumi.BoolPtrInput
}

func (ApplicationScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationScaleArgs)(nil)).Elem()
}

type ApplicationScaleInput interface {
	pulumi.Input

	ToApplicationScaleOutput() ApplicationScaleOutput
	ToApplicationScaleOutputWithContext(ctx context.Context) ApplicationScaleOutput
}

func (*ApplicationScale) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationScale)(nil)).Elem()
}

func (i *ApplicationScale) ToApplicationScaleOutput() ApplicationScaleOutput {
	return i.ToApplicationScaleOutputWithContext(context.Background())
}

func (i *ApplicationScale) ToApplicationScaleOutputWithContext(ctx context.Context) ApplicationScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScaleOutput)
}

// ApplicationScaleArrayInput is an input type that accepts ApplicationScaleArray and ApplicationScaleArrayOutput values.
// You can construct a concrete instance of `ApplicationScaleArrayInput` via:
//
//	ApplicationScaleArray{ ApplicationScaleArgs{...} }
type ApplicationScaleArrayInput interface {
	pulumi.Input

	ToApplicationScaleArrayOutput() ApplicationScaleArrayOutput
	ToApplicationScaleArrayOutputWithContext(context.Context) ApplicationScaleArrayOutput
}

type ApplicationScaleArray []ApplicationScaleInput

func (ApplicationScaleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationScale)(nil)).Elem()
}

func (i ApplicationScaleArray) ToApplicationScaleArrayOutput() ApplicationScaleArrayOutput {
	return i.ToApplicationScaleArrayOutputWithContext(context.Background())
}

func (i ApplicationScaleArray) ToApplicationScaleArrayOutputWithContext(ctx context.Context) ApplicationScaleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScaleArrayOutput)
}

// ApplicationScaleMapInput is an input type that accepts ApplicationScaleMap and ApplicationScaleMapOutput values.
// You can construct a concrete instance of `ApplicationScaleMapInput` via:
//
//	ApplicationScaleMap{ "key": ApplicationScaleArgs{...} }
type ApplicationScaleMapInput interface {
	pulumi.Input

	ToApplicationScaleMapOutput() ApplicationScaleMapOutput
	ToApplicationScaleMapOutputWithContext(context.Context) ApplicationScaleMapOutput
}

type ApplicationScaleMap map[string]ApplicationScaleInput

func (ApplicationScaleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationScale)(nil)).Elem()
}

func (i ApplicationScaleMap) ToApplicationScaleMapOutput() ApplicationScaleMapOutput {
	return i.ToApplicationScaleMapOutputWithContext(context.Background())
}

func (i ApplicationScaleMap) ToApplicationScaleMapOutputWithContext(ctx context.Context) ApplicationScaleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScaleMapOutput)
}

type ApplicationScaleOutput struct{ *pulumi.OutputState }

func (ApplicationScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationScale)(nil)).Elem()
}

func (o ApplicationScaleOutput) ToApplicationScaleOutput() ApplicationScaleOutput {
	return o
}

func (o ApplicationScaleOutput) ToApplicationScaleOutputWithContext(ctx context.Context) ApplicationScaleOutput {
	return o
}

// The ID of the application that you want to deploy.
func (o ApplicationScaleOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationScale) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The ID of the instance group to which you want to add ECS instances to scale out the application.
func (o ApplicationScaleOutput) DeployGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationScale) pulumi.StringOutput { return v.DeployGroup }).(pulumi.StringOutput)
}

// The ecc information of the resource supplied above. The value is formulated as `<ecc1,ecc2>`.
func (o ApplicationScaleOutput) EccInfo() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationScale) pulumi.StringOutput { return v.EccInfo }).(pulumi.StringOutput)
}

// The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
func (o ApplicationScaleOutput) EcuInfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationScale) pulumi.StringArrayOutput { return v.EcuInfos }).(pulumi.StringArrayOutput)
}

// This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified.
func (o ApplicationScaleOutput) ForceStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationScale) pulumi.BoolPtrOutput { return v.ForceStatus }).(pulumi.BoolPtrOutput)
}

type ApplicationScaleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationScaleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationScale)(nil)).Elem()
}

func (o ApplicationScaleArrayOutput) ToApplicationScaleArrayOutput() ApplicationScaleArrayOutput {
	return o
}

func (o ApplicationScaleArrayOutput) ToApplicationScaleArrayOutputWithContext(ctx context.Context) ApplicationScaleArrayOutput {
	return o
}

func (o ApplicationScaleArrayOutput) Index(i pulumi.IntInput) ApplicationScaleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationScale {
		return vs[0].([]*ApplicationScale)[vs[1].(int)]
	}).(ApplicationScaleOutput)
}

type ApplicationScaleMapOutput struct{ *pulumi.OutputState }

func (ApplicationScaleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationScale)(nil)).Elem()
}

func (o ApplicationScaleMapOutput) ToApplicationScaleMapOutput() ApplicationScaleMapOutput {
	return o
}

func (o ApplicationScaleMapOutput) ToApplicationScaleMapOutputWithContext(ctx context.Context) ApplicationScaleMapOutput {
	return o
}

func (o ApplicationScaleMapOutput) MapIndex(k pulumi.StringInput) ApplicationScaleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationScale {
		return vs[0].(map[string]*ApplicationScale)[vs[1].(string)]
	}).(ApplicationScaleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationScaleInput)(nil)).Elem(), &ApplicationScale{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationScaleArrayInput)(nil)).Elem(), ApplicationScaleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationScaleMapInput)(nil)).Elem(), ApplicationScaleMap{})
	pulumi.RegisterOutputType(ApplicationScaleOutput{})
	pulumi.RegisterOutputType(ApplicationScaleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationScaleMapOutput{})
}
