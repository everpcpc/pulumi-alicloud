// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eais

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Eais Instances of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.137.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eais"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := eais.GetInstances(ctx, &eais.GetInstancesArgs{
// 			Id: []string{
// 				"example_id",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eaisInstanceId1", ids.Instances[0].Id)
// 		nameRegex, err := eais.GetInstances(ctx, &eais.GetInstancesArgs{
// 			NameRegex: pulumi.StringRef("^my-Instance"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eaisInstanceId2", nameRegex.Instances[0].Id)
// 		return nil
// 	})
// }
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:eais/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of Instance IDs.
	Ids []string `pulumi:"ids"`
	// The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
	InstanceType *string `pulumi:"instanceType"`
	// A regex string to filter results by Instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string                 `pulumi:"id"`
	Ids          []string               `pulumi:"ids"`
	InstanceType *string                `pulumi:"instanceType"`
	Instances    []GetInstancesInstance `pulumi:"instances"`
	NameRegex    *string                `pulumi:"nameRegex"`
	Names        []string               `pulumi:"names"`
	OutputFile   *string                `pulumi:"outputFile"`
	Status       *string                `pulumi:"status"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			return *r, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// A list of Instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// A regex string to filter results by Instance name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
