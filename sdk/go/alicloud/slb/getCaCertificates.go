// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package slb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the CA certificate list.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_ca_certificates.html.markdown.
func GetCaCertificates(ctx *pulumi.Context, args *GetCaCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCaCertificatesResult, error) {
	var rv GetCaCertificatesResult
	err := ctx.Invoke("alicloud:slb/getCaCertificates:getCaCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCaCertificates.
type GetCaCertificatesArgs struct {
	// A list of ca certificates IDs to filter results.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by ca certificate name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which ca certificates belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getCaCertificates.
type GetCaCertificatesResult struct {
	// A list of SLB ca certificates. Each element contains the following attributes:
	Certificates []GetCaCertificatesCertificate `pulumi:"certificates"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB ca certificates IDs.
	Ids []string `pulumi:"ids"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of SLB ca certificates names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// The resource group Id of CA certificate.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// (Available in v1.66.0+) A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

