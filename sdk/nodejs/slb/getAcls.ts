// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the acls in the region.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const sampleDs = alicloud.slb.getAcls();
 * 
 * export const firstSlbAclId = sampleDs.acls[0].id;
 * ```
 * 
 * ## Entry Block
 * 
 * The entry mapping supports the following:
 * 
 * * `entry`   - An IP addresses or CIDR blocks.
 * * `comment` - the comment of the entry.
 * 
 * ## Listener Block
 * 
 * The Listener mapping supports the following:
 * 
 * * `loadBalancerId` - the id of load balancer instance, the listener belongs to.
 * * `frontendPort` - the listener port.
 * * `protocol`      - the listener protocol (such as tcp/udp/http/https, etc).
 * * `aclType`      - the type of acl (such as white/black).
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_acls.html.markdown.
 */
export function getAcls(args?: GetAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetAclsResult> & GetAclsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAclsResult> = pulumi.runtime.invoke("alicloud:slb/getAcls:getAcls", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsArgs {
    /**
     * A list of acls IDs to filter results.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by acl name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group which acl belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getAcls.
 */
export interface GetAclsResult {
    /**
     * A list of SLB  acls. Each element contains the following attributes:
     */
    readonly acls: outputs.slb.GetAclsAcl[];
    /**
     * A list of SLB acls IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of SLB acls names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * Resource group ID.
     */
    readonly resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
