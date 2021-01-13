// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Resource Manager Resource Group resource. If you need to group cloud resources according to business departments, projects, and other dimensions, you can create resource groups.
// For information about Resource Manager Resoource Group and how to use it, see [What is Resource Manager Resource Group](https://www.alibabacloud.com/help/en/doc-detail/94485.htm)
//
// > **NOTE:** Available in v1.82.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcemanager.NewResourceGroup(ctx, "example", &resourcemanager.ResourceGroupArgs{
// 			DisplayName: pulumi.String("testrd"),
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
// Resource Manager Resource Group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/resourceGroup:ResourceGroup example abc123456
// ```
type ResourceGroup struct {
	pulumi.CustomResourceState

	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The time when the resource group was created.
	// * `regionStatuses` -The status of the resource group in all regions.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
	Name           pulumi.StringOutput                  `pulumi:"name"`
	RegionStatuses ResourceGroupRegionStatusArrayOutput `pulumi:"regionStatuses"`
	// The status of the regional resource group.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource ResourceGroup
	err := ctx.RegisterResource("alicloud:resourcemanager/resourceGroup:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("alicloud:resourcemanager/resourceGroup:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId *string `pulumi:"accountId"`
	// The time when the resource group was created.
	// * `regionStatuses` -The status of the resource group in all regions.
	CreateDate *string `pulumi:"createDate"`
	// The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
	DisplayName *string `pulumi:"displayName"`
	// The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
	Name           *string                     `pulumi:"name"`
	RegionStatuses []ResourceGroupRegionStatus `pulumi:"regionStatuses"`
	// The status of the regional resource group.
	Status *string `pulumi:"status"`
}

type ResourceGroupState struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	AccountId pulumi.StringPtrInput
	// The time when the resource group was created.
	// * `regionStatuses` -The status of the resource group in all regions.
	CreateDate pulumi.StringPtrInput
	// The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
	DisplayName pulumi.StringPtrInput
	// The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
	Name           pulumi.StringPtrInput
	RegionStatuses ResourceGroupRegionStatusArrayInput
	// The status of the regional resource group.
	Status pulumi.StringPtrInput
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
	DisplayName string `pulumi:"displayName"`
	// The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-).
	DisplayName pulumi.StringInput
	// The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.
	Name pulumi.StringPtrInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroup)(nil)).Elem()
}

func (i ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

type ResourceGroupOutput struct {
	*pulumi.OutputState
}

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupOutput)(nil)).Elem()
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceGroupOutput{})
}
