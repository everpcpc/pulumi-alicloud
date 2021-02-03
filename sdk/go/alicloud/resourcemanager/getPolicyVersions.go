// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Resource Manager Policy Versions of the current Alibaba Cloud user.
//
// > **NOTE:**  Available in 1.85.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := resourcemanager.GetPolicyVersions(ctx, &resourcemanager.GetPolicyVersionsArgs{
// 			PolicyName: "tftest",
// 			PolicyType: "Custom",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstPolicyVersionId", _default.Versions[0].Id)
// 		return nil
// 	})
// }
// ```
func GetPolicyVersions(ctx *pulumi.Context, args *GetPolicyVersionsArgs, opts ...pulumi.InvokeOption) (*GetPolicyVersionsResult, error) {
	var rv GetPolicyVersionsResult
	err := ctx.Invoke("alicloud:resourcemanager/getPolicyVersions:getPolicyVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyVersions.
type GetPolicyVersionsArgs struct {
	// -(Optional, Available in v1.114.0+) Default to `false`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of policy version IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The name of the policy.
	PolicyName string `pulumi:"policyName"`
	// The type of the policy. Valid values:`Custom` and `System`.
	PolicyType string `pulumi:"policyType"`
}

// A collection of values returned by getPolicyVersions.
type GetPolicyVersionsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of policy version IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	PolicyName string   `pulumi:"policyName"`
	PolicyType string   `pulumi:"policyType"`
	// A list of policy versions. Each element contains the following attributes:
	Versions []GetPolicyVersionsVersion `pulumi:"versions"`
}
