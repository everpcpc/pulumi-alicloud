// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource provides a alarm rule resource and it can be used to monitor several cloud services according different metrics.
 * Details for [alarm rule](https://www.alibabacloud.com/help/doc-detail/28608.htm).
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const basic = new alicloud.cms.Alarm("basic", {
 *     contactGroups: ["test-group"],
 *     dimensions: {
 *         device: "/dev/vda1,/dev/vdb1",
 *         instanceId: "i-bp1247,i-bp11gd",
 *     },
 *     effectiveInterval: "0:00-2:00",
 *     escalationsCritical: {
 *         comparisonOperator: "<=",
 *         statistics: "Average",
 *         threshold: "35",
 *         times: 2,
 *     },
 *     metric: "disk_writebytes",
 *     period: 900,
 *     project: "acs_ecs_dashboard",
 *     webhook: pulumi.interpolate`https://${alicloud_account_current.id}.eu-central-1.fc.aliyuncs.com/2016-08-15/proxy/Terraform/AlarmEndpointMock/`,
 * });
 * ```
 *
 * ## Import
 *
 * Alarm rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cms/alarm:Alarm alarm abc12345
 * ```
 */
export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmState, opts?: pulumi.CustomResourceOptions): Alarm {
        return new Alarm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/alarm:Alarm';

    /**
     * Returns true if the given object is an instance of Alarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alarm.__pulumiType;
    }

    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    public readonly contactGroups!: pulumi.Output<string[]>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    public readonly dimensions!: pulumi.Output<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    public readonly effectiveInterval!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable alarm rule. Default to true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * A configuration of critical alarm (documented below).
     */
    public readonly escalationsCritical!: pulumi.Output<outputs.cms.AlarmEscalationsCritical | undefined>;
    /**
     * A configuration of critical info (documented below).
     */
    public readonly escalationsInfo!: pulumi.Output<outputs.cms.AlarmEscalationsInfo | undefined>;
    /**
     * A configuration of critical warn (documented below).
     */
    public readonly escalationsWarn!: pulumi.Output<outputs.cms.AlarmEscalationsWarn | undefined>;
    /**
     * Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    public readonly metric!: pulumi.Output<string>;
    /**
     * The alarm rule name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
     *
     * @deprecated Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
     */
    public readonly operator!: pulumi.Output<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
     */
    public readonly silenceTime!: pulumi.Output<number | undefined>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
     *
     * @deprecated Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
     */
    public readonly statistics!: pulumi.Output<string>;
    /**
     * The current alarm rule status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Critical level alarm threshold value, which must be a numeric value currently.
     *
     * @deprecated Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
     */
    public readonly threshold!: pulumi.Output<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
     *
     * @deprecated Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
     */
    public readonly triggeredCount!: pulumi.Output<number>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    public readonly webhook!: pulumi.Output<string | undefined>;

    /**
     * Create a Alarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmArgs | AlarmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmState | undefined;
            inputs["contactGroups"] = state ? state.contactGroups : undefined;
            inputs["dimensions"] = state ? state.dimensions : undefined;
            inputs["effectiveInterval"] = state ? state.effectiveInterval : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["escalationsCritical"] = state ? state.escalationsCritical : undefined;
            inputs["escalationsInfo"] = state ? state.escalationsInfo : undefined;
            inputs["escalationsWarn"] = state ? state.escalationsWarn : undefined;
            inputs["metric"] = state ? state.metric : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["operator"] = state ? state.operator : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["silenceTime"] = state ? state.silenceTime : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["statistics"] = state ? state.statistics : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["threshold"] = state ? state.threshold : undefined;
            inputs["triggeredCount"] = state ? state.triggeredCount : undefined;
            inputs["webhook"] = state ? state.webhook : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if ((!args || args.contactGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactGroups'");
            }
            if ((!args || args.dimensions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dimensions'");
            }
            if ((!args || args.metric === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metric'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["contactGroups"] = args ? args.contactGroups : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["effectiveInterval"] = args ? args.effectiveInterval : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["escalationsCritical"] = args ? args.escalationsCritical : undefined;
            inputs["escalationsInfo"] = args ? args.escalationsInfo : undefined;
            inputs["escalationsWarn"] = args ? args.escalationsWarn : undefined;
            inputs["metric"] = args ? args.metric : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["operator"] = args ? args.operator : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["silenceTime"] = args ? args.silenceTime : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["statistics"] = args ? args.statistics : undefined;
            inputs["threshold"] = args ? args.threshold : undefined;
            inputs["triggeredCount"] = args ? args.triggeredCount : undefined;
            inputs["webhook"] = args ? args.webhook : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Alarm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alarm resources.
 */
export interface AlarmState {
    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    dimensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    effectiveInterval?: pulumi.Input<string>;
    /**
     * Whether to enable alarm rule. Default to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    endTime?: pulumi.Input<number>;
    /**
     * A configuration of critical alarm (documented below).
     */
    escalationsCritical?: pulumi.Input<inputs.cms.AlarmEscalationsCritical>;
    /**
     * A configuration of critical info (documented below).
     */
    escalationsInfo?: pulumi.Input<inputs.cms.AlarmEscalationsInfo>;
    /**
     * A configuration of critical warn (documented below).
     */
    escalationsWarn?: pulumi.Input<inputs.cms.AlarmEscalationsWarn>;
    /**
     * Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metric?: pulumi.Input<string>;
    /**
     * The alarm rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
     *
     * @deprecated Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
     */
    operator?: pulumi.Input<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    period?: pulumi.Input<number>;
    /**
     * Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    project?: pulumi.Input<string>;
    /**
     * Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
     */
    silenceTime?: pulumi.Input<number>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
     *
     * @deprecated Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
     */
    statistics?: pulumi.Input<string>;
    /**
     * The current alarm rule status.
     */
    status?: pulumi.Input<string>;
    /**
     * Critical level alarm threshold value, which must be a numeric value currently.
     *
     * @deprecated Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
     */
    threshold?: pulumi.Input<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
     *
     * @deprecated Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
     */
    triggeredCount?: pulumi.Input<number>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    webhook?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    contactGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    dimensions: pulumi.Input<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    effectiveInterval?: pulumi.Input<string>;
    /**
     * Whether to enable alarm rule. Default to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'end_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    endTime?: pulumi.Input<number>;
    /**
     * A configuration of critical alarm (documented below).
     */
    escalationsCritical?: pulumi.Input<inputs.cms.AlarmEscalationsCritical>;
    /**
     * A configuration of critical info (documented below).
     */
    escalationsInfo?: pulumi.Input<inputs.cms.AlarmEscalationsInfo>;
    /**
     * A configuration of critical warn (documented below).
     */
    escalationsWarn?: pulumi.Input<inputs.cms.AlarmEscalationsWarn>;
    /**
     * Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metric: pulumi.Input<string>;
    /**
     * The alarm rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.comparison_operator' instead.
     *
     * @deprecated Field 'operator' has been deprecated from provider version 1.94.0. New field 'escalations_critical.comparison_operator' instead.
     */
    operator?: pulumi.Input<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    period?: pulumi.Input<number>;
    /**
     * Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    project: pulumi.Input<string>;
    /**
     * Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
     */
    silenceTime?: pulumi.Input<number>;
    /**
     * It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
     *
     * @deprecated Field 'start_time' has been deprecated from provider version 1.50.0. New field 'effective_interval' instead.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Critical level alarm statistics method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum", "Value", "ErrorCodeMaximum", "Sum", "Count"]. Default to "Average".
     *
     * @deprecated Field 'statistics' has been deprecated from provider version 1.94.0. New field 'escalations_critical.statistics' instead.
     */
    statistics?: pulumi.Input<string>;
    /**
     * Critical level alarm threshold value, which must be a numeric value currently.
     *
     * @deprecated Field 'threshold' has been deprecated from provider version 1.94.0. New field 'escalations_critical.threshold' instead.
     */
    threshold?: pulumi.Input<string>;
    /**
     * It has been deprecated from provider version 1.94.0 and 'escalations_critical.times' instead.
     *
     * @deprecated Field 'triggered_count' has been deprecated from provider version 1.94.0. New field 'escalations_critical.times' instead.
     */
    triggeredCount?: pulumi.Input<number>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    webhook?: pulumi.Input<string>;
}
