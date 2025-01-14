// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cloud Firewall Address Books of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.178.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cloudfirewall.getAddressBooks({});
 * export const cloudFirewallAddressBookId1 = ids.then(ids => ids.books?.[0]?.id);
 * ```
 */
export function getAddressBooks(args?: GetAddressBooksArgs, opts?: pulumi.InvokeOptions): Promise<GetAddressBooksResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cloudfirewall/getAddressBooks:getAddressBooks", {
        "groupType": args.groupType,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddressBooks.
 */
export interface GetAddressBooksArgs {
    /**
     * The type of the Address Book.
     */
    groupType?: string;
    /**
     * A list of Address Book IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results Address Book name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getAddressBooks.
 */
export interface GetAddressBooksResult {
    readonly books: outputs.cloudfirewall.GetAddressBooksBook[];
    readonly groupType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Cloud Firewall Address Books of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.178.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.cloudfirewall.getAddressBooks({});
 * export const cloudFirewallAddressBookId1 = ids.then(ids => ids.books?.[0]?.id);
 * ```
 */
export function getAddressBooksOutput(args?: GetAddressBooksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddressBooksResult> {
    return pulumi.output(args).apply((a: any) => getAddressBooks(a, opts))
}

/**
 * A collection of arguments for invoking getAddressBooks.
 */
export interface GetAddressBooksOutputArgs {
    /**
     * The type of the Address Book.
     */
    groupType?: pulumi.Input<string>;
    /**
     * A list of Address Book IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results Address Book name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
