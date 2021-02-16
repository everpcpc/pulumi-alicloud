// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Eipanycast Anycast Eip Address resource.
 *
 * For information about Eipanycast Anycast Eip Address and how to use it, see [What is Anycast Eip Address](https://help.aliyun.com/document_detail/169284.html).
 *
 * > **NOTE:** Available in v1.113.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.eipanycast.AnycastEipAddress("example", {
 *     serviceLocation: "international",
 * });
 * ```
 *
 * ## Import
 *
 * Eipanycast Anycast Eip Address can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eipanycast/anycastEipAddress:AnycastEipAddress example <id>
 * ```
 */
export class AnycastEipAddress extends pulumi.CustomResource {
    /**
     * Get an existing AnycastEipAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnycastEipAddressState, opts?: pulumi.CustomResourceOptions): AnycastEipAddress {
        return new AnycastEipAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eipanycast/anycastEipAddress:AnycastEipAddress';

    /**
     * Returns true if the given object is an instance of AnycastEipAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnycastEipAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnycastEipAddress.__pulumiType;
    }

    /**
     * Anycast EIP instance name.
     */
    public readonly anycastEipAddressName!: pulumi.Output<string | undefined>;
    /**
     * The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Anycast EIP instance description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
     */
    public readonly internetChargeType!: pulumi.Output<string | undefined>;
    /**
     * The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string | undefined>;
    /**
     * Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
     */
    public readonly serviceLocation!: pulumi.Output<string>;
    /**
     * The IP status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AnycastEipAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnycastEipAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnycastEipAddressArgs | AnycastEipAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AnycastEipAddressState | undefined;
            inputs["anycastEipAddressName"] = state ? state.anycastEipAddressName : undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            inputs["paymentType"] = state ? state.paymentType : undefined;
            inputs["serviceLocation"] = state ? state.serviceLocation : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AnycastEipAddressArgs | undefined;
            if ((!args || args.serviceLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceLocation'");
            }
            inputs["anycastEipAddressName"] = args ? args.anycastEipAddressName : undefined;
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            inputs["paymentType"] = args ? args.paymentType : undefined;
            inputs["serviceLocation"] = args ? args.serviceLocation : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AnycastEipAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnycastEipAddress resources.
 */
export interface AnycastEipAddressState {
    /**
     * Anycast EIP instance name.
     */
    readonly anycastEipAddressName?: pulumi.Input<string>;
    /**
     * The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Anycast EIP instance description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    readonly paymentType?: pulumi.Input<string>;
    /**
     * Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
     */
    readonly serviceLocation?: pulumi.Input<string>;
    /**
     * The IP status.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AnycastEipAddress resource.
 */
export interface AnycastEipAddressArgs {
    /**
     * Anycast EIP instance name.
     */
    readonly anycastEipAddressName?: pulumi.Input<string>;
    /**
     * The peak bandwidth of the Anycast EIP instance, in Mbps. It can not be changed when the internetChargeType is `PayByBandwidth` and the default value is 200.
     */
    readonly bandwidth?: pulumi.Input<number>;
    /**
     * Anycast EIP instance description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic. Valid value: `PayByBandwidth`.
     */
    readonly internetChargeType?: pulumi.Input<string>;
    /**
     * The payment model of Anycast EIP instance. `PayAsYouGo`: Refers to the post-paid mode. Valid value: `PayAsYouGo`. Default value is `PayAsYouGo`.
     */
    readonly paymentType?: pulumi.Input<string>;
    /**
     * Anycast EIP instance access area. `international`: Refers to areas outside of Mainland China.
     */
    readonly serviceLocation: pulumi.Input<string>;
}
