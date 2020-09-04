// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager Container Registry repositories.
 *
 * > **NOTE:** Available in v1.35.0+.
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
 * const my_repo = new alicloud.cr.Repo("my-repo", {
 *     namespace: my_namespace.name,
 *     summary: "this is summary of my new repo",
 *     repoType: "PUBLIC",
 *     detail: "this is a public repo",
 * });
 * ```
 */
export class Repo extends pulumi.CustomResource {
    /**
     * Get an existing Repo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepoState, opts?: pulumi.CustomResourceOptions): Repo {
        return new Repo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cr/repo:Repo';

    /**
     * Returns true if the given object is an instance of Repo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repo.__pulumiType;
    }

    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    public readonly detail!: pulumi.Output<string | undefined>;
    /**
     * The repository domain list.
     */
    public /*out*/ readonly domainList!: pulumi.Output<outputs.cr.RepoDomainList>;
    /**
     * Name of container registry repository.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of container registry namespace where repository is located.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    public readonly repoType!: pulumi.Output<string>;
    /**
     * The repository general information. It can contain 1 to 80 characters.
     */
    public readonly summary!: pulumi.Output<string>;

    /**
     * Create a Repo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepoArgs | RepoState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepoState | undefined;
            inputs["detail"] = state ? state.detail : undefined;
            inputs["domainList"] = state ? state.domainList : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
            inputs["repoType"] = state ? state.repoType : undefined;
            inputs["summary"] = state ? state.summary : undefined;
        } else {
            const args = argsOrState as RepoArgs | undefined;
            if (!args || args.namespace === undefined) {
                throw new Error("Missing required property 'namespace'");
            }
            if (!args || args.repoType === undefined) {
                throw new Error("Missing required property 'repoType'");
            }
            if (!args || args.summary === undefined) {
                throw new Error("Missing required property 'summary'");
            }
            inputs["detail"] = args ? args.detail : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["repoType"] = args ? args.repoType : undefined;
            inputs["summary"] = args ? args.summary : undefined;
            inputs["domainList"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Repo.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repo resources.
 */
export interface RepoState {
    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    readonly detail?: pulumi.Input<string>;
    /**
     * The repository domain list.
     */
    readonly domainList?: pulumi.Input<inputs.cr.RepoDomainList>;
    /**
     * Name of container registry repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of container registry namespace where repository is located.
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    readonly repoType?: pulumi.Input<string>;
    /**
     * The repository general information. It can contain 1 to 80 characters.
     */
    readonly summary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Repo resource.
 */
export interface RepoArgs {
    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    readonly detail?: pulumi.Input<string>;
    /**
     * Name of container registry repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of container registry namespace where repository is located.
     */
    readonly namespace: pulumi.Input<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    readonly repoType: pulumi.Input<string>;
    /**
     * The repository general information. It can contain 1 to 80 characters.
     */
    readonly summary: pulumi.Input<string>;
}
