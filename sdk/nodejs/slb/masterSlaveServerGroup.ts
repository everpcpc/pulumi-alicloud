// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccSlbMasterSlaveServerGroupVpc";
 * const number = config.get("number") || "1";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloudEfficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     eniAmount: 2,
 * });
 * const image = alicloud.ecs.getImages({
 *     mostRecent: true,
 *     nameRegex: "^ubuntu_18.*64",
 *     owners: "system",
 * });
 * const mainNetwork = new alicloud.vpc.Network("main", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const mainSwitch = new alicloud.vpc.Switch("main", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/16",
 *     vpcId: mainNetwork.id,
 * });
 * const groupSecurityGroup = new alicloud.ecs.SecurityGroup("group", {
 *     vpcId: mainNetwork.id,
 * });
 * const instanceInstance: alicloud.ecs.Instance[] = [];
 * for (let i = 0; i < 2; i++) {
 *     instanceInstance.push(new alicloud.ecs.Instance(`instance-${i}`, {
 *         availabilityZone: defaultZones.zones[0].id,
 *         imageId: image.images[0].id,
 *         instanceChargeType: "PostPaid",
 *         instanceName: name,
 *         instanceType: defaultInstanceTypes.instanceTypes[0].id,
 *         internetChargeType: "PayByTraffic",
 *         internetMaxBandwidthOut: 10,
 *         securityGroups: [groupSecurityGroup.id],
 *         systemDiskCategory: "cloudEfficiency",
 *         vswitchId: mainSwitch.id,
 *     }));
 * }
 * const instanceLoadBalancer = new alicloud.slb.LoadBalancer("instance", {
 *     specification: "slb.s2.small",
 *     vswitchId: mainSwitch.id,
 * });
 * const defaultNetworkInterface: alicloud.vpc.NetworkInterface[] = [];
 * for (let i = 0; i < number; i++) {
 *     defaultNetworkInterface.push(new alicloud.vpc.NetworkInterface(`default-${i}`, {
 *         securityGroups: [groupSecurityGroup.id],
 *         vswitchId: mainSwitch.id,
 *     }));
 * }
 * const defaultNetworkInterfaceAttachment: alicloud.vpc.NetworkInterfaceAttachment[] = [];
 * for (let i = 0; i < number; i++) {
 *     defaultNetworkInterfaceAttachment.push(new alicloud.vpc.NetworkInterfaceAttachment(`default-${i}`, {
 *         instanceId: instanceInstance[0].id,
 *         networkInterfaceId: pulumi.all(defaultNetworkInterface.map(v => v.id)).apply(id => id.map(v => v)[i]),
 *     }));
 * }
 * const groupMasterSlaveServerGroup = new alicloud.slb.MasterSlaveServerGroup("group", {
 *     loadBalancerId: instanceLoadBalancer.id,
 *     servers: [
 *         {
 *             port: 100,
 *             serverId: instanceInstance[0].id,
 *             serverType: "Master",
 *             weight: 100,
 *         },
 *         {
 *             port: 100,
 *             serverId: instanceInstance[1].id,
 *             serverType: "Slave",
 *             weight: 100,
 *         },
 *     ],
 * });
 * const tcp = new alicloud.slb.Listener("tcp", {
 *     bandwidth: 10,
 *     establishedTimeout: 600,
 *     frontendPort: 22,
 *     healthCheckConnectPort: 20,
 *     healthCheckHttpCode: "http2xx",
 *     healthCheckInterval: 5,
 *     healthCheckTimeout: 8,
 *     healthCheckType: "tcp",
 *     healthCheckUri: "/console",
 *     healthyThreshold: 8,
 *     loadBalancerId: instanceLoadBalancer.id,
 *     masterSlaveServerGroupId: groupMasterSlaveServerGroup.id,
 *     persistenceTimeout: 3600,
 *     protocol: "tcp",
 *     unhealthyThreshold: 8,
 * });
 * ```
 * 
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
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/slb_master_slave_server_group.html.markdown.
 */
export class MasterSlaveServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing MasterSlaveServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        if (opts && opts.id) {
            const state = argsOrState as MasterSlaveServerGroupState | undefined;
            inputs["deleteProtectionValidation"] = state ? state.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as MasterSlaveServerGroupArgs | undefined;
            if (!args || args.loadBalancerId === undefined) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            inputs["deleteProtectionValidation"] = args ? args.deleteProtectionValidation : undefined;
            inputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["servers"] = args ? args.servers : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
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
