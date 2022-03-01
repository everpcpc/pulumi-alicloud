// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecp Key Pairs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.130.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := ecp.GetKeyPairs(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ecpKeyPairId1", ids.Pairs[0].Id)
// 		nameRegex, err := ecp.GetKeyPairs(ctx, &ecp.GetKeyPairsArgs{
// 			NameRegex: pulumi.StringRef("^my-KeyPair"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ecpKeyPairId2", nameRegex.Pairs[0].Id)
// 		return nil
// 	})
// }
// ```
func GetKeyPairs(ctx *pulumi.Context, args *GetKeyPairsArgs, opts ...pulumi.InvokeOption) (*GetKeyPairsResult, error) {
	var rv GetKeyPairsResult
	err := ctx.Invoke("alicloud:ecp/getKeyPairs:getKeyPairs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeyPairs.
type GetKeyPairsArgs struct {
	// A list of Key Pair IDs. Its element value is same as Key Pair Name.
	Ids []string `pulumi:"ids"`
	// The Private Key of the Fingerprint.
	KeyPairFingerPrint *string `pulumi:"keyPairFingerPrint"`
	// A regex string to filter results by Key Pair name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getKeyPairs.
type GetKeyPairsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string            `pulumi:"id"`
	Ids                []string          `pulumi:"ids"`
	KeyPairFingerPrint *string           `pulumi:"keyPairFingerPrint"`
	NameRegex          *string           `pulumi:"nameRegex"`
	Names              []string          `pulumi:"names"`
	OutputFile         *string           `pulumi:"outputFile"`
	Pairs              []GetKeyPairsPair `pulumi:"pairs"`
}

func GetKeyPairsOutput(ctx *pulumi.Context, args GetKeyPairsOutputArgs, opts ...pulumi.InvokeOption) GetKeyPairsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKeyPairsResult, error) {
			args := v.(GetKeyPairsArgs)
			r, err := GetKeyPairs(ctx, &args, opts...)
			return *r, err
		}).(GetKeyPairsResultOutput)
}

// A collection of arguments for invoking getKeyPairs.
type GetKeyPairsOutputArgs struct {
	// A list of Key Pair IDs. Its element value is same as Key Pair Name.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Private Key of the Fingerprint.
	KeyPairFingerPrint pulumi.StringPtrInput `pulumi:"keyPairFingerPrint"`
	// A regex string to filter results by Key Pair name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetKeyPairsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyPairsArgs)(nil)).Elem()
}

// A collection of values returned by getKeyPairs.
type GetKeyPairsResultOutput struct{ *pulumi.OutputState }

func (GetKeyPairsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyPairsResult)(nil)).Elem()
}

func (o GetKeyPairsResultOutput) ToGetKeyPairsResultOutput() GetKeyPairsResultOutput {
	return o
}

func (o GetKeyPairsResultOutput) ToGetKeyPairsResultOutputWithContext(ctx context.Context) GetKeyPairsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKeyPairsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyPairsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKeyPairsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetKeyPairsResultOutput) KeyPairFingerPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.KeyPairFingerPrint }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetKeyPairsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeyPairsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetKeyPairsResultOutput) Pairs() GetKeyPairsPairArrayOutput {
	return o.ApplyT(func(v GetKeyPairsResult) []GetKeyPairsPair { return v.Pairs }).(GetKeyPairsPairArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKeyPairsResultOutput{})
}
