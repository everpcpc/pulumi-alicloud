// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Using this data source can enable OSS service automatically. If the service has been enabled, it will return `Opened`.
//
// For information about OSS and how to use it, see [What is OSS](https://www.alibabacloud.com/help/product/31815.htm).
//
// > **NOTE:** Available in v1.97.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/oss"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "On"
// 		_, err := oss.GetService(ctx, &oss.GetServiceArgs{
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
	err := ctx.Invoke("alicloud:oss/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: "On" or "Off".
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
