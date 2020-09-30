// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ots instance attachments of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const attachmentsDs = pulumi.output(alicloud.ots.getInstanceAttachments({
 *     instanceName: "sample-instance",
 *     nameRegex: "testvpc",
 *     outputFile: "attachments.txt",
 * }, { async: true }));
 *
 * export const firstOtsAttachmentId = attachmentsDs.attachments[0].id;
 * ```
 */
export function getInstanceAttachments(args: GetInstanceAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceAttachmentsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ots/getInstanceAttachments:getInstanceAttachments", {
        "instanceName": args.instanceName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceAttachments.
 */
export interface GetInstanceAttachmentsArgs {
    /**
     * The name of OTS instance.
     */
    readonly instanceName: string;
    /**
     * A regex string to filter results by vpc name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getInstanceAttachments.
 */
export interface GetInstanceAttachmentsResult {
    /**
     * A list of instance attachments. Each element contains the following attributes:
     */
    readonly attachments: outputs.ots.GetInstanceAttachmentsAttachment[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The instance name.
     */
    readonly instanceName: string;
    readonly nameRegex?: string;
    /**
     * A list of vpc names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of vpc ids.
     */
    readonly vpcIds: string[];
}