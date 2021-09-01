// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Sddp Rules of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultRule = new alicloud.sddp.Rule("defaultRule", {
 *     category: "0",
 *     content: "content",
 *     ruleName: "rule_name",
 *     riskLevelId: "4",
 *     productCode: "ODPS",
 * });
 * const defaultRules = defaultRule.id.apply(id => alicloud.sddp.getRules({
 *     ids: [id],
 * }));
 * export const sddpRuleId = defaultRules.id;
 * ```
 */
export function getRules(args?: GetRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetRulesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:sddp/getRules:getRules", {
        "category": args.category,
        "contentCategory": args.contentCategory,
        "customType": args.customType,
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "productId": args.productId,
        "riskLevelId": args.riskLevelId,
        "ruleType": args.ruleType,
        "status": args.status,
        "warnLevel": args.warnLevel,
    }, opts);
}

/**
 * A collection of arguments for invoking getRules.
 */
export interface GetRulesArgs {
    /**
     * Sensitive Data Identification Rules for the Type of.
     */
    readonly category?: number;
    /**
     * The Content Classification.
     */
    readonly contentCategory?: string;
    /**
     * Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
     */
    readonly customType?: number;
    readonly enableDetails?: boolean;
    /**
     * A list of Rule IDs.
     */
    readonly ids?: string[];
    /**
     * The name of rule.
     */
    readonly name?: string;
    /**
     * A regex string to filter results by Rule name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Product ID.
     */
    readonly productId?: string;
    /**
     * Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
     */
    readonly riskLevelId?: string;
    /**
     * Rule Type.
     */
    readonly ruleType?: number;
    /**
     * Sensitive Data Identification Rules Detection State of.
     */
    readonly status?: string;
    /**
     * The Level of Risk.
     */
    readonly warnLevel?: number;
}

/**
 * A collection of values returned by getRules.
 */
export interface GetRulesResult {
    readonly category?: number;
    readonly contentCategory?: string;
    readonly customType?: number;
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly name?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly productId?: string;
    readonly riskLevelId?: string;
    readonly ruleType?: number;
    readonly rules: outputs.sddp.GetRulesRule[];
    readonly status?: string;
    readonly warnLevel?: number;
}