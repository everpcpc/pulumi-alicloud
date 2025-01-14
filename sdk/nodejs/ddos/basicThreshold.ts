// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Ddos Basic Threshold resource.
 *
 * For information about Ddos Basic Threshold and how to use it, see [What is Threshold](https://www.alibabacloud.com/help/en/ddos-protection/latest/describe-ip-ddosthreshold).
 *
 * > **NOTE:** Available in v1.183.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "Instance",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "default-NODELETING",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?.[0],
 *     zoneId: defaultZones.zones?.[0]?.id,
 * }));
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {
 *     description: "New security group",
 *     vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?.[0]),
 * });
 * const defaultImages = alicloud.ecs.getImages({
 *     owners: "system",
 *     nameRegex: "^centos_8",
 *     mostRecent: true,
 * });
 * const defaultInstance = new alicloud.ecs.Instance("defaultInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     instanceName: _var.name,
 *     hostName: "tf-testAcc",
 *     internetMaxBandwidthOut: 10,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroups: [defaultSecurityGroup.id],
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 * });
 * const example = new alicloud.ddos.BasicThreshold("example", {
 *     pps: 60000,
 *     bps: 100,
 *     internetIp: defaultInstance.publicIp,
 *     instanceId: defaultInstance.id,
 *     instanceType: "ecs",
 * });
 * ```
 *
 * ## Import
 *
 * Ddos Basic Threshold can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ddos/basicThreshold:BasicThreshold example <instance_type>:<instance_id>:<internet_ip>
 * ```
 */
export class BasicThreshold extends pulumi.CustomResource {
    /**
     * Get an existing BasicThreshold resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BasicThresholdState, opts?: pulumi.CustomResourceOptions): BasicThreshold {
        return new BasicThreshold(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ddos/basicThreshold:BasicThreshold';

    /**
     * Returns true if the given object is an instance of BasicThreshold.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BasicThreshold {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BasicThreshold.__pulumiType;
    }

    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    public readonly bps!: pulumi.Output<number>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The IP address of the public IP address asset.
     */
    public readonly internetIp!: pulumi.Output<string>;
    /**
     * Maximum flow cleaning threshold. Unit: Mbps.
     */
    public /*out*/ readonly maxBps!: pulumi.Output<number>;
    /**
     * The maximum number of messages cleaning threshold. Unit: pps.
     */
    public /*out*/ readonly maxPps!: pulumi.Output<number>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    public readonly pps!: pulumi.Output<number>;

    /**
     * Create a BasicThreshold resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BasicThresholdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BasicThresholdArgs | BasicThresholdState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BasicThresholdState | undefined;
            resourceInputs["bps"] = state ? state.bps : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetIp"] = state ? state.internetIp : undefined;
            resourceInputs["maxBps"] = state ? state.maxBps : undefined;
            resourceInputs["maxPps"] = state ? state.maxPps : undefined;
            resourceInputs["pps"] = state ? state.pps : undefined;
        } else {
            const args = argsOrState as BasicThresholdArgs | undefined;
            if ((!args || args.bps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bps'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.internetIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internetIp'");
            }
            if ((!args || args.pps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pps'");
            }
            resourceInputs["bps"] = args ? args.bps : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetIp"] = args ? args.internetIp : undefined;
            resourceInputs["pps"] = args ? args.pps : undefined;
            resourceInputs["maxBps"] = undefined /*out*/;
            resourceInputs["maxPps"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BasicThreshold.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BasicThreshold resources.
 */
export interface BasicThresholdState {
    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    bps?: pulumi.Input<number>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The IP address of the public IP address asset.
     */
    internetIp?: pulumi.Input<string>;
    /**
     * Maximum flow cleaning threshold. Unit: Mbps.
     */
    maxBps?: pulumi.Input<number>;
    /**
     * The maximum number of messages cleaning threshold. Unit: pps.
     */
    maxPps?: pulumi.Input<number>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    pps?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BasicThreshold resource.
 */
export interface BasicThresholdArgs {
    /**
     * Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
     */
    bps: pulumi.Input<number>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The IP address of the public IP address asset.
     */
    internetIp: pulumi.Input<string>;
    /**
     * The current message number cleaning threshold. Unit: pps.
     */
    pps: pulumi.Input<number>;
}
