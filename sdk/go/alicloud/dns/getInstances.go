// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:dns/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	DomainType *string `pulumi:"domainType"`
	// A list of instance IDs.
	Ids          []string `pulumi:"ids"`
	Lang         *string  `pulumi:"lang"`
	OutputFile   *string  `pulumi:"outputFile"`
	UserClientIp *string  `pulumi:"userClientIp"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	DomainType *string `pulumi:"domainType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of instances. Each element contains the following attributes:
	Instances    []GetInstancesInstance `pulumi:"instances"`
	Lang         *string                `pulumi:"lang"`
	OutputFile   *string                `pulumi:"outputFile"`
	UserClientIp *string                `pulumi:"userClientIp"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			return *r, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	DomainType pulumi.StringPtrInput `pulumi:"domainType"`
	// A list of instance IDs.
	Ids          pulumi.StringArrayInput `pulumi:"ids"`
	Lang         pulumi.StringPtrInput   `pulumi:"lang"`
	OutputFile   pulumi.StringPtrInput   `pulumi:"outputFile"`
	UserClientIp pulumi.StringPtrInput   `pulumi:"userClientIp"`
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

func (o GetInstancesResultOutput) DomainType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.DomainType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of instance IDs.
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of instances. Each element contains the following attributes:
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

func (o GetInstancesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) UserClientIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.UserClientIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
