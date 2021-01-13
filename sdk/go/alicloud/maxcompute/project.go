// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maxcompute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The project is the basic unit of operation in maxcompute. It is similar to the concept of Database or Schema in traditional databases, and sets the boundary for maxcompute multi-user isolation and access control. [Refer to details](https://www.alibabacloud.com/help/doc-detail/27818.html).
//
// ->**NOTE:** Available in 1.77.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/maxcompute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := maxcompute.NewProject(ctx, "example", &maxcompute.ProjectArgs{
// 			OrderType:         pulumi.String("PayAsYouGo"),
// 			ProjectName:       pulumi.String("tf_maxcompute_project"),
// 			SpecificationType: pulumi.String("OdpsStandard"),
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
// MaxCompute project can be imported using the *name* or ID, e.g.
//
// ```sh
//  $ pulumi import alicloud:maxcompute/project:Project example tf_maxcompute_project
// ```
type Project struct {
	pulumi.CustomResourceState

	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringOutput `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringOutput `pulumi:"specificationType"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrderType == nil {
		return nil, errors.New("invalid value for required argument 'OrderType'")
	}
	if args.SpecificationType == nil {
		return nil, errors.New("invalid value for required argument 'SpecificationType'")
	}
	var resource Project
	err := ctx.RegisterResource("alicloud:maxcompute/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("alicloud:maxcompute/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name *string `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType *string `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName *string `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType *string `pulumi:"specificationType"`
}

type ProjectState struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringPtrInput
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringPtrInput
	// The name of the maxcompute project.
	ProjectName pulumi.StringPtrInput
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name *string `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType string `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName *string `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType string `pulumi:"specificationType"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringPtrInput
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringInput
	// The name of the maxcompute project.
	ProjectName pulumi.StringPtrInput
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (Project) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil)).Elem()
}

func (i Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct {
	*pulumi.OutputState
}

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectOutput)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
