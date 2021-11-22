// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Cms Monitor Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.113.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.cms.getMonitorGroups({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstCmsMonitorGroupId = example.then(example => example.groups?[0]?.id);
 * ```
 */
export function getMonitorGroups(args?: GetMonitorGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cms/getMonitorGroups:getMonitorGroups", {
        "dynamicTagRuleId": args.dynamicTagRuleId,
        "ids": args.ids,
        "includeTemplateHistory": args.includeTemplateHistory,
        "keyword": args.keyword,
        "monitorGroupName": args.monitorGroupName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "selectContactGroups": args.selectContactGroups,
        "tags": args.tags,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitorGroups.
 */
export interface GetMonitorGroupsArgs {
    /**
     * The ID of the tag rule.
     */
    dynamicTagRuleId?: string;
    /**
     * A list of Monitor Group IDs.
     */
    ids?: string[];
    /**
     * The include template history.
     */
    includeTemplateHistory?: boolean;
    /**
     * The keyword to be matched.
     */
    keyword?: string;
    /**
     * The name of the application group.
     */
    monitorGroupName?: string;
    /**
     * A regex string to filter results by Monitor Group name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The select contact groups.
     */
    selectContactGroups?: boolean;
    /**
     * A map of tags assigned to the Cms Monitor Group.
     */
    tags?: {[key: string]: any};
    /**
     * The type of the application group.
     */
    type?: string;
}

/**
 * A collection of values returned by getMonitorGroups.
 */
export interface GetMonitorGroupsResult {
    readonly dynamicTagRuleId?: string;
    readonly groups: outputs.cms.GetMonitorGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly includeTemplateHistory?: boolean;
    readonly keyword?: string;
    readonly monitorGroupName?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly selectContactGroups?: boolean;
    readonly tags?: {[key: string]: any};
    readonly type?: string;
}

export function getMonitorGroupsOutput(args?: GetMonitorGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitorGroupsResult> {
    return pulumi.output(args).apply(a => getMonitorGroups(a, opts))
}

/**
 * A collection of arguments for invoking getMonitorGroups.
 */
export interface GetMonitorGroupsOutputArgs {
    /**
     * The ID of the tag rule.
     */
    dynamicTagRuleId?: pulumi.Input<string>;
    /**
     * A list of Monitor Group IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The include template history.
     */
    includeTemplateHistory?: pulumi.Input<boolean>;
    /**
     * The keyword to be matched.
     */
    keyword?: pulumi.Input<string>;
    /**
     * The name of the application group.
     */
    monitorGroupName?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Monitor Group name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The select contact groups.
     */
    selectContactGroups?: pulumi.Input<boolean>;
    /**
     * A map of tags assigned to the Cms Monitor Group.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of the application group.
     */
    type?: pulumi.Input<string>;
}
