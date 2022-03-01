// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		queues, err := mns.GetQueues(ctx, &mns.GetQueuesArgs{
// 			NamePrefix: pulumi.StringRef("tf-"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstQueueId", queues.Queues[0].Id)
// 		return nil
// 	})
// }
// ```
func GetQueues(ctx *pulumi.Context, args *GetQueuesArgs, opts ...pulumi.InvokeOption) (*GetQueuesResult, error) {
	var rv GetQueuesResult
	err := ctx.Invoke("alicloud:mns/getQueues:getQueues", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueues.
type GetQueuesArgs struct {
	// A string to filter resulting queues by their name prefixs.
	NamePrefix *string `pulumi:"namePrefix"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getQueues.
type GetQueuesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of queue names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of queues. Each element contains the following attributes:
	Queues []GetQueuesQueue `pulumi:"queues"`
}

func GetQueuesOutput(ctx *pulumi.Context, args GetQueuesOutputArgs, opts ...pulumi.InvokeOption) GetQueuesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetQueuesResult, error) {
			args := v.(GetQueuesArgs)
			r, err := GetQueues(ctx, &args, opts...)
			return *r, err
		}).(GetQueuesResultOutput)
}

// A collection of arguments for invoking getQueues.
type GetQueuesOutputArgs struct {
	// A string to filter resulting queues by their name prefixs.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetQueuesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQueuesArgs)(nil)).Elem()
}

// A collection of values returned by getQueues.
type GetQueuesResultOutput struct{ *pulumi.OutputState }

func (GetQueuesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQueuesResult)(nil)).Elem()
}

func (o GetQueuesResultOutput) ToGetQueuesResultOutput() GetQueuesResultOutput {
	return o
}

func (o GetQueuesResultOutput) ToGetQueuesResultOutputWithContext(ctx context.Context) GetQueuesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetQueuesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueuesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetQueuesResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQueuesResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// A list of queue names.
func (o GetQueuesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetQueuesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetQueuesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQueuesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of queues. Each element contains the following attributes:
func (o GetQueuesResultOutput) Queues() GetQueuesQueueArrayOutput {
	return o.ApplyT(func(v GetQueuesResult) []GetQueuesQueue { return v.Queues }).(GetQueuesQueueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetQueuesResultOutput{})
}
