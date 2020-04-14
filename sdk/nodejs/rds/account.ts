// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an RDS account resource and used to manage databases.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "Rds";
 * const name = config.get("name") || "dbaccountmysql";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const instance = new alicloud.rds.Instance("instance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceName: name,
 *     instanceStorage: 10,
 *     instanceType: "rds.mysql.s1.small",
 *     vswitchId: defaultSwitch.id,
 * });
 * const account = new alicloud.rds.Account("account", {
 *     instanceId: instance.id,
 *     password: "Test12345",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_account.html.markdown.
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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

    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Id of instance in which account belongs.
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
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Account.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Id of instance in which account belongs.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Privilege type of account.
     * - Normal: Common privilege.
     * - Super: High privilege.
     */
    readonly type?: pulumi.Input<string>;
}
