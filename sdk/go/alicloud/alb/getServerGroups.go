// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Alb Server Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.131.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := alb.GetServerGroups(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albServerGroupId1", ids.Groups[0].Id)
// 		opt0 := "^my-ServerGroup"
// 		nameRegex, err := alb.GetServerGroups(ctx, &alb.GetServerGroupsArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albServerGroupId2", nameRegex.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetServerGroups(ctx *pulumi.Context, args *GetServerGroupsArgs, opts ...pulumi.InvokeOption) (*GetServerGroupsResult, error) {
	var rv GetServerGroupsResult
	err := ctx.Invoke("alicloud:alb/getServerGroups:getServerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerGroups.
type GetServerGroupsArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Server Group IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Server Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The server group ids.
	ServerGroupIds []string `pulumi:"serverGroupIds"`
	// The name of the resource.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The status of the resource. Valid values: `Provisioning`, `Available` and `Configuring`.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC that you want to access.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getServerGroups.
type GetServerGroupsResult struct {
	EnableDetails *bool                  `pulumi:"enableDetails"`
	Groups        []GetServerGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	NameRegex       *string                `pulumi:"nameRegex"`
	Names           []string               `pulumi:"names"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	ServerGroupIds  []string               `pulumi:"serverGroupIds"`
	ServerGroupName *string                `pulumi:"serverGroupName"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
	VpcId           *string                `pulumi:"vpcId"`
}

func GetServerGroupsOutput(ctx *pulumi.Context, args GetServerGroupsOutputArgs, opts ...pulumi.InvokeOption) GetServerGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerGroupsResult, error) {
			args := v.(GetServerGroupsArgs)
			r, err := GetServerGroups(ctx, &args, opts...)
			return *r, err
		}).(GetServerGroupsResultOutput)
}

// A collection of arguments for invoking getServerGroups.
type GetServerGroupsOutputArgs struct {
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Server Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Server Group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The server group ids.
	ServerGroupIds pulumi.StringArrayInput `pulumi:"serverGroupIds"`
	// The name of the resource.
	ServerGroupName pulumi.StringPtrInput `pulumi:"serverGroupName"`
	// The status of the resource. Valid values: `Provisioning`, `Available` and `Configuring`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	Tags   pulumi.MapInput       `pulumi:"tags"`
	// The ID of the VPC that you want to access.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetServerGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getServerGroups.
type GetServerGroupsResultOutput struct{ *pulumi.OutputState }

func (GetServerGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupsResult)(nil)).Elem()
}

func (o GetServerGroupsResultOutput) ToGetServerGroupsResultOutput() GetServerGroupsResultOutput {
	return o
}

func (o GetServerGroupsResultOutput) ToGetServerGroupsResultOutputWithContext(ctx context.Context) GetServerGroupsResultOutput {
	return o
}

func (o GetServerGroupsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetServerGroupsResultOutput) Groups() GetServerGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []GetServerGroupsGroup { return v.Groups }).(GetServerGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServerGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServerGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupsResultOutput) ServerGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerGroupsResult) []string { return v.ServerGroupIds }).(pulumi.StringArrayOutput)
}

func (o GetServerGroupsResultOutput) ServerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.ServerGroupName }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetServerGroupsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetServerGroupsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerGroupsResultOutput{})
}
