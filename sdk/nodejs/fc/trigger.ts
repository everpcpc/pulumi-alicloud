// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Function Compute trigger can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    public readonly config!: pulumi.Output<string | undefined>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    public readonly configMns!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute function name.
     */
    public readonly function!: pulumi.Output<string>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    public readonly sourceArn!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute trigger ID.
     */
    public /*out*/ readonly triggerId!: pulumi.Output<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["configMns"] = state ? state.configMns : undefined;
            inputs["function"] = state ? state.function : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["sourceArn"] = state ? state.sourceArn : undefined;
            inputs["triggerId"] = state ? state.triggerId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if ((!args || args.function === undefined) && !opts.urn) {
                throw new Error("Missing required property 'function'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["configMns"] = args ? args.configMns : undefined;
            inputs["function"] = args ? args.function : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["sourceArn"] = args ? args.sourceArn : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["lastModified"] = undefined /*out*/;
            inputs["triggerId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    lastModified?: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service?: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Function Compute trigger ID.
     */
    triggerId?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
     */
    config?: pulumi.Input<string>;
    /**
     * The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
     */
    configMns?: pulumi.Input<string>;
    /**
     * The Function Compute function name.
     */
    function: pulumi.Input<string>;
    /**
     * The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only trigger name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    role?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service: pulumi.Input<string>;
    /**
     * Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
     */
    sourceArn?: pulumi.Input<string>;
    /**
     * The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
     */
    type: pulumi.Input<string>;
}
