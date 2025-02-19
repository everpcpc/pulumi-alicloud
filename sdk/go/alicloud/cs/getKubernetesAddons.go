// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of available addons that the cluster can install.
//
// > **NOTE:** Available in 1.150.0+.
// **NOTE:** From version 1.166.0, support for returning custom configuration of kubernetes cluster addon.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := cs.GetKubernetesAddons(ctx, &cs.GetKubernetesAddonsArgs{
//				ClusterId: alicloud_cs_managed_kubernetes.Default[0].Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("addons", _default.Addons)
//			return nil
//		})
//	}
//
// ```
func GetKubernetesAddons(ctx *pulumi.Context, args *GetKubernetesAddonsArgs, opts ...pulumi.InvokeOption) (*GetKubernetesAddonsResult, error) {
	var rv GetKubernetesAddonsResult
	err := ctx.Invoke("alicloud:cs/getKubernetesAddons:getKubernetesAddons", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesAddons.
type GetKubernetesAddonsArgs struct {
	// A list of addons.
	Addons []GetKubernetesAddonsAddon `pulumi:"addons"`
	// The id of kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
	// A list of addon IDs. The id of addon consists of the cluster id and the addon name, with the structure <cluster_ud>:<addon_name>.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by addon name.
	NameRegex *string `pulumi:"nameRegex"`
}

// A collection of values returned by getKubernetesAddons.
type GetKubernetesAddonsResult struct {
	// A list of addons.
	Addons []GetKubernetesAddonsAddon `pulumi:"addons"`
	// The id of kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of addon names.
	Names []string `pulumi:"names"`
}

func GetKubernetesAddonsOutput(ctx *pulumi.Context, args GetKubernetesAddonsOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesAddonsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesAddonsResult, error) {
			args := v.(GetKubernetesAddonsArgs)
			r, err := GetKubernetesAddons(ctx, &args, opts...)
			var s GetKubernetesAddonsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesAddonsResultOutput)
}

// A collection of arguments for invoking getKubernetesAddons.
type GetKubernetesAddonsOutputArgs struct {
	// A list of addons.
	Addons GetKubernetesAddonsAddonArrayInput `pulumi:"addons"`
	// The id of kubernetes cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// A list of addon IDs. The id of addon consists of the cluster id and the addon name, with the structure <cluster_ud>:<addon_name>.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by addon name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
}

func (GetKubernetesAddonsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesAddonsArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesAddons.
type GetKubernetesAddonsResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesAddonsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesAddonsResult)(nil)).Elem()
}

func (o GetKubernetesAddonsResultOutput) ToGetKubernetesAddonsResultOutput() GetKubernetesAddonsResultOutput {
	return o
}

func (o GetKubernetesAddonsResultOutput) ToGetKubernetesAddonsResultOutputWithContext(ctx context.Context) GetKubernetesAddonsResultOutput {
	return o
}

// A list of addons.
func (o GetKubernetesAddonsResultOutput) Addons() GetKubernetesAddonsAddonArrayOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) []GetKubernetesAddonsAddon { return v.Addons }).(GetKubernetesAddonsAddonArrayOutput)
}

// The id of kubernetes cluster.
func (o GetKubernetesAddonsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubernetesAddonsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKubernetesAddonsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetKubernetesAddonsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of addon names.
func (o GetKubernetesAddonsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesAddonsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesAddonsResultOutput{})
}
