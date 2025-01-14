// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const config = new pulumi.Config();
 * const name = config.get("name") || "test_havip_attachment";
 * const fooNetwork = new alicloud.vpc.Network("fooNetwork", {cidrBlock: "172.16.0.0/12"});
 * const fooSwitch = new alicloud.vpc.Switch("fooSwitch", {
 *     vpcId: fooNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const fooHAVip = new alicloud.vpc.HAVip("fooHAVip", {
 *     vswitchId: fooSwitch.id,
 *     description: name,
 * });
 * const tfTestFoo = new alicloud.ecs.SecurityGroup("tfTestFoo", {
 *     description: "foo",
 *     vpcId: fooNetwork.id,
 * });
 * const fooInstance = new alicloud.ecs.Instance("fooInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchId: fooSwitch.id,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     systemDiskCategory: "cloud_efficiency",
 *     internetChargeType: "PayByTraffic",
 *     internetMaxBandwidthOut: 5,
 *     securityGroups: [tfTestFoo.id],
 *     instanceName: name,
 *     userData: "echo 'net.ipv4.ip_forward=1'>> /etc/sysctl.conf",
 * });
 * const fooHAVipAttachment = new alicloud.vpc.HAVipAttachment("fooHAVipAttachment", {
 *     havipId: fooHAVip.id,
 *     instanceId: fooInstance.id,
 * });
 * ```
 *
 * ## Import
 *
 * The havip attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/hAVipAttachment:HAVipAttachment foo havip-abc123456:i-abc123456
 * ```
 */
export class HAVipAttachment extends pulumi.CustomResource {
    /**
     * Get an existing HAVipAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HAVipAttachmentState, opts?: pulumi.CustomResourceOptions): HAVipAttachment {
        return new HAVipAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/hAVipAttachment:HAVipAttachment';

    /**
     * Returns true if the given object is an instance of HAVipAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HAVipAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HAVipAttachment.__pulumiType;
    }

    /**
     * Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
     */
    public readonly force!: pulumi.Output<string | undefined>;
    /**
     * The havipId of the havip attachment, the field can't be changed.
     */
    public readonly havipId!: pulumi.Output<string>;
    /**
     * The instanceId of the havip attachment, the field can't be changed.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * (Available in v1.201.0+) The status of the HaVip instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a HAVipAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HAVipAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HAVipAttachmentArgs | HAVipAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HAVipAttachmentState | undefined;
            resourceInputs["force"] = state ? state.force : undefined;
            resourceInputs["havipId"] = state ? state.havipId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as HAVipAttachmentArgs | undefined;
            if ((!args || args.havipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'havipId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["havipId"] = args ? args.havipId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HAVipAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HAVipAttachment resources.
 */
export interface HAVipAttachmentState {
    /**
     * Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
     */
    force?: pulumi.Input<string>;
    /**
     * The havipId of the havip attachment, the field can't be changed.
     */
    havipId?: pulumi.Input<string>;
    /**
     * The instanceId of the havip attachment, the field can't be changed.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * (Available in v1.201.0+) The status of the HaVip instance.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HAVipAttachment resource.
 */
export interface HAVipAttachmentArgs {
    /**
     * Specifies whether to forcefully disassociate the HAVIP from the ECS instance or ENI. Default value: `False`. Valid values: `True` and `False`.
     */
    force?: pulumi.Input<string>;
    /**
     * The havipId of the havip attachment, the field can't be changed.
     */
    havipId: pulumi.Input<string>;
    /**
     * The instanceId of the havip attachment, the field can't be changed.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The Type of instance to bind HaVip to. Valid values: `EcsInstance` and `NetworkInterface`. When the HaVip instance is bound to a resilient NIC, the resilient NIC instance must be filled in.
     */
    instanceType?: pulumi.Input<string>;
}
