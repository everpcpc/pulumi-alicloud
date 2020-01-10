// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The SSL-VPN servers data source lists lots of SSL-VPN servers resource information owned by an Alicloud account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ssl_vpn_servers.html.markdown.
func LookupSslVpnServers(ctx *pulumi.Context, args *GetSslVpnServersArgs) (*GetSslVpnServersResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ids"] = args.Ids
		inputs["nameRegex"] = args.NameRegex
		inputs["outputFile"] = args.OutputFile
		inputs["vpnGatewayId"] = args.VpnGatewayId
	}
	outputs, err := ctx.Invoke("alicloud:vpc/getSslVpnServers:getSslVpnServers", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSslVpnServersResult{
		Ids: outputs["ids"],
		NameRegex: outputs["nameRegex"],
		Names: outputs["names"],
		OutputFile: outputs["outputFile"],
		Servers: outputs["servers"],
		VpnGatewayId: outputs["vpnGatewayId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSslVpnServers.
type GetSslVpnServersArgs struct {
	// IDs of the SSL-VPN servers.
	Ids interface{}
	// A regex string of SSL-VPN server name.
	NameRegex interface{}
	// Save the result to the file.
	OutputFile interface{}
	// Use the VPN gateway ID as the search key.
	VpnGatewayId interface{}
}

// A collection of values returned by getSslVpnServers.
type GetSslVpnServersResult struct {
	// A list of SSL-VPN server IDs.
	Ids interface{}
	NameRegex interface{}
	// A list of SSL-VPN server names.
	Names interface{}
	OutputFile interface{}
	// A list of SSL-VPN servers. Each element contains the following attributes:
	Servers interface{}
	// The ID of the VPN gateway instance.
	VpnGatewayId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
