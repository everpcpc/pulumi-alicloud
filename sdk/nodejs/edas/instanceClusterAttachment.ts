// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an EDAS instance cluster attachment resource.
 *
 * > **NOTE:** Available in 1.82.0+
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.edas.InstanceClusterAttachment("default", {
 *     clusterId: _var.cluster_id,
 *     instanceIds: _var.instance_ids,
 * });
 * ```
 */
export class InstanceClusterAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InstanceClusterAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceClusterAttachmentState, opts?: pulumi.CustomResourceOptions): InstanceClusterAttachment {
        return new InstanceClusterAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:edas/instanceClusterAttachment:InstanceClusterAttachment';

    /**
     * Returns true if the given object is an instance of InstanceClusterAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceClusterAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceClusterAttachment.__pulumiType;
    }

    /**
     * The ID of the cluster that you want to create the application.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The cluster members map of the resource supplied above. The key is instanceId and the value is cluster_member_id.
     */
    public /*out*/ readonly clusterMemberIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ecu map of the resource supplied above. The key is instanceId and the value is ecu_id.
     */
    public /*out*/ readonly ecuMap!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of instance. Type: list.
     */
    public readonly instanceIds!: pulumi.Output<string[]>;
    /**
     * The status map of the resource supplied above. The key is instanceId and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     */
    public /*out*/ readonly statusMap!: pulumi.Output<{[key: string]: number}>;

    /**
     * Create a InstanceClusterAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceClusterAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceClusterAttachmentArgs | InstanceClusterAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceClusterAttachmentState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["clusterMemberIds"] = state ? state.clusterMemberIds : undefined;
            inputs["ecuMap"] = state ? state.ecuMap : undefined;
            inputs["instanceIds"] = state ? state.instanceIds : undefined;
            inputs["statusMap"] = state ? state.statusMap : undefined;
        } else {
            const args = argsOrState as InstanceClusterAttachmentArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.instanceIds === undefined) {
                throw new Error("Missing required property 'instanceIds'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["instanceIds"] = args ? args.instanceIds : undefined;
            inputs["clusterMemberIds"] = undefined /*out*/;
            inputs["ecuMap"] = undefined /*out*/;
            inputs["statusMap"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceClusterAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceClusterAttachment resources.
 */
export interface InstanceClusterAttachmentState {
    /**
     * The ID of the cluster that you want to create the application.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The cluster members map of the resource supplied above. The key is instanceId and the value is cluster_member_id.
     */
    readonly clusterMemberIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ecu map of the resource supplied above. The key is instanceId and the value is ecu_id.
     */
    readonly ecuMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of instance. Type: list.
     */
    readonly instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status map of the resource supplied above. The key is instanceId and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     */
    readonly statusMap?: pulumi.Input<{[key: string]: pulumi.Input<number>}>;
}

/**
 * The set of arguments for constructing a InstanceClusterAttachment resource.
 */
export interface InstanceClusterAttachmentArgs {
    /**
     * The ID of the cluster that you want to create the application.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * The ID of instance. Type: list.
     */
    readonly instanceIds: pulumi.Input<pulumi.Input<string>[]>;
}