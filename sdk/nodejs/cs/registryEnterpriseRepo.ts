// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource will help you to manager Container Registry Enterprise Edition repositories.
 *
 * For information about Container Registry Enterprise Edition repository and how to use it, see [Create a Repository](https://www.alibabacloud.com/help/doc-detail/145291.htm)
 *
 * > **NOTE:** Available in v1.86.0+.
 *
 * > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const my_namespace = new alicloud.cs.RegistryEnterpriseNamespace("my-namespace", {
 *     instanceId: "cri-xxx",
 *     autoCreate: false,
 *     defaultVisibility: "PUBLIC",
 * });
 * const my_repo = new alicloud.cs.RegistryEnterpriseRepo("my-repo", {
 *     instanceId: my_namespace.instanceId,
 *     namespace: my_namespace.name,
 *     summary: "this is summary of my new repo",
 *     repoType: "PUBLIC",
 *     detail: "this is a public repo",
 * });
 * ```
 *
 * ## Import
 *
 * Container Registry Enterprise Edition repository can be imported using the `{instance_id}:{namespace}:{repository}`, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo default `cri-xxx:my-namespace:my-repo`
 * ```
 */
export class RegistryEnterpriseRepo extends pulumi.CustomResource {
    /**
     * Get an existing RegistryEnterpriseRepo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryEnterpriseRepoState, opts?: pulumi.CustomResourceOptions): RegistryEnterpriseRepo {
        return new RegistryEnterpriseRepo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo';

    /**
     * Returns true if the given object is an instance of RegistryEnterpriseRepo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryEnterpriseRepo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryEnterpriseRepo.__pulumiType;
    }

    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    public readonly detail!: pulumi.Output<string | undefined>;
    /**
     * ID of Container Registry Enterprise Edition instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * The uuid of Container Registry Enterprise Edition repository.
     */
    public /*out*/ readonly repoId!: pulumi.Output<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    public readonly repoType!: pulumi.Output<string>;
    /**
     * The repository general information. It can contain 1 to 100 characters.
     */
    public readonly summary!: pulumi.Output<string>;

    /**
     * Create a RegistryEnterpriseRepo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryEnterpriseRepoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryEnterpriseRepoArgs | RegistryEnterpriseRepoState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegistryEnterpriseRepoState | undefined;
            inputs["detail"] = state ? state.detail : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
            inputs["repoId"] = state ? state.repoId : undefined;
            inputs["repoType"] = state ? state.repoType : undefined;
            inputs["summary"] = state ? state.summary : undefined;
        } else {
            const args = argsOrState as RegistryEnterpriseRepoArgs | undefined;
            if ((!args || args.instanceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.namespace === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.repoType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'repoType'");
            }
            if ((!args || args.summary === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'summary'");
            }
            inputs["detail"] = args ? args.detail : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["repoType"] = args ? args.repoType : undefined;
            inputs["summary"] = args ? args.summary : undefined;
            inputs["repoId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegistryEnterpriseRepo.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryEnterpriseRepo resources.
 */
export interface RegistryEnterpriseRepoState {
    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    readonly detail?: pulumi.Input<string>;
    /**
     * ID of Container Registry Enterprise Edition instance.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * The uuid of Container Registry Enterprise Edition repository.
     */
    readonly repoId?: pulumi.Input<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    readonly repoType?: pulumi.Input<string>;
    /**
     * The repository general information. It can contain 1 to 100 characters.
     */
    readonly summary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegistryEnterpriseRepo resource.
 */
export interface RegistryEnterpriseRepoArgs {
    /**
     * The repository specific information. MarkDown format is supported, and the length limit is 2000.
     */
    readonly detail?: pulumi.Input<string>;
    /**
     * ID of Container Registry Enterprise Edition instance.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
     */
    readonly namespace: pulumi.Input<string>;
    /**
     * `PUBLIC` or `PRIVATE`, repo's visibility.
     */
    readonly repoType: pulumi.Input<string>;
    /**
     * The repository general information. It can contain 1 to 100 characters.
     */
    readonly summary: pulumi.Input<string>;
}
