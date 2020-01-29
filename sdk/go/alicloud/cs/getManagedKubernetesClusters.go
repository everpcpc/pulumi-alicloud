// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list Container Service Managed Kubernetes Clusters on Alibaba Cloud.
// 
// > **NOTE:** Available in v1.35.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cs_managed_kubernetes_clusters.html.markdown.
func GetManagedKubernetesClusters(ctx *pulumi.Context, args *GetManagedKubernetesClustersArgs, opts ...pulumi.InvokeOption) (*GetManagedKubernetesClustersResult, error) {
	var rv GetManagedKubernetesClustersResult
	err := ctx.Invoke("alicloud:cs/getManagedKubernetesClusters:getManagedKubernetesClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedKubernetesClusters.
type GetManagedKubernetesClustersArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// Cluster IDs to filter.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by cluster name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getManagedKubernetesClusters.
type GetManagedKubernetesClustersResult struct {
	// A list of matched Kubernetes clusters. Each element contains the following attributes:
	Clusters []GetManagedKubernetesClustersCluster `pulumi:"clusters"`
	EnableDetails *bool `pulumi:"enableDetails"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Kubernetes clusters' ids.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of matched Kubernetes clusters' names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
}

