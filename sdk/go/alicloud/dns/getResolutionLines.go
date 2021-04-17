// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
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
