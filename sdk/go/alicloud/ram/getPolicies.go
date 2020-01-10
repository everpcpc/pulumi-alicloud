// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of RAM policies in an Alibaba Cloud account according to the specified filters.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ram_policies.html.markdown.
func LookupPolicies(ctx *pulumi.Context, args *GetPoliciesArgs) (*GetPoliciesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["groupName"] = args.GroupName
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["roleName"] = args.RoleName
		inputs["type"] = args.Type
		inputs["userName"] = args.UserName
	}
	outputs, err := ctx.Invoke("alicloud:ram/getPolicies:getPolicies", inputs)
	if err != nil {
		return nil, err
	}
	return &GetPoliciesResult{
		GroupName: outputs["groupName"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Policies: outputs["policies"],
		RoleName: outputs["roleName"],
		Type: outputs["type"],
		UserName: outputs["userName"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesArgs struct {
	// Filter results by a specific group name. Returned policies are attached to the specified group.
	GroupName interface{}
	// A regex string to filter resulting policies by name.
	NameRegex interface{}
	OutputFile interface{}
	// Filter results by a specific role name. Returned policies are attached to the specified role.
	RoleName interface{}
	// Filter results by a specific policy type. Valid values are `Custom` and `System`.
	Type interface{}
	// Filter results by a specific user name. Returned policies are attached to the specified user.
	UserName interface{}
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	GroupName interface{}
	NameRegex interface{}
	// A list of ram group names.
	Names interface{}
	OutputFile interface{}
	// A list of policies. Each element contains the following attributes:
	Policies interface{}
	RoleName interface{}
	// Type of the policy.
	Type interface{}
	UserName interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
