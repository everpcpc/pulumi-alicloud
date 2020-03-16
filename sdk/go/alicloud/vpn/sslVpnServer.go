// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpn

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SslVpnServer struct {
	pulumi.CustomResourceState

	// The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none
	Cipher pulumi.StringPtrOutput `pulumi:"cipher"`
	// The CIDR block from which access addresses are allocated to the virtual network interface card of the client.
	ClientIpPool pulumi.StringOutput `pulumi:"clientIpPool"`
	// Specify whether to compress the communication. Valid value: true (default) | false
	Compress pulumi.BoolPtrOutput `pulumi:"compress"`
	// The number of current connections.
	Connections pulumi.IntOutput `pulumi:"connections"`
	// The internet IP of the SSL-VPN server.
	InternetIp pulumi.StringOutput `pulumi:"internetIp"`
	// The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24`.
	LocalSubnet pulumi.StringOutput `pulumi:"localSubnet"`
	// The maximum number of connections.
	MaxConnections pulumi.IntOutput `pulumi:"maxConnections"`
	// The name of the SSL-VPN server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewSslVpnServer registers a new resource with the given unique name, arguments, and options.
func NewSslVpnServer(ctx *pulumi.Context,
	name string, args *SslVpnServerArgs, opts ...pulumi.ResourceOption) (*SslVpnServer, error) {
	if args == nil || args.ClientIpPool == nil {
		return nil, errors.New("missing required argument 'ClientIpPool'")
	}
	if args == nil || args.LocalSubnet == nil {
		return nil, errors.New("missing required argument 'LocalSubnet'")
	}
	if args == nil || args.VpnGatewayId == nil {
		return nil, errors.New("missing required argument 'VpnGatewayId'")
	}
	if args == nil {
		args = &SslVpnServerArgs{}
	}
	var resource SslVpnServer
	err := ctx.RegisterResource("alicloud:vpn/sslVpnServer:SslVpnServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslVpnServer gets an existing SslVpnServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslVpnServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslVpnServerState, opts ...pulumi.ResourceOption) (*SslVpnServer, error) {
	var resource SslVpnServer
	err := ctx.ReadResource("alicloud:vpn/sslVpnServer:SslVpnServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslVpnServer resources.
type sslVpnServerState struct {
	// The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none
	Cipher *string `pulumi:"cipher"`
	// The CIDR block from which access addresses are allocated to the virtual network interface card of the client.
	ClientIpPool *string `pulumi:"clientIpPool"`
	// Specify whether to compress the communication. Valid value: true (default) | false
	Compress *bool `pulumi:"compress"`
	// The number of current connections.
	Connections *int `pulumi:"connections"`
	// The internet IP of the SSL-VPN server.
	InternetIp *string `pulumi:"internetIp"`
	// The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24`.
	LocalSubnet *string `pulumi:"localSubnet"`
	// The maximum number of connections.
	MaxConnections *int `pulumi:"maxConnections"`
	// The name of the SSL-VPN server.
	Name *string `pulumi:"name"`
	// The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].
	Port *int `pulumi:"port"`
	// The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP
	Protocol *string `pulumi:"protocol"`
	// The ID of the VPN gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type SslVpnServerState struct {
	// The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none
	Cipher pulumi.StringPtrInput
	// The CIDR block from which access addresses are allocated to the virtual network interface card of the client.
	ClientIpPool pulumi.StringPtrInput
	// Specify whether to compress the communication. Valid value: true (default) | false
	Compress pulumi.BoolPtrInput
	// The number of current connections.
	Connections pulumi.IntPtrInput
	// The internet IP of the SSL-VPN server.
	InternetIp pulumi.StringPtrInput
	// The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24`.
	LocalSubnet pulumi.StringPtrInput
	// The maximum number of connections.
	MaxConnections pulumi.IntPtrInput
	// The name of the SSL-VPN server.
	Name pulumi.StringPtrInput
	// The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].
	Port pulumi.IntPtrInput
	// The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP
	Protocol pulumi.StringPtrInput
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringPtrInput
}

func (SslVpnServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslVpnServerState)(nil)).Elem()
}

type sslVpnServerArgs struct {
	// The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none
	Cipher *string `pulumi:"cipher"`
	// The CIDR block from which access addresses are allocated to the virtual network interface card of the client.
	ClientIpPool string `pulumi:"clientIpPool"`
	// Specify whether to compress the communication. Valid value: true (default) | false
	Compress *bool `pulumi:"compress"`
	// The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24`.
	LocalSubnet string `pulumi:"localSubnet"`
	// The name of the SSL-VPN server.
	Name *string `pulumi:"name"`
	// The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].
	Port *int `pulumi:"port"`
	// The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP
	Protocol *string `pulumi:"protocol"`
	// The ID of the VPN gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a SslVpnServer resource.
type SslVpnServerArgs struct {
	// The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none
	Cipher pulumi.StringPtrInput
	// The CIDR block from which access addresses are allocated to the virtual network interface card of the client.
	ClientIpPool pulumi.StringInput
	// Specify whether to compress the communication. Valid value: true (default) | false
	Compress pulumi.BoolPtrInput
	// The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24`.
	LocalSubnet pulumi.StringInput
	// The name of the SSL-VPN server.
	Name pulumi.StringPtrInput
	// The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].
	Port pulumi.IntPtrInput
	// The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP
	Protocol pulumi.StringPtrInput
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringInput
}

func (SslVpnServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslVpnServerArgs)(nil)).Elem()
}

