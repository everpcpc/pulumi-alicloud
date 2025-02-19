// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.109.0+.
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
//			example, err := privatelink.GetVpcEndpointServices(ctx, &privatelink.GetVpcEndpointServicesArgs{
//				Ids: []string{
//					"example_value",
//				},
//				NameRegex: pulumi.StringRef("the_resource_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstPrivatelinkVpcEndpointServiceId", example.Services[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetVpcEndpointServices(ctx *pulumi.Context, args *GetVpcEndpointServicesArgs, opts ...pulumi.InvokeOption) (*GetVpcEndpointServicesResult, error) {
	var rv GetVpcEndpointServicesResult
	err := ctx.Invoke("alicloud:privatelink/getVpcEndpointServices:getVpcEndpointServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEndpointServices.
type GetVpcEndpointServicesArgs struct {
	// Whether to automatically accept terminal node connections..
	AutoAcceptConnection *bool `pulumi:"autoAcceptConnection"`
	// A list of Vpc Endpoint Service IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Vpc Endpoint Service name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The business status of the terminal node service..
	ServiceBusinessStatus *string `pulumi:"serviceBusinessStatus"`
	// The Status of Vpc Endpoint Service.
	Status *string `pulumi:"status"`
	// The name of Vpc Endpoint Service.
	VpcEndpointServiceName *string `pulumi:"vpcEndpointServiceName"`
}

// A collection of values returned by getVpcEndpointServices.
type GetVpcEndpointServicesResult struct {
	AutoAcceptConnection *bool `pulumi:"autoAcceptConnection"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string                          `pulumi:"id"`
	Ids                    []string                        `pulumi:"ids"`
	NameRegex              *string                         `pulumi:"nameRegex"`
	Names                  []string                        `pulumi:"names"`
	OutputFile             *string                         `pulumi:"outputFile"`
	ServiceBusinessStatus  *string                         `pulumi:"serviceBusinessStatus"`
	Services               []GetVpcEndpointServicesService `pulumi:"services"`
	Status                 *string                         `pulumi:"status"`
	VpcEndpointServiceName *string                         `pulumi:"vpcEndpointServiceName"`
}

func GetVpcEndpointServicesOutput(ctx *pulumi.Context, args GetVpcEndpointServicesOutputArgs, opts ...pulumi.InvokeOption) GetVpcEndpointServicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcEndpointServicesResult, error) {
			args := v.(GetVpcEndpointServicesArgs)
			r, err := GetVpcEndpointServices(ctx, &args, opts...)
			var s GetVpcEndpointServicesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcEndpointServicesResultOutput)
}

// A collection of arguments for invoking getVpcEndpointServices.
type GetVpcEndpointServicesOutputArgs struct {
	// Whether to automatically accept terminal node connections..
	AutoAcceptConnection pulumi.BoolPtrInput `pulumi:"autoAcceptConnection"`
	// A list of Vpc Endpoint Service IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Vpc Endpoint Service name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The business status of the terminal node service..
	ServiceBusinessStatus pulumi.StringPtrInput `pulumi:"serviceBusinessStatus"`
	// The Status of Vpc Endpoint Service.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The name of Vpc Endpoint Service.
	VpcEndpointServiceName pulumi.StringPtrInput `pulumi:"vpcEndpointServiceName"`
}

func (GetVpcEndpointServicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesArgs)(nil)).Elem()
}

// A collection of values returned by getVpcEndpointServices.
type GetVpcEndpointServicesResultOutput struct{ *pulumi.OutputState }

func (GetVpcEndpointServicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcEndpointServicesResult)(nil)).Elem()
}

func (o GetVpcEndpointServicesResultOutput) ToGetVpcEndpointServicesResultOutput() GetVpcEndpointServicesResultOutput {
	return o
}

func (o GetVpcEndpointServicesResultOutput) ToGetVpcEndpointServicesResultOutputWithContext(ctx context.Context) GetVpcEndpointServicesResultOutput {
	return o
}

func (o GetVpcEndpointServicesResultOutput) AutoAcceptConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *bool { return v.AutoAcceptConnection }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcEndpointServicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcEndpointServicesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVpcEndpointServicesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetVpcEndpointServicesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVpcEndpointServicesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVpcEndpointServicesResultOutput) ServiceBusinessStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *string { return v.ServiceBusinessStatus }).(pulumi.StringPtrOutput)
}

func (o GetVpcEndpointServicesResultOutput) Services() GetVpcEndpointServicesServiceArrayOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) []GetVpcEndpointServicesService { return v.Services }).(GetVpcEndpointServicesServiceArrayOutput)
}

func (o GetVpcEndpointServicesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetVpcEndpointServicesResultOutput) VpcEndpointServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcEndpointServicesResult) *string { return v.VpcEndpointServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcEndpointServicesResultOutput{})
}
