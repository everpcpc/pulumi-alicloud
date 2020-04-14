// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getMongoInstances(args?: GetMongoInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetMongoInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:dds/getMongoInstances:getMongoInstances", {
        "availabilityZone": args.availabilityZone,
        "ids": args.ids,
        "instanceClass": args.instanceClass,
        "instanceType": args.instanceType,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getMongoInstances.
 */
export interface GetMongoInstancesArgs {
    readonly availabilityZone?: string;
    readonly ids?: string[];
    readonly instanceClass?: string;
    readonly instanceType?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getMongoInstances.
 */
export interface GetMongoInstancesResult {
    readonly availabilityZone?: string;
    readonly ids: string[];
    readonly instanceClass?: string;
    readonly instanceType?: string;
    readonly instances: outputs.dds.GetMongoInstancesInstance[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
