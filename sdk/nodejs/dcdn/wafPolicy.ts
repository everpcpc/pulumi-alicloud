// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DCDN Waf Policy resource.
 *
 * For information about DCDN Waf Policy and how to use it, see [What is Waf Policy](https://www.alibabacloud.com/help/en/dynamic-route-for-cdn/latest/set-the-protection-policies#doc-api-dcdn-CreateDcdnWafPolicy).
 *
 * > **NOTE:** Available in v1.184.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.dcdn.WafPolicy("example", {
 *     defenseScene: "waf_group",
 *     policyName: _var.name,
 *     policyType: "custom",
 *     status: "on",
 * });
 * ```
 *
 * ## Import
 *
 * DCDN Waf Policy can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dcdn/wafPolicy:WafPolicy example <id>
 * ```
 */
export class WafPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WafPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafPolicyState, opts?: pulumi.CustomResourceOptions): WafPolicy {
        return new WafPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dcdn/wafPolicy:WafPolicy';

    /**
     * Returns true if the given object is an instance of WafPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafPolicy.__pulumiType;
    }

    /**
     * The type of protection policy. Valid values: `wafGroup`, `customAcl`, `whitelist`.
     */
    public readonly defenseScene!: pulumi.Output<string>;
    /**
     * The name of the protection policy. The name must be 1 to 64 characters in length, and can contain letters, digits,and underscores (_).
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * The type of the protection policy. Valid values: `default`, `custom`.
     */
    public readonly policyType!: pulumi.Output<string>;
    /**
     * The status of the resource. Valid values: `on`, `off`.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a WafPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WafPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafPolicyArgs | WafPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafPolicyState | undefined;
            resourceInputs["defenseScene"] = state ? state.defenseScene : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as WafPolicyArgs | undefined;
            if ((!args || args.defenseScene === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defenseScene'");
            }
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            if ((!args || args.policyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyType'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["defenseScene"] = args ? args.defenseScene : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafPolicy resources.
 */
export interface WafPolicyState {
    /**
     * The type of protection policy. Valid values: `wafGroup`, `customAcl`, `whitelist`.
     */
    defenseScene?: pulumi.Input<string>;
    /**
     * The name of the protection policy. The name must be 1 to 64 characters in length, and can contain letters, digits,and underscores (_).
     */
    policyName?: pulumi.Input<string>;
    /**
     * The type of the protection policy. Valid values: `default`, `custom`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `on`, `off`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafPolicy resource.
 */
export interface WafPolicyArgs {
    /**
     * The type of protection policy. Valid values: `wafGroup`, `customAcl`, `whitelist`.
     */
    defenseScene: pulumi.Input<string>;
    /**
     * The name of the protection policy. The name must be 1 to 64 characters in length, and can contain letters, digits,and underscores (_).
     */
    policyName: pulumi.Input<string>;
    /**
     * The type of the protection policy. Valid values: `default`, `custom`.
     */
    policyType: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `on`, `off`.
     */
    status: pulumi.Input<string>;
}