// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package emr

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `emr.getMainVersions` data source provides a collection of emr 
// main versions available in Alibaba Cloud account when create a emr cluster.
//
// > **NOTE:** Available in 1.59.0+
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/emr_main_versions.html.markdown.
func GetMainVersions(ctx *pulumi.Context, args *GetMainVersionsArgs, opts ...pulumi.InvokeOption) (*GetMainVersionsResult, error) {
	var rv GetMainVersionsResult
	err := ctx.Invoke("alicloud:emr/getMainVersions:getMainVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMainVersions.
type GetMainVersionsArgs struct {
	// The supported clusterType of this emr version.
	// Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]
	ClusterTypes []string `pulumi:"clusterTypes"`
	// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
	EmrVersion *string `pulumi:"emrVersion"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getMainVersions.
type GetMainVersionsResult struct {
	ClusterTypes []string `pulumi:"clusterTypes"`
	// The version of the emr cluster instance.
	EmrVersion *string `pulumi:"emrVersion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of emr instance types IDs. 
	Ids []string `pulumi:"ids"`
	// A list of versions of the emr cluster instance. Each element contains the following attributes:
	MainVersions []GetMainVersionsMainVersion `pulumi:"mainVersions"`
	OutputFile *string `pulumi:"outputFile"`
}

