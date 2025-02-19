// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Eci Container Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.111.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eci"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := eci.GetContainerGroups(ctx, &eci.GetContainerGroupsArgs{
//				Ids: []string{
//					"example_value",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEciContainerGroupId", example.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetContainerGroups(ctx *pulumi.Context, args *GetContainerGroupsArgs, opts ...pulumi.InvokeOption) (*GetContainerGroupsResult, error) {
	var rv GetContainerGroupsResult
	err := ctx.Invoke("alicloud:eci/getContainerGroups:getContainerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerGroups.
type GetContainerGroupsArgs struct {
	// The name of ContainerGroup.
	ContainerGroupName *string `pulumi:"containerGroupName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Container Group IDs.
	Ids []string `pulumi:"ids"`
	// The maximum number of resources returned in the response. Default value is `20`. Maximum value: `20`. The number of returned results is no greater than the specified number.
	Limit *int `pulumi:"limit"`
	// A regex string to filter results by Container Group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group to which the container group belongs. If you have not specified a resource group for the container group, it is added to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of container.
	Status *string `pulumi:"status"`
	// The tags attached to the container group. Each tag is a key-value pair. You can attach up to 20 tags to a container group.
	Tags map[string]interface{} `pulumi:"tags"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
	WithEvent *bool   `pulumi:"withEvent"`
	// The IDs of the zones where the container groups are deployed. If this parameter is not set, the system automatically selects the zones. By default, no value is specified.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getContainerGroups.
type GetContainerGroupsResult struct {
	ContainerGroupName *string                   `pulumi:"containerGroupName"`
	EnableDetails      *bool                     `pulumi:"enableDetails"`
	Groups             []GetContainerGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	Limit           *int                   `pulumi:"limit"`
	NameRegex       *string                `pulumi:"nameRegex"`
	Names           []string               `pulumi:"names"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
	VswitchId       *string                `pulumi:"vswitchId"`
	WithEvent       *bool                  `pulumi:"withEvent"`
	ZoneId          *string                `pulumi:"zoneId"`
}

func GetContainerGroupsOutput(ctx *pulumi.Context, args GetContainerGroupsOutputArgs, opts ...pulumi.InvokeOption) GetContainerGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerGroupsResult, error) {
			args := v.(GetContainerGroupsArgs)
			r, err := GetContainerGroups(ctx, &args, opts...)
			var s GetContainerGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContainerGroupsResultOutput)
}

// A collection of arguments for invoking getContainerGroups.
type GetContainerGroupsOutputArgs struct {
	// The name of ContainerGroup.
	ContainerGroupName pulumi.StringPtrInput `pulumi:"containerGroupName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Container Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The maximum number of resources returned in the response. Default value is `20`. Maximum value: `20`. The number of returned results is no greater than the specified number.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// A regex string to filter results by Container Group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the resource group to which the container group belongs. If you have not specified a resource group for the container group, it is added to the default resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status of container.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The tags attached to the container group. Each tag is a key-value pair. You can attach up to 20 tags to a container group.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The vswitch id.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
	WithEvent pulumi.BoolPtrInput   `pulumi:"withEvent"`
	// The IDs of the zones where the container groups are deployed. If this parameter is not set, the system automatically selects the zones. By default, no value is specified.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetContainerGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getContainerGroups.
type GetContainerGroupsResultOutput struct{ *pulumi.OutputState }

func (GetContainerGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerGroupsResult)(nil)).Elem()
}

func (o GetContainerGroupsResultOutput) ToGetContainerGroupsResultOutput() GetContainerGroupsResultOutput {
	return o
}

func (o GetContainerGroupsResultOutput) ToGetContainerGroupsResultOutputWithContext(ctx context.Context) GetContainerGroupsResultOutput {
	return o
}

func (o GetContainerGroupsResultOutput) ContainerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.ContainerGroupName }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetContainerGroupsResultOutput) Groups() GetContainerGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) []GetContainerGroupsGroup { return v.Groups }).(GetContainerGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContainerGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetContainerGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetContainerGroupsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetContainerGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetContainerGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetContainerGroupsResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func (o GetContainerGroupsResultOutput) WithEvent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *bool { return v.WithEvent }).(pulumi.BoolPtrOutput)
}

func (o GetContainerGroupsResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerGroupsResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerGroupsResultOutput{})
}
