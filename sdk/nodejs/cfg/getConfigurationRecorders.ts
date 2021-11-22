// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Config Configuration Recorders of the current Alibaba Cloud user.
 *
 * > **NOTE:**  Available in 1.99.0+.
 *
 * > **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-northeast-1`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.cfg.getConfigurationRecorders({});
 * export const listOfResourceTypes = data.alicloud_config_configuration_recorders["this"].recorders[0].resource_types;
 * ```
 */
export function getConfigurationRecorders(args?: GetConfigurationRecordersArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationRecordersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cfg/getConfigurationRecorders:getConfigurationRecorders", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getConfigurationRecorders.
 */
export interface GetConfigurationRecordersArgs {
    outputFile?: string;
}

/**
 * A collection of values returned by getConfigurationRecorders.
 */
export interface GetConfigurationRecordersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * A list of Config Configuration Recorders. Each element contains the following attributes:
     */
    readonly recorders: outputs.cfg.GetConfigurationRecordersRecorder[];
}

export function getConfigurationRecordersOutput(args?: GetConfigurationRecordersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationRecordersResult> {
    return pulumi.output(args).apply(a => getConfigurationRecorders(a, opts))
}

/**
 * A collection of arguments for invoking getConfigurationRecorders.
 */
export interface GetConfigurationRecordersOutputArgs {
    outputFile?: pulumi.Input<string>;
}
