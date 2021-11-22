// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Serverless Workflow Flow resource.
 *
 * For information about Serverless Workflow Flow and how to use it, see [What is Flow](https://www.alibabacloud.com/help/en/doc-detail/123079.htm).
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
 * const _default = new alicloud.ram.Role("default", {document: `  {
 *     "Statement": [
 *       {
 *         "Action": "sts:AssumeRole",
 *         "Effect": "Allow",
 *         "Principal": {
 *           "Service": [
 *             "fnf.aliyuncs.com"
 *           ]
 *         }
 *       }
 *     ],
 *     "Version": "1"
 *   }
 * `});
 * const example = new alicloud.fnf.Flow("example", {
 *     definition: `  version: v1beta1
 *   type: flow
 *   steps:
 *     - type: pass
 *       name: helloworld
 * `,
 *     roleArn: _default.arn,
 *     description: "Test for terraform fnf_flow.",
 *     type: "FDL",
 * });
 * ```
 *
 * ## Import
 *
 * Serverless Workflow Flow can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fnf/flow:Flow example <name>
 * ```
 */
export class Flow extends pulumi.CustomResource {
    /**
     * Get an existing Flow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowState, opts?: pulumi.CustomResourceOptions): Flow {
        return new Flow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fnf/flow:Flow';

    /**
     * Returns true if the given object is an instance of Flow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Flow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Flow.__pulumiType;
    }

    /**
     * The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
     */
    public readonly definition!: pulumi.Output<string>;
    /**
     * The description of the flow.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The unique ID of the flow.
     */
    public /*out*/ readonly flowId!: pulumi.Output<string>;
    /**
     * The time when the flow was last modified.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The name of the flow. The name must be unique in an Alibaba Cloud account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The type of the flow. Valid values are `FDL` or `DEFAULT`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Flow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowArgs | FlowState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowState | undefined;
            inputs["definition"] = state ? state.definition : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["flowId"] = state ? state.flowId : undefined;
            inputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FlowArgs | undefined;
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["flowId"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Flow.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Flow resources.
 */
export interface FlowState {
    /**
     * The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
     */
    definition?: pulumi.Input<string>;
    /**
     * The description of the flow.
     */
    description?: pulumi.Input<string>;
    /**
     * The unique ID of the flow.
     */
    flowId?: pulumi.Input<string>;
    /**
     * The time when the flow was last modified.
     */
    lastModifiedTime?: pulumi.Input<string>;
    /**
     * The name of the flow. The name must be unique in an Alibaba Cloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The type of the flow. Valid values are `FDL` or `DEFAULT`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Flow resource.
 */
export interface FlowArgs {
    /**
     * The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
     */
    definition: pulumi.Input<string>;
    /**
     * The description of the flow.
     */
    description: pulumi.Input<string>;
    /**
     * The name of the flow. The name must be unique in an Alibaba Cloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The type of the flow. Valid values are `FDL` or `DEFAULT`.
     */
    type: pulumi.Input<string>;
}
