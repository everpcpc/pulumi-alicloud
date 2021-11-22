// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides availability zones for SLB that can be accessed by an Alibaba Cloud account within the region configured in the provider.
//
// > **NOTE:** Available in v1.73.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := slb.GetZones(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:slb/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType *string `pulumi:"availableSlbAddressType"`
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails *bool   `pulumi:"enableDetails"`
	OutputFile    *string `pulumi:"outputFile"`
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	AvailableSlbAddressType      *string `pulumi:"availableSlbAddressType"`
	EnableDetails                *bool   `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of availability zones. Each element contains the following attributes:
	Zones []GetZonesZone `pulumi:"zones"`
}

func GetZonesOutput(ctx *pulumi.Context, args GetZonesOutputArgs, opts ...pulumi.InvokeOption) GetZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZonesResult, error) {
			args := v.(GetZonesArgs)
			r, err := GetZones(ctx, &args, opts...)
			return *r, err
		}).(GetZonesResultOutput)
}

// A collection of arguments for invoking getZones.
type GetZonesOutputArgs struct {
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	AvailableSlbAddressIpVersion pulumi.StringPtrInput `pulumi:"availableSlbAddressIpVersion"`
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType pulumi.StringPtrInput `pulumi:"availableSlbAddressType"`
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	OutputFile    pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesArgs)(nil)).Elem()
}

// A collection of values returned by getZones.
type GetZonesResultOutput struct{ *pulumi.OutputState }

func (GetZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesResult)(nil)).Elem()
}

func (o GetZonesResultOutput) ToGetZonesResultOutput() GetZonesResultOutput {
	return o
}

func (o GetZonesResultOutput) ToGetZonesResultOutputWithContext(ctx context.Context) GetZonesResultOutput {
	return o
}

func (o GetZonesResultOutput) AvailableSlbAddressIpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableSlbAddressIpVersion }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) AvailableSlbAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableSlbAddressType }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zone IDs.
func (o GetZonesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of availability zones. Each element contains the following attributes:
func (o GetZonesResultOutput) Zones() GetZonesZoneArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []GetZonesZone { return v.Zones }).(GetZonesZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZonesResultOutput{})
}
