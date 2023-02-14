// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "Rds";
 * const name = config.get("name") || "dbaccountmysql";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 * });
 * const instance = new alicloud.rds.Instance("instance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceType: "rds.mysql.s1.small",
 *     instanceStorage: 10,
 *     vswitchId: defaultSwitch.id,
 *     instanceName: name,
 * });
 * const account = new alicloud.rds.Account("account", {
 *     instanceId: instance.id,
 *     password: "Test12345",
 * });
 * ```
 *
 * ## Import
 *
 * RDS account can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:rds/account:Account example "rm-12345:tf_account"
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    public readonly accountDescription!: pulumi.Output<string>;
    public readonly accountName!: pulumi.Output<string>;
    public readonly accountPassword!: pulumi.Output<string>;
    public readonly accountType!: pulumi.Output<string>;
    public readonly dbInstanceId!: pulumi.Output<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     *
     * @deprecated Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The Id of instance in which account belongs.
     *
     * @deprecated Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     *
     * @deprecated Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
     */
    public readonly password!: pulumi.Output<string>;
    public readonly resetPermissionFlag!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            resourceInputs["accountDescription"] = state ? state.accountDescription : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["accountType"] = state ? state.accountType : undefined;
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["resetPermissionFlag"] = state ? state.resetPermissionFlag : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            resourceInputs["accountDescription"] = args ? args.accountDescription : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["accountType"] = args ? args.accountType : undefined;
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["resetPermissionFlag"] = args ? args.resetPermissionFlag : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword", "password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    accountDescription?: pulumi.Input<string>;
    accountName?: pulumi.Input<string>;
    accountPassword?: pulumi.Input<string>;
    accountType?: pulumi.Input<string>;
    dbInstanceId?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     *
     * @deprecated Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
     */
    description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs.
     *
     * @deprecated Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     *
     * @deprecated Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
     */
    password?: pulumi.Input<string>;
    resetPermissionFlag?: pulumi.Input<boolean>;
    status?: pulumi.Input<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    accountDescription?: pulumi.Input<string>;
    accountName?: pulumi.Input<string>;
    accountPassword?: pulumi.Input<string>;
    accountType?: pulumi.Input<string>;
    dbInstanceId?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     *
     * @deprecated Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
     */
    description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs.
     *
     * @deprecated Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     *
     * @deprecated Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
     */
    password?: pulumi.Input<string>;
    resetPermissionFlag?: pulumi.Input<boolean>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     *
     * @deprecated Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
     */
    type?: pulumi.Input<string>;
}
