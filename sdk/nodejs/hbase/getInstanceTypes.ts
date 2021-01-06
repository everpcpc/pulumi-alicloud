// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides availability instanceTypes for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
 *
 * > **NOTE:** Available in v1.106.0+.
 */
export function getInstanceTypes(args?: GetInstanceTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:hbase/getInstanceTypes:getInstanceTypes", {
        "instanceType": args.instanceType,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypes.
 */
export interface GetInstanceTypesArgs {
    /**
     * The hbase instance type of create hbase cluster instance.
     */
    readonly instanceType?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getInstanceTypes.
 */
export interface GetInstanceTypesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instance types type IDs.
     */
    readonly ids: string[];
    readonly instanceType?: string;
    readonly outputFile?: string;
    /**
     * A list of instance types. Each element contains the following attributes:
     */
    readonly types: outputs.hbase.GetInstanceTypesType[];
}