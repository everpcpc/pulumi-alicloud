// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
// 
// > **NOTE:** Available in 1.37.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/forward_entries.html.markdown.
func LookupForwardEntries(ctx *pulumi.Context, args *GetForwardEntriesArgs) (*GetForwardEntriesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["externalIp"] = args.ExternalIp
		inputs["forwardTableId"] = args.ForwardTableId
		inputs["ids"] = args.Ids
		inputs["internalIp"] = args.InternalIp
		inputs["nameRegex"] = args.NameRegex
		inputs["names"] = args.Names
		inputs["outputFile"] = args.OutputFile
	}
	outputs, err := ctx.Invoke("alicloud:vpc/getForwardEntries:getForwardEntries", inputs)
	if err != nil {
		return nil, err
	}
	return &GetForwardEntriesResult{
		Entries: outputs["entries"],
		ExternalIp: outputs["externalIp"],
		ForwardTableId: outputs["forwardTableId"],
		Ids: outputs["ids"],
		InternalIp: outputs["internalIp"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getForwardEntries.
type GetForwardEntriesArgs struct {
	// The public IP address.
	ExternalIp interface{}
	// The ID of the Forward table.
	ForwardTableId interface{}
	// A list of Forward Entries IDs.
	Ids interface{}
	// The private IP address.
	InternalIp interface{}
	// A regex string to filter results by forward entry name.
	NameRegex interface{}
	Names interface{}
	OutputFile interface{}
}

// A collection of values returned by getForwardEntries.
type GetForwardEntriesResult struct {
	// A list of Forward Entries. Each element contains the following attributes:
	Entries interface{}
	// The public IP address.
	ExternalIp interface{}
	ForwardTableId interface{}
	// A list of Forward Entries IDs.
	Ids interface{}
	// The private IP address.
	InternalIp interface{}
	NameRegex interface{}
	// A list of Forward Entries names.
	Names interface{}
	OutputFile interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
