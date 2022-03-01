// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Global Accelerator (GA) Endpoint Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.113.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ga.GetEndpointGroups(ctx, &ga.GetEndpointGroupsArgs{
// 			AcceleratorId: "example_value",
// 			Ids: []string{
// 				"example_value",
// 			},
// 			NameRegex: pulumi.StringRef("the_resource_name"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstGaEndpointGroupId", example.Groups[0].Id)
// 		return nil
// 	})
// }
// ```
func GetEndpointGroups(ctx *pulumi.Context, args *GetEndpointGroupsArgs, opts ...pulumi.InvokeOption) (*GetEndpointGroupsResult, error) {
	var rv GetEndpointGroupsResult
	err := ctx.Invoke("alicloud:ga/getEndpointGroups:getEndpointGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointGroups.
type GetEndpointGroupsArgs struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
	EndpointGroupType *string `pulumi:"endpointGroupType"`
	// A list of Endpoint Group IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the listener that is associated with the endpoint group.
	ListenerId *string `pulumi:"listenerId"`
	// A regex string to filter results by Endpoint Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the endpoint group.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getEndpointGroups.
type GetEndpointGroupsResult struct {
	AcceleratorId     string                   `pulumi:"acceleratorId"`
	EndpointGroupType *string                  `pulumi:"endpointGroupType"`
	Groups            []GetEndpointGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	ListenerId *string  `pulumi:"listenerId"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetEndpointGroupsOutput(ctx *pulumi.Context, args GetEndpointGroupsOutputArgs, opts ...pulumi.InvokeOption) GetEndpointGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEndpointGroupsResult, error) {
			args := v.(GetEndpointGroupsArgs)
			r, err := GetEndpointGroups(ctx, &args, opts...)
			return *r, err
		}).(GetEndpointGroupsResultOutput)
}

// A collection of arguments for invoking getEndpointGroups.
type GetEndpointGroupsOutputArgs struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId pulumi.StringInput `pulumi:"acceleratorId"`
	// The endpoint group type. Valid values: `default`, `virtual`. Default value is `default`.
	EndpointGroupType pulumi.StringPtrInput `pulumi:"endpointGroupType"`
	// A list of Endpoint Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the listener that is associated with the endpoint group.
	ListenerId pulumi.StringPtrInput `pulumi:"listenerId"`
	// A regex string to filter results by Endpoint Group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the endpoint group.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetEndpointGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getEndpointGroups.
type GetEndpointGroupsResultOutput struct{ *pulumi.OutputState }

func (GetEndpointGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointGroupsResult)(nil)).Elem()
}

func (o GetEndpointGroupsResultOutput) ToGetEndpointGroupsResultOutput() GetEndpointGroupsResultOutput {
	return o
}

func (o GetEndpointGroupsResultOutput) ToGetEndpointGroupsResultOutputWithContext(ctx context.Context) GetEndpointGroupsResultOutput {
	return o
}

func (o GetEndpointGroupsResultOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) string { return v.AcceleratorId }).(pulumi.StringOutput)
}

func (o GetEndpointGroupsResultOutput) EndpointGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) *string { return v.EndpointGroupType }).(pulumi.StringPtrOutput)
}

func (o GetEndpointGroupsResultOutput) Groups() GetEndpointGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) []GetEndpointGroupsGroup { return v.Groups }).(GetEndpointGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEndpointGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEndpointGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEndpointGroupsResultOutput) ListenerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) *string { return v.ListenerId }).(pulumi.StringPtrOutput)
}

func (o GetEndpointGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEndpointGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEndpointGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEndpointGroupsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointGroupsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEndpointGroupsResultOutput{})
}
