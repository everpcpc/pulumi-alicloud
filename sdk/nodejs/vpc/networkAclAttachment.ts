// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a network acl attachment resource to associate network acls to vswitches.
 * 
 * > **NOTE:** Available in 1.44.0+. Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "NatGatewayConfigSpec";
 * 
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultNetworkAcl = new alicloud.vpc.NetworkAcl("default", {
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/21",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultNetworkAclAttachment = new alicloud.vpc.NetworkAclAttachment("default", {
 *     networkAclId: defaultNetworkAcl.id,
 *     resources: [{
 *         resourceId: defaultSwitch.id,
 *         resourceType: "VSwitch",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/network_acl_attachment.html.markdown.
 */
export class NetworkAclAttachment extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAclAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclAttachmentState, opts?: pulumi.CustomResourceOptions): NetworkAclAttachment {
        return new NetworkAclAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/networkAclAttachment:NetworkAclAttachment';

    /**
     * Returns true if the given object is an instance of NetworkAclAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAclAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAclAttachment.__pulumiType;
    }

    /**
     * The id of the network acl, the field can't be changed.
     */
    public readonly networkAclId!: pulumi.Output<string>;
    /**
     * List of the resources associated with the network acl. The details see Block Resources.
     */
    public readonly resources!: pulumi.Output<outputs.vpc.NetworkAclAttachmentResource[]>;

    /**
     * Create a NetworkAclAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAclAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclAttachmentArgs | NetworkAclAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkAclAttachmentState | undefined;
            inputs["networkAclId"] = state ? state.networkAclId : undefined;
            inputs["resources"] = state ? state.resources : undefined;
        } else {
            const args = argsOrState as NetworkAclAttachmentArgs | undefined;
            if (!args || args.networkAclId === undefined) {
                throw new Error("Missing required property 'networkAclId'");
            }
            if (!args || args.resources === undefined) {
                throw new Error("Missing required property 'resources'");
            }
            inputs["networkAclId"] = args ? args.networkAclId : undefined;
            inputs["resources"] = args ? args.resources : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkAclAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAclAttachment resources.
 */
export interface NetworkAclAttachmentState {
    /**
     * The id of the network acl, the field can't be changed.
     */
    readonly networkAclId?: pulumi.Input<string>;
    /**
     * List of the resources associated with the network acl. The details see Block Resources.
     */
    readonly resources?: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclAttachmentResource>[]>;
}

/**
 * The set of arguments for constructing a NetworkAclAttachment resource.
 */
export interface NetworkAclAttachmentArgs {
    /**
     * The id of the network acl, the field can't be changed.
     */
    readonly networkAclId: pulumi.Input<string>;
    /**
     * List of the resources associated with the network acl. The details see Block Resources.
     */
    readonly resources: pulumi.Input<pulumi.Input<inputs.vpc.NetworkAclAttachmentResource>[]>;
}
