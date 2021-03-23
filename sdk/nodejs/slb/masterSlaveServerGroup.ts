// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.
 *
 * > **NOTE:** One ECS instance can be added into multiple master slave server groups.
 *
 * > **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.
 *
 * > **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.
 *
 * > **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.
 *
 * > **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.
 *
 * > **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.
 *
 * > **NOTE:** Available in 1.54.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     eniAmount: 2,
 * }));
 * const image = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccSlbMasterSlaveServerGroupVpc";
 * const number = config.get("number") || "1";
 * const mainNetwork = new alicloud.vpc.Network("mainNetwork", {cidrBlock: "172.16.0.0/16"});
 * const mainSwitch = new alicloud.vpc.Switch("mainSwitch", {
 *     vpcId: mainNetwork.id,
 *     cidrBlock: "172.16.0.0/16",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *     vswitchName: name,
 * });
 * const groupSecurityGroup = new alicloud.ecs.SecurityGroup("groupSecurityGroup", {vpcId: mainNetwork.id});
 * const instanceInstance: alicloud.ecs.Instance[];
 * for (const range = {value: 0}; range.value < "2"; range.value++) {
 *     instanceInstance.push(new alicloud.ecs.Instance(`instanceInstance-${range.value}`, {
 *         imageId: image.then(image => image.images[0].id),
 *         instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes[0].id),
 *         instanceName: name,
 *         securityGroups: [groupSecurityGroup.id],
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: "10",
 *         availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *         instanceChargeType: "PostPaid",
 *         systemDiskCategory: "cloud_efficiency",
 *         vswitchId: mainSwitch.id,
 *     }));
 * }
 * const instanceLoadBalancer = new alicloud.slb.LoadBalancer("instanceLoadBalancer", {
 *     vswitchId: mainSwitch.id,
 *     specification: "slb.s2.small",
 * });
 * const defaultNetworkInterface: alicloud.vpc.NetworkInterface[];
 * for (const range = {value: 0}; range.value < number; range.value++) {
 *     defaultNetworkInterface.push(new alicloud.vpc.NetworkInterface(`defaultNetworkInterface-${range.value}`, {
 *         vswitchId: mainSwitch.id,
 *         securityGroups: [groupSecurityGroup.id],
 *     }));
 * }
 * const defaultNetworkInterfaceAttachment: alicloud.vpc.NetworkInterfaceAttachment[];
 * for (const range = {value: 0}; range.value < number; range.value++) {
 *     defaultNetworkInterfaceAttachment.push(new alicloud.vpc.NetworkInterfaceAttachment(`defaultNetworkInterfaceAttachment-${range.value}`, {
 *         instanceId: instanceInstance[0].id,
 *         networkInterfaceId: defaultNetworkInterface.map(__item => __item.id)[range.index],
 *     }));
 * }
 * const groupMasterSlaveServerGroup = new alicloud.slb.MasterSlaveServerGroup("groupMasterSlaveServerGroup", {
 *     loadBalancerId: instanceLoadBalancer.id,
 *     servers: [
 *         {
 *             serverId: instanceInstance[0].id,
 *             port: 100,
 *             weight: 100,
 *             serverType: "Master",
 *         },
 *         {
 *             serverId: instanceInstance[1].id,
 *             port: 100,
 *             weight: 100,
 *             serverType: "Slave",
 *         },
 *     ],
 * });
 * const tcp = new alicloud.slb.Listener("tcp", {
 *     loadBalancerId: instanceLoadBalancer.id,
 *     masterSlaveServerGroupId: groupMasterSlaveServerGroup.id,
 *     frontendPort: "22",
 *     protocol: "tcp",
 *     bandwidth: "10",
 *     healthCheckType: "tcp",
 *     persistenceTimeout: 3600,
 *     healthyThreshold: 8,
 *     unhealthyThreshold: 8,
 *     healthCheckTimeout: 8,
 *     healthCheckInterval: 5,
 *     healthCheckHttpCode: "http_2xx",
 *     healthCheckConnectPort: 20,
 *     healthCheckUri: "/console",
 *     establishedTimeout: 600,
 * });
 * ```
 * ## Block servers
 *
 * The servers mapping supports the following:
 *
 * * `serverIds` - (Required) A list backend server ID (ECS instance ID).
 * * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
 * * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
 * * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
 * * `serverType` - (Optional) The server type of the backend server. Valid value Master, Slave.
 * * `isBackup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.
 *
 * ## Import
 *
 * Load balancer master slave server group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
 * ```
 */
export class MasterSlaveServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing MasterSlaveServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MasterSlaveServerGroupState, opts?: pulumi.CustomResourceOptions): MasterSlaveServerGroup {
        return new MasterSlaveServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup';

    /**
     * Returns true if the given object is an instance of MasterSlaveServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MasterSlaveServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MasterSlaveServerGroup.__pulumiType;
    }

    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    public readonly deleteProtectionValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The Load Balancer ID which is used to launch a new master slave server group.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * Name of the master slave server group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
     */
    public readonly servers!: pulumi.Output<outputs.slb.MasterSlaveServerGroupServer[] | undefined>;

    /**
     * Create a MasterSlaveServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MasterSlaveServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MasterSlaveServerGroupArgs | MasterSlaveServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MasterSlaveServerGroupState | undefined;
            inputs["deleteProtectionValidation"] = state ? state.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as MasterSlaveServerGroupArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            inputs["deleteProtectionValidation"] = args ? args.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["servers"] = args ? args.servers : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MasterSlaveServerGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MasterSlaveServerGroup resources.
 */
export interface MasterSlaveServerGroupState {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new master slave server group.
     */
    readonly loadBalancerId?: pulumi.Input<string>;
    /**
     * Name of the master slave server group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
     */
    readonly servers?: pulumi.Input<pulumi.Input<inputs.slb.MasterSlaveServerGroupServer>[]>;
}

/**
 * The set of arguments for constructing a MasterSlaveServerGroup resource.
 */
export interface MasterSlaveServerGroupArgs {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The Load Balancer ID which is used to launch a new master slave server group.
     */
    readonly loadBalancerId: pulumi.Input<string>;
    /**
     * Name of the master slave server group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
     */
    readonly servers?: pulumi.Input<pulumi.Input<inputs.slb.MasterSlaveServerGroupServer>[]>;
}
