// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Quotas Quota Application resource.
 *
 * For information about Quotas Quota Application and how to use it, see [What is Quota Application](https://help.aliyun.com/document_detail/171289.html).
 *
 * > **NOTE:** Available in v1.117.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.quotas.QuotaApplication("example", {
 *     desireValue: 100,
 *     dimensions: [{
 *         key: "regionId",
 *         value: "cn-hangzhou",
 *     }],
 *     noticeType: 0,
 *     productCode: "ess",
 *     quotaActionCode: "q_db_instance",
 *     reason: "For Terraform Test",
 * });
 * ```
 *
 * ## Import
 *
 * Quotas Application Info can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:quotas/quotaApplication:QuotaApplication example <id>
 * ```
 */
export class QuotaApplication extends pulumi.CustomResource {
    /**
     * Get an existing QuotaApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuotaApplicationState, opts?: pulumi.CustomResourceOptions): QuotaApplication {
        return new QuotaApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:quotas/quotaApplication:QuotaApplication';

    /**
     * Returns true if the given object is an instance of QuotaApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QuotaApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QuotaApplication.__pulumiType;
    }

    /**
     * The approve value of the quota application.
     */
    public /*out*/ readonly approveValue!: pulumi.Output<string>;
    /**
     * The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
     */
    public readonly auditMode!: pulumi.Output<string | undefined>;
    /**
     * The audit reason.
     */
    public /*out*/ readonly auditReason!: pulumi.Output<string>;
    /**
     * The desire value of the quota application.
     */
    public readonly desireValue!: pulumi.Output<number>;
    /**
     * The quota dimensions.
     */
    public readonly dimensions!: pulumi.Output<outputs.quotas.QuotaApplicationDimension[] | undefined>;
    /**
     * The effective time of the quota application.
     */
    public /*out*/ readonly effectiveTime!: pulumi.Output<string>;
    /**
     * The expire time of the quota application.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * The notice type. Valid values: `0`, `1`, `2`, `3`.
     */
    public readonly noticeType!: pulumi.Output<number | undefined>;
    /**
     * The product code.
     */
    public readonly productCode!: pulumi.Output<string>;
    /**
     * The ID of quota action.
     */
    public readonly quotaActionCode!: pulumi.Output<string>;
    /**
     * The quota category. Valid values: `CommonQuota`, `FlowControl`.
     */
    public readonly quotaCategory!: pulumi.Output<string | undefined>;
    /**
     * The description of the quota application.
     */
    public /*out*/ readonly quotaDescription!: pulumi.Output<string>;
    /**
     * The name of the quota application.
     */
    public /*out*/ readonly quotaName!: pulumi.Output<string>;
    /**
     * The unit of the quota application.
     */
    public /*out*/ readonly quotaUnit!: pulumi.Output<string>;
    /**
     * The reason of the quota application.
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * The status of the quota application.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a QuotaApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuotaApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuotaApplicationArgs | QuotaApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuotaApplicationState | undefined;
            inputs["approveValue"] = state ? state.approveValue : undefined;
            inputs["auditMode"] = state ? state.auditMode : undefined;
            inputs["auditReason"] = state ? state.auditReason : undefined;
            inputs["desireValue"] = state ? state.desireValue : undefined;
            inputs["dimensions"] = state ? state.dimensions : undefined;
            inputs["effectiveTime"] = state ? state.effectiveTime : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["noticeType"] = state ? state.noticeType : undefined;
            inputs["productCode"] = state ? state.productCode : undefined;
            inputs["quotaActionCode"] = state ? state.quotaActionCode : undefined;
            inputs["quotaCategory"] = state ? state.quotaCategory : undefined;
            inputs["quotaDescription"] = state ? state.quotaDescription : undefined;
            inputs["quotaName"] = state ? state.quotaName : undefined;
            inputs["quotaUnit"] = state ? state.quotaUnit : undefined;
            inputs["reason"] = state ? state.reason : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as QuotaApplicationArgs | undefined;
            if ((!args || args.desireValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desireValue'");
            }
            if ((!args || args.productCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productCode'");
            }
            if ((!args || args.quotaActionCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'quotaActionCode'");
            }
            if ((!args || args.reason === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reason'");
            }
            inputs["auditMode"] = args ? args.auditMode : undefined;
            inputs["desireValue"] = args ? args.desireValue : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["noticeType"] = args ? args.noticeType : undefined;
            inputs["productCode"] = args ? args.productCode : undefined;
            inputs["quotaActionCode"] = args ? args.quotaActionCode : undefined;
            inputs["quotaCategory"] = args ? args.quotaCategory : undefined;
            inputs["reason"] = args ? args.reason : undefined;
            inputs["approveValue"] = undefined /*out*/;
            inputs["auditReason"] = undefined /*out*/;
            inputs["effectiveTime"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["quotaDescription"] = undefined /*out*/;
            inputs["quotaName"] = undefined /*out*/;
            inputs["quotaUnit"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(QuotaApplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QuotaApplication resources.
 */
export interface QuotaApplicationState {
    /**
     * The approve value of the quota application.
     */
    approveValue?: pulumi.Input<string>;
    /**
     * The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
     */
    auditMode?: pulumi.Input<string>;
    /**
     * The audit reason.
     */
    auditReason?: pulumi.Input<string>;
    /**
     * The desire value of the quota application.
     */
    desireValue?: pulumi.Input<number>;
    /**
     * The quota dimensions.
     */
    dimensions?: pulumi.Input<pulumi.Input<inputs.quotas.QuotaApplicationDimension>[]>;
    /**
     * The effective time of the quota application.
     */
    effectiveTime?: pulumi.Input<string>;
    /**
     * The expire time of the quota application.
     */
    expireTime?: pulumi.Input<string>;
    /**
     * The notice type. Valid values: `0`, `1`, `2`, `3`.
     */
    noticeType?: pulumi.Input<number>;
    /**
     * The product code.
     */
    productCode?: pulumi.Input<string>;
    /**
     * The ID of quota action.
     */
    quotaActionCode?: pulumi.Input<string>;
    /**
     * The quota category. Valid values: `CommonQuota`, `FlowControl`.
     */
    quotaCategory?: pulumi.Input<string>;
    /**
     * The description of the quota application.
     */
    quotaDescription?: pulumi.Input<string>;
    /**
     * The name of the quota application.
     */
    quotaName?: pulumi.Input<string>;
    /**
     * The unit of the quota application.
     */
    quotaUnit?: pulumi.Input<string>;
    /**
     * The reason of the quota application.
     */
    reason?: pulumi.Input<string>;
    /**
     * The status of the quota application.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QuotaApplication resource.
 */
export interface QuotaApplicationArgs {
    /**
     * The audit mode. Valid values: `Async`, `Sync`. Default to: `Async`.
     */
    auditMode?: pulumi.Input<string>;
    /**
     * The desire value of the quota application.
     */
    desireValue: pulumi.Input<number>;
    /**
     * The quota dimensions.
     */
    dimensions?: pulumi.Input<pulumi.Input<inputs.quotas.QuotaApplicationDimension>[]>;
    /**
     * The notice type. Valid values: `0`, `1`, `2`, `3`.
     */
    noticeType?: pulumi.Input<number>;
    /**
     * The product code.
     */
    productCode: pulumi.Input<string>;
    /**
     * The ID of quota action.
     */
    quotaActionCode: pulumi.Input<string>;
    /**
     * The quota category. Valid values: `CommonQuota`, `FlowControl`.
     */
    quotaCategory?: pulumi.Input<string>;
    /**
     * The reason of the quota application.
     */
    reason: pulumi.Input<string>;
}
