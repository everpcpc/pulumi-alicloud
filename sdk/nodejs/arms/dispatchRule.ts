// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Alert Dispatch Rule resource.
 *
 * For information about Application Real-Time Monitoring Service (ARMS) Alert Dispatch Rule and how to use it, see [What is Alert Dispatch_Rule](https://www.alibabacloud.com/help/en/doc-detail/203146.htm).
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAlertContact = new alicloud.arms.AlertContact("defaultAlertContact", {
 *     alertContactName: "example_value",
 *     email: "example_value@aaa.com",
 * });
 * const defaultAlertContactGroup = new alicloud.arms.AlertContactGroup("defaultAlertContactGroup", {
 *     alertContactGroupName: "example_value",
 *     contactIds: [defaultAlertContact.id],
 * });
 * const defaultDispatchRule = new alicloud.arms.DispatchRule("defaultDispatchRule", {
 *     dispatchRuleName: "example_value",
 *     dispatchType: "CREATE_ALERT",
 *     groupRules: [{
 *         groupWaitTime: 5,
 *         groupInterval: 15,
 *         repeatInterval: 100,
 *         groupingFields: ["alertname"],
 *     }],
 *     labelMatchExpressionGrids: [{
 *         labelMatchExpressionGroups: [{
 *             labelMatchExpressions: [{
 *                 key: "_aliyun_arms_involvedObject_kind",
 *                 value: "app",
 *                 operator: "eq",
 *             }],
 *         }],
 *     }],
 *     notifyRules: [{
 *         notifyObjects: [
 *             {
 *                 notifyObjectId: defaultAlertContact.id,
 *                 notifyType: "ARMS_CONTACT",
 *                 name: "example_value",
 *             },
 *             {
 *                 notifyObjectId: defaultAlertContactGroup.id,
 *                 notifyType: "ARMS_CONTACT_GROUP",
 *                 name: "example_value",
 *             },
 *         ],
 *         notifyChannels: [
 *             "dingTalk",
 *             "wechat",
 *         ],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Application Real-Time Monitoring Service (ARMS) Alert Contact can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:arms/dispatchRule:DispatchRule example <id>
 * ```
 */
export class DispatchRule extends pulumi.CustomResource {
    /**
     * Get an existing DispatchRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DispatchRuleState, opts?: pulumi.CustomResourceOptions): DispatchRule {
        return new DispatchRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:arms/dispatchRule:DispatchRule';

    /**
     * Returns true if the given object is an instance of DispatchRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DispatchRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DispatchRule.__pulumiType;
    }

    /**
     * The name of the dispatch policy.
     */
    public readonly dispatchRuleName!: pulumi.Output<string>;
    /**
     * The alert handling method. Valid values: CREATE_ALERT: generates an alert. DISCARD_ALERT: discards the alert event and generates no alert.
     */
    public readonly dispatchType!: pulumi.Output<string | undefined>;
    /**
     * Sets the event group. See the following `Block groupRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    public readonly groupRules!: pulumi.Output<outputs.arms.DispatchRuleGroupRule[]>;
    /**
     * Specifies whether to send the restored alert. Valid values: true: sends the alert. false: does not send the alert.
     */
    public readonly isRecover!: pulumi.Output<boolean | undefined>;
    /**
     * Sets the dispatch rule. See the following `Block labelMatchExpressionGrid`.
     */
    public readonly labelMatchExpressionGrids!: pulumi.Output<outputs.arms.DispatchRuleLabelMatchExpressionGrid[]>;
    /**
     * Sets the notification rule. See the following `Block notifyRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    public readonly notifyRules!: pulumi.Output<outputs.arms.DispatchRuleNotifyRule[]>;
    /**
     * The resource status of Alert Dispatch Rule.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a DispatchRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DispatchRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DispatchRuleArgs | DispatchRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DispatchRuleState | undefined;
            inputs["dispatchRuleName"] = state ? state.dispatchRuleName : undefined;
            inputs["dispatchType"] = state ? state.dispatchType : undefined;
            inputs["groupRules"] = state ? state.groupRules : undefined;
            inputs["isRecover"] = state ? state.isRecover : undefined;
            inputs["labelMatchExpressionGrids"] = state ? state.labelMatchExpressionGrids : undefined;
            inputs["notifyRules"] = state ? state.notifyRules : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DispatchRuleArgs | undefined;
            if ((!args || args.dispatchRuleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dispatchRuleName'");
            }
            if ((!args || args.groupRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupRules'");
            }
            if ((!args || args.labelMatchExpressionGrids === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labelMatchExpressionGrids'");
            }
            if ((!args || args.notifyRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifyRules'");
            }
            inputs["dispatchRuleName"] = args ? args.dispatchRuleName : undefined;
            inputs["dispatchType"] = args ? args.dispatchType : undefined;
            inputs["groupRules"] = args ? args.groupRules : undefined;
            inputs["isRecover"] = args ? args.isRecover : undefined;
            inputs["labelMatchExpressionGrids"] = args ? args.labelMatchExpressionGrids : undefined;
            inputs["notifyRules"] = args ? args.notifyRules : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DispatchRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DispatchRule resources.
 */
export interface DispatchRuleState {
    /**
     * The name of the dispatch policy.
     */
    dispatchRuleName?: pulumi.Input<string>;
    /**
     * The alert handling method. Valid values: CREATE_ALERT: generates an alert. DISCARD_ALERT: discards the alert event and generates no alert.
     */
    dispatchType?: pulumi.Input<string>;
    /**
     * Sets the event group. See the following `Block groupRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    groupRules?: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleGroupRule>[]>;
    /**
     * Specifies whether to send the restored alert. Valid values: true: sends the alert. false: does not send the alert.
     */
    isRecover?: pulumi.Input<boolean>;
    /**
     * Sets the dispatch rule. See the following `Block labelMatchExpressionGrid`.
     */
    labelMatchExpressionGrids?: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleLabelMatchExpressionGrid>[]>;
    /**
     * Sets the notification rule. See the following `Block notifyRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    notifyRules?: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleNotifyRule>[]>;
    /**
     * The resource status of Alert Dispatch Rule.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DispatchRule resource.
 */
export interface DispatchRuleArgs {
    /**
     * The name of the dispatch policy.
     */
    dispatchRuleName: pulumi.Input<string>;
    /**
     * The alert handling method. Valid values: CREATE_ALERT: generates an alert. DISCARD_ALERT: discards the alert event and generates no alert.
     */
    dispatchType?: pulumi.Input<string>;
    /**
     * Sets the event group. See the following `Block groupRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    groupRules: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleGroupRule>[]>;
    /**
     * Specifies whether to send the restored alert. Valid values: true: sends the alert. false: does not send the alert.
     */
    isRecover?: pulumi.Input<boolean>;
    /**
     * Sets the dispatch rule. See the following `Block labelMatchExpressionGrid`.
     */
    labelMatchExpressionGrids: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleLabelMatchExpressionGrid>[]>;
    /**
     * Sets the notification rule. See the following `Block notifyRules`. It will be ignored  when `dispatchType = "DISCARD_ALERT"`.
     */
    notifyRules: pulumi.Input<pulumi.Input<inputs.arms.DispatchRuleNotifyRule>[]>;
}
