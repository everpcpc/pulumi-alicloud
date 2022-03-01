// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "NetworkAclEntries";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/12"});
 * const defaultNetworkAcl = new alicloud.vpc.NetworkAcl("defaultNetworkAcl", {vpcId: defaultNetwork.id});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 * });
 * const defaultNetworkAclAttachment = new alicloud.vpc.NetworkAclAttachment("defaultNetworkAclAttachment", {
 *     networkAclId: defaultNetworkAcl.id,
 *     resources: [{
 *         resourceId: defaultSwitch.id,
 *         resourceType: "VSwitch",
 *     }],
 * });
 * const defaultNetworkAclEntries = new alicloud.vpc.NetworkAclEntries("defaultNetworkAclEntries", {
 *     networkAclId: defaultNetworkAcl.id,
 *     ingresses: [{
 *         protocol: "all",
 *         port: "-1/-1",
 *         sourceCidrIp: "0.0.0.0/32",
 *         name: name,
 *         entryType: "custom",
 *         policy: "accept",
 *         description: name,
 *     }],
 *     egresses: [{
 *         protocol: "all",
 *         port: "-1/-1",
 *         destinationCidrIp: "0.0.0.0/32",
 *         name: name,
 *         entryType: "custom",
 *         policy: "accept",
 *         description: name,
 *     }],
 * });
 * ```
 */
export class NetworkAclEntries extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAclEntries resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclEntriesState, opts?: pulumi.CustomResourceOptions): NetworkAclEntries {
        return new NetworkAclEntries(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/networkAclEntries:NetworkAclEntries';

    /**
     * Returns true if the given object is an instance of NetworkAclEntries.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAclEntries {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAclEntries.__pulumiType;
    }

    /**
     * List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
     */
    public readonly egresses!: pulumi.Output<outputs.vpc.NetworkAclEntriesEgress[] | undefined>;
    /**
     * List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
     */
    public readonly ingresses!: pulumi.Output<outputs.vpc.NetworkAclEntriesIngress[] | undefined>;
    /**
     * The id of the network acl, the field can't be changed.
     */
    public readonly networkAclId!: pulumi.Output<string>;

    /**
     * Create a NetworkAclEntries resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAclEntriesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclEntriesArgs | NetworkAclEntriesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAclEntriesState | undefined;
            resourceInputs["egresses"] = state ? state.egresses : undefined;
            resourceInputs["ingresses"] = state ? state.ingresses : undefined;
            resourceInputs["networkAclId"] = state ? state.networkAclId : undefined;
        } else {
            const args = argsOrState as NetworkAclEntriesArgs | undefined;
            if ((!args || args.networkAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkAclId'");
            }
            resourceInputs["egresses"] = args ? args.egresses : undefined;
            resourceInputs["ingresses"] = args ? args.ingresses : undefined;
            resourceInputs["networkAclId"] = args ? args.networkAclId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkAclEntries.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAclEntries resources.
 */
export interface NetworkAclEntriesState {
    /**
     * List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
     */
    egresses?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclEntriesEgress>[]>;
    /**
     * List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
     */
    ingresses?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclEntriesIngress>[]>;
    /**
     * The id of the network acl, the field can't be changed.
     */
    networkAclId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAclEntries resource.
 */
export interface NetworkAclEntriesArgs {
    /**
     * List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
     */
    egresses?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclEntriesEgress>[]>;
    /**
     * List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
     */
    ingresses?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclEntriesIngress>[]>;
    /**
     * The id of the network acl, the field can't be changed.
     */
    networkAclId: pulumi.Input<string>;
}
