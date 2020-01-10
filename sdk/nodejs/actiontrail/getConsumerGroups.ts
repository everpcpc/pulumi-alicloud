// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ALIKAFKA Consumer Groups in an Alibaba Cloud account according to the specified filters.
 * 
 * > **NOTE:** Available in 1.56.0+
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const consumerGroupsDs = alicloud.actiontrail.getConsumerGroups({
 *     consumerIdRegex: "CID-alikafkaGroupDatasourceName",
 *     instanceId: "xxx",
 *     outputFile: "consumerGroups.txt",
 * });
 * 
 * export const firstGroupName = consumerGroupsDs.consumerIds[0];
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/alikafka_consumer_groups.html.markdown.
 */
export function getConsumerGroups(args: GetConsumerGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetConsumerGroupsResult> & GetConsumerGroupsResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetConsumerGroupsResult> = pulumi.runtime.invoke("alicloud:actiontrail/getConsumerGroups:getConsumerGroups", {
        "consumerIdRegex": args.consumerIdRegex,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getConsumerGroups.
 */
export interface GetConsumerGroupsArgs {
    /**
     * A regex string to filter results by the consumer group id. 
     */
    readonly consumerIdRegex?: string;
    /**
     * ID of the ALIKAFKA Instance that owns the consumer groups.
     */
    readonly instanceId: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getConsumerGroups.
 */
export interface GetConsumerGroupsResult {
    readonly consumerIdRegex?: string;
    /**
     * A list of consumer group ids.
     */
    readonly consumerIds: string[];
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
