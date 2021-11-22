// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Sso Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.138.0+.
//
// > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudsso"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := cloudsso.GetGroups(ctx, &cloudsso.GetGroupsArgs{
// 			DirectoryId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 				"example_value-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoGroupId1", ids.Groups[0].Id)
// 		opt0 := "^my-Group"
// 		nameRegex, err := cloudsso.GetGroups(ctx, &cloudsso.GetGroupsArgs{
// 			DirectoryId: "example_value",
// 			NameRegex:   &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoGroupId2", nameRegex.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	var rv GetGroupsResult
	err := ctx.Invoke("alicloud:cloudsso/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// A list of Group IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Provision Type of the Group. Valid values: `Manual`, `Synchronized`.
	ProvisionType *string `pulumi:"provisionType"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	DirectoryId string           `pulumi:"directoryId"`
	Groups      []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	NameRegex     *string  `pulumi:"nameRegex"`
	Names         []string `pulumi:"names"`
	OutputFile    *string  `pulumi:"outputFile"`
	ProvisionType *string  `pulumi:"provisionType"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			return *r, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// The ID of the Directory.
	DirectoryId pulumi.StringInput `pulumi:"directoryId"`
	// A list of Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Provision Type of the Group. Valid values: `Manual`, `Synchronized`.
	ProvisionType pulumi.StringPtrInput `pulumi:"provisionType"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.DirectoryId }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) ProvisionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.ProvisionType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}
