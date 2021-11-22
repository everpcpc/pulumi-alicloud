// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Transit Router Route Table Associations available to the user.[What is Cen Transit Router Route Table Associations](https://help.aliyun.com/document_detail/261243.html)
 *
 * > **NOTE:** Available in 1.126.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cen.getTransitRouterRouteTableAssociations({
 *     transitRouterRouteTableId: "rtb-id1",
 * });
 * export const firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType = _default.then(_default => _default.associations?[0]?.resourceType);
 * ```
 */
export function getTransitRouterRouteTableAssociations(args: GetTransitRouterRouteTableAssociationsArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRouterRouteTableAssociationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouterRouteTableAssociations:getTransitRouterRouteTableAssociations", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "transitRouterRouteTableId": args.transitRouterRouteTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouterRouteTableAssociations.
 */
export interface GetTransitRouterRouteTableAssociationsArgs {
    /**
     * A list of CEN Transit Router Route Table Association IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of the route table, including `Active`, `Associating`, `Dissociating`.
     */
    status?: string;
    /**
     * ID of the route table of the VPC or VBR.
     */
    transitRouterRouteTableId: string;
}

/**
 * A collection of values returned by getTransitRouterRouteTableAssociations.
 */
export interface GetTransitRouterRouteTableAssociationsResult {
    /**
     * A list of CEN Transit Router Route Table Associations. Each element contains the following attributes:
     */
    readonly associations: outputs.cen.GetTransitRouterRouteTableAssociationsAssociation[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of CEN Transit Router Route Table Association IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * The status of the route table.
     */
    readonly status?: string;
    /**
     * ID of the transit router route table.
     */
    readonly transitRouterRouteTableId: string;
}

export function getTransitRouterRouteTableAssociationsOutput(args: GetTransitRouterRouteTableAssociationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRouterRouteTableAssociationsResult> {
    return pulumi.output(args).apply(a => getTransitRouterRouteTableAssociations(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouterRouteTableAssociations.
 */
export interface GetTransitRouterRouteTableAssociationsOutputArgs {
    /**
     * A list of CEN Transit Router Route Table Association IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the route table, including `Active`, `Associating`, `Dissociating`.
     */
    status?: pulumi.Input<string>;
    /**
     * ID of the route table of the VPC or VBR.
     */
    transitRouterRouteTableId: pulumi.Input<string>;
}
