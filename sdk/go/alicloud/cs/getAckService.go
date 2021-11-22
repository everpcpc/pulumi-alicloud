// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open Container Service (CS) service automatically. If the service has been opened, it will return opened.
//
// For information about Container Service (CS) and how to use it, see [What is Container Service (CS)](https://www.alibabacloud.com/help/en/product/85222.htm).
//
// > **NOTE:** Available in v1.113.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "On"
// 		_, err := cs.GetAckService(ctx, &cs.GetAckServiceArgs{
// 			Enable: &opt0,
// 			Type:   "propayasgo",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAckService(ctx *pulumi.Context, args *GetAckServiceArgs, opts ...pulumi.InvokeOption) (*GetAckServiceResult, error) {
	var rv GetAckServiceResult
	err := ctx.Invoke("alicloud:cs/getAckService:getAckService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAckService.
type GetAckServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
	// Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
	Type string `pulumi:"type"`
}

// A collection of values returned by getAckService.
type GetAckServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
	Type   string `pulumi:"type"`
}

func GetAckServiceOutput(ctx *pulumi.Context, args GetAckServiceOutputArgs, opts ...pulumi.InvokeOption) GetAckServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAckServiceResult, error) {
			args := v.(GetAckServiceArgs)
			r, err := GetAckService(ctx, &args, opts...)
			return *r, err
		}).(GetAckServiceResultOutput)
}

// A collection of arguments for invoking getAckService.
type GetAckServiceOutputArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable pulumi.StringPtrInput `pulumi:"enable"`
	// Types of services opened. Valid values: `propayasgo`: Container service ack Pro managed version, `edgepayasgo`: Edge container service, `gspayasgo`: Gene computing services.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetAckServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAckServiceArgs)(nil)).Elem()
}

// A collection of values returned by getAckService.
type GetAckServiceResultOutput struct{ *pulumi.OutputState }

func (GetAckServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAckServiceResult)(nil)).Elem()
}

func (o GetAckServiceResultOutput) ToGetAckServiceResultOutput() GetAckServiceResultOutput {
	return o
}

func (o GetAckServiceResultOutput) ToGetAckServiceResultOutputWithContext(ctx context.Context) GetAckServiceResultOutput {
	return o
}

func (o GetAckServiceResultOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAckServiceResult) *string { return v.Enable }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAckServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAckServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current service enable status.
func (o GetAckServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetAckServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetAckServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetAckServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAckServiceResultOutput{})
}
