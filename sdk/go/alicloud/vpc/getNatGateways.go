// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of Nat Gateways owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.37.0+.
func GetNatGateways(ctx *pulumi.Context, args *GetNatGatewaysArgs, opts ...pulumi.InvokeOption) (*GetNatGatewaysResult, error) {
	var rv GetNatGatewaysResult
	err := ctx.Invoke("alicloud:vpc/getNatGateways:getNatGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatGateways.
type GetNatGatewaysArgs struct {
	// A list of NAT gateways IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter nat gateways by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getNatGateways.
type GetNatGatewaysResult struct {
	// A list of Nat gateways. Each element contains the following attributes:
	Gateways []GetNatGatewaysGateway `pulumi:"gateways"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of Nat gateways IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Nat gateways names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}
