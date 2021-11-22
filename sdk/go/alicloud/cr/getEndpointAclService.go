// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the CR Endpoint Acl Service of the current Alibaba Cloud user.
//
// For information about Event Bridge and how to use it, see [What is CR Endpoint Acl](https://www.alibabacloud.com/help/en/doc-detail/142246.htm).
//
// > **NOTE:** Available in v1.139.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Registry"
// 		_, err := cr.GetEndpointAclService(ctx, &cr.GetEndpointAclServiceArgs{
// 			Enable:       true,
// 			EndpointType: "internet",
// 			InstanceId:   "example_id",
// 			ModuleName:   &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetEndpointAclService(ctx *pulumi.Context, args *GetEndpointAclServiceArgs, opts ...pulumi.InvokeOption) (*GetEndpointAclServiceResult, error) {
	var rv GetEndpointAclServiceResult
	err := ctx.Invoke("alicloud:cr/getEndpointAclService:getEndpointAclService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointAclService.
type GetEndpointAclServiceArgs struct {
	// Whether to enable Acl Service.  Valid values: `true` and `false`.
	Enable bool `pulumi:"enable"`
	// The type of endpoint. Valid values: `internet`.
	EndpointType string `pulumi:"endpointType"`
	// The ID of the CR Instance.
	InstanceId string `pulumi:"instanceId"`
	// The ModuleName. Valid values: `Registry`.
	ModuleName *string `pulumi:"moduleName"`
}

// A collection of values returned by getEndpointAclService.
type GetEndpointAclServiceResult struct {
	Enable       bool   `pulumi:"enable"`
	EndpointType string `pulumi:"endpointType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	ModuleName *string `pulumi:"moduleName"`
	Status     string  `pulumi:"status"`
}

func GetEndpointAclServiceOutput(ctx *pulumi.Context, args GetEndpointAclServiceOutputArgs, opts ...pulumi.InvokeOption) GetEndpointAclServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEndpointAclServiceResult, error) {
			args := v.(GetEndpointAclServiceArgs)
			r, err := GetEndpointAclService(ctx, &args, opts...)
			return *r, err
		}).(GetEndpointAclServiceResultOutput)
}

// A collection of arguments for invoking getEndpointAclService.
type GetEndpointAclServiceOutputArgs struct {
	// Whether to enable Acl Service.  Valid values: `true` and `false`.
	Enable pulumi.BoolInput `pulumi:"enable"`
	// The type of endpoint. Valid values: `internet`.
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// The ID of the CR Instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The ModuleName. Valid values: `Registry`.
	ModuleName pulumi.StringPtrInput `pulumi:"moduleName"`
}

func (GetEndpointAclServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointAclServiceArgs)(nil)).Elem()
}

// A collection of values returned by getEndpointAclService.
type GetEndpointAclServiceResultOutput struct{ *pulumi.OutputState }

func (GetEndpointAclServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointAclServiceResult)(nil)).Elem()
}

func (o GetEndpointAclServiceResultOutput) ToGetEndpointAclServiceResultOutput() GetEndpointAclServiceResultOutput {
	return o
}

func (o GetEndpointAclServiceResultOutput) ToGetEndpointAclServiceResultOutputWithContext(ctx context.Context) GetEndpointAclServiceResultOutput {
	return o
}

func (o GetEndpointAclServiceResultOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) bool { return v.Enable }).(pulumi.BoolOutput)
}

func (o GetEndpointAclServiceResultOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) string { return v.EndpointType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEndpointAclServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEndpointAclServiceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetEndpointAclServiceResultOutput) ModuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) *string { return v.ModuleName }).(pulumi.StringPtrOutput)
}

func (o GetEndpointAclServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEndpointAclServiceResultOutput{})
}
