// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ddos Bgp Ips of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.180.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ddos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ddos.GetDdosBgpIps(ctx, &ddos.GetDdosBgpIpsArgs{
//				InstanceId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ddosbgpIpId1", ids.Ips[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetDdosBgpIps(ctx *pulumi.Context, args *GetDdosBgpIpsArgs, opts ...pulumi.InvokeOption) (*GetDdosBgpIpsResult, error) {
	var rv GetDdosBgpIpsResult
	err := ctx.Invoke("alicloud:ddos/getDdosBgpIps:getDdosBgpIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDdosBgpIps.
type GetDdosBgpIpsArgs struct {
	// A list of Ip IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the native protection enterprise instance to be operated.
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
	ProductName *string `pulumi:"productName"`
	// The current state of the IP address.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDdosBgpIps.
type GetDdosBgpIpsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string            `pulumi:"id"`
	Ids         []string          `pulumi:"ids"`
	InstanceId  string            `pulumi:"instanceId"`
	Ips         []GetDdosBgpIpsIp `pulumi:"ips"`
	OutputFile  *string           `pulumi:"outputFile"`
	PageNumber  *int              `pulumi:"pageNumber"`
	PageSize    *int              `pulumi:"pageSize"`
	ProductName *string           `pulumi:"productName"`
	Status      *string           `pulumi:"status"`
}

func GetDdosBgpIpsOutput(ctx *pulumi.Context, args GetDdosBgpIpsOutputArgs, opts ...pulumi.InvokeOption) GetDdosBgpIpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDdosBgpIpsResult, error) {
			args := v.(GetDdosBgpIpsArgs)
			r, err := GetDdosBgpIps(ctx, &args, opts...)
			var s GetDdosBgpIpsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDdosBgpIpsResultOutput)
}

// A collection of arguments for invoking getDdosBgpIps.
type GetDdosBgpIpsOutputArgs struct {
	// A list of Ip IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the native protection enterprise instance to be operated.
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
	ProductName pulumi.StringPtrInput `pulumi:"productName"`
	// The current state of the IP address.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetDdosBgpIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosBgpIpsArgs)(nil)).Elem()
}

// A collection of values returned by getDdosBgpIps.
type GetDdosBgpIpsResultOutput struct{ *pulumi.OutputState }

func (GetDdosBgpIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDdosBgpIpsResult)(nil)).Elem()
}

func (o GetDdosBgpIpsResultOutput) ToGetDdosBgpIpsResultOutput() GetDdosBgpIpsResultOutput {
	return o
}

func (o GetDdosBgpIpsResultOutput) ToGetDdosBgpIpsResultOutputWithContext(ctx context.Context) GetDdosBgpIpsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDdosBgpIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDdosBgpIpsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDdosBgpIpsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDdosBgpIpsResultOutput) Ips() GetDdosBgpIpsIpArrayOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) []GetDdosBgpIpsIp { return v.Ips }).(GetDdosBgpIpsIpArrayOutput)
}

func (o GetDdosBgpIpsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDdosBgpIpsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetDdosBgpIpsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetDdosBgpIpsResultOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o GetDdosBgpIpsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDdosBgpIpsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDdosBgpIpsResultOutput{})
}