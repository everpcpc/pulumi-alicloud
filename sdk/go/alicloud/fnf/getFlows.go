// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fnf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Fnf Flows of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.105.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fnf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "the_resource_name"
// 		example, err := fnf.GetFlows(ctx, &fnf.GetFlowsArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstFnfFlowId", example.Flows[0].Id)
// 		return nil
// 	})
// }
// ```
func GetFlows(ctx *pulumi.Context, args *GetFlowsArgs, opts ...pulumi.InvokeOption) (*GetFlowsResult, error) {
	var rv GetFlowsResult
	err := ctx.Invoke("alicloud:fnf/getFlows:getFlows", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlows.
type GetFlowsArgs struct {
	// A list of Flow IDs.
	Ids []string `pulumi:"ids"`
	// The number of resource queries.
	Limit *int `pulumi:"limit"`
	// A regex string to filter results by Flow name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getFlows.
type GetFlowsResult struct {
	Flows []GetFlowsFlow `pulumi:"flows"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	Limit      *int     `pulumi:"limit"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetFlowsOutput(ctx *pulumi.Context, args GetFlowsOutputArgs, opts ...pulumi.InvokeOption) GetFlowsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlowsResult, error) {
			args := v.(GetFlowsArgs)
			r, err := GetFlows(ctx, &args, opts...)
			return *r, err
		}).(GetFlowsResultOutput)
}

// A collection of arguments for invoking getFlows.
type GetFlowsOutputArgs struct {
	// A list of Flow IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The number of resource queries.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// A regex string to filter results by Flow name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetFlowsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowsArgs)(nil)).Elem()
}

// A collection of values returned by getFlows.
type GetFlowsResultOutput struct{ *pulumi.OutputState }

func (GetFlowsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlowsResult)(nil)).Elem()
}

func (o GetFlowsResultOutput) ToGetFlowsResultOutput() GetFlowsResultOutput {
	return o
}

func (o GetFlowsResultOutput) ToGetFlowsResultOutputWithContext(ctx context.Context) GetFlowsResultOutput {
	return o
}

func (o GetFlowsResultOutput) Flows() GetFlowsFlowArrayOutput {
	return o.ApplyT(func(v GetFlowsResult) []GetFlowsFlow { return v.Flows }).(GetFlowsFlowArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlowsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlowsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFlowsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlowsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetFlowsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlowsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetFlowsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetFlowsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlowsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetFlowsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlowsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlowsResultOutput{})
}
