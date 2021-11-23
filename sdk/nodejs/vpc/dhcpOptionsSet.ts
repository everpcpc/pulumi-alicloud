// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a VPC Dhcp Options Set resource.
 *
 * For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 *
 * > **NOTE:** Available in v1.134.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.vpc.DhcpOptionsSet("example", {
 *     dhcpOptionsSetDescription: "example_value",
 *     dhcpOptionsSetName: "example_value",
 *     domainName: "example.com",
 *     domainNameServers: "100.100.2.136",
 * });
 * ```
 *
 * ## Import
 *
 * VPC Dhcp Options Set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet example <id>
 * ```
 */
export class DhcpOptionsSet extends pulumi.CustomResource {
    /**
     * Get an existing DhcpOptionsSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DhcpOptionsSetState, opts?: pulumi.CustomResourceOptions): DhcpOptionsSet {
        return new DhcpOptionsSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet';

    /**
     * Returns true if the given object is an instance of DhcpOptionsSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DhcpOptionsSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DhcpOptionsSet.__pulumiType;
    }

    /**
     * AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10.
     */
    public readonly associateVpcs!: pulumi.Output<outputs.vpc.DhcpOptionsSetAssociateVpc[] | undefined>;
    /**
     * The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    public readonly dhcpOptionsSetDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
     */
    public readonly dhcpOptionsSetName!: pulumi.Output<string | undefined>;
    /**
     * The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
     */
    public readonly domainName!: pulumi.Output<string | undefined>;
    /**
     * The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
     */
    public readonly domainNameServers!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` or `false`.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the account to which the DHCP options set belongs.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a DhcpOptionsSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DhcpOptionsSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DhcpOptionsSetArgs | DhcpOptionsSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DhcpOptionsSetState | undefined;
            inputs["associateVpcs"] = state ? state.associateVpcs : undefined;
            inputs["dhcpOptionsSetDescription"] = state ? state.dhcpOptionsSetDescription : undefined;
            inputs["dhcpOptionsSetName"] = state ? state.dhcpOptionsSetName : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["domainNameServers"] = state ? state.domainNameServers : undefined;
            inputs["dryRun"] = state ? state.dryRun : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DhcpOptionsSetArgs | undefined;
            inputs["associateVpcs"] = args ? args.associateVpcs : undefined;
            inputs["dhcpOptionsSetDescription"] = args ? args.dhcpOptionsSetDescription : undefined;
            inputs["dhcpOptionsSetName"] = args ? args.dhcpOptionsSetName : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainNameServers"] = args ? args.domainNameServers : undefined;
            inputs["dryRun"] = args ? args.dryRun : undefined;
            inputs["ownerId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DhcpOptionsSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DhcpOptionsSet resources.
 */
export interface DhcpOptionsSetState {
    /**
     * AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10.
     */
    associateVpcs?: pulumi.Input<pulumi.Input<inputs.vpc.DhcpOptionsSetAssociateVpc>[]>;
    /**
     * The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    dhcpOptionsSetDescription?: pulumi.Input<string>;
    /**
     * The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
     */
    dhcpOptionsSetName?: pulumi.Input<string>;
    /**
     * The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
     */
    domainNameServers?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` or `false`.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the account to which the DHCP options set belongs.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DhcpOptionsSet resource.
 */
export interface DhcpOptionsSetArgs {
    /**
     * AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10.
     */
    associateVpcs?: pulumi.Input<pulumi.Input<inputs.vpc.DhcpOptionsSetAssociateVpc>[]>;
    /**
     * The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    dhcpOptionsSetDescription?: pulumi.Input<string>;
    /**
     * The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
     */
    dhcpOptionsSetName?: pulumi.Input<string>;
    /**
     * The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
     */
    domainNameServers?: pulumi.Input<string>;
    /**
     * Specifies whether to precheck this request only. Valid values: `true` or `false`.
     */
    dryRun?: pulumi.Input<boolean>;
}