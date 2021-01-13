// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The project is the basic unit of operation in maxcompute. It is similar to the concept of Database or Schema in traditional databases, and sets the boundary for maxcompute multi-user isolation and access control. [Refer to details](https://www.alibabacloud.com/help/doc-detail/27818.html).
 *
 * ->**NOTE:** Available in 1.77.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.maxcompute.Project("example", {
 *     orderType: "PayAsYouGo",
 *     projectName: "tf_maxcompute_project",
 *     specificationType: "OdpsStandard",
 * });
 * ```
 *
 * ## Import
 *
 * MaxCompute project can be imported using the *name* or ID, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:maxcompute/project:Project example tf_maxcompute_project
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
    public static readonly __pulumiType = 'alicloud:maxcompute/project:Project';

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
     * It has been deprecated from provider version 1.110.0 and `projectName` instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of payment, only `PayAsYouGo` supported currently.
     */
    public readonly orderType!: pulumi.Output<string>;
    /**
     * The name of the maxcompute project.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The type of resource Specification, only `OdpsStandard` supported currently.
     */
    public readonly specificationType!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["orderType"] = state ? state.orderType : undefined;
            inputs["projectName"] = state ? state.projectName : undefined;
            inputs["specificationType"] = state ? state.specificationType : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.orderType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'orderType'");
            }
            if ((!args || args.specificationType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'specificationType'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["orderType"] = args ? args.orderType : undefined;
            inputs["projectName"] = args ? args.projectName : undefined;
            inputs["specificationType"] = args ? args.specificationType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * It has been deprecated from provider version 1.110.0 and `projectName` instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of payment, only `PayAsYouGo` supported currently.
     */
    readonly orderType?: pulumi.Input<string>;
    /**
     * The name of the maxcompute project.
     */
    readonly projectName?: pulumi.Input<string>;
    /**
     * The type of resource Specification, only `OdpsStandard` supported currently.
     */
    readonly specificationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * It has been deprecated from provider version 1.110.0 and `projectName` instead.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The type of payment, only `PayAsYouGo` supported currently.
     */
    readonly orderType: pulumi.Input<string>;
    /**
     * The name of the maxcompute project.
     */
    readonly projectName?: pulumi.Input<string>;
    /**
     * The type of resource Specification, only `OdpsStandard` supported currently.
     */
    readonly specificationType: pulumi.Input<string>;
}
