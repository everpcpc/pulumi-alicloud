// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the ots instances of the current Alibaba Cloud user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instancesDs, err := ots.GetInstances(ctx, &ots.GetInstancesArgs{
//				NameRegex:  pulumi.StringRef("sample-instance"),
//				OutputFile: pulumi.StringRef("instances.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstInstanceId", instancesDs.Instances[0].Id)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: alicloud.oss.getInstances has been deprecated in favor of alicloud.ots.getInstances
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:oss/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of instance IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// A map of tags assigned to the instance. It must be in the format:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		_, err := ots.GetInstances(ctx, &ots.GetInstancesArgs{
	// 			Tags: map[string]interface{}{
	// 				"tagKey1": "tagValue1",
	// 				"tagKey2": "tagValue2",
	// 			},
	// 		}, nil)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	})
	// }
	// ```
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of instances. Each element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// A list of instance names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The tags of the instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// A list of instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A map of tags assigned to the instance. It must be in the format:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		_, err := ots.GetInstances(ctx, &ots.GetInstancesArgs{
	// 			Tags: map[string]interface{}{
	// 				"tagKey1": "tagValue1",
	// 				"tagKey2": "tagValue2",
	// 			},
	// 		}, nil)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	})
	// }
	// ```
	Tags pulumi.MapInput `pulumi:"tags"`
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

// A list of instance IDs.
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of instances. Each element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of instance names.
func (o GetInstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The tags of the instance.
func (o GetInstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
