// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amqp

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Amqp Queues of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.127.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := amqp.GetQueues(ctx, &amqp.GetQueuesArgs{
// 			InstanceId:      "amqp-abc12345",
// 			VirtualHostName: "my-VirtualHost",
// 			Ids: []string{
// 				"my-Queue-1",
// 				"my-Queue-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("amqpQueueId1", ids.Queues[0].Id)
// 		opt0 := "^my-Queue"
// 		nameRegex, err := amqp.GetQueues(ctx, &amqp.GetQueuesArgs{
// 			InstanceId:      "amqp-abc12345",
// 			VirtualHostName: "my-VirtualHost",
// 			NameRegex:       &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("amqpQueueId2", nameRegex.Queues[0].Id)
// 		return nil
// 	})
// }
// ```
func GetQueues(ctx *pulumi.Context, args *GetQueuesArgs, opts ...pulumi.InvokeOption) (*GetQueuesResult, error) {
	var rv GetQueuesResult
	err := ctx.Invoke("alicloud:amqp/getQueues:getQueues", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueues.
type GetQueuesArgs struct {
	// A list of Queue IDs. Its element value is same as Queue Name.
	Ids []string `pulumi:"ids"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by Queue name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of the virtual host.
	VirtualHostName string `pulumi:"virtualHostName"`
}

// A collection of values returned by getQueues.
type GetQueuesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id              string           `pulumi:"id"`
	Ids             []string         `pulumi:"ids"`
	InstanceId      string           `pulumi:"instanceId"`
	NameRegex       *string          `pulumi:"nameRegex"`
	Names           []string         `pulumi:"names"`
	OutputFile      *string          `pulumi:"outputFile"`
	Queues          []GetQueuesQueue `pulumi:"queues"`
	VirtualHostName string           `pulumi:"virtualHostName"`
}