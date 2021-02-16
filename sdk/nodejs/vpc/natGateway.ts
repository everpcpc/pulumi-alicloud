// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Nat gateway can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/natGateway:NatGateway example ngw-abc123456
 * ```
 */
export class NatGateway extends pulumi.CustomResource {
    /**
     * Get an existing NatGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatGatewayState, opts?: pulumi.CustomResourceOptions): NatGateway {
        return new NatGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/natGateway:NatGateway';

    /**
     * Returns true if the given object is an instance of NatGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatGateway.__pulumiType;
    }

    /**
     * A list ID of the bandwidth packages, and split them with commas.
     */
    public /*out*/ readonly bandwidthPackageIds!: pulumi.Output<string>;
    /**
     * A list of bandwidth packages for the nat gatway. Only support nat gateway created before 00:00 on November 4, 2017. Available in v1.13.0+ and v1.7.1-.
     */
    public readonly bandwidthPackages!: pulumi.Output<outputs.vpc.NatGatewayBandwidthPackage[] | undefined>;
    /**
     * Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The nat gateway will auto create a snap and forward item, the `forwardTableIds` is the created one.
     */
    public /*out*/ readonly forwardTableIds!: pulumi.Output<string>;
    /**
     * The billing method of the nat gateway. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of nat gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
     */
    public readonly natType!: pulumi.Output<string | undefined>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The nat gateway will auto create a snap and forward item, the `snatTableIds` is the created one.
     */
    public /*out*/ readonly snatTableIds!: pulumi.Output<string>;
    /**
     * It has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     *
     * @deprecated Field 'spec' has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     */
    public readonly spec!: pulumi.Output<string | undefined>;
    /**
     * The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
     */
    public readonly specification!: pulumi.Output<string | undefined>;
    /**
     * The VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The id of VSwitch.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a NatGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatGatewayArgs | NatGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NatGatewayState | undefined;
            inputs["bandwidthPackageIds"] = state ? state.bandwidthPackageIds : undefined;
            inputs["bandwidthPackages"] = state ? state.bandwidthPackages : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["forwardTableIds"] = state ? state.forwardTableIds : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natType"] = state ? state.natType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["snatTableIds"] = state ? state.snatTableIds : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["specification"] = state ? state.specification : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as NatGatewayArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["bandwidthPackages"] = args ? args.bandwidthPackages : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natType"] = args ? args.natType : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["specification"] = args ? args.specification : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["bandwidthPackageIds"] = undefined /*out*/;
            inputs["forwardTableIds"] = undefined /*out*/;
            inputs["snatTableIds"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NatGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatGateway resources.
 */
export interface NatGatewayState {
    /**
     * A list ID of the bandwidth packages, and split them with commas.
     */
    readonly bandwidthPackageIds?: pulumi.Input<string>;
    /**
     * A list of bandwidth packages for the nat gatway. Only support nat gateway created before 00:00 on November 4, 2017. Available in v1.13.0+ and v1.7.1-.
     */
    readonly bandwidthPackages?: pulumi.Input<pulumi.Input<inputs.vpc.NatGatewayBandwidthPackage>[]>;
    /**
     * Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The nat gateway will auto create a snap and forward item, the `forwardTableIds` is the created one.
     */
    readonly forwardTableIds?: pulumi.Input<string>;
    /**
     * The billing method of the nat gateway. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of nat gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
     */
    readonly natType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The nat gateway will auto create a snap and forward item, the `snatTableIds` is the created one.
     */
    readonly snatTableIds?: pulumi.Input<string>;
    /**
     * It has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     *
     * @deprecated Field 'spec' has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     */
    readonly spec?: pulumi.Input<string>;
    /**
     * The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * The VPC ID.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The id of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NatGateway resource.
 */
export interface NatGatewayArgs {
    /**
     * A list of bandwidth packages for the nat gatway. Only support nat gateway created before 00:00 on November 4, 2017. Available in v1.13.0+ and v1.7.1-.
     */
    readonly bandwidthPackages?: pulumi.Input<pulumi.Input<inputs.vpc.NatGatewayBandwidthPackage>[]>;
    /**
     * Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The billing method of the nat gateway. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of nat gateway. Default to `Normal`. Valid values: [`Normal`, `Enhanced`].
     */
    readonly natType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * It has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     *
     * @deprecated Field 'spec' has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.
     */
    readonly spec?: pulumi.Input<string>;
    /**
     * The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Default to `Small`. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * The VPC ID.
     */
    readonly vpcId: pulumi.Input<string>;
    /**
     * The id of VSwitch.
     */
    readonly vswitchId?: pulumi.Input<string>;
}
