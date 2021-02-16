// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud ECS Elastic Network Interface Attachment as a resource to attach ENI to or detach ENI from ECS Instances.
 *
 * For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
 *
 * ## Example Usage
 *
 * Bacis Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "networkInterfaceAttachment";
 * const number = config.get("number") || "2";
 * const vpc = new alicloud.vpc.Network("vpc", {cidrBlock: "192.168.0.0/24"});
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const vswitch = new alicloud.vpc.Switch("vswitch", {
 *     cidrBlock: "192.168.0.0/24",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *     vpcId: vpc.id,
 * });
 * const group = new alicloud.ecs.SecurityGroup("group", {vpcId: vpc.id});
 * const instanceType = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     eniAmount: 2,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const instance: alicloud.ecs.Instance[];
 * for (const range = {value: 0}; range.value < number; range.value++) {
 *     instance.push(new alicloud.ecs.Instance(`instance-${range.value}`, {
 *         availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 *         securityGroups: [group.id],
 *         instanceType: instanceType.then(instanceType => instanceType.instanceTypes[0].id),
 *         systemDiskCategory: "cloud_efficiency",
 *         imageId: defaultImages.then(defaultImages => defaultImages.images[0].id),
 *         instanceName: name,
 *         vswitchId: vswitch.id,
 *         internetMaxBandwidthOut: 10,
 *     }));
 * }
 * const _interface: alicloud.vpc.NetworkInterface[];
 * for (const range = {value: 0}; range.value < number; range.value++) {
 *     _interface.push(new alicloud.vpc.NetworkInterface(`interface-${range.value}`, {
 *         vswitchId: vswitch.id,
 *         securityGroups: [group.id],
 *     }));
 * }
 * const attachment: alicloud.vpc.NetworkInterfaceAttachment[];
 * for (const range = {value: 0}; range.value < number; range.value++) {
 *     attachment.push(new alicloud.vpc.NetworkInterfaceAttachment(`attachment-${range.value}`, {
 *         instanceId: instance.map(__item => __item.id)[range.index],
 *         networkInterfaceId: _interface.map(__item => __item.id)[range.index],
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * Network Interfaces Attachment resource can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment eni eni-abc123456789000:i-abc123456789000
 * ```
 */
export class NetworkInterfaceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceAttachmentState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceAttachment {
        return new NetworkInterfaceAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceAttachment.__pulumiType;
    }

    /**
     * The instance ID to attach.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ENI ID to attach.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceAttachmentArgs | NetworkInterfaceAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceAttachmentState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceAttachmentArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.networkInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkInterfaceAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceAttachment resources.
 */
export interface NetworkInterfaceAttachmentState {
    /**
     * The instance ID to attach.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The ENI ID to attach.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceAttachment resource.
 */
export interface NetworkInterfaceAttachmentArgs {
    /**
     * The instance ID to attach.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The ENI ID to attach.
     */
    readonly networkInterfaceId: pulumi.Input<string>;
}
