// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventbridge

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Event Bridge Event Sources of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.130.0+.
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
// 		opt0 := "the_resource_name"
// 		example, err := eventbridge.GetEventSources(ctx, &eventbridge.GetEventSourcesArgs{
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstEventBridgeEventSourceId", example.Sources[0].Id)
// 		return nil
// 	})
// }
// ```
func GetEventSources(ctx *pulumi.Context, args *GetEventSourcesArgs, opts ...pulumi.InvokeOption) (*GetEventSourcesResult, error) {
	var rv GetEventSourcesResult
	err := ctx.Invoke("alicloud:eventbridge/getEventSources:getEventSources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventSources.
type GetEventSourcesArgs struct {
	// A list of Event Source IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Event Source name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEventSources.
type GetEventSourcesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                  `pulumi:"id"`
	Ids        []string                `pulumi:"ids"`
	NameRegex  *string                 `pulumi:"nameRegex"`
	Names      []string                `pulumi:"names"`
	OutputFile *string                 `pulumi:"outputFile"`
	Sources    []GetEventSourcesSource `pulumi:"sources"`
}
