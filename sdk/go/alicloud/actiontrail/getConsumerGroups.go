// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package actiontrail

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of ALIKAFKA Consumer Groups in an Alibaba Cloud account according to the specified filters.
// 
// > **NOTE:** Available in 1.56.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/alikafka_consumer_groups.html.markdown.
func GetConsumerGroups(ctx *pulumi.Context, args *GetConsumerGroupsArgs, opts ...pulumi.InvokeOption) (*GetConsumerGroupsResult, error) {
	var rv GetConsumerGroupsResult
	err := ctx.Invoke("alicloud:actiontrail/getConsumerGroups:getConsumerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConsumerGroups.
type GetConsumerGroupsArgs struct {
	// A regex string to filter results by the consumer group id. 
	ConsumerIdRegex *string `pulumi:"consumerIdRegex"`
	// ID of the ALIKAFKA Instance that owns the consumer groups.
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getConsumerGroups.
type GetConsumerGroupsResult struct {
	ConsumerIdRegex *string `pulumi:"consumerIdRegex"`
	// A list of consumer group ids.
	ConsumerIds []string `pulumi:"consumerIds"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
}

