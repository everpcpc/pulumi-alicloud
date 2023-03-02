// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Threat Detection Instance available to the user.[What is Instance](https://www.alibabacloud.com/help/en/security-center/latest/what-is-security-center)
//
// > **NOTE:** Available in 1.199.0+
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:threatdetection/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of Instance IDs.
	Ids []string `pulumi:"ids"`
	// The first ID of the resource
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The renewal status of the specified instance. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`.
	RenewStatus *string `pulumi:"renewStatus"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Instance IDs.
	Ids []string `pulumi:"ids"`
	// The first ID of the resource
	InstanceId *string `pulumi:"instanceId"`
	// A list of Instance Entries. Each element contains the following attributes:
	Instances   []GetInstancesInstance `pulumi:"instances"`
	OutputFile  *string                `pulumi:"outputFile"`
	PageNumber  *int                   `pulumi:"pageNumber"`
	PageSize    *int                   `pulumi:"pageSize"`
	RenewStatus *string                `pulumi:"renewStatus"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// A list of Instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The first ID of the resource
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The renewal status of the specified instance. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`.
	RenewStatus pulumi.StringPtrInput `pulumi:"renewStatus"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Instance IDs.
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The first ID of the resource
func (o GetInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// A list of Instance Entries. Each element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetInstancesResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetInstancesResultOutput) RenewStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.RenewStatus }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
