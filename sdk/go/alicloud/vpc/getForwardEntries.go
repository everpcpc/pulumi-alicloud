// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in 1.37.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "forward-entry-config-example-name"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 			CidrBlock:        pulumi.String("172.16.0.0/21"),
// 			VpcId:            defaultNetwork.ID(),
// 			VswitchName:      pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultNatGateway, err := vpc.NewNatGateway(ctx, "defaultNatGateway", &vpc.NatGatewayArgs{
// 			Specification: pulumi.String("Small"),
// 			VpcId:         defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultEip, err := ecs.NewEip(ctx, "defaultEip", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewEipAssociation(ctx, "defaultEipAssociation", &ecs.EipAssociationArgs{
// 			AllocationId: defaultEip.ID(),
// 			InstanceId:   defaultNatGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultForwardEntry, err := vpc.NewForwardEntry(ctx, "defaultForwardEntry", &vpc.ForwardEntryArgs{
// 			ExternalIp:     defaultEip.IpAddress,
// 			ExternalPort:   pulumi.String("80"),
// 			ForwardTableId: defaultNatGateway.ForwardTableIds,
// 			InternalIp:     pulumi.String("172.16.0.3"),
// 			InternalPort:   pulumi.String("8080"),
// 			IpProtocol:     pulumi.String("tcp"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetForwardEntries(ctx *pulumi.Context, args *GetForwardEntriesArgs, opts ...pulumi.InvokeOption) (*GetForwardEntriesResult, error) {
	var rv GetForwardEntriesResult
	err := ctx.Invoke("alicloud:vpc/getForwardEntries:getForwardEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getForwardEntries.
type GetForwardEntriesArgs struct {
	// The public IP address.
	ExternalIp *string `pulumi:"externalIp"`
	// The ID of the Forward table.
	ForwardTableId string `pulumi:"forwardTableId"`
	// A list of Forward Entries IDs.
	Ids []string `pulumi:"ids"`
	// The private IP address.
	InternalIp *string `pulumi:"internalIp"`
	// A regex string to filter results by forward entry name.
	NameRegex *string `pulumi:"nameRegex"`
	// A list of Forward Entries names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

// A collection of values returned by getForwardEntries.
type GetForwardEntriesResult struct {
	// A list of Forward Entries. Each element contains the following attributes:
	Entries []GetForwardEntriesEntry `pulumi:"entries"`
	// The public IP address.
	ExternalIp     *string `pulumi:"externalIp"`
	ForwardTableId string  `pulumi:"forwardTableId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Forward Entries IDs.
	Ids []string `pulumi:"ids"`
	// The private IP address.
	InternalIp *string `pulumi:"internalIp"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of Forward Entries names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
