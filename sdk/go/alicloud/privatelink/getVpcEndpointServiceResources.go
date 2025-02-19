// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Privatelink Vpc Endpoint Service Resources of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.110.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/privatelink"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := privatelink.GetVpcEndpointServiceResources(ctx, &privatelink.GetVpcEndpointServiceResourcesArgs{
//				ServiceId: "epsrv-gw8ii1xxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstPrivatelinkVpcEndpointServiceResourceId", example.Resources[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetVpcEndpointServiceResources(ctx *pulumi.Context, args *GetVpcEndpointServiceResourcesArgs, opts ...pulumi.InvokeOption) (*GetVpcEndpointServiceResourcesResult, error) {
	var rv GetVpcEndpointServiceResourcesResult
	err := ctx.Invoke("alicloud:privatelink/getVpcEndpointServiceResources:getVpcEndpointServiceResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEndpointServiceResources.
type GetVpcEndpointServiceResourcesArgs struct {
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of Vpc Endpoint Service.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by getVpcEndpointServiceResources.
type GetVpcEndpointServiceResourcesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                   `pulumi:"id"`
	Ids        []string                                 `pulumi:"ids"`
	OutputFile *string                                  `pulumi:"outputFile"`
	Resources  []GetVpcEndpointServiceResourcesResource `pulumi:"resources"`
	ServiceId  string                                   `pulumi:"serviceId"`
}

func GetVpcEndpointServiceResourcesOutput(ctx *pulumi.Context, args GetVpcEndpointServiceResourcesOutputArgs, opts ...pulumi.InvokeOption) GetVpcEndpointServiceResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcEndpointServiceResourcesResult, error) {
			args := v.(GetVpcEndpointServiceResourcesArgs)
			r, err := GetVpcEndpointServiceResources(ctx, &args, opts...)
			var s GetVpcEndpointServiceResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcEndpointServiceResourcesResultOutput)
}

// A collection of arguments for invoking getVpcEndpointServiceResources.
type GetVpcEndpointServiceResourcesOutputArgs struct {
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of Vpc Endpoint Service.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (GetVpcEndpointServiceResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServiceResourcesArgs)(nil)).Elem()
}

// A collection of values returned by getVpcEndpointServiceResources.
type GetVpcEndpointServiceResourcesResultOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServiceResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServiceResourcesResult)(nil)).Elem()
}

func (o GetVpcEndpointServiceResourcesResultOutput) ToGetVpcEndpointServiceResourcesResultOutput() GetVpcEndpointServiceResourcesResultOutput {
	return o
}

func (o GetVpcEndpointServiceResourcesResultOutput) ToGetVpcEndpointServiceResourcesResultOutputWithContext(ctx context.Context) GetVpcEndpointServiceResourcesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcEndpointServiceResourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServiceResourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcEndpointServiceResourcesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServiceResourcesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVpcEndpointServiceResourcesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServiceResourcesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVpcEndpointServiceResourcesResultOutput) Resources() GetVpcEndpointServiceResourcesResourceArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServiceResourcesResult) []GetVpcEndpointServiceResourcesResource {
		return v.Resources
	}).(GetVpcEndpointServiceResourcesResourceArrayOutput)
}

func (o GetVpcEndpointServiceResourcesResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServiceResourcesResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcEndpointServiceResourcesResultOutput{})
}
