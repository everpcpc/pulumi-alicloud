// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ots search index of the current Alibaba Cloud user.
 *
 * For information about OTS search index and how to use it, see [Search index overview](https://www.alibabacloud.com/help/en/tablestore/latest/search-index-overview).
 *
 * > **NOTE:** Available in v1.187.0+.
 */
export function getSearchIndexes(args: GetSearchIndexesArgs, opts?: pulumi.InvokeOptions): Promise<GetSearchIndexesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ots/getSearchIndexes:getSearchIndexes", {
        "ids": args.ids,
        "instanceName": args.instanceName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tableName": args.tableName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSearchIndexes.
 */
export interface GetSearchIndexesArgs {
    /**
     * A list of search index IDs.
     */
    ids?: string[];
    /**
     * The name of OTS instance.
     */
    instanceName: string;
    /**
     * A regex string to filter results by search index name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The name of OTS table.
     */
    tableName: string;
}

/**
 * A collection of values returned by getSearchIndexes.
 */
export interface GetSearchIndexesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of search index IDs.
     */
    readonly ids: string[];
    /**
     * A list of indexes. Each element contains the following attributes:
     */
    readonly indexes: outputs.ots.GetSearchIndexesIndex[];
    /**
     * The OTS instance name.
     */
    readonly instanceName: string;
    readonly nameRegex?: string;
    /**
     * A list of search index  names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The table name of the OTS which could not be changed.
     */
    readonly tableName: string;
}
/**
 * This data source provides the ots search index of the current Alibaba Cloud user.
 *
 * For information about OTS search index and how to use it, see [Search index overview](https://www.alibabacloud.com/help/en/tablestore/latest/search-index-overview).
 *
 * > **NOTE:** Available in v1.187.0+.
 */
export function getSearchIndexesOutput(args: GetSearchIndexesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSearchIndexesResult> {
    return pulumi.output(args).apply((a: any) => getSearchIndexes(a, opts))
}

/**
 * A collection of arguments for invoking getSearchIndexes.
 */
export interface GetSearchIndexesOutputArgs {
    /**
     * A list of search index IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of OTS instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * A regex string to filter results by search index name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The name of OTS table.
     */
    tableName: pulumi.Input<string>;
}
