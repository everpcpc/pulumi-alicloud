// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the details of the Kubernetes version supported by ACK.
 *
 * > **NOTE:** Available in 1.169.0+.
 */
export function getKubernetesVersion(args: GetKubernetesVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:cs/getKubernetesVersion:getKubernetesVersion", {
        "clusterType": args.clusterType,
        "kubernetesVersion": args.kubernetesVersion,
        "profile": args.profile,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesVersion.
 */
export interface GetKubernetesVersionArgs {
    /**
     * The type of cluster. Its valid value are `Kubernetes` and `ManagedKubernetes`.
     */
    clusterType: string;
    kubernetesVersion?: string;
    /**
     * The profile of cluster. Its valid value are `Default`, `Serverless` and `Edge`.
     */
    profile?: string;
}

/**
 * A collection of values returned by getKubernetesVersion.
 */
export interface GetKubernetesVersionResult {
    readonly clusterType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kubernetesVersion?: string;
    /**
     * A list of metadata of kubernetes version.
     */
    readonly metadatas: outputs.cs.GetKubernetesVersionMetadata[];
    readonly profile?: string;
}

export function getKubernetesVersionOutput(args: GetKubernetesVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubernetesVersionResult> {
    return pulumi.output(args).apply(a => getKubernetesVersion(a, opts))
}

/**
 * A collection of arguments for invoking getKubernetesVersion.
 */
export interface GetKubernetesVersionOutputArgs {
    /**
     * The type of cluster. Its valid value are `Kubernetes` and `ManagedKubernetes`.
     */
    clusterType: pulumi.Input<string>;
    kubernetesVersion?: pulumi.Input<string>;
    /**
     * The profile of cluster. Its valid value are `Default`, `Serverless` and `Edge`.
     */
    profile?: pulumi.Input<string>;
}