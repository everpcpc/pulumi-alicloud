// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an RDS database resource. A DB database deployed in a DB instance. A DB instance can own multiple databases.
 * 
 * > **NOTE:** This resource does not support creating 'PPAS' database. You have to login RDS instance to create manually.
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
 * const name = config.get("name") || "dbdatabasebasic";
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
 * const instance = new alicloud.rds.Instance("instance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceName: name,
 *     instanceStorage: 10,
 *     instanceType: "rds.mysql.s1.small",
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultDatabase = new alicloud.rds.Database("default", {
 *     instanceId: instance.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_database.html.markdown.
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Character set. The value range is limited to the following:
     * - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
     * - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
     * - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
     */
    public readonly characterSet!: pulumi.Output<string | undefined>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Id of instance that can run database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
     * and have no more than 64 characters.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["characterSet"] = state ? state.characterSet : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["characterSet"] = args ? args.characterSet : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * Character set. The value range is limited to the following:
     * - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
     * - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
     * - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
     */
    readonly characterSet?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
     * and have no more than 64 characters.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Character set. The value range is limited to the following:
     * - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
     * - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
     * - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
     */
    readonly characterSet?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Id of instance that can run database.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
     * and have no more than 64 characters.
     */
    readonly name?: pulumi.Input<string>;
}
