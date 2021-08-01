// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Event Bridge Event Bus resource.
 *
 * For information about Event Bridge Event Bus and how to use it, see [What is Event Bus](https://help.aliyun.com/document_detail/167863.html).
 *
 * > **NOTE:** Available in v1.129.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.eventbridge.EventBus("example", {
 *     eventBusName: "my-EventBus",
 * });
 * ```
 *
 * ## Import
 *
 * Event Bridge Event Bus can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eventbridge/eventBus:EventBus example <event_bus_name>
 * ```
 */
export class EventBus extends pulumi.CustomResource {
    /**
     * Get an existing EventBus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventBusState, opts?: pulumi.CustomResourceOptions): EventBus {
        return new EventBus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eventbridge/eventBus:EventBus';

    /**
     * Returns true if the given object is an instance of EventBus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventBus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventBus.__pulumiType;
    }

    /**
     * The description of event bus.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of event bus. The length is limited to 2 ~ 127 characters, which can be composed of letters, numbers or hyphens (-)
     */
    public readonly eventBusName!: pulumi.Output<string>;

    /**
     * Create a EventBus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventBusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventBusArgs | EventBusState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventBusState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventBusName"] = state ? state.eventBusName : undefined;
        } else {
            const args = argsOrState as EventBusArgs | undefined;
            if ((!args || args.eventBusName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventBusName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["eventBusName"] = args ? args.eventBusName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EventBus.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventBus resources.
 */
export interface EventBusState {
    /**
     * The description of event bus.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of event bus. The length is limited to 2 ~ 127 characters, which can be composed of letters, numbers or hyphens (-)
     */
    readonly eventBusName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventBus resource.
 */
export interface EventBusArgs {
    /**
     * The description of event bus.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of event bus. The length is limited to 2 ~ 127 characters, which can be composed of letters, numbers or hyphens (-)
     */
    readonly eventBusName: pulumi.Input<string>;
}
