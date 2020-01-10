// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/log_store_index.html.markdown.
 */
export class StoreIndex extends pulumi.CustomResource {
    /**
     * Get an existing StoreIndex resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StoreIndexState, opts?: pulumi.CustomResourceOptions): StoreIndex {
        return new StoreIndex(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/storeIndex:StoreIndex';

    /**
     * Returns true if the given object is an instance of StoreIndex.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StoreIndex {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StoreIndex.__pulumiType;
    }

    /**
     * List configurations of field search index. Valid item as follows:
     */
    public readonly fieldSearches!: pulumi.Output<outputs.log.StoreIndexFieldSearch[] | undefined>;
    /**
     * The configuration of full text index. Valid item as follows:
     */
    public readonly fullText!: pulumi.Output<outputs.log.StoreIndexFullText | undefined>;
    /**
     * The log store name to the query index belongs.
     */
    public readonly logstore!: pulumi.Output<string>;
    /**
     * The project name to the log store belongs.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a StoreIndex resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StoreIndexArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StoreIndexArgs | StoreIndexState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StoreIndexState | undefined;
            inputs["fieldSearches"] = state ? state.fieldSearches : undefined;
            inputs["fullText"] = state ? state.fullText : undefined;
            inputs["logstore"] = state ? state.logstore : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as StoreIndexArgs | undefined;
            if (!args || args.logstore === undefined) {
                throw new Error("Missing required property 'logstore'");
            }
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            inputs["fieldSearches"] = args ? args.fieldSearches : undefined;
            inputs["fullText"] = args ? args.fullText : undefined;
            inputs["logstore"] = args ? args.logstore : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StoreIndex.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StoreIndex resources.
 */
export interface StoreIndexState {
    /**
     * List configurations of field search index. Valid item as follows:
     */
    readonly fieldSearches?: pulumi.Input<pulumi.Input<inputs.log.StoreIndexFieldSearch>[]>;
    /**
     * The configuration of full text index. Valid item as follows:
     */
    readonly fullText?: pulumi.Input<inputs.log.StoreIndexFullText>;
    /**
     * The log store name to the query index belongs.
     */
    readonly logstore?: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StoreIndex resource.
 */
export interface StoreIndexArgs {
    /**
     * List configurations of field search index. Valid item as follows:
     */
    readonly fieldSearches?: pulumi.Input<pulumi.Input<inputs.log.StoreIndexFieldSearch>[]>;
    /**
     * The configuration of full text index. Valid item as follows:
     */
    readonly fullText?: pulumi.Input<inputs.log.StoreIndexFullText>;
    /**
     * The log store name to the query index belongs.
     */
    readonly logstore: pulumi.Input<string>;
    /**
     * The project name to the log store belongs.
     */
    readonly project: pulumi.Input<string>;
}
