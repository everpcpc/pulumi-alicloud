// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the VServer groups related to a server load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/slb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "slbservergroups"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := "cloud_efficiency"
// 		opt1 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableDiskCategory:     &opt0,
// 			AvailableResourceCreation: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			VpcName:   pulumi.String(name),
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/16"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 			VswitchName:      pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultLoadBalancer, err := slb.NewLoadBalancer(ctx, "defaultLoadBalancer", &slb.LoadBalancerArgs{
// 			VswitchId: defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = slb.NewServerGroup(ctx, "defaultServerGroup", &slb.ServerGroupArgs{
// 			LoadBalancerId: defaultLoadBalancer.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstSlbServerGroupId", sampleDs.ApplyT(func(sampleDs slb.GetServerGroupsResult) (string, error) {
// 			return sampleDs.SlbServerGroups[0].Id, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetServerGroups(ctx *pulumi.Context, args *GetServerGroupsArgs, opts ...pulumi.InvokeOption) (*GetServerGroupsResult, error) {
	var rv GetServerGroupsResult
	err := ctx.Invoke("alicloud:slb/getServerGroups:getServerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerGroups.
type GetServerGroupsArgs struct {
	// A list of VServer group IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by VServer group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getServerGroups.
type GetServerGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB VServer groups IDs.
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	// A list of SLB VServer groups names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of SLB VServer groups. Each element contains the following attributes:
	SlbServerGroups []GetServerGroupsSlbServerGroup `pulumi:"slbServerGroups"`
}
