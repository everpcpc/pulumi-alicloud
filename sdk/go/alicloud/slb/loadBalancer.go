// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Application Load Balancer resource.
// 
// > **NOTE:** At present, to avoid some unnecessary regulation confusion, SLB can not support alicloud international account to create "paybybandwidth" instance.
// 
// > **NOTE:** The supported specifications vary by region. Currently not all regions support guaranteed-performance instances.
// For more details about guaranteed-performance instance, see [Guaranteed-performance instances](https://www.alibabacloud.com/help/doc-detail/27657.htm).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb.html.markdown.
type LoadBalancer struct {
	s *pulumi.ResourceState
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["address"] = nil
		inputs["addressIpVersion"] = nil
		inputs["addressType"] = nil
		inputs["bandwidth"] = nil
		inputs["deleteProtection"] = nil
		inputs["instanceChargeType"] = nil
		inputs["internet"] = nil
		inputs["internetChargeType"] = nil
		inputs["masterZoneId"] = nil
		inputs["name"] = nil
		inputs["period"] = nil
		inputs["resourceGroupId"] = nil
		inputs["slaveZoneId"] = nil
		inputs["specification"] = nil
		inputs["tags"] = nil
		inputs["vswitchId"] = nil
	} else {
		inputs["address"] = args.Address
		inputs["addressIpVersion"] = args.AddressIpVersion
		inputs["addressType"] = args.AddressType
		inputs["bandwidth"] = args.Bandwidth
		inputs["deleteProtection"] = args.DeleteProtection
		inputs["instanceChargeType"] = args.InstanceChargeType
		inputs["internet"] = args.Internet
		inputs["internetChargeType"] = args.InternetChargeType
		inputs["masterZoneId"] = args.MasterZoneId
		inputs["name"] = args.Name
		inputs["period"] = args.Period
		inputs["resourceGroupId"] = args.ResourceGroupId
		inputs["slaveZoneId"] = args.SlaveZoneId
		inputs["specification"] = args.Specification
		inputs["tags"] = args.Tags
		inputs["vswitchId"] = args.VswitchId
	}
	s, err := ctx.RegisterResource("alicloud:slb/loadBalancer:LoadBalancer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerState, opts ...pulumi.ResourceOpt) (*LoadBalancer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["address"] = state.Address
		inputs["addressIpVersion"] = state.AddressIpVersion
		inputs["addressType"] = state.AddressType
		inputs["bandwidth"] = state.Bandwidth
		inputs["deleteProtection"] = state.DeleteProtection
		inputs["instanceChargeType"] = state.InstanceChargeType
		inputs["internet"] = state.Internet
		inputs["internetChargeType"] = state.InternetChargeType
		inputs["masterZoneId"] = state.MasterZoneId
		inputs["name"] = state.Name
		inputs["period"] = state.Period
		inputs["resourceGroupId"] = state.ResourceGroupId
		inputs["slaveZoneId"] = state.SlaveZoneId
		inputs["specification"] = state.Specification
		inputs["tags"] = state.Tags
		inputs["vswitchId"] = state.VswitchId
	}
	s, err := ctx.ReadResource("alicloud:slb/loadBalancer:LoadBalancer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LoadBalancer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LoadBalancer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
func (r *LoadBalancer) Address() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["address"])
}

// The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
func (r *LoadBalancer) AddressIpVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["addressIpVersion"])
}

// The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
// - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
// - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
func (r *LoadBalancer) AddressType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["addressType"])
}

// Valid
// value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
func (r *LoadBalancer) Bandwidth() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["bandwidth"])
}

// Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.   
func (r *LoadBalancer) DeleteProtection() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["deleteProtection"])
}

// The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
func (r *LoadBalancer) InstanceChargeType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceChargeType"])
}

// Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
func (r *LoadBalancer) Internet() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["internet"])
}

// Valid
// values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
// Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
func (r *LoadBalancer) InternetChargeType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internetChargeType"])
}

// The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
func (r *LoadBalancer) MasterZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["masterZoneId"])
}

func (r *LoadBalancer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
func (r *LoadBalancer) Period() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["period"])
}

// The Id of resource group which the SLB belongs.
func (r *LoadBalancer) ResourceGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupId"])
}

// The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
func (r *LoadBalancer) SlaveZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["slaveZoneId"])
}

// The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
// Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
// "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
func (r *LoadBalancer) Specification() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["specification"])
}

// A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
func (r *LoadBalancer) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
func (r *LoadBalancer) VswitchId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vswitchId"])
}

// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
	Address interface{}
	// The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
	AddressIpVersion interface{}
	// The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
	// - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
	// - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
	AddressType interface{}
	// Valid
	// value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
	Bandwidth interface{}
	// Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.   
	DeleteProtection interface{}
	// The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	InstanceChargeType interface{}
	// Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
	Internet interface{}
	// Valid
	// values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
	// Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
	InternetChargeType interface{}
	// The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
	MasterZoneId interface{}
	Name interface{}
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
	Period interface{}
	// The Id of resource group which the SLB belongs.
	ResourceGroupId interface{}
	// The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
	SlaveZoneId interface{}
	// The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
	// Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
	// "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
	Specification interface{}
	// A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
	Tags interface{}
	// The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
	VswitchId interface{}
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
	Address interface{}
	// The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
	AddressIpVersion interface{}
	// The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
	// - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
	// - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
	AddressType interface{}
	// Valid
	// value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
	Bandwidth interface{}
	// Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.   
	DeleteProtection interface{}
	// The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	InstanceChargeType interface{}
	// Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.
	Internet interface{}
	// Valid
	// values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
	// Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
	InternetChargeType interface{}
	// The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
	MasterZoneId interface{}
	Name interface{}
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36].
	Period interface{}
	// The Id of resource group which the SLB belongs.
	ResourceGroupId interface{}
	// The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
	SlaveZoneId interface{}
	// The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
	// Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
	// "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
	Specification interface{}
	// A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
	Tags interface{}
	// The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
	VswitchId interface{}
}
