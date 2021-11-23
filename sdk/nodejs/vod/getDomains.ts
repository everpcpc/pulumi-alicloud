// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Vod Domains of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultDomain = new alicloud.vod.Domain("defaultDomain", {
 *     domainName: "your_domain_name",
 *     scope: "domestic",
 *     sources: [{
 *         sourceType: "domain",
 *         sourceContent: "your_source_content",
 *         sourcePort: "80",
 *     }],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 * const defaultDomains = defaultDomain.id.apply(id => alicloud.vod.getDomains({
 *     ids: [id],
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * }));
 * export const vodDomain = defaultDomains.apply(defaultDomains => defaultDomains.domains?[0]);
 * ```
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vod/getDomains:getDomains", {
        "domainSearchType": args.domainSearchType,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * The search method. Valid values:
     * * `fuzzyMatch`: fuzzy match. This is the default value.
     * * `preMatch`: prefix match.
     * * `sufMatch`: suffix match.
     * * `fullMatch`: exact match
     */
    domainSearchType?: string;
    /**
     * A list of Domain IDs. Its element value is same as Domain Name.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Domain name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of the resource.
     */
    status?: string;
    /**
     * A mapping of tags to assign to the resource.
     * * `Key`: It can be up to 64 characters in length. It cannot be a null string.
     * * `Value`: It can be up to 128 characters in length. It can be a null string.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly domainSearchType?: string;
    readonly domains: outputs.vod.GetDomainsDomain[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
    readonly tags?: {[key: string]: any};
}

export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply(a => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * The search method. Valid values:
     * * `fuzzyMatch`: fuzzy match. This is the default value.
     * * `preMatch`: prefix match.
     * * `sufMatch`: suffix match.
     * * `fullMatch`: exact match
     */
    domainSearchType?: pulumi.Input<string>;
    /**
     * A list of Domain IDs. Its element value is same as Domain Name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Domain name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * * `Key`: It can be up to 64 characters in length. It cannot be a null string.
     * * `Value`: It can be up to 128 characters in length. It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}