// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the listeners related to a server load balancer of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.slb.LoadBalancer("default", {});
 * const tcp = new alicloud.slb.Listener("tcp", {
 *     loadBalancerId: _default.id,
 *     backendPort: "22",
 *     frontendPort: "22",
 *     protocol: "tcp",
 *     bandwidth: "10",
 *     healthCheckType: "tcp",
 *     persistenceTimeout: 3600,
 *     healthyThreshold: 8,
 *     unhealthyThreshold: 8,
 *     healthCheckTimeout: 8,
 *     healthCheckInterval: 5,
 *     healthCheckHttpCode: "http_2xx",
 *     healthCheckConnectPort: 20,
 *     healthCheckUri: "/console",
 *     establishedTimeout: 600,
 * });
 * const sampleDs = _default.id.apply(id => alicloud.slb.getListeners({
 *     loadBalancerId: id,
 * }));
 * export const firstSlbListenerProtocol = sampleDs.slbListeners[0].protocol;
 * ```
 */
export function getListeners(args: GetListenersArgs, opts?: pulumi.InvokeOptions): Promise<GetListenersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getListeners:getListeners", {
        "descriptionRegex": args.descriptionRegex,
        "frontendPort": args.frontendPort,
        "loadBalancerId": args.loadBalancerId,
        "outputFile": args.outputFile,
        "protocol": args.protocol,
    }, opts);
}

/**
 * A collection of arguments for invoking getListeners.
 */
export interface GetListenersArgs {
    /**
     * A regex string to filter results by SLB listener description.
     */
    readonly descriptionRegex?: string;
    /**
     * Filter listeners by the specified frontend port.
     */
    readonly frontendPort?: number;
    /**
     * ID of the SLB with listeners.
     */
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    /**
     * Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
     */
    readonly protocol?: string;
}

/**
 * A collection of values returned by getListeners.
 */
export interface GetListenersResult {
    readonly descriptionRegex?: string;
    /**
     * Frontend port used to receive incoming traffic and distribute it to the backend servers.
     */
    readonly frontendPort?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loadBalancerId: string;
    readonly outputFile?: string;
    /**
     * Listener protocol. Possible values: `http`, `https`, `tcp` and `udp`.
     */
    readonly protocol?: string;
    /**
     * A list of SLB listeners. Each element contains the following attributes:
     */
    readonly slbListeners: outputs.slb.GetListenersSlbListener[];
}
