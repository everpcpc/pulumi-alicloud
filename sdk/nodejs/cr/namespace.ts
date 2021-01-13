// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager Container Registry namespaces.
 *
 * > **NOTE:** Available in v1.34.0+.
 *
 * > **NOTE:** You need to set your registry password in Container Registry console before use this resource.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const my_namespace = new alicloud.cr.Namespace("my-namespace", {
 *     autoCreate: false,
 *     defaultVisibility: "PUBLIC",
 * });
 * ```
 *
 * ## Import
 *
 * Container Registry namespace can be imported using the namespace, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cr/namespace:Namespace default my-namespace
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cr/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
     */
    public readonly autoCreate!: pulumi.Output<boolean>;
    /**
     * `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
     */
    public readonly defaultVisibility!: pulumi.Output<string>;
    /**
     * Name of Container Registry namespace.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            inputs["autoCreate"] = state ? state.autoCreate : undefined;
            inputs["defaultVisibility"] = state ? state.defaultVisibility : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.autoCreate === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'autoCreate'");
            }
            if ((!args || args.defaultVisibility === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'defaultVisibility'");
            }
            inputs["autoCreate"] = args ? args.autoCreate : undefined;
            inputs["defaultVisibility"] = args ? args.defaultVisibility : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
     */
    readonly autoCreate?: pulumi.Input<boolean>;
    /**
     * `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
     */
    readonly defaultVisibility?: pulumi.Input<string>;
    /**
     * Name of Container Registry namespace.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
     */
    readonly autoCreate: pulumi.Input<boolean>;
    /**
     * `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
     */
    readonly defaultVisibility: pulumi.Input<string>;
    /**
     * Name of Container Registry namespace.
     */
    readonly name?: pulumi.Input<string>;
}
