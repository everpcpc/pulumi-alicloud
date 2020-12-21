// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides Private Zone Records resource information owned by an Alibaba Cloud account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const recordsDs = pulumi.all([alicloud_pvtz_zone_record_foo.value, alicloud_pvtz_zone_basic.id]).apply(([value, id]) => alicloud.pvtz.getZoneRecords({
 *     keyword: value,
 *     zoneId: id,
 * }, { async: true }));
 *
 * export const firstRecordId = recordsDs.records[0].id;
 * ```
 */
export function getZoneRecords(args: GetZoneRecordsArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneRecordsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:pvtz/getZoneRecords:getZoneRecords", {
        "ids": args.ids,
        "keyword": args.keyword,
        "lang": args.lang,
        "outputFile": args.outputFile,
        "searchMode": args.searchMode,
        "status": args.status,
        "tag": args.tag,
        "userClientIp": args.userClientIp,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getZoneRecords.
 */
export interface GetZoneRecordsArgs {
    /**
     * A list of Private Zone Record IDs.
     */
    readonly ids?: string[];
    /**
     * Keyword for record rr and value.
     */
    readonly keyword?: string;
    /**
     * User language.
     */
    readonly lang?: string;
    readonly outputFile?: string;
    /**
     * Search mode. Value: 
     * - LIKE: fuzzy search.
     * - EXACT: precise search. It is not filled in by default.
     */
    readonly searchMode?: string;
    /**
     * Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     */
    readonly status?: string;
    /**
     * It is not filled in by default, and queries the current zone resolution records. Fill in "ecs" to query the host name record list under the VPC associated with the current zone.
     */
    readonly tag?: string;
    /**
     * User ip.
     */
    readonly userClientIp?: string;
    /**
     * ID of the Private Zone.
     */
    readonly zoneId: string;
}

/**
 * A collection of values returned by getZoneRecords.
 */
export interface GetZoneRecordsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Private Zone Record IDs.
     */
    readonly ids: string[];
    readonly keyword?: string;
    readonly lang?: string;
    readonly outputFile?: string;
    /**
     * A list of zone records. Each element contains the following attributes:
     */
    readonly records: outputs.pvtz.GetZoneRecordsRecord[];
    readonly searchMode?: string;
    /**
     * Status of the Private Zone Record.
     */
    readonly status?: string;
    readonly tag?: string;
    readonly userClientIp?: string;
    readonly zoneId: string;
}
