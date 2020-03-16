// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get a list of elastic network interfaces according to the specified filters in an Alibaba Cloud account.
//
// For information about elastic network interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html)
//
// ##  Argument Reference
//
// The following arguments are supported:
//
// * `ids` - (Optional)  A list of ENI IDs.
// * `nameRegex` - (Optional) A regex string to filter results by ENI name.
// * `vpcId` - (Optional) The VPC ID linked to ENIs.
// * `vswitchId` - (Optional) The VSwitch ID linked to ENIs.
// * `privateIp` - (Optional) The primary private IP address of the ENI.
// * `securityGroupId` - (Optional) The security group ID linked to ENIs.
// * `name` - (Optional) The name of the ENIs.
// * `type` - (Optional) The type of ENIs, Only support for "Primary" or "Secondary".
// * `instanceId` - (Optional) The ECS instance ID that the ENI is attached to.
// * `tags` - (Optional) A map of tags assigned to ENIs.
// * `outputFile` - (Optional) The name of output file that saves the filter results.
// * `resourceGroupId` - (Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the network interface belongs.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/network_interfaces.html.markdown.
func GetNetworkInterfaces(ctx *pulumi.Context, args *GetNetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*GetNetworkInterfacesResult, error) {
	var rv GetNetworkInterfacesResult
	err := ctx.Invoke("alicloud:ecs/getNetworkInterfaces:getNetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterfaces.
type GetNetworkInterfacesArgs struct {
	Ids []string `pulumi:"ids"`
	InstanceId *string `pulumi:"instanceId"`
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	PrivateIp *string `pulumi:"privateIp"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	SecurityGroupId *string `pulumi:"securityGroupId"`
	Tags map[string]interface{} `pulumi:"tags"`
	Type *string `pulumi:"type"`
	VpcId *string `pulumi:"vpcId"`
	VswitchId *string `pulumi:"vswitchId"`
}


// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// ID of the instance that the ENI is attached to.
	InstanceId *string `pulumi:"instanceId"`
	// A list of ENIs. Each element contains the following attributes:
	Interfaces []GetNetworkInterfacesInterface `pulumi:"interfaces"`
	NameRegex *string `pulumi:"nameRegex"`
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	// Primary private IP of the ENI.
	PrivateIp *string `pulumi:"privateIp"`
	// The Id of resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// A map of tags assigned to the ENI.
	Tags map[string]interface{} `pulumi:"tags"`
	Type *string `pulumi:"type"`
	// ID of the VPC that the ENI belongs to.
	VpcId *string `pulumi:"vpcId"`
	// ID of the VSwitch that the ENI is linked to.
	VswitchId *string `pulumi:"vswitchId"`
}

