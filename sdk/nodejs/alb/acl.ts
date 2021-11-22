// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Application Load Balancer (ALB) Acl resource.
 *
 * For information about ALB Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/doc-detail/213617.htm).
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.alb.Acl("example", {
 *     aclName: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * ALB Acl can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:alb/acl:Acl example <id>
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alb/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * ACL Entries.
     */
    public readonly aclEntries!: pulumi.Output<outputs.alb.AclAclEntry[] | undefined>;
    /**
     * The name of the ACL. The name must be 2 to 128 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    public readonly aclName!: pulumi.Output<string>;
    /**
     * Specifies whether to precheck the API request. Valid values: `true`: only prechecks the API request. If you select this option, the specified endpoint service is not created after the request passes the precheck. The system prechecks the required parameters, request format, and service limits. If the request fails the precheck, the corresponding error message is returned. If the request passes the precheck, the DryRunOperation error code is returned. `false` (default): checks the request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Resource Group to Which the Number.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`.  `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            inputs["aclEntries"] = state ? state.aclEntries : undefined;
            inputs["aclName"] = state ? state.aclName : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.aclName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclName'");
            }
            inputs["aclEntries"] = args ? args.aclEntries : undefined;
            inputs["aclName"] = args ? args.aclName : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Acl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * ACL Entries.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be 2 to 128 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck the API request. Valid values: `true`: only prechecks the API request. If you select this option, the specified endpoint service is not created after the request passes the precheck. The system prechecks the required parameters, request format, and service limits. If the request fails the precheck, the corresponding error message is returned. If the request passes the precheck, the DryRunOperation error code is returned. `false` (default): checks the request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Resource Group to Which the Number.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The state of the ACL. Valid values:`Provisioning` , `Available` and `Configuring`.  `Provisioning`: The ACL is being created. `Available`: The ACL is available. `Configuring`: The ACL is being configured.
     */
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * ACL Entries.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of the ACL. The name must be 2 to 128 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). It must start with a letter.
     */
    aclName: pulumi.Input<string>;
    /**
     * Specifies whether to precheck the API request. Valid values: `true`: only prechecks the API request. If you select this option, the specified endpoint service is not created after the request passes the precheck. The system prechecks the required parameters, request format, and service limits. If the request fails the precheck, the corresponding error message is returned. If the request passes the precheck, the DryRunOperation error code is returned. `false` (default): checks the request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Resource Group to Which the Number.
     */
    resourceGroupId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}
