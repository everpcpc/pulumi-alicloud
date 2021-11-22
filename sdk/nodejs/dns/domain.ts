// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * DNS can be imported using the id or domain name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dns/domain:Domain example "aliyun.com"
 * ```
 *
 * @deprecated This resource has been deprecated in favour of DnsDomain
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        pulumi.log.warn("Domain is deprecated: This resource has been deprecated in favour of DnsDomain")
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * A list of the dns server name.
     */
    public /*out*/ readonly dnsServers!: pulumi.Output<string[]>;
    /**
     * The domain ID.
     */
    public /*out*/ readonly domainId!: pulumi.Output<string>;
    /**
     * Id of the group in which the domain will add. If not supplied, then use default group.
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Id of resource group which the dns belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated This resource has been deprecated in favour of DnsDomain */
    constructor(name: string, args?: DomainArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated This resource has been deprecated in favour of DnsDomain */
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Domain is deprecated: This resource has been deprecated in favour of DnsDomain")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            inputs["dnsServers"] = state ? state.dnsServers : undefined;
            inputs["domainId"] = state ? state.domainId : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["dnsServers"] = undefined /*out*/;
            inputs["domainId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * A list of the dns server name.
     */
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The domain ID.
     */
    domainId?: pulumi.Input<string>;
    /**
     * Id of the group in which the domain will add. If not supplied, then use default group.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the dns belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Id of the group in which the domain will add. If not supplied, then use default group.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    name?: pulumi.Input<string>;
    /**
     * The Id of resource group which the dns belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
}
