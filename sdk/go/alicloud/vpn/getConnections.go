// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The VPN connections data source lists lots of VPN connections resource information owned by an Alicloud account.
func GetConnections(ctx *pulumi.Context, args *GetConnectionsArgs, opts ...pulumi.InvokeOption) (*GetConnectionsResult, error) {
	var rv GetConnectionsResult
	err := ctx.Invoke("alicloud:vpn/getConnections:getConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConnections.
type GetConnectionsArgs struct {
	// Use the VPN customer gateway ID as the search key.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// IDs of the VPN connections.
	Ids []string `pulumi:"ids"`
	// A regex string of VPN connection name.
	NameRegex *string `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile *string `pulumi:"outputFile"`
	// Use the VPN gateway ID as the search key.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// A collection of values returned by getConnections.
type GetConnectionsResult struct {
	// A list of VPN connections. Each element contains the following attributes:
	Connections []GetConnectionsConnection `pulumi:"connections"`
	// ID of the VPN customer gateway.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) IDs of the VPN connections.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// (Optional) names of the VPN connections.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// ID of the VPN gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}
