// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Elastic IP address association can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ecs/eipAssociation:EipAssociation example eip-abc12345678:i-abc12355
 * ```
 */
export class EipAssociation extends pulumi.CustomResource {
    /**
     * Get an existing EipAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EipAssociationState, opts?: pulumi.CustomResourceOptions): EipAssociation {
        return new EipAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/eipAssociation:EipAssociation';

    /**
     * Returns true if the given object is an instance of EipAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EipAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EipAssociation.__pulumiType;
    }

    /**
     * The allocation EIP ID.
     */
    public readonly allocationId!: pulumi.Output<string>;
    /**
     * When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default to `false`.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The type of cloud product that the eip instance to bind. Valid values: `EcsInstance`, `SlbInstance`, `Nat`, `NetworkInterface` and `HaVip`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The private IP address in the network segment of the vswitch which has been assigned.
     */
    public readonly privateIpAddress!: pulumi.Output<string>;

    /**
     * Create a EipAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EipAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EipAssociationArgs | EipAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EipAssociationState | undefined;
            inputs["allocationId"] = state ? state.allocationId : undefined;
            inputs["force"] = state ? state.force : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
        } else {
            const args = argsOrState as EipAssociationArgs | undefined;
            if ((!args || args.allocationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allocationId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["allocationId"] = args ? args.allocationId : undefined;
            inputs["force"] = args ? args.force : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EipAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EipAssociation resources.
 */
export interface EipAssociationState {
    /**
     * The allocation EIP ID.
     */
    readonly allocationId?: pulumi.Input<string>;
    /**
     * When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default to `false`.
     */
    readonly force?: pulumi.Input<boolean>;
    /**
     * The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The type of cloud product that the eip instance to bind. Valid values: `EcsInstance`, `SlbInstance`, `Nat`, `NetworkInterface` and `HaVip`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The private IP address in the network segment of the vswitch which has been assigned.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EipAssociation resource.
 */
export interface EipAssociationArgs {
    /**
     * The allocation EIP ID.
     */
    readonly allocationId: pulumi.Input<string>;
    /**
     * When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default to `false`.
     */
    readonly force?: pulumi.Input<boolean>;
    /**
     * The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The type of cloud product that the eip instance to bind. Valid values: `EcsInstance`, `SlbInstance`, `Nat`, `NetworkInterface` and `HaVip`.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * The private IP address in the network segment of the vswitch which has been assigned.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
}
