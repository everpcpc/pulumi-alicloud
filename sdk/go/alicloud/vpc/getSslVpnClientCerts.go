// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The SSL-VPN client certificates data source lists lots of SSL-VPN client certificates resource information owned by an Alicloud account.
func GetSslVpnClientCerts(ctx *pulumi.Context, args *GetSslVpnClientCertsArgs, opts ...pulumi.InvokeOption) (*GetSslVpnClientCertsResult, error) {
	var rv GetSslVpnClientCertsResult
	err := ctx.Invoke("alicloud:vpc/getSslVpnClientCerts:getSslVpnClientCerts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSslVpnClientCerts.
type GetSslVpnClientCertsArgs struct {
	// IDs of the SSL-VPN client certificates.
	Ids []string `pulumi:"ids"`
	// A regex string of SSL-VPN client certificate name.
	NameRegex *string `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile *string `pulumi:"outputFile"`
	// Use the SSL-VPN server ID as the search key.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
}

// A collection of values returned by getSslVpnClientCerts.
type GetSslVpnClientCertsResult struct {
	Certs []GetSslVpnClientCertsCert `pulumi:"certs"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SSL-VPN client cert IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of SSL-VPN client cert names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// ID of the SSL-VPN Server.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
}
