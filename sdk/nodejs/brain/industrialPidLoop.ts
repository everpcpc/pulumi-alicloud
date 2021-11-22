// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Brain Industrial Pid Loop resource.
 *
 * > **NOTE:** Available in v1.117.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.brain.IndustrialPidLoop("example", {
 *     pidLoopConfiguration: "YourLoopConfiguration",
 *     pidLoopDcsType: "standard",
 *     pidLoopIsCrucial: true,
 *     pidLoopName: "tf-testAcc",
 *     pidLoopType: "0",
 *     pidProjectId: "856c6b8f-ca63-40a4-xxxx-xxxx",
 * });
 * ```
 *
 * ## Import
 *
 * Brain Industrial Pid Loop can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:brain/industrialPidLoop:IndustrialPidLoop example <id>
 * ```
 */
export class IndustrialPidLoop extends pulumi.CustomResource {
    /**
     * Get an existing IndustrialPidLoop resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IndustrialPidLoopState, opts?: pulumi.CustomResourceOptions): IndustrialPidLoop {
        return new IndustrialPidLoop(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:brain/industrialPidLoop:IndustrialPidLoop';

    /**
     * Returns true if the given object is an instance of IndustrialPidLoop.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IndustrialPidLoop {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IndustrialPidLoop.__pulumiType;
    }

    /**
     * The Pid Loop Configuration.
     */
    public readonly pidLoopConfiguration!: pulumi.Output<string>;
    /**
     * The dcs type of Pid Loop. Valid values: `standard`.
     */
    public readonly pidLoopDcsType!: pulumi.Output<string>;
    /**
     * The desc of Pid Loop.
     */
    public readonly pidLoopDesc!: pulumi.Output<string | undefined>;
    /**
     * Whether is crucial Pid Loop.
     */
    public readonly pidLoopIsCrucial!: pulumi.Output<boolean>;
    /**
     * The name of Pid Loop.
     */
    public readonly pidLoopName!: pulumi.Output<string>;
    /**
     * The type of Pid Loop. Valid values: `0`, `1`, `2`, `3`, `4`, `5`.
     */
    public readonly pidLoopType!: pulumi.Output<string>;
    /**
     * The pid project id.
     */
    public readonly pidProjectId!: pulumi.Output<string>;
    /**
     * The status of Pid Loop.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a IndustrialPidLoop resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IndustrialPidLoopArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IndustrialPidLoopArgs | IndustrialPidLoopState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IndustrialPidLoopState | undefined;
            inputs["pidLoopConfiguration"] = state ? state.pidLoopConfiguration : undefined;
            inputs["pidLoopDcsType"] = state ? state.pidLoopDcsType : undefined;
            inputs["pidLoopDesc"] = state ? state.pidLoopDesc : undefined;
            inputs["pidLoopIsCrucial"] = state ? state.pidLoopIsCrucial : undefined;
            inputs["pidLoopName"] = state ? state.pidLoopName : undefined;
            inputs["pidLoopType"] = state ? state.pidLoopType : undefined;
            inputs["pidProjectId"] = state ? state.pidProjectId : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as IndustrialPidLoopArgs | undefined;
            if ((!args || args.pidLoopConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidLoopConfiguration'");
            }
            if ((!args || args.pidLoopDcsType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidLoopDcsType'");
            }
            if ((!args || args.pidLoopIsCrucial === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidLoopIsCrucial'");
            }
            if ((!args || args.pidLoopName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidLoopName'");
            }
            if ((!args || args.pidLoopType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidLoopType'");
            }
            if ((!args || args.pidProjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pidProjectId'");
            }
            inputs["pidLoopConfiguration"] = args ? args.pidLoopConfiguration : undefined;
            inputs["pidLoopDcsType"] = args ? args.pidLoopDcsType : undefined;
            inputs["pidLoopDesc"] = args ? args.pidLoopDesc : undefined;
            inputs["pidLoopIsCrucial"] = args ? args.pidLoopIsCrucial : undefined;
            inputs["pidLoopName"] = args ? args.pidLoopName : undefined;
            inputs["pidLoopType"] = args ? args.pidLoopType : undefined;
            inputs["pidProjectId"] = args ? args.pidProjectId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IndustrialPidLoop.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IndustrialPidLoop resources.
 */
export interface IndustrialPidLoopState {
    /**
     * The Pid Loop Configuration.
     */
    pidLoopConfiguration?: pulumi.Input<string>;
    /**
     * The dcs type of Pid Loop. Valid values: `standard`.
     */
    pidLoopDcsType?: pulumi.Input<string>;
    /**
     * The desc of Pid Loop.
     */
    pidLoopDesc?: pulumi.Input<string>;
    /**
     * Whether is crucial Pid Loop.
     */
    pidLoopIsCrucial?: pulumi.Input<boolean>;
    /**
     * The name of Pid Loop.
     */
    pidLoopName?: pulumi.Input<string>;
    /**
     * The type of Pid Loop. Valid values: `0`, `1`, `2`, `3`, `4`, `5`.
     */
    pidLoopType?: pulumi.Input<string>;
    /**
     * The pid project id.
     */
    pidProjectId?: pulumi.Input<string>;
    /**
     * The status of Pid Loop.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IndustrialPidLoop resource.
 */
export interface IndustrialPidLoopArgs {
    /**
     * The Pid Loop Configuration.
     */
    pidLoopConfiguration: pulumi.Input<string>;
    /**
     * The dcs type of Pid Loop. Valid values: `standard`.
     */
    pidLoopDcsType: pulumi.Input<string>;
    /**
     * The desc of Pid Loop.
     */
    pidLoopDesc?: pulumi.Input<string>;
    /**
     * Whether is crucial Pid Loop.
     */
    pidLoopIsCrucial: pulumi.Input<boolean>;
    /**
     * The name of Pid Loop.
     */
    pidLoopName: pulumi.Input<string>;
    /**
     * The type of Pid Loop. Valid values: `0`, `1`, `2`, `3`, `4`, `5`.
     */
    pidLoopType: pulumi.Input<string>;
    /**
     * The pid project id.
     */
    pidProjectId: pulumi.Input<string>;
}
