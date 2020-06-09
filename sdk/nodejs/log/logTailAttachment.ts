// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class LogTailAttachment extends pulumi.CustomResource {
    /**
     * Get an existing LogTailAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogTailAttachmentState, opts?: pulumi.CustomResourceOptions): LogTailAttachment {
        return new LogTailAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/logTailAttachment:LogTailAttachment';

    /**
     * Returns true if the given object is an instance of LogTailAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogTailAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogTailAttachment.__pulumiType;
    }

    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    public readonly logtailConfigName!: pulumi.Output<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    public readonly machineGroupName!: pulumi.Output<string>;
    /**
     * The project name to the log store belongs.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a LogTailAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogTailAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogTailAttachmentArgs | LogTailAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LogTailAttachmentState | undefined;
            inputs["logtailConfigName"] = state ? state.logtailConfigName : undefined;
            inputs["machineGroupName"] = state ? state.machineGroupName : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as LogTailAttachmentArgs | undefined;
            if (!args || args.logtailConfigName === undefined) {
                throw new Error("Missing required property 'logtailConfigName'");
            }
            if (!args || args.machineGroupName === undefined) {
                throw new Error("Missing required property 'machineGroupName'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            inputs["logtailConfigName"] = args ? args.logtailConfigName : undefined;
            inputs["machineGroupName"] = args ? args.machineGroupName : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LogTailAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogTailAttachment resources.
 */
export interface LogTailAttachmentState {
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly logtailConfigName?: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    readonly machineGroupName?: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogTailAttachment resource.
 */
export interface LogTailAttachmentArgs {
    /**
     * The Logtail configuration name, which is unique in the same project.
     */
    readonly logtailConfigName: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    readonly machineGroupName: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project: pulumi.Input<string>;
}
