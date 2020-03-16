// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class LifecycleHook extends pulumi.CustomResource {
    /**
     * Get an existing LifecycleHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecycleHookState, opts?: pulumi.CustomResourceOptions): LifecycleHook {
        return new LifecycleHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/lifecycleHook:LifecycleHook';

    /**
     * Returns true if the given object is an instance of LifecycleHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecycleHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecycleHook.__pulumiType;
    }

    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
     */
    public readonly defaultResult!: pulumi.Output<string | undefined>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the defaultResult parameter. Default value: 600.
     */
    public readonly heartbeatTimeout!: pulumi.Output<number | undefined>;
    /**
     * Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     */
    public readonly lifecycleTransition!: pulumi.Output<string>;
    /**
     * The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Arn of notification target.
     */
    public readonly notificationArn!: pulumi.Output<string>;
    /**
     * Additional information that you want to include when Auto Scaling sends a message to the notification target.
     */
    public readonly notificationMetadata!: pulumi.Output<string>;
    /**
     * The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a LifecycleHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecycleHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecycleHookArgs | LifecycleHookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LifecycleHookState | undefined;
            inputs["defaultResult"] = state ? state.defaultResult : undefined;
            inputs["heartbeatTimeout"] = state ? state.heartbeatTimeout : undefined;
            inputs["lifecycleTransition"] = state ? state.lifecycleTransition : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationArn"] = state ? state.notificationArn : undefined;
            inputs["notificationMetadata"] = state ? state.notificationMetadata : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as LifecycleHookArgs | undefined;
            if (!args || args.lifecycleTransition === undefined) {
                throw new Error("Missing required property 'lifecycleTransition'");
            }
            if (!args || args.scalingGroupId === undefined) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            inputs["defaultResult"] = args ? args.defaultResult : undefined;
            inputs["heartbeatTimeout"] = args ? args.heartbeatTimeout : undefined;
            inputs["lifecycleTransition"] = args ? args.lifecycleTransition : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationArn"] = args ? args.notificationArn : undefined;
            inputs["notificationMetadata"] = args ? args.notificationMetadata : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LifecycleHook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecycleHook resources.
 */
export interface LifecycleHookState {
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
     */
    readonly defaultResult?: pulumi.Input<string>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the defaultResult parameter. Default value: 600.
     */
    readonly heartbeatTimeout?: pulumi.Input<number>;
    /**
     * Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     */
    readonly lifecycleTransition?: pulumi.Input<string>;
    /**
     * The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Arn of notification target.
     */
    readonly notificationArn?: pulumi.Input<string>;
    /**
     * Additional information that you want to include when Auto Scaling sends a message to the notification target.
     */
    readonly notificationMetadata?: pulumi.Input<string>;
    /**
     * The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LifecycleHook resource.
 */
export interface LifecycleHookArgs {
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.
     */
    readonly defaultResult?: pulumi.Input<string>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the defaultResult parameter. Default value: 600.
     */
    readonly heartbeatTimeout?: pulumi.Input<number>;
    /**
     * Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     */
    readonly lifecycleTransition: pulumi.Input<string>;
    /**
     * The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Arn of notification target.
     */
    readonly notificationArn?: pulumi.Input<string>;
    /**
     * Additional information that you want to include when Auto Scaling sends a message to the notification target.
     */
    readonly notificationMetadata?: pulumi.Input<string>;
    /**
     * The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     */
    readonly scalingGroupId: pulumi.Input<string>;
}
