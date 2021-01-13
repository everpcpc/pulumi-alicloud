// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Connection resource.
 *
 * For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://help.aliyun.com/document_detail/183551.html).
 *
 * > **NOTE:** Available in v1.110.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.privatelink.VpcEndpointServiceConnection("example", {
 *     bandwidth: 1024,
 *     endpointId: "example_value",
 *     serviceId: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Connection can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example <service_id>:<endpoint_id>
 * ```
 */
export class VpcEndpointServiceConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointServiceConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceConnectionState, opts?: pulumi.CustomResourceOptions): VpcEndpointServiceConnection {
        return new VpcEndpointServiceConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection';

    /**
     * Returns true if the given object is an instance of VpcEndpointServiceConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointServiceConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointServiceConnection.__pulumiType;
    }

    /**
     * The Bandwidth.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Vpc Endpoint.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * The ID of the Vpc Endpoint Service.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The status of Vpc Endpoint Connection.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointServiceConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointServiceConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceConnectionArgs | VpcEndpointServiceConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcEndpointServiceConnectionState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["endpointId"] = state ? state.endpointId : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceConnectionArgs | undefined;
            if ((!args || args.endpointId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.serviceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serviceId'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["endpointId"] = args ? args.endpointId : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcEndpointServiceConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointServiceConnection resources.
 */
export interface VpcEndpointServiceConnectionState {
    /**
     * The Bandwidth.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * The dry run.
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Vpc Endpoint.
     */
    readonly endpointId?: pulumi.Input<string>;
    /**
     * The ID of the Vpc Endpoint Service.
     */
    readonly serviceId?: pulumi.Input<string>;
    /**
     * The status of Vpc Endpoint Connection.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointServiceConnection resource.
 */
export interface VpcEndpointServiceConnectionArgs {
    /**
     * The Bandwidth.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * The dry run.
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Vpc Endpoint.
     */
    readonly endpointId: pulumi.Input<string>;
    /**
     * The ID of the Vpc Endpoint Service.
     */
    readonly serviceId: pulumi.Input<string>;
}
