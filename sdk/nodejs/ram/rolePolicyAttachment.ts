// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RAM Role attachment resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a RAM Role Policy attachment.
 * const role = new alicloud.ram.Role("role", {
 *     document: `    {
 *       "Statement": [
 *         {
 *           "Action": "sts:AssumeRole",
 *           "Effect": "Allow",
 *           "Principal": {
 *             "Service": [
 *               "apigateway.aliyuncs.com", 
 *               "ecs.aliyuncs.com"
 *             ]
 *           }
 *         }
 *       ],
 *       "Version": "1"
 *     }
 * `,
 *     description: "this is a role test.",
 *     force: true,
 * });
 * const policy = new alicloud.ram.Policy("policy", {
 *     document: `  {
 *     "Statement": [
 *       {
 *         "Action": [
 *           "oss:ListObjects",
 *           "oss:GetObject"
 *         ],
 *         "Effect": "Allow",
 *         "Resource": [
 *           "acs:oss:*:*:mybucket",
 *           "acs:oss:*:*:mybucket/*"
 *         ]
 *       }
 *     ],
 *       "Version": "1"
 *   }
 * `,
 *     description: "this is a policy test",
 *     force: true,
 * });
 * const attach = new alicloud.ram.RolePolicyAttachment("attach", {
 *     policyName: policy.name,
 *     policyType: policy.type,
 *     roleName: role.name,
 * });
 * ```
 *
 * ## Import
 *
 * RAM Role Policy attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ram/rolePolicyAttachment:RolePolicyAttachment example role:my-policy:Custom:my-role
 * ```
 */
export class RolePolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing RolePolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RolePolicyAttachmentState, opts?: pulumi.CustomResourceOptions): RolePolicyAttachment {
        return new RolePolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/rolePolicyAttachment:RolePolicyAttachment';

    /**
     * Returns true if the given object is an instance of RolePolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RolePolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RolePolicyAttachment.__pulumiType;
    }

    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    public readonly policyType!: pulumi.Output<string>;
    /**
     * Name of the RAM Role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
     */
    public readonly roleName!: pulumi.Output<string>;

    /**
     * Create a RolePolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RolePolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RolePolicyAttachmentArgs | RolePolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RolePolicyAttachmentState | undefined;
            inputs["policyName"] = state ? state.policyName : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
        } else {
            const args = argsOrState as RolePolicyAttachmentArgs | undefined;
            if ((!args || args.policyName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyName'");
            }
            if ((!args || args.policyType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policyType'");
            }
            if ((!args || args.roleName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RolePolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RolePolicyAttachment resources.
 */
export interface RolePolicyAttachmentState {
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly policyName?: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * Name of the RAM Role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
     */
    readonly roleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RolePolicyAttachment resource.
 */
export interface RolePolicyAttachmentArgs {
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly policyName: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    readonly policyType: pulumi.Input<string>;
    /**
     * Name of the RAM Role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
     */
    readonly roleName: pulumi.Input<string>;
}
