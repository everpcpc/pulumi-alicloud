// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Config Aggregate Deliveries of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.172.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cfg"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := cfg.GetAggregateDeliveries(ctx, &cfg.GetAggregateDeliveriesArgs{
// 			AggregatorId: "example_value",
// 			Ids: []string{
// 				"example_value-1",
// 				"example_value-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("configAggregateDeliveryId1", ids.Deliveries[0].Id)
// 		return nil
// 	})
// }
// ```
func GetAggregateDeliveries(ctx *pulumi.Context, args *GetAggregateDeliveriesArgs, opts ...pulumi.InvokeOption) (*GetAggregateDeliveriesResult, error) {
	var rv GetAggregateDeliveriesResult
	err := ctx.Invoke("alicloud:cfg/getAggregateDeliveries:getAggregateDeliveries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAggregateDeliveries.
type GetAggregateDeliveriesArgs struct {
	// The ID of the Aggregator.
	AggregatorId string `pulumi:"aggregatorId"`
	// A list of Aggregate Delivery IDs.
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getAggregateDeliveries.
type GetAggregateDeliveriesResult struct {
	AggregatorId string                           `pulumi:"aggregatorId"`
	Deliveries   []GetAggregateDeliveriesDelivery `pulumi:"deliveries"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *int     `pulumi:"status"`
}

func GetAggregateDeliveriesOutput(ctx *pulumi.Context, args GetAggregateDeliveriesOutputArgs, opts ...pulumi.InvokeOption) GetAggregateDeliveriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAggregateDeliveriesResult, error) {
			args := v.(GetAggregateDeliveriesArgs)
			r, err := GetAggregateDeliveries(ctx, &args, opts...)
			var s GetAggregateDeliveriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAggregateDeliveriesResultOutput)
}

// A collection of arguments for invoking getAggregateDeliveries.
type GetAggregateDeliveriesOutputArgs struct {
	// The ID of the Aggregator.
	AggregatorId pulumi.StringInput `pulumi:"aggregatorId"`
	// A list of Aggregate Delivery IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	NameRegex  pulumi.StringPtrInput   `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
	Status pulumi.IntPtrInput `pulumi:"status"`
}

func (GetAggregateDeliveriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAggregateDeliveriesArgs)(nil)).Elem()
}

// A collection of values returned by getAggregateDeliveries.
type GetAggregateDeliveriesResultOutput struct{ *pulumi.OutputState }

func (GetAggregateDeliveriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAggregateDeliveriesResult)(nil)).Elem()
}

func (o GetAggregateDeliveriesResultOutput) ToGetAggregateDeliveriesResultOutput() GetAggregateDeliveriesResultOutput {
	return o
}

func (o GetAggregateDeliveriesResultOutput) ToGetAggregateDeliveriesResultOutputWithContext(ctx context.Context) GetAggregateDeliveriesResultOutput {
	return o
}

func (o GetAggregateDeliveriesResultOutput) AggregatorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) string { return v.AggregatorId }).(pulumi.StringOutput)
}

func (o GetAggregateDeliveriesResultOutput) Deliveries() GetAggregateDeliveriesDeliveryArrayOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) []GetAggregateDeliveriesDelivery { return v.Deliveries }).(GetAggregateDeliveriesDeliveryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAggregateDeliveriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAggregateDeliveriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAggregateDeliveriesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAggregateDeliveriesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAggregateDeliveriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAggregateDeliveriesResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAggregateDeliveriesResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAggregateDeliveriesResultOutput{})
}
