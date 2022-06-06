// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a VPN Ipsec Server resource.
 *
 * For information about VPN Ipsec Server and how to use it, see [What is Ipsec Server](https://www.alibabacloud.com/help/en/doc-detail/205454.html).
 *
 * > **NOTE:** Available in v1.161.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "default-NODELETING",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?[0],
 *     zoneId: defaultZones.zones?[0]?.id,
 * }));
 * const vswitchId = defaultSwitches.then(defaultSwitches => defaultSwitches.ids?[0]);
 * const defaultGateway = new alicloud.vpn.Gateway("defaultGateway", {
 *     vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?[0]),
 *     bandwidth: 10,
 *     enableSsl: true,
 *     enableIpsec: true,
 *     sslConnections: 5,
 *     instanceChargeType: "PrePaid",
 *     vswitchId: vswitchId,
 * });
 * const example = new alicloud.vpn.IpsecServer("example", {
 *     clientIpPool: "example_value",
 *     ipsecServerName: "example_value",
 *     localSubnet: "example_value",
 *     vpnGatewayId: defaultGateway.id,
 * });
 * ```
 *
 * ## Import
 *
 * VPN Ipsec Server can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpn/ipsecServer:IpsecServer example <id>
 * ```
 */
export class IpsecServer extends pulumi.CustomResource {
    /**
     * Get an existing IpsecServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsecServerState, opts?: pulumi.CustomResourceOptions): IpsecServer {
        return new IpsecServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/ipsecServer:IpsecServer';

    /**
     * Returns true if the given object is an instance of IpsecServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsecServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsecServer.__pulumiType;
    }

    /**
     * The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
     */
    public readonly clientIpPool!: pulumi.Output<string>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether you want the configuration to immediately take effect.
     */
    public readonly effectImmediately!: pulumi.Output<boolean | undefined>;
    /**
     * The configuration of Phase 1 negotiations. See the following `Block ikeConfig`.
     */
    public readonly ikeConfigs!: pulumi.Output<outputs.vpn.IpsecServerIkeConfig[]>;
    /**
     * The configuration of Phase 2 negotiations. See the following `Block ipsecConfig`.
     */
    public readonly ipsecConfigs!: pulumi.Output<outputs.vpn.IpsecServerIpsecConfig[]>;
    /**
     * The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
     */
    public readonly ipsecServerName!: pulumi.Output<string | undefined>;
    /**
     * The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
     */
    public readonly localSubnet!: pulumi.Output<string>;
    /**
     * The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
     */
    public readonly psk!: pulumi.Output<string>;
    /**
     * Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
     */
    public readonly pskEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the VPN gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string>;

    /**
     * Create a IpsecServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpsecServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsecServerArgs | IpsecServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsecServerState | undefined;
            resourceInputs["clientIpPool"] = state ? state.clientIpPool : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["effectImmediately"] = state ? state.effectImmediately : undefined;
            resourceInputs["ikeConfigs"] = state ? state.ikeConfigs : undefined;
            resourceInputs["ipsecConfigs"] = state ? state.ipsecConfigs : undefined;
            resourceInputs["ipsecServerName"] = state ? state.ipsecServerName : undefined;
            resourceInputs["localSubnet"] = state ? state.localSubnet : undefined;
            resourceInputs["psk"] = state ? state.psk : undefined;
            resourceInputs["pskEnabled"] = state ? state.pskEnabled : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as IpsecServerArgs | undefined;
            if ((!args || args.clientIpPool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientIpPool'");
            }
            if ((!args || args.localSubnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localSubnet'");
            }
            if ((!args || args.vpnGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnGatewayId'");
            }
            resourceInputs["clientIpPool"] = args ? args.clientIpPool : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["effectImmediately"] = args ? args.effectImmediately : undefined;
            resourceInputs["ikeConfigs"] = args ? args.ikeConfigs : undefined;
            resourceInputs["ipsecConfigs"] = args ? args.ipsecConfigs : undefined;
            resourceInputs["ipsecServerName"] = args ? args.ipsecServerName : undefined;
            resourceInputs["localSubnet"] = args ? args.localSubnet : undefined;
            resourceInputs["psk"] = args ? args.psk : undefined;
            resourceInputs["pskEnabled"] = args ? args.pskEnabled : undefined;
            resourceInputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpsecServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsecServer resources.
 */
export interface IpsecServerState {
    /**
     * The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
     */
    clientIpPool?: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether you want the configuration to immediately take effect.
     */
    effectImmediately?: pulumi.Input<boolean>;
    /**
     * The configuration of Phase 1 negotiations. See the following `Block ikeConfig`.
     */
    ikeConfigs?: pulumi.Input<pulumi.Input<inputs.vpn.IpsecServerIkeConfig>[]>;
    /**
     * The configuration of Phase 2 negotiations. See the following `Block ipsecConfig`.
     */
    ipsecConfigs?: pulumi.Input<pulumi.Input<inputs.vpn.IpsecServerIpsecConfig>[]>;
    /**
     * The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
     */
    ipsecServerName?: pulumi.Input<string>;
    /**
     * The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
     */
    localSubnet?: pulumi.Input<string>;
    /**
     * The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
     */
    psk?: pulumi.Input<string>;
    /**
     * Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
     */
    pskEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsecServer resource.
 */
export interface IpsecServerArgs {
    /**
     * The client CIDR block. It refers to the CIDR block that is allocated to the virtual interface of the client.
     */
    clientIpPool: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Specifies whether you want the configuration to immediately take effect.
     */
    effectImmediately?: pulumi.Input<boolean>;
    /**
     * The configuration of Phase 1 negotiations. See the following `Block ikeConfig`.
     */
    ikeConfigs?: pulumi.Input<pulumi.Input<inputs.vpn.IpsecServerIkeConfig>[]>;
    /**
     * The configuration of Phase 2 negotiations. See the following `Block ipsecConfig`.
     */
    ipsecConfigs?: pulumi.Input<pulumi.Input<inputs.vpn.IpsecServerIpsecConfig>[]>;
    /**
     * The name of the IPsec server. The name must be `2` to `128` characters in length, and can contain digits, hyphens (-), and underscores (_). It must start with a letter.
     */
    ipsecServerName?: pulumi.Input<string>;
    /**
     * The local CIDR block. It refers to the CIDR block of the virtual private cloud (VPC) that is used to connect with the client. Separate multiple CIDR blocks with commas (,). Example: `192.168.1.0/24,192.168.2.0/24`.
     */
    localSubnet: pulumi.Input<string>;
    /**
     * The pre-shared key. The pre-shared key is used to authenticate the VPN gateway and the client. By default, the system generates a random string that is 16 bits in length. You can also specify the pre-shared key. It can contain at most 100 characters.
     */
    psk?: pulumi.Input<string>;
    /**
     * Whether to enable the pre-shared key authentication method. The value is only `true`, which indicates that the pre-shared key authentication method is enabled.
     */
    pskEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId: pulumi.Input<string>;
}