// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the apps of the current Alibaba Cloud user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "outapps"
// 		dataApigatway, err := apigateway.GetApps(ctx, &apigateway.GetAppsArgs{
// 			OutputFile: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstAppId", dataApigatway.Apps[0].Id)
// 		return nil
// 	})
// }
// ```
func GetApps(ctx *pulumi.Context, args *GetAppsArgs, opts ...pulumi.InvokeOption) (*GetAppsResult, error) {
	var rv GetAppsResult
	err := ctx.Invoke("alicloud:apigateway/getApps:getApps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApps.
type GetAppsArgs struct {
	// A list of app IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter apps by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getApps.
type GetAppsResult struct {
	// A list of apps. Each element contains the following attributes:
	Apps []GetAppsApp `pulumi:"apps"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of app IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of app names.
	Names      []string               `pulumi:"names"`
	OutputFile *string                `pulumi:"outputFile"`
	Tags       map[string]interface{} `pulumi:"tags"`
}

func GetAppsOutput(ctx *pulumi.Context, args GetAppsOutputArgs, opts ...pulumi.InvokeOption) GetAppsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsResult, error) {
			args := v.(GetAppsArgs)
			r, err := GetApps(ctx, &args, opts...)
			return *r, err
		}).(GetAppsResultOutput)
}

// A collection of arguments for invoking getApps.
type GetAppsOutputArgs struct {
	// A list of app IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter apps by name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetAppsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsArgs)(nil)).Elem()
}

// A collection of values returned by getApps.
type GetAppsResultOutput struct{ *pulumi.OutputState }

func (GetAppsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsResult)(nil)).Elem()
}

func (o GetAppsResultOutput) ToGetAppsResultOutput() GetAppsResultOutput {
	return o
}

func (o GetAppsResultOutput) ToGetAppsResultOutputWithContext(ctx context.Context) GetAppsResultOutput {
	return o
}

// A list of apps. Each element contains the following attributes:
func (o GetAppsResultOutput) Apps() GetAppsAppArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []GetAppsApp { return v.Apps }).(GetAppsAppArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of app IDs.
func (o GetAppsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAppsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of app names.
func (o GetAppsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAppsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAppsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetAppsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsResultOutput{})
}
