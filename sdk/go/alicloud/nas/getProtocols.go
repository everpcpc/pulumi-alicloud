// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide  a data source to retrieve the type of protocol used to create NAS file system.
//
// > **NOTE:** Available in 1.42.0
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "protocols.txt"
// 		opt1 := "cn-beijing-e"
// 		_default, err := nas.GetProtocols(ctx, &nas.GetProtocolsArgs{
// 			OutputFile: &opt0,
// 			Type:       "Performance",
// 			ZoneId:     &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("nasProtocolsProtocol", _default.Protocols[0])
// 		return nil
// 	})
// }
// ```
func GetProtocols(ctx *pulumi.Context, args *GetProtocolsArgs, opts ...pulumi.InvokeOption) (*GetProtocolsResult, error) {
	var rv GetProtocolsResult
	err := ctx.Invoke("alicloud:nas/getProtocols:getProtocols", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProtocols.
type GetProtocolsArgs struct {
	OutputFile *string `pulumi:"outputFile"`
	// The file system type. Valid Values: `Performance` and `Capacity`.
	Type string `pulumi:"type"`
	// String to filter results by zone id.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getProtocols.
type GetProtocolsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of supported protocol type..
	Protocols []string `pulumi:"protocols"`
	Type      string   `pulumi:"type"`
	ZoneId    *string  `pulumi:"zoneId"`
}
