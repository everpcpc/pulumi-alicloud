// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a snat resource.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/snat_entry.html.markdown.
 */
export class SnatEntry extends pulumi.CustomResource {
    /**
     * Get an existing SnatEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatEntryState, opts?: pulumi.CustomResourceOptions): SnatEntry {
        return new SnatEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/snatEntry:SnatEntry';

    /**
     * Returns true if the given object is an instance of SnatEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnatEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnatEntry.__pulumiType;
    }

    /**
     * The id of the snat entry on the server.
     */
    public /*out*/ readonly snatEntryId!: pulumi.Output<string>;
    /**
     * The name of snat entry.
     */
    public readonly snatEntryName!: pulumi.Output<string | undefined>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    public readonly snatIp!: pulumi.Output<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    public readonly snatTableId!: pulumi.Output<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    public readonly sourceCidr!: pulumi.Output<string | undefined>;
    /**
     * The vswitch ID.
     */
    public readonly sourceVswitchId!: pulumi.Output<string | undefined>;

    /**
     * Create a SnatEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnatEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnatEntryArgs | SnatEntryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SnatEntryState | undefined;
            inputs["snatEntryId"] = state ? state.snatEntryId : undefined;
            inputs["snatEntryName"] = state ? state.snatEntryName : undefined;
            inputs["snatIp"] = state ? state.snatIp : undefined;
            inputs["snatTableId"] = state ? state.snatTableId : undefined;
            inputs["sourceCidr"] = state ? state.sourceCidr : undefined;
            inputs["sourceVswitchId"] = state ? state.sourceVswitchId : undefined;
        } else {
            const args = argsOrState as SnatEntryArgs | undefined;
            if (!args || args.snatIp === undefined) {
                throw new Error("Missing required property 'snatIp'");
            }
            if (!args || args.snatTableId === undefined) {
                throw new Error("Missing required property 'snatTableId'");
            }
            inputs["snatEntryName"] = args ? args.snatEntryName : undefined;
            inputs["snatIp"] = args ? args.snatIp : undefined;
            inputs["snatTableId"] = args ? args.snatTableId : undefined;
            inputs["sourceCidr"] = args ? args.sourceCidr : undefined;
            inputs["sourceVswitchId"] = args ? args.sourceVswitchId : undefined;
            inputs["snatEntryId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SnatEntry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnatEntry resources.
 */
export interface SnatEntryState {
    /**
     * The id of the snat entry on the server.
     */
    readonly snatEntryId?: pulumi.Input<string>;
    /**
     * The name of snat entry.
     */
    readonly snatEntryName?: pulumi.Input<string>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    readonly snatIp?: pulumi.Input<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    readonly snatTableId?: pulumi.Input<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    readonly sourceCidr?: pulumi.Input<string>;
    /**
     * The vswitch ID.
     */
    readonly sourceVswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnatEntry resource.
 */
export interface SnatEntryArgs {
    /**
     * The name of snat entry.
     */
    readonly snatEntryName?: pulumi.Input<string>;
    /**
     * The SNAT ip address, the ip must along bandwidth package public ip which `alicloud.vpc.NatGateway` argument `bandwidthPackages`.
     */
    readonly snatIp: pulumi.Input<string>;
    /**
     * The value can get from `alicloud.vpc.NatGateway` Attributes "snatTableIds".
     */
    readonly snatTableId: pulumi.Input<string>;
    /**
     * The private network segment of Ecs. This parameter and the `sourceVswitchId` parameter are mutually exclusive and cannot appear at the same time.
     */
    readonly sourceCidr?: pulumi.Input<string>;
    /**
     * The vswitch ID.
     */
    readonly sourceVswitchId?: pulumi.Input<string>;
}
