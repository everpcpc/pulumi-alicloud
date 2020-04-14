// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `kvstore.getInstances` data source provides a collection of kvstore instances available in Alicloud account.
// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:kvstore/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of RKV instance IDs.
	Ids           []string `pulumi:"ids"`
	InstanceClass *string  `pulumi:"instanceClass"`
	// Database type. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
	InstanceType *string `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Status of the instance.
	// * `instanceClass`- (Optional) Type of the applied ApsaraDB for Redis instance.
	// For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
	Status *string `pulumi:"status"`
	// Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{"key1":"value1"}`.
	Tags map[string]interface{} `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId *string `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of RKV instance IDs.
	Ids           []string `pulumi:"ids"`
	InstanceClass *string  `pulumi:"instanceClass"`
	// (Optional) Database type. Options are `Memcache`, and `Redis`. If no value is specified, all types are returned.
	// * `instanceClass`- (Optional) Type of the applied ApsaraDB for Redis instance.
	// For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
	InstanceType *string `pulumi:"instanceType"`
	// A list of RKV instances. Its every element contains the following attributes:
	Instances  []GetInstancesInstance `pulumi:"instances"`
	NameRegex  *string                `pulumi:"nameRegex"`
	Names      []string               `pulumi:"names"`
	OutputFile *string                `pulumi:"outputFile"`
	// Status of the instance.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
	// VPC ID the instance belongs to.
	VpcId *string `pulumi:"vpcId"`
	// VSwitch ID the instance belongs to.
	VswitchId *string `pulumi:"vswitchId"`
}
