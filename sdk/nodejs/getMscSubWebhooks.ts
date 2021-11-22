// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the Msc Sub Webhooks of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.141.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.getMscSubWebhooks({
 *     ids: ["example_id"],
 * });
 * export const mscSubWebhookId1 = ids.then(ids => ids.webhooks?[0]?.id);
 * const nameRegex = alicloud.getMscSubWebhooks({
 *     nameRegex: "^my-Webhook",
 * });
 * export const mscSubWebhookId2 = nameRegex.then(nameRegex => nameRegex.webhooks?[0]?.id);
 * ```
 */
export function getMscSubWebhooks(args?: GetMscSubWebhooksArgs, opts?: pulumi.InvokeOptions): Promise<GetMscSubWebhooksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:index/getMscSubWebhooks:getMscSubWebhooks", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMscSubWebhooks.
 */
export interface GetMscSubWebhooksArgs {
    /**
     * A list of Webhook IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Webhook name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getMscSubWebhooks.
 */
export interface GetMscSubWebhooksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly webhooks: outputs.GetMscSubWebhooksWebhook[];
}

export function getMscSubWebhooksOutput(args?: GetMscSubWebhooksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMscSubWebhooksResult> {
    return pulumi.output(args).apply(a => getMscSubWebhooks(a, opts))
}

/**
 * A collection of arguments for invoking getMscSubWebhooks.
 */
export interface GetMscSubWebhooksOutputArgs {
    /**
     * A list of Webhook IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Webhook name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
