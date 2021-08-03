// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventbridge

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.129.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eventbridge"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := eventbridge.GetEventBuses(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eventBridgeEventBusId1", ids.Buses[0].Id)
// 		opt0 := "^my-EventBus"
// 		nameRegex, err := eventbridge.GetEventBuses(ctx, &eventbridge.GetEventBusesArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eventBridgeEventBusId2", nameRegex.Buses[0].Id)
// 		return nil
// 	})
// }
// ```
func GetEventBuses(ctx *pulumi.Context, args *GetEventBusesArgs, opts ...pulumi.InvokeOption) (*GetEventBusesResult, error) {
	var rv GetEventBusesResult
	err := ctx.Invoke("alicloud:eventbridge/getEventBuses:getEventBuses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventBuses.
type GetEventBusesArgs struct {
	// The event bus type.
	EventBusType *string `pulumi:"eventBusType"`
	// A list of Event Bus IDs. Its element value is same as Event Bus Name.
	Ids []string `pulumi:"ids"`
	// The name prefix.
	NamePrefix *string `pulumi:"namePrefix"`
	// A regex string to filter results by Event Bus name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEventBuses.
type GetEventBusesResult struct {
	Buses        []GetEventBusesBus `pulumi:"buses"`
	EventBusType *string            `pulumi:"eventBusType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NamePrefix *string  `pulumi:"namePrefix"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}