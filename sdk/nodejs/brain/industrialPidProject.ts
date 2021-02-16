// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Brain Industrial Pid Project resource.
 *
 * > **NOTE:** Available in v1.113.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.brain.IndustrialPidProject("example", {
 *     pidOrganisationId: "3e74e684-cbb5-xxxx",
 *     pidProjectName: "tf-testAcc",
 * });
 * ```
 *
 * ## Import
 *
 * Brain Industrial Pid Project can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:brain/industrialPidProject:IndustrialPidProject example <id>
 * ```
 */
export class IndustrialPidProject extends pulumi.CustomResource {
    /**
     * Get an existing IndustrialPidProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IndustrialPidProjectState, opts?: pulumi.CustomResourceOptions): IndustrialPidProject {
        return new IndustrialPidProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:brain/industrialPidProject:IndustrialPidProject';

    /**
     * Returns true if the given object is an instance of IndustrialPidProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IndustrialPidProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IndustrialPidProject.__pulumiType;
    }

    /**
     * The ID of Pid Organisation.
     */
    public readonly pidOrganisationId!: pulumi.Output<string>;
    /**
     * The description of Pid Project.
     */
    public readonly pidProjectDesc!: pulumi.Output<string | undefined>;
    /**
     * The name of Pid Project.
     */
    public readonly pidProjectName!: pulumi.Output<string>;

    /**
     * Create a IndustrialPidProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IndustrialPidProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IndustrialPidProjectArgs | IndustrialPidProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IndustrialPidProjectState | undefined;
            inputs["pidOrganisationId"] = state ? state.pidOrganisationId : undefined;
            inputs["pidProjectDesc"] = state ? state.pidProjectDesc : undefined;
            inputs["pidProjectName"] = state ? state.pidProjectName : undefined;
        } else {
            const args = argsOrState as IndustrialPidProjectArgs | undefined;
            if ((!args || args.pidOrganisationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidOrganisationId'");
            }
            if ((!args || args.pidProjectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidProjectName'");
            }
            inputs["pidOrganisationId"] = args ? args.pidOrganisationId : undefined;
            inputs["pidProjectDesc"] = args ? args.pidProjectDesc : undefined;
            inputs["pidProjectName"] = args ? args.pidProjectName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IndustrialPidProject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IndustrialPidProject resources.
 */
export interface IndustrialPidProjectState {
    /**
     * The ID of Pid Organisation.
     */
    readonly pidOrganisationId?: pulumi.Input<string>;
    /**
     * The description of Pid Project.
     */
    readonly pidProjectDesc?: pulumi.Input<string>;
    /**
     * The name of Pid Project.
     */
    readonly pidProjectName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IndustrialPidProject resource.
 */
export interface IndustrialPidProjectArgs {
    /**
     * The ID of Pid Organisation.
     */
    readonly pidOrganisationId: pulumi.Input<string>;
    /**
     * The description of Pid Project.
     */
    readonly pidProjectDesc?: pulumi.Input<string>;
    /**
     * The name of Pid Project.
     */
    readonly pidProjectName: pulumi.Input<string>;
}
