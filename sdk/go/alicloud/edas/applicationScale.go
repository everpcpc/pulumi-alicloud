// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This operation is provided to scale out an EDAS application.
//
// > **NOTE:** Available in 1.82.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/edas"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := edas.NewApplicationScale(ctx, "_default", &edas.ApplicationScaleArgs{
// 			AppId:       pulumi.Any(_var.App_id),
// 			DeployGroup: pulumi.Any(_var.Deploy_group),
// 			EcuInfos:    _var.Ecu_info,
// 			ForceStatus: pulumi.Any(_var.Force_status),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.DeployGroup == nil {
		return nil, errors.New("missing required argument 'DeployGroup'")
	}
	if args == nil || args.EcuInfos == nil {
		return nil, errors.New("missing required argument 'EcuInfos'")
	}
	if args == nil {
		args = &ApplicationScaleArgs{}
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
