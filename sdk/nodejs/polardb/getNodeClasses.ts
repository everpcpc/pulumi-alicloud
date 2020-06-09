// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the PolarDB node classes resource available info of Alibaba Cloud.
 *
 * > **NOTE:** Available in v1.81.0+
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const resourcesZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: "PolarDB",
 * }, { async: true }));
 * const resourcesNodeClasses = resourcesZones.apply(resourcesZones => alicloud.polardb.getNodeClasses({
 *     dbType: "MySQL",
 *     dbVersion: "5.6",
 *     payType: "Postpaid",
 *     zoneId: resourcesZones.zones[0].id,
 * }, { async: true }));
 *
 * export const firstPolardbNodeClass = resourcesNodeClasses.classes;
 * ```
 */
export function getNodeClasses(args: GetNodeClassesArgs, opts?: pulumi.InvokeOptions): Promise<GetNodeClassesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:polardb/getNodeClasses:getNodeClasses", {
        "dbNodeClass": args.dbNodeClass,
        "dbType": args.dbType,
        "dbVersion": args.dbVersion,
        "outputFile": args.outputFile,
        "payType": args.payType,
        "regionId": args.regionId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodeClasses.
 */
export interface GetNodeClassesArgs {
    /**
     * The PolarDB node class type by the user.
     */
    readonly dbNodeClass?: string;
    /**
     * Database type. Options are `MySQL`, `PostgreSQL`, `Oracle`. If dbType is set, dbVersion also needs to be set.
     */
    readonly dbType?: string;
    /**
     * Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) `DBVersion`. If dbVersion is set, dbType also needs to be set.
     */
    readonly dbVersion?: string;
    readonly outputFile?: string;
    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`.
     */
    readonly payType: string;
    /**
     * The Region to launch the PolarDB cluster.
     */
    readonly regionId?: string;
    /**
     * The Zone to launch the PolarDB cluster.
     */
    readonly zoneId?: string;
}

/**
 * A collection of values returned by getNodeClasses.
 */
export interface GetNodeClassesResult {
    /**
     * A list of PolarDB node classes. Each element contains the following attributes:
     */
    readonly classes: outputs.polardb.GetNodeClassesClass[];
    /**
     * PolarDB node available class.
     */
    readonly dbNodeClass?: string;
    readonly dbType?: string;
    readonly dbVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    readonly payType: string;
    readonly regionId?: string;
    /**
     * The Zone to launch the PolarDB cluster.
     */
    readonly zoneId?: string;
}
