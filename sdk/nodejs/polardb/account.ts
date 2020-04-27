// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB account resource and used to manage databases.
 * 
 * > **NOTE:** Available in v1.67.0+. 
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
 * const creation = config.get("creation") || "PolarDB";
 * const name = config.get("name") || "polardbaccountmysql";
 * 
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: creation,
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const cluster = new alicloud.polardb.Cluster("cluster", {
 *     dbNodeClass: "polar.mysql.x4.large",
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     description: name,
 *     payType: "PostPaid",
 *     vswitchId: defaultSwitch.id,
 * });
 * const account = new alicloud.rds.Account("account", {
 *     accountDescription: name,
 *     accountName: "tftestnormal",
 *     accountPassword: "Test12345",
 *     dbClusterId: alicloud_db_instance_cluster.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account.html.markdown.
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
    public static readonly __pulumiType = 'alicloud:polardb/account:Account';

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
     * Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    public readonly accountDescription!: pulumi.Output<string | undefined>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
     */
    public readonly accountPassword!: pulumi.Output<string>;
    /**
     * Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
     */
    public readonly accountType!: pulumi.Output<string | undefined>;
    /**
     * The Id of cluster in which account belongs.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;

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
            inputs["accountDescription"] = state ? state.accountDescription : undefined;
            inputs["accountName"] = state ? state.accountName : undefined;
            inputs["accountPassword"] = state ? state.accountPassword : undefined;
            inputs["accountType"] = state ? state.accountType : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.accountPassword === undefined) {
                throw new Error("Missing required property 'accountPassword'");
            }
            if (!args || args.dbClusterId === undefined) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            inputs["accountDescription"] = args ? args.accountDescription : undefined;
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["accountPassword"] = args ? args.accountPassword : undefined;
            inputs["accountType"] = args ? args.accountType : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            inputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
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
     * Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly accountDescription?: pulumi.Input<string>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     */
    readonly accountName?: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
     */
    readonly accountPassword?: pulumi.Input<string>;
    /**
     * Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
     */
    readonly accountType?: pulumi.Input<string>;
    /**
     * The Id of cluster in which account belongs.
     */
    readonly dbClusterId?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly accountDescription?: pulumi.Input<string>;
    /**
     * Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
     */
    readonly accountPassword: pulumi.Input<string>;
    /**
     * Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
     */
    readonly accountType?: pulumi.Input<string>;
    /**
     * The Id of cluster in which account belongs.
     */
    readonly dbClusterId: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
     */
    readonly kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    readonly kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
}
