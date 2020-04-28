// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides CEN flow logs available to the user.
 * 
 * > **NOTE:** Available in 1.78.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultFlowlogs = pulumi.output(alicloud.cen.getFlowlogs({
 *     ids: ["flowlog-tig1xxxxx"],
 *     nameRegex: "^foo",
 * }, { async: true }));
 * 
 * export const firstCenFlowlogId = alicloud_cen_instances_default.flowlogs.0.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cen_flowlogs.html.markdown.
 */
export function getFlowlogs(args?: GetFlowlogsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowlogsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cen/getFlowlogs:getFlowlogs", {
        "cenId": args.cenId,
        "description": args.description,
        "ids": args.ids,
        "logStoreName": args.logStoreName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlowlogs.
 */
export interface GetFlowlogsArgs {
    /**
     * The ID of the CEN Instance.
     */
    readonly cenId?: string;
    /**
     * The description of flowlog.
     */
    readonly description?: string;
    /**
     * A list of CEN flow log IDs.
     */
    readonly ids?: string[];
    /**
     * The name of the log store which is in the  `projectName` SLS project.
     */
    readonly logStoreName?: string;
    /**
     * A regex string to filter CEN flow logs by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of the SLS project.
     */
    readonly projectName?: string;
    /**
     * The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getFlowlogs.
 */
export interface GetFlowlogsResult {
    /**
     * The ID of the CEN Instance.
     */
    readonly cenId?: string;
    /**
     * The description of flowlog.
     */
    readonly description?: string;
    readonly flowlogs: outputs.cen.GetFlowlogsFlowlog[];
    /**
     * A list of CEN flow log IDs.
     */
    readonly ids: string[];
    /**
     * The name of the log store which is in the  `projectName` SLS project.
     */
    readonly logStoreName?: string;
    readonly nameRegex?: string;
    /**
     * A list of CEN flow log names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The name of the SLS project.
     */
    readonly projectName?: string;
    /**
     * The status of flowlog.
     */
    readonly status?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}