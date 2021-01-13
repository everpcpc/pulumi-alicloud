// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * ESS scaling rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ess/scalingRule:ScalingRule example abc123456
 * ```
 */
export class ScalingRule extends pulumi.CustomResource {
    /**
     * Get an existing ScalingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingRuleState, opts?: pulumi.CustomResourceOptions): ScalingRule {
        return new ScalingRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/scalingRule:ScalingRule';

    /**
     * Returns true if the given object is an instance of ScalingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingRule.__pulumiType;
    }

    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    public readonly adjustmentType!: pulumi.Output<string | undefined>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    public readonly adjustmentValue!: pulumi.Output<number | undefined>;
    public /*out*/ readonly ari!: pulumi.Output<string>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    public readonly cooldown!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    public readonly disableScaleIn!: pulumi.Output<boolean | undefined>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    public readonly estimatedInstanceWarmup!: pulumi.Output<number>;
    /**
     * A CloudMonitor metric name.
     */
    public readonly metricName!: pulumi.Output<string | undefined>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    public readonly scalingRuleName!: pulumi.Output<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
     */
    public readonly scalingRuleType!: pulumi.Output<string | undefined>;
    /**
     * Steps for StepScalingRule. See Block stepAdjustment below for details.
     */
    public readonly stepAdjustments!: pulumi.Output<outputs.ess.ScalingRuleStepAdjustment[] | undefined>;
    /**
     * The target value for the metric.
     */
    public readonly targetValue!: pulumi.Output<number | undefined>;

    /**
     * Create a ScalingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingRuleArgs | ScalingRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScalingRuleState | undefined;
            inputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            inputs["adjustmentValue"] = state ? state.adjustmentValue : undefined;
            inputs["ari"] = state ? state.ari : undefined;
            inputs["cooldown"] = state ? state.cooldown : undefined;
            inputs["disableScaleIn"] = state ? state.disableScaleIn : undefined;
            inputs["estimatedInstanceWarmup"] = state ? state.estimatedInstanceWarmup : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            inputs["scalingRuleName"] = state ? state.scalingRuleName : undefined;
            inputs["scalingRuleType"] = state ? state.scalingRuleType : undefined;
            inputs["stepAdjustments"] = state ? state.stepAdjustments : undefined;
            inputs["targetValue"] = state ? state.targetValue : undefined;
        } else {
            const args = argsOrState as ScalingRuleArgs | undefined;
            if ((!args || args.scalingGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            inputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            inputs["adjustmentValue"] = args ? args.adjustmentValue : undefined;
            inputs["cooldown"] = args ? args.cooldown : undefined;
            inputs["disableScaleIn"] = args ? args.disableScaleIn : undefined;
            inputs["estimatedInstanceWarmup"] = args ? args.estimatedInstanceWarmup : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            inputs["scalingRuleName"] = args ? args.scalingRuleName : undefined;
            inputs["scalingRuleType"] = args ? args.scalingRuleType : undefined;
            inputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            inputs["targetValue"] = args ? args.targetValue : undefined;
            inputs["ari"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScalingRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingRule resources.
 */
export interface ScalingRuleState {
    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    readonly adjustmentType?: pulumi.Input<string>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    readonly adjustmentValue?: pulumi.Input<number>;
    readonly ari?: pulumi.Input<string>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    readonly cooldown?: pulumi.Input<number>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    readonly disableScaleIn?: pulumi.Input<boolean>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    readonly estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * A CloudMonitor metric name.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    readonly scalingRuleName?: pulumi.Input<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
     */
    readonly scalingRuleType?: pulumi.Input<string>;
    /**
     * Steps for StepScalingRule. See Block stepAdjustment below for details.
     */
    readonly stepAdjustments?: pulumi.Input<pulumi.Input<inputs.ess.ScalingRuleStepAdjustment>[]>;
    /**
     * The target value for the metric.
     */
    readonly targetValue?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ScalingRule resource.
 */
export interface ScalingRuleArgs {
    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    readonly adjustmentType?: pulumi.Input<string>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    readonly adjustmentValue?: pulumi.Input<number>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    readonly cooldown?: pulumi.Input<number>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    readonly disableScaleIn?: pulumi.Input<boolean>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    readonly estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * A CloudMonitor metric name.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    readonly scalingGroupId: pulumi.Input<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    readonly scalingRuleName?: pulumi.Input<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
     */
    readonly scalingRuleType?: pulumi.Input<string>;
    /**
     * Steps for StepScalingRule. See Block stepAdjustment below for details.
     */
    readonly stepAdjustments?: pulumi.Input<pulumi.Input<inputs.ess.ScalingRuleStepAdjustment>[]>;
    /**
     * The target value for the metric.
     */
    readonly targetValue?: pulumi.Input<number>;
}
