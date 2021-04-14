// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	var rv GetClustersResult
	err := ctx.Invoke("alicloud:adb/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// The description of the ADB cluster.
	Description *string `pulumi:"description"`
	// A regex string to filter results by cluster description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	// A list of ADB cluster IDs.
	Ids             []string `pulumi:"ids"`
	OutputFile      *string  `pulumi:"outputFile"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	// The status of the cluster. Valid values: `Preparing`, `Creating`, `Restoring`, `Running`, `Deleting`, `ClassChanging`, `NetAddressCreating`, `NetAddressDeleting`. For more information, see [Cluster status](https://www.alibabacloud.com/help/doc-detail/143075.htm).
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// A list of ADB clusters. Each element contains the following attributes:
	Clusters []GetClustersCluster `pulumi:"clusters"`
	// The description of the ADB cluster.
	Description      *string `pulumi:"description"`
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// A list of ADB cluster descriptions.
	Descriptions  []string `pulumi:"descriptions"`
	EnableDetails *bool    `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of ADB cluster IDs.
	Ids             []string `pulumi:"ids"`
	OutputFile      *string  `pulumi:"outputFile"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	// Status of the cluster.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
}
