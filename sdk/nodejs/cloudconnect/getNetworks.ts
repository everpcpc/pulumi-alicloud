// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides Cloud Connect Networks available to the user.
 * 
 * > **NOTE:** Available in 1.59.0+
 * 
 * > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultNetworks = alicloud_cloud_connect_networks_default.id.apply(id => alicloud.cloudconnect.getNetworks({
 *     ids: [id],
 *     nameRegex: "^tf-testAcc.*",
 * }));
 * const defaultNetwork = new alicloud.cloudconnect.Network("default", {
 *     cidrBlock: "192.168.0.0/24",
 *     description: "tf-testAccCloudConnectNetworkDescription",
 *     isDefault: true,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/cloud_connect_networks.html.markdown.
 */
export function getNetworks(args?: GetNetworksArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cloudconnect/getNetworks:getNetworks", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworks.
 */
export interface GetNetworksArgs {
    /**
     * A list of CCN instances IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter CCN instances by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getNetworks.
 */
export interface GetNetworksResult {
    /**
     * A list of CCN instances IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of CCN instances names. 
     */
    readonly names: string[];
    /**
     * A list of CCN instances. Each element contains the following attributes:
     */
    readonly networks: outputs.cloudconnect.GetNetworksNetwork[];
    readonly outputFile?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
