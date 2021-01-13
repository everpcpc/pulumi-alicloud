// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a ROS Change Set resource.
 *
 * For information about ROS Change Set and how to use it, see [What is Change Set](https://www.alibabacloud.com/help/doc-detail/131051.htm).
 *
 * > **NOTE:** Available in v1.105.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ros.ChangeSet("example", {
 *     changeSetName: "example_value",
 *     changeSetType: "CREATE",
 *     description: "Test From Terraform",
 *     stackName: "tf-testacc",
 *     templateBody: "{\"ROSTemplateFormatVersion\":\"2015-09-01\"}",
 * });
 * ```
 *
 * ## Import
 *
 * ROS Change Set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ros/changeSet:ChangeSet example <change_set_id>
 * ```
 */
export class ChangeSet extends pulumi.CustomResource {
    /**
     * Get an existing ChangeSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ChangeSetState, opts?: pulumi.CustomResourceOptions): ChangeSet {
        return new ChangeSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ros/changeSet:ChangeSet';

    /**
     * Returns true if the given object is an instance of ChangeSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChangeSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChangeSet.__pulumiType;
    }

    /**
     * The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    public readonly changeSetName!: pulumi.Output<string>;
    /**
     * The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
     */
    public readonly changeSetType!: pulumi.Output<string | undefined>;
    /**
     * The description of the change set. The description can be up to 1,024 bytes in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    public readonly disableRollback!: pulumi.Output<boolean | undefined>;
    /**
     * The notification urls.
     */
    public readonly notificationUrls!: pulumi.Output<string[] | undefined>;
    /**
     * Parameters.
     */
    public readonly parameters!: pulumi.Output<outputs.ros.ChangeSetParameter[]>;
    /**
     * The ram role name.
     */
    public readonly ramRoleName!: pulumi.Output<string | undefined>;
    /**
     * The replacement option.
     */
    public readonly replacementOption!: pulumi.Output<string | undefined>;
    /**
     * The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    public readonly stackName!: pulumi.Output<string | undefined>;
    /**
     * The stack policy body.
     */
    public readonly stackPolicyBody!: pulumi.Output<string | undefined>;
    /**
     * The stack policy during update body.
     */
    public readonly stackPolicyDuringUpdateBody!: pulumi.Output<string | undefined>;
    /**
     * The stack policy during update url.
     */
    public readonly stackPolicyDuringUpdateUrl!: pulumi.Output<string | undefined>;
    /**
     * The stack policy url.
     */
    public readonly stackPolicyUrl!: pulumi.Output<string | undefined>;
    /**
     * The status of the change set.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
     */
    public readonly templateBody!: pulumi.Output<string | undefined>;
    /**
     * The template url.
     */
    public readonly templateUrl!: pulumi.Output<string | undefined>;
    /**
     * Timeout In Minutes.
     */
    public readonly timeoutInMinutes!: pulumi.Output<number>;
    /**
     * The use previous parameters.
     */
    public readonly usePreviousParameters!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ChangeSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChangeSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ChangeSetArgs | ChangeSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ChangeSetState | undefined;
            inputs["changeSetName"] = state ? state.changeSetName : undefined;
            inputs["changeSetType"] = state ? state.changeSetType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableRollback"] = state ? state.disableRollback : undefined;
            inputs["notificationUrls"] = state ? state.notificationUrls : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["ramRoleName"] = state ? state.ramRoleName : undefined;
            inputs["replacementOption"] = state ? state.replacementOption : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["stackName"] = state ? state.stackName : undefined;
            inputs["stackPolicyBody"] = state ? state.stackPolicyBody : undefined;
            inputs["stackPolicyDuringUpdateBody"] = state ? state.stackPolicyDuringUpdateBody : undefined;
            inputs["stackPolicyDuringUpdateUrl"] = state ? state.stackPolicyDuringUpdateUrl : undefined;
            inputs["stackPolicyUrl"] = state ? state.stackPolicyUrl : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["templateBody"] = state ? state.templateBody : undefined;
            inputs["templateUrl"] = state ? state.templateUrl : undefined;
            inputs["timeoutInMinutes"] = state ? state.timeoutInMinutes : undefined;
            inputs["usePreviousParameters"] = state ? state.usePreviousParameters : undefined;
        } else {
            const args = argsOrState as ChangeSetArgs | undefined;
            if ((!args || args.changeSetName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'changeSetName'");
            }
            inputs["changeSetName"] = args ? args.changeSetName : undefined;
            inputs["changeSetType"] = args ? args.changeSetType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableRollback"] = args ? args.disableRollback : undefined;
            inputs["notificationUrls"] = args ? args.notificationUrls : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["ramRoleName"] = args ? args.ramRoleName : undefined;
            inputs["replacementOption"] = args ? args.replacementOption : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["stackName"] = args ? args.stackName : undefined;
            inputs["stackPolicyBody"] = args ? args.stackPolicyBody : undefined;
            inputs["stackPolicyDuringUpdateBody"] = args ? args.stackPolicyDuringUpdateBody : undefined;
            inputs["stackPolicyDuringUpdateUrl"] = args ? args.stackPolicyDuringUpdateUrl : undefined;
            inputs["stackPolicyUrl"] = args ? args.stackPolicyUrl : undefined;
            inputs["templateBody"] = args ? args.templateBody : undefined;
            inputs["templateUrl"] = args ? args.templateUrl : undefined;
            inputs["timeoutInMinutes"] = args ? args.timeoutInMinutes : undefined;
            inputs["usePreviousParameters"] = args ? args.usePreviousParameters : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ChangeSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ChangeSet resources.
 */
export interface ChangeSetState {
    /**
     * The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    readonly changeSetName?: pulumi.Input<string>;
    /**
     * The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
     */
    readonly changeSetType?: pulumi.Input<string>;
    /**
     * The description of the change set. The description can be up to 1,024 bytes in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    readonly disableRollback?: pulumi.Input<boolean>;
    /**
     * The notification urls.
     */
    readonly notificationUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Parameters.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.ros.ChangeSetParameter>[]>;
    /**
     * The ram role name.
     */
    readonly ramRoleName?: pulumi.Input<string>;
    /**
     * The replacement option.
     */
    readonly replacementOption?: pulumi.Input<string>;
    /**
     * The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    readonly stackName?: pulumi.Input<string>;
    /**
     * The stack policy body.
     */
    readonly stackPolicyBody?: pulumi.Input<string>;
    /**
     * The stack policy during update body.
     */
    readonly stackPolicyDuringUpdateBody?: pulumi.Input<string>;
    /**
     * The stack policy during update url.
     */
    readonly stackPolicyDuringUpdateUrl?: pulumi.Input<string>;
    /**
     * The stack policy url.
     */
    readonly stackPolicyUrl?: pulumi.Input<string>;
    /**
     * The status of the change set.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
     */
    readonly templateBody?: pulumi.Input<string>;
    /**
     * The template url.
     */
    readonly templateUrl?: pulumi.Input<string>;
    /**
     * Timeout In Minutes.
     */
    readonly timeoutInMinutes?: pulumi.Input<number>;
    /**
     * The use previous parameters.
     */
    readonly usePreviousParameters?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ChangeSet resource.
 */
export interface ChangeSetArgs {
    /**
     * The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
     */
    readonly changeSetName: pulumi.Input<string>;
    /**
     * The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
     */
    readonly changeSetType?: pulumi.Input<string>;
    /**
     * The description of the change set. The description can be up to 1,024 bytes in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    readonly disableRollback?: pulumi.Input<boolean>;
    /**
     * The notification urls.
     */
    readonly notificationUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Parameters.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.ros.ChangeSetParameter>[]>;
    /**
     * The ram role name.
     */
    readonly ramRoleName?: pulumi.Input<string>;
    /**
     * The replacement option.
     */
    readonly replacementOption?: pulumi.Input<string>;
    /**
     * The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
     */
    readonly stackName?: pulumi.Input<string>;
    /**
     * The stack policy body.
     */
    readonly stackPolicyBody?: pulumi.Input<string>;
    /**
     * The stack policy during update body.
     */
    readonly stackPolicyDuringUpdateBody?: pulumi.Input<string>;
    /**
     * The stack policy during update url.
     */
    readonly stackPolicyDuringUpdateUrl?: pulumi.Input<string>;
    /**
     * The stack policy url.
     */
    readonly stackPolicyUrl?: pulumi.Input<string>;
    /**
     * The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
     */
    readonly templateBody?: pulumi.Input<string>;
    /**
     * The template url.
     */
    readonly templateUrl?: pulumi.Input<string>;
    /**
     * Timeout In Minutes.
     */
    readonly timeoutInMinutes?: pulumi.Input<number>;
    /**
     * The use previous parameters.
     */
    readonly usePreviousParameters?: pulumi.Input<boolean>;
}
