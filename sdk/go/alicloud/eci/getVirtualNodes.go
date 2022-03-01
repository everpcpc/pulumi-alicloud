// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.145.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := eci.GetVirtualNodes(ctx, &eci.GetVirtualNodesArgs{
// 			Ids: []string{
// 				"example_value-1",
// 				"example_value-2",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eciVirtualNodeId1", ids.Nodes[0].Id)
// 		nameRegex, err := eci.GetVirtualNodes(ctx, &eci.GetVirtualNodesArgs{
// 			NameRegex: pulumi.StringRef("^my-VirtualNode"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("eciVirtualNodeId2", nameRegex.Nodes[0].Id)
// 		return nil
// 	})
// }
// ```
func GetVirtualNodes(ctx *pulumi.Context, args *GetVirtualNodesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNodesResult, error) {
	var rv GetVirtualNodesResult
	err := ctx.Invoke("alicloud:eci/getVirtualNodes:getVirtualNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNodes.
type GetVirtualNodesArgs struct {
	// A list of Virtual Node IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Virtual Node name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security group ID.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The Status of the virtual node.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the virtual node.
	VirtualNodeName *string `pulumi:"virtualNodeName"`
	// The vswitch id.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getVirtualNodes.
type GetVirtualNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id              string                 `pulumi:"id"`
	Ids             []string               `pulumi:"ids"`
	NameRegex       *string                `pulumi:"nameRegex"`
	Names           []string               `pulumi:"names"`
	Nodes           []GetVirtualNodesNode  `pulumi:"nodes"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	SecurityGroupId *string                `pulumi:"securityGroupId"`
	Status          *string                `pulumi:"status"`
	Tags            map[string]interface{} `pulumi:"tags"`
	VirtualNodeName *string                `pulumi:"virtualNodeName"`
	VswitchId       *string                `pulumi:"vswitchId"`
}

func GetVirtualNodesOutput(ctx *pulumi.Context, args GetVirtualNodesOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNodesResult, error) {
			args := v.(GetVirtualNodesArgs)
			r, err := GetVirtualNodes(ctx, &args, opts...)
			return *r, err
		}).(GetVirtualNodesResultOutput)
}

// A collection of arguments for invoking getVirtualNodes.
type GetVirtualNodesOutputArgs struct {
	// A list of Virtual Node IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Virtual Node name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The resource group ID.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The security group ID.
	SecurityGroupId pulumi.StringPtrInput `pulumi:"securityGroupId"`
	// The Status of the virtual node.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The name of the virtual node.
	VirtualNodeName pulumi.StringPtrInput `pulumi:"virtualNodeName"`
	// The vswitch id.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetVirtualNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNodesArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualNodes.
type GetVirtualNodesResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNodesResult)(nil)).Elem()
}

func (o GetVirtualNodesResultOutput) ToGetVirtualNodesResultOutput() GetVirtualNodesResultOutput {
	return o
}

func (o GetVirtualNodesResultOutput) ToGetVirtualNodesResultOutputWithContext(ctx context.Context) GetVirtualNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVirtualNodesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVirtualNodesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVirtualNodesResultOutput) Nodes() GetVirtualNodesNodeArrayOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) []GetVirtualNodesNode { return v.Nodes }).(GetVirtualNodesNodeArrayOutput)
}

func (o GetVirtualNodesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetVirtualNodesResultOutput) VirtualNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.VirtualNodeName }).(pulumi.StringPtrOutput)
}

func (o GetVirtualNodesResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualNodesResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNodesResultOutput{})
}
