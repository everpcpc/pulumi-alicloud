// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Private Link Vpc Endpoint Service resource.
 *
 * For information about Private Link Vpc Endpoint Service and how to use it, see [What is Vpc Endpoint Service](https://help.aliyun.com/document_detail/183540.html).
 *
 * > **NOTE:** Available in v1.109.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.privatelink.VpcEndpointService("example", {
 *     autoAcceptConnection: false,
 *     connectBandwidth: 103,
 *     serviceDescription: "tftest",
 * });
 * ```
 *
 * ## Import
 *
 * Private Link Vpc Endpoint Service can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointService:VpcEndpointService example <service_id>
 * ```
 */
export class VpcEndpointService extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceState, opts?: pulumi.CustomResourceOptions): VpcEndpointService {
        return new VpcEndpointService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:privatelink/vpcEndpointService:VpcEndpointService';

    /**
     * Returns true if the given object is an instance of VpcEndpointService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointService.__pulumiType;
    }

    /**
     * Whether to automatically accept terminal node connections.
     */
    public readonly autoAcceptConnection!: pulumi.Output<boolean | undefined>;
    /**
     * The connection bandwidth.
     */
    public readonly connectBandwidth!: pulumi.Output<number | undefined>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
     */
    public readonly payer!: pulumi.Output<string | undefined>;
    /**
     * The business status of Vpc Endpoint Service.
     */
    public /*out*/ readonly serviceBusinessStatus!: pulumi.Output<string>;
    /**
     * The description of the terminal node service.
     */
    public readonly serviceDescription!: pulumi.Output<string | undefined>;
    /**
     * Service Domain.
     */
    public /*out*/ readonly serviceDomain!: pulumi.Output<string>;
    /**
     * The status of Vpc Endpoint Service.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcEndpointServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceArgs | VpcEndpointServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointServiceState | undefined;
            inputs["autoAcceptConnection"] = state ? state.autoAcceptConnection : undefined;
            inputs["connectBandwidth"] = state ? state.connectBandwidth : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["payer"] = state ? state.payer : undefined;
            inputs["serviceBusinessStatus"] = state ? state.serviceBusinessStatus : undefined;
            inputs["serviceDescription"] = state ? state.serviceDescription : undefined;
            inputs["serviceDomain"] = state ? state.serviceDomain : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceArgs | undefined;
            inputs["autoAcceptConnection"] = args ? args.autoAcceptConnection : undefined;
            inputs["connectBandwidth"] = args ? args.connectBandwidth : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["payer"] = args ? args.payer : undefined;
            inputs["serviceDescription"] = args ? args.serviceDescription : undefined;
            inputs["serviceBusinessStatus"] = undefined /*out*/;
            inputs["serviceDomain"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VpcEndpointService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointService resources.
 */
export interface VpcEndpointServiceState {
    /**
     * Whether to automatically accept terminal node connections.
     */
    readonly autoAcceptConnection?: pulumi.Input<boolean>;
    /**
     * The connection bandwidth.
     */
    readonly connectBandwidth?: pulumi.Input<number>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
     */
    readonly payer?: pulumi.Input<string>;
    /**
     * The business status of Vpc Endpoint Service.
     */
    readonly serviceBusinessStatus?: pulumi.Input<string>;
    /**
     * The description of the terminal node service.
     */
    readonly serviceDescription?: pulumi.Input<string>;
    /**
     * Service Domain.
     */
    readonly serviceDomain?: pulumi.Input<string>;
    /**
     * The status of Vpc Endpoint Service.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointService resource.
 */
export interface VpcEndpointServiceArgs {
    /**
     * Whether to automatically accept terminal node connections.
     */
    readonly autoAcceptConnection?: pulumi.Input<boolean>;
    /**
     * The connection bandwidth.
     */
    readonly connectBandwidth?: pulumi.Input<number>;
    /**
     * Whether to pre-check this request only. Default to: `false`
     */
    readonly dryRun?: pulumi.Input<boolean>;
    /**
     * The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
     */
    readonly payer?: pulumi.Input<string>;
    /**
     * The description of the terminal node service.
     */
    readonly serviceDescription?: pulumi.Input<string>;
}
