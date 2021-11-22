// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `emr.getInstanceTypes` data source provides a collection of ecs
// instance types available in Alibaba Cloud account when create a emr cluster.
//
// > **NOTE:** Available in 1.59.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/emr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "ecs.g5.2xlarge"
// 		opt1 := false
// 		_default, err := emr.GetInstanceTypes(ctx, &emr.GetInstanceTypesArgs{
// 			ClusterType:         "HADOOP",
// 			DestinationResource: "InstanceType",
// 			InstanceChargeType:  "PostPaid",
// 			InstanceType:        &opt0,
// 			SupportLocalStorage: &opt1,
// 			SupportNodeTypes: []string{
// 				"MASTER",
// 				"CORE",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstInstanceType", _default.Types[0].Id)
// 		return nil
// 	})
// }
// ```
func GetInstanceTypes(ctx *pulumi.Context, args *GetInstanceTypesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypesResult, error) {
	var rv GetInstanceTypesResult
	err := ctx.Invoke("alicloud:emr/getInstanceTypes:getInstanceTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesArgs struct {
	// The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
	ClusterType string `pulumi:"clusterType"`
	// The destination resource of emr cluster instance
	DestinationResource string `pulumi:"destinationResource"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType string `pulumi:"instanceChargeType"`
	// Filter the specific ecs instance type to create emr cluster.
	InstanceType *string `pulumi:"instanceType"`
	OutputFile   *string `pulumi:"outputFile"`
	// Whether the current storage disk is local or not.
	SupportLocalStorage *bool `pulumi:"supportLocalStorage"`
	// The specific supported node type list.
	// Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
	SupportNodeTypes []string `pulumi:"supportNodeTypes"`
	// The supported resources of specific zoneId.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResult struct {
	ClusterType         string `pulumi:"clusterType"`
	DestinationResource string `pulumi:"destinationResource"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of emr instance types IDs.
	Ids                 []string `pulumi:"ids"`
	InstanceChargeType  string   `pulumi:"instanceChargeType"`
	InstanceType        *string  `pulumi:"instanceType"`
	OutputFile          *string  `pulumi:"outputFile"`
	SupportLocalStorage *bool    `pulumi:"supportLocalStorage"`
	SupportNodeTypes    []string `pulumi:"supportNodeTypes"`
	// A list of emr instance types. Each element contains the following attributes:
	Types []GetInstanceTypesType `pulumi:"types"`
	// The available zone id in Alibaba Cloud account
	ZoneId *string `pulumi:"zoneId"`
}

func GetInstanceTypesOutput(ctx *pulumi.Context, args GetInstanceTypesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceTypesResult, error) {
			args := v.(GetInstanceTypesArgs)
			r, err := GetInstanceTypes(ctx, &args, opts...)
			return *r, err
		}).(GetInstanceTypesResultOutput)
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesOutputArgs struct {
	// The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
	ClusterType pulumi.StringInput `pulumi:"clusterType"`
	// The destination resource of emr cluster instance
	DestinationResource pulumi.StringInput `pulumi:"destinationResource"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType pulumi.StringInput `pulumi:"instanceChargeType"`
	// Filter the specific ecs instance type to create emr cluster.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	OutputFile   pulumi.StringPtrInput `pulumi:"outputFile"`
	// Whether the current storage disk is local or not.
	SupportLocalStorage pulumi.BoolPtrInput `pulumi:"supportLocalStorage"`
	// The specific supported node type list.
	// Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
	SupportNodeTypes pulumi.StringArrayInput `pulumi:"supportNodeTypes"`
	// The supported resources of specific zoneId.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetInstanceTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesResult)(nil)).Elem()
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutput() GetInstanceTypesResultOutput {
	return o
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutputWithContext(ctx context.Context) GetInstanceTypesResultOutput {
	return o
}

func (o GetInstanceTypesResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

func (o GetInstanceTypesResultOutput) DestinationResource() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.DestinationResource }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of emr instance types IDs.
func (o GetInstanceTypesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstanceTypesResultOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.InstanceChargeType }).(pulumi.StringOutput)
}

func (o GetInstanceTypesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) SupportLocalStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *bool { return v.SupportLocalStorage }).(pulumi.BoolPtrOutput)
}

func (o GetInstanceTypesResultOutput) SupportNodeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []string { return v.SupportNodeTypes }).(pulumi.StringArrayOutput)
}

// A list of emr instance types. Each element contains the following attributes:
func (o GetInstanceTypesResultOutput) Types() GetInstanceTypesTypeArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []GetInstanceTypesType { return v.Types }).(GetInstanceTypesTypeArrayOutput)
}

// The available zone id in Alibaba Cloud account
func (o GetInstanceTypesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceTypesResultOutput{})
}
