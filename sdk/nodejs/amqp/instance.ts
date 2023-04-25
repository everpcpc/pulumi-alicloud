// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RabbitMQ (AMQP) Instance resource.
 *
 * For information about RabbitMQ (AMQP) Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/doc-detail/101631.htm).
 *
 * > **NOTE:** Available in v1.128.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const professional = new alicloud.amqp.Instance("professional", {
 *     instanceType: "professional",
 *     maxEipTps: "128",
 *     maxTps: "1000",
 *     paymentType: "Subscription",
 *     period: 1,
 *     queueCapacity: "50",
 *     supportEip: true,
 * });
 * const vip = new alicloud.amqp.Instance("vip", {
 *     instanceType: "vip",
 *     maxEipTps: "128",
 *     maxTps: "5000",
 *     paymentType: "Subscription",
 *     period: 1,
 *     queueCapacity: "50",
 *     storageSize: "700",
 *     supportEip: true,
 * });
 * ```
 *
 * ## Import
 *
 * RabbitMQ (AMQP) Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:amqp/instance:Instance example <id>
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:amqp/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    public readonly logistics!: pulumi.Output<string | undefined>;
    /**
     * The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
     */
    public readonly maxEipTps!: pulumi.Output<string | undefined>;
    /**
     * The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     */
    public readonly maxTps!: pulumi.Output<string>;
    /**
     * The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     */
    public readonly modifyType!: pulumi.Output<string | undefined>;
    /**
     * The payment type. Valid values: `Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The queue capacity. The smallest value is 50 and the step size 5.
     */
    public readonly queueCapacity!: pulumi.Output<string>;
    /**
     * RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     */
    public readonly renewalDuration!: pulumi.Output<number | undefined>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    public readonly renewalDurationUnit!: pulumi.Output<string | undefined>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     */
    public readonly renewalStatus!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The storage size. It is valid when `instanceType` is vip.
     */
    public readonly storageSize!: pulumi.Output<string | undefined>;
    /**
     * Whether to support EIP.
     */
    public readonly supportEip!: pulumi.Output<boolean>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["logistics"] = state ? state.logistics : undefined;
            resourceInputs["maxEipTps"] = state ? state.maxEipTps : undefined;
            resourceInputs["maxTps"] = state ? state.maxTps : undefined;
            resourceInputs["modifyType"] = state ? state.modifyType : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["queueCapacity"] = state ? state.queueCapacity : undefined;
            resourceInputs["renewalDuration"] = state ? state.renewalDuration : undefined;
            resourceInputs["renewalDurationUnit"] = state ? state.renewalDurationUnit : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["supportEip"] = state ? state.supportEip : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.maxTps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxTps'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            if ((!args || args.queueCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queueCapacity'");
            }
            if ((!args || args.supportEip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'supportEip'");
            }
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["logistics"] = args ? args.logistics : undefined;
            resourceInputs["maxEipTps"] = args ? args.maxEipTps : undefined;
            resourceInputs["maxTps"] = args ? args.maxTps : undefined;
            resourceInputs["modifyType"] = args ? args.modifyType : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["queueCapacity"] = args ? args.queueCapacity : undefined;
            resourceInputs["renewalDuration"] = args ? args.renewalDuration : undefined;
            resourceInputs["renewalDurationUnit"] = args ? args.renewalDurationUnit : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["supportEip"] = args ? args.supportEip : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     */
    instanceType?: pulumi.Input<string>;
    logistics?: pulumi.Input<string>;
    /**
     * The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
     */
    maxEipTps?: pulumi.Input<string>;
    /**
     * The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     */
    maxTps?: pulumi.Input<string>;
    /**
     * The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     */
    modifyType?: pulumi.Input<string>;
    /**
     * The payment type. Valid values: `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     */
    period?: pulumi.Input<number>;
    /**
     * The queue capacity. The smallest value is 50 and the step size 5.
     */
    queueCapacity?: pulumi.Input<string>;
    /**
     * RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     */
    renewalDuration?: pulumi.Input<number>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    renewalDurationUnit?: pulumi.Input<string>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The storage size. It is valid when `instanceType` is vip.
     */
    storageSize?: pulumi.Input<string>;
    /**
     * Whether to support EIP.
     */
    supportEip?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     */
    instanceType: pulumi.Input<string>;
    logistics?: pulumi.Input<string>;
    /**
     * The max eip tps. It is valid when `supportEip` is true. The valid value is [128, 45000] with the step size 128.
     */
    maxEipTps?: pulumi.Input<string>;
    /**
     * The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     */
    maxTps: pulumi.Input<string>;
    /**
     * The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     */
    modifyType?: pulumi.Input<string>;
    /**
     * The payment type. Valid values: `Subscription`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     */
    period?: pulumi.Input<number>;
    /**
     * The queue capacity. The smallest value is 50 and the step size 5.
     */
    queueCapacity: pulumi.Input<string>;
    /**
     * RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     */
    renewalDuration?: pulumi.Input<number>;
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     */
    renewalDurationUnit?: pulumi.Input<string>;
    /**
     * Whether to renew an instance automatically or not. Default to "ManualRenewal".
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The storage size. It is valid when `instanceType` is vip.
     */
    storageSize?: pulumi.Input<string>;
    /**
     * Whether to support EIP.
     */
    supportEip: pulumi.Input<boolean>;
}
