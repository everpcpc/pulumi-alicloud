// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the KVStore instance classes resource available info of Alibaba Cloud.
//
// > **NOTE:** Available in v1.49.0+
func GetInstanceClasses(ctx *pulumi.Context, args *GetInstanceClassesArgs, opts ...pulumi.InvokeOption) (*GetInstanceClassesResult, error) {
	var rv GetInstanceClassesResult
	err := ctx.Invoke("alicloud:kvstore/getInstanceClasses:getInstanceClasses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceClasses.
type GetInstanceClassesArgs struct {
	// The KVStore instance system architecture required by the user. Valid values: `standard`, `cluster` and `rwsplit`.
	Architecture *string `pulumi:"architecture"`
	// The KVStore instance edition type required by the user. Valid values: `Community` and `Enterprise`.
	EditionType *string `pulumi:"editionType"`
	// Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
	Engine *string `pulumi:"engine"`
	// Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
	EngineVersion *string `pulumi:"engineVersion"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The KVStore instance node type required by the user. Valid values: `double`, `single`, `readone`, `readthree` and `readfive`.
	NodeType   *string `pulumi:"nodeType"`
	OutputFile *string `pulumi:"outputFile"`
	// It has been deprecated from 1.68.0.
	PackageType *string `pulumi:"packageType"`
	// It has been deprecated from 1.68.0.
	PerformanceType *string `pulumi:"performanceType"`
	// The KVStore instance series type required by the user. Valid values: `enhancedPerformanceType` and `hybridStorage`.
	SeriesType *string `pulumi:"seriesType"`
	// The number of shard.Valid values: `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, `256`.
	ShardNumber *int    `pulumi:"shardNumber"`
	SortedBy    *string `pulumi:"sortedBy"`
	// It has been deprecated from 1.68.0.
	StorageType *string `pulumi:"storageType"`
	// The Zone to launch the KVStore instance.
	ZoneId string `pulumi:"zoneId"`
}

// A collection of values returned by getInstanceClasses.
type GetInstanceClassesResult struct {
	Architecture *string `pulumi:"architecture"`
	// A list of KVStore available instance classes when the `sortedBy` is "Price". include:
	Classes       []GetInstanceClassesClass `pulumi:"classes"`
	EditionType   *string                   `pulumi:"editionType"`
	Engine        *string                   `pulumi:"engine"`
	EngineVersion *string                   `pulumi:"engineVersion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// A list of KVStore available instance classes.
	InstanceClasses []string `pulumi:"instanceClasses"`
	NodeType        *string  `pulumi:"nodeType"`
	OutputFile      *string  `pulumi:"outputFile"`
	PackageType     *string  `pulumi:"packageType"`
	PerformanceType *string  `pulumi:"performanceType"`
	SeriesType      *string  `pulumi:"seriesType"`
	ShardNumber     *int     `pulumi:"shardNumber"`
	SortedBy        *string  `pulumi:"sortedBy"`
	StorageType     *string  `pulumi:"storageType"`
	ZoneId          string   `pulumi:"zoneId"`
}
