// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Threat Detection Vul Whitelists of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := threatdetection.GetVulWhitelists(ctx, &threatdetection.GetVulWhitelistsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudThreatDetectionVulWhitelistExampleId", _default.Whitelists[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetVulWhitelists(ctx *pulumi.Context, args *GetVulWhitelistsArgs, opts ...pulumi.InvokeOption) (*GetVulWhitelistsResult, error) {
	var rv GetVulWhitelistsResult
	err := ctx.Invoke("alicloud:threatdetection/getVulWhitelists:getVulWhitelists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVulWhitelists.
type GetVulWhitelistsArgs struct {
	// A list of Threat Detection Vul Whitelist IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
}

// A collection of values returned by getVulWhitelists.
type GetVulWhitelistsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	// A list of Vul Whitelist Entries. Each element contains the following attributes:
	Whitelists []GetVulWhitelistsWhitelist `pulumi:"whitelists"`
}

func GetVulWhitelistsOutput(ctx *pulumi.Context, args GetVulWhitelistsOutputArgs, opts ...pulumi.InvokeOption) GetVulWhitelistsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVulWhitelistsResult, error) {
			args := v.(GetVulWhitelistsArgs)
			r, err := GetVulWhitelists(ctx, &args, opts...)
			var s GetVulWhitelistsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVulWhitelistsResultOutput)
}

// A collection of arguments for invoking getVulWhitelists.
type GetVulWhitelistsOutputArgs struct {
	// A list of Threat Detection Vul Whitelist IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput      `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput      `pulumi:"pageSize"`
}

func (GetVulWhitelistsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVulWhitelistsArgs)(nil)).Elem()
}

// A collection of values returned by getVulWhitelists.
type GetVulWhitelistsResultOutput struct{ *pulumi.OutputState }

func (GetVulWhitelistsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVulWhitelistsResult)(nil)).Elem()
}

func (o GetVulWhitelistsResultOutput) ToGetVulWhitelistsResultOutput() GetVulWhitelistsResultOutput {
	return o
}

func (o GetVulWhitelistsResultOutput) ToGetVulWhitelistsResultOutputWithContext(ctx context.Context) GetVulWhitelistsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVulWhitelistsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVulWhitelistsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVulWhitelistsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVulWhitelistsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetVulWhitelistsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// A list of Vul Whitelist Entries. Each element contains the following attributes:
func (o GetVulWhitelistsResultOutput) Whitelists() GetVulWhitelistsWhitelistArrayOutput {
	return o.ApplyT(func(v GetVulWhitelistsResult) []GetVulWhitelistsWhitelist { return v.Whitelists }).(GetVulWhitelistsWhitelistArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVulWhitelistsResultOutput{})
}