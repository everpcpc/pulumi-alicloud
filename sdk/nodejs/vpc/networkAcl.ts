// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a network acl resource to add network acls.
 *
 * > **NOTE:** Available in 1.43.0+. Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/12"});
 * const defaultNetworkAcl = new alicloud.vpc.NetworkAcl("defaultNetworkAcl", {
 *     vpcId: defaultNetwork.id,
 *     description: "network_acl",
 * });
 * ```
 *
 * ## Import
 *
 * The network acl can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/networkAcl:NetworkAcl default nacl-abc123456
 * ```
 */
export class NetworkAcl extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclState, opts?: pulumi.CustomResourceOptions): NetworkAcl {
        return new NetworkAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/networkAcl:NetworkAcl';

    /**
     * Returns true if the given object is an instance of NetworkAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAcl.__pulumiType;
    }

    /**
     * The description of the network acl instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the network acl.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The vpcId of the network acl, the field can't be changed.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a NetworkAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclArgs | NetworkAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAclState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as NetworkAclArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAcl resources.
 */
export interface NetworkAclState {
    /**
     * The description of the network acl instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the network acl.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The vpcId of the network acl, the field can't be changed.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAcl resource.
 */
export interface NetworkAclArgs {
    /**
     * The description of the network acl instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the network acl.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The vpcId of the network acl, the field can't be changed.
     */
    readonly vpcId: pulumi.Input<string>;
}
