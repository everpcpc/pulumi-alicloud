// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list Container Registry Enterprise Edition namespaces on Alibaba Cloud.
//
// > **NOTE:** Available in v1.86.0+
func GetRegistryEnterpriseNamespaces(ctx *pulumi.Context, args *GetRegistryEnterpriseNamespacesArgs, opts ...pulumi.InvokeOption) (*GetRegistryEnterpriseNamespacesResult, error) {
	var rv GetRegistryEnterpriseNamespacesResult
	err := ctx.Invoke("alicloud:cs/getRegistryEnterpriseNamespaces:getRegistryEnterpriseNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryEnterpriseNamespaces.
type GetRegistryEnterpriseNamespacesArgs struct {
	// A list of ids to filter results by namespace id.
	Ids []string `pulumi:"ids"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by namespace name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRegistryEnterpriseNamespaces.
type GetRegistryEnterpriseNamespacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.
	Ids []string `pulumi:"ids"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of namespace names.
	Names []string `pulumi:"names"`
	// A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:
	Namespaces []GetRegistryEnterpriseNamespacesNamespace `pulumi:"namespaces"`
	OutputFile *string                                    `pulumi:"outputFile"`
}