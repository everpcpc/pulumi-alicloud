// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Ad Connector Office Sites of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.176.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := eds.GetAdConnectorOfficeSites(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdAdConnectorOfficeSiteId1", ids.Sites[0].Id)
//			nameRegex, err := eds.GetAdConnectorOfficeSites(ctx, &eds.GetAdConnectorOfficeSitesArgs{
//				NameRegex: pulumi.StringRef("^my-AdConnectorOfficeSite"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdAdConnectorOfficeSiteId2", nameRegex.Sites[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAdConnectorOfficeSites(ctx *pulumi.Context, args *GetAdConnectorOfficeSitesArgs, opts ...pulumi.InvokeOption) (*GetAdConnectorOfficeSitesResult, error) {
	var rv GetAdConnectorOfficeSitesResult
	err := ctx.Invoke("alicloud:eds/getAdConnectorOfficeSites:getAdConnectorOfficeSites", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAdConnectorOfficeSites.
type GetAdConnectorOfficeSitesArgs struct {
	// A list of Ad Connector Office Site IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Ad Connector Office Site name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The workspace status.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAdConnectorOfficeSites.
type GetAdConnectorOfficeSitesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                          `pulumi:"id"`
	Ids        []string                        `pulumi:"ids"`
	NameRegex  *string                         `pulumi:"nameRegex"`
	Names      []string                        `pulumi:"names"`
	OutputFile *string                         `pulumi:"outputFile"`
	Sites      []GetAdConnectorOfficeSitesSite `pulumi:"sites"`
	Status     *string                         `pulumi:"status"`
}

func GetAdConnectorOfficeSitesOutput(ctx *pulumi.Context, args GetAdConnectorOfficeSitesOutputArgs, opts ...pulumi.InvokeOption) GetAdConnectorOfficeSitesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAdConnectorOfficeSitesResult, error) {
			args := v.(GetAdConnectorOfficeSitesArgs)
			r, err := GetAdConnectorOfficeSites(ctx, &args, opts...)
			var s GetAdConnectorOfficeSitesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAdConnectorOfficeSitesResultOutput)
}

// A collection of arguments for invoking getAdConnectorOfficeSites.
type GetAdConnectorOfficeSitesOutputArgs struct {
	// A list of Ad Connector Office Site IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Ad Connector Office Site name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The workspace status.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetAdConnectorOfficeSitesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdConnectorOfficeSitesArgs)(nil)).Elem()
}

// A collection of values returned by getAdConnectorOfficeSites.
type GetAdConnectorOfficeSitesResultOutput struct{ *pulumi.OutputState }

func (GetAdConnectorOfficeSitesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdConnectorOfficeSitesResult)(nil)).Elem()
}

func (o GetAdConnectorOfficeSitesResultOutput) ToGetAdConnectorOfficeSitesResultOutput() GetAdConnectorOfficeSitesResultOutput {
	return o
}

func (o GetAdConnectorOfficeSitesResultOutput) ToGetAdConnectorOfficeSitesResultOutputWithContext(ctx context.Context) GetAdConnectorOfficeSitesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAdConnectorOfficeSitesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) Sites() GetAdConnectorOfficeSitesSiteArrayOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) []GetAdConnectorOfficeSitesSite { return v.Sites }).(GetAdConnectorOfficeSitesSiteArrayOutput)
}

func (o GetAdConnectorOfficeSitesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorOfficeSitesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAdConnectorOfficeSitesResultOutput{})
}
