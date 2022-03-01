// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Resource Manager Policies of the current Alibaba Cloud user.
//
// > **NOTE:**  Available in 1.86.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := resourcemanager.GetPolicies(ctx, &resourcemanager.GetPoliciesArgs{
// 			DescriptionRegex: "tftest_policy",
// 			NameRegex:        pulumi.StringRef("tftest"),
// 			PolicyType:       pulumi.StringRef("Custom"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstPolicyId", example.Policies[0].Id)
// 		return nil
// 	})
// }
// ```
func GetPolicies(ctx *pulumi.Context, args *GetPoliciesArgs, opts ...pulumi.InvokeOption) (*GetPoliciesResult, error) {
	var rv GetPoliciesResult
	err := ctx.Invoke("alicloud:resourcemanager/getPolicies:getPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesArgs struct {
	// A list of Resource Manager Policy IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by policy name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values: `Custom` and `System`.
	PolicyType *string `pulumi:"policyType"`
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of policy IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of policy names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of policies. Each element contains the following attributes:
	Policies   []GetPoliciesPolicy `pulumi:"policies"`
	PolicyType *string             `pulumi:"policyType"`
}

func GetPoliciesOutput(ctx *pulumi.Context, args GetPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPoliciesResult, error) {
			args := v.(GetPoliciesArgs)
			r, err := GetPolicies(ctx, &args, opts...)
			return *r, err
		}).(GetPoliciesResultOutput)
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesOutputArgs struct {
	// A list of Resource Manager Policy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by policy name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values: `Custom` and `System`.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
}

func (GetPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getPolicies.
type GetPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesResult)(nil)).Elem()
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutput() GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutputWithContext(ctx context.Context) GetPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of policy IDs.
func (o GetPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPoliciesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of policy names.
func (o GetPoliciesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of policies. Each element contains the following attributes:
func (o GetPoliciesResultOutput) Policies() GetPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []GetPoliciesPolicy { return v.Policies }).(GetPoliciesPolicyArrayOutput)
}

func (o GetPoliciesResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoliciesResultOutput{})
}
