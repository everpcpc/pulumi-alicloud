// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Sae Instance Specifications of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.139.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sae"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := sae.GetInstanceSpecifications(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("saeInstanceSpecificationId1", ids.Specifications[0].Id)
// 		return nil
// 	})
// }
// ```
func GetInstanceSpecifications(ctx *pulumi.Context, args *GetInstanceSpecificationsArgs, opts ...pulumi.InvokeOption) (*GetInstanceSpecificationsResult, error) {
	var rv GetInstanceSpecificationsResult
	err := ctx.Invoke("alicloud:sae/getInstanceSpecifications:getInstanceSpecifications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceSpecifications.
type GetInstanceSpecificationsArgs struct {
	// A list of Instance Specification IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
}

// A collection of values returned by getInstanceSpecifications.
type GetInstanceSpecificationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id             string                                   `pulumi:"id"`
	Ids            []string                                 `pulumi:"ids"`
	OutputFile     *string                                  `pulumi:"outputFile"`
	Specifications []GetInstanceSpecificationsSpecification `pulumi:"specifications"`
}

func GetInstanceSpecificationsOutput(ctx *pulumi.Context, args GetInstanceSpecificationsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceSpecificationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceSpecificationsResult, error) {
			args := v.(GetInstanceSpecificationsArgs)
			r, err := GetInstanceSpecifications(ctx, &args, opts...)
			return *r, err
		}).(GetInstanceSpecificationsResultOutput)
}

// A collection of arguments for invoking getInstanceSpecifications.
type GetInstanceSpecificationsOutputArgs struct {
	// A list of Instance Specification IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
}

func (GetInstanceSpecificationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSpecificationsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceSpecifications.
type GetInstanceSpecificationsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceSpecificationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceSpecificationsResult)(nil)).Elem()
}

func (o GetInstanceSpecificationsResultOutput) ToGetInstanceSpecificationsResultOutput() GetInstanceSpecificationsResultOutput {
	return o
}

func (o GetInstanceSpecificationsResultOutput) ToGetInstanceSpecificationsResultOutputWithContext(ctx context.Context) GetInstanceSpecificationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceSpecificationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceSpecificationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceSpecificationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceSpecificationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstanceSpecificationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceSpecificationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstanceSpecificationsResultOutput) Specifications() GetInstanceSpecificationsSpecificationArrayOutput {
	return o.ApplyT(func(v GetInstanceSpecificationsResult) []GetInstanceSpecificationsSpecification {
		return v.Specifications
	}).(GetInstanceSpecificationsSpecificationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceSpecificationsResultOutput{})
}
