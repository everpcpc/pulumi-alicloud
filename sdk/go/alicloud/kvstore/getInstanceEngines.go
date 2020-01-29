// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kvstore

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the KVStore instance engines resource available info of Alibaba Cloud.
// 
// > **NOTE:** Available in v1.51.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kvstore_instance_engines.html.markdown.
func GetInstanceEngines(ctx *pulumi.Context, args *GetInstanceEnginesArgs, opts ...pulumi.InvokeOption) (*GetInstanceEnginesResult, error) {
	var rv GetInstanceEnginesResult
	err := ctx.Invoke("alicloud:kvstore/getInstanceEngines:getInstanceEngines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceEngines.
type GetInstanceEnginesArgs struct {
	// Database type. Options are `Redis`, `Memcache`. Default to `Redis`.
	Engine *string `pulumi:"engine"`
	// Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) `EngineVersion`. Value of Memcache should be empty.
	EngineVersion *string `pulumi:"engineVersion"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PrePaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	OutputFile *string `pulumi:"outputFile"`
	// The Zone to launch the KVStore instance.
	ZoneId string `pulumi:"zoneId"`
}


// A collection of values returned by getInstanceEngines.
type GetInstanceEnginesResult struct {
	// Database type.
	Engine *string `pulumi:"engine"`
	// KVStore Instance version.
	EngineVersion *string `pulumi:"engineVersion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// A list of KVStore available instance engines. Each element contains the following attributes:
	InstanceEngines []GetInstanceEnginesInstanceEngine `pulumi:"instanceEngines"`
	OutputFile *string `pulumi:"outputFile"`
	// The Zone to launch the KVStore instance.
	ZoneId string `pulumi:"zoneId"`
}

