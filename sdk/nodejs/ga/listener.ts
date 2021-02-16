// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Listener resource.
 *
 * For information about Global Accelerator (GA) Listener and how to use it, see [What is Listener](https://help.aliyun.com/document_detail/153253.html).
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleAccelerator = new alicloud.ga.Accelerator("exampleAccelerator", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "1",
 * });
 * const exampleListener = new alicloud.ga.Listener("exampleListener", {
 *     acceleratorId: exampleAccelerator.id,
 *     portRanges: [{
 *         fromPort: 60,
 *         toPort: 70,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Ga Listener can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/listener:Listener example <id>
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * The accelerator id.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The certificates of the listener.
     */
    public readonly certificates!: pulumi.Output<outputs.ga.ListenerCertificate[] | undefined>;
    /**
     * The clientAffinity of the listener. Default value is `NONE`. Valid values:
     * `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
     * `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
     */
    public readonly clientAffinity!: pulumi.Output<string | undefined>;
    /**
     * The description of the listener.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The portRanges of the listener.
     */
    public readonly portRanges!: pulumi.Output<outputs.ga.ListenerPortRange[]>;
    /**
     * Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The proxy protocol of the listener.
     */
    public readonly proxyProtocol!: pulumi.Output<boolean | undefined>;
    /**
     * The status of the listener.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerState | undefined;
            inputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            inputs["certificates"] = state ? state.certificates : undefined;
            inputs["clientAffinity"] = state ? state.clientAffinity : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portRanges"] = state ? state.portRanges : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["proxyProtocol"] = state ? state.proxyProtocol : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.portRanges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portRanges'");
            }
            inputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            inputs["certificates"] = args ? args.certificates : undefined;
            inputs["clientAffinity"] = args ? args.clientAffinity : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portRanges"] = args ? args.portRanges : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["proxyProtocol"] = args ? args.proxyProtocol : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Listener.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * The accelerator id.
     */
    readonly acceleratorId?: pulumi.Input<string>;
    /**
     * The certificates of the listener.
     */
    readonly certificates?: pulumi.Input<pulumi.Input<inputs.ga.ListenerCertificate>[]>;
    /**
     * The clientAffinity of the listener. Default value is `NONE`. Valid values:
     * `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
     * `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
     */
    readonly clientAffinity?: pulumi.Input<string>;
    /**
     * The description of the listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The portRanges of the listener.
     */
    readonly portRanges?: pulumi.Input<pulumi.Input<inputs.ga.ListenerPortRange>[]>;
    /**
     * Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The proxy protocol of the listener.
     */
    readonly proxyProtocol?: pulumi.Input<boolean>;
    /**
     * The status of the listener.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The accelerator id.
     */
    readonly acceleratorId: pulumi.Input<string>;
    /**
     * The certificates of the listener.
     */
    readonly certificates?: pulumi.Input<pulumi.Input<inputs.ga.ListenerCertificate>[]>;
    /**
     * The clientAffinity of the listener. Default value is `NONE`. Valid values:
     * `NONE`: client affinity is not maintained, that is, connection requests from the same client cannot always be directed to the same terminal node.
     * `SOURCE_IP`: maintain client affinity. When a client accesses a stateful application, all requests from the same client can be directed to the same terminal node, regardless of the source port and protocol.
     */
    readonly clientAffinity?: pulumi.Input<string>;
    /**
     * The description of the listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The portRanges of the listener.
     */
    readonly portRanges: pulumi.Input<pulumi.Input<inputs.ga.ListenerPortRange>[]>;
    /**
     * Type of network transport protocol monitored. Default value is `TCP`. Valid values: `TCP`, `UDP`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The proxy protocol of the listener.
     */
    readonly proxyProtocol?: pulumi.Input<boolean>;
}
