// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a collection of DCDN Domains to the specified filters.
 *
 * > **NOTE:** Available in 1.94.0+.
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:dcdn/getDomains:getDomains", {
        "changeEndTime": args.changeEndTime,
        "changeStartTime": args.changeStartTime,
        "checkDomainShow": args.checkDomainShow,
        "domainSearchType": args.domainSearchType,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "securityToken": args.securityToken,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * The end time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
     */
    changeEndTime?: string;
    /**
     * The start time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
     */
    changeStartTime?: string;
    /**
     * Specifies whether to display the domains in the checking, check_failed, or configureFailed status. Valid values: `true` or `false`.
     */
    checkDomainShow?: boolean;
    /**
     * The search method. Default value: `fuzzyMatch`. Valid values: `fuzzyMatch`, `preMatch`, `sufMatch`, `fullMatch`.
     */
    domainSearchType?: string;
    /**
     * Default to `false`. Set it to true can output more details.
     */
    enableDetails?: boolean;
    /**
     * A list ids of DCDN Domain.
     */
    ids?: string[];
    /**
     * A regex string to filter results by the DCDN Domain.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: string;
    securityToken?: string;
    /**
     * The status of DCDN Domain.
     */
    status?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly changeEndTime?: string;
    readonly changeStartTime?: string;
    readonly checkDomainShow?: boolean;
    readonly domainSearchType?: string;
    /**
     * A list of domains. Each element contains the following attributes:
     */
    readonly domains: outputs.dcdn.GetDomainsDomain[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list ids of DCDN Domain.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of DCDN Domain names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The ID of the resource group.
     */
    readonly resourceGroupId?: string;
    readonly securityToken?: string;
    /**
     * The status of DCDN Domain. Valid values: `online`, `offline`, `checkFailed`, `checking`, `configureFailed`, `configuring`.
     */
    readonly status?: string;
}

export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply(a => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * The end time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
     */
    changeEndTime?: pulumi.Input<string>;
    /**
     * The start time of the update. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
     */
    changeStartTime?: pulumi.Input<string>;
    /**
     * Specifies whether to display the domains in the checking, check_failed, or configureFailed status. Valid values: `true` or `false`.
     */
    checkDomainShow?: pulumi.Input<boolean>;
    /**
     * The search method. Default value: `fuzzyMatch`. Valid values: `fuzzyMatch`, `preMatch`, `sufMatch`, `fullMatch`.
     */
    domainSearchType?: pulumi.Input<string>;
    /**
     * Default to `false`. Set it to true can output more details.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list ids of DCDN Domain.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by the DCDN Domain.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    securityToken?: pulumi.Input<string>;
    /**
     * The status of DCDN Domain.
     */
    status?: pulumi.Input<string>;
}
