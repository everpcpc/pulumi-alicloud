// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Iot Device Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/iot"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := iot.GetDeviceGroups(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("iotDeviceGroupId1", ids.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetDeviceGroups(ctx *pulumi.Context, args *GetDeviceGroupsArgs, opts ...pulumi.InvokeOption) (*GetDeviceGroupsResult, error) {
	var rv GetDeviceGroupsResult
	err := ctx.Invoke("alicloud:iot/getDeviceGroups:getDeviceGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeviceGroups.
type GetDeviceGroupsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// The GroupName of the device group.
	GroupName *string `pulumi:"groupName"`
	// A list of device group IDs.
	Ids []string `pulumi:"ids"`
	// The id of the Iot Instance.
	IotInstanceId *string `pulumi:"iotInstanceId"`
	// A regex string to filter CEN instances by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The id of the SuperGroup.
	SuperGroupId *string `pulumi:"superGroupId"`
}

// A collection of values returned by getDeviceGroups.
type GetDeviceGroupsResult struct {
	EnableDetails *bool                  `pulumi:"enableDetails"`
	GroupName     *string                `pulumi:"groupName"`
	Groups        []GetDeviceGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	IotInstanceId *string  `pulumi:"iotInstanceId"`
	NameRegex     *string  `pulumi:"nameRegex"`
	OutputFile    *string  `pulumi:"outputFile"`
	SuperGroupId  *string  `pulumi:"superGroupId"`
}

func GetDeviceGroupsOutput(ctx *pulumi.Context, args GetDeviceGroupsOutputArgs, opts ...pulumi.InvokeOption) GetDeviceGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeviceGroupsResult, error) {
			args := v.(GetDeviceGroupsArgs)
			r, err := GetDeviceGroups(ctx, &args, opts...)
			return *r, err
		}).(GetDeviceGroupsResultOutput)
}

// A collection of arguments for invoking getDeviceGroups.
type GetDeviceGroupsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// The GroupName of the device group.
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// A list of device group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The id of the Iot Instance.
	IotInstanceId pulumi.StringPtrInput `pulumi:"iotInstanceId"`
	// A regex string to filter CEN instances by name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of the SuperGroup.
	SuperGroupId pulumi.StringPtrInput `pulumi:"superGroupId"`
}

func (GetDeviceGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getDeviceGroups.
type GetDeviceGroupsResultOutput struct{ *pulumi.OutputState }

func (GetDeviceGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceGroupsResult)(nil)).Elem()
}

func (o GetDeviceGroupsResultOutput) ToGetDeviceGroupsResultOutput() GetDeviceGroupsResultOutput {
	return o
}

func (o GetDeviceGroupsResultOutput) ToGetDeviceGroupsResultOutputWithContext(ctx context.Context) GetDeviceGroupsResultOutput {
	return o
}

func (o GetDeviceGroupsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetDeviceGroupsResultOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o GetDeviceGroupsResultOutput) Groups() GetDeviceGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) []GetDeviceGroupsGroup { return v.Groups }).(GetDeviceGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeviceGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeviceGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDeviceGroupsResultOutput) IotInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *string { return v.IotInstanceId }).(pulumi.StringPtrOutput)
}

func (o GetDeviceGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetDeviceGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDeviceGroupsResultOutput) SuperGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceGroupsResult) *string { return v.SuperGroupId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeviceGroupsResultOutput{})
}