// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * BGP-Line Anti-DDoS instance resource. "Ddoscoo" is the short term of this product. See [What is Anti-DDoS Pro](https://www.alibabacloud.com/help/doc-detail/69319.htm).
 *
 * > **NOTE:** The product region only support cn-hangzhou.
 *
 * > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
 *
 * > **NOTE:** Available in 1.37.0+ .
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const newInstance = new alicloud.ddos.DdosCooInstance("newInstance", {
 *     bandwidth: "30",
 *     baseBandwidth: "30",
 *     domainCount: "50",
 *     period: 1,
 *     portCount: "50",
 *     serviceBandwidth: "100",
 * });
 * ```
 *
 * ## Import
 *
 * Ddoscoo instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ddos/ddosCooInstance:DdosCooInstance example ddoscoo-cn-123456
 * ```
 */
export class DdosCooInstance extends pulumi.CustomResource {
    /**
     * Get an existing DdosCooInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DdosCooInstanceState, opts?: pulumi.CustomResourceOptions): DdosCooInstance {
        return new DdosCooInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ddos/ddosCooInstance:DdosCooInstance';

    /**
     * Returns true if the given object is an instance of DdosCooInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DdosCooInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DdosCooInstance.__pulumiType;
    }

    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    public readonly baseBandwidth!: pulumi.Output<string>;
    /**
     * Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    public readonly domainCount!: pulumi.Output<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    public readonly portCount!: pulumi.Output<string>;
    /**
     * Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
     */
    public readonly serviceBandwidth!: pulumi.Output<string>;

    /**
     * Create a DdosCooInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DdosCooInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DdosCooInstanceArgs | DdosCooInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DdosCooInstanceState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["baseBandwidth"] = state ? state.baseBandwidth : undefined;
            inputs["domainCount"] = state ? state.domainCount : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["portCount"] = state ? state.portCount : undefined;
            inputs["serviceBandwidth"] = state ? state.serviceBandwidth : undefined;
        } else {
            const args = argsOrState as DdosCooInstanceArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.baseBandwidth === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'baseBandwidth'");
            }
            if ((!args || args.domainCount === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'domainCount'");
            }
            if ((!args || args.portCount === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'portCount'");
            }
            if ((!args || args.serviceBandwidth === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serviceBandwidth'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["baseBandwidth"] = args ? args.baseBandwidth : undefined;
            inputs["domainCount"] = args ? args.domainCount : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["portCount"] = args ? args.portCount : undefined;
            inputs["serviceBandwidth"] = args ? args.serviceBandwidth : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "alicloud:dns/ddosCooInstance:DdosCooInstance" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DdosCooInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DdosCooInstance resources.
 */
export interface DdosCooInstanceState {
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    readonly bandwidth?: pulumi.Input<string>;
    /**
     * Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    readonly baseBandwidth?: pulumi.Input<string>;
    /**
     * Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    readonly domainCount?: pulumi.Input<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    readonly portCount?: pulumi.Input<string>;
    /**
     * Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
     */
    readonly serviceBandwidth?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DdosCooInstance resource.
 */
export interface DdosCooInstanceArgs {
    /**
     * Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    readonly bandwidth: pulumi.Input<string>;
    /**
     * Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
     */
    readonly baseBandwidth: pulumi.Input<string>;
    /**
     * Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    readonly domainCount: pulumi.Input<string>;
    /**
     * Name of the instance. This name can have a string of 1 to 63 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
     */
    readonly period?: pulumi.Input<number>;
    /**
     * Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
     */
    readonly portCount: pulumi.Input<string>;
    /**
     * Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
     */
    readonly serviceBandwidth: pulumi.Input<string>;
}
