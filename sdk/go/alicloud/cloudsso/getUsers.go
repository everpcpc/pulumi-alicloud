// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Sso Users of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.140.0+.
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
// 		ids, err := cloudsso.GetUsers(ctx, &cloudsso.GetUsersArgs{
// 			DirectoryId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 				"example_value-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoUserId1", ids.Users[0].Id)
// 		opt0 := "^my-User"
// 		nameRegex, err := cloudsso.GetUsers(ctx, &cloudsso.GetUsersArgs{
// 			DirectoryId: "example_value",
// 			NameRegex:   &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoUserId2", nameRegex.Users[0].Id)
// 		opt1 := "Manual"
// 		provisionType, err := cloudsso.GetUsers(ctx, &cloudsso.GetUsersArgs{
// 			DirectoryId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 			},
// 			ProvisionType: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoUserId3", provisionType.Users[0].Id)
// 		opt2 := "Enabled"
// 		status, err := cloudsso.GetUsers(ctx, &cloudsso.GetUsersArgs{
// 			DirectoryId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 			},
// 			Status: &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("cloudSsoUserId4", status.Users[0].Id)
// 		return nil
// 	})
// }
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("alicloud:cloudsso/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of User IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by User name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// ProvisionType.
	ProvisionType *string `pulumi:"provisionType"`
	// User status. Valid values: `Enabled` and `Disabled`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	DirectoryId   string `pulumi:"directoryId"`
	EnableDetails *bool  `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id            string         `pulumi:"id"`
	Ids           []string       `pulumi:"ids"`
	NameRegex     *string        `pulumi:"nameRegex"`
	Names         []string       `pulumi:"names"`
	OutputFile    *string        `pulumi:"outputFile"`
	ProvisionType *string        `pulumi:"provisionType"`
	Status        *string        `pulumi:"status"`
	Users         []GetUsersUser `pulumi:"users"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			return *r, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// The ID of the Directory.
	DirectoryId pulumi.StringInput `pulumi:"directoryId"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of User IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by User name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// ProvisionType.
	ProvisionType pulumi.StringPtrInput `pulumi:"provisionType"`
	// User status. Valid values: `Enabled` and `Disabled`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.DirectoryId }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetUsersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetUsersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) ProvisionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.ProvisionType }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
