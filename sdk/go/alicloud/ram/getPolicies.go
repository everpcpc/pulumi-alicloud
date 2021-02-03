// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of RAM policies in an Alibaba Cloud account according to the specified filters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ram"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "group1"
// 		opt1 := "policies.txt"
// 		opt2 := "System"
// 		opt3 := "user1"
// 		policiesDs, err := ram.GetPolicies(ctx, &ram.GetPoliciesArgs{
// 			GroupName:  &opt0,
// 			OutputFile: &opt1,
// 			Type:       &opt2,
// 			UserName:   &opt3,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstPolicyName", policiesDs.Policies[0].Name)
// 		return nil
// 	})
// }
// ```
func GetPolicies(ctx *pulumi.Context, args *GetPoliciesArgs, opts ...pulumi.InvokeOption) (*GetPoliciesResult, error) {
	var rv GetPoliciesResult
	err := ctx.Invoke("alicloud:ram/getPolicies:getPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesArgs struct {
	// Default to `true`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// Filter results by a specific group name. Returned policies are attached to the specified group.
	GroupName *string  `pulumi:"groupName"`
	Ids       []string `pulumi:"ids"`
	// A regex string to filter resulting policies by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Filter results by a specific role name. Returned policies are attached to the specified role.
	RoleName *string `pulumi:"roleName"`
	// Filter results by a specific policy type. Valid values are `Custom` and `System`.
	Type *string `pulumi:"type"`
	// Filter results by a specific user name. Returned policies are attached to the specified user.
	UserName *string `pulumi:"userName"`
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	EnableDetails *bool   `pulumi:"enableDetails"`
	GroupName     *string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of ram group names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of policies. Each element contains the following attributes:
	Policies []GetPoliciesPolicy `pulumi:"policies"`
	RoleName *string             `pulumi:"roleName"`
	// Type of the policy.
	Type     *string `pulumi:"type"`
	UserName *string `pulumi:"userName"`
}
