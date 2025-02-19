// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Waf Certificates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.waf.getCertificates({
 *     ids: ["your_certificate_id"],
 *     instanceId: "your_instance_id",
 *     domain: "your_domain_name",
 * });
 * export const wafCertificate = _default.then(_default => _default.certificates?.[0]);
 * ```
 */
export function getCertificates(args: GetCertificatesArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificatesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:waf/getCertificates:getCertificates", {
        "domain": args.domain,
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificates.
 */
export interface GetCertificatesArgs {
    /**
     * The domain that you want to add to WAF.
     */
    domain?: string;
    /**
     * A list of Certificate IDs.
     */
    ids?: string[];
    /**
     * WAF instance ID.
     */
    instanceId: string;
    /**
     * A regex string to filter results by Certificate name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getCertificates.
 */
export interface GetCertificatesResult {
    readonly certificates: outputs.waf.GetCertificatesCertificate[];
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Waf Certificates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.waf.getCertificates({
 *     ids: ["your_certificate_id"],
 *     instanceId: "your_instance_id",
 *     domain: "your_domain_name",
 * });
 * export const wafCertificate = _default.then(_default => _default.certificates?.[0]);
 * ```
 */
export function getCertificatesOutput(args: GetCertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificatesResult> {
    return pulumi.output(args).apply((a: any) => getCertificates(a, opts))
}

/**
 * A collection of arguments for invoking getCertificates.
 */
export interface GetCertificatesOutputArgs {
    /**
     * The domain that you want to add to WAF.
     */
    domain?: pulumi.Input<string>;
    /**
     * A list of Certificate IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * WAF instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A regex string to filter results by Certificate name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
