// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB database resource. A DB database deployed in a DB cluster. A DB cluster can own multiple databases.
 * 
 * > **NOTE:** Available in v1.66.0+.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const cluster = new alicloud.polardb.Cluster("cluster", {
 *     dbNodeClass: var_clusterclass,
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     description: "testDB",
 *     payType: "PostPaid",
 *     vswitchId: "polar.mysql.x4.large",
 * });
 * const defaultDatabase = new alicloud.polardb.Database("default", {
 *     dbClusterId: cluster.id,
 *     dbName: "tftestdatabase",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_database.html.markdown.
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
    public static readonly __pulumiType = 'alicloud:polardb/database:Database';

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
     * Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is "utf8" \(`utf8mb4` only supports versions 5.5 and 5.6\).
     */
    public readonly characterSetName!: pulumi.Output<string | undefined>;
    /**
     * The Id of cluster that can run database.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    public readonly dbDescription!: pulumi.Output<string | undefined>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
     */
    public readonly dbName!: pulumi.Output<string>;

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
            inputs["characterSetName"] = state ? state.characterSetName : undefined;
            inputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            inputs["dbDescription"] = state ? state.dbDescription : undefined;
            inputs["dbName"] = state ? state.dbName : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.dbClusterId === undefined) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            if (!args || args.dbName === undefined) {
                throw new Error("Missing required property 'dbName'");
            }
            inputs["characterSetName"] = args ? args.characterSetName : undefined;
            inputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            inputs["dbDescription"] = args ? args.dbDescription : undefined;
            inputs["dbName"] = args ? args.dbName : undefined;
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
     * Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is "utf8" \(`utf8mb4` only supports versions 5.5 and 5.6\).
     */
    readonly characterSetName?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId?: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly dbDescription?: pulumi.Input<string>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
     */
    readonly dbName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is "utf8" \(`utf8mb4` only supports versions 5.5 and 5.6\).
     */
    readonly characterSetName?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    readonly dbClusterId: pulumi.Input<string>;
    /**
     * Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
     */
    readonly dbDescription?: pulumi.Input<string>;
    /**
     * Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.
     */
    readonly dbName: pulumi.Input<string>;
}
