// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleState, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/schedule:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly launchExpirationTime!: pulumi.Output<number | undefined>;
    public readonly launchTime!: pulumi.Output<string>;
    public readonly recurrenceEndTime!: pulumi.Output<string>;
    public readonly recurrenceType!: pulumi.Output<string>;
    public readonly recurrenceValue!: pulumi.Output<string>;
    public readonly scheduledAction!: pulumi.Output<string>;
    public readonly scheduledTaskName!: pulumi.Output<string | undefined>;
    public readonly taskEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs | ScheduleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScheduleState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["launchExpirationTime"] = state ? state.launchExpirationTime : undefined;
            inputs["launchTime"] = state ? state.launchTime : undefined;
            inputs["recurrenceEndTime"] = state ? state.recurrenceEndTime : undefined;
            inputs["recurrenceType"] = state ? state.recurrenceType : undefined;
            inputs["recurrenceValue"] = state ? state.recurrenceValue : undefined;
            inputs["scheduledAction"] = state ? state.scheduledAction : undefined;
            inputs["scheduledTaskName"] = state ? state.scheduledTaskName : undefined;
            inputs["taskEnabled"] = state ? state.taskEnabled : undefined;
        } else {
            const args = argsOrState as ScheduleArgs | undefined;
            if (!args || args.launchTime === undefined) {
                throw new Error("Missing required property 'launchTime'");
            }
            if (!args || args.scheduledAction === undefined) {
                throw new Error("Missing required property 'scheduledAction'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["launchExpirationTime"] = args ? args.launchExpirationTime : undefined;
            inputs["launchTime"] = args ? args.launchTime : undefined;
            inputs["recurrenceEndTime"] = args ? args.recurrenceEndTime : undefined;
            inputs["recurrenceType"] = args ? args.recurrenceType : undefined;
            inputs["recurrenceValue"] = args ? args.recurrenceValue : undefined;
            inputs["scheduledAction"] = args ? args.scheduledAction : undefined;
            inputs["scheduledTaskName"] = args ? args.scheduledTaskName : undefined;
            inputs["taskEnabled"] = args ? args.taskEnabled : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Schedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Schedule resources.
 */
export interface ScheduleState {
    readonly description?: pulumi.Input<string>;
    readonly launchExpirationTime?: pulumi.Input<number>;
    readonly launchTime?: pulumi.Input<string>;
    readonly recurrenceEndTime?: pulumi.Input<string>;
    readonly recurrenceType?: pulumi.Input<string>;
    readonly recurrenceValue?: pulumi.Input<string>;
    readonly scheduledAction?: pulumi.Input<string>;
    readonly scheduledTaskName?: pulumi.Input<string>;
    readonly taskEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    readonly description?: pulumi.Input<string>;
    readonly launchExpirationTime?: pulumi.Input<number>;
    readonly launchTime: pulumi.Input<string>;
    readonly recurrenceEndTime?: pulumi.Input<string>;
    readonly recurrenceType?: pulumi.Input<string>;
    readonly recurrenceValue?: pulumi.Input<string>;
    readonly scheduledAction: pulumi.Input<string>;
    readonly scheduledTaskName?: pulumi.Input<string>;
    readonly taskEnabled?: pulumi.Input<boolean>;
}
