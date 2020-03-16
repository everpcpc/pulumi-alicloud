// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cloud Connect Network Attachment resource. This topic describes how to associate a Smart Access Gateway (SAG) instance with a network instance. You must associate an SAG instance with a network instance if you want to connect the SAG to Alibaba Cloud. You can connect an SAG to Alibaba Cloud through a leased line, the Internet, or the active and standby links.
//
// For information about Cloud Connect Network Attachment and how to use it, see [What is Cloud Connect Network Attachment](https://www.alibabacloud.com/help/doc-detail/124230.htm).
//
// > **NOTE:** Available in 1.64.0+
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cloud_connect_network_attachment.html.markdown.
type NetworkAttachment struct {
	pulumi.CustomResourceState

	// The ID of the CCN instance.
	CcnId pulumi.StringOutput `pulumi:"ccnId"`
	// The ID of the Smart Access Gateway instance.
	SagId pulumi.StringOutput `pulumi:"sagId"`
}

// NewNetworkAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkAttachment(ctx *pulumi.Context,
	name string, args *NetworkAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkAttachment, error) {
	if args == nil || args.CcnId == nil {
		return nil, errors.New("missing required argument 'CcnId'")
	}
	if args == nil || args.SagId == nil {
		return nil, errors.New("missing required argument 'SagId'")
	}
	if args == nil {
		args = &NetworkAttachmentArgs{}
	}
	var resource NetworkAttachment
	err := ctx.RegisterResource("alicloud:cloudconnect/networkAttachment:NetworkAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAttachment gets an existing NetworkAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAttachmentState, opts ...pulumi.ResourceOption) (*NetworkAttachment, error) {
	var resource NetworkAttachment
	err := ctx.ReadResource("alicloud:cloudconnect/networkAttachment:NetworkAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAttachment resources.
type networkAttachmentState struct {
	// The ID of the CCN instance.
	CcnId *string `pulumi:"ccnId"`
	// The ID of the Smart Access Gateway instance.
	SagId *string `pulumi:"sagId"`
}

type NetworkAttachmentState struct {
	// The ID of the CCN instance.
	CcnId pulumi.StringPtrInput
	// The ID of the Smart Access Gateway instance.
	SagId pulumi.StringPtrInput
}

func (NetworkAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAttachmentState)(nil)).Elem()
}

type networkAttachmentArgs struct {
	// The ID of the CCN instance.
	CcnId string `pulumi:"ccnId"`
	// The ID of the Smart Access Gateway instance.
	SagId string `pulumi:"sagId"`
}

// The set of arguments for constructing a NetworkAttachment resource.
type NetworkAttachmentArgs struct {
	// The ID of the CCN instance.
	CcnId pulumi.StringInput
	// The ID of the Smart Access Gateway instance.
	SagId pulumi.StringInput
}

func (NetworkAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAttachmentArgs)(nil)).Elem()
}

