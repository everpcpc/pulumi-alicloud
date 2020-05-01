// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Eip extends pulumi.CustomResource {
    /**
     * Get an existing Eip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipState, opts?: pulumi.CustomResourceOptions): Eip {
        return new Eip(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/eip:Eip';

    /**
     * Returns true if the given object is an instance of Eip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eip.__pulumiType;
    }

    /**
     * Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instanceChargeType` is PrePaid.
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    /**
     * The elastic ip address
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     */
    public readonly isp!: pulumi.Output<string>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The Id of resource group which the eip belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The EIP current status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Eip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipArgs | EipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EipState | undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["isp"] = state ? state.isp : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EipArgs | undefined;
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["isp"] = args ? args.isp : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ipAddress"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Eip.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Eip resources.
 */
export interface EipState {
    /**
     * Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instanceChargeType` is PrePaid.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The elastic ip address
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     */
    readonly isp?: pulumi.Input<string>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the eip belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * The EIP current status.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Eip resource.
 */
export interface EipArgs {
    /**
     * Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. From version `1.7.1`, default to `PayByTraffic`. It is only PayByBandwidth when `instanceChargeType` is PrePaid.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     */
    readonly isp?: pulumi.Input<string>;
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The Id of resource group which the eip belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
