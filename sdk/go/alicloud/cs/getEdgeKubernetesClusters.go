// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list Container Service Edge Kubernetes Clusters on Alibaba Cloud.
//
// > **NOTE:** Available in v1.103.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		k8sClusters, err := cs.GetEdgeKubernetesClusters(ctx, &cs.GetEdgeKubernetesClustersArgs{
// 			NameRegex:  pulumi.StringRef("my-first-k8s"),
// 			OutputFile: pulumi.StringRef("my-first-k8s-json"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("output", k8sClusters.Clusters)
// 		return nil
// 	})
// }
// ```
func GetEdgeKubernetesClusters(ctx *pulumi.Context, args *GetEdgeKubernetesClustersArgs, opts ...pulumi.InvokeOption) (*GetEdgeKubernetesClustersResult, error) {
	var rv GetEdgeKubernetesClustersResult
	err := ctx.Invoke("alicloud:cs/getEdgeKubernetesClusters:getEdgeKubernetesClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgeKubernetesClusters.
type GetEdgeKubernetesClustersArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// Cluster IDs to filter.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by cluster name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEdgeKubernetesClusters.
type GetEdgeKubernetesClustersResult struct {
	// A list of matched Kubernetes clusters. Each element contains the following attributes:
	Clusters      []GetEdgeKubernetesClustersCluster `pulumi:"clusters"`
	EnableDetails *bool                              `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Kubernetes clusters' ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of matched Kubernetes clusters' names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetEdgeKubernetesClustersOutput(ctx *pulumi.Context, args GetEdgeKubernetesClustersOutputArgs, opts ...pulumi.InvokeOption) GetEdgeKubernetesClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEdgeKubernetesClustersResult, error) {
			args := v.(GetEdgeKubernetesClustersArgs)
			r, err := GetEdgeKubernetesClusters(ctx, &args, opts...)
			return *r, err
		}).(GetEdgeKubernetesClustersResultOutput)
}

// A collection of arguments for invoking getEdgeKubernetesClusters.
type GetEdgeKubernetesClustersOutputArgs struct {
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// Cluster IDs to filter.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by cluster name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetEdgeKubernetesClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeKubernetesClustersArgs)(nil)).Elem()
}

// A collection of values returned by getEdgeKubernetesClusters.
type GetEdgeKubernetesClustersResultOutput struct{ *pulumi.OutputState }

func (GetEdgeKubernetesClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeKubernetesClustersResult)(nil)).Elem()
}

func (o GetEdgeKubernetesClustersResultOutput) ToGetEdgeKubernetesClustersResultOutput() GetEdgeKubernetesClustersResultOutput {
	return o
}

func (o GetEdgeKubernetesClustersResultOutput) ToGetEdgeKubernetesClustersResultOutputWithContext(ctx context.Context) GetEdgeKubernetesClustersResultOutput {
	return o
}

// A list of matched Kubernetes clusters. Each element contains the following attributes:
func (o GetEdgeKubernetesClustersResultOutput) Clusters() GetEdgeKubernetesClustersClusterArrayOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) []GetEdgeKubernetesClustersCluster { return v.Clusters }).(GetEdgeKubernetesClustersClusterArrayOutput)
}

func (o GetEdgeKubernetesClustersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEdgeKubernetesClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of matched Kubernetes clusters' ids.
func (o GetEdgeKubernetesClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEdgeKubernetesClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of matched Kubernetes clusters' names.
func (o GetEdgeKubernetesClustersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEdgeKubernetesClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgeKubernetesClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEdgeKubernetesClustersResultOutput{})
}
