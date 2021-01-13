// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an EDAS deploy group resource.
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
// 		_, err := edas.NewDeployGroup(ctx, "_default", &edas.DeployGroupArgs{
// 			AppId:     pulumi.Any(_var.App_id),
// 			GroupName: pulumi.Any(_var.Group_name),
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
// EDAS deploy group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:edas/deployGroup:DeployGroup group app_id:group_name:group_id
// ```
type DeployGroup struct {
	pulumi.CustomResourceState

	// The ID of the application that you want to deploy.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType pulumi.IntOutput `pulumi:"groupType"`
}

// NewDeployGroup registers a new resource with the given unique name, arguments, and options.
func NewDeployGroup(ctx *pulumi.Context,
	name string, args *DeployGroupArgs, opts ...pulumi.ResourceOption) (*DeployGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	var resource DeployGroup
	err := ctx.RegisterResource("alicloud:edas/deployGroup:DeployGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployGroup gets an existing DeployGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployGroupState, opts ...pulumi.ResourceOption) (*DeployGroup, error) {
	var resource DeployGroup
	err := ctx.ReadResource("alicloud:edas/deployGroup:DeployGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployGroup resources.
type deployGroupState struct {
	// The ID of the application that you want to deploy.
	AppId *string `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName *string `pulumi:"groupName"`
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType *int `pulumi:"groupType"`
}

type DeployGroupState struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringPtrInput
	// The name of the instance group that you want to create.
	GroupName pulumi.StringPtrInput
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType pulumi.IntPtrInput
}

func (DeployGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployGroupState)(nil)).Elem()
}

type deployGroupArgs struct {
	// The ID of the application that you want to deploy.
	AppId string `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName string `pulumi:"groupName"`
}

// The set of arguments for constructing a DeployGroup resource.
type DeployGroupArgs struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringInput
	// The name of the instance group that you want to create.
	GroupName pulumi.StringInput
}

func (DeployGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployGroupArgs)(nil)).Elem()
}

type DeployGroupInput interface {
	pulumi.Input

	ToDeployGroupOutput() DeployGroupOutput
	ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput
}

func (DeployGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployGroup)(nil)).Elem()
}

func (i DeployGroup) ToDeployGroupOutput() DeployGroupOutput {
	return i.ToDeployGroupOutputWithContext(context.Background())
}

func (i DeployGroup) ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployGroupOutput)
}

type DeployGroupOutput struct {
	*pulumi.OutputState
}

func (DeployGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployGroupOutput)(nil)).Elem()
}

func (o DeployGroupOutput) ToDeployGroupOutput() DeployGroupOutput {
	return o
}

func (o DeployGroupOutput) ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeployGroupOutput{})
}
