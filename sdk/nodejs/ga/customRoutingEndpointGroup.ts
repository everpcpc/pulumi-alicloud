// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Custom Routing Endpoint Group resource.
 *
 * For information about Global Accelerator (GA) Custom Routing Endpoint Group and how to use it, see [What is Custom Routing Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/createcustomroutingendpointgroups).
 *
 * > **NOTE:** Available in v1.197.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAccelerators = alicloud.ga.getAccelerators({
 *     status: "active",
 * });
 * const defaultBandwidthPackage = new alicloud.ga.BandwidthPackage("defaultBandwidthPackage", {
 *     bandwidth: 100,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     paymentType: "PayAsYouGo",
 *     billingType: "PayBy95",
 *     ratio: 30,
 * });
 * const defaultBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", {
 *     acceleratorId: defaultAccelerators.then(defaultAccelerators => defaultAccelerators.accelerators?.[0]?.id),
 *     bandwidthPackageId: defaultBandwidthPackage.id,
 * });
 * const defaultListener = new alicloud.ga.Listener("defaultListener", {
 *     acceleratorId: defaultBandwidthPackageAttachment.acceleratorId,
 *     listenerType: "CustomRouting",
 *     portRanges: [{
 *         fromPort: 10000,
 *         toPort: 16000,
 *     }],
 * });
 * const defaultCustomRoutingEndpointGroup = new alicloud.ga.CustomRoutingEndpointGroup("defaultCustomRoutingEndpointGroup", {
 *     acceleratorId: defaultListener.acceleratorId,
 *     listenerId: defaultListener.id,
 *     endpointGroupRegion: "cn-hangzhou",
 *     customRoutingEndpointGroupName: "example_value",
 *     description: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Custom Routing Endpoint Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/customRoutingEndpointGroup:CustomRoutingEndpointGroup example <id>
 * ```
 */
export class CustomRoutingEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing CustomRoutingEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomRoutingEndpointGroupState, opts?: pulumi.CustomResourceOptions): CustomRoutingEndpointGroup {
        return new CustomRoutingEndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/customRoutingEndpointGroup:CustomRoutingEndpointGroup';

    /**
     * Returns true if the given object is an instance of CustomRoutingEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomRoutingEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomRoutingEndpointGroup.__pulumiType;
    }

    /**
     * The ID of the GA instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The name of the endpoint group.
     */
    public readonly customRoutingEndpointGroupName!: pulumi.Output<string | undefined>;
    /**
     * The description of the endpoint group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the region in which to create the endpoint group.
     */
    public readonly endpointGroupRegion!: pulumi.Output<string>;
    /**
     * The ID of the custom routing listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The status of the Custom Routing Endpoint Group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a CustomRoutingEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomRoutingEndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomRoutingEndpointGroupArgs | CustomRoutingEndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomRoutingEndpointGroupState | undefined;
            resourceInputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            resourceInputs["customRoutingEndpointGroupName"] = state ? state.customRoutingEndpointGroupName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointGroupRegion"] = state ? state.endpointGroupRegion : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as CustomRoutingEndpointGroupArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.endpointGroupRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointGroupRegion'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            resourceInputs["customRoutingEndpointGroupName"] = args ? args.customRoutingEndpointGroupName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointGroupRegion"] = args ? args.endpointGroupRegion : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomRoutingEndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomRoutingEndpointGroup resources.
 */
export interface CustomRoutingEndpointGroupState {
    /**
     * The ID of the GA instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The name of the endpoint group.
     */
    customRoutingEndpointGroupName?: pulumi.Input<string>;
    /**
     * The description of the endpoint group.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the region in which to create the endpoint group.
     */
    endpointGroupRegion?: pulumi.Input<string>;
    /**
     * The ID of the custom routing listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The status of the Custom Routing Endpoint Group.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomRoutingEndpointGroup resource.
 */
export interface CustomRoutingEndpointGroupArgs {
    /**
     * The ID of the GA instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The name of the endpoint group.
     */
    customRoutingEndpointGroupName?: pulumi.Input<string>;
    /**
     * The description of the endpoint group.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the region in which to create the endpoint group.
     */
    endpointGroupRegion: pulumi.Input<string>;
    /**
     * The ID of the custom routing listener.
     */
    listenerId: pulumi.Input<string>;
}