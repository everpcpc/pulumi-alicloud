// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Snat Entries owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.37.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/snat_entries.html.markdown.
func GetSnatEntries(ctx *pulumi.Context, args *GetSnatEntriesArgs, opts ...pulumi.InvokeOption) (*GetSnatEntriesResult, error) {
	var rv GetSnatEntriesResult
	err := ctx.Invoke("alicloud:vpc/getSnatEntries:getSnatEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnatEntries.
type GetSnatEntriesArgs struct {
	// A list of Snat Entries IDs.
	Ids []string `pulumi:"ids"`
	OutputFile *string `pulumi:"outputFile"`
	// The public IP of the Snat Entry.
	SnatIp *string `pulumi:"snatIp"`
	// The ID of the Snat table.
	SnatTableId string `pulumi:"snatTableId"`
	// The source CIDR block of the Snat Entry.
	SourceCidr *string `pulumi:"sourceCidr"`
}


// A collection of values returned by getSnatEntries.
type GetSnatEntriesResult struct {
	// A list of Snat Entries. Each element contains the following attributes:
	Entries []GetSnatEntriesEntry `pulumi:"entries"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of Snat Entries IDs.
	Ids []string `pulumi:"ids"`
	OutputFile *string `pulumi:"outputFile"`
	// The public IP of the Snat Entry.
	SnatIp *string `pulumi:"snatIp"`
	SnatTableId string `pulumi:"snatTableId"`
	// The source CIDR block of the Snat Entry.
	SourceCidr *string `pulumi:"sourceCidr"`
}

