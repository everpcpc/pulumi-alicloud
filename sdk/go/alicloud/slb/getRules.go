// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package slb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the rules associated with a server load balancer listener.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_rules.html.markdown.
func GetRules(ctx *pulumi.Context, args *GetRulesArgs, opts ...pulumi.InvokeOption) (*GetRulesResult, error) {
	var rv GetRulesResult
	err := ctx.Invoke("alicloud:slb/getRules:getRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRules.
type GetRulesArgs struct {
	// SLB listener port.
	FrontendPort int `pulumi:"frontendPort"`
	// A list of rules IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB with listener rules.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by rule name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getRules.
type GetRulesResult struct {
	FrontendPort int `pulumi:"frontendPort"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB listener rules IDs.
	Ids []string `pulumi:"ids"`
	LoadBalancerId string `pulumi:"loadBalancerId"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of SLB listener rules names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of SLB listener rules. Each element contains the following attributes:
	SlbRules []GetRulesSlbRule `pulumi:"slbRules"`
}

