// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the ALB System Security Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.183.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaults, err := alb.GetSystemSecurityPolicies(ctx, &alb.GetSystemSecurityPoliciesArgs{
//				Ids: []string{
//					"tls_cipher_policy_1_0",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("albSystemSecurityPolicyId1", defaults.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetSystemSecurityPolicies(ctx *pulumi.Context, args *GetSystemSecurityPoliciesArgs, opts ...pulumi.InvokeOption) (*GetSystemSecurityPoliciesResult, error) {
	var rv GetSystemSecurityPoliciesResult
	err := ctx.Invoke("alicloud:alb/getSystemSecurityPolicies:getSystemSecurityPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemSecurityPolicies.
type GetSystemSecurityPoliciesArgs struct {
	// A list of System Security Policy IDs.
	Ids        []string               `pulumi:"ids"`
	OutputFile *string                `pulumi:"outputFile"`
	Tags       map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getSystemSecurityPolicies.
type GetSystemSecurityPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of System Security Policy IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of ALB Security Policies. Each element contains the following attributes:
	Policies []GetSystemSecurityPoliciesPolicy `pulumi:"policies"`
	Tags     map[string]interface{}            `pulumi:"tags"`
}

func GetSystemSecurityPoliciesOutput(ctx *pulumi.Context, args GetSystemSecurityPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetSystemSecurityPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemSecurityPoliciesResult, error) {
			args := v.(GetSystemSecurityPoliciesArgs)
			r, err := GetSystemSecurityPolicies(ctx, &args, opts...)
			var s GetSystemSecurityPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemSecurityPoliciesResultOutput)
}

// A collection of arguments for invoking getSystemSecurityPolicies.
type GetSystemSecurityPoliciesOutputArgs struct {
	// A list of System Security Policy IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	Tags       pulumi.MapInput         `pulumi:"tags"`
}

func (GetSystemSecurityPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSecurityPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getSystemSecurityPolicies.
type GetSystemSecurityPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetSystemSecurityPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemSecurityPoliciesResult)(nil)).Elem()
}

func (o GetSystemSecurityPoliciesResultOutput) ToGetSystemSecurityPoliciesResultOutput() GetSystemSecurityPoliciesResultOutput {
	return o
}

func (o GetSystemSecurityPoliciesResultOutput) ToGetSystemSecurityPoliciesResultOutputWithContext(ctx context.Context) GetSystemSecurityPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemSecurityPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemSecurityPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of System Security Policy IDs.
func (o GetSystemSecurityPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemSecurityPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSystemSecurityPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemSecurityPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of ALB Security Policies. Each element contains the following attributes:
func (o GetSystemSecurityPoliciesResultOutput) Policies() GetSystemSecurityPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetSystemSecurityPoliciesResult) []GetSystemSecurityPoliciesPolicy { return v.Policies }).(GetSystemSecurityPoliciesPolicyArrayOutput)
}

func (o GetSystemSecurityPoliciesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSystemSecurityPoliciesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemSecurityPoliciesResultOutput{})
}