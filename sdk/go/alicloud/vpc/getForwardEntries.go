// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.37.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/forward_entries.html.markdown.
func GetForwardEntries(ctx *pulumi.Context, args *GetForwardEntriesArgs, opts ...pulumi.InvokeOption) (*GetForwardEntriesResult, error) {
	var rv GetForwardEntriesResult
	err := ctx.Invoke("alicloud:vpc/getForwardEntries:getForwardEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForwardEntries.
type GetForwardEntriesArgs struct {
	// The public IP address.
	ExternalIp *string `pulumi:"externalIp"`
	// The ID of the Forward table.
	ForwardTableId string `pulumi:"forwardTableId"`
	// A list of Forward Entries IDs.
	Ids []string `pulumi:"ids"`
	// The private IP address.
	InternalIp *string `pulumi:"internalIp"`
	// A regex string to filter results by forward entry name.
	NameRegex *string `pulumi:"nameRegex"`
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
}


// A collection of values returned by getForwardEntries.
type GetForwardEntriesResult struct {
	// A list of Forward Entries. Each element contains the following attributes:
	Entries []GetForwardEntriesEntry `pulumi:"entries"`
	// The public IP address.
	ExternalIp *string `pulumi:"externalIp"`
	ForwardTableId string `pulumi:"forwardTableId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Forward Entries IDs.
	Ids []string `pulumi:"ids"`
	// The private IP address.
	InternalIp *string `pulumi:"internalIp"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of Forward Entries names.
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
}

