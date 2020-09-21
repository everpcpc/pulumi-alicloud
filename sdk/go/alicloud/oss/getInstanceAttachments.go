// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the ots instance attachments of the current Alibaba Cloud user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ots"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "testvpc"
// 		opt1 := "attachments.txt"
// 		attachmentsDs, err := ots.GetInstanceAttachments(ctx, &ots.GetInstanceAttachmentsArgs{
// 			InstanceName: "sample-instance",
// 			NameRegex:    &opt0,
// 			OutputFile:   &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstOtsAttachmentId", attachmentsDs.Attachments[0].Id)
// 		return nil
// 	})
// }
// ```
//
// Deprecated: alicloud.oss.getInstanceAttachments has been deprecated in favor of alicloud.ots.getInstanceAttachments
func GetInstanceAttachments(ctx *pulumi.Context, args *GetInstanceAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetInstanceAttachmentsResult, error) {
	var rv GetInstanceAttachmentsResult
	err := ctx.Invoke("alicloud:oss/getInstanceAttachments:getInstanceAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceAttachments.
type GetInstanceAttachmentsArgs struct {
	// The name of OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// A regex string to filter results by vpc name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInstanceAttachments.
type GetInstanceAttachmentsResult struct {
	// A list of instance attachments. Each element contains the following attributes:
	Attachments []GetInstanceAttachmentsAttachment `pulumi:"attachments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instance name.
	InstanceName string  `pulumi:"instanceName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of vpc names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of vpc ids.
	VpcIds []string `pulumi:"vpcIds"`
}
