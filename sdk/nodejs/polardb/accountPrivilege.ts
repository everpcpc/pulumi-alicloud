// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account_privilege.html.markdown.
 */
export class AccountPrivilege extends pulumi.CustomResource {
    /**
     * Get an existing AccountPrivilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountPrivilegeState, opts?: pulumi.CustomResourceOptions): AccountPrivilege {
        return new AccountPrivilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:polardb/accountPrivilege:AccountPrivilege';

    /**
     * Returns true if the given object is an instance of AccountPrivilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountPrivilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountPrivilege.__pulumiType;
    }

    /**
     * A specified account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
     */
    public readonly accountPrivilege!: pulumi.Output<string | undefined>;
    /**
     * The Id of cluster in which account belongs.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * List of specified database name.
     */
    public readonly dbNames!: pulumi.Output<string[]>;

    /**
     * Create a AccountPrivilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountPrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountPrivilegeArgs | AccountPrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountPrivilegeState | undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["accountPrivilege"] = state ? state.accountPrivilege : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["dbNames"] = state ? state.dbNames : undefined;
        } else {
            const args = argsOrState as AccountPrivilegeArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.dbClusterId === undefined) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            if (!args || args.dbNames === undefined) {
                throw new Error("Missing required property 'dbNames'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["accountPrivilege"] = args ? args.accountPrivilege : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["dbNames"] = args ? args.dbNames : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccountPrivilege.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountPrivilege resources.
 */
export interface AccountPrivilegeState {
    /**
     * A specified account name.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
     */
    readonly accountPrivilege?: pulumi.Input<string>;
    /**
     * The Id of cluster in which account belongs.
     */
    readonly dbClusterId?: pulumi.Input<string>;
    /**
     * List of specified database name.
     */
    readonly dbNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AccountPrivilege resource.
 */
export interface AccountPrivilegeArgs {
    /**
     * A specified account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
     */
    readonly accountPrivilege?: pulumi.Input<string>;
    /**
     * The Id of cluster in which account belongs.
     */
    readonly dbClusterId: pulumi.Input<string>;
    /**
     * List of specified database name.
     */
    readonly dbNames: pulumi.Input<pulumi.Input<string>[]>;
}
