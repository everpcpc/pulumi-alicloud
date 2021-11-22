// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Log machine group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:log/machineGroup:MachineGroup example tf-log:tf-machine-group
 * ```
 */
export class MachineGroup extends pulumi.CustomResource {
    /**
     * Get an existing MachineGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineGroupState, opts?: pulumi.CustomResourceOptions): MachineGroup {
        return new MachineGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/machineGroup:MachineGroup';

    /**
     * Returns true if the given object is an instance of MachineGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineGroup.__pulumiType;
    }

    /**
     * The specific machine identification, which can be an IP address or user-defined identity.
     */
    public readonly identifyLists!: pulumi.Output<string[]>;
    /**
     * The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
     */
    public readonly identifyType!: pulumi.Output<string | undefined>;
    /**
     * The machine group name, which is unique in the same project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project name to the machine group belongs.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The topic of a machine group.
     */
    public readonly topic!: pulumi.Output<string | undefined>;

    /**
     * Create a MachineGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachineGroupArgs | MachineGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineGroupState | undefined;
            inputs["identifyLists"] = state ? state.identifyLists : undefined;
            inputs["identifyType"] = state ? state.identifyType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as MachineGroupArgs | undefined;
            if ((!args || args.identifyLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifyLists'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["identifyLists"] = args ? args.identifyLists : undefined;
            inputs["identifyType"] = args ? args.identifyType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["topic"] = args ? args.topic : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MachineGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachineGroup resources.
 */
export interface MachineGroupState {
    /**
     * The specific machine identification, which can be an IP address or user-defined identity.
     */
    identifyLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
     */
    identifyType?: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    name?: pulumi.Input<string>;
    /**
     * The project name to the machine group belongs.
     */
    project?: pulumi.Input<string>;
    /**
     * The topic of a machine group.
     */
    topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MachineGroup resource.
 */
export interface MachineGroupArgs {
    /**
     * The specific machine identification, which can be an IP address or user-defined identity.
     */
    identifyLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
     */
    identifyType?: pulumi.Input<string>;
    /**
     * The machine group name, which is unique in the same project.
     */
    name?: pulumi.Input<string>;
    /**
     * The project name to the machine group belongs.
     */
    project: pulumi.Input<string>;
    /**
     * The topic of a machine group.
     */
    topic?: pulumi.Input<string>;
}
