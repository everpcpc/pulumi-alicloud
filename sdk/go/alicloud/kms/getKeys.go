// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of KMS keys in an Alibaba Cloud account according to the specified filters.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/kms_keys.html.markdown.
func GetKeys(ctx *pulumi.Context, args *GetKeysArgs, opts ...pulumi.InvokeOption) (*GetKeysResult, error) {
	var rv GetKeysResult
	err := ctx.Invoke("alicloud:kms/getKeys:getKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeys.
type GetKeysArgs struct {
	// A regex string to filter the results by the KMS key description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// A list of KMS key IDs.
	Ids []string `pulumi:"ids"`
	OutputFile *string `pulumi:"outputFile"`
	// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
	Status *string `pulumi:"status"`
}


// A collection of values returned by getKeys.
type GetKeysResult struct {
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of KMS key IDs.
	Ids []string `pulumi:"ids"`
	// A list of KMS keys. Each element contains the following attributes:
	Keys []GetKeysKey `pulumi:"keys"`
	OutputFile *string `pulumi:"outputFile"`
	// Status of the key. Possible values: `Enabled`, `Disabled` and `PendingDeletion`.
	Status *string `pulumi:"status"`
}

