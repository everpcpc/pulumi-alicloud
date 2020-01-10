// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source provides the listeners related to a server load balancer of the current Alibaba Cloud user.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_listeners.html.markdown.
func LookupListeners(ctx *pulumi.Context, args *GetListenersArgs) (*GetListenersResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["frontendPort"] = args.FrontendPort
		inputs["loadBalancerId"] = args.LoadBalancerId
		inputs["outputFile"] = args.OutputFile
		inputs["protocol"] = args.Protocol
	}
	outputs, err := ctx.Invoke("alicloud:slb/getListeners:getListeners", inputs)
	if err != nil {
		return nil, err
	}
	return &GetListenersResult{
		FrontendPort: outputs["frontendPort"],
		LoadBalancerId: outputs["loadBalancerId"],
		OutputFile: outputs["outputFile"],
		Protocol: outputs["protocol"],
		SlbListeners: outputs["slbListeners"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getListeners.
type GetListenersArgs struct {
	// Filter listeners by the specified frontend port.
	FrontendPort interface{}
	// ID of the SLB with listeners.
	LoadBalancerId interface{}
	OutputFile interface{}
	// Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
	Protocol interface{}
}

// A collection of values returned by getListeners.
type GetListenersResult struct {
	// Frontend port used to receive incoming traffic and distribute it to the backend servers.
	FrontendPort interface{}
	LoadBalancerId interface{}
	OutputFile interface{}
	// Listener protocol. Possible values: `http`, `https`, `tcp` and `udp`.
	Protocol interface{}
	// A list of SLB listeners. Each element contains the following attributes:
	SlbListeners interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
