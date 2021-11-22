// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of DNS Resolution Lines in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.60.0.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "support_lines.txt"
// 		resolutionLinesDs, err := dns.GetResolutionLines(ctx, &dns.GetResolutionLinesArgs{
// 			LineCodes: []string{
// 				"cn_unicom_shanxi",
// 			},
// 			OutputFile: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstLineCode", resolutionLinesDs.Lines[0].LineCode)
// 		return nil
// 	})
// }
// ```
func GetResolutionLines(ctx *pulumi.Context, args *GetResolutionLinesArgs, opts ...pulumi.InvokeOption) (*GetResolutionLinesResult, error) {
	var rv GetResolutionLinesResult
	err := ctx.Invoke("alicloud:dns/getResolutionLines:getResolutionLines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResolutionLines.
type GetResolutionLinesArgs struct {
	// Domain Name.
	DomainName *string `pulumi:"domainName"`
	// language.
	Lang *string `pulumi:"lang"`
	// A list of lines codes.
	LineCodes []string `pulumi:"lineCodes"`
	// A list of line display names.
	LineDisplayNames []string `pulumi:"lineDisplayNames"`
	LineNames        []string `pulumi:"lineNames"`
	OutputFile       *string  `pulumi:"outputFile"`
	// The ip of user client.
	UserClientIp *string `pulumi:"userClientIp"`
}

// A collection of values returned by getResolutionLines.
type GetResolutionLinesResult struct {
	DomainName *string `pulumi:"domainName"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Lang *string `pulumi:"lang"`
	// Line code.
	LineCodes []string `pulumi:"lineCodes"`
	// A list of line display names.
	LineDisplayNames []string `pulumi:"lineDisplayNames"`
	LineNames        []string `pulumi:"lineNames"`
	// A list of cloud resolution line. Each element contains the following attributes:
	Lines        []GetResolutionLinesLine `pulumi:"lines"`
	OutputFile   *string                  `pulumi:"outputFile"`
	UserClientIp *string                  `pulumi:"userClientIp"`
}

func GetResolutionLinesOutput(ctx *pulumi.Context, args GetResolutionLinesOutputArgs, opts ...pulumi.InvokeOption) GetResolutionLinesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetResolutionLinesResult, error) {
			args := v.(GetResolutionLinesArgs)
			r, err := GetResolutionLines(ctx, &args, opts...)
			return *r, err
		}).(GetResolutionLinesResultOutput)
}

// A collection of arguments for invoking getResolutionLines.
type GetResolutionLinesOutputArgs struct {
	// Domain Name.
	DomainName pulumi.StringPtrInput `pulumi:"domainName"`
	// language.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// A list of lines codes.
	LineCodes pulumi.StringArrayInput `pulumi:"lineCodes"`
	// A list of line display names.
	LineDisplayNames pulumi.StringArrayInput `pulumi:"lineDisplayNames"`
	LineNames        pulumi.StringArrayInput `pulumi:"lineNames"`
	OutputFile       pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The ip of user client.
	UserClientIp pulumi.StringPtrInput `pulumi:"userClientIp"`
}

func (GetResolutionLinesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResolutionLinesArgs)(nil)).Elem()
}

// A collection of values returned by getResolutionLines.
type GetResolutionLinesResultOutput struct{ *pulumi.OutputState }

func (GetResolutionLinesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResolutionLinesResult)(nil)).Elem()
}

func (o GetResolutionLinesResultOutput) ToGetResolutionLinesResultOutput() GetResolutionLinesResultOutput {
	return o
}

func (o GetResolutionLinesResultOutput) ToGetResolutionLinesResultOutputWithContext(ctx context.Context) GetResolutionLinesResultOutput {
	return o
}

func (o GetResolutionLinesResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResolutionLinesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetResolutionLinesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

// Line code.
func (o GetResolutionLinesResultOutput) LineCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) []string { return v.LineCodes }).(pulumi.StringArrayOutput)
}

// A list of line display names.
func (o GetResolutionLinesResultOutput) LineDisplayNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) []string { return v.LineDisplayNames }).(pulumi.StringArrayOutput)
}

func (o GetResolutionLinesResultOutput) LineNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) []string { return v.LineNames }).(pulumi.StringArrayOutput)
}

// A list of cloud resolution line. Each element contains the following attributes:
func (o GetResolutionLinesResultOutput) Lines() GetResolutionLinesLineArrayOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) []GetResolutionLinesLine { return v.Lines }).(GetResolutionLinesLineArrayOutput)
}

func (o GetResolutionLinesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetResolutionLinesResultOutput) UserClientIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetResolutionLinesResult) *string { return v.UserClientIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResolutionLinesResultOutput{})
}
