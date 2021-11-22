// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
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
    availabilityZone?: string;
    ids?: string[];
    instanceClass?: string;
    instanceType?: string;
    nameRegex?: string;
    outputFile?: string;
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getMongoInstances.
 */
export interface GetMongoInstancesResult {
    readonly availabilityZone?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceClass?: string;
    readonly instanceType?: string;
    readonly instances: outputs.dds.GetMongoInstancesInstance[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
}

export function getMongoInstancesOutput(args?: GetMongoInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMongoInstancesResult> {
    return pulumi.output(args).apply(a => getMongoInstances(a, opts))
}

/**
 * A collection of arguments for invoking getMongoInstances.
 */
export interface GetMongoInstancesOutputArgs {
    availabilityZone?: pulumi.Input<string>;
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    instanceClass?: pulumi.Input<string>;
    instanceType?: pulumi.Input<string>;
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}
