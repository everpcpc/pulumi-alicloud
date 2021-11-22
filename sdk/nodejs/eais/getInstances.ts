// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Eais Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.137.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.eais.getInstances({
 *     id: ["example_id"],
 * });
 * export const eaisInstanceId1 = ids.then(ids => ids.instances?[0]?.id);
 * const nameRegex = alicloud.eais.getInstances({
 *     nameRegex: "^my-Instance",
 * });
 * export const eaisInstanceId2 = nameRegex.then(nameRegex => nameRegex.instances?[0]?.id);
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:eais/getInstances:getInstances", {
        "ids": args.ids,
        "instanceType": args.instanceType,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * A list of Instance IDs.
     */
    ids?: string[];
    /**
     * The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
     */
    instanceType?: string;
    /**
     * A regex string to filter results by Instance name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
     */
    status?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceType?: string;
    readonly instances: outputs.eais.GetInstancesInstance[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}

export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * A list of Instance IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the resource. Valid values: `eais.ei-a6.4xlarge`, `eais.ei-a6.2xlarge`, `eais.ei-a6.xlarge`, `eais.ei-a6.large`, `eais.ei-a6.medium`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Instance name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Attaching`, `Available`, `Detaching`, `InUse`, `Starting`, `Unavailable`.
     */
    status?: pulumi.Input<string>;
}
