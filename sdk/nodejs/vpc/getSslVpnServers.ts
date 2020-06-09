// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The SSL-VPN servers data source lists lots of SSL-VPN servers resource information owned by an Alicloud account.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = pulumi.output(alicloud.vpc.getSslVpnServers({
 *     ids: ["fake-server-id"],
 *     nameRegex: "^foo",
 *     outputFile: "/tmp/sslserver",
 *     vpnGatewayId: "fake-vpn-id",
 * }, { async: true }));
 * ```
 */
export function getSslVpnServers(args?: GetSslVpnServersArgs, opts?: pulumi.InvokeOptions): Promise<GetSslVpnServersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getSslVpnServers:getSslVpnServers", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSslVpnServers.
 */
export interface GetSslVpnServersArgs {
    /**
     * IDs of the SSL-VPN servers.
     */
    readonly ids?: string[];
    /**
     * A regex string of SSL-VPN server name.
     */
    readonly nameRegex?: string;
    /**
     * Save the result to the file.
     */
    readonly outputFile?: string;
    /**
     * Use the VPN gateway ID as the search key.
     */
    readonly vpnGatewayId?: string;
}

/**
 * A collection of values returned by getSslVpnServers.
 */
export interface GetSslVpnServersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SSL-VPN server IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of SSL-VPN server names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of SSL-VPN servers. Each element contains the following attributes:
     */
    readonly servers: outputs.vpc.GetSslVpnServersServer[];
    /**
     * The ID of the VPN gateway instance.
     */
    readonly vpnGatewayId?: string;
}
