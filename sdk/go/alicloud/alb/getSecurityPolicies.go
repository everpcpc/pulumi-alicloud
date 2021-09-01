// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Alb Security Policies of the current Alibaba Cloud user.
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
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := alb.GetSecurityPolicies(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albSecurityPolicyId1", ids.Policies[0].Id)
// 		opt0 := "^my-SecurityPolicy"
// 		nameRegex, err := alb.GetSecurityPolicies(ctx, &alb.GetSecurityPoliciesArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albSecurityPolicyId2", nameRegex.Policies[0].Id)
// 		return nil
// 	})
// }
// ```
func GetSecurityPolicies(ctx *pulumi.Context, args *GetSecurityPoliciesArgs, opts ...pulumi.InvokeOption) (*GetSecurityPoliciesResult, error) {
	var rv GetSecurityPoliciesResult
	err := ctx.Invoke("alicloud:alb/getSecurityPolicies:getSecurityPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityPolicies.
type GetSecurityPoliciesArgs struct {
	// A list of Security Policy IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Security Policy name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security policy ids.
	SecurityPolicyIds []string `pulumi:"securityPolicyIds"`
	// The name of the resource. The name must be 2 to 128 characters in length and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
	SecurityPolicyName *string `pulumi:"securityPolicyName"`
	// The status of the resource.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getSecurityPolicies.
type GetSecurityPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                      `pulumi:"id"`
	Ids                []string                    `pulumi:"ids"`
	NameRegex          *string                     `pulumi:"nameRegex"`
	Names              []string                    `pulumi:"names"`
	OutputFile         *string                     `pulumi:"outputFile"`
	Policies           []GetSecurityPoliciesPolicy `pulumi:"policies"`
	ResourceGroupId    *string                     `pulumi:"resourceGroupId"`
	SecurityPolicyIds  []string                    `pulumi:"securityPolicyIds"`
	SecurityPolicyName *string                     `pulumi:"securityPolicyName"`
	Status             *string                     `pulumi:"status"`
	Tags               map[string]interface{}      `pulumi:"tags"`
}