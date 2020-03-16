// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of EIPs (Elastic IP address) owned by an Alibaba Cloud account.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const eipsDs = alicloud.ecs.getEips();
 * 
 * export const firstEipId = eipsDs.eips[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/eips.html.markdown.
 */
export function getEips(args?: GetEipsArgs, opts?: pulumi.InvokeOptions): Promise<GetEipsResult> & GetEipsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetEipsResult> = pulumi.runtime.invoke("alicloud:ecs/getEips:getEips", {
        "ids": args.ids,
        "inUse": args.inUse,
        "ipAddresses": args.ipAddresses,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getEips.
 */
export interface GetEipsArgs {
    /**
     * A list of EIP IDs.
     */
    readonly ids?: string[];
    /**
     * Deprecated since the version 1.8.0 of this provider.
     * 
     * @deprecated Field 'in_use' has been deprecated from provider version 1.8.0.
     */
    readonly inUse?: boolean;
    /**
     * A list of EIP public IP addresses.
     */
    readonly ipAddresses?: string[];
    readonly outputFile?: string;
    /**
     * The Id of resource group which the eips belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getEips.
 */
export interface GetEipsResult {
    /**
     * A list of EIPs. Each element contains the following attributes:
     */
    readonly eips: outputs.ecs.GetEipsEip[];
    /**
     * (Optional) A list of EIP IDs.
     */
    readonly ids: string[];
    readonly inUse?: boolean;
    readonly ipAddresses?: string[];
    /**
     * (Optional) A list of EIP names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The Id of resource group which the eips belongs.
     */
    readonly resourceGroupId?: string;
    readonly tags?: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
