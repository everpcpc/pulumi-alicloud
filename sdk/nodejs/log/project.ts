// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The project is the resource management unit in Log Service and is used to isolate and control resources.
 * You can manage all the logs and the related log sources of an application by using projects. [Refer to details](https://www.alibabacloud.com/help/doc-detail/48873.htm).
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.log.Project("example", {
 *     description: "created by terraform",
 *     tags: {
 *         test: "test",
 *     },
 * });
 * ```
 *
 * Project With Policy Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const examplePolicy = new alicloud.log.Project("examplePolicy", {
 *     description: "created by terraform",
 *     policy: `{
 *   "Statement": [
 *     {
 *       "Action": [
 *         "log:PostLogStoreLogs"
 *       ],
 *       "Condition": {
 *         "StringNotLike": {
 *           "acs:SourceVpc": [
 *             "vpc-*"
 *           ]
 *         }
 *       },
 *       "Effect": "Deny",
 *       "Resource": "acs:log:*:*:project/tf-log/*"
 *     }
 *   ],
 *   "Version": "1"
 * }
 *
 * `,
 * });
 * ```
 * ## Module Support
 *
 * You can use the existing sls module
 * to create SLS project, store and store index one-click, like ECS instances.
 *
 * ## Import
 *
 * Log project can be imported using the id or name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:log/project:Project example tf-log
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Description of the log project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * Log project tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Description of the log project.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    policy?: pulumi.Input<string>;
    /**
     * Log project tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Description of the log project.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    policy?: pulumi.Input<string>;
    /**
     * Log project tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
