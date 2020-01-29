// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package alicloud

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the identity of the current user.
// 
// > **NOTE:** Available in 1.65.0+.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/caller_identity.html.markdown.
func GetCallerIdentity(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCallerIdentityResult, error) {
	var rv GetCallerIdentityResult
	err := ctx.Invoke("alicloud:index/getCallerIdentity:getCallerIdentity", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCallerIdentity.
type GetCallerIdentityResult struct {
	// Account ID.
	AccountId string `pulumi:"accountId"`
	// The Alibaba Cloud Resource Name (ARN) of the user making the call.
	Arn string `pulumi:"arn"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The type of the princiapal. RAMUser for users.
	IdentityType string `pulumi:"identityType"`
}

