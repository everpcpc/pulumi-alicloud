// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ots_table.html.markdown.
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ots/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
     */
    public readonly deviationCellVersionInSec!: pulumi.Output<string | undefined>;
    /**
     * The name of the OTS instance in which table will located.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The maximum number of versions stored in this table. The valid value is 1-2147483647.
     */
    public readonly maxVersion!: pulumi.Output<number>;
    /**
     * The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
     */
    public readonly primaryKeys!: pulumi.Output<outputs.ots.TablePrimaryKey[]>;
    /**
     * The table name of the OTS instance. If changed, a new table would be created.
     */
    public readonly tableName!: pulumi.Output<string>;
    /**
     * The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
     */
    public readonly timeToLive!: pulumi.Output<number>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TableState | undefined;
            inputs["deviationCellVersionInSec"] = state ? state.deviationCellVersionInSec : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["maxVersion"] = state ? state.maxVersion : undefined;
            inputs["primaryKeys"] = state ? state.primaryKeys : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
            inputs["timeToLive"] = state ? state.timeToLive : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            if (!args || args.maxVersion === undefined) {
                throw new Error("Missing required property 'maxVersion'");
            }
            if (!args || args.primaryKeys === undefined) {
                throw new Error("Missing required property 'primaryKeys'");
            }
            if (!args || args.tableName === undefined) {
                throw new Error("Missing required property 'tableName'");
            }
            if (!args || args.timeToLive === undefined) {
                throw new Error("Missing required property 'timeToLive'");
            }
            inputs["deviationCellVersionInSec"] = args ? args.deviationCellVersionInSec : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["maxVersion"] = args ? args.maxVersion : undefined;
            inputs["primaryKeys"] = args ? args.primaryKeys : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
            inputs["timeToLive"] = args ? args.timeToLive : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Table.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
     */
    readonly deviationCellVersionInSec?: pulumi.Input<string>;
    /**
     * The name of the OTS instance in which table will located.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The maximum number of versions stored in this table. The valid value is 1-2147483647.
     */
    readonly maxVersion?: pulumi.Input<number>;
    /**
     * The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
     */
    readonly primaryKeys?: pulumi.Input<pulumi.Input<inputs.ots.TablePrimaryKey>[]>;
    /**
     * The table name of the OTS instance. If changed, a new table would be created.
     */
    readonly tableName?: pulumi.Input<string>;
    /**
     * The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
     */
    readonly timeToLive?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
     */
    readonly deviationCellVersionInSec?: pulumi.Input<string>;
    /**
     * The name of the OTS instance in which table will located.
     */
    readonly instanceName: pulumi.Input<string>;
    /**
     * The maximum number of versions stored in this table. The valid value is 1-2147483647.
     */
    readonly maxVersion: pulumi.Input<number>;
    /**
     * The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
     */
    readonly primaryKeys: pulumi.Input<pulumi.Input<inputs.ots.TablePrimaryKey>[]>;
    /**
     * The table name of the OTS instance. If changed, a new table would be created.
     */
    readonly tableName: pulumi.Input<string>;
    /**
     * The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
     */
    readonly timeToLive: pulumi.Input<number>;
}
