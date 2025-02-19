// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ecp Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.158.0+.
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ecp/getInstances:getInstances", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "imageId": args.imageId,
        "instanceName": args.instanceName,
        "instanceType": args.instanceType,
        "keyPairName": args.keyPairName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "paymentType": args.paymentType,
        "resolution": args.resolution,
        "status": args.status,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    enableDetails?: boolean;
    /**
     * A list of Ecp Instances IDs.
     */
    ids?: string[];
    /**
     * The ID Of The Image.
     */
    imageId?: string;
    /**
     * Instance Name.
     */
    instanceName?: string;
    /**
     * Instance Type.
     */
    instanceType?: string;
    /**
     * The Key Name.
     */
    keyPairName?: string;
    /**
     * A regex string to filter results by mobile phone name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     */
    paymentType?: string;
    /**
     * Resolution.
     */
    resolution?: string;
    /**
     * Instance Status.
     */
    status?: string;
    zoneId?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly imageId?: string;
    readonly instanceName?: string;
    readonly instanceType?: string;
    readonly instances: outputs.ecp.GetInstancesInstance[];
    readonly keyPairName?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly paymentType?: string;
    readonly resolution?: string;
    readonly status?: string;
    readonly zoneId?: string;
}
/**
 * This data source provides the Ecp Instances of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.158.0+.
 */
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Ecp Instances IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID Of The Image.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Instance Name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance Type.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The Key Name.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * A regex string to filter results by mobile phone name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Resolution.
     */
    resolution?: pulumi.Input<string>;
    /**
     * Instance Status.
     */
    status?: pulumi.Input<string>;
    zoneId?: pulumi.Input<string>;
}
