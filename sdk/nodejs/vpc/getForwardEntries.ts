// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of Forward Entries owned by an Alibaba Cloud account.
 *
 * > **NOTE:** Available in 1.37.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "forward-entry-config-example-name";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/12",
 *     vpcName: name,
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     cidrBlock: "172.16.0.0/21",
 *     vpcId: defaultNetwork.id,
 *     vswitchName: name,
 *     zoneId: defaultZones.zones[0].id,
 * });
 * const defaultNatGateway = new alicloud.vpc.NatGateway("default", {
 *     specification: "Small",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultEipAddress = new alicloud.ecs.EipAddress("default", {
 *     addressName: name,
 * });
 * const defaultEipAssociation = new alicloud.ecs.EipAssociation("default", {
 *     allocationId: defaultEipAddress.id,
 *     instanceId: defaultNatGateway.id,
 * });
 * const defaultForwardEntry = new alicloud.vpc.ForwardEntry("default", {
 *     externalIp: defaultEipAddress.ipAddress,
 *     externalPort: "80",
 *     forwardTableId: defaultNatGateway.forwardTableIds,
 *     internalIp: "172.16.0.3",
 *     internalPort: "8080",
 *     ipProtocol: "tcp",
 * });
 * const defaultForwardEntries = defaultForwardEntry.forwardTableId.apply(forwardTableId => alicloud.vpc.getForwardEntries({
 *     forwardTableId: forwardTableId,
 * }));
 * ```
 */
export function getForwardEntries(args: GetForwardEntriesArgs, opts?: pulumi.InvokeOptions): Promise<GetForwardEntriesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getForwardEntries:getForwardEntries", {
        "externalIp": args.externalIp,
        "externalPort": args.externalPort,
        "forwardEntryName": args.forwardEntryName,
        "forwardTableId": args.forwardTableId,
        "ids": args.ids,
        "internalIp": args.internalIp,
        "internalPort": args.internalPort,
        "ipProtocol": args.ipProtocol,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getForwardEntries.
 */
export interface GetForwardEntriesArgs {
    /**
     * The public IP address.
     */
    externalIp?: string;
    /**
     * The public port.
     */
    externalPort?: string;
    /**
     * The name of forward entry.
     */
    forwardEntryName?: string;
    /**
     * The ID of the Forward table.
     */
    forwardTableId: string;
    /**
     * A list of Forward Entries IDs.
     */
    ids?: string[];
    /**
     * The private IP address.
     */
    internalIp?: string;
    /**
     * The internal port.
     */
    internalPort?: string;
    /**
     * The ip protocol. Valid values: `any`,`tcp` and `udp`.
     */
    ipProtocol?: string;
    /**
     * A regex string to filter results by forward entry name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
     */
    status?: string;
}

/**
 * A collection of values returned by getForwardEntries.
 */
export interface GetForwardEntriesResult {
    /**
     * A list of Forward Entries. Each element contains the following attributes:
     */
    readonly entries: outputs.vpc.GetForwardEntriesEntry[];
    /**
     * The public IP address.
     */
    readonly externalIp?: string;
    /**
     * The public port.
     */
    readonly externalPort?: string;
    /**
     * The name of forward entry.
     */
    readonly forwardEntryName?: string;
    readonly forwardTableId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Forward Entries IDs.
     */
    readonly ids: string[];
    /**
     * The private IP address.
     */
    readonly internalIp?: string;
    /**
     * The private port.
     */
    readonly internalPort?: string;
    /**
     * The protocol type.
     */
    readonly ipProtocol?: string;
    readonly nameRegex?: string;
    /**
     * A list of Forward Entries names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of forward entry.
     */
    readonly status?: string;
}

export function getForwardEntriesOutput(args: GetForwardEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetForwardEntriesResult> {
    return pulumi.output(args).apply(a => getForwardEntries(a, opts))
}

/**
 * A collection of arguments for invoking getForwardEntries.
 */
export interface GetForwardEntriesOutputArgs {
    /**
     * The public IP address.
     */
    externalIp?: pulumi.Input<string>;
    /**
     * The public port.
     */
    externalPort?: pulumi.Input<string>;
    /**
     * The name of forward entry.
     */
    forwardEntryName?: pulumi.Input<string>;
    /**
     * The ID of the Forward table.
     */
    forwardTableId: pulumi.Input<string>;
    /**
     * A list of Forward Entries IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The private IP address.
     */
    internalIp?: pulumi.Input<string>;
    /**
     * The internal port.
     */
    internalPort?: pulumi.Input<string>;
    /**
     * The ip protocol. Valid values: `any`,`tcp` and `udp`.
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * A regex string to filter results by forward entry name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
     */
    status?: pulumi.Input<string>;
}
