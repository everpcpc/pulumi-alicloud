// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/adb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		adbClustersDs, err := adb.GetClusters(ctx, &adb.GetClustersArgs{
// 			DescriptionRegex: pulumi.StringRef("am-\\w+"),
// 			Status:           pulumi.StringRef("Running"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstAdbClusterId", adbClustersDs.Clusters[0].Id)
// 		return nil
// 	})
// }
// ```
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
	PageNumber      *int     `pulumi:"pageNumber"`
	PageSize        *int     `pulumi:"pageSize"`
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
	PageNumber      *int     `pulumi:"pageNumber"`
	PageSize        *int     `pulumi:"pageSize"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	// Status of the cluster.
	Status     *string                `pulumi:"status"`
	Tags       map[string]interface{} `pulumi:"tags"`
	TotalCount int                    `pulumi:"totalCount"`
}

func GetClustersOutput(ctx *pulumi.Context, args GetClustersOutputArgs, opts ...pulumi.InvokeOption) GetClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClustersResult, error) {
			args := v.(GetClustersArgs)
			r, err := GetClusters(ctx, &args, opts...)
			return *r, err
		}).(GetClustersResultOutput)
}

// A collection of arguments for invoking getClusters.
type GetClustersOutputArgs struct {
	// The description of the ADB cluster.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A regex string to filter results by cluster description.
	DescriptionRegex pulumi.StringPtrInput `pulumi:"descriptionRegex"`
	EnableDetails    pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// A list of ADB cluster IDs.
	Ids             pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile      pulumi.StringPtrInput   `pulumi:"outputFile"`
	PageNumber      pulumi.IntPtrInput      `pulumi:"pageNumber"`
	PageSize        pulumi.IntPtrInput      `pulumi:"pageSize"`
	ResourceGroupId pulumi.StringPtrInput   `pulumi:"resourceGroupId"`
	// The status of the cluster. Valid values: `Preparing`, `Creating`, `Restoring`, `Running`, `Deleting`, `ClassChanging`, `NetAddressCreating`, `NetAddressDeleting`. For more information, see [Cluster status](https://www.alibabacloud.com/help/doc-detail/143075.htm).
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersArgs)(nil)).Elem()
}

// A collection of values returned by getClusters.
type GetClustersResultOutput struct{ *pulumi.OutputState }

func (GetClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersResult)(nil)).Elem()
}

func (o GetClustersResultOutput) ToGetClustersResultOutput() GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ToGetClustersResultOutputWithContext(ctx context.Context) GetClustersResultOutput {
	return o
}

// A list of ADB clusters. Each element contains the following attributes:
func (o GetClustersResultOutput) Clusters() GetClustersClusterArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []GetClustersCluster { return v.Clusters }).(GetClustersClusterArrayOutput)
}

// The description of the ADB cluster.
func (o GetClustersResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) DescriptionRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.DescriptionRegex }).(pulumi.StringPtrOutput)
}

// A list of ADB cluster descriptions.
func (o GetClustersResultOutput) Descriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Descriptions }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of ADB cluster IDs.
func (o GetClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetClustersResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetClustersResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// Status of the cluster.
func (o GetClustersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetClustersResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetClustersResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetClustersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersResultOutput{})
}
