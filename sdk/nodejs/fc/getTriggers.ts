// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Function Compute triggers of the current Alibaba Cloud user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const fcTriggersDs = alicloud.fc.getTriggers({
 *     functionName: "sampleFunction",
 *     nameRegex: "sampleFcTrigger",
 *     serviceName: "sampleService",
 * });
 * 
 * export const firstFcTriggerName = fcTriggersDs.triggers[0].name;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/fc_triggers.html.markdown.
 */
export function getTriggers(args: GetTriggersArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:fc/getTriggers:getTriggers", {
        "functionName": args.functionName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTriggers.
 */
export interface GetTriggersArgs {
    /**
     * FC function name.
     */
    readonly functionName: string;
    /**
     * A list of FC triggers ids.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by FC trigger name.
     * * `ids` (Optional, Available in 1.53.0+) - A list of FC triggers ids.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * FC service name.
     */
    readonly serviceName: string;
}

/**
 * A collection of values returned by getTriggers.
 */
export interface GetTriggersResult {
    readonly functionName: string;
    /**
     * A list of FC triggers ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of FC triggers names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly serviceName: string;
    /**
     * A list of FC triggers. Each element contains the following attributes:
     */
    readonly triggers: outputs.fc.GetTriggersTrigger[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
