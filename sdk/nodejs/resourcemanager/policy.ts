// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Resource Manager Policy resource.\
 * For information about Resource Manager Policy and how to use it, see [What is Resource Manager Policy](https://www.alibabacloud.com/help/en/doc-detail/93732.htm).
 *
 * > **NOTE:** Available in v1.83.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.resourcemanager.Policy("example", {
 *     policyDocument: `		{
 * 			"Statement": [{
 * 				"Action": ["oss:*"],
 * 				"Effect": "Allow",
 * 				"Resource": ["acs:oss:*:*:*"]
 * 			}],
 * 			"Version": "1"
 * 		}
 *     `,
 *     policyName: "abc12345",
 * });
 * ```
 *
 * ## Import
 *
 * Resource Manager Policy can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:resourcemanager/policy:Policy example abc12345
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:resourcemanager/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * The version of the policy. Default to v1.
     *
     * @deprecated Field 'default_version' has been deprecated from provider version 1.90.0
     */
    public readonly defaultVersion!: pulumi.Output<string>;
    /**
     * The description of the policy. The description must be 1 to 1,024 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The content of the policy. The content must be 1 to 2,048 characters in length.
     */
    public readonly policyDocument!: pulumi.Output<string>;
    /**
     * The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * The type of the policy. Valid values: `Custom`, `System`.
     */
    public /*out*/ readonly policyType!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["policyDocument"] = state ? state.policyDocument : undefined;
            inputs["policyName"] = state ? state.policyName : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.policyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDocument'");
            }
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            inputs["defaultVersion"] = args ? args.defaultVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["policyDocument"] = args ? args.policyDocument : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policyType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * The version of the policy. Default to v1.
     *
     * @deprecated Field 'default_version' has been deprecated from provider version 1.90.0
     */
    readonly defaultVersion?: pulumi.Input<string>;
    /**
     * The description of the policy. The description must be 1 to 1,024 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The content of the policy. The content must be 1 to 2,048 characters in length.
     */
    readonly policyDocument?: pulumi.Input<string>;
    /**
     * The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     */
    readonly policyName?: pulumi.Input<string>;
    /**
     * The type of the policy. Valid values: `Custom`, `System`.
     */
    readonly policyType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The version of the policy. Default to v1.
     *
     * @deprecated Field 'default_version' has been deprecated from provider version 1.90.0
     */
    readonly defaultVersion?: pulumi.Input<string>;
    /**
     * The description of the policy. The description must be 1 to 1,024 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The content of the policy. The content must be 1 to 2,048 characters in length.
     */
    readonly policyDocument: pulumi.Input<string>;
    /**
     * The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
     */
    readonly policyName: pulumi.Input<string>;
}
