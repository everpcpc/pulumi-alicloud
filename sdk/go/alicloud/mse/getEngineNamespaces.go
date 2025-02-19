// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Mse Engine Namespaces of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.166.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mse"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := mse.GetEngineNamespaces(ctx, &mse.GetEngineNamespacesArgs{
//				ClusterId: "example_value",
//				Ids: []string{
//					"example_value",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("mseEngineNamespaceId1", ids.Namespaces[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEngineNamespaces(ctx *pulumi.Context, args *GetEngineNamespacesArgs, opts ...pulumi.InvokeOption) (*GetEngineNamespacesResult, error) {
	var rv GetEngineNamespacesResult
	err := ctx.Invoke("alicloud:mse/getEngineNamespaces:getEngineNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEngineNamespaces.
type GetEngineNamespacesArgs struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The id of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// A list of Engine Namespace IDs. It is formatted to `<cluster_id>:<namespace_id>`.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEngineNamespaces.
type GetEngineNamespacesResult struct {
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	ClusterId      string  `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                         `pulumi:"id"`
	Ids        []string                       `pulumi:"ids"`
	Namespaces []GetEngineNamespacesNamespace `pulumi:"namespaces"`
	OutputFile *string                        `pulumi:"outputFile"`
}

func GetEngineNamespacesOutput(ctx *pulumi.Context, args GetEngineNamespacesOutputArgs, opts ...pulumi.InvokeOption) GetEngineNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEngineNamespacesResult, error) {
			args := v.(GetEngineNamespacesArgs)
			r, err := GetEngineNamespaces(ctx, &args, opts...)
			var s GetEngineNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEngineNamespacesResultOutput)
}

// A collection of arguments for invoking getEngineNamespaces.
type GetEngineNamespacesOutputArgs struct {
	// The language type of the returned information. Valid values: `zh`, `en`.
	AcceptLanguage pulumi.StringPtrInput `pulumi:"acceptLanguage"`
	// The id of the cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// A list of Engine Namespace IDs. It is formatted to `<cluster_id>:<namespace_id>`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetEngineNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEngineNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getEngineNamespaces.
type GetEngineNamespacesResultOutput struct{ *pulumi.OutputState }

func (GetEngineNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEngineNamespacesResult)(nil)).Elem()
}

func (o GetEngineNamespacesResultOutput) ToGetEngineNamespacesResultOutput() GetEngineNamespacesResultOutput {
	return o
}

func (o GetEngineNamespacesResultOutput) ToGetEngineNamespacesResultOutputWithContext(ctx context.Context) GetEngineNamespacesResultOutput {
	return o
}

func (o GetEngineNamespacesResultOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o GetEngineNamespacesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEngineNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEngineNamespacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEngineNamespacesResultOutput) Namespaces() GetEngineNamespacesNamespaceArrayOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) []GetEngineNamespacesNamespace { return v.Namespaces }).(GetEngineNamespacesNamespaceArrayOutput)
}

func (o GetEngineNamespacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEngineNamespacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEngineNamespacesResultOutput{})
}
