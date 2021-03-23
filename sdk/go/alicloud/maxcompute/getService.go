// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package maxcompute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **NOTE:** When you open MaxCompute service, you'd better open [DataWorks service](https://www.alibabacloud.com/help/en/product/72772.htm) as well.
//
// Using this data source can open Maxcompute service automatically. If the service has been opened, it will return opened.
//
// For information about Maxcompute and how to use it, see [What is Maxcompute](https://www.alibabacloud.com/help/en/product/27797.htm).
//
// > **NOTE:** Available in v1.117.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/maxcompute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "On"
// 		_, err := maxcompute.GetService(ctx, &maxcompute.GetServiceArgs{
// 			Enable: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("alicloud:maxcompute/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}
