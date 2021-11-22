// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Resource Manager Roles of the current Alibaba Cloud user.
//
// > **NOTE:**  Available in 1.86.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "tftest"
// 		example, err := resourcemanager.GetRoles(ctx, &resourcemanager.GetRolesArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstRoleId", example.Roles[0].Id)
// 		return nil
// 	})
// }
// ```
func GetRoles(ctx *pulumi.Context, args *GetRolesArgs, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	var rv GetRolesResult
	err := ctx.Invoke("alicloud:resourcemanager/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type GetRolesArgs struct {
	// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Resource Manager Role IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by role name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of role IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of role names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of roles. Each element contains the following attributes:
	Roles []GetRolesRole `pulumi:"roles"`
}

func GetRolesOutput(ctx *pulumi.Context, args GetRolesOutputArgs, opts ...pulumi.InvokeOption) GetRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolesResult, error) {
			args := v.(GetRolesArgs)
			r, err := GetRoles(ctx, &args, opts...)
			return *r, err
		}).(GetRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type GetRolesOutputArgs struct {
	// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Resource Manager Role IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by role name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type GetRolesResultOutput struct{ *pulumi.OutputState }

func (GetRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesResult)(nil)).Elem()
}

func (o GetRolesResultOutput) ToGetRolesResultOutput() GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) ToGetRolesResultOutputWithContext(ctx context.Context) GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of role IDs.
func (o GetRolesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRolesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of role names.
func (o GetRolesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRolesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of roles. Each element contains the following attributes:
func (o GetRolesResultOutput) Roles() GetRolesRoleArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []GetRolesRole { return v.Roles }).(GetRolesRoleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesResultOutput{})
}
