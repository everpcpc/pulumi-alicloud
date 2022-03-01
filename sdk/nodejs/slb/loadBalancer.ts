// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraformtestslbconfig";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/12"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultLoadBalancer = new alicloud.slb.LoadBalancer("defaultLoadBalancer", {
 *     specification: "slb.s2.small",
 *     vswitchId: defaultSwitch.id,
 *     tags: {
 *         tag_a: 1,
 *         tag_b: 2,
 *         tag_c: 3,
 *         tag_d: 4,
 *         tag_e: 5,
 *         tag_f: 6,
 *         tag_g: 7,
 *         tag_h: 8,
 *         tag_i: 9,
 *         tag_j: 10,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Load balancer can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:slb/loadBalancer:LoadBalancer example lb-abc123456
 * ```
 *
 * @deprecated This resource has been deprecated in favour of the ApplicationLoadBalancer resource
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        pulumi.log.warn("LoadBalancer is deprecated: This resource has been deprecated in favour of the ApplicationLoadBalancer resource")
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    public readonly addressIpVersion!: pulumi.Output<string | undefined>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    public readonly deleteProtection!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    public readonly loadBalancerName!: pulumi.Output<string>;
    public readonly loadBalancerSpec!: pulumi.Output<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    public readonly masterZoneId!: pulumi.Output<string>;
    public readonly modificationProtectionReason!: pulumi.Output<string | undefined>;
    public readonly modificationProtectionStatus!: pulumi.Output<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    public readonly name!: pulumi.Output<string>;
    public readonly paymentType!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    public readonly slaveZoneId!: pulumi.Output<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     *
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    public readonly specification!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated This resource has been deprecated in favour of the ApplicationLoadBalancer resource */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated This resource has been deprecated in favour of the ApplicationLoadBalancer resource */
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadBalancer is deprecated: This resource has been deprecated in favour of the ApplicationLoadBalancer resource")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["addressType"] = state ? state.addressType : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["deleteProtection"] = state ? state.deleteProtection : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["loadBalancerSpec"] = state ? state.loadBalancerSpec : undefined;
            resourceInputs["masterZoneId"] = state ? state.masterZoneId : undefined;
            resourceInputs["modificationProtectionReason"] = state ? state.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = state ? state.modificationProtectionStatus : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["slaveZoneId"] = state ? state.slaveZoneId : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["addressType"] = args ? args.addressType : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["deleteProtection"] = args ? args.deleteProtection : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["loadBalancerSpec"] = args ? args.loadBalancerSpec : undefined;
            resourceInputs["masterZoneId"] = args ? args.masterZoneId : undefined;
            resourceInputs["modificationProtectionReason"] = args ? args.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = args ? args.modificationProtectionStatus : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["slaveZoneId"] = args ? args.slaveZoneId : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    addressType?: pulumi.Input<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    deleteProtection?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    internetChargeType?: pulumi.Input<string>;
    loadBalancerName?: pulumi.Input<string>;
    loadBalancerSpec?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    masterZoneId?: pulumi.Input<string>;
    modificationProtectionReason?: pulumi.Input<string>;
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    name?: pulumi.Input<string>;
    paymentType?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     *
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    specification?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.
     */
    address?: pulumi.Input<string>;
    /**
     * The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet".
     * - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet.
     * - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.
     */
    addressType?: pulumi.Input<string>;
    /**
     * Valid
     * value is between 1 and 1000, If argument "internetChargeType" is "paybytraffic", then this value will be ignore.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.
     */
    deleteProtection?: pulumi.Input<string>;
    /**
     * The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     *
     * @deprecated Field 'instance_charge_type' has been deprecated from provider version 1.124. Use 'payment_type' replaces it.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Valid
     * values are `PayByBandwidth`, `PayByTraffic`. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic".
     * Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".
     */
    internetChargeType?: pulumi.Input<string>;
    loadBalancerName?: pulumi.Input<string>;
    loadBalancerSpec?: pulumi.Input<string>;
    /**
     * The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    masterZoneId?: pulumi.Input<string>;
    modificationProtectionReason?: pulumi.Input<string>;
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * @deprecated Field 'name' has been deprecated from provider version 1.123.1. New field 'load_balancer_name' instead
     */
    name?: pulumi.Input<string>;
    paymentType?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the SLB belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.
     * Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium",
     * "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".
     *
     * @deprecated Field 'specification' has been deprecated from provider version 1.123.1. New field 'load_balancer_spec' instead
     */
    specification?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. The `tags` can have a maximum of 10 tag for every load balancer instance.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VSwitch ID to launch in. If `addressType` is internet, it will be ignore.
     */
    vswitchId?: pulumi.Input<string>;
}
