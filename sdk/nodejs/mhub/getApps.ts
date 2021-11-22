// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Mhub Apps of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.138.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "example_value";
 * const _default = new alicloud.mhub.App("default", {
 *     appName: name,
 *     productId: alicloud_mhub_product["default"].id,
 *     packageName: "com.test.android",
 *     type: "2",
 * });
 * const ids = alicloud.mhub.getApps({});
 * export const mhubAppId1 = ids.then(ids => ids.apps?[0]?.id);
 * const nameRegex = alicloud.mhub.getApps({
 *     nameRegex: "^my-App",
 * });
 * export const mhubAppId2 = nameRegex.then(nameRegex => nameRegex.apps?[0]?.id);
 * ```
 */
export function getApps(args: GetAppsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:mhub/getApps:getApps", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "osType": args.osType,
        "outputFile": args.outputFile,
        "productId": args.productId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApps.
 */
export interface GetAppsArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of App IDs. The value formats as `<product_id>:<app_key>`
     */
    ids?: string[];
    /**
     * A regex string to filter results by App name.
     */
    nameRegex?: string;
    /**
     * The os type. Valid values: `Android` and `iOS`.
     */
    osType?: string;
    outputFile?: string;
    /**
     * The ID of the Product.
     */
    productId: string;
}

/**
 * A collection of values returned by getApps.
 */
export interface GetAppsResult {
    readonly apps: outputs.mhub.GetAppsApp[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly osType?: string;
    readonly outputFile?: string;
    readonly productId: string;
}

export function getAppsOutput(args: GetAppsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppsResult> {
    return pulumi.output(args).apply(a => getApps(a, opts))
}

/**
 * A collection of arguments for invoking getApps.
 */
export interface GetAppsOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of App IDs. The value formats as `<product_id>:<app_key>`
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by App name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The os type. Valid values: `Android` and `iOS`.
     */
    osType?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the Product.
     */
    productId: pulumi.Input<string>;
}
