// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the ots instance attachments of the current Alibaba Cloud user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ots_instance_attachments.html.markdown.
func LookupInstanceAttachments(ctx *pulumi.Context, args *GetInstanceAttachmentsArgs) (*GetInstanceAttachmentsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["instanceName"] = args.InstanceName
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
	}
	outputs, err := ctx.Invoke("alicloud:oss/getInstanceAttachments:getInstanceAttachments", inputs)
	if err != nil {
		return nil, err
	}
	return &GetInstanceAttachmentsResult{
		Attachments: outputs["attachments"],
		InstanceName: outputs["instanceName"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		VpcIds: outputs["vpcIds"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getInstanceAttachments.
type GetInstanceAttachmentsArgs struct {
	// The name of OTS instance.
	InstanceName interface{}
	// A regex string to filter results by vpc name.
	NameRegex interface{}
	OutputFile interface{}
}

// A collection of values returned by getInstanceAttachments.
type GetInstanceAttachmentsResult struct {
	// A list of instance attachments. Each element contains the following attributes:
	Attachments interface{}
	// The instance name.
	InstanceName interface{}
	NameRegex interface{}
	// A list of vpc names.
	Names interface{}
	OutputFile interface{}
	// A list of vpc ids.
	VpcIds interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
