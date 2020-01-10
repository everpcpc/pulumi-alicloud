// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Sag DnatEntry resource. This topic describes how to add a DNAT entry to a Smart Access Gateway (SAG) instance to enable the DNAT function. By using the DNAT function, you can forward requests received by public IP addresses to Alibaba Cloud instances according to custom mapping rules.
// 
// For information about Sag DnatEntry and how to use it, see [What is Sag DnatEntry](https://www.alibabacloud.com/help/doc-detail/124312.htm).
// 
// > **NOTE:** Available in 1.63.0+
// 
// > **NOTE:** Only the following regions suppor. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_dnat_entry.html.markdown.
type DnatEntry struct {
	s *pulumi.ResourceState
}

// NewDnatEntry registers a new resource with the given unique name, arguments, and options.
func NewDnatEntry(ctx *pulumi.Context,
	name string, args *DnatEntryArgs, opts ...pulumi.ResourceOpt) (*DnatEntry, error) {
	if args == nil || args.ExternalPort == nil {
		return nil, errors.New("missing required argument 'ExternalPort'")
	}
	if args == nil || args.InternalIp == nil {
		return nil, errors.New("missing required argument 'InternalIp'")
	}
	if args == nil || args.InternalPort == nil {
		return nil, errors.New("missing required argument 'InternalPort'")
	}
	if args == nil || args.IpProtocol == nil {
		return nil, errors.New("missing required argument 'IpProtocol'")
	}
	if args == nil || args.SagId == nil {
		return nil, errors.New("missing required argument 'SagId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["externalIp"] = nil
		inputs["externalPort"] = nil
		inputs["internalIp"] = nil
		inputs["internalPort"] = nil
		inputs["ipProtocol"] = nil
		inputs["sagId"] = nil
		inputs["type"] = nil
	} else {
		inputs["externalIp"] = args.ExternalIp
		inputs["externalPort"] = args.ExternalPort
		inputs["internalIp"] = args.InternalIp
		inputs["internalPort"] = args.InternalPort
		inputs["ipProtocol"] = args.IpProtocol
		inputs["sagId"] = args.SagId
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("alicloud:rocketmq/dnatEntry:DnatEntry", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DnatEntry{s: s}, nil
}

// GetDnatEntry gets an existing DnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnatEntry(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DnatEntryState, opts ...pulumi.ResourceOpt) (*DnatEntry, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["externalIp"] = state.ExternalIp
		inputs["externalPort"] = state.ExternalPort
		inputs["internalIp"] = state.InternalIp
		inputs["internalPort"] = state.InternalPort
		inputs["ipProtocol"] = state.IpProtocol
		inputs["sagId"] = state.SagId
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("alicloud:rocketmq/dnatEntry:DnatEntry", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DnatEntry{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DnatEntry) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DnatEntry) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The external public IP address.when "type" is "Internet",automatically identify the external ip.
func (r *DnatEntry) ExternalIp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["externalIp"])
}

// The public port.Value range: 1 to 65535 or "any".
func (r *DnatEntry) ExternalPort() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["externalPort"])
}

// The destination private IP address.
func (r *DnatEntry) InternalIp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalIp"])
}

// The destination private port.Value range: 1 to 65535 or "any".
func (r *DnatEntry) InternalPort() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalPort"])
}

// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
func (r *DnatEntry) IpProtocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipProtocol"])
}

// The ID of the SAG instance.
func (r *DnatEntry) SagId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sagId"])
}

// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
func (r *DnatEntry) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering DnatEntry resources.
type DnatEntryState struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp interface{}
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort interface{}
	// The destination private IP address.
	InternalIp interface{}
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort interface{}
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol interface{}
	// The ID of the SAG instance.
	SagId interface{}
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type interface{}
}

// The set of arguments for constructing a DnatEntry resource.
type DnatEntryArgs struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp interface{}
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort interface{}
	// The destination private IP address.
	InternalIp interface{}
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort interface{}
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol interface{}
	// The ID of the SAG instance.
	SagId interface{}
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type interface{}
}
