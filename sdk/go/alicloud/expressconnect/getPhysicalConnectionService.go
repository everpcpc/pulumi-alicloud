// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can enable outbound traffic for an Express Connect circuit automatically. If the service has been opened, it will return opened.
//
// For information about Express Connect and how to use it, see [What is Express Connect](https://www.alibabacloud.com/help/doc-detail/275179.htm).
//
// > **NOTE:** Available in v1.132.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "On"
// 		_, err := expressconnect.GetPhysicalConnectionService(ctx, &expressconnect.GetPhysicalConnectionServiceArgs{
// 			Enable: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPhysicalConnectionService(ctx *pulumi.Context, args *GetPhysicalConnectionServiceArgs, opts ...pulumi.InvokeOption) (*GetPhysicalConnectionServiceResult, error) {
	var rv GetPhysicalConnectionServiceResult
	err := ctx.Invoke("alicloud:expressconnect/getPhysicalConnectionService:getPhysicalConnectionService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPhysicalConnectionService.
type GetPhysicalConnectionServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getPhysicalConnectionService.
type GetPhysicalConnectionServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}