// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the rules associated with a server load balancer listener.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "slbrulebasicconfig"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
// 			AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
// 			AvailableResourceCreation: pulumi.StringRef("VSwitch"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:       defaultNetwork.ID(),
// 			CidrBlock:   pulumi.String("172.16.0.0/16"),
// 			ZoneId:      pulumi.String(defaultZones.Zones[0].Id),
// 			VswitchName: pulumi.String(name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "defaultApplicationLoadBalancer", &slb.ApplicationLoadBalancerArgs{
// 			LoadBalancerName: pulumi.String(name),
// 			VswitchId:        defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultListener, err := slb.NewListener(ctx, "defaultListener", &slb.ListenerArgs{
// 			LoadBalancerId:         defaultApplicationLoadBalancer.ID(),
// 			BackendPort:            pulumi.Int(22),
// 			FrontendPort:           pulumi.Int(22),
// 			Protocol:               pulumi.String("http"),
// 			Bandwidth:              pulumi.Int(5),
// 			HealthCheckConnectPort: pulumi.Int(20),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultServerGroup, err := slb.NewServerGroup(ctx, "defaultServerGroup", &slb.ServerGroupArgs{
// 			LoadBalancerId: defaultApplicationLoadBalancer.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = slb.NewRule(ctx, "defaultRule", &slb.RuleArgs{
// 			LoadBalancerId: defaultApplicationLoadBalancer.ID(),
// 			FrontendPort:   defaultListener.FrontendPort,
// 			Domain:         pulumi.String("*.aliyun.com"),
// 			Url:            pulumi.String("/image"),
// 			ServerGroupId:  defaultServerGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstSlbRuleId", sampleDs.ApplyT(func(sampleDs slb.GetRulesResult) (string, error) {
// 			return sampleDs.SlbRules[0].Id, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetRules(ctx *pulumi.Context, args *GetRulesArgs, opts ...pulumi.InvokeOption) (*GetRulesResult, error) {
	var rv GetRulesResult
	err := ctx.Invoke("alicloud:slb/getRules:getRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRules.
type GetRulesArgs struct {
	// SLB listener port.
	FrontendPort int `pulumi:"frontendPort"`
	// A list of rules IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB with listener rules.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by rule name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRules.
type GetRulesResult struct {
	FrontendPort int `pulumi:"frontendPort"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB listener rules IDs.
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	// A list of SLB listener rules names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of SLB listener rules. Each element contains the following attributes:
	SlbRules []GetRulesSlbRule `pulumi:"slbRules"`
}

func GetRulesOutput(ctx *pulumi.Context, args GetRulesOutputArgs, opts ...pulumi.InvokeOption) GetRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRulesResult, error) {
			args := v.(GetRulesArgs)
			r, err := GetRules(ctx, &args, opts...)
			return *r, err
		}).(GetRulesResultOutput)
}

// A collection of arguments for invoking getRules.
type GetRulesOutputArgs struct {
	// SLB listener port.
	FrontendPort pulumi.IntInput `pulumi:"frontendPort"`
	// A list of rules IDs to filter results.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the SLB with listener rules.
	LoadBalancerId pulumi.StringInput `pulumi:"loadBalancerId"`
	// A regex string to filter results by rule name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesArgs)(nil)).Elem()
}

// A collection of values returned by getRules.
type GetRulesResultOutput struct{ *pulumi.OutputState }

func (GetRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesResult)(nil)).Elem()
}

func (o GetRulesResultOutput) ToGetRulesResultOutput() GetRulesResultOutput {
	return o
}

func (o GetRulesResultOutput) ToGetRulesResultOutputWithContext(ctx context.Context) GetRulesResultOutput {
	return o
}

func (o GetRulesResultOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesResult) int { return v.FrontendPort }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SLB listener rules IDs.
func (o GetRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRulesResultOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o GetRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of SLB listener rules names.
func (o GetRulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of SLB listener rules. Each element contains the following attributes:
func (o GetRulesResultOutput) SlbRules() GetRulesSlbRuleArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []GetRulesSlbRule { return v.SlbRules }).(GetRulesSlbRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRulesResultOutput{})
}
