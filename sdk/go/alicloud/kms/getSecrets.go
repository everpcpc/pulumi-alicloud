// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of KMS Secrets in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in v1.86.0+.
func GetSecrets(ctx *pulumi.Context, args *GetSecretsArgs, opts ...pulumi.InvokeOption) (*GetSecretsResult, error) {
	var rv GetSecretsResult
	err := ctx.Invoke("alicloud:kms/getSecrets:getSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecrets.
type GetSecretsArgs struct {
	// Whether to include the predetermined resource tag in the return value. Default to `false`.
	FetchTags *bool `pulumi:"fetchTags"`
	// A list of KMS Secret ids. The value is same as KMS secret_name.
	Ids []string `pulumi:"ids"`
	// A regex string to filter the results by the KMS secret_name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getSecrets.
type GetSecretsResult struct {
	FetchTags *bool `pulumi:"fetchTags"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Kms Secret ids. The value is same as KMS secret_name.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of KMS Secret names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of KMS Secrets. Each element contains the following attributes:
	Secrets []GetSecretsSecret `pulumi:"secrets"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}