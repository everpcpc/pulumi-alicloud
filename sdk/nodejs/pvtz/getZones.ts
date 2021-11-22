// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const pvtzZonesDs = alicloud_pvtz_zone_basic.zoneName.apply(zoneName => alicloud.pvtz.getZones({
 *     keyword: zoneName,
 * }));
 *
 * export const firstZoneId = pvtzZonesDs.zones[0].id;
 * ```
 */
export function getZones(args?: GetZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:pvtz/getZones:getZones", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "keyword": args.keyword,
        "lang": args.lang,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "queryRegionId": args.queryRegionId,
        "queryVpcId": args.queryVpcId,
        "resourceGroupId": args.resourceGroupId,
        "searchMode": args.searchMode,
    }, opts);
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesArgs {
    /**
     * -(Optional, Available 1.107.0+) Default to `false`. Set it to true can output more details.
     */
    enableDetails?: boolean;
    /**
     * A list of zone IDs.
     */
    ids?: string[];
    /**
     * keyword for zone name.
     */
    keyword?: string;
    /**
     * User language.
     */
    lang?: string;
    nameRegex?: string;
    outputFile?: string;
    /**
     * query_region_id for zone regionId.
     */
    queryRegionId?: string;
    /**
     * query_vpc_id for zone vpcId.
     */
    queryVpcId?: string;
    /**
     * resource_group_id for zone resourceGroupId.
     */
    resourceGroupId?: string;
    /**
     * Search mode. Value: 
     * - LIKE: fuzzy search.
     * - EXACT: precise search. It is not filled in by default.
     */
    searchMode?: string;
}

/**
 * A collection of values returned by getZones.
 */
export interface GetZonesResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of zone IDs.
     */
    readonly ids: string[];
    readonly keyword?: string;
    readonly lang?: string;
    readonly nameRegex?: string;
    /**
     * A list of zone names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly queryRegionId?: string;
    readonly queryVpcId?: string;
    /**
     * The Id of resource group which the Private Zone belongs.
     */
    readonly resourceGroupId?: string;
    readonly searchMode?: string;
    /**
     * A list of zones. Each element contains the following attributes:
     */
    readonly zones: outputs.pvtz.GetZonesZone[];
}

export function getZonesOutput(args?: GetZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZonesResult> {
    return pulumi.output(args).apply(a => getZones(a, opts))
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesOutputArgs {
    /**
     * -(Optional, Available 1.107.0+) Default to `false`. Set it to true can output more details.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of zone IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * keyword for zone name.
     */
    keyword?: pulumi.Input<string>;
    /**
     * User language.
     */
    lang?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * query_region_id for zone regionId.
     */
    queryRegionId?: pulumi.Input<string>;
    /**
     * query_vpc_id for zone vpcId.
     */
    queryVpcId?: pulumi.Input<string>;
    /**
     * resource_group_id for zone resourceGroupId.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Search mode. Value: 
     * - LIKE: fuzzy search.
     * - EXACT: precise search. It is not filled in by default.
     */
    searchMode?: pulumi.Input<string>;
}
